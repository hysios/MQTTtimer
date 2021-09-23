//go:build !linux
// +build !linux

package mqtttimer

import (
	"time"
)

type rtc struct {
}

// NewRTC opens a real-time clock device.
func NewRTC(dev string) (*rtc, error) {

	return &rtc{}, nil
}

// Close closes a real-time clock device.
func (c *rtc) Close() (err error) {
	return nil
}

// Epoch returns the real-time clock's epoch.
func (c *rtc) Epoch() (epoch uint, err error) {

	return 0, nil
}

// SetEpoch sets the real-time clock's epoch.
func (c *rtc) SetEpoch(epoch uint) (err error) {

	return nil
}

// Time returns the specified real-time clock device time.
func (c *rtc) Time() (t time.Time, err error) {

	return time.Now(), nil
}

// SetTime sets the time for the specified real-time clock device.
func (c *rtc) SetTime(t time.Time) (err error) {

	return nil
}

// Frequency returns the periodic interrupt frequency.
func (c *rtc) Frequency() (frequency uint, err error) {

	return 0, nil
}

// SetFrequency sets the frequency of the real-time clock's periodic interrupt.
func (c *rtc) SetFrequency(frequency uint) (err error) {

	return nil
}

// SetPeriodicInterrupt enables or disables the real-time clock's periodic interrupts.
func (c *rtc) SetPeriodicInterrupt(enable bool) (err error) {

	return nil
}

// SetAlarmInterrupt enables or disables the real-time clock's alarm interrupt.
func (c *rtc) SetAlarmInterrupt(enable bool) (err error) {

	return nil
}

// SetUpdateInterrupt enables or disables the real-time clock's update interrupt.
func (c *rtc) SetUpdateInterrupt(enable bool) (err error) {

	return nil
}

// Alarm returns the real-time clock's alarm time.
func (c *rtc) Alarm() (t time.Time, err error) {

	return time.Time{}, nil
}

// SetAlarm sets the real-time clock's alarm time.
func (c *rtc) SetAlarm(t time.Time) (err error) {

	return nil
}

// WakeAlarm returns the real-time clock's wake alarm time.
func (c *rtc) WakeAlarm() (enabled bool, pending bool, t time.Time, err error) {
	enabled = false
	pending = false
	t = time.Time{}
	err = nil
	return
}

// SetWakeAlarm sets the real-time clock's wake alarm time.
func (c *rtc) SetWakeAlarm(t time.Time) (err error) {

	return nil
}

// CancelWakeAlarm cancels the real-time clock's wake alarm.
func (c *rtc) CancelWakeAlarm() (err error) {

	return nil
}
