package main

import "encoding/json"

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func (r Result) Get() (bool, string) {
	var ret string
	if bytes, err := json.Marshal(r); err == nil {
		return true, string(bytes)
	}
	return false, ret
}
