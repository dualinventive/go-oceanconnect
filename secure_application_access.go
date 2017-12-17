// Copyright 2017 The go-oceanconnect authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package oceanconnect

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
)

// loginResponse struct with response data
type loginResponse struct {
	AccessToken string
	TokenType   string
	ExpiresIn   int64
	Scope       string
}

// Login with the client to oceanconnect
func (c *Client) Login() error {
	v := url.Values{}
	v.Set("appId", c.cfg.AppID)
	v.Set("Secret", c.cfg.Secret)

	resp, err := c.c.PostForm(c.cfg.URL+"/iocm/app/sec/v1.1.0/login", v)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("invalid response code: " + resp.Status)
	}
	l := loginResponse{}
	err = json.NewDecoder(resp.Body).Decode(&l)
	if err == nil {
		c.token = l.AccessToken
		c.tokenExpires = time.Now().Add(time.Second * time.Duration(l.ExpiresIn))
		logrus.Infof("Token retrieved, expires: %v", c.tokenExpires)
	}
	return err
}
