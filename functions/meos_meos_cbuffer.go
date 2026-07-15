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

// CbufferAsEWKT wraps MEOS C function cbuffer_as_ewkt.
func CbufferAsEWKT(cb *Cbuffer, maxdd int) string {
	_cret := C.cbuffer_as_ewkt(cb._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// CbufferAsHexwkb wraps MEOS C function cbuffer_as_hexwkb.
func CbufferAsHexwkb(cb *Cbuffer, variant uint8, size_out unsafe.Pointer) string {
	_cret := C.cbuffer_as_hexwkb(cb._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return C.GoString(_cret)
}


// CbufferAsText wraps MEOS C function cbuffer_as_text.
func CbufferAsText(cb *Cbuffer, maxdd int) string {
	_cret := C.cbuffer_as_text(cb._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// CbufferAsWKB wraps MEOS C function cbuffer_as_wkb.
func CbufferAsWKB(cb *Cbuffer, variant uint8, size_out unsafe.Pointer) unsafe.Pointer {
	_cret := C.cbuffer_as_wkb(cb._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return unsafe.Pointer(_cret)
}


// CbufferFromHexwkb wraps MEOS C function cbuffer_from_hexwkb.
func CbufferFromHexwkb(hexwkb string) *Cbuffer {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	_cret := C.cbuffer_from_hexwkb(_c_hexwkb)
	return &Cbuffer{_inner: _cret}
}


// CbufferFromWKB wraps MEOS C function cbuffer_from_wkb.
func CbufferFromWKB(wkb unsafe.Pointer, size uint) *Cbuffer {
	_cret := C.cbuffer_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	return &Cbuffer{_inner: _cret}
}


// CbufferIn wraps MEOS C function cbuffer_in.
func CbufferIn(str string) *Cbuffer {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.cbuffer_in(_c_str)
	return &Cbuffer{_inner: _cret}
}


// CbufferOut wraps MEOS C function cbuffer_out.
func CbufferOut(cb *Cbuffer, maxdd int) string {
	_cret := C.cbuffer_out(cb._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// CbufferCopy wraps MEOS C function cbuffer_copy.
func CbufferCopy(cb *Cbuffer) *Cbuffer {
	_cret := C.cbuffer_copy(cb._inner)
	return &Cbuffer{_inner: _cret}
}


// CbufferMake wraps MEOS C function cbuffer_make.
func CbufferMake(point *Geom, radius float64) *Cbuffer {
	_cret := C.cbuffer_make(point._inner, C.double(radius))
	return &Cbuffer{_inner: _cret}
}


// CbufferToGeom wraps MEOS C function cbuffer_to_geom.
func CbufferToGeom(cb *Cbuffer) *Geom {
	_cret := C.cbuffer_to_geom(cb._inner)
	return &Geom{_inner: _cret}
}


// CbufferToSTBOX wraps MEOS C function cbuffer_to_stbox.
func CbufferToSTBOX(cb *Cbuffer) *STBox {
	_cret := C.cbuffer_to_stbox(cb._inner)
	return &STBox{_inner: _cret}
}


// CbufferarrToGeom wraps MEOS C function cbufferarr_to_geom.
func CbufferarrToGeom(cbarr unsafe.Pointer, count int) *Geom {
	_cret := C.cbufferarr_to_geom((**C.Cbuffer)(unsafe.Pointer(cbarr)), C.int(count))
	return &Geom{_inner: _cret}
}


// GeomToCbuffer wraps MEOS C function geom_to_cbuffer.
func GeomToCbuffer(gs *Geom) *Cbuffer {
	_cret := C.geom_to_cbuffer(gs._inner)
	return &Cbuffer{_inner: _cret}
}


// CbufferHash wraps MEOS C function cbuffer_hash.
func CbufferHash(cb *Cbuffer) uint32 {
	_cret := C.cbuffer_hash(cb._inner)
	return uint32(_cret)
}


// CbufferHashExtended wraps MEOS C function cbuffer_hash_extended.
func CbufferHashExtended(cb *Cbuffer, seed uint64) uint64 {
	_cret := C.cbuffer_hash_extended(cb._inner, C.uint64_t(seed))
	return uint64(_cret)
}


// CbufferPoint wraps MEOS C function cbuffer_point.
func CbufferPoint(cb *Cbuffer) *Geom {
	_cret := C.cbuffer_point(cb._inner)
	return &Geom{_inner: _cret}
}


// CbufferRadius wraps MEOS C function cbuffer_radius.
func CbufferRadius(cb *Cbuffer) float64 {
	_cret := C.cbuffer_radius(cb._inner)
	return float64(_cret)
}


// CbufferRound wraps MEOS C function cbuffer_round.
func CbufferRound(cb *Cbuffer, maxdd int) *Cbuffer {
	_cret := C.cbuffer_round(cb._inner, C.int(maxdd))
	return &Cbuffer{_inner: _cret}
}


// CbufferarrRound wraps MEOS C function cbufferarr_round.
func CbufferarrRound(cbarr unsafe.Pointer, count int, maxdd int) unsafe.Pointer {
	_cret := C.cbufferarr_round((**C.Cbuffer)(unsafe.Pointer(cbarr)), C.int(count), C.int(maxdd))
	return unsafe.Pointer(_cret)
}


// CbufferSetSRID wraps MEOS C function cbuffer_set_srid.
func CbufferSetSRID(cb *Cbuffer, srid int32) {
	C.cbuffer_set_srid(cb._inner, C.int32_t(srid))
}


// CbufferSRID wraps MEOS C function cbuffer_srid.
func CbufferSRID(cb *Cbuffer) int32 {
	_cret := C.cbuffer_srid(cb._inner)
	return int32(_cret)
}


// CbufferTransform wraps MEOS C function cbuffer_transform.
func CbufferTransform(cb *Cbuffer, srid int32) *Cbuffer {
	_cret := C.cbuffer_transform(cb._inner, C.int32_t(srid))
	return &Cbuffer{_inner: _cret}
}


// CbufferTransformPipeline wraps MEOS C function cbuffer_transform_pipeline.
func CbufferTransformPipeline(cb *Cbuffer, pipelinestr string, srid int32, is_forward bool) *Cbuffer {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	_cret := C.cbuffer_transform_pipeline(cb._inner, _c_pipelinestr, C.int32_t(srid), C.bool(is_forward))
	return &Cbuffer{_inner: _cret}
}


// ContainsCbufferCbuffer wraps MEOS C function contains_cbuffer_cbuffer.
func ContainsCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer) int {
	_cret := C.contains_cbuffer_cbuffer(cb1._inner, cb2._inner)
	return int(_cret)
}


// CoversCbufferCbuffer wraps MEOS C function covers_cbuffer_cbuffer.
func CoversCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer) int {
	_cret := C.covers_cbuffer_cbuffer(cb1._inner, cb2._inner)
	return int(_cret)
}


// DisjointCbufferCbuffer wraps MEOS C function disjoint_cbuffer_cbuffer.
func DisjointCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer) int {
	_cret := C.disjoint_cbuffer_cbuffer(cb1._inner, cb2._inner)
	return int(_cret)
}


// DwithinCbufferCbuffer wraps MEOS C function dwithin_cbuffer_cbuffer.
func DwithinCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer, dist float64) int {
	_cret := C.dwithin_cbuffer_cbuffer(cb1._inner, cb2._inner, C.double(dist))
	return int(_cret)
}


// IntersectsCbufferCbuffer wraps MEOS C function intersects_cbuffer_cbuffer.
func IntersectsCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer) int {
	_cret := C.intersects_cbuffer_cbuffer(cb1._inner, cb2._inner)
	return int(_cret)
}


// TouchesCbufferCbuffer wraps MEOS C function touches_cbuffer_cbuffer.
func TouchesCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer) int {
	_cret := C.touches_cbuffer_cbuffer(cb1._inner, cb2._inner)
	return int(_cret)
}


// CbufferTstzspanToSTBOX wraps MEOS C function cbuffer_tstzspan_to_stbox.
func CbufferTstzspanToSTBOX(cb *Cbuffer, s *Span) *STBox {
	_cret := C.cbuffer_tstzspan_to_stbox(cb._inner, s._inner)
	return &STBox{_inner: _cret}
}


// CbufferTimestamptzToSTBOX wraps MEOS C function cbuffer_timestamptz_to_stbox.
func CbufferTimestamptzToSTBOX(cb *Cbuffer, t int64) *STBox {
	_cret := C.cbuffer_timestamptz_to_stbox(cb._inner, C.TimestampTz(t))
	return &STBox{_inner: _cret}
}


// DistanceCbufferCbuffer wraps MEOS C function distance_cbuffer_cbuffer.
func DistanceCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer) float64 {
	_cret := C.distance_cbuffer_cbuffer(cb1._inner, cb2._inner)
	return float64(_cret)
}


// DistanceCbufferGeo wraps MEOS C function distance_cbuffer_geo.
func DistanceCbufferGeo(cb *Cbuffer, gs *Geom) float64 {
	_cret := C.distance_cbuffer_geo(cb._inner, gs._inner)
	return float64(_cret)
}


// DistanceCbufferSTBOX wraps MEOS C function distance_cbuffer_stbox.
func DistanceCbufferSTBOX(cb *Cbuffer, box *STBox) float64 {
	_cret := C.distance_cbuffer_stbox(cb._inner, box._inner)
	return float64(_cret)
}


// NadCbufferSTBOX wraps MEOS C function nad_cbuffer_stbox.
func NadCbufferSTBOX(cb *Cbuffer, box *STBox) float64 {
	_cret := C.nad_cbuffer_stbox(cb._inner, box._inner)
	return float64(_cret)
}


// CbufferCmp wraps MEOS C function cbuffer_cmp.
func CbufferCmp(cb1 *Cbuffer, cb2 *Cbuffer) int {
	_cret := C.cbuffer_cmp(cb1._inner, cb2._inner)
	return int(_cret)
}


// CbufferEq wraps MEOS C function cbuffer_eq.
func CbufferEq(cb1 *Cbuffer, cb2 *Cbuffer) bool {
	_cret := C.cbuffer_eq(cb1._inner, cb2._inner)
	return bool(_cret)
}


// CbufferGe wraps MEOS C function cbuffer_ge.
func CbufferGe(cb1 *Cbuffer, cb2 *Cbuffer) bool {
	_cret := C.cbuffer_ge(cb1._inner, cb2._inner)
	return bool(_cret)
}


// CbufferGt wraps MEOS C function cbuffer_gt.
func CbufferGt(cb1 *Cbuffer, cb2 *Cbuffer) bool {
	_cret := C.cbuffer_gt(cb1._inner, cb2._inner)
	return bool(_cret)
}


// CbufferLe wraps MEOS C function cbuffer_le.
func CbufferLe(cb1 *Cbuffer, cb2 *Cbuffer) bool {
	_cret := C.cbuffer_le(cb1._inner, cb2._inner)
	return bool(_cret)
}


// CbufferLt wraps MEOS C function cbuffer_lt.
func CbufferLt(cb1 *Cbuffer, cb2 *Cbuffer) bool {
	_cret := C.cbuffer_lt(cb1._inner, cb2._inner)
	return bool(_cret)
}


// CbufferNe wraps MEOS C function cbuffer_ne.
func CbufferNe(cb1 *Cbuffer, cb2 *Cbuffer) bool {
	_cret := C.cbuffer_ne(cb1._inner, cb2._inner)
	return bool(_cret)
}


// CbufferNsame wraps MEOS C function cbuffer_nsame.
func CbufferNsame(cb1 *Cbuffer, cb2 *Cbuffer) bool {
	_cret := C.cbuffer_nsame(cb1._inner, cb2._inner)
	return bool(_cret)
}


// CbufferSame wraps MEOS C function cbuffer_same.
func CbufferSame(cb1 *Cbuffer, cb2 *Cbuffer) bool {
	_cret := C.cbuffer_same(cb1._inner, cb2._inner)
	return bool(_cret)
}


// CbuffersetIn wraps MEOS C function cbufferset_in.
func CbuffersetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.cbufferset_in(_c_str)
	return &Set{_inner: _cret}
}


// CbuffersetOut wraps MEOS C function cbufferset_out.
func CbuffersetOut(s *Set, maxdd int) string {
	_cret := C.cbufferset_out(s._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// CbuffersetMake wraps MEOS C function cbufferset_make.
func CbuffersetMake(values unsafe.Pointer, count int) *Set {
	_cret := C.cbufferset_make((**C.Cbuffer)(unsafe.Pointer(values)), C.int(count))
	return &Set{_inner: _cret}
}


// CbufferToSet wraps MEOS C function cbuffer_to_set.
func CbufferToSet(cb *Cbuffer) *Set {
	_cret := C.cbuffer_to_set(cb._inner)
	return &Set{_inner: _cret}
}


// CbuffersetEndValue wraps MEOS C function cbufferset_end_value.
func CbuffersetEndValue(s *Set) *Cbuffer {
	_cret := C.cbufferset_end_value(s._inner)
	return &Cbuffer{_inner: _cret}
}


// CbuffersetStartValue wraps MEOS C function cbufferset_start_value.
func CbuffersetStartValue(s *Set) *Cbuffer {
	_cret := C.cbufferset_start_value(s._inner)
	return &Cbuffer{_inner: _cret}
}


// CbuffersetValueN wraps MEOS C function cbufferset_value_n.
func CbuffersetValueN(s *Set, n int) (bool, *Cbuffer) {
	var _out_result *C.Cbuffer
	_cret := C.cbufferset_value_n(s._inner, C.int(n), &_out_result)
	return bool(_cret), &Cbuffer{_inner: _out_result}
}


// CbuffersetValues wraps MEOS C function cbufferset_values.
func CbuffersetValues(s *Set, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.cbufferset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// CbufferUnionTransfn wraps MEOS C function cbuffer_union_transfn.
func CbufferUnionTransfn(state *Set, cb *Cbuffer) *Set {
	_cret := C.cbuffer_union_transfn(state._inner, cb._inner)
	return &Set{_inner: _cret}
}


// ContainedCbufferSet wraps MEOS C function contained_cbuffer_set.
func ContainedCbufferSet(cb *Cbuffer, s *Set) bool {
	_cret := C.contained_cbuffer_set(cb._inner, s._inner)
	return bool(_cret)
}


// ContainsSetCbuffer wraps MEOS C function contains_set_cbuffer.
func ContainsSetCbuffer(s *Set, cb *Cbuffer) bool {
	_cret := C.contains_set_cbuffer(s._inner, cb._inner)
	return bool(_cret)
}


// IntersectionCbufferSet wraps MEOS C function intersection_cbuffer_set.
func IntersectionCbufferSet(cb *Cbuffer, s *Set) *Set {
	_cret := C.intersection_cbuffer_set(cb._inner, s._inner)
	return &Set{_inner: _cret}
}


// IntersectionSetCbuffer wraps MEOS C function intersection_set_cbuffer.
func IntersectionSetCbuffer(s *Set, cb *Cbuffer) *Set {
	_cret := C.intersection_set_cbuffer(s._inner, cb._inner)
	return &Set{_inner: _cret}
}


// MinusCbufferSet wraps MEOS C function minus_cbuffer_set.
func MinusCbufferSet(cb *Cbuffer, s *Set) *Set {
	_cret := C.minus_cbuffer_set(cb._inner, s._inner)
	return &Set{_inner: _cret}
}


// MinusSetCbuffer wraps MEOS C function minus_set_cbuffer.
func MinusSetCbuffer(s *Set, cb *Cbuffer) *Set {
	_cret := C.minus_set_cbuffer(s._inner, cb._inner)
	return &Set{_inner: _cret}
}


// UnionCbufferSet wraps MEOS C function union_cbuffer_set.
func UnionCbufferSet(cb *Cbuffer, s *Set) *Set {
	_cret := C.gunion_cbuffer_set(cb._inner, s._inner)
	return &Set{_inner: _cret}
}


// UnionSetCbuffer wraps MEOS C function union_set_cbuffer.
func UnionSetCbuffer(s *Set, cb *Cbuffer) *Set {
	_cret := C.gunion_set_cbuffer(s._inner, cb._inner)
	return &Set{_inner: _cret}
}


// TcbufferIn wraps MEOS C function tcbuffer_in.
func TcbufferIn(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tcbuffer_in(_c_str)
	return &Temporal{_inner: _cret}
}


// TcbufferFromMFJSON wraps MEOS C function tcbuffer_from_mfjson.
func TcbufferFromMFJSON(mfjson string) *Temporal {
	_c_mfjson := C.CString(mfjson)
	defer C.free(unsafe.Pointer(_c_mfjson))
	_cret := C.tcbuffer_from_mfjson(_c_mfjson)
	return &Temporal{_inner: _cret}
}


// TcbufferinstMake wraps MEOS C function tcbufferinst_make.
func TcbufferinstMake(cb *Cbuffer, t int64) *TInstant {
	_cret := C.tcbufferinst_make(cb._inner, C.TimestampTz(t))
	return &TInstant{_inner: _cret}
}


// TcbufferMake wraps MEOS C function tcbuffer_make.
func TcbufferMake(tpoint *Temporal, tfloat *Temporal) *Temporal {
	_cret := C.tcbuffer_make(tpoint._inner, tfloat._inner)
	return &Temporal{_inner: _cret}
}


// TcbufferFromBaseTemp wraps MEOS C function tcbuffer_from_base_temp.
func TcbufferFromBaseTemp(cb *Cbuffer, temp *Temporal) *Temporal {
	_cret := C.tcbuffer_from_base_temp(cb._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TcbufferseqFromBaseTstzset wraps MEOS C function tcbufferseq_from_base_tstzset.
func TcbufferseqFromBaseTstzset(cb *Cbuffer, s *Set) *TSequence {
	_cret := C.tcbufferseq_from_base_tstzset(cb._inner, s._inner)
	return &TSequence{_inner: _cret}
}


// TcbufferseqFromBaseTstzspan wraps MEOS C function tcbufferseq_from_base_tstzspan.
func TcbufferseqFromBaseTstzspan(cb *Cbuffer, s *Span, interp Interpolation) *TSequence {
	_cret := C.tcbufferseq_from_base_tstzspan(cb._inner, s._inner, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TcbufferseqsetFromBaseTstzspanset wraps MEOS C function tcbufferseqset_from_base_tstzspanset.
func TcbufferseqsetFromBaseTstzspanset(cb *Cbuffer, ss *SpanSet, interp Interpolation) *TSequenceSet {
	_cret := C.tcbufferseqset_from_base_tstzspanset(cb._inner, ss._inner, C.interpType(interp))
	return &TSequenceSet{_inner: _cret}
}


// TcbufferEndValue wraps MEOS C function tcbuffer_end_value.
func TcbufferEndValue(temp *Temporal) *Cbuffer {
	_cret := C.tcbuffer_end_value(temp._inner)
	return &Cbuffer{_inner: _cret}
}


// TcbufferPoints wraps MEOS C function tcbuffer_points.
func TcbufferPoints(temp *Temporal) *Set {
	_cret := C.tcbuffer_points(temp._inner)
	return &Set{_inner: _cret}
}


// TcbufferRadius wraps MEOS C function tcbuffer_radius.
func TcbufferRadius(temp *Temporal) *Set {
	_cret := C.tcbuffer_radius(temp._inner)
	return &Set{_inner: _cret}
}


// TcbufferTraversedArea wraps MEOS C function tcbuffer_traversed_area.
func TcbufferTraversedArea(temp *Temporal, unary_union bool) *Geom {
	_cret := C.tcbuffer_traversed_area(temp._inner, C.bool(unary_union))
	return &Geom{_inner: _cret}
}


// TcbufferConvexHull wraps MEOS C function tcbuffer_convex_hull.
func TcbufferConvexHull(temp *Temporal) *Geom {
	_cret := C.tcbuffer_convex_hull(temp._inner)
	return &Geom{_inner: _cret}
}


// TcbufferStartValue wraps MEOS C function tcbuffer_start_value.
func TcbufferStartValue(temp *Temporal) *Cbuffer {
	_cret := C.tcbuffer_start_value(temp._inner)
	return &Cbuffer{_inner: _cret}
}


// TcbufferValueAtTimestamptz wraps MEOS C function tcbuffer_value_at_timestamptz.
func TcbufferValueAtTimestamptz(temp *Temporal, t int64, strict bool) (bool, *Cbuffer) {
	var _out_value *C.Cbuffer
	_cret := C.tcbuffer_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_value)
	return bool(_cret), &Cbuffer{_inner: _out_value}
}


// TcbufferValueN wraps MEOS C function tcbuffer_value_n.
func TcbufferValueN(temp *Temporal, n int) (bool, *Cbuffer) {
	var _out_result *C.Cbuffer
	_cret := C.tcbuffer_value_n(temp._inner, C.int(n), &_out_result)
	return bool(_cret), &Cbuffer{_inner: _out_result}
}


// TcbufferValues wraps MEOS C function tcbuffer_values.
func TcbufferValues(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tcbuffer_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TcbufferToTfloat wraps MEOS C function tcbuffer_to_tfloat.
func TcbufferToTfloat(temp *Temporal) *Temporal {
	_cret := C.tcbuffer_to_tfloat(temp._inner)
	return &Temporal{_inner: _cret}
}


// TcbufferToTgeompoint wraps MEOS C function tcbuffer_to_tgeompoint.
func TcbufferToTgeompoint(temp *Temporal) *Temporal {
	_cret := C.tcbuffer_to_tgeompoint(temp._inner)
	return &Temporal{_inner: _cret}
}


// TgeometryToTcbuffer wraps MEOS C function tgeometry_to_tcbuffer.
func TgeometryToTcbuffer(temp *Temporal) *Temporal {
	_cret := C.tgeometry_to_tcbuffer(temp._inner)
	return &Temporal{_inner: _cret}
}


// TcbufferExpand wraps MEOS C function tcbuffer_expand.
func TcbufferExpand(temp *Temporal, dist float64) *Temporal {
	_cret := C.tcbuffer_expand(temp._inner, C.double(dist))
	return &Temporal{_inner: _cret}
}


// TcbufferAtCbuffer wraps MEOS C function tcbuffer_at_cbuffer.
func TcbufferAtCbuffer(temp *Temporal, cb *Cbuffer) *Temporal {
	_cret := C.tcbuffer_at_cbuffer(temp._inner, cb._inner)
	return &Temporal{_inner: _cret}
}


// TcbufferAtGeom wraps MEOS C function tcbuffer_at_geom.
func TcbufferAtGeom(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tcbuffer_at_geom(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TcbufferAtSTBOX wraps MEOS C function tcbuffer_at_stbox.
func TcbufferAtSTBOX(temp *Temporal, box *STBox, border_inc bool) *Temporal {
	_cret := C.tcbuffer_at_stbox(temp._inner, box._inner, C.bool(border_inc))
	return &Temporal{_inner: _cret}
}


// TcbufferMinusCbuffer wraps MEOS C function tcbuffer_minus_cbuffer.
func TcbufferMinusCbuffer(temp *Temporal, cb *Cbuffer) *Temporal {
	_cret := C.tcbuffer_minus_cbuffer(temp._inner, cb._inner)
	return &Temporal{_inner: _cret}
}


// TcbufferMinusGeom wraps MEOS C function tcbuffer_minus_geom.
func TcbufferMinusGeom(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tcbuffer_minus_geom(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TcbufferMinusSTBOX wraps MEOS C function tcbuffer_minus_stbox.
func TcbufferMinusSTBOX(temp *Temporal, box *STBox, border_inc bool) *Temporal {
	_cret := C.tcbuffer_minus_stbox(temp._inner, box._inner, C.bool(border_inc))
	return &Temporal{_inner: _cret}
}


// TdistanceTcbufferCbuffer wraps MEOS C function tdistance_tcbuffer_cbuffer.
func TdistanceTcbufferCbuffer(temp *Temporal, cb *Cbuffer) *Temporal {
	_cret := C.tdistance_tcbuffer_cbuffer(temp._inner, cb._inner)
	return &Temporal{_inner: _cret}
}


// TdistanceTcbufferGeo wraps MEOS C function tdistance_tcbuffer_geo.
func TdistanceTcbufferGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tdistance_tcbuffer_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TdistanceTcbufferTcbuffer wraps MEOS C function tdistance_tcbuffer_tcbuffer.
func TdistanceTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tdistance_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// NadTcbufferCbuffer wraps MEOS C function nad_tcbuffer_cbuffer.
func NadTcbufferCbuffer(temp *Temporal, cb *Cbuffer) float64 {
	_cret := C.nad_tcbuffer_cbuffer(temp._inner, cb._inner)
	return float64(_cret)
}


// NadTcbufferGeo wraps MEOS C function nad_tcbuffer_geo.
func NadTcbufferGeo(temp *Temporal, gs *Geom) float64 {
	_cret := C.nad_tcbuffer_geo(temp._inner, gs._inner)
	return float64(_cret)
}


// NadTcbufferSTBOX wraps MEOS C function nad_tcbuffer_stbox.
func NadTcbufferSTBOX(temp *Temporal, box *STBox) float64 {
	_cret := C.nad_tcbuffer_stbox(temp._inner, box._inner)
	return float64(_cret)
}


// NadTcbufferTcbuffer wraps MEOS C function nad_tcbuffer_tcbuffer.
func NadTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) float64 {
	_cret := C.nad_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return float64(_cret)
}


// MindistanceTcbufferTcbuffer wraps MEOS C function mindistance_tcbuffer_tcbuffer.
func MindistanceTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal, threshold float64) float64 {
	_cret := C.mindistance_tcbuffer_tcbuffer(temp1._inner, temp2._inner, C.double(threshold))
	return float64(_cret)
}


// NaiTcbufferCbuffer wraps MEOS C function nai_tcbuffer_cbuffer.
func NaiTcbufferCbuffer(temp *Temporal, cb *Cbuffer) *TInstant {
	_cret := C.nai_tcbuffer_cbuffer(temp._inner, cb._inner)
	return &TInstant{_inner: _cret}
}


// NaiTcbufferGeo wraps MEOS C function nai_tcbuffer_geo.
func NaiTcbufferGeo(temp *Temporal, gs *Geom) *TInstant {
	_cret := C.nai_tcbuffer_geo(temp._inner, gs._inner)
	return &TInstant{_inner: _cret}
}


// NaiTcbufferTcbuffer wraps MEOS C function nai_tcbuffer_tcbuffer.
func NaiTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) *TInstant {
	_cret := C.nai_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return &TInstant{_inner: _cret}
}


// ShortestlineTcbufferCbuffer wraps MEOS C function shortestline_tcbuffer_cbuffer.
func ShortestlineTcbufferCbuffer(temp *Temporal, cb *Cbuffer) *Geom {
	_cret := C.shortestline_tcbuffer_cbuffer(temp._inner, cb._inner)
	return &Geom{_inner: _cret}
}


// ShortestlineTcbufferGeo wraps MEOS C function shortestline_tcbuffer_geo.
func ShortestlineTcbufferGeo(temp *Temporal, gs *Geom) *Geom {
	_cret := C.shortestline_tcbuffer_geo(temp._inner, gs._inner)
	return &Geom{_inner: _cret}
}


// ShortestlineTcbufferTcbuffer wraps MEOS C function shortestline_tcbuffer_tcbuffer.
func ShortestlineTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) *Geom {
	_cret := C.shortestline_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return &Geom{_inner: _cret}
}


// AlwaysEqCbufferTcbuffer wraps MEOS C function always_eq_cbuffer_tcbuffer.
func AlwaysEqCbufferTcbuffer(cb *Cbuffer, temp *Temporal) int {
	_cret := C.always_eq_cbuffer_tcbuffer(cb._inner, temp._inner)
	return int(_cret)
}


// AlwaysEqTcbufferCbuffer wraps MEOS C function always_eq_tcbuffer_cbuffer.
func AlwaysEqTcbufferCbuffer(temp *Temporal, cb *Cbuffer) int {
	_cret := C.always_eq_tcbuffer_cbuffer(temp._inner, cb._inner)
	return int(_cret)
}


// AlwaysEqTcbufferTcbuffer wraps MEOS C function always_eq_tcbuffer_tcbuffer.
func AlwaysEqTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_eq_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return int(_cret)
}


// AlwaysNeCbufferTcbuffer wraps MEOS C function always_ne_cbuffer_tcbuffer.
func AlwaysNeCbufferTcbuffer(cb *Cbuffer, temp *Temporal) int {
	_cret := C.always_ne_cbuffer_tcbuffer(cb._inner, temp._inner)
	return int(_cret)
}


// AlwaysNeTcbufferCbuffer wraps MEOS C function always_ne_tcbuffer_cbuffer.
func AlwaysNeTcbufferCbuffer(temp *Temporal, cb *Cbuffer) int {
	_cret := C.always_ne_tcbuffer_cbuffer(temp._inner, cb._inner)
	return int(_cret)
}


// AlwaysNeTcbufferTcbuffer wraps MEOS C function always_ne_tcbuffer_tcbuffer.
func AlwaysNeTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_ne_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverEqCbufferTcbuffer wraps MEOS C function ever_eq_cbuffer_tcbuffer.
func EverEqCbufferTcbuffer(cb *Cbuffer, temp *Temporal) int {
	_cret := C.ever_eq_cbuffer_tcbuffer(cb._inner, temp._inner)
	return int(_cret)
}


// EverEqTcbufferCbuffer wraps MEOS C function ever_eq_tcbuffer_cbuffer.
func EverEqTcbufferCbuffer(temp *Temporal, cb *Cbuffer) int {
	_cret := C.ever_eq_tcbuffer_cbuffer(temp._inner, cb._inner)
	return int(_cret)
}


// EverEqTcbufferTcbuffer wraps MEOS C function ever_eq_tcbuffer_tcbuffer.
func EverEqTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_eq_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverNeCbufferTcbuffer wraps MEOS C function ever_ne_cbuffer_tcbuffer.
func EverNeCbufferTcbuffer(cb *Cbuffer, temp *Temporal) int {
	_cret := C.ever_ne_cbuffer_tcbuffer(cb._inner, temp._inner)
	return int(_cret)
}


// EverNeTcbufferCbuffer wraps MEOS C function ever_ne_tcbuffer_cbuffer.
func EverNeTcbufferCbuffer(temp *Temporal, cb *Cbuffer) int {
	_cret := C.ever_ne_tcbuffer_cbuffer(temp._inner, cb._inner)
	return int(_cret)
}


// EverNeTcbufferTcbuffer wraps MEOS C function ever_ne_tcbuffer_tcbuffer.
func EverNeTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_ne_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return int(_cret)
}


// TeqCbufferTcbuffer wraps MEOS C function teq_cbuffer_tcbuffer.
func TeqCbufferTcbuffer(cb *Cbuffer, temp *Temporal) *Temporal {
	_cret := C.teq_cbuffer_tcbuffer(cb._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TeqTcbufferCbuffer wraps MEOS C function teq_tcbuffer_cbuffer.
func TeqTcbufferCbuffer(temp *Temporal, cb *Cbuffer) *Temporal {
	_cret := C.teq_tcbuffer_cbuffer(temp._inner, cb._inner)
	return &Temporal{_inner: _cret}
}


// TneCbufferTcbuffer wraps MEOS C function tne_cbuffer_tcbuffer.
func TneCbufferTcbuffer(cb *Cbuffer, temp *Temporal) *Temporal {
	_cret := C.tne_cbuffer_tcbuffer(cb._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TneTcbufferCbuffer wraps MEOS C function tne_tcbuffer_cbuffer.
func TneTcbufferCbuffer(temp *Temporal, cb *Cbuffer) *Temporal {
	_cret := C.tne_tcbuffer_cbuffer(temp._inner, cb._inner)
	return &Temporal{_inner: _cret}
}


// AcontainsCbufferTcbuffer wraps MEOS C function acontains_cbuffer_tcbuffer.
func AcontainsCbufferTcbuffer(cb *Cbuffer, temp *Temporal) int {
	_cret := C.acontains_cbuffer_tcbuffer(cb._inner, temp._inner)
	return int(_cret)
}


// AcontainsGeoTcbuffer wraps MEOS C function acontains_geo_tcbuffer.
func AcontainsGeoTcbuffer(gs *Geom, temp *Temporal) int {
	_cret := C.acontains_geo_tcbuffer(gs._inner, temp._inner)
	return int(_cret)
}


// AcontainsTcbufferCbuffer wraps MEOS C function acontains_tcbuffer_cbuffer.
func AcontainsTcbufferCbuffer(temp *Temporal, cb *Cbuffer) int {
	_cret := C.acontains_tcbuffer_cbuffer(temp._inner, cb._inner)
	return int(_cret)
}


// AcontainsTcbufferGeo wraps MEOS C function acontains_tcbuffer_geo.
func AcontainsTcbufferGeo(temp *Temporal, gs *Geom) int {
	_cret := C.acontains_tcbuffer_geo(temp._inner, gs._inner)
	return int(_cret)
}


// AcoversCbufferTcbuffer wraps MEOS C function acovers_cbuffer_tcbuffer.
func AcoversCbufferTcbuffer(cb *Cbuffer, temp *Temporal) int {
	_cret := C.acovers_cbuffer_tcbuffer(cb._inner, temp._inner)
	return int(_cret)
}


// AcoversGeoTcbuffer wraps MEOS C function acovers_geo_tcbuffer.
func AcoversGeoTcbuffer(gs *Geom, temp *Temporal) int {
	_cret := C.acovers_geo_tcbuffer(gs._inner, temp._inner)
	return int(_cret)
}


// AcoversTcbufferCbuffer wraps MEOS C function acovers_tcbuffer_cbuffer.
func AcoversTcbufferCbuffer(temp *Temporal, cb *Cbuffer) int {
	_cret := C.acovers_tcbuffer_cbuffer(temp._inner, cb._inner)
	return int(_cret)
}


// AcoversTcbufferGeo wraps MEOS C function acovers_tcbuffer_geo.
func AcoversTcbufferGeo(temp *Temporal, gs *Geom) int {
	_cret := C.acovers_tcbuffer_geo(temp._inner, gs._inner)
	return int(_cret)
}


// AcoversTcbufferTcbuffer wraps MEOS C function acovers_tcbuffer_tcbuffer.
func AcoversTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.acovers_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return int(_cret)
}


// AdisjointTcbufferGeo wraps MEOS C function adisjoint_tcbuffer_geo.
func AdisjointTcbufferGeo(temp *Temporal, gs *Geom) int {
	_cret := C.adisjoint_tcbuffer_geo(temp._inner, gs._inner)
	return int(_cret)
}


// AdisjointTcbufferCbuffer wraps MEOS C function adisjoint_tcbuffer_cbuffer.
func AdisjointTcbufferCbuffer(temp *Temporal, cb *Cbuffer) int {
	_cret := C.adisjoint_tcbuffer_cbuffer(temp._inner, cb._inner)
	return int(_cret)
}


// AdisjointTcbufferTcbuffer wraps MEOS C function adisjoint_tcbuffer_tcbuffer.
func AdisjointTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.adisjoint_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return int(_cret)
}


// AdwithinTcbufferGeo wraps MEOS C function adwithin_tcbuffer_geo.
func AdwithinTcbufferGeo(temp *Temporal, gs *Geom, dist float64) int {
	_cret := C.adwithin_tcbuffer_geo(temp._inner, gs._inner, C.double(dist))
	return int(_cret)
}


// AdwithinTcbufferCbuffer wraps MEOS C function adwithin_tcbuffer_cbuffer.
func AdwithinTcbufferCbuffer(temp *Temporal, cb *Cbuffer, dist float64) int {
	_cret := C.adwithin_tcbuffer_cbuffer(temp._inner, cb._inner, C.double(dist))
	return int(_cret)
}


// AdwithinTcbufferTcbuffer wraps MEOS C function adwithin_tcbuffer_tcbuffer.
func AdwithinTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal, dist float64) int {
	_cret := C.adwithin_tcbuffer_tcbuffer(temp1._inner, temp2._inner, C.double(dist))
	return int(_cret)
}


// AintersectsTcbufferGeo wraps MEOS C function aintersects_tcbuffer_geo.
func AintersectsTcbufferGeo(temp *Temporal, gs *Geom) int {
	_cret := C.aintersects_tcbuffer_geo(temp._inner, gs._inner)
	return int(_cret)
}


// AintersectsTcbufferCbuffer wraps MEOS C function aintersects_tcbuffer_cbuffer.
func AintersectsTcbufferCbuffer(temp *Temporal, cb *Cbuffer) int {
	_cret := C.aintersects_tcbuffer_cbuffer(temp._inner, cb._inner)
	return int(_cret)
}


// AintersectsTcbufferTcbuffer wraps MEOS C function aintersects_tcbuffer_tcbuffer.
func AintersectsTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.aintersects_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return int(_cret)
}


// AtouchesTcbufferGeo wraps MEOS C function atouches_tcbuffer_geo.
func AtouchesTcbufferGeo(temp *Temporal, gs *Geom) int {
	_cret := C.atouches_tcbuffer_geo(temp._inner, gs._inner)
	return int(_cret)
}


// AtouchesTcbufferCbuffer wraps MEOS C function atouches_tcbuffer_cbuffer.
func AtouchesTcbufferCbuffer(temp *Temporal, cb *Cbuffer) int {
	_cret := C.atouches_tcbuffer_cbuffer(temp._inner, cb._inner)
	return int(_cret)
}


// AtouchesTcbufferTcbuffer wraps MEOS C function atouches_tcbuffer_tcbuffer.
func AtouchesTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.atouches_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return int(_cret)
}


// EcontainsCbufferTcbuffer wraps MEOS C function econtains_cbuffer_tcbuffer.
func EcontainsCbufferTcbuffer(cb *Cbuffer, temp *Temporal) int {
	_cret := C.econtains_cbuffer_tcbuffer(cb._inner, temp._inner)
	return int(_cret)
}


// EcontainsTcbufferCbuffer wraps MEOS C function econtains_tcbuffer_cbuffer.
func EcontainsTcbufferCbuffer(temp *Temporal, cb *Cbuffer) int {
	_cret := C.econtains_tcbuffer_cbuffer(temp._inner, cb._inner)
	return int(_cret)
}


// EcontainsTcbufferGeo wraps MEOS C function econtains_tcbuffer_geo.
func EcontainsTcbufferGeo(temp *Temporal, gs *Geom) int {
	_cret := C.econtains_tcbuffer_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EcoversCbufferTcbuffer wraps MEOS C function ecovers_cbuffer_tcbuffer.
func EcoversCbufferTcbuffer(cb *Cbuffer, temp *Temporal) int {
	_cret := C.ecovers_cbuffer_tcbuffer(cb._inner, temp._inner)
	return int(_cret)
}


// EcoversGeoTcbuffer wraps MEOS C function ecovers_geo_tcbuffer.
func EcoversGeoTcbuffer(gs *Geom, temp *Temporal) int {
	_cret := C.ecovers_geo_tcbuffer(gs._inner, temp._inner)
	return int(_cret)
}


// EcoversTcbufferCbuffer wraps MEOS C function ecovers_tcbuffer_cbuffer.
func EcoversTcbufferCbuffer(temp *Temporal, cb *Cbuffer) int {
	_cret := C.ecovers_tcbuffer_cbuffer(temp._inner, cb._inner)
	return int(_cret)
}


// EcoversTcbufferGeo wraps MEOS C function ecovers_tcbuffer_geo.
func EcoversTcbufferGeo(temp *Temporal, gs *Geom) int {
	_cret := C.ecovers_tcbuffer_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EcoversTcbufferTcbuffer wraps MEOS C function ecovers_tcbuffer_tcbuffer.
func EcoversTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ecovers_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return int(_cret)
}


// EdisjointTcbufferGeo wraps MEOS C function edisjoint_tcbuffer_geo.
func EdisjointTcbufferGeo(temp *Temporal, gs *Geom) int {
	_cret := C.edisjoint_tcbuffer_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EdisjointTcbufferCbuffer wraps MEOS C function edisjoint_tcbuffer_cbuffer.
func EdisjointTcbufferCbuffer(temp *Temporal, cb *Cbuffer) int {
	_cret := C.edisjoint_tcbuffer_cbuffer(temp._inner, cb._inner)
	return int(_cret)
}


// EdwithinTcbufferGeo wraps MEOS C function edwithin_tcbuffer_geo.
func EdwithinTcbufferGeo(temp *Temporal, gs *Geom, dist float64) int {
	_cret := C.edwithin_tcbuffer_geo(temp._inner, gs._inner, C.double(dist))
	return int(_cret)
}


// EdwithinTcbufferCbuffer wraps MEOS C function edwithin_tcbuffer_cbuffer.
func EdwithinTcbufferCbuffer(temp *Temporal, cb *Cbuffer, dist float64) int {
	_cret := C.edwithin_tcbuffer_cbuffer(temp._inner, cb._inner, C.double(dist))
	return int(_cret)
}


// EdwithinTcbufferTcbuffer wraps MEOS C function edwithin_tcbuffer_tcbuffer.
func EdwithinTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal, dist float64) int {
	_cret := C.edwithin_tcbuffer_tcbuffer(temp1._inner, temp2._inner, C.double(dist))
	return int(_cret)
}


// EintersectsTcbufferGeo wraps MEOS C function eintersects_tcbuffer_geo.
func EintersectsTcbufferGeo(temp *Temporal, gs *Geom) int {
	_cret := C.eintersects_tcbuffer_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EintersectsTcbufferCbuffer wraps MEOS C function eintersects_tcbuffer_cbuffer.
func EintersectsTcbufferCbuffer(temp *Temporal, cb *Cbuffer) int {
	_cret := C.eintersects_tcbuffer_cbuffer(temp._inner, cb._inner)
	return int(_cret)
}


// EintersectsTcbufferTcbuffer wraps MEOS C function eintersects_tcbuffer_tcbuffer.
func EintersectsTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.eintersects_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return int(_cret)
}


// EtouchesTcbufferGeo wraps MEOS C function etouches_tcbuffer_geo.
func EtouchesTcbufferGeo(temp *Temporal, gs *Geom) int {
	_cret := C.etouches_tcbuffer_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EtouchesTcbufferCbuffer wraps MEOS C function etouches_tcbuffer_cbuffer.
func EtouchesTcbufferCbuffer(temp *Temporal, cb *Cbuffer) int {
	_cret := C.etouches_tcbuffer_cbuffer(temp._inner, cb._inner)
	return int(_cret)
}


// EtouchesTcbufferTcbuffer wraps MEOS C function etouches_tcbuffer_tcbuffer.
func EtouchesTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.etouches_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return int(_cret)
}


// TcontainsCbufferTcbuffer wraps MEOS C function tcontains_cbuffer_tcbuffer.
func TcontainsCbufferTcbuffer(cb *Cbuffer, temp *Temporal) *Temporal {
	_cret := C.tcontains_cbuffer_tcbuffer(cb._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TcontainsGeoTcbuffer wraps MEOS C function tcontains_geo_tcbuffer.
func TcontainsGeoTcbuffer(gs *Geom, temp *Temporal) *Temporal {
	_cret := C.tcontains_geo_tcbuffer(gs._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TcontainsTcbufferGeo wraps MEOS C function tcontains_tcbuffer_geo.
func TcontainsTcbufferGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tcontains_tcbuffer_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TcontainsTcbufferCbuffer wraps MEOS C function tcontains_tcbuffer_cbuffer.
func TcontainsTcbufferCbuffer(temp *Temporal, cb *Cbuffer) *Temporal {
	_cret := C.tcontains_tcbuffer_cbuffer(temp._inner, cb._inner)
	return &Temporal{_inner: _cret}
}


// TcontainsTcbufferTcbuffer wraps MEOS C function tcontains_tcbuffer_tcbuffer.
func TcontainsTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tcontains_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TcoversCbufferTcbuffer wraps MEOS C function tcovers_cbuffer_tcbuffer.
func TcoversCbufferTcbuffer(cb *Cbuffer, temp *Temporal) *Temporal {
	_cret := C.tcovers_cbuffer_tcbuffer(cb._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TcoversGeoTcbuffer wraps MEOS C function tcovers_geo_tcbuffer.
func TcoversGeoTcbuffer(gs *Geom, temp *Temporal) *Temporal {
	_cret := C.tcovers_geo_tcbuffer(gs._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TcoversTcbufferGeo wraps MEOS C function tcovers_tcbuffer_geo.
func TcoversTcbufferGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tcovers_tcbuffer_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TcoversTcbufferCbuffer wraps MEOS C function tcovers_tcbuffer_cbuffer.
func TcoversTcbufferCbuffer(temp *Temporal, cb *Cbuffer) *Temporal {
	_cret := C.tcovers_tcbuffer_cbuffer(temp._inner, cb._inner)
	return &Temporal{_inner: _cret}
}


// TcoversTcbufferTcbuffer wraps MEOS C function tcovers_tcbuffer_tcbuffer.
func TcoversTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tcovers_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TdwithinGeoTcbuffer wraps MEOS C function tdwithin_geo_tcbuffer.
func TdwithinGeoTcbuffer(gs *Geom, temp *Temporal, dist float64) *Temporal {
	_cret := C.tdwithin_geo_tcbuffer(gs._inner, temp._inner, C.double(dist))
	return &Temporal{_inner: _cret}
}


// TdwithinTcbufferGeo wraps MEOS C function tdwithin_tcbuffer_geo.
func TdwithinTcbufferGeo(temp *Temporal, gs *Geom, dist float64) *Temporal {
	_cret := C.tdwithin_tcbuffer_geo(temp._inner, gs._inner, C.double(dist))
	return &Temporal{_inner: _cret}
}


// TdwithinTcbufferCbuffer wraps MEOS C function tdwithin_tcbuffer_cbuffer.
func TdwithinTcbufferCbuffer(temp *Temporal, cb *Cbuffer, dist float64) *Temporal {
	_cret := C.tdwithin_tcbuffer_cbuffer(temp._inner, cb._inner, C.double(dist))
	return &Temporal{_inner: _cret}
}


// TdwithinTcbufferTcbuffer wraps MEOS C function tdwithin_tcbuffer_tcbuffer.
func TdwithinTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal, dist float64) *Temporal {
	_cret := C.tdwithin_tcbuffer_tcbuffer(temp1._inner, temp2._inner, C.double(dist))
	return &Temporal{_inner: _cret}
}


// TdisjointCbufferTcbuffer wraps MEOS C function tdisjoint_cbuffer_tcbuffer.
func TdisjointCbufferTcbuffer(cb *Cbuffer, temp *Temporal) *Temporal {
	_cret := C.tdisjoint_cbuffer_tcbuffer(cb._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TdisjointGeoTcbuffer wraps MEOS C function tdisjoint_geo_tcbuffer.
func TdisjointGeoTcbuffer(gs *Geom, temp *Temporal) *Temporal {
	_cret := C.tdisjoint_geo_tcbuffer(gs._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TdisjointTcbufferGeo wraps MEOS C function tdisjoint_tcbuffer_geo.
func TdisjointTcbufferGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tdisjoint_tcbuffer_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TdisjointTcbufferCbuffer wraps MEOS C function tdisjoint_tcbuffer_cbuffer.
func TdisjointTcbufferCbuffer(temp *Temporal, cb *Cbuffer) *Temporal {
	_cret := C.tdisjoint_tcbuffer_cbuffer(temp._inner, cb._inner)
	return &Temporal{_inner: _cret}
}


// TdisjointTcbufferTcbuffer wraps MEOS C function tdisjoint_tcbuffer_tcbuffer.
func TdisjointTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tdisjoint_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TintersectsCbufferTcbuffer wraps MEOS C function tintersects_cbuffer_tcbuffer.
func TintersectsCbufferTcbuffer(cb *Cbuffer, temp *Temporal) *Temporal {
	_cret := C.tintersects_cbuffer_tcbuffer(cb._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TintersectsGeoTcbuffer wraps MEOS C function tintersects_geo_tcbuffer.
func TintersectsGeoTcbuffer(gs *Geom, temp *Temporal) *Temporal {
	_cret := C.tintersects_geo_tcbuffer(gs._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TintersectsTcbufferGeo wraps MEOS C function tintersects_tcbuffer_geo.
func TintersectsTcbufferGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tintersects_tcbuffer_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TintersectsTcbufferCbuffer wraps MEOS C function tintersects_tcbuffer_cbuffer.
func TintersectsTcbufferCbuffer(temp *Temporal, cb *Cbuffer) *Temporal {
	_cret := C.tintersects_tcbuffer_cbuffer(temp._inner, cb._inner)
	return &Temporal{_inner: _cret}
}


// TintersectsTcbufferTcbuffer wraps MEOS C function tintersects_tcbuffer_tcbuffer.
func TintersectsTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tintersects_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TtouchesGeoTcbuffer wraps MEOS C function ttouches_geo_tcbuffer.
func TtouchesGeoTcbuffer(gs *Geom, temp *Temporal) *Temporal {
	_cret := C.ttouches_geo_tcbuffer(gs._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TtouchesTcbufferGeo wraps MEOS C function ttouches_tcbuffer_geo.
func TtouchesTcbufferGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.ttouches_tcbuffer_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TtouchesCbufferTcbuffer wraps MEOS C function ttouches_cbuffer_tcbuffer.
func TtouchesCbufferTcbuffer(cb *Cbuffer, temp *Temporal) *Temporal {
	_cret := C.ttouches_cbuffer_tcbuffer(cb._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TtouchesTcbufferCbuffer wraps MEOS C function ttouches_tcbuffer_cbuffer.
func TtouchesTcbufferCbuffer(temp *Temporal, cb *Cbuffer) *Temporal {
	_cret := C.ttouches_tcbuffer_cbuffer(temp._inner, cb._inner)
	return &Temporal{_inner: _cret}
}


// TtouchesTcbufferTcbuffer wraps MEOS C function ttouches_tcbuffer_tcbuffer.
func TtouchesTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.ttouches_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}

