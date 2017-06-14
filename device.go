package oceanconnect

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// Device struct with device data
type Device struct {
	DeviceID         string
	GatewayID        string
	NodeType         string
	CreateTime       OcTime
	LastModifiedTime OcTime
	DeviceInfo       `json:"deviceInfo"`
	Services         []Services
	client           *Client
}

// DeviceInfo struct with device info data
type DeviceInfo struct {
	NodeID            string
	Name              string
	Description       string
	ManufacturerID    string
	ManufacturerName  string
	Mac               string
	Location          string
	DeviceType        string
	Model             string
	Swversion         string
	FwVersion         string
	HwVersion         string
	ProtocolType      string
	BridgeID          string
	Status            string
	StatusDetail      string
	Mute              string
	SupportedSecurity string
	IsSecurity        string
	SignalStrength    string
	SigVersion        string
	SerialNumber      string
}

// deviceHistory struct with response data
type deviceHistory struct {
	TotalCount int
	PageNo     int
	PageSize   int
	DeviceData []DeviceData `json:"deviceDataHistoryDTOs"`
}

// DeviceData struct with response data
type DeviceData struct {
	DeviceID  string
	GatewayID string
	Appid     string
	ServiceIS string
	Data      `json:"data"`
	Timestamp OcTime
}

// GetHistoricalData returns data from specific device
func (d *Device) GetHistoricalData() ([]DeviceData, error) {
	r, err := http.NewRequest(http.MethodGet, d.client.cfg.URL+"/iocm/app/data/v1.1.0/deviceDataHistory?deviceId="+d.DeviceID+"&gatewayId="+d.GatewayID, nil)
	if err != nil {
		return nil, err
	}

	resp, err := d.client.doRequest(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("invalid response code: " + resp.Status)
	}

	// save device response
	dh := deviceHistory{}
	if err := json.NewDecoder(resp.Body).Decode(&dh); err != nil {
		return nil, err
	}
	return dh.DeviceData, nil
}

func CalculateDeviceUID(imei string) string {
	return "04" + strings.Repeat("0", 30-len(imei)) + imei
}
