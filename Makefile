build: 
	@go build -o bin/block-chain-golang
run: build
	@./bin/docker	 
test:
	@go test -v ./...