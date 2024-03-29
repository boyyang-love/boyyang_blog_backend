package like

import (
	"blog_server/common/respx"
	"blog_server/internal/logic/like"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func LikeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddLikesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := like.NewLikeLogic(r.Context(), svcCtx)
		err, msg := l.Like(&req)
		respx.Response(w, nil, err, msg)
	}
}
