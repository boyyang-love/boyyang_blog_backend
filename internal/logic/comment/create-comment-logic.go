package comment

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateBlogCommentReq) (resp *types.CreateBlogCommentRes, err error, msg respx.SucMsg) {

	userId, err := l.ctx.Value("Uid").(json.Number).Int64()
	comment := models.Comment{
		Content: req.Content,
		BlogId:  uint32(req.BlogId),
		UserId:  uint32(userId),
	}
	res := l.svcCtx.DB.Model(&models.Comment{}).Create(&comment)
	if res.Error == nil {
		return &types.CreateBlogCommentRes{Msg: "评论成功"}, nil, msg
	} else {
		return nil, res.Error, msg
	}
}
