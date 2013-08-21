/*
A wrapper for libnova -- Celestial Mechanics, Astrometry and Astrodynamics Library.

http://libnova.sourceforge.net/

Implemented so far:

 Julian date  (to/from time.Time)
 Lunar
 Solar

Tested with libnova v 0.15.0
*/
package gonova

/*
#cgo LDFLAGS: -lnova
*/
import "C"
