package helper

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
)

func CosFileUpload(client *cos.Client, fileHeader *multipart.FileHeader, path string) (url string, err error) {
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "image/jpeg",
		},
	}

	cloudPath := fmt.Sprintf("%s/%s", path, fileHeader.Filename)
	f, _ := fileHeader.Open()
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	_, err = client.Object.Put(context.Background(), cloudPath, f, opt)

	if err == nil {
		return cloudPath, nil
	} else {
		return "", err
	}
}
