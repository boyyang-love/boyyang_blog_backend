package exhibition

import (
	"blog_server/common/respx"
	"blog_server/models"
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"strconv"
	"strings"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExhibitionInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExhibitionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExhibitionInfoLogic {
	return &ExhibitionInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExhibitionInfoLogic) ExhibitionInfo(req *types.ExhibitionInfoReq) (resp *types.ExhibitionInfoRes, err error, msg respx.SucMsg) {
	userid, _ := l.ctx.Value("Id").(json.Number).Int64()
	DB := l.svcCtx.DB
	ids := strings.Split(req.Ids, ",")

	var count int64
	var exInfo []types.ExhibitionInfo

	status, err := l.getStatus(userid)
	if err != nil {
		return nil, err, msg
	}

	likes, err := l.getLikesIds(userid)

	if len(ids) > 0 && req.Ids != "" {
		if err := DB.
			Model(&models.Exhibition{}).
			Preload("UserInfo", func(db *gorm.DB) *gorm.DB {
				return db.Select("id", "username", "gender", "avatar_url", "tel")
			}).
			Where("status = ? and user_id = ?", req.Type, userid).
			Order("created_at desc").
			Find(&exInfo, ids).
			Count(&count).
			Error; err == nil {
			return &types.ExhibitionInfoRes{
					Exhibitions:    exInfo,
					Count:          int(count),
					InReview:       int(status[0]),
					Approved:       int(status[1]),
					ReviewRjection: int(status[2]),
					LikesIds:       likes,
				},
				nil,
				respx.SucMsg{Msg: "获取成功"}
		} else {
			return nil, err, msg
		}
	} else {
		if req.Page != "" && req.Limit != "" {
			page, _ := strconv.Atoi(req.Page)
			limit, _ := strconv.Atoi(req.Limit)
			DB = DB.
				Offset((page - 1) * limit).
				Limit(limit)
		}

		if err := DB.
			Model(&models.Exhibition{}).
			Preload("UserInfo", func(db *gorm.DB) *gorm.DB {
				return db.Select("id", "username", "gender", "avatar_url", "tel")
			}).
			Where("status = ? and user_id = ?", req.Type, userid).
			Order("created_at desc").
			Find(&exInfo).
			Offset(-1).
			Count(&count).
			Error; err == nil {
			return &types.ExhibitionInfoRes{
					Exhibitions:    exInfo,
					Count:          int(count),
					InReview:       int(status[0]),
					Approved:       int(status[1]),
					ReviewRjection: int(status[2]),
					LikesIds:       likes,
				},
				nil,
				respx.SucMsg{Msg: "获取成功"}
		} else {
			return nil, err, msg
		}
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
		Select("exhibition_id").
		Where("user_id = ? and likes_type = ?", userid, true).
		Scan(&likesIds).
		Error; err != nil {
		return nil, err
	} else {
		return likesIds, nil
	}
}
