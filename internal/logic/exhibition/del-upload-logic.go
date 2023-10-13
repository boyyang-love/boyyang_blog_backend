package exhibition

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type DelUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUploadLogic {
	return &DelUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelUploadLogic) DelUpload(req *types.DelUploadReq) (err error, msg respx.SucMsg) {
	err, isExhibitionsUse := l.exhibitionNoUse(req)
	if err != nil {
		return err, msg
	}
	err, isBlogUse := l.blogNoUse(req)
	if err != nil {
		return err, msg
	}

	err, isArticleUse := l.articleNoUse(req)
	if err != nil {
		return err, msg
	}

	if isExhibitionsUse && isBlogUse && isArticleUse {
		_, err = l.svcCtx.Client.Object.Delete(context.Background(), req.Key)
		if err != nil {
			return err, msg
		}
		return nil, respx.SucMsg{Msg: "图片删除成功"}
	} else {
		return nil, respx.SucMsg{Msg: "当前图片有其它引用"}
	}
}

func (l *DelUploadLogic) exhibitionNoUse(req *types.DelUploadReq) (err error, isNoUse bool) {
	var count int64
	DB := l.svcCtx.DB
	if err = DB.
		Model(&models.Exhibition{}).
		Where("cover = ?", req.Key).
		Count(&count).Error; err != nil {
		return err, false
	} else {
		if count >= 1 {
			return nil, false
		} else {
			return nil, true
		}
	}
}

func (l *DelUploadLogic) blogNoUse(req *types.DelUploadReq) (err error, isNoUse bool) {
	var count int64
	DB := l.svcCtx.DB
	if err = DB.
		Model(&models.Blog{}).
		Where("cover = ?", req.Key).
		Count(&count).Error; err != nil {
		return err, false
	} else {
		if count >= 1 {
			return nil, false
		} else {
			return nil, true
		}
	}
}

func (l *DelUploadLogic) articleNoUse(req *types.DelUploadReq) (err error, isNoUse bool) {
	var count int64
	DB := l.svcCtx.DB
	if err = DB.
		Model(&models.Article{}).
		Where("cover = ? or images like ?", req.Key, "%"+req.Key+"%").
		Count(&count).Error; err != nil {
		return err, false
	} else {
		if count >= 1 {
			return nil, false
		} else {
			return nil, true
		}
	}
}
