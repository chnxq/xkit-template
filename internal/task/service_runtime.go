package task

import (
	"context"

	taskruntime "xkit-template-v01/internal/task/runtime"
)

type ServiceLogger interface {
	Errorf(format string, args ...any)
}

func RegisterServices(
	ctx context.Context,
	taskServiceScheduler *taskruntime.Scheduler,
	taskGroupServiceScheduler *taskruntime.Scheduler,
	logger ServiceLogger,
) (func(), error) {
	_ = ctx
	_ = taskServiceScheduler
	_ = taskGroupServiceScheduler
	_ = logger
	return func() {}, nil
}
