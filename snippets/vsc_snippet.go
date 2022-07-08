// Package snippets implements support for
// building snippets for vscode
package snippets

import (
	"embed"
	"io"
	"io/fs"
	"strings"

	"github.com/ettle/strcase"
	"github.com/pkg/errors"
)

var _ Generator = VSCode{}

type VSCode struct{}

//go:embed package.json
var EXTENSION_PACKAGE_FILE []byte

// Item is the content of a named VSCode snippet.
// Items are placed into maps of map[string]*VSCodeSnippet
// to create a full named VSCodeSnippet.
type VSCodeSnippet struct {
	// name is an internal field used to match
	// Snippet item with its index in a Snippet Map
	name string

	// Scope scopes the snippet so that only relevant snippets are suggested.
	Scope string `json:"scope,omitempty"`

	// Prefix defines one or more trigger words that
	// display the snippet in IntelliSense. Substring
	// matching is performed on prefixes.
	Prefix string `json:"prefix"`

	// Description is an optional description of the snippet
	// displayed by IntelliSense.
	Description string `json:"description"`

	// Body is one or more lines of content, which will be joined as multiple lines upon insertion.
	// Newlines and embedded tabs will be formatted according to the context in which the snippet is inserted.
	Body []string `json:"body"`
}

// VSCodeOption is a modifying function that applies
// a property to an Item.
type VSCodeOption func(*VSCodeSnippet)

// AddItem adds an Item to the target Snippet map.
func (v VSCode) AddItem(target *map[string]*VSCodeSnippet, snippet *VSCodeSnippet) {
	m := make(map[string]*VSCodeSnippet)
	for key, val := range *target {
		m[key] = val
	}

	m[snippet.name] = snippet
	*target = m
}

// NewSnippetMap creates a new VSCode snippet map.
//
// SnippetMaps are used to build a vscode snippet file
// using a map of named Items.
func (v VSCode) NewSnippetMap() map[string]*VSCodeSnippet {
	m := make(map[string]*VSCodeSnippet)
	return m
}

// NewItem creates a new snippet Item, the building block for
// VSCode snippets.
//
// All Item bodies are generated from parsing the provided Go template.
func (v VSCode) NewItem(name string, template string, opts ...VSCodeOption) *VSCodeSnippet {
	name = strcase.ToSnake(name)

	var (
		defaultPrefix      = "fsn-" + name
		defaultScope       = "global"
		defaultDescription = "fusion snippet"
	)

	snippet := &VSCodeSnippet{
		name: name,

		Scope:       defaultScope,
		Prefix:      defaultPrefix,
		Description: defaultDescription,
		Body:        []string{""},
	}

	snippet.applyGoTemplate(template)

	for _, opt := range opts {
		opt(snippet)
	}

	return snippet
}

// WithPrefix applies an intellisense prefix to an Item.
func (v VSCode) WithPrefix(prefix string) VSCodeOption {
	return func(s *VSCodeSnippet) {
		s.Prefix = prefix
	}
}

// WithDescription applies a description to an Item.
func (v VSCode) WithDescription(description string) VSCodeOption {
	return func(s *VSCodeSnippet) {
		s.Description = description
	}
}

// WithBody applies the snippet body lines to an Item.
func (v VSCode) WithBody(body []string) VSCodeOption {
	return func(s *VSCodeSnippet) {
		s.Body = body
	}
}

// applyGoTemplate parses the Go template string
// and converts it into the vscode new-line delimited
// body format.
func (i *VSCodeSnippet) applyGoTemplate(template string) {
	v := VSCode{}
	if template == "" {
		i.Body = []string{""}
		return
	}
	tpl, varCount := v.replaceBlockNames(template, 1)
	tpl, _ = v.replaceDirectInsertions(tpl, varCount)
	lines := strings.Split(tpl, "\n")

	var bodyLines []string
	for i, v := range lines {
		if i != len(lines)-1 {
			bodyLines = append(bodyLines, strings.TrimRight(v, "\r\n"))
		} else {
			bodyLines = append(bodyLines, v)
		}
	}

	i.Body = bodyLines
}

// Generate generates a snippets file from all
// Go template files in the provided filesystem.
//
// Generate supports rendering output to an io.Writer.
func (v VSCode) Generate(w io.Writer, filesystem embed.FS) error {
	s := v.NewSnippetMap()

	err := fs.WalkDir(filesystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return errors.Wrap(err, "walkdir was passed an error")
		}
		name, ok := nameFromPath(path)
		if !ok {
			return nil
		}

		f, err := filesystem.ReadFile(path)
		if err != nil {
			return errors.Wrapf(err, "failed to read file %q", path)
		}

		snippet := v.NewItem(name, string(f))
		v.AddItem(&s, snippet)

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to iterate over templates")
	}

	return JSON(w, s)
}
