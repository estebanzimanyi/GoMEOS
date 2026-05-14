package generated

// #include <stddef.h>
import "C"
import (
	"unsafe"

	"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

// TrgeoOut wraps MEOS C function trgeo_out.
func TrgeoOut(temp Temporal) string {
	res := C.trgeo_out(temp.Inner())
	return C.GoString(res)
}


// TODO trgeoinst_make: unsupported param const Pose *
// func TrgeoinstMake(...) { /* not yet handled by codegen */ }


// GeoTposeToTrgeo wraps MEOS C function geo_tpose_to_trgeo.
func GeoTposeToTrgeo(gs *Geom, temp Temporal) Temporal {
	res := C.geo_tpose_to_trgeo(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TrgeoToTpose wraps MEOS C function trgeo_to_tpose.
func TrgeoToTpose(temp Temporal) Temporal {
	res := C.trgeo_to_tpose(temp.Inner())
	return CreateTemporal(res)
}


// TrgeoToTpoint wraps MEOS C function trgeo_to_tpoint.
func TrgeoToTpoint(temp Temporal) Temporal {
	res := C.trgeo_to_tpoint(temp.Inner())
	return CreateTemporal(res)
}


// TrgeoEndInstant wraps MEOS C function trgeo_end_instant.
func TrgeoEndInstant(temp Temporal) TInstant {
	res := C.trgeo_end_instant(temp.Inner())
	return TInstant{_inner: res}
}


// TrgeoEndSequence wraps MEOS C function trgeo_end_sequence.
func TrgeoEndSequence(temp Temporal) TSequence {
	res := C.trgeo_end_sequence(temp.Inner())
	return TSequence{_inner: res}
}


// TrgeoEndValue wraps MEOS C function trgeo_end_value.
func TrgeoEndValue(temp Temporal) *Geom {
	res := C.trgeo_end_value(temp.Inner())
	return &Geom{_inner: res}
}


// TrgeoGeom wraps MEOS C function trgeo_geom.
func TrgeoGeom(temp Temporal) *Geom {
	res := C.trgeo_geom(temp.Inner())
	return &Geom{_inner: res}
}


// TrgeoInstantN wraps MEOS C function trgeo_instant_n.
func TrgeoInstantN(temp Temporal, n int) TInstant {
	res := C.trgeo_instant_n(temp.Inner(), C.int(n))
	return TInstant{_inner: res}
}


// TrgeoInstants wraps MEOS C function trgeo_instants.
func TrgeoInstants(temp Temporal) []TInstant {
	var _out_count C.int
	res := C.trgeo_instants(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.TInstant)(unsafe.Pointer(res)), _n)
	_out := make([]TInstant, _n)
	for _i, _e := range _slice {
		_out[_i] = TInstant{_inner: _e}
	}
	return _out
}


// TrgeoPoints wraps MEOS C function trgeo_points.
func TrgeoPoints(temp Temporal) *Set {
	res := C.trgeo_points(temp.Inner())
	return &Set{_inner: res}
}


// TrgeoRotation wraps MEOS C function trgeo_rotation.
func TrgeoRotation(temp Temporal) Temporal {
	res := C.trgeo_rotation(temp.Inner())
	return CreateTemporal(res)
}


// TrgeoSegments wraps MEOS C function trgeo_segments.
func TrgeoSegments(temp Temporal) []TSequence {
	var _out_count C.int
	res := C.trgeo_segments(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.TSequence)(unsafe.Pointer(res)), _n)
	_out := make([]TSequence, _n)
	for _i, _e := range _slice {
		_out[_i] = TSequence{_inner: _e}
	}
	return _out
}


// TrgeoSequenceN wraps MEOS C function trgeo_sequence_n.
func TrgeoSequenceN(temp Temporal, i int) TSequence {
	res := C.trgeo_sequence_n(temp.Inner(), C.int(i))
	return TSequence{_inner: res}
}


// TrgeoSequences wraps MEOS C function trgeo_sequences.
func TrgeoSequences(temp Temporal) []TSequence {
	var _out_count C.int
	res := C.trgeo_sequences(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.TSequence)(unsafe.Pointer(res)), _n)
	_out := make([]TSequence, _n)
	for _i, _e := range _slice {
		_out[_i] = TSequence{_inner: _e}
	}
	return _out
}


// TrgeoStartInstant wraps MEOS C function trgeo_start_instant.
func TrgeoStartInstant(temp Temporal) TInstant {
	res := C.trgeo_start_instant(temp.Inner())
	return TInstant{_inner: res}
}


// TrgeoStartSequence wraps MEOS C function trgeo_start_sequence.
func TrgeoStartSequence(temp Temporal) TSequence {
	res := C.trgeo_start_sequence(temp.Inner())
	return TSequence{_inner: res}
}


// TrgeoStartValue wraps MEOS C function trgeo_start_value.
func TrgeoStartValue(temp Temporal) *Geom {
	res := C.trgeo_start_value(temp.Inner())
	return &Geom{_inner: res}
}


// TrgeoValueN wraps MEOS C function trgeo_value_n.
func TrgeoValueN(temp Temporal, n int) (bool, *Geom) {
	var _out_result *C.GSERIALIZED
	res := C.trgeo_value_n(temp.Inner(), C.int(n), &_out_result)
	return bool(res), &Geom{_inner: _out_result}
}


// TrgeoTraversedArea wraps MEOS C function trgeo_traversed_area.
func TrgeoTraversedArea(temp Temporal, unary_union bool) *Geom {
	res := C.trgeo_traversed_area(temp.Inner(), C.bool(unary_union))
	return &Geom{_inner: res}
}


// TrgeoAppendTinstant wraps MEOS C function trgeo_append_tinstant.
func TrgeoAppendTinstant(temp Temporal, inst TInstant, interp Interpolation, maxdist float64, maxt timeutil.Timedelta, expand bool) Temporal {
	res := C.trgeo_append_tinstant(temp.Inner(), inst.Inner(), C.interpType(interp), C.double(maxdist), maxt.Inner(), C.bool(expand))
	return CreateTemporal(res)
}


// TrgeoAppendTsequence wraps MEOS C function trgeo_append_tsequence.
func TrgeoAppendTsequence(temp Temporal, seq TSequence, expand bool) Temporal {
	res := C.trgeo_append_tsequence(temp.Inner(), seq.Inner(), C.bool(expand))
	return CreateTemporal(res)
}


// TrgeoDeleteTimestamptz wraps MEOS C function trgeo_delete_timestamptz.
func TrgeoDeleteTimestamptz(temp Temporal, t int64, connect bool) Temporal {
	res := C.trgeo_delete_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(connect))
	return CreateTemporal(res)
}


// TrgeoDeleteTstzset wraps MEOS C function trgeo_delete_tstzset.
func TrgeoDeleteTstzset(temp Temporal, s *Set, connect bool) Temporal {
	res := C.trgeo_delete_tstzset(temp.Inner(), s._inner, C.bool(connect))
	return CreateTemporal(res)
}


// TrgeoDeleteTstzspan wraps MEOS C function trgeo_delete_tstzspan.
func TrgeoDeleteTstzspan(temp Temporal, s *Span, connect bool) Temporal {
	res := C.trgeo_delete_tstzspan(temp.Inner(), s._inner, C.bool(connect))
	return CreateTemporal(res)
}


// TrgeoDeleteTstzspanset wraps MEOS C function trgeo_delete_tstzspanset.
func TrgeoDeleteTstzspanset(temp Temporal, ss *SpanSet, connect bool) Temporal {
	res := C.trgeo_delete_tstzspanset(temp.Inner(), ss._inner, C.bool(connect))
	return CreateTemporal(res)
}


// TrgeoRound wraps MEOS C function trgeo_round.
func TrgeoRound(temp Temporal, maxdd int) Temporal {
	res := C.trgeo_round(temp.Inner(), C.int(maxdd))
	return CreateTemporal(res)
}


// TrgeoSetInterp wraps MEOS C function trgeo_set_interp.
func TrgeoSetInterp(temp Temporal, interp Interpolation) Temporal {
	res := C.trgeo_set_interp(temp.Inner(), C.interpType(interp))
	return CreateTemporal(res)
}


// TrgeoToTinstant wraps MEOS C function trgeo_to_tinstant.
func TrgeoToTinstant(temp Temporal) TInstant {
	res := C.trgeo_to_tinstant(temp.Inner())
	return TInstant{_inner: res}
}


// TrgeoAfterTimestamptz wraps MEOS C function trgeo_after_timestamptz.
func TrgeoAfterTimestamptz(temp Temporal, t int64, strict bool) Temporal {
	res := C.trgeo_after_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(strict))
	return CreateTemporal(res)
}


// TrgeoBeforeTimestamptz wraps MEOS C function trgeo_before_timestamptz.
func TrgeoBeforeTimestamptz(temp Temporal, t int64, strict bool) Temporal {
	res := C.trgeo_before_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(strict))
	return CreateTemporal(res)
}


// TrgeoRestrictValues wraps MEOS C function trgeo_restrict_values.
func TrgeoRestrictValues(temp Temporal, s *Set, atfunc bool) Temporal {
	res := C.trgeo_restrict_values(temp.Inner(), s._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TrgeoRestrictTimestamptz wraps MEOS C function trgeo_restrict_timestamptz.
func TrgeoRestrictTimestamptz(temp Temporal, t int64, atfunc bool) Temporal {
	res := C.trgeo_restrict_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(atfunc))
	return CreateTemporal(res)
}


// TrgeoRestrictTstzset wraps MEOS C function trgeo_restrict_tstzset.
func TrgeoRestrictTstzset(temp Temporal, s *Set, atfunc bool) Temporal {
	res := C.trgeo_restrict_tstzset(temp.Inner(), s._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TrgeoRestrictTstzspan wraps MEOS C function trgeo_restrict_tstzspan.
func TrgeoRestrictTstzspan(temp Temporal, s *Span, atfunc bool) Temporal {
	res := C.trgeo_restrict_tstzspan(temp.Inner(), s._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TrgeoRestrictTstzspanset wraps MEOS C function trgeo_restrict_tstzspanset.
func TrgeoRestrictTstzspanset(temp Temporal, ss *SpanSet, atfunc bool) Temporal {
	res := C.trgeo_restrict_tstzspanset(temp.Inner(), ss._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TdistanceTrgeoGeo wraps MEOS C function tdistance_trgeo_geo.
func TdistanceTrgeoGeo(temp Temporal, gs *Geom) Temporal {
	res := C.tdistance_trgeo_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TdistanceTrgeoTpoint wraps MEOS C function tdistance_trgeo_tpoint.
func TdistanceTrgeoTpoint(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tdistance_trgeo_tpoint(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TdistanceTrgeoTrgeo wraps MEOS C function tdistance_trgeo_trgeo.
func TdistanceTrgeoTrgeo(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tdistance_trgeo_trgeo(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// NadSTBOXTrgeo wraps MEOS C function nad_stbox_trgeo.
func NadSTBOXTrgeo(box *STBox, temp Temporal) float64 {
	res := C.nad_stbox_trgeo(box._inner, temp.Inner())
	return float64(res)
}


// NadTrgeoGeo wraps MEOS C function nad_trgeo_geo.
func NadTrgeoGeo(temp Temporal, gs *Geom) float64 {
	res := C.nad_trgeo_geo(temp.Inner(), gs._inner)
	return float64(res)
}


// NadTrgeoSTBOX wraps MEOS C function nad_trgeo_stbox.
func NadTrgeoSTBOX(temp Temporal, box *STBox) float64 {
	res := C.nad_trgeo_stbox(temp.Inner(), box._inner)
	return float64(res)
}


// NadTrgeoTpoint wraps MEOS C function nad_trgeo_tpoint.
func NadTrgeoTpoint(temp1 Temporal, temp2 Temporal) float64 {
	res := C.nad_trgeo_tpoint(temp1.Inner(), temp2.Inner())
	return float64(res)
}


// NadTrgeoTrgeo wraps MEOS C function nad_trgeo_trgeo.
func NadTrgeoTrgeo(temp1 Temporal, temp2 Temporal) float64 {
	res := C.nad_trgeo_trgeo(temp1.Inner(), temp2.Inner())
	return float64(res)
}


// NaiTrgeoGeo wraps MEOS C function nai_trgeo_geo.
func NaiTrgeoGeo(temp Temporal, gs *Geom) TInstant {
	res := C.nai_trgeo_geo(temp.Inner(), gs._inner)
	return TInstant{_inner: res}
}


// NaiTrgeoTpoint wraps MEOS C function nai_trgeo_tpoint.
func NaiTrgeoTpoint(temp1 Temporal, temp2 Temporal) TInstant {
	res := C.nai_trgeo_tpoint(temp1.Inner(), temp2.Inner())
	return TInstant{_inner: res}
}


// NaiTrgeoTrgeo wraps MEOS C function nai_trgeo_trgeo.
func NaiTrgeoTrgeo(temp1 Temporal, temp2 Temporal) TInstant {
	res := C.nai_trgeo_trgeo(temp1.Inner(), temp2.Inner())
	return TInstant{_inner: res}
}


// ShortestlineTrgeoGeo wraps MEOS C function shortestline_trgeo_geo.
func ShortestlineTrgeoGeo(temp Temporal, gs *Geom) *Geom {
	res := C.shortestline_trgeo_geo(temp.Inner(), gs._inner)
	return &Geom{_inner: res}
}


// ShortestlineTrgeoTpoint wraps MEOS C function shortestline_trgeo_tpoint.
func ShortestlineTrgeoTpoint(temp1 Temporal, temp2 Temporal) *Geom {
	res := C.shortestline_trgeo_tpoint(temp1.Inner(), temp2.Inner())
	return &Geom{_inner: res}
}


// ShortestlineTrgeoTrgeo wraps MEOS C function shortestline_trgeo_trgeo.
func ShortestlineTrgeoTrgeo(temp1 Temporal, temp2 Temporal) *Geom {
	res := C.shortestline_trgeo_trgeo(temp1.Inner(), temp2.Inner())
	return &Geom{_inner: res}
}


// AlwaysEqGeoTrgeo wraps MEOS C function always_eq_geo_trgeo.
func AlwaysEqGeoTrgeo(gs *Geom, temp Temporal) int {
	res := C.always_eq_geo_trgeo(gs._inner, temp.Inner())
	return int(res)
}


// AlwaysEqTrgeoGeo wraps MEOS C function always_eq_trgeo_geo.
func AlwaysEqTrgeoGeo(temp Temporal, gs *Geom) int {
	res := C.always_eq_trgeo_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AlwaysEqTrgeoTrgeo wraps MEOS C function always_eq_trgeo_trgeo.
func AlwaysEqTrgeoTrgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.always_eq_trgeo_trgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AlwaysNeGeoTrgeo wraps MEOS C function always_ne_geo_trgeo.
func AlwaysNeGeoTrgeo(gs *Geom, temp Temporal) int {
	res := C.always_ne_geo_trgeo(gs._inner, temp.Inner())
	return int(res)
}


// AlwaysNeTrgeoGeo wraps MEOS C function always_ne_trgeo_geo.
func AlwaysNeTrgeoGeo(temp Temporal, gs *Geom) int {
	res := C.always_ne_trgeo_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AlwaysNeTrgeoTrgeo wraps MEOS C function always_ne_trgeo_trgeo.
func AlwaysNeTrgeoTrgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.always_ne_trgeo_trgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverEqGeoTrgeo wraps MEOS C function ever_eq_geo_trgeo.
func EverEqGeoTrgeo(gs *Geom, temp Temporal) int {
	res := C.ever_eq_geo_trgeo(gs._inner, temp.Inner())
	return int(res)
}


// EverEqTrgeoGeo wraps MEOS C function ever_eq_trgeo_geo.
func EverEqTrgeoGeo(temp Temporal, gs *Geom) int {
	res := C.ever_eq_trgeo_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EverEqTrgeoTrgeo wraps MEOS C function ever_eq_trgeo_trgeo.
func EverEqTrgeoTrgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_eq_trgeo_trgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverNeGeoTrgeo wraps MEOS C function ever_ne_geo_trgeo.
func EverNeGeoTrgeo(gs *Geom, temp Temporal) int {
	res := C.ever_ne_geo_trgeo(gs._inner, temp.Inner())
	return int(res)
}


// EverNeTrgeoGeo wraps MEOS C function ever_ne_trgeo_geo.
func EverNeTrgeoGeo(temp Temporal, gs *Geom) int {
	res := C.ever_ne_trgeo_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EverNeTrgeoTrgeo wraps MEOS C function ever_ne_trgeo_trgeo.
func EverNeTrgeoTrgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_ne_trgeo_trgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TeqGeoTrgeo wraps MEOS C function teq_geo_trgeo.
func TeqGeoTrgeo(gs *Geom, temp Temporal) Temporal {
	res := C.teq_geo_trgeo(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TeqTrgeoGeo wraps MEOS C function teq_trgeo_geo.
func TeqTrgeoGeo(temp Temporal, gs *Geom) Temporal {
	res := C.teq_trgeo_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TneGeoTrgeo wraps MEOS C function tne_geo_trgeo.
func TneGeoTrgeo(gs *Geom, temp Temporal) Temporal {
	res := C.tne_geo_trgeo(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TneTrgeoGeo wraps MEOS C function tne_trgeo_geo.
func TneTrgeoGeo(temp Temporal, gs *Geom) Temporal {
	res := C.tne_trgeo_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}

