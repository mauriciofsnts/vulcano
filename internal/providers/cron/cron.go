package cron

import (
	"time"

	"github.com/go-co-op/gocron"
)

type Cron interface {
	AddJob(spec string, job func()) error
	Start()
	Stop()
	List() []*gocron.Job
}

type gocronImpl struct {
	scheduler *gocron.Scheduler
}

func (c *gocronImpl) AddJob(spec string, job func()) error {
	_, err := c.scheduler.Cron(spec).Do(job)
	return err
}

func (c *gocronImpl) Start() {
	c.scheduler.StartAsync()
}

func (c *gocronImpl) Stop() {
	c.scheduler.Stop()
}

func (c *gocronImpl) List() []*gocron.Job {
	return c.scheduler.Jobs()
}

func New() Cron {
	s := gocron.NewScheduler(time.UTC)

	return &gocronImpl{
		scheduler: s,
	}
}
