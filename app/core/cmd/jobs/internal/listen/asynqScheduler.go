package listen

import (
	"ark-admin-zero/app/core/cmd/jobs/internal/svc"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/jsonx"
	"gopkg.in/yaml.v2"
)

type AsynqScheduler struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAsynqScheduler(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqScheduler {
	return &AsynqScheduler{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AsynqScheduler) Start() {
	fmt.Println("AsynqScheduler start ")
	// provider := &FileBasedConfigProvider{filename: "etc/periodic_task_config.yml"}
	provider := &DbBasedConfigProvider{ctx: l.ctx, svcCtx: l.svcCtx}

	location, _ := time.LoadLocation("Asia/Shanghai")
	schedulerOpts := &asynq.SchedulerOpts{
		Location: location,
	}
	mgr, err := asynq.NewPeriodicTaskManager(
		asynq.PeriodicTaskManagerOpts{
			RedisConnOpt: asynq.RedisClientOpt{
				Addr:     l.svcCtx.Config.Redis.Host,
				Password: l.svcCtx.Config.Redis.Pass,
			},
			PeriodicTaskConfigProvider: provider,         // this provider object is the interface to your config source
			SyncInterval:               10 * time.Second, // this field specifies how often sync should happen
			SchedulerOpts:              schedulerOpts,
		})
	if err != nil {
		log.Fatal(err)
	}
	if err := mgr.Run(); err != nil {
		log.Fatal(err)
	}
}

func (l *AsynqScheduler) Stop() {

	fmt.Println("AsynqScheduler stop ")
}

type DbBasedConfigProvider struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (p *DbBasedConfigProvider) GetConfigs() ([]*asynq.PeriodicTaskConfig, error) {
	builder := p.svcCtx.TaskJobModel.SelectBuilder("", "").Where("status=?", 1)
	data, err := p.svcCtx.TaskJobModel.GetAll(p.ctx, builder)
	if err != nil {
		return nil, err
	}
	var configs []*asynq.PeriodicTaskConfig
	for _, cfg := range data {
		typeName := "crontab"
		payload, _ := jsonx.Marshal(cfg)
		t := asynq.NewTask(typeName, payload)
		configs = append(configs, &asynq.PeriodicTaskConfig{Cronspec: cfg.CronExpression, Task: t})
	}
	return configs, nil
}

// FileBasedConfigProvider implements asynq.PeriodicTaskConfigProvider interface.
type FileBasedConfigProvider struct {
	filename string
}

// Parses the yaml file and return a list of PeriodicTaskConfigs.
func (p *FileBasedConfigProvider) GetConfigs() ([]*asynq.PeriodicTaskConfig, error) {
	data, err := os.ReadFile(p.filename)
	if err != nil {
		return nil, err
	}
	var c PeriodicTaskConfigContainer
	if err := yaml.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	var configs []*asynq.PeriodicTaskConfig
	for _, cfg := range c.Configs {
		configs = append(configs, &asynq.PeriodicTaskConfig{Cronspec: cfg.Cronspec, Task: asynq.NewTask(cfg.TaskType, nil)})
	}
	return configs, nil
}

type PeriodicTaskConfigContainer struct {
	Configs []*Config `yaml:"configs"`
}

type Config struct {
	Cronspec string `yaml:"cronspec"`
	TaskType string `yaml:"task_type"`
}
