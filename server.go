// Copyright 2017 The go-oceanconnect authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package oceanconnect

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//DeviceDataChanged stuct with response data
type DeviceDataChanged struct {
	NotifyType string
	DeviceID   string
	GatewayID  string
	RequestID  string
	Service    `json:"service"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got request:")
	fmt.Printf("%+v\n", r.Header)
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n\n", string(contents))

	var data DeviceDataChanged
	if err := json.Unmarshal(contents, &data); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("data: %+v\n", data)

}

// NewServer creates new server
func NewServer() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
