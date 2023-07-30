package awscmd

import (
	commands "github.com/slalombuild/fusion/internal/commands"
	"github.com/slalombuild/fusion/templates/aws/stacks"
)

// NewAWSServerlessWebsiteCmd creates a new ServerlessWebsite
type NewAWSServerlessWebsiteCmd struct {
	Globals
	*stacks.ServerlessWebsite `embed:""`
}

func (cmd *NewAWSServerlessWebsiteCmd) Run(ctx *commands.Context) error {
	return cmd.Render(ctx.Output, ctx.NoColor.Bool())
}
