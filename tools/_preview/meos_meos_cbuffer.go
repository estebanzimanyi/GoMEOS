package generated

// #include <stddef.h>
import "C"
import (
	"unsafe"

	"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

// TODO cbuffer_as_ewkt: unsupported param const Cbuffer *
// func CbufferAsEWKT(...) { /* not yet handled by codegen */ }


// TODO cbuffer_as_hexwkb: unsupported param const Cbuffer *
// func CbufferAsHexwkb(...) { /* not yet handled by codegen */ }


// TODO cbuffer_as_text: unsupported param const Cbuffer *
// func CbufferAsText(...) { /* not yet handled by codegen */ }


// TODO cbuffer_as_wkb: unsupported param const Cbuffer *
// func CbufferAsWKB(...) { /* not yet handled by codegen */ }


// TODO cbuffer_from_hexwkb: unsupported return type Cbuffer *
// func CbufferFromHexwkb(...) { /* not yet handled by codegen */ }


// TODO cbuffer_from_wkb: unsupported return type Cbuffer *
// func CbufferFromWKB(...) { /* not yet handled by codegen */ }


// TODO cbuffer_in: unsupported return type Cbuffer *
// func CbufferIn(...) { /* not yet handled by codegen */ }


// TODO cbuffer_out: unsupported param const Cbuffer *
// func CbufferOut(...) { /* not yet handled by codegen */ }


// TODO cbuffer_copy: unsupported return type Cbuffer *
// func CbufferCopy(...) { /* not yet handled by codegen */ }


// TODO cbuffer_make: unsupported return type Cbuffer *
// func CbufferMake(...) { /* not yet handled by codegen */ }


// TODO cbuffer_to_geom: unsupported return type int *
// func CbufferToGeom(...) { /* not yet handled by codegen */ }


// TODO cbuffer_to_stbox: unsupported param const Cbuffer *
// func CbufferToSTBOX(...) { /* not yet handled by codegen */ }


// TODO cbufferarr_to_geom: unsupported return type int *
// func CbufferarrToGeom(...) { /* not yet handled by codegen */ }


// TODO geom_to_cbuffer: unsupported return type Cbuffer *
// func GeomToCbuffer(...) { /* not yet handled by codegen */ }


// TODO cbuffer_hash: unsupported param const Cbuffer *
// func CbufferHash(...) { /* not yet handled by codegen */ }


// TODO cbuffer_hash_extended: unsupported param const Cbuffer *
// func CbufferHashExtended(...) { /* not yet handled by codegen */ }


// TODO cbuffer_point: unsupported return type int *
// func CbufferPoint(...) { /* not yet handled by codegen */ }


// TODO cbuffer_radius: unsupported param const Cbuffer *
// func CbufferRadius(...) { /* not yet handled by codegen */ }


// TODO cbuffer_round: unsupported return type Cbuffer *
// func CbufferRound(...) { /* not yet handled by codegen */ }


// TODO cbufferarr_round: unsupported return type Cbuffer **
// func CbufferarrRound(...) { /* not yet handled by codegen */ }


// TODO cbuffer_set_srid: unsupported param Cbuffer *
// func CbufferSetSRID(...) { /* not yet handled by codegen */ }


// TODO cbuffer_srid: unsupported param const Cbuffer *
// func CbufferSRID(...) { /* not yet handled by codegen */ }


// TODO cbuffer_transform: unsupported return type Cbuffer *
// func CbufferTransform(...) { /* not yet handled by codegen */ }


// TODO cbuffer_transform_pipeline: unsupported return type Cbuffer *
// func CbufferTransformPipeline(...) { /* not yet handled by codegen */ }


// TODO contains_cbuffer_cbuffer: unsupported param const Cbuffer *
// func ContainsCbufferCbuffer(...) { /* not yet handled by codegen */ }


// TODO covers_cbuffer_cbuffer: unsupported param const Cbuffer *
// func CoversCbufferCbuffer(...) { /* not yet handled by codegen */ }


// TODO disjoint_cbuffer_cbuffer: unsupported param const Cbuffer *
// func DisjointCbufferCbuffer(...) { /* not yet handled by codegen */ }


// TODO dwithin_cbuffer_cbuffer: unsupported param const Cbuffer *
// func DwithinCbufferCbuffer(...) { /* not yet handled by codegen */ }


// TODO intersects_cbuffer_cbuffer: unsupported param const Cbuffer *
// func IntersectsCbufferCbuffer(...) { /* not yet handled by codegen */ }


// TODO touches_cbuffer_cbuffer: unsupported param const Cbuffer *
// func TouchesCbufferCbuffer(...) { /* not yet handled by codegen */ }


// TODO cbuffer_tstzspan_to_stbox: unsupported param const Cbuffer *
// func CbufferTstzspanToSTBOX(...) { /* not yet handled by codegen */ }


// TODO cbuffer_timestamptz_to_stbox: unsupported param const Cbuffer *
// func CbufferTimestamptzToSTBOX(...) { /* not yet handled by codegen */ }


// TODO distance_cbuffer_cbuffer: unsupported param const Cbuffer *
// func DistanceCbufferCbuffer(...) { /* not yet handled by codegen */ }


// TODO distance_cbuffer_geo: unsupported param const Cbuffer *
// func DistanceCbufferGeo(...) { /* not yet handled by codegen */ }


// TODO distance_cbuffer_stbox: unsupported param const Cbuffer *
// func DistanceCbufferSTBOX(...) { /* not yet handled by codegen */ }


// TODO nad_cbuffer_stbox: unsupported param const Cbuffer *
// func NadCbufferSTBOX(...) { /* not yet handled by codegen */ }


// TODO cbuffer_cmp: unsupported param const Cbuffer *
// func CbufferCmp(...) { /* not yet handled by codegen */ }


// TODO cbuffer_eq: unsupported param const Cbuffer *
// func CbufferEq(...) { /* not yet handled by codegen */ }


// TODO cbuffer_ge: unsupported param const Cbuffer *
// func CbufferGe(...) { /* not yet handled by codegen */ }


// TODO cbuffer_gt: unsupported param const Cbuffer *
// func CbufferGt(...) { /* not yet handled by codegen */ }


// TODO cbuffer_le: unsupported param const Cbuffer *
// func CbufferLe(...) { /* not yet handled by codegen */ }


// TODO cbuffer_lt: unsupported param const Cbuffer *
// func CbufferLt(...) { /* not yet handled by codegen */ }


// TODO cbuffer_ne: unsupported param const Cbuffer *
// func CbufferNe(...) { /* not yet handled by codegen */ }


// TODO cbuffer_nsame: unsupported param const Cbuffer *
// func CbufferNsame(...) { /* not yet handled by codegen */ }


// TODO cbuffer_same: unsupported param const Cbuffer *
// func CbufferSame(...) { /* not yet handled by codegen */ }


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


// TODO cbufferset_make: unsupported param Cbuffer **
// func CbuffersetMake(...) { /* not yet handled by codegen */ }


// TODO cbuffer_to_set: unsupported param const Cbuffer *
// func CbufferToSet(...) { /* not yet handled by codegen */ }


// TODO cbufferset_end_value: unsupported return type Cbuffer *
// func CbuffersetEndValue(...) { /* not yet handled by codegen */ }


// TODO cbufferset_start_value: unsupported return type Cbuffer *
// func CbuffersetStartValue(...) { /* not yet handled by codegen */ }


// TODO cbufferset_value_n: unhandled OUTPUT_SCALAR shape Cbuffer **
// func CbuffersetValueN(...) { /* not yet handled by codegen */ }


// TODO cbufferset_values: unsupported return type Cbuffer **
// func CbuffersetValues(...) { /* not yet handled by codegen */ }


// TODO cbuffer_union_transfn: unsupported param const Cbuffer *
// func CbufferUnionTransfn(...) { /* not yet handled by codegen */ }


// TODO contained_cbuffer_set: unsupported param const Cbuffer *
// func ContainedCbufferSet(...) { /* not yet handled by codegen */ }


// TODO contains_set_cbuffer: unsupported param Cbuffer *
// func ContainsSetCbuffer(...) { /* not yet handled by codegen */ }


// TODO intersection_cbuffer_set: unsupported param const Cbuffer *
// func IntersectionCbufferSet(...) { /* not yet handled by codegen */ }


// TODO intersection_set_cbuffer: unsupported param const Cbuffer *
// func IntersectionSetCbuffer(...) { /* not yet handled by codegen */ }


// TODO minus_cbuffer_set: unsupported param const Cbuffer *
// func MinusCbufferSet(...) { /* not yet handled by codegen */ }


// TODO minus_set_cbuffer: unsupported param const Cbuffer *
// func MinusSetCbuffer(...) { /* not yet handled by codegen */ }


// TODO union_cbuffer_set: unsupported param const Cbuffer *
// func UnionCbufferSet(...) { /* not yet handled by codegen */ }


// TODO union_set_cbuffer: unsupported param const Cbuffer *
// func UnionSetCbuffer(...) { /* not yet handled by codegen */ }


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


// TODO tcbuffer_trav_area: unsupported return type int *
// func TcbufferTravArea(...) { /* not yet handled by codegen */ }


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


// TODO tcbuffer_at_cbuffer: unsupported param const Cbuffer *
// func TcbufferAtCbuffer(...) { /* not yet handled by codegen */ }


// TODO tcbuffer_at_geom: unsupported param const int *
// func TcbufferAtGeom(...) { /* not yet handled by codegen */ }


// TcbufferAtSTBOX wraps MEOS C function tcbuffer_at_stbox.
func TcbufferAtSTBOX(temp Temporal, box *STBox, border_inc bool) Temporal {
	res := C.tcbuffer_at_stbox(temp.Inner(), box._inner, C.bool(border_inc))
	return CreateTemporal(res)
}


// TODO tcbuffer_minus_cbuffer: unsupported param const Cbuffer *
// func TcbufferMinusCbuffer(...) { /* not yet handled by codegen */ }


// TODO tcbuffer_minus_geom: unsupported param const int *
// func TcbufferMinusGeom(...) { /* not yet handled by codegen */ }


// TcbufferMinusSTBOX wraps MEOS C function tcbuffer_minus_stbox.
func TcbufferMinusSTBOX(temp Temporal, box *STBox, border_inc bool) Temporal {
	res := C.tcbuffer_minus_stbox(temp.Inner(), box._inner, C.bool(border_inc))
	return CreateTemporal(res)
}


// TODO tdistance_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func TdistanceTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// TODO tdistance_tcbuffer_geo: unsupported param const int *
// func TdistanceTcbufferGeo(...) { /* not yet handled by codegen */ }


// TdistanceTcbufferTcbuffer wraps MEOS C function tdistance_tcbuffer_tcbuffer.
func TdistanceTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tdistance_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TODO nad_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func NadTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// TODO nad_tcbuffer_geo: unsupported param const int *
// func NadTcbufferGeo(...) { /* not yet handled by codegen */ }


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


// TODO nai_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func NaiTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// TODO nai_tcbuffer_geo: unsupported param const int *
// func NaiTcbufferGeo(...) { /* not yet handled by codegen */ }


// NaiTcbufferTcbuffer wraps MEOS C function nai_tcbuffer_tcbuffer.
func NaiTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) TInstant {
	res := C.nai_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return TInstant{_inner: res}
}


// TODO shortestline_tcbuffer_cbuffer: unsupported return type int *
// func ShortestlineTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// TODO shortestline_tcbuffer_geo: unsupported return type int *
// func ShortestlineTcbufferGeo(...) { /* not yet handled by codegen */ }


// TODO shortestline_tcbuffer_tcbuffer: unsupported return type int *
// func ShortestlineTcbufferTcbuffer(...) { /* not yet handled by codegen */ }


// TODO always_eq_cbuffer_tcbuffer: unsupported param const Cbuffer *
// func AlwaysEqCbufferTcbuffer(...) { /* not yet handled by codegen */ }


// TODO always_eq_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func AlwaysEqTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// AlwaysEqTcbufferTcbuffer wraps MEOS C function always_eq_tcbuffer_tcbuffer.
func AlwaysEqTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.always_eq_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO always_ne_cbuffer_tcbuffer: unsupported param const Cbuffer *
// func AlwaysNeCbufferTcbuffer(...) { /* not yet handled by codegen */ }


// TODO always_ne_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func AlwaysNeTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// AlwaysNeTcbufferTcbuffer wraps MEOS C function always_ne_tcbuffer_tcbuffer.
func AlwaysNeTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.always_ne_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO ever_eq_cbuffer_tcbuffer: unsupported param const Cbuffer *
// func EverEqCbufferTcbuffer(...) { /* not yet handled by codegen */ }


// TODO ever_eq_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func EverEqTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// EverEqTcbufferTcbuffer wraps MEOS C function ever_eq_tcbuffer_tcbuffer.
func EverEqTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_eq_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO ever_ne_cbuffer_tcbuffer: unsupported param const Cbuffer *
// func EverNeCbufferTcbuffer(...) { /* not yet handled by codegen */ }


// TODO ever_ne_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func EverNeTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// EverNeTcbufferTcbuffer wraps MEOS C function ever_ne_tcbuffer_tcbuffer.
func EverNeTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_ne_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO teq_cbuffer_tcbuffer: unsupported param const Cbuffer *
// func TeqCbufferTcbuffer(...) { /* not yet handled by codegen */ }


// TODO teq_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func TeqTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// TODO tne_cbuffer_tcbuffer: unsupported param const Cbuffer *
// func TneCbufferTcbuffer(...) { /* not yet handled by codegen */ }


// TODO tne_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func TneTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// TODO acontains_cbuffer_tcbuffer: unsupported param const Cbuffer *
// func AcontainsCbufferTcbuffer(...) { /* not yet handled by codegen */ }


// TODO acontains_geo_tcbuffer: unsupported param const int *
// func AcontainsGeoTcbuffer(...) { /* not yet handled by codegen */ }


// TODO acontains_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func AcontainsTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// TODO acontains_tcbuffer_geo: unsupported param const int *
// func AcontainsTcbufferGeo(...) { /* not yet handled by codegen */ }


// TODO acovers_cbuffer_tcbuffer: unsupported param const Cbuffer *
// func AcoversCbufferTcbuffer(...) { /* not yet handled by codegen */ }


// TODO acovers_geo_tcbuffer: unsupported param const int *
// func AcoversGeoTcbuffer(...) { /* not yet handled by codegen */ }


// TODO acovers_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func AcoversTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// TODO acovers_tcbuffer_geo: unsupported param const int *
// func AcoversTcbufferGeo(...) { /* not yet handled by codegen */ }


// TODO adisjoint_tcbuffer_geo: unsupported param const int *
// func AdisjointTcbufferGeo(...) { /* not yet handled by codegen */ }


// TODO adisjoint_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func AdisjointTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// AdisjointTcbufferTcbuffer wraps MEOS C function adisjoint_tcbuffer_tcbuffer.
func AdisjointTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.adisjoint_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO adwithin_tcbuffer_geo: unsupported param const int *
// func AdwithinTcbufferGeo(...) { /* not yet handled by codegen */ }


// TODO adwithin_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func AdwithinTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// AdwithinTcbufferTcbuffer wraps MEOS C function adwithin_tcbuffer_tcbuffer.
func AdwithinTcbufferTcbuffer(temp1 Temporal, temp2 Temporal, dist float64) int {
	res := C.adwithin_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner(), C.double(dist))
	return int(res)
}


// TODO aintersects_tcbuffer_geo: unsupported param const int *
// func AintersectsTcbufferGeo(...) { /* not yet handled by codegen */ }


// TODO aintersects_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func AintersectsTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// AintersectsTcbufferTcbuffer wraps MEOS C function aintersects_tcbuffer_tcbuffer.
func AintersectsTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.aintersects_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO atouches_tcbuffer_geo: unsupported param const int *
// func AtouchesTcbufferGeo(...) { /* not yet handled by codegen */ }


// TODO atouches_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func AtouchesTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// AtouchesTcbufferTcbuffer wraps MEOS C function atouches_tcbuffer_tcbuffer.
func AtouchesTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.atouches_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO econtains_cbuffer_tcbuffer: unsupported param const Cbuffer *
// func EcontainsCbufferTcbuffer(...) { /* not yet handled by codegen */ }


// TODO econtains_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func EcontainsTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// TODO econtains_tcbuffer_geo: unsupported param const int *
// func EcontainsTcbufferGeo(...) { /* not yet handled by codegen */ }


// TODO ecovers_cbuffer_tcbuffer: unsupported param const Cbuffer *
// func EcoversCbufferTcbuffer(...) { /* not yet handled by codegen */ }


// TODO ecovers_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func EcoversTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// TODO ecovers_tcbuffer_geo: unsupported param const int *
// func EcoversTcbufferGeo(...) { /* not yet handled by codegen */ }


// EcoversTcbufferTcbuffer wraps MEOS C function ecovers_tcbuffer_tcbuffer.
func EcoversTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.ecovers_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO edisjoint_tcbuffer_geo: unsupported param const int *
// func EdisjointTcbufferGeo(...) { /* not yet handled by codegen */ }


// TODO edisjoint_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func EdisjointTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// TODO edwithin_tcbuffer_geo: unsupported param const int *
// func EdwithinTcbufferGeo(...) { /* not yet handled by codegen */ }


// TODO edwithin_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func EdwithinTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// EdwithinTcbufferTcbuffer wraps MEOS C function edwithin_tcbuffer_tcbuffer.
func EdwithinTcbufferTcbuffer(temp1 Temporal, temp2 Temporal, dist float64) int {
	res := C.edwithin_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner(), C.double(dist))
	return int(res)
}


// TODO eintersects_tcbuffer_geo: unsupported param const int *
// func EintersectsTcbufferGeo(...) { /* not yet handled by codegen */ }


// TODO eintersects_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func EintersectsTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// EintersectsTcbufferTcbuffer wraps MEOS C function eintersects_tcbuffer_tcbuffer.
func EintersectsTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.eintersects_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO etouches_tcbuffer_geo: unsupported param const int *
// func EtouchesTcbufferGeo(...) { /* not yet handled by codegen */ }


// TODO etouches_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func EtouchesTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// EtouchesTcbufferTcbuffer wraps MEOS C function etouches_tcbuffer_tcbuffer.
func EtouchesTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) int {
	res := C.etouches_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO tcontains_cbuffer_tcbuffer: unsupported param const Cbuffer *
// func TcontainsCbufferTcbuffer(...) { /* not yet handled by codegen */ }


// TODO tcontains_geo_tcbuffer: unsupported param const int *
// func TcontainsGeoTcbuffer(...) { /* not yet handled by codegen */ }


// TODO tcontains_tcbuffer_geo: unsupported param const int *
// func TcontainsTcbufferGeo(...) { /* not yet handled by codegen */ }


// TODO tcontains_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func TcontainsTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// TcontainsTcbufferTcbuffer wraps MEOS C function tcontains_tcbuffer_tcbuffer.
func TcontainsTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tcontains_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TODO tcovers_cbuffer_tcbuffer: unsupported param const Cbuffer *
// func TcoversCbufferTcbuffer(...) { /* not yet handled by codegen */ }


// TODO tcovers_geo_tcbuffer: unsupported param const int *
// func TcoversGeoTcbuffer(...) { /* not yet handled by codegen */ }


// TODO tcovers_tcbuffer_geo: unsupported param const int *
// func TcoversTcbufferGeo(...) { /* not yet handled by codegen */ }


// TODO tcovers_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func TcoversTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// TcoversTcbufferTcbuffer wraps MEOS C function tcovers_tcbuffer_tcbuffer.
func TcoversTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tcovers_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TODO tdwithin_geo_tcbuffer: unsupported param const int *
// func TdwithinGeoTcbuffer(...) { /* not yet handled by codegen */ }


// TODO tdwithin_tcbuffer_geo: unsupported param const int *
// func TdwithinTcbufferGeo(...) { /* not yet handled by codegen */ }


// TODO tdwithin_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func TdwithinTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// TdwithinTcbufferTcbuffer wraps MEOS C function tdwithin_tcbuffer_tcbuffer.
func TdwithinTcbufferTcbuffer(temp1 Temporal, temp2 Temporal, dist float64) Temporal {
	res := C.tdwithin_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner(), C.double(dist))
	return CreateTemporal(res)
}


// TODO tdisjoint_cbuffer_tcbuffer: unsupported param const Cbuffer *
// func TdisjointCbufferTcbuffer(...) { /* not yet handled by codegen */ }


// TODO tdisjoint_geo_tcbuffer: unsupported param const int *
// func TdisjointGeoTcbuffer(...) { /* not yet handled by codegen */ }


// TODO tdisjoint_tcbuffer_geo: unsupported param const int *
// func TdisjointTcbufferGeo(...) { /* not yet handled by codegen */ }


// TODO tdisjoint_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func TdisjointTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// TdisjointTcbufferTcbuffer wraps MEOS C function tdisjoint_tcbuffer_tcbuffer.
func TdisjointTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tdisjoint_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TODO tintersects_cbuffer_tcbuffer: unsupported param const Cbuffer *
// func TintersectsCbufferTcbuffer(...) { /* not yet handled by codegen */ }


// TODO tintersects_geo_tcbuffer: unsupported param const int *
// func TintersectsGeoTcbuffer(...) { /* not yet handled by codegen */ }


// TODO tintersects_tcbuffer_geo: unsupported param const int *
// func TintersectsTcbufferGeo(...) { /* not yet handled by codegen */ }


// TODO tintersects_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func TintersectsTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// TintersectsTcbufferTcbuffer wraps MEOS C function tintersects_tcbuffer_tcbuffer.
func TintersectsTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tintersects_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TODO ttouches_geo_tcbuffer: unsupported param const int *
// func TtouchesGeoTcbuffer(...) { /* not yet handled by codegen */ }


// TODO ttouches_tcbuffer_geo: unsupported param const int *
// func TtouchesTcbufferGeo(...) { /* not yet handled by codegen */ }


// TODO ttouches_cbuffer_tcbuffer: unsupported param const Cbuffer *
// func TtouchesCbufferTcbuffer(...) { /* not yet handled by codegen */ }


// TODO ttouches_tcbuffer_cbuffer: unsupported param const Cbuffer *
// func TtouchesTcbufferCbuffer(...) { /* not yet handled by codegen */ }


// TtouchesTcbufferTcbuffer wraps MEOS C function ttouches_tcbuffer_tcbuffer.
func TtouchesTcbufferTcbuffer(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.ttouches_tcbuffer_tcbuffer(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}

