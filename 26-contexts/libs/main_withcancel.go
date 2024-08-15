package libs

import (
	"context"
	"time"
)

func InitWithCancel() {
	ctx, cancel := context.WithCancel(
		context.Background(),
	)

	go printUntilCancel(ctx)

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
