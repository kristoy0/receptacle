PROTO = server/proto/tasks.proto
GOCMD = go
GOGET = $(GOCMD) get
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOARCH = amd64
RELEASE = release/receptacle
PROJECTROOT = github.com/kristoy0/receptacle

proto:
	protoc --proto_path=$(GOPATH)/src:. --micro_out=. --go_out=. $(PROTO)

deps:
	$(GOGET) -v ...

build-server:
	GOOS=linux GOARCH=$(GOARCH) CGO_ENABLED=0 $(GOBUILD) -o $(RELEASE)-server $(PROJECTROOT)/server

build-agent:
	GOOS=linux GOARCH=$(GOARCH) CGO_ENABLED=0 $(GOBUILD) -o $(RELEASE)-agent $(PROJECTROOT)/agent

build-client:
	GOOS=linux GOARCH=$(GOARCH) CGO_ENABLED=0 $(GOBUILD) -o $(RELEASE)-client-linux-$(GOARCH) $(PROJECTROOT)/client
	GOOS=linux GOARCH=$(GOARCH) CGO_ENABLED=0 $(GOBUILD) -o $(RELEASE)-client-windows-$(GOARCH) $(PROJECTROOT)/client