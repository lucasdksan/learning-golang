package libs

import (
	"context"
	"fmt"
)

func printUntilCancelText(ctx context.Context) {
	fmt.Println(ctx.Value("testKey"))
}
