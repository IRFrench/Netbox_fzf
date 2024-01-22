package main

import (
	"os"
	"ssh_tool/cfg"
	"ssh_tool/netbox"
	"ssh_tool/template"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	err := runService()
	if err != nil {
		log.Error().Err(err).Msg("an error occured")
		os.Exit(1)
	}
	os.Exit(0)
}

func runService() error {
	config, err := cfg.ParseFlags()
	if err != nil {
		log.Error().Err(err).Msg("could not gather environment")
		return err
	}

	// Read Netbox Urls
	log.Debug().Msg("reading netbox urls file")
	urls, err := os.ReadFile(config.Netbox_urls)
	if err != nil {
		log.Error().Err(err).Str("file", config.Netbox_urls).Msg("could not read netbox url file")
		return err
	}
	url_list := strings.Split(string(urls), "\n")

	// Send Requests for those URLs
	allConfig := make([]template.NetboxConfigLists, len(url_list))

	log.Debug().Msg("running through netbox urls")
	netboxClient := netbox.NewClient(config.Token)
	for index := range url_list {
		log.Info().Str("url", url_list[index]).Msg("collecting config")
		deviceConfigs, err := netboxClient.RunRequest(url_list[index])
		if err != nil {
			log.Error().Err(err).Str("url", url_list[index]).Msg("could not run netbox request")
			return err
		}
		allConfig[index] = template.NetboxConfigLists{
			Url:    url_list[index],
			Config: deviceConfigs,
		}
	}

	// Write these to the template
	log.Debug().Msg("building configuration template")
	err = template.BuildConfig(allConfig, config.Template, config.Output)
	if err != nil {
		log.Error().Err(err).Msg("could not create the ssh config file")
		return err
	}

	log.Info().Msg("successfully created config file")
	return nil
}
