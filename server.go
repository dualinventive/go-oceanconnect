// Copyright 2017 The go-oceanconnect authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package oceanconnect

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/Sirupsen/logrus"
)

type Server struct {
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Got request:")
	rd, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(rd))

	dec := json.NewDecoder(r.Body)
	var n struct {
		NotifyType string `json:"notifyType"`
	}

	if err := dec.Decode(&n); err != nil {
		logrus.Errorf("error decoding notification type")
	}

	switch Notification(n.NotifyType) {
	case NotificationDeviceAdded:
		logrus.Infof("Notification received: %d", NotificationDeviceAdded)
	case NotificationDeviceInfoChanged:
		logrus.Infof("Notification received: %d", NotificationDeviceInfoChanged)
	case NotificationDeviceDataChanged:
		logrus.Infof("Notification received: %d", NotificationDeviceDataChanged)
	case NotificationDeviceDeleted:
		logrus.Infof("Notification received: %d", NotificationDeviceDeleted)
	case NotificationMessageConfirm:
		logrus.Infof("Notification received: %d", NotificationMessageConfirm)
	case NotificationCommandResponse:
		logrus.Infof("Notification received: %d", NotificationCommandResponse)
	case NotificationDeviceEvent:
		logrus.Infof("Notification received: %d", NotificationDeviceEvent)
	case NotificationServiceInfoChanged:
		logrus.Infof("Notification received: %d", NotificationServiceInfoChanged)
	case NotificationRuleEvent:
		logrus.Infof("Notification received: %d", NotificationRuleEvent)
	}

	/*	contents, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logrus.Errorf("err: %v", err)
			return
		}

		var data DeviceDataChanged
		if err := json.Unmarshal(contents, &data); err != nil {
			logrus.Errorf("err: %v", err)
			return
		}
		fmt.Printf("data: %+v\n", data)*/
}

func (s *Server) ListenAndServe(uri string) error {
	http.HandleFunc("/", s.handler)
	return http.ListenAndServe(uri, nil)
}

// NewServer creates new server
func NewServer() *Server {
	return &Server{}
}
