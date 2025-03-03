package redis

import (
	"context"
	"time"

	"biyard.co/common/logger"
	"github.com/go-redsync/redsync/v4"
)

type Lock struct {
	cli *redsync.Mutex
	ctx context.Context
}

func NewLock(ctx context.Context, cli *redsync.Mutex) *Lock {
	return &Lock{cli: cli, ctx: ctx}
}

func (r *Lock) Lock() {
	for err := r.cli.Lock(); err != nil; {
		logger.Errorf(r.ctx, "error locking(redis): %v", err)
		time.Sleep(1 * time.Second)
		err = r.cli.Lock()
	}

	return
}

func (r *Lock) Unlock() {
	for status, err := r.cli.Unlock(); !status && err != nil; {
		logger.Errorf(r.ctx, "error unlocking(redis): %v", err)
		time.Sleep(1 * time.Second)
		status, err = r.cli.Unlock()
	}

	return
}
