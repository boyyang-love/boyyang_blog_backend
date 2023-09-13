package main

import (
	"blog_server/internal/config"
	"blog_server/internal/handler"
	"blog_server/internal/svc"
	"blog_server/job"
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
			"https://prod-2g5hif5wbec83baa-1301921121.tcloudbaseapp.com",
		),
	)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)

	handler.RegisterHandlers(server, ctx)

	go job.CronRun(ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func notAllowedFn(http.ResponseWriter) {

}
