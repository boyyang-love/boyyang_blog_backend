package logic

import (
	"context"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikesInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikesInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikesInfoLogic {
	return &LikesInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikesInfoLogic) LikesInfo(req *types.GetLikesReq) (resp *types.GetLikesRes, err error) {
	// todo: add your logic here and delete this line

	return
}
