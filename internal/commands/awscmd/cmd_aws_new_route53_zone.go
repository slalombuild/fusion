package awscmd

import (
	commands "github.com/slalombuild/fusion/internal/commands"
	aws "github.com/slalombuild/fusion/templates/aws"
)

// NewRoute53ZoneCmd creates a new route53_zone
type NewRoute53ZoneCmd struct {
	Globals
	*aws.Route53HostedZone `embed:""`
}

func (cmd *NewRoute53ZoneCmd) Run(ctx *commands.Context) error {
	return cmd.Render(ctx.Output, ctx.NoColor.Bool())
}
