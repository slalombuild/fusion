package ctlcmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/slalombuild/fusion/internal/commands"
)

type NewStackCmd struct {
	Questions []*survey.Question `kong:"-"`
}

func (cmd *NewStackCmd) Run(ctx *commands.Context) error {

	return nil
}
