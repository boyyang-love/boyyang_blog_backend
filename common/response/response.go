package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err == nil {
		body.Code = 1
		body.Msg = "OK"
		body.Data = resp
	} else {
		body.Code = 0
		body.Msg = err.Error()
	}

	httpx.OkJson(w, body)
}
