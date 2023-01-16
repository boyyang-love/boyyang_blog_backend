package logic

import (
	"blog_server/common/response"
	"blog_server/models"
	"context"
	"errors"

	"blog_server/internal/svc"
	"blog_server/internal/types"

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

func (l *ApprovalExhibitionLogic) ApprovalExhibition(req *types.ApprovalReq) (resp *types.ApprovalRes, err error, msg response.SuccessMsg) {
	db := l.svcCtx.DB

	if err = db.
		Model(&models.Exhibition{}).
		Where("id = ?", req.Id).
		Updates(&models.Exhibition{
			Status:    req.Status,
			RejectRes: req.Reason,
		}).Error; err == nil {
		return &types.ApprovalRes{Id: req.Id}, err, response.SuccessMsg{Msg: "状态更新成功"}
	} else {
		return nil, errors.New("状态更新失败"), msg
	}
}
