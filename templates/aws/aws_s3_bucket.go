package aws

import (
	_ "embed"
	"io"

	"github.com/SlalomBuild/fusion/templates"
	errors "github.com/pkg/errors"
)

//go:embed aws_s3_bucket.tmpl
var TEMPLATE_AWS_S3_BUCKET string

/*
View the Kong CLI docs to see options for your command's flags and arguments
https://github.com/alecthomas/kong#flags
*/
type S3Bucket struct {
	Name         string `help:"Name of the S3 Bucket" default:"fusion-bucket" required:"true"`
	StaticSite   bool   `help:"Enable static website hosting from bucket"`
	ForceDestroy bool   `help:"Destroy objects on bucket deletion" default:"false"`
}

// Render generates the Terraform code for the S3Bucket
func (resource *S3Bucket) Render(w io.Writer, skipColor bool) error {
	output, err := templates.Execute(TEMPLATE_AWS_S3_BUCKET, &resource)
	if err != nil {
		return errors.Wrap(err, "failed to generate template")
	}
	err = templates.Highlight(w, output.String(), templates.HCL, skipColor)
	if err != nil {
		return errors.Wrap(err, "failed to highlight source")
	}
	return nil
}
