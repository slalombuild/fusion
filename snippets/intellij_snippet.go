package snippets

import (
	"embed"
	"encoding/xml"
	"fmt"
	"io"
	"io/fs"

	"github.com/ettle/strcase"
	"github.com/pkg/errors"
)

// The IntelliJTemplateGroup contains all of the Templates that will
// be contained within the Template group
type IntelliJTemplateGroup struct {
	XMLName   xml.Name           `xml:"templateSet"`
	Group     string             `xml:"group,attr"`
	Templates []IntelliJTemplate `xml:"template"`
}

var _ Generator = IntelliJ{}

type IntelliJ struct{}

// Templates contain information about the live template
// including information about any variables it uses,
// the base template itself, and other metadata
type IntelliJTemplate struct {
	// XMLName serves as the XML tag which this struct will populate
	XMLName xml.Name `xml:"template"`

	// Name serves as the abbreviation for a template in the IDE
	Name string `xml:"name,attr"`

	// Value is the template as a string, which is to be encoded to XML at a later time
	Value string `xml:"value,attr"`

	// Description is an optional description of the live template displayed by the IDE
	Description string `xml:"description,attr"`

	// Variable definitions are needed in order to allow their usage within templates
	// - see Variable struct
	Variables []templateVariable `xml:"variable"`

	ToReformat string `xml:"toReformat,attr"`

	ToShortenFQNames string `xml:"toShortenFQNames,attr"`

	// Context contains Options for a Template - see Context struct
	Context templateContext `xml:"context"`
}

// Variables hold info about the variables that are used
// within the live template such as their default value
type templateVariable struct {
	// XMLName serves as the XML tag which this struct will populate
	XMLName xml.Name `xml:"variable"`

	// Name of the variable within the Template Value string
	Name string `xml:"name,attr"`

	// Expression can utilize functions provided by jetbrains
	// see https://www.jetbrains.com/help/idea/template-variables.html#predefined_functions
	// for fusion we default to surrounding the variable name in quotes
	Expression string `xml:"expression,attr"`

	// DefaultValue will be used if the expression fails to evaluate
	DefaultValue string `xml:"defaultValue,attr"`

	AlwaysStopAt string `xml:"alwaysStopAt,attr"`
}

// Options describe which language(s) and in what contexts
// the live template is used
type templateOption struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

// templateContext contains Options for a Template
type templateContext struct {
	Options []templateOption `xml:"option"`
}

// IntelliJOption is a modifying function that applies
// a property to an Template.
type IntelliJOption func(*IntelliJTemplate)

// AddTemplate adds a Tempate to the target TemplateSet struct
func (i IntelliJ) AddTemplate(target *IntelliJTemplateGroup, template IntelliJTemplate) {
	target.Templates = append(target.Templates, template)
}

// NewTemplate creates a new live template Template, the building block for
// IntelliJ live templates.
//
// All Template bodies are generated from parsing the provided Go template.
func (i IntelliJ) NewTemplate(name string, template string, opts ...IntelliJOption) *IntelliJTemplate {
	name = "fsn-" + strcase.ToSnake(name)

	defaultOptions := []templateOption{
		{Name: "HTML", Value: "false"},
		{Name: "JSON", Value: "false"},
		{Name: "OTHER", Value: "true"},
		{Name: "PYTHON", Value: "false"},
		{Name: "SHELL_SCRIPT", Value: "false"},
		{Name: "XML", Value: "false"},
	}
	defaultContext := templateContext{Options: defaultOptions}

	var (
		n                       = &xml.Name{Local: name}
		defaultValue            = ""
		defaultDescription      = "fusion live template"
		defaultToReformat       = "false"
		defaultToShortenFQNames = "true"
	)

	v := &IntelliJTemplate{
		XMLName:          *n,
		Name:             name,
		Value:            defaultValue,
		Description:      defaultDescription,
		ToReformat:       defaultToReformat,
		ToShortenFQNames: defaultToShortenFQNames,
		Context:          defaultContext,
	}

	// variables are returned from the parser so they
	// can be included in the Variables struct to be encoded
	vars := v.applyGoTemplate(template)

	xmlVars := []templateVariable{}
	for i, value := range vars {
		varStruct := templateVariable{
			Name:         fmt.Sprintf("%s_%v", value, i+1),
			Expression:   fmt.Sprintf(`"%s"`, value),
			DefaultValue: value,
			AlwaysStopAt: "true",
		}
		xmlVars = append(xmlVars, varStruct)
	}
	v.Variables = xmlVars

	for _, opt := range opts {
		opt(v)
	}

	return v
}

func (i IntelliJ) WithVariables(vars []string) IntelliJOption {
	return func(s *IntelliJTemplate) {
		xmlVars := []templateVariable{}
		for i, v := range vars {
			varStruct := templateVariable{
				Name:         fmt.Sprintf("%s_%v", v, i),
				Expression:   "",
				DefaultValue: "",
				AlwaysStopAt: "true",
			}
			xmlVars = append(xmlVars, varStruct)
		}
		s.Variables = xmlVars
	}
}

// WithDescription applies a description to a Template.
func (i IntelliJ) WithDescription(description string) IntelliJOption {
	return func(s *IntelliJTemplate) {
		s.Description = description
	}
}

// WithValue applies the live template value lines to a Template.
func (i IntelliJ) WithValue(value string) IntelliJOption {
	return func(s *IntelliJTemplate) {
		s.Value = value
	}
}

// applyGoTemplate parses the Go template string
// and converts it into the intellij new-line delimited
// value format.
func (i *IntelliJTemplate) applyGoTemplate(template string) []string {
	ij := IntelliJ{}
	vars := []string{}
	if template == "" {
		i.Value = ""
		return vars
	}
	tpl, varCount, vars := ij.replaceBlockNames(template, 1, vars)
	tpl, _, vars = ij.replaceDirectInsertions(tpl, varCount, vars)
	i.Value = tpl
	return vars
}

// Generate generates an XML file which contains
// an IntelliJ live template group built from
// the fusion templates
//
// Generate supports rendering output to an io.Writer.
func (i IntelliJ) Generate(w io.Writer, filesystem embed.FS) error {
	s := IntelliJTemplateGroup{Group: "fusion"}

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

		template := i.NewTemplate(name, string(f))
		i.AddTemplate(&s, *template)

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to iterate over templates")
	}

	// attempt at creating zip file for use with
	// intellij editors' import settings feature

	// buf := new(bytes.Buffer)
	// zipWriter := zip.NewWriter(buf)
	// zipFile, err := zipWriter.Create("fusion.xml")
	// if err != nil{
	// 	return errors.Wrap(err, "failed to create zip archive")
	// }
	// xmlBytes, err := xml.MarshalIndent(s, "", "  ")
	// _, err = zipFile.Write(xmlBytes)
	// if err != nil{
	// 	return errors.Wrap(err, "failed to write to fusion.xml")
	// }
	// err = zipWriter.Close()
	// if err != nil{
	// 	return errors.Wrap(err, "failed closing zipWriter")
	// }
	// err = ioutil.WriteFile("fusion_templates.zip", buf.Bytes(), 0740)
	// if err != nil{
	// 	return errors.Wrap(err, "failed to write to zip archive")
	// }
	// return err
	return XML(w, s)
}
