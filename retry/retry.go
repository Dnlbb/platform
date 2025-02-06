package retry

import (
	"fmt"
	"time"
)

var ErrMaxAttempts = fmt.Errorf("maximum number attempts has been reached")

func DoRetry(fn func() error, needRetry func(error) bool, delay time.Duration, maxDelay time.Duration) error {
	tm := time.Now().Add(maxDelay)
	var attempt int

	for tm.After(time.Now()) {
		var err error
		if err = fn(); err == nil {
			return nil
		}

		if needRetry(err) {
			attempt++
			time.Sleep(time.Duration(attempt) * delay)
			continue
		} else {
			return err
		}
	}

	return ErrMaxAttempts
}
