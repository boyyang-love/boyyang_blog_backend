package article

import (
	"blog_server/common/respx"
	"blog_server/models"
	"context"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateArticleLogic) UpdateArticle(req *types.UpdateArticleReq) (err error, msg respx.SucMsg) {
	DB := l.svcCtx.DB
	if err = DB.
		Model(&models.Article{}).
		Where("uid = ?", req.Uid).
		Updates(&models.Article{
			Title:    req.Title,
			SubTitle: req.SubTitle,
			Content:  req.Content,
			Tag:      req.Tag,
		}).
		Error; err != nil {
		return err, msg
	} else {
		return nil, respx.SucMsg{Msg: "文章更新成功!"}
	}
}
