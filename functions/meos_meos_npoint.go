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

// NpointAsEWKT wraps MEOS C function npoint_as_ewkt.
func NpointAsEWKT(np *Npoint, maxdd int) string {
	_cret := C.npoint_as_ewkt(np._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// NpointAsHexwkb wraps MEOS C function npoint_as_hexwkb.
func NpointAsHexwkb(np *Npoint, variant uint8, size_out unsafe.Pointer) string {
	_cret := C.npoint_as_hexwkb(np._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return C.GoString(_cret)
}


// NpointAsText wraps MEOS C function npoint_as_text.
func NpointAsText(np *Npoint, maxdd int) string {
	_cret := C.npoint_as_text(np._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// NpointAsWKB wraps MEOS C function npoint_as_wkb.
func NpointAsWKB(np *Npoint, variant uint8, size_out unsafe.Pointer) unsafe.Pointer {
	_cret := C.npoint_as_wkb(np._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return unsafe.Pointer(_cret)
}


// NpointFromHexwkb wraps MEOS C function npoint_from_hexwkb.
func NpointFromHexwkb(hexwkb string) *Npoint {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	_cret := C.npoint_from_hexwkb(_c_hexwkb)
	return &Npoint{_inner: _cret}
}


// NpointFromWKB wraps MEOS C function npoint_from_wkb.
func NpointFromWKB(wkb unsafe.Pointer, size uint) *Npoint {
	_cret := C.npoint_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	return &Npoint{_inner: _cret}
}


// NpointIn wraps MEOS C function npoint_in.
func NpointIn(str string) *Npoint {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.npoint_in(_c_str)
	return &Npoint{_inner: _cret}
}


// NpointOut wraps MEOS C function npoint_out.
func NpointOut(np *Npoint, maxdd int) string {
	_cret := C.npoint_out(np._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// NsegmentIn wraps MEOS C function nsegment_in.
func NsegmentIn(str string) *Nsegment {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.nsegment_in(_c_str)
	return &Nsegment{_inner: _cret}
}


// NsegmentOut wraps MEOS C function nsegment_out.
func NsegmentOut(ns *Nsegment, maxdd int) string {
	_cret := C.nsegment_out(ns._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// NpointMake wraps MEOS C function npoint_make.
func NpointMake(rid int64, pos float64) *Npoint {
	_cret := C.npoint_make(C.int64_t(rid), C.double(pos))
	return &Npoint{_inner: _cret}
}


// NsegmentMake wraps MEOS C function nsegment_make.
func NsegmentMake(rid int64, pos1 float64, pos2 float64) *Nsegment {
	_cret := C.nsegment_make(C.int64_t(rid), C.double(pos1), C.double(pos2))
	return &Nsegment{_inner: _cret}
}


// GeompointToNpoint wraps MEOS C function geompoint_to_npoint.
func GeompointToNpoint(gs *Geom) *Npoint {
	_cret := C.geompoint_to_npoint(gs._inner)
	return &Npoint{_inner: _cret}
}


// GeomToNsegment wraps MEOS C function geom_to_nsegment.
func GeomToNsegment(gs *Geom) *Nsegment {
	_cret := C.geom_to_nsegment(gs._inner)
	return &Nsegment{_inner: _cret}
}


// NpointToGeompoint wraps MEOS C function npoint_to_geompoint.
func NpointToGeompoint(np *Npoint) *Geom {
	_cret := C.npoint_to_geompoint(np._inner)
	return &Geom{_inner: _cret}
}


// NpointToNsegment wraps MEOS C function npoint_to_nsegment.
func NpointToNsegment(np *Npoint) *Nsegment {
	_cret := C.npoint_to_nsegment(np._inner)
	return &Nsegment{_inner: _cret}
}


// NpointToSTBOX wraps MEOS C function npoint_to_stbox.
func NpointToSTBOX(np *Npoint) *STBox {
	_cret := C.npoint_to_stbox(np._inner)
	return &STBox{_inner: _cret}
}


// NsegmentToGeom wraps MEOS C function nsegment_to_geom.
func NsegmentToGeom(ns *Nsegment) *Geom {
	_cret := C.nsegment_to_geom(ns._inner)
	return &Geom{_inner: _cret}
}


// NsegmentToSTBOX wraps MEOS C function nsegment_to_stbox.
func NsegmentToSTBOX(ns *Nsegment) *STBox {
	_cret := C.nsegment_to_stbox(ns._inner)
	return &STBox{_inner: _cret}
}


// NpointHash wraps MEOS C function npoint_hash.
func NpointHash(np *Npoint) uint32 {
	_cret := C.npoint_hash(np._inner)
	return uint32(_cret)
}


// NpointHashExtended wraps MEOS C function npoint_hash_extended.
func NpointHashExtended(np *Npoint, seed uint64) uint64 {
	_cret := C.npoint_hash_extended(np._inner, C.uint64_t(seed))
	return uint64(_cret)
}


// NpointPosition wraps MEOS C function npoint_position.
func NpointPosition(np *Npoint) float64 {
	_cret := C.npoint_position(np._inner)
	return float64(_cret)
}


// NpointRoute wraps MEOS C function npoint_route.
func NpointRoute(np *Npoint) int64 {
	_cret := C.npoint_route(np._inner)
	return int64(_cret)
}


// NsegmentEndPosition wraps MEOS C function nsegment_end_position.
func NsegmentEndPosition(ns *Nsegment) float64 {
	_cret := C.nsegment_end_position(ns._inner)
	return float64(_cret)
}


// NsegmentRoute wraps MEOS C function nsegment_route.
func NsegmentRoute(ns *Nsegment) int64 {
	_cret := C.nsegment_route(ns._inner)
	return int64(_cret)
}


// NsegmentStartPosition wraps MEOS C function nsegment_start_position.
func NsegmentStartPosition(ns *Nsegment) float64 {
	_cret := C.nsegment_start_position(ns._inner)
	return float64(_cret)
}


// RouteExists wraps MEOS C function route_exists.
func RouteExists(rid int64) bool {
	_cret := C.route_exists(C.int64_t(rid))
	return bool(_cret)
}


// RouteGeom wraps MEOS C function route_geom.
func RouteGeom(rid int64) *Geom {
	_cret := C.route_geom(C.int64_t(rid))
	return &Geom{_inner: _cret}
}


// RouteLength wraps MEOS C function route_length.
func RouteLength(rid int64) float64 {
	_cret := C.route_length(C.int64_t(rid))
	return float64(_cret)
}


// NpointRound wraps MEOS C function npoint_round.
func NpointRound(np *Npoint, maxdd int) *Npoint {
	_cret := C.npoint_round(np._inner, C.int(maxdd))
	return &Npoint{_inner: _cret}
}


// NsegmentRound wraps MEOS C function nsegment_round.
func NsegmentRound(ns *Nsegment, maxdd int) *Nsegment {
	_cret := C.nsegment_round(ns._inner, C.int(maxdd))
	return &Nsegment{_inner: _cret}
}


// GetSRIDWays wraps MEOS C function get_srid_ways.
func GetSRIDWays() int32 {
	_cret := C.get_srid_ways()
	return int32(_cret)
}


// NpointSRID wraps MEOS C function npoint_srid.
func NpointSRID(np *Npoint) int32 {
	_cret := C.npoint_srid(np._inner)
	return int32(_cret)
}


// NsegmentSRID wraps MEOS C function nsegment_srid.
func NsegmentSRID(ns *Nsegment) int32 {
	_cret := C.nsegment_srid(ns._inner)
	return int32(_cret)
}


// NpointTimestamptzToSTBOX wraps MEOS C function npoint_timestamptz_to_stbox.
func NpointTimestamptzToSTBOX(np *Npoint, t int64) *STBox {
	_cret := C.npoint_timestamptz_to_stbox(np._inner, C.TimestampTz(t))
	return &STBox{_inner: _cret}
}


// NpointTstzspanToSTBOX wraps MEOS C function npoint_tstzspan_to_stbox.
func NpointTstzspanToSTBOX(np *Npoint, s *Span) *STBox {
	_cret := C.npoint_tstzspan_to_stbox(np._inner, s._inner)
	return &STBox{_inner: _cret}
}


// NpointCmp wraps MEOS C function npoint_cmp.
func NpointCmp(np1 *Npoint, np2 *Npoint) int {
	_cret := C.npoint_cmp(np1._inner, np2._inner)
	return int(_cret)
}


// NpointEq wraps MEOS C function npoint_eq.
func NpointEq(np1 *Npoint, np2 *Npoint) bool {
	_cret := C.npoint_eq(np1._inner, np2._inner)
	return bool(_cret)
}


// NpointGe wraps MEOS C function npoint_ge.
func NpointGe(np1 *Npoint, np2 *Npoint) bool {
	_cret := C.npoint_ge(np1._inner, np2._inner)
	return bool(_cret)
}


// NpointGt wraps MEOS C function npoint_gt.
func NpointGt(np1 *Npoint, np2 *Npoint) bool {
	_cret := C.npoint_gt(np1._inner, np2._inner)
	return bool(_cret)
}


// NpointLe wraps MEOS C function npoint_le.
func NpointLe(np1 *Npoint, np2 *Npoint) bool {
	_cret := C.npoint_le(np1._inner, np2._inner)
	return bool(_cret)
}


// NpointLt wraps MEOS C function npoint_lt.
func NpointLt(np1 *Npoint, np2 *Npoint) bool {
	_cret := C.npoint_lt(np1._inner, np2._inner)
	return bool(_cret)
}


// NpointNe wraps MEOS C function npoint_ne.
func NpointNe(np1 *Npoint, np2 *Npoint) bool {
	_cret := C.npoint_ne(np1._inner, np2._inner)
	return bool(_cret)
}


// NpointSame wraps MEOS C function npoint_same.
func NpointSame(np1 *Npoint, np2 *Npoint) bool {
	_cret := C.npoint_same(np1._inner, np2._inner)
	return bool(_cret)
}


// NsegmentCmp wraps MEOS C function nsegment_cmp.
func NsegmentCmp(ns1 *Nsegment, ns2 *Nsegment) int {
	_cret := C.nsegment_cmp(ns1._inner, ns2._inner)
	return int(_cret)
}


// NsegmentEq wraps MEOS C function nsegment_eq.
func NsegmentEq(ns1 *Nsegment, ns2 *Nsegment) bool {
	_cret := C.nsegment_eq(ns1._inner, ns2._inner)
	return bool(_cret)
}


// NsegmentGe wraps MEOS C function nsegment_ge.
func NsegmentGe(ns1 *Nsegment, ns2 *Nsegment) bool {
	_cret := C.nsegment_ge(ns1._inner, ns2._inner)
	return bool(_cret)
}


// NsegmentGt wraps MEOS C function nsegment_gt.
func NsegmentGt(ns1 *Nsegment, ns2 *Nsegment) bool {
	_cret := C.nsegment_gt(ns1._inner, ns2._inner)
	return bool(_cret)
}


// NsegmentLe wraps MEOS C function nsegment_le.
func NsegmentLe(ns1 *Nsegment, ns2 *Nsegment) bool {
	_cret := C.nsegment_le(ns1._inner, ns2._inner)
	return bool(_cret)
}


// NsegmentLt wraps MEOS C function nsegment_lt.
func NsegmentLt(ns1 *Nsegment, ns2 *Nsegment) bool {
	_cret := C.nsegment_lt(ns1._inner, ns2._inner)
	return bool(_cret)
}


// NsegmentNe wraps MEOS C function nsegment_ne.
func NsegmentNe(ns1 *Nsegment, ns2 *Nsegment) bool {
	_cret := C.nsegment_ne(ns1._inner, ns2._inner)
	return bool(_cret)
}


// NpointsetIn wraps MEOS C function npointset_in.
func NpointsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.npointset_in(_c_str)
	return &Set{_inner: _cret}
}


// NpointsetOut wraps MEOS C function npointset_out.
func NpointsetOut(s *Set, maxdd int) string {
	_cret := C.npointset_out(s._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// NpointsetMake wraps MEOS C function npointset_make.
func NpointsetMake(values unsafe.Pointer, count int) *Set {
	_cret := C.npointset_make((**C.Npoint)(unsafe.Pointer(values)), C.int(count))
	return &Set{_inner: _cret}
}


// NpointToSet wraps MEOS C function npoint_to_set.
func NpointToSet(np *Npoint) *Set {
	_cret := C.npoint_to_set(np._inner)
	return &Set{_inner: _cret}
}


// NpointsetEndValue wraps MEOS C function npointset_end_value.
func NpointsetEndValue(s *Set) *Npoint {
	_cret := C.npointset_end_value(s._inner)
	return &Npoint{_inner: _cret}
}


// NpointsetRoutes wraps MEOS C function npointset_routes.
func NpointsetRoutes(s *Set) *Set {
	_cret := C.npointset_routes(s._inner)
	return &Set{_inner: _cret}
}


// NpointsetStartValue wraps MEOS C function npointset_start_value.
func NpointsetStartValue(s *Set) *Npoint {
	_cret := C.npointset_start_value(s._inner)
	return &Npoint{_inner: _cret}
}


// NpointsetValueN wraps MEOS C function npointset_value_n.
func NpointsetValueN(s *Set, n int) (bool, *Npoint) {
	var _out_result *C.Npoint
	_cret := C.npointset_value_n(s._inner, C.int(n), &_out_result)
	return bool(_cret), &Npoint{_inner: _out_result}
}


// NpointsetValues wraps MEOS C function npointset_values.
func NpointsetValues(s *Set, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.npointset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// ContainedNpointSet wraps MEOS C function contained_npoint_set.
func ContainedNpointSet(np *Npoint, s *Set) bool {
	_cret := C.contained_npoint_set(np._inner, s._inner)
	return bool(_cret)
}


// ContainsSetNpoint wraps MEOS C function contains_set_npoint.
func ContainsSetNpoint(s *Set, np *Npoint) bool {
	_cret := C.contains_set_npoint(s._inner, np._inner)
	return bool(_cret)
}


// IntersectionNpointSet wraps MEOS C function intersection_npoint_set.
func IntersectionNpointSet(np *Npoint, s *Set) *Set {
	_cret := C.intersection_npoint_set(np._inner, s._inner)
	return &Set{_inner: _cret}
}


// IntersectionSetNpoint wraps MEOS C function intersection_set_npoint.
func IntersectionSetNpoint(s *Set, np *Npoint) *Set {
	_cret := C.intersection_set_npoint(s._inner, np._inner)
	return &Set{_inner: _cret}
}


// MinusNpointSet wraps MEOS C function minus_npoint_set.
func MinusNpointSet(np *Npoint, s *Set) *Set {
	_cret := C.minus_npoint_set(np._inner, s._inner)
	return &Set{_inner: _cret}
}


// MinusSetNpoint wraps MEOS C function minus_set_npoint.
func MinusSetNpoint(s *Set, np *Npoint) *Set {
	_cret := C.minus_set_npoint(s._inner, np._inner)
	return &Set{_inner: _cret}
}


// NpointUnionTransfn wraps MEOS C function npoint_union_transfn.
func NpointUnionTransfn(state *Set, np *Npoint) *Set {
	_cret := C.npoint_union_transfn(state._inner, np._inner)
	return &Set{_inner: _cret}
}


// UnionNpointSet wraps MEOS C function union_npoint_set.
func UnionNpointSet(np *Npoint, s *Set) *Set {
	_cret := C.gunion_npoint_set(np._inner, s._inner)
	return &Set{_inner: _cret}
}


// UnionSetNpoint wraps MEOS C function union_set_npoint.
func UnionSetNpoint(s *Set, np *Npoint) *Set {
	_cret := C.gunion_set_npoint(s._inner, np._inner)
	return &Set{_inner: _cret}
}


// TnpointIn wraps MEOS C function tnpoint_in.
func TnpointIn(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tnpoint_in(_c_str)
	return &Temporal{_inner: _cret}
}


// TnpointFromMFJSON wraps MEOS C function tnpoint_from_mfjson.
func TnpointFromMFJSON(mfjson string) *Temporal {
	_c_mfjson := C.CString(mfjson)
	defer C.free(unsafe.Pointer(_c_mfjson))
	_cret := C.tnpoint_from_mfjson(_c_mfjson)
	return &Temporal{_inner: _cret}
}


// TnpointOut wraps MEOS C function tnpoint_out.
func TnpointOut(temp *Temporal, maxdd int) string {
	_cret := C.tnpoint_out(temp._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// TnpointinstMake wraps MEOS C function tnpointinst_make.
func TnpointinstMake(np *Npoint, t int64) *TInstant {
	_cret := C.tnpointinst_make(np._inner, C.TimestampTz(t))
	return &TInstant{_inner: _cret}
}


// TnpointFromBaseTemp wraps MEOS C function tnpoint_from_base_temp.
func TnpointFromBaseTemp(np *Npoint, temp *Temporal) *Temporal {
	_cret := C.tnpoint_from_base_temp(np._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TnpointseqFromBaseTstzset wraps MEOS C function tnpointseq_from_base_tstzset.
func TnpointseqFromBaseTstzset(np *Npoint, s *Set) *TSequence {
	_cret := C.tnpointseq_from_base_tstzset(np._inner, s._inner)
	return &TSequence{_inner: _cret}
}


// TnpointseqFromBaseTstzspan wraps MEOS C function tnpointseq_from_base_tstzspan.
func TnpointseqFromBaseTstzspan(np *Npoint, s *Span, interp Interpolation) *TSequence {
	_cret := C.tnpointseq_from_base_tstzspan(np._inner, s._inner, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TnpointseqsetFromBaseTstzspanset wraps MEOS C function tnpointseqset_from_base_tstzspanset.
func TnpointseqsetFromBaseTstzspanset(np *Npoint, ss *SpanSet, interp Interpolation) *TSequenceSet {
	_cret := C.tnpointseqset_from_base_tstzspanset(np._inner, ss._inner, C.interpType(interp))
	return &TSequenceSet{_inner: _cret}
}


// TgeompointToTnpoint wraps MEOS C function tgeompoint_to_tnpoint.
func TgeompointToTnpoint(temp *Temporal) *Temporal {
	_cret := C.tgeompoint_to_tnpoint(temp._inner)
	return &Temporal{_inner: _cret}
}


// TnpointToTgeompoint wraps MEOS C function tnpoint_to_tgeompoint.
func TnpointToTgeompoint(temp *Temporal) *Temporal {
	_cret := C.tnpoint_to_tgeompoint(temp._inner)
	return &Temporal{_inner: _cret}
}


// TnpointCumulativeLength wraps MEOS C function tnpoint_cumulative_length.
func TnpointCumulativeLength(temp *Temporal) *Temporal {
	_cret := C.tnpoint_cumulative_length(temp._inner)
	return &Temporal{_inner: _cret}
}


// TnpointEndValue wraps MEOS C function tnpoint_end_value.
func TnpointEndValue(temp *Temporal) *Npoint {
	_cret := C.tnpoint_end_value(temp._inner)
	return &Npoint{_inner: _cret}
}


// TnpointLength wraps MEOS C function tnpoint_length.
func TnpointLength(temp *Temporal) float64 {
	_cret := C.tnpoint_length(temp._inner)
	return float64(_cret)
}


// TnpointPositions wraps MEOS C function tnpoint_positions.
func TnpointPositions(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tnpoint_positions(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TnpointRoute wraps MEOS C function tnpoint_route.
func TnpointRoute(temp *Temporal) int64 {
	_cret := C.tnpoint_route(temp._inner)
	return int64(_cret)
}


// TnpointRoutes wraps MEOS C function tnpoint_routes.
func TnpointRoutes(temp *Temporal) *Set {
	_cret := C.tnpoint_routes(temp._inner)
	return &Set{_inner: _cret}
}


// TnpointSpeed wraps MEOS C function tnpoint_speed.
func TnpointSpeed(temp *Temporal) *Temporal {
	_cret := C.tnpoint_speed(temp._inner)
	return &Temporal{_inner: _cret}
}


// TnpointStartValue wraps MEOS C function tnpoint_start_value.
func TnpointStartValue(temp *Temporal) *Npoint {
	_cret := C.tnpoint_start_value(temp._inner)
	return &Npoint{_inner: _cret}
}


// TnpointTrajectory wraps MEOS C function tnpoint_trajectory.
func TnpointTrajectory(temp *Temporal) *Geom {
	_cret := C.tnpoint_trajectory(temp._inner)
	return &Geom{_inner: _cret}
}


// TnpointValueAtTimestamptz wraps MEOS C function tnpoint_value_at_timestamptz.
func TnpointValueAtTimestamptz(temp *Temporal, t int64, strict bool) (bool, *Npoint) {
	var _out_value *C.Npoint
	_cret := C.tnpoint_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_value)
	return bool(_cret), &Npoint{_inner: _out_value}
}


// TnpointValueN wraps MEOS C function tnpoint_value_n.
func TnpointValueN(temp *Temporal, n int) (bool, *Npoint) {
	var _out_result *C.Npoint
	_cret := C.tnpoint_value_n(temp._inner, C.int(n), &_out_result)
	return bool(_cret), &Npoint{_inner: _out_result}
}


// TnpointValues wraps MEOS C function tnpoint_values.
func TnpointValues(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tnpoint_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TnpointTwcentroid wraps MEOS C function tnpoint_twcentroid.
func TnpointTwcentroid(temp *Temporal) *Geom {
	_cret := C.tnpoint_twcentroid(temp._inner)
	return &Geom{_inner: _cret}
}


// TnpointAtGeom wraps MEOS C function tnpoint_at_geom.
func TnpointAtGeom(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tnpoint_at_geom(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TnpointAtNpoint wraps MEOS C function tnpoint_at_npoint.
func TnpointAtNpoint(temp *Temporal, np *Npoint) *Temporal {
	_cret := C.tnpoint_at_npoint(temp._inner, np._inner)
	return &Temporal{_inner: _cret}
}


// TnpointAtNpointset wraps MEOS C function tnpoint_at_npointset.
func TnpointAtNpointset(temp *Temporal, s *Set) *Temporal {
	_cret := C.tnpoint_at_npointset(temp._inner, s._inner)
	return &Temporal{_inner: _cret}
}


// TnpointAtSTBOX wraps MEOS C function tnpoint_at_stbox.
func TnpointAtSTBOX(temp *Temporal, box *STBox, border_inc bool) *Temporal {
	_cret := C.tnpoint_at_stbox(temp._inner, box._inner, C.bool(border_inc))
	return &Temporal{_inner: _cret}
}


// TnpointMinusGeom wraps MEOS C function tnpoint_minus_geom.
func TnpointMinusGeom(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tnpoint_minus_geom(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TnpointMinusNpoint wraps MEOS C function tnpoint_minus_npoint.
func TnpointMinusNpoint(temp *Temporal, np *Npoint) *Temporal {
	_cret := C.tnpoint_minus_npoint(temp._inner, np._inner)
	return &Temporal{_inner: _cret}
}


// TnpointMinusNpointset wraps MEOS C function tnpoint_minus_npointset.
func TnpointMinusNpointset(temp *Temporal, s *Set) *Temporal {
	_cret := C.tnpoint_minus_npointset(temp._inner, s._inner)
	return &Temporal{_inner: _cret}
}


// TnpointMinusSTBOX wraps MEOS C function tnpoint_minus_stbox.
func TnpointMinusSTBOX(temp *Temporal, box *STBox, border_inc bool) *Temporal {
	_cret := C.tnpoint_minus_stbox(temp._inner, box._inner, C.bool(border_inc))
	return &Temporal{_inner: _cret}
}


// TdistanceTnpointNpoint wraps MEOS C function tdistance_tnpoint_npoint.
func TdistanceTnpointNpoint(temp *Temporal, np *Npoint) *Temporal {
	_cret := C.tdistance_tnpoint_npoint(temp._inner, np._inner)
	return &Temporal{_inner: _cret}
}


// TdistanceTnpointGeo wraps MEOS C function tdistance_tnpoint_geo.
func TdistanceTnpointGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tdistance_tnpoint_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TdistanceTnpointTnpoint wraps MEOS C function tdistance_tnpoint_tnpoint.
func TdistanceTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tdistance_tnpoint_tnpoint(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// NadTnpointGeo wraps MEOS C function nad_tnpoint_geo.
func NadTnpointGeo(temp *Temporal, gs *Geom) float64 {
	_cret := C.nad_tnpoint_geo(temp._inner, gs._inner)
	return float64(_cret)
}


// NadTnpointNpoint wraps MEOS C function nad_tnpoint_npoint.
func NadTnpointNpoint(temp *Temporal, np *Npoint) float64 {
	_cret := C.nad_tnpoint_npoint(temp._inner, np._inner)
	return float64(_cret)
}


// NadTnpointSTBOX wraps MEOS C function nad_tnpoint_stbox.
func NadTnpointSTBOX(temp *Temporal, box *STBox) float64 {
	_cret := C.nad_tnpoint_stbox(temp._inner, box._inner)
	return float64(_cret)
}


// NadTnpointTnpoint wraps MEOS C function nad_tnpoint_tnpoint.
func NadTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) float64 {
	_cret := C.nad_tnpoint_tnpoint(temp1._inner, temp2._inner)
	return float64(_cret)
}


// NaiTnpointGeo wraps MEOS C function nai_tnpoint_geo.
func NaiTnpointGeo(temp *Temporal, gs *Geom) *TInstant {
	_cret := C.nai_tnpoint_geo(temp._inner, gs._inner)
	return &TInstant{_inner: _cret}
}


// NaiTnpointNpoint wraps MEOS C function nai_tnpoint_npoint.
func NaiTnpointNpoint(temp *Temporal, np *Npoint) *TInstant {
	_cret := C.nai_tnpoint_npoint(temp._inner, np._inner)
	return &TInstant{_inner: _cret}
}


// NaiTnpointTnpoint wraps MEOS C function nai_tnpoint_tnpoint.
func NaiTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) *TInstant {
	_cret := C.nai_tnpoint_tnpoint(temp1._inner, temp2._inner)
	return &TInstant{_inner: _cret}
}


// ShortestlineTnpointGeo wraps MEOS C function shortestline_tnpoint_geo.
func ShortestlineTnpointGeo(temp *Temporal, gs *Geom) *Geom {
	_cret := C.shortestline_tnpoint_geo(temp._inner, gs._inner)
	return &Geom{_inner: _cret}
}


// ShortestlineTnpointNpoint wraps MEOS C function shortestline_tnpoint_npoint.
func ShortestlineTnpointNpoint(temp *Temporal, np *Npoint) *Geom {
	_cret := C.shortestline_tnpoint_npoint(temp._inner, np._inner)
	return &Geom{_inner: _cret}
}


// ShortestlineTnpointTnpoint wraps MEOS C function shortestline_tnpoint_tnpoint.
func ShortestlineTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) *Geom {
	_cret := C.shortestline_tnpoint_tnpoint(temp1._inner, temp2._inner)
	return &Geom{_inner: _cret}
}


// TnpointTcentroidTransfn wraps MEOS C function tnpoint_tcentroid_transfn.
func TnpointTcentroidTransfn(state *SkipList, temp *Temporal) *SkipList {
	_cret := C.tnpoint_tcentroid_transfn(state._inner, temp._inner)
	return &SkipList{_inner: _cret}
}


// AlwaysEqNpointTnpoint wraps MEOS C function always_eq_npoint_tnpoint.
func AlwaysEqNpointTnpoint(np *Npoint, temp *Temporal) int {
	_cret := C.always_eq_npoint_tnpoint(np._inner, temp._inner)
	return int(_cret)
}


// AlwaysEqTnpointNpoint wraps MEOS C function always_eq_tnpoint_npoint.
func AlwaysEqTnpointNpoint(temp *Temporal, np *Npoint) int {
	_cret := C.always_eq_tnpoint_npoint(temp._inner, np._inner)
	return int(_cret)
}


// AlwaysEqTnpointTnpoint wraps MEOS C function always_eq_tnpoint_tnpoint.
func AlwaysEqTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_eq_tnpoint_tnpoint(temp1._inner, temp2._inner)
	return int(_cret)
}


// AlwaysNeNpointTnpoint wraps MEOS C function always_ne_npoint_tnpoint.
func AlwaysNeNpointTnpoint(np *Npoint, temp *Temporal) int {
	_cret := C.always_ne_npoint_tnpoint(np._inner, temp._inner)
	return int(_cret)
}


// AlwaysNeTnpointNpoint wraps MEOS C function always_ne_tnpoint_npoint.
func AlwaysNeTnpointNpoint(temp *Temporal, np *Npoint) int {
	_cret := C.always_ne_tnpoint_npoint(temp._inner, np._inner)
	return int(_cret)
}


// AlwaysNeTnpointTnpoint wraps MEOS C function always_ne_tnpoint_tnpoint.
func AlwaysNeTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_ne_tnpoint_tnpoint(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverEqNpointTnpoint wraps MEOS C function ever_eq_npoint_tnpoint.
func EverEqNpointTnpoint(np *Npoint, temp *Temporal) int {
	_cret := C.ever_eq_npoint_tnpoint(np._inner, temp._inner)
	return int(_cret)
}


// EverEqTnpointNpoint wraps MEOS C function ever_eq_tnpoint_npoint.
func EverEqTnpointNpoint(temp *Temporal, np *Npoint) int {
	_cret := C.ever_eq_tnpoint_npoint(temp._inner, np._inner)
	return int(_cret)
}


// EverEqTnpointTnpoint wraps MEOS C function ever_eq_tnpoint_tnpoint.
func EverEqTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_eq_tnpoint_tnpoint(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverNeNpointTnpoint wraps MEOS C function ever_ne_npoint_tnpoint.
func EverNeNpointTnpoint(np *Npoint, temp *Temporal) int {
	_cret := C.ever_ne_npoint_tnpoint(np._inner, temp._inner)
	return int(_cret)
}


// EverNeTnpointNpoint wraps MEOS C function ever_ne_tnpoint_npoint.
func EverNeTnpointNpoint(temp *Temporal, np *Npoint) int {
	_cret := C.ever_ne_tnpoint_npoint(temp._inner, np._inner)
	return int(_cret)
}


// EverNeTnpointTnpoint wraps MEOS C function ever_ne_tnpoint_tnpoint.
func EverNeTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_ne_tnpoint_tnpoint(temp1._inner, temp2._inner)
	return int(_cret)
}


// TeqTnpointNpoint wraps MEOS C function teq_tnpoint_npoint.
func TeqTnpointNpoint(temp *Temporal, np *Npoint) *Temporal {
	_cret := C.teq_tnpoint_npoint(temp._inner, np._inner)
	return &Temporal{_inner: _cret}
}


// TneTnpointNpoint wraps MEOS C function tne_tnpoint_npoint.
func TneTnpointNpoint(temp *Temporal, np *Npoint) *Temporal {
	_cret := C.tne_tnpoint_npoint(temp._inner, np._inner)
	return &Temporal{_inner: _cret}
}

