package logic

import (
	"blog_server/common/response"
	"blog_server/models"
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"strings"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThumbsUpBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThumbsUpBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbsUpBlogLogic {
	return &ThumbsUpBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThumbsUpBlogLogic) ThumbsUpBlog(req *types.ThumbsUpBlogReq) (resp *types.ThumbsUpBlogRes, err error, msg response.SuccessMsg) {
	id, err := l.ctx.Value("Id").(json.Number).Int64()
	var blogInfo models.Blog
	DB := l.svcCtx.DB

	DB = DB.
		Model(&models.Blog{}).
		Where("id = ?", req.Id).
		Find(&blogInfo)

	if DB.Error == nil {
		thumbsIds := strings.Split(blogInfo.ThumbsUpList, ",")
		isThumbed := false
		for _, thumbsId := range thumbsIds {
			if thumbsId == string(id) {
				isThumbed = true
				break
			}
		}

		if isThumbed {
			return nil, nil, response.SuccessMsg{Msg: "您已经点赞过啦"}
		} else {
			thumbsIds = append(thumbsIds, "1")
			DB.Update("thumbs_up", gorm.Expr("thumbs_up + ?", 1))
			blogInfo.ThumbsUpList = strings.Join(thumbsIds, ",")
			DB.Save(&blogInfo)
			return nil, nil, response.SuccessMsg{Msg: "点赞成功"}
		}
	} else {
		return nil, err, msg
	}

}