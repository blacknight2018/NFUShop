package Utils

import (
	"encoding/json"
	"strconv"
)

const EmptyString = ""

func GetJSONStringArray(jsonString string) []string {
	var ret []string
	json.Unmarshal([]byte(jsonString), &ret)
	return ret
}
func StringArrayToJSON(strArray []string) string {
	if bytes, err := json.Marshal(strArray); err == nil {
		return string(bytes)
	}
	return EmptyString
}

func StrToInt(str string) int {
	var ret int
	ret = 0
	if data, err := strconv.Atoi(str); err == nil {
		ret = data
	}
	return ret
}
