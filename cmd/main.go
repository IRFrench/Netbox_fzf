package main

import (
	"errors"
	"os"
	"ssh_tool/netbox"
	"ssh_tool/template"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	NoConfig = errors.New("TOKEN environment variable is not set")
)

type config struct {
	netbox_urls string
	template    string
	token       string
}

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
	config, err := checkEnv()
	if err != nil {
		log.Error().Err(err).Msg("could not gather environment")
		return err
	}

	// Read Netbox Urls
	urls, err := os.ReadFile(config.netbox_urls)
	if err != nil {
		log.Error().Err(err).Str("file", config.netbox_urls).Msg("could not read netbox url file")
		return err
	}
	url_list := strings.Split(string(urls), "\n")

	// Send Requests for those URLs
	allConfig := make([]template.NetboxConfigLists, len(url_list))

	netboxClient := netbox.NewClient(config.token)
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
	err = template.BuildConfig(allConfig, config.template)
	if err != nil {
		log.Error().Err(err).Msg("could not create the ssh config file")
		return err
	}

	log.Info().Msg("successfully created config file")
	return nil
}

func checkEnv() (*config, error) {
	token, ok := os.LookupEnv("TOKEN")
	if !ok {
		return nil, NoConfig
	}

	netbox_urls, ok := os.LookupEnv("NETBOX_URLS")
	if !ok {
		netbox_urls = "netbox/netbox.txt"
		log.Warn().Str("environment", "NETBOX_URLS").Str("default", netbox_urls).Msg("unset environment")
	}

	template, ok := os.LookupEnv("TEMPLATE")
	if !ok {
		template = "template/ssh_template"
		log.Warn().Str("environment", "TEMPLATE").Str("default", template).Msg("unset environment")
	}

	return &config{
		token:       token,
		netbox_urls: netbox_urls,
		template:    template,
	}, nil
}
