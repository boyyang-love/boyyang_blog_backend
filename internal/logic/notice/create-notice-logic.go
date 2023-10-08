package notice

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNoticeLogic {
	return &CreateNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateNoticeLogic) CreateNotice(req *types.NoticeCreateReq) (err error, msg respx.SucMsg) {

	DB := l.svcCtx.DB
	id, err := l.ctx.Value("Uid").(json.Number).Int64() // 用户id

	if err = DB.
		Model(&models.Notice{}).
		Create(&models.Notice{
			UserId:  uint32(id),
			Content: req.Content,
		}).
		Error; err != nil {
		return err, msg
	} else {
		return nil, respx.SucMsg{Msg: "公告发布成功!"}
	}
}
