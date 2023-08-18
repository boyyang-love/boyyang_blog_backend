package exhibition

import (
	"blog_server/common/respx"
	"context"

	"blog_server/internal/svc"
	"blog_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUploadLogic {
	return &DelUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelUploadLogic) DelUpload(req *types.DelUploadReq) (err error, msg respx.SucMsg) {
	_, err = l.svcCtx.Client.Object.Delete(context.Background(), req.Key)
	if err != nil {
		return err, msg
	}
	return nil, respx.SucMsg{Msg: "图片删除成功"}
}
