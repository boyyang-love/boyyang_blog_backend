package exhibition

import (
	"blog_server/common/respx"
	"blog_server/models"
	"context"
	"database/sql"
	"gorm.io/gorm"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SimilarExhibitionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSimilarExhibitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SimilarExhibitionLogic {
	return &SimilarExhibitionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SimilarExhibitionLogic) SimilarExhibition(req *types.SimilarReq) (resp *types.SimilarRes, err error, msg respx.SucMsg) {

	DB := l.svcCtx.DB
	var ex []types.ExhibitionInfo

	if err = DB.
		Model(models.Exhibition{}).
		Limit(10).
		Preload("UserInfo", func(db *gorm.DB) *gorm.DB {
			return db.Select("uid", "username", "gender", "avatar_url", "tel")
		}).
		Where(
			"tags like @tags",
			sql.Named("tags", "%"+req.Tag+"%"),
		).
		Find(&ex).
		Order("created desc").
		Error; err != nil {
		return nil, err, msg
	} else {
		return &types.SimilarRes{Infos: ex}, nil, respx.SucMsg{Msg: "获取成功!"}
	}
}
