package exhibition

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"database/sql"
	"encoding/json"
	"gorm.io/gorm"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExhibitionInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type Params struct {
	Uids     string
	Page     int
	Limit    int
	Public   bool
	Type     int
	UserId   uint
	Sort     string
	Keywords string
}

func NewExhibitionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExhibitionInfoLogic {
	return &ExhibitionInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExhibitionInfoLogic) ExhibitionInfo(req *types.ExhibitionInfoReq) (resp *types.ExhibitionInfoRes, err error, msg respx.SucMsg) {
	userid, _ := l.ctx.Value("Uid").(json.Number).Int64()

	params := Params{
		Uids:     req.Uids,
		Limit:    req.Limit,
		Page:     req.Page,
		Type:     req.Type,
		Public:   req.Public,
		UserId:   uint(userid),
		Sort:     req.Sort,
		Keywords: req.Keywords,
	}

	status, err := l.getStatus(userid)
	if err != nil {
		return nil, err, msg
	}

	likes, err := l.getLikesIds(userid)
	if err != nil {
		return nil, err, msg
	}

	star, err := l.getStarIds(userid)
	if err != nil {
		return nil, err, msg
	}

	exhibitions, count, err := l.getExhibitionInfo(params)

	if err != nil {
		return nil, err, msg
	} else {
		return &types.ExhibitionInfoRes{
			Count:          int(count),
			Exhibitions:    exhibitions,
			InReview:       int(status[0]),
			Approved:       int(status[1]),
			ReviewRjection: int(status[2]),
			LikesIds:       likes,
			StarIds:        star,
		}, nil, msg
	}

}

func (l *ExhibitionInfoLogic) getExhibitionInfo(params Params) (exhibitions []types.ExhibitionInfo, count int64, err error) {
	DB := l.svcCtx.DB

	DB = DB.Model(&models.Exhibition{})

	if params.Page > 0 && params.Limit > 0 {
		DB = DB.Offset((params.Page - 1) * params.Limit).
			Limit(params.Limit)
	}

	if params.Public {
		DB = DB.Where("status = ?", params.Type)
	} else {
		DB = DB.Where("status = ? and user_id = ?", params.Type, params.UserId)
	}

	if params.Uids != "" {
		DB = DB.Where("id in (?)", strings.Split(params.Uids, ","))
	}

	if params.Keywords != "" {
		DB = DB.Where("tags like @keywords or title like @keywords", sql.Named("keywords", "%"+params.Keywords+"%"))
	}

	DB = DB.
		Preload("UserInfo", func(db *gorm.DB) *gorm.DB {
			return db.Select("uid", "username", "gender", "avatar_url", "tel")
		}).
		Order(params.Sort)

	if err = DB.
		Model(&models.Exhibition{}).
		Find(&exhibitions).
		Offset(-1).
		Count(&count).
		Error; err != nil {
		return nil, count, err
	} else {
		return exhibitions, count, nil
	}

}

func (l *ExhibitionInfoLogic) getStatus(userid int64) (status []int64, err error) {
	var counts []int64
	var count1 int64
	var count2 int64
	var count3 int64
	err = l.svcCtx.DB.
		Model(&models.Exhibition{}).
		Where("status = ? and user_id = ?", 1, userid).
		Count(&count1).
		Error

	err = l.svcCtx.DB.
		Model(&models.Exhibition{}).
		Where("status = ? and user_id = ?", 2, userid).
		Count(&count2).
		Error

	err = l.svcCtx.DB.
		Model(&models.Exhibition{}).
		Where("status = ? and user_id = ?", 3, userid).
		Count(&count3).
		Error

	if err == nil {
		return append(counts, count1, count2, count3), nil
	} else {
		return nil, err
	}

}

func (l *ExhibitionInfoLogic) getLikesIds(userid int64) (likesIds []int, err error) {
	if err = l.svcCtx.DB.
		Model(&models.Likes{}).
		Select("likes_id").
		Where("user_id = ? and likes_type = ?", userid, true).
		Scan(&likesIds).
		Error; err != nil {
		return nil, err
	} else {
		return likesIds, nil
	}
}

func (l *ExhibitionInfoLogic) getStarIds(userid int64) (starIds []int, err error) {
	if err = l.svcCtx.DB.
		Model(&models.Star{}).
		Select("star_id").
		Where("user_id = ? and star_type = ?", userid, 1).
		Scan(&starIds).
		Error; err != nil {
		return nil, err
	} else {
		return starIds, nil
	}
}
