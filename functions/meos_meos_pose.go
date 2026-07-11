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

// PoseAsEWKT wraps MEOS C function pose_as_ewkt.
func PoseAsEWKT(pose *Pose, maxdd int) string {
	_cret := C.pose_as_ewkt(pose._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// PoseAsHexwkb wraps MEOS C function pose_as_hexwkb.
func PoseAsHexwkb(pose *Pose, variant uint8, size_out unsafe.Pointer) string {
	_cret := C.pose_as_hexwkb(pose._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return C.GoString(_cret)
}


// PoseAsText wraps MEOS C function pose_as_text.
func PoseAsText(pose *Pose, maxdd int) string {
	_cret := C.pose_as_text(pose._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// PoseAsWKB wraps MEOS C function pose_as_wkb.
func PoseAsWKB(pose *Pose, variant uint8, size_out unsafe.Pointer) unsafe.Pointer {
	_cret := C.pose_as_wkb(pose._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return unsafe.Pointer(_cret)
}


// PoseFromWKB wraps MEOS C function pose_from_wkb.
func PoseFromWKB(wkb unsafe.Pointer, size uint) *Pose {
	_cret := C.pose_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	return &Pose{_inner: _cret}
}


// PoseFromHexwkb wraps MEOS C function pose_from_hexwkb.
func PoseFromHexwkb(hexwkb string) *Pose {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	_cret := C.pose_from_hexwkb(_c_hexwkb)
	return &Pose{_inner: _cret}
}


// PoseIn wraps MEOS C function pose_in.
func PoseIn(str string) *Pose {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.pose_in(_c_str)
	return &Pose{_inner: _cret}
}


// PoseOut wraps MEOS C function pose_out.
func PoseOut(pose *Pose, maxdd int) string {
	_cret := C.pose_out(pose._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// PoseFromGeopose wraps MEOS C function pose_from_geopose.
func PoseFromGeopose(json string) *Pose {
	_c_json := C.CString(json)
	defer C.free(unsafe.Pointer(_c_json))
	_cret := C.pose_from_geopose(_c_json)
	return &Pose{_inner: _cret}
}


// PoseAsGeopose wraps MEOS C function pose_as_geopose.
func PoseAsGeopose(pose *Pose, conformance int, precision int) string {
	_cret := C.pose_as_geopose(pose._inner, C.int(conformance), C.int(precision))
	return C.GoString(_cret)
}


// TposeFromGeopose wraps MEOS C function tpose_from_geopose.
func TposeFromGeopose(json string) *Temporal {
	_c_json := C.CString(json)
	defer C.free(unsafe.Pointer(_c_json))
	_cret := C.tpose_from_geopose(_c_json)
	return &Temporal{_inner: _cret}
}


// TposeAsGeopose wraps MEOS C function tpose_as_geopose.
func TposeAsGeopose(temp *Temporal, conformance int, precision int) string {
	_cret := C.tpose_as_geopose(temp._inner, C.int(conformance), C.int(precision))
	return C.GoString(_cret)
}


// PoseApplyGeo wraps MEOS C function pose_apply_geo.
func PoseApplyGeo(pose *Pose, body *Geom) *Geom {
	_cret := C.pose_apply_geo(pose._inner, body._inner)
	return &Geom{_inner: _cret}
}


// TposeApplyGeo wraps MEOS C function tpose_apply_geo.
func TposeApplyGeo(temp *Temporal, body *Geom) *Temporal {
	_cret := C.tpose_apply_geo(temp._inner, body._inner)
	return &Temporal{_inner: _cret}
}


// PoseCopy wraps MEOS C function pose_copy.
func PoseCopy(pose *Pose) *Pose {
	_cret := C.pose_copy(pose._inner)
	return &Pose{_inner: _cret}
}


// PoseMake2d wraps MEOS C function pose_make_2d.
func PoseMake2d(x float64, y float64, theta float64, geodetic bool, srid int32) *Pose {
	_cret := C.pose_make_2d(C.double(x), C.double(y), C.double(theta), C.bool(geodetic), C.int32_t(srid))
	return &Pose{_inner: _cret}
}


// PoseMake3d wraps MEOS C function pose_make_3d.
func PoseMake3d(x float64, y float64, z float64, W float64, X float64, Y float64, Z float64, geodetic bool, srid int32) *Pose {
	_cret := C.pose_make_3d(C.double(x), C.double(y), C.double(z), C.double(W), C.double(X), C.double(Y), C.double(Z), C.bool(geodetic), C.int32_t(srid))
	return &Pose{_inner: _cret}
}


// PoseMakePoint2d wraps MEOS C function pose_make_point2d.
func PoseMakePoint2d(gs *Geom, theta float64) *Pose {
	_cret := C.pose_make_point2d(gs._inner, C.double(theta))
	return &Pose{_inner: _cret}
}


// PoseMakePoint3d wraps MEOS C function pose_make_point3d.
func PoseMakePoint3d(gs *Geom, W float64, X float64, Y float64, Z float64) *Pose {
	_cret := C.pose_make_point3d(gs._inner, C.double(W), C.double(X), C.double(Y), C.double(Z))
	return &Pose{_inner: _cret}
}


// PoseToPoint wraps MEOS C function pose_to_point.
func PoseToPoint(pose *Pose) *Geom {
	_cret := C.pose_to_point(pose._inner)
	return &Geom{_inner: _cret}
}


// PoseToSTBOX wraps MEOS C function pose_to_stbox.
func PoseToSTBOX(pose *Pose) *STBox {
	_cret := C.pose_to_stbox(pose._inner)
	return &STBox{_inner: _cret}
}


// PoseHash wraps MEOS C function pose_hash.
func PoseHash(pose *Pose) uint32 {
	_cret := C.pose_hash(pose._inner)
	return uint32(_cret)
}


// PoseHashExtended wraps MEOS C function pose_hash_extended.
func PoseHashExtended(pose *Pose, seed uint64) uint64 {
	_cret := C.pose_hash_extended(pose._inner, C.uint64_t(seed))
	return uint64(_cret)
}


// PoseOrientation wraps MEOS C function pose_orientation.
func PoseOrientation(pose *Pose, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.pose_orientation(pose._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// PoseRotation wraps MEOS C function pose_rotation.
func PoseRotation(pose *Pose) float64 {
	_cret := C.pose_rotation(pose._inner)
	return float64(_cret)
}


// PoseYaw wraps MEOS C function pose_yaw.
func PoseYaw(pose *Pose) float64 {
	_cret := C.pose_yaw(pose._inner)
	return float64(_cret)
}


// PosePitch wraps MEOS C function pose_pitch.
func PosePitch(pose *Pose) float64 {
	_cret := C.pose_pitch(pose._inner)
	return float64(_cret)
}


// PoseRoll wraps MEOS C function pose_roll.
func PoseRoll(pose *Pose) float64 {
	_cret := C.pose_roll(pose._inner)
	return float64(_cret)
}


// PoseAngularDistance wraps MEOS C function pose_angular_distance.
func PoseAngularDistance(pose1 *Pose, pose2 *Pose) float64 {
	_cret := C.pose_angular_distance(pose1._inner, pose2._inner)
	return float64(_cret)
}


// PoseNormalize wraps MEOS C function pose_normalize.
func PoseNormalize(pose *Pose) *Pose {
	_cret := C.pose_normalize(pose._inner)
	return &Pose{_inner: _cret}
}


// PoseRound wraps MEOS C function pose_round.
func PoseRound(pose *Pose, maxdd int) *Pose {
	_cret := C.pose_round(pose._inner, C.int(maxdd))
	return &Pose{_inner: _cret}
}


// PosearrRound wraps MEOS C function posearr_round.
func PosearrRound(posearr unsafe.Pointer, count int, maxdd int) unsafe.Pointer {
	_cret := C.posearr_round((**C.Pose)(unsafe.Pointer(posearr)), C.int(count), C.int(maxdd))
	return unsafe.Pointer(_cret)
}


// PoseSetSRID wraps MEOS C function pose_set_srid.
func PoseSetSRID(pose *Pose, srid int32) {
	C.pose_set_srid(pose._inner, C.int32_t(srid))
}


// PoseSRID wraps MEOS C function pose_srid.
func PoseSRID(pose *Pose) int32 {
	_cret := C.pose_srid(pose._inner)
	return int32(_cret)
}


// PoseTransform wraps MEOS C function pose_transform.
func PoseTransform(pose *Pose, srid int32) *Pose {
	_cret := C.pose_transform(pose._inner, C.int32_t(srid))
	return &Pose{_inner: _cret}
}


// PoseTransformPipeline wraps MEOS C function pose_transform_pipeline.
func PoseTransformPipeline(pose *Pose, pipelinestr string, srid int32, is_forward bool) *Pose {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	_cret := C.pose_transform_pipeline(pose._inner, _c_pipelinestr, C.int32_t(srid), C.bool(is_forward))
	return &Pose{_inner: _cret}
}


// PoseTstzspanToSTBOX wraps MEOS C function pose_tstzspan_to_stbox.
func PoseTstzspanToSTBOX(pose *Pose, s *Span) *STBox {
	_cret := C.pose_tstzspan_to_stbox(pose._inner, s._inner)
	return &STBox{_inner: _cret}
}


// PoseTimestamptzToSTBOX wraps MEOS C function pose_timestamptz_to_stbox.
func PoseTimestamptzToSTBOX(pose *Pose, t int64) *STBox {
	_cret := C.pose_timestamptz_to_stbox(pose._inner, C.TimestampTz(t))
	return &STBox{_inner: _cret}
}


// DistancePoseGeo wraps MEOS C function distance_pose_geo.
func DistancePoseGeo(pose *Pose, gs *Geom) float64 {
	_cret := C.distance_pose_geo(pose._inner, gs._inner)
	return float64(_cret)
}


// DistancePosePose wraps MEOS C function distance_pose_pose.
func DistancePosePose(pose1 *Pose, pose2 *Pose) float64 {
	_cret := C.distance_pose_pose(pose1._inner, pose2._inner)
	return float64(_cret)
}


// DistancePoseSTBOX wraps MEOS C function distance_pose_stbox.
func DistancePoseSTBOX(pose *Pose, box *STBox) float64 {
	_cret := C.distance_pose_stbox(pose._inner, box._inner)
	return float64(_cret)
}


// PoseCmp wraps MEOS C function pose_cmp.
func PoseCmp(pose1 *Pose, pose2 *Pose) int {
	_cret := C.pose_cmp(pose1._inner, pose2._inner)
	return int(_cret)
}


// PoseEq wraps MEOS C function pose_eq.
func PoseEq(pose1 *Pose, pose2 *Pose) bool {
	_cret := C.pose_eq(pose1._inner, pose2._inner)
	return bool(_cret)
}


// PoseGe wraps MEOS C function pose_ge.
func PoseGe(pose1 *Pose, pose2 *Pose) bool {
	_cret := C.pose_ge(pose1._inner, pose2._inner)
	return bool(_cret)
}


// PoseGt wraps MEOS C function pose_gt.
func PoseGt(pose1 *Pose, pose2 *Pose) bool {
	_cret := C.pose_gt(pose1._inner, pose2._inner)
	return bool(_cret)
}


// PoseLe wraps MEOS C function pose_le.
func PoseLe(pose1 *Pose, pose2 *Pose) bool {
	_cret := C.pose_le(pose1._inner, pose2._inner)
	return bool(_cret)
}


// PoseLt wraps MEOS C function pose_lt.
func PoseLt(pose1 *Pose, pose2 *Pose) bool {
	_cret := C.pose_lt(pose1._inner, pose2._inner)
	return bool(_cret)
}


// PoseNe wraps MEOS C function pose_ne.
func PoseNe(pose1 *Pose, pose2 *Pose) bool {
	_cret := C.pose_ne(pose1._inner, pose2._inner)
	return bool(_cret)
}


// PoseNsame wraps MEOS C function pose_nsame.
func PoseNsame(pose1 *Pose, pose2 *Pose) bool {
	_cret := C.pose_nsame(pose1._inner, pose2._inner)
	return bool(_cret)
}


// PoseSame wraps MEOS C function pose_same.
func PoseSame(pose1 *Pose, pose2 *Pose) bool {
	_cret := C.pose_same(pose1._inner, pose2._inner)
	return bool(_cret)
}


// PosesetIn wraps MEOS C function poseset_in.
func PosesetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.poseset_in(_c_str)
	return &Set{_inner: _cret}
}


// PosesetOut wraps MEOS C function poseset_out.
func PosesetOut(s *Set, maxdd int) string {
	_cret := C.poseset_out(s._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// PosesetMake wraps MEOS C function poseset_make.
func PosesetMake(values unsafe.Pointer, count int) *Set {
	_cret := C.poseset_make((**C.Pose)(unsafe.Pointer(values)), C.int(count))
	return &Set{_inner: _cret}
}


// PoseToSet wraps MEOS C function pose_to_set.
func PoseToSet(pose *Pose) *Set {
	_cret := C.pose_to_set(pose._inner)
	return &Set{_inner: _cret}
}


// PosesetEndValue wraps MEOS C function poseset_end_value.
func PosesetEndValue(s *Set) *Pose {
	_cret := C.poseset_end_value(s._inner)
	return &Pose{_inner: _cret}
}


// PosesetStartValue wraps MEOS C function poseset_start_value.
func PosesetStartValue(s *Set) *Pose {
	_cret := C.poseset_start_value(s._inner)
	return &Pose{_inner: _cret}
}


// PosesetValueN wraps MEOS C function poseset_value_n.
func PosesetValueN(s *Set, n int) (bool, *Pose) {
	var _out_result *C.Pose
	_cret := C.poseset_value_n(s._inner, C.int(n), &_out_result)
	return bool(_cret), &Pose{_inner: _out_result}
}


// PosesetValues wraps MEOS C function poseset_values.
func PosesetValues(s *Set, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.poseset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// ContainedPoseSet wraps MEOS C function contained_pose_set.
func ContainedPoseSet(pose *Pose, s *Set) bool {
	_cret := C.contained_pose_set(pose._inner, s._inner)
	return bool(_cret)
}


// ContainsSetPose wraps MEOS C function contains_set_pose.
func ContainsSetPose(s *Set, pose *Pose) bool {
	_cret := C.contains_set_pose(s._inner, pose._inner)
	return bool(_cret)
}


// IntersectionPoseSet wraps MEOS C function intersection_pose_set.
func IntersectionPoseSet(pose *Pose, s *Set) *Set {
	_cret := C.intersection_pose_set(pose._inner, s._inner)
	return &Set{_inner: _cret}
}


// IntersectionSetPose wraps MEOS C function intersection_set_pose.
func IntersectionSetPose(s *Set, pose *Pose) *Set {
	_cret := C.intersection_set_pose(s._inner, pose._inner)
	return &Set{_inner: _cret}
}


// MinusPoseSet wraps MEOS C function minus_pose_set.
func MinusPoseSet(pose *Pose, s *Set) *Set {
	_cret := C.minus_pose_set(pose._inner, s._inner)
	return &Set{_inner: _cret}
}


// MinusSetPose wraps MEOS C function minus_set_pose.
func MinusSetPose(s *Set, pose *Pose) *Set {
	_cret := C.minus_set_pose(s._inner, pose._inner)
	return &Set{_inner: _cret}
}


// PoseUnionTransfn wraps MEOS C function pose_union_transfn.
func PoseUnionTransfn(state *Set, pose *Pose) *Set {
	_cret := C.pose_union_transfn(state._inner, pose._inner)
	return &Set{_inner: _cret}
}


// UnionPoseSet wraps MEOS C function union_pose_set.
func UnionPoseSet(pose *Pose, s *Set) *Set {
	_cret := C.gunion_pose_set(pose._inner, s._inner)
	return &Set{_inner: _cret}
}


// UnionSetPose wraps MEOS C function union_set_pose.
func UnionSetPose(s *Set, pose *Pose) *Set {
	_cret := C.gunion_set_pose(s._inner, pose._inner)
	return &Set{_inner: _cret}
}


// TposeFromMFJSON wraps MEOS C function tpose_from_mfjson.
func TposeFromMFJSON(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tpose_from_mfjson(_c_str)
	return &Temporal{_inner: _cret}
}


// TposeIn wraps MEOS C function tpose_in.
func TposeIn(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tpose_in(_c_str)
	return &Temporal{_inner: _cret}
}


// TposeinstMake wraps MEOS C function tposeinst_make.
func TposeinstMake(pose *Pose, t int64) *TInstant {
	_cret := C.tposeinst_make(pose._inner, C.TimestampTz(t))
	return &TInstant{_inner: _cret}
}


// TposeFromBaseTemp wraps MEOS C function tpose_from_base_temp.
func TposeFromBaseTemp(pose *Pose, temp *Temporal) *Temporal {
	_cret := C.tpose_from_base_temp(pose._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TposeseqFromBaseTstzset wraps MEOS C function tposeseq_from_base_tstzset.
func TposeseqFromBaseTstzset(pose *Pose, s *Set) *TSequence {
	_cret := C.tposeseq_from_base_tstzset(pose._inner, s._inner)
	return &TSequence{_inner: _cret}
}


// TposeseqFromBaseTstzspan wraps MEOS C function tposeseq_from_base_tstzspan.
func TposeseqFromBaseTstzspan(pose *Pose, s *Span, interp Interpolation) *TSequence {
	_cret := C.tposeseq_from_base_tstzspan(pose._inner, s._inner, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TposeseqsetFromBaseTstzspanset wraps MEOS C function tposeseqset_from_base_tstzspanset.
func TposeseqsetFromBaseTstzspanset(pose *Pose, ss *SpanSet, interp Interpolation) *TSequenceSet {
	_cret := C.tposeseqset_from_base_tstzspanset(pose._inner, ss._inner, C.interpType(interp))
	return &TSequenceSet{_inner: _cret}
}


// TposeMake wraps MEOS C function tpose_make.
func TposeMake(tpoint *Temporal, tradius *Temporal) *Temporal {
	_cret := C.tpose_make(tpoint._inner, tradius._inner)
	return &Temporal{_inner: _cret}
}


// TposeToTpoint wraps MEOS C function tpose_to_tpoint.
func TposeToTpoint(temp *Temporal) *Temporal {
	_cret := C.tpose_to_tpoint(temp._inner)
	return &Temporal{_inner: _cret}
}


// TposeEndValue wraps MEOS C function tpose_end_value.
func TposeEndValue(temp *Temporal) *Pose {
	_cret := C.tpose_end_value(temp._inner)
	return &Pose{_inner: _cret}
}


// TposePoints wraps MEOS C function tpose_points.
func TposePoints(temp *Temporal) *Set {
	_cret := C.tpose_points(temp._inner)
	return &Set{_inner: _cret}
}


// TposeRotation wraps MEOS C function tpose_rotation.
func TposeRotation(temp *Temporal) *Temporal {
	_cret := C.tpose_rotation(temp._inner)
	return &Temporal{_inner: _cret}
}


// TposeYaw wraps MEOS C function tpose_yaw.
func TposeYaw(temp *Temporal) *Temporal {
	_cret := C.tpose_yaw(temp._inner)
	return &Temporal{_inner: _cret}
}


// TposePitch wraps MEOS C function tpose_pitch.
func TposePitch(temp *Temporal) *Temporal {
	_cret := C.tpose_pitch(temp._inner)
	return &Temporal{_inner: _cret}
}


// TposeRoll wraps MEOS C function tpose_roll.
func TposeRoll(temp *Temporal) *Temporal {
	_cret := C.tpose_roll(temp._inner)
	return &Temporal{_inner: _cret}
}


// TposeSpeed wraps MEOS C function tpose_speed.
func TposeSpeed(temp *Temporal) *Temporal {
	_cret := C.tpose_speed(temp._inner)
	return &Temporal{_inner: _cret}
}


// TposeAngularSpeed wraps MEOS C function tpose_angular_speed.
func TposeAngularSpeed(temp *Temporal) *Temporal {
	_cret := C.tpose_angular_speed(temp._inner)
	return &Temporal{_inner: _cret}
}


// TposeStartValue wraps MEOS C function tpose_start_value.
func TposeStartValue(temp *Temporal) *Pose {
	_cret := C.tpose_start_value(temp._inner)
	return &Pose{_inner: _cret}
}


// TposeTrajectory wraps MEOS C function tpose_trajectory.
func TposeTrajectory(temp *Temporal) *Geom {
	_cret := C.tpose_trajectory(temp._inner)
	return &Geom{_inner: _cret}
}


// TposeValueAtTimestamptz wraps MEOS C function tpose_value_at_timestamptz.
func TposeValueAtTimestamptz(temp *Temporal, t int64, strict bool) (bool, *Pose) {
	var _out_result *C.Pose
	_cret := C.tpose_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_result)
	return bool(_cret), &Pose{_inner: _out_result}
}


// TposeValueN wraps MEOS C function tpose_value_n.
func TposeValueN(temp *Temporal, n int) (bool, *Pose) {
	var _out_result *C.Pose
	_cret := C.tpose_value_n(temp._inner, C.int(n), &_out_result)
	return bool(_cret), &Pose{_inner: _out_result}
}


// TposeValues wraps MEOS C function tpose_values.
func TposeValues(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tpose_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TposeAtGeom wraps MEOS C function tpose_at_geom.
func TposeAtGeom(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tpose_at_geom(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TposeAtSTBOX wraps MEOS C function tpose_at_stbox.
func TposeAtSTBOX(temp *Temporal, box *STBox, border_inc bool) *Temporal {
	_cret := C.tpose_at_stbox(temp._inner, box._inner, C.bool(border_inc))
	return &Temporal{_inner: _cret}
}


// TposeAtPose wraps MEOS C function tpose_at_pose.
func TposeAtPose(temp *Temporal, pose *Pose) *Temporal {
	_cret := C.tpose_at_pose(temp._inner, pose._inner)
	return &Temporal{_inner: _cret}
}


// TposeMinusGeom wraps MEOS C function tpose_minus_geom.
func TposeMinusGeom(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tpose_minus_geom(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TposeMinusPose wraps MEOS C function tpose_minus_pose.
func TposeMinusPose(temp *Temporal, pose *Pose) *Temporal {
	_cret := C.tpose_minus_pose(temp._inner, pose._inner)
	return &Temporal{_inner: _cret}
}


// TposeMinusSTBOX wraps MEOS C function tpose_minus_stbox.
func TposeMinusSTBOX(temp *Temporal, box *STBox, border_inc bool) *Temporal {
	_cret := C.tpose_minus_stbox(temp._inner, box._inner, C.bool(border_inc))
	return &Temporal{_inner: _cret}
}


// TdistanceTposePose wraps MEOS C function tdistance_tpose_pose.
func TdistanceTposePose(temp *Temporal, pose *Pose) *Temporal {
	_cret := C.tdistance_tpose_pose(temp._inner, pose._inner)
	return &Temporal{_inner: _cret}
}


// TdistanceTposeGeo wraps MEOS C function tdistance_tpose_geo.
func TdistanceTposeGeo(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tdistance_tpose_geo(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TdistanceTposeTpose wraps MEOS C function tdistance_tpose_tpose.
func TdistanceTposeTpose(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tdistance_tpose_tpose(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// NadTposeGeo wraps MEOS C function nad_tpose_geo.
func NadTposeGeo(temp *Temporal, gs *Geom) float64 {
	_cret := C.nad_tpose_geo(temp._inner, gs._inner)
	return float64(_cret)
}


// NadTposePose wraps MEOS C function nad_tpose_pose.
func NadTposePose(temp *Temporal, pose *Pose) float64 {
	_cret := C.nad_tpose_pose(temp._inner, pose._inner)
	return float64(_cret)
}


// NadTposeSTBOX wraps MEOS C function nad_tpose_stbox.
func NadTposeSTBOX(temp *Temporal, box *STBox) float64 {
	_cret := C.nad_tpose_stbox(temp._inner, box._inner)
	return float64(_cret)
}


// NadTposeTpose wraps MEOS C function nad_tpose_tpose.
func NadTposeTpose(temp1 *Temporal, temp2 *Temporal) float64 {
	_cret := C.nad_tpose_tpose(temp1._inner, temp2._inner)
	return float64(_cret)
}


// NaiTposeGeo wraps MEOS C function nai_tpose_geo.
func NaiTposeGeo(temp *Temporal, gs *Geom) *TInstant {
	_cret := C.nai_tpose_geo(temp._inner, gs._inner)
	return &TInstant{_inner: _cret}
}


// NaiTposePose wraps MEOS C function nai_tpose_pose.
func NaiTposePose(temp *Temporal, pose *Pose) *TInstant {
	_cret := C.nai_tpose_pose(temp._inner, pose._inner)
	return &TInstant{_inner: _cret}
}


// NaiTposeTpose wraps MEOS C function nai_tpose_tpose.
func NaiTposeTpose(temp1 *Temporal, temp2 *Temporal) *TInstant {
	_cret := C.nai_tpose_tpose(temp1._inner, temp2._inner)
	return &TInstant{_inner: _cret}
}


// ShortestlineTposeGeo wraps MEOS C function shortestline_tpose_geo.
func ShortestlineTposeGeo(temp *Temporal, gs *Geom) *Geom {
	_cret := C.shortestline_tpose_geo(temp._inner, gs._inner)
	return &Geom{_inner: _cret}
}


// ShortestlineTposePose wraps MEOS C function shortestline_tpose_pose.
func ShortestlineTposePose(temp *Temporal, pose *Pose) *Geom {
	_cret := C.shortestline_tpose_pose(temp._inner, pose._inner)
	return &Geom{_inner: _cret}
}


// ShortestlineTposeTpose wraps MEOS C function shortestline_tpose_tpose.
func ShortestlineTposeTpose(temp1 *Temporal, temp2 *Temporal) *Geom {
	_cret := C.shortestline_tpose_tpose(temp1._inner, temp2._inner)
	return &Geom{_inner: _cret}
}


// AlwaysEqPoseTpose wraps MEOS C function always_eq_pose_tpose.
func AlwaysEqPoseTpose(pose *Pose, temp *Temporal) int {
	_cret := C.always_eq_pose_tpose(pose._inner, temp._inner)
	return int(_cret)
}


// AlwaysEqTposePose wraps MEOS C function always_eq_tpose_pose.
func AlwaysEqTposePose(temp *Temporal, pose *Pose) int {
	_cret := C.always_eq_tpose_pose(temp._inner, pose._inner)
	return int(_cret)
}


// AlwaysEqTposeTpose wraps MEOS C function always_eq_tpose_tpose.
func AlwaysEqTposeTpose(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_eq_tpose_tpose(temp1._inner, temp2._inner)
	return int(_cret)
}


// AlwaysNePoseTpose wraps MEOS C function always_ne_pose_tpose.
func AlwaysNePoseTpose(pose *Pose, temp *Temporal) int {
	_cret := C.always_ne_pose_tpose(pose._inner, temp._inner)
	return int(_cret)
}


// AlwaysNeTposePose wraps MEOS C function always_ne_tpose_pose.
func AlwaysNeTposePose(temp *Temporal, pose *Pose) int {
	_cret := C.always_ne_tpose_pose(temp._inner, pose._inner)
	return int(_cret)
}


// AlwaysNeTposeTpose wraps MEOS C function always_ne_tpose_tpose.
func AlwaysNeTposeTpose(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_ne_tpose_tpose(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverEqPoseTpose wraps MEOS C function ever_eq_pose_tpose.
func EverEqPoseTpose(pose *Pose, temp *Temporal) int {
	_cret := C.ever_eq_pose_tpose(pose._inner, temp._inner)
	return int(_cret)
}


// EverEqTposePose wraps MEOS C function ever_eq_tpose_pose.
func EverEqTposePose(temp *Temporal, pose *Pose) int {
	_cret := C.ever_eq_tpose_pose(temp._inner, pose._inner)
	return int(_cret)
}


// EverEqTposeTpose wraps MEOS C function ever_eq_tpose_tpose.
func EverEqTposeTpose(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_eq_tpose_tpose(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverNePoseTpose wraps MEOS C function ever_ne_pose_tpose.
func EverNePoseTpose(pose *Pose, temp *Temporal) int {
	_cret := C.ever_ne_pose_tpose(pose._inner, temp._inner)
	return int(_cret)
}


// EverNeTposePose wraps MEOS C function ever_ne_tpose_pose.
func EverNeTposePose(temp *Temporal, pose *Pose) int {
	_cret := C.ever_ne_tpose_pose(temp._inner, pose._inner)
	return int(_cret)
}


// EverNeTposeTpose wraps MEOS C function ever_ne_tpose_tpose.
func EverNeTposeTpose(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_ne_tpose_tpose(temp1._inner, temp2._inner)
	return int(_cret)
}


// TeqPoseTpose wraps MEOS C function teq_pose_tpose.
func TeqPoseTpose(pose *Pose, temp *Temporal) *Temporal {
	_cret := C.teq_pose_tpose(pose._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TeqTposePose wraps MEOS C function teq_tpose_pose.
func TeqTposePose(temp *Temporal, pose *Pose) *Temporal {
	_cret := C.teq_tpose_pose(temp._inner, pose._inner)
	return &Temporal{_inner: _cret}
}


// TnePoseTpose wraps MEOS C function tne_pose_tpose.
func TnePoseTpose(pose *Pose, temp *Temporal) *Temporal {
	_cret := C.tne_pose_tpose(pose._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TneTposePose wraps MEOS C function tne_tpose_pose.
func TneTposePose(temp *Temporal, pose *Pose) *Temporal {
	_cret := C.tne_tpose_pose(temp._inner, pose._inner)
	return &Temporal{_inner: _cret}
}

