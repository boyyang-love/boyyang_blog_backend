package logic

import (
	"blog_server/models"
	"context"
	"gorm.io/gorm"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThumbsUpCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThumbsUpCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbsUpCommentLogic {
	return &ThumbsUpCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThumbsUpCommentLogic) ThumbsUpComment(req *types.ThumbsUpBlogCommentReq) (resp *types.ThumbsUpBlogCommentRes, err error) {
	res := l.svcCtx.DB.
		Model(&models.Comment{}).
		Where("id = ?", req.Id).
		Update("thumbs_up", gorm.Expr("thumbs_up + ?", 1))
	if res.Error == nil {
		return &types.ThumbsUpBlogCommentRes{Msg: "点赞成功"}, nil
	} else {
		return nil, res.Error
	}
}
