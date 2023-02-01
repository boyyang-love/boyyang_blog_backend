package logic

import (
	"blog_server/common/response"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeLogic) Like(req *types.AddLikesReq) (err error, msg response.SuccessMsg) {

	DB := l.svcCtx.DB

	userId, err := l.ctx.Value("Id").(json.Number).Int64()
	if err != nil {
		return err, msg
	}

	if err = DB.Model(&models.Likes{}).
		FirstOrCreate(
			&models.Likes{},
			&models.Likes{
				ExhibitionId: req.LikesId,
				UserId:       uint(userId),
			}).
		Where("user_id = ? and exhibition_id = ?", userId, req.LikesId).
		Error;
		err == nil {
		return nil, response.SuccessMsg{Msg: "收藏成功"}
	} else {
		return err, msg
	}

}
