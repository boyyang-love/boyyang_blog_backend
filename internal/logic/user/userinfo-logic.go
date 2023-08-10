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

	err, otherInfo := l.userOtherInfo(req.Id)
	if err != nil {
		return nil, err, msg
	}

	return &types.UserInfoRes{
		UserInfo:      *userInfo,
		UserOtherInfo: *otherInfo,
	}, err, msg
}

func (l *UserInfoLogic) userInfo(userId uint) (err error, info *types.User) {
	DB := l.svcCtx.DB
	var userInfo *types.User

	if err = DB.Model(&models.User{}).
		Where("id = ?", userId).
		First(&userInfo).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("不存在该用户"), nil
		}
		return err, nil
	}

	return nil, userInfo
}

func (l *UserInfoLogic) userOtherInfo(userId uint) (err error, info *types.UserOtherInfo) {
	DB := l.svcCtx.DB

	var Publish int64
	var Likes int64
	var Follows int64
	var ThumbsUpNum int
	var ThumbsUp []int

	if err = DB.Model(&models.Exhibition{}).
		Where("user_id = ?", userId).
		Count(&Publish).
		Error; err != nil {
		return err, nil
	}

	if err = DB.Model(&models.Likes{}).
		Where("user_id = ?", userId).
		Count(&Likes).
		Error; err != nil {
		return err, nil
	}

	if err = DB.Model(&models.Follow{}).
		Where("follow_user_id", userId).
		Count(&Follows).
		Error; err != nil {
		return err, nil
	}

	if err = DB.Model(&models.Exhibition{}).
		Select("ThumbsUp").
		Where("user_id = ?", userId).
		Scan(&ThumbsUp).
		Error; err != nil {
		return err, nil
	}

	for _, num := range ThumbsUp {
		ThumbsUpNum += num
	}

	return nil, &types.UserOtherInfo{
		Publish:  int(Publish),
		Likes:    int(Likes),
		Follows:  int(Follows),
		ThumbsUp: ThumbsUpNum,
	}
}
