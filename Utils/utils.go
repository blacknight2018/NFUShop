package Utils

import "encoding/json"

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
