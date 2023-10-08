package notice

import (
	"blog_server/common/respx"
	"net/http"

	"blog_server/internal/logic/notice"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteNoticeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NoticeDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := notice.NewDeleteNoticeLogic(r.Context(), svcCtx)
		err, msg := l.DeleteNotice(&req)
		respx.Response(w, nil, err, msg)
	}
}
