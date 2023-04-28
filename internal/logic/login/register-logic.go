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

	res := l.svcCtx.DB.
		Where("username = ? or tel = ?", req.Username, req.Tel).
		First(&models.User{})

	if res.RowsAffected == 0 {
		info := models.User{
			Username: req.Username,
			Password: helper.MakeHash(req.Password),
			Tel:      &req.Tel,
		}

		l.svcCtx.DB.
			Model(&models.User{}).
			Create(&info)
		return &types.RegisterRes{
			Id: int(info.Id),
		}, nil, respx.SucMsg{Msg: "账号注册成功"}

	} else {
		return nil, errors.New("该用户已经存在"), msg
	}
}
