package clicker

import (
	"encoding/base64"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func base64Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// decode ...
func base64Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return []byte("")
	}
	return data
}
