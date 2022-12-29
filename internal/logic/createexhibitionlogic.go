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

type CreateExhibitionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateExhibitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateExhibitionLogic {
	return &CreateExhibitionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateExhibitionLogic) CreateExhibition(req *types.CreateExhibitionReq) (resp *types.CreateExhibitionRes, err error, msg response.SuccessMsg) {
	userId, _ := l.ctx.Value("Id").(json.Number).Int64()
	ex := models.Exhibition{
		Title:    req.Title,
		SubTitle: req.SubTitle,
		Des:      req.Des,
		Cover:    req.Cover,
		UserId:   uint(userId),
	}
	res := l.svcCtx.DB.
		Model(&models.Exhibition{}).
		Create(&ex)

	if res.Error == nil {
		return &types.CreateExhibitionRes{Id: int(ex.Id)}, nil, msg
	} else {
		return nil, res.Error, msg
	}
}
