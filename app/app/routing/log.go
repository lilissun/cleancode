package main

import (
	"encoding/json"
	"time"
)

type routelog struct {
	Tag      string `json:"tag,omitempty"`
	Route    string `json:"route,omitempty"`
	Request  string `json:"request,omitempty"`
	Response string `json:"response,omitempty"`
	Error    string `json:"error,omitempty"`
	Duration int64  `json:"duration,omitempty"`
}

func newErrorLog(route string, request string, err error, duration time.Duration) routelog {
	return routelog{Tag: "ERROR", Route: route, Request: request, Error: err.Error(), Duration: duration.Nanoseconds()}
}

func newInfoLog(route string, request string, response string, duration time.Duration) routelog {
	return routelog{Tag: "INFO", Route: route, Request: request, Response: response, Duration: duration.Nanoseconds()}
}

func (rlog routelog) String() string {
	data, _ := json.Marshal(rlog)
	return string(data)
}
