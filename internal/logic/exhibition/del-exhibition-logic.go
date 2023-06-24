package logic

import (
	"blog_server/common/respx"
	"blog_server/models"
	"context"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelExhibitionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelExhibitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelExhibitionLogic {
	return &DelExhibitionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelExhibitionLogic) DelExhibition(req *types.DelExhibitionReq) (err error, msg respx.SucMsg) {
	if err = l.svcCtx.DB.
		Model(&models.Exhibition{}).
		Where("id = ?", req.Id).
		Delete(&models.Exhibition{}).
		Error; err != nil {
		return err, msg
	} else {
		return nil, respx.SucMsg{Msg: "图片删除成功"}
	}
}
