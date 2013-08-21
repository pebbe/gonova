package gonova

/*
#include <libnova/lunar.h>

*/
import "C"

import (
	"errors"
)

// double LIBNOVA_EXPORT ln_get_lunar_sdiam (double JD);

// Calculate the semidiameter of the Moon in arc seconds.
func GetLunarSdiam(JD Julian) float64 {
	return float64(C.ln_get_lunar_sdiam(C.double(JD)))
}

// int LIBNOVA_EXPORT ln_get_lunar_rst (double JD, struct ln_lnlat_posn * observer, struct ln_rst_time * rst);

// Calculate the time of rise, set and transit for the Moon.
func GetLunarRST(JD Julian, lat, long float64) (rise, set, transit Julian, err error) {
	var rst C.struct_ln_rst_time
	observer := C.struct_ln_lnlat_posn{lng: C.double(long), lat: C.double(lat)}
	e := C.ln_get_lunar_rst(C.double(JD), &observer, &rst)
	if e != 0 {
		err = errors.New("No moonrise/moonset on this day at this position")
		return
	}
	rise = Julian(rst.rise)
	set = Julian(rst.set)
	transit = Julian(rst.transit)
	return
}

// void LIBNOVA_EXPORT ln_get_lunar_geo_posn (double JD, struct ln_rect_posn * moon, double precision);

// Calculate the rectangular geocentric lunar cordinates.
// ptrecision: The truncation level of the series in radians for longitude and latitude
// and in km for distance. (Valid range 0 - 0.01, 0 being highest accuracy)
func GetLunarGeoPosn(JD Julian, precision float64) (x, y, z float64) {
	var moon C.struct_ln_rect_posn
	if precision < 0 {
		precision = 0
	}
	if precision > 0.01 {
		precision = 0.01
	}
	C.ln_get_lunar_geo_posn(C.double(JD), &moon, C.double(precision))
	x, y, z = float64(moon.X), float64(moon.Y), float64(moon.Z)
	return
}

// void LIBNOVA_EXPORT ln_get_lunar_equ_coords_prec (double JD, struct ln_equ_posn * position, double precision);

// Calculate the lunar RA and DEC for Julian day JD.
// Accuracy is better than 10 arcsecs in right ascension and 4 arcsecs in declination.
// precision: The truncation level of the series in radians for longitude and latitude
// and in km for distance. (Valid range 0 - 0.01, 0 being highest accuracy)
func GetLunarEquCoordsPrec(JD Julian, precision float64) (ra, dec float64) {
	var r C.struct_ln_equ_posn
	if precision < 0 {
		precision = 0
	}
	if precision > 0.01 {
		precision = 0.01
	}
	C.ln_get_lunar_equ_coords_prec(C.double(JD), &r, C.double(precision))
	return float64(r.ra), float64(r.dec)
}

// void LIBNOVA_EXPORT ln_get_lunar_equ_coords (double JD, struct ln_equ_posn * position);

// Calculate the lunar RA and DEC for Julian day JD.
// Accuracy is better than 10 arcsecs in right ascension and 4 arcsecs in declination.
func GetLunarEquCoords(JD Julian) (ra, dec float64) {
	var r C.struct_ln_equ_posn
	C.ln_get_lunar_equ_coords(C.double(JD), &r)
	return float64(r.ra), float64(r.dec)
}

// void LIBNOVA_EXPORT ln_get_lunar_ecl_coords (double JD, struct ln_lnlat_posn * position, double precision);

// Calculate the lunar longitude and latitude for Julian day JD.
// Accuracy is better than 10 arcsecs in longitude and 4 arcsecs in latitude.
// precision: The truncation level of the series in radians for longitude
// and latitude and in km for distance. (Valid range 0 - 0.01, 0 being highest accuracy)
func GetLunarEclCoords(JD Julian, precision float64) (lng, lat float64) {
    var r C.struct_ln_lnlat_posn
	if precision < 0 {
		precision = 0
	}
	if precision > 0.01 {
		precision = 0.01
	}
	C.ln_get_lunar_ecl_coords(C.double(JD), &r, C.double(precision))
	lng, lat = float64(r.lng), float64(r.lat)
	return
}

// double LIBNOVA_EXPORT ln_get_lunar_phase (double JD);

// Calculate the phase angle of the Moon.
func GetLunarPhase(JD Julian) (angle float64) {
	return float64(C.ln_get_lunar_phase(C.double(JD)))
}

// double LIBNOVA_EXPORT ln_get_lunar_disk (double JD);

// Calculates the illuminated fraction of the Moon's disk.
func GetLunarDisk(JD Julian) float64 {
	return float64(C.ln_get_lunar_disk(C.double(JD)))
}

// double LIBNOVA_EXPORT ln_get_lunar_earth_dist (double JD);

// Calculates the distance between the centre of the Earth and the centre of the Moon in km.
func GetLunarEarthDist(JD Julian) float64 {
	return float64(C.ln_get_lunar_earth_dist(C.double(JD)))
}

// double LIBNOVA_EXPORT ln_get_lunar_bright_limb (double JD);

// Calculates the position angle of the midpoint of the illuminated
// limb of the moon, reckoned eastward from the north point of the disk.
func GetLunarBrightLimb(JD Julian) float64 {
	return float64(C.ln_get_lunar_bright_limb(C.double(JD)))
}

// double LIBNOVA_EXPORT ln_get_lunar_long_asc_node (double JD);

// Calculate the longitude of the Moon's mean ascending node.
func GetLunarLongAscNode(JD Julian) float64 {
	return float64(C.ln_get_lunar_long_asc_node(C.double(JD)))
}

// double LIBNOVA_EXPORT ln_get_lunar_long_perigee (double JD);

// Calculate the longitude of the Moon's mean perigee.
func GetLunarLongPerigee(JD Julian) float64 {
	return float64(C.ln_get_lunar_long_perigee(C.double(JD)))
}
