package generated

// #include <stddef.h>
import "C"
import (
	"unsafe"

	"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

// TODO pose_as_ewkt: unsupported param const Pose *
// func PoseAsEWKT(...) { /* not yet handled by codegen */ }


// TODO pose_as_hexwkb: unsupported param const Pose *
// func PoseAsHexwkb(...) { /* not yet handled by codegen */ }


// TODO pose_as_text: unsupported param const Pose *
// func PoseAsText(...) { /* not yet handled by codegen */ }


// TODO pose_as_wkb: unsupported param const Pose *
// func PoseAsWKB(...) { /* not yet handled by codegen */ }


// TODO pose_from_wkb: unsupported return type Pose *
// func PoseFromWKB(...) { /* not yet handled by codegen */ }


// TODO pose_from_hexwkb: unsupported return type Pose *
// func PoseFromHexwkb(...) { /* not yet handled by codegen */ }


// TODO pose_in: unsupported return type Pose *
// func PoseIn(...) { /* not yet handled by codegen */ }


// TODO pose_out: unsupported param const Pose *
// func PoseOut(...) { /* not yet handled by codegen */ }


// TODO pose_copy: unsupported return type Pose *
// func PoseCopy(...) { /* not yet handled by codegen */ }


// TODO pose_make_2d: unsupported return type Pose *
// func PoseMake2d(...) { /* not yet handled by codegen */ }


// TODO pose_make_3d: unsupported return type Pose *
// func PoseMake3d(...) { /* not yet handled by codegen */ }


// TODO pose_make_point2d: unsupported return type Pose *
// func PoseMakePoint2d(...) { /* not yet handled by codegen */ }


// TODO pose_make_point3d: unsupported return type Pose *
// func PoseMakePoint3d(...) { /* not yet handled by codegen */ }


// TODO pose_to_point: unsupported param const Pose *
// func PoseToPoint(...) { /* not yet handled by codegen */ }


// TODO pose_to_stbox: unsupported param const Pose *
// func PoseToSTBOX(...) { /* not yet handled by codegen */ }


// TODO pose_hash: unsupported param const Pose *
// func PoseHash(...) { /* not yet handled by codegen */ }


// TODO pose_hash_extended: unsupported param const Pose *
// func PoseHashExtended(...) { /* not yet handled by codegen */ }


// TODO pose_orientation: unsupported return type double *
// func PoseOrientation(...) { /* not yet handled by codegen */ }


// TODO pose_rotation: unsupported param const Pose *
// func PoseRotation(...) { /* not yet handled by codegen */ }


// TODO pose_round: unsupported return type Pose *
// func PoseRound(...) { /* not yet handled by codegen */ }


// TODO posearr_round: unsupported return type Pose **
// func PosearrRound(...) { /* not yet handled by codegen */ }


// TODO pose_set_srid: unsupported param Pose *
// func PoseSetSRID(...) { /* not yet handled by codegen */ }


// TODO pose_srid: unsupported param const Pose *
// func PoseSRID(...) { /* not yet handled by codegen */ }


// TODO pose_transform: unsupported return type Pose *
// func PoseTransform(...) { /* not yet handled by codegen */ }


// TODO pose_transform_pipeline: unsupported return type Pose *
// func PoseTransformPipeline(...) { /* not yet handled by codegen */ }


// TODO pose_tstzspan_to_stbox: unsupported param const Pose *
// func PoseTstzspanToSTBOX(...) { /* not yet handled by codegen */ }


// TODO pose_timestamptz_to_stbox: unsupported param const Pose *
// func PoseTimestamptzToSTBOX(...) { /* not yet handled by codegen */ }


// TODO distance_pose_geo: unsupported param const Pose *
// func DistancePoseGeo(...) { /* not yet handled by codegen */ }


// TODO distance_pose_pose: unsupported param const Pose *
// func DistancePosePose(...) { /* not yet handled by codegen */ }


// TODO distance_pose_stbox: unsupported param const Pose *
// func DistancePoseSTBOX(...) { /* not yet handled by codegen */ }


// TODO pose_cmp: unsupported param const Pose *
// func PoseCmp(...) { /* not yet handled by codegen */ }


// TODO pose_eq: unsupported param const Pose *
// func PoseEq(...) { /* not yet handled by codegen */ }


// TODO pose_ge: unsupported param const Pose *
// func PoseGe(...) { /* not yet handled by codegen */ }


// TODO pose_gt: unsupported param const Pose *
// func PoseGt(...) { /* not yet handled by codegen */ }


// TODO pose_le: unsupported param const Pose *
// func PoseLe(...) { /* not yet handled by codegen */ }


// TODO pose_lt: unsupported param const Pose *
// func PoseLt(...) { /* not yet handled by codegen */ }


// TODO pose_ne: unsupported param const Pose *
// func PoseNe(...) { /* not yet handled by codegen */ }


// TODO pose_nsame: unsupported param const Pose *
// func PoseNsame(...) { /* not yet handled by codegen */ }


// TODO pose_same: unsupported param const Pose *
// func PoseSame(...) { /* not yet handled by codegen */ }


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


// TODO poseset_make: unsupported param const Pose **
// func PosesetMake(...) { /* not yet handled by codegen */ }


// TODO pose_to_set: unsupported param const Pose *
// func PoseToSet(...) { /* not yet handled by codegen */ }


// TODO poseset_end_value: unsupported return type Pose *
// func PosesetEndValue(...) { /* not yet handled by codegen */ }


// TODO poseset_start_value: unsupported return type Pose *
// func PosesetStartValue(...) { /* not yet handled by codegen */ }


// TODO poseset_value_n: unhandled OUTPUT_SCALAR shape Pose **
// func PosesetValueN(...) { /* not yet handled by codegen */ }


// TODO poseset_values: unsupported return type Pose **
// func PosesetValues(...) { /* not yet handled by codegen */ }


// TODO contained_pose_set: unsupported param const Pose *
// func ContainedPoseSet(...) { /* not yet handled by codegen */ }


// TODO contains_set_pose: unsupported param Pose *
// func ContainsSetPose(...) { /* not yet handled by codegen */ }


// TODO intersection_pose_set: unsupported param const Pose *
// func IntersectionPoseSet(...) { /* not yet handled by codegen */ }


// TODO intersection_set_pose: unsupported param const Pose *
// func IntersectionSetPose(...) { /* not yet handled by codegen */ }


// TODO minus_pose_set: unsupported param const Pose *
// func MinusPoseSet(...) { /* not yet handled by codegen */ }


// TODO minus_set_pose: unsupported param const Pose *
// func MinusSetPose(...) { /* not yet handled by codegen */ }


// TODO pose_union_transfn: unsupported param const Pose *
// func PoseUnionTransfn(...) { /* not yet handled by codegen */ }


// TODO union_pose_set: unsupported param const Pose *
// func UnionPoseSet(...) { /* not yet handled by codegen */ }


// TODO union_set_pose: unsupported param const Pose *
// func UnionSetPose(...) { /* not yet handled by codegen */ }


// TposeIn wraps MEOS C function tpose_in.
func TposeIn(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tpose_in(_c_str)
	return CreateTemporal(res)
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


// TODO tpose_end_value: unsupported return type Pose *
// func TposeEndValue(...) { /* not yet handled by codegen */ }


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


// TODO tpose_start_value: unsupported return type Pose *
// func TposeStartValue(...) { /* not yet handled by codegen */ }


// TposeTrajectory wraps MEOS C function tpose_trajectory.
func TposeTrajectory(temp Temporal) *Geom {
	res := C.tpose_trajectory(temp.Inner())
	return &Geom{_inner: res}
}


// TODO tpose_value_at_timestamptz: unhandled OUTPUT_SCALAR shape Pose **
// func TposeValueAtTimestamptz(...) { /* not yet handled by codegen */ }


// TODO tpose_value_n: unhandled OUTPUT_SCALAR shape Pose **
// func TposeValueN(...) { /* not yet handled by codegen */ }


// TODO tpose_values: unsupported return type Pose **
// func TposeValues(...) { /* not yet handled by codegen */ }


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


// TODO tpose_at_pose: unsupported param const Pose *
// func TposeAtPose(...) { /* not yet handled by codegen */ }


// TposeMinusGeom wraps MEOS C function tpose_minus_geom.
func TposeMinusGeom(temp Temporal, gs *Geom) Temporal {
	res := C.tpose_minus_geom(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}


// TODO tpose_minus_pose: unsupported param const Pose *
// func TposeMinusPose(...) { /* not yet handled by codegen */ }


// TposeMinusSTBOX wraps MEOS C function tpose_minus_stbox.
func TposeMinusSTBOX(temp Temporal, box *STBox, border_inc bool) Temporal {
	res := C.tpose_minus_stbox(temp.Inner(), box._inner, C.bool(border_inc))
	return CreateTemporal(res)
}


// TODO tdistance_tpose_pose: unsupported param const Pose *
// func TdistanceTposePose(...) { /* not yet handled by codegen */ }


// TdistanceTposePoint wraps MEOS C function tdistance_tpose_point.
func TdistanceTposePoint(temp Temporal, gs *Geom) Temporal {
	res := C.tdistance_tpose_point(temp.Inner(), gs._inner)
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


// TODO nad_tpose_pose: unsupported param const Pose *
// func NadTposePose(...) { /* not yet handled by codegen */ }


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


// TODO nai_tpose_pose: unsupported param const Pose *
// func NaiTposePose(...) { /* not yet handled by codegen */ }


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


// TODO shortestline_tpose_pose: unsupported param const Pose *
// func ShortestlineTposePose(...) { /* not yet handled by codegen */ }


// ShortestlineTposeTpose wraps MEOS C function shortestline_tpose_tpose.
func ShortestlineTposeTpose(temp1 Temporal, temp2 Temporal) *Geom {
	res := C.shortestline_tpose_tpose(temp1.Inner(), temp2.Inner())
	return &Geom{_inner: res}
}


// TODO always_eq_pose_tpose: unsupported param const Pose *
// func AlwaysEqPoseTpose(...) { /* not yet handled by codegen */ }


// TODO always_eq_tpose_pose: unsupported param const Pose *
// func AlwaysEqTposePose(...) { /* not yet handled by codegen */ }


// AlwaysEqTposeTpose wraps MEOS C function always_eq_tpose_tpose.
func AlwaysEqTposeTpose(temp1 Temporal, temp2 Temporal) int {
	res := C.always_eq_tpose_tpose(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO always_ne_pose_tpose: unsupported param const Pose *
// func AlwaysNePoseTpose(...) { /* not yet handled by codegen */ }


// TODO always_ne_tpose_pose: unsupported param const Pose *
// func AlwaysNeTposePose(...) { /* not yet handled by codegen */ }


// AlwaysNeTposeTpose wraps MEOS C function always_ne_tpose_tpose.
func AlwaysNeTposeTpose(temp1 Temporal, temp2 Temporal) int {
	res := C.always_ne_tpose_tpose(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO ever_eq_pose_tpose: unsupported param const Pose *
// func EverEqPoseTpose(...) { /* not yet handled by codegen */ }


// TODO ever_eq_tpose_pose: unsupported param const Pose *
// func EverEqTposePose(...) { /* not yet handled by codegen */ }


// EverEqTposeTpose wraps MEOS C function ever_eq_tpose_tpose.
func EverEqTposeTpose(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_eq_tpose_tpose(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO ever_ne_pose_tpose: unsupported param const Pose *
// func EverNePoseTpose(...) { /* not yet handled by codegen */ }


// TODO ever_ne_tpose_pose: unsupported param const Pose *
// func EverNeTposePose(...) { /* not yet handled by codegen */ }


// EverNeTposeTpose wraps MEOS C function ever_ne_tpose_tpose.
func EverNeTposeTpose(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_ne_tpose_tpose(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TODO teq_pose_tpose: unsupported param const Pose *
// func TeqPoseTpose(...) { /* not yet handled by codegen */ }


// TODO teq_tpose_pose: unsupported param const Pose *
// func TeqTposePose(...) { /* not yet handled by codegen */ }


// TODO tne_pose_tpose: unsupported param const Pose *
// func TnePoseTpose(...) { /* not yet handled by codegen */ }


// TODO tne_tpose_pose: unsupported param const Pose *
// func TneTposePose(...) { /* not yet handled by codegen */ }

