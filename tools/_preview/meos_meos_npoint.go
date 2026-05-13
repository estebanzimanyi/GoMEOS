package generated

// #include <stddef.h>
import "C"
import (
	"unsafe"

	"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

// NpointAsEWKT wraps MEOS C function npoint_as_ewkt.
func NpointAsEWKT(np *Npoint, maxdd int) string {
	res := C.npoint_as_ewkt(np._inner, C.int(maxdd))
	return C.GoString(res)
}


// NpointAsHexwkb wraps MEOS C function npoint_as_hexwkb.
func NpointAsHexwkb(np *Npoint, variant uint8) (string, uint) {
	var _out_size_out C.size_t
	res := C.npoint_as_hexwkb(np._inner, C.uint8_t(variant), &_out_size_out)
	return C.GoString(res), uint(_out_size_out)
}


// NpointAsText wraps MEOS C function npoint_as_text.
func NpointAsText(np *Npoint, maxdd int) string {
	res := C.npoint_as_text(np._inner, C.int(maxdd))
	return C.GoString(res)
}


// NpointAsWKB wraps MEOS C function npoint_as_wkb.
func NpointAsWKB(np *Npoint, variant uint8) []uint8 {
	var _out_size_out C.size_t
	res := C.npoint_as_wkb(np._inner, C.uint8_t(variant), &_out_size_out)
	_n := int(_out_size_out)
	_slice := unsafe.Slice((*C.uint8_t)(unsafe.Pointer(res)), _n)
	_out := make([]uint8, _n)
	for _i, _e := range _slice {
		_out[_i] = uint8(_e)
	}
	return _out
}


// NpointFromHexwkb wraps MEOS C function npoint_from_hexwkb.
func NpointFromHexwkb(hexwkb string) *Npoint {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	res := C.npoint_from_hexwkb(_c_hexwkb)
	return &Npoint{_inner: res}
}


// NpointFromWKB wraps MEOS C function npoint_from_wkb.
func NpointFromWKB(wkb []byte) *Npoint {
	res := C.npoint_from_wkb((*C.uint8_t)(unsafe.Pointer(&wkb[0])), C.size_t(len(wkb)))
	return &Npoint{_inner: res}
}


// NpointIn wraps MEOS C function npoint_in.
func NpointIn(str string) *Npoint {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.npoint_in(_c_str)
	return &Npoint{_inner: res}
}


// NpointOut wraps MEOS C function npoint_out.
func NpointOut(np *Npoint, maxdd int) string {
	res := C.npoint_out(np._inner, C.int(maxdd))
	return C.GoString(res)
}


// NsegmentIn wraps MEOS C function nsegment_in.
func NsegmentIn(str string) *Nsegment {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.nsegment_in(_c_str)
	return &Nsegment{_inner: res}
}


// NsegmentOut wraps MEOS C function nsegment_out.
func NsegmentOut(ns *Nsegment, maxdd int) string {
	res := C.nsegment_out(ns._inner, C.int(maxdd))
	return C.GoString(res)
}


// NpointMake wraps MEOS C function npoint_make.
func NpointMake(rid int64, pos float64) *Npoint {
	res := C.npoint_make(C.int64(rid), C.double(pos))
	return &Npoint{_inner: res}
}


// NsegmentMake wraps MEOS C function nsegment_make.
func NsegmentMake(rid int64, pos1 float64, pos2 float64) *Nsegment {
	res := C.nsegment_make(C.int64(rid), C.double(pos1), C.double(pos2))
	return &Nsegment{_inner: res}
}


// GeompointToNpoint wraps MEOS C function geompoint_to_npoint.
func GeompointToNpoint(gs *Geom) *Npoint {
	res := C.geompoint_to_npoint(gs._inner)
	return &Npoint{_inner: res}
}


// GeomToNsegment wraps MEOS C function geom_to_nsegment.
func GeomToNsegment(gs *Geom) *Nsegment {
	res := C.geom_to_nsegment(gs._inner)
	return &Nsegment{_inner: res}
}


// NpointToGeompoint wraps MEOS C function npoint_to_geompoint.
func NpointToGeompoint(np *Npoint) *Geom {
	res := C.npoint_to_geompoint(np._inner)
	return &Geom{_inner: res}
}


// NpointToNsegment wraps MEOS C function npoint_to_nsegment.
func NpointToNsegment(np *Npoint) *Nsegment {
	res := C.npoint_to_nsegment(np._inner)
	return &Nsegment{_inner: res}
}


// NpointToSTBOX wraps MEOS C function npoint_to_stbox.
func NpointToSTBOX(np *Npoint) *STBox {
	res := C.npoint_to_stbox(np._inner)
	return &STBox{_inner: res}
}


// NsegmentToGeom wraps MEOS C function nsegment_to_geom.
func NsegmentToGeom(ns *Nsegment) *Geom {
	res := C.nsegment_to_geom(ns._inner)
	return &Geom{_inner: res}
}


// NsegmentToSTBOX wraps MEOS C function nsegment_to_stbox.
func NsegmentToSTBOX(np *Nsegment) *STBox {
	res := C.nsegment_to_stbox(np._inner)
	return &STBox{_inner: res}
}


// NpointHash wraps MEOS C function npoint_hash.
func NpointHash(np *Npoint) uint32 {
	res := C.npoint_hash(np._inner)
	return uint32(res)
}


// NpointHashExtended wraps MEOS C function npoint_hash_extended.
func NpointHashExtended(np *Npoint, seed uint64) uint64 {
	res := C.npoint_hash_extended(np._inner, C.uint64(seed))
	return uint64(res)
}


// NpointPosition wraps MEOS C function npoint_position.
func NpointPosition(np *Npoint) float64 {
	res := C.npoint_position(np._inner)
	return float64(res)
}


// NpointRoute wraps MEOS C function npoint_route.
func NpointRoute(np *Npoint) int64 {
	res := C.npoint_route(np._inner)
	return int64(res)
}


// NsegmentEndPosition wraps MEOS C function nsegment_end_position.
func NsegmentEndPosition(ns *Nsegment) float64 {
	res := C.nsegment_end_position(ns._inner)
	return float64(res)
}


// NsegmentRoute wraps MEOS C function nsegment_route.
func NsegmentRoute(ns *Nsegment) int64 {
	res := C.nsegment_route(ns._inner)
	return int64(res)
}


// NsegmentStartPosition wraps MEOS C function nsegment_start_position.
func NsegmentStartPosition(ns *Nsegment) float64 {
	res := C.nsegment_start_position(ns._inner)
	return float64(res)
}


// RouteExists wraps MEOS C function route_exists.
func RouteExists(rid int64) bool {
	res := C.route_exists(C.int64(rid))
	return bool(res)
}


// RouteGeom wraps MEOS C function route_geom.
func RouteGeom(rid int64) *Geom {
	res := C.route_geom(C.int64(rid))
	return &Geom{_inner: res}
}


// RouteLength wraps MEOS C function route_length.
func RouteLength(rid int64) float64 {
	res := C.route_length(C.int64(rid))
	return float64(res)
}


// NpointRound wraps MEOS C function npoint_round.
func NpointRound(np *Npoint, maxdd int) *Npoint {
	res := C.npoint_round(np._inner, C.int(maxdd))
	return &Npoint{_inner: res}
}


// NsegmentRound wraps MEOS C function nsegment_round.
func NsegmentRound(ns *Nsegment, maxdd int) *Nsegment {
	res := C.nsegment_round(ns._inner, C.int(maxdd))
	return &Nsegment{_inner: res}
}


// GetSRIDWays wraps MEOS C function get_srid_ways.
func GetSRIDWays() int32 {
	res := C.get_srid_ways()
	return int32(res)
}


// NpointSRID wraps MEOS C function npoint_srid.
func NpointSRID(np *Npoint) int32 {
	res := C.npoint_srid(np._inner)
	return int32(res)
}


// NsegmentSRID wraps MEOS C function nsegment_srid.
func NsegmentSRID(ns *Nsegment) int32 {
	res := C.nsegment_srid(ns._inner)
	return int32(res)
}


// NpointTimestamptzToSTBOX wraps MEOS C function npoint_timestamptz_to_stbox.
func NpointTimestamptzToSTBOX(np *Npoint, t int64) *STBox {
	res := C.npoint_timestamptz_to_stbox(np._inner, C.TimestampTz(t))
	return &STBox{_inner: res}
}


// NpointTstzspanToSTBOX wraps MEOS C function npoint_tstzspan_to_stbox.
func NpointTstzspanToSTBOX(np *Npoint, s *Span) *STBox {
	res := C.npoint_tstzspan_to_stbox(np._inner, s._inner)
	return &STBox{_inner: res}
}


// NpointCmp wraps MEOS C function npoint_cmp.
func NpointCmp(np1 *Npoint, np2 *Npoint) int {
	res := C.npoint_cmp(np1._inner, np2._inner)
	return int(res)
}


// NpointEq wraps MEOS C function npoint_eq.
func NpointEq(np1 *Npoint, np2 *Npoint) bool {
	res := C.npoint_eq(np1._inner, np2._inner)
	return bool(res)
}


// NpointGe wraps MEOS C function npoint_ge.
func NpointGe(np1 *Npoint, np2 *Npoint) bool {
	res := C.npoint_ge(np1._inner, np2._inner)
	return bool(res)
}


// NpointGt wraps MEOS C function npoint_gt.
func NpointGt(np1 *Npoint, np2 *Npoint) bool {
	res := C.npoint_gt(np1._inner, np2._inner)
	return bool(res)
}


// NpointLe wraps MEOS C function npoint_le.
func NpointLe(np1 *Npoint, np2 *Npoint) bool {
	res := C.npoint_le(np1._inner, np2._inner)
	return bool(res)
}


// NpointLt wraps MEOS C function npoint_lt.
func NpointLt(np1 *Npoint, np2 *Npoint) bool {
	res := C.npoint_lt(np1._inner, np2._inner)
	return bool(res)
}


// NpointNe wraps MEOS C function npoint_ne.
func NpointNe(np1 *Npoint, np2 *Npoint) bool {
	res := C.npoint_ne(np1._inner, np2._inner)
	return bool(res)
}


// NpointSame wraps MEOS C function npoint_same.
func NpointSame(np1 *Npoint, np2 *Npoint) bool {
	res := C.npoint_same(np1._inner, np2._inner)
	return bool(res)
}


// NsegmentCmp wraps MEOS C function nsegment_cmp.
func NsegmentCmp(ns1 *Nsegment, ns2 *Nsegment) int {
	res := C.nsegment_cmp(ns1._inner, ns2._inner)
	return int(res)
}


// NsegmentEq wraps MEOS C function nsegment_eq.
func NsegmentEq(ns1 *Nsegment, ns2 *Nsegment) bool {
	res := C.nsegment_eq(ns1._inner, ns2._inner)
	return bool(res)
}


// NsegmentGe wraps MEOS C function nsegment_ge.
func NsegmentGe(ns1 *Nsegment, ns2 *Nsegment) bool {
	res := C.nsegment_ge(ns1._inner, ns2._inner)
	return bool(res)
}


// NsegmentGt wraps MEOS C function nsegment_gt.
func NsegmentGt(ns1 *Nsegment, ns2 *Nsegment) bool {
	res := C.nsegment_gt(ns1._inner, ns2._inner)
	return bool(res)
}


// NsegmentLe wraps MEOS C function nsegment_le.
func NsegmentLe(ns1 *Nsegment, ns2 *Nsegment) bool {
	res := C.nsegment_le(ns1._inner, ns2._inner)
	return bool(res)
}


// NsegmentLt wraps MEOS C function nsegment_lt.
func NsegmentLt(ns1 *Nsegment, ns2 *Nsegment) bool {
	res := C.nsegment_lt(ns1._inner, ns2._inner)
	return bool(res)
}


// NsegmentNe wraps MEOS C function nsegment_ne.
func NsegmentNe(ns1 *Nsegment, ns2 *Nsegment) bool {
	res := C.nsegment_ne(ns1._inner, ns2._inner)
	return bool(res)
}


// NpointsetIn wraps MEOS C function npointset_in.
func NpointsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.npointset_in(_c_str)
	return &Set{_inner: res}
}


// NpointsetOut wraps MEOS C function npointset_out.
func NpointsetOut(s *Set, maxdd int) string {
	res := C.npointset_out(s._inner, C.int(maxdd))
	return C.GoString(res)
}


// NpointsetMake wraps MEOS C function npointset_make.
func NpointsetMake(values []*Npoint) *Set {
	_c_values := make([]*C.Npoint, len(values))
	for _i, _v := range values { _c_values[_i] = _v._inner }
	res := C.npointset_make((**C.Npoint)(unsafe.Pointer(&_c_values[0])), C.int(len(values)))
	return &Set{_inner: res}
}


// NpointToSet wraps MEOS C function npoint_to_set.
func NpointToSet(np *Npoint) *Set {
	res := C.npoint_to_set(np._inner)
	return &Set{_inner: res}
}


// NpointsetEndValue wraps MEOS C function npointset_end_value.
func NpointsetEndValue(s *Set) *Npoint {
	res := C.npointset_end_value(s._inner)
	return &Npoint{_inner: res}
}


// NpointsetRoutes wraps MEOS C function npointset_routes.
func NpointsetRoutes(s *Set) *Set {
	res := C.npointset_routes(s._inner)
	return &Set{_inner: res}
}


// NpointsetStartValue wraps MEOS C function npointset_start_value.
func NpointsetStartValue(s *Set) *Npoint {
	res := C.npointset_start_value(s._inner)
	return &Npoint{_inner: res}
}


// NpointsetValueN wraps MEOS C function npointset_value_n.
func NpointsetValueN(s *Set, n int) (bool, *Npoint) {
	var _out_result *C.Npoint
	res := C.npointset_value_n(s._inner, C.int(n), &_out_result)
	return bool(res), &Npoint{_inner: _out_result}
}


// TODO npointset_values: unsupported return type Npoint **
// func NpointsetValues(...) { /* not yet handled by codegen */ }


// ContainedNpointSet wraps MEOS C function contained_npoint_set.
func ContainedNpointSet(np *Npoint, s *Set) bool {
	res := C.contained_npoint_set(np._inner, s._inner)
	return bool(res)
}


// ContainsSetNpoint wraps MEOS C function contains_set_npoint.
func ContainsSetNpoint(s *Set, np *Npoint) bool {
	res := C.contains_set_npoint(s._inner, np._inner)
	return bool(res)
}


// IntersectionNpointSet wraps MEOS C function intersection_npoint_set.
func IntersectionNpointSet(np *Npoint, s *Set) *Set {
	res := C.intersection_npoint_set(np._inner, s._inner)
	return &Set{_inner: res}
}


// IntersectionSetNpoint wraps MEOS C function intersection_set_npoint.
func IntersectionSetNpoint(s *Set, np *Npoint) *Set {
	res := C.intersection_set_npoint(s._inner, np._inner)
	return &Set{_inner: res}
}


// MinusNpointSet wraps MEOS C function minus_npoint_set.
func MinusNpointSet(np *Npoint, s *Set) *Set {
	res := C.minus_npoint_set(np._inner, s._inner)
	return &Set{_inner: res}
}


// MinusSetNpoint wraps MEOS C function minus_set_npoint.
func MinusSetNpoint(s *Set, np *Npoint) *Set {
	res := C.minus_set_npoint(s._inner, np._inner)
	return &Set{_inner: res}
}


// NpointUnionTransfn wraps MEOS C function npoint_union_transfn.
func NpointUnionTransfn(state *Set, np *Npoint) *Set {
	res := C.npoint_union_transfn(state._inner, np._inner)
	return &Set{_inner: res}
}


// UnionNpointSet wraps MEOS C function union_npoint_set.
func UnionNpointSet(np *Npoint, s *Set) *Set {
	res := C.union_npoint_set(np._inner, s._inner)
	return &Set{_inner: res}
}


// UnionSetNpoint wraps MEOS C function union_set_npoint.
func UnionSetNpoint(s *Set, np *Npoint) *Set {
	res := C.union_set_npoint(s._inner, np._inner)
	return &Set{_inner: res}
}


// TnpointIn wraps MEOS C function tnpoint_in.
func TnpointIn(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tnpoint_in(_c_str)
	return CreateTemporal(res)
}


// TnpointOut wraps MEOS C function tnpoint_out.
func TnpointOut(temp Temporal, maxdd int) string {
	res := C.tnpoint_out(temp.Inner(), C.int(maxdd))
	return C.GoString(res)
}


// TnpointinstMake wraps MEOS C function tnpointinst_make.
func TnpointinstMake(np *Npoint, t int64) TInstant {
	res := C.tnpointinst_make(np._inner, C.TimestampTz(t))
	return TInstant{_inner: res}
}


// TgeompointToTnpoint wraps MEOS C function tgeompoint_to_tnpoint.
func TgeompointToTnpoint(temp Temporal) Temporal {
	res := C.tgeompoint_to_tnpoint(temp.Inner())
	return CreateTemporal(res)
}


// TnpointToTgeompoint wraps MEOS C function tnpoint_to_tgeompoint.
func TnpointToTgeompoint(temp Temporal) Temporal {
	res := C.tnpoint_to_tgeompoint(temp.Inner())
	return CreateTemporal(res)
}


// TnpointCumulativeLength wraps MEOS C function tnpoint_cumulative_length.
func TnpointCumulativeLength(temp Temporal) Temporal {
	res := C.tnpoint_cumulative_length(temp.Inner())
	return CreateTemporal(res)
}


// TnpointLength wraps MEOS C function tnpoint_length.
func TnpointLength(temp Temporal) float64 {
	res := C.tnpoint_length(temp.Inner())
	return float64(res)
}


// TnpointPositions wraps MEOS C function tnpoint_positions.
func TnpointPositions(temp Temporal) []*Nsegment {
	var _out_count C.int
	res := C.tnpoint_positions(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.Nsegment)(unsafe.Pointer(res)), _n)
	_out := make([]*Nsegment, _n)
	for _i, _e := range _slice {
		_out[_i] = &Nsegment{_inner: _e}
	}
	return _out
}


// TnpointRoute wraps MEOS C function tnpoint_route.
func TnpointRoute(temp Temporal) int64 {
	res := C.tnpoint_route(temp.Inner())
	return int64(res)
}


// TnpointRoutes wraps MEOS C function tnpoint_routes.
func TnpointRoutes(temp Temporal) *Set {
	res := C.tnpoint_routes(temp.Inner())
	return &Set{_inner: res}
}


// TnpointSpeed wraps MEOS C function tnpoint_speed.
func TnpointSpeed(temp Temporal) Temporal {
	res := C.tnpoint_speed(temp.Inner())
	return CreateTemporal(res)
}


// TnpointTrajectory wraps MEOS C function tnpoint_trajectory.
func TnpointTrajectory(temp Temporal) *Geom {
	res := C.tnpoint_trajectory(temp.Inner())
	return &Geom{_inner: res}
}


// TnpointTwcentroid wraps MEOS C function tnpoint_twcentroid.
func TnpointTwcentroid(temp Temporal) *Geom {
	res := C.tnpoint_twcentroid(temp.Inner())
	return &Geom{_inner: res}
}


// TnpointAtGeom wraps MEOS C function tnpoint_at_geom.
func TnpointAtGeom(temp Temporal, gs *Geom) Temporal {
	res := C.tnpoint_at_geom(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TnpointAtNpoint wraps MEOS C function tnpoint_at_npoint.
func TnpointAtNpoint(temp Temporal, np *Npoint) Temporal {
	res := C.tnpoint_at_npoint(temp.Inner(), np._inner)
	return CreateTemporal(res)
}


// TnpointAtNpointset wraps MEOS C function tnpoint_at_npointset.
func TnpointAtNpointset(temp Temporal, s *Set) Temporal {
	res := C.tnpoint_at_npointset(temp.Inner(), s._inner)
	return CreateTemporal(res)
}


// TnpointAtSTBOX wraps MEOS C function tnpoint_at_stbox.
func TnpointAtSTBOX(temp Temporal, box *STBox, border_inc bool) Temporal {
	res := C.tnpoint_at_stbox(temp.Inner(), box._inner, C.bool(border_inc))
	return CreateTemporal(res)
}


// TnpointMinusGeom wraps MEOS C function tnpoint_minus_geom.
func TnpointMinusGeom(temp Temporal, gs *Geom) Temporal {
	res := C.tnpoint_minus_geom(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TnpointMinusNpoint wraps MEOS C function tnpoint_minus_npoint.
func TnpointMinusNpoint(temp Temporal, np *Npoint) Temporal {
	res := C.tnpoint_minus_npoint(temp.Inner(), np._inner)
	return CreateTemporal(res)
}


// TnpointMinusNpointset wraps MEOS C function tnpoint_minus_npointset.
func TnpointMinusNpointset(temp Temporal, s *Set) Temporal {
	res := C.tnpoint_minus_npointset(temp.Inner(), s._inner)
	return CreateTemporal(res)
}


// TnpointMinusSTBOX wraps MEOS C function tnpoint_minus_stbox.
func TnpointMinusSTBOX(temp Temporal, box *STBox, border_inc bool) Temporal {
	res := C.tnpoint_minus_stbox(temp.Inner(), box._inner, C.bool(border_inc))
	return CreateTemporal(res)
}


// TdistanceTnpointNpoint wraps MEOS C function tdistance_tnpoint_npoint.
func TdistanceTnpointNpoint(temp Temporal, np *Npoint) Temporal {
	res := C.tdistance_tnpoint_npoint(temp.Inner(), np._inner)
	return CreateTemporal(res)
}


// TdistanceTnpointPoint wraps MEOS C function tdistance_tnpoint_point.
func TdistanceTnpointPoint(temp Temporal, gs *Geom) Temporal {
	res := C.tdistance_tnpoint_point(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TdistanceTnpointTnpoint wraps MEOS C function tdistance_tnpoint_tnpoint.
func TdistanceTnpointTnpoint(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tdistance_tnpoint_tnpoint(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// NadTnpointGeo wraps MEOS C function nad_tnpoint_geo.
func NadTnpointGeo(temp Temporal, gs *Geom) float64 {
	res := C.nad_tnpoint_geo(temp.Inner(), gs._inner)
	return float64(res)
}


// NadTnpointNpoint wraps MEOS C function nad_tnpoint_npoint.
func NadTnpointNpoint(temp Temporal, np *Npoint) float64 {
	res := C.nad_tnpoint_npoint(temp.Inner(), np._inner)
	return float64(res)
}


// NadTnpointSTBOX wraps MEOS C function nad_tnpoint_stbox.
func NadTnpointSTBOX(temp Temporal, box *STBox) float64 {
	res := C.nad_tnpoint_stbox(temp.Inner(), box._inner)
	return float64(res)
}


// NadTnpointTnpoint wraps MEOS C function nad_tnpoint_tnpoint.
func NadTnpointTnpoint(temp1 Temporal, temp2 Temporal) float64 {
	res := C.nad_tnpoint_tnpoint(temp1.Inner(), temp2.Inner())
	return float64(res)
}


// NaiTnpointGeo wraps MEOS C function nai_tnpoint_geo.
func NaiTnpointGeo(temp Temporal, gs *Geom) TInstant {
	res := C.nai_tnpoint_geo(temp.Inner(), gs._inner)
	return TInstant{_inner: res}
}


// NaiTnpointNpoint wraps MEOS C function nai_tnpoint_npoint.
func NaiTnpointNpoint(temp Temporal, np *Npoint) TInstant {
	res := C.nai_tnpoint_npoint(temp.Inner(), np._inner)
	return TInstant{_inner: res}
}


// NaiTnpointTnpoint wraps MEOS C function nai_tnpoint_tnpoint.
func NaiTnpointTnpoint(temp1 Temporal, temp2 Temporal) TInstant {
	res := C.nai_tnpoint_tnpoint(temp1.Inner(), temp2.Inner())
	return TInstant{_inner: res}
}


// ShortestlineTnpointGeo wraps MEOS C function shortestline_tnpoint_geo.
func ShortestlineTnpointGeo(temp Temporal, gs *Geom) *Geom {
	res := C.shortestline_tnpoint_geo(temp.Inner(), gs._inner)
	return &Geom{_inner: res}
}


// ShortestlineTnpointNpoint wraps MEOS C function shortestline_tnpoint_npoint.
func ShortestlineTnpointNpoint(temp Temporal, np *Npoint) *Geom {
	res := C.shortestline_tnpoint_npoint(temp.Inner(), np._inner)
	return &Geom{_inner: res}
}


// ShortestlineTnpointTnpoint wraps MEOS C function shortestline_tnpoint_tnpoint.
func ShortestlineTnpointTnpoint(temp1 Temporal, temp2 Temporal) *Geom {
	res := C.shortestline_tnpoint_tnpoint(temp1.Inner(), temp2.Inner())
	return &Geom{_inner: res}
}


// TnpointTcentroidTransfn wraps MEOS C function tnpoint_tcentroid_transfn.
func TnpointTcentroidTransfn(state *SkipList, temp Temporal) *SkipList {
	res := C.tnpoint_tcentroid_transfn(state._inner, temp.Inner())
	return &SkipList{_inner: res}
}


// AlwaysEqNpointTnpoint wraps MEOS C function always_eq_npoint_tnpoint.
func AlwaysEqNpointTnpoint(np *Npoint, temp Temporal) int {
	res := C.always_eq_npoint_tnpoint(np._inner, temp.Inner())
	return int(res)
}


// AlwaysEqTnpointNpoint wraps MEOS C function always_eq_tnpoint_npoint.
func AlwaysEqTnpointNpoint(temp Temporal, np *Npoint) int {
	res := C.always_eq_tnpoint_npoint(temp.Inner(), np._inner)
	return int(res)
}


// AlwaysEqTnpointTnpoint wraps MEOS C function always_eq_tnpoint_tnpoint.
func AlwaysEqTnpointTnpoint(temp1 Temporal, temp2 Temporal) int {
	res := C.always_eq_tnpoint_tnpoint(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AlwaysNeNpointTnpoint wraps MEOS C function always_ne_npoint_tnpoint.
func AlwaysNeNpointTnpoint(np *Npoint, temp Temporal) int {
	res := C.always_ne_npoint_tnpoint(np._inner, temp.Inner())
	return int(res)
}


// AlwaysNeTnpointNpoint wraps MEOS C function always_ne_tnpoint_npoint.
func AlwaysNeTnpointNpoint(temp Temporal, np *Npoint) int {
	res := C.always_ne_tnpoint_npoint(temp.Inner(), np._inner)
	return int(res)
}


// AlwaysNeTnpointTnpoint wraps MEOS C function always_ne_tnpoint_tnpoint.
func AlwaysNeTnpointTnpoint(temp1 Temporal, temp2 Temporal) int {
	res := C.always_ne_tnpoint_tnpoint(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverEqNpointTnpoint wraps MEOS C function ever_eq_npoint_tnpoint.
func EverEqNpointTnpoint(np *Npoint, temp Temporal) int {
	res := C.ever_eq_npoint_tnpoint(np._inner, temp.Inner())
	return int(res)
}


// EverEqTnpointNpoint wraps MEOS C function ever_eq_tnpoint_npoint.
func EverEqTnpointNpoint(temp Temporal, np *Npoint) int {
	res := C.ever_eq_tnpoint_npoint(temp.Inner(), np._inner)
	return int(res)
}


// EverEqTnpointTnpoint wraps MEOS C function ever_eq_tnpoint_tnpoint.
func EverEqTnpointTnpoint(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_eq_tnpoint_tnpoint(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverNeNpointTnpoint wraps MEOS C function ever_ne_npoint_tnpoint.
func EverNeNpointTnpoint(np *Npoint, temp Temporal) int {
	res := C.ever_ne_npoint_tnpoint(np._inner, temp.Inner())
	return int(res)
}


// EverNeTnpointNpoint wraps MEOS C function ever_ne_tnpoint_npoint.
func EverNeTnpointNpoint(temp Temporal, np *Npoint) int {
	res := C.ever_ne_tnpoint_npoint(temp.Inner(), np._inner)
	return int(res)
}


// EverNeTnpointTnpoint wraps MEOS C function ever_ne_tnpoint_tnpoint.
func EverNeTnpointTnpoint(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_ne_tnpoint_tnpoint(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TeqTnpointNpoint wraps MEOS C function teq_tnpoint_npoint.
func TeqTnpointNpoint(temp Temporal, np *Npoint) Temporal {
	res := C.teq_tnpoint_npoint(temp.Inner(), np._inner)
	return CreateTemporal(res)
}


// TneTnpointNpoint wraps MEOS C function tne_tnpoint_npoint.
func TneTnpointNpoint(temp Temporal, np *Npoint) Temporal {
	res := C.tne_tnpoint_npoint(temp.Inner(), np._inner)
	return CreateTemporal(res)
}

