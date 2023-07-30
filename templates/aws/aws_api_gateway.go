package aws

import (
	_ "embed"
	"io"

	errors "github.com/pkg/errors"
	"github.com/slalombuild/fusion/templates"
)

//go:embed aws_api_gateway.tmpl
var TEMPLATE_AWS_API_GATEWAY string

/*
View the Kong CLI docs to see options for your command's flags and arguments
https://github.com/alecthomas/kong#flags
*/
type APIGateway struct {
	LogGroupArn    string `help:"If a CloudWatch log group already exists, put its arn here" short:"a" xor:"LogGroupArn,CreateLogGroup" required:""`
	Name           string `help:"API gateway name" required:"" default:"Fusion API gateway"`
	CreateLogGroup bool   `help:"If a CloudWatch log group does not already exist, set this to true to create one" short:"c" xor:"LogGroupArn,CreateLogGroup" enum:"true,false" default:"true" required:""`
}

// Render generates the Terraform code for the APIGatewayv2
func (resource *APIGateway) Render(w io.Writer, skipColor bool) error {
	output, err := templates.Execute(TEMPLATE_AWS_API_GATEWAY, &resource)
	if err != nil {
		return errors.Wrap(err, "failed to generate template")
	}
	err = templates.Highlight(w, output.String(), templates.HCL, skipColor)
	if err != nil {
		return errors.Wrap(err, "failed to highlight source")
	}
	return nil
}
