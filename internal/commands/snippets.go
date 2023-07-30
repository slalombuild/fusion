package commands

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	snippets "github.com/slalombuild/fusion/snippets"
	"github.com/slalombuild/fusion/templates"
)

const (
	EDITOR_VSCODE   = "vscode"
	EDITOR_INTELLIJ = "intellij"
)

type GenerateSnippetsCommand struct {
	Output  string `short:"o" xor:"output,install" required:"" type:"path" help:"Accepts '' for stdout or a filepath (Default stdout)"`
	Install bool   `short:"i" xor:"output,install" help:"Use this flag to install or update your fusion snippets in VS Code (intellij not supported)"`
	Editor  string `short:"e" required:"" enum:"vscode,intellij" help:"Specify which editor to generate snippets for. Current options are 'vscode' or 'intellij'"`
}

func (g *GenerateSnippetsCommand) Run() error {
	if g.Install {
		return installSnippets(g)
	}
	return writeSnippets(g)
}

func (g *GenerateSnippetsCommand) Validate() error {
	if !strings.HasSuffix(g.Output, ".json") && g.Output != "" && g.Editor == EDITOR_VSCODE {
		return errors.New("file must be of type '.json'")
	}
	if !strings.HasSuffix(g.Output, ".xml") && g.Output != "" && g.Editor == EDITOR_INTELLIJ {
		return errors.New("file must be of type '.xml'")
	}
	return nil
}

// writeSnippets either writes snippets to a given file
// or to stdout if given ” for --output
func writeSnippets(g *GenerateSnippetsCommand) error {
	var err error

	w := os.Stdout
	if g.Output != "" {
		w, err = os.Create(g.Output)
		if err != nil {
			return errors.Wrapf(err, "failed to create file %q", g.Output)
		}
	}

	switch g.Editor {
	case EDITOR_VSCODE:
		return snippets.VSCode{}.Generate(w, templates.ALL_TEMPLATES)
	case EDITOR_INTELLIJ:
		return snippets.IntelliJ{}.Generate(w, templates.ALL_TEMPLATES)
	default:
		return errors.Wrap(errors.New("unrecognized editor"), "failed to writeSnippets to unrecognized editor, check --editor enum or writeSnippets")
	}
}

func installSnippets(g *GenerateSnippetsCommand) error {
	if g.Editor == EDITOR_INTELLIJ {
		return errors.New("installation for IntelliJ not currently supported due to variance in installation location please use --output and place the .xml file in the IDE's templates folder")
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "could not find user's home directory")
	}

	extensionDir := filepath.Join(homeDir, ".vscode", "extensions", "fusion-snippets")
	err = os.Mkdir(extensionDir, 0740)
	if err != nil && !os.IsExist(err) {
		return errors.Wrapf(err, "failed to create extension folder %s, could not find %s/.vscode", extensionDir, homeDir)
	}

	packagePath := filepath.Join(extensionDir, "package.json")
	err = os.WriteFile(packagePath, snippets.EXTENSION_PACKAGE_FILE, 0740)
	if err != nil {
		return errors.Wrapf(err, "failed to write package.json to %s", extensionDir)
	}

	g.Output = filepath.Join(extensionDir, "snippets.json")
	err = writeSnippets(g)
	if err != nil {
		return errors.Wrap(err, "failed to write snippets when attempting to --install")
	}
	log.Info().Msg("Fusion snippets sucessfully installed to " + extensionDir)
	return nil
}
