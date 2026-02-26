PROTO_PATH := C:/Users/sadra/Desktop/code/block-chain-golang/proto

build: 
	@go build -o bin/block-chain-golang
run: build
	@./bin/docker	 
test:
	@go test -v ./...
proto:
	protoc \
	  --go_out=. \
	  --proto_path=$(PROTO_PATH) \
	  $(PROTO_PATH)/*.proto
.PHONY: proto	