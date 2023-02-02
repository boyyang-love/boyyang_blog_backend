package logic

import (
	"blog_server/common/response"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(req *types.AddAndUnFollowReq) (err error, msg response.SuccessMsg) {
	DB := l.svcCtx.DB
	id, err := l.ctx.Value("Id").(json.Number).Int64()
	if err != nil {
		return err, msg
	}

	if req.FollowId == uint(id) {
		return nil, response.SuccessMsg{Msg: "不能关注自己"}
	}

	if req.FollowType == 1 {
		if err = DB.
			Debug().
			Model(&models.Follow{}).
			FirstOrCreate(
				&models.Follow{},
				&models.Follow{
					FollowUserId: req.FollowId,
					UserId:       uint(id),
				}).
			Where("follow_user_id = ? and user_id = ?", req.FollowId, id).
			Error; err == nil {
			return nil, response.SuccessMsg{Msg: "关注成功"}
		}
	} else {
		if err = DB.
			Model(&models.Follow{}).
			Delete(&models.Follow{}).
			Where("follow_user_id = ? and user_id = ?", req.FollowId, id).
			Error; err == nil {
			return nil, response.SuccessMsg{Msg: "取消关注成功"}
		}
	}

	return err, msg
}
