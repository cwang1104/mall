package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"mall/apps/app/api/internal/config"
	"mall/apps/app/api/internal/handler"
	"mall/apps/app/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api-api.yaml", "the config file")

func main() {
	flag.Parse()
	logx.DisableStat()
	logx.MustSetup(logx.LogConf{
		Encoding: "plain",
	})
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
