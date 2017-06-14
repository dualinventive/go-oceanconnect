// Copyright 2017 The go-oceanconnect authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package oceanconnect

import (
	"strings"
	"time"
)

const ocTimeLayout = "20060102T150405Z07:00"

// OcTime is used for unmarshalling the times communicated via the API
// to time.Time
type OcTime struct {
	time.Time
}

// UnmarshalJSON reads the times to time.Time
func (ct *OcTime) UnmarshalJSON(b []byte) error {
	var err error
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return nil
	}
	ct.Time, err = time.Parse(ocTimeLayout, s)
	return err
}
