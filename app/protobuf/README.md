## 前提
- 必须要有**protoc**环境，通过`protoc --version`来检测是否存在，1否则去对应的github上面进行下载
- protobuf 官网地址 https://github.com/protocolbuffers/protobuf/releases
- 安装protoc-gen-go，通过protoc-gen-go --version检测是否存在，不存在则通过下面命令进行安装，这个命令是安装再gopath下面的，因为gopath被导入的系统环境，所以能够直接作为命令工具使用
- go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

## 命令
在项目根目录环境下执行下面命令
- protoc --go_out=paths=source_relative:./  ./app/protobuf/city/city.proto
- protoc --go_out=paths=source_relative:./  ./app/protobuf/person/person.proto

## 配置 --proto_path之后的影响
-  protoc  --proto_path=./third_party --proto_path=./internal --proto_path=./app --go_out=paths=source_relative:./app/  ./app/protobuf/person/person.proto