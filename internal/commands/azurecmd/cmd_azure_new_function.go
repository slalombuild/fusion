package azurecmd

import (
	"fmt"

	"github.com/SlalomBuild/fusion/templates"
	"github.com/SlalomBuild/fusion/templates/azure"

	"github.com/SlalomBuild/fusion/internal/commands"
)

type NewFunctionCommand struct {
	Globals
	Data *azure.AzureFunction `embed:""`
}

func (cmd *NewFunctionCommand) Run(ctx *commands.Context) error {
	output, err := templates.Execute(azure.TEMPLATE_AZURE_FUNCTION, &cmd.Data)
	if err != nil {
		return err
	}

	fmt.Println(output.String())

	return nil
}
