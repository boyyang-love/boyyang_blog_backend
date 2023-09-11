package tray

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type TrayExhibitionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTrayExhibitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TrayExhibitionLogic {
	return &TrayExhibitionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TrayExhibitionLogic) TrayExhibition(req *types.TrayReq) (resp *types.TrayRes, err error, msg respx.SucMsg) {
	ex, count, err := l.exhibitions(*req)
	if err != nil {
		return nil, err, msg
	}

	return &types.TrayRes{
		Count:           count,
		TrayExhibitions: ex,
	}, err, respx.SucMsg{Msg: "获取成功"}
}

func (l *TrayExhibitionLogic) exhibitions(req types.TrayReq) (resp []types.TrayExhibitionInfo, count int64, err error) {
	DB := l.svcCtx.DB

	if req.Page != 0 && req.Limit != 0 {
		DB = DB.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	if err = DB.
		Model(&models.Exhibition{}).
		Where("status = ?", 2).
		Order("created desc").
		Find(&resp).
		Count(&count).Error; err != nil {
		return nil, count, err
	} else {
		return resp, count, err
	}
}
