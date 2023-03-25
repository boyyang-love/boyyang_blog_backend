package logic

import (
	"blog_server/common/response"
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

func (l *DeleteBlogLogic) DeleteBlog(req *types.DeleteBlogReq) (resp *types.DeleteBlogRes, err error, msg response.SuccessMsg) {
	res := l.svcCtx.DB.
		Model(&models.Blog{}).
		Where("id = ?", req.Id).
		Delete(&models.Blog{})
	if res.Error == nil {
		return &types.DeleteBlogRes{Msg: "删除成功"}, nil, msg
	} else {
		return nil, res.Error, msg
	}
}
