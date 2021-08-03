package mqtttimer

var (
	ADJ_OFFSET    = 0x0001 /* time offset */
	ADJ_FREQUENCY = 0x0002 /* frequency offset */
	ADJ_MAXERROR  = 0x0004 /* maximum time error */
	ADJ_ESTERROR  = 0x0008 /* estimated time error */
	ADJ_STATUS    = 0x0010 /* clock status */
	ADJ_TIMECONST = 0x0020 /* pll time constant */
	ADJ_TAI       = 0x0080 /* set TAI offset */
	ADJ_SETOFFSET = 0x0100 /* add 'time' to current time */
	ADJ_MICRO     = 0x1000 /* select microsecond resolution */
	ADJ_NANO      = 0x2000 /* select nanosecond resolution */
	ADJ_TICK      = 0x4000 /* tick value */
)
