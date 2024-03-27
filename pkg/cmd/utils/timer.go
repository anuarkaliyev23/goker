package cmd

import "time"

func MeasureTime(fn func() error) (error, int64) {
	start := time.Now()
	err := fn()
	finish := time.Now()

	measuredTime := finish.Sub(start).Milliseconds()

	return err, measuredTime
}
