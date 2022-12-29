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

type CreateBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBlogLogic {
	return &CreateBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateBlogLogic) CreateBlog(req *types.CreateBlogReq) (resp *types.CreateBlogRes, err error, msg response.SuccessMsg) {
	userId, _ := l.ctx.Value("Id").(json.Number).Int64()
	blog := models.Blog{
		Title:    req.Title,
		SubTitle: req.SubTitle,
		Content:  req.Content,
		Cover:    req.Cover,
		UserId:   uint(userId),
	}
	res := l.svcCtx.DB.
		Model(&models.Blog{}).
		Create(&blog)
	if res.Error == nil {
		return &types.CreateBlogRes{Id: int(blog.Id)}, nil, msg
	} else {
		return nil, res.Error, msg
	}
}
