<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# templates

```go
import "github.com/slalombuild/fusion/templates"
```

Package templates includes go text/templates for terraform resources

## Index

- [Variables](<#variables>)
- [func Execute(templateString string, data interface{}) (*bytes.Buffer, error)](<#func-execute>)
- [func FormatHCL(template string) string](<#func-formathcl>)
- [func Highlight(w io.Writer, text string, lang Language, skipColor bool) error](<#func-highlight>)
- [type Language](<#type-language>)
- [type Renderer](<#type-renderer>)


## Variables

```go
var (
    ErrTemplateParse = "failed to parse text/template during template.Execute"
    ErrTemplateExec  = "failed to execute template during exec()"
)
```

```go
var (
    ErrNoColor   = "failed to print output with NO_COLOR"
    ErrHighlight = "failed to highlight %s source"
    ErrFmt       = "failed to format %s source"
)
```

go:embed aws/\*\.tmpl gcp/\*\.tmpl azure/\*\.tmpl

```go
var ALL_TEMPLATES embed.FS
```

## func [Execute](<https://github.com/slalombuild/fusion/blob/main/templates/execute.go#L19>)

```go
func Execute(templateString string, data interface{}) (*bytes.Buffer, error)
```

Execute renders the template string into a buffer and includes some useful helper functions

## func [FormatHCL](<https://github.com/slalombuild/fusion/blob/main/templates/templates.go#L24>)

```go
func FormatHCL(template string) string
```

FormatHCL formats an HCL string and trims whitespace

## func [Highlight](<https://github.com/slalombuild/fusion/blob/main/templates/highlight.go#L31>)

```go
func Highlight(w io.Writer, text string, lang Language, skipColor bool) error
```

Highlight and format text

## type [Language](<https://github.com/slalombuild/fusion/blob/main/templates/highlight.go#L18>)

```go
type Language int
```

```go
const (
    GO  Language = iota
    HCL
    PLAINTEXT
)
```

## type [Renderer](<https://github.com/slalombuild/fusion/blob/main/templates/templates.go#L18-L20>)

Renderer is the required interface for terraform templates to be rendered

```go
type Renderer interface {
    Render(w io.Writer, skipColor bool) error
}
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
