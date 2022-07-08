package commands

import (
	"fmt"
	"io"

	"github.com/alecthomas/kong"
)

type Context struct {
	Version  VersionFlag
	Verbose  VerboseFlag
	NoColor  NoColorFlag
	NoFormat NoFormatFlag
	Output   OutputFlag
}

type OutputFlag io.Writer

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Printf("Version: %q\n", vars["version"])
	fmt.Printf("Commit: %q\n", vars["commit"])
	fmt.Printf("Date: %q\n", vars["date"])
	app.Exit(0)
	return nil
}

// VerboseFlag is a flag with a hook that, if triggered,
// will set the debug loggers output to stdout.
type VerboseFlag bool

func (v VerboseFlag) Bool() bool {
	return bool(v)
}

// NoColorFlag supports the no-color environment
// variable standard to skip colorful output
// https://no-color.org
type NoColorFlag bool

func (n NoColorFlag) Bool() bool {
	return bool(n)
}

// NoFormatFlag supports skipping
// the formatting
type NoFormatFlag bool

func (n NoFormatFlag) Bool() bool {
	return bool(n)
}
