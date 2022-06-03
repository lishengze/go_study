package loadgen

import (
	"context"
	"time"
)

type MyGenerator struct {
	ctx         context.Context
	cancel_func context.CancelFunc

	eclapsed time.Duration
	caller   struct{}

	resultCh   chan struct{}
	sourcePool struct{}
}
