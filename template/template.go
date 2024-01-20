package template

import (
	"os"
	"ssh_tool/netbox"
	"text/template"
)

type templateData struct {
	Config [][]netbox.SshDeviceSettings
	Title  string
}

func BuildConfig(sshConfig [][]netbox.SshDeviceSettings, templateFile string) error {
	newTemplate, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	file, err := os.Create("config")
	if err != nil {
		return err
	}
	defer file.Close()

	if err := newTemplate.Execute(file, templateData{
		Config: sshConfig,
		Title:  netboxTitle,
	}); err != nil {
		return err
	}

	return nil
}
