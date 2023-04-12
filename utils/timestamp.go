package utils

import "time"

func SetTimestamp(data interface{}, isInsert bool) {
	dict := Struct2Map(data)

	_, hasCreatedAt := dict["CreatedAt"]
	_, hasUpdatedAt := dict["UpdatedAt"]
	if !hasCreatedAt || !hasUpdatedAt {
		return
	}

	now := time.Now()
	if isInsert {
		dict["CreatedAt"] = now
	}
	dict["UpdatedAt"] = now

  Map2Struct(dict, data)
}
