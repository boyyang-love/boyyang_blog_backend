package tag

import (
	"context"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TagsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTagsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TagsInfoLogic {
	return &TagsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TagsInfoLogic) TagsInfo() (resp *types.TagsInfoRes, err error) {
	// todo: add your logic here and delete this line

	return
}
