package logic

import (
	"blog_server/common/respx"
	"blog_server/models"
	"context"
	"errors"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoRes, err error, msg respx.SucMsg) {
	userInfo := models.User{}
	res := l.svcCtx.DB.
		Model(&models.User{}).
		Where("id = ?", req.Id).
		First(&userInfo)

	if res.RowsAffected == 0 {
		return nil, errors.New("不存在该用户"), msg
	} else {
		return &types.UserInfoRes{
			Id:        int(userInfo.Id),
			Username:  userInfo.Username,
			Age:       userInfo.Age,
			Gender:    userInfo.Gender,
			AvatarUrl: userInfo.AvatarUrl,
			Tel:       *userInfo.Tel,
			Email:     *userInfo.Email,
			Address:   userInfo.Address,
		}, nil, msg
	}
}
