package retry

import (
	"context"
	"time"
)

type Operation func(ctx context.Context) (string, error)

func Retry(ctx context.Context, attempts int, delay time.Duration, operation Operation) (string, error) {
	var lastErr error

	for i := 0; i < attempts; i++ {
		if ctx.Err() != nil {
			return "", ctx.Err()
		}

		result, err := operation(ctx)
		if err == nil {
			return result, nil
		}
		lastErr = err

		if i == attempts-1 {
			break
		}

		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-time.After(delay):
		}
	}

	return "", lastErr
}
