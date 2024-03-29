package blog

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBlogLogic {
	return &UpdateBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBlogLogic) UpdateBlog(req *types.UpdateBlogReq) (resp *types.UpdateBlogRes, err error, msg respx.SucMsg) {
	blogInfo := models.Blog{
		Title:    req.Title,
		SubTitle: req.SubTitle,
		Content:  req.Content,
		Cover:    req.Cover,
	}
	res := l.svcCtx.DB.
		Model(&models.Blog{}).
		Where("uid = ?", req.Uid).
		Updates(&blogInfo)
	if res.Error == nil {
		return &types.UpdateBlogRes{Msg: "更新成功"}, err, respx.SucMsg{Msg: "更新成功"}
	} else {
		return nil, res.Error, msg
	}
}
