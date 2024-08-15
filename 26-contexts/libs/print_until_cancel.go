package libs

import (
	"context"
	"fmt"
	"time"
)

func printUntilCancel(ctx context.Context) {
	count := 0

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cancel signal received, existing")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Printf("Printing until cancel, number: %d \n", count)
			count += 1
		}

	}
}
