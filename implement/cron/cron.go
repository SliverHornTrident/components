package cron

import (
	"context"
	"fmt"
	"github.com/SliverHornTrident/components/interfaces"
	"github.com/pkg/errors"
)

type Cron struct {
	task map[string]interfaces.Task
	cron *cron.Cron
}

func NewCron() *Cron {
	c := cron.New(cron.WithSeconds())
	task := make(map[string]interfaces.Task)
	return &Cron{task: task, cron: c}
}

func (c *Cron) Start() {
	for name, task := range c.task {
		fmt.Printf("[components][cron][name:%s][entryID:%d][spec:%s] Initialization success.\n", name, task.EntryID, task.Spec)
	}
	c.cron.Start()
}

func (c *Cron) Stop() context.Context {
	return c.cron.Stop()
}

func (c *Cron) Task() map[string]interfaces.Task {
	return c.task
}

func (c *Cron) Find(name string) (interfaces.Task, cron.Entry, bool) {
	value, ok := c.task[name]
	entry := c.cron.Entry(value.EntryID)
	return value, entry, ok
}

func (c *Cron) Remove(name string) {
	task, ok := c.task[name]
	if ok {
		c.cron.Remove(task.EntryID)
		delete(c.task, name)
	}
}

func (c *Cron) AddJob(name, spec string, job cron.Job) error {
	_, ok := c.task[name]
	if ok {
		return errors.Errorf("[cron][name:%s]任务已存在!", name)
	}
	entryID, err := c.cron.AddJob(spec, job)
	if err != nil {
		return err
	}
	task := interfaces.Task{
		Name:    name,
		Spec:    spec,
		Job:     job,
		EntryID: entryID,
	}
	c.task[name] = task
	return nil
}

func (c *Cron) AddFunc(name, spec string, cmd cron.FuncJob) error {
	_, ok := c.task[name]
	if ok {
		return errors.Errorf("[cron][name:%s]任务已存在!", name)
	}
	entryID, err := c.cron.AddFunc(spec, cmd)
	if err != nil {
		return err
	}
	task := interfaces.Task{
		Name:    name,
		Spec:    spec,
		Cmd:     cmd,
		EntryID: entryID,
	}
	c.task[name] = task
	return nil
}
