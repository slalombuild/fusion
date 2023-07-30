package azurecmd

import (
	"github.com/pkg/errors"
	"github.com/slalombuild/fusion/internal/commands"
	"github.com/slalombuild/fusion/templates/azure"

	"github.com/rs/zerolog/log"
)

var (
	ErrGenVnet   = "failed to generate terraform for vnet"
	ErrHighlight = "error highlighting terraform"
)

// NewVnetCmd creates a new vnet
type NewVnetCommand struct {
	*azure.Vnet `embed:""`
}

func (cmd *NewVnetCommand) Run(ctx *commands.Context) error {
	log.Warn().Msg("The following template contains a variable that should be separated out into a different file (ex. variables.tf).")
	log.Info().Str("provider", "azure").Str("resource", "vnet").Interface("data", cmd.Vnet).Send()

	err := cmd.Render(ctx.Output, ctx.NoColor.Bool())
	if err != nil {
		return errors.Wrap(err, ErrGenVnet)
	}
	return nil
}
