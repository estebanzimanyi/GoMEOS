package generated

// #include <stddef.h>
import "C"
import (
	"unsafe"

	"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

// Box3dFromGbox wraps MEOS C function box3d_from_gbox.
func Box3dFromGbox(box *GBox) *Box3D {
	res := C.box3d_from_gbox(box._inner)
	return &Box3D{_inner: res}
}


// Box3dMake wraps MEOS C function box3d_make.
func Box3dMake(xmin float64, xmax float64, ymin float64, ymax float64, zmin float64, zmax float64, srid int32) *Box3D {
	res := C.box3d_make(C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.double(zmin), C.double(zmax), C.int32_t(srid))
	return &Box3D{_inner: res}
}


// Box3dIn wraps MEOS C function box3d_in.
func Box3dIn(str string) *Box3D {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.box3d_in(_c_str)
	return &Box3D{_inner: res}
}


// Box3dOut wraps MEOS C function box3d_out.
func Box3dOut(box *Box3D, maxdd int) string {
	res := C.box3d_out(box._inner, C.int(maxdd))
	return C.GoString(res)
}


// GboxMake wraps MEOS C function gbox_make.
func GboxMake(hasz bool, hasm bool, geodetic bool, xmin float64, xmax float64, ymin float64, ymax float64, zmin float64, zmax float64, mmin float64, mmax float64) *GBox {
	res := C.gbox_make(C.bool(hasz), C.bool(hasm), C.bool(geodetic), C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.double(zmin), C.double(zmax), C.double(mmin), C.double(mmax))
	return &GBox{_inner: res}
}


// GboxIn wraps MEOS C function gbox_in.
func GboxIn(str string) *GBox {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.gbox_in(_c_str)
	return &GBox{_inner: res}
}


// GboxOut wraps MEOS C function gbox_out.
func GboxOut(box *GBox, maxdd int) string {
	res := C.gbox_out(box._inner, C.int(maxdd))
	return C.GoString(res)
}


// GeoAsEWKB wraps MEOS C function geo_as_ewkb.
func GeoAsEWKB(gs *Geom, endian string) []uint8 {
	_c_endian := C.CString(endian)
	defer C.free(unsafe.Pointer(_c_endian))
	var _out_size C.size_t
	res := C.geo_as_ewkb(gs._inner, _c_endian, &_out_size)
	_n := int(_out_size)
	_slice := unsafe.Slice((*C.uint8_t)(unsafe.Pointer(res)), _n)
	_out := make([]uint8, _n)
	for _i, _e := range _slice {
		_out[_i] = uint8(_e)
	}
	return _out
}


// GeoAsEWKT wraps MEOS C function geo_as_ewkt.
func GeoAsEWKT(gs *Geom, precision int) string {
	res := C.geo_as_ewkt(gs._inner, C.int(precision))
	return C.GoString(res)
}


// GeoAsGeojson wraps MEOS C function geo_as_geojson.
func GeoAsGeojson(gs *Geom, option int, precision int, srs string) string {
	_c_srs := C.CString(srs)
	defer C.free(unsafe.Pointer(_c_srs))
	res := C.geo_as_geojson(gs._inner, C.int(option), C.int(precision), _c_srs)
	return C.GoString(res)
}


// GeoAsHexewkb wraps MEOS C function geo_as_hexewkb.
func GeoAsHexewkb(gs *Geom, endian string) string {
	_c_endian := C.CString(endian)
	defer C.free(unsafe.Pointer(_c_endian))
	res := C.geo_as_hexewkb(gs._inner, _c_endian)
	return C.GoString(res)
}


// GeoAsText wraps MEOS C function geo_as_text.
func GeoAsText(gs *Geom, precision int) string {
	res := C.geo_as_text(gs._inner, C.int(precision))
	return C.GoString(res)
}


// GeoFromEWKB wraps MEOS C function geo_from_ewkb.
func GeoFromEWKB(wkb []byte, srid int32) *Geom {
	res := C.geo_from_ewkb((*C.uint8_t)(unsafe.Pointer(&wkb[0])), C.size_t(len(wkb)), C.int32_t(srid))
	return &Geom{_inner: res}
}


// GeoFromGeojson wraps MEOS C function geo_from_geojson.
func GeoFromGeojson(geojson string) *Geom {
	_c_geojson := C.CString(geojson)
	defer C.free(unsafe.Pointer(_c_geojson))
	res := C.geo_from_geojson(_c_geojson)
	return &Geom{_inner: res}
}


// GeoFromText wraps MEOS C function geo_from_text.
func GeoFromText(wkt string, srid int32) *Geom {
	_c_wkt := C.CString(wkt)
	defer C.free(unsafe.Pointer(_c_wkt))
	res := C.geo_from_text(_c_wkt, C.int32_t(srid))
	return &Geom{_inner: res}
}


// GeoOut wraps MEOS C function geo_out.
func GeoOut(gs *Geom) string {
	res := C.geo_out(gs._inner)
	return C.GoString(res)
}


// GeogFromHexewkb wraps MEOS C function geog_from_hexewkb.
func GeogFromHexewkb(wkt string) *Geom {
	_c_wkt := C.CString(wkt)
	defer C.free(unsafe.Pointer(_c_wkt))
	res := C.geog_from_hexewkb(_c_wkt)
	return &Geom{_inner: res}
}


// GeogIn wraps MEOS C function geog_in.
func GeogIn(str string, typmod int) *Geom {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.geog_in(_c_str, C.int(typmod))
	return &Geom{_inner: res}
}


// GeomFromHexewkb wraps MEOS C function geom_from_hexewkb.
func GeomFromHexewkb(wkt string) *Geom {
	_c_wkt := C.CString(wkt)
	defer C.free(unsafe.Pointer(_c_wkt))
	res := C.geom_from_hexewkb(_c_wkt)
	return &Geom{_inner: res}
}


// GeomIn wraps MEOS C function geom_in.
func GeomIn(str string, typmod int) *Geom {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.geom_in(_c_str, C.int(typmod))
	return &Geom{_inner: res}
}


// GeoCopy wraps MEOS C function geo_copy.
func GeoCopy(gs *Geom) *Geom {
	res := C.geo_copy(gs._inner)
	return &Geom{_inner: res}
}


// GeogpointMake2d wraps MEOS C function geogpoint_make2d.
func GeogpointMake2d(srid int32, x float64, y float64) *Geom {
	res := C.geogpoint_make2d(C.int32_t(srid), C.double(x), C.double(y))
	return &Geom{_inner: res}
}


// GeogpointMake3dz wraps MEOS C function geogpoint_make3dz.
func GeogpointMake3dz(srid int32, x float64, y float64, z float64) *Geom {
	res := C.geogpoint_make3dz(C.int32_t(srid), C.double(x), C.double(y), C.double(z))
	return &Geom{_inner: res}
}


// GeompointMake2d wraps MEOS C function geompoint_make2d.
func GeompointMake2d(srid int32, x float64, y float64) *Geom {
	res := C.geompoint_make2d(C.int32_t(srid), C.double(x), C.double(y))
	return &Geom{_inner: res}
}


// GeompointMake3dz wraps MEOS C function geompoint_make3dz.
func GeompointMake3dz(srid int32, x float64, y float64, z float64) *Geom {
	res := C.geompoint_make3dz(C.int32_t(srid), C.double(x), C.double(y), C.double(z))
	return &Geom{_inner: res}
}


// GeomToGeog wraps MEOS C function geom_to_geog.
func GeomToGeog(geom *Geom) *Geom {
	res := C.geom_to_geog(geom._inner)
	return &Geom{_inner: res}
}


// GeogToGeom wraps MEOS C function geog_to_geom.
func GeogToGeom(geog *Geom) *Geom {
	res := C.geog_to_geom(geog._inner)
	return &Geom{_inner: res}
}


// GeoIsEmpty wraps MEOS C function geo_is_empty.
func GeoIsEmpty(gs *Geom) bool {
	res := C.geo_is_empty(gs._inner)
	return bool(res)
}


// GeoIsUnitary wraps MEOS C function geo_is_unitary.
func GeoIsUnitary(gs *Geom) bool {
	res := C.geo_is_unitary(gs._inner)
	return bool(res)
}


// GeoTypename wraps MEOS C function geo_typename.
func GeoTypename(type_ int) string {
	res := C.geo_typename(C.int(type_))
	return C.GoString(res)
}


// GeogArea wraps MEOS C function geog_area.
func GeogArea(gs *Geom, use_spheroid bool) float64 {
	res := C.geog_area(gs._inner, C.bool(use_spheroid))
	return float64(res)
}


// GeogCentroid wraps MEOS C function geog_centroid.
func GeogCentroid(gs *Geom, use_spheroid bool) *Geom {
	res := C.geog_centroid(gs._inner, C.bool(use_spheroid))
	return &Geom{_inner: res}
}


// GeogLength wraps MEOS C function geog_length.
func GeogLength(gs *Geom, use_spheroid bool) float64 {
	res := C.geog_length(gs._inner, C.bool(use_spheroid))
	return float64(res)
}


// GeogPerimeter wraps MEOS C function geog_perimeter.
func GeogPerimeter(gs *Geom, use_spheroid bool) float64 {
	res := C.geog_perimeter(gs._inner, C.bool(use_spheroid))
	return float64(res)
}


// GeomAzimuth wraps MEOS C function geom_azimuth.
func GeomAzimuth(gs1 *Geom, gs2 *Geom) (bool, float64) {
	var _out_result C.double
	res := C.geom_azimuth(gs1._inner, gs2._inner, &_out_result)
	return bool(res), float64(_out_result)
}


// GeomLength wraps MEOS C function geom_length.
func GeomLength(gs *Geom) float64 {
	res := C.geom_length(gs._inner)
	return float64(res)
}


// GeomPerimeter wraps MEOS C function geom_perimeter.
func GeomPerimeter(gs *Geom) float64 {
	res := C.geom_perimeter(gs._inner)
	return float64(res)
}


// LineNumpoints wraps MEOS C function line_numpoints.
func LineNumpoints(gs *Geom) int {
	res := C.line_numpoints(gs._inner)
	return int(res)
}


// LinePointN wraps MEOS C function line_point_n.
func LinePointN(geom *Geom, n int) *Geom {
	res := C.line_point_n(geom._inner, C.int(n))
	return &Geom{_inner: res}
}


// TODO geo_wlof: unsupported param uint32_t *
// func GeoWlof(...) { /* not yet handled by codegen */ }


// GeoReverse wraps MEOS C function geo_reverse.
func GeoReverse(gs *Geom) *Geom {
	res := C.geo_reverse(gs._inner)
	return &Geom{_inner: res}
}


// GeoRound wraps MEOS C function geo_round.
func GeoRound(gs *Geom, maxdd int) *Geom {
	res := C.geo_round(gs._inner, C.int(maxdd))
	return &Geom{_inner: res}
}


// GeoSetSRID wraps MEOS C function geo_set_srid.
func GeoSetSRID(gs *Geom, srid int32) *Geom {
	res := C.geo_set_srid(gs._inner, C.int32_t(srid))
	return &Geom{_inner: res}
}


// GeoSRID wraps MEOS C function geo_srid.
func GeoSRID(gs *Geom) int32 {
	res := C.geo_srid(gs._inner)
	return int32(res)
}


// GeoTransform wraps MEOS C function geo_transform.
func GeoTransform(geom *Geom, srid_to int32) *Geom {
	res := C.geo_transform(geom._inner, C.int32_t(srid_to))
	return &Geom{_inner: res}
}


// GeoTransformPipeline wraps MEOS C function geo_transform_pipeline.
func GeoTransformPipeline(gs *Geom, pipeline string, srid_to int32, is_forward bool) *Geom {
	_c_pipeline := C.CString(pipeline)
	defer C.free(unsafe.Pointer(_c_pipeline))
	res := C.geo_transform_pipeline(gs._inner, _c_pipeline, C.int32_t(srid_to), C.bool(is_forward))
	return &Geom{_inner: res}
}


// GeoCollectGarray wraps MEOS C function geo_collect_garray.
func GeoCollectGarray(gsarr []*Geom) *Geom {
	_c_gsarr := make([]*C.GSERIALIZED, len(gsarr))
	for _i, _v := range gsarr { _c_gsarr[_i] = _v._inner }
	res := C.geo_collect_garray((**C.GSERIALIZED)(unsafe.Pointer(&_c_gsarr[0])), C.int(len(gsarr)))
	return &Geom{_inner: res}
}


// GeoMakelineGarray wraps MEOS C function geo_makeline_garray.
func GeoMakelineGarray(gsarr []*Geom) *Geom {
	_c_gsarr := make([]*C.GSERIALIZED, len(gsarr))
	for _i, _v := range gsarr { _c_gsarr[_i] = _v._inner }
	res := C.geo_makeline_garray((**C.GSERIALIZED)(unsafe.Pointer(&_c_gsarr[0])), C.int(len(gsarr)))
	return &Geom{_inner: res}
}


// GeoNumPoints wraps MEOS C function geo_num_points.
func GeoNumPoints(gs *Geom) int {
	res := C.geo_num_points(gs._inner)
	return int(res)
}


// GeoNumGeos wraps MEOS C function geo_num_geos.
func GeoNumGeos(gs *Geom) int {
	res := C.geo_num_geos(gs._inner)
	return int(res)
}


// GeoGeoN wraps MEOS C function geo_geo_n.
func GeoGeoN(geom *Geom, n int) *Geom {
	res := C.geo_geo_n(geom._inner, C.int(n))
	return &Geom{_inner: res}
}


// GeoPointarr wraps MEOS C function geo_pointarr.
func GeoPointarr(gs *Geom) []*Geom {
	var _out_count C.int
	res := C.geo_pointarr(gs._inner, &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.GSERIALIZED)(unsafe.Pointer(res)), _n)
	_out := make([]*Geom, _n)
	for _i, _e := range _slice {
		_out[_i] = &Geom{_inner: _e}
	}
	return _out
}


// GeoPoints wraps MEOS C function geo_points.
func GeoPoints(gs *Geom) *Geom {
	res := C.geo_points(gs._inner)
	return &Geom{_inner: res}
}


// GeomArrayUnion wraps MEOS C function geom_array_union.
func GeomArrayUnion(gsarr []*Geom) *Geom {
	_c_gsarr := make([]*C.GSERIALIZED, len(gsarr))
	for _i, _v := range gsarr { _c_gsarr[_i] = _v._inner }
	res := C.geom_array_union((**C.GSERIALIZED)(unsafe.Pointer(&_c_gsarr[0])), C.int(len(gsarr)))
	return &Geom{_inner: res}
}


// GeomBoundary wraps MEOS C function geom_boundary.
func GeomBoundary(gs *Geom) *Geom {
	res := C.geom_boundary(gs._inner)
	return &Geom{_inner: res}
}


// GeomBuffer wraps MEOS C function geom_buffer.
func GeomBuffer(gs *Geom, size float64, params string) *Geom {
	_c_params := C.CString(params)
	defer C.free(unsafe.Pointer(_c_params))
	res := C.geom_buffer(gs._inner, C.double(size), _c_params)
	return &Geom{_inner: res}
}


// GeomCentroid wraps MEOS C function geom_centroid.
func GeomCentroid(gs *Geom) *Geom {
	res := C.geom_centroid(gs._inner)
	return &Geom{_inner: res}
}


// GeomConvexHull wraps MEOS C function geom_convex_hull.
func GeomConvexHull(gs *Geom) *Geom {
	res := C.geom_convex_hull(gs._inner)
	return &Geom{_inner: res}
}


// GeomDifference2d wraps MEOS C function geom_difference2d.
func GeomDifference2d(gs1 *Geom, gs2 *Geom) *Geom {
	res := C.geom_difference2d(gs1._inner, gs2._inner)
	return &Geom{_inner: res}
}


// GeomIntersection2d wraps MEOS C function geom_intersection2d.
func GeomIntersection2d(gs1 *Geom, gs2 *Geom) *Geom {
	res := C.geom_intersection2d(gs1._inner, gs2._inner)
	return &Geom{_inner: res}
}


// GeomIntersection2dColl wraps MEOS C function geom_intersection2d_coll.
func GeomIntersection2dColl(gs1 *Geom, gs2 *Geom) *Geom {
	res := C.geom_intersection2d_coll(gs1._inner, gs2._inner)
	return &Geom{_inner: res}
}


// TODO geom_min_bounding_radius: unsupported param double *
// func GeomMinBoundingRadius(...) { /* not yet handled by codegen */ }


// GeomShortestline2d wraps MEOS C function geom_shortestline2d.
func GeomShortestline2d(gs1 *Geom, gs2 *Geom) *Geom {
	res := C.geom_shortestline2d(gs1._inner, gs2._inner)
	return &Geom{_inner: res}
}


// GeomShortestline3d wraps MEOS C function geom_shortestline3d.
func GeomShortestline3d(gs1 *Geom, gs2 *Geom) *Geom {
	res := C.geom_shortestline3d(gs1._inner, gs2._inner)
	return &Geom{_inner: res}
}


// GeomUnaryUnion wraps MEOS C function geom_unary_union.
func GeomUnaryUnion(gs *Geom, prec float64) *Geom {
	res := C.geom_unary_union(gs._inner, C.double(prec))
	return &Geom{_inner: res}
}


// LineInterpolatePoint wraps MEOS C function line_interpolate_point.
func LineInterpolatePoint(gs *Geom, distance_fraction float64, repeat bool) *Geom {
	res := C.line_interpolate_point(gs._inner, C.double(distance_fraction), C.bool(repeat))
	return &Geom{_inner: res}
}


// LineLocatePoint wraps MEOS C function line_locate_point.
func LineLocatePoint(gs1 *Geom, gs2 *Geom) float64 {
	res := C.line_locate_point(gs1._inner, gs2._inner)
	return float64(res)
}


// LineSubstring wraps MEOS C function line_substring.
func LineSubstring(gs *Geom, from float64, to float64) *Geom {
	res := C.line_substring(gs._inner, C.double(from), C.double(to))
	return &Geom{_inner: res}
}


// GeogDwithin wraps MEOS C function geog_dwithin.
func GeogDwithin(g1 *Geom, g2 *Geom, tolerance float64, use_spheroid bool) bool {
	res := C.geog_dwithin(g1._inner, g2._inner, C.double(tolerance), C.bool(use_spheroid))
	return bool(res)
}


// GeogIntersects wraps MEOS C function geog_intersects.
func GeogIntersects(gs1 *Geom, gs2 *Geom, use_spheroid bool) bool {
	res := C.geog_intersects(gs1._inner, gs2._inner, C.bool(use_spheroid))
	return bool(res)
}


// GeomContains wraps MEOS C function geom_contains.
func GeomContains(gs1 *Geom, gs2 *Geom) bool {
	res := C.geom_contains(gs1._inner, gs2._inner)
	return bool(res)
}


// GeomCovers wraps MEOS C function geom_covers.
func GeomCovers(gs1 *Geom, gs2 *Geom) bool {
	res := C.geom_covers(gs1._inner, gs2._inner)
	return bool(res)
}


// GeomDisjoint2d wraps MEOS C function geom_disjoint2d.
func GeomDisjoint2d(gs1 *Geom, gs2 *Geom) bool {
	res := C.geom_disjoint2d(gs1._inner, gs2._inner)
	return bool(res)
}


// GeomDwithin wraps MEOS C function geom_dwithin.
func GeomDwithin(gs1 *Geom, gs2 *Geom, tolerance float64) bool {
	res := C.geom_dwithin(gs1._inner, gs2._inner, C.double(tolerance))
	return bool(res)
}


// GeomDwithin2d wraps MEOS C function geom_dwithin2d.
func GeomDwithin2d(gs1 *Geom, gs2 *Geom, tolerance float64) bool {
	res := C.geom_dwithin2d(gs1._inner, gs2._inner, C.double(tolerance))
	return bool(res)
}


// GeomDwithin3d wraps MEOS C function geom_dwithin3d.
func GeomDwithin3d(gs1 *Geom, gs2 *Geom, tolerance float64) bool {
	res := C.geom_dwithin3d(gs1._inner, gs2._inner, C.double(tolerance))
	return bool(res)
}


// GeomIntersects wraps MEOS C function geom_intersects.
func GeomIntersects(gs1 *Geom, gs2 *Geom) bool {
	res := C.geom_intersects(gs1._inner, gs2._inner)
	return bool(res)
}


// GeomIntersects2d wraps MEOS C function geom_intersects2d.
func GeomIntersects2d(gs1 *Geom, gs2 *Geom) bool {
	res := C.geom_intersects2d(gs1._inner, gs2._inner)
	return bool(res)
}


// GeomIntersects3d wraps MEOS C function geom_intersects3d.
func GeomIntersects3d(gs1 *Geom, gs2 *Geom) bool {
	res := C.geom_intersects3d(gs1._inner, gs2._inner)
	return bool(res)
}


// GeomRelatePattern wraps MEOS C function geom_relate_pattern.
func GeomRelatePattern(gs1 *Geom, gs2 *Geom, patt string) bool {
	_c_patt := C.CString(patt)
	defer C.free(unsafe.Pointer(_c_patt))
	res := C.geom_relate_pattern(gs1._inner, gs2._inner, _c_patt)
	return bool(res)
}


// GeomTouches wraps MEOS C function geom_touches.
func GeomTouches(gs1 *Geom, gs2 *Geom) bool {
	res := C.geom_touches(gs1._inner, gs2._inner)
	return bool(res)
}


// GeoStboxes wraps MEOS C function geo_stboxes.
func GeoStboxes(gs *Geom) (*STBox, int) {
	var _out_count C.int
	res := C.geo_stboxes(gs._inner, &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// GeoSplitEachNStboxes wraps MEOS C function geo_split_each_n_stboxes.
func GeoSplitEachNStboxes(gs *Geom, elem_count int) (*STBox, int) {
	var _out_count C.int
	res := C.geo_split_each_n_stboxes(gs._inner, C.int(elem_count), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// GeoSplitNStboxes wraps MEOS C function geo_split_n_stboxes.
func GeoSplitNStboxes(gs *Geom, box_count int) (*STBox, int) {
	var _out_count C.int
	res := C.geo_split_n_stboxes(gs._inner, C.int(box_count), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// GeogDistance wraps MEOS C function geog_distance.
func GeogDistance(g1 *Geom, g2 *Geom) float64 {
	res := C.geog_distance(g1._inner, g2._inner)
	return float64(res)
}


// GeomDistance2d wraps MEOS C function geom_distance2d.
func GeomDistance2d(gs1 *Geom, gs2 *Geom) float64 {
	res := C.geom_distance2d(gs1._inner, gs2._inner)
	return float64(res)
}


// GeomDistance3d wraps MEOS C function geom_distance3d.
func GeomDistance3d(gs1 *Geom, gs2 *Geom) float64 {
	res := C.geom_distance3d(gs1._inner, gs2._inner)
	return float64(res)
}


// GeoEquals wraps MEOS C function geo_equals.
func GeoEquals(gs1 *Geom, gs2 *Geom) int {
	res := C.geo_equals(gs1._inner, gs2._inner)
	return int(res)
}


// GeoSame wraps MEOS C function geo_same.
func GeoSame(gs1 *Geom, gs2 *Geom) bool {
	res := C.geo_same(gs1._inner, gs2._inner)
	return bool(res)
}


// GeogsetIn wraps MEOS C function geogset_in.
func GeogsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.geogset_in(_c_str)
	return &Set{_inner: res}
}


// GeomsetIn wraps MEOS C function geomset_in.
func GeomsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.geomset_in(_c_str)
	return &Set{_inner: res}
}


// SpatialsetOut wraps MEOS C function spatialset_out.
func SpatialsetOut(s *Set, maxdd int) string {
	res := C.spatialset_out(s._inner, C.int(maxdd))
	return C.GoString(res)
}


// SpatialsetAsText wraps MEOS C function spatialset_as_text.
func SpatialsetAsText(set *Set, maxdd int) string {
	res := C.spatialset_as_text(set._inner, C.int(maxdd))
	return C.GoString(res)
}


// SpatialsetAsEWKT wraps MEOS C function spatialset_as_ewkt.
func SpatialsetAsEWKT(set *Set, maxdd int) string {
	res := C.spatialset_as_ewkt(set._inner, C.int(maxdd))
	return C.GoString(res)
}


// GeosetMake wraps MEOS C function geoset_make.
func GeosetMake(values []*Geom) *Set {
	_c_values := make([]*C.GSERIALIZED, len(values))
	for _i, _v := range values { _c_values[_i] = _v._inner }
	res := C.geoset_make((**C.GSERIALIZED)(unsafe.Pointer(&_c_values[0])), C.int(len(values)))
	return &Set{_inner: res}
}


// GeoToSet wraps MEOS C function geo_to_set.
func GeoToSet(gs *Geom) *Set {
	res := C.geo_to_set(gs._inner)
	return &Set{_inner: res}
}


// GeosetEndValue wraps MEOS C function geoset_end_value.
func GeosetEndValue(s *Set) *Geom {
	res := C.geoset_end_value(s._inner)
	return &Geom{_inner: res}
}


// GeosetStartValue wraps MEOS C function geoset_start_value.
func GeosetStartValue(s *Set) *Geom {
	res := C.geoset_start_value(s._inner)
	return &Geom{_inner: res}
}


// GeosetValueN wraps MEOS C function geoset_value_n.
func GeosetValueN(s *Set, n int) (bool, *Geom) {
	var _out_result *C.GSERIALIZED
	res := C.geoset_value_n(s._inner, C.int(n), &_out_result)
	return bool(res), &Geom{_inner: _out_result}
}


// GeosetValues wraps MEOS C function geoset_values.
func GeosetValues(s *Set) []*Geom {
	var _out_count C.int
	res := C.geoset_values(s._inner, &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.GSERIALIZED)(unsafe.Pointer(res)), _n)
	_out := make([]*Geom, _n)
	for _i, _e := range _slice {
		_out[_i] = &Geom{_inner: _e}
	}
	return _out
}


// ContainedGeoSet wraps MEOS C function contained_geo_set.
func ContainedGeoSet(gs *Geom, s *Set) bool {
	res := C.contained_geo_set(gs._inner, s._inner)
	return bool(res)
}


// ContainsSetGeo wraps MEOS C function contains_set_geo.
func ContainsSetGeo(s *Set, gs *Geom) bool {
	res := C.contains_set_geo(s._inner, gs._inner)
	return bool(res)
}


// GeoUnionTransfn wraps MEOS C function geo_union_transfn.
func GeoUnionTransfn(state *Set, gs *Geom) *Set {
	res := C.geo_union_transfn(state._inner, gs._inner)
	return &Set{_inner: res}
}


// IntersectionGeoSet wraps MEOS C function intersection_geo_set.
func IntersectionGeoSet(gs *Geom, s *Set) *Set {
	res := C.intersection_geo_set(gs._inner, s._inner)
	return &Set{_inner: res}
}


// IntersectionSetGeo wraps MEOS C function intersection_set_geo.
func IntersectionSetGeo(s *Set, gs *Geom) *Set {
	res := C.intersection_set_geo(s._inner, gs._inner)
	return &Set{_inner: res}
}


// MinusGeoSet wraps MEOS C function minus_geo_set.
func MinusGeoSet(gs *Geom, s *Set) *Set {
	res := C.minus_geo_set(gs._inner, s._inner)
	return &Set{_inner: res}
}


// MinusSetGeo wraps MEOS C function minus_set_geo.
func MinusSetGeo(s *Set, gs *Geom) *Set {
	res := C.minus_set_geo(s._inner, gs._inner)
	return &Set{_inner: res}
}


// UnionGeoSet wraps MEOS C function union_geo_set.
func UnionGeoSet(gs *Geom, s *Set) *Set {
	res := C.union_geo_set(gs._inner, s._inner)
	return &Set{_inner: res}
}


// UnionSetGeo wraps MEOS C function union_set_geo.
func UnionSetGeo(s *Set, gs *Geom) *Set {
	res := C.union_set_geo(s._inner, gs._inner)
	return &Set{_inner: res}
}


// SpatialsetSetSRID wraps MEOS C function spatialset_set_srid.
func SpatialsetSetSRID(s *Set, srid int32) *Set {
	res := C.spatialset_set_srid(s._inner, C.int32_t(srid))
	return &Set{_inner: res}
}


// SpatialsetSRID wraps MEOS C function spatialset_srid.
func SpatialsetSRID(s *Set) int32 {
	res := C.spatialset_srid(s._inner)
	return int32(res)
}


// SpatialsetTransform wraps MEOS C function spatialset_transform.
func SpatialsetTransform(s *Set, srid int32) *Set {
	res := C.spatialset_transform(s._inner, C.int32_t(srid))
	return &Set{_inner: res}
}


// SpatialsetTransformPipeline wraps MEOS C function spatialset_transform_pipeline.
func SpatialsetTransformPipeline(s *Set, pipelinestr string, srid int32, is_forward bool) *Set {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	res := C.spatialset_transform_pipeline(s._inner, _c_pipelinestr, C.int32_t(srid), C.bool(is_forward))
	return &Set{_inner: res}
}


// STBOXAsHexwkb wraps MEOS C function stbox_as_hexwkb.
func STBOXAsHexwkb(box *STBox, variant uint8) (string, uint) {
	var _out_size C.size_t
	res := C.stbox_as_hexwkb(box._inner, C.uint8_t(variant), &_out_size)
	return C.GoString(res), uint(_out_size)
}


// STBOXAsWKB wraps MEOS C function stbox_as_wkb.
func STBOXAsWKB(box *STBox, variant uint8) []uint8 {
	var _out_size_out C.size_t
	res := C.stbox_as_wkb(box._inner, C.uint8_t(variant), &_out_size_out)
	_n := int(_out_size_out)
	_slice := unsafe.Slice((*C.uint8_t)(unsafe.Pointer(res)), _n)
	_out := make([]uint8, _n)
	for _i, _e := range _slice {
		_out[_i] = uint8(_e)
	}
	return _out
}


// STBOXFromHexwkb wraps MEOS C function stbox_from_hexwkb.
func STBOXFromHexwkb(hexwkb string) *STBox {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	res := C.stbox_from_hexwkb(_c_hexwkb)
	return &STBox{_inner: res}
}


// STBOXFromWKB wraps MEOS C function stbox_from_wkb.
func STBOXFromWKB(wkb []byte) *STBox {
	res := C.stbox_from_wkb((*C.uint8_t)(unsafe.Pointer(&wkb[0])), C.size_t(len(wkb)))
	return &STBox{_inner: res}
}


// STBOXIn wraps MEOS C function stbox_in.
func STBOXIn(str string) *STBox {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.stbox_in(_c_str)
	return &STBox{_inner: res}
}


// STBOXOut wraps MEOS C function stbox_out.
func STBOXOut(box *STBox, maxdd int) string {
	res := C.stbox_out(box._inner, C.int(maxdd))
	return C.GoString(res)
}


// GeoTimestamptzToSTBOX wraps MEOS C function geo_timestamptz_to_stbox.
func GeoTimestamptzToSTBOX(gs *Geom, t int64) *STBox {
	res := C.geo_timestamptz_to_stbox(gs._inner, C.TimestampTz(t))
	return &STBox{_inner: res}
}


// GeoTstzspanToSTBOX wraps MEOS C function geo_tstzspan_to_stbox.
func GeoTstzspanToSTBOX(gs *Geom, s *Span) *STBox {
	res := C.geo_tstzspan_to_stbox(gs._inner, s._inner)
	return &STBox{_inner: res}
}


// STBOXCopy wraps MEOS C function stbox_copy.
func STBOXCopy(box *STBox) *STBox {
	res := C.stbox_copy(box._inner)
	return &STBox{_inner: res}
}


// STBOXMake wraps MEOS C function stbox_make.
func STBOXMake(hasx bool, hasz bool, geodetic bool, srid int32, xmin float64, xmax float64, ymin float64, ymax float64, zmin float64, zmax float64, s *Span) *STBox {
	res := C.stbox_make(C.bool(hasx), C.bool(hasz), C.bool(geodetic), C.int32_t(srid), C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.double(zmin), C.double(zmax), s._inner)
	return &STBox{_inner: res}
}


// GeoToSTBOX wraps MEOS C function geo_to_stbox.
func GeoToSTBOX(gs *Geom) *STBox {
	res := C.geo_to_stbox(gs._inner)
	return &STBox{_inner: res}
}


// SpatialsetToSTBOX wraps MEOS C function spatialset_to_stbox.
func SpatialsetToSTBOX(s *Set) *STBox {
	res := C.spatialset_to_stbox(s._inner)
	return &STBox{_inner: res}
}


// STBOXToBox3d wraps MEOS C function stbox_to_box3d.
func STBOXToBox3d(box *STBox) *Box3D {
	res := C.stbox_to_box3d(box._inner)
	return &Box3D{_inner: res}
}


// STBOXToGbox wraps MEOS C function stbox_to_gbox.
func STBOXToGbox(box *STBox) *GBox {
	res := C.stbox_to_gbox(box._inner)
	return &GBox{_inner: res}
}


// STBOXToGeo wraps MEOS C function stbox_to_geo.
func STBOXToGeo(box *STBox) *Geom {
	res := C.stbox_to_geo(box._inner)
	return &Geom{_inner: res}
}


// STBOXToTstzspan wraps MEOS C function stbox_to_tstzspan.
func STBOXToTstzspan(box *STBox) *Span {
	res := C.stbox_to_tstzspan(box._inner)
	return &Span{_inner: res}
}


// TimestamptzToSTBOX wraps MEOS C function timestamptz_to_stbox.
func TimestamptzToSTBOX(t int64) *STBox {
	res := C.timestamptz_to_stbox(C.TimestampTz(t))
	return &STBox{_inner: res}
}


// TstzsetToSTBOX wraps MEOS C function tstzset_to_stbox.
func TstzsetToSTBOX(s *Set) *STBox {
	res := C.tstzset_to_stbox(s._inner)
	return &STBox{_inner: res}
}


// TstzspanToSTBOX wraps MEOS C function tstzspan_to_stbox.
func TstzspanToSTBOX(s *Span) *STBox {
	res := C.tstzspan_to_stbox(s._inner)
	return &STBox{_inner: res}
}


// TstzspansetToSTBOX wraps MEOS C function tstzspanset_to_stbox.
func TstzspansetToSTBOX(ss *SpanSet) *STBox {
	res := C.tstzspanset_to_stbox(ss._inner)
	return &STBox{_inner: res}
}


// STBOXArea wraps MEOS C function stbox_area.
func STBOXArea(box *STBox, spheroid bool) float64 {
	res := C.stbox_area(box._inner, C.bool(spheroid))
	return float64(res)
}


// STBOXHash wraps MEOS C function stbox_hash.
func STBOXHash(box *STBox) int {
	res := C.stbox_hash(box._inner)
	return int(res)
}


// STBOXHashExtended wraps MEOS C function stbox_hash_extended.
func STBOXHashExtended(box *STBox, seed int) int {
	res := C.stbox_hash_extended(box._inner, C.int(seed))
	return int(res)
}


// STBOXHast wraps MEOS C function stbox_hast.
func STBOXHast(box *STBox) bool {
	res := C.stbox_hast(box._inner)
	return bool(res)
}


// STBOXHasx wraps MEOS C function stbox_hasx.
func STBOXHasx(box *STBox) bool {
	res := C.stbox_hasx(box._inner)
	return bool(res)
}


// STBOXHasz wraps MEOS C function stbox_hasz.
func STBOXHasz(box *STBox) bool {
	res := C.stbox_hasz(box._inner)
	return bool(res)
}


// STBOXIsgeodetic wraps MEOS C function stbox_isgeodetic.
func STBOXIsgeodetic(box *STBox) bool {
	res := C.stbox_isgeodetic(box._inner)
	return bool(res)
}


// STBOXPerimeter wraps MEOS C function stbox_perimeter.
func STBOXPerimeter(box *STBox, spheroid bool) float64 {
	res := C.stbox_perimeter(box._inner, C.bool(spheroid))
	return float64(res)
}


// STBOXTmax wraps MEOS C function stbox_tmax.
func STBOXTmax(box *STBox) (bool, int64) {
	var _out_result C.TimestampTz
	res := C.stbox_tmax(box._inner, &_out_result)
	return bool(res), int64(_out_result)
}


// STBOXTmaxInc wraps MEOS C function stbox_tmax_inc.
func STBOXTmaxInc(box *STBox) (bool, bool) {
	var _out_result C.bool
	res := C.stbox_tmax_inc(box._inner, &_out_result)
	return bool(res), bool(_out_result)
}


// STBOXTmin wraps MEOS C function stbox_tmin.
func STBOXTmin(box *STBox) (bool, int64) {
	var _out_result C.TimestampTz
	res := C.stbox_tmin(box._inner, &_out_result)
	return bool(res), int64(_out_result)
}


// STBOXTminInc wraps MEOS C function stbox_tmin_inc.
func STBOXTminInc(box *STBox) (bool, bool) {
	var _out_result C.bool
	res := C.stbox_tmin_inc(box._inner, &_out_result)
	return bool(res), bool(_out_result)
}


// STBOXVolume wraps MEOS C function stbox_volume.
func STBOXVolume(box *STBox) float64 {
	res := C.stbox_volume(box._inner)
	return float64(res)
}


// STBOXXmax wraps MEOS C function stbox_xmax.
func STBOXXmax(box *STBox) (bool, float64) {
	var _out_result C.double
	res := C.stbox_xmax(box._inner, &_out_result)
	return bool(res), float64(_out_result)
}


// STBOXXmin wraps MEOS C function stbox_xmin.
func STBOXXmin(box *STBox) (bool, float64) {
	var _out_result C.double
	res := C.stbox_xmin(box._inner, &_out_result)
	return bool(res), float64(_out_result)
}


// STBOXYmax wraps MEOS C function stbox_ymax.
func STBOXYmax(box *STBox) (bool, float64) {
	var _out_result C.double
	res := C.stbox_ymax(box._inner, &_out_result)
	return bool(res), float64(_out_result)
}


// STBOXYmin wraps MEOS C function stbox_ymin.
func STBOXYmin(box *STBox) (bool, float64) {
	var _out_result C.double
	res := C.stbox_ymin(box._inner, &_out_result)
	return bool(res), float64(_out_result)
}


// STBOXZmax wraps MEOS C function stbox_zmax.
func STBOXZmax(box *STBox) (bool, float64) {
	var _out_result C.double
	res := C.stbox_zmax(box._inner, &_out_result)
	return bool(res), float64(_out_result)
}


// STBOXZmin wraps MEOS C function stbox_zmin.
func STBOXZmin(box *STBox) (bool, float64) {
	var _out_result C.double
	res := C.stbox_zmin(box._inner, &_out_result)
	return bool(res), float64(_out_result)
}


// STBOXExpandSpace wraps MEOS C function stbox_expand_space.
func STBOXExpandSpace(box *STBox, d float64) *STBox {
	res := C.stbox_expand_space(box._inner, C.double(d))
	return &STBox{_inner: res}
}


// STBOXExpandTime wraps MEOS C function stbox_expand_time.
func STBOXExpandTime(box *STBox, interv timeutil.Timedelta) *STBox {
	res := C.stbox_expand_time(box._inner, interv.Inner())
	return &STBox{_inner: res}
}


// STBOXGetSpace wraps MEOS C function stbox_get_space.
func STBOXGetSpace(box *STBox) *STBox {
	res := C.stbox_get_space(box._inner)
	return &STBox{_inner: res}
}


// STBOXQuadSplit wraps MEOS C function stbox_quad_split.
func STBOXQuadSplit(box *STBox) (*STBox, int) {
	var _out_count C.int
	res := C.stbox_quad_split(box._inner, &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// STBOXRound wraps MEOS C function stbox_round.
func STBOXRound(box *STBox, maxdd int) *STBox {
	res := C.stbox_round(box._inner, C.int(maxdd))
	return &STBox{_inner: res}
}


// STBOXShiftScaleTime wraps MEOS C function stbox_shift_scale_time.
func STBOXShiftScaleTime(box *STBox, shift timeutil.Timedelta, duration timeutil.Timedelta) *STBox {
	res := C.stbox_shift_scale_time(box._inner, shift.Inner(), duration.Inner())
	return &STBox{_inner: res}
}


// StboxarrRound wraps MEOS C function stboxarr_round.
func StboxarrRound(boxarr *STBox, count int, maxdd int) *STBox {
	res := C.stboxarr_round(boxarr._inner, C.int(count), C.int(maxdd))
	return &STBox{_inner: res}
}


// STBOXSetSRID wraps MEOS C function stbox_set_srid.
func STBOXSetSRID(box *STBox, srid int32) *STBox {
	res := C.stbox_set_srid(box._inner, C.int32_t(srid))
	return &STBox{_inner: res}
}


// STBOXSRID wraps MEOS C function stbox_srid.
func STBOXSRID(box *STBox) int32 {
	res := C.stbox_srid(box._inner)
	return int32(res)
}


// STBOXTransform wraps MEOS C function stbox_transform.
func STBOXTransform(box *STBox, srid int32) *STBox {
	res := C.stbox_transform(box._inner, C.int32_t(srid))
	return &STBox{_inner: res}
}


// STBOXTransformPipeline wraps MEOS C function stbox_transform_pipeline.
func STBOXTransformPipeline(box *STBox, pipelinestr string, srid int32, is_forward bool) *STBox {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	res := C.stbox_transform_pipeline(box._inner, _c_pipelinestr, C.int32_t(srid), C.bool(is_forward))
	return &STBox{_inner: res}
}


// AdjacentSTBOXSTBOX wraps MEOS C function adjacent_stbox_stbox.
func AdjacentSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.adjacent_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// ContainedSTBOXSTBOX wraps MEOS C function contained_stbox_stbox.
func ContainedSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.contained_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// ContainsSTBOXSTBOX wraps MEOS C function contains_stbox_stbox.
func ContainsSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.contains_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// OverlapsSTBOXSTBOX wraps MEOS C function overlaps_stbox_stbox.
func OverlapsSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.overlaps_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// SameSTBOXSTBOX wraps MEOS C function same_stbox_stbox.
func SameSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.same_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// AboveSTBOXSTBOX wraps MEOS C function above_stbox_stbox.
func AboveSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.above_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// AfterSTBOXSTBOX wraps MEOS C function after_stbox_stbox.
func AfterSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.after_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// BackSTBOXSTBOX wraps MEOS C function back_stbox_stbox.
func BackSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.back_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// BeforeSTBOXSTBOX wraps MEOS C function before_stbox_stbox.
func BeforeSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.before_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// BelowSTBOXSTBOX wraps MEOS C function below_stbox_stbox.
func BelowSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.below_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// FrontSTBOXSTBOX wraps MEOS C function front_stbox_stbox.
func FrontSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.front_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// LeftSTBOXSTBOX wraps MEOS C function left_stbox_stbox.
func LeftSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.left_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// OveraboveSTBOXSTBOX wraps MEOS C function overabove_stbox_stbox.
func OveraboveSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.overabove_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// OverafterSTBOXSTBOX wraps MEOS C function overafter_stbox_stbox.
func OverafterSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.overafter_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// OverbackSTBOXSTBOX wraps MEOS C function overback_stbox_stbox.
func OverbackSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.overback_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// OverbeforeSTBOXSTBOX wraps MEOS C function overbefore_stbox_stbox.
func OverbeforeSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.overbefore_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// OverbelowSTBOXSTBOX wraps MEOS C function overbelow_stbox_stbox.
func OverbelowSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.overbelow_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// OverfrontSTBOXSTBOX wraps MEOS C function overfront_stbox_stbox.
func OverfrontSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.overfront_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// OverleftSTBOXSTBOX wraps MEOS C function overleft_stbox_stbox.
func OverleftSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.overleft_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// OverrightSTBOXSTBOX wraps MEOS C function overright_stbox_stbox.
func OverrightSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.overright_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// RightSTBOXSTBOX wraps MEOS C function right_stbox_stbox.
func RightSTBOXSTBOX(box1 *STBox, box2 *STBox) bool {
	res := C.right_stbox_stbox(box1._inner, box2._inner)
	return bool(res)
}


// UnionSTBOXSTBOX wraps MEOS C function union_stbox_stbox.
func UnionSTBOXSTBOX(box1 *STBox, box2 *STBox, strict bool) *STBox {
	res := C.union_stbox_stbox(box1._inner, box2._inner, C.bool(strict))
	return &STBox{_inner: res}
}


// IntersectionSTBOXSTBOX wraps MEOS C function intersection_stbox_stbox.
func IntersectionSTBOXSTBOX(box1 *STBox, box2 *STBox) *STBox {
	res := C.intersection_stbox_stbox(box1._inner, box2._inner)
	return &STBox{_inner: res}
}


// STBOXCmp wraps MEOS C function stbox_cmp.
func STBOXCmp(box1 *STBox, box2 *STBox) int {
	res := C.stbox_cmp(box1._inner, box2._inner)
	return int(res)
}


// STBOXEq wraps MEOS C function stbox_eq.
func STBOXEq(box1 *STBox, box2 *STBox) bool {
	res := C.stbox_eq(box1._inner, box2._inner)
	return bool(res)
}


// STBOXGe wraps MEOS C function stbox_ge.
func STBOXGe(box1 *STBox, box2 *STBox) bool {
	res := C.stbox_ge(box1._inner, box2._inner)
	return bool(res)
}


// STBOXGt wraps MEOS C function stbox_gt.
func STBOXGt(box1 *STBox, box2 *STBox) bool {
	res := C.stbox_gt(box1._inner, box2._inner)
	return bool(res)
}


// STBOXLe wraps MEOS C function stbox_le.
func STBOXLe(box1 *STBox, box2 *STBox) bool {
	res := C.stbox_le(box1._inner, box2._inner)
	return bool(res)
}


// STBOXLt wraps MEOS C function stbox_lt.
func STBOXLt(box1 *STBox, box2 *STBox) bool {
	res := C.stbox_lt(box1._inner, box2._inner)
	return bool(res)
}


// STBOXNe wraps MEOS C function stbox_ne.
func STBOXNe(box1 *STBox, box2 *STBox) bool {
	res := C.stbox_ne(box1._inner, box2._inner)
	return bool(res)
}


// TspatialOut wraps MEOS C function tspatial_out.
func TspatialOut(temp Temporal, maxdd int) string {
	res := C.tspatial_out(temp.Inner(), C.int(maxdd))
	return C.GoString(res)
}


// TgeogpointFromMFJSON wraps MEOS C function tgeogpoint_from_mfjson.
func TgeogpointFromMFJSON(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeogpoint_from_mfjson(_c_str)
	return CreateTemporal(res)
}


// TgeogpointIn wraps MEOS C function tgeogpoint_in.
func TgeogpointIn(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeogpoint_in(_c_str)
	return CreateTemporal(res)
}


// TgeographyFromMFJSON wraps MEOS C function tgeography_from_mfjson.
func TgeographyFromMFJSON(mfjson string) Temporal {
	_c_mfjson := C.CString(mfjson)
	defer C.free(unsafe.Pointer(_c_mfjson))
	res := C.tgeography_from_mfjson(_c_mfjson)
	return CreateTemporal(res)
}


// TgeographyIn wraps MEOS C function tgeography_in.
func TgeographyIn(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeography_in(_c_str)
	return CreateTemporal(res)
}


// TgeometryFromMFJSON wraps MEOS C function tgeometry_from_mfjson.
func TgeometryFromMFJSON(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeometry_from_mfjson(_c_str)
	return CreateTemporal(res)
}


// TgeometryIn wraps MEOS C function tgeometry_in.
func TgeometryIn(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeometry_in(_c_str)
	return CreateTemporal(res)
}


// TgeompointFromMFJSON wraps MEOS C function tgeompoint_from_mfjson.
func TgeompointFromMFJSON(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeompoint_from_mfjson(_c_str)
	return CreateTemporal(res)
}


// TgeompointIn wraps MEOS C function tgeompoint_in.
func TgeompointIn(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeompoint_in(_c_str)
	return CreateTemporal(res)
}


// TspatialAsEWKT wraps MEOS C function tspatial_as_ewkt.
func TspatialAsEWKT(temp Temporal, maxdd int) string {
	res := C.tspatial_as_ewkt(temp.Inner(), C.int(maxdd))
	return C.GoString(res)
}


// TspatialAsText wraps MEOS C function tspatial_as_text.
func TspatialAsText(temp Temporal, maxdd int) string {
	res := C.tspatial_as_text(temp.Inner(), C.int(maxdd))
	return C.GoString(res)
}


// TgeoFromBaseTemp wraps MEOS C function tgeo_from_base_temp.
func TgeoFromBaseTemp(gs *Geom, temp Temporal) Temporal {
	res := C.tgeo_from_base_temp(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TgeoinstMake wraps MEOS C function tgeoinst_make.
func TgeoinstMake(gs *Geom, t int64) TInstant {
	res := C.tgeoinst_make(gs._inner, C.TimestampTz(t))
	return TInstant{_inner: res}
}


// TgeoseqFromBaseTstzset wraps MEOS C function tgeoseq_from_base_tstzset.
func TgeoseqFromBaseTstzset(gs *Geom, s *Set) TSequence {
	res := C.tgeoseq_from_base_tstzset(gs._inner, s._inner)
	return TSequence{_inner: res}
}


// TgeoseqFromBaseTstzspan wraps MEOS C function tgeoseq_from_base_tstzspan.
func TgeoseqFromBaseTstzspan(gs *Geom, s *Span, interp Interpolation) TSequence {
	res := C.tgeoseq_from_base_tstzspan(gs._inner, s._inner, C.interpType(interp))
	return TSequence{_inner: res}
}


// TgeoseqsetFromBaseTstzspanset wraps MEOS C function tgeoseqset_from_base_tstzspanset.
func TgeoseqsetFromBaseTstzspanset(gs *Geom, ss *SpanSet, interp Interpolation) TSequenceSet {
	res := C.tgeoseqset_from_base_tstzspanset(gs._inner, ss._inner, C.interpType(interp))
	return TSequenceSet{_inner: res}
}


// TpointFromBaseTemp wraps MEOS C function tpoint_from_base_temp.
func TpointFromBaseTemp(gs *Geom, temp Temporal) Temporal {
	res := C.tpoint_from_base_temp(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TpointinstMake wraps MEOS C function tpointinst_make.
func TpointinstMake(gs *Geom, t int64) TInstant {
	res := C.tpointinst_make(gs._inner, C.TimestampTz(t))
	return TInstant{_inner: res}
}


// TpointseqFromBaseTstzset wraps MEOS C function tpointseq_from_base_tstzset.
func TpointseqFromBaseTstzset(gs *Geom, s *Set) TSequence {
	res := C.tpointseq_from_base_tstzset(gs._inner, s._inner)
	return TSequence{_inner: res}
}


// TpointseqFromBaseTstzspan wraps MEOS C function tpointseq_from_base_tstzspan.
func TpointseqFromBaseTstzspan(gs *Geom, s *Span, interp Interpolation) TSequence {
	res := C.tpointseq_from_base_tstzspan(gs._inner, s._inner, C.interpType(interp))
	return TSequence{_inner: res}
}


// TODO tpointseq_make_coords: unsupported param const double *
// func TpointseqMakeCoords(...) { /* not yet handled by codegen */ }


// TpointseqsetFromBaseTstzspanset wraps MEOS C function tpointseqset_from_base_tstzspanset.
func TpointseqsetFromBaseTstzspanset(gs *Geom, ss *SpanSet, interp Interpolation) TSequenceSet {
	res := C.tpointseqset_from_base_tstzspanset(gs._inner, ss._inner, C.interpType(interp))
	return TSequenceSet{_inner: res}
}


// Box3dToSTBOX wraps MEOS C function box3d_to_stbox.
func Box3dToSTBOX(box *Box3D) *STBox {
	res := C.box3d_to_stbox(box._inner)
	return &STBox{_inner: res}
}


// GboxToSTBOX wraps MEOS C function gbox_to_stbox.
func GboxToSTBOX(box *GBox) *STBox {
	res := C.gbox_to_stbox(box._inner)
	return &STBox{_inner: res}
}


// GeomeasToTpoint wraps MEOS C function geomeas_to_tpoint.
func GeomeasToTpoint(gs *Geom) Temporal {
	res := C.geomeas_to_tpoint(gs._inner)
	return CreateTemporal(res)
}


// TgeogpointToTgeography wraps MEOS C function tgeogpoint_to_tgeography.
func TgeogpointToTgeography(temp Temporal) Temporal {
	res := C.tgeogpoint_to_tgeography(temp.Inner())
	return CreateTemporal(res)
}


// TgeographyToTgeogpoint wraps MEOS C function tgeography_to_tgeogpoint.
func TgeographyToTgeogpoint(temp Temporal) Temporal {
	res := C.tgeography_to_tgeogpoint(temp.Inner())
	return CreateTemporal(res)
}


// TgeographyToTgeometry wraps MEOS C function tgeography_to_tgeometry.
func TgeographyToTgeometry(temp Temporal) Temporal {
	res := C.tgeography_to_tgeometry(temp.Inner())
	return CreateTemporal(res)
}


// TgeometryToTgeography wraps MEOS C function tgeometry_to_tgeography.
func TgeometryToTgeography(temp Temporal) Temporal {
	res := C.tgeometry_to_tgeography(temp.Inner())
	return CreateTemporal(res)
}


// TgeometryToTgeompoint wraps MEOS C function tgeometry_to_tgeompoint.
func TgeometryToTgeompoint(temp Temporal) Temporal {
	res := C.tgeometry_to_tgeompoint(temp.Inner())
	return CreateTemporal(res)
}


// TgeompointToTgeometry wraps MEOS C function tgeompoint_to_tgeometry.
func TgeompointToTgeometry(temp Temporal) Temporal {
	res := C.tgeompoint_to_tgeometry(temp.Inner())
	return CreateTemporal(res)
}


// TODO tpoint_as_mvtgeom: unsupported return type MvtGeom
// func TpointAsMvtgeom(...) { /* not yet handled by codegen */ }


// TpointTfloatToGeomeas wraps MEOS C function tpoint_tfloat_to_geomeas.
func TpointTfloatToGeomeas(tpoint Temporal, measure Temporal, segmentize bool) (bool, *Geom) {
	var _out_result *C.GSERIALIZED
	res := C.tpoint_tfloat_to_geomeas(tpoint.Inner(), measure.Inner(), C.bool(segmentize), &_out_result)
	return bool(res), &Geom{_inner: _out_result}
}


// TspatialToSTBOX wraps MEOS C function tspatial_to_stbox.
func TspatialToSTBOX(temp Temporal) *STBox {
	res := C.tspatial_to_stbox(temp.Inner())
	return &STBox{_inner: res}
}


// BearingPointPoint wraps MEOS C function bearing_point_point.
func BearingPointPoint(gs1 *Geom, gs2 *Geom) (bool, float64) {
	var _out_result C.double
	res := C.bearing_point_point(gs1._inner, gs2._inner, &_out_result)
	return bool(res), float64(_out_result)
}


// BearingTpointPoint wraps MEOS C function bearing_tpoint_point.
func BearingTpointPoint(temp Temporal, gs *Geom, invert bool) Temporal {
	res := C.bearing_tpoint_point(temp.Inner(), gs._inner, C.bool(invert))
	return CreateTemporal(res)
}


// BearingTpointTpoint wraps MEOS C function bearing_tpoint_tpoint.
func BearingTpointTpoint(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.bearing_tpoint_tpoint(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TgeoCentroid wraps MEOS C function tgeo_centroid.
func TgeoCentroid(temp Temporal) Temporal {
	res := C.tgeo_centroid(temp.Inner())
	return CreateTemporal(res)
}


// TgeoConvexHull wraps MEOS C function tgeo_convex_hull.
func TgeoConvexHull(temp Temporal) *Geom {
	res := C.tgeo_convex_hull(temp.Inner())
	return &Geom{_inner: res}
}


// TgeoEndValue wraps MEOS C function tgeo_end_value.
func TgeoEndValue(temp Temporal) *Geom {
	res := C.tgeo_end_value(temp.Inner())
	return &Geom{_inner: res}
}


// TgeoStartValue wraps MEOS C function tgeo_start_value.
func TgeoStartValue(temp Temporal) *Geom {
	res := C.tgeo_start_value(temp.Inner())
	return &Geom{_inner: res}
}


// TgeoTraversedArea wraps MEOS C function tgeo_traversed_area.
func TgeoTraversedArea(temp Temporal, unary_union bool) *Geom {
	res := C.tgeo_traversed_area(temp.Inner(), C.bool(unary_union))
	return &Geom{_inner: res}
}


// TgeoValueAtTimestamptz wraps MEOS C function tgeo_value_at_timestamptz.
func TgeoValueAtTimestamptz(temp Temporal, t int64, strict bool) (bool, *Geom) {
	var _out_result *C.GSERIALIZED
	res := C.tgeo_value_at_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(strict), &_out_result)
	return bool(res), &Geom{_inner: _out_result}
}


// TgeoValueN wraps MEOS C function tgeo_value_n.
func TgeoValueN(temp Temporal, n int) (bool, *Geom) {
	var _out_result *C.GSERIALIZED
	res := C.tgeo_value_n(temp.Inner(), C.int(n), &_out_result)
	return bool(res), &Geom{_inner: _out_result}
}


// TgeoValues wraps MEOS C function tgeo_values.
func TgeoValues(temp Temporal) []*Geom {
	var _out_count C.int
	res := C.tgeo_values(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.GSERIALIZED)(unsafe.Pointer(res)), _n)
	_out := make([]*Geom, _n)
	for _i, _e := range _slice {
		_out[_i] = &Geom{_inner: _e}
	}
	return _out
}


// TpointAngularDifference wraps MEOS C function tpoint_angular_difference.
func TpointAngularDifference(temp Temporal) Temporal {
	res := C.tpoint_angular_difference(temp.Inner())
	return CreateTemporal(res)
}


// TpointAzimuth wraps MEOS C function tpoint_azimuth.
func TpointAzimuth(temp Temporal) Temporal {
	res := C.tpoint_azimuth(temp.Inner())
	return CreateTemporal(res)
}


// TpointCumulativeLength wraps MEOS C function tpoint_cumulative_length.
func TpointCumulativeLength(temp Temporal) Temporal {
	res := C.tpoint_cumulative_length(temp.Inner())
	return CreateTemporal(res)
}


// TpointDirection wraps MEOS C function tpoint_direction.
func TpointDirection(temp Temporal) (bool, float64) {
	var _out_result C.double
	res := C.tpoint_direction(temp.Inner(), &_out_result)
	return bool(res), float64(_out_result)
}


// TpointGetX wraps MEOS C function tpoint_get_x.
func TpointGetX(temp Temporal) Temporal {
	res := C.tpoint_get_x(temp.Inner())
	return CreateTemporal(res)
}


// TpointGetY wraps MEOS C function tpoint_get_y.
func TpointGetY(temp Temporal) Temporal {
	res := C.tpoint_get_y(temp.Inner())
	return CreateTemporal(res)
}


// TpointGetZ wraps MEOS C function tpoint_get_z.
func TpointGetZ(temp Temporal) Temporal {
	res := C.tpoint_get_z(temp.Inner())
	return CreateTemporal(res)
}


// TpointIsSimple wraps MEOS C function tpoint_is_simple.
func TpointIsSimple(temp Temporal) bool {
	res := C.tpoint_is_simple(temp.Inner())
	return bool(res)
}


// TpointLength wraps MEOS C function tpoint_length.
func TpointLength(temp Temporal) float64 {
	res := C.tpoint_length(temp.Inner())
	return float64(res)
}


// TpointSpeed wraps MEOS C function tpoint_speed.
func TpointSpeed(temp Temporal) Temporal {
	res := C.tpoint_speed(temp.Inner())
	return CreateTemporal(res)
}


// TpointTrajectory wraps MEOS C function tpoint_trajectory.
func TpointTrajectory(temp Temporal, unary_union bool) *Geom {
	res := C.tpoint_trajectory(temp.Inner(), C.bool(unary_union))
	return &Geom{_inner: res}
}


// TpointTwcentroid wraps MEOS C function tpoint_twcentroid.
func TpointTwcentroid(temp Temporal) *Geom {
	res := C.tpoint_twcentroid(temp.Inner())
	return &Geom{_inner: res}
}


// TgeoAffine wraps MEOS C function tgeo_affine.
func TgeoAffine(temp Temporal, a *AFFINE) Temporal {
	res := C.tgeo_affine(temp.Inner(), a._inner)
	return CreateTemporal(res)
}


// TgeoScale wraps MEOS C function tgeo_scale.
func TgeoScale(temp Temporal, scale *Geom, sorigin *Geom) Temporal {
	res := C.tgeo_scale(temp.Inner(), scale._inner, sorigin._inner)
	return CreateTemporal(res)
}


// TpointMakeSimple wraps MEOS C function tpoint_make_simple.
func TpointMakeSimple(temp Temporal) []Temporal {
	var _out_count C.int
	res := C.tpoint_make_simple(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.Temporal)(unsafe.Pointer(res)), _n)
	_out := make([]Temporal, _n)
	for _i, _e := range _slice {
		_out[_i] = CreateTemporal(_e)
	}
	return _out
}


// TspatialSRID wraps MEOS C function tspatial_srid.
func TspatialSRID(temp Temporal) int32 {
	res := C.tspatial_srid(temp.Inner())
	return int32(res)
}


// TspatialSetSRID wraps MEOS C function tspatial_set_srid.
func TspatialSetSRID(temp Temporal, srid int32) Temporal {
	res := C.tspatial_set_srid(temp.Inner(), C.int32_t(srid))
	return CreateTemporal(res)
}


// TspatialTransform wraps MEOS C function tspatial_transform.
func TspatialTransform(temp Temporal, srid int32) Temporal {
	res := C.tspatial_transform(temp.Inner(), C.int32_t(srid))
	return CreateTemporal(res)
}


// TspatialTransformPipeline wraps MEOS C function tspatial_transform_pipeline.
func TspatialTransformPipeline(temp Temporal, pipelinestr string, srid int32, is_forward bool) Temporal {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	res := C.tspatial_transform_pipeline(temp.Inner(), _c_pipelinestr, C.int32_t(srid), C.bool(is_forward))
	return CreateTemporal(res)
}


// TgeoAtGeom wraps MEOS C function tgeo_at_geom.
func TgeoAtGeom(temp Temporal, gs *Geom) Temporal {
	res := C.tgeo_at_geom(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TgeoAtSTBOX wraps MEOS C function tgeo_at_stbox.
func TgeoAtSTBOX(temp Temporal, box *STBox, border_inc bool) Temporal {
	res := C.tgeo_at_stbox(temp.Inner(), box._inner, C.bool(border_inc))
	return CreateTemporal(res)
}


// TgeoAtValue wraps MEOS C function tgeo_at_value.
func TgeoAtValue(temp Temporal, gs *Geom) Temporal {
	res := C.tgeo_at_value(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TgeoMinusGeom wraps MEOS C function tgeo_minus_geom.
func TgeoMinusGeom(temp Temporal, gs *Geom) Temporal {
	res := C.tgeo_minus_geom(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TgeoMinusSTBOX wraps MEOS C function tgeo_minus_stbox.
func TgeoMinusSTBOX(temp Temporal, box *STBox, border_inc bool) Temporal {
	res := C.tgeo_minus_stbox(temp.Inner(), box._inner, C.bool(border_inc))
	return CreateTemporal(res)
}


// TgeoMinusValue wraps MEOS C function tgeo_minus_value.
func TgeoMinusValue(temp Temporal, gs *Geom) Temporal {
	res := C.tgeo_minus_value(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TpointAtElevation wraps MEOS C function tpoint_at_elevation.
func TpointAtElevation(temp Temporal, s *Span) Temporal {
	res := C.tpoint_at_elevation(temp.Inner(), s._inner)
	return CreateTemporal(res)
}


// TpointAtGeom wraps MEOS C function tpoint_at_geom.
func TpointAtGeom(temp Temporal, gs *Geom) Temporal {
	res := C.tpoint_at_geom(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TpointAtValue wraps MEOS C function tpoint_at_value.
func TpointAtValue(temp Temporal, gs *Geom) Temporal {
	res := C.tpoint_at_value(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TpointMinusElevation wraps MEOS C function tpoint_minus_elevation.
func TpointMinusElevation(temp Temporal, s *Span) Temporal {
	res := C.tpoint_minus_elevation(temp.Inner(), s._inner)
	return CreateTemporal(res)
}


// TpointMinusGeom wraps MEOS C function tpoint_minus_geom.
func TpointMinusGeom(temp Temporal, gs *Geom) Temporal {
	res := C.tpoint_minus_geom(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TpointMinusValue wraps MEOS C function tpoint_minus_value.
func TpointMinusValue(temp Temporal, gs *Geom) Temporal {
	res := C.tpoint_minus_value(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// AlwaysEqGeoTgeo wraps MEOS C function always_eq_geo_tgeo.
func AlwaysEqGeoTgeo(gs *Geom, temp Temporal) int {
	res := C.always_eq_geo_tgeo(gs._inner, temp.Inner())
	return int(res)
}


// AlwaysEqTgeoGeo wraps MEOS C function always_eq_tgeo_geo.
func AlwaysEqTgeoGeo(temp Temporal, gs *Geom) int {
	res := C.always_eq_tgeo_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AlwaysEqTgeoTgeo wraps MEOS C function always_eq_tgeo_tgeo.
func AlwaysEqTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.always_eq_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AlwaysNeGeoTgeo wraps MEOS C function always_ne_geo_tgeo.
func AlwaysNeGeoTgeo(gs *Geom, temp Temporal) int {
	res := C.always_ne_geo_tgeo(gs._inner, temp.Inner())
	return int(res)
}


// AlwaysNeTgeoGeo wraps MEOS C function always_ne_tgeo_geo.
func AlwaysNeTgeoGeo(temp Temporal, gs *Geom) int {
	res := C.always_ne_tgeo_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AlwaysNeTgeoTgeo wraps MEOS C function always_ne_tgeo_tgeo.
func AlwaysNeTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.always_ne_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverEqGeoTgeo wraps MEOS C function ever_eq_geo_tgeo.
func EverEqGeoTgeo(gs *Geom, temp Temporal) int {
	res := C.ever_eq_geo_tgeo(gs._inner, temp.Inner())
	return int(res)
}


// EverEqTgeoGeo wraps MEOS C function ever_eq_tgeo_geo.
func EverEqTgeoGeo(temp Temporal, gs *Geom) int {
	res := C.ever_eq_tgeo_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EverEqTgeoTgeo wraps MEOS C function ever_eq_tgeo_tgeo.
func EverEqTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_eq_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverNeGeoTgeo wraps MEOS C function ever_ne_geo_tgeo.
func EverNeGeoTgeo(gs *Geom, temp Temporal) int {
	res := C.ever_ne_geo_tgeo(gs._inner, temp.Inner())
	return int(res)
}


// EverNeTgeoGeo wraps MEOS C function ever_ne_tgeo_geo.
func EverNeTgeoGeo(temp Temporal, gs *Geom) int {
	res := C.ever_ne_tgeo_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EverNeTgeoTgeo wraps MEOS C function ever_ne_tgeo_tgeo.
func EverNeTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_ne_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TeqGeoTgeo wraps MEOS C function teq_geo_tgeo.
func TeqGeoTgeo(gs *Geom, temp Temporal) Temporal {
	res := C.teq_geo_tgeo(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TeqTgeoGeo wraps MEOS C function teq_tgeo_geo.
func TeqTgeoGeo(temp Temporal, gs *Geom) Temporal {
	res := C.teq_tgeo_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TneGeoTgeo wraps MEOS C function tne_geo_tgeo.
func TneGeoTgeo(gs *Geom, temp Temporal) Temporal {
	res := C.tne_geo_tgeo(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TneTgeoGeo wraps MEOS C function tne_tgeo_geo.
func TneTgeoGeo(temp Temporal, gs *Geom) Temporal {
	res := C.tne_tgeo_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TgeoStboxes wraps MEOS C function tgeo_stboxes.
func TgeoStboxes(temp Temporal) (*STBox, int) {
	var _out_count C.int
	res := C.tgeo_stboxes(temp.Inner(), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// TgeoSpaceBoxes wraps MEOS C function tgeo_space_boxes.
func TgeoSpaceBoxes(temp Temporal, xsize float64, ysize float64, zsize float64, sorigin *Geom, bitmatrix bool, border_inc bool) (*STBox, int) {
	var _out_count C.int
	res := C.tgeo_space_boxes(temp.Inner(), C.double(xsize), C.double(ysize), C.double(zsize), sorigin._inner, C.bool(bitmatrix), C.bool(border_inc), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// TgeoSpaceTimeBoxes wraps MEOS C function tgeo_space_time_boxes.
func TgeoSpaceTimeBoxes(temp Temporal, xsize float64, ysize float64, zsize float64, duration timeutil.Timedelta, sorigin *Geom, torigin int64, bitmatrix bool, border_inc bool) (*STBox, int) {
	var _out_count C.int
	res := C.tgeo_space_time_boxes(temp.Inner(), C.double(xsize), C.double(ysize), C.double(zsize), duration.Inner(), sorigin._inner, C.TimestampTz(torigin), C.bool(bitmatrix), C.bool(border_inc), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// TgeoSplitEachNStboxes wraps MEOS C function tgeo_split_each_n_stboxes.
func TgeoSplitEachNStboxes(temp Temporal, elem_count int) (*STBox, int) {
	var _out_count C.int
	res := C.tgeo_split_each_n_stboxes(temp.Inner(), C.int(elem_count), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// TgeoSplitNStboxes wraps MEOS C function tgeo_split_n_stboxes.
func TgeoSplitNStboxes(temp Temporal, box_count int) (*STBox, int) {
	var _out_count C.int
	res := C.tgeo_split_n_stboxes(temp.Inner(), C.int(box_count), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// AdjacentSTBOXTspatial wraps MEOS C function adjacent_stbox_tspatial.
func AdjacentSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.adjacent_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// AdjacentTspatialSTBOX wraps MEOS C function adjacent_tspatial_stbox.
func AdjacentTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.adjacent_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// AdjacentTspatialTspatial wraps MEOS C function adjacent_tspatial_tspatial.
func AdjacentTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.adjacent_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// ContainedSTBOXTspatial wraps MEOS C function contained_stbox_tspatial.
func ContainedSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.contained_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// ContainedTspatialSTBOX wraps MEOS C function contained_tspatial_stbox.
func ContainedTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.contained_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// ContainedTspatialTspatial wraps MEOS C function contained_tspatial_tspatial.
func ContainedTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.contained_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// ContainsSTBOXTspatial wraps MEOS C function contains_stbox_tspatial.
func ContainsSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.contains_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// ContainsTspatialSTBOX wraps MEOS C function contains_tspatial_stbox.
func ContainsTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.contains_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// ContainsTspatialTspatial wraps MEOS C function contains_tspatial_tspatial.
func ContainsTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.contains_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// OverlapsSTBOXTspatial wraps MEOS C function overlaps_stbox_tspatial.
func OverlapsSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.overlaps_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// OverlapsTspatialSTBOX wraps MEOS C function overlaps_tspatial_stbox.
func OverlapsTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.overlaps_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// OverlapsTspatialTspatial wraps MEOS C function overlaps_tspatial_tspatial.
func OverlapsTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.overlaps_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// SameSTBOXTspatial wraps MEOS C function same_stbox_tspatial.
func SameSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.same_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// SameTspatialSTBOX wraps MEOS C function same_tspatial_stbox.
func SameTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.same_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// SameTspatialTspatial wraps MEOS C function same_tspatial_tspatial.
func SameTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.same_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// AboveSTBOXTspatial wraps MEOS C function above_stbox_tspatial.
func AboveSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.above_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// AboveTspatialSTBOX wraps MEOS C function above_tspatial_stbox.
func AboveTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.above_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// AboveTspatialTspatial wraps MEOS C function above_tspatial_tspatial.
func AboveTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.above_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// AfterSTBOXTspatial wraps MEOS C function after_stbox_tspatial.
func AfterSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.after_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// AfterTspatialSTBOX wraps MEOS C function after_tspatial_stbox.
func AfterTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.after_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// AfterTspatialTspatial wraps MEOS C function after_tspatial_tspatial.
func AfterTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.after_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// BackSTBOXTspatial wraps MEOS C function back_stbox_tspatial.
func BackSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.back_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// BackTspatialSTBOX wraps MEOS C function back_tspatial_stbox.
func BackTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.back_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// BackTspatialTspatial wraps MEOS C function back_tspatial_tspatial.
func BackTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.back_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// BeforeSTBOXTspatial wraps MEOS C function before_stbox_tspatial.
func BeforeSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.before_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// BeforeTspatialSTBOX wraps MEOS C function before_tspatial_stbox.
func BeforeTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.before_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// BeforeTspatialTspatial wraps MEOS C function before_tspatial_tspatial.
func BeforeTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.before_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// BelowSTBOXTspatial wraps MEOS C function below_stbox_tspatial.
func BelowSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.below_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// BelowTspatialSTBOX wraps MEOS C function below_tspatial_stbox.
func BelowTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.below_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// BelowTspatialTspatial wraps MEOS C function below_tspatial_tspatial.
func BelowTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.below_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// FrontSTBOXTspatial wraps MEOS C function front_stbox_tspatial.
func FrontSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.front_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// FrontTspatialSTBOX wraps MEOS C function front_tspatial_stbox.
func FrontTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.front_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// FrontTspatialTspatial wraps MEOS C function front_tspatial_tspatial.
func FrontTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.front_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// LeftSTBOXTspatial wraps MEOS C function left_stbox_tspatial.
func LeftSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.left_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// LeftTspatialSTBOX wraps MEOS C function left_tspatial_stbox.
func LeftTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.left_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// LeftTspatialTspatial wraps MEOS C function left_tspatial_tspatial.
func LeftTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.left_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// OveraboveSTBOXTspatial wraps MEOS C function overabove_stbox_tspatial.
func OveraboveSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.overabove_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// OveraboveTspatialSTBOX wraps MEOS C function overabove_tspatial_stbox.
func OveraboveTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.overabove_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// OveraboveTspatialTspatial wraps MEOS C function overabove_tspatial_tspatial.
func OveraboveTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.overabove_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// OverafterSTBOXTspatial wraps MEOS C function overafter_stbox_tspatial.
func OverafterSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.overafter_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// OverafterTspatialSTBOX wraps MEOS C function overafter_tspatial_stbox.
func OverafterTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.overafter_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// OverafterTspatialTspatial wraps MEOS C function overafter_tspatial_tspatial.
func OverafterTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.overafter_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// OverbackSTBOXTspatial wraps MEOS C function overback_stbox_tspatial.
func OverbackSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.overback_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// OverbackTspatialSTBOX wraps MEOS C function overback_tspatial_stbox.
func OverbackTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.overback_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// OverbackTspatialTspatial wraps MEOS C function overback_tspatial_tspatial.
func OverbackTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.overback_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// OverbeforeSTBOXTspatial wraps MEOS C function overbefore_stbox_tspatial.
func OverbeforeSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.overbefore_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// OverbeforeTspatialSTBOX wraps MEOS C function overbefore_tspatial_stbox.
func OverbeforeTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.overbefore_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// OverbeforeTspatialTspatial wraps MEOS C function overbefore_tspatial_tspatial.
func OverbeforeTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.overbefore_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// OverbelowSTBOXTspatial wraps MEOS C function overbelow_stbox_tspatial.
func OverbelowSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.overbelow_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// OverbelowTspatialSTBOX wraps MEOS C function overbelow_tspatial_stbox.
func OverbelowTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.overbelow_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// OverbelowTspatialTspatial wraps MEOS C function overbelow_tspatial_tspatial.
func OverbelowTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.overbelow_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// OverfrontSTBOXTspatial wraps MEOS C function overfront_stbox_tspatial.
func OverfrontSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.overfront_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// OverfrontTspatialSTBOX wraps MEOS C function overfront_tspatial_stbox.
func OverfrontTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.overfront_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// OverfrontTspatialTspatial wraps MEOS C function overfront_tspatial_tspatial.
func OverfrontTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.overfront_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// OverleftSTBOXTspatial wraps MEOS C function overleft_stbox_tspatial.
func OverleftSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.overleft_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// OverleftTspatialSTBOX wraps MEOS C function overleft_tspatial_stbox.
func OverleftTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.overleft_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// OverleftTspatialTspatial wraps MEOS C function overleft_tspatial_tspatial.
func OverleftTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.overleft_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// OverrightSTBOXTspatial wraps MEOS C function overright_stbox_tspatial.
func OverrightSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.overright_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// OverrightTspatialSTBOX wraps MEOS C function overright_tspatial_stbox.
func OverrightTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.overright_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// OverrightTspatialTspatial wraps MEOS C function overright_tspatial_tspatial.
func OverrightTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.overright_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// RightSTBOXTspatial wraps MEOS C function right_stbox_tspatial.
func RightSTBOXTspatial(box *STBox, temp Temporal) bool {
	res := C.right_stbox_tspatial(box._inner, temp.Inner())
	return bool(res)
}


// RightTspatialSTBOX wraps MEOS C function right_tspatial_stbox.
func RightTspatialSTBOX(temp Temporal, box *STBox) bool {
	res := C.right_tspatial_stbox(temp.Inner(), box._inner)
	return bool(res)
}


// RightTspatialTspatial wraps MEOS C function right_tspatial_tspatial.
func RightTspatialTspatial(temp1 Temporal, temp2 Temporal) bool {
	res := C.right_tspatial_tspatial(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// AcontainsGeoTgeo wraps MEOS C function acontains_geo_tgeo.
func AcontainsGeoTgeo(gs *Geom, temp Temporal) int {
	res := C.acontains_geo_tgeo(gs._inner, temp.Inner())
	return int(res)
}


// AcontainsTgeoGeo wraps MEOS C function acontains_tgeo_geo.
func AcontainsTgeoGeo(temp Temporal, gs *Geom) int {
	res := C.acontains_tgeo_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AcontainsTgeoTgeo wraps MEOS C function acontains_tgeo_tgeo.
func AcontainsTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.acontains_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AcoversGeoTgeo wraps MEOS C function acovers_geo_tgeo.
func AcoversGeoTgeo(gs *Geom, temp Temporal) int {
	res := C.acovers_geo_tgeo(gs._inner, temp.Inner())
	return int(res)
}


// AcoversTgeoGeo wraps MEOS C function acovers_tgeo_geo.
func AcoversTgeoGeo(temp Temporal, gs *Geom) int {
	res := C.acovers_tgeo_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AcoversTgeoTgeo wraps MEOS C function acovers_tgeo_tgeo.
func AcoversTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.acovers_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AdisjointTgeoGeo wraps MEOS C function adisjoint_tgeo_geo.
func AdisjointTgeoGeo(temp Temporal, gs *Geom) int {
	res := C.adisjoint_tgeo_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AdisjointTgeoTgeo wraps MEOS C function adisjoint_tgeo_tgeo.
func AdisjointTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.adisjoint_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AdwithinTgeoGeo wraps MEOS C function adwithin_tgeo_geo.
func AdwithinTgeoGeo(temp Temporal, gs *Geom, dist float64) int {
	res := C.adwithin_tgeo_geo(temp.Inner(), gs._inner, C.double(dist))
	return int(res)
}


// AdwithinTgeoTgeo wraps MEOS C function adwithin_tgeo_tgeo.
func AdwithinTgeoTgeo(temp1 Temporal, temp2 Temporal, dist float64) int {
	res := C.adwithin_tgeo_tgeo(temp1.Inner(), temp2.Inner(), C.double(dist))
	return int(res)
}


// AintersectsTgeoGeo wraps MEOS C function aintersects_tgeo_geo.
func AintersectsTgeoGeo(temp Temporal, gs *Geom) int {
	res := C.aintersects_tgeo_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AintersectsTgeoTgeo wraps MEOS C function aintersects_tgeo_tgeo.
func AintersectsTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.aintersects_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AtouchesTgeoGeo wraps MEOS C function atouches_tgeo_geo.
func AtouchesTgeoGeo(temp Temporal, gs *Geom) int {
	res := C.atouches_tgeo_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AtouchesTgeoTgeo wraps MEOS C function atouches_tgeo_tgeo.
func AtouchesTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.atouches_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AtouchesTpointGeo wraps MEOS C function atouches_tpoint_geo.
func AtouchesTpointGeo(temp Temporal, gs *Geom) int {
	res := C.atouches_tpoint_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EcontainsGeoTgeo wraps MEOS C function econtains_geo_tgeo.
func EcontainsGeoTgeo(gs *Geom, temp Temporal) int {
	res := C.econtains_geo_tgeo(gs._inner, temp.Inner())
	return int(res)
}


// EcontainsTgeoGeo wraps MEOS C function econtains_tgeo_geo.
func EcontainsTgeoGeo(temp Temporal, gs *Geom) int {
	res := C.econtains_tgeo_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EcontainsTgeoTgeo wraps MEOS C function econtains_tgeo_tgeo.
func EcontainsTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.econtains_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EcoversGeoTgeo wraps MEOS C function ecovers_geo_tgeo.
func EcoversGeoTgeo(gs *Geom, temp Temporal) int {
	res := C.ecovers_geo_tgeo(gs._inner, temp.Inner())
	return int(res)
}


// EcoversTgeoGeo wraps MEOS C function ecovers_tgeo_geo.
func EcoversTgeoGeo(temp Temporal, gs *Geom) int {
	res := C.ecovers_tgeo_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EcoversTgeoTgeo wraps MEOS C function ecovers_tgeo_tgeo.
func EcoversTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.ecovers_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EdisjointTgeoGeo wraps MEOS C function edisjoint_tgeo_geo.
func EdisjointTgeoGeo(temp Temporal, gs *Geom) int {
	res := C.edisjoint_tgeo_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EdisjointTgeoTgeo wraps MEOS C function edisjoint_tgeo_tgeo.
func EdisjointTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.edisjoint_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EdwithinTgeoGeo wraps MEOS C function edwithin_tgeo_geo.
func EdwithinTgeoGeo(temp Temporal, gs *Geom, dist float64) int {
	res := C.edwithin_tgeo_geo(temp.Inner(), gs._inner, C.double(dist))
	return int(res)
}


// EdwithinTgeoTgeo wraps MEOS C function edwithin_tgeo_tgeo.
func EdwithinTgeoTgeo(temp1 Temporal, temp2 Temporal, dist float64) int {
	res := C.edwithin_tgeo_tgeo(temp1.Inner(), temp2.Inner(), C.double(dist))
	return int(res)
}


// EintersectsTgeoGeo wraps MEOS C function eintersects_tgeo_geo.
func EintersectsTgeoGeo(temp Temporal, gs *Geom) int {
	res := C.eintersects_tgeo_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EintersectsTgeoTgeo wraps MEOS C function eintersects_tgeo_tgeo.
func EintersectsTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.eintersects_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EtouchesTgeoGeo wraps MEOS C function etouches_tgeo_geo.
func EtouchesTgeoGeo(temp Temporal, gs *Geom) int {
	res := C.etouches_tgeo_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EtouchesTgeoTgeo wraps MEOS C function etouches_tgeo_tgeo.
func EtouchesTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.etouches_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EtouchesTpointGeo wraps MEOS C function etouches_tpoint_geo.
func EtouchesTpointGeo(temp Temporal, gs *Geom) int {
	res := C.etouches_tpoint_geo(temp.Inner(), gs._inner)
	return int(res)
}


// TcontainsGeoTgeo wraps MEOS C function tcontains_geo_tgeo.
func TcontainsGeoTgeo(gs *Geom, temp Temporal) Temporal {
	res := C.tcontains_geo_tgeo(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TcontainsTgeoGeo wraps MEOS C function tcontains_tgeo_geo.
func TcontainsTgeoGeo(temp Temporal, gs *Geom) Temporal {
	res := C.tcontains_tgeo_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TcontainsTgeoTgeo wraps MEOS C function tcontains_tgeo_tgeo.
func TcontainsTgeoTgeo(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tcontains_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TcoversGeoTgeo wraps MEOS C function tcovers_geo_tgeo.
func TcoversGeoTgeo(gs *Geom, temp Temporal) Temporal {
	res := C.tcovers_geo_tgeo(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TcoversTgeoGeo wraps MEOS C function tcovers_tgeo_geo.
func TcoversTgeoGeo(temp Temporal, gs *Geom) Temporal {
	res := C.tcovers_tgeo_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TcoversTgeoTgeo wraps MEOS C function tcovers_tgeo_tgeo.
func TcoversTgeoTgeo(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tcovers_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TdisjointGeoTgeo wraps MEOS C function tdisjoint_geo_tgeo.
func TdisjointGeoTgeo(gs *Geom, temp Temporal) Temporal {
	res := C.tdisjoint_geo_tgeo(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TdisjointTgeoGeo wraps MEOS C function tdisjoint_tgeo_geo.
func TdisjointTgeoGeo(temp Temporal, gs *Geom) Temporal {
	res := C.tdisjoint_tgeo_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TdisjointTgeoTgeo wraps MEOS C function tdisjoint_tgeo_tgeo.
func TdisjointTgeoTgeo(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tdisjoint_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TdwithinGeoTgeo wraps MEOS C function tdwithin_geo_tgeo.
func TdwithinGeoTgeo(gs *Geom, temp Temporal, dist float64) Temporal {
	res := C.tdwithin_geo_tgeo(gs._inner, temp.Inner(), C.double(dist))
	return CreateTemporal(res)
}


// TdwithinTgeoGeo wraps MEOS C function tdwithin_tgeo_geo.
func TdwithinTgeoGeo(temp Temporal, gs *Geom, dist float64) Temporal {
	res := C.tdwithin_tgeo_geo(temp.Inner(), gs._inner, C.double(dist))
	return CreateTemporal(res)
}


// TdwithinTgeoTgeo wraps MEOS C function tdwithin_tgeo_tgeo.
func TdwithinTgeoTgeo(temp1 Temporal, temp2 Temporal, dist float64) Temporal {
	res := C.tdwithin_tgeo_tgeo(temp1.Inner(), temp2.Inner(), C.double(dist))
	return CreateTemporal(res)
}


// TintersectsGeoTgeo wraps MEOS C function tintersects_geo_tgeo.
func TintersectsGeoTgeo(gs *Geom, temp Temporal) Temporal {
	res := C.tintersects_geo_tgeo(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TintersectsTgeoGeo wraps MEOS C function tintersects_tgeo_geo.
func TintersectsTgeoGeo(temp Temporal, gs *Geom) Temporal {
	res := C.tintersects_tgeo_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TintersectsTgeoTgeo wraps MEOS C function tintersects_tgeo_tgeo.
func TintersectsTgeoTgeo(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tintersects_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TtouchesGeoTgeo wraps MEOS C function ttouches_geo_tgeo.
func TtouchesGeoTgeo(gs *Geom, temp Temporal) Temporal {
	res := C.ttouches_geo_tgeo(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TtouchesTgeoGeo wraps MEOS C function ttouches_tgeo_geo.
func TtouchesTgeoGeo(temp Temporal, gs *Geom) Temporal {
	res := C.ttouches_tgeo_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TtouchesTgeoTgeo wraps MEOS C function ttouches_tgeo_tgeo.
func TtouchesTgeoTgeo(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.ttouches_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TODO edwithin_tgeoarr_tgeoarr: unsupported param const Temporal **
// func EdwithinTgeoarrTgeoarr(...) { /* not yet handled by codegen */ }


// TODO adwithin_tgeoarr_tgeoarr: unsupported param const Temporal **
// func AdwithinTgeoarrTgeoarr(...) { /* not yet handled by codegen */ }


// TODO eintersects_tgeoarr_tgeoarr: unsupported param const Temporal **
// func EintersectsTgeoarrTgeoarr(...) { /* not yet handled by codegen */ }


// TODO aintersects_tgeoarr_tgeoarr: unsupported param const Temporal **
// func AintersectsTgeoarrTgeoarr(...) { /* not yet handled by codegen */ }


// TODO etouches_tgeoarr_tgeoarr: unsupported param const Temporal **
// func EtouchesTgeoarrTgeoarr(...) { /* not yet handled by codegen */ }


// TODO atouches_tgeoarr_tgeoarr: unsupported param const Temporal **
// func AtouchesTgeoarrTgeoarr(...) { /* not yet handled by codegen */ }


// TODO edisjoint_tgeoarr_tgeoarr: unsupported param const Temporal **
// func EdisjointTgeoarrTgeoarr(...) { /* not yet handled by codegen */ }


// TODO adisjoint_tgeoarr_tgeoarr: unsupported param const Temporal **
// func AdisjointTgeoarrTgeoarr(...) { /* not yet handled by codegen */ }


// TODO tdwithin_tgeoarr_tgeoarr: unsupported param const Temporal **
// func TdwithinTgeoarrTgeoarr(...) { /* not yet handled by codegen */ }


// TODO tintersects_tgeoarr_tgeoarr: unsupported param const Temporal **
// func TintersectsTgeoarrTgeoarr(...) { /* not yet handled by codegen */ }


// TODO ttouches_tgeoarr_tgeoarr: unsupported param const Temporal **
// func TtouchesTgeoarrTgeoarr(...) { /* not yet handled by codegen */ }


// TODO tdisjoint_tgeoarr_tgeoarr: unsupported param const Temporal **
// func TdisjointTgeoarrTgeoarr(...) { /* not yet handled by codegen */ }


// TdistanceTgeoGeo wraps MEOS C function tdistance_tgeo_geo.
func TdistanceTgeoGeo(temp Temporal, gs *Geom) Temporal {
	res := C.tdistance_tgeo_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TdistanceTgeoTgeo wraps MEOS C function tdistance_tgeo_tgeo.
func TdistanceTgeoTgeo(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tdistance_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// NadSTBOXGeo wraps MEOS C function nad_stbox_geo.
func NadSTBOXGeo(box *STBox, gs *Geom) float64 {
	res := C.nad_stbox_geo(box._inner, gs._inner)
	return float64(res)
}


// NadSTBOXSTBOX wraps MEOS C function nad_stbox_stbox.
func NadSTBOXSTBOX(box1 *STBox, box2 *STBox) float64 {
	res := C.nad_stbox_stbox(box1._inner, box2._inner)
	return float64(res)
}


// NadTgeoGeo wraps MEOS C function nad_tgeo_geo.
func NadTgeoGeo(temp Temporal, gs *Geom) float64 {
	res := C.nad_tgeo_geo(temp.Inner(), gs._inner)
	return float64(res)
}


// NadTgeoSTBOX wraps MEOS C function nad_tgeo_stbox.
func NadTgeoSTBOX(temp Temporal, box *STBox) float64 {
	res := C.nad_tgeo_stbox(temp.Inner(), box._inner)
	return float64(res)
}


// NadTgeoTgeo wraps MEOS C function nad_tgeo_tgeo.
func NadTgeoTgeo(temp1 Temporal, temp2 Temporal) float64 {
	res := C.nad_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return float64(res)
}


// NaiTgeoGeo wraps MEOS C function nai_tgeo_geo.
func NaiTgeoGeo(temp Temporal, gs *Geom) TInstant {
	res := C.nai_tgeo_geo(temp.Inner(), gs._inner)
	return TInstant{_inner: res}
}


// NaiTgeoTgeo wraps MEOS C function nai_tgeo_tgeo.
func NaiTgeoTgeo(temp1 Temporal, temp2 Temporal) TInstant {
	res := C.nai_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return TInstant{_inner: res}
}


// ShortestlineTgeoGeo wraps MEOS C function shortestline_tgeo_geo.
func ShortestlineTgeoGeo(temp Temporal, gs *Geom) *Geom {
	res := C.shortestline_tgeo_geo(temp.Inner(), gs._inner)
	return &Geom{_inner: res}
}


// ShortestlineTgeoTgeo wraps MEOS C function shortestline_tgeo_tgeo.
func ShortestlineTgeoTgeo(temp1 Temporal, temp2 Temporal) *Geom {
	res := C.shortestline_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return &Geom{_inner: res}
}


// TODO tgeoarr_tgeoarr_mindist: unsupported param const Temporal **
// func TgeoarrTgeoarrMindist(...) { /* not yet handled by codegen */ }


// MindistanceTgeoTgeo wraps MEOS C function mindistance_tgeo_tgeo.
func MindistanceTgeoTgeo(temp1 Temporal, temp2 Temporal, threshold float64) float64 {
	res := C.mindistance_tgeo_tgeo(temp1.Inner(), temp2.Inner(), C.double(threshold))
	return float64(res)
}


// TpointTcentroidFinalfn wraps MEOS C function tpoint_tcentroid_finalfn.
func TpointTcentroidFinalfn(state *SkipList) Temporal {
	res := C.tpoint_tcentroid_finalfn(state._inner)
	return CreateTemporal(res)
}


// TpointTcentroidTransfn wraps MEOS C function tpoint_tcentroid_transfn.
func TpointTcentroidTransfn(state *SkipList, temp Temporal) *SkipList {
	res := C.tpoint_tcentroid_transfn(state._inner, temp.Inner())
	return &SkipList{_inner: res}
}


// TspatialExtentTransfn wraps MEOS C function tspatial_extent_transfn.
func TspatialExtentTransfn(box *STBox, temp Temporal) *STBox {
	res := C.tspatial_extent_transfn(box._inner, temp.Inner())
	return &STBox{_inner: res}
}


// STBOXGetSpaceTile wraps MEOS C function stbox_get_space_tile.
func STBOXGetSpaceTile(point *Geom, xsize float64, ysize float64, zsize float64, sorigin *Geom) *STBox {
	res := C.stbox_get_space_tile(point._inner, C.double(xsize), C.double(ysize), C.double(zsize), sorigin._inner)
	return &STBox{_inner: res}
}


// STBOXGetSpaceTimeTile wraps MEOS C function stbox_get_space_time_tile.
func STBOXGetSpaceTimeTile(point *Geom, t int64, xsize float64, ysize float64, zsize float64, duration timeutil.Timedelta, sorigin *Geom, torigin int64) *STBox {
	res := C.stbox_get_space_time_tile(point._inner, C.TimestampTz(t), C.double(xsize), C.double(ysize), C.double(zsize), duration.Inner(), sorigin._inner, C.TimestampTz(torigin))
	return &STBox{_inner: res}
}


// STBOXGetTimeTile wraps MEOS C function stbox_get_time_tile.
func STBOXGetTimeTile(t int64, duration timeutil.Timedelta, torigin int64) *STBox {
	res := C.stbox_get_time_tile(C.TimestampTz(t), duration.Inner(), C.TimestampTz(torigin))
	return &STBox{_inner: res}
}


// STBOXSpaceTiles wraps MEOS C function stbox_space_tiles.
func STBOXSpaceTiles(bounds *STBox, xsize float64, ysize float64, zsize float64, sorigin *Geom, border_inc bool) (*STBox, int) {
	var _out_count C.int
	res := C.stbox_space_tiles(bounds._inner, C.double(xsize), C.double(ysize), C.double(zsize), sorigin._inner, C.bool(border_inc), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// STBOXSpaceTimeTiles wraps MEOS C function stbox_space_time_tiles.
func STBOXSpaceTimeTiles(bounds *STBox, xsize float64, ysize float64, zsize float64, duration timeutil.Timedelta, sorigin *Geom, torigin int64, border_inc bool) (*STBox, int) {
	var _out_count C.int
	res := C.stbox_space_time_tiles(bounds._inner, C.double(xsize), C.double(ysize), C.double(zsize), duration.Inner(), sorigin._inner, C.TimestampTz(torigin), C.bool(border_inc), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// STBOXTimeTiles wraps MEOS C function stbox_time_tiles.
func STBOXTimeTiles(bounds *STBox, duration timeutil.Timedelta, torigin int64, border_inc bool) (*STBox, int) {
	var _out_count C.int
	res := C.stbox_time_tiles(bounds._inner, duration.Inner(), C.TimestampTz(torigin), C.bool(border_inc), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// TODO tgeo_space_split: unsupported return type SpaceSplit
// func TgeoSpaceSplit(...) { /* not yet handled by codegen */ }


// TODO tgeo_space_time_split: unsupported return type SpaceTimeSplit
// func TgeoSpaceTimeSplit(...) { /* not yet handled by codegen */ }


// GeoClusterKmeans wraps MEOS C function geo_cluster_kmeans.
func GeoClusterKmeans(geoms []*Geom, k uint32) []int {
	_c_geoms := make([]*C.GSERIALIZED, len(geoms))
	for _i, _v := range geoms { _c_geoms[_i] = _v._inner }
	var _out_count C.int
	res := C.geo_cluster_kmeans((**C.GSERIALIZED)(unsafe.Pointer(&_c_geoms[0])), C.uint32_t(len(geoms)), C.uint32_t(k), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((*C.int)(unsafe.Pointer(res)), _n)
	_out := make([]int, _n)
	for _i, _e := range _slice {
		_out[_i] = int(_e)
	}
	return _out
}


// GeoClusterDbscan wraps MEOS C function geo_cluster_dbscan.
func GeoClusterDbscan(geoms []*Geom, tolerance float64, minpoints int) []uint32 {
	_c_geoms := make([]*C.GSERIALIZED, len(geoms))
	for _i, _v := range geoms { _c_geoms[_i] = _v._inner }
	var _out_count C.int
	res := C.geo_cluster_dbscan((**C.GSERIALIZED)(unsafe.Pointer(&_c_geoms[0])), C.uint32_t(len(geoms)), C.double(tolerance), C.int(minpoints), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((*C.uint32_t)(unsafe.Pointer(res)), _n)
	_out := make([]uint32, _n)
	for _i, _e := range _slice {
		_out[_i] = uint32(_e)
	}
	return _out
}


// GeoClusterIntersecting wraps MEOS C function geo_cluster_intersecting.
func GeoClusterIntersecting(geoms []*Geom) []*Geom {
	_c_geoms := make([]*C.GSERIALIZED, len(geoms))
	for _i, _v := range geoms { _c_geoms[_i] = _v._inner }
	var _out_count C.int
	res := C.geo_cluster_intersecting((**C.GSERIALIZED)(unsafe.Pointer(&_c_geoms[0])), C.uint32_t(len(geoms)), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.GSERIALIZED)(unsafe.Pointer(res)), _n)
	_out := make([]*Geom, _n)
	for _i, _e := range _slice {
		_out[_i] = &Geom{_inner: _e}
	}
	return _out
}


// GeoClusterWithin wraps MEOS C function geo_cluster_within.
func GeoClusterWithin(geoms []*Geom, tolerance float64) []*Geom {
	_c_geoms := make([]*C.GSERIALIZED, len(geoms))
	for _i, _v := range geoms { _c_geoms[_i] = _v._inner }
	var _out_count C.int
	res := C.geo_cluster_within((**C.GSERIALIZED)(unsafe.Pointer(&_c_geoms[0])), C.uint32_t(len(geoms)), C.double(tolerance), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.GSERIALIZED)(unsafe.Pointer(res)), _n)
	_out := make([]*Geom, _n)
	for _i, _e := range _slice {
		_out[_i] = &Geom{_inner: _e}
	}
	return _out
}

