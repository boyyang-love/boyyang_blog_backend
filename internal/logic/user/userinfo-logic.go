package user

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
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
	// 如果不传ID 则使用token中的ID
	if req.Id == 0 {
		id, _ := l.ctx.Value("Id").(json.Number).Int64()
		req.Id = uint(id)
	}

	err, userInfo := l.userInfo(req.Id)
	if err != nil {
		return nil, err, msg
	}

	return userInfo, nil, msg
}

func (l *UserInfoLogic) userInfo(userId uint) (err error, userInfo *types.UserInfoRes) {
	DB := l.svcCtx.DB

	DB = DB.
		Model(&models.User{}).
		Where("id = ?", userId).
		First(&userInfo)

	if err = DB.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("不存在该用户"), userInfo
		}

		return err, nil
	}

	return err, userInfo
}
