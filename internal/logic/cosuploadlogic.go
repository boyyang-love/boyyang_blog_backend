package logic

import (
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"strings"
)

type CosUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCosUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CosUploadLogic {
	return &CosUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CosUploadLogic) CosUpload(req *types.CosUploadReq) (resp *types.CosUploadRes, err error) {
	token, err := getAccessToken()
	if err == nil {
		resp, err = getDownloadUrl(token, fmt.Sprintf("images/%s", req.FileName))
		resp.Key = fmt.Sprintf("images/%s", req.FileName)
		return resp, err
	} else {
		return nil, err
	}
}

func getAccessToken() (token string, err error) {

	type Token struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	var accessToken Token

	res, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wx20773192bbf7b3b8&secret=309f0b28ace40739a7f15d4772537774")
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		return "", err
	}

	return accessToken.AccessToken, nil
}

func getDownloadUrl(token string, uploadPath string) (resp *types.CosUploadRes, err error) {

	type QueryParams struct {
		Env   string `json:"env"`
		Token string `json:"token,omitempty"`
		Path  string `json:"path"`
	}

	query := QueryParams{
		Env:  "prod-2g5hif5wbec83baa",
		Path: uploadPath,
	}

	queryStr, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	postUrl := fmt.Sprintf("https://api.weixin.qq.com/tcb/uploadfile?access_token=%s", token)

	//  获取下载链接
	res, err := http.Post(
		postUrl,
		"application/json",
		strings.NewReader(string(queryStr)),
	)

	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	fmt.Println(string(body))
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
