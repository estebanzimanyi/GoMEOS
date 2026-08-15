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

// ProjGetContext wraps MEOS C function proj_get_context.
func ProjGetContext() *PJContext {
	_cret := C.proj_get_context()
	return &PJContext{_inner: _cret}
}


// PointRound wraps MEOS C function point_round.
func PointRound(gs *Geom, maxdd int) *Geom {
	_cret := C.point_round(gs._inner, C.int(maxdd))
	return &Geom{_inner: _cret}
}


// GeoSetSRIDInt wraps MEOS C function geo_set_srid_int.
func GeoSetSRIDInt(gs *Geom, srid int32) {
	C.geo_set_srid_int(gs._inner, C.int32_t(srid))
}


// STBOXSet wraps MEOS C function stbox_set.
func STBOXSet(hasx bool, hasz bool, geodetic bool, srid int32, xmin float64, xmax float64, ymin float64, ymax float64, zmin float64, zmax float64, s *Span) *STBox {
	var _out_result C.STBox
	C.stbox_set(C.bool(hasx), C.bool(hasz), C.bool(geodetic), C.int32(srid), C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.double(zmin), C.double(zmax), s._inner, &_out_result)
	return &STBox{_inner: &_out_result}
}


// GboxSetSTBOX wraps MEOS C function gbox_set_stbox.
func GboxSetSTBOX(box *GBox, srid int32) *STBox {
	var _out_result C.STBox
	C.gbox_set_stbox(box._inner, C.int32_t(srid), &_out_result)
	return &STBox{_inner: &_out_result}
}


// GeoSetSTBOX wraps MEOS C function geo_set_stbox.
func GeoSetSTBOX(gs *Geom) (bool, *STBox) {
	var _out_result C.STBox
	_cret := C.geo_set_stbox(gs._inner, &_out_result)
	return bool(_cret), &STBox{_inner: &_out_result}
}


// SpatialsetSetSTBOX wraps MEOS C function spatialset_set_stbox.
func SpatialsetSetSTBOX(set *Set) *STBox {
	var _out_result C.STBox
	C.spatialset_set_stbox(set._inner, &_out_result)
	return &STBox{_inner: &_out_result}
}


// STBOXSetBox3d wraps MEOS C function stbox_set_box3d.
func STBOXSetBox3d(box *STBox) *Box3D {
	var _out_result C.BOX3D
	C.stbox_set_box3d(box._inner, &_out_result)
	return &Box3D{_inner: &_out_result}
}


// STBOXSetGbox wraps MEOS C function stbox_set_gbox.
func STBOXSetGbox(box *STBox) *GBox {
	var _out_result C.GBOX
	C.stbox_set_gbox(box._inner, &_out_result)
	return &GBox{_inner: &_out_result}
}


// TstzsetSetSTBOX wraps MEOS C function tstzset_set_stbox.
func TstzsetSetSTBOX(s *Set) *STBox {
	var _out_result C.STBox
	C.tstzset_set_stbox(s._inner, &_out_result)
	return &STBox{_inner: &_out_result}
}


// TimestamptzSetSTBOX wraps MEOS C function timestamptz_set_stbox.
func TimestamptzSetSTBOX(t int64) *STBox {
	var _out_result C.STBox
	C.timestamptz_set_stbox(C.TimestampTz(t), &_out_result)
	return &STBox{_inner: &_out_result}
}


// TstzspanSetSTBOX wraps MEOS C function tstzspan_set_stbox.
func TstzspanSetSTBOX(s *Span) *STBox {
	var _out_result C.STBox
	C.tstzspan_set_stbox(s._inner, &_out_result)
	return &STBox{_inner: &_out_result}
}


// TstzspansetSetSTBOX wraps MEOS C function tstzspanset_set_stbox.
func TstzspansetSetSTBOX(s *SpanSet) *STBox {
	var _out_result C.STBox
	C.tstzspanset_set_stbox(s._inner, &_out_result)
	return &STBox{_inner: &_out_result}
}


// STBOXExpand wraps MEOS C function stbox_expand.
func STBOXExpand(box1 *STBox, box2 *STBox) {
	C.stbox_expand(box1._inner, box2._inner)
}


// STBOXExpandSpaceSet wraps MEOS C function stbox_expand_space_set.
func STBOXExpandSpaceSet(box *STBox, d float64) (bool, *STBox) {
	var _out_result C.STBox
	_cret := C.stbox_expand_space_set(box._inner, C.double(d), &_out_result)
	return bool(_cret), &STBox{_inner: &_out_result}
}


// InterSTBOXSTBOX wraps MEOS C function inter_stbox_stbox.
func InterSTBOXSTBOX(box1 *STBox, box2 *STBox) (bool, *STBox) {
	var _out_result C.STBox
	_cret := C.inter_stbox_stbox(box1._inner, box2._inner, &_out_result)
	return bool(_cret), &STBox{_inner: &_out_result}
}


// TgeogpointinstIn wraps MEOS C function tgeogpointinst_in.
func TgeogpointinstIn(str string) *TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeogpointinst_in(_c_str)
	return &TInstant{_inner: _cret}
}


// TgeogpointseqIn wraps MEOS C function tgeogpointseq_in.
func TgeogpointseqIn(str string, interp Interpolation) *TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeogpointseq_in(_c_str, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TgeogpointseqsetIn wraps MEOS C function tgeogpointseqset_in.
func TgeogpointseqsetIn(str string) *TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeogpointseqset_in(_c_str)
	return &TSequenceSet{_inner: _cret}
}


// TgeompointinstIn wraps MEOS C function tgeompointinst_in.
func TgeompointinstIn(str string) *TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeompointinst_in(_c_str)
	return &TInstant{_inner: _cret}
}


// TgeompointseqIn wraps MEOS C function tgeompointseq_in.
func TgeompointseqIn(str string, interp Interpolation) *TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeompointseq_in(_c_str, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TgeompointseqsetIn wraps MEOS C function tgeompointseqset_in.
func TgeompointseqsetIn(str string) *TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeompointseqset_in(_c_str)
	return &TSequenceSet{_inner: _cret}
}


// TgeographyinstIn wraps MEOS C function tgeographyinst_in.
func TgeographyinstIn(str string) *TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeographyinst_in(_c_str)
	return &TInstant{_inner: _cret}
}


// TgeographyseqIn wraps MEOS C function tgeographyseq_in.
func TgeographyseqIn(str string, interp Interpolation) *TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeographyseq_in(_c_str, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TgeographyseqsetIn wraps MEOS C function tgeographyseqset_in.
func TgeographyseqsetIn(str string) *TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeographyseqset_in(_c_str)
	return &TSequenceSet{_inner: _cret}
}


// TgeometryinstIn wraps MEOS C function tgeometryinst_in.
func TgeometryinstIn(str string) *TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeometryinst_in(_c_str)
	return &TInstant{_inner: _cret}
}


// TgeometryseqIn wraps MEOS C function tgeometryseq_in.
func TgeometryseqIn(str string, interp Interpolation) *TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeometryseq_in(_c_str, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TgeometryseqsetIn wraps MEOS C function tgeometryseqset_in.
func TgeometryseqsetIn(str string) *TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tgeometryseqset_in(_c_str)
	return &TSequenceSet{_inner: _cret}
}


// TspatialSetSTBOX wraps MEOS C function tspatial_set_stbox.
func TspatialSetSTBOX(temp *Temporal) *STBox {
	var _out_result C.STBox
	C.tspatial_set_stbox(temp._inner, &_out_result)
	return &STBox{_inner: &_out_result}
}


// TspatialseqSetSTBOX wraps MEOS C function tspatialseq_set_stbox.
func TspatialseqSetSTBOX(seq *TSequence, box *STBox) {
	C.tspatialseq_set_stbox(seq._inner, box._inner)
}


// TspatialseqsetSetSTBOX wraps MEOS C function tspatialseqset_set_stbox.
func TspatialseqsetSetSTBOX(ss *TSequenceSet, box *STBox) {
	C.tspatialseqset_set_stbox(ss._inner, box._inner)
}


// TgeoRestrictElevation wraps MEOS C function tgeo_restrict_elevation.
func TgeoRestrictElevation(temp *Temporal, s *Span, atfunc bool) *Temporal {
	_cret := C.tgeo_restrict_elevation(temp._inner, s._inner, C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TgeoRestrictGeom wraps MEOS C function tgeo_restrict_geom.
func TgeoRestrictGeom(temp *Temporal, gs *Geom, atfunc bool) *Temporal {
	_cret := C.tgeo_restrict_geom(temp._inner, gs._inner, C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TgeoRestrictSTBOX wraps MEOS C function tgeo_restrict_stbox.
func TgeoRestrictSTBOX(temp *Temporal, box *STBox, border_inc bool, atfunc bool) *Temporal {
	_cret := C.tgeo_restrict_stbox(temp._inner, box._inner, C.bool(border_inc), C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TgeoinstRestrictGeom wraps MEOS C function tgeoinst_restrict_geom.
func TgeoinstRestrictGeom(inst *TInstant, gs *Geom, atfunc bool) *TInstant {
	_cret := C.tgeoinst_restrict_geom(inst._inner, gs._inner, C.bool(atfunc))
	return &TInstant{_inner: _cret}
}


// TgeoinstRestrictSTBOX wraps MEOS C function tgeoinst_restrict_stbox.
func TgeoinstRestrictSTBOX(inst *TInstant, box *STBox, border_inc bool, atfunc bool) *TInstant {
	_cret := C.tgeoinst_restrict_stbox(inst._inner, box._inner, C.bool(border_inc), C.bool(atfunc))
	return &TInstant{_inner: _cret}
}


// TgeoseqRestrictGeom wraps MEOS C function tgeoseq_restrict_geom.
func TgeoseqRestrictGeom(seq *TSequence, gs *Geom, atfunc bool) *Temporal {
	_cret := C.tgeoseq_restrict_geom(seq._inner, gs._inner, C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TgeoseqRestrictSTBOX wraps MEOS C function tgeoseq_restrict_stbox.
func TgeoseqRestrictSTBOX(seq *TSequence, box *STBox, border_inc bool, atfunc bool) *Temporal {
	_cret := C.tgeoseq_restrict_stbox(seq._inner, box._inner, C.bool(border_inc), C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TgeoseqsetRestrictGeom wraps MEOS C function tgeoseqset_restrict_geom.
func TgeoseqsetRestrictGeom(ss *TSequenceSet, gs *Geom, atfunc bool) *TSequenceSet {
	_cret := C.tgeoseqset_restrict_geom(ss._inner, gs._inner, C.bool(atfunc))
	return &TSequenceSet{_inner: _cret}
}


// TgeoseqsetRestrictSTBOX wraps MEOS C function tgeoseqset_restrict_stbox.
func TgeoseqsetRestrictSTBOX(ss *TSequenceSet, box *STBox, border_inc bool, atfunc bool) *TSequenceSet {
	_cret := C.tgeoseqset_restrict_stbox(ss._inner, box._inner, C.bool(border_inc), C.bool(atfunc))
	return &TSequenceSet{_inner: _cret}
}


// GeoClipCtxMake wraps MEOS C function geo_clip_ctx_make.
func GeoClipCtxMake(gs *Geom) unsafe.Pointer {
	_cret := C.geo_clip_ctx_make(gs._inner)
	return unsafe.Pointer(_cret)
}


// GeoClipCtxFree wraps MEOS C function geo_clip_ctx_free.
func GeoClipCtxFree(ctx unsafe.Pointer) {
	C.geo_clip_ctx_free(unsafe.Pointer(ctx))
}


// GeoIntersects2d wraps MEOS C function geo_intersects2d.
func GeoIntersects2d(gs1 *Geom, gs2 *Geom) bool {
	_cret := C.geo_intersects2d(gs1._inner, gs2._inner)
	return bool(_cret)
}


// GeoIntersects2dCtx wraps MEOS C function geo_intersects2d_ctx.
func GeoIntersects2dCtx(gs *Geom, ctx unsafe.Pointer) bool {
	_cret := C.geo_intersects2d_ctx(gs._inner, unsafe.Pointer(ctx))
	return bool(_cret)
}


// GeoCovers2d wraps MEOS C function geo_covers2d.
func GeoCovers2d(gs1 *Geom, gs2 *Geom) bool {
	_cret := C.geo_covers2d(gs1._inner, gs2._inner)
	return bool(_cret)
}


// TpointLinearInterGeom wraps MEOS C function tpoint_linear_inter_geom.
func TpointLinearInterGeom(temp *Temporal, gs *Geom, clip bool) *Temporal {
	_cret := C.tpoint_linear_inter_geom(temp._inner, gs._inner, C.bool(clip))
	return &Temporal{_inner: _cret}
}


// TpointLinearInterGeomCtx wraps MEOS C function tpoint_linear_inter_geom_ctx.
func TpointLinearInterGeomCtx(temp *Temporal, ctx unsafe.Pointer, clip bool) *Temporal {
	_cret := C.tpoint_linear_inter_geom_ctx(temp._inner, unsafe.Pointer(ctx), C.bool(clip))
	return &Temporal{_inner: _cret}
}


// TpointLinearDwithinGeom wraps MEOS C function tpoint_linear_dwithin_geom.
func TpointLinearDwithinGeom(temp *Temporal, gs *Geom, dist float64) *Temporal {
	_cret := C.tpoint_linear_dwithin_geom(temp._inner, gs._inner, C.double(dist))
	return &Temporal{_inner: _cret}
}


// TpointLinearDwithinGeomCtx wraps MEOS C function tpoint_linear_dwithin_geom_ctx.
func TpointLinearDwithinGeomCtx(temp *Temporal, ctx unsafe.Pointer, dist float64) *Temporal {
	_cret := C.tpoint_linear_dwithin_geom_ctx(temp._inner, unsafe.Pointer(ctx), C.double(dist))
	return &Temporal{_inner: _cret}
}


// TpointLinearDistanceGeom wraps MEOS C function tpoint_linear_distance_geom.
func TpointLinearDistanceGeom(temp *Temporal, gs *Geom) *Temporal {
	_cret := C.tpoint_linear_distance_geom(temp._inner, gs._inner)
	return &Temporal{_inner: _cret}
}


// TpointLinearRestrictGeom wraps MEOS C function tpoint_linear_restrict_geom.
func TpointLinearRestrictGeom(temp *Temporal, gs *Geom, atfunc bool) *Temporal {
	_cret := C.tpoint_linear_restrict_geom(temp._inner, gs._inner, C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// GeomClipSupported wraps MEOS C function geom_clip_supported.
func GeomClipSupported(geom unsafe.Pointer) bool {
	_cret := C.geom_clip_supported((*C.LWGEOM)(unsafe.Pointer(geom)))
	return bool(_cret)
}


// TspatialinstSRID wraps MEOS C function tspatialinst_srid.
func TspatialinstSRID(inst *TInstant) int {
	_cret := C.tspatialinst_srid(inst._inner)
	return int(_cret)
}


// TpointseqAzimuth wraps MEOS C function tpointseq_azimuth.
func TpointseqAzimuth(seq *TSequence) *TSequenceSet {
	_cret := C.tpointseq_azimuth(seq._inner)
	return &TSequenceSet{_inner: _cret}
}


// TpointseqCumulativeLength wraps MEOS C function tpointseq_cumulative_length.
func TpointseqCumulativeLength(seq *TSequence, prevlength float64) *TSequence {
	_cret := C.tpointseq_cumulative_length(seq._inner, C.double(prevlength))
	return &TSequence{_inner: _cret}
}


// TpointseqIsSimple wraps MEOS C function tpointseq_is_simple.
func TpointseqIsSimple(seq *TSequence) bool {
	_cret := C.tpointseq_is_simple(seq._inner)
	return bool(_cret)
}


// TpointseqLength wraps MEOS C function tpointseq_length.
func TpointseqLength(seq *TSequence) float64 {
	_cret := C.tpointseq_length(seq._inner)
	return float64(_cret)
}


// TpointseqLinearTrajectory wraps MEOS C function tpointseq_linear_trajectory.
func TpointseqLinearTrajectory(seq *TSequence, unary_union bool) *Geom {
	_cret := C.tpointseq_linear_trajectory(seq._inner, C.bool(unary_union))
	return &Geom{_inner: _cret}
}


// TgeoseqStboxes wraps MEOS C function tgeoseq_stboxes.
func TgeoseqStboxes(seq *TSequence, count unsafe.Pointer) *STBox {
	_cret := C.tgeoseq_stboxes(seq._inner, (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// TgeoseqSplitNStboxes wraps MEOS C function tgeoseq_split_n_stboxes.
func TgeoseqSplitNStboxes(seq *TSequence, max_count int, count unsafe.Pointer) *STBox {
	_cret := C.tgeoseq_split_n_stboxes(seq._inner, C.int(max_count), (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// TpointseqsetAzimuth wraps MEOS C function tpointseqset_azimuth.
func TpointseqsetAzimuth(ss *TSequenceSet) *TSequenceSet {
	_cret := C.tpointseqset_azimuth(ss._inner)
	return &TSequenceSet{_inner: _cret}
}


// TpointseqsetCumulativeLength wraps MEOS C function tpointseqset_cumulative_length.
func TpointseqsetCumulativeLength(ss *TSequenceSet) *TSequenceSet {
	_cret := C.tpointseqset_cumulative_length(ss._inner)
	return &TSequenceSet{_inner: _cret}
}


// TpointseqsetIsSimple wraps MEOS C function tpointseqset_is_simple.
func TpointseqsetIsSimple(ss *TSequenceSet) bool {
	_cret := C.tpointseqset_is_simple(ss._inner)
	return bool(_cret)
}


// TpointseqsetLength wraps MEOS C function tpointseqset_length.
func TpointseqsetLength(ss *TSequenceSet) float64 {
	_cret := C.tpointseqset_length(ss._inner)
	return float64(_cret)
}


// TgeoseqsetStboxes wraps MEOS C function tgeoseqset_stboxes.
func TgeoseqsetStboxes(ss *TSequenceSet, count unsafe.Pointer) *STBox {
	_cret := C.tgeoseqset_stboxes(ss._inner, (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// TgeoseqsetSplitNStboxes wraps MEOS C function tgeoseqset_split_n_stboxes.
func TgeoseqsetSplitNStboxes(ss *TSequenceSet, max_count int, count unsafe.Pointer) *STBox {
	_cret := C.tgeoseqset_split_n_stboxes(ss._inner, C.int(max_count), (*C.int)(unsafe.Pointer(count)))
	return &STBox{_inner: _cret}
}


// TgeominstTgeoginst wraps MEOS C function tgeominst_tgeoginst.
func TgeominstTgeoginst(inst *TInstant, oper bool) *TInstant {
	_cret := C.tgeominst_tgeoginst(inst._inner, C.bool(oper))
	return &TInstant{_inner: _cret}
}


// TgeomseqTgeogseq wraps MEOS C function tgeomseq_tgeogseq.
func TgeomseqTgeogseq(seq *TSequence, oper bool) *TSequence {
	_cret := C.tgeomseq_tgeogseq(seq._inner, C.bool(oper))
	return &TSequence{_inner: _cret}
}


// TgeomseqsetTgeogseqset wraps MEOS C function tgeomseqset_tgeogseqset.
func TgeomseqsetTgeogseqset(ss *TSequenceSet, oper bool) *TSequenceSet {
	_cret := C.tgeomseqset_tgeogseqset(ss._inner, C.bool(oper))
	return &TSequenceSet{_inner: _cret}
}


// TgeomTgeog wraps MEOS C function tgeom_tgeog.
func TgeomTgeog(temp *Temporal, oper bool) *Temporal {
	_cret := C.tgeom_tgeog(temp._inner, C.bool(oper))
	return &Temporal{_inner: _cret}
}


// TgeoTpoint wraps MEOS C function tgeo_tpoint.
func TgeoTpoint(temp *Temporal, oper bool) *Temporal {
	_cret := C.tgeo_tpoint(temp._inner, C.bool(oper))
	return &Temporal{_inner: _cret}
}


// TspatialinstSetSRID wraps MEOS C function tspatialinst_set_srid.
func TspatialinstSetSRID(inst *TInstant, srid int32) {
	C.tspatialinst_set_srid(inst._inner, C.int32_t(srid))
}


// TpointseqMakeSimple wraps MEOS C function tpointseq_make_simple.
func TpointseqMakeSimple(seq *TSequence, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tpointseq_make_simple(seq._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TspatialseqSetSRID wraps MEOS C function tspatialseq_set_srid.
func TspatialseqSetSRID(seq *TSequence, srid int32) {
	C.tspatialseq_set_srid(seq._inner, C.int32_t(srid))
}


// TpointseqsetMakeSimple wraps MEOS C function tpointseqset_make_simple.
func TpointseqsetMakeSimple(ss *TSequenceSet, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tpointseqset_make_simple(ss._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TspatialseqsetSetSRID wraps MEOS C function tspatialseqset_set_srid.
func TspatialseqsetSetSRID(ss *TSequenceSet, srid int32) {
	C.tspatialseqset_set_srid(ss._inner, C.int32_t(srid))
}


// TpointseqTwcentroid wraps MEOS C function tpointseq_twcentroid.
func TpointseqTwcentroid(seq *TSequence) *Geom {
	_cret := C.tpointseq_twcentroid(seq._inner)
	return &Geom{_inner: _cret}
}


// TpointseqsetTwcentroid wraps MEOS C function tpointseqset_twcentroid.
func TpointseqsetTwcentroid(ss *TSequenceSet) *Geom {
	_cret := C.tpointseqset_twcentroid(ss._inner)
	return &Geom{_inner: _cret}
}

