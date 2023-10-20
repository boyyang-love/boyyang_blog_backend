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
	userId, err := l.ctx.Value("Uid").(json.Number).Int64()
	if err != nil {
		return nil, err, msg
	}
	DB := l.svcCtx.DB.
		Model(&models.Article{}).
		Preload("UserInfo").
		Order(req.Sort)

	var count int64
	var articleInfo []types.ArticleInfo
	var cardInfo types.CardInfo

	if req.UserId != 0 {
		DB = DB.Where("user_id = ?", req.UserId)
		_, cardInfo = l.getCardInfo(req.UserId)
	} else {
		_, cardInfo = l.getCardInfo(uint32(userId))
	}

	if req.Type == 1 {
		DB = DB.Where("user_id = ?", userId)
	}

	if req.Type == 2 {
		DB = DB.Where("user_id != ?", userId)
	}

	if req.Page != 0 && req.Limit != 0 {
		DB = DB.
			Preload("UserInfo").
			Offset((req.Page - 1) * req.Limit).
			Limit(req.Limit).
			Find(&articleInfo).
			Offset(-1).
			Count(&count)
	}

	if req.Uid != 0 {
		DB = DB.
			Where("uid = ?", req.Uid).
			First(&articleInfo)
	}

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

	currentUserId, _ := l.ctx.Value("Uid").(json.Number).Int64()
	var followIds []int64
	if err = l.svcCtx.DB.
		Model(&models.Follow{}).
		Select("follow_user_id").
		Where("user_id = ? and follow_type = ?", currentUserId, true).
		Find(&followIds).
		Error; err != nil {
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
		Select("IFNULL(sum(thumbs_up), 0) as thumbs").
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
		Where("content_id in ?", articleIds).
		Count(&comments).
		Error; err != nil {
		return err, cardInfo
	}

	var starIds []int64
	if err = l.svcCtx.DB.
		Model(&models.Star{}).
		Select("star_id").
		Where("user_id = ? and star_type = ? and type = ?", currentUserId, true, 3).
		Find(&starIds).
		Error; err != nil {
		return err, cardInfo
	}

	return nil, types.CardInfo{
		Follow:    follow,
		Fans:      fans,
		Thumb:     thumbs,
		Article:   articles,
		Comment:   comments,
		FollowIds: followIds,
		StarIds:   starIds,
	}
}
