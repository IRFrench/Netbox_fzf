.DEFAULT_GOAL := help
.PHONY: build

run: ##	Run the service
	TOKEN=test CGO_ENABLED=0 go run cmd/main.go

build: ##	Build a binary for the service
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o build/ssh_builder cmd/main.go

help: ##	Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_.-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)