package helper

import (
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

func InitCloudBase(clientUrl string, clientSecretId string, clientSecretKey string) (client *cos.Client) {
	u, _ := url.Parse(clientUrl)
	b := &cos.BaseURL{BucketURL: u}
	client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  clientSecretId,
			SecretKey: clientSecretKey,
		},
	})

	return client
}
