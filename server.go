// Copyright 2017 The go-oceanconnect authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package oceanconnect

import (
	"encoding/json"
	"net/http"

	"sync"

	"io/ioutil"

	"github.com/Sirupsen/logrus"
)

type NotificationFunc func(interface{}) error

type Server struct {
	cbsLock sync.RWMutex
	cbs     map[Notification]NotificationFunc
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var n struct {
		NotifyType string `json:"notifyType"`
	}
	if err := json.Unmarshal(buf, &n); err != nil {
		logrus.Errorf("error decoding notification type")
		return
	}

	if err := s.runCallback(Notification(n.NotifyType), buf); err != nil {
		logrus.Errorf("Error running callback: %v", err)
		return
	}
}

func (s *Server) runCallback(not Notification, dec []byte) error {
	s.cbsLock.RLock()
	defer s.cbsLock.RUnlock()

	if s.cbs == nil {
		logrus.Infof("no callbacks registered, callback received")
	}
	cb, ok := s.cbs[not]
	if ok {
		v, err := notificationDeserializer(not, dec)
		if err != nil {
			return err
		}
		return cb(v)
	}
	logrus.Debugf("no callback registered for %s", string(not))
	return nil
}

func (s *Server) ListenAndServe(uri string) error {
	http.HandleFunc("/", s.handler)
	return http.ListenAndServe(uri, nil)
}

func (s *Server) RegisterCallback(not Notification, cb NotificationFunc) {
	s.cbsLock.Lock()
	if s.cbs == nil {
		s.cbs = make(map[Notification]NotificationFunc)
	}
	s.cbs[not] = cb
	s.cbsLock.Unlock()
}
