package main

import (
	"context"
	"fmt"

	"golang.org/x/time/rate"
)

// simpleLimiter is a wrapper around a rate.Limiter (Token Bucket Rate Limiter)
type simpleLimiter struct {
	limiter *rate.Limiter
}

// Limit tells wether or not to allow a certain request
// depending on the underlying limiter decision.
func (l *simpleLimiter) Limit(_ context.Context) error {
	if !l.limiter.Allow() {
		return fmt.Errorf("reached Rate-Limiting %v", l.limiter.Limit())
	}

	return nil
}
