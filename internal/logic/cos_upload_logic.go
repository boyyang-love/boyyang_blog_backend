package logic

import (
	"blog_server/common/response"
	"blog_server/internal/config"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"context"
	"crypto/tls"
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

func (l *CosUploadLogic) CosUpload(req *types.CosUploadReq) (resp *types.CosUploadRes, err error, msg response.SuccessMsg) {
	// 获取 cos token
	token, err := getAccessToken(l.svcCtx.Config)
	// 用户不传 路径 默认储存在 images 文件夹下
	if req.Path == "" {
		req.Path = "images"
	}

	if err == nil {
		// 获取 cos 前台 上传路径 相关信息
		resp, err = getUploadMsg(
			token,
			fmt.Sprintf("%s/%s", req.Path, req.FileName),
		)
		resp.Key = fmt.Sprintf("%s/%s", req.Path, req.FileName)
		return resp, err, msg
	} else {
		return nil, err, msg
	}
}

func getAccessToken(config config.Config) (token string, err error) {

	type Token struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	type AppidAndSecret struct {
		Appid     string
		Secret    string
		GrantType string
	}

	var accessToken Token
	appidAndSecret := AppidAndSecret{
		Appid:     config.AppidAndSecret.AppId,
		Secret:    config.AppidAndSecret.Secret,
		GrantType: "client_credential",
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	res, err := client.Get(
		fmt.Sprintf(
			"https://api.weixin.qq.com/cgi-bin/token?grant_type=%s&appid=%s&secret=%s",
			appidAndSecret.GrantType,
			appidAndSecret.Appid,
			appidAndSecret.Secret,
		),
	)

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

func getUploadMsg(token string, uploadPath string) (resp *types.CosUploadRes, err error) {

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
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	//  获取上传链接 相关信息
	res, err := client.Post(
		postUrl,
		"application/json",
		strings.NewReader(string(queryStr)),
	)

	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
