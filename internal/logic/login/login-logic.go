package logic

import (
	"blog_server/common/helper"
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error, msg respx.SucMsg) {
	var info types.User
	res := l.svcCtx.DB.
		Model(&models.User{}).
		Where("username = ? and password = ?", req.Username, helper.MakeHash(req.Password)).
		First(&info)

	if res.RowsAffected == 0 {
		return nil, errors.New("请检查账号或密码是否正确"), msg
	} else {
		token, _ := helper.GenerateJwtToken(
			&helper.GenerateJwtStruct{
				Id:       int(info.Id),
				Username: info.Username,
			},
			l.svcCtx.Config.Auth.AccessSecret,
			l.svcCtx.Config.Auth.AccessExpire,
		)

		return &types.LoginRes{
			Info:  info,
			Token: token,
		}, nil, respx.SucMsg{Msg: "登录成功"}
	}
}
