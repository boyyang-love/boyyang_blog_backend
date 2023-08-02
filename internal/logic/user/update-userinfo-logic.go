package user

import (
	"blog_server/common/respx"
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

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoReq) (err error, msg respx.SucMsg) {
	DB := l.svcCtx.DB
	var userInfo models.User
	if err = DB.
		Model(&models.User{}).
		Where("id = ?", req.Id).
		First(&userInfo).
		Error; err != nil {
		return errors.New("不存在该用户"), msg
	} else {
		var count int64
		DB.
			Select("id", "username").
			Model(&models.User{}).
			Where("username = ? and id != ?", req.Username, req.Id).
			Count(&count)
		if count == 0 { // 更新用户信息 用户名不能重复
			if err = l.svcCtx.DB.
				Model(&models.User{}).
				Where("id = ?", req.Id).
				Updates(&models.User{
					Username:        req.Username,
					Gender:          req.Gender,
					Age:             req.Age,
					Address:         req.Address,
					Tel:             &req.Tel,
					Email:           &req.Email,
					AvatarUrl:       req.AvatarUrl,
					BackgroundImage: req.BackgroundImage,
				}).Error; err != nil {
				return err, msg
			} else {
				return nil, respx.SucMsg{
					Msg: "用户信息更新成功",
				}
			}
		} else {
			return errors.New("该用户名已经存在"), msg
		}
	}
}
