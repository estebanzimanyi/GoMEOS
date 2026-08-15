package functions

/*
// The installed libmeos describes itself: `pkg-config --cflags meos` reports
// its include dir together with the family macros it is compiled with, and
// `--libs` reports where to link it from. Both matter, because the public
// headers gate declarations on those macros (meos.h holds #if MEOS, #if
// POINTCLOUD and #if JSON blocks), so a hand-kept -D list that disagrees with
// the library either declares a symbol that cannot link or hides a family the
// library provides.
//
// Asking pkg-config also makes the binding relocatable: PKG_CONFIG_PATH selects
// which libmeos to build against, so a build points at a private prefix rather
// than inheriting whatever occupies a machine-wide directory. A prefix outside
// the loader's default search path needs LD_LIBRARY_PATH (or ldconfig).
#cgo pkg-config: meos

// H3 ships no pkg-config file, so the one dependency whose header the public
// MEOS headers expose (meos_h3.h includes <h3api.h>) names its include dir
// here. GEOS and GSL, the other exposed dependencies, resolve from the default
// include path.
#cgo darwin CFLAGS: -I/opt/homebrew/include/h3
#cgo linux CFLAGS: -I/usr/include/h3

#include <stddef.h>
#include "meos.h"
#include "meos_catalog.h"
#include "meos_geo.h"
#include "meos_internal.h"
#include "meos_internal_geo.h"
#include "meos_npoint.h"
#include "meos_cbuffer.h"
#include "meos_pose.h"
#include "meos_rgeo.h"
#include "meos_h3.h"
#include "meos_quadbin.h"
#include "meos_json.h"
#include "meos_pointcloud.h"
*/
import "C"
