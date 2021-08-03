package mqtttimer

import (
	"encoding/json"
	"errors"
	"syscall"
	"time"

	"github.com/segmentio/ksuid"
)

func UID() string {
	return ksuid.New().String()
}

func pack(pkt NtpPackage) []byte {
	b, _ := json.Marshal(pkt)
	return b
}

func unpack(b []byte) NtpPackage {
	var p NtpPackage
	json.Unmarshal(b, &p)
	return p
}

func utc() time.Time {
	return time.Now().UTC()
}

func now() time.Time {
	return time.Now()
}

func SetSystemDate(newTime time.Time) error {
	tv := syscall.NsecToTimeval(newTime.UnixNano())
	if err := syscall.Settimeofday(&tv); err != nil {
		return errors.New("settimeofday: " + err.Error())
	}
	return nil
}

func Adjtime(delta int64) error {
	var (
		intval = syscall.NsecToTimeval(delta)
		oldval = syscall.Timeval{}
	)
	return syscall.Adjtime(&intval, &oldval)
}
