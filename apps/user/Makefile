
serviceName := user
env := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
.PHONY: proto build sql

proto:
	goctl rpc protoc ./$(serviceName)/$(serviceName).proto --go_out=./$(serviceName) --go-grpc_out=./$(serviceName) --zrpc_out=./$(serviceName)

build:
	$(env) go build -o $(serviceName) ./$(serviceName)/*.go
	@echo " $(serviceName) build success"

sql:
	goctl model mysql ddl -src ./model/user.sql -dir ./model -c