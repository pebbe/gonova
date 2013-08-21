package gonova

/*
#include <libnova/solar.h>
*/
import "C"

import (
	"errors"
)

// Calculate the time of rise, set and transit for the Sun.
func GetSolarRST(JD Julian, lat, long float64) (rise, set, transit Julian, err error) {
	var rst C.struct_ln_rst_time
	observer := C.struct_ln_lnlat_posn{lng: C.double(long), lat: C.double(lat)}
	e := C.ln_get_solar_rst(C.double(JD), &observer, &rst)
	if e != 0 {
		err = errors.New("No sunrise/sunset on this day at this position")
		return
	}
	rise = Julian(rst.rise)
	set = Julian(rst.set)
	transit = Julian(rst.transit)
	return
}
