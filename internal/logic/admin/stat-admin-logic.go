package admin

import (
	"blog_server/common/respx"
	"blog_server/models"
	"context"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatAdminLogic {
	return &StatAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatAdminLogic) StatAdmin() (resp *types.AdminStatRes, err error, msg respx.SucMsg) {

	userCount, err := l.getUserCount()
	if err != nil {
		return nil, err, msg
	}

	imageCount, imageDownloadCount, err := l.getImageCount()
	if err != nil {
		return nil, err, msg
	}

	blogCount, err := l.getBlogCount()
	if err != nil {
		return nil, err, msg
	}

	articleCount, err := l.getArticleCount()
	if err != nil {
		return nil, err, msg
	}

	return &types.AdminStatRes{
		UserCount:          userCount,
		ImageCount:         imageCount,
		ImageDownloadCount: imageDownloadCount,
		BlogCount:          blogCount,
		ArticleCount:       articleCount,
	}, nil, respx.SucMsg{Msg: "获取成功!"}
}

func (l *StatAdminLogic) getUserCount() (count int64, err error) {
	if err = l.svcCtx.DB.
		Model(&models.User{}).
		Count(&count).
		Error; err != nil {
		return count, err
	} else {
		return count, nil
	}
}

func (l *StatAdminLogic) getImageCount() (imageCount int64, imageDownloadCount int64, err error) {
	err = l.svcCtx.DB.
		Model(&models.Exhibition{}).
		Count(&imageCount).Error
	if err != nil {
		return imageCount, imageDownloadCount, err
	}

	err = l.svcCtx.DB.
		Model(&models.Exhibition{}).
		Select("sum(download) as count").
		Where("download != ?", 0).
		Find(&imageDownloadCount).
		Error
	if err != nil {
		return imageCount, imageDownloadCount, err
	}
	return imageCount, imageDownloadCount, nil
}

func (l *StatAdminLogic) getBlogCount() (count int64, err error) {
	if err = l.svcCtx.DB.
		Model(&models.Blog{}).
		Count(&count).
		Error; err != nil {
		return count, err
	} else {
		return count, nil
	}
}

func (l *StatAdminLogic) getArticleCount() (count int64, err error) {
	if err = l.svcCtx.DB.
		Model(&models.Article{}).
		Count(&count).
		Error; err != nil {
		return count, err
	} else {
		return count, nil
	}
}
