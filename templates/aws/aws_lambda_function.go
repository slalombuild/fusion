package aws

import (
	_ "embed"
	"io"

	"github.com/pkg/errors"
	"github.com/slalombuild/fusion/templates"
)

//go:embed aws_lambda_function.tmpl
var TEMPLATE_AWS_LAMBDA_FUNCTION string

// LambdaFunction is the template data object used to create
// a lambda function
type LambdaFunction struct {
	Name     string `help:"" default:"my_lambda"`
	Filename string `help:"" default:"my_lambda.zip"`
	Handler  string `help:"" default:"handler.index.js"`
	Runtime  string `help:"" default:"nodejs14.x"`
}

// NewLambdaFunction creates new lambda function
func NewLambdaFunction(name, filename, handler, runtime string) *LambdaFunction {
	return &LambdaFunction{
		Name:     name,
		Filename: filename,
		Handler:  handler,
		Runtime:  runtime,
	}
}

// Render generates the Terraform code for the LambdaFunction
func (resource *LambdaFunction) Render(w io.Writer, skipColor bool) error {
	output, err := templates.Execute(TEMPLATE_AWS_LAMBDA_FUNCTION, &resource)
	if err != nil {
		return errors.Wrap(err, "failed to generate template")
	}

	err = templates.Highlight(w, output.String(), templates.HCL, skipColor)
	if err != nil {
		return errors.Wrap(err, "failed to highlight source")
	}

	return nil
}
