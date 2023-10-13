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
			Model(&models.Article{}).
			Order("created desc").
			Preload("UserInfo").
			Offset((req.Page - 1) * req.Limit).
			Limit(req.Limit).
			Find(&articleInfo)
		isPage = true
	}

	if req.Uid != 0 {
		DB = DB.
			Model(&models.Article{}).
			Order("created desc").
			Preload("UserInfo").
			Where("uid = ?", req.Uid).
			First(&articleInfo)
	}

	if isPage {
		DB = DB.Count(&count).Offset(-1)
	}

	var userId uint32
	if req.UserId != 0 {
		userId = req.UserId
	} else {
		userId = uint32(articleInfo[0].UserId)
	}
	_, cardInfo := l.getCardInfo(userId)

	if err = DB.Error; err != nil {
		return nil, err, msg
	} else {
		return &types.InfoArticleRes{
			Count:       count,
			ArticleInfo: articleInfo,
			CardInfo:    cardInfo,
		}, nil, respx.SucMsg{Msg: "获取成功!"}
	}
}

func (l *InfoArticleLogic) getCardInfo(userId uint32) (err error, cardInfo types.CardInfo) {

	var follow int64
	if err = l.svcCtx.DB.
		Model(&models.Follow{}).
		Where("user_id = ?", userId).
		Count(&follow).Error; err != nil {
		return err, cardInfo
	}

	var fans int64
	if err = l.svcCtx.DB.
		Model(&models.Follow{}).
		Where("follow_user_id = ?", userId).
		Count(&fans).
		Error; err != nil {
		return err, cardInfo
	}

	var thumbs int64
	if err = l.svcCtx.DB.
		Model(&models.Article{}).
		Select("sum(thumbs_up) as thumbs").
		Where("user_id = ?", userId).
		Scan(&thumbs).
		Error; err != nil {
		return err, cardInfo
	}

	var articleIds []uint32
	var articles int64
	if err = l.svcCtx.DB.
		Model(&models.Article{}).
		Select("uid").
		Where("user_id = ?", userId).
		Scan(&articleIds).
		Count(&articles).
		Error; err != nil {
		return err, cardInfo
	}

	var comments int64
	if err = l.svcCtx.DB.
		Model(&models.Comment{}).
		Where("blog_id in ?", articleIds).
		Count(&comments).
		Error; err != nil {
		return err, cardInfo
	}

	return nil, types.CardInfo{
		Follow:  follow,
		Fans:    fans,
		Thumb:   thumbs,
		Article: articles,
		Comment: comments,
	}
}
