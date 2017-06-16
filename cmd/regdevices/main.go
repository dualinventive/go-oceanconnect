// Copyright 2017 The go-oceanconnect authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/csv"
	"flag"
	"io/ioutil"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/dualinventive/go-oceanconnect"
	"gopkg.in/yaml.v2"
)

var (
	imei    = flag.String("imei", "", "IMEI number of the device to register")
	csvFile = flag.String("csv", "", "CSV-file to import devices from")
	csvSep  = flag.String("csvsep", ",", "Separator in CSV files")
	devID   = flag.String("devid", "", "Device ID to update device parameters")
	devName = flag.String("name", "", "Device name in OceanConnect (defaults to IMEI)")

	cfgFile = flag.String("config", "config.yml", "config-file for the API-settings")
)

func registerDevice(client *oceanconnect.Client, deviceID, imei, name string) {
	if len(deviceID) == 0 {
		reg, err := client.RegisterDevice(imei, 3600)
		if err != nil {
			logrus.Fatalf("register device failed: %v\n", err)
		}
		deviceID = reg.DeviceID
		logrus.Infof("Registration successful for %s: %s", imei, deviceID)
	}

	if name == "" {
		name = imei
	}
	err := client.SetDeviceInfo(deviceID, name)
	if err != nil {
		logrus.Fatalf("setting device-info failed: %v\n", err)
	}
	logrus.Infof("Set device info for %s succeeded", imei)
}

func main() {
	flag.Parse()
	d, err := ioutil.ReadFile(*cfgFile)
	if err != nil {
		logrus.Fatalf("error reading config-file: %v", err)
	}
	c := oceanconnect.Config{
		CertFile:    "cert.crt",
		CertKeyFile: "key.key",
	}
	if err := yaml.Unmarshal(d, &c); err != nil {
		logrus.Fatalf("reading config-file failed: %v", err)
	}

	if len(*imei) == 0 || len(*csvFile) == 0 {
		logrus.Fatalf("no imei number nor CSV-file provided, use -imei or -csv")
	}

	client, err := oceanconnect.NewClient(c)
	if err != nil {
		logrus.Fatalf("client not created: %v", err)
	}

	if len(*csvFile) != 0 {
		f, err := os.Open(*csvFile)
		if err != nil {
			logrus.Fatalf("error opening CSV file: %v", err)
		}
		defer f.Close()
		c := csv.NewReader(f)
		c.Comma = []rune(*csvSep)[0]
		c.TrimLeadingSpace = true
		for {
			record, err := c.Read()
			if err != nil {
				logrus.Errorf("error reading csv")
				break
			}
			switch len(record) {
			case 1:
				registerDevice(client, "", record[0], "")
			case 2:
				registerDevice(client, "", record[0], record[1])
			default:
				logrus.Warnf("invalid CSV line: %v", record)
			}
		}
	} else {
		registerDevice(client, *devID, *imei, *devName)
	}
}
