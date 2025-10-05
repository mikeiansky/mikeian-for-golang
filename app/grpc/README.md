## 环境准备
- 检查 protoc --version
- 安装地址官网地址 https://github.com/protocolbuffers/protobuf/releases
- 检查 protoc-gen-go --version
- 安装 go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
- 检查 protoc-gen-go-grpc --version
- 安装 go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

## 生成代码
- 在项目根目录下面执行protoc --go_out=paths=source_relative:./ --go-grpc_out=paths=source_relative:./ ./app/grpc/api/api.proto
