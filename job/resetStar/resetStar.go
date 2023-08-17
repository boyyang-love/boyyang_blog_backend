package resetStar

import (
	"blog_server/internal/svc"
	"blog_server/models"
	"fmt"
)

type ResetStartLogic struct {
	svcCtx *svc.ServiceContext
}

func NewResetStartLogic(svcCtx *svc.ServiceContext) *ResetStartLogic {
	return &ResetStartLogic{
		svcCtx: svcCtx,
	}
}

func (l *ResetStartLogic) Start() (err error) {
	if err := l.svcCtx.DB.
		Model(&models.Star{}).
		Select("star_type").
		Where("star_type = ?", true).
		Updates(map[string]interface{}{"star_type": false}).Error; err != nil {
		return err
	} else {
		fmt.Println("star 状态更新成功！")
		return nil
	}
}
