package aws

import (
	_ "embed"
	"io"

	errors "github.com/pkg/errors"
	"github.com/slalombuild/fusion/templates"
)

//go:embed aws_route53_hosted_zone.tmpl
var TEMPLATE_AWS_ROUTE53_ZONE string

/*
View the Kong CLI docs to see options for your command's flags and arguments
https://github.com/alecthomas/kong#flags
*/
type Route53HostedZone struct {
	Target string `help:"Target for the zone"`
	Zone   string `help:"Zone name"`
}

// Render generates the Terraform code for the Route53Stack
func (resource *Route53HostedZone) Render(w io.Writer, skipColor bool) error {
	output, err := templates.Execute(TEMPLATE_AWS_ROUTE53_ZONE, &resource)
	if err != nil {
		return errors.Wrap(err, "failed to generate template")
	}
	err = templates.Highlight(w, output.String(), templates.HCL, skipColor)
	if err != nil {
		return errors.Wrap(err, "failed to highlight source")
	}
	return nil
}
