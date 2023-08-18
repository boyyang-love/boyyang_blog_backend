package user

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
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
	// 判断是否存在被修改用户
	err, isExist := l.isExistUser(req.Uid)
	if err != nil {
		return err, msg
	}

	if isExist {
		// 判断新用户名是否重复
		err, isSameName := l.isSameName(req.Uid, req.Username)
		if err != nil {
			return err, msg
		}
		if isSameName {
			return errors.New("该用户名已经被注册"), msg
		} else {
			err := l.updateUserInfo(req)
			if err != nil {
				return err, msg
			} else {
				return nil, respx.SucMsg{Msg: "用户信息修改成功"}
			}
		}
	} else {
		return errors.New("不存在该用户"), msg
	}
}

// 查看当前用户是否存在
func (l *UpdateUserInfoLogic) isExistUser(userId uint) (err error, isExist bool) {
	var user struct {
		Uid int `json:"uid"`
	}
	DB := l.svcCtx.DB.Model(&models.User{}).
		Select("uid").
		Where("uid = ?", userId).
		First(&user)

	if err = DB.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("不存在该用户"), isExist
		}
		return err, isExist
	} else {
		return nil, true
	}
}

// 判断是否重名
func (l *UpdateUserInfoLogic) isSameName(userId uint, name string) (err error, isSameName bool) {
	var count int64
	DB := l.svcCtx.DB.
		Select("uid", "username").
		Model(&models.User{}).
		Where("username = ? and uid != ?", name, userId).
		Count(&count)
	if err = DB.Error; err != nil {
		return err, isSameName
	}
	if count == 0 {
		return nil, false
	} else {
		return nil, true
	}
}

// 修改用户信息
func (l *UpdateUserInfoLogic) updateUserInfo(req *types.UpdateUserInfoReq) (err error) {
	DB := l.
		svcCtx.DB.
		Model(&models.User{}).
		Where("uid =?", req.Uid).
		Updates(&models.User{
			Username:        req.Username,
			Age:             req.Age,
			Gender:          req.Gender,
			AvatarUrl:       req.AvatarUrl,
			Tel:             req.Tel,
			Email:           req.Email,
			Qq:              req.Qq,
			Wechat:          req.Wechat,
			GitHub:          req.GitHub,
			Address:         req.Address,
			BackgroundImage: req.BackgroundImage,
			Motto:           req.Motto,
		})

	if err = DB.Error; err != nil {
		return err
	} else {
		return nil
	}
}
