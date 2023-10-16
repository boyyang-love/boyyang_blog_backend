package article

import (
	"blog_server/common/respx"
	"blog_server/models"
	"context"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleLogic) DeleteArticle(req *types.DeleteArticleReq) (err error, msg respx.SucMsg) {
	DB := l.svcCtx.DB

	if err = DB.
		Model(&models.Article{}).
		Where("uid = ?", req.Uid).
		Delete(&models.Article{}).Error; err != nil {
		return err, msg
	} else {
		l.svcCtx.DB.
			Model(&models.Comment{}).
			Where("content_id = ?", req.Uid).
			Delete(&models.Comment{})
		return nil, respx.SucMsg{Msg: "删除成功!"}
	}
}
