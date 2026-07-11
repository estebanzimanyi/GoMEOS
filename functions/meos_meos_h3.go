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

// H3indexIn wraps MEOS C function h3index_in.
func H3indexIn(str string) uint64 {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.h3index_in(_c_str)
	return uint64(_cret)
}


// H3indexOut wraps MEOS C function h3index_out.
func H3indexOut(cell uint64) string {
	_cret := C.h3index_out(C.uint64_t(cell))
	return C.GoString(_cret)
}


// H3indexFromWKB wraps MEOS C function h3index_from_wkb.
func H3indexFromWKB(wkb unsafe.Pointer, size uint) uint64 {
	_cret := C.h3index_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	return uint64(_cret)
}


// H3indexFromHexwkb wraps MEOS C function h3index_from_hexwkb.
func H3indexFromHexwkb(hexwkb string) uint64 {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	_cret := C.h3index_from_hexwkb(_c_hexwkb)
	return uint64(_cret)
}


// H3indexAsWKB wraps MEOS C function h3index_as_wkb.
func H3indexAsWKB(cell uint64, variant uint8, size_out unsafe.Pointer) unsafe.Pointer {
	_cret := C.h3index_as_wkb(C.uint64_t(cell), C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return unsafe.Pointer(_cret)
}


// H3indexAsHexwkb wraps MEOS C function h3index_as_hexwkb.
func H3indexAsHexwkb(cell uint64, variant uint8, size_out unsafe.Pointer) string {
	_cret := C.h3index_as_hexwkb(C.uint64_t(cell), C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return C.GoString(_cret)
}


// H3indexEq wraps MEOS C function h3index_eq.
func H3indexEq(a uint64, b uint64) bool {
	_cret := C.h3index_eq(C.uint64_t(a), C.uint64_t(b))
	return bool(_cret)
}


// H3indexNe wraps MEOS C function h3index_ne.
func H3indexNe(a uint64, b uint64) bool {
	_cret := C.h3index_ne(C.uint64_t(a), C.uint64_t(b))
	return bool(_cret)
}


// H3indexLt wraps MEOS C function h3index_lt.
func H3indexLt(a uint64, b uint64) bool {
	_cret := C.h3index_lt(C.uint64_t(a), C.uint64_t(b))
	return bool(_cret)
}


// H3indexLe wraps MEOS C function h3index_le.
func H3indexLe(a uint64, b uint64) bool {
	_cret := C.h3index_le(C.uint64_t(a), C.uint64_t(b))
	return bool(_cret)
}


// H3indexGt wraps MEOS C function h3index_gt.
func H3indexGt(a uint64, b uint64) bool {
	_cret := C.h3index_gt(C.uint64_t(a), C.uint64_t(b))
	return bool(_cret)
}


// H3indexGe wraps MEOS C function h3index_ge.
func H3indexGe(a uint64, b uint64) bool {
	_cret := C.h3index_ge(C.uint64_t(a), C.uint64_t(b))
	return bool(_cret)
}


// H3indexCmp wraps MEOS C function h3index_cmp.
func H3indexCmp(a uint64, b uint64) int {
	_cret := C.h3index_cmp(C.uint64_t(a), C.uint64_t(b))
	return int(_cret)
}


// H3indexHash wraps MEOS C function h3index_hash.
func H3indexHash(cell uint64) uint32 {
	_cret := C.h3index_hash(C.uint64_t(cell))
	return uint32(_cret)
}


// H3GridDisk wraps MEOS C function h3_grid_disk.
func H3GridDisk(origin uint64, k int) *Set {
	_cret := C.h3_grid_disk(C.uint64_t(origin), C.int(k))
	return &Set{_inner: _cret}
}


// H3CellToChildren wraps MEOS C function h3_cell_to_children.
func H3CellToChildren(origin uint64, childRes int) *Set {
	_cret := C.h3_cell_to_children(C.uint64_t(origin), C.int(childRes))
	return &Set{_inner: _cret}
}


// H3CompactCells wraps MEOS C function h3_compact_cells.
func H3CompactCells(cells *Set) *Set {
	_cret := C.h3_compact_cells(cells._inner)
	return &Set{_inner: _cret}
}


// H3UncompactCells wraps MEOS C function h3_uncompact_cells.
func H3UncompactCells(cells *Set, res int) *Set {
	_cret := C.h3_uncompact_cells(cells._inner, C.int(res))
	return &Set{_inner: _cret}
}


// Th3indexIn wraps MEOS C function th3index_in.
func Th3indexIn(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.th3index_in(_c_str)
	return &Temporal{_inner: _cret}
}


// Th3indexinstIn wraps MEOS C function th3indexinst_in.
func Th3indexinstIn(str string) *TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.th3indexinst_in(_c_str)
	return &TInstant{_inner: _cret}
}


// Th3indexseqIn wraps MEOS C function th3indexseq_in.
func Th3indexseqIn(str string, interp Interpolation) *TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.th3indexseq_in(_c_str, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// Th3indexseqsetIn wraps MEOS C function th3indexseqset_in.
func Th3indexseqsetIn(str string) *TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.th3indexseqset_in(_c_str)
	return &TSequenceSet{_inner: _cret}
}


// Th3indexMake wraps MEOS C function th3index_make.
func Th3indexMake(value uint64, t int64) *Temporal {
	_cret := C.th3index_make(C.uint64_t(value), C.TimestampTz(t))
	return &Temporal{_inner: _cret}
}


// Th3indexinstMake wraps MEOS C function th3indexinst_make.
func Th3indexinstMake(value uint64, t int64) *TInstant {
	_cret := C.th3indexinst_make(C.uint64_t(value), C.TimestampTz(t))
	return &TInstant{_inner: _cret}
}


// Th3indexseqMake wraps MEOS C function th3indexseq_make.
func Th3indexseqMake(values unsafe.Pointer, times unsafe.Pointer, count int, lower_inc bool, upper_inc bool) *TSequence {
	_cret := C.th3indexseq_make((*C.uint64_t)(unsafe.Pointer(values)), (*C.TimestampTz)(unsafe.Pointer(times)), C.int(count), C.bool(lower_inc), C.bool(upper_inc))
	return &TSequence{_inner: _cret}
}


// Th3indexseqsetMake wraps MEOS C function th3indexseqset_make.
func Th3indexseqsetMake(sequences unsafe.Pointer, count int) *TSequenceSet {
	_cret := C.th3indexseqset_make((**C.TSequence)(unsafe.Pointer(sequences)), C.int(count))
	return &TSequenceSet{_inner: _cret}
}


// Th3indexStartValue wraps MEOS C function th3index_start_value.
func Th3indexStartValue(temp *Temporal) uint64 {
	_cret := C.th3index_start_value(temp._inner)
	return uint64(_cret)
}


// Th3indexEndValue wraps MEOS C function th3index_end_value.
func Th3indexEndValue(temp *Temporal) uint64 {
	_cret := C.th3index_end_value(temp._inner)
	return uint64(_cret)
}


// Th3indexValueN wraps MEOS C function th3index_value_n.
func Th3indexValueN(temp *Temporal, n int) (bool, uint64) {
	var _out_result C.uint64_t
	_cret := C.th3index_value_n(temp._inner, C.int(n), &_out_result)
	return bool(_cret), uint64(_out_result)
}


// Th3indexValues wraps MEOS C function th3index_values.
func Th3indexValues(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.th3index_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// Th3indexValueAtTimestamptz wraps MEOS C function th3index_value_at_timestamptz.
func Th3indexValueAtTimestamptz(temp *Temporal, t int64, strict bool) (bool, uint64) {
	var _out_result C.uint64_t
	_cret := C.th3index_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_result)
	return bool(_cret), uint64(_out_result)
}


// TbigintToTh3index wraps MEOS C function tbigint_to_th3index.
func TbigintToTh3index(temp *Temporal) *Temporal {
	_cret := C.tbigint_to_th3index(temp._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexToTbigint wraps MEOS C function th3index_to_tbigint.
func Th3indexToTbigint(temp *Temporal) *Temporal {
	_cret := C.th3index_to_tbigint(temp._inner)
	return &Temporal{_inner: _cret}
}


// EverEqH3indexTh3index wraps MEOS C function ever_eq_h3index_th3index.
func EverEqH3indexTh3index(cell uint64, temp *Temporal) int {
	_cret := C.ever_eq_h3index_th3index(C.uint64_t(cell), temp._inner)
	return int(_cret)
}


// EverEqTh3indexH3index wraps MEOS C function ever_eq_th3index_h3index.
func EverEqTh3indexH3index(temp *Temporal, cell uint64) int {
	_cret := C.ever_eq_th3index_h3index(temp._inner, C.uint64_t(cell))
	return int(_cret)
}


// EverNeH3indexTh3index wraps MEOS C function ever_ne_h3index_th3index.
func EverNeH3indexTh3index(cell uint64, temp *Temporal) int {
	_cret := C.ever_ne_h3index_th3index(C.uint64_t(cell), temp._inner)
	return int(_cret)
}


// EverNeTh3indexH3index wraps MEOS C function ever_ne_th3index_h3index.
func EverNeTh3indexH3index(temp *Temporal, cell uint64) int {
	_cret := C.ever_ne_th3index_h3index(temp._inner, C.uint64_t(cell))
	return int(_cret)
}


// AlwaysEqH3indexTh3index wraps MEOS C function always_eq_h3index_th3index.
func AlwaysEqH3indexTh3index(cell uint64, temp *Temporal) int {
	_cret := C.always_eq_h3index_th3index(C.uint64_t(cell), temp._inner)
	return int(_cret)
}


// AlwaysEqTh3indexH3index wraps MEOS C function always_eq_th3index_h3index.
func AlwaysEqTh3indexH3index(temp *Temporal, cell uint64) int {
	_cret := C.always_eq_th3index_h3index(temp._inner, C.uint64_t(cell))
	return int(_cret)
}


// AlwaysNeH3indexTh3index wraps MEOS C function always_ne_h3index_th3index.
func AlwaysNeH3indexTh3index(cell uint64, temp *Temporal) int {
	_cret := C.always_ne_h3index_th3index(C.uint64_t(cell), temp._inner)
	return int(_cret)
}


// AlwaysNeTh3indexH3index wraps MEOS C function always_ne_th3index_h3index.
func AlwaysNeTh3indexH3index(temp *Temporal, cell uint64) int {
	_cret := C.always_ne_th3index_h3index(temp._inner, C.uint64_t(cell))
	return int(_cret)
}


// EverEqTh3indexTh3index wraps MEOS C function ever_eq_th3index_th3index.
func EverEqTh3indexTh3index(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_eq_th3index_th3index(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverNeTh3indexTh3index wraps MEOS C function ever_ne_th3index_th3index.
func EverNeTh3indexTh3index(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_ne_th3index_th3index(temp1._inner, temp2._inner)
	return int(_cret)
}


// AlwaysEqTh3indexTh3index wraps MEOS C function always_eq_th3index_th3index.
func AlwaysEqTh3indexTh3index(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_eq_th3index_th3index(temp1._inner, temp2._inner)
	return int(_cret)
}


// AlwaysNeTh3indexTh3index wraps MEOS C function always_ne_th3index_th3index.
func AlwaysNeTh3indexTh3index(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_ne_th3index_th3index(temp1._inner, temp2._inner)
	return int(_cret)
}


// TeqH3indexTh3index wraps MEOS C function teq_h3index_th3index.
func TeqH3indexTh3index(cell uint64, temp *Temporal) *Temporal {
	_cret := C.teq_h3index_th3index(C.uint64_t(cell), temp._inner)
	return &Temporal{_inner: _cret}
}


// TeqTh3indexH3index wraps MEOS C function teq_th3index_h3index.
func TeqTh3indexH3index(temp *Temporal, cell uint64) *Temporal {
	_cret := C.teq_th3index_h3index(temp._inner, C.uint64_t(cell))
	return &Temporal{_inner: _cret}
}


// TeqTh3indexTh3index wraps MEOS C function teq_th3index_th3index.
func TeqTh3indexTh3index(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.teq_th3index_th3index(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TneH3indexTh3index wraps MEOS C function tne_h3index_th3index.
func TneH3indexTh3index(cell uint64, temp *Temporal) *Temporal {
	_cret := C.tne_h3index_th3index(C.uint64_t(cell), temp._inner)
	return &Temporal{_inner: _cret}
}


// TneTh3indexH3index wraps MEOS C function tne_th3index_h3index.
func TneTh3indexH3index(temp *Temporal, cell uint64) *Temporal {
	_cret := C.tne_th3index_h3index(temp._inner, C.uint64_t(cell))
	return &Temporal{_inner: _cret}
}


// TneTh3indexTh3index wraps MEOS C function tne_th3index_th3index.
func TneTh3indexTh3index(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tne_th3index_th3index(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexGetResolution wraps MEOS C function th3index_get_resolution.
func Th3indexGetResolution(temp *Temporal) *Temporal {
	_cret := C.th3index_get_resolution(temp._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexGetBaseCellNumber wraps MEOS C function th3index_get_base_cell_number.
func Th3indexGetBaseCellNumber(temp *Temporal) *Temporal {
	_cret := C.th3index_get_base_cell_number(temp._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexIsValidCell wraps MEOS C function th3index_is_valid_cell.
func Th3indexIsValidCell(temp *Temporal) *Temporal {
	_cret := C.th3index_is_valid_cell(temp._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexIsResClassIii wraps MEOS C function th3index_is_res_class_iii.
func Th3indexIsResClassIii(temp *Temporal) *Temporal {
	_cret := C.th3index_is_res_class_iii(temp._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexIsPentagon wraps MEOS C function th3index_is_pentagon.
func Th3indexIsPentagon(temp *Temporal) *Temporal {
	_cret := C.th3index_is_pentagon(temp._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexCellToParent wraps MEOS C function th3index_cell_to_parent.
func Th3indexCellToParent(temp *Temporal, resolution int32) *Temporal {
	_cret := C.th3index_cell_to_parent(temp._inner, C.int32(resolution))
	return &Temporal{_inner: _cret}
}


// Th3indexCellToParentNext wraps MEOS C function th3index_cell_to_parent_next.
func Th3indexCellToParentNext(temp *Temporal) *Temporal {
	_cret := C.th3index_cell_to_parent_next(temp._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexCellToCenterChild wraps MEOS C function th3index_cell_to_center_child.
func Th3indexCellToCenterChild(temp *Temporal, resolution int32) *Temporal {
	_cret := C.th3index_cell_to_center_child(temp._inner, C.int32(resolution))
	return &Temporal{_inner: _cret}
}


// Th3indexCellToCenterChildNext wraps MEOS C function th3index_cell_to_center_child_next.
func Th3indexCellToCenterChildNext(temp *Temporal) *Temporal {
	_cret := C.th3index_cell_to_center_child_next(temp._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexCellToChildPos wraps MEOS C function th3index_cell_to_child_pos.
func Th3indexCellToChildPos(temp *Temporal, parent_res int32) *Temporal {
	_cret := C.th3index_cell_to_child_pos(temp._inner, C.int32(parent_res))
	return &Temporal{_inner: _cret}
}


// Th3indexChildPosToCell wraps MEOS C function th3index_child_pos_to_cell.
func Th3indexChildPosToCell(child_pos *Temporal, parent *Temporal, child_res int32) *Temporal {
	_cret := C.th3index_child_pos_to_cell(child_pos._inner, parent._inner, C.int32(child_res))
	return &Temporal{_inner: _cret}
}


// TgeogpointToTh3index wraps MEOS C function tgeogpoint_to_th3index.
func TgeogpointToTh3index(temp *Temporal, resolution int32) *Temporal {
	_cret := C.tgeogpoint_to_th3index(temp._inner, C.int32(resolution))
	return &Temporal{_inner: _cret}
}


// TgeompointToTh3index wraps MEOS C function tgeompoint_to_th3index.
func TgeompointToTh3index(temp *Temporal, resolution int32) *Temporal {
	_cret := C.tgeompoint_to_th3index(temp._inner, C.int32(resolution))
	return &Temporal{_inner: _cret}
}


// Th3indexToTgeogpoint wraps MEOS C function th3index_to_tgeogpoint.
func Th3indexToTgeogpoint(temp *Temporal) *Temporal {
	_cret := C.th3index_to_tgeogpoint(temp._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexToTgeompoint wraps MEOS C function th3index_to_tgeompoint.
func Th3indexToTgeompoint(temp *Temporal) *Temporal {
	_cret := C.th3index_to_tgeompoint(temp._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexCellToBoundary wraps MEOS C function th3index_cell_to_boundary.
func Th3indexCellToBoundary(temp *Temporal) *Temporal {
	_cret := C.th3index_cell_to_boundary(temp._inner)
	return &Temporal{_inner: _cret}
}


// H3GsPointToCell wraps MEOS C function h3_gs_point_to_cell.
func H3GsPointToCell(point *Geom, resolution int32) uint64 {
	_cret := C.h3_gs_point_to_cell(point._inner, C.int32(resolution))
	return uint64(_cret)
}


// GeoToH3indexSet wraps MEOS C function geo_to_h3index_set.
func GeoToH3indexSet(gs *Geom, resolution int32) *Set {
	_cret := C.geo_to_h3index_set(gs._inner, C.int32(resolution))
	return &Set{_inner: _cret}
}


// EverEqH3indexsetTh3index wraps MEOS C function ever_eq_h3indexset_th3index.
func EverEqH3indexsetTh3index(cells *Set, th3idx *Temporal) int {
	_cret := C.ever_eq_h3indexset_th3index(cells._inner, th3idx._inner)
	return int(_cret)
}


// Th3indexAreNeighborCells wraps MEOS C function th3index_are_neighbor_cells.
func Th3indexAreNeighborCells(origin *Temporal, dest *Temporal) *Temporal {
	_cret := C.th3index_are_neighbor_cells(origin._inner, dest._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexCellsToDirectedEdge wraps MEOS C function th3index_cells_to_directed_edge.
func Th3indexCellsToDirectedEdge(origin *Temporal, dest *Temporal) *Temporal {
	_cret := C.th3index_cells_to_directed_edge(origin._inner, dest._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexIsValidDirectedEdge wraps MEOS C function th3index_is_valid_directed_edge.
func Th3indexIsValidDirectedEdge(edge *Temporal) *Temporal {
	_cret := C.th3index_is_valid_directed_edge(edge._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexGetDirectedEdgeOrigin wraps MEOS C function th3index_get_directed_edge_origin.
func Th3indexGetDirectedEdgeOrigin(edge *Temporal) *Temporal {
	_cret := C.th3index_get_directed_edge_origin(edge._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexGetDirectedEdgeDestination wraps MEOS C function th3index_get_directed_edge_destination.
func Th3indexGetDirectedEdgeDestination(edge *Temporal) *Temporal {
	_cret := C.th3index_get_directed_edge_destination(edge._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexDirectedEdgeToBoundary wraps MEOS C function th3index_directed_edge_to_boundary.
func Th3indexDirectedEdgeToBoundary(edge *Temporal) *Temporal {
	_cret := C.th3index_directed_edge_to_boundary(edge._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexCellToVertex wraps MEOS C function th3index_cell_to_vertex.
func Th3indexCellToVertex(temp *Temporal, vertex_num int32) *Temporal {
	_cret := C.th3index_cell_to_vertex(temp._inner, C.int32(vertex_num))
	return &Temporal{_inner: _cret}
}


// Th3indexVertexToLatlng wraps MEOS C function th3index_vertex_to_latlng.
func Th3indexVertexToLatlng(temp *Temporal) *Temporal {
	_cret := C.th3index_vertex_to_latlng(temp._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexIsValidVertex wraps MEOS C function th3index_is_valid_vertex.
func Th3indexIsValidVertex(temp *Temporal) *Temporal {
	_cret := C.th3index_is_valid_vertex(temp._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexGridDistance wraps MEOS C function th3index_grid_distance.
func Th3indexGridDistance(origin *Temporal, dest *Temporal) *Temporal {
	_cret := C.th3index_grid_distance(origin._inner, dest._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexCellToLocalIj wraps MEOS C function th3index_cell_to_local_ij.
func Th3indexCellToLocalIj(origin *Temporal, cell *Temporal) *Temporal {
	_cret := C.th3index_cell_to_local_ij(origin._inner, cell._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexLocalIjToCell wraps MEOS C function th3index_local_ij_to_cell.
func Th3indexLocalIjToCell(origin *Temporal, coord *Temporal) *Temporal {
	_cret := C.th3index_local_ij_to_cell(origin._inner, coord._inner)
	return &Temporal{_inner: _cret}
}


// Th3indexCellArea wraps MEOS C function th3index_cell_area.
func Th3indexCellArea(temp *Temporal, unit string) *Temporal {
	_c_unit := C.CString(unit)
	defer C.free(unsafe.Pointer(_c_unit))
	_cret := C.th3index_cell_area(temp._inner, _c_unit)
	return &Temporal{_inner: _cret}
}


// Th3indexEdgeLength wraps MEOS C function th3index_edge_length.
func Th3indexEdgeLength(temp *Temporal, unit string) *Temporal {
	_c_unit := C.CString(unit)
	defer C.free(unsafe.Pointer(_c_unit))
	_cret := C.th3index_edge_length(temp._inner, _c_unit)
	return &Temporal{_inner: _cret}
}


// TgeogpointGreatCircleDistance wraps MEOS C function tgeogpoint_great_circle_distance.
func TgeogpointGreatCircleDistance(a *Temporal, b *Temporal, unit string) *Temporal {
	_c_unit := C.CString(unit)
	defer C.free(unsafe.Pointer(_c_unit))
	_cret := C.tgeogpoint_great_circle_distance(a._inner, b._inner, _c_unit)
	return &Temporal{_inner: _cret}
}

