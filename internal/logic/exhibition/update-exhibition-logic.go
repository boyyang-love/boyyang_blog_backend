package logic

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateExhibitionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateExhibitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateExhibitionLogic {
	return &UpdateExhibitionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateExhibitionLogic) UpdateExhibition(req *types.UpdateExhibitionReq) (resp *types.UpdateExhibitionRes, err error, msg respx.SucMsg) {
	DB := l.svcCtx.DB

	if err = DB.
		Model(&models.Exhibition{}).
		Where("id = ?", req.Id).
		Updates(&models.Exhibition{
			Title:    req.Title,
			SubTitle: req.SubTitle,
			Des:      req.Des,
		}).Error; err == nil {
		return &types.UpdateExhibitionRes{Id: req.Id}, nil, respx.SucMsg{Msg: "更新成功"}
	} else {
		return nil, err, msg
	}
}
