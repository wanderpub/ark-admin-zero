package main

import (
	"ark-admin-zero/app/core/cmd/jobs/internal/config"
	"ark-admin-zero/app/core/cmd/jobs/internal/handlers"
	"ark-admin-zero/app/core/cmd/jobs/internal/listen"
	"ark-admin-zero/app/core/cmd/jobs/internal/svc"
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/job-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	handlers.SetUp(ctx)
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range listen.Mqs(ctx) {
		serviceGroup.Add(mq)
	}
	serviceGroup.Start()
}
