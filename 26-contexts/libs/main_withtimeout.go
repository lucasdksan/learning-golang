package libs

import (
	"context"
	"time"
)

func InitWithTimeout() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	printUntilCancel(ctx)
}
