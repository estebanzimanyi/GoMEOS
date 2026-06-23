package generated

// #include <stddef.h>
import "C"
import (
	"unsafe"

	"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

// Th3indexIn wraps MEOS C function th3index_in.
func Th3indexIn(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.th3index_in(_c_str)
	return CreateTemporal(res)
}


// Th3indexinstIn wraps MEOS C function th3indexinst_in.
func Th3indexinstIn(str string) TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.th3indexinst_in(_c_str)
	return TInstant{_inner: res}
}


// Th3indexseqIn wraps MEOS C function th3indexseq_in.
func Th3indexseqIn(str string, interp Interpolation) TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.th3indexseq_in(_c_str, C.interpType(interp))
	return TSequence{_inner: res}
}


// Th3indexseqsetIn wraps MEOS C function th3indexseqset_in.
func Th3indexseqsetIn(str string) TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.th3indexseqset_in(_c_str)
	return TSequenceSet{_inner: res}
}


// Th3indexMake wraps MEOS C function th3index_make.
func Th3indexMake(value uint64, t int64) Temporal {
	res := C.th3index_make(C.uint64_t(value), C.TimestampTz(t))
	return CreateTemporal(res)
}


// Th3indexinstMake wraps MEOS C function th3indexinst_make.
func Th3indexinstMake(value uint64, t int64) TInstant {
	res := C.th3indexinst_make(C.uint64_t(value), C.TimestampTz(t))
	return TInstant{_inner: res}
}


// TODO th3indexseq_make: unsupported param const uint64_t *
// func Th3indexseqMake(...) { /* not yet handled by codegen */ }


// Th3indexseqsetMake wraps MEOS C function th3indexseqset_make.
func Th3indexseqsetMake(sequences []TSequence) TSequenceSet {
	_c_sequences := make([]*C.TSequence, len(sequences))
	for _i, _v := range sequences { _c_sequences[_i] = _v._inner }
	res := C.th3indexseqset_make((**C.TSequence)(unsafe.Pointer(&_c_sequences[0])), C.int(len(sequences)))
	return TSequenceSet{_inner: res}
}


// Th3indexStartValue wraps MEOS C function th3index_start_value.
func Th3indexStartValue(temp Temporal) uint64 {
	res := C.th3index_start_value(temp.Inner())
	return uint64(res)
}


// Th3indexEndValue wraps MEOS C function th3index_end_value.
func Th3indexEndValue(temp Temporal) uint64 {
	res := C.th3index_end_value(temp.Inner())
	return uint64(res)
}


// Th3indexValueN wraps MEOS C function th3index_value_n.
func Th3indexValueN(temp Temporal, n int) (bool, uint64) {
	var _out_result C.uint64_t
	res := C.th3index_value_n(temp.Inner(), C.int(n), &_out_result)
	return bool(res), uint64(_out_result)
}


// Th3indexValues wraps MEOS C function th3index_values.
func Th3indexValues(temp Temporal) []uint64 {
	var _out_count C.int
	res := C.th3index_values(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((*C.uint64_t)(unsafe.Pointer(res)), _n)
	_out := make([]uint64, _n)
	for _i, _e := range _slice {
		_out[_i] = uint64(_e)
	}
	return _out
}


// Th3indexValueAtTimestamptz wraps MEOS C function th3index_value_at_timestamptz.
func Th3indexValueAtTimestamptz(temp Temporal, t int64, strict bool) (bool, uint64) {
	var _out_result C.uint64_t
	res := C.th3index_value_at_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(strict), &_out_result)
	return bool(res), uint64(_out_result)
}


// TbigintToTh3index wraps MEOS C function tbigint_to_th3index.
func TbigintToTh3index(temp Temporal) Temporal {
	res := C.tbigint_to_th3index(temp.Inner())
	return CreateTemporal(res)
}


// Th3indexToTbigint wraps MEOS C function th3index_to_tbigint.
func Th3indexToTbigint(temp Temporal) Temporal {
	res := C.th3index_to_tbigint(temp.Inner())
	return CreateTemporal(res)
}


// EverEqH3indexTh3index wraps MEOS C function ever_eq_h3index_th3index.
func EverEqH3indexTh3index(cell uint64, temp Temporal) int {
	res := C.ever_eq_h3index_th3index(C.uint64_t(cell), temp.Inner())
	return int(res)
}


// EverEqTh3indexH3index wraps MEOS C function ever_eq_th3index_h3index.
func EverEqTh3indexH3index(temp Temporal, cell uint64) int {
	res := C.ever_eq_th3index_h3index(temp.Inner(), C.uint64_t(cell))
	return int(res)
}


// EverNeH3indexTh3index wraps MEOS C function ever_ne_h3index_th3index.
func EverNeH3indexTh3index(cell uint64, temp Temporal) int {
	res := C.ever_ne_h3index_th3index(C.uint64_t(cell), temp.Inner())
	return int(res)
}


// EverNeTh3indexH3index wraps MEOS C function ever_ne_th3index_h3index.
func EverNeTh3indexH3index(temp Temporal, cell uint64) int {
	res := C.ever_ne_th3index_h3index(temp.Inner(), C.uint64_t(cell))
	return int(res)
}


// AlwaysEqH3indexTh3index wraps MEOS C function always_eq_h3index_th3index.
func AlwaysEqH3indexTh3index(cell uint64, temp Temporal) int {
	res := C.always_eq_h3index_th3index(C.uint64_t(cell), temp.Inner())
	return int(res)
}


// AlwaysEqTh3indexH3index wraps MEOS C function always_eq_th3index_h3index.
func AlwaysEqTh3indexH3index(temp Temporal, cell uint64) int {
	res := C.always_eq_th3index_h3index(temp.Inner(), C.uint64_t(cell))
	return int(res)
}


// AlwaysNeH3indexTh3index wraps MEOS C function always_ne_h3index_th3index.
func AlwaysNeH3indexTh3index(cell uint64, temp Temporal) int {
	res := C.always_ne_h3index_th3index(C.uint64_t(cell), temp.Inner())
	return int(res)
}


// AlwaysNeTh3indexH3index wraps MEOS C function always_ne_th3index_h3index.
func AlwaysNeTh3indexH3index(temp Temporal, cell uint64) int {
	res := C.always_ne_th3index_h3index(temp.Inner(), C.uint64_t(cell))
	return int(res)
}


// EverEqTh3indexTh3index wraps MEOS C function ever_eq_th3index_th3index.
func EverEqTh3indexTh3index(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_eq_th3index_th3index(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverNeTh3indexTh3index wraps MEOS C function ever_ne_th3index_th3index.
func EverNeTh3indexTh3index(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_ne_th3index_th3index(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AlwaysEqTh3indexTh3index wraps MEOS C function always_eq_th3index_th3index.
func AlwaysEqTh3indexTh3index(temp1 Temporal, temp2 Temporal) int {
	res := C.always_eq_th3index_th3index(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AlwaysNeTh3indexTh3index wraps MEOS C function always_ne_th3index_th3index.
func AlwaysNeTh3indexTh3index(temp1 Temporal, temp2 Temporal) int {
	res := C.always_ne_th3index_th3index(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TeqH3indexTh3index wraps MEOS C function teq_h3index_th3index.
func TeqH3indexTh3index(cell uint64, temp Temporal) Temporal {
	res := C.teq_h3index_th3index(C.uint64_t(cell), temp.Inner())
	return CreateTemporal(res)
}


// TeqTh3indexH3index wraps MEOS C function teq_th3index_h3index.
func TeqTh3indexH3index(temp Temporal, cell uint64) Temporal {
	res := C.teq_th3index_h3index(temp.Inner(), C.uint64_t(cell))
	return CreateTemporal(res)
}


// TeqTh3indexTh3index wraps MEOS C function teq_th3index_th3index.
func TeqTh3indexTh3index(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.teq_th3index_th3index(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TneH3indexTh3index wraps MEOS C function tne_h3index_th3index.
func TneH3indexTh3index(cell uint64, temp Temporal) Temporal {
	res := C.tne_h3index_th3index(C.uint64_t(cell), temp.Inner())
	return CreateTemporal(res)
}


// TneTh3indexH3index wraps MEOS C function tne_th3index_h3index.
func TneTh3indexH3index(temp Temporal, cell uint64) Temporal {
	res := C.tne_th3index_h3index(temp.Inner(), C.uint64_t(cell))
	return CreateTemporal(res)
}


// TneTh3indexTh3index wraps MEOS C function tne_th3index_th3index.
func TneTh3indexTh3index(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tne_th3index_th3index(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// Th3indexGetResolution wraps MEOS C function th3index_get_resolution.
func Th3indexGetResolution(temp Temporal) Temporal {
	res := C.th3index_get_resolution(temp.Inner())
	return CreateTemporal(res)
}


// Th3indexGetBaseCellNumber wraps MEOS C function th3index_get_base_cell_number.
func Th3indexGetBaseCellNumber(temp Temporal) Temporal {
	res := C.th3index_get_base_cell_number(temp.Inner())
	return CreateTemporal(res)
}


// Th3indexIsValidCell wraps MEOS C function th3index_is_valid_cell.
func Th3indexIsValidCell(temp Temporal) Temporal {
	res := C.th3index_is_valid_cell(temp.Inner())
	return CreateTemporal(res)
}


// Th3indexIsResClassIii wraps MEOS C function th3index_is_res_class_iii.
func Th3indexIsResClassIii(temp Temporal) Temporal {
	res := C.th3index_is_res_class_iii(temp.Inner())
	return CreateTemporal(res)
}


// Th3indexIsPentagon wraps MEOS C function th3index_is_pentagon.
func Th3indexIsPentagon(temp Temporal) Temporal {
	res := C.th3index_is_pentagon(temp.Inner())
	return CreateTemporal(res)
}


// Th3indexCellToParent wraps MEOS C function th3index_cell_to_parent.
func Th3indexCellToParent(temp Temporal, resolution int) Temporal {
	res := C.th3index_cell_to_parent(temp.Inner(), C.int(resolution))
	return CreateTemporal(res)
}


// Th3indexCellToParentNext wraps MEOS C function th3index_cell_to_parent_next.
func Th3indexCellToParentNext(temp Temporal) Temporal {
	res := C.th3index_cell_to_parent_next(temp.Inner())
	return CreateTemporal(res)
}


// Th3indexCellToCenterChild wraps MEOS C function th3index_cell_to_center_child.
func Th3indexCellToCenterChild(temp Temporal, resolution int) Temporal {
	res := C.th3index_cell_to_center_child(temp.Inner(), C.int(resolution))
	return CreateTemporal(res)
}


// Th3indexCellToCenterChildNext wraps MEOS C function th3index_cell_to_center_child_next.
func Th3indexCellToCenterChildNext(temp Temporal) Temporal {
	res := C.th3index_cell_to_center_child_next(temp.Inner())
	return CreateTemporal(res)
}


// Th3indexCellToChildPos wraps MEOS C function th3index_cell_to_child_pos.
func Th3indexCellToChildPos(temp Temporal, parent_res int) Temporal {
	res := C.th3index_cell_to_child_pos(temp.Inner(), C.int(parent_res))
	return CreateTemporal(res)
}


// Th3indexChildPosToCell wraps MEOS C function th3index_child_pos_to_cell.
func Th3indexChildPosToCell(child_pos Temporal, parent Temporal, child_res int) Temporal {
	res := C.th3index_child_pos_to_cell(child_pos.Inner(), parent.Inner(), C.int(child_res))
	return CreateTemporal(res)
}


// TgeogpointToTh3index wraps MEOS C function tgeogpoint_to_th3index.
func TgeogpointToTh3index(temp Temporal, resolution int) Temporal {
	res := C.tgeogpoint_to_th3index(temp.Inner(), C.int(resolution))
	return CreateTemporal(res)
}


// TgeompointToTh3index wraps MEOS C function tgeompoint_to_th3index.
func TgeompointToTh3index(temp Temporal, resolution int) Temporal {
	res := C.tgeompoint_to_th3index(temp.Inner(), C.int(resolution))
	return CreateTemporal(res)
}


// Th3indexToTgeogpoint wraps MEOS C function th3index_to_tgeogpoint.
func Th3indexToTgeogpoint(temp Temporal) Temporal {
	res := C.th3index_to_tgeogpoint(temp.Inner())
	return CreateTemporal(res)
}


// Th3indexToTgeompoint wraps MEOS C function th3index_to_tgeompoint.
func Th3indexToTgeompoint(temp Temporal) Temporal {
	res := C.th3index_to_tgeompoint(temp.Inner())
	return CreateTemporal(res)
}


// Th3indexCellToBoundary wraps MEOS C function th3index_cell_to_boundary.
func Th3indexCellToBoundary(temp Temporal) Temporal {
	res := C.th3index_cell_to_boundary(temp.Inner())
	return CreateTemporal(res)
}


// GeoToH3indexSet wraps MEOS C function geo_to_h3index_set.
func GeoToH3indexSet(gs *Geom, resolution int) *Set {
	res := C.geo_to_h3index_set(gs._inner, C.int(resolution))
	return &Set{_inner: res}
}


// EverEqH3indexsetTh3index wraps MEOS C function ever_eq_h3indexset_th3index.
func EverEqH3indexsetTh3index(cells *Set, th3idx Temporal) int {
	res := C.ever_eq_h3indexset_th3index(cells._inner, th3idx.Inner())
	return int(res)
}


// Th3indexAreNeighborCells wraps MEOS C function th3index_are_neighbor_cells.
func Th3indexAreNeighborCells(origin Temporal, dest Temporal) Temporal {
	res := C.th3index_are_neighbor_cells(origin.Inner(), dest.Inner())
	return CreateTemporal(res)
}


// Th3indexCellsToDirectedEdge wraps MEOS C function th3index_cells_to_directed_edge.
func Th3indexCellsToDirectedEdge(origin Temporal, dest Temporal) Temporal {
	res := C.th3index_cells_to_directed_edge(origin.Inner(), dest.Inner())
	return CreateTemporal(res)
}


// Th3indexIsValidDirectedEdge wraps MEOS C function th3index_is_valid_directed_edge.
func Th3indexIsValidDirectedEdge(edge Temporal) Temporal {
	res := C.th3index_is_valid_directed_edge(edge.Inner())
	return CreateTemporal(res)
}


// Th3indexGetDirectedEdgeOrigin wraps MEOS C function th3index_get_directed_edge_origin.
func Th3indexGetDirectedEdgeOrigin(edge Temporal) Temporal {
	res := C.th3index_get_directed_edge_origin(edge.Inner())
	return CreateTemporal(res)
}


// Th3indexGetDirectedEdgeDestination wraps MEOS C function th3index_get_directed_edge_destination.
func Th3indexGetDirectedEdgeDestination(edge Temporal) Temporal {
	res := C.th3index_get_directed_edge_destination(edge.Inner())
	return CreateTemporal(res)
}


// Th3indexDirectedEdgeToBoundary wraps MEOS C function th3index_directed_edge_to_boundary.
func Th3indexDirectedEdgeToBoundary(edge Temporal) Temporal {
	res := C.th3index_directed_edge_to_boundary(edge.Inner())
	return CreateTemporal(res)
}


// Th3indexCellToVertex wraps MEOS C function th3index_cell_to_vertex.
func Th3indexCellToVertex(temp Temporal, vertex_num int) Temporal {
	res := C.th3index_cell_to_vertex(temp.Inner(), C.int(vertex_num))
	return CreateTemporal(res)
}


// Th3indexVertexToLatlng wraps MEOS C function th3index_vertex_to_latlng.
func Th3indexVertexToLatlng(temp Temporal) Temporal {
	res := C.th3index_vertex_to_latlng(temp.Inner())
	return CreateTemporal(res)
}


// Th3indexIsValidVertex wraps MEOS C function th3index_is_valid_vertex.
func Th3indexIsValidVertex(temp Temporal) Temporal {
	res := C.th3index_is_valid_vertex(temp.Inner())
	return CreateTemporal(res)
}


// Th3indexGridDistance wraps MEOS C function th3index_grid_distance.
func Th3indexGridDistance(origin Temporal, dest Temporal) Temporal {
	res := C.th3index_grid_distance(origin.Inner(), dest.Inner())
	return CreateTemporal(res)
}


// Th3indexCellToLocalIj wraps MEOS C function th3index_cell_to_local_ij.
func Th3indexCellToLocalIj(origin Temporal, cell Temporal) Temporal {
	res := C.th3index_cell_to_local_ij(origin.Inner(), cell.Inner())
	return CreateTemporal(res)
}


// Th3indexLocalIjToCell wraps MEOS C function th3index_local_ij_to_cell.
func Th3indexLocalIjToCell(origin Temporal, coord Temporal) Temporal {
	res := C.th3index_local_ij_to_cell(origin.Inner(), coord.Inner())
	return CreateTemporal(res)
}


// Th3indexCellArea wraps MEOS C function th3index_cell_area.
func Th3indexCellArea(temp Temporal, unit string) Temporal {
	_c_unit := C.CString(unit)
	defer C.free(unsafe.Pointer(_c_unit))
	res := C.th3index_cell_area(temp.Inner(), _c_unit)
	return CreateTemporal(res)
}


// Th3indexEdgeLength wraps MEOS C function th3index_edge_length.
func Th3indexEdgeLength(temp Temporal, unit string) Temporal {
	_c_unit := C.CString(unit)
	defer C.free(unsafe.Pointer(_c_unit))
	res := C.th3index_edge_length(temp.Inner(), _c_unit)
	return CreateTemporal(res)
}


// TgeogpointGreatCircleDistance wraps MEOS C function tgeogpoint_great_circle_distance.
func TgeogpointGreatCircleDistance(a Temporal, b Temporal, unit string) Temporal {
	_c_unit := C.CString(unit)
	defer C.free(unsafe.Pointer(_c_unit))
	res := C.tgeogpoint_great_circle_distance(a.Inner(), b.Inner(), _c_unit)
	return CreateTemporal(res)
}

