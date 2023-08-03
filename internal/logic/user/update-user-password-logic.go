package user

import (
	"blog_server/common/helper"
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"encoding/json"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPasswordLogic {
	return &UpdateUserPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Info struct {
	Password string `json:"password"`
}

func (l *UpdateUserPasswordLogic) UpdateUserPassword(req *types.UpdateUserPasswordReq) (err error, msg respx.SucMsg) {
	userId, err := l.ctx.Value("Id").(json.Number).Int64()
	if err != nil {
		return errors.New("token 解析失败"), msg
	}

	if err = l.updatePassword(userId, req.Password); err != nil {
		return err, msg
	}

	return nil, respx.SucMsg{Msg: "修改密码成功"}
}

func (l *UpdateUserPasswordLogic) updatePassword(userId int64, newPassword string) (err error) {
	DB := l.svcCtx.DB
	var info struct {
		Password string `json:"password"`
	}

	DB = DB.
		Model(&models.User{}).
		Select("password").
		Where("id=?", userId).
		First(&info)

	if err = DB.Error; err != nil {
		return err
	} else {
		hashNewPassword := helper.MakeHash(newPassword)
		if info.Password != hashNewPassword {
			info.Password = hashNewPassword
			if err = DB.Save(&info).Error; err != nil {
				return err
			}

			return nil
		}

		return errors.New("新密码不能和旧密码一致")
	}
}
