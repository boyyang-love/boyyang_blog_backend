package logic

import (
	"blog_server/common/response"
	"context"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowInfoLogic {
	return &FollowInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowInfoLogic) FollowInfo() (resp *types.FollowInfoRes, err error, msg response.SuccessMsg) {
	// todo: add your logic here and delete this line

	return
}
