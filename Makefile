PROTO=server/proto/tasks.proto
GOCMD=go
GOGET=$(GOCMD) get

proto:
	protoc --proto_path=$(GOPATH)/src:. --micro_out=. --go_out=. $(PROTO)

deps:
	$(GOGET) -v ...