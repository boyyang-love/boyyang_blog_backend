package logic

import (
	"blog_server/common/response"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoReq) (resp *types.UpdateUserInfoRes, err error, msg response.SuccessMsg) {
	isExist := l.svcCtx.DB.
		Model(&models.User{}).
		Where("username = ? and id != ?", req.Username, req.Id).First(&models.User{})
	if isExist.RowsAffected != 0 {
		return nil, errors.New("该用户名已经存在"), msg
	}

	err = l.svcCtx.DB.
		Model(&models.User{}).
		Where("id = ?", req.Id).
		Updates(&models.User{
			Username:  req.Username,
			Gender:    req.Gender,
			Age:       req.Age,
			Address:   req.Address,
			Tel:       &req.Tel,
			Email:     &req.Email,
			AvatarUrl: req.AvatarUrl,
		}).Error

	if err == nil {
		return &types.UpdateUserInfoRes{Msg: "更新成功"}, nil, msg
	} else {
		return nil, err, msg
	}
}
