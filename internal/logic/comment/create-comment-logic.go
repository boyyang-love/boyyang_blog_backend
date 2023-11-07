package comment

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
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
		Content:   req.Content,
		ContentId: req.ContentId,
		UserId:    uint32(userId),
	}

	if err = l.svcCtx.DB.
		Model(&models.Comment{}).
		Create(&comment).
		Error; err != nil {
		return nil, err, msg
	}

	if err = l.updateComment(*req); err != nil {
		return nil, err, msg
	}

	return &types.CreateBlogCommentRes{Msg: "评论成功"}, err, respx.SucMsg{
		Msg: "评论成功",
	}
}

func (l *CreateCommentLogic) updateComment(req types.CreateBlogCommentReq) (err error) {
	DB := l.svcCtx.DB
	if req.Type == "image" {
		DB = DB.Model(&models.Exhibition{})
	}

	if req.Type == "article" {
		DB = DB.Model(&models.Article{})
	}

	if req.Type == "blog" {
		DB = DB.Model(&models.Blog{})
	}

	if err = DB.
		Where("uid = ?", req.ContentId).
		Update("comment", gorm.Expr("comment + ?", 1)).
		Error; err != nil {
		return err
	}

	return
}
