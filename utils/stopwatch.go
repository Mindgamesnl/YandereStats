package utils

import "time"

var TimingLogs       = make(map[string]time.Duration)

type Stopwatch struct {
	start time.Time
	name string
}

func NewStopwatch(name string) Stopwatch {
	return Stopwatch{
		start: time.Now(),
		name:  name,
	}
}

func (sw Stopwatch) Stop() time.Duration {
	elapsed := time.Since(sw.start)
	TimingLogs[sw.name] = elapsed
	return elapsed
}