package exhibition

import (
	"blog_server/common/respx"
	"blog_server/models"
	"context"
	"gorm.io/gorm"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDownloadLogic {
	return &UpdateDownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDownloadLogic) UpdateDownload(req *types.UpdateDownloadReq) (err error, msg respx.SucMsg) {
	DB := l.svcCtx.DB

	if err := DB.
		Model(&models.Exhibition{}).
		Where("uid = ?", req.Uid).
		Update("download", gorm.Expr("download + ?", 1)).
		Error; err != nil {
		return err, msg
	} else {
		return nil, respx.SucMsg{Msg: "更新成功!"}
	}
}
