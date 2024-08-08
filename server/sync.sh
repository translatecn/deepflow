# /usr/bin/env zsh
ARCH=$(arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/)
if [[ ${ARCH} == "arm64" ]]; then
  export https_proxy=http://10.211.51.30:7890 http_proxy=http://10.211.51.30:7890 all_proxy=socks5://10.211.51.30:7890
fi

if [[ ${ARCH} == "amd64" ]]; then
  export https_proxy=http://10.230.205.190:7890 http_proxy=http://10.230.205.190:7890 all_proxy=socks5://10.230.205.190:7890
fi

[[ -s "/Users/acejilam/.gvm/scripts/gvm" ]] && source "/Users/acejilam/.gvm/scripts/gvm"
[[ -s "/root/.gvm/scripts/gvm" ]] && source "/root/.gvm/scripts/gvm"

gvm use go1.18
go version
if [[ $(go env GOHOSTOS) == "linux" ]]; then
  GO111MODULE="off" GOPATH='/root/.go' go get github.com/gogo/protobuf/protoc-gen-gofast
  GO111MODULE="off" GOPATH='/root/.go' go get github.com/gogo/protobuf/proto
  GO111MODULE="off" GOPATH='/root/.go' go get github.com/gogo/protobuf/jsonpb
  GO111MODULE="off" GOPATH='/root/.go' go get github.com/gogo/protobuf/protoc-gen-gogo
  GO111MODULE="off" GOPATH='/root/.go' go get github.com/gogo/protobuf/gogoproto
  GO111MODULE="off" GOPATH='/root/.go' go get github.com/golang/protobuf/protoc-gen-go
fi

if [[ $(go env GOHOSTOS) == "darwin" ]]; then
  which gvm
  GO111MODULE="off" GOPATH='/Users/acejilam/.gopath' go get github.com/gogo/protobuf/protoc-gen-gofast
  GO111MODULE="off" GOPATH='/Users/acejilam/.gopath' go get github.com/gogo/protobuf/proto
  GO111MODULE="off" GOPATH='/Users/acejilam/.gopath' go get github.com/gogo/protobuf/jsonpb
  GO111MODULE="off" GOPATH='/Users/acejilam/.gopath' go get github.com/gogo/protobuf/protoc-gen-gogo
  GO111MODULE="off" GOPATH='/Users/acejilam/.gopath' go get github.com/gogo/protobuf/gogoproto
  GO111MODULE="off" GOPATH='/Users/acejilam/.gopath' go get github.com/golang/protobuf/protoc-gen-go
fi

gvm use system
go version

unset https_proxy && unset http_proxy && unset all_proxy
