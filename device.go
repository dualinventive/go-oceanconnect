// Copyright 2017 The go-oceanconnect authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package oceanconnect

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Device struct with device data
type Device struct {
	DeviceID         string     `json:"deviceId"`
	GatewayID        string     `json:"gatewayId"`
	NodeType         string     `json:"nodeType"`
	CreateTime       OcTime     `json:"creationTime"`
	LastModifiedTime OcTime     `json:"lastModifiedTime"`
	DeviceInfo       DeviceInfo `json:"deviceInfo"`
	Services         []Service  `json:"services"`
	client           *Client
}

// Service struct which holds service information data
type Service struct {
	ServiceID   string `json:"serviceId"`
	ServiceType string `json:"serviceType"`
	Data        []byte `json:"data"`
	EventTime   OcTime `json:"eventTime`
	ServiceInfo string `json:"serviceInfo"`
}

func (u *Service) UnmarshalJSON(data []byte) error {

	type Alias Service

	aux := &struct {
		Data interface{} `json:"data"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	var err error
	u.Data, err = json.Marshal(aux.Data)
	return err
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
	ServiceID string
	Data      []byte `json:"data"`
	Timestamp OcTime
}

func (u *DeviceData) UnmarshalJSON(data []byte) error {
	type Alias DeviceData
	aux := &struct {
		Data interface{} `json:"data"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	var err error
	u.Data, err = json.Marshal(aux.Data)
	return err
}

// GetHistoricalData returns data from specific device
func (d *Device) GetHistoricalData() ([]DeviceData, error) {
	resp, err := d.client.request(http.MethodGet, "/iocm/app/data/v1.1.0/deviceDataHistory?deviceId="+d.DeviceID+"&gatewayId="+d.GatewayID, nil)
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

// Command send command to device
func (d *Device) Command(serviceID string, idata interface{}, timeoutSec int64) error {
	return d.client.SendCommand(d.DeviceID, serviceID, idata, timeoutSec)
}
