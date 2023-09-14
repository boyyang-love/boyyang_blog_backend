package user

import (
	"blog_server/common/respx"
	"blog_server/internal/logic/user"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UpdateUserPasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserPasswordReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUpdateUserPasswordLogic(r.Context(), svcCtx)
		err, msg := l.UpdateUserPassword(&req)

		respx.Response(w, nil, err, msg)
	}
}
