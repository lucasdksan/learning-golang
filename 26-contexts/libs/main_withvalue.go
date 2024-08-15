package libs

import (
	"context"
)

func InitWithValue() {
	ctx := context.WithValue(
		context.Background(),
		"testKey",
		"testvalue",
	)

	printUntilCancelText(ctx)
}
