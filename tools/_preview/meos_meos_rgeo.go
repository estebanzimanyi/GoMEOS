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


// TODO trgeoinst_make: unsupported param const int *
// func TrgeoinstMake(...) { /* not yet handled by codegen */ }


// TODO geo_tpose_to_trgeo: unsupported param const int *
// func GeoTposeToTrgeo(...) { /* not yet handled by codegen */ }


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


// TODO trgeo_end_value: unsupported return type int *
// func TrgeoEndValue(...) { /* not yet handled by codegen */ }


// TODO trgeo_geom: unsupported return type int *
// func TrgeoGeom(...) { /* not yet handled by codegen */ }


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


// TODO trgeo_start_value: unsupported return type int *
// func TrgeoStartValue(...) { /* not yet handled by codegen */ }


// TODO trgeo_value_n: unhandled OUTPUT_SCALAR shape int **
// func TrgeoValueN(...) { /* not yet handled by codegen */ }


// TODO trgeo_traversed_area: unsupported return type int *
// func TrgeoTraversedArea(...) { /* not yet handled by codegen */ }


// TODO trgeo_append_tinstant: unsupported param const int *
// func TrgeoAppendTinstant(...) { /* not yet handled by codegen */ }


// TrgeoAppendTsequence wraps MEOS C function trgeo_append_tsequence.
func TrgeoAppendTsequence(temp Temporal, seq TSequence, expand bool) Temporal {
	res := C.trgeo_append_tsequence(temp.Inner(), seq.Inner(), C.bool(expand))
	return CreateTemporal(res)
}


// TrgeoDeleteTimestamptz wraps MEOS C function trgeo_delete_timestamptz.
func TrgeoDeleteTimestamptz(temp Temporal, t int, connect bool) Temporal {
	res := C.trgeo_delete_timestamptz(temp.Inner(), C.int(t), C.bool(connect))
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
func TrgeoAfterTimestamptz(temp Temporal, t int, strict bool) Temporal {
	res := C.trgeo_after_timestamptz(temp.Inner(), C.int(t), C.bool(strict))
	return CreateTemporal(res)
}


// TrgeoBeforeTimestamptz wraps MEOS C function trgeo_before_timestamptz.
func TrgeoBeforeTimestamptz(temp Temporal, t int, strict bool) Temporal {
	res := C.trgeo_before_timestamptz(temp.Inner(), C.int(t), C.bool(strict))
	return CreateTemporal(res)
}


// TrgeoRestrictValues wraps MEOS C function trgeo_restrict_values.
func TrgeoRestrictValues(temp Temporal, s *Set, atfunc bool) Temporal {
	res := C.trgeo_restrict_values(temp.Inner(), s._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TrgeoRestrictTimestamptz wraps MEOS C function trgeo_restrict_timestamptz.
func TrgeoRestrictTimestamptz(temp Temporal, t int, atfunc bool) Temporal {
	res := C.trgeo_restrict_timestamptz(temp.Inner(), C.int(t), C.bool(atfunc))
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


// TODO tdistance_trgeo_geo: unsupported param const int *
// func TdistanceTrgeoGeo(...) { /* not yet handled by codegen */ }


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


// TODO nad_trgeo_geo: unsupported param const int *
// func NadTrgeoGeo(...) { /* not yet handled by codegen */ }


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


// TODO nai_trgeo_geo: unsupported param const int *
// func NaiTrgeoGeo(...) { /* not yet handled by codegen */ }


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


// TODO shortestline_trgeo_geo: unsupported return type int *
// func ShortestlineTrgeoGeo(...) { /* not yet handled by codegen */ }


// TODO shortestline_trgeo_tpoint: unsupported return type int *
// func ShortestlineTrgeoTpoint(...) { /* not yet handled by codegen */ }


// TODO shortestline_trgeo_trgeo: unsupported return type int *
// func ShortestlineTrgeoTrgeo(...) { /* not yet handled by codegen */ }


// TODO always_eq_geo_trgeo: unsupported param const int *
// func AlwaysEqGeoTrgeo(...) { /* not yet handled by codegen */ }


// TODO always_eq_trgeo_geo: unsupported param const int *
// func AlwaysEqTrgeoGeo(...) { /* not yet handled by codegen */ }


// AlwaysEqTrgeoTrgeo wraps MEOS C function always_eq_trgeo_trgeo.
func AlwaysEqTrgeoTrgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.always_eq_trgeo_trgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO always_ne_geo_trgeo: unsupported param const int *
// func AlwaysNeGeoTrgeo(...) { /* not yet handled by codegen */ }


// TODO always_ne_trgeo_geo: unsupported param const int *
// func AlwaysNeTrgeoGeo(...) { /* not yet handled by codegen */ }


// AlwaysNeTrgeoTrgeo wraps MEOS C function always_ne_trgeo_trgeo.
func AlwaysNeTrgeoTrgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.always_ne_trgeo_trgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO ever_eq_geo_trgeo: unsupported param const int *
// func EverEqGeoTrgeo(...) { /* not yet handled by codegen */ }


// TODO ever_eq_trgeo_geo: unsupported param const int *
// func EverEqTrgeoGeo(...) { /* not yet handled by codegen */ }


// EverEqTrgeoTrgeo wraps MEOS C function ever_eq_trgeo_trgeo.
func EverEqTrgeoTrgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_eq_trgeo_trgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO ever_ne_geo_trgeo: unsupported param const int *
// func EverNeGeoTrgeo(...) { /* not yet handled by codegen */ }


// TODO ever_ne_trgeo_geo: unsupported param const int *
// func EverNeTrgeoGeo(...) { /* not yet handled by codegen */ }


// EverNeTrgeoTrgeo wraps MEOS C function ever_ne_trgeo_trgeo.
func EverNeTrgeoTrgeo(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_ne_trgeo_trgeo(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO teq_geo_trgeo: unsupported param const int *
// func TeqGeoTrgeo(...) { /* not yet handled by codegen */ }


// TODO teq_trgeo_geo: unsupported param const int *
// func TeqTrgeoGeo(...) { /* not yet handled by codegen */ }


// TODO tne_geo_trgeo: unsupported param const int *
// func TneGeoTrgeo(...) { /* not yet handled by codegen */ }


// TODO tne_trgeo_geo: unsupported param const int *
// func TneTrgeoGeo(...) { /* not yet handled by codegen */ }

