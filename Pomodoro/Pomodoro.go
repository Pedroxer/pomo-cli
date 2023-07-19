package pomodoro

import (
	"time"
)

type Tomato struct {
	SubjectName string        // name of subject
	Timer       time.Duration // how long tomato is gonna be
	Done        bool
}

func Constructor(name string, dur string) Tomato {
	t, err := time.ParseDuration(dur)
	if err != nil {
		return Tomato{}
	}
	return Tomato{
		SubjectName: name,
		Timer:       t,
		Done:        false,
	}
}

