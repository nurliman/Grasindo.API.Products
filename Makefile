.PHONY: compile
compile: ## Compile the proto file.
	protoc -I protos protos/product.proto --go_out=plugins=grpc:protos/
 
.PHONY: server
server: ## Build and run server.
	go build -race -ldflags "-s -w" -o bin/server server/main.go
	bin/server