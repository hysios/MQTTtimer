package mqtttimer

import (
	"errors"
	"github.com/jcollie/w32"
	"time"
)

func SetSystemDate(newTime time.Time) error {
	w32.SetSystemTime(w32.SYSTEMTIME{
		Year:         uint16(newTime.Year()),
		Month:        uint16(newTime.Month()),
		Day:          uint16(newTime.Day()),
		Hour:         uint16(newTime.Hour()),
		Minute:       uint16(newTime.Minute()),
		Second:       uint16(newTime.Second()),
		Milliseconds: 0,
	})
	return nil
}

func Adjtime(delta int64) error {

	return errors.New("nonimplement")
}
