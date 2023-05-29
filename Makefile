# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
WEB_SERVER_PATH=./cmd/web-server/
BOT_PATH=./cmd/bot/
BINARY_PATH=./build/
WEB_SERVER_BINARY_NAME=server
WEB_SERVER_BINARY_UNIX=$(WEB_SERVER_BINARY_NAME)_unix
BOT_BINARY_NAME=bot
BOT_BINARY_UNIX=$(BOT_BINARY_NAME)_unix


all: test build

test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	@$(GORUN) $(CMD_PATH)*.go
run-local:
	@$(GORUN) $(CMD_PATH)*.go -c config/local.yaml -metrics_server_addr :9091
deps:
	$(GOGET) -u ./...

gen-api:
	protoc -I api/ --go_out=api/ --go-grpc_out=api/ api/bot_service.proto


# Cross compilation
.PHONY: build
build-web-server:
	CGO_ENABLED=0 \
		$(GOBUILD) \
			-installsuffix cgo \
			-o $(BINARY_PATH)$(WEB_SERVER_BINARY_NAME) \
			$(WEB_SERVER_PATH)*.go

build-linux-web-server:
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
		$(GOBUILD) \
			-installsuffix cgo \
			-o $(BINARY_PATH)$(WEB_SERVER_BINARY_NAME).linux \
			-ldflags "-X main.Version=$(APP_VERSION)" \
			$(WEB_SERVER_PATH)*.go


build-bot:
	CGO_ENABLED=0 \
		$(GOBUILD) \
			-installsuffix cgo \
			-o $(BINARY_PATH)$(BOT_BINARY_NAME) \
			-ldflags "-X main.Version=$(APP_VERSION)" \
			$(BOT_PATH)*.go

build-linux-bot:
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
		$(GOBUILD) \
			-installsuffix cgo \
			-o $(BINARY_PATH)$(BOT_BINARY_NAME).linux \
			-ldflags "-X main.Version=$(APP_VERSION)" \
			$(BOT_PATH)*.go


