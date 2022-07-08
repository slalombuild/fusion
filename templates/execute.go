package templates

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/pkg/errors"
)

var (
	ErrTemplateParse = "failed to parse text/template during template.Execute"
	ErrTemplateExec  = "failed to execute template during exec()"
)

// Execute renders the template string into a buffer and includes
// some useful helper functions
func Execute(templateString string, data interface{}) (*bytes.Buffer, error) {
	t, err := template.New("template").Funcs(sprig.TxtFuncMap()).Parse(templateString)
	if err != nil {
		return nil, errors.Wrap(err, ErrTemplateParse)
	}

	buffer := bytes.NewBuffer(nil)
	if err := t.Execute(buffer, data); err != nil {
		return nil, errors.Wrap(err, ErrTemplateExec)
	}

	return fmt(buffer)
}

func fmt(b *bytes.Buffer) (*bytes.Buffer, error) {
	fmt := hclwrite.Format(b.Bytes())
	if fmt == nil {
		return nil, errors.New("formatted template returned empty")
	}

	buf := bytes.NewBuffer(fmt)
	return buf, nil
}
