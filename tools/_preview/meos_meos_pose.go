package generated

// #include <stddef.h>
import "C"
import (
	"unsafe"

	"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

// PoseAsEWKT wraps MEOS C function pose_as_ewkt.
func PoseAsEWKT(pose *Pose, maxdd int) string {
	res := C.pose_as_ewkt(pose._inner, C.int(maxdd))
	return C.GoString(res)
}


// PoseAsHexwkb wraps MEOS C function pose_as_hexwkb.
func PoseAsHexwkb(pose *Pose, variant uint8) (string, uint) {
	var _out_size C.size_t
	res := C.pose_as_hexwkb(pose._inner, C.uint8_t(variant), &_out_size)
	return C.GoString(res), uint(_out_size)
}


// PoseAsText wraps MEOS C function pose_as_text.
func PoseAsText(pose *Pose, maxdd int) string {
	res := C.pose_as_text(pose._inner, C.int(maxdd))
	return C.GoString(res)
}


// PoseAsWKB wraps MEOS C function pose_as_wkb.
func PoseAsWKB(pose *Pose, variant uint8) []uint8 {
	var _out_size_out C.size_t
	res := C.pose_as_wkb(pose._inner, C.uint8_t(variant), &_out_size_out)
	_n := int(_out_size_out)
	_slice := unsafe.Slice((*C.uint8_t)(unsafe.Pointer(res)), _n)
	_out := make([]uint8, _n)
	for _i, _e := range _slice {
		_out[_i] = uint8(_e)
	}
	return _out
}


// PoseFromWKB wraps MEOS C function pose_from_wkb.
func PoseFromWKB(wkb []byte) *Pose {
	res := C.pose_from_wkb((*C.uint8_t)(unsafe.Pointer(&wkb[0])), C.size_t(len(wkb)))
	return &Pose{_inner: res}
}


// PoseFromHexwkb wraps MEOS C function pose_from_hexwkb.
func PoseFromHexwkb(hexwkb string) *Pose {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	res := C.pose_from_hexwkb(_c_hexwkb)
	return &Pose{_inner: res}
}


// PoseIn wraps MEOS C function pose_in.
func PoseIn(str string) *Pose {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.pose_in(_c_str)
	return &Pose{_inner: res}
}


// PoseOut wraps MEOS C function pose_out.
func PoseOut(pose *Pose, maxdd int) string {
	res := C.pose_out(pose._inner, C.int(maxdd))
	return C.GoString(res)
}


// PoseFromGeopose wraps MEOS C function pose_from_geopose.
func PoseFromGeopose(json string) *Pose {
	_c_json := C.CString(json)
	defer C.free(unsafe.Pointer(_c_json))
	res := C.pose_from_geopose(_c_json)
	return &Pose{_inner: res}
}


// PoseAsGeopose wraps MEOS C function pose_as_geopose.
func PoseAsGeopose(pose *Pose, conformance int, precision int) string {
	res := C.pose_as_geopose(pose._inner, C.int(conformance), C.int(precision))
	return C.GoString(res)
}


// TposeFromGeopose wraps MEOS C function tpose_from_geopose.
func TposeFromGeopose(json string) Temporal {
	_c_json := C.CString(json)
	defer C.free(unsafe.Pointer(_c_json))
	res := C.tpose_from_geopose(_c_json)
	return CreateTemporal(res)
}


// TposeAsGeopose wraps MEOS C function tpose_as_geopose.
func TposeAsGeopose(temp Temporal, conformance int, precision int) string {
	res := C.tpose_as_geopose(temp.Inner(), C.int(conformance), C.int(precision))
	return C.GoString(res)
}


// PoseApplyGeo wraps MEOS C function pose_apply_geo.
func PoseApplyGeo(pose *Pose, body *Geom) *Geom {
	res := C.pose_apply_geo(pose._inner, body._inner)
	return &Geom{_inner: res}
}


// TposeApplyGeo wraps MEOS C function tpose_apply_geo.
func TposeApplyGeo(temp Temporal, body *Geom) Temporal {
	res := C.tpose_apply_geo(temp.Inner(), body._inner)
	return CreateTemporal(res)
}


// PoseCopy wraps MEOS C function pose_copy.
func PoseCopy(pose *Pose) *Pose {
	res := C.pose_copy(pose._inner)
	return &Pose{_inner: res}
}


// PoseMake2d wraps MEOS C function pose_make_2d.
func PoseMake2d(x float64, y float64, theta float64, geodetic bool, srid int32) *Pose {
	res := C.pose_make_2d(C.double(x), C.double(y), C.double(theta), C.bool(geodetic), C.int32_t(srid))
	return &Pose{_inner: res}
}


// PoseMake3d wraps MEOS C function pose_make_3d.
func PoseMake3d(x float64, y float64, z float64, W float64, X float64, Y float64, Z float64, geodetic bool, srid int32) *Pose {
	res := C.pose_make_3d(C.double(x), C.double(y), C.double(z), C.double(W), C.double(X), C.double(Y), C.double(Z), C.bool(geodetic), C.int32_t(srid))
	return &Pose{_inner: res}
}


// PoseMakePoint2d wraps MEOS C function pose_make_point2d.
func PoseMakePoint2d(gs *Geom, theta float64) *Pose {
	res := C.pose_make_point2d(gs._inner, C.double(theta))
	return &Pose{_inner: res}
}


// PoseMakePoint3d wraps MEOS C function pose_make_point3d.
func PoseMakePoint3d(gs *Geom, W float64, X float64, Y float64, Z float64) *Pose {
	res := C.pose_make_point3d(gs._inner, C.double(W), C.double(X), C.double(Y), C.double(Z))
	return &Pose{_inner: res}
}


// PoseToPoint wraps MEOS C function pose_to_point.
func PoseToPoint(pose *Pose) *Geom {
	res := C.pose_to_point(pose._inner)
	return &Geom{_inner: res}
}


// PoseToSTBOX wraps MEOS C function pose_to_stbox.
func PoseToSTBOX(pose *Pose) *STBox {
	res := C.pose_to_stbox(pose._inner)
	return &STBox{_inner: res}
}


// PoseHash wraps MEOS C function pose_hash.
func PoseHash(pose *Pose) int {
	res := C.pose_hash(pose._inner)
	return int(res)
}


// PoseHashExtended wraps MEOS C function pose_hash_extended.
func PoseHashExtended(pose *Pose, seed int) int {
	res := C.pose_hash_extended(pose._inner, C.int(seed))
	return int(res)
}


// PoseOrientation wraps MEOS C function pose_orientation.
func PoseOrientation(pose *Pose) []float64 {
	var _out_count C.int
	res := C.pose_orientation(pose._inner, &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((*C.double)(unsafe.Pointer(res)), _n)
	_out := make([]float64, _n)
	for _i, _e := range _slice {
		_out[_i] = float64(_e)
	}
	return _out
}


// PoseRotation wraps MEOS C function pose_rotation.
func PoseRotation(pose *Pose) float64 {
	res := C.pose_rotation(pose._inner)
	return float64(res)
}


// PoseYaw wraps MEOS C function pose_yaw.
func PoseYaw(pose *Pose) float64 {
	res := C.pose_yaw(pose._inner)
	return float64(res)
}


// PosePitch wraps MEOS C function pose_pitch.
func PosePitch(pose *Pose) float64 {
	res := C.pose_pitch(pose._inner)
	return float64(res)
}


// PoseRoll wraps MEOS C function pose_roll.
func PoseRoll(pose *Pose) float64 {
	res := C.pose_roll(pose._inner)
	return float64(res)
}


// PoseAngularDistance wraps MEOS C function pose_angular_distance.
func PoseAngularDistance(pose1 *Pose, pose2 *Pose) float64 {
	res := C.pose_angular_distance(pose1._inner, pose2._inner)
	return float64(res)
}


// PoseNormalise wraps MEOS C function pose_normalise.
func PoseNormalise(pose *Pose) *Pose {
	res := C.pose_normalise(pose._inner)
	return &Pose{_inner: res}
}


// PoseRound wraps MEOS C function pose_round.
func PoseRound(pose *Pose, maxdd int) *Pose {
	res := C.pose_round(pose._inner, C.int(maxdd))
	return &Pose{_inner: res}
}


// PosearrRound wraps MEOS C function posearr_round.
func PosearrRound(posearr []*Pose, maxdd int) []*Pose {
	_c_posearr := make([]*C.Pose, len(posearr))
	for _i, _v := range posearr { _c_posearr[_i] = _v._inner }
	res := C.posearr_round((**C.Pose)(unsafe.Pointer(&_c_posearr[0])), C.int(len(posearr)), C.int(maxdd))
	_n := len(posearr)
	_slice := unsafe.Slice((**C.Pose)(unsafe.Pointer(res)), _n)
	_out := make([]*Pose, _n)
	for _i, _e := range _slice {
		_out[_i] = &Pose{_inner: _e}
	}
	return _out
}


// PoseSetSRID wraps MEOS C function pose_set_srid.
func PoseSetSRID(pose *Pose, srid int32) {
	C.pose_set_srid(pose._inner, C.int32_t(srid))
}


// PoseSRID wraps MEOS C function pose_srid.
func PoseSRID(pose *Pose) int32 {
	res := C.pose_srid(pose._inner)
	return int32(res)
}


// PoseTransform wraps MEOS C function pose_transform.
func PoseTransform(pose *Pose, srid int32) *Pose {
	res := C.pose_transform(pose._inner, C.int32_t(srid))
	return &Pose{_inner: res}
}


// PoseTransformPipeline wraps MEOS C function pose_transform_pipeline.
func PoseTransformPipeline(pose *Pose, pipelinestr string, srid int32, is_forward bool) *Pose {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	res := C.pose_transform_pipeline(pose._inner, _c_pipelinestr, C.int32_t(srid), C.bool(is_forward))
	return &Pose{_inner: res}
}


// PoseTstzspanToSTBOX wraps MEOS C function pose_tstzspan_to_stbox.
func PoseTstzspanToSTBOX(pose *Pose, s *Span) *STBox {
	res := C.pose_tstzspan_to_stbox(pose._inner, s._inner)
	return &STBox{_inner: res}
}


// PoseTimestamptzToSTBOX wraps MEOS C function pose_timestamptz_to_stbox.
func PoseTimestamptzToSTBOX(pose *Pose, t int64) *STBox {
	res := C.pose_timestamptz_to_stbox(pose._inner, C.TimestampTz(t))
	return &STBox{_inner: res}
}


// DistancePoseGeo wraps MEOS C function distance_pose_geo.
func DistancePoseGeo(pose *Pose, gs *Geom) float64 {
	res := C.distance_pose_geo(pose._inner, gs._inner)
	return float64(res)
}


// DistancePosePose wraps MEOS C function distance_pose_pose.
func DistancePosePose(pose1 *Pose, pose2 *Pose) float64 {
	res := C.distance_pose_pose(pose1._inner, pose2._inner)
	return float64(res)
}


// DistancePoseSTBOX wraps MEOS C function distance_pose_stbox.
func DistancePoseSTBOX(pose *Pose, box *STBox) float64 {
	res := C.distance_pose_stbox(pose._inner, box._inner)
	return float64(res)
}


// PoseCmp wraps MEOS C function pose_cmp.
func PoseCmp(pose1 *Pose, pose2 *Pose) int {
	res := C.pose_cmp(pose1._inner, pose2._inner)
	return int(res)
}


// PoseEq wraps MEOS C function pose_eq.
func PoseEq(pose1 *Pose, pose2 *Pose) bool {
	res := C.pose_eq(pose1._inner, pose2._inner)
	return bool(res)
}


// PoseGe wraps MEOS C function pose_ge.
func PoseGe(pose1 *Pose, pose2 *Pose) bool {
	res := C.pose_ge(pose1._inner, pose2._inner)
	return bool(res)
}


// PoseGt wraps MEOS C function pose_gt.
func PoseGt(pose1 *Pose, pose2 *Pose) bool {
	res := C.pose_gt(pose1._inner, pose2._inner)
	return bool(res)
}


// PoseLe wraps MEOS C function pose_le.
func PoseLe(pose1 *Pose, pose2 *Pose) bool {
	res := C.pose_le(pose1._inner, pose2._inner)
	return bool(res)
}


// PoseLt wraps MEOS C function pose_lt.
func PoseLt(pose1 *Pose, pose2 *Pose) bool {
	res := C.pose_lt(pose1._inner, pose2._inner)
	return bool(res)
}


// PoseNe wraps MEOS C function pose_ne.
func PoseNe(pose1 *Pose, pose2 *Pose) bool {
	res := C.pose_ne(pose1._inner, pose2._inner)
	return bool(res)
}


// PoseNsame wraps MEOS C function pose_nsame.
func PoseNsame(pose1 *Pose, pose2 *Pose) bool {
	res := C.pose_nsame(pose1._inner, pose2._inner)
	return bool(res)
}


// PoseSame wraps MEOS C function pose_same.
func PoseSame(pose1 *Pose, pose2 *Pose) bool {
	res := C.pose_same(pose1._inner, pose2._inner)
	return bool(res)
}


// PosesetIn wraps MEOS C function poseset_in.
func PosesetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.poseset_in(_c_str)
	return &Set{_inner: res}
}


// PosesetOut wraps MEOS C function poseset_out.
func PosesetOut(s *Set, maxdd int) string {
	res := C.poseset_out(s._inner, C.int(maxdd))
	return C.GoString(res)
}


// PosesetMake wraps MEOS C function poseset_make.
func PosesetMake(values []*Pose) *Set {
	_c_values := make([]*C.Pose, len(values))
	for _i, _v := range values { _c_values[_i] = _v._inner }
	res := C.poseset_make((**C.Pose)(unsafe.Pointer(&_c_values[0])), C.int(len(values)))
	return &Set{_inner: res}
}


// PoseToSet wraps MEOS C function pose_to_set.
func PoseToSet(pose *Pose) *Set {
	res := C.pose_to_set(pose._inner)
	return &Set{_inner: res}
}


// PosesetEndValue wraps MEOS C function poseset_end_value.
func PosesetEndValue(s *Set) *Pose {
	res := C.poseset_end_value(s._inner)
	return &Pose{_inner: res}
}


// PosesetStartValue wraps MEOS C function poseset_start_value.
func PosesetStartValue(s *Set) *Pose {
	res := C.poseset_start_value(s._inner)
	return &Pose{_inner: res}
}


// PosesetValueN wraps MEOS C function poseset_value_n.
func PosesetValueN(s *Set, n int) (bool, *Pose) {
	var _out_result *C.Pose
	res := C.poseset_value_n(s._inner, C.int(n), &_out_result)
	return bool(res), &Pose{_inner: _out_result}
}


// PosesetValues wraps MEOS C function poseset_values.
func PosesetValues(s *Set) []*Pose {
	var _out_count C.int
	res := C.poseset_values(s._inner, &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.Pose)(unsafe.Pointer(res)), _n)
	_out := make([]*Pose, _n)
	for _i, _e := range _slice {
		_out[_i] = &Pose{_inner: _e}
	}
	return _out
}


// ContainedPoseSet wraps MEOS C function contained_pose_set.
func ContainedPoseSet(pose *Pose, s *Set) bool {
	res := C.contained_pose_set(pose._inner, s._inner)
	return bool(res)
}


// ContainsSetPose wraps MEOS C function contains_set_pose.
func ContainsSetPose(s *Set, pose *Pose) bool {
	res := C.contains_set_pose(s._inner, pose._inner)
	return bool(res)
}


// IntersectionPoseSet wraps MEOS C function intersection_pose_set.
func IntersectionPoseSet(pose *Pose, s *Set) *Set {
	res := C.intersection_pose_set(pose._inner, s._inner)
	return &Set{_inner: res}
}


// IntersectionSetPose wraps MEOS C function intersection_set_pose.
func IntersectionSetPose(s *Set, pose *Pose) *Set {
	res := C.intersection_set_pose(s._inner, pose._inner)
	return &Set{_inner: res}
}


// MinusPoseSet wraps MEOS C function minus_pose_set.
func MinusPoseSet(pose *Pose, s *Set) *Set {
	res := C.minus_pose_set(pose._inner, s._inner)
	return &Set{_inner: res}
}


// MinusSetPose wraps MEOS C function minus_set_pose.
func MinusSetPose(s *Set, pose *Pose) *Set {
	res := C.minus_set_pose(s._inner, pose._inner)
	return &Set{_inner: res}
}


// PoseUnionTransfn wraps MEOS C function pose_union_transfn.
func PoseUnionTransfn(state *Set, pose *Pose) *Set {
	res := C.pose_union_transfn(state._inner, pose._inner)
	return &Set{_inner: res}
}


// UnionPoseSet wraps MEOS C function union_pose_set.
func UnionPoseSet(pose *Pose, s *Set) *Set {
	res := C.union_pose_set(pose._inner, s._inner)
	return &Set{_inner: res}
}


// UnionSetPose wraps MEOS C function union_set_pose.
func UnionSetPose(s *Set, pose *Pose) *Set {
	res := C.union_set_pose(s._inner, pose._inner)
	return &Set{_inner: res}
}


// TposeFromMFJSON wraps MEOS C function tpose_from_mfjson.
func TposeFromMFJSON(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tpose_from_mfjson(_c_str)
	return CreateTemporal(res)
}


// TposeIn wraps MEOS C function tpose_in.
func TposeIn(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tpose_in(_c_str)
	return CreateTemporal(res)
}


// TposeinstMake wraps MEOS C function tposeinst_make.
func TposeinstMake(pose *Pose, t int64) TInstant {
	res := C.tposeinst_make(pose._inner, C.TimestampTz(t))
	return TInstant{_inner: res}
}


// TposeFromBaseTemp wraps MEOS C function tpose_from_base_temp.
func TposeFromBaseTemp(pose *Pose, temp Temporal) Temporal {
	res := C.tpose_from_base_temp(pose._inner, temp.Inner())
	return CreateTemporal(res)
}


// TposeseqFromBaseTstzset wraps MEOS C function tposeseq_from_base_tstzset.
func TposeseqFromBaseTstzset(pose *Pose, s *Set) TSequence {
	res := C.tposeseq_from_base_tstzset(pose._inner, s._inner)
	return TSequence{_inner: res}
}


// TposeseqFromBaseTstzspan wraps MEOS C function tposeseq_from_base_tstzspan.
func TposeseqFromBaseTstzspan(pose *Pose, s *Span, interp Interpolation) TSequence {
	res := C.tposeseq_from_base_tstzspan(pose._inner, s._inner, C.interpType(interp))
	return TSequence{_inner: res}
}


// TposeseqsetFromBaseTstzspanset wraps MEOS C function tposeseqset_from_base_tstzspanset.
func TposeseqsetFromBaseTstzspanset(pose *Pose, ss *SpanSet, interp Interpolation) TSequenceSet {
	res := C.tposeseqset_from_base_tstzspanset(pose._inner, ss._inner, C.interpType(interp))
	return TSequenceSet{_inner: res}
}


// TposeMake wraps MEOS C function tpose_make.
func TposeMake(tpoint Temporal, tradius Temporal) Temporal {
	res := C.tpose_make(tpoint.Inner(), tradius.Inner())
	return CreateTemporal(res)
}


// TposeToTpoint wraps MEOS C function tpose_to_tpoint.
func TposeToTpoint(temp Temporal) Temporal {
	res := C.tpose_to_tpoint(temp.Inner())
	return CreateTemporal(res)
}


// TposeEndValue wraps MEOS C function tpose_end_value.
func TposeEndValue(temp Temporal) *Pose {
	res := C.tpose_end_value(temp.Inner())
	return &Pose{_inner: res}
}


// TposePoints wraps MEOS C function tpose_points.
func TposePoints(temp Temporal) *Set {
	res := C.tpose_points(temp.Inner())
	return &Set{_inner: res}
}


// TposeRotation wraps MEOS C function tpose_rotation.
func TposeRotation(temp Temporal) Temporal {
	res := C.tpose_rotation(temp.Inner())
	return CreateTemporal(res)
}


// TposeYaw wraps MEOS C function tpose_yaw.
func TposeYaw(temp Temporal) Temporal {
	res := C.tpose_yaw(temp.Inner())
	return CreateTemporal(res)
}


// TposePitch wraps MEOS C function tpose_pitch.
func TposePitch(temp Temporal) Temporal {
	res := C.tpose_pitch(temp.Inner())
	return CreateTemporal(res)
}


// TposeRoll wraps MEOS C function tpose_roll.
func TposeRoll(temp Temporal) Temporal {
	res := C.tpose_roll(temp.Inner())
	return CreateTemporal(res)
}


// TposeSpeed wraps MEOS C function tpose_speed.
func TposeSpeed(temp Temporal) Temporal {
	res := C.tpose_speed(temp.Inner())
	return CreateTemporal(res)
}


// TposeAngularSpeed wraps MEOS C function tpose_angular_speed.
func TposeAngularSpeed(temp Temporal) Temporal {
	res := C.tpose_angular_speed(temp.Inner())
	return CreateTemporal(res)
}


// TposeStartValue wraps MEOS C function tpose_start_value.
func TposeStartValue(temp Temporal) *Pose {
	res := C.tpose_start_value(temp.Inner())
	return &Pose{_inner: res}
}


// TposeTrajectory wraps MEOS C function tpose_trajectory.
func TposeTrajectory(temp Temporal) *Geom {
	res := C.tpose_trajectory(temp.Inner())
	return &Geom{_inner: res}
}


// TposeValueAtTimestamptz wraps MEOS C function tpose_value_at_timestamptz.
func TposeValueAtTimestamptz(temp Temporal, t int64, strict bool) (bool, *Pose) {
	var _out_result *C.Pose
	res := C.tpose_value_at_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(strict), &_out_result)
	return bool(res), &Pose{_inner: _out_result}
}


// TposeValueN wraps MEOS C function tpose_value_n.
func TposeValueN(temp Temporal, n int) (bool, *Pose) {
	var _out_result *C.Pose
	res := C.tpose_value_n(temp.Inner(), C.int(n), &_out_result)
	return bool(res), &Pose{_inner: _out_result}
}


// TposeValues wraps MEOS C function tpose_values.
func TposeValues(temp Temporal) []*Pose {
	var _out_count C.int
	res := C.tpose_values(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.Pose)(unsafe.Pointer(res)), _n)
	_out := make([]*Pose, _n)
	for _i, _e := range _slice {
		_out[_i] = &Pose{_inner: _e}
	}
	return _out
}


// TposeAtGeom wraps MEOS C function tpose_at_geom.
func TposeAtGeom(temp Temporal, gs *Geom) Temporal {
	res := C.tpose_at_geom(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TposeAtSTBOX wraps MEOS C function tpose_at_stbox.
func TposeAtSTBOX(temp Temporal, box *STBox, border_inc bool) Temporal {
	res := C.tpose_at_stbox(temp.Inner(), box._inner, C.bool(border_inc))
	return CreateTemporal(res)
}


// TposeAtPose wraps MEOS C function tpose_at_pose.
func TposeAtPose(temp Temporal, pose *Pose) Temporal {
	res := C.tpose_at_pose(temp.Inner(), pose._inner)
	return CreateTemporal(res)
}


// TposeMinusGeom wraps MEOS C function tpose_minus_geom.
func TposeMinusGeom(temp Temporal, gs *Geom) Temporal {
	res := C.tpose_minus_geom(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TposeMinusPose wraps MEOS C function tpose_minus_pose.
func TposeMinusPose(temp Temporal, pose *Pose) Temporal {
	res := C.tpose_minus_pose(temp.Inner(), pose._inner)
	return CreateTemporal(res)
}


// TposeMinusSTBOX wraps MEOS C function tpose_minus_stbox.
func TposeMinusSTBOX(temp Temporal, box *STBox, border_inc bool) Temporal {
	res := C.tpose_minus_stbox(temp.Inner(), box._inner, C.bool(border_inc))
	return CreateTemporal(res)
}


// TdistanceTposePose wraps MEOS C function tdistance_tpose_pose.
func TdistanceTposePose(temp Temporal, pose *Pose) Temporal {
	res := C.tdistance_tpose_pose(temp.Inner(), pose._inner)
	return CreateTemporal(res)
}


// TdistanceTposeGeo wraps MEOS C function tdistance_tpose_geo.
func TdistanceTposeGeo(temp Temporal, gs *Geom) Temporal {
	res := C.tdistance_tpose_geo(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TdistanceTposeTpose wraps MEOS C function tdistance_tpose_tpose.
func TdistanceTposeTpose(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tdistance_tpose_tpose(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// NadTposeGeo wraps MEOS C function nad_tpose_geo.
func NadTposeGeo(temp Temporal, gs *Geom) float64 {
	res := C.nad_tpose_geo(temp.Inner(), gs._inner)
	return float64(res)
}


// NadTposePose wraps MEOS C function nad_tpose_pose.
func NadTposePose(temp Temporal, pose *Pose) float64 {
	res := C.nad_tpose_pose(temp.Inner(), pose._inner)
	return float64(res)
}


// NadTposeSTBOX wraps MEOS C function nad_tpose_stbox.
func NadTposeSTBOX(temp Temporal, box *STBox) float64 {
	res := C.nad_tpose_stbox(temp.Inner(), box._inner)
	return float64(res)
}


// NadTposeTpose wraps MEOS C function nad_tpose_tpose.
func NadTposeTpose(temp1 Temporal, temp2 Temporal) float64 {
	res := C.nad_tpose_tpose(temp1.Inner(), temp2.Inner())
	return float64(res)
}


// NaiTposeGeo wraps MEOS C function nai_tpose_geo.
func NaiTposeGeo(temp Temporal, gs *Geom) TInstant {
	res := C.nai_tpose_geo(temp.Inner(), gs._inner)
	return TInstant{_inner: res}
}


// NaiTposePose wraps MEOS C function nai_tpose_pose.
func NaiTposePose(temp Temporal, pose *Pose) TInstant {
	res := C.nai_tpose_pose(temp.Inner(), pose._inner)
	return TInstant{_inner: res}
}


// NaiTposeTpose wraps MEOS C function nai_tpose_tpose.
func NaiTposeTpose(temp1 Temporal, temp2 Temporal) TInstant {
	res := C.nai_tpose_tpose(temp1.Inner(), temp2.Inner())
	return TInstant{_inner: res}
}


// ShortestlineTposeGeo wraps MEOS C function shortestline_tpose_geo.
func ShortestlineTposeGeo(temp Temporal, gs *Geom) *Geom {
	res := C.shortestline_tpose_geo(temp.Inner(), gs._inner)
	return &Geom{_inner: res}
}


// ShortestlineTposePose wraps MEOS C function shortestline_tpose_pose.
func ShortestlineTposePose(temp Temporal, pose *Pose) *Geom {
	res := C.shortestline_tpose_pose(temp.Inner(), pose._inner)
	return &Geom{_inner: res}
}


// ShortestlineTposeTpose wraps MEOS C function shortestline_tpose_tpose.
func ShortestlineTposeTpose(temp1 Temporal, temp2 Temporal) *Geom {
	res := C.shortestline_tpose_tpose(temp1.Inner(), temp2.Inner())
	return &Geom{_inner: res}
}


// AlwaysEqPoseTpose wraps MEOS C function always_eq_pose_tpose.
func AlwaysEqPoseTpose(pose *Pose, temp Temporal) int {
	res := C.always_eq_pose_tpose(pose._inner, temp.Inner())
	return int(res)
}


// AlwaysEqTposePose wraps MEOS C function always_eq_tpose_pose.
func AlwaysEqTposePose(temp Temporal, pose *Pose) int {
	res := C.always_eq_tpose_pose(temp.Inner(), pose._inner)
	return int(res)
}


// AlwaysEqTposeTpose wraps MEOS C function always_eq_tpose_tpose.
func AlwaysEqTposeTpose(temp1 Temporal, temp2 Temporal) int {
	res := C.always_eq_tpose_tpose(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AlwaysNePoseTpose wraps MEOS C function always_ne_pose_tpose.
func AlwaysNePoseTpose(pose *Pose, temp Temporal) int {
	res := C.always_ne_pose_tpose(pose._inner, temp.Inner())
	return int(res)
}


// AlwaysNeTposePose wraps MEOS C function always_ne_tpose_pose.
func AlwaysNeTposePose(temp Temporal, pose *Pose) int {
	res := C.always_ne_tpose_pose(temp.Inner(), pose._inner)
	return int(res)
}


// AlwaysNeTposeTpose wraps MEOS C function always_ne_tpose_tpose.
func AlwaysNeTposeTpose(temp1 Temporal, temp2 Temporal) int {
	res := C.always_ne_tpose_tpose(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverEqPoseTpose wraps MEOS C function ever_eq_pose_tpose.
func EverEqPoseTpose(pose *Pose, temp Temporal) int {
	res := C.ever_eq_pose_tpose(pose._inner, temp.Inner())
	return int(res)
}


// EverEqTposePose wraps MEOS C function ever_eq_tpose_pose.
func EverEqTposePose(temp Temporal, pose *Pose) int {
	res := C.ever_eq_tpose_pose(temp.Inner(), pose._inner)
	return int(res)
}


// EverEqTposeTpose wraps MEOS C function ever_eq_tpose_tpose.
func EverEqTposeTpose(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_eq_tpose_tpose(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverNePoseTpose wraps MEOS C function ever_ne_pose_tpose.
func EverNePoseTpose(pose *Pose, temp Temporal) int {
	res := C.ever_ne_pose_tpose(pose._inner, temp.Inner())
	return int(res)
}


// EverNeTposePose wraps MEOS C function ever_ne_tpose_pose.
func EverNeTposePose(temp Temporal, pose *Pose) int {
	res := C.ever_ne_tpose_pose(temp.Inner(), pose._inner)
	return int(res)
}


// EverNeTposeTpose wraps MEOS C function ever_ne_tpose_tpose.
func EverNeTposeTpose(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_ne_tpose_tpose(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TeqPoseTpose wraps MEOS C function teq_pose_tpose.
func TeqPoseTpose(pose *Pose, temp Temporal) Temporal {
	res := C.teq_pose_tpose(pose._inner, temp.Inner())
	return CreateTemporal(res)
}


// TeqTposePose wraps MEOS C function teq_tpose_pose.
func TeqTposePose(temp Temporal, pose *Pose) Temporal {
	res := C.teq_tpose_pose(temp.Inner(), pose._inner)
	return CreateTemporal(res)
}


// TnePoseTpose wraps MEOS C function tne_pose_tpose.
func TnePoseTpose(pose *Pose, temp Temporal) Temporal {
	res := C.tne_pose_tpose(pose._inner, temp.Inner())
	return CreateTemporal(res)
}


// TneTposePose wraps MEOS C function tne_tpose_pose.
func TneTposePose(temp Temporal, pose *Pose) Temporal {
	res := C.tne_tpose_pose(temp.Inner(), pose._inner)
	return CreateTemporal(res)
}

