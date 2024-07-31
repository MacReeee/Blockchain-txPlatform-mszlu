package task

import (
	"github.com/go-co-op/gocron"
	"jobcenter/internal/kline"
	"jobcenter/internal/svc"
	"time"
)

type Task struct {
	s   *gocron.Scheduler
	ctx *svc.ServiceContext
}

func NewTask(ctx *svc.ServiceContext) *Task {
	return &Task{
		s:   gocron.NewScheduler(time.UTC),
		ctx: ctx,
	}
}

func (t *Task) Run() {
	t.s.Every(1).Minute().Do(func() {
		kline.NewKline(t.ctx.Config.Okx, t.ctx.MongoClient).Do("1m")
	})
	t.s.Every(1).Hour().Do(func() {
		kline.NewKline(t.ctx.Config.Okx, t.ctx.MongoClient).Do("1H")
	})
}

func (t *Task) StartBlocking() {
	t.s.StartBlocking()
}

func (t *Task) Stop() {
	t.s.Stop()
}
