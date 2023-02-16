package logic

import (
	"blog_server/common/response"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"sort"
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
	var userInfo types.DashboardUser
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

	// dashboard
	dashboard, err = getDashboardData(DB, uint(id))

	// 用户统计图
	err = DB.
		Table(
			"(?) as l, (?) as b, (?) as f",
			DB.
				Model(&models.Likes{}).
				Select("count(*) as likes").
				Where("user_id = ? and likes_type = ?", id, true),
			DB.
				Model(&models.Blog{}).
				Select("sum(thumbs_up) as thumbs_up", "count(*) as publish").
				Where("user_id = ?", id),
			DB.
				Model(&models.Follow{}).
				Select("count(*) as following").
				Where("follow_user_id = ? and follow_type = ?", id, true),
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
		Order("created_at desc").
		Find(&exhibitionInfo).
		Error

	if err == nil {
		return &types.DashboardRes{
			UserInfo: types.DashboardUserInfo{
				ThumbsUp:      &thumbsUp.ThumbsUp,
				Like:          &thumbsUp.Likes,
				Publish:       &thumbsUp.Publish,
				Following:     &thumbsUp.Following,
				DashboardUser: userInfo,
			},
			Dashboard:   dashboard,
			Exhibitions: exhibitionInfo,
		}, nil, response.SuccessMsg{Msg: "获取成功"}
	} else {
		return nil, err, msg
	}
}

func getDashboardData(DB *gorm.DB, id uint) (dashboard []types.Dashboard, err error) {
	var blogDashboard []types.Dashboard
	var publishDashboard []types.Dashboard
	m := make(map[string]types.Dashboard)

	DB.
		Debug().
		Model(&models.Blog{}).
		Where("user_id", id).
		Select("count(*) blog_publish_value", "DATE_FORMAT(created_at,'%Y-%m-%d') name").
		Group("name").
		Scan(&blogDashboard)

	DB.
		Model(&models.Exhibition{}).
		Where("user_id", id).
		Select("count(*) exhibitions_publish_value", "DATE_FORMAT(created_at,'%Y-%m-%d') name").
		Group("name").
		Scan(&publishDashboard)

	dashboard = append(dashboard, append(blogDashboard, publishDashboard...)...)

	for _, d := range dashboard {
		if d.BlogPublishValue != "" {
			m[d.Name] = types.Dashboard{
				Name:                    d.Name,
				BlogPublishValue:        d.BlogPublishValue,
				ExhibitionsPublishValue: m[d.Name].ExhibitionsPublishValue,
			}
		}
		if d.ExhibitionsPublishValue != "" {
			m[d.Name] = types.Dashboard{
				Name:                    d.Name,
				BlogPublishValue:        m[d.Name].BlogPublishValue,
				ExhibitionsPublishValue: d.ExhibitionsPublishValue,
			}
		}
	}
	dashboard = []types.Dashboard{}
	for _, t := range m {
		dashboard = append(dashboard, t)
	}

	sort.Slice(dashboard, func(i, j int) bool {
		return dashboard[i].Name < dashboard[j].Name
	})

	return dashboard, nil
}
