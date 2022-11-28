package listen

import (
	"ark-admin-zero/app/core/cmd/jobs/internal/config"
	"ark-admin-zero/app/core/cmd/jobs/internal/svc"
	"context"

	"github.com/zeromicro/go-zero/core/service"
)

// 返回所有消费者
func Mqs(svcCtx *svc.ServiceContext) []service.Service {
	ctx := context.Background()
	var services []service.Service
	//asynq 定时任务/延迟任务
	services = append(services, AsynqMqs(svcCtx.Config, ctx, svcCtx)...)
	return services
}

// 定时任务/延迟任务
func AsynqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {
	return []service.Service{
		//监听【队列/延迟队列】
		NewAsynqTask(ctx, svcContext),
		//监听【定时任务】
		NewAsynqScheduler(ctx, svcContext),
	}
}
