package main

import (
	"os"

	"github.com/SlalomBuild/fusion/internal/commands"
	"github.com/SlalomBuild/fusion/internal/commands/awscmd"
	"github.com/SlalomBuild/fusion/internal/commands/azurecmd"
	"github.com/SlalomBuild/fusion/internal/commands/gcpcmd"
	"github.com/SlalomBuild/fusion/internal/resolver"
	"github.com/alecthomas/kong"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/willabides/kongplete"
)

var (
	ErrRun = "failed during fusion run"
)

var (
	date    string
	commit  string
	version string
)

var CLI struct {
	Version  commands.VersionFlag `help:"Show version information" name:"version"`
	Verbose  bool                 `help:"Enable verbose logging" short:"v" default:"false"`
	NoColor  bool                 `help:"Disable colorful output" short:"n" env:"NO_COLOR" default:"false"`
	NoFormat bool                 `help:"Disable code formatting" env:"NO_FORMAT" default:"false"`
	Config   kong.ConfigFlag      `help:"Provide a JSON file which will populate flags and their values"`

	New struct {
		AWS   awscmd.AWS     `cmd:"" help:"Create AWS resources with Terraform"`
		GCP   gcpcmd.GCP     `cmd:"" help:"Create GCP resources with Terraform"`
		AZURE azurecmd.Azure `cmd:"" help:"Create Azure resources with Terraform"`
	} `cmd:"" help:"Create new cloud resources with Terraform"`

	Gen struct {
		Snippets    commands.GenerateSnippetsCommand `cmd:"" help:"Generate terraform vscode snippets"`
		Completions kongplete.InstallCompletions     `cmd:"" help:"Generate shell completions"`
	} `cmd:"" help:"Generate snippets and shell completions"`
}

func Run() (*kong.Context, error) {
	parser := kong.Must(&CLI,
		kong.Name("fusion"),
		kong.Description("Generate secure by default cloud infrastructure configuration"),
		kong.ConfigureHelp(kong.HelpOptions{
			NoExpandSubcommands: true,
			Compact:             true,
		}),
		kong.Bind(&commands.Context{}),
		kong.Vars{
			"version": version,
			"date":    date,
			"commit":  commit,
		},
		kong.ShortUsageOnError(),
		kong.NamedMapper("jsonfile", resolver.JSONFileMapper),
		kong.Configuration(resolver.JSON),
	)

	// Enable auto shell completions
	kongplete.Complete(parser)

	// Bind global flags to subcommand
	// context
	ctx, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)

	// Bind global flags to subcommand
	// context
	ctx.Bind(&commands.Context{
		Version:  commands.VersionFlag(ctx.Model.Vars()["version"]),
		Verbose:  commands.VerboseFlag(CLI.Verbose),
		NoColor:  commands.NoColorFlag(CLI.NoColor),
		NoFormat: commands.NoFormatFlag(CLI.NoFormat),
		Output:   commands.OutputFlag(os.Stdout),
	})

	// Enable logger if --verbose
	// Disable colorful output if --no-color
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: CLI.NoColor})
	if !CLI.Verbose {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}

	err = ctx.Run()
	return ctx, errors.Wrap(err, ErrRun)
}

func main() {
	ctx, err := Run()
	ctx.FatalIfErrorf(err)
}
