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

type DashboardLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDashboardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DashboardLogic {
	return &DashboardLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DashboardLogic) Dashboard() (resp *types.DashboardRes, err error, msg response.SuccessMsg) {

	type thumbs struct {
		ThumbsUp  int
		Likes     int
		Publish   int
		Following int
	}

	DB := l.svcCtx.DB
	var userInfo types.User
	var dashboard []types.Dashboard
	var thumbsUp thumbs
	var exhibitionInfo []types.DashboardExhibition

	id, err := l.ctx.Value("Id").(json.Number).Int64()
	if err != nil {
		return nil, err, msg
	}

	// userinfo
	err = DB.
		Model(models.User{}).
		Where("id = ?", id).
		Scan(&userInfo).
		Error

	// 统计图
	err = DB.
		Model(&models.Exhibition{}).
		Select("count(*) as value", "DATE_FORMAT(created_at, '%Y-%m-%d') as name").
		Group("name").
		Where("user_id = ?", id).
		Scan(&dashboard).
		Error

	// 用户统计图
	err = DB.
		Table(
			"(?) as l, (?) as b, (?) as f",
			DB.
				Model(&models.Likes{}).
				Select("count(*) as likes").
				Where("user_id = ?", id),
			DB.
				Model(&models.Blog{}).
				Select("sum(thumbs_up) as thumbs_up", "count(*) as publish").
				Where("user_id = ?", id),
			DB.
				Model(&models.Follow{}).
				Select("count(*) as following").
				Where("follow_user_id = ?", id),
		).
		Select("l.likes", "b.thumbs_up", "b.publish", "f.following").
		Scan(&thumbsUp).
		Error

	// 图片列表
	err = DB.
		Model(&models.Exhibition{}).
		Offset(0).
		Limit(10).
		Where("user_id", id).
		Find(&exhibitionInfo).
		Error

	if err == nil {
		return &types.DashboardRes{
			UserInfo: types.DashboardUserInfo{
				User:      userInfo,
				ThumbsUp:  &thumbsUp.ThumbsUp,
				Like:      &thumbsUp.Likes,
				Publish:   &thumbsUp.Publish,
				Following: &thumbsUp.Following,
			},
			Dashboard:   dashboard,
			Exhibitions: exhibitionInfo,
		}, nil, response.SuccessMsg{Msg: "获取成功"}
	} else {
		return nil, err, msg
	}
}
