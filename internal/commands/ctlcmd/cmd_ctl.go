// Package ctlcmd contains all commands for the fusionctl cli
package ctlcmd

import (
	"path/filepath"
)

type Ctl struct {
	Resource NewResourceCmd `cmd:"" help:"Create new fusion resources from existing terraform"`
	Stack    NewStackCmd    `cmd:"" help:"Create new fusion stacks from existing terraform" hidden:"true"`
}

func suggestFiles(toComplete string) []string {
	files, _ := filepath.Glob(toComplete + "*.tf")
	return files
}
