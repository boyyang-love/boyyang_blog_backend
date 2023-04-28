package logic

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlogInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBlogInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlogInfoLogic {
	return &BlogInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BlogInfoLogic) BlogInfo(req *types.BlogInfoReq) (resp *types.BlogInfoRes, err error, msg respx.SucMsg) {
	DB := l.svcCtx.DB
	var count int64
	var blogInfo []models.Blog

	userId, _ := l.ctx.Value("Id").(json.Number).Int64()

	ids := strings.Split(req.Ids, ",")

	if len(ids) > 0 && req.Ids != "" {
		res := DB.
			Model(&models.Blog{}).
			Preload("UserInfo").
			Preload("Comments").
			Preload("Comments.UserInfo").
			Where("user_id", userId).
			Find(&blogInfo, ids).
			Count(&count)
		err = res.Error
	} else {
		// 分页
		if req.Page != "" || req.Limit != "" {
			page, _ := strconv.Atoi(req.Page)
			limit, _ := strconv.Atoi(req.Limit)
			DB = DB.
				Offset((page - 1) * limit).
				Limit(limit)
		}

		res := DB.
			Model(&models.Blog{}).
			Preload("UserInfo").
			Preload("Comments").
			Preload("Comments.UserInfo").
			Where("user_id", userId).
			Find(&blogInfo).
			Offset(-1).
			Count(&count)
		err = res.Error
	}

	if err == nil {
		var info []types.BlogInfo
		_ = copier.Copy(&info, &blogInfo)
		return &types.BlogInfoRes{BlogInfo: info, Count: int(count)}, nil, msg
	} else {
		return nil, err, msg
	}
}
