package evaluation

import (
	"context"
	"github.com/mraron/njudge/pkg/problems"
	"sync"
	"time"
)

type IgnoreStatusUpdate struct{}

func (i IgnoreStatusUpdate) UpdateStatus(ctx context.Context, testcase string, status problems.Status) error {
	return nil
}

func (i IgnoreStatusUpdate) Done(ctx context.Context) error {
	return nil
}

type Status struct {
	Testcase string
	Status   problems.Status
}

type ChanStatusUpdate struct {
	res chan Status
}

func NewChanStatusUpdate() (ChanStatusUpdate, <-chan Status) {
	u := ChanStatusUpdate{make(chan Status)}
	return u, u.res
}

func (c ChanStatusUpdate) UpdateStatus(_ context.Context, testcase string, status problems.Status) error {
	c.res <- Status{testcase, DeepCopyStatus(status)}
	return nil
}

func (c ChanStatusUpdate) Done(_ context.Context) error {
	close(c.res)
	return nil
}

type RateLimitStatusUpdate struct {
	su    problems.StatusUpdater
	every time.Duration

	mutex sync.Mutex
	last  time.Time
}

// NewRateLimitStatusUpdate drops the updates which come sooner than every after the last forwarded one.
func NewRateLimitStatusUpdate(innerUpdater problems.StatusUpdater, every time.Duration) *RateLimitStatusUpdate {
	return &RateLimitStatusUpdate{
		su:    innerUpdater,
		every: every,
	}
}

func (r *RateLimitStatusUpdate) allow() bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	now := time.Now()
	if !r.last.IsZero() && now.Sub(r.last) < r.every {
		return false
	}
	r.last = now
	return true
}

func (r *RateLimitStatusUpdate) UpdateStatus(ctx context.Context, testcase string, status problems.Status) error {
	if r.allow() {
		return r.su.UpdateStatus(ctx, testcase, status)
	}
	return nil
}

func (r *RateLimitStatusUpdate) Done(ctx context.Context) error {
	return r.su.Done(ctx)
}
