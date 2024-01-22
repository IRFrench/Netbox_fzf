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
}

func BuildConfig(sshConfig []NetboxConfigLists, templateFile string, outputPath string) error {
	newTemplate, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := newTemplate.Execute(file, templateData{
		NetboxConfigs: sshConfig,
	}); err != nil {
		return err
	}

	return nil
}
