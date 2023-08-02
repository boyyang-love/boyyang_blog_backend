package login

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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error, msg respx.SucMsg) {

	if err = l.svcCtx.DB.
		Where("username = ? or tel = ?", req.Username, req.Tel).
		First(&models.User{}).Error; err != nil {
		info := models.User{
			Username: req.Username,
			Password: helper.MakeHash(req.Password),
			Tel:      &req.Tel,
		}
		if err = l.svcCtx.DB.
			Model(&models.User{}).
			Create(&info).Error; err != nil {
			return nil, err, msg
		} else {
			return &types.RegisterRes{
				Id: info.Id,
			}, nil, respx.SucMsg{Msg: "账号注册成功"}
		}
	} else {
		return nil, errors.New("用户名已经存在或者该手机号已经注册"), msg
	}
}
