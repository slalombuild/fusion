package gcpcmd

import (
	"github.com/SlalomBuild/fusion/templates/gcp"
	"github.com/pkg/errors"

	"github.com/SlalomBuild/fusion/internal/commands"
	"github.com/rs/zerolog/log"
)

var (
	ErrGenLoadbalancer = "failed to generate terraform for loadbalancer"
	ErrHighlight       = "error highlighting terraform"
)

type NewLoadBalancerCommand struct {
	*gcp.InternalLoadBalancer `embed:""`
}

func (cmd *NewLoadBalancerCommand) Run(ctx *commands.Context) error {
	log.Info().Str("provider", "gcp").Str("resource", "loadbalancer").Interface("data", cmd.InternalLoadBalancer).Send()
	err := cmd.Render(ctx.Output, ctx.NoColor.Bool())
	if err != nil {
		return errors.Wrap(err, ErrGenLoadbalancer)
	}

	return nil
}
