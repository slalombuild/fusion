package snippets

import (
	"embed"
	"io"
)

type Generator interface {
	Generate(io.Writer, embed.FS) error
}
