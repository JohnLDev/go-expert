package utils

import (
	"time"
)

type Timer struct {
	timer  time.Time
	Result int64 // milliseconds
}

func (t *Timer) StopTimer() *Timer {
	t.Result = time.Since(t.timer).Milliseconds()
	return t
}

func StartTimer() *Timer {
	return &Timer{
		timer: time.Now(),
	}
}
