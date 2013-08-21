package gonova

/*
#include <libnova/julian_day.h>
*/
import "C"

import (
	"time"
)

type Julian C.double

func TimeToJulian(t time.Time) Julian {
	t1 := t.UTC()
	lnd := C.struct_ln_date{
		years:   C.int(t1.Year()),
		months:  C.int(t1.Month()),
		days:    C.int(t1.Day()),
		hours:   C.int(t1.Hour()),
		minutes: C.int(t1.Minute()),
		seconds: C.double(t1.Second()) + C.double(t1.Nanosecond())*1e-9,
	}
	return Julian(C.ln_get_julian_day(&lnd))
}

func JulianToTime(JD Julian) time.Time {
	lnd := C.struct_ln_date{}
	C.ln_get_date(C.double(JD), &lnd)
	sec := int(lnd.seconds)
	nano := int(1e9 * (float64(lnd.seconds) - float64(sec)))
	return time.Date(
		int(lnd.years),
		time.Month(lnd.months),
		int(lnd.days),
		int(lnd.hours),
		int(lnd.minutes),
		sec,
		nano,
		time.UTC,
	).Local()
}
