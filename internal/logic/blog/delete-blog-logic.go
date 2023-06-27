package logic

import (
	"blog_server/common/respx"
	"blog_server/models"
	"context"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBlogLogic {
	return &DeleteBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBlogLogic) DeleteBlog(req *types.DeleteBlogReq) (err error, msg respx.SucMsg) {
	if err = l.svcCtx.DB.
		Model(&models.Blog{}).
		Where("id = ?", req.Id).
		Delete(&models.Blog{}).
		Error; err == nil {
		return nil, respx.SucMsg{Msg: "博客删除成功"}
	} else {
		return err, msg
	}
}
