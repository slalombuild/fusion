package awscmd

import (
	commands "github.com/SlalomBuild/fusion/internal/commands"
	aws "github.com/SlalomBuild/fusion/templates/aws"
)

// NewSecurityGroupCmd creates a new SecurityGroup
type NewSecurityGroupCmd struct {
	Globals
	*aws.SecurityGroup `embed:""`
}

func (cmd *NewSecurityGroupCmd) Run(ctx *commands.Context) error {
	return cmd.Render(ctx.Output, ctx.NoColor.Bool())
}
