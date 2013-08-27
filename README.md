A wrapper for [libnova](http://libnova.sourceforge.net/) -- Celestial Mechanics, Astrometry and Astrodynamics Library.

The library is not very accurate. If you want something more precise,
try [novas](https://github.com/pebbe/novas)

Implemented so far:
 * Julian date  (to/from time.Time)
 * Lunar
 * Solar

Tested with libnova v 0.15.0

Keywords: astronomy, astrometry, celestial mechanics, sun, moon, planets

## Install

You need libnova installed, with libs and headers where gcc can find them.

    go get github.com/pebbe/gonova

## Docs

 * [package help](http://godoc.org/github.com/pebbe/gonova)
