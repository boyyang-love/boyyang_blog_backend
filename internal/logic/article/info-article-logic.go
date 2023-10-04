package article

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type InfoArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoArticleLogic {
	return &InfoArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoArticleLogic) InfoArticle(req *types.InfoArticleReq) (resp *types.InfoArticleRes, err error, msg respx.SucMsg) {
	DB := l.svcCtx.DB

	isPage := false
	var count int64
	var articleInfo []types.ArticleInfo
	if req.Page != 0 && req.Limit != 0 {
		DB = DB.
			Debug().
			Model(&models.Article{}).
			Preload("UserInfo").
			Offset((req.Page - 1) * req.Limit).
			Limit(req.Limit).
			Find(&articleInfo)
		isPage = true
	}

	if req.Uid != 0 {
		DB = DB.
			Debug().
			Model(&models.Article{}).
			Preload("UserInfo").
			Where("uid = ?", req.Uid).
			First(&articleInfo)
	}

	if isPage {
		DB = DB.Count(&count).Offset(-1)
	}

	if err = DB.Error; err != nil {
		return nil, err, msg
	} else {
		return &types.InfoArticleRes{
			Count:       count,
			ArticleInfo: articleInfo,
		}, nil, respx.SucMsg{Msg: "获取成功!"}
	}
}
