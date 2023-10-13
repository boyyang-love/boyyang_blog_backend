package article

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArticleLogic {
	return &CreateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateArticleLogic) CreateArticle(req *types.CreateArticleReq) (err error, msg respx.SucMsg) {
	DB := l.svcCtx.DB
	userId, _ := l.ctx.Value("Uid").(json.Number).Int64()
	if err = DB.
		Model(&models.Article{}).
		Create(&models.Article{
			Title:    req.Title,
			SubTitle: req.SubTitle,
			Content:  req.Content,
			Cover:    req.Cover,
			Images:   req.Images,
			UserId:   uint32(userId),
			Tag:      req.Tag,
		}).
		Error; err != nil {
		return err, msg
	} else {
		return nil, respx.SucMsg{Msg: "文章发布成功!"}
	}
}
