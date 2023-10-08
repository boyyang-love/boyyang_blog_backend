package notice

import (
	"blog_server/common/respx"
	"blog_server/internal/logic/notice"
	"net/http"

	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func InfoNoticeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NoticeInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := notice.NewInfoNoticeLogic(r.Context(), svcCtx)
		resp, err, msg := l.InfoNotice(&req)

		respx.Response(w, resp, err, msg)
	}
}
