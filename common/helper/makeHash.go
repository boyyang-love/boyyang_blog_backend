package helper

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MakeHash(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return hash
}
