package aws

import (
	_ "embed"
	"io"

	errors "github.com/pkg/errors"
	"github.com/slalombuild/fusion/templates"
)

//go:embed aws_iam_policy.tmpl
var TEMPLATE_AWS_IAM_POLICY string

/*
View the Kong CLI docs to see options for your command's flags and arguments
https://github.com/alecthomas/kong#flags
*/
type IamPolicy struct {
	Name        string      `help:"Name of the IAM policy" required:""`
	Description string      `help:"Description of the IAM policy" default:"" short:"d"`
	Path        string      `help:"Path of the IAM Policy" default:"/" short:"p"`
	PolicyJSON  interface{} `help:"Body of the IAM policy via json file" required:"" short:"j" type:"jsonfile"`
}

// Render generates the Terraform code for the IamPolicy
func (resource *IamPolicy) Render(w io.Writer, skipColor bool) error {
	output, err := templates.Execute(TEMPLATE_AWS_IAM_POLICY, &resource)
	if err != nil {
		return errors.Wrap(err, "failed to generate template")
	}
	err = templates.Highlight(w, output.String(), templates.HCL, skipColor)
	if err != nil {
		return errors.Wrap(err, "failed to highlight source")
	}
	return nil
}
