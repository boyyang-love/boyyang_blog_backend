package like

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"encoding/json"
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

	var ids []uint
	var likesInfos []types.LikesInfo
	if req.ExhibitionId == 0 {
		ids, err = l.getLikesIds()
	} else {
		ids = append(ids, req.ExhibitionId)
	}

	if err = l.svcCtx.DB.
		Model(&models.Exhibition{}).
		Debug().
		Where("id in (?)", ids).
		Find(&likesInfos).
		Error; err != nil {
		return nil, err, msg
	} else {
		return &types.LikesInfoRes{
			LikesInfo: likesInfos,
		}, nil, msg
	}
}

func (l *LikesInfoLogic) getLikesIds() (ids []uint, err error) {
	userid, _ := l.ctx.Value("Id").(json.Number).Int64()
	DB := l.svcCtx.DB
	var likesIds []uint

	if err = DB.
		Model(&models.Likes{}).
		Select("exhibition_id").
		Where("user_id=?", userid).
		Scan(&likesIds).
		Error; err != nil {
		return nil, err
	} else {
		return likesIds, nil
	}
}
