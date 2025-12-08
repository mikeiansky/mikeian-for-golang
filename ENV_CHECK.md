# 开发环境检查与代码质量控制指南

本文档介绍如何通过 `Makefile`、`golangci-lint` 和 `pre-commit` 建立完整的开发环境检查和代码质量控制流程。

---

## 目录

- [Go 版本检测](#go-版本检测)
- [golangci-lint 版本检测](#golangci-lint-版本检测)
- [pre-commit 安装检测](#pre-commit-安装检测)
- [环境检查命令](#环境检查命令)
- [golangci-lint 配置](#golangci-lint-配置)
- [pre-commit 配置](#pre-commit-配置)
- [常见问题](#常见问题)

---

## Go 版本检测

### 概述

统一 Go 版本可以避免版本差异导致的兼容性问题和不可预期的行为。通过 Makefile 配置，开发人员在开发前可以快速验证本地 Go 版本是否满足项目要求。

### Makefile 配置

```makefile
# 目标 Go 版本
GO_VERSION := 1.25.1

# 获取当前系统 Go 版本（例如：go version go1.25.1 linux/amd64 → 提取 1.25.1）
CURRENT_GO_VERSION := $(shell go version | awk '{print $$3}' | sed 's/go//')

check-go-version:
    @if [ "$(CURRENT_GO_VERSION)" != "$(GO_VERSION)" ]; then \
        echo "❌ Error: Go version mismatch!"; \
        echo "   Expected: $(GO_VERSION)"; \
        echo "   Found:    $(CURRENT_GO_VERSION)"; \
        echo "   Please install or switch to Go $(GO_VERSION)."; \
        exit 1; \
    fi
    @echo "✅ Go version check passed: $(CURRENT_GO_VERSION)"
```

### 使用方法

```bash
make check-go-version
```

**预期输出（成功）：**
```
✅ Go version check passed: 1.25.1
```

**预期输出（失败）：**
```
❌ Error: Go version mismatch!
   Expected: 1.25.1
   Found:    1.24.0
   Please install or switch to Go 1.25.1.
```

---

## golangci-lint 版本检测

### 概述

`golangci-lint` 是一个强大的 Go 代码静态分析工具集，整合了多个 linter。通过版本检测确保团队使用统一的代码检查规则，保证代码质量的一致性。

### Makefile 配置

```makefile
# 目标 golangci-lint 最低版本
GOLANGCI_LINT_REQUIRED_VERSION := 2.7.1

# 获取 golangci-lint 版本号，例如从输出中提取 "2.7.1"
GOLANGCI_LINT_VERSION := $(shell golangci-lint --version 2>/dev/null | grep -oE '[0-9]+\.[0-9]+\.[0-9]+' | head -1)

check-golangci-lint-version:
    @if [ -z "$(GOLANGCI_LINT_VERSION)" ]; then \
        echo "❌ Error: golangci-lint 未安装或无法获取版本！"; \
        exit 1; \
    fi
    @echo "🔍 当前 golangci-lint 版本: $(GOLANGCI_LINT_VERSION)"
    @echo "✅ 要求最低版本: $(GOLANGCI_LINT_REQUIRED_VERSION)"

    @if [ "$(shell printf '%s\n%s' "$(GOLANGCI_LINT_REQUIRED_VERSION)" "$(GOLANGCI_LINT_VERSION)" | sort -V | head -n1)" != "$(GOLANGCI_LINT_REQUIRED_VERSION)" ]; then \
        echo "❌ 错误：golangci-lint 版本过低！"; \
        echo "   当前版本: $(GOLANGCI_LINT_VERSION)"; \
        echo "   要求最低: $(GOLANGCI_LINT_REQUIRED_VERSION)"; \
        echo "   请升级 golangci-lint，例如："; \
        exit 1; \
    fi

    @echo "✅ golangci-lint 版本检查通过: $(GOLANGCI_LINT_VERSION)"
```

### 安装指南

#### 方式一：使用 Go 命令

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

#### 方式三：下载二进制文件

访问 [golangci-lint Releases](https://github.com/golangci/golangci-lint/releases) 下载对应操作系统的版本。

### 使用方法

```bash
# 验证版本
make check-golangci-lint-version
```

---

## pre-commit 安装检测

### 概述

`pre-commit` 是一个框架，用于在 Git 提交前自动执行检查脚本。通过检测 pre-commit 是否已安装，确保团队的代码审查流程能够正常运行。

### Makefile 配置

```makefile
check-pre-commit-installed:
    @if command -v pre-commit >/dev/null 2>&1; then \
        echo "✅ pre-commit 已安装"; \
    else \
        echo "❌ pre-commit 未安装"; \
        exit 1; \
    fi
```

### 安装指南

#### 方式一：使用 pip

```bash
pip install pre-commit
```

#### 方式二：使用 Homebrew（macOS）

```bash
brew install pre-commit
```
#### 方式三：使用 bash进行安装（ubuntu）
```bash
sudo apt install pre-commit
```


### 使用方法

```bash
make check-pre-commit-installed
```

---

## 环境检查命令

### 完整检查

将以上三个检查命令组合，创建一个统一的环境检查命令：

```makefile
check-env: check-pre-commit-installed check-golangci-lint-version check-go-version
    @echo ""
    @echo "╔════════════════════════════════════════════╗"
    @echo "║  ✅ 所有环境检查通过                        ║"
    @echo "║  可以安全执行后续开发操作                  ║"
    @echo "╚════════════════════════════════════════════╝"
```

### 执行命令

```bash
make check-env
```

**预期输出：**
```
✅ pre-commit 已安装
🔍 当前 golangci-lint 版本: 2.7.1
✅ 要求最低版本: 2.7.1
✅ golangci-lint 版本检查通过: 2.7.1
✅ Go version check passed: 1.25.1

╔════════════════════════════════════════════╗
║  ✅ 所有环境检查通过                        ║
║  可以安全执行后续开发操作                  ║
╚════════════════════════════════════════════╝
```

---

## golangci-lint 配置

### 创建配置文件

在项目根目录创建 `.golangci.yaml` 文件。

### 配置示例

```yaml
# .golangci.yaml
version: 2

# 运行配置
run:
  timeout: 5m
  issues-exit-code: 1

# linter 配置
linters:
  # 启用的 linter
  enable:
    - misspell      # 检查拼写错误
    - govet         # Go 官方静态分析工具
  
  # 对启用的 linter 进行详细配置
  settings:
    misspell:
      locale: US    # 使用美式英语拼写规则
```

### 常用 Linter 说明

| Linter | 说明 |
|--------|------|
| `govet` | Go 官方的静态分析工具，检查常见编程错误 |
| `misspell` | 检查常见的英文单词拼写错误 |
| `revive` | 代码风格和命名规范检查 |
| `unused` | 检查未使用的变量、函数和类型 |
| `ineffassign` | 检查无效的赋值语句 |

### 官方文档

详细配置选项请参考 [golangci-lint 官方文档](https://golangci-lint.run/docs/welcome/install/)

### 运行检查

```bash
# 检查整个项目
golangci-lint run

# 检查特定目录
golangci-lint run ./pkg/...

# 使用指定配置文件
golangci-lint run --config=.golangci.yaml
```

---

## 在代码中跳过检查

在某些情况下，需要在代码中忽略某些 linter 的检查。使用 `nolint` 注释实现。

### 跳过单行检查

```go
var bad_name int //nolint:all
var bad_name int //nolint:misspell,unused
```

### 跳过多行检查

```go
//nolint:all
func allIssuesInThisFunctionAreExcluded() *string {
  // ...
}

//nolint:govet
var (
  a int
  b int
)
```

### 跳过整个文件检查

```go
//nolint:unparam
package pkg
```

> ⚠️ **建议：** 尽量避免使用 `nolint` 注释，优先通过修改代码来满足 linter 规范。频繁跳过检查会降低代码质量。

---

## pre-commit 配置

### 创建配置文件

在项目根目录创建 `.pre-commit-config.yaml` 文件，定义提交前需要执行的检查。

### 配置示例

```yaml
# .pre-commit-config.yaml

repos:
  # 使用本地已安装的 make 命令
  - repo: local
    hooks:
      - id: check-env-via-make
        name: Check Env (via Makefile)
        entry: make
        language: system
        types: []              # 对所有文件类型触发
        pass_filenames: false  # 不传递文件名给命令
        args:
          - check-env
        verbose: true

  # 使用 golangci-lint 检查
  - repo: local
    hooks:
      - id: golangci-lint
        name: GolangCI-Lint Check
        entry: golangci-lint
        language: system
        types: [go]            # 仅对 Go 文件触发
        pass_filenames: false  # 不传递文件名给命令
        args: [run, --config=.golangci.yaml]
```

### 配置参数说明

| 参数 | 说明 |
|------|------|
| `repo` | 指定钩子来源，`local` 表示使用本地命令 |
| `id` | 钩子的唯一标识符 |
| `name` | 钩子的显示名称 |
| `entry` | 要执行的命令 |
| `language` | 命令的运行环境 |
| `types` | 触发条件，为空表示所有文件类型 |
| `pass_filenames` | 是否将文件名作为参数传递给命令 |
| `args` | 传递给命令的参数 |
| `verbose` | 是否显示详细输出 |

### 安装 pre-commit 钩子

配置完成后，需要将 pre-commit 配置安装到 Git 钩子中：

```bash
pre-commit install
```

此命令会在 `.git/hooks/pre-commit` 中创建钩子脚本，后续每次执行 `git commit` 时都会自动触发配置中定义的检查。

### 使用方法

#### 1. 自动运行（提交时）

```bash
git add .
git commit -m "feat: add new feature"
```

提交时会自动运行 pre-commit 钩子。如果检查失败，提交会被阻止。

#### 2. 手动运行所有钩子

```bash
# 对所有文件运行钩子
pre-commit run --all-files

# 仅运行特定钩子
pre-commit run check-env-via-make --all-files
```

#### 3. 跳过钩子（不推荐）

```bash
git commit --no-verify
```

#### 4. 更新 pre-commit 框架

```bash
pre-commit autoupdate
```

#### 5. 卸载钩子

```bash
pre-commit uninstall
```

---

## 快速开始

### 步骤一：验证环境

```bash
# 检查所有环境
make check-env
```

### 步骤二：配置项目

在项目根目录添加以下文件：
- `Makefile`（包含上述检查命令）
- `.golangci.yaml`（golangci-lint 配置）
- `.pre-commit-config.yaml`（pre-commit 配置）

### 步骤三：安装 Git 钩子

```bash
pre-commit install
```

### 步骤五：开发与提交

```bash
# 编写代码
vi your_file.go

# 提交代码（自动运行 pre-commit 钩子）
git add .
git commit -m "feat: implement new feature"

# 如果检查失败，修复问题后重新提交
git add .
git commit -m "feat: implement new feature"
```

### 常用命令速查

```bash
# 检查本地环境
make check-env

# 运行 golangci-lint 检查
golangci-lint run

# 手动运行所有 pre-commit 钩子
pre-commit run --all-files

# 查看 pre-commit 配置
cat .pre-commit-config.yaml

# 跳过钩子提交（谨慎使用）
git commit --no-verify
```

---

## 常见问题

### Q1：golangci-lint 无法找到或安装失败怎么办？

**A：** 首先检查安装路径，然后重新安装：

```bash
which golangci-lint  # 查看是否已安装及安装位置
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

如果仍无法解决，请访问 [golangci-lint 官方网站](https://golangci-lint.run/docs/welcome/install/) 查看详细安装说明。

### Q2：pre-commit 钩子不执行怎么办？

**A：** 重新安装 pre-commit 钩子：

```bash
pre-commit uninstall
pre-commit install
```

### Q3：如何在特定情况下跳过某个检查？

**A：** 有两种方式：

1. **在代码中添加注释**（针对单个检查）：
   ```go
   var bad_name int //nolint:misspell
   ```

2. **临时跳过 Git 钩子**（针对整个提交）：
   ```bash
   git commit --no-verify
   ```

> ⚠️ 建议优先修改代码而不是跳过检查。

### Q4：Go 版本需要更改怎么办？

**A：** 修改 `Makefile` 中的 `GO_VERSION` 变量，然后重新运行检查：

```makefile
GO_VERSION := 1.26.0  # 修改为新版本
```

```bash
make check-go-version
```

### Q5：如何更新 golangci-lint 到最新版本？

**A：** 取决于安装方式：

```bash
# Homebrew
brew upgrade golangci-lint

# Go install
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

---

## 相关资源

- **golangci-lint 官网：** https://golangci-lint.run/
- **pre-commit 官网：** https://pre-commit.com/
- **Go 官方网站：** https://golang.org/
- **golangci-lint GitHub：** https://github.com/golangci/golangci-lint
- **pre-commit GitHub：** https://github.com/pre-commit/pre-commit

---

**文档版本：** v1.0  