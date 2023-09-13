package like

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeLogic) Like(req *types.AddLikesReq) (err error, msg respx.SucMsg) {

	DB := l.svcCtx.DB

	uid, err := l.ctx.Value("Uid").(json.Number).Int64()
	if err != nil {
		return err, msg
	}

	if req.LikesType == 1 { // 添加收藏
		if err = DB.
			Model(&models.Likes{}).
			Where("user_id = ? and likes_id = ? and type = ?", uid, req.Uid, req.Type).
			Assign(&models.Likes{LikesType: true}).
			FirstOrCreate(
				&models.Likes{
					LikesId:   req.Uid,
					UserId:    uint32(uid),
					LikesType: true,
					Type:      req.Type,
				}).
			Error; err == nil {
			return nil, respx.SucMsg{Msg: "收藏成功"}
		} else {
			return err, msg
		}
	} else {
		if err = DB.
			Model(&models.Likes{}).
			Where("user_id = ? and likes_id = ? and type = ?", uid, req.Uid, req.Type).
			Update("likes_type", false).
			Error; err == nil {
			return nil, respx.SucMsg{Msg: "取消收藏成功"}
		} else {
			return err, msg
		}
	}
}
