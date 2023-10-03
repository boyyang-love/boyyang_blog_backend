package tag

import (
	"blog_server/common/respx"
	"blog_server/models"
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

func (l *TagsInfoLogic) TagsInfo(req *types.TagsInfoReq) (resp *types.TagsInfoRes, err error, msg respx.SucMsg) {
	DB := l.svcCtx.DB
	var tagsInfo []types.TagInfo
	if err = DB.
		Debug().
		Model(&models.Tag{}).
		Select("uid", "name", "type").
		Where("type = ?", req.Type).
		Find(&tagsInfo).
		Error; err != nil {
		return nil, err, msg
	} else {
		return &types.TagsInfoRes{
				TagsInfo: tagsInfo,
			}, nil, respx.SucMsg{
				Msg: "获取成功!",
			}
	}
}
