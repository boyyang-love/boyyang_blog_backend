package logic

import (
	"blog_server/models"
	"context"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *types.DeleateBlogCommentReq) (resp *types.DeleateBlogCommentRes, err error) {
	res := l.svcCtx.DB.
		Model(&models.Comment{}).
		Where("id = ?", req.Id).
		Delete(&models.Comment{})
	if res.Error == nil {
		return &types.DeleateBlogCommentRes{Msg: "删除评论成功"}, nil
	} else {
		return nil, res.Error
	}
}
