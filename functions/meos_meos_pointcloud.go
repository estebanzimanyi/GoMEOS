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

// PcpointHexIn wraps MEOS C function pcpoint_hex_in.
func PcpointHexIn(str string) *Pcpoint {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.pcpoint_hex_in(_c_str)
	return &Pcpoint{_inner: _cret}
}


// PcpointHexOut wraps MEOS C function pcpoint_hex_out.
func PcpointHexOut(pt *Pcpoint, maxdd int) string {
	_cret := C.pcpoint_hex_out(pt._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// PcpointFromHexwkb wraps MEOS C function pcpoint_from_hexwkb.
func PcpointFromHexwkb(hexwkb string) *Pcpoint {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	_cret := C.pcpoint_from_hexwkb(_c_hexwkb)
	return &Pcpoint{_inner: _cret}
}


// PcpointAsHexwkb wraps MEOS C function pcpoint_as_hexwkb.
func PcpointAsHexwkb(pt *Pcpoint) string {
	_cret := C.pcpoint_as_hexwkb(pt._inner)
	return C.GoString(_cret)
}


// PcpointCopy wraps MEOS C function pcpoint_copy.
func PcpointCopy(pt *Pcpoint) *Pcpoint {
	_cret := C.pcpoint_copy(pt._inner)
	return &Pcpoint{_inner: _cret}
}


// PcpointGetPcid wraps MEOS C function pcpoint_get_pcid.
func PcpointGetPcid(pt *Pcpoint) uint32 {
	_cret := C.pcpoint_get_pcid(pt._inner)
	return uint32(_cret)
}


// PcpointHash wraps MEOS C function pcpoint_hash.
func PcpointHash(pt *Pcpoint) uint32 {
	_cret := C.pcpoint_hash(pt._inner)
	return uint32(_cret)
}


// PcpointHashExtended wraps MEOS C function pcpoint_hash_extended.
func PcpointHashExtended(pt *Pcpoint, seed uint64) uint64 {
	_cret := C.pcpoint_hash_extended(pt._inner, C.uint64_t(seed))
	return uint64(_cret)
}


// PcpointGetX wraps MEOS C function pcpoint_get_x.
func PcpointGetX(pt *Pcpoint, schema *PCSchema, out unsafe.Pointer) bool {
	_cret := C.pcpoint_get_x(pt._inner, schema._inner, (*C.double)(unsafe.Pointer(out)))
	return bool(_cret)
}


// PcpointGetY wraps MEOS C function pcpoint_get_y.
func PcpointGetY(pt *Pcpoint, schema *PCSchema, out unsafe.Pointer) bool {
	_cret := C.pcpoint_get_y(pt._inner, schema._inner, (*C.double)(unsafe.Pointer(out)))
	return bool(_cret)
}


// PcpointGetZ wraps MEOS C function pcpoint_get_z.
func PcpointGetZ(pt *Pcpoint, schema *PCSchema, out unsafe.Pointer) bool {
	_cret := C.pcpoint_get_z(pt._inner, schema._inner, (*C.double)(unsafe.Pointer(out)))
	return bool(_cret)
}


// PcpointGetDim wraps MEOS C function pcpoint_get_dim.
func PcpointGetDim(pt *Pcpoint, schema *PCSchema, name string, out unsafe.Pointer) bool {
	_c_name := C.CString(name)
	defer C.free(unsafe.Pointer(_c_name))
	_cret := C.pcpoint_get_dim(pt._inner, schema._inner, _c_name, (*C.double)(unsafe.Pointer(out)))
	return bool(_cret)
}


// PcpointToTpcbox wraps MEOS C function pcpoint_to_tpcbox.
func PcpointToTpcbox(pt *Pcpoint, schema *PCSchema) *TPCBox {
	_cret := C.pcpoint_to_tpcbox(pt._inner, schema._inner)
	return &TPCBox{_inner: _cret}
}


// MeosPcSchema wraps MEOS C function meos_pc_schema.
func MeosPcSchema(pcid uint32) *PCSchema {
	_cret := C.meos_pc_schema(C.uint32_t(pcid))
	return &PCSchema{_inner: _cret}
}


// MeosPcSchemaRegister wraps MEOS C function meos_pc_schema_register.
func MeosPcSchemaRegister(pcid uint32, schema *PCSchema) {
	C.meos_pc_schema_register(C.uint32_t(pcid), schema._inner)
}


// MeosPcSchemaRegisterXml wraps MEOS C function meos_pc_schema_register_xml.
func MeosPcSchemaRegisterXml(pcid uint32, schema *PCSchema, xml_text string) {
	_c_xml_text := C.CString(xml_text)
	defer C.free(unsafe.Pointer(_c_xml_text))
	C.meos_pc_schema_register_xml(C.uint32_t(pcid), schema._inner, _c_xml_text)
}


// MeosPcSchemaXml wraps MEOS C function meos_pc_schema_xml.
func MeosPcSchemaXml(pcid uint32) string {
	_cret := C.meos_pc_schema_xml(C.uint32_t(pcid))
	return C.GoString(_cret)
}


// MeosPcSchemaClear wraps MEOS C function meos_pc_schema_clear.
func MeosPcSchemaClear() {
	C.meos_pc_schema_clear()
}


// PcpointCmp wraps MEOS C function pcpoint_cmp.
func PcpointCmp(pt1 *Pcpoint, pt2 *Pcpoint) int {
	_cret := C.pcpoint_cmp(pt1._inner, pt2._inner)
	return int(_cret)
}


// PcpointEq wraps MEOS C function pcpoint_eq.
func PcpointEq(pt1 *Pcpoint, pt2 *Pcpoint) bool {
	_cret := C.pcpoint_eq(pt1._inner, pt2._inner)
	return bool(_cret)
}


// PcpointNe wraps MEOS C function pcpoint_ne.
func PcpointNe(pt1 *Pcpoint, pt2 *Pcpoint) bool {
	_cret := C.pcpoint_ne(pt1._inner, pt2._inner)
	return bool(_cret)
}


// PcpointLt wraps MEOS C function pcpoint_lt.
func PcpointLt(pt1 *Pcpoint, pt2 *Pcpoint) bool {
	_cret := C.pcpoint_lt(pt1._inner, pt2._inner)
	return bool(_cret)
}


// PcpointLe wraps MEOS C function pcpoint_le.
func PcpointLe(pt1 *Pcpoint, pt2 *Pcpoint) bool {
	_cret := C.pcpoint_le(pt1._inner, pt2._inner)
	return bool(_cret)
}


// PcpointGt wraps MEOS C function pcpoint_gt.
func PcpointGt(pt1 *Pcpoint, pt2 *Pcpoint) bool {
	_cret := C.pcpoint_gt(pt1._inner, pt2._inner)
	return bool(_cret)
}


// PcpointGe wraps MEOS C function pcpoint_ge.
func PcpointGe(pt1 *Pcpoint, pt2 *Pcpoint) bool {
	_cret := C.pcpoint_ge(pt1._inner, pt2._inner)
	return bool(_cret)
}


// PcpatchHexIn wraps MEOS C function pcpatch_hex_in.
func PcpatchHexIn(str string) *Pcpatch {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.pcpatch_hex_in(_c_str)
	return &Pcpatch{_inner: _cret}
}


// PcpatchHexOut wraps MEOS C function pcpatch_hex_out.
func PcpatchHexOut(pa *Pcpatch, maxdd int) string {
	_cret := C.pcpatch_hex_out(pa._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// PcpatchFromHexwkb wraps MEOS C function pcpatch_from_hexwkb.
func PcpatchFromHexwkb(hexwkb string) *Pcpatch {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	_cret := C.pcpatch_from_hexwkb(_c_hexwkb)
	return &Pcpatch{_inner: _cret}
}


// PcpatchAsHexwkb wraps MEOS C function pcpatch_as_hexwkb.
func PcpatchAsHexwkb(pa *Pcpatch) string {
	_cret := C.pcpatch_as_hexwkb(pa._inner)
	return C.GoString(_cret)
}


// PcpatchCopy wraps MEOS C function pcpatch_copy.
func PcpatchCopy(pa *Pcpatch) *Pcpatch {
	_cret := C.pcpatch_copy(pa._inner)
	return &Pcpatch{_inner: _cret}
}


// PcpatchGetPcid wraps MEOS C function pcpatch_get_pcid.
func PcpatchGetPcid(pa *Pcpatch) uint32 {
	_cret := C.pcpatch_get_pcid(pa._inner)
	return uint32(_cret)
}


// PcpatchNpoints wraps MEOS C function pcpatch_npoints.
func PcpatchNpoints(pa *Pcpatch) uint32 {
	_cret := C.pcpatch_npoints(pa._inner)
	return uint32(_cret)
}


// PcpatchHash wraps MEOS C function pcpatch_hash.
func PcpatchHash(pa *Pcpatch) uint32 {
	_cret := C.pcpatch_hash(pa._inner)
	return uint32(_cret)
}


// PcpatchHashExtended wraps MEOS C function pcpatch_hash_extended.
func PcpatchHashExtended(pa *Pcpatch, seed uint64) uint64 {
	_cret := C.pcpatch_hash_extended(pa._inner, C.uint64_t(seed))
	return uint64(_cret)
}


// PcpatchCmp wraps MEOS C function pcpatch_cmp.
func PcpatchCmp(pa1 *Pcpatch, pa2 *Pcpatch) int {
	_cret := C.pcpatch_cmp(pa1._inner, pa2._inner)
	return int(_cret)
}


// PcpatchEq wraps MEOS C function pcpatch_eq.
func PcpatchEq(pa1 *Pcpatch, pa2 *Pcpatch) bool {
	_cret := C.pcpatch_eq(pa1._inner, pa2._inner)
	return bool(_cret)
}


// PcpatchNe wraps MEOS C function pcpatch_ne.
func PcpatchNe(pa1 *Pcpatch, pa2 *Pcpatch) bool {
	_cret := C.pcpatch_ne(pa1._inner, pa2._inner)
	return bool(_cret)
}


// PcpatchLt wraps MEOS C function pcpatch_lt.
func PcpatchLt(pa1 *Pcpatch, pa2 *Pcpatch) bool {
	_cret := C.pcpatch_lt(pa1._inner, pa2._inner)
	return bool(_cret)
}


// PcpatchLe wraps MEOS C function pcpatch_le.
func PcpatchLe(pa1 *Pcpatch, pa2 *Pcpatch) bool {
	_cret := C.pcpatch_le(pa1._inner, pa2._inner)
	return bool(_cret)
}


// PcpatchGt wraps MEOS C function pcpatch_gt.
func PcpatchGt(pa1 *Pcpatch, pa2 *Pcpatch) bool {
	_cret := C.pcpatch_gt(pa1._inner, pa2._inner)
	return bool(_cret)
}


// PcpatchGe wraps MEOS C function pcpatch_ge.
func PcpatchGe(pa1 *Pcpatch, pa2 *Pcpatch) bool {
	_cret := C.pcpatch_ge(pa1._inner, pa2._inner)
	return bool(_cret)
}


// PcpointsetIn wraps MEOS C function pcpointset_in.
func PcpointsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.pcpointset_in(_c_str)
	return &Set{_inner: _cret}
}


// PcpointsetOut wraps MEOS C function pcpointset_out.
func PcpointsetOut(s *Set, maxdd int) string {
	_cret := C.pcpointset_out(s._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// PcpointsetMake wraps MEOS C function pcpointset_make.
func PcpointsetMake(values unsafe.Pointer, count int) *Set {
	_cret := C.pcpointset_make((**C.Pcpoint)(unsafe.Pointer(values)), C.int(count))
	return &Set{_inner: _cret}
}


// PcpointToSet wraps MEOS C function pcpoint_to_set.
func PcpointToSet(pt *Pcpoint) *Set {
	_cret := C.pcpoint_to_set(pt._inner)
	return &Set{_inner: _cret}
}


// PcpointsetStartValue wraps MEOS C function pcpointset_start_value.
func PcpointsetStartValue(s *Set) *Pcpoint {
	_cret := C.pcpointset_start_value(s._inner)
	return &Pcpoint{_inner: _cret}
}


// PcpointsetEndValue wraps MEOS C function pcpointset_end_value.
func PcpointsetEndValue(s *Set) *Pcpoint {
	_cret := C.pcpointset_end_value(s._inner)
	return &Pcpoint{_inner: _cret}
}


// PcpointsetValueN wraps MEOS C function pcpointset_value_n.
func PcpointsetValueN(s *Set, n int) (bool, *Pcpoint) {
	var _out_result *C.Pcpoint
	_cret := C.pcpointset_value_n(s._inner, C.int(n), &_out_result)
	return bool(_cret), &Pcpoint{_inner: _out_result}
}


// PcpointsetValues wraps MEOS C function pcpointset_values.
func PcpointsetValues(s *Set, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.pcpointset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// ContainsSetPcpoint wraps MEOS C function contains_set_pcpoint.
func ContainsSetPcpoint(s *Set, pt *Pcpoint) bool {
	_cret := C.contains_set_pcpoint(s._inner, pt._inner)
	return bool(_cret)
}


// ContainedPcpointSet wraps MEOS C function contained_pcpoint_set.
func ContainedPcpointSet(pt *Pcpoint, s *Set) bool {
	_cret := C.contained_pcpoint_set(pt._inner, s._inner)
	return bool(_cret)
}


// IntersectionPcpointSet wraps MEOS C function intersection_pcpoint_set.
func IntersectionPcpointSet(pt *Pcpoint, s *Set) *Set {
	_cret := C.intersection_pcpoint_set(pt._inner, s._inner)
	return &Set{_inner: _cret}
}


// IntersectionSetPcpoint wraps MEOS C function intersection_set_pcpoint.
func IntersectionSetPcpoint(s *Set, pt *Pcpoint) *Set {
	_cret := C.intersection_set_pcpoint(s._inner, pt._inner)
	return &Set{_inner: _cret}
}


// MinusPcpointSet wraps MEOS C function minus_pcpoint_set.
func MinusPcpointSet(pt *Pcpoint, s *Set) *Set {
	_cret := C.minus_pcpoint_set(pt._inner, s._inner)
	return &Set{_inner: _cret}
}


// MinusSetPcpoint wraps MEOS C function minus_set_pcpoint.
func MinusSetPcpoint(s *Set, pt *Pcpoint) *Set {
	_cret := C.minus_set_pcpoint(s._inner, pt._inner)
	return &Set{_inner: _cret}
}


// UnionPcpointSet wraps MEOS C function union_pcpoint_set.
func UnionPcpointSet(pt *Pcpoint, s *Set) *Set {
	_cret := C.gunion_pcpoint_set(pt._inner, s._inner)
	return &Set{_inner: _cret}
}


// UnionSetPcpoint wraps MEOS C function union_set_pcpoint.
func UnionSetPcpoint(s *Set, pt *Pcpoint) *Set {
	_cret := C.gunion_set_pcpoint(s._inner, pt._inner)
	return &Set{_inner: _cret}
}


// PcpointUnionTransfn wraps MEOS C function pcpoint_union_transfn.
func PcpointUnionTransfn(state *Set, pt *Pcpoint) *Set {
	_cret := C.pcpoint_union_transfn(state._inner, pt._inner)
	return &Set{_inner: _cret}
}


// PcpatchsetIn wraps MEOS C function pcpatchset_in.
func PcpatchsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.pcpatchset_in(_c_str)
	return &Set{_inner: _cret}
}


// PcpatchsetOut wraps MEOS C function pcpatchset_out.
func PcpatchsetOut(s *Set, maxdd int) string {
	_cret := C.pcpatchset_out(s._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// PcpatchsetMake wraps MEOS C function pcpatchset_make.
func PcpatchsetMake(values unsafe.Pointer, count int) *Set {
	_cret := C.pcpatchset_make((**C.Pcpatch)(unsafe.Pointer(values)), C.int(count))
	return &Set{_inner: _cret}
}


// PcpatchToSet wraps MEOS C function pcpatch_to_set.
func PcpatchToSet(pa *Pcpatch) *Set {
	_cret := C.pcpatch_to_set(pa._inner)
	return &Set{_inner: _cret}
}


// PcpatchsetStartValue wraps MEOS C function pcpatchset_start_value.
func PcpatchsetStartValue(s *Set) *Pcpatch {
	_cret := C.pcpatchset_start_value(s._inner)
	return &Pcpatch{_inner: _cret}
}


// PcpatchsetEndValue wraps MEOS C function pcpatchset_end_value.
func PcpatchsetEndValue(s *Set) *Pcpatch {
	_cret := C.pcpatchset_end_value(s._inner)
	return &Pcpatch{_inner: _cret}
}


// PcpatchsetValueN wraps MEOS C function pcpatchset_value_n.
func PcpatchsetValueN(s *Set, n int) (bool, *Pcpatch) {
	var _out_result *C.Pcpatch
	_cret := C.pcpatchset_value_n(s._inner, C.int(n), &_out_result)
	return bool(_cret), &Pcpatch{_inner: _out_result}
}


// PcpatchsetValues wraps MEOS C function pcpatchset_values.
func PcpatchsetValues(s *Set, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.pcpatchset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// ContainsSetPcpatch wraps MEOS C function contains_set_pcpatch.
func ContainsSetPcpatch(s *Set, pa *Pcpatch) bool {
	_cret := C.contains_set_pcpatch(s._inner, pa._inner)
	return bool(_cret)
}


// ContainedPcpatchSet wraps MEOS C function contained_pcpatch_set.
func ContainedPcpatchSet(pa *Pcpatch, s *Set) bool {
	_cret := C.contained_pcpatch_set(pa._inner, s._inner)
	return bool(_cret)
}


// IntersectionPcpatchSet wraps MEOS C function intersection_pcpatch_set.
func IntersectionPcpatchSet(pa *Pcpatch, s *Set) *Set {
	_cret := C.intersection_pcpatch_set(pa._inner, s._inner)
	return &Set{_inner: _cret}
}


// IntersectionSetPcpatch wraps MEOS C function intersection_set_pcpatch.
func IntersectionSetPcpatch(s *Set, pa *Pcpatch) *Set {
	_cret := C.intersection_set_pcpatch(s._inner, pa._inner)
	return &Set{_inner: _cret}
}


// MinusPcpatchSet wraps MEOS C function minus_pcpatch_set.
func MinusPcpatchSet(pa *Pcpatch, s *Set) *Set {
	_cret := C.minus_pcpatch_set(pa._inner, s._inner)
	return &Set{_inner: _cret}
}


// MinusSetPcpatch wraps MEOS C function minus_set_pcpatch.
func MinusSetPcpatch(s *Set, pa *Pcpatch) *Set {
	_cret := C.minus_set_pcpatch(s._inner, pa._inner)
	return &Set{_inner: _cret}
}


// UnionPcpatchSet wraps MEOS C function union_pcpatch_set.
func UnionPcpatchSet(pa *Pcpatch, s *Set) *Set {
	_cret := C.gunion_pcpatch_set(pa._inner, s._inner)
	return &Set{_inner: _cret}
}


// UnionSetPcpatch wraps MEOS C function union_set_pcpatch.
func UnionSetPcpatch(s *Set, pa *Pcpatch) *Set {
	_cret := C.gunion_set_pcpatch(s._inner, pa._inner)
	return &Set{_inner: _cret}
}


// PcpatchUnionTransfn wraps MEOS C function pcpatch_union_transfn.
func PcpatchUnionTransfn(state *Set, pa *Pcpatch) *Set {
	_cret := C.pcpatch_union_transfn(state._inner, pa._inner)
	return &Set{_inner: _cret}
}


// TpcboxIn wraps MEOS C function tpcbox_in.
func TpcboxIn(str string) *TPCBox {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tpcbox_in(_c_str)
	return &TPCBox{_inner: _cret}
}


// TpcboxOut wraps MEOS C function tpcbox_out.
func TpcboxOut(box *TPCBox, maxdd int) string {
	_cret := C.tpcbox_out(box._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// TpcboxMake wraps MEOS C function tpcbox_make.
func TpcboxMake(hasx bool, hasz bool, hast bool, geodetic bool, srid int32, pcid uint32, xmin float64, xmax float64, ymin float64, ymax float64, zmin float64, zmax float64, period *Span) *TPCBox {
	_cret := C.tpcbox_make(C.bool(hasx), C.bool(hasz), C.bool(hast), C.bool(geodetic), C.int32_t(srid), C.uint32_t(pcid), C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.double(zmin), C.double(zmax), period._inner)
	return &TPCBox{_inner: _cret}
}


// TpcboxCopy wraps MEOS C function tpcbox_copy.
func TpcboxCopy(box *TPCBox) *TPCBox {
	_cret := C.tpcbox_copy(box._inner)
	return &TPCBox{_inner: _cret}
}


// PcpatchToTpcbox wraps MEOS C function pcpatch_to_tpcbox.
func PcpatchToTpcbox(pa *Pcpatch, srid int32) *TPCBox {
	_cret := C.pcpatch_to_tpcbox(pa._inner, C.int32_t(srid))
	return &TPCBox{_inner: _cret}
}


// TpcboxHasx wraps MEOS C function tpcbox_hasx.
func TpcboxHasx(box *TPCBox) bool {
	_cret := C.tpcbox_hasx(box._inner)
	return bool(_cret)
}


// TpcboxHasz wraps MEOS C function tpcbox_hasz.
func TpcboxHasz(box *TPCBox) bool {
	_cret := C.tpcbox_hasz(box._inner)
	return bool(_cret)
}


// TpcboxHast wraps MEOS C function tpcbox_hast.
func TpcboxHast(box *TPCBox) bool {
	_cret := C.tpcbox_hast(box._inner)
	return bool(_cret)
}


// TpcboxGeodetic wraps MEOS C function tpcbox_geodetic.
func TpcboxGeodetic(box *TPCBox) bool {
	_cret := C.tpcbox_geodetic(box._inner)
	return bool(_cret)
}


// TpcboxXmin wraps MEOS C function tpcbox_xmin.
func TpcboxXmin(box *TPCBox) (bool, float64) {
	var _out_result C.double
	_cret := C.tpcbox_xmin(box._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// TpcboxXmax wraps MEOS C function tpcbox_xmax.
func TpcboxXmax(box *TPCBox) (bool, float64) {
	var _out_result C.double
	_cret := C.tpcbox_xmax(box._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// TpcboxYmin wraps MEOS C function tpcbox_ymin.
func TpcboxYmin(box *TPCBox) (bool, float64) {
	var _out_result C.double
	_cret := C.tpcbox_ymin(box._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// TpcboxYmax wraps MEOS C function tpcbox_ymax.
func TpcboxYmax(box *TPCBox) (bool, float64) {
	var _out_result C.double
	_cret := C.tpcbox_ymax(box._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// TpcboxZmin wraps MEOS C function tpcbox_zmin.
func TpcboxZmin(box *TPCBox) (bool, float64) {
	var _out_result C.double
	_cret := C.tpcbox_zmin(box._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// TpcboxZmax wraps MEOS C function tpcbox_zmax.
func TpcboxZmax(box *TPCBox) (bool, float64) {
	var _out_result C.double
	_cret := C.tpcbox_zmax(box._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// TpcboxTmin wraps MEOS C function tpcbox_tmin.
func TpcboxTmin(box *TPCBox) (bool, int64) {
	var _out_result C.TimestampTz
	_cret := C.tpcbox_tmin(box._inner, &_out_result)
	return bool(_cret), int64(_out_result)
}


// TpcboxTminInc wraps MEOS C function tpcbox_tmin_inc.
func TpcboxTminInc(box *TPCBox) (bool, bool) {
	var _out_result C.bool
	_cret := C.tpcbox_tmin_inc(box._inner, &_out_result)
	return bool(_cret), bool(_out_result)
}


// TpcboxTmax wraps MEOS C function tpcbox_tmax.
func TpcboxTmax(box *TPCBox) (bool, int64) {
	var _out_result C.TimestampTz
	_cret := C.tpcbox_tmax(box._inner, &_out_result)
	return bool(_cret), int64(_out_result)
}


// TpcboxTmaxInc wraps MEOS C function tpcbox_tmax_inc.
func TpcboxTmaxInc(box *TPCBox) (bool, bool) {
	var _out_result C.bool
	_cret := C.tpcbox_tmax_inc(box._inner, &_out_result)
	return bool(_cret), bool(_out_result)
}


// TpcboxSRID wraps MEOS C function tpcbox_srid.
func TpcboxSRID(box *TPCBox) int32 {
	_cret := C.tpcbox_srid(box._inner)
	return int32(_cret)
}


// TpcboxPcid wraps MEOS C function tpcbox_pcid.
func TpcboxPcid(box *TPCBox) uint32 {
	_cret := C.tpcbox_pcid(box._inner)
	return uint32(_cret)
}


// TpcboxToSTBOX wraps MEOS C function tpcbox_to_stbox.
func TpcboxToSTBOX(box *TPCBox) *STBox {
	_cret := C.tpcbox_to_stbox(box._inner)
	return &STBox{_inner: _cret}
}


// TpcboxExpand wraps MEOS C function tpcbox_expand.
func TpcboxExpand(box1 *TPCBox, box2 *TPCBox) {
	C.tpcbox_expand(box1._inner, box2._inner)
}


// TpcboxRound wraps MEOS C function tpcbox_round.
func TpcboxRound(box *TPCBox, maxdd int) *TPCBox {
	_cret := C.tpcbox_round(box._inner, C.int(maxdd))
	return &TPCBox{_inner: _cret}
}


// TpcboxSetSRID wraps MEOS C function tpcbox_set_srid.
func TpcboxSetSRID(box *TPCBox, srid int32) *TPCBox {
	_cret := C.tpcbox_set_srid(box._inner, C.int32_t(srid))
	return &TPCBox{_inner: _cret}
}


// UnionTpcboxTpcbox wraps MEOS C function union_tpcbox_tpcbox.
func UnionTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox, strict bool) *TPCBox {
	_cret := C.gunion_tpcbox_tpcbox(box1._inner, box2._inner, C.bool(strict))
	return &TPCBox{_inner: _cret}
}


// InterTpcboxTpcbox wraps MEOS C function inter_tpcbox_tpcbox.
func InterTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (bool, *TPCBox) {
	var _out_result C.TPCBox
	_cret := C.inter_tpcbox_tpcbox(box1._inner, box2._inner, &_out_result)
	return bool(_cret), &TPCBox{_inner: &_out_result}
}


// IntersectionTpcboxTpcbox wraps MEOS C function intersection_tpcbox_tpcbox.
func IntersectionTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) *TPCBox {
	_cret := C.intersection_tpcbox_tpcbox(box1._inner, box2._inner)
	return &TPCBox{_inner: _cret}
}


// ContainsTpcboxTpcbox wraps MEOS C function contains_tpcbox_tpcbox.
func ContainsTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.contains_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// ContainedTpcboxTpcbox wraps MEOS C function contained_tpcbox_tpcbox.
func ContainedTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.contained_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverlapsTpcboxTpcbox wraps MEOS C function overlaps_tpcbox_tpcbox.
func OverlapsTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.overlaps_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// SameTpcboxTpcbox wraps MEOS C function same_tpcbox_tpcbox.
func SameTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.same_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// AdjacentTpcboxTpcbox wraps MEOS C function adjacent_tpcbox_tpcbox.
func AdjacentTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.adjacent_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// TpcboxCmp wraps MEOS C function tpcbox_cmp.
func TpcboxCmp(box1 *TPCBox, box2 *TPCBox) int {
	_cret := C.tpcbox_cmp(box1._inner, box2._inner)
	return int(_cret)
}


// TpcboxEq wraps MEOS C function tpcbox_eq.
func TpcboxEq(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.tpcbox_eq(box1._inner, box2._inner)
	return bool(_cret)
}


// TpcboxNe wraps MEOS C function tpcbox_ne.
func TpcboxNe(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.tpcbox_ne(box1._inner, box2._inner)
	return bool(_cret)
}


// TpcboxLt wraps MEOS C function tpcbox_lt.
func TpcboxLt(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.tpcbox_lt(box1._inner, box2._inner)
	return bool(_cret)
}


// TpcboxLe wraps MEOS C function tpcbox_le.
func TpcboxLe(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.tpcbox_le(box1._inner, box2._inner)
	return bool(_cret)
}


// TpcboxGt wraps MEOS C function tpcbox_gt.
func TpcboxGt(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.tpcbox_gt(box1._inner, box2._inner)
	return bool(_cret)
}


// TpcboxGe wraps MEOS C function tpcbox_ge.
func TpcboxGe(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.tpcbox_ge(box1._inner, box2._inner)
	return bool(_cret)
}


// LeftTpcboxTpcbox wraps MEOS C function left_tpcbox_tpcbox.
func LeftTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.left_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverleftTpcboxTpcbox wraps MEOS C function overleft_tpcbox_tpcbox.
func OverleftTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.overleft_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// RightTpcboxTpcbox wraps MEOS C function right_tpcbox_tpcbox.
func RightTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.right_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverrightTpcboxTpcbox wraps MEOS C function overright_tpcbox_tpcbox.
func OverrightTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.overright_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// BelowTpcboxTpcbox wraps MEOS C function below_tpcbox_tpcbox.
func BelowTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.below_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverbelowTpcboxTpcbox wraps MEOS C function overbelow_tpcbox_tpcbox.
func OverbelowTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.overbelow_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// AboveTpcboxTpcbox wraps MEOS C function above_tpcbox_tpcbox.
func AboveTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.above_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OveraboveTpcboxTpcbox wraps MEOS C function overabove_tpcbox_tpcbox.
func OveraboveTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.overabove_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// FrontTpcboxTpcbox wraps MEOS C function front_tpcbox_tpcbox.
func FrontTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.front_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverfrontTpcboxTpcbox wraps MEOS C function overfront_tpcbox_tpcbox.
func OverfrontTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.overfront_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// BackTpcboxTpcbox wraps MEOS C function back_tpcbox_tpcbox.
func BackTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.back_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverbackTpcboxTpcbox wraps MEOS C function overback_tpcbox_tpcbox.
func OverbackTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.overback_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// BeforeTpcboxTpcbox wraps MEOS C function before_tpcbox_tpcbox.
func BeforeTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.before_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverbeforeTpcboxTpcbox wraps MEOS C function overbefore_tpcbox_tpcbox.
func OverbeforeTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.overbefore_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// AfterTpcboxTpcbox wraps MEOS C function after_tpcbox_tpcbox.
func AfterTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.after_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverafterTpcboxTpcbox wraps MEOS C function overafter_tpcbox_tpcbox.
func OverafterTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.overafter_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// EnsureSamePcidTpcbox wraps MEOS C function ensure_same_pcid_tpcbox.
func EnsureSamePcidTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	_cret := C.ensure_same_pcid_tpcbox(box1._inner, box2._inner)
	return bool(_cret)
}


// TpointcloudinstMake wraps MEOS C function tpointcloudinst_make.
func TpointcloudinstMake(pt *Pcpoint, t int64) *TInstant {
	_cret := C.tpointcloudinst_make(pt._inner, C.TimestampTz(t))
	return &TInstant{_inner: _cret}
}


// EintersectsTpcpointGeo wraps MEOS C function eintersects_tpcpoint_geo.
func EintersectsTpcpointGeo(temp *Temporal, gs *Geom) bool {
	_cret := C.eintersects_tpcpoint_geo(temp._inner, gs._inner)
	return bool(_cret)
}


// NadTpcpointGeo wraps MEOS C function nad_tpcpoint_geo.
func NadTpcpointGeo(temp *Temporal, gs *Geom) float64 {
	_cret := C.nad_tpcpoint_geo(temp._inner, gs._inner)
	return float64(_cret)
}

