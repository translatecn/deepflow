MESSAGE = github.com/deepflowio/deepflow/message
REV_COUNT = $(shell git rev-list --count HEAD)
COMMIT_DATE = $(shell git show -s --format=%cd --date=short HEAD)
REVISION = $(shell git rev-parse HEAD)
BRANCH = $(shell git branch  --show-current)
COMPILE_TIME= $(shell date +"%Y-%m-%d %H:%M:%S")
# build tags: [-tags=go-json] used for [gin-gonic/gin] marshal json
FLAGS = -gcflags "-l -l" -ldflags "-X main.Branch=${BRANCH} -X main.RevCount=${REV_COUNT} -X main.Revision=${REVISION} -X main.CommitDate=${COMMIT_DATE} \
		-X 'main.goVersion=$(shell go version)' -X 'main.CompileTime=${COMPILE_TIME}'" \
		-tags=go_json
BINARY_SUFFIX :=

.PHONY: all
all: server



generate_message:
	go generate ../message/common/stub.go
	go generate ../message/trident/stub.go
	go generate ../message/controller/stub.go
	go generate ../message/alarm_event/stub.go
	go generate ../message/k8s_event/stub.go

	cp -r ../message/flow_log.proto libs/datatype/pb/
	# delete first 3 char of line 6-10 for go language
	gsed -i '6,10s/^...//' libs/datatype/pb/flow_log.proto

	cp -r ../message/metric.proto libs/flow-metrics/pb
	# delete first 3 char of line 6-10 for go language
	gsed -i '6,10s/^...//' libs/flow-metrics/pb/metric.proto

	cp ../message/stats.proto libs/stats/pb/stats.proto
	# delete first 3 char of line 6-10 for go language
	gsed -i '6,10s/^...//' libs/stats/pb/stats.proto

	go generate ./...




libs/geo/ip_info.go: libs/geo/ip_info_mini.json libs/geo/ip_info.py
	go generate ./...

vendor:
#	go install github.com/gogo/protobuf/protoc-gen-gofast@latest
#	go install github.com/gogo/protobuf/protoc-gen-gogo@latest
#	go install github.com/golang/protobuf/protoc-gen-go@latest
# 	go install github.com/benbjohnson/tmpl@latest

#	 GO111MODULE="off" GOPATH='/Users/acejilam/.gopath' /Users/acejilam/.gvm/gos/go1.18/bin/go get github.com/gogo/protobuf/protoc-gen-gofast
#	 GO111MODULE="off" GOPATH='/Users/acejilam/.gopath' /Users/acejilam/.gvm/gos/go1.18/bin/go get github.com/gogo/protobuf/proto
#	 GO111MODULE="off" GOPATH='/Users/acejilam/.gopath' /Users/acejilam/.gvm/gos/go1.18/bin/go get github.com/gogo/protobuf/jsonpb
#	 GO111MODULE="off" GOPATH='/Users/acejilam/.gopath' /Users/acejilam/.gvm/gos/go1.18/bin/go get github.com/gogo/protobuf/protoc-gen-gogo
#	 GO111MODULE="off" GOPATH='/Users/acejilam/.gopath' /Users/acejilam/.gvm/gos/go1.18/bin/go get github.com/gogo/protobuf/gogoproto
#	 GO111MODULE="off" GOPATH='/Users/acejilam/.gopath' /Users/acejilam/.gvm/gos/go1.18/bin/go get github.com/golang/protobuf/protoc-gen-go

	go mod tidy && go mod download && go mod vendor
	gfind ../vendor -type d -exec chmod +w {} \;
	# optimize the write performance of Array(string) data type
	patch -p1 -d ../vendor/github.com/ClickHouse/clickhouse-go/v2/ < patch/clickhouse-go/write_improve.patch
	# reset the pointer to the nullable field
	patch -p1 -d ../vendor/github.com/ClickHouse/clickhouse-go/v2/ < patch/clickhouse-go/nullable.patch
	# optimize write datatime
	patch -p1 -d ../vendor/github.com/ClickHouse/clickhouse-go/v2/ < patch/clickhouse-go/datetime_uint32.patch
	# optimize the performance of hashmap when writing a large amount of data
	patch -p1 -d ../vendor/github.com/cornelk/hashmap < patch/cornelk_hashmap/perf.patch
	# hashmap support key [2]uint64
	patch -p1 -d ../vendor/github.com/cornelk/hashmap < patch/cornelk_hashmap/complex128.patch

.PHONY: test
test: vendor generate_message $(generated_libs)
	go test -mod vendor -short ./... -timeout 5s -coverprofile .test-coverage.txt
	go tool cover -func=.test-coverage.txt

.PHONY: server
server: vendor generate_message
	go build ${FLAGS} -o bin/deepflow-server${BINARY_SUFFIX} cmd/server/

.PHONY: querier
querier: vendor generate_message
	go build -mod vendor ${FLAGS} -o bin/querier${BINARY_SUFFIX} querier/cmd/querier/main.go

.PHONY: clean
clean:
	touch vendor
	chmod -R 777 vendor
	rm -rf vendor
	rm -rf bin
	rm -rf .test-coverage.txt
	gfind . -name '*.pb.go' | grep -v libs/datatype/prompb/ | xargs -i rm -f {}
