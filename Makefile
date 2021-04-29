## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

test: 
	export API_URL=http://localhost:3001; \
	go test -count=1 -v ./... 

## build-wasm: builds the go code to wasm
build-wasm: 
	GOOS=js GOARCH=wasm go build -o ./wasm/assets/main.wasm ./wasm/main.go

## run-wasm: runs a server with wasm loaded on :8080
run-wasm: build-wasm
	@go run ./wasm/server/main.go
	