
serviceName := api
env := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
.PHONY: api build

api:
	goctl api go -api ./$(serviceName)/$(serviceName).api -dir ./$(serviceName)

build:
	$(env) go build -o $(serviceName) ./$(serviceName)/*.go
	@echo " $(serviceName) build success"