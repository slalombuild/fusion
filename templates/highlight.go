package templates

import (
	"go/format"
	"io"

	"github.com/alecthomas/chroma/quick"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/pkg/errors"
)

var (
	ErrNoColor   = "failed to print output with NO_COLOR"
	ErrHighlight = "failed to highlight %s source"
	ErrFmt       = "failed to format %s source"
)

type Language int

const (
	GO Language = iota
	HCL
	PLAINTEXT
)

func (l Language) string() string {
	return [...]string{"go", "hcl", "plaintext"}[l]
}

// Highlight and format text
func Highlight(w io.Writer, text string, lang Language, skipColor bool) error {
	formatter := "terminal"
	if skipColor {
		formatter = "noop"
	}

	switch lang {
	case HCL:
		fmt := hclwrite.Format([]byte(text))
		if fmt == nil {
			return nil
		}
		text = string(fmt)
	case GO:
		fmt, err := format.Source([]byte(text))
		if err != nil {
			return errors.Wrapf(err, ErrFmt, lang.string())
		}
		text = string(fmt)
	}

	err := quick.Highlight(w, text, lang.string(), formatter, "fallback")
	if err != nil {
		return errors.Wrapf(err, ErrHighlight, lang.string())
	}

	return nil
}
