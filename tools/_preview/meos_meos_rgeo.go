package generated

// #include <stddef.h>
import "C"
import (
	"unsafe"

	"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

// TrgeometryOut wraps MEOS C function trgeometry_out.
func TrgeometryOut(temp Temporal) string {
	res := C.trgeometry_out(temp.Inner())
	return C.GoString(res)
}


// TrgeometryinstMake wraps MEOS C function trgeometryinst_make.
func TrgeometryinstMake(geom *Geom, pose *Pose, t int64) TInstant {
	res := C.trgeometryinst_make(geom._inner, pose._inner, C.TimestampTz(t))
	return TInstant{_inner: res}
}


// GeoTposeToTrgeometry wraps MEOS C function geo_tpose_to_trgeometry.
func GeoTposeToTrgeometry(gs *Geom, temp Temporal) Temporal {
	res := C.geo_tpose_to_trgeometry(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TrgeometryToTpose wraps MEOS C function trgeometry_to_tpose.
func TrgeometryToTpose(temp Temporal) Temporal {
	res := C.trgeometry_to_tpose(temp.Inner())
	return CreateTemporal(res)
}


// TrgeometryToTpoint wraps MEOS C function trgeometry_to_tpoint.
func TrgeometryToTpoint(temp Temporal) Temporal {
	res := C.trgeometry_to_tpoint(temp.Inner())
	return CreateTemporal(res)
}


// TrgeometryToTgeometry wraps MEOS C function trgeometry_to_tgeometry.
func TrgeometryToTgeometry(temp Temporal) Temporal {
	res := C.trgeometry_to_tgeometry(temp.Inner())
	return CreateTemporal(res)
}


// TrgeometryEndInstant wraps MEOS C function trgeometry_end_instant.
func TrgeometryEndInstant(temp Temporal) TInstant {
	res := C.trgeometry_end_instant(temp.Inner())
	return TInstant{_inner: res}
}


// TrgeometryEndSequence wraps MEOS C function trgeometry_end_sequence.
func TrgeometryEndSequence(temp Temporal) TSequence {
	res := C.trgeometry_end_sequence(temp.Inner())
	return TSequence{_inner: res}
}


// TrgeometryEndValue wraps MEOS C function trgeometry_end_value.
func TrgeometryEndValue(temp Temporal) *Geom {
	res := C.trgeometry_end_value(temp.Inner())
	return &Geom{_inner: res}
}


// TrgeometryGeom wraps MEOS C function trgeometry_geom.
func TrgeometryGeom(temp Temporal) *Geom {
	res := C.trgeometry_geom(temp.Inner())
	return &Geom{_inner: res}
}


// TrgeometryInstantN wraps MEOS C function trgeometry_instant_n.
func TrgeometryInstantN(temp Temporal, n int) TInstant {
	res := C.trgeometry_instant_n(temp.Inner(), C.int(n))
	return TInstant{_inner: res}
}


// TrgeometryInstants wraps MEOS C function trgeometry_instants.
func TrgeometryInstants(temp Temporal) []TInstant {
	var _out_count C.int
	res := C.trgeometry_instants(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.TInstant)(unsafe.Pointer(res)), _n)
	_out := make([]TInstant, _n)
	for _i, _e := range _slice {
		_out[_i] = TInstant{_inner: _e}
	}
	return _out
}


// TrgeometryPoints wraps MEOS C function trgeometry_points.
func TrgeometryPoints(temp Temporal) *Set {
	res := C.trgeometry_points(temp.Inner())
	return &Set{_inner: res}
}


// TrgeometryRotation wraps MEOS C function trgeometry_rotation.
func TrgeometryRotation(temp Temporal) Temporal {
	res := C.trgeometry_rotation(temp.Inner())
	return CreateTemporal(res)
}


// TrgeometrySegments wraps MEOS C function trgeometry_segments.
func TrgeometrySegments(temp Temporal) []TSequence {
	var _out_count C.int
	res := C.trgeometry_segments(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.TSequence)(unsafe.Pointer(res)), _n)
	_out := make([]TSequence, _n)
	for _i, _e := range _slice {
		_out[_i] = TSequence{_inner: _e}
	}
	return _out
}


// TrgeometrySequenceN wraps MEOS C function trgeometry_sequence_n.
func TrgeometrySequenceN(temp Temporal, i int) TSequence {
	res := C.trgeometry_sequence_n(temp.Inner(), C.int(i))
	return TSequence{_inner: res}
}


// TrgeometrySequences wraps MEOS C function trgeometry_sequences.
func TrgeometrySequences(temp Temporal) []TSequence {
	var _out_count C.int
	res := C.trgeometry_sequences(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.TSequence)(unsafe.Pointer(res)), _n)
	_out := make([]TSequence, _n)
	for _i, _e := range _slice {
		_out[_i] = TSequence{_inner: _e}
	}
	return _out
}


// TrgeometryStartInstant wraps MEOS C function trgeometry_start_instant.
func TrgeometryStartInstant(temp Temporal) TInstant {
	res := C.trgeometry_start_instant(temp.Inner())
	return TInstant{_inner: res}
}


// TrgeometryStartSequence wraps MEOS C function trgeometry_start_sequence.
func TrgeometryStartSequence(temp Temporal) TSequence {
	res := C.trgeometry_start_sequence(temp.Inner())
	return TSequence{_inner: res}
}


// TrgeometryStartValue wraps MEOS C function trgeometry_start_value.
func TrgeometryStartValue(temp Temporal) *Geom {
	res := C.trgeometry_start_value(temp.Inner())
	return &Geom{_inner: res}
}


// TrgeometryValueN wraps MEOS C function trgeometry_value_n.
func TrgeometryValueN(temp Temporal, n int) (bool, *Geom) {
	var _out_result *C.GSERIALIZED
	res := C.trgeometry_value_n(temp.Inner(), C.int(n), &_out_result)
	return bool(res), &Geom{_inner: _out_result}
}


// TrgeometryTraversedArea wraps MEOS C function trgeometry_traversed_area.
func TrgeometryTraversedArea(temp Temporal, unary_union bool) *Geom {
	res := C.trgeometry_traversed_area(temp.Inner(), C.bool(unary_union))
	return &Geom{_inner: res}
}


// TrgeometryCentroid wraps MEOS C function trgeometry_centroid.
func TrgeometryCentroid(temp Temporal) Temporal {
	res := C.trgeometry_centroid(temp.Inner())
	return CreateTemporal(res)
}


// TrgeometryConvexHull wraps MEOS C function trgeometry_convex_hull.
func TrgeometryConvexHull(temp Temporal) *Geom {
	res := C.trgeometry_convex_hull(temp.Inner())
	return &Geom{_inner: res}
}


// TrgeometryBodyPointTrajectory wraps MEOS C function trgeometry_body_point_trajectory.
func TrgeometryBodyPointTrajectory(temp Temporal, gs *Geom) Temporal {
	res := C.trgeometry_body_point_trajectory(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TrgeometrySpaceBoxes wraps MEOS C function trgeometry_space_boxes.
func TrgeometrySpaceBoxes(temp Temporal, xsize float64, ysize float64, zsize float64, sorigin *Geom, bitmatrix bool, border_inc bool) (*STBox, int) {
	var _out_count C.int
	res := C.trgeometry_space_boxes(temp.Inner(), C.double(xsize), C.double(ysize), C.double(zsize), sorigin._inner, C.bool(bitmatrix), C.bool(border_inc), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// TrgeometrySpaceTimeBoxes wraps MEOS C function trgeometry_space_time_boxes.
func TrgeometrySpaceTimeBoxes(temp Temporal, xsize float64, ysize float64, zsize float64, duration timeutil.Timedelta, sorigin *Geom, torigin int64, bitmatrix bool, border_inc bool) (*STBox, int) {
	var _out_count C.int
	res := C.trgeometry_space_time_boxes(temp.Inner(), C.double(xsize), C.double(ysize), C.double(zsize), duration.Inner(), sorigin._inner, C.TimestampTz(torigin), C.bool(bitmatrix), C.bool(border_inc), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// TrgeometryStboxes wraps MEOS C function trgeometry_stboxes.
func TrgeometryStboxes(temp Temporal) (*STBox, int) {
	var _out_count C.int
	res := C.trgeometry_stboxes(temp.Inner(), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// TrgeometrySplitNStboxes wraps MEOS C function trgeometry_split_n_stboxes.
func TrgeometrySplitNStboxes(temp Temporal, box_count int) (*STBox, int) {
	var _out_count C.int
	res := C.trgeometry_split_n_stboxes(temp.Inner(), C.int(box_count), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// TrgeometrySplitEachNStboxes wraps MEOS C function trgeometry_split_each_n_stboxes.
func TrgeometrySplitEachNStboxes(temp Temporal, elem_count int) (*STBox, int) {
	var _out_count C.int
	res := C.trgeometry_split_each_n_stboxes(temp.Inner(), C.int(elem_count), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// TrgeometryHausdorffDistance wraps MEOS C function trgeometry_hausdorff_distance.
func TrgeometryHausdorffDistance(temp1 Temporal, temp2 Temporal) float64 {
	res := C.trgeometry_hausdorff_distance(temp1.Inner(), temp2.Inner())
	return float64(res)
}


// TrgeometryFrechetDistance wraps MEOS C function trgeometry_frechet_distance.
func TrgeometryFrechetDistance(temp1 Temporal, temp2 Temporal) float64 {
	res := C.trgeometry_frechet_distance(temp1.Inner(), temp2.Inner())
	return float64(res)
}


// TrgeometryDyntimewarpDistance wraps MEOS C function trgeometry_dyntimewarp_distance.
func TrgeometryDyntimewarpDistance(temp1 Temporal, temp2 Temporal) float64 {
	res := C.trgeometry_dyntimewarp_distance(temp1.Inner(), temp2.Inner())
	return float64(res)
}


// TrgeometryFrechetPath wraps MEOS C function trgeometry_frechet_path.
func TrgeometryFrechetPath(temp1 Temporal, temp2 Temporal) (*Match, int) {
	var _out_count C.int
	res := C.trgeometry_frechet_path(temp1.Inner(), temp2.Inner(), &_out_count)
	return &Match{_inner: res}, int(_out_count)
}


// TrgeometryDyntimewarpPath wraps MEOS C function trgeometry_dyntimewarp_path.
func TrgeometryDyntimewarpPath(temp1 Temporal, temp2 Temporal) (*Match, int) {
	var _out_count C.int
	res := C.trgeometry_dyntimewarp_path(temp1.Inner(), temp2.Inner(), &_out_count)
	return &Match{_inner: res}, int(_out_count)
}


// TrgeometryLength wraps MEOS C function trgeometry_length.
func TrgeometryLength(temp Temporal) float64 {
	res := C.trgeometry_length(temp.Inner())
	return float64(res)
}


// TrgeometryCumulativeLength wraps MEOS C function trgeometry_cumulative_length.
func TrgeometryCumulativeLength(temp Temporal) Temporal {
	res := C.trgeometry_cumulative_length(temp.Inner())
	return CreateTemporal(res)
}


// TrgeometrySpeed wraps MEOS C function trgeometry_speed.
func TrgeometrySpeed(temp Temporal) Temporal {
	res := C.trgeometry_speed(temp.Inner())
	return CreateTemporal(res)
}


// TrgeometryTwcentroid wraps MEOS C function trgeometry_twcentroid.
func TrgeometryTwcentroid(temp Temporal) *Geom {
	res := C.trgeometry_twcentroid(temp.Inner())
	return &Geom{_inner: res}
}


// TrgeometryAppendTinstant wraps MEOS C function trgeometry_append_tinstant.
func TrgeometryAppendTinstant(temp Temporal, inst TInstant, interp Interpolation, maxdist float64, maxt timeutil.Timedelta, expand bool) Temporal {
	res := C.trgeometry_append_tinstant(temp.Inner(), inst.Inner(), C.interpType(interp), C.double(maxdist), maxt.Inner(), C.bool(expand))
	return CreateTemporal(res)
}


// TrgeometryAppendTsequence wraps MEOS C function trgeometry_append_tsequence.
func TrgeometryAppendTsequence(temp Temporal, seq TSequence, expand bool) Temporal {
	res := C.trgeometry_append_tsequence(temp.Inner(), seq.Inner(), C.bool(expand))
	return CreateTemporal(res)
}


// TrgeometryDeleteTimestamptz wraps MEOS C function trgeometry_delete_timestamptz.
func TrgeometryDeleteTimestamptz(temp Temporal, t int64, connect bool) Temporal {
	res := C.trgeometry_delete_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(connect))
	return CreateTemporal(res)
}


// TrgeometryDeleteTstzset wraps MEOS C function trgeometry_delete_tstzset.
func TrgeometryDeleteTstzset(temp Temporal, s *Set, connect bool) Temporal {
	res := C.trgeometry_delete_tstzset(temp.Inner(), s._inner, C.bool(connect))
	return CreateTemporal(res)
}


// TrgeometryDeleteTstzspan wraps MEOS C function trgeometry_delete_tstzspan.
func TrgeometryDeleteTstzspan(temp Temporal, s *Span, connect bool) Temporal {
	res := C.trgeometry_delete_tstzspan(temp.Inner(), s._inner, C.bool(connect))
	return CreateTemporal(res)
}


// TrgeometryDeleteTstzspanset wraps MEOS C function trgeometry_delete_tstzspanset.
func TrgeometryDeleteTstzspanset(temp Temporal, ss *SpanSet, connect bool) Temporal {
	res := C.trgeometry_delete_tstzspanset(temp.Inner(), ss._inner, C.bool(connect))
	return CreateTemporal(res)
}


// TrgeometryRound wraps MEOS C function trgeometry_round.
func TrgeometryRound(temp Temporal, maxdd int) Temporal {
	res := C.trgeometry_round(temp.Inner(), C.int(maxdd))
	return CreateTemporal(res)
}


// TrgeometrySetInterp wraps MEOS C function trgeometry_set_interp.
func TrgeometrySetInterp(temp Temporal, interp Interpolation) Temporal {
	res := C.trgeometry_set_interp(temp.Inner(), C.interpType(interp))
	return CreateTemporal(res)
}


// TrgeometryToTinstant wraps MEOS C function trgeometry_to_tinstant.
func TrgeometryToTinstant(temp Temporal) TInstant {
	res := C.trgeometry_to_tinstant(temp.Inner())
	return TInstant{_inner: res}
}


// TrgeometryAfterTimestamptz wraps MEOS C function trgeometry_after_timestamptz.
func TrgeometryAfterTimestamptz(temp Temporal, t int64, strict bool) Temporal {
	res := C.trgeometry_after_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(strict))
	return CreateTemporal(res)
}


// TrgeometryBeforeTimestamptz wraps MEOS C function trgeometry_before_timestamptz.
func TrgeometryBeforeTimestamptz(temp Temporal, t int64, strict bool) Temporal {
	res := C.trgeometry_before_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(strict))
	return CreateTemporal(res)
}


// TrgeometryRestrictValues wraps MEOS C function trgeometry_restrict_values.
func TrgeometryRestrictValues(temp Temporal, s *Set, atfunc bool) Temporal {
	res := C.trgeometry_restrict_values(temp.Inner(), s._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TrgeometryRestrictTimestamptz wraps MEOS C function trgeometry_restrict_timestamptz.
func TrgeometryRestrictTimestamptz(temp Temporal, t int64, atfunc bool) Temporal {
	res := C.trgeometry_restrict_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(atfunc))
	return CreateTemporal(res)
}


// TrgeometryRestrictTstzset wraps MEOS C function trgeometry_restrict_tstzset.
func TrgeometryRestrictTstzset(temp Temporal, s *Set, atfunc bool) Temporal {
	res := C.trgeometry_restrict_tstzset(temp.Inner(), s._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TrgeometryRestrictTstzspan wraps MEOS C function trgeometry_restrict_tstzspan.
func TrgeometryRestrictTstzspan(temp Temporal, s *Span, atfunc bool) Temporal {
	res := C.trgeometry_restrict_tstzspan(temp.Inner(), s._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TrgeometryRestrictTstzspanset wraps MEOS C function trgeometry_restrict_tstzspanset.
func TrgeometryRestrictTstzspanset(temp Temporal, ss *SpanSet, atfunc bool) Temporal {
	res := C.trgeometry_restrict_tstzspanset(temp.Inner(), ss._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TrgeometryAtGeom wraps MEOS C function trgeometry_at_geom.
func TrgeometryAtGeom(temp Temporal, gs *Geom) Temporal {
	res := C.trgeometry_at_geom(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TrgeometryMinusGeom wraps MEOS C function trgeometry_minus_geom.
func TrgeometryMinusGeom(temp Temporal, gs *Geom) Temporal {
	res := C.trgeometry_minus_geom(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TrgeometryAtSTBOX wraps MEOS C function trgeometry_at_stbox.
func TrgeometryAtSTBOX(temp Temporal, box *STBox, border_inc bool) Temporal {
	res := C.trgeometry_at_stbox(temp.Inner(), box._inner, C.bool(border_inc))
	return CreateTemporal(res)
}


// TrgeometryMinusSTBOX wraps MEOS C function trgeometry_minus_stbox.
func TrgeometryMinusSTBOX(temp Temporal, box *STBox, border_inc bool) Temporal {
	res := C.trgeometry_minus_stbox(temp.Inner(), box._inner, C.bool(border_inc))
	return CreateTemporal(res)
}


// TdistanceTrgeometryGeo wraps MEOS C function tdistance_trgeometry_geo.
func TdistanceTrgeometryGeo(temp Temporal, gs *Geom) Temporal {
	res := C.tdistance_trgeometry_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TdistanceTrgeometryTpoint wraps MEOS C function tdistance_trgeometry_tpoint.
func TdistanceTrgeometryTpoint(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tdistance_trgeometry_tpoint(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TdistanceTrgeometryTrgeometry wraps MEOS C function tdistance_trgeometry_trgeometry.
func TdistanceTrgeometryTrgeometry(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tdistance_trgeometry_trgeometry(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// NadSTBOXTrgeometry wraps MEOS C function nad_stbox_trgeometry.
func NadSTBOXTrgeometry(box *STBox, temp Temporal) float64 {
	res := C.nad_stbox_trgeometry(box._inner, temp.Inner())
	return float64(res)
}


// NadTrgeometryGeo wraps MEOS C function nad_trgeometry_geo.
func NadTrgeometryGeo(temp Temporal, gs *Geom) float64 {
	res := C.nad_trgeometry_geo(temp.Inner(), gs._inner)
	return float64(res)
}


// NadTrgeometrySTBOX wraps MEOS C function nad_trgeometry_stbox.
func NadTrgeometrySTBOX(temp Temporal, box *STBox) float64 {
	res := C.nad_trgeometry_stbox(temp.Inner(), box._inner)
	return float64(res)
}


// NadTrgeometryTpoint wraps MEOS C function nad_trgeometry_tpoint.
func NadTrgeometryTpoint(temp1 Temporal, temp2 Temporal) float64 {
	res := C.nad_trgeometry_tpoint(temp1.Inner(), temp2.Inner())
	return float64(res)
}


// NadTrgeometryTrgeometry wraps MEOS C function nad_trgeometry_trgeometry.
func NadTrgeometryTrgeometry(temp1 Temporal, temp2 Temporal) float64 {
	res := C.nad_trgeometry_trgeometry(temp1.Inner(), temp2.Inner())
	return float64(res)
}


// NaiTrgeometryGeo wraps MEOS C function nai_trgeometry_geo.
func NaiTrgeometryGeo(temp Temporal, gs *Geom) TInstant {
	res := C.nai_trgeometry_geo(temp.Inner(), gs._inner)
	return TInstant{_inner: res}
}


// NaiTrgeometryTpoint wraps MEOS C function nai_trgeometry_tpoint.
func NaiTrgeometryTpoint(temp1 Temporal, temp2 Temporal) TInstant {
	res := C.nai_trgeometry_tpoint(temp1.Inner(), temp2.Inner())
	return TInstant{_inner: res}
}


// NaiTrgeometryTrgeometry wraps MEOS C function nai_trgeometry_trgeometry.
func NaiTrgeometryTrgeometry(temp1 Temporal, temp2 Temporal) TInstant {
	res := C.nai_trgeometry_trgeometry(temp1.Inner(), temp2.Inner())
	return TInstant{_inner: res}
}


// ShortestlineTrgeometryGeo wraps MEOS C function shortestline_trgeometry_geo.
func ShortestlineTrgeometryGeo(temp Temporal, gs *Geom) *Geom {
	res := C.shortestline_trgeometry_geo(temp.Inner(), gs._inner)
	return &Geom{_inner: res}
}


// ShortestlineTrgeometryTpoint wraps MEOS C function shortestline_trgeometry_tpoint.
func ShortestlineTrgeometryTpoint(temp1 Temporal, temp2 Temporal) *Geom {
	res := C.shortestline_trgeometry_tpoint(temp1.Inner(), temp2.Inner())
	return &Geom{_inner: res}
}


// ShortestlineTrgeometryTrgeometry wraps MEOS C function shortestline_trgeometry_trgeometry.
func ShortestlineTrgeometryTrgeometry(temp1 Temporal, temp2 Temporal) *Geom {
	res := C.shortestline_trgeometry_trgeometry(temp1.Inner(), temp2.Inner())
	return &Geom{_inner: res}
}


// AlwaysEqGeoTrgeometry wraps MEOS C function always_eq_geo_trgeometry.
func AlwaysEqGeoTrgeometry(gs *Geom, temp Temporal) int {
	res := C.always_eq_geo_trgeometry(gs._inner, temp.Inner())
	return int(res)
}


// AlwaysEqTrgeometryGeo wraps MEOS C function always_eq_trgeometry_geo.
func AlwaysEqTrgeometryGeo(temp Temporal, gs *Geom) int {
	res := C.always_eq_trgeometry_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AlwaysEqTrgeometryTrgeometry wraps MEOS C function always_eq_trgeometry_trgeometry.
func AlwaysEqTrgeometryTrgeometry(temp1 Temporal, temp2 Temporal) int {
	res := C.always_eq_trgeometry_trgeometry(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AlwaysNeGeoTrgeometry wraps MEOS C function always_ne_geo_trgeometry.
func AlwaysNeGeoTrgeometry(gs *Geom, temp Temporal) int {
	res := C.always_ne_geo_trgeometry(gs._inner, temp.Inner())
	return int(res)
}


// AlwaysNeTrgeometryGeo wraps MEOS C function always_ne_trgeometry_geo.
func AlwaysNeTrgeometryGeo(temp Temporal, gs *Geom) int {
	res := C.always_ne_trgeometry_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AlwaysNeTrgeometryTrgeometry wraps MEOS C function always_ne_trgeometry_trgeometry.
func AlwaysNeTrgeometryTrgeometry(temp1 Temporal, temp2 Temporal) int {
	res := C.always_ne_trgeometry_trgeometry(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverEqGeoTrgeometry wraps MEOS C function ever_eq_geo_trgeometry.
func EverEqGeoTrgeometry(gs *Geom, temp Temporal) int {
	res := C.ever_eq_geo_trgeometry(gs._inner, temp.Inner())
	return int(res)
}


// EverEqTrgeometryGeo wraps MEOS C function ever_eq_trgeometry_geo.
func EverEqTrgeometryGeo(temp Temporal, gs *Geom) int {
	res := C.ever_eq_trgeometry_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EverEqTrgeometryTrgeometry wraps MEOS C function ever_eq_trgeometry_trgeometry.
func EverEqTrgeometryTrgeometry(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_eq_trgeometry_trgeometry(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverNeGeoTrgeometry wraps MEOS C function ever_ne_geo_trgeometry.
func EverNeGeoTrgeometry(gs *Geom, temp Temporal) int {
	res := C.ever_ne_geo_trgeometry(gs._inner, temp.Inner())
	return int(res)
}


// EverNeTrgeometryGeo wraps MEOS C function ever_ne_trgeometry_geo.
func EverNeTrgeometryGeo(temp Temporal, gs *Geom) int {
	res := C.ever_ne_trgeometry_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EverNeTrgeometryTrgeometry wraps MEOS C function ever_ne_trgeometry_trgeometry.
func EverNeTrgeometryTrgeometry(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_ne_trgeometry_trgeometry(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TeqGeoTrgeometry wraps MEOS C function teq_geo_trgeometry.
func TeqGeoTrgeometry(gs *Geom, temp Temporal) Temporal {
	res := C.teq_geo_trgeometry(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TeqTrgeometryGeo wraps MEOS C function teq_trgeometry_geo.
func TeqTrgeometryGeo(temp Temporal, gs *Geom) Temporal {
	res := C.teq_trgeometry_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TneGeoTrgeometry wraps MEOS C function tne_geo_trgeometry.
func TneGeoTrgeometry(gs *Geom, temp Temporal) Temporal {
	res := C.tne_geo_trgeometry(gs._inner, temp.Inner())
	return CreateTemporal(res)
}


// TneTrgeometryGeo wraps MEOS C function tne_trgeometry_geo.
func TneTrgeometryGeo(temp Temporal, gs *Geom) Temporal {
	res := C.tne_trgeometry_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// EcontainsGeoTrgeometry wraps MEOS C function econtains_geo_trgeometry.
func EcontainsGeoTrgeometry(gs *Geom, temp Temporal) int {
	res := C.econtains_geo_trgeometry(gs._inner, temp.Inner())
	return int(res)
}


// AcontainsGeoTrgeometry wraps MEOS C function acontains_geo_trgeometry.
func AcontainsGeoTrgeometry(gs *Geom, temp Temporal) int {
	res := C.acontains_geo_trgeometry(gs._inner, temp.Inner())
	return int(res)
}


// EcoversGeoTrgeometry wraps MEOS C function ecovers_geo_trgeometry.
func EcoversGeoTrgeometry(gs *Geom, temp Temporal) int {
	res := C.ecovers_geo_trgeometry(gs._inner, temp.Inner())
	return int(res)
}


// AcoversGeoTrgeometry wraps MEOS C function acovers_geo_trgeometry.
func AcoversGeoTrgeometry(gs *Geom, temp Temporal) int {
	res := C.acovers_geo_trgeometry(gs._inner, temp.Inner())
	return int(res)
}


// EcoversTrgeometryGeo wraps MEOS C function ecovers_trgeometry_geo.
func EcoversTrgeometryGeo(temp Temporal, gs *Geom) int {
	res := C.ecovers_trgeometry_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AcoversTrgeometryGeo wraps MEOS C function acovers_trgeometry_geo.
func AcoversTrgeometryGeo(temp Temporal, gs *Geom) int {
	res := C.acovers_trgeometry_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EdisjointTrgeometryGeo wraps MEOS C function edisjoint_trgeometry_geo.
func EdisjointTrgeometryGeo(temp Temporal, gs *Geom) int {
	res := C.edisjoint_trgeometry_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AdisjointTrgeometryGeo wraps MEOS C function adisjoint_trgeometry_geo.
func AdisjointTrgeometryGeo(temp Temporal, gs *Geom) int {
	res := C.adisjoint_trgeometry_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EintersectsTrgeometryGeo wraps MEOS C function eintersects_trgeometry_geo.
func EintersectsTrgeometryGeo(temp Temporal, gs *Geom) int {
	res := C.eintersects_trgeometry_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AintersectsTrgeometryGeo wraps MEOS C function aintersects_trgeometry_geo.
func AintersectsTrgeometryGeo(temp Temporal, gs *Geom) int {
	res := C.aintersects_trgeometry_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EtouchesTrgeometryGeo wraps MEOS C function etouches_trgeometry_geo.
func EtouchesTrgeometryGeo(temp Temporal, gs *Geom) int {
	res := C.etouches_trgeometry_geo(temp.Inner(), gs._inner)
	return int(res)
}


// AtouchesTrgeometryGeo wraps MEOS C function atouches_trgeometry_geo.
func AtouchesTrgeometryGeo(temp Temporal, gs *Geom) int {
	res := C.atouches_trgeometry_geo(temp.Inner(), gs._inner)
	return int(res)
}


// EdwithinTrgeometryGeo wraps MEOS C function edwithin_trgeometry_geo.
func EdwithinTrgeometryGeo(temp Temporal, gs *Geom, dist float64) int {
	res := C.edwithin_trgeometry_geo(temp.Inner(), gs._inner, C.double(dist))
	return int(res)
}


// AdwithinTrgeometryGeo wraps MEOS C function adwithin_trgeometry_geo.
func AdwithinTrgeometryGeo(temp Temporal, gs *Geom, dist float64) int {
	res := C.adwithin_trgeometry_geo(temp.Inner(), gs._inner, C.double(dist))
	return int(res)
}


// EdisjointTrgeometryTrgeometry wraps MEOS C function edisjoint_trgeometry_trgeometry.
func EdisjointTrgeometryTrgeometry(temp1 Temporal, temp2 Temporal) int {
	res := C.edisjoint_trgeometry_trgeometry(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AdisjointTrgeometryTrgeometry wraps MEOS C function adisjoint_trgeometry_trgeometry.
func AdisjointTrgeometryTrgeometry(temp1 Temporal, temp2 Temporal) int {
	res := C.adisjoint_trgeometry_trgeometry(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EintersectsTrgeometryTrgeometry wraps MEOS C function eintersects_trgeometry_trgeometry.
func EintersectsTrgeometryTrgeometry(temp1 Temporal, temp2 Temporal) int {
	res := C.eintersects_trgeometry_trgeometry(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AintersectsTrgeometryTrgeometry wraps MEOS C function aintersects_trgeometry_trgeometry.
func AintersectsTrgeometryTrgeometry(temp1 Temporal, temp2 Temporal) int {
	res := C.aintersects_trgeometry_trgeometry(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EdwithinTrgeometryTrgeometry wraps MEOS C function edwithin_trgeometry_trgeometry.
func EdwithinTrgeometryTrgeometry(temp1 Temporal, temp2 Temporal, dist float64) int {
	res := C.edwithin_trgeometry_trgeometry(temp1.Inner(), temp2.Inner(), C.double(dist))
	return int(res)
}


// AdwithinTrgeometryTrgeometry wraps MEOS C function adwithin_trgeometry_trgeometry.
func AdwithinTrgeometryTrgeometry(temp1 Temporal, temp2 Temporal, dist float64) int {
	res := C.adwithin_trgeometry_trgeometry(temp1.Inner(), temp2.Inner(), C.double(dist))
	return int(res)
}

