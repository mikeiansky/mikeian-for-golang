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


# 目标 golangci-lint 版本（比如要求最低是 2.7.1）
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
		echo "   - 使用 brew upgrade golangci-lint"; \
		echo "   - 或下载最新 release: https://github.com/golangci/golangci-lint/releases"; \
		exit 1; \
	fi

	@echo "✅ golangci-lint 版本检查通过: $(GOLANGCI_LINT_VERSION)"



check-pre-commit-installed:
	@if command -v pre-commit >/dev/null 2>&1; then \
		echo "✅ pre-commit 已安装"; \
	else \
		echo "❌ pre-commit 未安装"; \
		echo "🔧 请通过以下命令安装："; \
		echo "   pip install pre-commit"; \
		echo "   或访问：https://pre-commit.com/#installation"; \
		exit 1; \
	fi


check-env: check-pre-commit-installed check-golangci-lint-version check-go-version
	@echo "✅ 所有环境检查通过，可以安全执行后续操作"

