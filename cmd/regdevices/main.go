// Copyright 2017 The go-oceanconnect authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/dualinventive/go-oceanconnect"
)

var (
	imei    = flag.String("imei", "", "IMEI number of the device to register")
	devID   = flag.String("devid", "", "Device ID to update device parameters")
	devName = flag.String("name", "", "Device name in OceanConnect (defaults to IMEI)")

	cfgFile = flag.String("config", "config.yml", "config-file for the API-settings")
)

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

	client, err := oceanconnect.NewClient(c)
	if err != nil {
		logrus.Fatalf("client not created: %v", err)
	}

	if len(*imei) == 0 {
		logrus.Fatalf("no imei number provided, use -imei")
	}

	deviceID := *devID

	if len(*devID) == 0 {
		reg, err := client.RegisterDevice(*imei, 3600)
		if err != nil {
			logrus.Fatalf("register device failed: %v\n", err)
		}
		deviceID = reg.DeviceID
		logrus.Infof("Registration successful: %s", deviceID)
	}

	name := *devName
	if name == "" {
		name = *imei
	}
	err = client.SetDeviceInfo(deviceID, name)
	if err != nil {
		logrus.Fatalf("setting device-info failed: %v\n", err)
	}
	logrus.Infof("Set device info succeeded")
}
