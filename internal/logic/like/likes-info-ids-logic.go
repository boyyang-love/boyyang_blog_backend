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

type LikesInfoIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikesInfoIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikesInfoIdsLogic {
	return &LikesInfoIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikesInfoIdsLogic) LikesInfoIds() (resp *types.LikesInfoIds, err error, msg respx.SucMsg) {
	userid, _ := l.ctx.Value("Uid").(json.Number).Int64()
	var ids []int
	if err = l.svcCtx.DB.
		Model(&models.Likes{}).
		Select("exhibition_id").
		Where("user_id = ? and likes_type = ?", userid, true).
		Find(&models.Likes{}).
		Scan(&ids).Error; err != nil {
		return nil, err, msg
	} else {
		return &types.LikesInfoIds{Uids: ids}, nil, msg
	}
}
