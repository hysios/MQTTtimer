package mqtttimer

import "syscall"

func Adjtime(delta int64) error {
	var (
		intval = syscall.NsecToTimeval(delta)
		oldval = syscall.Timeval{}
	)
	return syscall.Adjtime(&intval, &oldval)
}
