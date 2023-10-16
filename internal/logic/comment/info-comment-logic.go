package comment

import (
	"blog_server/common/respx"
	"blog_server/models"
	"context"
	"gorm.io/gorm"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoCommentLogic {
	return &InfoCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoCommentLogic) InfoComment(req *types.InfoCommentReq) (resp *types.InfoCommentRes, err error, msg respx.SucMsg) {
	var count int64
	var infos []types.CommentInfo
	if err = l.svcCtx.DB.
		Model(&models.Comment{}).
		Preload("UserInfo", func(db *gorm.DB) *gorm.DB {
			return db.Select("uid", "username", "gender", "avatar_url", "tel")
		}).
		Where("content_id = ?", req.ContentId).
		Offset((req.Page - 1) * req.Limit).
		Limit(req.Limit).
		Find(&infos).
		Count(&count).Error; err != nil {
		return nil, err, msg
	} else {
		return &types.InfoCommentRes{
			Count: count,
			Infos: infos,
		}, nil, respx.SucMsg{Msg: "获取成功"}
	}
}
