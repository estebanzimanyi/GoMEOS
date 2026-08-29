package gomeos

/*
// The installed libmeos describes itself: `pkg-config --cflags meos` reports its
// include dir together with the family macros it carries, and `--libs` reports
// where to link it from. A hand-kept path pair cannot report the macros, and the
// public headers gate declarations on them (meos.h holds #if MEOS, #if POINTCLOUD
// and #if JSON blocks), so a build that misses them sees a different API than the
// library exports. PKG_CONFIG_PATH then selects WHICH libmeos, so a build points
// at a private prefix rather than inheriting whatever occupies a machine-wide
// directory; a prefix outside the loader's search path needs LD_LIBRARY_PATH.
// This is the form functions/cgo.go already uses.
#cgo pkg-config: meos
#include "meos.h"
#include "meos_catalog.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"

type Interpolation C.int

const (
	INTERP_NONE    Interpolation = C.INTERP_NONE
	DISCRETE       Interpolation = C.DISCRETE
	STEP           Interpolation = C.STEP
	LINEAR                       = C.LINEAR
	ANYTEMPSUBTYPE               = C.ANYTEMPSUBTYPE /**< Any temporal subtype */
	TINSTANT                     = C.TINSTANT       /**< Temporal instant subtype */
	TSEQUENCE                    = C.TSEQUENCE      /**< Temporal sequence subtype */
	TSEQUENCESET                 = C.TSEQUENCESET   /**< Temporal sequence set subtype */
	T_TBOOL                      = C.T_TBOOL        /**< temporal boolean type */
	T_TFLOAT                     = C.T_TFLOAT       /**< temporal float type */
	T_TINT                       = C.T_TINT         /**< temporal integer type */
	T_TTEXT                      = C.T_TTEXT        /**< temporal text type */
	T_TGEOMPOINT                 = C.T_TGEOMPOINT   /**< temporal geometry point type */
	T_TGEOGPOINT                 = C.T_TGEOGPOINT   /**< temporal geography point type */
)

func init() {
}
