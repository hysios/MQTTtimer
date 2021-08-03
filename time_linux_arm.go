package mqtttimer

import "golang.org/x/sys/unix"

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
