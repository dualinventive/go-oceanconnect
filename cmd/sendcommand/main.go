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
	devID  = flag.String("devid", "", "Device ID to read data from")
	name   = flag.String("name", "", "Device name to read data from")
	txData = flag.String("data", "Hello World", "Data to send")

	cfgFile = flag.String("config", "config.yml", "config-file for the API-settings")
)

func sendCmd(dat *oceanconnect.Device) {
	err := dat.Command([]byte(*txData), 150)
	if err != nil {
		logrus.Fatalf("command error: %v", err)
	}
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

	client, err := oceanconnect.NewClient(c)
	if err != nil {
		logrus.Fatalf("client not created: %v", err)
	}

	if len(*name) == 0 && len(*devID) == 0 {
		logrus.Fatalf("no device name nor device ID is present")
	}

	if len(*name) != 0 && len(*devID) != 0 {
		logrus.Fatalf("not both name and device ID can be present")
	}

	devs, err := client.GetDevices(oceanconnect.GetDevicesStruct{PageNo: 0, PageSize: 100})
	if err != nil {
		logrus.Fatalf("problem while retrieving devices: %v", err)
	}

	deviceFound := false
	for _, dev := range devs {
		if (len(*name) > 0 && *name == dev.DeviceInfo.Name) ||
			(len(*devID) > 0 && *devID == dev.DeviceID) {
			sendCmd(&dev)
			deviceFound = true
		}
	}
	if !deviceFound {
		logrus.Fatalf("device not found")
	}

}
