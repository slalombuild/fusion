package awscmd

import (
	commands "github.com/SlalomBuild/fusion/internal/commands"
	aws "github.com/SlalomBuild/fusion/templates/aws"
)

// NewIamPolicyCmd creates a new iam_policy
type NewIamPolicyCmd struct {
	Globals
	*aws.IamPolicy `embed:""`
}

func (cmd *NewIamPolicyCmd) Run(ctx *commands.Context) error {
	return cmd.Render(ctx.Output, ctx.NoColor.Bool())
}
