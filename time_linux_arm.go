package mqtttimer

import (
	"errors"
	"golang.org/x/sys/unix"
	"syscall"
	"time"
)

func Adjtime(delta int64) error {
	var buf unix.Timex
	buf.Modes = uint32(ADJ_OFFSET | ADJ_NANO)
	buf.Offset = int32(delta / 1000)

	sts, err := unix.Adjtimex(&buf)
	if sts == 0 {
		return nil
	}
	return err
}

func SetSystemDate(newTime time.Time) error {
	tv := syscall.NsecToTimeval(newTime.UnixNano())
	if err := syscall.Settimeofday(&tv); err != nil {
		return errors.New("settimeofday: " + err.Error())
	}
	return nil
}
