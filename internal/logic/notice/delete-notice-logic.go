package notice

import (
	"blog_server/common/respx"
	"blog_server/models"
	"context"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteNoticeLogic {
	return &DeleteNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteNoticeLogic) DeleteNotice(req *types.NoticeDeleteReq) (err error, msg respx.SucMsg) {

	DB := l.svcCtx.DB

	if err = DB.
		Model(&models.Notice{}).
		Where("uid = ?", req.Uid).
		Error; err != nil {
		return err, msg
	} else {
		return nil, respx.SucMsg{Msg: "删除成功!"}
	}
}
