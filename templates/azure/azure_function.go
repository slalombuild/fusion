package azure

import (
	_ "embed"
)

//go:embed azure_function.tmpl
var TEMPLATE_AZURE_FUNCTION string

// AzureFunction is the template data object used to create
// an azure function
type AzureFunction struct {
	// Implement me
	// Name     string `help:"" default:"my_function"`
	// Filename string `help:"" default:"my_lambda.zip"`
	// Handler  string `help:"" default:"handler.index.js"`
	// Runtime  string `help:"" default:"NODE_14.x"`
}

// NewAzureFunction creates new azure function data
func NewAzureFunction() *AzureFunction {
	// implement me
	return &AzureFunction{}
}
