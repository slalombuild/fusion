BUILD_DIR=bin

.PHONY : help
help : Makefile
	@sed -n 's/^##//p' $<

.PHONY: install
## install : Install fusion cli
install: 
	@echo "Installing fusion cli"
	go build -o "${GOPATH}/bin/fusion" -ldflags="-s -w" cmd/fusion/main.go
	@echo "Installing fusionctl cli"
	go build -o "${GOPATH}/bin/fusionctl" -ldflags="-s -w" cmd/fusionctl/main.go

.PHONY: fmt
 ## fmt    : Format Go source code
fmt:
	@go fmt ./...

.PHONY: tools
## tools   : Install development tools
tools:
	@echo "Installing dev tools"
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %
	@echo ""

.PHONY: clean
## clean   : Clean generated files and test cache
clean:
	@rm -rf $(BUILD_DIR)
	@go clean -testcache

.PHONT: lint
## lint	   : Lint code
lint: tools
	ls-lint
	golangci-lint run ./...

.PHONY: test
## test    : Run unit tests
test: clean tools
	@echo "🧪\tTesting" && go test ./... -coverprofile cover.out
	@echo
	@echo "⛱️\tCoverage" && go tool cover -func cover.out

.PHONY: build
## build   : Build fusion executable
build: tools
	goreleaser build --rm-dist --skip-validate

.PHONY: demo
## demo	   : Run a demo of the fusion project
demo:
	@go run cmd/demo/main.go --all