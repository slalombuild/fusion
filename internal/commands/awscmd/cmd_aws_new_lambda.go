package awscmd

import (
	"github.com/SlalomBuild/fusion/internal/commands"
	"github.com/SlalomBuild/fusion/templates/aws"
	"github.com/pkg/errors"

	"github.com/rs/zerolog/log"
)

var (
	ErrGenLambda = "failed to generate terraform for lamba"
	ErrHighlight = "error highlighting terraform"
)

type NewLambdaCommand struct {
	Globals
	Graph               bool `help:"Generate graphviz of terraform resource" default:"false"`
	*aws.LambdaFunction `embed:""`
}

func (cmd *NewLambdaCommand) Run(ctx *commands.Context) error {
	log.Info().Str("provider", "aws").Str("resource", "lambda_function").Interface("data", cmd.LambdaFunction).Send()
	if cmd.Graph {
		return commands.Graph(ctx, cmd)
	}

	err := cmd.Render(ctx.Output, ctx.NoColor.Bool())
	if err != nil {
		return errors.Wrap(err, ErrGenLambda)
	}

	return nil
}
