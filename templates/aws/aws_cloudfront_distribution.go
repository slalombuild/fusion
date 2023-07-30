package aws

import (
	_ "embed"
	"io"

	errors "github.com/pkg/errors"
	"github.com/slalombuild/fusion/templates"
)

//go:embed aws_cloudfront_distribution.tmpl
var TEMPLATE_AWS_CLOUDFRONT_DISTRIBUTION string

/*
View the Kong CLI docs to see options for your command's flags and arguments
https://github.com/alecthomas/kong#flags
*/
type CloudfrontDistribution struct {
	LoggingBucketName string `help:"Name for a new s3 bucket that will contain all cloudfront logs"`
	OriginBucketName  string `help:"Name of the origin s3 bucket"`
	CloudfrontAlias   string `help:"DNS alias for cloudfront"`
}

// Render generates the Terraform code for the CloudFront
func (resource *CloudfrontDistribution) Render(w io.Writer, skipColor bool) error {
	output, err := templates.Execute(TEMPLATE_AWS_CLOUDFRONT_DISTRIBUTION, &resource)
	if err != nil {
		return errors.Wrap(err, "failed to generate template")
	}
	err = templates.Highlight(w, output.String(), templates.HCL, skipColor)
	if err != nil {
		return errors.Wrap(err, "failed to highlight source")
	}
	return nil
}
