package awscmd

import (
	commands "github.com/SlalomBuild/fusion/internal/commands"
	aws "github.com/SlalomBuild/fusion/templates/aws"
)

// NewVpcStackCmd creates a new vpc_stack
type NewVpcStackCmd struct {
	Globals
	*aws.VPC `embed:""`
}

func (cmd *NewVpcStackCmd) Run(ctx *commands.Context) error {
	return cmd.Render(ctx.Output, ctx.NoColor.Bool())
}
