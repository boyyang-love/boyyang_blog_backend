package tag

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTagLogic {
	return &CreateTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTagLogic) CreateTag(req *types.CreateTagReq) (err error, msg respx.SucMsg) {
	DB := l.svcCtx.DB
	if err := DB.
		Model(&models.Tag{}).
		Where("name = ? and type = ?", req.Name, req.Type).
		FirstOrCreate(&models.Tag{
			Name: req.Name,
			Type: req.Type,
		}).Error; err != nil {
		return err, msg
	} else {
		return nil, respx.SucMsg{Msg: "标签创建成功!"}
	}
}
