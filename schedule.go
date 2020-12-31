package schedule

import (
	"context"
	"sync"
	"time"
)

// Scheduler type defines a structure containing a WaitGroup and
// cancellation functions.
type Scheduler struct {
	wg            *sync.WaitGroup
	cancellations []context.CancelFunc
}

// Job type is any function with a context.
type Job func(ctx context.Context)

// New returns a new scheduler with a wait group and process control context.
func New() *Scheduler {
	return &Scheduler{
		wg:            new(sync.WaitGroup),
		cancellations: make([]context.CancelFunc, 0),
	}
}

// process runs provided job on the time interval provided.
func (s *Scheduler) process(ctx context.Context, j Job, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			j(ctx)
		case <-ctx.Done():
			s.wg.Done()
			return
		}
	}
}

// Add starts a goroutine containing the provided job to run on provided interval.
func (s *Scheduler) Add(ctx context.Context, j Job, interval time.Duration) {
	ctx, cancel := context.WithCancel(ctx)
	s.cancellations = append(s.cancellations, cancel)
	s.wg.Add(1)
	go s.process(ctx, j, interval)
}

// Stop cancels all running jobs of a scheduled process.
func (s *Scheduler) Stop() {
	for _, cancel := range s.cancellations {
		cancel()
	}
	s.wg.Wait()
}
