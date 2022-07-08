package aws

import (
	_ "embed"
	"io"

	templates "github.com/SlalomBuild/fusion/templates"
	errors "github.com/pkg/errors"
)

//go:embed aws_securitygroup.tmpl
var TEMPLATE_AWS_SECURITY_GROUP string

/*
View the Kong CLI docs to see options for your command's flags and arguments
https://github.com/alecthomas/kong#flags
*/
type SecurityGroup struct {
	PublicIngressPort     int    `help:"Public ingress port for the security group"`
	PublicIngressProtocol string `help:"Public ingress protocol for the security group"`
}

// Render generates the Terraform code for the SecurityGroup
func (resource *SecurityGroup) Render(w io.Writer, skipColor bool) error {
	output, err := templates.Execute(TEMPLATE_AWS_SECURITY_GROUP, &resource)
	if err != nil {
		return errors.Wrap(err, "failed to generate template")
	}
	err = templates.Highlight(w, output.String(), templates.HCL, skipColor)
	if err != nil {
		return errors.Wrap(err, "failed to highlight source")
	}
	return nil
}
