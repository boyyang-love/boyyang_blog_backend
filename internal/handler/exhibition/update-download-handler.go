package exhibition

import (
	"blog_server/common/respx"
	"net/http"

	"blog_server/internal/logic/exhibition"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateDownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateDownloadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := exhibition.NewUpdateDownloadLogic(r.Context(), svcCtx)
		err, msg := l.UpdateDownload(&req)

		respx.Response(w, nil, err, msg)
	}
}
