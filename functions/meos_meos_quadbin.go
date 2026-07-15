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

// QuadbinIsValidIndex wraps MEOS C function quadbin_is_valid_index.
func QuadbinIsValidIndex(index uint64) bool {
	_cret := C.quadbin_is_valid_index(C.uint64_t(index))
	return bool(_cret)
}


// QuadbinIsValidCell wraps MEOS C function quadbin_is_valid_cell.
func QuadbinIsValidCell(cell uint64) bool {
	_cret := C.quadbin_is_valid_cell(C.uint64_t(cell))
	return bool(_cret)
}


// QuadbinTileToCell wraps MEOS C function quadbin_tile_to_cell.
func QuadbinTileToCell(x uint32, y uint32, z uint32) uint64 {
	_cret := C.quadbin_tile_to_cell(C.uint32_t(x), C.uint32_t(y), C.uint32_t(z))
	return uint64(_cret)
}


// QuadbinCellToTile wraps MEOS C function quadbin_cell_to_tile.
func QuadbinCellToTile(cell uint64, x unsafe.Pointer, y unsafe.Pointer, z unsafe.Pointer) {
	C.quadbin_cell_to_tile(C.uint64_t(cell), (*C.uint32_t)(unsafe.Pointer(x)), (*C.uint32_t)(unsafe.Pointer(y)), (*C.uint32_t)(unsafe.Pointer(z)))
}


// QuadbinGetResolution wraps MEOS C function quadbin_get_resolution.
func QuadbinGetResolution(cell uint64) uint32 {
	_cret := C.quadbin_get_resolution(C.uint64_t(cell))
	return uint32(_cret)
}


// QuadbinCellToParent wraps MEOS C function quadbin_cell_to_parent.
func QuadbinCellToParent(cell uint64, parent_resolution uint32) uint64 {
	_cret := C.quadbin_cell_to_parent(C.uint64_t(cell), C.uint32_t(parent_resolution))
	return uint64(_cret)
}


// QuadbinCellToChildren wraps MEOS C function quadbin_cell_to_children.
func QuadbinCellToChildren(cell uint64, children_resolution uint32, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.quadbin_cell_to_children(C.uint64_t(cell), C.uint32_t(children_resolution), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// QuadbinCellSibling wraps MEOS C function quadbin_cell_sibling.
func QuadbinCellSibling(cell uint64, direction string) uint64 {
	_c_direction := C.CString(direction)
	defer C.free(unsafe.Pointer(_c_direction))
	_cret := C.quadbin_cell_sibling(C.uint64_t(cell), _c_direction)
	return uint64(_cret)
}


// QuadbinKRing wraps MEOS C function quadbin_k_ring.
func QuadbinKRing(cell uint64, k int, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.quadbin_k_ring(C.uint64_t(cell), C.int(k), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// QuadbinPointToCell wraps MEOS C function quadbin_point_to_cell.
func QuadbinPointToCell(longitude float64, latitude float64, resolution uint32) uint64 {
	_cret := C.quadbin_point_to_cell(C.double(longitude), C.double(latitude), C.uint32_t(resolution))
	return uint64(_cret)
}


// QuadbinCellToPoint wraps MEOS C function quadbin_cell_to_point.
func QuadbinCellToPoint(cell uint64, longitude unsafe.Pointer, latitude unsafe.Pointer) {
	C.quadbin_cell_to_point(C.uint64_t(cell), (*C.double)(unsafe.Pointer(longitude)), (*C.double)(unsafe.Pointer(latitude)))
}


// QuadbinCellToBoundingBox wraps MEOS C function quadbin_cell_to_bounding_box.
func QuadbinCellToBoundingBox(cell uint64, xmin unsafe.Pointer, ymin unsafe.Pointer, xmax unsafe.Pointer, ymax unsafe.Pointer) {
	C.quadbin_cell_to_bounding_box(C.uint64_t(cell), (*C.double)(unsafe.Pointer(xmin)), (*C.double)(unsafe.Pointer(ymin)), (*C.double)(unsafe.Pointer(xmax)), (*C.double)(unsafe.Pointer(ymax)))
}


// QuadbinCellArea wraps MEOS C function quadbin_cell_area.
func QuadbinCellArea(cell uint64) float64 {
	_cret := C.quadbin_cell_area(C.uint64_t(cell))
	return float64(_cret)
}


// QuadbinIndexToString wraps MEOS C function quadbin_index_to_string.
func QuadbinIndexToString(index uint64) string {
	_cret := C.quadbin_index_to_string(C.uint64_t(index))
	return C.GoString(_cret)
}


// QuadbinStringToIndex wraps MEOS C function quadbin_string_to_index.
func QuadbinStringToIndex(str string) uint64 {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.quadbin_string_to_index(_c_str)
	return uint64(_cret)
}


// QuadbinCellToQuadkey wraps MEOS C function quadbin_cell_to_quadkey.
func QuadbinCellToQuadkey(cell uint64) string {
	_cret := C.quadbin_cell_to_quadkey(C.uint64_t(cell))
	return C.GoString(_cret)
}


// QuadbinParse wraps MEOS C function quadbin_parse.
func QuadbinParse(str string) uint64 {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.quadbin_parse(_c_str)
	return uint64(_cret)
}


// QuadbinEq wraps MEOS C function quadbin_eq.
func QuadbinEq(a uint64, b uint64) bool {
	_cret := C.quadbin_eq(C.uint64_t(a), C.uint64_t(b))
	return bool(_cret)
}


// QuadbinNe wraps MEOS C function quadbin_ne.
func QuadbinNe(a uint64, b uint64) bool {
	_cret := C.quadbin_ne(C.uint64_t(a), C.uint64_t(b))
	return bool(_cret)
}


// QuadbinLt wraps MEOS C function quadbin_lt.
func QuadbinLt(a uint64, b uint64) bool {
	_cret := C.quadbin_lt(C.uint64_t(a), C.uint64_t(b))
	return bool(_cret)
}


// QuadbinLe wraps MEOS C function quadbin_le.
func QuadbinLe(a uint64, b uint64) bool {
	_cret := C.quadbin_le(C.uint64_t(a), C.uint64_t(b))
	return bool(_cret)
}


// QuadbinGt wraps MEOS C function quadbin_gt.
func QuadbinGt(a uint64, b uint64) bool {
	_cret := C.quadbin_gt(C.uint64_t(a), C.uint64_t(b))
	return bool(_cret)
}


// QuadbinGe wraps MEOS C function quadbin_ge.
func QuadbinGe(a uint64, b uint64) bool {
	_cret := C.quadbin_ge(C.uint64_t(a), C.uint64_t(b))
	return bool(_cret)
}


// QuadbinCmp wraps MEOS C function quadbin_cmp.
func QuadbinCmp(a uint64, b uint64) int {
	_cret := C.quadbin_cmp(C.uint64_t(a), C.uint64_t(b))
	return int(_cret)
}


// QuadbinHash wraps MEOS C function quadbin_hash.
func QuadbinHash(cell uint64) uint32 {
	_cret := C.quadbin_hash(C.uint64_t(cell))
	return uint32(_cret)
}


// QuadbinGridDisk wraps MEOS C function quadbin_grid_disk.
func QuadbinGridDisk(origin uint64, k int) *Set {
	_cret := C.quadbin_grid_disk(C.uint64_t(origin), C.int(k))
	return &Set{_inner: _cret}
}


// QuadbinCellToChildrenSet wraps MEOS C function quadbin_cell_to_children_set.
func QuadbinCellToChildrenSet(origin uint64, children_resolution int) *Set {
	_cret := C.quadbin_cell_to_children_set(C.uint64_t(origin), C.int(children_resolution))
	return &Set{_inner: _cret}
}


// TquadbinIn wraps MEOS C function tquadbin_in.
func TquadbinIn(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tquadbin_in(_c_str)
	return &Temporal{_inner: _cret}
}


// TquadbininstIn wraps MEOS C function tquadbininst_in.
func TquadbininstIn(str string) *TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tquadbininst_in(_c_str)
	return &TInstant{_inner: _cret}
}


// TquadbinseqIn wraps MEOS C function tquadbinseq_in.
func TquadbinseqIn(str string, interp Interpolation) *TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tquadbinseq_in(_c_str, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TquadbinseqsetIn wraps MEOS C function tquadbinseqset_in.
func TquadbinseqsetIn(str string) *TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tquadbinseqset_in(_c_str)
	return &TSequenceSet{_inner: _cret}
}


// TquadbinMake wraps MEOS C function tquadbin_make.
func TquadbinMake(value uint64, t int64) *Temporal {
	_cret := C.tquadbin_make(C.uint64_t(value), C.TimestampTz(t))
	return &Temporal{_inner: _cret}
}


// TquadbininstMake wraps MEOS C function tquadbininst_make.
func TquadbininstMake(value uint64, t int64) *TInstant {
	_cret := C.tquadbininst_make(C.uint64_t(value), C.TimestampTz(t))
	return &TInstant{_inner: _cret}
}


// TquadbinseqMake wraps MEOS C function tquadbinseq_make.
func TquadbinseqMake(values unsafe.Pointer, times unsafe.Pointer, count int, lower_inc bool, upper_inc bool) *TSequence {
	_cret := C.tquadbinseq_make((*C.uint64_t)(unsafe.Pointer(values)), (*C.TimestampTz)(unsafe.Pointer(times)), C.int(count), C.bool(lower_inc), C.bool(upper_inc))
	return &TSequence{_inner: _cret}
}


// TquadbinseqsetMake wraps MEOS C function tquadbinseqset_make.
func TquadbinseqsetMake(sequences unsafe.Pointer, count int) *TSequenceSet {
	_cret := C.tquadbinseqset_make((**C.TSequence)(unsafe.Pointer(sequences)), C.int(count))
	return &TSequenceSet{_inner: _cret}
}


// TquadbinStartValue wraps MEOS C function tquadbin_start_value.
func TquadbinStartValue(temp *Temporal) uint64 {
	_cret := C.tquadbin_start_value(temp._inner)
	return uint64(_cret)
}


// TquadbinEndValue wraps MEOS C function tquadbin_end_value.
func TquadbinEndValue(temp *Temporal) uint64 {
	_cret := C.tquadbin_end_value(temp._inner)
	return uint64(_cret)
}


// TquadbinValueN wraps MEOS C function tquadbin_value_n.
func TquadbinValueN(temp *Temporal, n int) (bool, uint64) {
	var _out_result C.uint64_t
	_cret := C.tquadbin_value_n(temp._inner, C.int(n), &_out_result)
	return bool(_cret), uint64(_out_result)
}


// TquadbinValues wraps MEOS C function tquadbin_values.
func TquadbinValues(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tquadbin_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TquadbinValueAtTimestamptz wraps MEOS C function tquadbin_value_at_timestamptz.
func TquadbinValueAtTimestamptz(temp *Temporal, t int64, strict bool) (bool, uint64) {
	var _out_result C.uint64_t
	_cret := C.tquadbin_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_result)
	return bool(_cret), uint64(_out_result)
}


// TbigintToTquadbin wraps MEOS C function tbigint_to_tquadbin.
func TbigintToTquadbin(temp *Temporal) *Temporal {
	_cret := C.tbigint_to_tquadbin(temp._inner)
	return &Temporal{_inner: _cret}
}


// TquadbinToTbigint wraps MEOS C function tquadbin_to_tbigint.
func TquadbinToTbigint(temp *Temporal) *Temporal {
	_cret := C.tquadbin_to_tbigint(temp._inner)
	return &Temporal{_inner: _cret}
}


// EverEqQuadbinTquadbin wraps MEOS C function ever_eq_quadbin_tquadbin.
func EverEqQuadbinTquadbin(cell uint64, temp *Temporal) int {
	_cret := C.ever_eq_quadbin_tquadbin(C.uint64_t(cell), temp._inner)
	return int(_cret)
}


// EverEqTquadbinQuadbin wraps MEOS C function ever_eq_tquadbin_quadbin.
func EverEqTquadbinQuadbin(temp *Temporal, cell uint64) int {
	_cret := C.ever_eq_tquadbin_quadbin(temp._inner, C.uint64_t(cell))
	return int(_cret)
}


// EverNeQuadbinTquadbin wraps MEOS C function ever_ne_quadbin_tquadbin.
func EverNeQuadbinTquadbin(cell uint64, temp *Temporal) int {
	_cret := C.ever_ne_quadbin_tquadbin(C.uint64_t(cell), temp._inner)
	return int(_cret)
}


// EverNeTquadbinQuadbin wraps MEOS C function ever_ne_tquadbin_quadbin.
func EverNeTquadbinQuadbin(temp *Temporal, cell uint64) int {
	_cret := C.ever_ne_tquadbin_quadbin(temp._inner, C.uint64_t(cell))
	return int(_cret)
}


// AlwaysEqQuadbinTquadbin wraps MEOS C function always_eq_quadbin_tquadbin.
func AlwaysEqQuadbinTquadbin(cell uint64, temp *Temporal) int {
	_cret := C.always_eq_quadbin_tquadbin(C.uint64_t(cell), temp._inner)
	return int(_cret)
}


// AlwaysEqTquadbinQuadbin wraps MEOS C function always_eq_tquadbin_quadbin.
func AlwaysEqTquadbinQuadbin(temp *Temporal, cell uint64) int {
	_cret := C.always_eq_tquadbin_quadbin(temp._inner, C.uint64_t(cell))
	return int(_cret)
}


// AlwaysNeQuadbinTquadbin wraps MEOS C function always_ne_quadbin_tquadbin.
func AlwaysNeQuadbinTquadbin(cell uint64, temp *Temporal) int {
	_cret := C.always_ne_quadbin_tquadbin(C.uint64_t(cell), temp._inner)
	return int(_cret)
}


// AlwaysNeTquadbinQuadbin wraps MEOS C function always_ne_tquadbin_quadbin.
func AlwaysNeTquadbinQuadbin(temp *Temporal, cell uint64) int {
	_cret := C.always_ne_tquadbin_quadbin(temp._inner, C.uint64_t(cell))
	return int(_cret)
}


// EverEqTquadbinTquadbin wraps MEOS C function ever_eq_tquadbin_tquadbin.
func EverEqTquadbinTquadbin(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_eq_tquadbin_tquadbin(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverNeTquadbinTquadbin wraps MEOS C function ever_ne_tquadbin_tquadbin.
func EverNeTquadbinTquadbin(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_ne_tquadbin_tquadbin(temp1._inner, temp2._inner)
	return int(_cret)
}


// AlwaysEqTquadbinTquadbin wraps MEOS C function always_eq_tquadbin_tquadbin.
func AlwaysEqTquadbinTquadbin(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_eq_tquadbin_tquadbin(temp1._inner, temp2._inner)
	return int(_cret)
}


// AlwaysNeTquadbinTquadbin wraps MEOS C function always_ne_tquadbin_tquadbin.
func AlwaysNeTquadbinTquadbin(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_ne_tquadbin_tquadbin(temp1._inner, temp2._inner)
	return int(_cret)
}


// TeqQuadbinTquadbin wraps MEOS C function teq_quadbin_tquadbin.
func TeqQuadbinTquadbin(cell uint64, temp *Temporal) *Temporal {
	_cret := C.teq_quadbin_tquadbin(C.uint64_t(cell), temp._inner)
	return &Temporal{_inner: _cret}
}


// TeqTquadbinQuadbin wraps MEOS C function teq_tquadbin_quadbin.
func TeqTquadbinQuadbin(temp *Temporal, cell uint64) *Temporal {
	_cret := C.teq_tquadbin_quadbin(temp._inner, C.uint64_t(cell))
	return &Temporal{_inner: _cret}
}


// TeqTquadbinTquadbin wraps MEOS C function teq_tquadbin_tquadbin.
func TeqTquadbinTquadbin(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.teq_tquadbin_tquadbin(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TneQuadbinTquadbin wraps MEOS C function tne_quadbin_tquadbin.
func TneQuadbinTquadbin(cell uint64, temp *Temporal) *Temporal {
	_cret := C.tne_quadbin_tquadbin(C.uint64_t(cell), temp._inner)
	return &Temporal{_inner: _cret}
}


// TneTquadbinQuadbin wraps MEOS C function tne_tquadbin_quadbin.
func TneTquadbinQuadbin(temp *Temporal, cell uint64) *Temporal {
	_cret := C.tne_tquadbin_quadbin(temp._inner, C.uint64_t(cell))
	return &Temporal{_inner: _cret}
}


// TneTquadbinTquadbin wraps MEOS C function tne_tquadbin_tquadbin.
func TneTquadbinTquadbin(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tne_tquadbin_tquadbin(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TquadbinCellToQuadkey wraps MEOS C function tquadbin_cell_to_quadkey.
func TquadbinCellToQuadkey(temp *Temporal) *Temporal {
	_cret := C.tquadbin_cell_to_quadkey(temp._inner)
	return &Temporal{_inner: _cret}
}

