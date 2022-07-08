package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/dave/jennifer/jen"
	"github.com/ettle/strcase"
)

const (
	MODULE_FUSION    = "github.com/SlalomBuild/fusion/"
	MODULE_ZEROLOG   = "github.com/rs/zerolog/log"
	MODULE_ERRORS    = "github.com/pkg/errors"
	MODULE_TEMPLATES = MODULE_FUSION + "templates"
	MODULE_COMMANDS  = MODULE_FUSION + "internal/commands"
)

// Command generates source code for a new
// resource command
func Command(provider, resource string) string {
	f := jen.NewFile(strings.ToLower(provider + "cmd"))
	fnName := CommandName(resource)

	// Create command struct
	f.Commentf("%s creates a new %s", fnName, resource)
	f.Type().Id(fnName).Struct(
		jen.Id("Globals"),
		jen.Op("*").Qual(MODULE_TEMPLATES+"/"+provider, strcase.ToGoPascal(resource)).Tag(map[string]string{"embed": ""}),
	)

	// Create run function
	f.Add(CommandRunFunc(resource))

	return f.GoString()
}

func CommandRunFunc(resource string) jen.Code {
	// Build run function with all
	// components
	return jen.Func().Params(
		jen.Id("cmd").Op("*").Id(CommandName(resource))).Id("Run").Params(jen.Id("ctx").Op("*").Qual(MODULE_COMMANDS, "Context")).Params(jen.Id("error")).
		Block(
			jen.Return(jen.Id("cmd").Dot("Render").Call(jen.Id("ctx").Dot("Output"), jen.Id("ctx").Dot("NoColor").Dot("Bool").Call())),
		)
}

// parseJenniferTypeFromString finds the relevant jen.Statement
// type from the provided string argument
func parseJenniferTypeFromString(s string) *jen.Statement {
	switch strings.ToLower(s) {
	case "string":
		return jen.String()
	case "number", "int":
		return jen.Int()
	case "list", "list:string", "array", "array:string", "slice", "slice:string":
		return jen.Index().String()
	case "list:int", "array:int", "slice:int":
		return jen.Index().Int()
	case "object", "object:string", "map", "map:string":
		return jen.Map(jen.String()).String()
	case "object:int", "map:int":
		return jen.Map(jen.String()).Int()
	case "bool":
		return jen.Bool()
	default:
		return jen.String()
	}
}

// TemplateData generates source code for rendering template data
func TemplateData(fields map[string]string, provider, resource string) string {
	f := jen.NewFile(strings.ToLower(provider))
	f.Anon("embed")

	// Go embed comment
	file := fmt.Sprintf("%s_%s", provider, resource)
	f.Commentf("//go:embed %s.tmpl", file)

	// Go embed
	embedVar := strcase.ToSNAKE("TEMPLATE_" + file)
	f.Var().Id(embedVar).String()

	// Generate template arguments
	var args []jen.Code
	for key, val := range fields {
		fieldType := parseJenniferTypeFromString(val).Tag(map[string]string{
			"help": "",
		}).Comment("Update default value and help message")
		args = append(args, jen.Id(strcase.ToGoPascal(key)).Add(fieldType))
	}

	f.Comment("View the Kong CLI docs to see options for your command's flags and arguments\nhttps://github.com/alecthomas/kong#flags")
	f.Type().Id(strcase.ToGoPascal(resource)).Struct(args...)

	f.Commentf("Render generates the Terraform code for the %s", strcase.ToGoPascal(resource))
	f.Func().Params(jen.Id("resource").Op("*").Id(strcase.ToGoPascal(resource))).Id("Render").Params(
		jen.Id("w").Qual("io", "Writer"),
		jen.Id("skipColor").Bool(),
	).Error().Block(
		jen.List(
			jen.Id("output"),
			jen.Id("err"),
		).Op(":=").Qual(MODULE_TEMPLATES, "Execute").Call(jen.Id(embedVar), jen.Op("&").Id("resource")),
		jen.If(
			jen.Err().Op("!=").Nil(),
		).Block(
			jen.Return(
				jen.Qual("github.com/pkg/errors", "Wrap").Call(jen.Id("err"), jen.Lit("failed to generate template")),
			),
		),
		jen.Id("err").Op("=").Qual(MODULE_TEMPLATES, "Highlight").Call(
			jen.Id("w"),
			jen.Id("output").Dot("String").Call(),
			jen.Qual(MODULE_TEMPLATES, "HCL"),
			jen.Id("skipColor"),
		),
		jen.If(
			jen.Err().Op("!=").Nil(),
		).Block(
			jen.Return(
				jen.Qual("github.com/pkg/errors", "Wrap").Call(jen.Id("err"), jen.Lit("failed to highlight source")),
			),
		),
		jen.Return(jen.Nil()),
	)

	return f.GoString()
}

func Template(fields map[string]string, sourceFile string) []byte {
	// Read the imported terraform file
	source, err := os.ReadFile(sourceFile)
	if err != nil {
		log.Fatal().Err(err).Msg(ErrReadFile)
	}

	// Prepend template with template variables
	var result string
	result += "# Available Template Data - for development only\n"
	for key := range fields {
		result += fmt.Sprintf("# %-4s = {{ .%s }}\n", strcase.ToGoPascal(key), strcase.ToGoPascal(key))
	}

	// Append template with source
	result += fmt.Sprintf("\n\n%s\n", string(source))

	return []byte(result)
}
