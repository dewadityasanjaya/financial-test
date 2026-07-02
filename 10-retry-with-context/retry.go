package retry

import (
	"context"
	"time"
)

type Operation func(ctx context.Context) (string, error)

func Retry(ctx context.Context, attempts int, delay time.Duration, operation Operation) (string, error) {
	// TODO: implement.
	return "", nil
}

