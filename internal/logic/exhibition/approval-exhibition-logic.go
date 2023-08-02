package exhibition

import (
	"blog_server/common/errorx"
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApprovalExhibitionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApprovalExhibitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApprovalExhibitionLogic {
	return &ApprovalExhibitionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApprovalExhibitionLogic) ApprovalExhibition(req *types.ApprovalReq) (resp *types.ApprovalRes, err error, msg respx.SucMsg) {
	DB := l.svcCtx.DB

	if err = DB.
		Model(&models.Exhibition{}).
		Where("id = ?", req.Id).
		Updates(&models.Exhibition{
			Status:    req.Status,
			RejectRes: req.Reason,
		}).Error; err == nil {
		return &types.ApprovalRes{Id: req.Id}, err, respx.SucMsg{Msg: "状态更新成功"}
	} else {
		return nil, errorx.NewDefaultError(err.Error()), msg
	}
}
