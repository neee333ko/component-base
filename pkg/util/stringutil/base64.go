package stringutil

import "encoding/base64"

func DecodeBase64(i string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(i)
}
