package generated

// #include <stddef.h>
import "C"
import (
	"unsafe"

	"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

// CbufferAsEWKT wraps MEOS C function cbuffer_as_ewkt.
func CbufferAsEWKT(cb *Cbuffer, maxdd int) string {
	res := C.cbuffer_as_ewkt(cb._inner, C.int(maxdd))
	return C.GoString(res)
}


// CbufferAsHexwkb wraps MEOS C function cbuffer_as_hexwkb.
func CbufferAsHexwkb(cb *Cbuffer, variant uint8) (string, uint) {
	var _out_size C.size_t
	res := C.cbuffer_as_hexwkb(cb._inner, C.uint8_t(variant), &_out_size)
	return C.GoString(res), uint(_out_size)
}


// CbufferAsText wraps MEOS C function cbuffer_as_text.
func CbufferAsText(cb *Cbuffer, maxdd int) string {
	res := C.cbuffer_as_text(cb._inner, C.int(maxdd))
	return C.GoString(res)
}


// CbufferAsWKB wraps MEOS C function cbuffer_as_wkb.
func CbufferAsWKB(cb *Cbuffer, variant uint8) []uint8 {
	var _out_size_out C.size_t
	res := C.cbuffer_as_wkb(cb._inner, C.uint8_t(variant), &_out_size_out)
	_n := int(_out_size_out)
	_slice := unsafe.Slice((*C.uint8_t)(unsafe.Pointer(res)), _n)
	_out := make([]uint8, _n)
	for _i, _e := range _slice {
		_out[_i] = uint8(_e)
	}
	return _out
}


// CbufferFromHexwkb wraps MEOS C function cbuffer_from_hexwkb.
func CbufferFromHexwkb(hexwkb string) *Cbuffer {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	res := C.cbuffer_from_hexwkb(_c_hexwkb)
	return &Cbuffer{_inner: res}
}


// CbufferFromWKB wraps MEOS C function cbuffer_from_wkb.
func CbufferFromWKB(wkb []byte) *Cbuffer {
	res := C.cbuffer_from_wkb((*C.uint8_t)(unsafe.Pointer(&wkb[0])), C.size_t(len(wkb)))
	return &Cbuffer{_inner: res}
}


// CbufferIn wraps MEOS C function cbuffer_in.
func CbufferIn(str string) *Cbuffer {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.cbuffer_in(_c_str)
	return &Cbuffer{_inner: res}
}


// CbufferOut wraps MEOS C function cbuffer_out.
func CbufferOut(cb *Cbuffer, maxdd int) string {
	res := C.cbuffer_out(cb._inner, C.int(maxdd))
	return C.GoString(res)
}


// CbufferCopy wraps MEOS C function cbuffer_copy.
func CbufferCopy(cb *Cbuffer) *Cbuffer {
	res := C.cbuffer_copy(cb._inner)
	return &Cbuffer{_inner: res}
}


// CbufferMake wraps MEOS C function cbuffer_make.
func CbufferMake(point *Geom, radius float64) *Cbuffer {
	res := C.cbuffer_make(point._inner, C.double(radius))
	return &Cbuffer{_inner: res}
}


// CbufferToGeom wraps MEOS C function cbuffer_to_geom.
func CbufferToGeom(cb *Cbuffer) *Geom {
	res := C.cbuffer_to_geom(cb._inner)
	return &Geom{_inner: res}
}


// CbufferToSTBOX wraps MEOS C function cbuffer_to_stbox.
func CbufferToSTBOX(cb *Cbuffer) *STBox {
	res := C.cbuffer_to_stbox(cb._inner)
	return &STBox{_inner: res}
}


// CbufferarrToGeom wraps MEOS C function cbufferarr_to_geom.
func CbufferarrToGeom(cbarr []*Cbuffer) *Geom {
	_c_cbarr := make([]*C.Cbuffer, len(cbarr))
	for _i, _v := range cbarr { _c_cbarr[_i] = _v._inner }
	res := C.cbufferarr_to_geom((**C.Cbuffer)(unsafe.Pointer(&_c_cbarr[0])), C.int(len(cbarr)))
	return &Geom{_inner: res}
}


// GeomToCbuffer wraps MEOS C function geom_to_cbuffer.
func GeomToCbuffer(gs *Geom) *Cbuffer {
	res := C.geom_to_cbuffer(gs._inner)
	return &Cbuffer{_inner: res}
}


// CbufferHash wraps MEOS C function cbuffer_hash.
func CbufferHash(cb *Cbuffer) uint32 {
	res := C.cbuffer_hash(cb._inner)
	return uint32(res)
}


// CbufferHashExtended wraps MEOS C function cbuffer_hash_extended.
func CbufferHashExtended(cb *Cbuffer, seed uint64) uint64 {
	res := C.cbuffer_hash_extended(cb._inner, C.uint64(seed))
	return uint64(res)
}


// CbufferPoint wraps MEOS C function cbuffer_point.
func CbufferPoint(cb *Cbuffer) *Geom {
	res := C.cbuffer_point(cb._inner)
	return &Geom{_inner: res}
}


// CbufferRadius wraps MEOS C function cbuffer_radius.
func CbufferRadius(cb *Cbuffer) float64 {
	res := C.cbuffer_radius(cb._inner)
	return float64(res)
}


// CbufferRound wraps MEOS C function cbuffer_round.
func CbufferRound(cb *Cbuffer, maxdd int) *Cbuffer {
	res := C.cbuffer_round(cb._inner, C.int(maxdd))
	return &Cbuffer{_inner: res}
}


// CbufferarrRound wraps MEOS C function cbufferarr_round.
func CbufferarrRound(cbarr []*Cbuffer, maxdd int) []*Cbuffer {
	_c_cbarr := make([]*C.Cbuffer, len(cbarr))
	for _i, _v := range cbarr { _c_cbarr[_i] = _v._inner }
	res := C.cbufferarr_round((**C.Cbuffer)(unsafe.Pointer(&_c_cbarr[0])), C.int(len(cbarr)), C.int(maxdd))
	_n := len(cbarr)
	_slice := unsafe.Slice((**C.Cbuffer)(unsafe.Pointer(res)), _n)
	_out := make([]*Cbuffer, _n)
	for _i, _e := range _slice {
		_out[_i] = &Cbuffer{_inner: _e}
	}
	return _out
}


// CbufferSetSRID wraps MEOS C function cbuffer_set_srid.
func CbufferSetSRID(cb *Cbuffer, srid int32) {
	C.cbuffer_set_srid(cb._inner, C.int32_t(srid))
}


// CbufferSRID wraps MEOS C function cbuffer_srid.
func CbufferSRID(cb *Cbuffer) int32 {
	res := C.cbuffer_srid(cb._inner)
	return int32(res)
}


// CbufferTransform wraps MEOS C function cbuffer_transform.
func CbufferTransform(cb *Cbuffer, srid int32) *Cbuffer {
	res := C.cbuffer_transform(cb._inner, C.int32_t(srid))
	return &Cbuffer{_inner: res}
}


// CbufferTransformPipeline wraps MEOS C function cbuffer_transform_pipeline.
func CbufferTransformPipeline(cb *Cbuffer, pipelinestr string, srid int32, is_forward bool) *Cbuffer {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	res := C.cbuffer_transform_pipeline(cb._inner, _c_pipelinestr, C.int32_t(srid), C.bool(is_forward))
	return &Cbuffer{_inner: res}
}


// ContainsCbufferCbuffer wraps MEOS C function contains_cbuffer_cbuffer.
func ContainsCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer) int {
	res := C.contains_cbuffer_cbuffer(cb1._inner, cb2._inner)
	return int(res)
}


// CoversCbufferCbuffer wraps MEOS C function covers_cbuffer_cbuffer.
func CoversCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer) int {
	res := C.covers_cbuffer_cbuffer(cb1._inner, cb2._inner)
	return int(res)
}


// DisjointCbufferCbuffer wraps MEOS C function disjoint_cbuffer_cbuffer.
func DisjointCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer) int {
	res := C.disjoint_cbuffer_cbuffer(cb1._inner, cb2._inner)
	return int(res)
}


// DwithinCbufferCbuffer wraps MEOS C function dwithin_cbuffer_cbuffer.
func DwithinCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer, dist float64) int {
	res := C.dwithin_cbuffer_cbuffer(cb1._inner, cb2._inner, C.double(dist))
	return int(res)
}


// IntersectsCbufferCbuffer wraps MEOS C function intersects_cbuffer_cbuffer.
func IntersectsCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer) int {
	res := C.intersects_cbuffer_cbuffer(cb1._inner, cb2._inner)
	return int(res)
}


// TouchesCbufferCbuffer wraps MEOS C function touches_cbuffer_cbuffer.
func TouchesCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer) int {
	res := C.touches_cbuffer_cbuffer(cb1._inner, cb2._inner)
	return int(res)
}


// CbufferTstzspanToSTBOX wraps MEOS C function cbuffer_tstzspan_to_stbox.
func CbufferTstzspanToSTBOX(cb *Cbuffer, s *Span) *STBox {
	res := C.cbuffer_tstzspan_to_stbox(cb._inner, s._inner)
	return &STBox{_inner: res}
}


// CbufferTimestamptzToSTBOX wraps MEOS C function cbuffer_timestamptz_to_stbox.
func CbufferTimestamptzToSTBOX(cb *Cbuffer, t int64) *STBox {
	res := C.cbuffer_timestamptz_to_stbox(cb._inner, C.TimestampTz(t))
	return &STBox{_inner: res}
}


// DistanceCbufferCbuffer wraps MEOS C function distance_cbuffer_cbuffer.
func DistanceCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer) float64 {
	res := C.distance_cbuffer_cbuffer(cb1._inner, cb2._inner)
	return float64(res)
}


// DistanceCbufferGeo wraps MEOS C function distance_cbuffer_geo.
func DistanceCbufferGeo(cb *Cbuffer, gs *Geom) float64 {
	res := C.distance_cbuffer_geo(cb._inner, gs._inner)
	return float64(res)
}


// DistanceCbufferSTBOX wraps MEOS C function distance_cbuffer_stbox.
func DistanceCbufferSTBOX(cb *Cbuffer, box *STBox) float64 {
	res := C.distance_cbuffer_stbox(cb._inner, box._inner)
	return float64(res)
}


// NadCbufferSTBOX wraps MEOS C function nad_cbuffer_stbox.
func NadCbufferSTBOX(cb *Cbuffer, box *STBox) float64 {
	res := C.nad_cbuffer_stbox(cb._inner, box._inner)
	return float64(res)
}


// CbufferCmp wraps MEOS C function cbuffer_cmp.
func CbufferCmp(cb1 *Cbuffer, cb2 *Cbuffer) int {
	res := C.cbuffer_cmp(cb1._inner, cb2._inner)
	return int(res)
}


// CbufferEq wraps MEOS C function cbuffer_eq.
func CbufferEq(cb1 *Cbuffer, cb2 *Cbuffer) bool {
	res := C.cbuffer_eq(cb1._inner, cb2._inner)
	return bool(res)
}


// CbufferGe wraps MEOS C function cbuffer_ge.
func CbufferGe(cb1 *Cbuffer, cb2 *Cbuffer) bool {
	res := C.cbuffer_ge(cb1._inner, cb2._inner)
	return bool(res)
}


// CbufferGt wraps MEOS C function cbuffer_gt.
func CbufferGt(cb1 *Cbuffer, cb2 *Cbuffer) bool {
	res := C.cbuffer_gt(cb1._inner, cb2._inner)
	return bool(res)
}


// CbufferLe wraps MEOS C function cbuffer_le.
func CbufferLe(cb1 *Cbuffer, cb2 *Cbuffer) bool {
	res := C.cbuffer_le(cb1._inner, cb2._inner)
	return bool(res)
}


// CbufferLt wraps MEOS C function cbuffer_lt.
func CbufferLt(cb1 *Cbuffer, cb2 *Cbuffer) bool {
	res := C.cbuffer_lt(cb1._inner, cb2._inner)
	return bool(res)
}


// CbufferNe wraps MEOS C function cbuffer_ne.
func CbufferNe(cb1 *Cbuffer, cb2 *Cbuffer) bool {
	res := C.cbuffer_ne(cb1._inner, cb2._inner)
	return bool(res)
}


// CbufferNsame wraps MEOS C function cbuffer_nsame.
func CbufferNsame(cb1 *Cbuffer, cb2 *Cbuffer) bool {
	res := C.cbuffer_nsame(cb1._inner, cb2._inner)
	return bool(res)
}


// CbufferSame wraps MEOS C function cbuffer_same.
func CbufferSame(cb1 *Cbuffer, cb2 *Cbuffer) bool {
	res := C.cbuffer_same(cb1._inner, cb2._inner)
	return bool(res)
}


// CbuffersetIn wraps MEOS C function cbufferset_in.
func CbuffersetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.cbufferset_in(_c_str)
	return &Set{_inner: res}
}


// CbuffersetOut wraps MEOS C function cbufferset_out.
func CbuffersetOut(s *Set, maxdd int) string {
	res := C.cbufferset_out(s._inner, C.int(maxdd))
	return C.GoString(res)
}


// CbuffersetMake wraps MEOS C function cbufferset_make.
func CbuffersetMake(values []*Cbuffer) *Set {
	_c_values := make([]*C.Cbuffer, len(values))
	for _i, _v := range values { _c_values[_i] = _v._inner }
	res := C.cbufferset_make((**C.Cbuffer)(unsafe.Pointer(&_c_values[0])), C.int(len(values)))
	return &Set{_inner: res}
}


// CbufferToSet wraps MEOS C function cbuffer_to_set.
func CbufferToSet(cb *Cbuffer) *Set {
	res := C.cbuffer_to_set(cb._inner)
	return &Set{_inner: res}
}


// CbuffersetEndValue wraps MEOS C function cbufferset_end_value.
func CbuffersetEndValue(s *Set) *Cbuffer {
	res := C.cbufferset_end_value(s._inner)
	return &Cbuffer{_inner: res}
}


// CbuffersetStartValue wraps MEOS C function cbufferset_start_value.
func CbuffersetStartValue(s *Set) *Cbuffer {
	res := C.cbufferset_start_value(s._inner)
	return &Cbuffer{_inner: res}
}


// CbuffersetValueN wraps MEOS C function cbufferset_value_n.
func CbuffersetValueN(s *Set, n int) (bool, *Cbuffer) {
	var _out_result *C.Cbuffer
	res := C.cbufferset_value_n(s._inner, C.int(n), &_out_result)
	return bool(res), &Cbuffer{_inner: _out_result}
}


// TODO cbufferset_values: unsupported return type Cbuffer **
// func CbuffersetValues(...) { /* not yet handled by codegen */ }


// CbufferUnionTransfn wraps MEOS C function cbuffer_union_transfn.
func CbufferUnionTransfn(state *Set, cb *Cbuffer) *Set {
	res := C.cbuffer_union_transfn(state._inner, cb._inner)
	return &Set{_inner: res}
}


// ContainedCbufferSet wraps MEOS C function contained_cbuffer_set.
func ContainedCbufferSet(cb *Cbuffer, s *Set) bool {
	res := C.contained_cbuffer_set(cb._inner, s._inner)
	return bool(res)
}


// ContainsSetCbuffer wraps MEOS C function contains_set_cbuffer.
func ContainsSetCbuffer(s *Set, cb *Cbuffer) bool {
	res := C.contains_set_cbuffer(s._inner, cb._inner)
	return bool(res)
}


// IntersectionCbufferSet wraps MEOS C function intersection_cbuffer_set.
func IntersectionCbufferSet(cb *Cbuffer, s *Set) *Set {
	res := C.intersection_cbuffer_set(cb._inner, s._inner)
	return &Set{_inner: res}
}


// IntersectionSetCbuffer wraps MEOS C function intersection_set_cbuffer.
func IntersectionSetCbuffer(s *Set, cb *Cbuffer) *Set {
	res := C.intersection_set_cbuffer(s._inner, cb._inner)
	return &Set{_inner: res}
}


// MinusCbufferSet wraps MEOS C function minus_cbuffer_set.
func MinusCbufferSet(cb *Cbuffer, s *Set) *Set {
	res := C.minus_cbuffer_set(cb._inner, s._inner)
	return &Set{_inner: res}
}


// MinusSetCbuffer wraps MEOS C function minus_set_cbuffer.
func MinusSetCbuffer(s *Set, cb *Cbuffer) *Set {
	res := C.minus_set_cbuffer(s._inner, cb._inner)
	return &Set{_inner: res}
}


// UnionCbufferSet wraps MEOS C function union_cbuffer_set.
func UnionCbufferSet(cb *Cbuffer, s *Set) *Set {
	res := C.union_cbuffer_set(cb._inner, s._inner)
	return &Set{_inner: res}
}


// UnionSetCbuffer wraps MEOS C function union_set_cbuffer.
func UnionSetCbuffer(s *Set, cb *Cbuffer) *Set {
	res := C.union_set_cbuffer(s._inner, cb._inner)
	return &Set{_inner: res}
}


// TcbufferIn wraps MEOS C function tcbuffer_in.
func TcbufferIn(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tcbuffer_in(_c_str)
	return CreateTemporal(res)
}


// TcbufferMake wraps MEOS C function tcbuffer_make.
func TcbufferMake(tpoint Temporal, tfloat Temporal) Temporal {
	res := C.tcbuffer_make(tpoint.Inner(), tfloat.Inner())
	return CreateTemporal(res)
}


// TcbufferPoints wraps MEOS C function tcbuffer_points.
func TcbufferPoints(temp Temporal) *Set {
	res := C.tcbuffer_points(temp.Inner())
	return &Set{_inner: res}
}


// TcbufferRadius wraps MEOS C function tcbuffer_radius.
func TcbufferRadius(temp Temporal) *Set {
	res := C.tcbuffer_radius(temp.Inner())
	return &Set{_inner: res}
}


// TcbufferTravArea wraps MEOS C function tcbuffer_trav_area.
func TcbufferTravArea(temp Temporal, merge_union bool) *Geom {
	res := C.tcbuffer_trav_area(temp.Inner(), C.bool(merge_union))
	return &Geom{_inner: res}
}


// TcbufferToTfloat wraps MEOS C function tcbuffer_to_tfloat.
func TcbufferToTfloat(temp Temporal) Temporal {
	res := C.tcbuffer_to_tfloat(temp.Inner())
	return CreateTemporal(res)
}


// TcbufferToTgeompoint wraps MEOS C function tcbuffer_to_tgeompoint.
func TcbufferToTgeompoint(temp Temporal) Temporal {
	res := C.tcbuffer_to_tgeompoint(temp.Inner())
	return CreateTemporal(res)
}


// TgeometryToTcbuffer wraps MEOS C function tgeometry_to_tcbuffer.
func TgeometryToTcbuffer(temp Temporal) Temporal {
	res := C.tgeometry_to_tcbuffer(temp.Inner())
	return CreateTemporal(res)
}


// TcbufferExpand wraps MEOS C function tcbuffer_expand.
func TcbufferExpand(temp Temporal, dist float64) Temporal {
	res := C.tcbuffer_expand(temp.Inner(), C.double(dist))
	return CreateTemporal(res)
}


// TcbufferAtCbuffer wraps MEOS C function tcbuffer_at_cbuffer.
func TcbufferAtCbuffer(temp Temporal, cb *Cbuffer) Temporal {
	res := C.tcbuffer_at_cbuffer(temp.Inner(), cb._inner)
	return CreateTemporal(res)
}


// TcbufferAtGeom wraps MEOS C function tcbuffer_at_geom.
func TcbufferAtGeom(temp Temporal, gs *Geom) Temporal {
	res := C.tcbuffer_at_geom(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TcbufferAtSTBOX wraps MEOS C function tcbuffer_at_stbox.
func TcbufferAtSTBOX(temp Temporal, box *STBox, border_inc bool) Temporal {
	res := C.tcbuffer_at_stbox(temp.Inner(), box._inner, C.bool(border_inc))
	return CreateTemporal(res)
}


// TcbufferMinusCbuffer wraps MEOS C function tcbuffer_minus_cbuffer.
func TcbufferMinusCbuffer(temp Temporal, cb *Cbuffer) Temporal {
	res := C.tcbuffer_minus_cbuffer(temp.Inner(), cb._inner)
	return CreateTemporal(res)
}


// TcbufferMinusGeom wraps MEOS C function tcbuffer_minus_geom.
func TcbufferMinusGeom(temp Temporal, gs *Geom) Temporal {
	res := C.tcbuffer_minus_geom(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TcbufferMinusSTBOX wraps MEOS C function tcbuffer_minus_stbox.
func TcbufferMinusSTBOX(temp Temporal, box *STBox, border_inc bool) Temporal {
	res := C.tcbuffer_minus_stbox(temp.Inner(), box._inner, C.bool(border_inc))
	return CreateTemporal(res)
}


// TdistanceTcbufferCbuffer wraps MEOS C function tdistance_tcbuffer_cbuffer.
func TdistanceTcbufferCbuffer(temp Temporal, cb *Cbuffer) Temporal {
	res := C.tdistance_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return CreateTemporal(res)
}


// TdistanceTcbufferGeo wraps MEOS C function tdistance_tcbuffer_geo.
func TdistanceTcbufferGeo(temp Temporal, gs *Geom) Temporal {
	res := C.tdistance_tcbuffer_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TdistanceTcbufferTcbuffer wraps MEOS C function tdistance_tcbuffer_tcbuffer.
func TdistanceTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tdistance_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// NadTcbufferCbuffer wraps MEOS C function nad_tcbuffer_cbuffer.
func NadTcbufferCbuffer(temp Temporal, cb *Cbuffer) float64 {
	res := C.nad_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return float64(res)
}


// NadTcbufferGeo wraps MEOS C function nad_tcbuffer_geo.
func NadTcbufferGeo(temp Temporal, gs *Geom) float64 {
	res := C.nad_tcbuffer_geo(temp.Inner(), gs._inner)
	return float64(res)
}


// NadTcbufferSTBOX wraps MEOS C function nad_tcbuffer_stbox.
func NadTcbufferSTBOX(temp Temporal, box *STBox) float64 {
	res := C.nad_tcbuffer_stbox(temp.Inner(), box._inner)
	return float64(res)
}


// NadTcbufferTcbuffer wraps MEOS C function nad_tcbuffer_tcbuffer.
func NadTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) float64 {
	res := C.nad_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return float64(res)
}


// NaiTcbufferCbuffer wraps MEOS C function nai_tcbuffer_cbuffer.
func NaiTcbufferCbuffer(temp Temporal, cb *Cbuffer) TInstant {
	res := C.nai_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return TInstant{_inner: res}
}


// NaiTcbufferGeo wraps MEOS C function nai_tcbuffer_geo.
func NaiTcbufferGeo(temp Temporal, gs *Geom) TInstant {
	res := C.nai_tcbuffer_geo(temp.Inner(), gs._inner)
	return TInstant{_inner: res}
}


// NaiTcbufferTcbuffer wraps MEOS C function nai_tcbuffer_tcbuffer.
func NaiTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) TInstant {
	res := C.nai_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return TInstant{_inner: res}
}


// ShortestlineTcbufferCbuffer wraps MEOS C function shortestline_tcbuffer_cbuffer.
func ShortestlineTcbufferCbuffer(temp Temporal, cb *Cbuffer) *Geom {
	res := C.shortestline_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return &Geom{_inner: res}
}


// ShortestlineTcbufferGeo wraps MEOS C function shortestline_tcbuffer_geo.
func ShortestlineTcbufferGeo(temp Temporal, gs *Geom) *Geom {
	res := C.shortestline_tcbuffer_geo(temp.Inner(), gs._inner)
	return &Geom{_inner: res}
}


// ShortestlineTcbufferTcbuffer wraps MEOS C function shortestline_tcbuffer_tcbuffer.
func ShortestlineTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) *Geom {
	res := C.shortestline_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return &Geom{_inner: res}
}


// AlwaysEqCbufferTcbuffer wraps MEOS C function always_eq_cbuffer_tcbuffer.
func AlwaysEqCbufferTcbuffer(cb *Cbuffer, temp Temporal) int {
	res := C.always_eq_cbuffer_tcbuffer(cb._inner, temp.Inner())
	return int(res)
}


// AlwaysEqTcbufferCbuffer wraps MEOS C function always_eq_tcbuffer_cbuffer.
func AlwaysEqTcbufferCbuffer(temp Temporal, cb *Cbuffer) int {
	res := C.always_eq_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return int(res)
}


// AlwaysEqTcbufferTcbuffer wraps MEOS C function always_eq_tcbuffer_tcbuffer.
func AlwaysEqTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.always_eq_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AlwaysNeCbufferTcbuffer wraps MEOS C function always_ne_cbuffer_tcbuffer.
func AlwaysNeCbufferTcbuffer(cb *Cbuffer, temp Temporal) int {
	res := C.always_ne_cbuffer_tcbuffer(cb._inner, temp.Inner())
	return int(res)
}


// AlwaysNeTcbufferCbuffer wraps MEOS C function always_ne_tcbuffer_cbuffer.
func AlwaysNeTcbufferCbuffer(temp Temporal, cb *Cbuffer) int {
	res := C.always_ne_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return int(res)
}


// AlwaysNeTcbufferTcbuffer wraps MEOS C function always_ne_tcbuffer_tcbuffer.
func AlwaysNeTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.always_ne_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverEqCbufferTcbuffer wraps MEOS C function ever_eq_cbuffer_tcbuffer.
func EverEqCbufferTcbuffer(cb *Cbuffer, temp Temporal) int {
	res := C.ever_eq_cbuffer_tcbuffer(cb._inner, temp.Inner())
	return int(res)
}


// EverEqTcbufferCbuffer wraps MEOS C function ever_eq_tcbuffer_cbuffer.
func EverEqTcbufferCbuffer(temp Temporal, cb *Cbuffer) int {
	res := C.ever_eq_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return int(res)
}


// EverEqTcbufferTcbuffer wraps MEOS C function ever_eq_tcbuffer_tcbuffer.
func EverEqTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_eq_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverNeCbufferTcbuffer wraps MEOS C function ever_ne_cbuffer_tcbuffer.
func EverNeCbufferTcbuffer(cb *Cbuffer, temp Temporal) int {
	res := C.ever_ne_cbuffer_tcbuffer(cb._inner, temp.Inner())
	return int(res)
}


// EverNeTcbufferCbuffer wraps MEOS C function ever_ne_tcbuffer_cbuffer.
func EverNeTcbufferCbuffer(temp Temporal, cb *Cbuffer) int {
	res := C.ever_ne_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return int(res)
}


// EverNeTcbufferTcbuffer wraps MEOS C function ever_ne_tcbuffer_tcbuffer.
func EverNeTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_ne_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TeqCbufferTcbuffer wraps MEOS C function teq_cbuffer_tcbuffer.
func TeqCbufferTcbuffer(cb *Cbuffer, temp Temporal) Temporal {
	res := C.teq_cbuffer_tcbuffer(cb._inner, temp.Inner())
	return CreateTemporal(res)
}


// TeqTcbufferCbuffer wraps MEOS C function teq_tcbuffer_cbuffer.
func TeqTcbufferCbuffer(temp Temporal, cb *Cbuffer) Temporal {
	res := C.teq_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return CreateTemporal(res)
}


// TneCbufferTcbuffer wraps MEOS C function tne_cbuffer_tcbuffer.
func TneCbufferTcbuffer(cb *Cbuffer, temp Temporal) Temporal {
	res := C.tne_cbuffer_tcbuffer(cb._inner, temp.Inner())
	return CreateTemporal(res)
}


// TneTcbufferCbuffer wraps MEOS C function tne_tcbuffer_cbuffer.
func TneTcbufferCbuffer(temp Temporal, cb *Cbuffer) Temporal {
	res := C.tne_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return CreateTemporal(res)
}


// AcontainsCbufferTcbuffer wraps MEOS C function acontains_cbuffer_tcbuffer.
func AcontainsCbufferTcbuffer(cb *Cbuffer, temp Temporal) int {
	res := C.acontains_cbuffer_tcbuffer(cb._inner, temp.Inner())
	return int(res)
}


// AcontainsGeoTcbuffer wraps MEOS C function acontains_geo_tcbuffer.
func AcontainsGeoTcbuffer(gs *Geom, temp Temporal) int {
	res := C.acontains_geo_tcbuffer(gs._inner, temp.Inner())
	return int(res)
}


// AcontainsTcbufferCbuffer wraps MEOS C function acontains_tcbuffer_cbuffer.
func AcontainsTcbufferCbuffer(temp Temporal, cb *Cbuffer) int {
	res := C.acontains_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return int(res)
}


// AcontainsTcbufferGeo wraps MEOS C function acontains_tcbuffer_geo.
func AcontainsTcbufferGeo(temp Temporal, gs *Geom) int {
	res := C.acontains_tcbuffer_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AcoversCbufferTcbuffer wraps MEOS C function acovers_cbuffer_tcbuffer.
func AcoversCbufferTcbuffer(cb *Cbuffer, temp Temporal) int {
	res := C.acovers_cbuffer_tcbuffer(cb._inner, temp.Inner())
	return int(res)
}


// AcoversGeoTcbuffer wraps MEOS C function acovers_geo_tcbuffer.
func AcoversGeoTcbuffer(gs *Geom, temp Temporal) int {
	res := C.acovers_geo_tcbuffer(gs._inner, temp.Inner())
	return int(res)
}


// AcoversTcbufferCbuffer wraps MEOS C function acovers_tcbuffer_cbuffer.
func AcoversTcbufferCbuffer(temp Temporal, cb *Cbuffer) int {
	res := C.acovers_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return int(res)
}


// AcoversTcbufferGeo wraps MEOS C function acovers_tcbuffer_geo.
func AcoversTcbufferGeo(temp Temporal, gs *Geom) int {
	res := C.acovers_tcbuffer_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AdisjointTcbufferGeo wraps MEOS C function adisjoint_tcbuffer_geo.
func AdisjointTcbufferGeo(temp Temporal, gs *Geom) int {
	res := C.adisjoint_tcbuffer_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AdisjointTcbufferCbuffer wraps MEOS C function adisjoint_tcbuffer_cbuffer.
func AdisjointTcbufferCbuffer(temp Temporal, cb *Cbuffer) int {
	res := C.adisjoint_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return int(res)
}


// AdisjointTcbufferTcbuffer wraps MEOS C function adisjoint_tcbuffer_tcbuffer.
func AdisjointTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.adisjoint_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AdwithinTcbufferGeo wraps MEOS C function adwithin_tcbuffer_geo.
func AdwithinTcbufferGeo(temp Temporal, gs *Geom, dist float64) int {
	res := C.adwithin_tcbuffer_geo(temp.Inner(), gs._inner, C.double(dist))
	return int(res)
}


// AdwithinTcbufferCbuffer wraps MEOS C function adwithin_tcbuffer_cbuffer.
func AdwithinTcbufferCbuffer(temp Temporal, cb *Cbuffer, dist float64) int {
	res := C.adwithin_tcbuffer_cbuffer(temp.Inner(), cb._inner, C.double(dist))
	return int(res)
}


// AdwithinTcbufferTcbuffer wraps MEOS C function adwithin_tcbuffer_tcbuffer.
func AdwithinTcbufferTcbuffer(temp1 Temporal, temp2 Temporal, dist float64) int {
	res := C.adwithin_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner(), C.double(dist))
	return int(res)
}


// AintersectsTcbufferGeo wraps MEOS C function aintersects_tcbuffer_geo.
func AintersectsTcbufferGeo(temp Temporal, gs *Geom) int {
	res := C.aintersects_tcbuffer_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AintersectsTcbufferCbuffer wraps MEOS C function aintersects_tcbuffer_cbuffer.
func AintersectsTcbufferCbuffer(temp Temporal, cb *Cbuffer) int {
	res := C.aintersects_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return int(res)
}


// AintersectsTcbufferTcbuffer wraps MEOS C function aintersects_tcbuffer_tcbuffer.
func AintersectsTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.aintersects_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AtouchesTcbufferGeo wraps MEOS C function atouches_tcbuffer_geo.
func AtouchesTcbufferGeo(temp Temporal, gs *Geom) int {
	res := C.atouches_tcbuffer_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AtouchesTcbufferCbuffer wraps MEOS C function atouches_tcbuffer_cbuffer.
func AtouchesTcbufferCbuffer(temp Temporal, cb *Cbuffer) int {
	res := C.atouches_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return int(res)
}


// AtouchesTcbufferTcbuffer wraps MEOS C function atouches_tcbuffer_tcbuffer.
func AtouchesTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.atouches_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EcontainsCbufferTcbuffer wraps MEOS C function econtains_cbuffer_tcbuffer.
func EcontainsCbufferTcbuffer(cb *Cbuffer, temp Temporal) int {
	res := C.econtains_cbuffer_tcbuffer(cb._inner, temp.Inner())
	return int(res)
}


// EcontainsTcbufferCbuffer wraps MEOS C function econtains_tcbuffer_cbuffer.
func EcontainsTcbufferCbuffer(temp Temporal, cb *Cbuffer) int {
	res := C.econtains_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return int(res)
}


// EcontainsTcbufferGeo wraps MEOS C function econtains_tcbuffer_geo.
func EcontainsTcbufferGeo(temp Temporal, gs *Geom) int {
	res := C.econtains_tcbuffer_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EcoversCbufferTcbuffer wraps MEOS C function ecovers_cbuffer_tcbuffer.
func EcoversCbufferTcbuffer(cb *Cbuffer, temp Temporal) int {
	res := C.ecovers_cbuffer_tcbuffer(cb._inner, temp.Inner())
	return int(res)
}


// EcoversTcbufferCbuffer wraps MEOS C function ecovers_tcbuffer_cbuffer.
func EcoversTcbufferCbuffer(temp Temporal, cb *Cbuffer) int {
	res := C.ecovers_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return int(res)
}


// EcoversTcbufferGeo wraps MEOS C function ecovers_tcbuffer_geo.
func EcoversTcbufferGeo(temp Temporal, gs *Geom) int {
	res := C.ecovers_tcbuffer_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EcoversTcbufferTcbuffer wraps MEOS C function ecovers_tcbuffer_tcbuffer.
func EcoversTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.ecovers_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EdisjointTcbufferGeo wraps MEOS C function edisjoint_tcbuffer_geo.
func EdisjointTcbufferGeo(temp Temporal, gs *Geom) int {
	res := C.edisjoint_tcbuffer_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EdisjointTcbufferCbuffer wraps MEOS C function edisjoint_tcbuffer_cbuffer.
func EdisjointTcbufferCbuffer(temp Temporal, cb *Cbuffer) int {
	res := C.edisjoint_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return int(res)
}


// EdwithinTcbufferGeo wraps MEOS C function edwithin_tcbuffer_geo.
func EdwithinTcbufferGeo(temp Temporal, gs *Geom, dist float64) int {
	res := C.edwithin_tcbuffer_geo(temp.Inner(), gs._inner, C.double(dist))
	return int(res)
}


// EdwithinTcbufferCbuffer wraps MEOS C function edwithin_tcbuffer_cbuffer.
func EdwithinTcbufferCbuffer(temp Temporal, cb *Cbuffer, dist float64) int {
	res := C.edwithin_tcbuffer_cbuffer(temp.Inner(), cb._inner, C.double(dist))
	return int(res)
}


// EdwithinTcbufferTcbuffer wraps MEOS C function edwithin_tcbuffer_tcbuffer.
func EdwithinTcbufferTcbuffer(temp1 Temporal, temp2 Temporal, dist float64) int {
	res := C.edwithin_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner(), C.double(dist))
	return int(res)
}


// EintersectsTcbufferGeo wraps MEOS C function eintersects_tcbuffer_geo.
func EintersectsTcbufferGeo(temp Temporal, gs *Geom) int {
	res := C.eintersects_tcbuffer_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EintersectsTcbufferCbuffer wraps MEOS C function eintersects_tcbuffer_cbuffer.
func EintersectsTcbufferCbuffer(temp Temporal, cb *Cbuffer) int {
	res := C.eintersects_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return int(res)
}


// EintersectsTcbufferTcbuffer wraps MEOS C function eintersects_tcbuffer_tcbuffer.
func EintersectsTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.eintersects_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EtouchesTcbufferGeo wraps MEOS C function etouches_tcbuffer_geo.
func EtouchesTcbufferGeo(temp Temporal, gs *Geom) int {
	res := C.etouches_tcbuffer_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EtouchesTcbufferCbuffer wraps MEOS C function etouches_tcbuffer_cbuffer.
func EtouchesTcbufferCbuffer(temp Temporal, cb *Cbuffer) int {
	res := C.etouches_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return int(res)
}


// EtouchesTcbufferTcbuffer wraps MEOS C function etouches_tcbuffer_tcbuffer.
func EtouchesTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.etouches_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TcontainsCbufferTcbuffer wraps MEOS C function tcontains_cbuffer_tcbuffer.
func TcontainsCbufferTcbuffer(cb *Cbuffer, temp Temporal) Temporal {
	res := C.tcontains_cbuffer_tcbuffer(cb._inner, temp.Inner())
	return CreateTemporal(res)
}


// TcontainsGeoTcbuffer wraps MEOS C function tcontains_geo_tcbuffer.
func TcontainsGeoTcbuffer(gs *Geom, temp Temporal) Temporal {
	res := C.tcontains_geo_tcbuffer(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TcontainsTcbufferGeo wraps MEOS C function tcontains_tcbuffer_geo.
func TcontainsTcbufferGeo(temp Temporal, gs *Geom) Temporal {
	res := C.tcontains_tcbuffer_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TcontainsTcbufferCbuffer wraps MEOS C function tcontains_tcbuffer_cbuffer.
func TcontainsTcbufferCbuffer(temp Temporal, cb *Cbuffer) Temporal {
	res := C.tcontains_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return CreateTemporal(res)
}


// TcontainsTcbufferTcbuffer wraps MEOS C function tcontains_tcbuffer_tcbuffer.
func TcontainsTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tcontains_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TcoversCbufferTcbuffer wraps MEOS C function tcovers_cbuffer_tcbuffer.
func TcoversCbufferTcbuffer(cb *Cbuffer, temp Temporal) Temporal {
	res := C.tcovers_cbuffer_tcbuffer(cb._inner, temp.Inner())
	return CreateTemporal(res)
}


// TcoversGeoTcbuffer wraps MEOS C function tcovers_geo_tcbuffer.
func TcoversGeoTcbuffer(gs *Geom, temp Temporal) Temporal {
	res := C.tcovers_geo_tcbuffer(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TcoversTcbufferGeo wraps MEOS C function tcovers_tcbuffer_geo.
func TcoversTcbufferGeo(temp Temporal, gs *Geom) Temporal {
	res := C.tcovers_tcbuffer_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TcoversTcbufferCbuffer wraps MEOS C function tcovers_tcbuffer_cbuffer.
func TcoversTcbufferCbuffer(temp Temporal, cb *Cbuffer) Temporal {
	res := C.tcovers_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return CreateTemporal(res)
}


// TcoversTcbufferTcbuffer wraps MEOS C function tcovers_tcbuffer_tcbuffer.
func TcoversTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tcovers_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TdwithinGeoTcbuffer wraps MEOS C function tdwithin_geo_tcbuffer.
func TdwithinGeoTcbuffer(gs *Geom, temp Temporal, dist float64) Temporal {
	res := C.tdwithin_geo_tcbuffer(gs._inner, temp.Inner(), C.double(dist))
	return CreateTemporal(res)
}


// TdwithinTcbufferGeo wraps MEOS C function tdwithin_tcbuffer_geo.
func TdwithinTcbufferGeo(temp Temporal, gs *Geom, dist float64) Temporal {
	res := C.tdwithin_tcbuffer_geo(temp.Inner(), gs._inner, C.double(dist))
	return CreateTemporal(res)
}


// TdwithinTcbufferCbuffer wraps MEOS C function tdwithin_tcbuffer_cbuffer.
func TdwithinTcbufferCbuffer(temp Temporal, cb *Cbuffer, dist float64) Temporal {
	res := C.tdwithin_tcbuffer_cbuffer(temp.Inner(), cb._inner, C.double(dist))
	return CreateTemporal(res)
}


// TdwithinTcbufferTcbuffer wraps MEOS C function tdwithin_tcbuffer_tcbuffer.
func TdwithinTcbufferTcbuffer(temp1 Temporal, temp2 Temporal, dist float64) Temporal {
	res := C.tdwithin_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner(), C.double(dist))
	return CreateTemporal(res)
}


// TdisjointCbufferTcbuffer wraps MEOS C function tdisjoint_cbuffer_tcbuffer.
func TdisjointCbufferTcbuffer(cb *Cbuffer, temp Temporal) Temporal {
	res := C.tdisjoint_cbuffer_tcbuffer(cb._inner, temp.Inner())
	return CreateTemporal(res)
}


// TdisjointGeoTcbuffer wraps MEOS C function tdisjoint_geo_tcbuffer.
func TdisjointGeoTcbuffer(gs *Geom, temp Temporal) Temporal {
	res := C.tdisjoint_geo_tcbuffer(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TdisjointTcbufferGeo wraps MEOS C function tdisjoint_tcbuffer_geo.
func TdisjointTcbufferGeo(temp Temporal, gs *Geom) Temporal {
	res := C.tdisjoint_tcbuffer_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TdisjointTcbufferCbuffer wraps MEOS C function tdisjoint_tcbuffer_cbuffer.
func TdisjointTcbufferCbuffer(temp Temporal, cb *Cbuffer) Temporal {
	res := C.tdisjoint_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return CreateTemporal(res)
}


// TdisjointTcbufferTcbuffer wraps MEOS C function tdisjoint_tcbuffer_tcbuffer.
func TdisjointTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tdisjoint_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TintersectsCbufferTcbuffer wraps MEOS C function tintersects_cbuffer_tcbuffer.
func TintersectsCbufferTcbuffer(cb *Cbuffer, temp Temporal) Temporal {
	res := C.tintersects_cbuffer_tcbuffer(cb._inner, temp.Inner())
	return CreateTemporal(res)
}


// TintersectsGeoTcbuffer wraps MEOS C function tintersects_geo_tcbuffer.
func TintersectsGeoTcbuffer(gs *Geom, temp Temporal) Temporal {
	res := C.tintersects_geo_tcbuffer(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TintersectsTcbufferGeo wraps MEOS C function tintersects_tcbuffer_geo.
func TintersectsTcbufferGeo(temp Temporal, gs *Geom) Temporal {
	res := C.tintersects_tcbuffer_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TintersectsTcbufferCbuffer wraps MEOS C function tintersects_tcbuffer_cbuffer.
func TintersectsTcbufferCbuffer(temp Temporal, cb *Cbuffer) Temporal {
	res := C.tintersects_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return CreateTemporal(res)
}


// TintersectsTcbufferTcbuffer wraps MEOS C function tintersects_tcbuffer_tcbuffer.
func TintersectsTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tintersects_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TtouchesGeoTcbuffer wraps MEOS C function ttouches_geo_tcbuffer.
func TtouchesGeoTcbuffer(gs *Geom, temp Temporal) Temporal {
	res := C.ttouches_geo_tcbuffer(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TtouchesTcbufferGeo wraps MEOS C function ttouches_tcbuffer_geo.
func TtouchesTcbufferGeo(temp Temporal, gs *Geom) Temporal {
	res := C.ttouches_tcbuffer_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TtouchesCbufferTcbuffer wraps MEOS C function ttouches_cbuffer_tcbuffer.
func TtouchesCbufferTcbuffer(cb *Cbuffer, temp Temporal) Temporal {
	res := C.ttouches_cbuffer_tcbuffer(cb._inner, temp.Inner())
	return CreateTemporal(res)
}


// TtouchesTcbufferCbuffer wraps MEOS C function ttouches_tcbuffer_cbuffer.
func TtouchesTcbufferCbuffer(temp Temporal, cb *Cbuffer) Temporal {
	res := C.ttouches_tcbuffer_cbuffer(temp.Inner(), cb._inner)
	return CreateTemporal(res)
}


// TtouchesTcbufferTcbuffer wraps MEOS C function ttouches_tcbuffer_tcbuffer.
func TtouchesTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.ttouches_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}

