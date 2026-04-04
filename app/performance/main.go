package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"runtime/debug"
	"sync"
)

const blockSize = 256 * 1024

var (
	mu      sync.Mutex
	buffers [][]byte

	// 复用 MemProfile 缓冲区，避免每次 /api/stats 都 make 一大块 slice（否则 inuse 会被轮询自己顶上去）
	memProfMu  sync.Mutex
	memProfBuf []runtime.MemProfileRecord
)

// 与 pprof 同源：对 runtime.MemProfile 各条 Stack 合计，等价于把 heap/allocs profile 里的 sample 加总。
type statsResponse struct {
	Blocks int `json:"blocks"`

	HeapInuseObjects int64  `json:"heap_inuse_objects"` // Σ InUseObjects → 对应 heap 的 inuse 对象
	HeapInuseBytes   int64  `json:"heap_inuse_bytes"`   // Σ InUseBytes
	AllocObjects     uint64 `json:"alloc_objects"`      // Σ AllocObjects → 对应 allocs
	AllocBytes       uint64 `json:"alloc_bytes"`        // Σ AllocBytes
	FreeObjects      uint64 `json:"free_objects"`
	FreeBytes        uint64 `json:"free_bytes"`
}

func sumMemProfile() (inUseObj, inUseBytes int64, allocObj, allocBytes, freeObj, freeBytes uint64) {
	memProfMu.Lock()
	defer memProfMu.Unlock()

	n := len(memProfBuf)
	if n < 32 {
		n = 32
	}
	for {
		if cap(memProfBuf) < n {
			memProfBuf = make([]runtime.MemProfileRecord, n)
		} else {
			memProfBuf = memProfBuf[:n]
		}
		got, ok := runtime.MemProfile(memProfBuf, true)
		if ok {
			for _, r := range memProfBuf[:got] {
				inUseObj += r.InUseObjects()
				inUseBytes += r.InUseBytes()
				allocObj += uint64(r.AllocObjects)
				allocBytes += uint64(r.AllocBytes)
				freeObj += uint64(r.FreeObjects)
				freeBytes += uint64(r.FreeBytes)
			}
			return
		}
		if got > n {
			n = got + got/4
		} else {
			n *= 2
		}
	}
}

func collectStats() statsResponse {
	mu.Lock()
	bn := len(buffers)
	mu.Unlock()

	// 演示用：先 GC 再读 profile，减弱「轮询 / JSON」造成的 inuse 背景上涨；生产勿照搬
	runtime.GC()

	hiObj, hiBy, aObj, aBy, fObj, fBy := sumMemProfile()
	return statsResponse{
		Blocks:           bn,
		HeapInuseObjects: hiObj,
		HeapInuseBytes:   hiBy,
		AllocObjects:     aObj,
		AllocBytes:       aBy,
		FreeObjects:      fObj,
		FreeBytes:        fBy,
	}
}

func handleStats(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(collectStats())
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, blockSize)
	for i := range b {
		b[i] = byte(i & 0xff)
	}
	mu.Lock()
	buffers = append(buffers, b)
	mu.Unlock()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleRemove(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	mu.Lock()
	if n := len(buffers); n > 0 {
		buffers[n-1] = nil
		buffers = buffers[:n-1]
	}
	mu.Unlock()
	runtime.GC()
	debug.FreeOSMemory()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = pageTpl.Execute(w, map[string]any{"BlockSize": blockSize, "MemProfileRate": runtime.MemProfileRate})
}

var pageTpl = template.Must(template.New("").Parse(`<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>memprofile / pprof</title>
<style>
body { font-family: system-ui, sans-serif; max-width: 32rem; margin: 2rem auto; padding: 0 1rem; }
.row { margin: 0.75rem 0; }
label { color: #555; font-size: 0.85rem; display: block; }
.val { font-size: 1.2rem; font-variant-numeric: tabular-nums; }
.delta { font-size: 0.8rem; color: #2a6; margin-left: 0.35rem; }
.delta.neg { color: #c42; }
.sec { margin-top: 1.1rem; padding-top: 0.65rem; border-top: 1px solid #ddd; font-size: 0.8rem; color: #444; }
a, button { font: inherit; margin-right: 0.5rem; }
button { padding: 0.35rem 0.75rem; }
</style>
</head>
<body>
<p><a href="/add">add</a>
<form method="post" action="/remove" style="display:inline"><button type="submit">remove</button></form></p>
<div class="row"><label>blocks（业务持有块数，非 pprof）</label><div class="val" id="blocks">—</div></div>

<div class="sec">与 <code>/debug/pprof/heap</code> 同源：<code>runtime.MemProfile</code> 各栈 InUse 合计</div>
<div class="row"><label>inuse 对象数</label><div class="val"><span id="hiObj">—</span><span class="delta" id="dHiObj"></span></div></div>
<div class="row"><label>inuse 字节</label><div class="val"><span id="hiBy">—</span><span class="delta" id="dHiBy"></span></div></div>

<div class="sec">与 <code>/debug/pprof/allocs</code> 同源：同一批 <code>MemProfileRecord</code> 的 Alloc / Free 合计</div>
<div class="row"><label>alloc 对象（累计）</label><div class="val"><span id="aObj">—</span><span class="delta" id="dAObj"></span></div></div>
<div class="row"><label>alloc 字节（累计）</label><div class="val"><span id="aBy">—</span><span class="delta" id="dABy"></span></div></div>
<div class="row"><label>free 对象（累计）</label><div class="val"><span id="fObj">—</span><span class="delta" id="dFObj"></span></div></div>
<div class="row"><label>free 字节（累计）</label><div class="val"><span id="fBy">—</span><span class="delta" id="dFBy"></span></div></div>

<p style="font-size:0.75rem;color:#666"><strong>读数：</strong><code>alloc*</code> / <code>free*</code> 是<strong>累计</strong>，只会增加；remove 后应看到 <strong>free 变大、inuse 变小</strong>，而不是 alloc 下降。之前 inuse 跟着涨，多半是每次轮询 <code>MemProfile</code> 临时分配；已复用缓冲区并在采样前 GC 降噪。</p>
<p style="font-size:0.75rem;color:#666">MemProfileRate={{.MemProfileRate}}。单次 add {{.BlockSize}} 字节。Δ=与上次轮询差。<a href="/debug/pprof/heap">heap</a> · <a href="/debug/pprof/allocs">allocs</a></p>
<script>
function fmt(n) {
  if (n < 1024) return n + ' B';
  if (n < 1048576) return (n/1024).toFixed(1) + ' KiB';
  return (n/1048576).toFixed(2) + ' MiB';
}
function fmtInt(n) { return Number(n).toLocaleString(); }
function setDelta(el, cur, prev, isBytes) {
  if (prev == null) { el.textContent = ''; el.className = 'delta'; return; }
  const d = Number(cur) - Number(prev);
  if (d === 0) { el.textContent = ' Δ 0'; el.className = 'delta'; return; }
  const sign = d > 0 ? '+' : '';
  const txt = isBytes ? fmt(Math.abs(d)) : fmtInt(Math.abs(d));
  el.textContent = ' Δ ' + sign + (d > 0 ? '' : '-') + txt;
  el.className = 'delta' + (d < 0 ? ' neg' : '');
}
let prev = null;
async function tick() {
  const r = await fetch('/api/stats');
  if (!r.ok) return;
  const s = await r.json();
  document.getElementById('blocks').textContent = s.blocks;
  document.getElementById('hiObj').textContent = fmtInt(s.heap_inuse_objects);
  document.getElementById('hiBy').textContent = fmt(s.heap_inuse_bytes);
  document.getElementById('aObj').textContent = fmtInt(s.alloc_objects);
  document.getElementById('aBy').textContent = fmt(s.alloc_bytes);
  document.getElementById('fObj').textContent = fmtInt(s.free_objects);
  document.getElementById('fBy').textContent = fmt(s.free_bytes);
  if (prev) {
    setDelta(document.getElementById('dHiObj'), s.heap_inuse_objects, prev.heap_inuse_objects, false);
    setDelta(document.getElementById('dHiBy'), s.heap_inuse_bytes, prev.heap_inuse_bytes, true);
    setDelta(document.getElementById('dAObj'), s.alloc_objects, prev.alloc_objects, false);
    setDelta(document.getElementById('dABy'), s.alloc_bytes, prev.alloc_bytes, true);
    setDelta(document.getElementById('dFObj'), s.free_objects, prev.free_objects, false);
    setDelta(document.getElementById('dFBy'), s.free_bytes, prev.free_bytes, true);
  }
  prev = s;
}
tick();
setInterval(tick, 800);
</script>
</body>
</html>
`))

func main() {
	// 与 pprof 使用同一采样率；设为 1 时几乎所有分配进 profile（演示开销大）。与 blockSize 对齐时通常每次 add 会留下一条栈记录。
	runtime.MemProfileRate = blockSize

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)
	mux.HandleFunc("/add", handleAdd)
	mux.HandleFunc("/remove", handleRemove)
	mux.HandleFunc("/api/stats", handleStats)
	mux.Handle("/debug/pprof/", http.DefaultServeMux)

	addr := ":8080"
	log.Printf("http://localhost%s  MemProfileRate=%d\n", addr, runtime.MemProfileRate)
	log.Fatal(http.ListenAndServe(addr, mux))
}
