package libs

import (
	"context"
	"time"
)

func InitWithDeadline() {
	ctx, cancel := context.WithDeadline(
		context.Background(),
		time.Now().Add(10*time.Second),
	)

	defer cancel()

	printUntilCancel(ctx)
}
