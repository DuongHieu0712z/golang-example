package utils

import "encoding/json"

func Struct2Map(data interface{}) map[string]interface{} {
	arr, _ := json.Marshal(data)
	var res map[string]interface{}
	json.Unmarshal(arr, &res)
	return res
}

func Map2Struct(dict map[string]interface{}, data interface{}) {
	arr, _ := json.Marshal(dict)
	json.Unmarshal(arr, data)
}

func Bson2Map() {

}

func Map2Bson() {

}
