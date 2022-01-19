package utils

import "encoding/json"

// Convert 转换结构体数据
func Convert(from, to interface{}) error {
	data, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &to)
}
