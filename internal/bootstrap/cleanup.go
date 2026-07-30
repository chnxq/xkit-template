package bootstrap

import (
	"context"
	"errors"
	"sync"
)

type CleanupFunc func(context.Context) error

type CleanupStack struct {
	mu    sync.Mutex
	funcs []CleanupFunc
	run   bool
}

func (s *CleanupStack) Add(cleanup func()) {
	if cleanup == nil {
		return
	}
	s.AddContext(func(context.Context) error {
		cleanup()
		return nil
	})
}

func (s *CleanupStack) AddContext(cleanup CleanupFunc) {
	if cleanup == nil {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.run {
		return
	}
	s.funcs = append(s.funcs, cleanup)
}

func (s *CleanupStack) Run(ctx context.Context) error {
	s.mu.Lock()
	if s.run {
		s.mu.Unlock()
		return nil
	}
	s.run = true
	funcs := s.funcs
	s.funcs = nil
	s.mu.Unlock()

	var joined error
	for i := len(funcs) - 1; i >= 0; i-- {
		joined = errors.Join(joined, funcs[i](ctx))
	}
	return joined
}
