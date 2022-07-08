//go:build tools
// +build tools

package main

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/goreleaser/goreleaser"
	_ "github.com/loeffel-io/ls-lint"
	_ "github.com/princjef/gomarkdoc"
)
