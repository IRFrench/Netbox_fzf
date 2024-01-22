package netbox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

type NetboxClient struct {
	client *http.Client
	token  string
}

func (n *NetboxClient) buildRequest(url string) (*http.Request, error) {
	newRequest, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	newRequest.Header.Add("Authorization", fmt.Sprintf("Token %v", n.token))
	return newRequest, nil
}

func (n *NetboxClient) parseResponse(response *http.Response) ([]SshDeviceSettings, error) {
	parsedResponse := NetboxResponse{}

	if err := json.NewDecoder(response.Body).Decode(&parsedResponse); err != nil {
		return nil, err
	}

	sshConfigDevices := []SshDeviceSettings{}

	for index := range parsedResponse.Results {
		if parsedResponse.Results[index].PrimaryIp.Address == "" || parsedResponse.Results[index].Name == "" {
			continue
		}

		sshConfigDevices = append(sshConfigDevices, SshDeviceSettings{
			Name: parsedResponse.Results[index].Name,
			Ip:   strings.Split(parsedResponse.Results[index].PrimaryIp.Address, "/")[0],
		})

		log.Debug().Str("hostname", parsedResponse.Results[index].Name).
			Str("ip", parsedResponse.Results[index].PrimaryIp.Address).
			Str("parsed_ip", strings.Split(parsedResponse.Results[index].PrimaryIp.Address, "/")[0]).
			Msg("parsed config")
	}

	return sshConfigDevices, nil
}

func (n *NetboxClient) RunRequest(url string) ([]SshDeviceSettings, error) {
	log.Debug().Str("url", url).Msg("building request")
	newRequest, err := n.buildRequest(url)
	if err != nil {
		return nil, err
	}

	log.Debug().Msg("sending request")
	response, err := n.client.Do(newRequest)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	log.Debug().Msg("parsing response")
	return n.parseResponse(response)
}

func NewClient(token string) *NetboxClient {
	newClient := http.Client{
		Timeout: 10 * time.Second,
	}

	return &NetboxClient{
		client: &newClient,
		token:  token,
	}
}
