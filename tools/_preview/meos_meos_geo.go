package generated

// #include <stddef.h>
import "C"
import (
	"unsafe"

	"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

// TODO geo_as_ewkb: unsupported param const int *
// func GeoAsEWKB(...) { /* not yet handled by codegen */ }


// TODO geo_as_ewkt: unsupported param const int *
// func GeoAsEWKT(...) { /* not yet handled by codegen */ }


// TODO geo_as_geojson: unsupported param const int *
// func GeoAsGeojson(...) { /* not yet handled by codegen */ }


// TODO geo_as_hexewkb: unsupported param const int *
// func GeoAsHexewkb(...) { /* not yet handled by codegen */ }


// TODO geo_as_text: unsupported param const int *
// func GeoAsText(...) { /* not yet handled by codegen */ }


// GeoFromEWKB wraps MEOS C function geo_from_ewkb.
func GeoFromEWKB(wkb []byte, srid int) []int {
	res := C.geo_from_ewkb((*C.uint8_t)(unsafe.Pointer(&wkb[0])), C.size_t(len(wkb)), C.int(srid))
	_n := len(wkb)
	_slice := unsafe.Slice((*C.int)(unsafe.Pointer(res)), _n)
	_out := make([]int, _n)
	for _i, _e := range _slice {
		_out[_i] = int(_e)
	}
	return _out
}


// TODO geo_from_geojson: unsupported return type int *
// func GeoFromGeojson(...) { /* not yet handled by codegen */ }


// TODO geo_from_text: unsupported return type int *
// func GeoFromText(...) { /* not yet handled by codegen */ }


// TODO geo_out: unsupported param const int *
// func GeoOut(...) { /* not yet handled by codegen */ }


// TODO geog_from_binary: unsupported return type int *
// func GeogFromBinary(...) { /* not yet handled by codegen */ }


// TODO geog_from_hexewkb: unsupported return type int *
// func GeogFromHexewkb(...) { /* not yet handled by codegen */ }


// TODO geog_in: unsupported return type int *
// func GeogIn(...) { /* not yet handled by codegen */ }


// TODO geom_from_hexewkb: unsupported return type int *
// func GeomFromHexewkb(...) { /* not yet handled by codegen */ }


// TODO geom_in: unsupported return type int *
// func GeomIn(...) { /* not yet handled by codegen */ }


// TODO box3d_make: unsupported return type int *
// func Box3dMake(...) { /* not yet handled by codegen */ }


// TODO box3d_out: unsupported param const int *
// func Box3dOut(...) { /* not yet handled by codegen */ }


// TODO gbox_make: unsupported return type int *
// func GboxMake(...) { /* not yet handled by codegen */ }


// TODO gbox_out: unsupported param const int *
// func GboxOut(...) { /* not yet handled by codegen */ }


// TODO geo_copy: unsupported return type int *
// func GeoCopy(...) { /* not yet handled by codegen */ }


// TODO geogpoint_make2d: unsupported return type int *
// func GeogpointMake2d(...) { /* not yet handled by codegen */ }


// TODO geogpoint_make3dz: unsupported return type int *
// func GeogpointMake3dz(...) { /* not yet handled by codegen */ }


// TODO geompoint_make2d: unsupported return type int *
// func GeompointMake2d(...) { /* not yet handled by codegen */ }


// TODO geompoint_make3dz: unsupported return type int *
// func GeompointMake3dz(...) { /* not yet handled by codegen */ }


// TODO geom_to_geog: unsupported return type int *
// func GeomToGeog(...) { /* not yet handled by codegen */ }


// TODO geog_to_geom: unsupported return type int *
// func GeogToGeom(...) { /* not yet handled by codegen */ }


// TODO geo_is_empty: unsupported param const int *
// func GeoIsEmpty(...) { /* not yet handled by codegen */ }


// TODO geo_is_unitary: unsupported param const int *
// func GeoIsUnitary(...) { /* not yet handled by codegen */ }


// GeoTypename wraps MEOS C function geo_typename.
func GeoTypename(type_ int) string {
	res := C.geo_typename(C.int(type_))
	return C.GoString(res)
}


// TODO geog_area: unsupported param const int *
// func GeogArea(...) { /* not yet handled by codegen */ }


// TODO geog_centroid: unsupported return type int *
// func GeogCentroid(...) { /* not yet handled by codegen */ }


// TODO geog_length: unsupported param const int *
// func GeogLength(...) { /* not yet handled by codegen */ }


// TODO geog_perimeter: unsupported param const int *
// func GeogPerimeter(...) { /* not yet handled by codegen */ }


// TODO geom_azimuth: unsupported param const int *
// func GeomAzimuth(...) { /* not yet handled by codegen */ }


// TODO geom_length: unsupported param const int *
// func GeomLength(...) { /* not yet handled by codegen */ }


// TODO geom_perimeter: unsupported param const int *
// func GeomPerimeter(...) { /* not yet handled by codegen */ }


// TODO line_numpoints: unsupported param const int *
// func LineNumpoints(...) { /* not yet handled by codegen */ }


// TODO line_point_n: unsupported return type int *
// func LinePointN(...) { /* not yet handled by codegen */ }


// TODO geo_reverse: unsupported return type int *
// func GeoReverse(...) { /* not yet handled by codegen */ }


// TODO geo_round: unsupported return type int *
// func GeoRound(...) { /* not yet handled by codegen */ }


// TODO geo_set_srid: unsupported return type int *
// func GeoSetSRID(...) { /* not yet handled by codegen */ }


// TODO geo_srid: unsupported param const int *
// func GeoSRID(...) { /* not yet handled by codegen */ }


// TODO geo_transform: unsupported return type int *
// func GeoTransform(...) { /* not yet handled by codegen */ }


// TODO geo_transform_pipeline: unsupported return type int *
// func GeoTransformPipeline(...) { /* not yet handled by codegen */ }


// TODO geo_collect_garray: unsupported return type int *
// func GeoCollectGarray(...) { /* not yet handled by codegen */ }


// TODO geo_makeline_garray: unsupported return type int *
// func GeoMakelineGarray(...) { /* not yet handled by codegen */ }


// TODO geo_num_points: unsupported param const int *
// func GeoNumPoints(...) { /* not yet handled by codegen */ }


// TODO geo_num_geos: unsupported param const int *
// func GeoNumGeos(...) { /* not yet handled by codegen */ }


// TODO geo_geo_n: unsupported return type int *
// func GeoGeoN(...) { /* not yet handled by codegen */ }


// TODO geo_pointarr: unsupported return type int **
// func GeoPointarr(...) { /* not yet handled by codegen */ }


// TODO geo_points: unsupported return type int *
// func GeoPoints(...) { /* not yet handled by codegen */ }


// TODO geom_array_union: unsupported return type int *
// func GeomArrayUnion(...) { /* not yet handled by codegen */ }


// TODO geom_boundary: unsupported return type int *
// func GeomBoundary(...) { /* not yet handled by codegen */ }


// TODO geom_buffer: unsupported return type int *
// func GeomBuffer(...) { /* not yet handled by codegen */ }


// TODO geom_centroid: unsupported return type int *
// func GeomCentroid(...) { /* not yet handled by codegen */ }


// TODO geom_convex_hull: unsupported return type int *
// func GeomConvexHull(...) { /* not yet handled by codegen */ }


// TODO geom_difference2d: unsupported return type int *
// func GeomDifference2d(...) { /* not yet handled by codegen */ }


// TODO geom_intersection2d: unsupported return type int *
// func GeomIntersection2d(...) { /* not yet handled by codegen */ }


// TODO geom_intersection2d_coll: unsupported return type int *
// func GeomIntersection2dColl(...) { /* not yet handled by codegen */ }


// TODO geom_min_bounding_radius: unsupported return type int *
// func GeomMinBoundingRadius(...) { /* not yet handled by codegen */ }


// TODO geom_shortestline2d: unsupported return type int *
// func GeomShortestline2d(...) { /* not yet handled by codegen */ }


// TODO geom_shortestline3d: unsupported return type int *
// func GeomShortestline3d(...) { /* not yet handled by codegen */ }


// TODO geom_unary_union: unsupported return type int *
// func GeomUnaryUnion(...) { /* not yet handled by codegen */ }


// TODO line_interpolate_point: unsupported return type int *
// func LineInterpolatePoint(...) { /* not yet handled by codegen */ }


// TODO line_locate_point: unsupported param const int *
// func LineLocatePoint(...) { /* not yet handled by codegen */ }


// TODO line_substring: unsupported return type int *
// func LineSubstring(...) { /* not yet handled by codegen */ }


// TODO geog_dwithin: unsupported param const int *
// func GeogDwithin(...) { /* not yet handled by codegen */ }


// TODO geog_intersects: unsupported param const int *
// func GeogIntersects(...) { /* not yet handled by codegen */ }


// TODO geom_contains: unsupported param const int *
// func GeomContains(...) { /* not yet handled by codegen */ }


// TODO geom_covers: unsupported param const int *
// func GeomCovers(...) { /* not yet handled by codegen */ }


// TODO geom_disjoint2d: unsupported param const int *
// func GeomDisjoint2d(...) { /* not yet handled by codegen */ }


// TODO geom_dwithin2d: unsupported param const int *
// func GeomDwithin2d(...) { /* not yet handled by codegen */ }


// TODO geom_dwithin3d: unsupported param const int *
// func GeomDwithin3d(...) { /* not yet handled by codegen */ }


// TODO geom_intersects2d: unsupported param const int *
// func GeomIntersects2d(...) { /* not yet handled by codegen */ }


// TODO geom_intersects3d: unsupported param const int *
// func GeomIntersects3d(...) { /* not yet handled by codegen */ }


// TODO geom_relate_pattern: unsupported param const int *
// func GeomRelatePattern(...) { /* not yet handled by codegen */ }


// TODO geom_touches: unsupported param const int *
// func GeomTouches(...) { /* not yet handled by codegen */ }


// TODO geo_stboxes: unsupported param const int *
// func GeoStboxes(...) { /* not yet handled by codegen */ }


// TODO geo_split_each_n_stboxes: unsupported param const int *
// func GeoSplitEachNStboxes(...) { /* not yet handled by codegen */ }


// TODO geo_split_n_stboxes: unsupported param const int *
// func GeoSplitNStboxes(...) { /* not yet handled by codegen */ }


// TODO geog_distance: unsupported param const int *
// func GeogDistance(...) { /* not yet handled by codegen */ }


// TODO geom_distance2d: unsupported param const int *
// func GeomDistance2d(...) { /* not yet handled by codegen */ }


// TODO geom_distance3d: unsupported param const int *
// func GeomDistance3d(...) { /* not yet handled by codegen */ }


// TODO geo_equals: unsupported param const int *
// func GeoEquals(...) { /* not yet handled by codegen */ }


// TODO geo_same: unsupported param const int *
// func GeoSame(...) { /* not yet handled by codegen */ }


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


// TODO geoset_make: unsupported param int **
// func GeosetMake(...) { /* not yet handled by codegen */ }


// TODO geo_to_set: unsupported param const int *
// func GeoToSet(...) { /* not yet handled by codegen */ }


// TODO geoset_end_value: unsupported return type int *
// func GeosetEndValue(...) { /* not yet handled by codegen */ }


// TODO geoset_start_value: unsupported return type int *
// func GeosetStartValue(...) { /* not yet handled by codegen */ }


// TODO geoset_value_n: unhandled OUTPUT_SCALAR shape int **
// func GeosetValueN(...) { /* not yet handled by codegen */ }


// TODO geoset_values: unsupported return type int **
// func GeosetValues(...) { /* not yet handled by codegen */ }


// TODO contained_geo_set: unsupported param const int *
// func ContainedGeoSet(...) { /* not yet handled by codegen */ }


// TODO contains_set_geo: unsupported param int *
// func ContainsSetGeo(...) { /* not yet handled by codegen */ }


// TODO geo_union_transfn: unsupported param const int *
// func GeoUnionTransfn(...) { /* not yet handled by codegen */ }


// TODO intersection_geo_set: unsupported param const int *
// func IntersectionGeoSet(...) { /* not yet handled by codegen */ }


// TODO intersection_set_geo: unsupported param const int *
// func IntersectionSetGeo(...) { /* not yet handled by codegen */ }


// TODO minus_geo_set: unsupported param const int *
// func MinusGeoSet(...) { /* not yet handled by codegen */ }


// TODO minus_set_geo: unsupported param const int *
// func MinusSetGeo(...) { /* not yet handled by codegen */ }


// TODO union_geo_set: unsupported param const int *
// func UnionGeoSet(...) { /* not yet handled by codegen */ }


// TODO union_set_geo: unsupported param const int *
// func UnionSetGeo(...) { /* not yet handled by codegen */ }


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


// TODO geo_timestamptz_to_stbox: unsupported param const int *
// func GeoTimestamptzToSTBOX(...) { /* not yet handled by codegen */ }


// TODO geo_tstzspan_to_stbox: unsupported param const int *
// func GeoTstzspanToSTBOX(...) { /* not yet handled by codegen */ }


// STBOXCopy wraps MEOS C function stbox_copy.
func STBOXCopy(box *STBox) *STBox {
	res := C.stbox_copy(box._inner)
	return &STBox{_inner: res}
}


// STBOXMake wraps MEOS C function stbox_make.
func STBOXMake(hasx bool, hasz bool, geodetic bool, srid int, xmin float64, xmax float64, ymin float64, ymax float64, zmin float64, zmax float64, s *Span) *STBox {
	res := C.stbox_make(C.bool(hasx), C.bool(hasz), C.bool(geodetic), C.int(srid), C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.double(zmin), C.double(zmax), s._inner)
	return &STBox{_inner: res}
}


// TODO geo_to_stbox: unsupported param const int *
// func GeoToSTBOX(...) { /* not yet handled by codegen */ }


// SpatialsetToSTBOX wraps MEOS C function spatialset_to_stbox.
func SpatialsetToSTBOX(s *Set) *STBox {
	res := C.spatialset_to_stbox(s._inner)
	return &STBox{_inner: res}
}


// TODO stbox_to_box3d: unsupported return type int *
// func STBOXToBox3d(...) { /* not yet handled by codegen */ }


// TODO stbox_to_gbox: unsupported return type int *
// func STBOXToGbox(...) { /* not yet handled by codegen */ }


// TODO stbox_to_geo: unsupported return type int *
// func STBOXToGeo(...) { /* not yet handled by codegen */ }


// STBOXToTstzspan wraps MEOS C function stbox_to_tstzspan.
func STBOXToTstzspan(box *STBox) *Span {
	res := C.stbox_to_tstzspan(box._inner)
	return &Span{_inner: res}
}


// TimestamptzToSTBOX wraps MEOS C function timestamptz_to_stbox.
func TimestamptzToSTBOX(t int) *STBox {
	res := C.timestamptz_to_stbox(C.int(t))
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
func STBOXTmax(box *STBox) (bool, int) {
	var _out_result C.int
	res := C.stbox_tmax(box._inner, &_out_result)
	return bool(res), int(_out_result)
}


// STBOXTmaxInc wraps MEOS C function stbox_tmax_inc.
func STBOXTmaxInc(box *STBox) (bool, bool) {
	var _out_result C.bool
	res := C.stbox_tmax_inc(box._inner, &_out_result)
	return bool(res), bool(_out_result)
}


// STBOXTmin wraps MEOS C function stbox_tmin.
func STBOXTmin(box *STBox) (bool, int) {
	var _out_result C.int
	res := C.stbox_tmin(box._inner, &_out_result)
	return bool(res), int(_out_result)
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


// TODO stbox_expand_time: unsupported param const int *
// func STBOXExpandTime(...) { /* not yet handled by codegen */ }


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


// TODO stbox_shift_scale_time: unsupported param const int *
// func STBOXShiftScaleTime(...) { /* not yet handled by codegen */ }


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


// TspatialOut wraps MEOS C function tspatial_out.
func TspatialOut(temp Temporal, maxdd int) string {
	res := C.tspatial_out(temp.Inner(), C.int(maxdd))
	return C.GoString(res)
}


// TODO tgeo_from_base_temp: unsupported param const int *
// func TgeoFromBaseTemp(...) { /* not yet handled by codegen */ }


// TODO tgeoinst_make: unsupported param const int *
// func TgeoinstMake(...) { /* not yet handled by codegen */ }


// TODO tgeoseq_from_base_tstzset: unsupported param const int *
// func TgeoseqFromBaseTstzset(...) { /* not yet handled by codegen */ }


// TODO tgeoseq_from_base_tstzspan: unsupported param const int *
// func TgeoseqFromBaseTstzspan(...) { /* not yet handled by codegen */ }


// TODO tgeoseqset_from_base_tstzspanset: unsupported param const int *
// func TgeoseqsetFromBaseTstzspanset(...) { /* not yet handled by codegen */ }


// TODO tpoint_from_base_temp: unsupported param const int *
// func TpointFromBaseTemp(...) { /* not yet handled by codegen */ }


// TODO tpointinst_make: unsupported param const int *
// func TpointinstMake(...) { /* not yet handled by codegen */ }


// TODO tpointseq_from_base_tstzset: unsupported param const int *
// func TpointseqFromBaseTstzset(...) { /* not yet handled by codegen */ }


// TODO tpointseq_from_base_tstzspan: unsupported param const int *
// func TpointseqFromBaseTstzspan(...) { /* not yet handled by codegen */ }


// TpointseqMakeCoords wraps MEOS C function tpointseq_make_coords.
func TpointseqMakeCoords(xcoords []float64, ycoords []float64, zcoords []float64, times []int, srid int, geodetic bool, lower_inc bool, upper_inc bool, interp Interpolation, normalize bool) TSequence {
	_c_xcoords := make([]C.double, len(xcoords))
	for _i, _v := range xcoords { _c_xcoords[_i] = C.double(_v) }
	_c_ycoords := make([]C.double, len(ycoords))
	for _i, _v := range ycoords { _c_ycoords[_i] = C.double(_v) }
	var _c_zcoords []C.double
	if zcoords != nil {
		_c_zcoords = make([]C.double, len(zcoords))
		for _i, _v := range zcoords { _c_zcoords[_i] = C.double(_v) }
	}
	var _c_times []C.int
	if times != nil {
		_c_times = make([]C.int, len(times))
		for _i, _v := range times { _c_times[_i] = C.int(_v) }
	}
	res := C.tpointseq_make_coords(&_c_xcoords[0], &_c_ycoords[0], _ptr_or_nil_double(_c_zcoords), _ptr_or_nil_int(_c_times), C.int(len(xcoords)), C.int(srid), C.bool(geodetic), C.bool(lower_inc), C.bool(upper_inc), C.interpType(interp), C.bool(normalize))
	return TSequence{_inner: res}
}


// TODO tpointseqset_from_base_tstzspanset: unsupported param const int *
// func TpointseqsetFromBaseTstzspanset(...) { /* not yet handled by codegen */ }


// TODO box3d_to_stbox: unsupported param const int *
// func Box3dToSTBOX(...) { /* not yet handled by codegen */ }


// TODO gbox_to_stbox: unsupported param const int *
// func GboxToSTBOX(...) { /* not yet handled by codegen */ }


// TODO geomeas_to_tpoint: unsupported param const int *
// func GeomeasToTpoint(...) { /* not yet handled by codegen */ }


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


// TpointAsMvtgeom wraps MEOS C function tpoint_as_mvtgeom.
func TpointAsMvtgeom(temp Temporal, bounds *STBox, extent int32, buffer int32, clip_geom bool) (bool, []int, []int, int) {
	var _out_gsarr *C.int
	var _out_timesarr *C.int
	var _out_count C.int
	res := C.tpoint_as_mvtgeom(temp.Inner(), bounds._inner, C.int32_t(extent), C.int32_t(buffer), C.bool(clip_geom), &_out_gsarr, &_out_timesarr, &_out_count)
	_slice__out_gsarr := unsafe.Slice(_out_gsarr, None)
	_out_gsarr_go := make([]int, None)
	for _i, _e := range _slice__out_gsarr { _out_gsarr_go[_i] = int(_e) }
	_slice__out_timesarr := unsafe.Slice(_out_timesarr, None)
	_out_timesarr_go := make([]int, None)
	for _i, _e := range _slice__out_timesarr { _out_timesarr_go[_i] = int(_e) }
	return bool(res), _out_gsarr_go, _out_timesarr_go, int(_out_count)
}


// TODO tpoint_tfloat_to_geomeas: unhandled OUTPUT_SCALAR shape int **
// func TpointTfloatToGeomeas(...) { /* not yet handled by codegen */ }


// TspatialToSTBOX wraps MEOS C function tspatial_to_stbox.
func TspatialToSTBOX(temp Temporal) *STBox {
	res := C.tspatial_to_stbox(temp.Inner())
	return &STBox{_inner: res}
}


// TODO bearing_point_point: unsupported param const int *
// func BearingPointPoint(...) { /* not yet handled by codegen */ }


// TODO bearing_tpoint_point: unsupported param const int *
// func BearingTpointPoint(...) { /* not yet handled by codegen */ }


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


// TODO tgeo_convex_hull: unsupported return type int *
// func TgeoConvexHull(...) { /* not yet handled by codegen */ }


// TODO tgeo_end_value: unsupported return type int *
// func TgeoEndValue(...) { /* not yet handled by codegen */ }


// TODO tgeo_start_value: unsupported return type int *
// func TgeoStartValue(...) { /* not yet handled by codegen */ }


// TODO tgeo_traversed_area: unsupported return type int *
// func TgeoTraversedArea(...) { /* not yet handled by codegen */ }


// TODO tgeo_value_at_timestamptz: unhandled OUTPUT_SCALAR shape int **
// func TgeoValueAtTimestamptz(...) { /* not yet handled by codegen */ }


// TODO tgeo_value_n: unhandled OUTPUT_SCALAR shape int **
// func TgeoValueN(...) { /* not yet handled by codegen */ }


// TODO tgeo_values: unsupported return type int **
// func TgeoValues(...) { /* not yet handled by codegen */ }


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


// TODO tpoint_trajectory: unsupported return type int *
// func TpointTrajectory(...) { /* not yet handled by codegen */ }


// TODO tpoint_twcentroid: unsupported return type int *
// func TpointTwcentroid(...) { /* not yet handled by codegen */ }


// TODO tgeo_affine: unsupported param const int *
// func TgeoAffine(...) { /* not yet handled by codegen */ }


// TODO tgeo_scale: unsupported param const int *
// func TgeoScale(...) { /* not yet handled by codegen */ }


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


// TODO tgeo_at_geom: unsupported param const int *
// func TgeoAtGeom(...) { /* not yet handled by codegen */ }


// TgeoAtSTBOX wraps MEOS C function tgeo_at_stbox.
func TgeoAtSTBOX(temp Temporal, box *STBox, border_inc bool) Temporal {
	res := C.tgeo_at_stbox(temp.Inner(), box._inner, C.bool(border_inc))
	return CreateTemporal(res)
}


// TODO tgeo_at_value: unsupported param int *
// func TgeoAtValue(...) { /* not yet handled by codegen */ }


// TODO tgeo_minus_geom: unsupported param const int *
// func TgeoMinusGeom(...) { /* not yet handled by codegen */ }


// TgeoMinusSTBOX wraps MEOS C function tgeo_minus_stbox.
func TgeoMinusSTBOX(temp Temporal, box *STBox, border_inc bool) Temporal {
	res := C.tgeo_minus_stbox(temp.Inner(), box._inner, C.bool(border_inc))
	return CreateTemporal(res)
}


// TODO tgeo_minus_value: unsupported param int *
// func TgeoMinusValue(...) { /* not yet handled by codegen */ }


// TpointAtElevation wraps MEOS C function tpoint_at_elevation.
func TpointAtElevation(temp Temporal, s *Span) Temporal {
	res := C.tpoint_at_elevation(temp.Inner(), s._inner)
	return CreateTemporal(res)
}


// TODO tpoint_at_geom: unsupported param const int *
// func TpointAtGeom(...) { /* not yet handled by codegen */ }


// TODO tpoint_at_value: unsupported param int *
// func TpointAtValue(...) { /* not yet handled by codegen */ }


// TpointMinusElevation wraps MEOS C function tpoint_minus_elevation.
func TpointMinusElevation(temp Temporal, s *Span) Temporal {
	res := C.tpoint_minus_elevation(temp.Inner(), s._inner)
	return CreateTemporal(res)
}


// TODO tpoint_minus_geom: unsupported param const int *
// func TpointMinusGeom(...) { /* not yet handled by codegen */ }


// TODO tpoint_minus_value: unsupported param int *
// func TpointMinusValue(...) { /* not yet handled by codegen */ }


// TODO always_eq_geo_tgeo: unsupported param const int *
// func AlwaysEqGeoTgeo(...) { /* not yet handled by codegen */ }


// TODO always_eq_tgeo_geo: unsupported param const int *
// func AlwaysEqTgeoGeo(...) { /* not yet handled by codegen */ }


// AlwaysEqTgeoTgeo wraps MEOS C function always_eq_tgeo_tgeo.
func AlwaysEqTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.always_eq_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO always_ne_geo_tgeo: unsupported param const int *
// func AlwaysNeGeoTgeo(...) { /* not yet handled by codegen */ }


// TODO always_ne_tgeo_geo: unsupported param const int *
// func AlwaysNeTgeoGeo(...) { /* not yet handled by codegen */ }


// AlwaysNeTgeoTgeo wraps MEOS C function always_ne_tgeo_tgeo.
func AlwaysNeTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.always_ne_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO ever_eq_geo_tgeo: unsupported param const int *
// func EverEqGeoTgeo(...) { /* not yet handled by codegen */ }


// TODO ever_eq_tgeo_geo: unsupported param const int *
// func EverEqTgeoGeo(...) { /* not yet handled by codegen */ }


// EverEqTgeoTgeo wraps MEOS C function ever_eq_tgeo_tgeo.
func EverEqTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_eq_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO ever_ne_geo_tgeo: unsupported param const int *
// func EverNeGeoTgeo(...) { /* not yet handled by codegen */ }


// TODO ever_ne_tgeo_geo: unsupported param const int *
// func EverNeTgeoGeo(...) { /* not yet handled by codegen */ }


// EverNeTgeoTgeo wraps MEOS C function ever_ne_tgeo_tgeo.
func EverNeTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_ne_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO teq_geo_tgeo: unsupported param const int *
// func TeqGeoTgeo(...) { /* not yet handled by codegen */ }


// TODO teq_tgeo_geo: unsupported param const int *
// func TeqTgeoGeo(...) { /* not yet handled by codegen */ }


// TODO tne_geo_tgeo: unsupported param const int *
// func TneGeoTgeo(...) { /* not yet handled by codegen */ }


// TODO tne_tgeo_geo: unsupported param const int *
// func TneTgeoGeo(...) { /* not yet handled by codegen */ }


// TgeoStboxes wraps MEOS C function tgeo_stboxes.
func TgeoStboxes(temp Temporal) (*STBox, int) {
	var _out_count C.int
	res := C.tgeo_stboxes(temp.Inner(), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// TODO tgeo_space_boxes: unsupported param const int *
// func TgeoSpaceBoxes(...) { /* not yet handled by codegen */ }


// TODO tgeo_space_time_boxes: unsupported param const int *
// func TgeoSpaceTimeBoxes(...) { /* not yet handled by codegen */ }


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


// TODO acontains_geo_tgeo: unsupported param const int *
// func AcontainsGeoTgeo(...) { /* not yet handled by codegen */ }


// TODO acontains_tgeo_geo: unsupported param const int *
// func AcontainsTgeoGeo(...) { /* not yet handled by codegen */ }


// AcontainsTgeoTgeo wraps MEOS C function acontains_tgeo_tgeo.
func AcontainsTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.acontains_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO adisjoint_tgeo_geo: unsupported param const int *
// func AdisjointTgeoGeo(...) { /* not yet handled by codegen */ }


// AdisjointTgeoTgeo wraps MEOS C function adisjoint_tgeo_tgeo.
func AdisjointTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.adisjoint_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO adwithin_tgeo_geo: unsupported param const int *
// func AdwithinTgeoGeo(...) { /* not yet handled by codegen */ }


// AdwithinTgeoTgeo wraps MEOS C function adwithin_tgeo_tgeo.
func AdwithinTgeoTgeo(temp1 Temporal, temp2 Temporal, dist float64) int {
	res := C.adwithin_tgeo_tgeo(temp1.Inner(), temp2.Inner(), C.double(dist))
	return int(res)
}


// TODO aintersects_tgeo_geo: unsupported param const int *
// func AintersectsTgeoGeo(...) { /* not yet handled by codegen */ }


// AintersectsTgeoTgeo wraps MEOS C function aintersects_tgeo_tgeo.
func AintersectsTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.aintersects_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO atouches_tgeo_geo: unsupported param const int *
// func AtouchesTgeoGeo(...) { /* not yet handled by codegen */ }


// AtouchesTgeoTgeo wraps MEOS C function atouches_tgeo_tgeo.
func AtouchesTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.atouches_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO atouches_tpoint_geo: unsupported param const int *
// func AtouchesTpointGeo(...) { /* not yet handled by codegen */ }


// TODO econtains_geo_tgeo: unsupported param const int *
// func EcontainsGeoTgeo(...) { /* not yet handled by codegen */ }


// TODO econtains_tgeo_geo: unsupported param const int *
// func EcontainsTgeoGeo(...) { /* not yet handled by codegen */ }


// EcontainsTgeoTgeo wraps MEOS C function econtains_tgeo_tgeo.
func EcontainsTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.econtains_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO ecovers_geo_tgeo: unsupported param const int *
// func EcoversGeoTgeo(...) { /* not yet handled by codegen */ }


// TODO ecovers_tgeo_geo: unsupported param const int *
// func EcoversTgeoGeo(...) { /* not yet handled by codegen */ }


// EcoversTgeoTgeo wraps MEOS C function ecovers_tgeo_tgeo.
func EcoversTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.ecovers_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO edisjoint_tgeo_geo: unsupported param const int *
// func EdisjointTgeoGeo(...) { /* not yet handled by codegen */ }


// EdisjointTgeoTgeo wraps MEOS C function edisjoint_tgeo_tgeo.
func EdisjointTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.edisjoint_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO edwithin_tgeo_geo: unsupported param const int *
// func EdwithinTgeoGeo(...) { /* not yet handled by codegen */ }


// EdwithinTgeoTgeo wraps MEOS C function edwithin_tgeo_tgeo.
func EdwithinTgeoTgeo(temp1 Temporal, temp2 Temporal, dist float64) int {
	res := C.edwithin_tgeo_tgeo(temp1.Inner(), temp2.Inner(), C.double(dist))
	return int(res)
}


// TODO eintersects_tgeo_geo: unsupported param const int *
// func EintersectsTgeoGeo(...) { /* not yet handled by codegen */ }


// EintersectsTgeoTgeo wraps MEOS C function eintersects_tgeo_tgeo.
func EintersectsTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.eintersects_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO etouches_tgeo_geo: unsupported param const int *
// func EtouchesTgeoGeo(...) { /* not yet handled by codegen */ }


// EtouchesTgeoTgeo wraps MEOS C function etouches_tgeo_tgeo.
func EtouchesTgeoTgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.etouches_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO etouches_tpoint_geo: unsupported param const int *
// func EtouchesTpointGeo(...) { /* not yet handled by codegen */ }


// TODO tcontains_geo_tgeo: unsupported param const int *
// func TcontainsGeoTgeo(...) { /* not yet handled by codegen */ }


// TODO tcontains_tgeo_geo: unsupported param const int *
// func TcontainsTgeoGeo(...) { /* not yet handled by codegen */ }


// TcontainsTgeoTgeo wraps MEOS C function tcontains_tgeo_tgeo.
func TcontainsTgeoTgeo(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tcontains_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TODO tcovers_geo_tgeo: unsupported param const int *
// func TcoversGeoTgeo(...) { /* not yet handled by codegen */ }


// TODO tcovers_tgeo_geo: unsupported param const int *
// func TcoversTgeoGeo(...) { /* not yet handled by codegen */ }


// TcoversTgeoTgeo wraps MEOS C function tcovers_tgeo_tgeo.
func TcoversTgeoTgeo(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tcovers_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TODO tdisjoint_geo_tgeo: unsupported param const int *
// func TdisjointGeoTgeo(...) { /* not yet handled by codegen */ }


// TODO tdisjoint_tgeo_geo: unsupported param const int *
// func TdisjointTgeoGeo(...) { /* not yet handled by codegen */ }


// TdisjointTgeoTgeo wraps MEOS C function tdisjoint_tgeo_tgeo.
func TdisjointTgeoTgeo(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tdisjoint_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TODO tdwithin_geo_tgeo: unsupported param const int *
// func TdwithinGeoTgeo(...) { /* not yet handled by codegen */ }


// TODO tdwithin_tgeo_geo: unsupported param const int *
// func TdwithinTgeoGeo(...) { /* not yet handled by codegen */ }


// TdwithinTgeoTgeo wraps MEOS C function tdwithin_tgeo_tgeo.
func TdwithinTgeoTgeo(temp1 Temporal, temp2 Temporal, dist float64) Temporal {
	res := C.tdwithin_tgeo_tgeo(temp1.Inner(), temp2.Inner(), C.double(dist))
	return CreateTemporal(res)
}


// TODO tintersects_geo_tgeo: unsupported param const int *
// func TintersectsGeoTgeo(...) { /* not yet handled by codegen */ }


// TODO tintersects_tgeo_geo: unsupported param const int *
// func TintersectsTgeoGeo(...) { /* not yet handled by codegen */ }


// TintersectsTgeoTgeo wraps MEOS C function tintersects_tgeo_tgeo.
func TintersectsTgeoTgeo(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tintersects_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TODO ttouches_geo_tgeo: unsupported param const int *
// func TtouchesGeoTgeo(...) { /* not yet handled by codegen */ }


// TODO ttouches_tgeo_geo: unsupported param const int *
// func TtouchesTgeoGeo(...) { /* not yet handled by codegen */ }


// TtouchesTgeoTgeo wraps MEOS C function ttouches_tgeo_tgeo.
func TtouchesTgeoTgeo(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.ttouches_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TODO tdistance_tgeo_geo: unsupported param const int *
// func TdistanceTgeoGeo(...) { /* not yet handled by codegen */ }


// TdistanceTgeoTgeo wraps MEOS C function tdistance_tgeo_tgeo.
func TdistanceTgeoTgeo(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tdistance_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TODO nad_stbox_geo: unsupported param const int *
// func NadSTBOXGeo(...) { /* not yet handled by codegen */ }


// NadSTBOXSTBOX wraps MEOS C function nad_stbox_stbox.
func NadSTBOXSTBOX(box1 *STBox, box2 *STBox) float64 {
	res := C.nad_stbox_stbox(box1._inner, box2._inner)
	return float64(res)
}


// TODO nad_tgeo_geo: unsupported param const int *
// func NadTgeoGeo(...) { /* not yet handled by codegen */ }


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


// TODO nai_tgeo_geo: unsupported param const int *
// func NaiTgeoGeo(...) { /* not yet handled by codegen */ }


// NaiTgeoTgeo wraps MEOS C function nai_tgeo_tgeo.
func NaiTgeoTgeo(temp1 Temporal, temp2 Temporal) TInstant {
	res := C.nai_tgeo_tgeo(temp1.Inner(), temp2.Inner())
	return TInstant{_inner: res}
}


// TODO shortestline_tgeo_geo: unsupported return type int *
// func ShortestlineTgeoGeo(...) { /* not yet handled by codegen */ }


// TODO shortestline_tgeo_tgeo: unsupported return type int *
// func ShortestlineTgeoTgeo(...) { /* not yet handled by codegen */ }


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


// TODO stbox_get_space_tile: unsupported param const int *
// func STBOXGetSpaceTile(...) { /* not yet handled by codegen */ }


// TODO stbox_get_space_time_tile: unsupported param const int *
// func STBOXGetSpaceTimeTile(...) { /* not yet handled by codegen */ }


// TODO stbox_get_time_tile: unsupported param const int *
// func STBOXGetTimeTile(...) { /* not yet handled by codegen */ }


// TODO stbox_space_tiles: unsupported param const int *
// func STBOXSpaceTiles(...) { /* not yet handled by codegen */ }


// TODO stbox_space_time_tiles: unsupported param const int *
// func STBOXSpaceTimeTiles(...) { /* not yet handled by codegen */ }


// TODO stbox_time_tiles: unsupported param const int *
// func STBOXTimeTiles(...) { /* not yet handled by codegen */ }


// TODO tgeo_space_split: unsupported param const int *
// func TgeoSpaceSplit(...) { /* not yet handled by codegen */ }


// TODO tgeo_space_time_split: unsupported param const int *
// func TgeoSpaceTimeSplit(...) { /* not yet handled by codegen */ }


// TODO geo_cluster_kmeans: unsupported return type int *
// func GeoClusterKmeans(...) { /* not yet handled by codegen */ }


// TODO geo_cluster_dbscan: unsupported param const int **
// func GeoClusterDbscan(...) { /* not yet handled by codegen */ }


// TODO geo_cluster_intersecting: unsupported return type int **
// func GeoClusterIntersecting(...) { /* not yet handled by codegen */ }


// TODO geo_cluster_within: unsupported return type int **
// func GeoClusterWithin(...) { /* not yet handled by codegen */ }

