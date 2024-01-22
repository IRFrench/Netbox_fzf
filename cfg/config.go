package cfg

import "errors"

const (
	defaultNetboxUrls = "netbox/netbox.txt"
	defaultTemplate   = "template/ssh_template"
	defaultOutput     = "config"
)

var (
	ErrNoToken = errors.New("No Token has been provided to the service. Please use the -t flag or the TOKEN environment")
)

type Config struct {
	Netbox_urls string
	Template    string
	Token       string
	Output      string
	Debug       bool
}
