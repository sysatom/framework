package types

import (
	"context"
	"time"
)

type Context struct {
	// ctx is the context
	ctx context.Context
	// cancel function
	cancel context.CancelFunc

	// Sender's UserId as string.
	AsUser Uid
}

func (c *Context) Context() context.Context {
	if c.ctx == nil {
		return context.Background()
	}
	return c.ctx
}

func (c *Context) SetTimeout(timeout time.Duration) {
	// If there is an existing cancel function, call it first to avoid resource leaks
	if c.cancel != nil {
		c.cancel()
	}

	// If the context is nil, create a new context with timeout
	if c.ctx == nil {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		c.ctx = ctx
		c.cancel = cancel
		return
	}

	// If the context already exists but has no deadline, create a new context with timeout based on the existing context
	if _, ok := c.ctx.Deadline(); !ok {
		ctx, cancel := context.WithTimeout(c.ctx, timeout)
		c.ctx = ctx
		c.cancel = cancel
		return
	}
}

func (c *Context) Cancel() context.CancelFunc {
	return c.cancel
}
