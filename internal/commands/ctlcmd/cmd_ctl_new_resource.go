package ctlcmd

import (
	"io"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/slalombuild/fusion/internal/commands"
	"github.com/slalombuild/fusion/internal/generate"
	"github.com/slalombuild/fusion/templates"
)

var (
	ErrPrompt = "failed to ask survey prompt"
)

type NewResourceCmd struct {
	Questions []*survey.Question `kong:"-"`
	Answers   Answers            `embed:""`
	Fields    map[string]string  `help:"Map of fields and their types to be used in template" short:"f" default:"name=string;description=string"`
	Save      bool               `help:"Save output" default:"false"`
}

type Answers struct {
	Resource string `arg:"" help:"Terraform resource name (e.g. loadbalancer, security_group)"`
	Provider string `help:"Cloud provider name (e.g. aws, gcp, azure)" enum:"aws,gcp,azure," default:"" short:"p"`
	Import   string `help:"Terraform file to import" short:"i" type:"existingfile"`
}

func (cmd *NewResourceCmd) Run(ctx *commands.Context) error {
	// Setup questions to skip prompt if answer already provided
	if cmd.Answers.Provider == "" {
		err := survey.AskOne(&survey.Select{
			Message: "Choose a cloud provider:",
			Options: []string{"aws", "gcp", "azure"},
		}, &cmd.Answers.Provider)
		if err != nil {
			return errors.Wrap(err, ErrPrompt)
		}
	}

	if cmd.Answers.Resource == "" {
		err := survey.AskOne(
			&survey.Input{
				Message: "What is your resource name?",
			}, &cmd.Answers.Resource,
		)
		if err != nil {
			return errors.Wrap(err, ErrPrompt)
		}
	}

	if cmd.Answers.Import == "" {
		err := survey.AskOne(&survey.Input{
			Message: "Import a terraform file to templatize",
			Suggest: suggestFiles,
		}, &cmd.Answers.Import)
		if err != nil {
			return errors.Wrap(err, ErrPrompt)
		}
	}

	// Log user action required to wire
	// up CLI
	if cmd.Save {
		log.
			Info().
			Msgf("Add your '%s' into the '%s' struct in 'internal/commands/%scmd/cmd_%s.go' when ready to implement your new resource",
				generate.CommandName(cmd.Answers.Resource),
				strings.ToUpper(cmd.Answers.Provider),
				cmd.Answers.Provider,
				cmd.Answers.Provider,
			)
	}

	err := cmd.generateGoCmd(ctx.Output, ctx.NoColor.Bool())
	if err != nil {
		return err
	}

	err = cmd.generateGoTemplateData(ctx.Output, cmd.Fields, ctx.NoColor.Bool())
	if err != nil {
		return err
	}

	err = cmd.generateTerraformTemplate(ctx.Output, cmd.Fields, ctx.NoColor.Bool())
	if err != nil {
		return err
	}

	return nil
}

// generateGoCmd Generates syntax highlighted and formatted
// Go source code from answers
func (cmd *NewResourceCmd) generateGoCmd(w io.Writer, skipColor bool) error {
	// Prefix source with comment
	outputPath := generate.OutputPath(generate.DESTINATION_COMMAND, cmd.Answers.Provider, cmd.Answers.Resource)

	// Generate code
	source := generate.Command(cmd.Answers.Provider, cmd.Answers.Resource)
	if !cmd.Save {
		log.Info().Str("language", "GO").Str("output", outputPath).Str("category", "cli").Send()

		// Render source
		err := templates.Highlight(w, source, templates.GO, skipColor)
		if err != nil {
			return err
		}

		return nil
	}

	log.Info().Str("output", outputPath).Msg("generating output")
	err := generate.Save(outputPath, []byte(source))
	if err != nil {
		return err
	}

	return nil
}

// generateGoTemplateData Generates syntax highlighted and formatted
// Go source code from answers for rendering template data
func (cmd *NewResourceCmd) generateGoTemplateData(w io.Writer, fields map[string]string, skipColor bool) error {
	// Generate code
	source := generate.TemplateData(fields, cmd.Answers.Provider, cmd.Answers.Resource)

	// Prefix source with comment
	outputPath := generate.OutputPath(generate.DESTINATION_TEMPLATE_DATA, cmd.Answers.Provider, cmd.Answers.Resource)

	if !cmd.Save {
		log.Info().Str("language", "GO").Str("output", outputPath).Str("category", "template").Send()
		// Render source
		err := templates.Highlight(w, source, templates.GO, skipColor)
		if err != nil {
			return err
		}

		return nil
	}

	log.Info().Str("output", outputPath).Msg("generating output")
	err := generate.Save(outputPath, []byte(source))
	if err != nil {
		return err
	}

	return nil
}

// generateTerraformTemplate Generates a syntax highlighted and formatted
// Terraform template from answers
func (cmd *NewResourceCmd) generateTerraformTemplate(w io.Writer, fields map[string]string, skipColor bool) error {
	// Prefix source with comment
	outputPath := generate.OutputPath(generate.DESTINATION_TEMPLATE, cmd.Answers.Provider, cmd.Answers.Resource)

	source := generate.Template(fields, cmd.Answers.Import)

	if !cmd.Save {
		log.Info().Str("language", "HCL").Str("output", outputPath).Str("category", "template").Send()

		// Render source
		err := templates.Highlight(w, string(source), templates.HCL, skipColor)
		if err != nil {
			return err
		}

		return nil
	}

	log.Info().Str("output", outputPath).Msg("generating output")
	err := generate.Save(outputPath, source)
	if err != nil {
		return err
	}

	return nil
}
