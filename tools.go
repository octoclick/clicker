package clicker

import (
	"encoding/base64"
	"strconv"
)

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

// Str2Int ...
func str2Int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
