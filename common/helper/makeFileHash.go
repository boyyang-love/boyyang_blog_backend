package helper

import (
	"crypto/md5"
	"fmt"
	"mime/multipart"
)

func MakeFileHash(file multipart.File, fileHeader *multipart.FileHeader) (hash string, err error) {
	h := make([]byte, fileHeader.Size)
	_, err = file.Read(h)

	if err == nil {
		hash := fmt.Sprintf("%x", md5.Sum(h))
		return hash, nil
	} else {
		return "", err
	}
}
