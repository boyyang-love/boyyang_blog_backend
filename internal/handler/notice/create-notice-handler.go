package notice

import (
	"blog_server/common/respx"
	"net/http"

	"blog_server/internal/logic/notice"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateNoticeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NoticeCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := notice.NewCreateNoticeLogic(r.Context(), svcCtx)
		err, msg := l.CreateNotice(&req)
		respx.Response(w, nil, err, msg)
	}
}
