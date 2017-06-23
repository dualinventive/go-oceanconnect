// Copyright 2017 The go-oceanconnect authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package oceanconnect

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"github.com/Sirupsen/logrus"
)

type Server struct {
}

//DeviceDataChanged stuct with response data
type DeviceDataChanged struct {
	NotifyType string
	DeviceID   string
	GatewayID  string
	RequestID  string
	Service    `json:"service"`
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Got request:")
	rd, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(rd))

	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("err: %v", err)
		return
	}

	var data DeviceDataChanged
	if err := json.Unmarshal(contents, &data); err != nil {
		logrus.Errorf("err: %v", err)
		return
	}
	fmt.Printf("data: %+v\n", data)

}

func (s *Server) ListenAndServe(uri string) error {
	http.HandleFunc("/", s.handler)
	return http.ListenAndServe(uri, nil)
}

// NewServer creates new server
func NewServer() *Server {
	return &Server{}
}
