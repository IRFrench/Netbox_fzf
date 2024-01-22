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

func (n *NetboxClient) passResponse(response *http.Response) ([]SshDeviceSettings, error) {
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
	newRequest, err := n.buildRequest(url)
	if err != nil {
		return nil, err
	}

	response, err := n.client.Do(newRequest)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return n.passResponse(response)
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
