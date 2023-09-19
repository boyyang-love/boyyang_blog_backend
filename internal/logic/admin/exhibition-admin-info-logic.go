package admin

import (
	"blog_server/common/respx"
	"blog_server/models"
	"context"
	"gorm.io/gorm"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExhibitionAdminInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExhibitionAdminInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExhibitionAdminInfoLogic {
	return &ExhibitionAdminInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExhibitionAdminInfoLogic) ExhibitionAdminInfo(req *types.AdminExhibitionsReq) (resp *types.AdminExhibitionsRes, err error, msg respx.SucMsg) {
	exhibitions, count, err := l.getExhibitions(req)
	if err != nil {
		return nil, err, msg
	}

	return &types.AdminExhibitionsRes{
		Count:       count,
		Exhibitions: exhibitions,
	}, err, respx.SucMsg{Msg: "获取成功"}
}

func (l *ExhibitionAdminInfoLogic) getExhibitions(req *types.AdminExhibitionsReq) (exhibitions []types.AdminExhibitionInfo, count int64, err error) {
	DB := l.svcCtx.DB

	if req.Sort != "" {
		DB = DB.Order(req.Sort)
	}

	if err = DB.
		Debug().
		Preload("UserInfo", func(db *gorm.DB) *gorm.DB {
			return db.Select("uid", "username", "gender", "avatar_url", "tel")
		}).
		Model(&models.Exhibition{}).
		Offset((req.Page-1)*req.Limit).
		Limit(req.Limit).
		Where("status = ?", req.Type).
		Find(&exhibitions).
		Offset(-1).
		Count(&count).
		Error; err != nil {
		return nil, count, err
	} else {
		return exhibitions, count, nil
	}
}
