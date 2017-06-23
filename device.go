// Copyright 2017 The go-oceanconnect authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package oceanconnect

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
)

// Device struct with device data
type Device struct {
	DeviceID         string
	GatewayID        string
	NodeType         string
	CreateTime       OcTime
	LastModifiedTime OcTime
	DeviceInfo       `json:"deviceInfo"`
	Services         []Service
	client           *Client
}

// Service struct which holds service information data
type Service struct {
	ServiceID   string
	ServiceType string
	Data        `json:"data"`
	EventTime   OcTime
	ServiceInfo string `json:",omitEmpty"`
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

func (d *Device) Command(data []byte, timeoutSec int64) error {
	type devCmdBodyRawData struct {
		RawData string `json:"rawData"`
	}
	type devCmdBodyCommand struct {
		ServiceID string            `json:"serviceId"`
		Method    string            `json:"method"`
		Params    devCmdBodyRawData `json:"paras"`
	}
	type devCmdBody struct {
		//RequestID   string            `json:"requestId"`
		Command devCmdBodyCommand `json:"command"`
		//CallbackURL string            `json:"callbackUrl"`
		ExpireTime int64 `json:"expireTime"`
	}

	cmd := devCmdBody{
		//RequestID: "1234567890",
		Command: devCmdBodyCommand{
			ServiceID: "RawData",
			Method:    "RawData",
			Params: devCmdBodyRawData{
				RawData: base64.StdEncoding.EncodeToString(data),
			},
		},
		//CallbackURL: "https://www.google.com/",
		ExpireTime: timeoutSec,
	}
	body, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	r, err := http.NewRequest(http.MethodPost, d.client.cfg.URL+"/iocm/app/cmd/v1.2.0/devices/"+d.DeviceID+"/commands", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	resp, err := d.client.doRequest(r)
	if err != nil {
		return err
	}

	rs, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("==== response =====\n%s\n", string(rs))

	if resp.StatusCode != http.StatusOK {
		return errors.New("invalid response code: " + resp.Status)
	}

	// save device response
	/*	dh := deviceHistory{}
		if err := json.NewDecoder(resp.Body).Decode(&dh); err != nil {
			return nil, err
		}
		return dh.DeviceData, nil*/
	return nil
}
