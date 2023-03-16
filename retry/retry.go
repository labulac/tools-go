package retry

import (
	"time"
)

func Retry(tries int, delay time.Duration, fn func() error) error {
	var err error

	for i := 0; i < tries; i++ {
		if err = fn(); err == nil {
			return nil
		}

		time.Sleep(delay)
	}

	return err
}
