package stacks

import (
	_ "embed"
	"io"

	templates "github.com/SlalomBuild/fusion/templates"
	errors "github.com/pkg/errors"
)

//go:embed aws_serverless_website.tmpl
var TEMPLATE_AWS_SERVERLESS_WEBSITE string

/*
View the Kong CLI docs to see options for your command's flags and arguments
https://github.com/alecthomas/kong#flags
*/
type ServerlessWebsite struct {
	Route53Zone             string `help:"Name of the Route53 Zone" default:"Route53Zone"`
	Route53Target           string `help:"Target for Route53" default:"Route53Target"`
	APIGatewayName          string `help:"Name of the API Gateway" default:"APIGatewayName"`
	LambdaName              string `help:"Name of the Lambda function" default:"LambdaName"`
	VpcName                 string `help:"Name of the VPC" default:"fusion-vpc"`
	CloudfrontLoggingBucket string `help:"Name of the Cloudfront S3 logging bucket" default:"CloudfrontLoggingBucket"`
	CloudfrontAlias         string `help:"Alias for Cloudfront" default:"CloudfrontAlias"`
	AwsRegion               string `help:"Provider Region" default:"us-east-1"`
	VpcCiderBlock           string `help:"Initial CIDR block for the VPC" default:"VPCCiderBlock"`
	VpcAzCount              int    `help:"Number of AZs" default:"2"`
	VpcSubnetCapacity       int    `help:"Capacity in the created subnet" default:"256"`
	LambdaFilePath          string `help:"Filepath for the Lambda Function" default:"LambdaFilePath"`
	CloudfrontOriginBucket  string `help:"Name of the Cloudfront S3 origin bucket" default:"CloudfrontOriginBucket"`
}

// Render generates the Terraform code for the ServerlessStack
func (resource *ServerlessWebsite) Render(w io.Writer, skipColor bool) error {
	output, err := templates.Execute(TEMPLATE_AWS_SERVERLESS_WEBSITE, &resource)
	if err != nil {
		return errors.Wrap(err, "failed to generate template")
	}
	err = templates.Highlight(w, output.String(), templates.HCL, skipColor)
	if err != nil {
		return errors.Wrap(err, "failed to highlight source")
	}
	return nil
}
