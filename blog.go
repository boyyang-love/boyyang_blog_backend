package main

import (
	"blog_server/internal/config"
	"blog_server/internal/handler"
	"blog_server/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

var configFile = flag.String("f", "etc/blog-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(
		c.RestConf,
		rest.WithCustomCors(
			nil,
			notAllowedFn,
			"http://localhost:3000",
			"http://boyyanglove.web3v.vip",
			"https://prod-2g5hif5wbec83baa-1301921121.tcloudbaseapp.com",
			"http://boyyang.3vkj.club",
			"http://111.67.195.4:8085",
			"http://111.67.195.4:8081",
		),
	)
	defer server.Stop()

	//hub := wss.NewHub()
	//go hub.Run()
	//
	//server.AddRoute(rest.Route{
	//	Method: http.MethodGet,
	//	Path:   "/wss",
	//	Handler: func(w http.ResponseWriter, r *http.Request) {
	//		wss.ServeWs(hub, w, r)
	//	},
	//})

	ctx := svc.NewServiceContext(c)
	go ctx.Hub.HubRun()
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func notAllowedFn(w http.ResponseWriter) {

}
