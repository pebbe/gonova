package gonova

/*
#include <libnova/solar.h>

const double
	C_SOLAR_STANDART_HORIZON      = LN_SOLAR_STANDART_HORIZON,
	C_SOLAR_CIVIL_HORIZON         = LN_SOLAR_CIVIL_HORIZON,
	C_SOLAR_NAUTIC_HORIZON        = LN_SOLAR_NAUTIC_HORIZON,
	C_SOLAR_ASTRONOMICAL_HORIZON  = LN_SOLAR_ASTRONOMICAL_HORIZON;
*/
import "C"

import (
	"errors"
	"math"
)

var (
	SOLAR_STANDART_HORIZON     = float64(C.C_SOLAR_STANDART_HORIZON)
	SOLAR_CIVIL_HORIZON        = float64(C.C_SOLAR_CIVIL_HORIZON)
	SOLAR_NAUTIC_HORIZON       = float64(C.C_SOLAR_NAUTIC_HORIZON)
	SOLAR_ASTRONOMICAL_HORIZON = float64(C.C_SOLAR_ASTRONOMICAL_HORIZON)
)

// int LIBNOVA_EXPORT ln_get_solar_rst_horizon (double JD, struct ln_lnlat_posn * observer, double horizon, struct ln_rst_time * rst);

// Calculate the time of rise, set and transit for the Sun.
// Return solar rise/set time over local horizon (specified in degrees).
func GetSolarRSTHorizon(JD Julian, long, lat, horizon float64) (rise, set, transit Julian, err error) {
	var rst C.struct_ln_rst_time
	observer := C.struct_ln_lnlat_posn{lng: C.double(long), lat: C.double(lat)}
	e := C.ln_get_solar_rst_horizon(C.double(JD), &observer, C.double(horizon), &rst)
	if e != 0 {
		err = errors.New("No sunrise/sunset on this day at this position")
		return
	}
	rise = Julian(rst.rise)
	set = Julian(rst.set)
	transit = Julian(rst.transit)
	return
}

// int LIBNOVA_EXPORT ln_get_solar_rst (double JD, struct ln_lnlat_posn * observer, struct ln_rst_time * rst);

// Calculate the time of rise, set and transit for the Sun.
// Calls GetSolarRSTHorizon with horizon set to SOLAR_STANDART_HORIZON.
func GetSolarRST(JD Julian, long, lat float64) (rise, set, transit Julian, err error) {
	return GetSolarRSTHorizon(JD, long, lat, SOLAR_STANDART_HORIZON)
}

// void LIBNOVA_EXPORT ln_get_solar_geom_coords (double JD, struct ln_helio_posn * position);

// Calculate solar geometric coordinates.
func GetSolarGeomCoords(JD Julian) (l, b, r float64) {
	var posn C.struct_ln_helio_posn
	C.ln_get_solar_geom_coords(C.double(JD), &posn)
	l, b, r = float64(posn.L), float64(posn.B), float64(posn.R)
	return
}

// void LIBNOVA_EXPORT ln_get_solar_equ_coords (double JD, struct ln_equ_posn * position);

// Calculate apparent equatorial coordinates.
func GetSolarEquCoords(JD Julian) (ra, dec float64) {
	var posn C.struct_ln_equ_posn
	C.ln_get_solar_equ_coords(C.double(JD), &posn)
	ra, dec = float64(posn.ra), float64(posn.dec)
	return
}

// void LIBNOVA_EXPORT ln_get_solar_ecl_coords (double JD, struct ln_lnlat_posn * position);

// Calculate apparent ecliptical coordinates.
func GetSolarEclCoords(JD Julian) (long, lat float64) {
	var posn C.struct_ln_lnlat_posn
	C.ln_get_solar_ecl_coords(C.double(JD), &posn)
	long, lat = float64(posn.lng), float64(posn.lat)
	return
}

// void LIBNOVA_EXPORT ln_get_solar_geo_coords (double JD, struct ln_rect_posn * position);

// Calculate geocentric coordinates (rectangular).
func GetSolarGeoCoords(JD Julian) (x, y, z float64) {
	var posn C.struct_ln_rect_posn
	C.ln_get_solar_geo_coords(C.double(JD), &posn)
	x, y, z = float64(posn.X), float64(posn.Y), float64(posn.Z)
	return
}

// double LIBNOVA_EXPORT ln_get_solar_sdiam (double JD);

// Calculate the semidiameter of the Sun in arc seconds.
func GetSolarSdiam(JD Julian) float64 {
	return float64(C.ln_get_solar_sdiam(C.double(JD)))
}

//
// Convenience functions
//

// Calculate distance of Sun from Earth in AU.
func GetSolarEarthDist(JD Julian) float64 {
	x, y, z := GetSolarGeoCoords(JD)
	return math.Sqrt(x*x + y*y + z*z)
}
