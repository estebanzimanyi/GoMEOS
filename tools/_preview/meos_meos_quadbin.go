package generated

// #include <stddef.h>
import "C"
import (
	"unsafe"

	"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

// QuadbinIsValidIndex wraps MEOS C function quadbin_is_valid_index.
func QuadbinIsValidIndex(index uint64) bool {
	res := C.quadbin_is_valid_index(C.Quadbin(index))
	return bool(res)
}


// QuadbinIsValidCell wraps MEOS C function quadbin_is_valid_cell.
func QuadbinIsValidCell(cell uint64) bool {
	res := C.quadbin_is_valid_cell(C.Quadbin(cell))
	return bool(res)
}


// QuadbinTileToCell wraps MEOS C function quadbin_tile_to_cell.
func QuadbinTileToCell(x uint32, y uint32, z uint32) uint64 {
	res := C.quadbin_tile_to_cell(C.uint32_t(x), C.uint32_t(y), C.uint32_t(z))
	return uint64(res)
}


// TODO quadbin_cell_to_tile: unsupported param uint32_t *
// func QuadbinCellToTile(...) { /* not yet handled by codegen */ }


// QuadbinGetResolution wraps MEOS C function quadbin_get_resolution.
func QuadbinGetResolution(cell uint64) uint32 {
	res := C.quadbin_get_resolution(C.Quadbin(cell))
	return uint32(res)
}


// QuadbinCellToParent wraps MEOS C function quadbin_cell_to_parent.
func QuadbinCellToParent(cell uint64, parent_resolution uint32) uint64 {
	res := C.quadbin_cell_to_parent(C.Quadbin(cell), C.uint32_t(parent_resolution))
	return uint64(res)
}


// QuadbinCellToChildren wraps MEOS C function quadbin_cell_to_children.
func QuadbinCellToChildren(cell uint64, children_resolution uint32) []uint64 {
	var _out_count C.int
	res := C.quadbin_cell_to_children(C.Quadbin(cell), C.uint32_t(children_resolution), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((*C.Quadbin)(unsafe.Pointer(res)), _n)
	_out := make([]uint64, _n)
	for _i, _e := range _slice {
		_out[_i] = uint64(_e)
	}
	return _out
}


// QuadbinCellSibling wraps MEOS C function quadbin_cell_sibling.
func QuadbinCellSibling(cell uint64, direction string) uint64 {
	_c_direction := C.CString(direction)
	defer C.free(unsafe.Pointer(_c_direction))
	res := C.quadbin_cell_sibling(C.Quadbin(cell), _c_direction)
	return uint64(res)
}


// QuadbinKRing wraps MEOS C function quadbin_k_ring.
func QuadbinKRing(cell uint64, k int) []uint64 {
	var _out_count C.int
	res := C.quadbin_k_ring(C.Quadbin(cell), C.int(k), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((*C.Quadbin)(unsafe.Pointer(res)), _n)
	_out := make([]uint64, _n)
	for _i, _e := range _slice {
		_out[_i] = uint64(_e)
	}
	return _out
}


// QuadbinPointToCell wraps MEOS C function quadbin_point_to_cell.
func QuadbinPointToCell(longitude float64, latitude float64, resolution uint32) uint64 {
	res := C.quadbin_point_to_cell(C.double(longitude), C.double(latitude), C.uint32_t(resolution))
	return uint64(res)
}


// TODO quadbin_cell_to_point: unsupported param double *
// func QuadbinCellToPoint(...) { /* not yet handled by codegen */ }


// TODO quadbin_cell_to_bounding_box: unsupported param double *
// func QuadbinCellToBoundingBox(...) { /* not yet handled by codegen */ }


// QuadbinCellArea wraps MEOS C function quadbin_cell_area.
func QuadbinCellArea(cell uint64) float64 {
	res := C.quadbin_cell_area(C.Quadbin(cell))
	return float64(res)
}


// QuadbinIndexToString wraps MEOS C function quadbin_index_to_string.
func QuadbinIndexToString(index uint64) string {
	res := C.quadbin_index_to_string(C.Quadbin(index))
	return C.GoString(res)
}


// QuadbinStringToIndex wraps MEOS C function quadbin_string_to_index.
func QuadbinStringToIndex(str string) uint64 {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.quadbin_string_to_index(_c_str)
	return uint64(res)
}


// QuadbinCellToQuadkey wraps MEOS C function quadbin_cell_to_quadkey.
func QuadbinCellToQuadkey(cell uint64) string {
	res := C.quadbin_cell_to_quadkey(C.Quadbin(cell))
	return C.GoString(res)
}


// QuadbinParse wraps MEOS C function quadbin_parse.
func QuadbinParse(str string) uint64 {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.quadbin_parse(_c_str)
	return uint64(res)
}


// QuadbinEq wraps MEOS C function quadbin_eq.
func QuadbinEq(a uint64, b uint64) bool {
	res := C.quadbin_eq(C.Quadbin(a), C.Quadbin(b))
	return bool(res)
}


// QuadbinNe wraps MEOS C function quadbin_ne.
func QuadbinNe(a uint64, b uint64) bool {
	res := C.quadbin_ne(C.Quadbin(a), C.Quadbin(b))
	return bool(res)
}


// QuadbinLt wraps MEOS C function quadbin_lt.
func QuadbinLt(a uint64, b uint64) bool {
	res := C.quadbin_lt(C.Quadbin(a), C.Quadbin(b))
	return bool(res)
}


// QuadbinLe wraps MEOS C function quadbin_le.
func QuadbinLe(a uint64, b uint64) bool {
	res := C.quadbin_le(C.Quadbin(a), C.Quadbin(b))
	return bool(res)
}


// QuadbinGt wraps MEOS C function quadbin_gt.
func QuadbinGt(a uint64, b uint64) bool {
	res := C.quadbin_gt(C.Quadbin(a), C.Quadbin(b))
	return bool(res)
}


// QuadbinGe wraps MEOS C function quadbin_ge.
func QuadbinGe(a uint64, b uint64) bool {
	res := C.quadbin_ge(C.Quadbin(a), C.Quadbin(b))
	return bool(res)
}


// QuadbinCmp wraps MEOS C function quadbin_cmp.
func QuadbinCmp(a uint64, b uint64) int {
	res := C.quadbin_cmp(C.Quadbin(a), C.Quadbin(b))
	return int(res)
}


// QuadbinHash wraps MEOS C function quadbin_hash.
func QuadbinHash(cell uint64) uint32 {
	res := C.quadbin_hash(C.Quadbin(cell))
	return uint32(res)
}


// QuadbinGridDisk wraps MEOS C function quadbin_grid_disk.
func QuadbinGridDisk(origin uint64, k int) *Set {
	res := C.quadbin_grid_disk(C.Quadbin(origin), C.int(k))
	return &Set{_inner: res}
}


// QuadbinCellToChildrenSet wraps MEOS C function quadbin_cell_to_children_set.
func QuadbinCellToChildrenSet(origin uint64, children_resolution int) *Set {
	res := C.quadbin_cell_to_children_set(C.Quadbin(origin), C.int(children_resolution))
	return &Set{_inner: res}
}


// TquadbinIn wraps MEOS C function tquadbin_in.
func TquadbinIn(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tquadbin_in(_c_str)
	return CreateTemporal(res)
}


// TquadbininstIn wraps MEOS C function tquadbininst_in.
func TquadbininstIn(str string) TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tquadbininst_in(_c_str)
	return TInstant{_inner: res}
}


// TquadbinseqIn wraps MEOS C function tquadbinseq_in.
func TquadbinseqIn(str string, interp Interpolation) TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tquadbinseq_in(_c_str, C.interpType(interp))
	return TSequence{_inner: res}
}


// TquadbinseqsetIn wraps MEOS C function tquadbinseqset_in.
func TquadbinseqsetIn(str string) TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tquadbinseqset_in(_c_str)
	return TSequenceSet{_inner: res}
}


// TquadbinMake wraps MEOS C function tquadbin_make.
func TquadbinMake(value uint64, t int64) Temporal {
	res := C.tquadbin_make(C.Quadbin(value), C.TimestampTz(t))
	return CreateTemporal(res)
}


// TquadbininstMake wraps MEOS C function tquadbininst_make.
func TquadbininstMake(value uint64, t int64) TInstant {
	res := C.tquadbininst_make(C.Quadbin(value), C.TimestampTz(t))
	return TInstant{_inner: res}
}


// TODO tquadbinseq_make: unsupported param const Quadbin *
// func TquadbinseqMake(...) { /* not yet handled by codegen */ }


// TquadbinseqsetMake wraps MEOS C function tquadbinseqset_make.
func TquadbinseqsetMake(sequences []TSequence) TSequenceSet {
	_c_sequences := make([]*C.TSequence, len(sequences))
	for _i, _v := range sequences { _c_sequences[_i] = _v._inner }
	res := C.tquadbinseqset_make((**C.TSequence)(unsafe.Pointer(&_c_sequences[0])), C.int(len(sequences)))
	return TSequenceSet{_inner: res}
}


// TquadbinStartValue wraps MEOS C function tquadbin_start_value.
func TquadbinStartValue(temp Temporal) uint64 {
	res := C.tquadbin_start_value(temp.Inner())
	return uint64(res)
}


// TquadbinEndValue wraps MEOS C function tquadbin_end_value.
func TquadbinEndValue(temp Temporal) uint64 {
	res := C.tquadbin_end_value(temp.Inner())
	return uint64(res)
}


// TquadbinValueN wraps MEOS C function tquadbin_value_n.
func TquadbinValueN(temp Temporal, n int) (bool, uint64) {
	var _out_result C.Quadbin
	res := C.tquadbin_value_n(temp.Inner(), C.int(n), &_out_result)
	return bool(res), uint64(_out_result)
}


// TquadbinValues wraps MEOS C function tquadbin_values.
func TquadbinValues(temp Temporal) []uint64 {
	var _out_count C.int
	res := C.tquadbin_values(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((*C.Quadbin)(unsafe.Pointer(res)), _n)
	_out := make([]uint64, _n)
	for _i, _e := range _slice {
		_out[_i] = uint64(_e)
	}
	return _out
}


// TquadbinValueAtTimestamptz wraps MEOS C function tquadbin_value_at_timestamptz.
func TquadbinValueAtTimestamptz(temp Temporal, t int64, strict bool) (bool, uint64) {
	var _out_result C.Quadbin
	res := C.tquadbin_value_at_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(strict), &_out_result)
	return bool(res), uint64(_out_result)
}


// TbigintToTquadbin wraps MEOS C function tbigint_to_tquadbin.
func TbigintToTquadbin(temp Temporal) Temporal {
	res := C.tbigint_to_tquadbin(temp.Inner())
	return CreateTemporal(res)
}


// TquadbinToTbigint wraps MEOS C function tquadbin_to_tbigint.
func TquadbinToTbigint(temp Temporal) Temporal {
	res := C.tquadbin_to_tbigint(temp.Inner())
	return CreateTemporal(res)
}


// EverEqQuadbinTquadbin wraps MEOS C function ever_eq_quadbin_tquadbin.
func EverEqQuadbinTquadbin(cell uint64, temp Temporal) int {
	res := C.ever_eq_quadbin_tquadbin(C.Quadbin(cell), temp.Inner())
	return int(res)
}


// EverEqTquadbinQuadbin wraps MEOS C function ever_eq_tquadbin_quadbin.
func EverEqTquadbinQuadbin(temp Temporal, cell uint64) int {
	res := C.ever_eq_tquadbin_quadbin(temp.Inner(), C.Quadbin(cell))
	return int(res)
}


// EverNeQuadbinTquadbin wraps MEOS C function ever_ne_quadbin_tquadbin.
func EverNeQuadbinTquadbin(cell uint64, temp Temporal) int {
	res := C.ever_ne_quadbin_tquadbin(C.Quadbin(cell), temp.Inner())
	return int(res)
}


// EverNeTquadbinQuadbin wraps MEOS C function ever_ne_tquadbin_quadbin.
func EverNeTquadbinQuadbin(temp Temporal, cell uint64) int {
	res := C.ever_ne_tquadbin_quadbin(temp.Inner(), C.Quadbin(cell))
	return int(res)
}


// AlwaysEqQuadbinTquadbin wraps MEOS C function always_eq_quadbin_tquadbin.
func AlwaysEqQuadbinTquadbin(cell uint64, temp Temporal) int {
	res := C.always_eq_quadbin_tquadbin(C.Quadbin(cell), temp.Inner())
	return int(res)
}


// AlwaysEqTquadbinQuadbin wraps MEOS C function always_eq_tquadbin_quadbin.
func AlwaysEqTquadbinQuadbin(temp Temporal, cell uint64) int {
	res := C.always_eq_tquadbin_quadbin(temp.Inner(), C.Quadbin(cell))
	return int(res)
}


// AlwaysNeQuadbinTquadbin wraps MEOS C function always_ne_quadbin_tquadbin.
func AlwaysNeQuadbinTquadbin(cell uint64, temp Temporal) int {
	res := C.always_ne_quadbin_tquadbin(C.Quadbin(cell), temp.Inner())
	return int(res)
}


// AlwaysNeTquadbinQuadbin wraps MEOS C function always_ne_tquadbin_quadbin.
func AlwaysNeTquadbinQuadbin(temp Temporal, cell uint64) int {
	res := C.always_ne_tquadbin_quadbin(temp.Inner(), C.Quadbin(cell))
	return int(res)
}


// EverEqTquadbinTquadbin wraps MEOS C function ever_eq_tquadbin_tquadbin.
func EverEqTquadbinTquadbin(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_eq_tquadbin_tquadbin(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverNeTquadbinTquadbin wraps MEOS C function ever_ne_tquadbin_tquadbin.
func EverNeTquadbinTquadbin(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_ne_tquadbin_tquadbin(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AlwaysEqTquadbinTquadbin wraps MEOS C function always_eq_tquadbin_tquadbin.
func AlwaysEqTquadbinTquadbin(temp1 Temporal, temp2 Temporal) int {
	res := C.always_eq_tquadbin_tquadbin(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AlwaysNeTquadbinTquadbin wraps MEOS C function always_ne_tquadbin_tquadbin.
func AlwaysNeTquadbinTquadbin(temp1 Temporal, temp2 Temporal) int {
	res := C.always_ne_tquadbin_tquadbin(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TeqQuadbinTquadbin wraps MEOS C function teq_quadbin_tquadbin.
func TeqQuadbinTquadbin(cell uint64, temp Temporal) Temporal {
	res := C.teq_quadbin_tquadbin(C.Quadbin(cell), temp.Inner())
	return CreateTemporal(res)
}


// TeqTquadbinQuadbin wraps MEOS C function teq_tquadbin_quadbin.
func TeqTquadbinQuadbin(temp Temporal, cell uint64) Temporal {
	res := C.teq_tquadbin_quadbin(temp.Inner(), C.Quadbin(cell))
	return CreateTemporal(res)
}


// TeqTquadbinTquadbin wraps MEOS C function teq_tquadbin_tquadbin.
func TeqTquadbinTquadbin(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.teq_tquadbin_tquadbin(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TneQuadbinTquadbin wraps MEOS C function tne_quadbin_tquadbin.
func TneQuadbinTquadbin(cell uint64, temp Temporal) Temporal {
	res := C.tne_quadbin_tquadbin(C.Quadbin(cell), temp.Inner())
	return CreateTemporal(res)
}


// TneTquadbinQuadbin wraps MEOS C function tne_tquadbin_quadbin.
func TneTquadbinQuadbin(temp Temporal, cell uint64) Temporal {
	res := C.tne_tquadbin_quadbin(temp.Inner(), C.Quadbin(cell))
	return CreateTemporal(res)
}


// TneTquadbinTquadbin wraps MEOS C function tne_tquadbin_tquadbin.
func TneTquadbinTquadbin(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tne_tquadbin_tquadbin(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TquadbinCellToQuadkey wraps MEOS C function tquadbin_cell_to_quadkey.
func TquadbinCellToQuadkey(temp Temporal) Temporal {
	res := C.tquadbin_cell_to_quadkey(temp.Inner())
	return CreateTemporal(res)
}

