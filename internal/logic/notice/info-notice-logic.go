package notice

import (
	"blog_server/common/respx"
	"blog_server/models"
	"context"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoNoticeLogic {
	return &InfoNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoNoticeLogic) InfoNotice(req *types.NoticeInfoReq) (resp *types.NoticeInfoRes, err error, msg respx.SucMsg) {
	var infos []types.NoticeInfo
	var count int64
	if err = l.svcCtx.DB.
		Model(&models.Notice{}).
		Order("created desc").
		Select("uid", "content", "created", "user_id").
		Offset((req.Page - 1) * req.Limit).
		Limit(req.Limit).
		Find(&infos).
		Offset(-1).
		Count(&count).
		Error; err != nil {
		return nil, err, msg
	} else {
		return &types.NoticeInfoRes{Count: count, Infos: infos}, nil, respx.SucMsg{Msg: "获取成功!"}
	}
}
