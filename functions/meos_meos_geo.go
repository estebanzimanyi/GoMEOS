package functions

/*
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

// cgo reads C.union_* as a union TYPE; alias the union operators so C.g<name> resolves them as functions.
#define gunion_bigint_set union_bigint_set
#define gunion_bigint_span union_bigint_span
#define gunion_bigint_spanset union_bigint_spanset
#define gunion_cbuffer_set union_cbuffer_set
#define gunion_date_set union_date_set
#define gunion_date_span union_date_span
#define gunion_date_spanset union_date_spanset
#define gunion_float_set union_float_set
#define gunion_float_span union_float_span
#define gunion_float_spanset union_float_spanset
#define gunion_geo_set union_geo_set
#define gunion_int_set union_int_set
#define gunion_int_span union_int_span
#define gunion_int_spanset union_int_spanset
#define gunion_jsonb_set union_jsonb_set
#define gunion_npoint_set union_npoint_set
#define gunion_pcpatch_set union_pcpatch_set
#define gunion_pcpoint_set union_pcpoint_set
#define gunion_pose_set union_pose_set
#define gunion_set_bigint union_set_bigint
#define gunion_set_cbuffer union_set_cbuffer
#define gunion_set_date union_set_date
#define gunion_set_float union_set_float
#define gunion_set_geo union_set_geo
#define gunion_set_int union_set_int
#define gunion_set_jsonb union_set_jsonb
#define gunion_set_npoint union_set_npoint
#define gunion_set_pcpatch union_set_pcpatch
#define gunion_set_pcpoint union_set_pcpoint
#define gunion_set_pose union_set_pose
#define gunion_set_set union_set_set
#define gunion_set_text union_set_text
#define gunion_set_timestamptz union_set_timestamptz
#define gunion_set_value union_set_value
#define gunion_span_bigint union_span_bigint
#define gunion_span_date union_span_date
#define gunion_span_float union_span_float
#define gunion_span_int union_span_int
#define gunion_span_span union_span_span
#define gunion_span_spanset union_span_spanset
#define gunion_span_timestamptz union_span_timestamptz
#define gunion_span_value union_span_value
#define gunion_spanset_bigint union_spanset_bigint
#define gunion_spanset_date union_spanset_date
#define gunion_spanset_float union_spanset_float
#define gunion_spanset_int union_spanset_int
#define gunion_spanset_span union_spanset_span
#define gunion_spanset_spanset union_spanset_spanset
#define gunion_spanset_timestamptz union_spanset_timestamptz
#define gunion_spanset_value union_spanset_value
#define gunion_stbox_stbox union_stbox_stbox
#define gunion_tbox_tbox union_tbox_tbox
#define gunion_text_set union_text_set
#define gunion_timestamptz_set union_timestamptz_set
#define gunion_timestamptz_span union_timestamptz_span
#define gunion_timestamptz_spanset union_timestamptz_spanset
#define gunion_tpcbox_tpcbox union_tpcbox_tpcbox
#define gunion_value_set union_value_set
#define gunion_value_span union_value_span
#define gunion_value_spanset union_value_spanset
*/
import "C"
import (
	"unsafe"
)

var _ = unsafe.Pointer(nil)

// Box3dFromGbox wraps MEOS C function box3d_from_gbox.
func Box3dFromGbox(box *GBox) *Box3D {
	_cret := C.box3d_from_gbox(box._inner)
	return &Box3D{_inner: _cret}
}


// Box3dMake wraps MEOS C function box3d_make.
func Box3dMake(xmin float64, xmax float64, ymin float64, ymax float64, zmin float64, zmax float64, srid int32) *Box3D {
	_cret := C.box3d_make(C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.double(zmin), C.double(zmax), C.int32_t(srid))
	return &Box3D{_inner: _cret}
}


// Box3dIn wraps MEOS C function box3d_in.
func Box3dIn(str string) *Box3D {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.box3d_in(_c_str)
	return &Box3D{_inner: _cret}
}


// Box3dOut wraps MEOS C function box3d_out.
func Box3dOut(box *Box3D, maxdd int) string {
	_cret := C.box3d_out(box._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// GboxMake wraps MEOS C function gbox_make.
func GboxMake(hasz bool, hasm bool, geodetic bool, xmin float64, xmax float64, ymin float64, ymax float64, zmin float64, zmax float64, mmin float64, mmax float64) *GBox {
	_cret := C.gbox_make(C.bool(hasz), C.bool(hasm), C.bool(geodetic), C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.double(zmin), C.double(zmax), C.double(mmin), C.double(mmax))
	return &GBox{_inner: _cret}
}


// GboxIn wraps MEOS C function gbox_in.
func GboxIn(str string) *GBox {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.gbox_in(_c_str)
	return &GBox{_inner: _cret}
}


// GboxOut wraps MEOS C function gbox_out.
func GboxOut(box *GBox, maxdd int) string {
	_cret := C.gbox_out(box._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// GeoAsEWKB wraps MEOS C function geo_as_ewkb.
func GeoAsEWKB(gs *Geom, endian string, size unsafe.Pointer) unsafe.Pointer {
	_c_endian := C.CString(endian)
	defer C.free(unsafe.Pointer(_c_endian))
	_cret := C.geo_as_ewkb(gs._inner, _c_endian, (*C.size_t)(unsafe.Pointer(size)))
	return unsafe.Pointer(_cret)
}


// GeoAsEWKT wraps MEOS C function geo_as_ewkt.
func GeoAsEWKT(gs *Geom, precision int) string {
	_cret := C.geo_as_ewkt(gs._inner, C.int(precision))
	return C.GoString(_cret)
}


// GeoAsGeojson wraps MEOS C function geo_as_geojson.
func GeoAsGeojson(gs *Geom, option int, precision int, srs string) string {
	_c_srs := C.CString(srs)
	defer C.free(unsafe.Pointer(_c_srs))
	_cret := C.geo_as_geojson(gs._inner, C.int(option), C.int(precision), _c_srs)
	return C.GoString(_cret)
}


// GeoAsHexewkb wraps MEOS C function geo_as_hexewkb.
func GeoAsHexewkb(gs *Geom, endian string) string {
	_c_endian := C.CString(endian)
	defer C.free(unsafe.Pointer(_c_endian))
	_cret := C.geo_as_hexewkb(gs._inner, _c_endian)
	return C.GoString(_cret)
}


// GeoAsText wraps MEOS C function geo_as_text.
func GeoAsText(gs *Geom, precision int) string {
	_cret := C.geo_as_text(gs._inner, C.int(precision))
	return C.GoString(_cret)
}


// GeoFromEWKB wraps MEOS C function geo_from_ewkb.
func GeoFromEWKB(wkb unsafe.Pointer, wkb_size uint, srid int32) *Geom {
	_cret := C.geo_from_ewkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(wkb_size), C.int32_t(srid))
	return &Geom{_inner: _cret}
}


// GeoFromGeojson wraps MEOS C function geo_from_geojson.
func GeoFromGeojson(geojson string) *Geom {
	_c_geojson := C.CString(geojson)
	defer C.free(unsafe.Pointer(_c_geojson))
	_cret := C.geo_from_geojson(_c_geojson)
	return &Geom{_inner: _cret}
}


// GeoFromText wraps MEOS C function geo_from_text.
func GeoFromText(wkt string, srid int32) *Geom {
	_c_wkt := C.CString(wkt)
	defer C.free(unsafe.Pointer(_c_wkt))
	_cret := C.geo_from_text(_c_wkt, C.int32_t(srid))
	return &Geom{_inner: _cret}
}


// GeoOut wraps MEOS C function geo_out.
func GeoOut(gs *Geom) string {
	_cret := C.geo_out(gs._inner)
	return C.GoString(_cret)
}


// GeogFromHexewkb wraps MEOS C function geog_from_hexewkb.
func GeogFromHexewkb(wkt string) *Geom {
	_c_wkt := C.CString(wkt)
	defer C.free(unsafe.Pointer(_c_wkt))
	_cret := C.geog_from_hexewkb(_c_wkt)
	return &Geom{_inner: _cret}
}


// GeogIn wraps MEOS C function geog_in.
func GeogIn(str string, typmod int32) *Geom {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.geog_in(_c_str, C.int32(typmod))
	return &Geom{_inner: _cret}
}


// GeomFromHexewkb wraps MEOS C function geom_from_hexewkb.
func GeomFromHexewkb(wkt string) *Geom {
	_c_wkt := C.CString(wkt)
	defer C.free(unsafe.Pointer(_c_wkt))
	_cret := C.geom_from_hexewkb(_c_wkt)
	return &Geom{_inner: _cret}
}


// GeomIn wraps MEOS C function geom_in.
func GeomIn(str string, typmod int32) *Geom {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.geom_in(_c_str, C.int32(typmod))
	return &Geom{_inner: _cret}
}


// GeoCopy wraps MEOS C function geo_copy.
func GeoCopy(gs *Geom) *Geom {
	_cret := C.geo_copy(gs._inner)
	return &Geom{_inner: _cret}
}


// GeogpointMake2d wraps MEOS C function geogpoint_make2d.
func GeogpointMake2d(srid int32, x float64, y float64) *Geom {
	_cret := C.geogpoint_make2d(C.int32_t(srid), C.double(x), C.double(y))
	return &Geom{_inner: _cret}
}


// GeogpointMake3dz wraps MEOS C function geogpoint_make3dz.
func GeogpointMake3dz(srid int32, x float64, y float64, z float64) *Geom {
	_cret := C.geogpoint_make3dz(C.int32_t(srid), C.double(x), C.double(y), C.double(z))
	return &Geom{_inner: _cret}
}


// GeompointMake2d wraps MEOS C function geompoint_make2d.
func GeompointMake2d(srid int32, x float64, y float64) *Geom {
	_cret := C.geompoint_make2d(C.int32_t(srid), C.double(x), C.double(y))
	return &Geom{_inner: _cret}
}


// GeompointMake3dz wraps MEOS C function geompoint_make3dz.
func GeompointMake3dz(srid int32, x float64, y float64, z float64) *Geom {
	_cret := C.geompoint_make3dz(C.int32_t(srid), C.double(x), C.double(y), C.double(z))
	return &Geom{_inner: _cret}
}


// GeomToGeog wraps MEOS C function geom_to_geog.
func GeomToGeog(geom *Geom) *Geom {
	_cret := C.geom_to_geog(geom._inner)
	return &Geom{_inner: _cret}
}


// GeogToGeom wraps MEOS C function geog_to_geom.
func GeogToGeom(geog *Geom) *Geom {
	_cret := C.geog_to_geom(geog._inner)
	return &Geom{_inner: _cret}
}


// GeoIsEmpty wraps MEOS C function geo_is_empty.
func GeoIsEmpty(gs *Geom) bool {
	_cret := C.geo_is_empty(gs._inner)
	return bool(_cret)
}


// GeoIsUnitary wraps MEOS C function geo_is_unitary.
func GeoIsUnitary(gs *Geom) bool {
	_cret := C.geo_is_unitary(gs._inner)
	return bool(_cret)
}


// GeoTypename wraps MEOS C function geo_typename.
func GeoTypename(type_ int) string {
	_cret := C.geo_typename(C.int(type_))
	return C.GoString(_cret)
}


// GeogArea wraps MEOS C function geog_area.
func GeogArea(gs *Geom, use_spheroid bool) float64 {
	_cret := C.geog_area(gs._inner, C.bool(use_spheroid))
	return float64(_cret)
}


// GeogCentroid wraps MEOS C function geog_centroid.
func GeogCentroid(gs *Geom, use_spheroid bool) *Geom {
	_cret := C.geog_centroid(gs._inner, C.bool(use_spheroid))
	return &Geom{_inner: _cret}
}


// GeogLength wraps MEOS C function geog_length.
func GeogLength(gs *Geom, use_spheroid bool) float64 {
	_cret := C.geog_length(gs._inner, C.bool(use_spheroid))
	return float64(_cret)
}


// GeogPerimeter wraps MEOS C function geog_perimeter.
func GeogPerimeter(gs *Geom, use_spheroid bool) float64 {
	_cret := C.geog_perimeter(gs._inner, C.bool(use_spheroid))
	return float64(_cret)
}


// GeomAzimuth wraps MEOS C function geom_azimuth.
func GeomAzimuth(gs1 *Geom, gs2 *Geom) (bool, float64) {
	var _out_result C.double
	_cret := C.geom_azimuth(gs1._inner, gs2._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// GeomLength wraps MEOS C function geom_length.
func GeomLength(gs *Geom) float64 {
	_cret := C.geom_length(gs._inner)
	return float64(_cret)
}


// GeomPerimeter wraps MEOS C function geom_perimeter.
func GeomPerimeter(gs *Geom) float64 {
	_cret := C.geom_perimeter(gs._inner)
	return float64(_cret)
}


// LineNumpoints wraps MEOS C function line_numpoints.
func LineNumpoints(gs *Geom) int {
	_cret := C.line_numpoints(gs._inner)
	return int(_cret)
}


// LinePointN wraps MEOS C function line_point_n.
func LinePointN(geom *Geom, n int) *Geom {
	_cret := C.line_point_n(geom._inner, C.int(n))
	return &Geom{_inner: _cret}
}


// GeoReverse wraps MEOS C function geo_reverse.
func GeoReverse(gs *Geom) *Geom {
	_cret := C.geo_reverse(gs._inner)
	return &Geom{_inner: _cret}
}


// GeoRound wraps MEOS C function geo_round.
func GeoRound(gs *Geom, maxdd int) *Geom {
	_cret := C.geo_round(gs._inner, C.int(maxdd))
	return &Geom{_inner: _cret}
}


// GeoSetSRID wraps MEOS C function geo_set_srid.
func GeoSetSRID(gs *Geom, srid int32) *Geom {
	_cret := C.geo_set_srid(gs._inner, C.int32_t(srid))
	return &Geom{_inner: _cret}
}


// GeoSRID wraps MEOS C function geo_srid.
func GeoSRID(gs *Geom) int32 {
	_cret := C.geo_srid(gs._inner)
	return int32(_cret)
}


// GeoTransform wraps MEOS C function geo_transform.
func GeoTransform(geom *Geom, srid_to int32) *Geom {
	_cret := C.geo_transform(geom._inner, C.int32_t(srid_to))
	return &Geom{_inner: _cret}
}


// GeoTransformPipeline wraps MEOS C function geo_transform_pipeline.
func GeoTransformPipeline(gs *Geom, pipeline string, srid_to int32, is_forward bool) *Geom {
	_c_pipeline := C.CString(pipeline)
	defer C.free(unsafe.Pointer(_c_pipeline))
	_cret := C.geo_transform_pipeline(gs._inner, _c_pipeline, C.int32_t(srid_to), C.bool(is_forward))
	return &Geom{_inner: _cret}
}


// GeoCollectGarray wraps MEOS C function geo_collect_garray.
func GeoCollectGarray(gsarr unsafe.Pointer, count int) *Geom {
	_cret := C.geo_collect_garray((**C.GSERIALIZED)(unsafe.Pointer(gsarr)), C.int(count))
	return &Geom{_inner: _cret}
}


// GeoMakelineGarray wraps MEOS C function geo_makeline_garray.
func GeoMakelineGarray(gsarr unsafe.Pointer, count int) *Geom {
	_cret := C.geo_makeline_garray((**C.GSERIALIZED)(unsafe.Pointer(gsarr)), C.int(count))
	return &Geom{_inner: _cret}
}


// GeoNumPoints wraps MEOS C function geo_num_points.
func GeoNumPoints(gs *Geom) int {
	_cret := C.geo_num_points(gs._inner)
	return int(_cret)
}


// GeoNumGeos wraps MEOS C function geo_num_geos.
func GeoNumGeos(gs *Geom) int {
	_cret := C.geo_num_geos(gs._inner)
	return int(_cret)
}


// GeoGeoN wraps MEOS C function geo_geo_n.
func GeoGeoN(geom *Geom, n int) *Geom {
	_cret := C.geo_geo_n(geom._inner, C.int(n))
	return &Geom{_inner: _cret}
}


// GeoPointarr wraps MEOS C function geo_pointarr.
func GeoPointarr(gs *Geom, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.geo_pointarr(gs._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// GeoPoints wraps MEOS C function geo_points.
func GeoPoints(gs *Geom) *Geom {
	_cret := C.geo_points(gs._inner)
	return &Geom{_inner: _cret}
}


// GeomArrayUnion wraps MEOS C function geom_array_union.
func GeomArrayUnion(gsarr unsafe.Pointer, count int) *Geom {
	_cret := C.geom_array_union((**C.GSERIALIZED)(unsafe.Pointer(gsarr)), C.int(count))
	return &Geom{_inner: _cret}
}


// GeomBoundary wraps MEOS C function geom_boundary.
func GeomBoundary(gs *Geom) *Geom {
	_cret := C.geom_boundary(gs._inner)
	return &Geom{_inner: _cret}
}


// GeomBuffer wraps MEOS C function geom_buffer.
func GeomBuffer(gs *Geom, size float64, params string) *Geom {
	_c_params := C.CString(params)
	defer C.free(unsafe.Pointer(_c_params))
	_cret := C.geom_buffer(gs._inner, C.double(size), _c_params)
	return &Geom{_inner: _cret}
}


// GeomCentroid wraps MEOS C function geom_centroid.
func GeomCentroid(gs *Geom) *Geom {
	_cret := C.geom_centroid(gs._inner)
	return &Geom{_inner: _cret}
}


// GeomConvexHull wraps MEOS C function geom_convex_hull.
func GeomConvexHull(gs *Geom) *Geom {
	_cret := C.geom_convex_hull(gs._inner)
	return &Geom{_inner: _cret}
}


// GeomDifference2d wraps MEOS C function geom_difference2d.
func GeomDifference2d(gs1 *Geom, gs2 *Geom) *Geom {
	_cret := C.geom_difference2d(gs1._inner, gs2._inner)
	return &Geom{_inner: _cret}
}


// GeomIntersection2d wraps MEOS C function geom_intersection2d.
func GeomIntersection2d(gs1 *Geom, gs2 *Geom) *Geom {
	_cret := C.geom_intersection2d(gs1._inner, gs2._inner)
	return &Geom{_inner: _cret}
}


// GeomIntersection2dColl wraps MEOS C function geom_intersection2d_coll.
func GeomIntersection2dColl(gs1 *Geom, gs2 *Geom) *Geom {
	_cret := C.geom_intersection2d_coll(gs1._inner, gs2._inner)
	return &Geom{_inner: _cret}
}


// GeomMinBoundingRadius wraps MEOS C function geom_min_bounding_radius.
func GeomMinBoundingRadius(geom *Geom, radius unsafe.Pointer) *Geom {
	_cret := C.geom_min_bounding_radius(geom._inner, (*C.double)(unsafe.Pointer(radius)))
	return &Geom{_inner: _cret}
}


// GeomShortestline2d wraps MEOS C function geom_shortestline2d.
func GeomShortestline2d(gs1 *Geom, gs2 *Geom) *Geom {
	_cret := C.geom_shortestline2d(gs1._inner, gs2._inner)
	return &Geom{_inner: _cret}
}


// GeomShortestline3d wraps MEOS C function geom_shortestline3d.
func GeomShortestline3d(gs1 *Geom, gs2 *Geom) *Geom {
	_cret := C.geom_shortestline3d(gs1._inner, gs2._inner)
	return &Geom{_inner: _cret}
}


// GeomUnaryUnion wraps MEOS C function geom_unary_union.
func GeomUnaryUnion(gs *Geom, prec float64) *Geom {
	_cret := C.geom_unary_union(gs._inner, C.double(prec))
	return &Geom{_inner: _cret}
}


// LineInterpolatePoint wraps MEOS C function line_interpolate_point.
func LineInterpolatePoint(gs *Geom, distance_fraction float64, repeat bool) *Geom {
	_cret := C.line_interpolate_point(gs._inner, C.double(distance_fraction), C.bool(repeat))
	return &Geom{_inner: _cret}
}


// LineLocatePoint wraps MEOS C function line_locate_point.
func LineLocatePoint(gs1 *Geom, gs2 *Geom) float64 {
	_cret := C.line_locate_point(gs1._inner, gs2._inner)
	return float64(_cret)
}


// LineSubstring wraps MEOS C function line_substring.
func LineSubstring(gs *Geom, from float64, to float64) *Geom {
	_cret := C.line_substring(gs._inner, C.double(from), C.double(to))
	return &Geom{_inner: _cret}
}


// GeogDwithin wraps MEOS C function geog_dwithin.
func GeogDwithin(g1 *Geom, g2 *Geom, tolerance float64, use_spheroid bool) bool {
	_cret := C.geog_dwithin(g1._inner, g2._inner, C.double(tolerance), C.bool(use_spheroid))
	return bool(_cret)
}


// GeogIntersects wraps MEOS C function geog_intersects.
func GeogIntersects(gs1 *Geom, gs2 *Geom, use_spheroid bool) bool {
	_cret := C.geog_intersects(gs1._inner, gs2._inner, C.bool(use_spheroid))
	return bool(_cret)
}


// GeomContains wraps MEOS C function geom_contains.
func GeomContains(gs1 *Geom, gs2 *Geom) bool {
	_cret := C.geom_contains(gs1._inner, gs2._inner)
	return bool(_cret)
}


// GeomCovers wraps MEOS C function geom_covers.
func GeomCovers(gs1 *Geom, gs2 *Geom) bool {
	_cret := C.geom_covers(gs1._inner, gs2._inner)
	return bool(_cret)
}


// GeomDisjoint2d wraps MEOS C function geom_disjoint2d.
func GeomDisjoint2d(gs1 *Geom, gs2 *Geom) bool {
	_cret := C.geom_disjoint2d(gs1._inner, gs2._inner)
	return bool(_cret)
}


// GeomDwithin wraps MEOS C function geom_dwithin.
func GeomDwithin(gs1 *Geom, gs2 *Geom, tolerance float64) bool {
	_cret := C.geom_dwithin(gs1._inner, gs2._inner, C.double(tolerance))
	return bool(_cret)
}


// GeomDwithin2d wraps MEOS C function geom_dwithin2d.
func GeomDwithin2d(gs1 *Geom, gs2 *Geom, tolerance float64) bool {
	_cret := C.geom_dwithin2d(gs1._inner, gs2._inner, C.double(tolerance))
	return bool(_cret)
}


// GeomDwithin3d wraps MEOS C function geom_dwithin3d.
func GeomDwithin3d(gs1 *Geom, gs2 *Geom, tolerance float64) bool {
	_cret := C.geom_dwithin3d(gs1._inner, gs2._inner, C.double(tolerance))
	return bool(_cret)
}


// GeomIntersects wraps MEOS C function geom_intersects.
func GeomIntersects(gs1 *Geom, gs2 *Geom) bool {
	_cret := C.geom_intersects(gs1._inner, gs2._inner)
	return bool(_cret)
}


// GeomIntersects2d wraps MEOS C function geom_intersects2d.
func GeomIntersects2d(gs1 *Geom, gs2 *Geom) bool {
	_cret := C.geom_intersects2d(gs1._inner, gs2._inner)
	return bool(_cret)
}


// GeomIntersects3d wraps MEOS C function geom_intersects3d.
func GeomIntersects3d(gs1 *Geom, gs2 *Geom) bool {
	_cret := C.geom_intersects3d(gs1._inner, gs2._inner)
	return bool(_cret)
}


// GeomRelatePattern wraps MEOS C function geom_relate_pattern.
func GeomRelatePattern(gs1 *Geom, gs2 *Geom, patt string) bool {
	_c_patt := C.CString(patt)
	defer C.free(unsafe.Pointer(_c_patt))
	_cret := C.geom_relate_pattern(gs1._inner, gs2._inner, _c_patt)
	return bool(_cret)
}


// GeomTouches wraps MEOS C function geom_touches.
func GeomTouches(gs1 *Geom, gs2 *Geom) bool {
	_cret := C.geom_touches(gs1._inner, gs2._inner)
	return bool(_cret)
}


// GeoStboxes wraps MEOS C function geo_stboxes.
func GeoStboxes(gs *Geom, count unsafe.Pointer) *STBox {
	_cret := C.geo_stboxes(gs._inner, (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// GeoSplitEachNStboxes wraps MEOS C function geo_split_each_n_stboxes.
func GeoSplitEachNStboxes(gs *Geom, elem_count int, count unsafe.Pointer) *STBox {
	_cret := C.geo_split_each_n_stboxes(gs._inner, C.int(elem_count), (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// GeoSplitNStboxes wraps MEOS C function geo_split_n_stboxes.
func GeoSplitNStboxes(gs *Geom, box_count int, count unsafe.Pointer) *STBox {
	_cret := C.geo_split_n_stboxes(gs._inner, C.int(box_count), (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// GeogDistance wraps MEOS C function geog_distance.
func GeogDistance(g1 *Geom, g2 *Geom) float64 {
	_cret := C.geog_distance(g1._inner, g2._inner)
	return float64(_cret)
}


// GeomDistance2d wraps MEOS C function geom_distance2d.
func GeomDistance2d(gs1 *Geom, gs2 *Geom) float64 {
	_cret := C.geom_distance2d(gs1._inner, gs2._inner)
	return float64(_cret)
}


// GeomDistance3d wraps MEOS C function geom_distance3d.
func GeomDistance3d(gs1 *Geom, gs2 *Geom) float64 {
	_cret := C.geom_distance3d(gs1._inner, gs2._inner)
	return float64(_cret)
}


// GeoEquals wraps MEOS C function geo_equals.
func GeoEquals(gs1 *Geom, gs2 *Geom) int {
	_cret := C.geo_equals(gs1._inner, gs2._inner)
	return int(_cret)
}


// GeoSame wraps MEOS C function geo_same.
func GeoSame(gs1 *Geom, gs2 *Geom) bool {
	_cret := C.geo_same(gs1._inner, gs2._inner)
	return bool(_cret)
}


// GeogsetIn wraps MEOS C function geogset_in.
func GeogsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.geogset_in(_c_str)
	return &Set{_inner: _cret}
}


// GeomsetIn wraps MEOS C function geomset_in.
func GeomsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.geomset_in(_c_str)
	return &Set{_inner: _cret}
}


// SpatialsetOut wraps MEOS C function spatialset_out.
func SpatialsetOut(s *Set, maxdd int) string {
	_cret := C.spatialset_out(s._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// SpatialsetAsText wraps MEOS C function spatialset_as_text.
func SpatialsetAsText(set *Set, maxdd int) string {
	_cret := C.spatialset_as_text(set._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// SpatialsetAsEWKT wraps MEOS C function spatialset_as_ewkt.
func SpatialsetAsEWKT(set *Set, maxdd int) string {
	_cret := C.spatialset_as_ewkt(set._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// GeosetMake wraps MEOS C function geoset_make.
func GeosetMake(values unsafe.Pointer, count int) *Set {
	_cret := C.geoset_make((**C.GSERIALIZED)(unsafe.Pointer(values)), C.int(count))
	return &Set{_inner: _cret}
}


// GeoToSet wraps MEOS C function geo_to_set.
func GeoToSet(gs *Geom) *Set {
	_cret := C.geo_to_set(gs._inner)
	return &Set{_inner: _cret}
}


// GeosetEndValue wraps MEOS C function geoset_end_value.
func GeosetEndValue(s *Set) *Geom {
	_cret := C.geoset_end_value(s._inner)
	return &Geom{_inner: _cret}
}


// GeosetStartValue wraps MEOS C function geoset_start_value.
func GeosetStartValue(s *Set) *Geom {
	_cret := C.geoset_start_value(s._inner)
	return &Geom{_inner: _cret}
}


// GeosetValueN wraps MEOS C function geoset_value_n.
func GeosetValueN(s *Set, n int) (bool, *Geom) {
	var _out_result *C.GSERIALIZED
	_cret := C.geoset_value_n(s._inner, C.int(n), &_out_result)
	return bool(_cret), &Geom{_inner: _out_result}
}


// GeosetValues wraps MEOS C function geoset_values.
func GeosetValues(s *Set, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.geoset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// ContainedGeoSet wraps MEOS C function contained_geo_set.
func ContainedGeoSet(gs *Geom, s *Set) bool {
	_cret := C.contained_geo_set(gs._inner, s._inner)
	return bool(_cret)
}


// ContainsSetGeo wraps MEOS C function contains_set_geo.
func ContainsSetGeo(s *Set, gs *Geom) bool {
	_cret := C.contains_set_geo(s._inner, gs._inner)
	return bool(_cret)
}


// GeoUnionTransfn wraps MEOS C function geo_union_transfn.
func GeoUnionTransfn(state *Set, gs *Geom) *Set {
	_cret := C.geo_union_transfn(state._inner, gs._inner)
	return &Set{_inner: _cret}
}


// IntersectionGeoSet wraps MEOS C function intersection_geo_set.
func IntersectionGeoSet(gs *Geom, s *Set) *Set {
	_cret := C.intersection_geo_set(gs._inner, s._inner)
	return &Set{_inner: _cret}
}


// IntersectionSetGeo wraps MEOS C function intersection_set_geo.
func IntersectionSetGeo(s *Set, gs *Geom) *Set {
	_cret := C.intersection_set_geo(s._inner, gs._inner)
	return &Set{_inner: _cret}
}


// MinusGeoSet wraps MEOS C function minus_geo_set.
func MinusGeoSet(gs *Geom, s *Set) *Set {
	_cret := C.minus_geo_set(gs._inner, s._inner)
	return &Set{_inner: _cret}
}


// MinusSetGeo wraps MEOS C function minus_set_geo.
func MinusSetGeo(s *Set, gs *Geom) *Set {
	_cret := C.minus_set_geo(s._inner, gs._inner)
	return &Set{_inner: _cret}
}


// UnionGeoSet wraps MEOS C function union_geo_set.
func UnionGeoSet(gs *Geom, s *Set) *Set {
	_cret := C.gunion_geo_set(gs._inner, s._inner)
	return &Set{_inner: _cret}
}


// UnionSetGeo wraps MEOS C function union_set_geo.
func UnionSetGeo(s *Set, gs *Geom) *Set {
	_cret := C.gunion_set_geo(s._inner, gs._inner)
	return &Set{_inner: _cret}
}


// SpatialsetSetSRID wraps MEOS C function spatialset_set_srid.
func SpatialsetSetSRID(s *Set, srid int32) *Set {
	_cret := C.spatialset_set_srid(s._inner, C.int32_t(srid))
	return &Set{_inner: _cret}
}


// SpatialsetSRID wraps MEOS C function spatialset_srid.
func SpatialsetSRID(s *Set) int32 {
	_cret := C.spatialset_srid(s._inner)
	return int32(_cret)
}


// SpatialsetTransform wraps MEOS C function spatialset_transform.
func SpatialsetTransform(s *Set, srid int32) *Set {
	_cret := C.spatialset_transform(s._inner, C.int32_t(srid))
	return &Set{_inner: _cret}
}


// SpatialsetTransformPipeline wraps MEOS C function spatialset_transform_pipeline.
func SpatialsetTransformPipeline(s *Set, pipelinestr string, srid int32, is_forward bool) *Set {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	_cret := C.spatialset_transform_pipeline(s._inner, _c_pipelinestr, C.int32_t(srid), C.bool(is_forward))
	return &Set{_inner: _cret}
}


// STBOXAsHexwkb wraps MEOS C function stbox_as_hexwkb.
func STBOXAsHexwkb(box *STBox, variant uint8, size_out unsafe.Pointer) string {
	_cret := C.stbox_as_hexwkb(box._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return C.GoString(_cret)
}


// STBOXAsWKB wraps MEOS C function stbox_as_wkb.
func STBOXAsWKB(box *STBox, variant uint8, size_out unsafe.Pointer) unsafe.Pointer {
	_cret := C.stbox_as_wkb(box._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return unsafe.Pointer(_cret)
}


// STBOXFromHexwkb wraps MEOS C function stbox_from_hexwkb.
func STBOXFromHexwkb(hexwkb string) *STBox {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	_cret := C.stbox_from_hexwkb(_c_hexwkb)
	return &STBox{_inner: _cret}
}


// STBOXFromWKB wraps MEOS C function stbox_from_wkb.
func STBOXFromWKB(wkb unsafe.Pointer, size uint) *STBox {
	_cret := C.stbox_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	return &STBox{_inner: _cret}
}


// STBOXIn wraps MEOS C function stbox_in.
func STBOXIn(str string) *STBox {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.stbox_in(_c_str)
	return &STBox{_inner: _cret}
}


// STBOXOut wraps MEOS C function stbox_out.
func STBOXOut(box *STBox, maxdd int) string {
	_cret := C.stbox_out(box._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// GeoTimestamptzToSTBOX wraps MEOS C function geo_timestamptz_to_stbox.
func GeoTimestamptzToSTBOX(gs *Geom, t int64) *STBox {
	_cret := C.geo_timestamptz_to_stbox(gs._inner, C.TimestampTz(t))
	return &STBox{_inner: _cret}
}


// GeoTstzspanToSTBOX wraps MEOS C function geo_tstzspan_to_stbox.
func GeoTstzspanToSTBOX(gs *Geom, s *Span) *STBox {
	_cret := C.geo_tstzspan_to_stbox(gs._inner, s._inner)
	return &STBox{_inner: _cret}
}


// STBOXCopy wraps MEOS C function stbox_copy.
func STBOXCopy(box *STBox) *STBox {
	_cret := C.stbox_copy(box._inner)
	return &STBox{_inner: _cret}
}


// STBOXMake wraps MEOS C function stbox_make.
func STBOXMake(hasx bool, hasz bool, geodetic bool, srid int32, xmin float64, xmax float64, ymin float64, ymax float64, zmin float64, zmax float64, s *Span) *STBox {
	_cret := C.stbox_make(C.bool(hasx), C.bool(hasz), C.bool(geodetic), C.int32_t(srid), C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.double(zmin), C.double(zmax), s._inner)
	return &STBox{_inner: _cret}
}


// GeoToSTBOX wraps MEOS C function geo_to_stbox.
func GeoToSTBOX(gs *Geom) *STBox {
	_cret := C.geo_to_stbox(gs._inner)
	return &STBox{_inner: _cret}
}


// SpatialsetToSTBOX wraps MEOS C function spatialset_to_stbox.
func SpatialsetToSTBOX(s *Set) *STBox {
	_cret := C.spatialset_to_stbox(s._inner)
	return &STBox{_inner: _cret}
}


// STBOXToBox3d wraps MEOS C function stbox_to_box3d.
func STBOXToBox3d(box *STBox) *Box3D {
	_cret := C.stbox_to_box3d(box._inner)
	return &Box3D{_inner: _cret}
}


// STBOXToGbox wraps MEOS C function stbox_to_gbox.
func STBOXToGbox(box *STBox) *GBox {
	_cret := C.stbox_to_gbox(box._inner)
	return &GBox{_inner: _cret}
}


// STBOXToGeo wraps MEOS C function stbox_to_geo.
func STBOXToGeo(box *STBox) *Geom {
	_cret := C.stbox_to_geo(box._inner)
	return &Geom{_inner: _cret}
}


// STBOXToTstzspan wraps MEOS C function stbox_to_tstzspan.
func STBOXToTstzspan(box *STBox) *Span {
	_cret := C.stbox_to_tstzspan(box._inner)
	return &Span{_inner: _cret}
}


// TimestamptzToSTBOX wraps MEOS C function timestamptz_to_stbox.
func TimestamptzToSTBOX(t int64) *STBox {
	_cret := C.timestamptz_to_stbox(C.TimestampTz(t))
	return &STBox{_inner: _cret}
}


// TstzsetToSTBOX wraps MEOS C function tstzset_to_stbox.
func TstzsetToSTBOX(s *Set) *STBox {
	_cret := C.tstzset_to_stbox(s._inner)
	return &STBox{_inner: _cret}
}


// TstzspanToSTBOX wraps MEOS C function tstzspan_to_stbox.
func TstzspanToSTBOX(s *Span) *STBox {
	_cret := C.tstzspan_to_stbox(s._inner)
	return &STBox{_inner: _cret}
}


// TstzspansetToSTBOX wraps MEOS C function tstzspanset_to_stbox.
func TstzspansetToSTBOX(ss *SpanSet) *STBox {
	_cret := C.tstzspanset_to_stbox(ss._inner)
	return &STBox{_inner: _cret}
}


// STBOXArea wraps MEOS C function stbox_area.
func STBOXArea(box *STBox, spheroid bool) float64 {
	_cret := C.stbox_area(box._inner, C.bool(spheroid))
	return float64(_cret)
}


// STBOXHash wraps MEOS C function stbox_hash.
func STBOXHash(box *STBox) uint32 {
	_cret := C.stbox_hash(box._inner)
	return uint32(_cret)
}


// STBOXHashExtended wraps MEOS C function stbox_hash_extended.
func STBOXHashExtended(box *STBox, seed uint64) uint64 {
	_cret := C.stbox_hash_extended(box._inner, C.uint64_t(seed))
	return uint64(_cret)
}


// STBOXHast wraps MEOS C function stbox_hast.
func STBOXHast(box *STBox) bool {
	_cret := C.stbox_hast(box._inner)
	return bool(_cret)
}


// STBOXHasx wraps MEOS C function stbox_hasx.
func STBOXHasx(box *STBox) bool {
	_cret := C.stbox_hasx(box._inner)
	return bool(_cret)
}


// STBOXHasz wraps MEOS C function stbox_hasz.
func STBOXHasz(box *STBox) bool {
	_cret := C.stbox_hasz(box._inner)
	return bool(_cret)
}


// STBOXIsgeodetic wraps MEOS C function stbox_isgeodetic.
func STBOXIsgeodetic(box *STBox) bool {
	_cret := C.stbox_isgeodetic(box._inner)
	return bool(_cret)
}


// STBOXPerimeter wraps MEOS C function stbox_perimeter.
func STBOXPerimeter(box *STBox, spheroid bool) float64 {
	_cret := C.stbox_perimeter(box._inner, C.bool(spheroid))
	return float64(_cret)
}


// STBOXTmax wraps MEOS C function stbox_tmax.
func STBOXTmax(box *STBox) (bool, int64) {
	var _out_result C.TimestampTz
	_cret := C.stbox_tmax(box._inner, &_out_result)
	return bool(_cret), int64(_out_result)
}


// STBOXTmaxInc wraps MEOS C function stbox_tmax_inc.
func STBOXTmaxInc(box *STBox) (bool, bool) {
	var _out_result C.bool
	_cret := C.stbox_tmax_inc(box._inner, &_out_result)
	return bool(_cret), bool(_out_result)
}


// STBOXTmin wraps MEOS C function stbox_tmin.
func STBOXTmin(box *STBox) (bool, int64) {
	var _out_result C.TimestampTz
	_cret := C.stbox_tmin(box._inner, &_out_result)
	return bool(_cret), int64(_out_result)
}


// STBOXTminInc wraps MEOS C function stbox_tmin_inc.
func STBOXTminInc(box *STBox) (bool, bool) {
	var _out_result C.bool
	_cret := C.stbox_tmin_inc(box._inner, &_out_result)
	return bool(_cret), bool(_out_result)
}


// STBOXVolume wraps MEOS C function stbox_volume.
func STBOXVolume(box *STBox) float64 {
	_cret := C.stbox_volume(box._inner)
	return float64(_cret)
}


// STBOXXmax wraps MEOS C function stbox_xmax.
func STBOXXmax(box *STBox) (bool, float64) {
	var _out_result C.double
	_cret := C.stbox_xmax(box._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// STBOXXmin wraps MEOS C function stbox_xmin.
func STBOXXmin(box *STBox) (bool, float64) {
	var _out_result C.double
	_cret := C.stbox_xmin(box._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// STBOXYmax wraps MEOS C function stbox_ymax.
func STBOXYmax(box *STBox) (bool, float64) {
	var _out_result C.double
	_cret := C.stbox_ymax(box._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// STBOXYmin wraps MEOS C function stbox_ymin.
func STBOXYmin(box *STBox) (bool, float64) {
	var _out_result C.double
	_cret := C.stbox_ymin(box._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// STBOXZmax wraps MEOS C function stbox_zmax.
func STBOXZmax(box *STBox) (bool, float64) {
	var _out_result C.double
	_cret := C.stbox_zmax(box._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// STBOXZmin wraps MEOS C function stbox_zmin.
func STBOXZmin(box *STBox) (bool, float64) {
	var _out_result C.double
	_cret := C.stbox_zmin(box._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// STBOXExpandSpace wraps MEOS C function stbox_expand_space.
func STBOXExpandSpace(box *STBox, d float64) *STBox {
	_cret := C.stbox_expand_space(box._inner, C.double(d))
	return &STBox{_inner: _cret}
}


// STBOXExpandTime wraps MEOS C function stbox_expand_time.
func STBOXExpandTime(box *STBox, interv *Interval) *STBox {
	_cret := C.stbox_expand_time(box._inner, interv._inner)
	return &STBox{_inner: _cret}
}


// STBOXGetSpace wraps MEOS C function stbox_get_space.
func STBOXGetSpace(box *STBox) *STBox {
	_cret := C.stbox_get_space(box._inner)
	return &STBox{_inner: _cret}
}


// STBOXQuadSplit wraps MEOS C function stbox_quad_split.
func STBOXQuadSplit(box *STBox, count unsafe.Pointer) *STBox {
	_cret := C.stbox_quad_split(box._inner, (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// STBOXRound wraps MEOS C function stbox_round.
func STBOXRound(box *STBox, maxdd int) *STBox {
	_cret := C.stbox_round(box._inner, C.int(maxdd))
	return &STBox{_inner: _cret}
}


// STBOXShiftScaleTime wraps MEOS C function stbox_shift_scale_time.
func STBOXShiftScaleTime(box *STBox, shift *Interval, duration *Interval) *STBox {
	_cret := C.stbox_shift_scale_time(box._inner, shift._inner, duration._inner)
	return &STBox{_inner: _cret}
}


// StboxarrRound wraps MEOS C function stboxarr_round.
func StboxarrRound(boxarr *STBox, count int, maxdd int) *STBox {
	_cret := C.stboxarr_round(boxarr._inner, C.int(count), C.int(maxdd))
	return &STBox{_inner: _cret}
}


// STBOXSetSRID wraps MEOS C function stbox_set_srid.
func STBOXSetSRID(box *STBox, srid int32) *STBox {
	_cret := C.stbox_set_srid(box._inner, C.int32_t(srid))
	return &STBox{_inner: _cret}
}


// STBOXSRID wraps MEOS C function stbox_srid.
func STBOXSRID(box *STBox) int32 {
	_cret := C.stbox_srid(box._inner)
	return int32(_cret)
}


// STBOXTransform wraps MEOS C function stbox_transform.
func STBOXTransform(box *STBox, srid int32) *STBox {
	_cret := C.stbox_transform(box._inner, C.int32_t(srid))
	return &STBox{_inner: _cret}
}


// STBOXTransformPipeline wraps MEOS C function stbox_transform_pipeline.
func STBOXTransformPipeline(box *STBox, pipelinestr string, srid int32, is_forward bool) *STBox {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	_cret := C.stbox_transform_pipeline(box._inner, _c_pipelinestr, C.int32_t(srid), C.bool(is_forward))
	return &STBox{_inner: _cret}
}


// AdjacentSTBOXSTBOX wraps MEOS C function adjacent_stbox_stbox.
func AdjacentSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.adjacent_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// ContainedSTBOXSTBOX wraps MEOS C function contained_stbox_stbox.
func ContainedSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.contained_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// ContainsSTBOXSTBOX wraps MEOS C function contains_stbox_stbox.
func ContainsSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.contains_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverlapsSTBOXSTBOX wraps MEOS C function overlaps_stbox_stbox.
func OverlapsSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.overlaps_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// SameSTBOXSTBOX wraps MEOS C function same_stbox_stbox.
func SameSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.same_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// AboveSTBOXSTBOX wraps MEOS C function above_stbox_stbox.
func AboveSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.above_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// AfterSTBOXSTBOX wraps MEOS C function after_stbox_stbox.
func AfterSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.after_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// BackSTBOXSTBOX wraps MEOS C function back_stbox_stbox.
func BackSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.back_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// BeforeSTBOXSTBOX wraps MEOS C function before_stbox_stbox.
func BeforeSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.before_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// BelowSTBOXSTBOX wraps MEOS C function below_stbox_stbox.
func BelowSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.below_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// FrontSTBOXSTBOX wraps MEOS C function front_stbox_stbox.
func FrontSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.front_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// LeftSTBOXSTBOX wraps MEOS C function left_stbox_stbox.
func LeftSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.left_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OveraboveSTBOXSTBOX wraps MEOS C function overabove_stbox_stbox.
func OveraboveSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.overabove_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverafterSTBOXSTBOX wraps MEOS C function overafter_stbox_stbox.
func OverafterSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.overafter_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverbackSTBOXSTBOX wraps MEOS C function overback_stbox_stbox.
func OverbackSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.overback_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverbeforeSTBOXSTBOX wraps MEOS C function overbefore_stbox_stbox.
func OverbeforeSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.overbefore_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverbelowSTBOXSTBOX wraps MEOS C function overbelow_stbox_stbox.
func OverbelowSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.overbelow_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverfrontSTBOXSTBOX wraps MEOS C function overfront_stbox_stbox.
func OverfrontSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.overfront_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverleftSTBOXSTBOX wraps MEOS C function overleft_stbox_stbox.
func OverleftSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.overleft_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverrightSTBOXSTBOX wraps MEOS C function overright_stbox_stbox.
func OverrightSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.overright_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// RightSTBOXSTBOX wraps MEOS C function right_stbox_stbox.
func RightSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	_cret := C.right_stbox_stbox(box1._inner, box2._inner)
	return bool(_cret)
}


// UnionSTBOXSTBOX wraps MEOS C function union_stbox_stbox.
func UnionSTBOXSTBOX(box1 *STBox, box2 *STBox, strict bool) *STBox {
	_cret := C.gunion_stbox_stbox(box1._inner, box2._inner, C.bool(strict))
	return &STBox{_inner: _cret}
}


// IntersectionSTBOXSTBOX wraps MEOS C function intersection_stbox_stbox.
func IntersectionSTBOXSTBOX(box1 *STBox, box2 *STBox) *STBox {
	_cret := C.intersection_stbox_stbox(box1._inner, box2._inner)
	return &STBox{_inner: _cret}
}


// STBOXCmp wraps MEOS C function stbox_cmp.
func STBOXCmp(box1 *STBox, box2 *STBox) int {
	_cret := C.stbox_cmp(box1._inner, box2._inner)
	return int(_cret)
}


// STBOXEq wraps MEOS C function stbox_eq.
func STBOXEq(box1 *STBox, box2 *STBox) bool {
	_cret := C.stbox_eq(box1._inner, box2._inner)
	return bool(_cret)
}


// STBOXGe wraps MEOS C function stbox_ge.
func STBOXGe(box1 *STBox, box2 *STBox) bool {
	_cret := C.stbox_ge(box1._inner, box2._inner)
	return bool(_cret)
}


// STBOXGt wraps MEOS C function stbox_gt.
func STBOXGt(box1 *STBox, box2 *STBox) bool {
	_cret := C.stbox_gt(box1._inner, box2._inner)
	return bool(_cret)
}


// STBOXLe wraps MEOS C function stbox_le.
func STBOXLe(box1 *STBox, box2 *STBox) bool {
	_cret := C.stbox_le(box1._inner, box2._inner)
	return bool(_cret)
}


// STBOXLt wraps MEOS C function stbox_lt.
func STBOXLt(box1 *STBox, box2 *STBox) bool {
	_cret := C.stbox_lt(box1._inner, box2._inner)
	return bool(_cret)
}


// STBOXNe wraps MEOS C function stbox_ne.
func STBOXNe(box1 *STBox, box2 *STBox) bool {
	_cret := C.stbox_ne(box1._inner, box2._inner)
	return bool(_cret)
}


// TspatialOut wraps MEOS C function tspatial_out.
func TspatialOut(temp *Temporal, maxdd int) string {
	_cret := C.tspatial_out(temp._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// TgeogpointFromMFJSON wraps MEOS C function tgeogpoint_from_mfjson.
func TgeogpointFromMFJSON(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeogpoint_from_mfjson(_c_str)
	return &Temporal{_inner: _cret}
}


// TgeogpointIn wraps MEOS C function tgeogpoint_in.
func TgeogpointIn(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeogpoint_in(_c_str)
	return &Temporal{_inner: _cret}
}


// TgeographyFromMFJSON wraps MEOS C function tgeography_from_mfjson.
func TgeographyFromMFJSON(mfjson string) *Temporal {
	_c_mfjson := C.CString(mfjson)
	defer C.free(unsafe.Pointer(_c_mfjson))
	_cret := C.tgeography_from_mfjson(_c_mfjson)
	return &Temporal{_inner: _cret}
}


// TgeographyIn wraps MEOS C function tgeography_in.
func TgeographyIn(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeography_in(_c_str)
	return &Temporal{_inner: _cret}
}


// TgeometryFromMFJSON wraps MEOS C function tgeometry_from_mfjson.
func TgeometryFromMFJSON(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeometry_from_mfjson(_c_str)
	return &Temporal{_inner: _cret}
}


// TgeometryIn wraps MEOS C function tgeometry_in.
func TgeometryIn(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeometry_in(_c_str)
	return &Temporal{_inner: _cret}
}


// TgeompointFromMFJSON wraps MEOS C function tgeompoint_from_mfjson.
func TgeompointFromMFJSON(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeompoint_from_mfjson(_c_str)
	return &Temporal{_inner: _cret}
}


// TgeompointIn wraps MEOS C function tgeompoint_in.
func TgeompointIn(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeompoint_in(_c_str)
	return &Temporal{_inner: _cret}
}


// TspatialAsEWKT wraps MEOS C function tspatial_as_ewkt.
func TspatialAsEWKT(temp *Temporal, maxdd int) string {
	_cret := C.tspatial_as_ewkt(temp._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// TspatialAsText wraps MEOS C function tspatial_as_text.
func TspatialAsText(temp *Temporal, maxdd int) string {
	_cret := C.tspatial_as_text(temp._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// TgeoFromBaseTemp wraps MEOS C function tgeo_from_base_temp.
func TgeoFromBaseTemp(gs *Geom, temp *Temporal) *Temporal {
	_cret := C.tgeo_from_base_temp(gs._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TgeoinstMake wraps MEOS C function tgeoinst_make.
func TgeoinstMake(gs *Geom, t int64) *TInstant {
	_cret := C.tgeoinst_make(gs._inner, C.TimestampTz(t))
	return &TInstant{_inner: _cret}
}


// TgeoseqFromBaseTstzset wraps MEOS C function tgeoseq_from_base_tstzset.
func TgeoseqFromBaseTstzset(gs *Geom, s *Set) *TSequence {
	_cret := C.tgeoseq_from_base_tstzset(gs._inner, s._inner)
	return &TSequence{_inner: _cret}
}


// TgeoseqFromBaseTstzspan wraps MEOS C function tgeoseq_from_base_tstzspan.
func TgeoseqFromBaseTstzspan(gs *Geom, s *Span, interp Interpolation) *TSequence {
	_cret := C.tgeoseq_from_base_tstzspan(gs._inner, s._inner, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TgeoseqsetFromBaseTstzspanset wraps MEOS C function tgeoseqset_from_base_tstzspanset.
func TgeoseqsetFromBaseTstzspanset(gs *Geom, ss *SpanSet, interp Interpolation) *TSequenceSet {
	_cret := C.tgeoseqset_from_base_tstzspanset(gs._inner, ss._inner, C.interpType(interp))
	return &TSequenceSet{_inner: _cret}
}


// TpointFromBaseTemp wraps MEOS C function tpoint_from_base_temp.
func TpointFromBaseTemp(gs *Geom, temp *Temporal) *Temporal {
	_cret := C.tpoint_from_base_temp(gs._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TpointinstMake wraps MEOS C function tpointinst_make.
func TpointinstMake(gs *Geom, t int64) *TInstant {
	_cret := C.tpointinst_make(gs._inner, C.TimestampTz(t))
	return &TInstant{_inner: _cret}
}


// TpointseqFromBaseTstzset wraps MEOS C function tpointseq_from_base_tstzset.
func TpointseqFromBaseTstzset(gs *Geom, s *Set) *TSequence {
	_cret := C.tpointseq_from_base_tstzset(gs._inner, s._inner)
	return &TSequence{_inner: _cret}
}


// TpointseqFromBaseTstzspan wraps MEOS C function tpointseq_from_base_tstzspan.
func TpointseqFromBaseTstzspan(gs *Geom, s *Span, interp Interpolation) *TSequence {
	_cret := C.tpointseq_from_base_tstzspan(gs._inner, s._inner, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TpointseqMakeCoords wraps MEOS C function tpointseq_make_coords.
func TpointseqMakeCoords(xcoords unsafe.Pointer, ycoords unsafe.Pointer, zcoords unsafe.Pointer, times unsafe.Pointer, count int, srid int32, geodetic bool, lower_inc bool, upper_inc bool, interp Interpolation, normalize bool) *TSequence {
	_cret := C.tpointseq_make_coords((*C.double)(unsafe.Pointer(xcoords)), (*C.double)(unsafe.Pointer(ycoords)), (*C.double)(unsafe.Pointer(zcoords)), (*C.TimestampTz)(unsafe.Pointer(times)), C.int(count), C.int32_t(srid), C.bool(geodetic), C.bool(lower_inc), C.bool(upper_inc), C.interpType(interp), C.bool(normalize))
	return &TSequence{_inner: _cret}
}


// TpointseqsetFromBaseTstzspanset wraps MEOS C function tpointseqset_from_base_tstzspanset.
func TpointseqsetFromBaseTstzspanset(gs *Geom, ss *SpanSet, interp Interpolation) *TSequenceSet {
	_cret := C.tpointseqset_from_base_tstzspanset(gs._inner, ss._inner, C.interpType(interp))
	return &TSequenceSet{_inner: _cret}
}


// Box3dToSTBOX wraps MEOS C function box3d_to_stbox.
func Box3dToSTBOX(box *Box3D) *STBox {
	_cret := C.box3d_to_stbox(box._inner)
	return &STBox{_inner: _cret}
}


// GboxToSTBOX wraps MEOS C function gbox_to_stbox.
func GboxToSTBOX(box *GBox) *STBox {
	_cret := C.gbox_to_stbox(box._inner)
	return &STBox{_inner: _cret}
}


// GeomeasToTpoint wraps MEOS C function geomeas_to_tpoint.
func GeomeasToTpoint(gs *Geom) *Temporal {
	_cret := C.geomeas_to_tpoint(gs._inner)
	return &Temporal{_inner: _cret}
}


// TgeogpointToTgeography wraps MEOS C function tgeogpoint_to_tgeography.
func TgeogpointToTgeography(temp *Temporal) *Temporal {
	_cret := C.tgeogpoint_to_tgeography(temp._inner)
	return &Temporal{_inner: _cret}
}


// TgeographyToTgeogpoint wraps MEOS C function tgeography_to_tgeogpoint.
func TgeographyToTgeogpoint(temp *Temporal) *Temporal {
	_cret := C.tgeography_to_tgeogpoint(temp._inner)
	return &Temporal{_inner: _cret}
}


// TgeographyToTgeometry wraps MEOS C function tgeography_to_tgeometry.
func TgeographyToTgeometry(temp *Temporal) *Temporal {
	_cret := C.tgeography_to_tgeometry(temp._inner)
	return &Temporal{_inner: _cret}
}


// TgeometryToTgeography wraps MEOS C function tgeometry_to_tgeography.
func TgeometryToTgeography(temp *Temporal) *Temporal {
	_cret := C.tgeometry_to_tgeography(temp._inner)
	return &Temporal{_inner: _cret}
}


// TgeometryToTgeompoint wraps MEOS C function tgeometry_to_tgeompoint.
func TgeometryToTgeompoint(temp *Temporal) *Temporal {
	_cret := C.tgeometry_to_tgeompoint(temp._inner)
	return &Temporal{_inner: _cret}
}


// TgeompointToTgeometry wraps MEOS C function tgeompoint_to_tgeometry.
func TgeompointToTgeometry(temp *Temporal) *Temporal {
	_cret := C.tgeompoint_to_tgeometry(temp._inner)
	return &Temporal{_inner: _cret}
}


// TODO tpoint_as_mvtgeom: unsupported return type MvtGeom
// func TpointAsMvtgeom(...) { /* not yet handled by codegen */ }


// TpointTfloatToGeomeas wraps MEOS C function tpoint_tfloat_to_geomeas.
func TpointTfloatToGeomeas(tpoint *Temporal, measure *Temporal, segmentize bool) (bool, *Geom) {
	var _out_result *C.GSERIALIZED
	_cret := C.tpoint_tfloat_to_geomeas(tpoint._inner, measure._inner, C.bool(segmentize), &_out_result)
	return bool(_cret), &Geom{_inner: _out_result}
}


// TspatialToSTBOX wraps MEOS C function tspatial_to_stbox.
func TspatialToSTBOX(temp *Temporal) *STBox {
	_cret := C.tspatial_to_stbox(temp._inner)
	return &STBox{_inner: _cret}
}


// BearingPointPoint wraps MEOS C function bearing_point_point.
func BearingPointPoint(gs1 *Geom, gs2 *Geom) (bool, float64) {
	var _out_result C.double
	_cret := C.bearing_point_point(gs1._inner, gs2._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// BearingTpointPoint wraps MEOS C function bearing_tpoint_point.
func BearingTpointPoint(temp *Temporal, gs *Geom, invert bool) *Temporal {
	_cret := C.bearing_tpoint_point(temp._inner, gs._inner, C.bool(invert))
	return &Temporal{_inner: _cret}
}


// BearingTpointTpoint wraps MEOS C function bearing_tpoint_tpoint.
func BearingTpointTpoint(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.bearing_tpoint_tpoint(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TgeoCentroid wraps MEOS C function tgeo_centroid.
func TgeoCentroid(temp *Temporal) *Temporal {
	_cret := C.tgeo_centroid(temp._inner)
	return &Temporal{_inner: _cret}
}


// TgeoConvexHull wraps MEOS C function tgeo_convex_hull.
func TgeoConvexHull(temp *Temporal) *Geom {
	_cret := C.tgeo_convex_hull(temp._inner)
	return &Geom{_inner: _cret}
}


// TgeoEndValue wraps MEOS C function tgeo_end_value.
func TgeoEndValue(temp *Temporal) *Geom {
	_cret := C.tgeo_end_value(temp._inner)
	return &Geom{_inner: _cret}
}


// TgeoStartValue wraps MEOS C function tgeo_start_value.
func TgeoStartValue(temp *Temporal) *Geom {
	_cret := C.tgeo_start_value(temp._inner)
	return &Geom{_inner: _cret}
}


// TgeoTraversedArea wraps MEOS C function tgeo_traversed_area.
func TgeoTraversedArea(temp *Temporal, unary_union bool) *Geom {
	_cret := C.tgeo_traversed_area(temp._inner, C.bool(unary_union))
	return &Geom{_inner: _cret}
}


// TgeoValueAtTimestamptz wraps MEOS C function tgeo_value_at_timestamptz.
func TgeoValueAtTimestamptz(temp *Temporal, t int64, strict bool) (bool, *Geom) {
	var _out_result *C.GSERIALIZED
	_cret := C.tgeo_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_result)
	return bool(_cret), &Geom{_inner: _out_result}
}


// TgeoValueN wraps MEOS C function tgeo_value_n.
func TgeoValueN(temp *Temporal, n int) (bool, *Geom) {
	var _out_result *C.GSERIALIZED
	_cret := C.tgeo_value_n(temp._inner, C.int(n), &_out_result)
	return bool(_cret), &Geom{_inner: _out_result}
}


// TgeoValues wraps MEOS C function tgeo_values.
func TgeoValues(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tgeo_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TpointAngularDifference wraps MEOS C function tpoint_angular_difference.
func TpointAngularDifference(temp *Temporal) *Temporal {
	_cret := C.tpoint_angular_difference(temp._inner)
	return &Temporal{_inner: _cret}
}


// TpointAzimuth wraps MEOS C function tpoint_azimuth.
func TpointAzimuth(temp *Temporal) *Temporal {
	_cret := C.tpoint_azimuth(temp._inner)
	return &Temporal{_inner: _cret}
}


// TpointCumulativeLength wraps MEOS C function tpoint_cumulative_length.
func TpointCumulativeLength(temp *Temporal) *Temporal {
	_cret := C.tpoint_cumulative_length(temp._inner)
	return &Temporal{_inner: _cret}
}


// TpointDirection wraps MEOS C function tpoint_direction.
func TpointDirection(temp *Temporal) (bool, float64) {
	var _out_result C.double
	_cret := C.tpoint_direction(temp._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// TpointGetX wraps MEOS C function tpoint_get_x.
func TpointGetX(temp *Temporal) *Temporal {
	_cret := C.tpoint_get_x(temp._inner)
	return &Temporal{_inner: _cret}
}


// TpointGetY wraps MEOS C function tpoint_get_y.
func TpointGetY(temp *Temporal) *Temporal {
	_cret := C.tpoint_get_y(temp._inner)
	return &Temporal{_inner: _cret}
}


// TpointGetZ wraps MEOS C function tpoint_get_z.
func TpointGetZ(temp *Temporal) *Temporal {
	_cret := C.tpoint_get_z(temp._inner)
	return &Temporal{_inner: _cret}
}


// TpointIsSimple wraps MEOS C function tpoint_is_simple.
func TpointIsSimple(temp *Temporal) bool {
	_cret := C.tpoint_is_simple(temp._inner)
	return bool(_cret)
}


// TpointLength wraps MEOS C function tpoint_length.
func TpointLength(temp *Temporal) float64 {
	_cret := C.tpoint_length(temp._inner)
	return float64(_cret)
}


// TpointSpeed wraps MEOS C function tpoint_speed.
func TpointSpeed(temp *Temporal) *Temporal {
	_cret := C.tpoint_speed(temp._inner)
	return &Temporal{_inner: _cret}
}


// TpointTrajectory wraps MEOS C function tpoint_trajectory.
func TpointTrajectory(temp *Temporal, unary_union bool) *Geom {
	_cret := C.tpoint_trajectory(temp._inner, C.bool(unary_union))
	return &Geom{_inner: _cret}
}


// TpointTwcentroid wraps MEOS C function tpoint_twcentroid.
func TpointTwcentroid(temp *Temporal) *Geom {
	_cret := C.tpoint_twcentroid(temp._inner)
	return &Geom{_inner: _cret}
}


// TgeoAffine wraps MEOS C function tgeo_affine.
func TgeoAffine(temp *Temporal, a *AFFINE) *Temporal {
	_cret := C.tgeo_affine(temp._inner, a._inner)
	return &Temporal{_inner: _cret}
}


// TgeoScale wraps MEOS C function tgeo_scale.
func TgeoScale(temp *Temporal, scale *Geom, sorigin *Geom) *Temporal {
	_cret := C.tgeo_scale(temp._inner, scale._inner, sorigin._inner)
	return &Temporal{_inner: _cret}
}


// TpointMakeSimple wraps MEOS C function tpoint_make_simple.
func TpointMakeSimple(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tpoint_make_simple(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TspatialSRID wraps MEOS C function tspatial_srid.
func TspatialSRID(temp *Temporal) int32 {
	_cret := C.tspatial_srid(temp._inner)
	return int32(_cret)
}


// TspatialSetSRID wraps MEOS C function tspatial_set_srid.
func TspatialSetSRID(temp *Temporal, srid int32) *Temporal {
	_cret := C.tspatial_set_srid(temp._inner, C.int32_t(srid))
	return &Temporal{_inner: _cret}
}


// TspatialTransform wraps MEOS C function tspatial_transform.
func TspatialTransform(temp *Temporal, srid int32) *Temporal {
	_cret := C.tspatial_transform(temp._inner, C.int32_t(srid))
	return &Temporal{_inner: _cret}
}


// TspatialTransformPipeline wraps MEOS C function tspatial_transform_pipeline.
func TspatialTransformPipeline(temp *Temporal, pipelinestr string, srid int32, is_forward bool) *Temporal {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	_cret := C.tspatial_transform_pipeline(temp._inner, _c_pipelinestr, C.int32_t(srid), C.bool(is_forward))
	return &Temporal{_inner: _cret}
}


// TgeoAtGeom wraps MEOS C function tgeo_at_geom.
func TgeoAtGeom(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tgeo_at_geom(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TgeoAtSTBOX wraps MEOS C function tgeo_at_stbox.
func TgeoAtSTBOX(temp *Temporal, box *STBox, border_inc bool) *Temporal {
	_cret := C.tgeo_at_stbox(temp._inner, box._inner, C.bool(border_inc))
	return &Temporal{_inner: _cret}
}


// TgeoAtValue wraps MEOS C function tgeo_at_value.
func TgeoAtValue(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tgeo_at_value(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TgeoMinusGeom wraps MEOS C function tgeo_minus_geom.
func TgeoMinusGeom(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tgeo_minus_geom(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TgeoMinusSTBOX wraps MEOS C function tgeo_minus_stbox.
func TgeoMinusSTBOX(temp *Temporal, box *STBox, border_inc bool) *Temporal {
	_cret := C.tgeo_minus_stbox(temp._inner, box._inner, C.bool(border_inc))
	return &Temporal{_inner: _cret}
}


// TgeoMinusValue wraps MEOS C function tgeo_minus_value.
func TgeoMinusValue(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tgeo_minus_value(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TpointAtElevation wraps MEOS C function tpoint_at_elevation.
func TpointAtElevation(temp *Temporal, s *Span) *Temporal {
	_cret := C.tpoint_at_elevation(temp._inner, s._inner)
	return &Temporal{_inner: _cret}
}


// TpointAtGeom wraps MEOS C function tpoint_at_geom.
func TpointAtGeom(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tpoint_at_geom(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TpointAtValue wraps MEOS C function tpoint_at_value.
func TpointAtValue(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tpoint_at_value(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TpointMinusElevation wraps MEOS C function tpoint_minus_elevation.
func TpointMinusElevation(temp *Temporal, s *Span) *Temporal {
	_cret := C.tpoint_minus_elevation(temp._inner, s._inner)
	return &Temporal{_inner: _cret}
}


// TpointMinusGeom wraps MEOS C function tpoint_minus_geom.
func TpointMinusGeom(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tpoint_minus_geom(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TpointMinusValue wraps MEOS C function tpoint_minus_value.
func TpointMinusValue(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tpoint_minus_value(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// AlwaysEqGeoTgeo wraps MEOS C function always_eq_geo_tgeo.
func AlwaysEqGeoTgeo(gs *Geom, temp *Temporal) int {
	_cret := C.always_eq_geo_tgeo(gs._inner, temp._inner)
	return int(_cret)
}


// AlwaysEqTgeoGeo wraps MEOS C function always_eq_tgeo_geo.
func AlwaysEqTgeoGeo(temp *Temporal, gs *Geom) int {
	_cret := C.always_eq_tgeo_geo(temp._inner, gs._inner)
	return int(_cret)
}


// AlwaysEqTgeoTgeo wraps MEOS C function always_eq_tgeo_tgeo.
func AlwaysEqTgeoTgeo(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_eq_tgeo_tgeo(temp1._inner, temp2._inner)
	return int(_cret)
}


// AlwaysNeGeoTgeo wraps MEOS C function always_ne_geo_tgeo.
func AlwaysNeGeoTgeo(gs *Geom, temp *Temporal) int {
	_cret := C.always_ne_geo_tgeo(gs._inner, temp._inner)
	return int(_cret)
}


// AlwaysNeTgeoGeo wraps MEOS C function always_ne_tgeo_geo.
func AlwaysNeTgeoGeo(temp *Temporal, gs *Geom) int {
	_cret := C.always_ne_tgeo_geo(temp._inner, gs._inner)
	return int(_cret)
}


// AlwaysNeTgeoTgeo wraps MEOS C function always_ne_tgeo_tgeo.
func AlwaysNeTgeoTgeo(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_ne_tgeo_tgeo(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverEqGeoTgeo wraps MEOS C function ever_eq_geo_tgeo.
func EverEqGeoTgeo(gs *Geom, temp *Temporal) int {
	_cret := C.ever_eq_geo_tgeo(gs._inner, temp._inner)
	return int(_cret)
}


// EverEqTgeoGeo wraps MEOS C function ever_eq_tgeo_geo.
func EverEqTgeoGeo(temp *Temporal, gs *Geom) int {
	_cret := C.ever_eq_tgeo_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EverEqTgeoTgeo wraps MEOS C function ever_eq_tgeo_tgeo.
func EverEqTgeoTgeo(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_eq_tgeo_tgeo(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverNeGeoTgeo wraps MEOS C function ever_ne_geo_tgeo.
func EverNeGeoTgeo(gs *Geom, temp *Temporal) int {
	_cret := C.ever_ne_geo_tgeo(gs._inner, temp._inner)
	return int(_cret)
}


// EverNeTgeoGeo wraps MEOS C function ever_ne_tgeo_geo.
func EverNeTgeoGeo(temp *Temporal, gs *Geom) int {
	_cret := C.ever_ne_tgeo_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EverNeTgeoTgeo wraps MEOS C function ever_ne_tgeo_tgeo.
func EverNeTgeoTgeo(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_ne_tgeo_tgeo(temp1._inner, temp2._inner)
	return int(_cret)
}


// TeqGeoTgeo wraps MEOS C function teq_geo_tgeo.
func TeqGeoTgeo(gs *Geom, temp *Temporal) *Temporal {
	_cret := C.teq_geo_tgeo(gs._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TeqTgeoGeo wraps MEOS C function teq_tgeo_geo.
func TeqTgeoGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.teq_tgeo_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TneGeoTgeo wraps MEOS C function tne_geo_tgeo.
func TneGeoTgeo(gs *Geom, temp *Temporal) *Temporal {
	_cret := C.tne_geo_tgeo(gs._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TneTgeoGeo wraps MEOS C function tne_tgeo_geo.
func TneTgeoGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tne_tgeo_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TgeoStboxes wraps MEOS C function tgeo_stboxes.
func TgeoStboxes(temp *Temporal, count unsafe.Pointer) *STBox {
	_cret := C.tgeo_stboxes(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// TgeoSpaceBoxes wraps MEOS C function tgeo_space_boxes.
func TgeoSpaceBoxes(temp *Temporal, xsize float64, ysize float64, zsize float64, sorigin *Geom, bitmatrix bool, border_inc bool, count unsafe.Pointer) *STBox {
	_cret := C.tgeo_space_boxes(temp._inner, C.double(xsize), C.double(ysize), C.double(zsize), sorigin._inner, C.bool(bitmatrix), C.bool(border_inc), (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// TgeoSpaceTimeBoxes wraps MEOS C function tgeo_space_time_boxes.
func TgeoSpaceTimeBoxes(temp *Temporal, xsize float64, ysize float64, zsize float64, duration *Interval, sorigin *Geom, torigin int64, bitmatrix bool, border_inc bool, count unsafe.Pointer) *STBox {
	_cret := C.tgeo_space_time_boxes(temp._inner, C.double(xsize), C.double(ysize), C.double(zsize), duration._inner, sorigin._inner, C.TimestampTz(torigin), C.bool(bitmatrix), C.bool(border_inc), (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// TgeoSplitEachNStboxes wraps MEOS C function tgeo_split_each_n_stboxes.
func TgeoSplitEachNStboxes(temp *Temporal, elem_count int, count unsafe.Pointer) *STBox {
	_cret := C.tgeo_split_each_n_stboxes(temp._inner, C.int(elem_count), (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// TgeoSplitNStboxes wraps MEOS C function tgeo_split_n_stboxes.
func TgeoSplitNStboxes(temp *Temporal, box_count int, count unsafe.Pointer) *STBox {
	_cret := C.tgeo_split_n_stboxes(temp._inner, C.int(box_count), (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// AdjacentSTBOXTspatial wraps MEOS C function adjacent_stbox_tspatial.
func AdjacentSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.adjacent_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// AdjacentTspatialSTBOX wraps MEOS C function adjacent_tspatial_stbox.
func AdjacentTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.adjacent_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// AdjacentTspatialTspatial wraps MEOS C function adjacent_tspatial_tspatial.
func AdjacentTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.adjacent_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// ContainedSTBOXTspatial wraps MEOS C function contained_stbox_tspatial.
func ContainedSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.contained_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// ContainedTspatialSTBOX wraps MEOS C function contained_tspatial_stbox.
func ContainedTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.contained_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// ContainedTspatialTspatial wraps MEOS C function contained_tspatial_tspatial.
func ContainedTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.contained_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// ContainsSTBOXTspatial wraps MEOS C function contains_stbox_tspatial.
func ContainsSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.contains_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// ContainsTspatialSTBOX wraps MEOS C function contains_tspatial_stbox.
func ContainsTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.contains_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// ContainsTspatialTspatial wraps MEOS C function contains_tspatial_tspatial.
func ContainsTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.contains_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// OverlapsSTBOXTspatial wraps MEOS C function overlaps_stbox_tspatial.
func OverlapsSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.overlaps_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// OverlapsTspatialSTBOX wraps MEOS C function overlaps_tspatial_stbox.
func OverlapsTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.overlaps_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// OverlapsTspatialTspatial wraps MEOS C function overlaps_tspatial_tspatial.
func OverlapsTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.overlaps_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// SameSTBOXTspatial wraps MEOS C function same_stbox_tspatial.
func SameSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.same_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// SameTspatialSTBOX wraps MEOS C function same_tspatial_stbox.
func SameTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.same_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// SameTspatialTspatial wraps MEOS C function same_tspatial_tspatial.
func SameTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.same_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// AboveSTBOXTspatial wraps MEOS C function above_stbox_tspatial.
func AboveSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.above_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// AboveTspatialSTBOX wraps MEOS C function above_tspatial_stbox.
func AboveTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.above_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// AboveTspatialTspatial wraps MEOS C function above_tspatial_tspatial.
func AboveTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.above_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// AfterSTBOXTspatial wraps MEOS C function after_stbox_tspatial.
func AfterSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.after_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// AfterTspatialSTBOX wraps MEOS C function after_tspatial_stbox.
func AfterTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.after_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// AfterTspatialTspatial wraps MEOS C function after_tspatial_tspatial.
func AfterTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.after_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// BackSTBOXTspatial wraps MEOS C function back_stbox_tspatial.
func BackSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.back_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// BackTspatialSTBOX wraps MEOS C function back_tspatial_stbox.
func BackTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.back_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// BackTspatialTspatial wraps MEOS C function back_tspatial_tspatial.
func BackTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.back_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// BeforeSTBOXTspatial wraps MEOS C function before_stbox_tspatial.
func BeforeSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.before_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// BeforeTspatialSTBOX wraps MEOS C function before_tspatial_stbox.
func BeforeTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.before_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// BeforeTspatialTspatial wraps MEOS C function before_tspatial_tspatial.
func BeforeTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.before_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// BelowSTBOXTspatial wraps MEOS C function below_stbox_tspatial.
func BelowSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.below_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// BelowTspatialSTBOX wraps MEOS C function below_tspatial_stbox.
func BelowTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.below_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// BelowTspatialTspatial wraps MEOS C function below_tspatial_tspatial.
func BelowTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.below_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// FrontSTBOXTspatial wraps MEOS C function front_stbox_tspatial.
func FrontSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.front_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// FrontTspatialSTBOX wraps MEOS C function front_tspatial_stbox.
func FrontTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.front_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// FrontTspatialTspatial wraps MEOS C function front_tspatial_tspatial.
func FrontTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.front_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// LeftSTBOXTspatial wraps MEOS C function left_stbox_tspatial.
func LeftSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.left_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// LeftTspatialSTBOX wraps MEOS C function left_tspatial_stbox.
func LeftTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.left_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// LeftTspatialTspatial wraps MEOS C function left_tspatial_tspatial.
func LeftTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.left_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// OveraboveSTBOXTspatial wraps MEOS C function overabove_stbox_tspatial.
func OveraboveSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.overabove_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// OveraboveTspatialSTBOX wraps MEOS C function overabove_tspatial_stbox.
func OveraboveTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.overabove_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// OveraboveTspatialTspatial wraps MEOS C function overabove_tspatial_tspatial.
func OveraboveTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.overabove_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// OverafterSTBOXTspatial wraps MEOS C function overafter_stbox_tspatial.
func OverafterSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.overafter_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// OverafterTspatialSTBOX wraps MEOS C function overafter_tspatial_stbox.
func OverafterTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.overafter_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// OverafterTspatialTspatial wraps MEOS C function overafter_tspatial_tspatial.
func OverafterTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.overafter_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// OverbackSTBOXTspatial wraps MEOS C function overback_stbox_tspatial.
func OverbackSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.overback_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// OverbackTspatialSTBOX wraps MEOS C function overback_tspatial_stbox.
func OverbackTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.overback_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// OverbackTspatialTspatial wraps MEOS C function overback_tspatial_tspatial.
func OverbackTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.overback_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// OverbeforeSTBOXTspatial wraps MEOS C function overbefore_stbox_tspatial.
func OverbeforeSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.overbefore_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// OverbeforeTspatialSTBOX wraps MEOS C function overbefore_tspatial_stbox.
func OverbeforeTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.overbefore_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// OverbeforeTspatialTspatial wraps MEOS C function overbefore_tspatial_tspatial.
func OverbeforeTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.overbefore_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// OverbelowSTBOXTspatial wraps MEOS C function overbelow_stbox_tspatial.
func OverbelowSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.overbelow_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// OverbelowTspatialSTBOX wraps MEOS C function overbelow_tspatial_stbox.
func OverbelowTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.overbelow_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// OverbelowTspatialTspatial wraps MEOS C function overbelow_tspatial_tspatial.
func OverbelowTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.overbelow_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// OverfrontSTBOXTspatial wraps MEOS C function overfront_stbox_tspatial.
func OverfrontSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.overfront_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// OverfrontTspatialSTBOX wraps MEOS C function overfront_tspatial_stbox.
func OverfrontTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.overfront_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// OverfrontTspatialTspatial wraps MEOS C function overfront_tspatial_tspatial.
func OverfrontTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.overfront_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// OverleftSTBOXTspatial wraps MEOS C function overleft_stbox_tspatial.
func OverleftSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.overleft_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// OverleftTspatialSTBOX wraps MEOS C function overleft_tspatial_stbox.
func OverleftTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.overleft_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// OverleftTspatialTspatial wraps MEOS C function overleft_tspatial_tspatial.
func OverleftTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.overleft_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// OverrightSTBOXTspatial wraps MEOS C function overright_stbox_tspatial.
func OverrightSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.overright_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// OverrightTspatialSTBOX wraps MEOS C function overright_tspatial_stbox.
func OverrightTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.overright_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// OverrightTspatialTspatial wraps MEOS C function overright_tspatial_tspatial.
func OverrightTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.overright_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// RightSTBOXTspatial wraps MEOS C function right_stbox_tspatial.
func RightSTBOXTspatial(box *STBox, temp *Temporal) bool {
	_cret := C.right_stbox_tspatial(box._inner, temp._inner)
	return bool(_cret)
}


// RightTspatialSTBOX wraps MEOS C function right_tspatial_stbox.
func RightTspatialSTBOX(temp *Temporal, box *STBox) bool {
	_cret := C.right_tspatial_stbox(temp._inner, box._inner)
	return bool(_cret)
}


// RightTspatialTspatial wraps MEOS C function right_tspatial_tspatial.
func RightTspatialTspatial(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.right_tspatial_tspatial(temp1._inner, temp2._inner)
	return bool(_cret)
}


// AcontainsGeoTgeo wraps MEOS C function acontains_geo_tgeo.
func AcontainsGeoTgeo(gs *Geom, temp *Temporal) int {
	_cret := C.acontains_geo_tgeo(gs._inner, temp._inner)
	return int(_cret)
}


// AcontainsTgeoGeo wraps MEOS C function acontains_tgeo_geo.
func AcontainsTgeoGeo(temp *Temporal, gs *Geom) int {
	_cret := C.acontains_tgeo_geo(temp._inner, gs._inner)
	return int(_cret)
}


// AcontainsTgeoTgeo wraps MEOS C function acontains_tgeo_tgeo.
func AcontainsTgeoTgeo(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.acontains_tgeo_tgeo(temp1._inner, temp2._inner)
	return int(_cret)
}


// AcoversGeoTgeo wraps MEOS C function acovers_geo_tgeo.
func AcoversGeoTgeo(gs *Geom, temp *Temporal) int {
	_cret := C.acovers_geo_tgeo(gs._inner, temp._inner)
	return int(_cret)
}


// AcoversTgeoGeo wraps MEOS C function acovers_tgeo_geo.
func AcoversTgeoGeo(temp *Temporal, gs *Geom) int {
	_cret := C.acovers_tgeo_geo(temp._inner, gs._inner)
	return int(_cret)
}


// AcoversTgeoTgeo wraps MEOS C function acovers_tgeo_tgeo.
func AcoversTgeoTgeo(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.acovers_tgeo_tgeo(temp1._inner, temp2._inner)
	return int(_cret)
}


// AdisjointGeoTgeo wraps MEOS C function adisjoint_geo_tgeo.
func AdisjointGeoTgeo(gs *Geom, temp *Temporal) int {
	_cret := C.adisjoint_geo_tgeo(gs._inner, temp._inner)
	return int(_cret)
}


// AdisjointTgeoGeo wraps MEOS C function adisjoint_tgeo_geo.
func AdisjointTgeoGeo(temp *Temporal, gs *Geom) int {
	_cret := C.adisjoint_tgeo_geo(temp._inner, gs._inner)
	return int(_cret)
}


// AdisjointTgeoTgeo wraps MEOS C function adisjoint_tgeo_tgeo.
func AdisjointTgeoTgeo(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.adisjoint_tgeo_tgeo(temp1._inner, temp2._inner)
	return int(_cret)
}


// AdwithinGeoTgeo wraps MEOS C function adwithin_geo_tgeo.
func AdwithinGeoTgeo(gs *Geom, temp *Temporal, dist float64) int {
	_cret := C.adwithin_geo_tgeo(gs._inner, temp._inner, C.double(dist))
	return int(_cret)
}


// AdwithinTgeoGeo wraps MEOS C function adwithin_tgeo_geo.
func AdwithinTgeoGeo(temp *Temporal, gs *Geom, dist float64) int {
	_cret := C.adwithin_tgeo_geo(temp._inner, gs._inner, C.double(dist))
	return int(_cret)
}


// AdwithinTgeoTgeo wraps MEOS C function adwithin_tgeo_tgeo.
func AdwithinTgeoTgeo(temp1 *Temporal, temp2 *Temporal, dist float64) int {
	_cret := C.adwithin_tgeo_tgeo(temp1._inner, temp2._inner, C.double(dist))
	return int(_cret)
}


// AintersectsGeoTgeo wraps MEOS C function aintersects_geo_tgeo.
func AintersectsGeoTgeo(gs *Geom, temp *Temporal) int {
	_cret := C.aintersects_geo_tgeo(gs._inner, temp._inner)
	return int(_cret)
}


// AintersectsTgeoGeo wraps MEOS C function aintersects_tgeo_geo.
func AintersectsTgeoGeo(temp *Temporal, gs *Geom) int {
	_cret := C.aintersects_tgeo_geo(temp._inner, gs._inner)
	return int(_cret)
}


// AintersectsTgeoTgeo wraps MEOS C function aintersects_tgeo_tgeo.
func AintersectsTgeoTgeo(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.aintersects_tgeo_tgeo(temp1._inner, temp2._inner)
	return int(_cret)
}


// AtouchesGeoTgeo wraps MEOS C function atouches_geo_tgeo.
func AtouchesGeoTgeo(gs *Geom, temp *Temporal) int {
	_cret := C.atouches_geo_tgeo(gs._inner, temp._inner)
	return int(_cret)
}


// AtouchesTgeoGeo wraps MEOS C function atouches_tgeo_geo.
func AtouchesTgeoGeo(temp *Temporal, gs *Geom) int {
	_cret := C.atouches_tgeo_geo(temp._inner, gs._inner)
	return int(_cret)
}


// AtouchesTgeoTgeo wraps MEOS C function atouches_tgeo_tgeo.
func AtouchesTgeoTgeo(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.atouches_tgeo_tgeo(temp1._inner, temp2._inner)
	return int(_cret)
}


// AtouchesTpointGeo wraps MEOS C function atouches_tpoint_geo.
func AtouchesTpointGeo(temp *Temporal, gs *Geom) int {
	_cret := C.atouches_tpoint_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EcontainsGeoTgeo wraps MEOS C function econtains_geo_tgeo.
func EcontainsGeoTgeo(gs *Geom, temp *Temporal) int {
	_cret := C.econtains_geo_tgeo(gs._inner, temp._inner)
	return int(_cret)
}


// EcontainsTgeoGeo wraps MEOS C function econtains_tgeo_geo.
func EcontainsTgeoGeo(temp *Temporal, gs *Geom) int {
	_cret := C.econtains_tgeo_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EcontainsTgeoTgeo wraps MEOS C function econtains_tgeo_tgeo.
func EcontainsTgeoTgeo(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.econtains_tgeo_tgeo(temp1._inner, temp2._inner)
	return int(_cret)
}


// EcoversGeoTgeo wraps MEOS C function ecovers_geo_tgeo.
func EcoversGeoTgeo(gs *Geom, temp *Temporal) int {
	_cret := C.ecovers_geo_tgeo(gs._inner, temp._inner)
	return int(_cret)
}


// EcoversTgeoGeo wraps MEOS C function ecovers_tgeo_geo.
func EcoversTgeoGeo(temp *Temporal, gs *Geom) int {
	_cret := C.ecovers_tgeo_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EcoversTgeoTgeo wraps MEOS C function ecovers_tgeo_tgeo.
func EcoversTgeoTgeo(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ecovers_tgeo_tgeo(temp1._inner, temp2._inner)
	return int(_cret)
}


// EdisjointGeoTgeo wraps MEOS C function edisjoint_geo_tgeo.
func EdisjointGeoTgeo(gs *Geom, temp *Temporal) int {
	_cret := C.edisjoint_geo_tgeo(gs._inner, temp._inner)
	return int(_cret)
}


// EdisjointTgeoGeo wraps MEOS C function edisjoint_tgeo_geo.
func EdisjointTgeoGeo(temp *Temporal, gs *Geom) int {
	_cret := C.edisjoint_tgeo_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EdisjointTgeoTgeo wraps MEOS C function edisjoint_tgeo_tgeo.
func EdisjointTgeoTgeo(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.edisjoint_tgeo_tgeo(temp1._inner, temp2._inner)
	return int(_cret)
}


// EdwithinGeoTgeo wraps MEOS C function edwithin_geo_tgeo.
func EdwithinGeoTgeo(gs *Geom, temp *Temporal, dist float64) int {
	_cret := C.edwithin_geo_tgeo(gs._inner, temp._inner, C.double(dist))
	return int(_cret)
}


// EdwithinTgeoGeo wraps MEOS C function edwithin_tgeo_geo.
func EdwithinTgeoGeo(temp *Temporal, gs *Geom, dist float64) int {
	_cret := C.edwithin_tgeo_geo(temp._inner, gs._inner, C.double(dist))
	return int(_cret)
}


// EdwithinTgeoTgeo wraps MEOS C function edwithin_tgeo_tgeo.
func EdwithinTgeoTgeo(temp1 *Temporal, temp2 *Temporal, dist float64) int {
	_cret := C.edwithin_tgeo_tgeo(temp1._inner, temp2._inner, C.double(dist))
	return int(_cret)
}


// EintersectsGeoTgeo wraps MEOS C function eintersects_geo_tgeo.
func EintersectsGeoTgeo(gs *Geom, temp *Temporal) int {
	_cret := C.eintersects_geo_tgeo(gs._inner, temp._inner)
	return int(_cret)
}


// EintersectsTgeoGeo wraps MEOS C function eintersects_tgeo_geo.
func EintersectsTgeoGeo(temp *Temporal, gs *Geom) int {
	_cret := C.eintersects_tgeo_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EintersectsTgeoTgeo wraps MEOS C function eintersects_tgeo_tgeo.
func EintersectsTgeoTgeo(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.eintersects_tgeo_tgeo(temp1._inner, temp2._inner)
	return int(_cret)
}


// EtouchesGeoTgeo wraps MEOS C function etouches_geo_tgeo.
func EtouchesGeoTgeo(gs *Geom, temp *Temporal) int {
	_cret := C.etouches_geo_tgeo(gs._inner, temp._inner)
	return int(_cret)
}


// EtouchesTgeoGeo wraps MEOS C function etouches_tgeo_geo.
func EtouchesTgeoGeo(temp *Temporal, gs *Geom) int {
	_cret := C.etouches_tgeo_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EtouchesTgeoTgeo wraps MEOS C function etouches_tgeo_tgeo.
func EtouchesTgeoTgeo(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.etouches_tgeo_tgeo(temp1._inner, temp2._inner)
	return int(_cret)
}


// EtouchesTpointGeo wraps MEOS C function etouches_tpoint_geo.
func EtouchesTpointGeo(temp *Temporal, gs *Geom) int {
	_cret := C.etouches_tpoint_geo(temp._inner, gs._inner)
	return int(_cret)
}


// TcontainsGeoTgeo wraps MEOS C function tcontains_geo_tgeo.
func TcontainsGeoTgeo(gs *Geom, temp *Temporal) *Temporal {
	_cret := C.tcontains_geo_tgeo(gs._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TcontainsTgeoGeo wraps MEOS C function tcontains_tgeo_geo.
func TcontainsTgeoGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tcontains_tgeo_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TcontainsTgeoTgeo wraps MEOS C function tcontains_tgeo_tgeo.
func TcontainsTgeoTgeo(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tcontains_tgeo_tgeo(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TcoversGeoTgeo wraps MEOS C function tcovers_geo_tgeo.
func TcoversGeoTgeo(gs *Geom, temp *Temporal) *Temporal {
	_cret := C.tcovers_geo_tgeo(gs._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TcoversTgeoGeo wraps MEOS C function tcovers_tgeo_geo.
func TcoversTgeoGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tcovers_tgeo_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TcoversTgeoTgeo wraps MEOS C function tcovers_tgeo_tgeo.
func TcoversTgeoTgeo(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tcovers_tgeo_tgeo(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TdisjointGeoTgeo wraps MEOS C function tdisjoint_geo_tgeo.
func TdisjointGeoTgeo(gs *Geom, temp *Temporal) *Temporal {
	_cret := C.tdisjoint_geo_tgeo(gs._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TdisjointTgeoGeo wraps MEOS C function tdisjoint_tgeo_geo.
func TdisjointTgeoGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tdisjoint_tgeo_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TdisjointTgeoTgeo wraps MEOS C function tdisjoint_tgeo_tgeo.
func TdisjointTgeoTgeo(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tdisjoint_tgeo_tgeo(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TdwithinGeoTgeo wraps MEOS C function tdwithin_geo_tgeo.
func TdwithinGeoTgeo(gs *Geom, temp *Temporal, dist float64) *Temporal {
	_cret := C.tdwithin_geo_tgeo(gs._inner, temp._inner, C.double(dist))
	return &Temporal{_inner: _cret}
}


// TdwithinTgeoGeo wraps MEOS C function tdwithin_tgeo_geo.
func TdwithinTgeoGeo(temp *Temporal, gs *Geom, dist float64) *Temporal {
	_cret := C.tdwithin_tgeo_geo(temp._inner, gs._inner, C.double(dist))
	return &Temporal{_inner: _cret}
}


// TdwithinTgeoTgeo wraps MEOS C function tdwithin_tgeo_tgeo.
func TdwithinTgeoTgeo(temp1 *Temporal, temp2 *Temporal, dist float64) *Temporal {
	_cret := C.tdwithin_tgeo_tgeo(temp1._inner, temp2._inner, C.double(dist))
	return &Temporal{_inner: _cret}
}


// TintersectsGeoTgeo wraps MEOS C function tintersects_geo_tgeo.
func TintersectsGeoTgeo(gs *Geom, temp *Temporal) *Temporal {
	_cret := C.tintersects_geo_tgeo(gs._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TintersectsTgeoGeo wraps MEOS C function tintersects_tgeo_geo.
func TintersectsTgeoGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tintersects_tgeo_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TintersectsTgeoTgeo wraps MEOS C function tintersects_tgeo_tgeo.
func TintersectsTgeoTgeo(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tintersects_tgeo_tgeo(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TtouchesGeoTgeo wraps MEOS C function ttouches_geo_tgeo.
func TtouchesGeoTgeo(gs *Geom, temp *Temporal) *Temporal {
	_cret := C.ttouches_geo_tgeo(gs._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TtouchesTgeoGeo wraps MEOS C function ttouches_tgeo_geo.
func TtouchesTgeoGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.ttouches_tgeo_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TtouchesTgeoTgeo wraps MEOS C function ttouches_tgeo_tgeo.
func TtouchesTgeoTgeo(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.ttouches_tgeo_tgeo(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// EdwithinTgeoarrTgeoarr wraps MEOS C function edwithin_tgeoarr_tgeoarr.
func EdwithinTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, dist float64, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.edwithin_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), C.double(dist), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// AdwithinTgeoarrTgeoarr wraps MEOS C function adwithin_tgeoarr_tgeoarr.
func AdwithinTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, dist float64, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.adwithin_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), C.double(dist), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// EintersectsTgeoarrTgeoarr wraps MEOS C function eintersects_tgeoarr_tgeoarr.
func EintersectsTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.eintersects_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// AintersectsTgeoarrTgeoarr wraps MEOS C function aintersects_tgeoarr_tgeoarr.
func AintersectsTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.aintersects_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// EtouchesTgeoarrTgeoarr wraps MEOS C function etouches_tgeoarr_tgeoarr.
func EtouchesTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.etouches_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// AtouchesTgeoarrTgeoarr wraps MEOS C function atouches_tgeoarr_tgeoarr.
func AtouchesTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.atouches_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// EdisjointTgeoarrTgeoarr wraps MEOS C function edisjoint_tgeoarr_tgeoarr.
func EdisjointTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.edisjoint_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// AdisjointTgeoarrTgeoarr wraps MEOS C function adisjoint_tgeoarr_tgeoarr.
func AdisjointTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.adisjoint_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TdwithinTgeoarrTgeoarr wraps MEOS C function tdwithin_tgeoarr_tgeoarr.
func TdwithinTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, dist float64, count unsafe.Pointer, periods unsafe.Pointer) unsafe.Pointer {
	_cret := C.tdwithin_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), C.double(dist), (*C.int)(unsafe.Pointer(count)), (***C.SpanSet)(unsafe.Pointer(periods)))
	return unsafe.Pointer(_cret)
}


// TintersectsTgeoarrTgeoarr wraps MEOS C function tintersects_tgeoarr_tgeoarr.
func TintersectsTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, count unsafe.Pointer, periods unsafe.Pointer) unsafe.Pointer {
	_cret := C.tintersects_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), (*C.int)(unsafe.Pointer(count)), (***C.SpanSet)(unsafe.Pointer(periods)))
	return unsafe.Pointer(_cret)
}


// TtouchesTgeoarrTgeoarr wraps MEOS C function ttouches_tgeoarr_tgeoarr.
func TtouchesTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, count unsafe.Pointer, periods unsafe.Pointer) unsafe.Pointer {
	_cret := C.ttouches_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), (*C.int)(unsafe.Pointer(count)), (***C.SpanSet)(unsafe.Pointer(periods)))
	return unsafe.Pointer(_cret)
}


// TdisjointTgeoarrTgeoarr wraps MEOS C function tdisjoint_tgeoarr_tgeoarr.
func TdisjointTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, count unsafe.Pointer, periods unsafe.Pointer) unsafe.Pointer {
	_cret := C.tdisjoint_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), (*C.int)(unsafe.Pointer(count)), (***C.SpanSet)(unsafe.Pointer(periods)))
	return unsafe.Pointer(_cret)
}


// TdistanceTgeoGeo wraps MEOS C function tdistance_tgeo_geo.
func TdistanceTgeoGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tdistance_tgeo_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TdistanceTgeoTgeo wraps MEOS C function tdistance_tgeo_tgeo.
func TdistanceTgeoTgeo(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tdistance_tgeo_tgeo(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// NadSTBOXGeo wraps MEOS C function nad_stbox_geo.
func NadSTBOXGeo(box *STBox, gs *Geom) float64 {
	_cret := C.nad_stbox_geo(box._inner, gs._inner)
	return float64(_cret)
}


// NadSTBOXSTBOX wraps MEOS C function nad_stbox_stbox.
func NadSTBOXSTBOX(box1 *STBox, box2 *STBox) float64 {
	_cret := C.nad_stbox_stbox(box1._inner, box2._inner)
	return float64(_cret)
}


// STBOXSpatialDistance wraps MEOS C function stbox_spatial_distance.
func STBOXSpatialDistance(box1 *STBox, box2 *STBox) float64 {
	_cret := C.stbox_spatial_distance(box1._inner, box2._inner)
	return float64(_cret)
}


// NadTgeoGeo wraps MEOS C function nad_tgeo_geo.
func NadTgeoGeo(temp *Temporal, gs *Geom) float64 {
	_cret := C.nad_tgeo_geo(temp._inner, gs._inner)
	return float64(_cret)
}


// NadTgeoSTBOX wraps MEOS C function nad_tgeo_stbox.
func NadTgeoSTBOX(temp *Temporal, box *STBox) float64 {
	_cret := C.nad_tgeo_stbox(temp._inner, box._inner)
	return float64(_cret)
}


// NadTgeoTgeo wraps MEOS C function nad_tgeo_tgeo.
func NadTgeoTgeo(temp1 *Temporal, temp2 *Temporal) float64 {
	_cret := C.nad_tgeo_tgeo(temp1._inner, temp2._inner)
	return float64(_cret)
}


// NaiTgeoGeo wraps MEOS C function nai_tgeo_geo.
func NaiTgeoGeo(temp *Temporal, gs *Geom) *TInstant {
	_cret := C.nai_tgeo_geo(temp._inner, gs._inner)
	return &TInstant{_inner: _cret}
}


// NaiTgeoTgeo wraps MEOS C function nai_tgeo_tgeo.
func NaiTgeoTgeo(temp1 *Temporal, temp2 *Temporal) *TInstant {
	_cret := C.nai_tgeo_tgeo(temp1._inner, temp2._inner)
	return &TInstant{_inner: _cret}
}


// ShortestlineTgeoGeo wraps MEOS C function shortestline_tgeo_geo.
func ShortestlineTgeoGeo(temp *Temporal, gs *Geom) *Geom {
	_cret := C.shortestline_tgeo_geo(temp._inner, gs._inner)
	return &Geom{_inner: _cret}
}


// ShortestlineTgeoTgeo wraps MEOS C function shortestline_tgeo_tgeo.
func ShortestlineTgeoTgeo(temp1 *Temporal, temp2 *Temporal) *Geom {
	_cret := C.shortestline_tgeo_tgeo(temp1._inner, temp2._inner)
	return &Geom{_inner: _cret}
}


// MindistanceTgeoTgeo wraps MEOS C function mindistance_tgeo_tgeo.
func MindistanceTgeoTgeo(temp1 *Temporal, temp2 *Temporal, threshold float64) float64 {
	_cret := C.mindistance_tgeo_tgeo(temp1._inner, temp2._inner, C.double(threshold))
	return float64(_cret)
}


// MindistanceTgeoarrTgeoarr wraps MEOS C function mindistance_tgeoarr_tgeoarr.
func MindistanceTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int) float64 {
	_cret := C.mindistance_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2))
	return float64(_cret)
}


// TpointTcentroidFinalfn wraps MEOS C function tpoint_tcentroid_finalfn.
func TpointTcentroidFinalfn(state *SkipList) *Temporal {
	_cret := C.tpoint_tcentroid_finalfn(state._inner)
	return &Temporal{_inner: _cret}
}


// TpointTcentroidTransfn wraps MEOS C function tpoint_tcentroid_transfn.
func TpointTcentroidTransfn(state *SkipList, temp *Temporal) *SkipList {
	_cret := C.tpoint_tcentroid_transfn(state._inner, temp._inner)
	return &SkipList{_inner: _cret}
}


// TspatialExtentTransfn wraps MEOS C function tspatial_extent_transfn.
func TspatialExtentTransfn(box *STBox, temp *Temporal) *STBox {
	_cret := C.tspatial_extent_transfn(box._inner, temp._inner)
	return &STBox{_inner: _cret}
}


// STBOXGetSpaceTile wraps MEOS C function stbox_get_space_tile.
func STBOXGetSpaceTile(point *Geom, xsize float64, ysize float64, zsize float64, sorigin *Geom) *STBox {
	_cret := C.stbox_get_space_tile(point._inner, C.double(xsize), C.double(ysize), C.double(zsize), sorigin._inner)
	return &STBox{_inner: _cret}
}


// STBOXGetSpaceTimeTile wraps MEOS C function stbox_get_space_time_tile.
func STBOXGetSpaceTimeTile(point *Geom, t int64, xsize float64, ysize float64, zsize float64, duration *Interval, sorigin *Geom, torigin int64) *STBox {
	_cret := C.stbox_get_space_time_tile(point._inner, C.TimestampTz(t), C.double(xsize), C.double(ysize), C.double(zsize), duration._inner, sorigin._inner, C.TimestampTz(torigin))
	return &STBox{_inner: _cret}
}


// STBOXGetTimeTile wraps MEOS C function stbox_get_time_tile.
func STBOXGetTimeTile(t int64, duration *Interval, torigin int64) *STBox {
	_cret := C.stbox_get_time_tile(C.TimestampTz(t), duration._inner, C.TimestampTz(torigin))
	return &STBox{_inner: _cret}
}


// STBOXSpaceTiles wraps MEOS C function stbox_space_tiles.
func STBOXSpaceTiles(bounds *STBox, xsize float64, ysize float64, zsize float64, sorigin *Geom, border_inc bool, count unsafe.Pointer) *STBox {
	_cret := C.stbox_space_tiles(bounds._inner, C.double(xsize), C.double(ysize), C.double(zsize), sorigin._inner, C.bool(border_inc), (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// STBOXSpaceTimeTiles wraps MEOS C function stbox_space_time_tiles.
func STBOXSpaceTimeTiles(bounds *STBox, xsize float64, ysize float64, zsize float64, duration *Interval, sorigin *Geom, torigin int64, border_inc bool, count unsafe.Pointer) *STBox {
	_cret := C.stbox_space_time_tiles(bounds._inner, C.double(xsize), C.double(ysize), C.double(zsize), duration._inner, sorigin._inner, C.TimestampTz(torigin), C.bool(border_inc), (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// STBOXTimeTiles wraps MEOS C function stbox_time_tiles.
func STBOXTimeTiles(bounds *STBox, duration *Interval, torigin int64, border_inc bool, count unsafe.Pointer) *STBox {
	_cret := C.stbox_time_tiles(bounds._inner, duration._inner, C.TimestampTz(torigin), C.bool(border_inc), (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// TODO tgeo_space_split: unsupported return type SpaceSplit
// func TgeoSpaceSplit(...) { /* not yet handled by codegen */ }


// TODO tgeo_space_time_split: unsupported return type SpaceTimeSplit
// func TgeoSpaceTimeSplit(...) { /* not yet handled by codegen */ }


// GeoClusterKmeans wraps MEOS C function geo_cluster_kmeans.
func GeoClusterKmeans(geoms unsafe.Pointer, ngeoms uint32, k uint32, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.geo_cluster_kmeans((**C.GSERIALIZED)(unsafe.Pointer(geoms)), C.uint32_t(ngeoms), C.uint32_t(k), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// GeoClusterDbscan wraps MEOS C function geo_cluster_dbscan.
func GeoClusterDbscan(geoms unsafe.Pointer, ngeoms uint32, tolerance float64, minpoints int, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.geo_cluster_dbscan((**C.GSERIALIZED)(unsafe.Pointer(geoms)), C.uint32_t(ngeoms), C.double(tolerance), C.int(minpoints), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// GeoClusterIntersecting wraps MEOS C function geo_cluster_intersecting.
func GeoClusterIntersecting(geoms unsafe.Pointer, ngeoms uint32, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.geo_cluster_intersecting((**C.GSERIALIZED)(unsafe.Pointer(geoms)), C.uint32_t(ngeoms), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// GeoClusterWithin wraps MEOS C function geo_cluster_within.
func GeoClusterWithin(geoms unsafe.Pointer, ngeoms uint32, tolerance float64, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.geo_cluster_within((**C.GSERIALIZED)(unsafe.Pointer(geoms)), C.uint32_t(ngeoms), C.double(tolerance), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}

