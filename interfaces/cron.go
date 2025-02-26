package interfaces

import (
	"context"
	"github.com/robfig/cron/v3"
)

type Task struct {
	Name    string
	Spec    string
	Cmd     cron.FuncJob
	Job     cron.Job
	EntryID cron.EntryID
}

type Cron interface {
	// Stop 停止定时任务
	Stop() context.Context
	// Start 启动定时任务
	Start()
	// Find 查找任务
	Find(name string) (Task, cron.Entry, bool)
	// Remove 移除任务
	Remove(name string)
	// AddJob 添加任务
	AddJob(name, spec string, job cron.Job) error
	// AddFunc 添加任务
	AddFunc(name, spec string, cmd cron.FuncJob) error
}
