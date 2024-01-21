package template

import (
	"os"
	"ssh_tool/netbox"
	"text/template"
)

type NetboxConfigLists struct {
	Url    string
	Config []netbox.SshDeviceSettings
}

type templateData struct {
	NetboxConfigs []NetboxConfigLists
	Title         string
}

func BuildConfig(sshConfig []NetboxConfigLists, templateFile string) error {
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
		NetboxConfigs: sshConfig,
		Title:         netboxTitle,
	}); err != nil {
		return err
	}

	return nil
}
