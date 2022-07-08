## Development

View the available Makefile commands for building, testing, installing, and formatting the application with `make help`

### Requirements

Before running the test suite or building the application, install development tools with: 

```
# Install the dev tools
make tools
```

```
make help

install : Install fusion and fusionctl cli
tools   : Install development tools
clean   : Clean generated files and test cache
test    : Run unit tests
build   : Build executables with goreleaser
```