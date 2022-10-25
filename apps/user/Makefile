
serviceName := user
env := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
.PHONY: proto build

proto:
	goctl rpc protoc ./$(serviceName)/$(serviceName).proto --go_out=./$(serviceName) --go-grpc_out=./$(serviceName) --zrpc_out=./$(serviceName)

build:
	$(env) go build -o $(serviceName) ./$(serviceName)/*.go
	@echo " $(serviceName) build success"