package logic

import (
	"blog_server/common/response"
	"blog_server/models"
	"context"
	"gorm.io/gorm"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThumbsUpBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThumbsUpBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbsUpBlogLogic {
	return &ThumbsUpBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThumbsUpBlogLogic) ThumbsUpBlog(req *types.ThumbsUpBlogReq) (resp *types.ThumbsUpBlogRes, err error, msg response.SuccessMsg) {
	res := l.svcCtx.DB.
		Model(&models.Blog{}).
		Where("id = ?", req.Id).
		Update("thumbs_up", gorm.Expr("thumbs_up + ?", 1))
	if res.Error == nil {
		return &types.ThumbsUpBlogRes{Msg: "点赞成功"}, nil, msg
	} else {
		return nil, res.Error, msg
	}
}
