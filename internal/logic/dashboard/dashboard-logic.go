package dashboard

import (
	"blog_server/common/respx"
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

func (l *DashboardLogic) Dashboard() (resp *types.DashboardRes, err error, msg respx.SucMsg) {

	DB := l.svcCtx.DB
	uid, err := l.ctx.Value("Uid").(json.Number).Int64()
	if err != nil {
		return nil, err, msg
	}

	// 用户信息
	userInfo, err := getDashboardUserInfo(DB, uint(uid))
	// dashboard
	dashboard, err := getDashboardData(DB, uint(uid))
	// cardData
	cardData, err := getDashboardCard(DB, uint(uid))
	// exhibition
	exhibition, err := getDashboardExhibition(DB, uint(uid))

	if err == nil {
		return &types.DashboardRes{
			UserInfo: types.DashboardUserInfo{
				ThumbsUp:      &cardData.ThumbsUp,
				Like:          &cardData.Likes,
				Publish:       &cardData.Publish,
				Following:     &cardData.Following,
				DashboardUser: userInfo,
			},
			Dashboard:   dashboard,
			Exhibitions: exhibition,
		}, nil, respx.SucMsg{Msg: "获取成功"}
	} else {
		return nil, err, msg
	}
}

// 仪表盘数据
func getDashboardData(DB *gorm.DB, id uint) (dashboard []types.Dashboard, err error) {
	var blogDashboard []types.Dashboard
	var publishDashboard []types.Dashboard
	m := make(map[string]types.Dashboard)

	DB.
		Model(&models.Blog{}).
		Where("user_id = ?", id).
		Select("count(*) blog_publish_value", "DATE_FORMAT(created_at,'%Y-%m-%d') name").
		Group("name").
		Scan(&blogDashboard)

	DB.
		Model(&models.Exhibition{}).
		Where("user_id = ? and status = ?", id, 2).
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

// 用户信息
func getDashboardUserInfo(DB *gorm.DB, id uint) (userInfo types.User, err error) {
	// userinfo
	err = DB.
		Model(models.User{}).
		Where("uid = ?", id).
		Find(&userInfo).
		Error

	return userInfo, err
}

type CardData struct {
	ThumbsUp  int
	Likes     int
	Publish   int
	Following int
}

// 用户卡片数据
func getDashboardCard(DB *gorm.DB, id uint) (cardData CardData, err error) {
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
		Scan(&cardData).
		Error

	return cardData, err
}

// 图片墙数据
func getDashboardExhibition(DB *gorm.DB, id uint) (exhibition []types.DashboardExhibition, err error) {
	// 图片列表
	err = DB.
		Model(&models.Exhibition{}).
		Offset(0).
		Limit(10).
		Where("user_id", id).
		Order("created_at desc").
		Find(&exhibition).
		Error

	return exhibition, err
}
