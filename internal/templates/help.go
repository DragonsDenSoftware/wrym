package templates

import (
	"github.com/syke99/wyvrn-cli/internal/constants"
	"github.com/syke99/wyvrn-cli/internal/pkg/models"
	"html/template"
	"os"
)

var newCommand = models.Command{
	Name:  constants.NewName,
	Usage: constants.HelpUsage,
	Flags: []models.Flag{
		{
			Name:  constants.HelpName,
			Alias: constants.HelpAlias,
			Usage: constants.HelpUsage,
		},
		{
			Name:  constants.ModuleName,
			Alias: constants.ModuleAlias,
			Usage: constants.ModuleUsage,
		},
		{
			Name:  constants.LanguageName,
			Alias: constants.LanguageAlias,
			Usage: constants.LanguageUsage,
		},
		{
			Name:  constants.DirectoryName,
			Alias: constants.DirectoryAlias,
			Usage: constants.DirectoryUsage,
		},
	},
}

var runCommand = models.Command{
	Name:  constants.RunName,
	Usage: constants.RunUsage,
	Flags: []models.Flag{
		{
			Name:  constants.HelpName,
			Alias: constants.HelpAlias,
			Usage: constants.HelpUsage,
		},
		{
			Name:  constants.DevName,
			Alias: constants.DevAlias,
			Usage: constants.DevUsage,
		},
		{
			Name:  constants.StagingName,
			Alias: constants.StagingAlias,
			Usage: constants.StagingUsage,
		},
		{
			Name:  constants.ProdName,
			Alias: constants.ProdAlias,
			Usage: constants.ProdUsage,
		},
	},
}

type Commands int8

const (
	TopLevel Commands = iota
	New
	Run
)

var helpTemplate = `{{ range . }}
Name:  {{ .Name }}

Usage:   {{ .Usage }}

Flags: {{ range .Flags }}
		-{{ .Alias }}|--{{ .Name }} : {{ .Usage }}
       {{ end }}
{{ end }}`

func Help(cmd Commands) error {
	tmpl := template.Must(template.New("help").Parse(helpTemplate))

	command := make([]models.Command, 0, 2)

	switch cmd {
	case New:
		command[0] = newCommand
	case Run:
		command[1] = runCommand
	case TopLevel:
		command[0] = newCommand
		command[1] = runCommand
	}

	return tmpl.Execute(os.Stdout, command)
}
