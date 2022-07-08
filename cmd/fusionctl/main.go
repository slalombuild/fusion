package main

import (
	"os"

	"github.com/SlalomBuild/fusion/internal/commands"
	"github.com/SlalomBuild/fusion/internal/commands/ctlcmd"
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

	New ctlcmd.Ctl `cmd:"" help:"Generate new fusion commands and templates"`

	Completions kongplete.InstallCompletions `cmd:"" alias:"complete" help:"generate shell completions"`
}

func Run() (*kong.Context, error) {
	parser := kong.Must(&CLI,
		kong.Name("fusionctl"),
		kong.Description("Generate code and import terraform resources into fusion"),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Bind(&commands.Context{}),
		kong.Vars{
			"version": version,
			"date":    date,
			"commit":  commit,
		},
		kong.ShortUsageOnError(),
	)

	// Generate autocompletions
	kongplete.Complete(parser)
	// Bind global flags to subcommand
	// context
	ctx, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)

	// Bind global flags to subcommand
	// context
	ctx.Bind(&commands.Context{
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
