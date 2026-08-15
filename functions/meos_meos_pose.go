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
func PoseAsEWKT(pose *Pose, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_as_ewkt(pose._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// PoseAsHexwkb wraps MEOS C function pose_as_hexwkb.
func PoseAsHexwkb(pose *Pose, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_as_hexwkb(pose._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// PoseAsText wraps MEOS C function pose_as_text.
func PoseAsText(pose *Pose, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_as_text(pose._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// PoseAsWKB wraps MEOS C function pose_as_wkb.
func PoseAsWKB(pose *Pose, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_as_wkb(pose._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PoseFromWKB wraps MEOS C function pose_from_wkb.
func PoseFromWKB(wkb unsafe.Pointer, size uint) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseFromHexwkb wraps MEOS C function pose_from_hexwkb.
func PoseFromHexwkb(hexwkb string) (_r0 *Pose, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.pose_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseIn wraps MEOS C function pose_in.
func PoseIn(str string) (_r0 *Pose, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.pose_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseOut wraps MEOS C function pose_out.
func PoseOut(pose *Pose, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_out(pose._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// PoseFromGeopose wraps MEOS C function pose_from_geopose.
func PoseFromGeopose(json string) (_r0 *Pose, _err error) {
	_c_json := C.CString(json)
	defer C.free(unsafe.Pointer(_c_json))
	C.meos_errno_reset()
	_cret := C.pose_from_geopose(_c_json)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseAsGeopose wraps MEOS C function pose_as_geopose.
func PoseAsGeopose(pose *Pose, conformance int, precision int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_as_geopose(pose._inner, C.int(conformance), C.int(precision))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// TposeFromGeopose wraps MEOS C function tpose_from_geopose.
func TposeFromGeopose(json string) (_r0 *Temporal, _err error) {
	_c_json := C.CString(json)
	defer C.free(unsafe.Pointer(_c_json))
	C.meos_errno_reset()
	_cret := C.tpose_from_geopose(_c_json)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeAsGeopose wraps MEOS C function tpose_as_geopose.
func TposeAsGeopose(temp *Temporal, conformance int, precision int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_as_geopose(temp._inner, C.int(conformance), C.int(precision))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// TposeAsGeoposeStreamHeader wraps MEOS C function tpose_as_geopose_stream_header.
func TposeAsGeoposeStreamHeader(temp *Temporal, precision int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_as_geopose_stream_header(temp._inner, C.int(precision))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// TposeAsGeoposeStreamElement wraps MEOS C function tpose_as_geopose_stream_element.
func TposeAsGeoposeStreamElement(temp *Temporal, inst *TInstant, precision int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_as_geopose_stream_element(temp._inner, inst._inner, C.int(precision))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// PoseApplyGeo wraps MEOS C function pose_apply_geo.
func PoseApplyGeo(pose *Pose, body *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_apply_geo(pose._inner, body._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TposeApplyGeo wraps MEOS C function tpose_apply_geo.
func TposeApplyGeo(temp *Temporal, body *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_apply_geo(temp._inner, body._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// PoseCopy wraps MEOS C function pose_copy.
func PoseCopy(pose *Pose) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_copy(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseMake2d wraps MEOS C function pose_make_2d.
func PoseMake2d(x float64, y float64, theta float64, geodetic bool, srid int32) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_make_2d(C.double(x), C.double(y), C.double(theta), C.bool(geodetic), C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseMake3d wraps MEOS C function pose_make_3d.
func PoseMake3d(x float64, y float64, z float64, W float64, X float64, Y float64, Z float64, geodetic bool, srid int32) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_make_3d(C.double(x), C.double(y), C.double(z), C.double(W), C.double(X), C.double(Y), C.double(Z), C.bool(geodetic), C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseMakePoint2d wraps MEOS C function pose_make_point2d.
func PoseMakePoint2d(gs *Geom, theta float64) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_make_point2d(gs._inner, C.double(theta))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseMakePoint3d wraps MEOS C function pose_make_point3d.
func PoseMakePoint3d(gs *Geom, W float64, X float64, Y float64, Z float64) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_make_point3d(gs._inner, C.double(W), C.double(X), C.double(Y), C.double(Z))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseToPoint wraps MEOS C function pose_to_point.
func PoseToPoint(pose *Pose) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_to_point(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// PoseToSTBOX wraps MEOS C function pose_to_stbox.
func PoseToSTBOX(pose *Pose) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_to_stbox(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// PoseHash wraps MEOS C function pose_hash.
func PoseHash(pose *Pose) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_hash(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// PoseHashExtended wraps MEOS C function pose_hash_extended.
func PoseHashExtended(pose *Pose, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_hash_extended(pose._inner, C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// PoseOrientation wraps MEOS C function pose_orientation.
func PoseOrientation(pose *Pose, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_orientation(pose._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PoseRotation wraps MEOS C function pose_rotation.
func PoseRotation(pose *Pose) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_rotation(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// PoseYaw wraps MEOS C function pose_yaw.
func PoseYaw(pose *Pose) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_yaw(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// PosePitch wraps MEOS C function pose_pitch.
func PosePitch(pose *Pose) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_pitch(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// PoseRoll wraps MEOS C function pose_roll.
func PoseRoll(pose *Pose) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_roll(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// PoseAngularDistance wraps MEOS C function pose_angular_distance.
func PoseAngularDistance(pose1 *Pose, pose2 *Pose) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_angular_distance(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// PoseNormalize wraps MEOS C function pose_normalize.
func PoseNormalize(pose *Pose) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_normalize(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseRound wraps MEOS C function pose_round.
func PoseRound(pose *Pose, maxdd int) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_round(pose._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PosearrRound wraps MEOS C function posearr_round.
func PosearrRound(posearr unsafe.Pointer, count int, maxdd int) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.posearr_round((**C.Pose)(unsafe.Pointer(posearr)), C.int(count), C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PoseSetSRID wraps MEOS C function pose_set_srid.
func PoseSetSRID(pose *Pose, srid int32) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_set_srid(pose._inner, C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseSRID wraps MEOS C function pose_srid.
func PoseSRID(pose *Pose) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_srid(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// PoseTransform wraps MEOS C function pose_transform.
func PoseTransform(pose *Pose, srid int32) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_transform(pose._inner, C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseTransformPipeline wraps MEOS C function pose_transform_pipeline.
func PoseTransformPipeline(pose *Pose, pipelinestr string, srid int32, is_forward bool) (_r0 *Pose, _err error) {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	C.meos_errno_reset()
	_cret := C.pose_transform_pipeline(pose._inner, _c_pipelinestr, C.int32_t(srid), C.bool(is_forward))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseTstzspanToSTBOX wraps MEOS C function pose_tstzspan_to_stbox.
func PoseTstzspanToSTBOX(pose *Pose, s *Span) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_tstzspan_to_stbox(pose._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// PoseTimestamptzToSTBOX wraps MEOS C function pose_timestamptz_to_stbox.
func PoseTimestamptzToSTBOX(pose *Pose, t int64) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_timestamptz_to_stbox(pose._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// DistancePoseGeo wraps MEOS C function distance_pose_geo.
func DistancePoseGeo(pose *Pose, gs *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_pose_geo(pose._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistancePosePose wraps MEOS C function distance_pose_pose.
func DistancePosePose(pose1 *Pose, pose2 *Pose) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_pose_pose(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistancePoseSTBOX wraps MEOS C function distance_pose_stbox.
func DistancePoseSTBOX(pose *Pose, box *STBox) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_pose_stbox(pose._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// PoseCmp wraps MEOS C function pose_cmp.
func PoseCmp(pose1 *Pose, pose2 *Pose) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_cmp(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// PoseEq wraps MEOS C function pose_eq.
func PoseEq(pose1 *Pose, pose2 *Pose) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_eq(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PoseGe wraps MEOS C function pose_ge.
func PoseGe(pose1 *Pose, pose2 *Pose) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_ge(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PoseGt wraps MEOS C function pose_gt.
func PoseGt(pose1 *Pose, pose2 *Pose) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_gt(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PoseLe wraps MEOS C function pose_le.
func PoseLe(pose1 *Pose, pose2 *Pose) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_le(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PoseLt wraps MEOS C function pose_lt.
func PoseLt(pose1 *Pose, pose2 *Pose) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_lt(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PoseNe wraps MEOS C function pose_ne.
func PoseNe(pose1 *Pose, pose2 *Pose) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_ne(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PoseNsame wraps MEOS C function pose_nsame.
func PoseNsame(pose1 *Pose, pose2 *Pose) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_nsame(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PoseSame wraps MEOS C function pose_same.
func PoseSame(pose1 *Pose, pose2 *Pose) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_same(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PosesetIn wraps MEOS C function poseset_in.
func PosesetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.poseset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PosesetOut wraps MEOS C function poseset_out.
func PosesetOut(s *Set, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.poseset_out(s._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// PosesetMake wraps MEOS C function poseset_make.
func PosesetMake(values unsafe.Pointer, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.poseset_make((**C.Pose)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PoseToSet wraps MEOS C function pose_to_set.
func PoseToSet(pose *Pose) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_to_set(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PosesetEndValue wraps MEOS C function poseset_end_value.
func PosesetEndValue(s *Set) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.poseset_end_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PosesetStartValue wraps MEOS C function poseset_start_value.
func PosesetStartValue(s *Set) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.poseset_start_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PosesetValueN wraps MEOS C function poseset_value_n.
func PosesetValueN(s *Set, n int) (_r0 bool, _r1 *Pose, _err error) {
	var _out_result *C.Pose
	C.meos_errno_reset()
	_cret := C.poseset_value_n(s._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Pose{_inner: _out_result}, nil
}


// PosesetValues wraps MEOS C function poseset_values.
func PosesetValues(s *Set, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.poseset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// ContainedPoseSet wraps MEOS C function contained_pose_set.
func ContainedPoseSet(pose *Pose, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_pose_set(pose._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSetPose wraps MEOS C function contains_set_pose.
func ContainsSetPose(s *Set, pose *Pose) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_set_pose(s._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// IntersectionPoseSet wraps MEOS C function intersection_pose_set.
func IntersectionPoseSet(pose *Pose, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_pose_set(pose._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetPose wraps MEOS C function intersection_set_pose.
func IntersectionSetPose(s *Set, pose *Pose) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_set_pose(s._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusPoseSet wraps MEOS C function minus_pose_set.
func MinusPoseSet(pose *Pose, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_pose_set(pose._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSetPose wraps MEOS C function minus_set_pose.
func MinusSetPose(s *Set, pose *Pose) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_set_pose(s._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PoseUnionTransfn wraps MEOS C function pose_union_transfn.
func PoseUnionTransfn(state *Set, pose *Pose) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_union_transfn(state._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionPoseSet wraps MEOS C function union_pose_set.
func UnionPoseSet(pose *Pose, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_pose_set(pose._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSetPose wraps MEOS C function union_set_pose.
func UnionSetPose(s *Set, pose *Pose) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_set_pose(s._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TposeFromMFJSON wraps MEOS C function tpose_from_mfjson.
func TposeFromMFJSON(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tpose_from_mfjson(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeIn wraps MEOS C function tpose_in.
func TposeIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tpose_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeinstMake wraps MEOS C function tposeinst_make.
func TposeinstMake(pose *Pose, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tposeinst_make(pose._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TposeFromBaseTemp wraps MEOS C function tpose_from_base_temp.
func TposeFromBaseTemp(pose *Pose, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_from_base_temp(pose._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeseqFromBaseTstzset wraps MEOS C function tposeseq_from_base_tstzset.
func TposeseqFromBaseTstzset(pose *Pose, s *Set) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tposeseq_from_base_tstzset(pose._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TposeseqFromBaseTstzspan wraps MEOS C function tposeseq_from_base_tstzspan.
func TposeseqFromBaseTstzspan(pose *Pose, s *Span, interp Interpolation) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tposeseq_from_base_tstzspan(pose._inner, s._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TposeseqsetFromBaseTstzspanset wraps MEOS C function tposeseqset_from_base_tstzspanset.
func TposeseqsetFromBaseTstzspanset(pose *Pose, ss *SpanSet, interp Interpolation) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tposeseqset_from_base_tstzspanset(pose._inner, ss._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TposeMake wraps MEOS C function tpose_make.
func TposeMake(tpoint *Temporal, tradius *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_make(tpoint._inner, tradius._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeToTpoint wraps MEOS C function tpose_to_tpoint.
func TposeToTpoint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_to_tpoint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeEndValue wraps MEOS C function tpose_end_value.
func TposeEndValue(temp *Temporal) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_end_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// TposePoints wraps MEOS C function tpose_points.
func TposePoints(temp *Temporal) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_points(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TposeRotation wraps MEOS C function tpose_rotation.
func TposeRotation(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_rotation(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeYaw wraps MEOS C function tpose_yaw.
func TposeYaw(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_yaw(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposePitch wraps MEOS C function tpose_pitch.
func TposePitch(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_pitch(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeRoll wraps MEOS C function tpose_roll.
func TposeRoll(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_roll(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeSpeed wraps MEOS C function tpose_speed.
func TposeSpeed(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_speed(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeAngularSpeed wraps MEOS C function tpose_angular_speed.
func TposeAngularSpeed(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_angular_speed(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeStartValue wraps MEOS C function tpose_start_value.
func TposeStartValue(temp *Temporal) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_start_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// TposeTrajectory wraps MEOS C function tpose_trajectory.
func TposeTrajectory(temp *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_trajectory(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TposeValueAtTimestamptz wraps MEOS C function tpose_value_at_timestamptz.
func TposeValueAtTimestamptz(temp *Temporal, t int64, strict bool) (_r0 bool, _r1 *Pose, _err error) {
	var _out_result *C.Pose
	C.meos_errno_reset()
	_cret := C.tpose_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Pose{_inner: _out_result}, nil
}


// TposeValueN wraps MEOS C function tpose_value_n.
func TposeValueN(temp *Temporal, n int) (_r0 bool, _r1 *Pose, _err error) {
	var _out_result *C.Pose
	C.meos_errno_reset()
	_cret := C.tpose_value_n(temp._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Pose{_inner: _out_result}, nil
}


// TposeValues wraps MEOS C function tpose_values.
func TposeValues(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TposeAtGeom wraps MEOS C function tpose_at_geom.
func TposeAtGeom(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_at_geom(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeAtSTBOX wraps MEOS C function tpose_at_stbox.
func TposeAtSTBOX(temp *Temporal, box *STBox, border_inc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_at_stbox(temp._inner, box._inner, C.bool(border_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeAtPose wraps MEOS C function tpose_at_pose.
func TposeAtPose(temp *Temporal, pose *Pose) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_at_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeMinusGeom wraps MEOS C function tpose_minus_geom.
func TposeMinusGeom(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_minus_geom(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeMinusPose wraps MEOS C function tpose_minus_pose.
func TposeMinusPose(temp *Temporal, pose *Pose) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_minus_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeMinusSTBOX wraps MEOS C function tpose_minus_stbox.
func TposeMinusSTBOX(temp *Temporal, box *STBox, border_inc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_minus_stbox(temp._inner, box._inner, C.bool(border_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTposePose wraps MEOS C function tdistance_tpose_pose.
func TdistanceTposePose(temp *Temporal, pose *Pose) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTposeGeo wraps MEOS C function tdistance_tpose_geo.
func TdistanceTposeGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_tpose_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTposeTpose wraps MEOS C function tdistance_tpose_tpose.
func TdistanceTposeTpose(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_tpose_tpose(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// NadTposeGeo wraps MEOS C function nad_tpose_geo.
func NadTposeGeo(temp *Temporal, gs *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tpose_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTposePose wraps MEOS C function nad_tpose_pose.
func NadTposePose(temp *Temporal, pose *Pose) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTposeSTBOX wraps MEOS C function nad_tpose_stbox.
func NadTposeSTBOX(temp *Temporal, box *STBox) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tpose_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTposeTpose wraps MEOS C function nad_tpose_tpose.
func NadTposeTpose(temp1 *Temporal, temp2 *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tpose_tpose(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NaiTposeGeo wraps MEOS C function nai_tpose_geo.
func NaiTposeGeo(temp *Temporal, gs *Geom) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.nai_tpose_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// NaiTposePose wraps MEOS C function nai_tpose_pose.
func NaiTposePose(temp *Temporal, pose *Pose) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.nai_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// NaiTposeTpose wraps MEOS C function nai_tpose_tpose.
func NaiTposeTpose(temp1 *Temporal, temp2 *Temporal) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.nai_tpose_tpose(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// ShortestlineTposeGeo wraps MEOS C function shortestline_tpose_geo.
func ShortestlineTposeGeo(temp *Temporal, gs *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.shortestline_tpose_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// ShortestlineTposePose wraps MEOS C function shortestline_tpose_pose.
func ShortestlineTposePose(temp *Temporal, pose *Pose) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.shortestline_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// ShortestlineTposeTpose wraps MEOS C function shortestline_tpose_tpose.
func ShortestlineTposeTpose(temp1 *Temporal, temp2 *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.shortestline_tpose_tpose(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// AlwaysEqPoseTpose wraps MEOS C function always_eq_pose_tpose.
func AlwaysEqPoseTpose(pose *Pose, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_pose_tpose(pose._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTposePose wraps MEOS C function always_eq_tpose_pose.
func AlwaysEqTposePose(temp *Temporal, pose *Pose) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTposeTpose wraps MEOS C function always_eq_tpose_tpose.
func AlwaysEqTposeTpose(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tpose_tpose(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNePoseTpose wraps MEOS C function always_ne_pose_tpose.
func AlwaysNePoseTpose(pose *Pose, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_pose_tpose(pose._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTposePose wraps MEOS C function always_ne_tpose_pose.
func AlwaysNeTposePose(temp *Temporal, pose *Pose) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTposeTpose wraps MEOS C function always_ne_tpose_tpose.
func AlwaysNeTposeTpose(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tpose_tpose(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqPoseTpose wraps MEOS C function ever_eq_pose_tpose.
func EverEqPoseTpose(pose *Pose, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_pose_tpose(pose._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTposePose wraps MEOS C function ever_eq_tpose_pose.
func EverEqTposePose(temp *Temporal, pose *Pose) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTposeTpose wraps MEOS C function ever_eq_tpose_tpose.
func EverEqTposeTpose(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tpose_tpose(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNePoseTpose wraps MEOS C function ever_ne_pose_tpose.
func EverNePoseTpose(pose *Pose, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_pose_tpose(pose._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTposePose wraps MEOS C function ever_ne_tpose_pose.
func EverNeTposePose(temp *Temporal, pose *Pose) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTposeTpose wraps MEOS C function ever_ne_tpose_tpose.
func EverNeTposeTpose(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tpose_tpose(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TeqPoseTpose wraps MEOS C function teq_pose_tpose.
func TeqPoseTpose(pose *Pose, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_pose_tpose(pose._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTposePose wraps MEOS C function teq_tpose_pose.
func TeqTposePose(temp *Temporal, pose *Pose) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnePoseTpose wraps MEOS C function tne_pose_tpose.
func TnePoseTpose(pose *Pose, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_pose_tpose(pose._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTposePose wraps MEOS C function tne_tpose_pose.
func TneTposePose(temp *Temporal, pose *Pose) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}

