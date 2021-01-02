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

// New returns a new scheduler containing a wait group and process control contexts.
func New() *Scheduler {
	return &Scheduler{
		wg:            new(sync.WaitGroup),
		cancellations: make([]context.CancelFunc, 0),
	}
}

// process runs the provided job on the time interval.
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

// Add spawns a goroutine to run the job at the provided interval.
func (s *Scheduler) Add(ctx context.Context, job Job, interval time.Duration) {
	ctx, cancel := context.WithCancel(ctx)
	s.cancellations = append(s.cancellations, cancel)
	s.wg.Add(1)
	go s.process(ctx, job, interval)
}

// Stop cancels all running jobs of a scheduled process.
func (s *Scheduler) Stop() {
	for _, cancel := range s.cancellations {
		cancel()
	}
	s.wg.Wait()
}
