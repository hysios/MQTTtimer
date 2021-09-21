package mqtttimer

import (
	"errors"
	"syscall"
	"time"
)

func Adjtime(delta int64) error {
	var (
		intval = syscall.NsecToTimeval(delta)
		oldval = syscall.Timeval{}
	)
	return syscall.Adjtime(&intval, &oldval)
}

func SetSystemDate(newTime time.Time) error {
	tv := syscall.NsecToTimeval(newTime.UnixNano())
	if err := syscall.Settimeofday(&tv); err != nil {
		return errors.New("settimeofday: " + err.Error())
	}
	return nil
}
