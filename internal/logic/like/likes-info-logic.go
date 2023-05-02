package logic

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"fmt"

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

func (l *LikesInfoLogic) LikesInfo(req *types.LikesInfoReq) (resp *types.LikesInfoRes, err error, msg respx.SucMsg) {
	fmt.Println(req.ExhibitionId)
	DB := l.svcCtx.DB
	var likesInfo []types.LikesInfo
	if err = DB.
		Model(&models.Likes{}).
		Where("exhibition_id = ? and likes_type = ?", req.ExhibitionId, true).
		Find(&likesInfo).Error; err != nil {
		return nil, err, msg
	} else {
		fmt.Println(likesInfo)
		return &types.LikesInfoRes{
			LikesInfo: likesInfo,
		}, nil, msg
	}
}
