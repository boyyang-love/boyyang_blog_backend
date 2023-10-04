package blog

import (
	"blog_server/common/respx"
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

func (l *CreateBlogLogic) CreateBlog(req *types.CreateBlogReq) (resp *types.CreateBlogRes, err error, msg respx.SucMsg) {
	userId, _ := l.ctx.Value("Uid").(json.Number).Int64()
	blog := &models.Blog{
		Title:    req.Title,
		SubTitle: req.SubTitle,
		Content:  req.Content,
		Cover:    req.Cover,
		UserId:   uint32(userId),
		Tag:      req.Tags,
	}

	if err := l.svcCtx.DB.
		Model(&models.Blog{}).
		Create(&blog).
		Error; err != nil {
		return nil, err, msg
	} else {
		return &types.CreateBlogRes{Uid: uint32(uint(blog.Uid))},
			nil,
			respx.SucMsg{
				Msg: "博客发布成功!",
			}
	}
}
