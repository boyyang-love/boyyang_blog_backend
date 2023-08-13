package exhibition

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"encoding/json"
	"strings"

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

func (l *CreateExhibitionLogic) CreateExhibition(req *types.CreateExhibitionReq) (resp *types.CreateExhibitionRes, err error, msg respx.SucMsg) {
	userId, _ := l.ctx.Value("Id").(json.Number).Int64()
	exhibition := models.Exhibition{
		Title:    req.Title,
		SubTitle: req.SubTitle,
		Des:      req.Des,
		Cover:    req.Cover,
		UserId:   uint(userId),
		Tags:     &req.Tags,
	}
	if err = l.svcCtx.DB.
		Model(&models.Exhibition{}).
		Create(&exhibition).Error; err != nil {
		return nil, err, msg
	} else {
		if strings.Trim(req.Tags, " ") != "" {
			// 创建 tags
			for _, tag := range strings.Split(req.Tags, ",") {
				tags := models.Tag{
					Name:    tag,
					ImageId: exhibition.Id,
					UserId:  uint(userId),
				}
				if err = l.svcCtx.DB.
					Model(&models.Tag{}).
					FirstOrCreate(&tags, &tags).
					Error; err != nil {
					return nil, err, msg
				}
			}
		}

		return &types.CreateExhibitionRes{Id: exhibition.Id},
			nil,
			respx.SucMsg{Msg: "图片上传成功"}
	}
}
