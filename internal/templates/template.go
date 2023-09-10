package templates

import (
	"github.com/iancoleman/strcase"
	"io"
	"text/template"
)

type Language int8

const (
	C Language = iota
	Haskell
	Rust
	Go
	AssemblyScript
	CSharp
	Zig
	JavaScript
	Python
)

type stepCases struct {
	Snake  string
	Pascal string
	Camel  string
}

func Template(w io.Writer, lang Language, step *string) error {
	s := stepCases{
		Snake:  strcase.ToSnake(*step),
		Pascal: strcase.ToCamel(*step),
		Camel:  strcase.ToLowerCamel(*step),
	}

	var tmpl *template.Template

	switch lang {
	case C:
		tmpl = template.Must(template.ParseGlob(CTemplate))
	case Haskell:
		tmpl = template.Must(template.ParseGlob(HaskellTemplate))
	case Rust:
		tmpl = template.Must(template.ParseGlob(RustTemplate))
	case Go:
		tmpl = template.Must(template.ParseGlob(GoTemplate))
	case AssemblyScript:
		tmpl = template.Must(template.ParseGlob(AssemblyScriptTemplate))
	case CSharp:
		tmpl = template.Must(template.ParseGlob(CSharpTemplate))
	case Zig:
		tmpl = template.Must(template.ParseGlob(ZigTemplate))
	case JavaScript:
		tmpl = template.Must(template.ParseGlob(JsTemplate))
	case Python:
		tmpl = template.Must(template.ParseGlob(PythonTemplate))
	}

	return tmpl.Execute(w, s)
}
