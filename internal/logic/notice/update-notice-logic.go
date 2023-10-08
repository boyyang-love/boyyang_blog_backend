package notice

import (
	"blog_server/common/respx"
	"blog_server/models"
	"context"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoticeLogic {
	return &UpdateNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateNoticeLogic) UpdateNotice(req *types.NoticeUpdateReq) (err error, msg respx.SucMsg) {
	DB := l.svcCtx.DB

	if err = DB.
		Model(&models.Notice{}).
		Where("uid = ?", req.Uid).
		Update("content", req.Content).
		Error; err != nil {
		return err, msg
	} else {
		return nil, respx.SucMsg{Msg: "更新成功!"}
	}
}
