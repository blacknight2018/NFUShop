package Result

import "encoding/json"

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func (r Result) Get() string {
	var ret string
	if bytes, err := json.Marshal(r); err == nil {
		ret = string(bytes)
	}
	return ret
}
