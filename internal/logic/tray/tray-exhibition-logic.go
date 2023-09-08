package tray

import (
	"context"

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

func (l *TrayExhibitionLogic) TrayExhibition(req *types.TrayReq) (resp *types.TrayRes, err error) {
	// todo: add your logic here and delete this line
	return
}
