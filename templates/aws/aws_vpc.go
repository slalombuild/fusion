package aws

import (
	_ "embed"
	"io"

	errors "github.com/pkg/errors"
	"github.com/slalombuild/fusion/templates"
)

//go:embed aws_vpc.tmpl
var TEMPLATE_AWS_VPC string

/*
View the Kong CLI docs to see options for your command's flags and arguments
https://github.com/alecthomas/kong#flags
*/
type VPC struct {
	CidrBlock      string `help:"CIDR block for the VPC, must be able to accomodate all the hosts the AZs will have" default:"10.0.0.0/16"`
	AzCount        int    `help:"Number of AZs to deploy in, will automatically distribute across AZs in that region" default:"4"`
	SubnetCapacity int    `help:"Number of ip addresses per subnet" default:"256"`
}

// Render generates the Terraform code for the VpcStackModule
func (resource *VPC) Render(w io.Writer, skipColor bool) error {
	output, err := templates.Execute(TEMPLATE_AWS_VPC, &resource)
	if err != nil {
		return errors.Wrap(err, "failed to generate template")
	}
	err = templates.Highlight(w, output.String(), templates.HCL, skipColor)
	if err != nil {
		return errors.Wrap(err, "failed to highlight source")
	}
	return nil
}
