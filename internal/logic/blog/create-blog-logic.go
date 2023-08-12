package blog

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
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
	userId, _ := l.ctx.Value("Id").(json.Number).Int64()
	blog := &models.Blog{
		Title:    req.Title,
		SubTitle: req.SubTitle,
		Content:  req.Content,
		Cover:    req.Cover,
		UserId:   uint(userId),
		Tag:      req.Tags,
	}

	if err := l.svcCtx.DB.
		Model(&models.Blog{}).
		Create(&blog).
		Error; err != nil {
		return nil, err, msg
	} else {
		if strings.Trim(req.Tags, " ") != "" {
			// 创建tags
			var tags []models.Tag
			for _, tag := range strings.Split(req.Tags, ",") {
				tags = append(tags, models.Tag{
					Name:   tag,
					BlogId: blog.Id,
					UserId: uint(userId),
				})
			}
			if err = l.svcCtx.DB.
				Model(&models.Tag{}).
				FirstOrCreate(
					&models.Tag{},
					&tags,
				).
				Error; err != nil {
				return nil, err, msg
			}
		}
		return &types.CreateBlogRes{Id: blog.Id},
			nil,
			respx.SucMsg{
				Msg: "博客发布成功！",
			}
	}

}
