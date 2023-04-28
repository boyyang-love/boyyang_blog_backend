package respx

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type SucMsg struct {
	Msg string `json:"msg"`
}

func Response(w http.ResponseWriter, resp interface{}, err error, msg interface{}) {
	var body Body
	if err == nil {
		body.Code = 0
		body.Data = resp
		if msg.(SucMsg).Msg != "" {
			body.Msg = msg.(SucMsg).Msg
		} else {
			body.Msg = "ok"
		}
	} else {
		body.Code = 1
		body.Msg = err.Error()
	}

	httpx.OkJson(w, body)
}
