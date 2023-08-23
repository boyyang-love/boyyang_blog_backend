package search

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (resp *types.SearchRes, err error, msg respx.SucMsg) {
	if req.Type == 1 {
		res, count, err := l.searchExhibitionInfos(req)
		if err != nil {
			return nil, err, msg
		} else {
			return &types.SearchRes{
				ExhibitionInfo: res,
				BlogInfo:       nil,
				Count:          count,
			}, nil, msg
		}
	} else {
		res, count, err := l.searchBlogInfos(req)
		if err != nil {
			return nil, err, msg
		} else {
			return &types.SearchRes{
				ExhibitionInfo: nil,
				BlogInfo:       res,
				Count:          count,
			}, nil, msg
		}
	}
}

func (l *SearchLogic) searchExhibitionInfos(req *types.SearchReq) (resp []types.SearchExhibitionInfos, count int64, err error) {
	if err = l.svcCtx.DB.
		Debug().
		Model(&models.Exhibition{}).
		Offset((req.Page-1)*req.Limit).
		Limit(req.Limit).
		Where("title like @keyword or tags like @keyword", sql.Named("keyword", "%"+req.Keyword+"%")).
		Find(&resp).
		Offset(-1).
		Count(&count).
		Error; err != nil {
		return nil, count, err
	} else {
		return resp, count, nil
	}
}

func (l *SearchLogic) searchBlogInfos(req *types.SearchReq) (resp []types.SearchBlogInfos, count int64, err error) {
	if err = l.
		svcCtx.
		DB.
		Model(&models.Blog{}).
		Offset((req.Page-1)*req.Limit).
		Limit(req.Limit).
		Where("tag like ?", "%"+req.Keyword+"%").
		Find(&resp).
		Offset(-1).
		Count(&count).
		Error; err != nil {
		return nil, count, err
	} else {
		return resp, count, nil
	}
}
