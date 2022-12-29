package logic

import (
	"blog_server/common/response"
	"blog_server/models"
	"context"
	"github.com/jinzhu/copier"
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

func (l *ExhibitionInfoLogic) ExhibitionInfo(req *types.ExhibitionInfoReq) (resp *types.ExhibitionInfoRes, err error, msg response.SuccessMsg) {
	DB := l.svcCtx.DB

	ids := strings.Split(req.Ids, ",")

	var ex []models.Exhibition
	var exInfo []types.ExhibitionInfo
	var count int64

	if len(ids) > 0 && req.Ids != "" {
		if err := DB.
			Model(&models.Exhibition{}).
			Preload("UserInfo").
			Find(&ex, ids).
			Count(&count).
			Error; err == nil {
			copier.Copy(&exInfo, &ex)
			return &types.ExhibitionInfoRes{Exhibitions: exInfo}, nil, msg
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
			Preload("UserInfo").
			Find(&ex).
			Offset(-1).
			Count(&count).
			Error; err == nil {
			copier.Copy(&exInfo, &ex)
			return &types.ExhibitionInfoRes{Exhibitions: exInfo, Count: int(count)}, nil, msg
		} else {
			return nil, err, msg
		}
	}
}
