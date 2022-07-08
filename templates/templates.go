// Package templates includes go text/templates for terraform resources
package templates

import (
	"bytes"
	"embed"
	"io"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

//go:embed aws/*.tmpl gcp/*.tmpl azure/*.tmpl
var ALL_TEMPLATES embed.FS

// Renderer is the required
// interface for terraform templates
// to be rendered
type Renderer interface {
	Render(w io.Writer, skipColor bool) error
}

// FormatHCL formats an HCL string and trims
// whitespace
func FormatHCL(template string) string {
	got := hclwrite.Format(bytes.TrimSpace([]byte(template)))
	return string(got)
}
