package cfg

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ParseFlags() (*Config, error) {
	var Debug bool
	var Token string
	var Template string
	var Netbox_urls string
	var Output string

	flag.BoolVar(&Debug, "v", false, "Verbose mode. Adds debug statements to the output")
	flag.StringVar(&Token, "t", "", "A Netbox API Token")
	flag.StringVar(&Template, "s", defaultTemplate, "The path of your ssh configuration template")
	flag.StringVar(&Netbox_urls, "n", defaultNetboxUrls, "The path of your netbox urls txt file")
	flag.StringVar(&Output, "o", defaultOutput, "The path for your output file")

	flag.Parse()

	newConfig := Config{
		Netbox_urls: Netbox_urls,
		Template:    Template,
		Token:       Token,
		Output:      Output,
		Debug:       Debug,
	}

	return &newConfig, checkFlags(&newConfig)
}

func checkFlags(config *Config) error {
	if config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Debug().Msg("running in verbose mode")

	if config.Token == "" {
		token, ok := os.LookupEnv("TOKEN")
		if !ok {
			return ErrNoToken
		}
		config.Token = token
	}
	log.Debug().Str("token", config.Token).Msg("token has been set")

	if config.Netbox_urls == defaultNetboxUrls {
		log.Warn().Str("flag", "n").Str("default", defaultNetboxUrls).Msg("using default netbox urls path")
	}
	log.Debug().Str("netbox_urls", config.Netbox_urls).Msg("tetbox urls have been set")

	if config.Template == defaultTemplate {
		log.Warn().Str("flag", "s").Str("default", defaultTemplate).Msg("using default ssh config template path")
	}
	log.Debug().Str("template", config.Template).Msg("template has been set")

	if config.Output == defaultOutput {
		log.Warn().Str("flag", "o").Str("default", defaultOutput).Msg("using default output file path")
	}
	log.Debug().Str("output", config.Output).Msg("output has been set")

	return nil
}
