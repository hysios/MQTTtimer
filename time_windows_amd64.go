package mqtttimer

import (
	"errors"
	"github.com/jcollie/w32"
	"time"
)

func SetSystemDate(newTime time.Time) error {
	return w32.SetSystemTime(w32.SystemTime{
		Year:         newTime.Year(),
		Month:        newTime.Month(),
		Day:          newTime.Day(),
		Hour:         newTime.Hour(),
		Minute:       newTime.Minute(),
		Second:       newTime.Second(),
		Milliseconds: 0,
	})
}

func Adjtime(delta int64) error {

	return errors.New("nonimplement")
}
