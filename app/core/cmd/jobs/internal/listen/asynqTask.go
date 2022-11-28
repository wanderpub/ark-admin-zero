package listen

import (
	"ark-admin-zero/app/core/cmd/jobs/internal/svc"
	"ark-admin-zero/app/core/model"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

/*
*
监听关闭订单
*/
type AsynqTask struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAsynqTask(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqTask {
	return &AsynqTask{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AsynqTask) Start() {

	fmt.Println("AsynqTask start ")

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: l.svcCtx.Config.Redis.Host, Password: l.svcCtx.Config.Redis.Pass},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			IsFailure:      func(err error) bool { return !IsRateLimitError(err) },
			RetryDelayFunc: retryDelay, //限速
		},
	)

	mux := asynq.NewServeMux()
	mux.Use(loggingMiddleware)
	mux.HandleFunc(TypeSend, l.handleSendProcess)
	mux.HandleFunc(TypeEmail, l.handleSendProcess)
	mux.HandleFunc(TypeSms, l.handleSendProcess)
	mux.HandleFunc(TypeAliSms, l.handleSendProcess)
	mux.HandleFunc(TypeTencentSms, l.handleSendProcess)

	// mux1 := asynq.NewServeMux()
	// mux1.Use(loggingMiddleware)
	mux.HandleFunc(TypeOffiaccount, l.handleSendProcess)
	mux.HandleFunc(TypeMiniprogram, l.handleSendProcess)
	mux.HandleFunc(TypeEnterpriseWechat, l.handleSendProcess)
	mux.HandleFunc(TypeDingRobot, l.handleSendProcess)
	mux.HandleFunc(TypeDingWork, l.handleSendProcess)
	mux.HandleFunc(TypeTencentPush, l.handleSendProcess)
	mux.HandleFunc("crontab", handleCrontab)
	mux.HandleFunc("bar", handleCrontab)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func handleCrontab(ctx context.Context, t *asynq.Task) error {
	var job model.SysTaskJob
	jsonx.Unmarshal(t.Payload(), &job)
	logx.Infof("消息内容 %v, now is %v\n", job, time.Now())
	return nil
}

// Using Middleware
func loggingMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		location, _ := time.LoadLocation("Asia/Shanghai")
		fmt.Println("处理时间：", time.Now().In(location).Format("2006-01-02 15:04:05"))
		start := time.Now().In(location)
		log.Printf("Start processing %q", t.Type())
		err := h.ProcessTask(ctx, t)
		if err != nil {
			return err
		}
		log.Printf("Finished processing %q: Elapsed Time = %v", t.Type(), time.Since(start))
		return nil
	})
}
func (l *AsynqTask) Stop() {
	fmt.Println("AsynqTask stop")
}
