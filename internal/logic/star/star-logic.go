package star

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"encoding/json"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type StarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StarLogic {
	return &StarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StarLogic) Star(req *types.StarReq) (err error, msg respx.SucMsg) {
	userId, err := l.ctx.Value("Uid").(json.Number).Int64()
	if err != nil {
		return err, msg
	}
	if req.StarType == 1 {
		if err = l.svcCtx.DB.
			Model(&models.Star{}).
			Where("star_id = ?", req.Uid).
			Assign("star_type", req.StarType).
			FirstOrCreate(&models.Star{
				UserId:   uint32(userId),
				StarId:   req.Uid,
				StarType: true,
				Type:     req.Type,
			}).
			Error; err != nil {
			return err, msg
		} else {
			err = l.updateStar(uint(req.Uid), req.Type, 1)
			if err != nil {
				return err, msg
			}
			return nil, respx.SucMsg{Msg: "点赞成功"}
		}
	} else {
		if err = l.svcCtx.DB.
			Model(&models.Star{}).
			Where("star_id = ?", req.Uid).
			Update("star_type", req.StarType).
			Error; err != nil {
			return err, msg
		} else {
			err = l.updateStar(uint(req.Uid), req.Type, -1)
			if err != nil {
				return err, msg
			}
			return nil, respx.SucMsg{Msg: "取消点赞成功"}
		}
	}
}

func (l *StarLogic) updateStar(starId uint, starType int, val int) (err error) {
	if starType == 1 {
		if err = l.svcCtx.DB.
			Model(&models.Exhibition{}).
			Where("uid = ?", starId).
			Update("thumbs_up", gorm.Expr("thumbs_up + ?", val)).
			Error; err != nil {
			return err
		}
	}
	if starType == 2 {
		if err = l.svcCtx.DB.
			Model(&models.Blog{}).
			Where("uid = ?", starId).
			Update("thumbs_up", gorm.Expr("thumbs_up + ?", val)).
			Error; err != nil {
			return err
		}
	}

	if starType == 3 {
		if err = l.svcCtx.DB.
			Model(&models.Article{}).
			Where("uid = ?", starId).
			Update("thumbs_up", gorm.Expr("thumbs_up + ?", val)).
			Error; err != nil {
			return err
		}
	}
	return err
}
