package runtime

type Scheduler struct{}

type Logger interface {
	Errorf(format string, args ...any)
}
