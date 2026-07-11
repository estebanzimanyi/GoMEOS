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

// TrgeometryOut wraps MEOS C function trgeometry_out.
func TrgeometryOut(temp *Temporal) string {
	_cret := C.trgeometry_out(temp._inner)
	return C.GoString(_cret)
}


// TrgeometryinstMake wraps MEOS C function trgeometryinst_make.
func TrgeometryinstMake(geom *Geom, pose *Pose, t int64) *TInstant {
	_cret := C.trgeometryinst_make(geom._inner, pose._inner, C.TimestampTz(t))
	return &TInstant{_inner: _cret}
}


// GeoTposeToTrgeometry wraps MEOS C function geo_tpose_to_trgeometry.
func GeoTposeToTrgeometry(gs *Geom, temp *Temporal) *Temporal {
	_cret := C.geo_tpose_to_trgeometry(gs._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TrgeometryToTpose wraps MEOS C function trgeometry_to_tpose.
func TrgeometryToTpose(temp *Temporal) *Temporal {
	_cret := C.trgeometry_to_tpose(temp._inner)
	return &Temporal{_inner: _cret}
}


// TrgeometryToTpoint wraps MEOS C function trgeometry_to_tpoint.
func TrgeometryToTpoint(temp *Temporal) *Temporal {
	_cret := C.trgeometry_to_tpoint(temp._inner)
	return &Temporal{_inner: _cret}
}


// TrgeometryToTgeometry wraps MEOS C function trgeometry_to_tgeometry.
func TrgeometryToTgeometry(temp *Temporal) *Temporal {
	_cret := C.trgeometry_to_tgeometry(temp._inner)
	return &Temporal{_inner: _cret}
}


// TrgeometryEndInstant wraps MEOS C function trgeometry_end_instant.
func TrgeometryEndInstant(temp *Temporal) *TInstant {
	_cret := C.trgeometry_end_instant(temp._inner)
	return &TInstant{_inner: _cret}
}


// TrgeometryEndSequence wraps MEOS C function trgeometry_end_sequence.
func TrgeometryEndSequence(temp *Temporal) *TSequence {
	_cret := C.trgeometry_end_sequence(temp._inner)
	return &TSequence{_inner: _cret}
}


// TrgeometryEndValue wraps MEOS C function trgeometry_end_value.
func TrgeometryEndValue(temp *Temporal) *Geom {
	_cret := C.trgeometry_end_value(temp._inner)
	return &Geom{_inner: _cret}
}


// TrgeometryGeom wraps MEOS C function trgeometry_geom.
func TrgeometryGeom(temp *Temporal) *Geom {
	_cret := C.trgeometry_geom(temp._inner)
	return &Geom{_inner: _cret}
}


// TrgeometryInstantN wraps MEOS C function trgeometry_instant_n.
func TrgeometryInstantN(temp *Temporal, n int) *TInstant {
	_cret := C.trgeometry_instant_n(temp._inner, C.int(n))
	return &TInstant{_inner: _cret}
}


// TrgeometryInstants wraps MEOS C function trgeometry_instants.
func TrgeometryInstants(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.trgeometry_instants(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TrgeometryPoints wraps MEOS C function trgeometry_points.
func TrgeometryPoints(temp *Temporal) *Set {
	_cret := C.trgeometry_points(temp._inner)
	return &Set{_inner: _cret}
}


// TrgeometryRotation wraps MEOS C function trgeometry_rotation.
func TrgeometryRotation(temp *Temporal) *Temporal {
	_cret := C.trgeometry_rotation(temp._inner)
	return &Temporal{_inner: _cret}
}


// TrgeometrySegments wraps MEOS C function trgeometry_segments.
func TrgeometrySegments(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.trgeometry_segments(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TrgeometrySequenceN wraps MEOS C function trgeometry_sequence_n.
func TrgeometrySequenceN(temp *Temporal, i int) *TSequence {
	_cret := C.trgeometry_sequence_n(temp._inner, C.int(i))
	return &TSequence{_inner: _cret}
}


// TrgeometrySequences wraps MEOS C function trgeometry_sequences.
func TrgeometrySequences(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.trgeometry_sequences(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TrgeometryStartInstant wraps MEOS C function trgeometry_start_instant.
func TrgeometryStartInstant(temp *Temporal) *TInstant {
	_cret := C.trgeometry_start_instant(temp._inner)
	return &TInstant{_inner: _cret}
}


// TrgeometryStartSequence wraps MEOS C function trgeometry_start_sequence.
func TrgeometryStartSequence(temp *Temporal) *TSequence {
	_cret := C.trgeometry_start_sequence(temp._inner)
	return &TSequence{_inner: _cret}
}


// TrgeometryStartValue wraps MEOS C function trgeometry_start_value.
func TrgeometryStartValue(temp *Temporal) *Geom {
	_cret := C.trgeometry_start_value(temp._inner)
	return &Geom{_inner: _cret}
}


// TrgeometryValueN wraps MEOS C function trgeometry_value_n.
func TrgeometryValueN(temp *Temporal, n int) (bool, *Geom) {
	var _out_result *C.GSERIALIZED
	_cret := C.trgeometry_value_n(temp._inner, C.int(n), &_out_result)
	return bool(_cret), &Geom{_inner: _out_result}
}


// TrgeometryTraversedArea wraps MEOS C function trgeometry_traversed_area.
func TrgeometryTraversedArea(temp *Temporal, unary_union bool) *Geom {
	_cret := C.trgeometry_traversed_area(temp._inner, C.bool(unary_union))
	return &Geom{_inner: _cret}
}


// TrgeometryCentroid wraps MEOS C function trgeometry_centroid.
func TrgeometryCentroid(temp *Temporal) *Temporal {
	_cret := C.trgeometry_centroid(temp._inner)
	return &Temporal{_inner: _cret}
}


// TrgeometryConvexHull wraps MEOS C function trgeometry_convex_hull.
func TrgeometryConvexHull(temp *Temporal) *Geom {
	_cret := C.trgeometry_convex_hull(temp._inner)
	return &Geom{_inner: _cret}
}


// TrgeometryBodyPointTrajectory wraps MEOS C function trgeometry_body_point_trajectory.
func TrgeometryBodyPointTrajectory(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.trgeometry_body_point_trajectory(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TrgeometrySpaceBoxes wraps MEOS C function trgeometry_space_boxes.
func TrgeometrySpaceBoxes(temp *Temporal, xsize float64, ysize float64, zsize float64, sorigin *Geom, bitmatrix bool, border_inc bool, count unsafe.Pointer) *STBox {
	_cret := C.trgeometry_space_boxes(temp._inner, C.double(xsize), C.double(ysize), C.double(zsize), sorigin._inner, C.bool(bitmatrix), C.bool(border_inc), (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// TrgeometrySpaceTimeBoxes wraps MEOS C function trgeometry_space_time_boxes.
func TrgeometrySpaceTimeBoxes(temp *Temporal, xsize float64, ysize float64, zsize float64, duration *Interval, sorigin *Geom, torigin int64, bitmatrix bool, border_inc bool, count unsafe.Pointer) *STBox {
	_cret := C.trgeometry_space_time_boxes(temp._inner, C.double(xsize), C.double(ysize), C.double(zsize), duration._inner, sorigin._inner, C.TimestampTz(torigin), C.bool(bitmatrix), C.bool(border_inc), (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// TrgeometryStboxes wraps MEOS C function trgeometry_stboxes.
func TrgeometryStboxes(temp *Temporal, count unsafe.Pointer) *STBox {
	_cret := C.trgeometry_stboxes(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// TrgeometrySplitNStboxes wraps MEOS C function trgeometry_split_n_stboxes.
func TrgeometrySplitNStboxes(temp *Temporal, box_count int, count unsafe.Pointer) *STBox {
	_cret := C.trgeometry_split_n_stboxes(temp._inner, C.int(box_count), (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// TrgeometrySplitEachNStboxes wraps MEOS C function trgeometry_split_each_n_stboxes.
func TrgeometrySplitEachNStboxes(temp *Temporal, elem_count int, count unsafe.Pointer) *STBox {
	_cret := C.trgeometry_split_each_n_stboxes(temp._inner, C.int(elem_count), (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// TrgeometryHausdorffDistance wraps MEOS C function trgeometry_hausdorff_distance.
func TrgeometryHausdorffDistance(temp1 *Temporal, temp2 *Temporal) float64 {
	_cret := C.trgeometry_hausdorff_distance(temp1._inner, temp2._inner)
	return float64(_cret)
}


// TrgeometryFrechetDistance wraps MEOS C function trgeometry_frechet_distance.
func TrgeometryFrechetDistance(temp1 *Temporal, temp2 *Temporal) float64 {
	_cret := C.trgeometry_frechet_distance(temp1._inner, temp2._inner)
	return float64(_cret)
}


// TrgeometryDyntimewarpDistance wraps MEOS C function trgeometry_dyntimewarp_distance.
func TrgeometryDyntimewarpDistance(temp1 *Temporal, temp2 *Temporal) float64 {
	_cret := C.trgeometry_dyntimewarp_distance(temp1._inner, temp2._inner)
	return float64(_cret)
}


// TrgeometryFrechetPath wraps MEOS C function trgeometry_frechet_path.
func TrgeometryFrechetPath(temp1 *Temporal, temp2 *Temporal, count unsafe.Pointer) *Match {
	_cret := C.trgeometry_frechet_path(temp1._inner, temp2._inner, (*C.int)(unsafe.Pointer(count)))
	return &Match{_inner: _cret}
}


// TrgeometryDyntimewarpPath wraps MEOS C function trgeometry_dyntimewarp_path.
func TrgeometryDyntimewarpPath(temp1 *Temporal, temp2 *Temporal, count unsafe.Pointer) *Match {
	_cret := C.trgeometry_dyntimewarp_path(temp1._inner, temp2._inner, (*C.int)(unsafe.Pointer(count)))
	return &Match{_inner: _cret}
}


// TrgeometryLength wraps MEOS C function trgeometry_length.
func TrgeometryLength(temp *Temporal) float64 {
	_cret := C.trgeometry_length(temp._inner)
	return float64(_cret)
}


// TrgeometryCumulativeLength wraps MEOS C function trgeometry_cumulative_length.
func TrgeometryCumulativeLength(temp *Temporal) *Temporal {
	_cret := C.trgeometry_cumulative_length(temp._inner)
	return &Temporal{_inner: _cret}
}


// TrgeometrySpeed wraps MEOS C function trgeometry_speed.
func TrgeometrySpeed(temp *Temporal) *Temporal {
	_cret := C.trgeometry_speed(temp._inner)
	return &Temporal{_inner: _cret}
}


// TrgeometryTwcentroid wraps MEOS C function trgeometry_twcentroid.
func TrgeometryTwcentroid(temp *Temporal) *Geom {
	_cret := C.trgeometry_twcentroid(temp._inner)
	return &Geom{_inner: _cret}
}


// TrgeometryAppendTinstant wraps MEOS C function trgeometry_append_tinstant.
func TrgeometryAppendTinstant(temp *Temporal, inst *TInstant, interp Interpolation, maxdist float64, maxt *Interval, expand bool) *Temporal {
	_cret := C.trgeometry_append_tinstant(temp._inner, inst._inner, C.interpType(interp), C.double(maxdist), maxt._inner, C.bool(expand))
	return &Temporal{_inner: _cret}
}


// TrgeometryAppendTsequence wraps MEOS C function trgeometry_append_tsequence.
func TrgeometryAppendTsequence(temp *Temporal, seq *TSequence, expand bool) *Temporal {
	_cret := C.trgeometry_append_tsequence(temp._inner, seq._inner, C.bool(expand))
	return &Temporal{_inner: _cret}
}


// TrgeometryDeleteTimestamptz wraps MEOS C function trgeometry_delete_timestamptz.
func TrgeometryDeleteTimestamptz(temp *Temporal, t int64, connect bool) *Temporal {
	_cret := C.trgeometry_delete_timestamptz(temp._inner, C.TimestampTz(t), C.bool(connect))
	return &Temporal{_inner: _cret}
}


// TrgeometryDeleteTstzset wraps MEOS C function trgeometry_delete_tstzset.
func TrgeometryDeleteTstzset(temp *Temporal, s *Set, connect bool) *Temporal {
	_cret := C.trgeometry_delete_tstzset(temp._inner, s._inner, C.bool(connect))
	return &Temporal{_inner: _cret}
}


// TrgeometryDeleteTstzspan wraps MEOS C function trgeometry_delete_tstzspan.
func TrgeometryDeleteTstzspan(temp *Temporal, s *Span, connect bool) *Temporal {
	_cret := C.trgeometry_delete_tstzspan(temp._inner, s._inner, C.bool(connect))
	return &Temporal{_inner: _cret}
}


// TrgeometryDeleteTstzspanset wraps MEOS C function trgeometry_delete_tstzspanset.
func TrgeometryDeleteTstzspanset(temp *Temporal, ss *SpanSet, connect bool) *Temporal {
	_cret := C.trgeometry_delete_tstzspanset(temp._inner, ss._inner, C.bool(connect))
	return &Temporal{_inner: _cret}
}


// TrgeometryRound wraps MEOS C function trgeometry_round.
func TrgeometryRound(temp *Temporal, maxdd int) *Temporal {
	_cret := C.trgeometry_round(temp._inner, C.int(maxdd))
	return &Temporal{_inner: _cret}
}


// TrgeometrySetInterp wraps MEOS C function trgeometry_set_interp.
func TrgeometrySetInterp(temp *Temporal, interp Interpolation) *Temporal {
	_cret := C.trgeometry_set_interp(temp._inner, C.interpType(interp))
	return &Temporal{_inner: _cret}
}


// TrgeometryToTinstant wraps MEOS C function trgeometry_to_tinstant.
func TrgeometryToTinstant(temp *Temporal) *TInstant {
	_cret := C.trgeometry_to_tinstant(temp._inner)
	return &TInstant{_inner: _cret}
}


// TrgeometryAfterTimestamptz wraps MEOS C function trgeometry_after_timestamptz.
func TrgeometryAfterTimestamptz(temp *Temporal, t int64, strict bool) *Temporal {
	_cret := C.trgeometry_after_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict))
	return &Temporal{_inner: _cret}
}


// TrgeometryBeforeTimestamptz wraps MEOS C function trgeometry_before_timestamptz.
func TrgeometryBeforeTimestamptz(temp *Temporal, t int64, strict bool) *Temporal {
	_cret := C.trgeometry_before_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict))
	return &Temporal{_inner: _cret}
}


// TrgeometryRestrictValues wraps MEOS C function trgeometry_restrict_values.
func TrgeometryRestrictValues(temp *Temporal, s *Set, atfunc bool) *Temporal {
	_cret := C.trgeometry_restrict_values(temp._inner, s._inner, C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TrgeometryRestrictTimestamptz wraps MEOS C function trgeometry_restrict_timestamptz.
func TrgeometryRestrictTimestamptz(temp *Temporal, t int64, atfunc bool) *Temporal {
	_cret := C.trgeometry_restrict_timestamptz(temp._inner, C.TimestampTz(t), C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TrgeometryRestrictTstzset wraps MEOS C function trgeometry_restrict_tstzset.
func TrgeometryRestrictTstzset(temp *Temporal, s *Set, atfunc bool) *Temporal {
	_cret := C.trgeometry_restrict_tstzset(temp._inner, s._inner, C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TrgeometryRestrictTstzspan wraps MEOS C function trgeometry_restrict_tstzspan.
func TrgeometryRestrictTstzspan(temp *Temporal, s *Span, atfunc bool) *Temporal {
	_cret := C.trgeometry_restrict_tstzspan(temp._inner, s._inner, C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TrgeometryRestrictTstzspanset wraps MEOS C function trgeometry_restrict_tstzspanset.
func TrgeometryRestrictTstzspanset(temp *Temporal, ss *SpanSet, atfunc bool) *Temporal {
	_cret := C.trgeometry_restrict_tstzspanset(temp._inner, ss._inner, C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TrgeometryAtGeom wraps MEOS C function trgeometry_at_geom.
func TrgeometryAtGeom(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.trgeometry_at_geom(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TrgeometryMinusGeom wraps MEOS C function trgeometry_minus_geom.
func TrgeometryMinusGeom(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.trgeometry_minus_geom(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TrgeometryAtSTBOX wraps MEOS C function trgeometry_at_stbox.
func TrgeometryAtSTBOX(temp *Temporal, box *STBox, border_inc bool) *Temporal {
	_cret := C.trgeometry_at_stbox(temp._inner, box._inner, C.bool(border_inc))
	return &Temporal{_inner: _cret}
}


// TrgeometryMinusSTBOX wraps MEOS C function trgeometry_minus_stbox.
func TrgeometryMinusSTBOX(temp *Temporal, box *STBox, border_inc bool) *Temporal {
	_cret := C.trgeometry_minus_stbox(temp._inner, box._inner, C.bool(border_inc))
	return &Temporal{_inner: _cret}
}


// TdistanceTrgeometryGeo wraps MEOS C function tdistance_trgeometry_geo.
func TdistanceTrgeometryGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tdistance_trgeometry_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TdistanceTrgeometryTpoint wraps MEOS C function tdistance_trgeometry_tpoint.
func TdistanceTrgeometryTpoint(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tdistance_trgeometry_tpoint(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TdistanceTrgeometryTrgeometry wraps MEOS C function tdistance_trgeometry_trgeometry.
func TdistanceTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tdistance_trgeometry_trgeometry(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// NadSTBOXTrgeometry wraps MEOS C function nad_stbox_trgeometry.
func NadSTBOXTrgeometry(box *STBox, temp *Temporal) float64 {
	_cret := C.nad_stbox_trgeometry(box._inner, temp._inner)
	return float64(_cret)
}


// NadTrgeometryGeo wraps MEOS C function nad_trgeometry_geo.
func NadTrgeometryGeo(temp *Temporal, gs *Geom) float64 {
	_cret := C.nad_trgeometry_geo(temp._inner, gs._inner)
	return float64(_cret)
}


// NadTrgeometrySTBOX wraps MEOS C function nad_trgeometry_stbox.
func NadTrgeometrySTBOX(temp *Temporal, box *STBox) float64 {
	_cret := C.nad_trgeometry_stbox(temp._inner, box._inner)
	return float64(_cret)
}


// NadTrgeometryTpoint wraps MEOS C function nad_trgeometry_tpoint.
func NadTrgeometryTpoint(temp1 *Temporal, temp2 *Temporal) float64 {
	_cret := C.nad_trgeometry_tpoint(temp1._inner, temp2._inner)
	return float64(_cret)
}


// NadTrgeometryTrgeometry wraps MEOS C function nad_trgeometry_trgeometry.
func NadTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) float64 {
	_cret := C.nad_trgeometry_trgeometry(temp1._inner, temp2._inner)
	return float64(_cret)
}


// NaiTrgeometryGeo wraps MEOS C function nai_trgeometry_geo.
func NaiTrgeometryGeo(temp *Temporal, gs *Geom) *TInstant {
	_cret := C.nai_trgeometry_geo(temp._inner, gs._inner)
	return &TInstant{_inner: _cret}
}


// NaiTrgeometryTpoint wraps MEOS C function nai_trgeometry_tpoint.
func NaiTrgeometryTpoint(temp1 *Temporal, temp2 *Temporal) *TInstant {
	_cret := C.nai_trgeometry_tpoint(temp1._inner, temp2._inner)
	return &TInstant{_inner: _cret}
}


// NaiTrgeometryTrgeometry wraps MEOS C function nai_trgeometry_trgeometry.
func NaiTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) *TInstant {
	_cret := C.nai_trgeometry_trgeometry(temp1._inner, temp2._inner)
	return &TInstant{_inner: _cret}
}


// ShortestlineTrgeometryGeo wraps MEOS C function shortestline_trgeometry_geo.
func ShortestlineTrgeometryGeo(temp *Temporal, gs *Geom) *Geom {
	_cret := C.shortestline_trgeometry_geo(temp._inner, gs._inner)
	return &Geom{_inner: _cret}
}


// ShortestlineTrgeometryTpoint wraps MEOS C function shortestline_trgeometry_tpoint.
func ShortestlineTrgeometryTpoint(temp1 *Temporal, temp2 *Temporal) *Geom {
	_cret := C.shortestline_trgeometry_tpoint(temp1._inner, temp2._inner)
	return &Geom{_inner: _cret}
}


// ShortestlineTrgeometryTrgeometry wraps MEOS C function shortestline_trgeometry_trgeometry.
func ShortestlineTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) *Geom {
	_cret := C.shortestline_trgeometry_trgeometry(temp1._inner, temp2._inner)
	return &Geom{_inner: _cret}
}


// AlwaysEqGeoTrgeometry wraps MEOS C function always_eq_geo_trgeometry.
func AlwaysEqGeoTrgeometry(gs *Geom, temp *Temporal) int {
	_cret := C.always_eq_geo_trgeometry(gs._inner, temp._inner)
	return int(_cret)
}


// AlwaysEqTrgeometryGeo wraps MEOS C function always_eq_trgeometry_geo.
func AlwaysEqTrgeometryGeo(temp *Temporal, gs *Geom) int {
	_cret := C.always_eq_trgeometry_geo(temp._inner, gs._inner)
	return int(_cret)
}


// AlwaysEqTrgeometryTrgeometry wraps MEOS C function always_eq_trgeometry_trgeometry.
func AlwaysEqTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_eq_trgeometry_trgeometry(temp1._inner, temp2._inner)
	return int(_cret)
}


// AlwaysNeGeoTrgeometry wraps MEOS C function always_ne_geo_trgeometry.
func AlwaysNeGeoTrgeometry(gs *Geom, temp *Temporal) int {
	_cret := C.always_ne_geo_trgeometry(gs._inner, temp._inner)
	return int(_cret)
}


// AlwaysNeTrgeometryGeo wraps MEOS C function always_ne_trgeometry_geo.
func AlwaysNeTrgeometryGeo(temp *Temporal, gs *Geom) int {
	_cret := C.always_ne_trgeometry_geo(temp._inner, gs._inner)
	return int(_cret)
}


// AlwaysNeTrgeometryTrgeometry wraps MEOS C function always_ne_trgeometry_trgeometry.
func AlwaysNeTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_ne_trgeometry_trgeometry(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverEqGeoTrgeometry wraps MEOS C function ever_eq_geo_trgeometry.
func EverEqGeoTrgeometry(gs *Geom, temp *Temporal) int {
	_cret := C.ever_eq_geo_trgeometry(gs._inner, temp._inner)
	return int(_cret)
}


// EverEqTrgeometryGeo wraps MEOS C function ever_eq_trgeometry_geo.
func EverEqTrgeometryGeo(temp *Temporal, gs *Geom) int {
	_cret := C.ever_eq_trgeometry_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EverEqTrgeometryTrgeometry wraps MEOS C function ever_eq_trgeometry_trgeometry.
func EverEqTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_eq_trgeometry_trgeometry(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverNeGeoTrgeometry wraps MEOS C function ever_ne_geo_trgeometry.
func EverNeGeoTrgeometry(gs *Geom, temp *Temporal) int {
	_cret := C.ever_ne_geo_trgeometry(gs._inner, temp._inner)
	return int(_cret)
}


// EverNeTrgeometryGeo wraps MEOS C function ever_ne_trgeometry_geo.
func EverNeTrgeometryGeo(temp *Temporal, gs *Geom) int {
	_cret := C.ever_ne_trgeometry_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EverNeTrgeometryTrgeometry wraps MEOS C function ever_ne_trgeometry_trgeometry.
func EverNeTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_ne_trgeometry_trgeometry(temp1._inner, temp2._inner)
	return int(_cret)
}


// TeqGeoTrgeometry wraps MEOS C function teq_geo_trgeometry.
func TeqGeoTrgeometry(gs *Geom, temp *Temporal) *Temporal {
	_cret := C.teq_geo_trgeometry(gs._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TeqTrgeometryGeo wraps MEOS C function teq_trgeometry_geo.
func TeqTrgeometryGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.teq_trgeometry_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TneGeoTrgeometry wraps MEOS C function tne_geo_trgeometry.
func TneGeoTrgeometry(gs *Geom, temp *Temporal) *Temporal {
	_cret := C.tne_geo_trgeometry(gs._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TneTrgeometryGeo wraps MEOS C function tne_trgeometry_geo.
func TneTrgeometryGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tne_trgeometry_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// EcontainsGeoTrgeometry wraps MEOS C function econtains_geo_trgeometry.
func EcontainsGeoTrgeometry(gs *Geom, temp *Temporal) int {
	_cret := C.econtains_geo_trgeometry(gs._inner, temp._inner)
	return int(_cret)
}


// AcontainsGeoTrgeometry wraps MEOS C function acontains_geo_trgeometry.
func AcontainsGeoTrgeometry(gs *Geom, temp *Temporal) int {
	_cret := C.acontains_geo_trgeometry(gs._inner, temp._inner)
	return int(_cret)
}


// EcoversGeoTrgeometry wraps MEOS C function ecovers_geo_trgeometry.
func EcoversGeoTrgeometry(gs *Geom, temp *Temporal) int {
	_cret := C.ecovers_geo_trgeometry(gs._inner, temp._inner)
	return int(_cret)
}


// AcoversGeoTrgeometry wraps MEOS C function acovers_geo_trgeometry.
func AcoversGeoTrgeometry(gs *Geom, temp *Temporal) int {
	_cret := C.acovers_geo_trgeometry(gs._inner, temp._inner)
	return int(_cret)
}


// EcoversTrgeometryGeo wraps MEOS C function ecovers_trgeometry_geo.
func EcoversTrgeometryGeo(temp *Temporal, gs *Geom) int {
	_cret := C.ecovers_trgeometry_geo(temp._inner, gs._inner)
	return int(_cret)
}


// AcoversTrgeometryGeo wraps MEOS C function acovers_trgeometry_geo.
func AcoversTrgeometryGeo(temp *Temporal, gs *Geom) int {
	_cret := C.acovers_trgeometry_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EdisjointTrgeometryGeo wraps MEOS C function edisjoint_trgeometry_geo.
func EdisjointTrgeometryGeo(temp *Temporal, gs *Geom) int {
	_cret := C.edisjoint_trgeometry_geo(temp._inner, gs._inner)
	return int(_cret)
}


// AdisjointTrgeometryGeo wraps MEOS C function adisjoint_trgeometry_geo.
func AdisjointTrgeometryGeo(temp *Temporal, gs *Geom) int {
	_cret := C.adisjoint_trgeometry_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EintersectsTrgeometryGeo wraps MEOS C function eintersects_trgeometry_geo.
func EintersectsTrgeometryGeo(temp *Temporal, gs *Geom) int {
	_cret := C.eintersects_trgeometry_geo(temp._inner, gs._inner)
	return int(_cret)
}


// AintersectsTrgeometryGeo wraps MEOS C function aintersects_trgeometry_geo.
func AintersectsTrgeometryGeo(temp *Temporal, gs *Geom) int {
	_cret := C.aintersects_trgeometry_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EtouchesTrgeometryGeo wraps MEOS C function etouches_trgeometry_geo.
func EtouchesTrgeometryGeo(temp *Temporal, gs *Geom) int {
	_cret := C.etouches_trgeometry_geo(temp._inner, gs._inner)
	return int(_cret)
}


// AtouchesTrgeometryGeo wraps MEOS C function atouches_trgeometry_geo.
func AtouchesTrgeometryGeo(temp *Temporal, gs *Geom) int {
	_cret := C.atouches_trgeometry_geo(temp._inner, gs._inner)
	return int(_cret)
}


// EdwithinTrgeometryGeo wraps MEOS C function edwithin_trgeometry_geo.
func EdwithinTrgeometryGeo(temp *Temporal, gs *Geom, dist float64) int {
	_cret := C.edwithin_trgeometry_geo(temp._inner, gs._inner, C.double(dist))
	return int(_cret)
}


// AdwithinTrgeometryGeo wraps MEOS C function adwithin_trgeometry_geo.
func AdwithinTrgeometryGeo(temp *Temporal, gs *Geom, dist float64) int {
	_cret := C.adwithin_trgeometry_geo(temp._inner, gs._inner, C.double(dist))
	return int(_cret)
}


// EdisjointTrgeometryTrgeometry wraps MEOS C function edisjoint_trgeometry_trgeometry.
func EdisjointTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.edisjoint_trgeometry_trgeometry(temp1._inner, temp2._inner)
	return int(_cret)
}


// AdisjointTrgeometryTrgeometry wraps MEOS C function adisjoint_trgeometry_trgeometry.
func AdisjointTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.adisjoint_trgeometry_trgeometry(temp1._inner, temp2._inner)
	return int(_cret)
}


// EintersectsTrgeometryTrgeometry wraps MEOS C function eintersects_trgeometry_trgeometry.
func EintersectsTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.eintersects_trgeometry_trgeometry(temp1._inner, temp2._inner)
	return int(_cret)
}


// AintersectsTrgeometryTrgeometry wraps MEOS C function aintersects_trgeometry_trgeometry.
func AintersectsTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.aintersects_trgeometry_trgeometry(temp1._inner, temp2._inner)
	return int(_cret)
}


// EdwithinTrgeometryTrgeometry wraps MEOS C function edwithin_trgeometry_trgeometry.
func EdwithinTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal, dist float64) int {
	_cret := C.edwithin_trgeometry_trgeometry(temp1._inner, temp2._inner, C.double(dist))
	return int(_cret)
}


// AdwithinTrgeometryTrgeometry wraps MEOS C function adwithin_trgeometry_trgeometry.
func AdwithinTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal, dist float64) int {
	_cret := C.adwithin_trgeometry_trgeometry(temp1._inner, temp2._inner, C.double(dist))
	return int(_cret)
}

