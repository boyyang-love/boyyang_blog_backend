package tray

import (
	"blog_server/common/respx"
	"blog_server/models"
	"context"
	"fmt"

	"blog_server/internal/svc"
	"blog_server/internal/types"

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
	fmt.Println(req.Page, req.Limit)
	ex, err := l.exhibitions()
	if err != nil {
		return nil, err, msg
	}

	return &types.TrayRes{TrayExhibitions: ex}, err, respx.SucMsg{Msg: "获取成功"}
}

func (l *TrayExhibitionLogic) exhibitions() (resp []types.TrayExhibitionInfo, err error) {
	DB := l.svcCtx.DB
	if err = DB.
		Model(&models.Exhibition{}).
		Order("created desc").
		Find(&resp).Error; err != nil {
		return nil, err
	} else {
		return resp, err
	}
}
