package like

import (
	"blog_server/common/respx"
	"blog_server/models"
	"context"
	"encoding/json"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowInfoLogic {
	return &FollowInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowInfoLogic) FollowInfo() (resp *types.FollowInfoRes, err error, msg respx.SucMsg) {
	id, err := l.ctx.Value("Uid").(json.Number).Int64() // 用户id
	var followInfo []models.Follow
	DB := l.svcCtx.DB
	if err = DB.
		Model(&models.Follow{}).
		Where("user_id = ? and  follow_type = ? ", id, true).
		Scan(&followInfo).Error; err != nil {
		return nil, err, msg
	} else {
		var followIds []uint      // 用户关注列表 following id
		var userInfo []types.User // 关注列表用户信息
		for _, follow := range followInfo {
			followIds = append(followIds, follow.FollowUserId)
		}
		if err =
			DB.
				Select("uid", "username", "gender", "avatar_url", "tel").
				Model(&models.User{}).
				Where("uid in ?", followIds).
				Scan(&userInfo).
				Error; err != nil {
			return nil, err, msg
		} else {
			return &types.FollowInfoRes{
				FollowingUser: userInfo,
			}, nil, msg
		}
	}
}
