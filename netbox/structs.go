package netbox

// Device settings that we care about, so not alot

type SshDeviceSettings struct {
	Name string
	Ip   string
}

// Netbox responses

type NetboxResponse struct {
	Count   int              `json:"count"`
	Results []DeviceResponse `json:"results"`
}

type DeviceResponse struct {
	Name      string     `json:"name"`
	PrimaryIp IPResponse `json:"primary_ip"`
}

type IPResponse struct {
	Address string `json:"address"`
}
