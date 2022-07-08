package snippets

import (
	"encoding/json"
	"encoding/xml"
	"io"

	"github.com/pkg/errors"
)

// XML renders the live templates to a pretty-printed
// object at the destination io.Writer (w).
func XML(w io.Writer, v interface{}) error {
	enc := xml.NewEncoder(w)
	enc.Indent("", "    ")
	err := enc.Encode(v)
	return errors.Wrap(err, "failed to encode XML")
}

// JSON renders the snippets to a pretty-printed
// object at the destination io.Writer (w).
func JSON(w io.Writer, v interface{}) error {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	err := enc.Encode(v)
	return errors.Wrap(err, "failed to encode JSON")
}
