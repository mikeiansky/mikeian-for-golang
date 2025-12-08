# 目标 Go 版本
GO_VERSION := 1.25.1

# 获取当前系统 Go 版本（例如：go version go1.25.1 linux/amd64 → 提取 1.25.1）
CURRENT_GO_VERSION := $(shell go version | awk '{print $$3}' | sed 's/go//')


# 主构建目标，依赖于 check-env
build: check-env
	@echo "✅ Go version is correct: $(CURRENT_GO_VERSION)"
	@echo "🚀 Starting build..."
	# 这里放你的实际构建命令，例如：
	# go build -o myapp .

# 运行目标，也依赖版本检查
run: check-env
	@echo "✅ Running application with Go $(CURRENT_GO_VERSION)"
	# go run main.go

check-env:
	@if [ "$(CURRENT_GO_VERSION)" != "$(GO_VERSION)" ]; then \
    		echo "❌ Error: Go version mismatch!"; \
    		echo "   Expected: $(GO_VERSION)"; \
    		echo "   Found:    $(CURRENT_GO_VERSION)"; \
    		echo "   Please install or switch to Go $(GO_VERSION)."; \
    		exit 1; \
	fi
	@echo "✅ Go version check passed: $(CURRENT_GO_VERSION)"

# 测试目标
test: check-env
	@echo "✅ Running tests with Go $(CURRENT_GO_VERSION)"

.PHONY: hello
hello:
	@echo hello

