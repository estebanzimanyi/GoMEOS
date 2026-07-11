package functions

/*
// Every optional MEOS family is enabled so the ``#if <FAMILY>``-guarded
// declarations in the core headers (e.g. meos_initialize_pointcloud) are
// visible -- matching the all-families libmeos the catalog is derived from, in
// sync with MobilityDB CMakeLists.txt's ``if(ALL)`` loop (alphabetical). The
// families expose external library types (H3's h3api.h) through their own
// headers, so their include dirs are added too.
#cgo darwin CFLAGS: -I/opt/homebrew/include -I/opt/homebrew/include/h3 -DMEOS=1 -DARROW=1 -DCBUFFER=1 -DH3=1 -DJSON=1 -DNPOINT=1 -DPOINTCLOUD=1 -DPOSE=1 -DQUADBIN=1 -DRASTER=1 -DRGEO=1
#cgo darwin LDFLAGS: -L/opt/homebrew/lib -lmeos -Wl,-rpath,/opt/homebrew/lib

#cgo linux CFLAGS: -I/usr/local/include/ -I/usr/include/h3 -DMEOS=1 -DARROW=1 -DCBUFFER=1 -DH3=1 -DJSON=1 -DNPOINT=1 -DPOINTCLOUD=1 -DPOSE=1 -DQUADBIN=1 -DRASTER=1 -DRGEO=1
#cgo linux LDFLAGS: -L/usr/local/lib -lmeos -Wl,-rpath,/usr/local/lib

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
#include "meos_arrow.h"
*/
import "C"
