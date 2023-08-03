package helper

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MakeHash(s string) string {
	h := md5.New()
	_, err := io.WriteString(h, s)
	if err != nil {
		return ""
	}
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return hash
}
