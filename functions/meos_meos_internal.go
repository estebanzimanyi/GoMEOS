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

// GslGetGenerationRng wraps MEOS C function gsl_get_generation_rng.
func GslGetGenerationRng() *GslRng {
	_cret := C.gsl_get_generation_rng()
	return &GslRng{_inner: _cret}
}


// GslGetAggregationRng wraps MEOS C function gsl_get_aggregation_rng.
func GslGetAggregationRng() *GslRng {
	_cret := C.gsl_get_aggregation_rng()
	return &GslRng{_inner: _cret}
}


// FloatspanRoundSet wraps MEOS C function floatspan_round_set.
func FloatspanRoundSet(s *Span, maxdd int) *Span {
	var _out_result C.Span
	C.floatspan_round_set(s._inner, C.int(maxdd), &_out_result)
	return &Span{_inner: &_out_result}
}


// SetIn wraps MEOS C function set_in.
func SetIn(str string, basetype MeosType) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.set_in(_c_str, C.MeosType(basetype))
	return &Set{_inner: _cret}
}


// SetOut wraps MEOS C function set_out.
func SetOut(s *Set, maxdd int) string {
	_cret := C.set_out(s._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// SpanIn wraps MEOS C function span_in.
func SpanIn(str string, spantype MeosType) *Span {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.span_in(_c_str, C.MeosType(spantype))
	return &Span{_inner: _cret}
}


// SpanOut wraps MEOS C function span_out.
func SpanOut(s *Span, maxdd int) string {
	_cret := C.span_out(s._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// SpansetIn wraps MEOS C function spanset_in.
func SpansetIn(str string, spantype MeosType) *SpanSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.spanset_in(_c_str, C.MeosType(spantype))
	return &SpanSet{_inner: _cret}
}


// SpansetOut wraps MEOS C function spanset_out.
func SpansetOut(ss *SpanSet, maxdd int) string {
	_cret := C.spanset_out(ss._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// SpansetMakeExp wraps MEOS C function spanset_make_exp.
func SpansetMakeExp(spans *Span, count int, maxcount int, normalize bool, order bool) *SpanSet {
	_cret := C.spanset_make_exp(spans._inner, C.int(count), C.int(maxcount), C.bool(normalize), C.bool(order))
	return &SpanSet{_inner: _cret}
}


// SpansetMakeFree wraps MEOS C function spanset_make_free.
func SpansetMakeFree(spans *Span, count int, normalize bool, order bool) *SpanSet {
	_cret := C.spanset_make_free(spans._inner, C.int(count), C.bool(normalize), C.bool(order))
	return &SpanSet{_inner: _cret}
}


// SetSpan wraps MEOS C function set_span.
func SetSpan(s *Set) *Span {
	_cret := C.set_span(s._inner)
	return &Span{_inner: _cret}
}


// SetSpanset wraps MEOS C function set_spanset.
func SetSpanset(s *Set) *SpanSet {
	_cret := C.set_spanset(s._inner)
	return &SpanSet{_inner: _cret}
}


// SetMemSize wraps MEOS C function set_mem_size.
func SetMemSize(s *Set) int {
	_cret := C.set_mem_size(s._inner)
	return int(_cret)
}


// SetSetSubspan wraps MEOS C function set_set_subspan.
func SetSetSubspan(s *Set, minidx int, maxidx int) *Span {
	var _out_result C.Span
	C.set_set_subspan(s._inner, C.int(minidx), C.int(maxidx), &_out_result)
	return &Span{_inner: &_out_result}
}


// SetSetSpan wraps MEOS C function set_set_span.
func SetSetSpan(s *Set) *Span {
	var _out_result C.Span
	C.set_set_span(s._inner, &_out_result)
	return &Span{_inner: &_out_result}
}


// SpansetMemSize wraps MEOS C function spanset_mem_size.
func SpansetMemSize(ss *SpanSet) int {
	_cret := C.spanset_mem_size(ss._inner)
	return int(_cret)
}


// SpansetSps wraps MEOS C function spanset_sps.
func SpansetSps(ss *SpanSet, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.spanset_sps(ss._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// DatespanSetTstzspan wraps MEOS C function datespan_set_tstzspan.
func DatespanSetTstzspan(s1 *Span, s2 *Span) {
	C.datespan_set_tstzspan(s1._inner, s2._inner)
}


// BigintspanSetFloatspan wraps MEOS C function bigintspan_set_floatspan.
func BigintspanSetFloatspan(s1 *Span, s2 *Span) {
	C.bigintspan_set_floatspan(s1._inner, s2._inner)
}


// BigintspanSetIntspan wraps MEOS C function bigintspan_set_intspan.
func BigintspanSetIntspan(s1 *Span, s2 *Span) {
	C.bigintspan_set_intspan(s1._inner, s2._inner)
}


// FloatspanSetBigintspan wraps MEOS C function floatspan_set_bigintspan.
func FloatspanSetBigintspan(s1 *Span, s2 *Span) {
	C.floatspan_set_bigintspan(s1._inner, s2._inner)
}


// FloatspanSetIntspan wraps MEOS C function floatspan_set_intspan.
func FloatspanSetIntspan(s1 *Span, s2 *Span) {
	C.floatspan_set_intspan(s1._inner, s2._inner)
}


// IntspanSetBigintspan wraps MEOS C function intspan_set_bigintspan.
func IntspanSetBigintspan(s1 *Span, s2 *Span) {
	C.intspan_set_bigintspan(s1._inner, s2._inner)
}


// IntspanSetFloatspan wraps MEOS C function intspan_set_floatspan.
func IntspanSetFloatspan(s1 *Span, s2 *Span) {
	C.intspan_set_floatspan(s1._inner, s2._inner)
}


// SetCompact wraps MEOS C function set_compact.
func SetCompact(s *Set) *Set {
	_cret := C.set_compact(s._inner)
	return &Set{_inner: _cret}
}


// SpanExpand wraps MEOS C function span_expand.
func SpanExpand(s1 *Span, s2 *Span) {
	C.span_expand(s1._inner, s2._inner)
}


// SuperUnionSpanSpan wraps MEOS C function super_union_span_span.
func SuperUnionSpanSpan(s1 *Span, s2 *Span) *Span {
	_cret := C.super_union_span_span(s1._inner, s2._inner)
	return &Span{_inner: _cret}
}


// SpansetCompact wraps MEOS C function spanset_compact.
func SpansetCompact(ss *SpanSet) *SpanSet {
	_cret := C.spanset_compact(ss._inner)
	return &SpanSet{_inner: _cret}
}


// TextcatTextsetTextCommon wraps MEOS C function textcat_textset_text_common.
func TextcatTextsetTextCommon(s *Set, txt string, invert bool) *Set {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.textcat_textset_text_common(s._inner, _c_txt, C.bool(invert))
	return &Set{_inner: _cret}
}


// TstzspanSetDatespan wraps MEOS C function tstzspan_set_datespan.
func TstzspanSetDatespan(s1 *Span, s2 *Span) {
	C.tstzspan_set_datespan(s1._inner, s2._inner)
}


// OvadjSpanSpan wraps MEOS C function ovadj_span_span.
func OvadjSpanSpan(s1 *Span, s2 *Span) bool {
	_cret := C.ovadj_span_span(s1._inner, s2._inner)
	return bool(_cret)
}


// LfnadjSpanSpan wraps MEOS C function lfnadj_span_span.
func LfnadjSpanSpan(s1 *Span, s2 *Span) bool {
	_cret := C.lfnadj_span_span(s1._inner, s2._inner)
	return bool(_cret)
}


// BboxType wraps MEOS C function bbox_type.
func BboxType(bboxtype MeosType) bool {
	_cret := C.bbox_type(C.MeosType(bboxtype))
	return bool(_cret)
}


// BboxGetSize wraps MEOS C function bbox_get_size.
func BboxGetSize(bboxtype MeosType) uint {
	_cret := C.bbox_get_size(C.MeosType(bboxtype))
	return uint(_cret)
}


// BboxMaxDims wraps MEOS C function bbox_max_dims.
func BboxMaxDims(bboxtype MeosType) int {
	_cret := C.bbox_max_dims(C.MeosType(bboxtype))
	return int(_cret)
}


// TemporalBboxEq wraps MEOS C function temporal_bbox_eq.
func TemporalBboxEq(box1 unsafe.Pointer, box2 unsafe.Pointer, temptype MeosType) bool {
	_cret := C.temporal_bbox_eq(unsafe.Pointer(box1), unsafe.Pointer(box2), C.MeosType(temptype))
	return bool(_cret)
}


// TemporalBboxCmp wraps MEOS C function temporal_bbox_cmp.
func TemporalBboxCmp(box1 unsafe.Pointer, box2 unsafe.Pointer, temptype MeosType) int {
	_cret := C.temporal_bbox_cmp(unsafe.Pointer(box1), unsafe.Pointer(box2), C.MeosType(temptype))
	return int(_cret)
}


// EnsureBboxTemporalCompatible wraps MEOS C function ensure_bbox_temporal_compatible.
func EnsureBboxTemporalCompatible(bboxtype MeosType, temp *Temporal) bool {
	_cret := C.ensure_bbox_temporal_compatible(C.MeosType(bboxtype), temp._inner)
	return bool(_cret)
}


// BboxTemporalSplitBoxes wraps MEOS C function bbox_temporal_split_boxes.
func BboxTemporalSplitBoxes(bboxtype MeosType, boxsize uint, temp *Temporal, maxboxes int, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.bbox_temporal_split_boxes(C.MeosType(bboxtype), C.size_t(boxsize), temp._inner, C.int(maxboxes), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// BboxUnionSpanSpan wraps MEOS C function bbox_union_span_span.
func BboxUnionSpanSpan(s1 *Span, s2 *Span) *Span {
	var _out_result C.Span
	C.bbox_union_span_span(s1._inner, s2._inner, &_out_result)
	return &Span{_inner: &_out_result}
}


// InterSpanSpan wraps MEOS C function inter_span_span.
func InterSpanSpan(s1 *Span, s2 *Span) (bool, *Span) {
	var _out_result C.Span
	_cret := C.inter_span_span(s1._inner, s2._inner, &_out_result)
	return bool(_cret), &Span{_inner: &_out_result}
}


// MiSpanSpan wraps MEOS C function mi_span_span.
func MiSpanSpan(s1 *Span, s2 *Span) (int, *Span) {
	var _out_result C.Span
	_cret := C.mi_span_span(s1._inner, s2._inner, &_out_result)
	return int(_cret), &Span{_inner: &_out_result}
}


// TBOXSet wraps MEOS C function tbox_set.
func TBOXSet(s *Span, p *Span, box *TBox) {
	C.tbox_set(s._inner, p._inner, box._inner)
}


// FloatSetTBOX wraps MEOS C function float_set_tbox.
func FloatSetTBOX(d float64, box *TBox) {
	C.float_set_tbox(C.double(d), box._inner)
}


// IntSetTBOX wraps MEOS C function int_set_tbox.
func IntSetTBOX(i int, box *TBox) {
	C.int_set_tbox(C.int(i), box._inner)
}


// NumsetSetTBOX wraps MEOS C function numset_set_tbox.
func NumsetSetTBOX(s *Set, box *TBox) {
	C.numset_set_tbox(s._inner, box._inner)
}


// NumspanSetTBOX wraps MEOS C function numspan_set_tbox.
func NumspanSetTBOX(span *Span, box *TBox) {
	C.numspan_set_tbox(span._inner, box._inner)
}


// TimestamptzSetTBOX wraps MEOS C function timestamptz_set_tbox.
func TimestamptzSetTBOX(t int64, box *TBox) {
	C.timestamptz_set_tbox(C.TimestampTz(t), box._inner)
}


// TstzsetSetTBOX wraps MEOS C function tstzset_set_tbox.
func TstzsetSetTBOX(s *Set, box *TBox) {
	C.tstzset_set_tbox(s._inner, box._inner)
}


// TstzspanSetTBOX wraps MEOS C function tstzspan_set_tbox.
func TstzspanSetTBOX(s *Span, box *TBox) {
	C.tstzspan_set_tbox(s._inner, box._inner)
}


// TBOXExpand wraps MEOS C function tbox_expand.
func TBOXExpand(box1 *TBox, box2 *TBox) {
	C.tbox_expand(box1._inner, box2._inner)
}


// InterTBOXTBOX wraps MEOS C function inter_tbox_tbox.
func InterTBOXTBOX(box1 *TBox, box2 *TBox) (bool, *TBox) {
	var _out_result C.TBox
	_cret := C.inter_tbox_tbox(box1._inner, box2._inner, &_out_result)
	return bool(_cret), &TBox{_inner: &_out_result}
}


// TboolinstIn wraps MEOS C function tboolinst_in.
func TboolinstIn(str string) *TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tboolinst_in(_c_str)
	return &TInstant{_inner: _cret}
}


// TboolseqIn wraps MEOS C function tboolseq_in.
func TboolseqIn(str string, interp Interpolation) *TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tboolseq_in(_c_str, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TboolseqsetIn wraps MEOS C function tboolseqset_in.
func TboolseqsetIn(str string) *TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tboolseqset_in(_c_str)
	return &TSequenceSet{_inner: _cret}
}


// TemporalIn wraps MEOS C function temporal_in.
func TemporalIn(str string, temptype MeosType) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.temporal_in(_c_str, C.MeosType(temptype))
	return &Temporal{_inner: _cret}
}


// TemporalOut wraps MEOS C function temporal_out.
func TemporalOut(temp *Temporal, maxdd int) string {
	_cret := C.temporal_out(temp._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// TemparrOut wraps MEOS C function temparr_out.
func TemparrOut(temparr unsafe.Pointer, count int, maxdd int) unsafe.Pointer {
	_cret := C.temparr_out((**C.Temporal)(unsafe.Pointer(temparr)), C.int(count), C.int(maxdd))
	return unsafe.Pointer(_cret)
}


// TfloatinstIn wraps MEOS C function tfloatinst_in.
func TfloatinstIn(str string) *TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tfloatinst_in(_c_str)
	return &TInstant{_inner: _cret}
}


// TfloatseqIn wraps MEOS C function tfloatseq_in.
func TfloatseqIn(str string, interp Interpolation) *TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tfloatseq_in(_c_str, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TfloatseqsetIn wraps MEOS C function tfloatseqset_in.
func TfloatseqsetIn(str string) *TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tfloatseqset_in(_c_str)
	return &TSequenceSet{_inner: _cret}
}


// TinstantIn wraps MEOS C function tinstant_in.
func TinstantIn(str string, temptype MeosType) *TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tinstant_in(_c_str, C.MeosType(temptype))
	return &TInstant{_inner: _cret}
}


// TinstantOut wraps MEOS C function tinstant_out.
func TinstantOut(inst *TInstant, maxdd int) string {
	_cret := C.tinstant_out(inst._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// TbigintinstIn wraps MEOS C function tbigintinst_in.
func TbigintinstIn(str string) *TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tbigintinst_in(_c_str)
	return &TInstant{_inner: _cret}
}


// TbigintseqsetIn wraps MEOS C function tbigintseqset_in.
func TbigintseqsetIn(str string) *TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tbigintseqset_in(_c_str)
	return &TSequenceSet{_inner: _cret}
}


// TintinstIn wraps MEOS C function tintinst_in.
func TintinstIn(str string) *TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tintinst_in(_c_str)
	return &TInstant{_inner: _cret}
}


// TintseqIn wraps MEOS C function tintseq_in.
func TintseqIn(str string, interp Interpolation) *TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tintseq_in(_c_str, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TintseqsetIn wraps MEOS C function tintseqset_in.
func TintseqsetIn(str string) *TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tintseqset_in(_c_str)
	return &TSequenceSet{_inner: _cret}
}


// TsequenceIn wraps MEOS C function tsequence_in.
func TsequenceIn(str string, temptype MeosType, interp Interpolation) *TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tsequence_in(_c_str, C.MeosType(temptype), C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TsequenceOut wraps MEOS C function tsequence_out.
func TsequenceOut(seq *TSequence, maxdd int) string {
	_cret := C.tsequence_out(seq._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// TsequencesetIn wraps MEOS C function tsequenceset_in.
func TsequencesetIn(str string, temptype MeosType, interp Interpolation) *TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tsequenceset_in(_c_str, C.MeosType(temptype), C.interpType(interp))
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetOut wraps MEOS C function tsequenceset_out.
func TsequencesetOut(ss *TSequenceSet, maxdd int) string {
	_cret := C.tsequenceset_out(ss._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// TtextinstIn wraps MEOS C function ttextinst_in.
func TtextinstIn(str string) *TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.ttextinst_in(_c_str)
	return &TInstant{_inner: _cret}
}


// TtextseqIn wraps MEOS C function ttextseq_in.
func TtextseqIn(str string, interp Interpolation) *TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.ttextseq_in(_c_str, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TtextseqsetIn wraps MEOS C function ttextseqset_in.
func TtextseqsetIn(str string) *TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.ttextseqset_in(_c_str)
	return &TSequenceSet{_inner: _cret}
}


// TemporalFromMFJSON wraps MEOS C function temporal_from_mfjson.
func TemporalFromMFJSON(mfjson string, temptype MeosType) *Temporal {
	_c_mfjson := C.CString(mfjson)
	defer C.free(unsafe.Pointer(_c_mfjson))
	_cret := C.temporal_from_mfjson(_c_mfjson, C.MeosType(temptype))
	return &Temporal{_inner: _cret}
}


// TinstantCopy wraps MEOS C function tinstant_copy.
func TinstantCopy(inst *TInstant) *TInstant {
	_cret := C.tinstant_copy(inst._inner)
	return &TInstant{_inner: _cret}
}


// TsequenceCopy wraps MEOS C function tsequence_copy.
func TsequenceCopy(seq *TSequence) *TSequence {
	_cret := C.tsequence_copy(seq._inner)
	return &TSequence{_inner: _cret}
}


// TsequenceMakeExp wraps MEOS C function tsequence_make_exp.
func TsequenceMakeExp(instants unsafe.Pointer, count int, maxcount int, lower_inc bool, upper_inc bool, interp Interpolation, normalize bool) *TSequence {
	_cret := C.tsequence_make_exp((**C.TInstant)(unsafe.Pointer(instants)), C.int(count), C.int(maxcount), C.bool(lower_inc), C.bool(upper_inc), C.interpType(interp), C.bool(normalize))
	return &TSequence{_inner: _cret}
}


// TsequenceMakeFree wraps MEOS C function tsequence_make_free.
func TsequenceMakeFree(instants unsafe.Pointer, count int, lower_inc bool, upper_inc bool, interp Interpolation, normalize bool) *TSequence {
	_cret := C.tsequence_make_free((**C.TInstant)(unsafe.Pointer(instants)), C.int(count), C.bool(lower_inc), C.bool(upper_inc), C.interpType(interp), C.bool(normalize))
	return &TSequence{_inner: _cret}
}


// TsequencesetCopy wraps MEOS C function tsequenceset_copy.
func TsequencesetCopy(ss *TSequenceSet) *TSequenceSet {
	_cret := C.tsequenceset_copy(ss._inner)
	return &TSequenceSet{_inner: _cret}
}


// TseqsetarrToTseqset wraps MEOS C function tseqsetarr_to_tseqset.
func TseqsetarrToTseqset(seqsets unsafe.Pointer, count int, totalseqs int) *TSequenceSet {
	_cret := C.tseqsetarr_to_tseqset((**C.TSequenceSet)(unsafe.Pointer(seqsets)), C.int(count), C.int(totalseqs))
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetMakeExp wraps MEOS C function tsequenceset_make_exp.
func TsequencesetMakeExp(sequences unsafe.Pointer, count int, maxcount int, normalize bool) *TSequenceSet {
	_cret := C.tsequenceset_make_exp((**C.TSequence)(unsafe.Pointer(sequences)), C.int(count), C.int(maxcount), C.bool(normalize))
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetMakeFree wraps MEOS C function tsequenceset_make_free.
func TsequencesetMakeFree(sequences unsafe.Pointer, count int, normalize bool) *TSequenceSet {
	_cret := C.tsequenceset_make_free((**C.TSequence)(unsafe.Pointer(sequences)), C.int(count), C.bool(normalize))
	return &TSequenceSet{_inner: _cret}
}


// TemporalSetTstzspan wraps MEOS C function temporal_set_tstzspan.
func TemporalSetTstzspan(temp *Temporal, s *Span) {
	C.temporal_set_tstzspan(temp._inner, s._inner)
}


// TemporalTimeOverlaps wraps MEOS C function temporal_time_overlaps.
func TemporalTimeOverlaps(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.temporal_time_overlaps(temp1._inner, temp2._inner)
	return bool(_cret)
}


// TinstantSetTstzspan wraps MEOS C function tinstant_set_tstzspan.
func TinstantSetTstzspan(inst *TInstant, s *Span) {
	C.tinstant_set_tstzspan(inst._inner, s._inner)
}


// TnumberSetTBOX wraps MEOS C function tnumber_set_tbox.
func TnumberSetTBOX(temp *Temporal, box *TBox) {
	C.tnumber_set_tbox(temp._inner, box._inner)
}


// TnumberinstSetTBOX wraps MEOS C function tnumberinst_set_tbox.
func TnumberinstSetTBOX(inst *TInstant, box *TBox) {
	C.tnumberinst_set_tbox(inst._inner, box._inner)
}


// TnumberseqSetTBOX wraps MEOS C function tnumberseq_set_tbox.
func TnumberseqSetTBOX(seq *TSequence, box *TBox) {
	C.tnumberseq_set_tbox(seq._inner, box._inner)
}


// TnumberseqsetSetTBOX wraps MEOS C function tnumberseqset_set_tbox.
func TnumberseqsetSetTBOX(ss *TSequenceSet, box *TBox) {
	C.tnumberseqset_set_tbox(ss._inner, box._inner)
}


// TsequenceSetTstzspan wraps MEOS C function tsequence_set_tstzspan.
func TsequenceSetTstzspan(seq *TSequence, s *Span) {
	C.tsequence_set_tstzspan(seq._inner, s._inner)
}


// TsequencesetSetTstzspan wraps MEOS C function tsequenceset_set_tstzspan.
func TsequencesetSetTstzspan(ss *TSequenceSet, s *Span) {
	C.tsequenceset_set_tstzspan(ss._inner, s._inner)
}


// TemporalEndInst wraps MEOS C function temporal_end_inst.
func TemporalEndInst(temp *Temporal) *TInstant {
	_cret := C.temporal_end_inst(temp._inner)
	return &TInstant{_inner: _cret}
}


// TemporalInstN wraps MEOS C function temporal_inst_n.
func TemporalInstN(temp *Temporal, n int) *TInstant {
	_cret := C.temporal_inst_n(temp._inner, C.int(n))
	return &TInstant{_inner: _cret}
}


// TemporalInstsP wraps MEOS C function temporal_insts_p.
func TemporalInstsP(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.temporal_insts_p(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TemporalMaxInstP wraps MEOS C function temporal_max_inst_p.
func TemporalMaxInstP(temp *Temporal) *TInstant {
	_cret := C.temporal_max_inst_p(temp._inner)
	return &TInstant{_inner: _cret}
}


// TemporalMemSize wraps MEOS C function temporal_mem_size.
func TemporalMemSize(temp *Temporal) uint {
	_cret := C.temporal_mem_size(temp._inner)
	return uint(_cret)
}


// TemporalMinInstP wraps MEOS C function temporal_min_inst_p.
func TemporalMinInstP(temp *Temporal) *TInstant {
	_cret := C.temporal_min_inst_p(temp._inner)
	return &TInstant{_inner: _cret}
}


// TemporalSequencesP wraps MEOS C function temporal_sequences_p.
func TemporalSequencesP(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.temporal_sequences_p(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TemporalSetBbox wraps MEOS C function temporal_set_bbox.
func TemporalSetBbox(temp *Temporal, box unsafe.Pointer) {
	C.temporal_set_bbox(temp._inner, unsafe.Pointer(box))
}


// TemporalStartInst wraps MEOS C function temporal_start_inst.
func TemporalStartInst(temp *Temporal) *TInstant {
	_cret := C.temporal_start_inst(temp._inner)
	return &TInstant{_inner: _cret}
}


// TinstantHash wraps MEOS C function tinstant_hash.
func TinstantHash(inst *TInstant) uint32 {
	_cret := C.tinstant_hash(inst._inner)
	return uint32(_cret)
}


// TinstantHashExtended wraps MEOS C function tinstant_hash_extended.
func TinstantHashExtended(inst *TInstant, seed uint64) uint64 {
	_cret := C.tinstant_hash_extended(inst._inner, C.uint64_t(seed))
	return uint64(_cret)
}


// TinstantInsts wraps MEOS C function tinstant_insts.
func TinstantInsts(inst *TInstant, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tinstant_insts(inst._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TinstantSetBbox wraps MEOS C function tinstant_set_bbox.
func TinstantSetBbox(inst *TInstant, box unsafe.Pointer) {
	C.tinstant_set_bbox(inst._inner, unsafe.Pointer(box))
}


// TinstantTime wraps MEOS C function tinstant_time.
func TinstantTime(inst *TInstant) *SpanSet {
	_cret := C.tinstant_time(inst._inner)
	return &SpanSet{_inner: _cret}
}


// TinstantTimestamps wraps MEOS C function tinstant_timestamps.
func TinstantTimestamps(inst *TInstant, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tinstant_timestamps(inst._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TnumberSetSpan wraps MEOS C function tnumber_set_span.
func TnumberSetSpan(temp *Temporal, s *Span) {
	C.tnumber_set_span(temp._inner, s._inner)
}


// TnumberinstValuespans wraps MEOS C function tnumberinst_valuespans.
func TnumberinstValuespans(inst *TInstant) *SpanSet {
	_cret := C.tnumberinst_valuespans(inst._inner)
	return &SpanSet{_inner: _cret}
}


// TnumberseqAvgVal wraps MEOS C function tnumberseq_avg_val.
func TnumberseqAvgVal(seq *TSequence) float64 {
	_cret := C.tnumberseq_avg_val(seq._inner)
	return float64(_cret)
}


// TnumberseqValuespans wraps MEOS C function tnumberseq_valuespans.
func TnumberseqValuespans(seq *TSequence) *SpanSet {
	_cret := C.tnumberseq_valuespans(seq._inner)
	return &SpanSet{_inner: _cret}
}


// TnumberseqsetAvgVal wraps MEOS C function tnumberseqset_avg_val.
func TnumberseqsetAvgVal(ss *TSequenceSet) float64 {
	_cret := C.tnumberseqset_avg_val(ss._inner)
	return float64(_cret)
}


// TnumberseqsetValuespans wraps MEOS C function tnumberseqset_valuespans.
func TnumberseqsetValuespans(ss *TSequenceSet) *SpanSet {
	_cret := C.tnumberseqset_valuespans(ss._inner)
	return &SpanSet{_inner: _cret}
}


// TsequenceDuration wraps MEOS C function tsequence_duration.
func TsequenceDuration(seq *TSequence) *Interval {
	_cret := C.tsequence_duration(seq._inner)
	return &Interval{_inner: _cret}
}


// TsequenceEndTimestamptz wraps MEOS C function tsequence_end_timestamptz.
func TsequenceEndTimestamptz(seq *TSequence) int64 {
	_cret := C.tsequence_end_timestamptz(seq._inner)
	return int64(_cret)
}


// TsequenceHash wraps MEOS C function tsequence_hash.
func TsequenceHash(seq *TSequence) uint32 {
	_cret := C.tsequence_hash(seq._inner)
	return uint32(_cret)
}


// TsequenceHashExtended wraps MEOS C function tsequence_hash_extended.
func TsequenceHashExtended(seq *TSequence, seed uint64) uint64 {
	_cret := C.tsequence_hash_extended(seq._inner, C.uint64_t(seed))
	return uint64(_cret)
}


// TsequenceInstsP wraps MEOS C function tsequence_insts_p.
func TsequenceInstsP(seq *TSequence, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tsequence_insts_p(seq._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TsequenceMaxInstP wraps MEOS C function tsequence_max_inst_p.
func TsequenceMaxInstP(seq *TSequence) *TInstant {
	_cret := C.tsequence_max_inst_p(seq._inner)
	return &TInstant{_inner: _cret}
}


// TsequenceMinInstP wraps MEOS C function tsequence_min_inst_p.
func TsequenceMinInstP(seq *TSequence) *TInstant {
	_cret := C.tsequence_min_inst_p(seq._inner)
	return &TInstant{_inner: _cret}
}


// TsequenceSegments wraps MEOS C function tsequence_segments.
func TsequenceSegments(seq *TSequence, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tsequence_segments(seq._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TsequenceSeqs wraps MEOS C function tsequence_seqs.
func TsequenceSeqs(seq *TSequence, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tsequence_seqs(seq._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TsequenceStartTimestamptz wraps MEOS C function tsequence_start_timestamptz.
func TsequenceStartTimestamptz(seq *TSequence) int64 {
	_cret := C.tsequence_start_timestamptz(seq._inner)
	return int64(_cret)
}


// TsequenceTime wraps MEOS C function tsequence_time.
func TsequenceTime(seq *TSequence) *SpanSet {
	_cret := C.tsequence_time(seq._inner)
	return &SpanSet{_inner: _cret}
}


// TsequenceTimestamps wraps MEOS C function tsequence_timestamps.
func TsequenceTimestamps(seq *TSequence, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tsequence_timestamps(seq._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TsequencesetDuration wraps MEOS C function tsequenceset_duration.
func TsequencesetDuration(ss *TSequenceSet, boundspan bool) *Interval {
	_cret := C.tsequenceset_duration(ss._inner, C.bool(boundspan))
	return &Interval{_inner: _cret}
}


// TsequencesetEndTimestamptz wraps MEOS C function tsequenceset_end_timestamptz.
func TsequencesetEndTimestamptz(ss *TSequenceSet) int64 {
	_cret := C.tsequenceset_end_timestamptz(ss._inner)
	return int64(_cret)
}


// TsequencesetHash wraps MEOS C function tsequenceset_hash.
func TsequencesetHash(ss *TSequenceSet) uint32 {
	_cret := C.tsequenceset_hash(ss._inner)
	return uint32(_cret)
}


// TsequencesetHashExtended wraps MEOS C function tsequenceset_hash_extended.
func TsequencesetHashExtended(ss *TSequenceSet, seed uint64) uint64 {
	_cret := C.tsequenceset_hash_extended(ss._inner, C.uint64_t(seed))
	return uint64(_cret)
}


// TsequencesetInstN wraps MEOS C function tsequenceset_inst_n.
func TsequencesetInstN(ss *TSequenceSet, n int) *TInstant {
	_cret := C.tsequenceset_inst_n(ss._inner, C.int(n))
	return &TInstant{_inner: _cret}
}


// TsequencesetInstsP wraps MEOS C function tsequenceset_insts_p.
func TsequencesetInstsP(ss *TSequenceSet, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tsequenceset_insts_p(ss._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TsequencesetMaxInstP wraps MEOS C function tsequenceset_max_inst_p.
func TsequencesetMaxInstP(ss *TSequenceSet) *TInstant {
	_cret := C.tsequenceset_max_inst_p(ss._inner)
	return &TInstant{_inner: _cret}
}


// TsequencesetMinInstP wraps MEOS C function tsequenceset_min_inst_p.
func TsequencesetMinInstP(ss *TSequenceSet) *TInstant {
	_cret := C.tsequenceset_min_inst_p(ss._inner)
	return &TInstant{_inner: _cret}
}


// TsequencesetNumInstants wraps MEOS C function tsequenceset_num_instants.
func TsequencesetNumInstants(ss *TSequenceSet) int {
	_cret := C.tsequenceset_num_instants(ss._inner)
	return int(_cret)
}


// TsequencesetNumTimestamps wraps MEOS C function tsequenceset_num_timestamps.
func TsequencesetNumTimestamps(ss *TSequenceSet) int {
	_cret := C.tsequenceset_num_timestamps(ss._inner)
	return int(_cret)
}


// TsequencesetSegments wraps MEOS C function tsequenceset_segments.
func TsequencesetSegments(ss *TSequenceSet, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tsequenceset_segments(ss._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TsequencesetSequencesP wraps MEOS C function tsequenceset_sequences_p.
func TsequencesetSequencesP(ss *TSequenceSet) unsafe.Pointer {
	_cret := C.tsequenceset_sequences_p(ss._inner)
	return unsafe.Pointer(_cret)
}


// TsequencesetStartTimestamptz wraps MEOS C function tsequenceset_start_timestamptz.
func TsequencesetStartTimestamptz(ss *TSequenceSet) int64 {
	_cret := C.tsequenceset_start_timestamptz(ss._inner)
	return int64(_cret)
}


// TsequencesetTime wraps MEOS C function tsequenceset_time.
func TsequencesetTime(ss *TSequenceSet) *SpanSet {
	_cret := C.tsequenceset_time(ss._inner)
	return &SpanSet{_inner: _cret}
}


// TsequencesetTimestamptzN wraps MEOS C function tsequenceset_timestamptz_n.
func TsequencesetTimestamptzN(ss *TSequenceSet, n int) (bool, int64) {
	var _out_result C.TimestampTz
	_cret := C.tsequenceset_timestamptz_n(ss._inner, C.int(n), &_out_result)
	return bool(_cret), int64(_out_result)
}


// TsequencesetTimestamps wraps MEOS C function tsequenceset_timestamps.
func TsequencesetTimestamps(ss *TSequenceSet, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tsequenceset_timestamps(ss._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TemporalRestart wraps MEOS C function temporal_restart.
func TemporalRestart(temp *Temporal, count int) {
	C.temporal_restart(temp._inner, C.int(count))
}


// TemporalTsequence wraps MEOS C function temporal_tsequence.
func TemporalTsequence(temp *Temporal, interp Interpolation) *TSequence {
	_cret := C.temporal_tsequence(temp._inner, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TemporalTsequenceset wraps MEOS C function temporal_tsequenceset.
func TemporalTsequenceset(temp *Temporal, interp Interpolation) *TSequenceSet {
	_cret := C.temporal_tsequenceset(temp._inner, C.interpType(interp))
	return &TSequenceSet{_inner: _cret}
}


// TinstantShiftTime wraps MEOS C function tinstant_shift_time.
func TinstantShiftTime(inst *TInstant, interv *Interval) *TInstant {
	_cret := C.tinstant_shift_time(inst._inner, interv._inner)
	return &TInstant{_inner: _cret}
}


// TinstantAsTsequence wraps MEOS C function tinstant_as_tsequence.
func TinstantAsTsequence(inst *TInstant, interp Interpolation) *TSequence {
	_cret := C.tinstant_as_tsequence(inst._inner, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TinstantToTsequenceFree wraps MEOS C function tinstant_to_tsequence_free.
func TinstantToTsequenceFree(inst *TInstant, interp Interpolation) *TSequence {
	_cret := C.tinstant_to_tsequence_free(inst._inner, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TinstantAsTsequenceset wraps MEOS C function tinstant_as_tsequenceset.
func TinstantAsTsequenceset(inst *TInstant, interp Interpolation) *TSequenceSet {
	_cret := C.tinstant_as_tsequenceset(inst._inner, C.interpType(interp))
	return &TSequenceSet{_inner: _cret}
}


// TsequenceRestart wraps MEOS C function tsequence_restart.
func TsequenceRestart(seq *TSequence, count int) {
	C.tsequence_restart(seq._inner, C.int(count))
}


// TsequenceSetInterp wraps MEOS C function tsequence_set_interp.
func TsequenceSetInterp(seq *TSequence, interp Interpolation) *Temporal {
	_cret := C.tsequence_set_interp(seq._inner, C.interpType(interp))
	return &Temporal{_inner: _cret}
}


// TsequenceShiftScaleTime wraps MEOS C function tsequence_shift_scale_time.
func TsequenceShiftScaleTime(seq *TSequence, shift *Interval, duration *Interval) *TSequence {
	_cret := C.tsequence_shift_scale_time(seq._inner, shift._inner, duration._inner)
	return &TSequence{_inner: _cret}
}


// TsequenceSubseq wraps MEOS C function tsequence_subseq.
func TsequenceSubseq(seq *TSequence, from int, to int, lower_inc bool, upper_inc bool) *TSequence {
	_cret := C.tsequence_subseq(seq._inner, C.int(from), C.int(to), C.bool(lower_inc), C.bool(upper_inc))
	return &TSequence{_inner: _cret}
}


// TsequenceAsTinstant wraps MEOS C function tsequence_as_tinstant.
func TsequenceAsTinstant(seq *TSequence) *TInstant {
	_cret := C.tsequence_as_tinstant(seq._inner)
	return &TInstant{_inner: _cret}
}


// TsequenceAsTsequenceset wraps MEOS C function tsequence_as_tsequenceset.
func TsequenceAsTsequenceset(seq *TSequence) *TSequenceSet {
	_cret := C.tsequence_as_tsequenceset(seq._inner)
	return &TSequenceSet{_inner: _cret}
}


// TsequenceToTsequencesetFree wraps MEOS C function tsequence_to_tsequenceset_free.
func TsequenceToTsequencesetFree(seq *TSequence) *TSequenceSet {
	_cret := C.tsequence_to_tsequenceset_free(seq._inner)
	return &TSequenceSet{_inner: _cret}
}


// TsequenceToTsequencesetInterp wraps MEOS C function tsequence_to_tsequenceset_interp.
func TsequenceToTsequencesetInterp(seq *TSequence, interp Interpolation) *TSequenceSet {
	_cret := C.tsequence_to_tsequenceset_interp(seq._inner, C.interpType(interp))
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetRestart wraps MEOS C function tsequenceset_restart.
func TsequencesetRestart(ss *TSequenceSet, count int) {
	C.tsequenceset_restart(ss._inner, C.int(count))
}


// TsequencesetSetInterp wraps MEOS C function tsequenceset_set_interp.
func TsequencesetSetInterp(ss *TSequenceSet, interp Interpolation) *Temporal {
	_cret := C.tsequenceset_set_interp(ss._inner, C.interpType(interp))
	return &Temporal{_inner: _cret}
}


// TsequencesetShiftScaleTime wraps MEOS C function tsequenceset_shift_scale_time.
func TsequencesetShiftScaleTime(ss *TSequenceSet, start *Interval, duration *Interval) *TSequenceSet {
	_cret := C.tsequenceset_shift_scale_time(ss._inner, start._inner, duration._inner)
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetToDiscrete wraps MEOS C function tsequenceset_to_discrete.
func TsequencesetToDiscrete(ss *TSequenceSet) *TSequence {
	_cret := C.tsequenceset_to_discrete(ss._inner)
	return &TSequence{_inner: _cret}
}


// TsequencesetToLinear wraps MEOS C function tsequenceset_to_linear.
func TsequencesetToLinear(ss *TSequenceSet) *TSequenceSet {
	_cret := C.tsequenceset_to_linear(ss._inner)
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetToStep wraps MEOS C function tsequenceset_to_step.
func TsequencesetToStep(ss *TSequenceSet) *TSequenceSet {
	_cret := C.tsequenceset_to_step(ss._inner)
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetAsTinstant wraps MEOS C function tsequenceset_as_tinstant.
func TsequencesetAsTinstant(ss *TSequenceSet) *TInstant {
	_cret := C.tsequenceset_as_tinstant(ss._inner)
	return &TInstant{_inner: _cret}
}


// TsequencesetAsTsequence wraps MEOS C function tsequenceset_as_tsequence.
func TsequencesetAsTsequence(ss *TSequenceSet) *TSequence {
	_cret := C.tsequenceset_as_tsequence(ss._inner)
	return &TSequence{_inner: _cret}
}


// TinstantMerge wraps MEOS C function tinstant_merge.
func TinstantMerge(inst1 *TInstant, inst2 *TInstant) *Temporal {
	_cret := C.tinstant_merge(inst1._inner, inst2._inner)
	return &Temporal{_inner: _cret}
}


// TinstantMergeArray wraps MEOS C function tinstant_merge_array.
func TinstantMergeArray(instants unsafe.Pointer, count int) *Temporal {
	_cret := C.tinstant_merge_array((**C.TInstant)(unsafe.Pointer(instants)), C.int(count))
	return &Temporal{_inner: _cret}
}


// TsequenceAppendTinstant wraps MEOS C function tsequence_append_tinstant.
func TsequenceAppendTinstant(seq *TSequence, inst *TInstant, maxdist float64, maxt *Interval, expand bool) *Temporal {
	_cret := C.tsequence_append_tinstant(seq._inner, inst._inner, C.double(maxdist), maxt._inner, C.bool(expand))
	return &Temporal{_inner: _cret}
}


// TsequenceAppendTsequence wraps MEOS C function tsequence_append_tsequence.
func TsequenceAppendTsequence(seq1 *TSequence, seq2 *TSequence, expand bool) *Temporal {
	_cret := C.tsequence_append_tsequence(seq1._inner, seq2._inner, C.bool(expand))
	return &Temporal{_inner: _cret}
}


// TsequenceDeleteTimestamptz wraps MEOS C function tsequence_delete_timestamptz.
func TsequenceDeleteTimestamptz(seq *TSequence, t int64, connect bool) *Temporal {
	_cret := C.tsequence_delete_timestamptz(seq._inner, C.TimestampTz(t), C.bool(connect))
	return &Temporal{_inner: _cret}
}


// TsequenceDeleteTstzset wraps MEOS C function tsequence_delete_tstzset.
func TsequenceDeleteTstzset(seq *TSequence, s *Set, connect bool) *Temporal {
	_cret := C.tsequence_delete_tstzset(seq._inner, s._inner, C.bool(connect))
	return &Temporal{_inner: _cret}
}


// TsequenceDeleteTstzspan wraps MEOS C function tsequence_delete_tstzspan.
func TsequenceDeleteTstzspan(seq *TSequence, s *Span, connect bool) *Temporal {
	_cret := C.tsequence_delete_tstzspan(seq._inner, s._inner, C.bool(connect))
	return &Temporal{_inner: _cret}
}


// TsequenceDeleteTstzspanset wraps MEOS C function tsequence_delete_tstzspanset.
func TsequenceDeleteTstzspanset(seq *TSequence, ss *SpanSet, connect bool) *Temporal {
	_cret := C.tsequence_delete_tstzspanset(seq._inner, ss._inner, C.bool(connect))
	return &Temporal{_inner: _cret}
}


// TsequenceInsert wraps MEOS C function tsequence_insert.
func TsequenceInsert(seq1 *TSequence, seq2 *TSequence, connect bool) *Temporal {
	_cret := C.tsequence_insert(seq1._inner, seq2._inner, C.bool(connect))
	return &Temporal{_inner: _cret}
}


// TsequenceMerge wraps MEOS C function tsequence_merge.
func TsequenceMerge(seq1 *TSequence, seq2 *TSequence) *Temporal {
	_cret := C.tsequence_merge(seq1._inner, seq2._inner)
	return &Temporal{_inner: _cret}
}


// TsequenceMergeArray wraps MEOS C function tsequence_merge_array.
func TsequenceMergeArray(sequences unsafe.Pointer, count int) *Temporal {
	_cret := C.tsequence_merge_array((**C.TSequence)(unsafe.Pointer(sequences)), C.int(count))
	return &Temporal{_inner: _cret}
}


// TsequencesetAppendTinstant wraps MEOS C function tsequenceset_append_tinstant.
func TsequencesetAppendTinstant(ss *TSequenceSet, inst *TInstant, maxdist float64, maxt *Interval, expand bool) *TSequenceSet {
	_cret := C.tsequenceset_append_tinstant(ss._inner, inst._inner, C.double(maxdist), maxt._inner, C.bool(expand))
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetAppendTsequence wraps MEOS C function tsequenceset_append_tsequence.
func TsequencesetAppendTsequence(ss *TSequenceSet, seq *TSequence, expand bool) *TSequenceSet {
	_cret := C.tsequenceset_append_tsequence(ss._inner, seq._inner, C.bool(expand))
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetDeleteTimestamptz wraps MEOS C function tsequenceset_delete_timestamptz.
func TsequencesetDeleteTimestamptz(ss *TSequenceSet, t int64) *TSequenceSet {
	_cret := C.tsequenceset_delete_timestamptz(ss._inner, C.TimestampTz(t))
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetDeleteTstzset wraps MEOS C function tsequenceset_delete_tstzset.
func TsequencesetDeleteTstzset(ss *TSequenceSet, s *Set) *TSequenceSet {
	_cret := C.tsequenceset_delete_tstzset(ss._inner, s._inner)
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetDeleteTstzspan wraps MEOS C function tsequenceset_delete_tstzspan.
func TsequencesetDeleteTstzspan(ss *TSequenceSet, s *Span) *TSequenceSet {
	_cret := C.tsequenceset_delete_tstzspan(ss._inner, s._inner)
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetDeleteTstzspanset wraps MEOS C function tsequenceset_delete_tstzspanset.
func TsequencesetDeleteTstzspanset(ss *TSequenceSet, ps *SpanSet) *TSequenceSet {
	_cret := C.tsequenceset_delete_tstzspanset(ss._inner, ps._inner)
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetInsert wraps MEOS C function tsequenceset_insert.
func TsequencesetInsert(ss1 *TSequenceSet, ss2 *TSequenceSet) *TSequenceSet {
	_cret := C.tsequenceset_insert(ss1._inner, ss2._inner)
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetMerge wraps MEOS C function tsequenceset_merge.
func TsequencesetMerge(ss1 *TSequenceSet, ss2 *TSequenceSet) *TSequenceSet {
	_cret := C.tsequenceset_merge(ss1._inner, ss2._inner)
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetMergeArray wraps MEOS C function tsequenceset_merge_array.
func TsequencesetMergeArray(seqsets unsafe.Pointer, count int) *TSequenceSet {
	_cret := C.tsequenceset_merge_array((**C.TSequenceSet)(unsafe.Pointer(seqsets)), C.int(count))
	return &TSequenceSet{_inner: _cret}
}


// TsequenceExpandBbox wraps MEOS C function tsequence_expand_bbox.
func TsequenceExpandBbox(seq *TSequence, inst *TInstant) {
	C.tsequence_expand_bbox(seq._inner, inst._inner)
}


// TsequenceSetBbox wraps MEOS C function tsequence_set_bbox.
func TsequenceSetBbox(seq *TSequence, box unsafe.Pointer) {
	C.tsequence_set_bbox(seq._inner, unsafe.Pointer(box))
}


// TsequencesetExpandBbox wraps MEOS C function tsequenceset_expand_bbox.
func TsequencesetExpandBbox(ss *TSequenceSet, seq *TSequence) {
	C.tsequenceset_expand_bbox(ss._inner, seq._inner)
}


// TsequencesetSetBbox wraps MEOS C function tsequenceset_set_bbox.
func TsequencesetSetBbox(ss *TSequenceSet, box unsafe.Pointer) {
	C.tsequenceset_set_bbox(ss._inner, unsafe.Pointer(box))
}


// TcontseqAfterTimestamptz wraps MEOS C function tcontseq_after_timestamptz.
func TcontseqAfterTimestamptz(seq *TSequence, t int64, strict bool) *TSequence {
	_cret := C.tcontseq_after_timestamptz(seq._inner, C.TimestampTz(t), C.bool(strict))
	return &TSequence{_inner: _cret}
}


// TcontseqBeforeTimestamptz wraps MEOS C function tcontseq_before_timestamptz.
func TcontseqBeforeTimestamptz(seq *TSequence, t int64, strict bool) *TSequence {
	_cret := C.tcontseq_before_timestamptz(seq._inner, C.TimestampTz(t), C.bool(strict))
	return &TSequence{_inner: _cret}
}


// TcontseqRestrictMinmax wraps MEOS C function tcontseq_restrict_minmax.
func TcontseqRestrictMinmax(seq *TSequence, min bool, atfunc bool) *TSequenceSet {
	_cret := C.tcontseq_restrict_minmax(seq._inner, C.bool(min), C.bool(atfunc))
	return &TSequenceSet{_inner: _cret}
}


// TdiscseqAfterTimestamptz wraps MEOS C function tdiscseq_after_timestamptz.
func TdiscseqAfterTimestamptz(seq *TSequence, t int64, strict bool) *TSequence {
	_cret := C.tdiscseq_after_timestamptz(seq._inner, C.TimestampTz(t), C.bool(strict))
	return &TSequence{_inner: _cret}
}


// TdiscseqBeforeTimestamptz wraps MEOS C function tdiscseq_before_timestamptz.
func TdiscseqBeforeTimestamptz(seq *TSequence, t int64, strict bool) *TSequence {
	_cret := C.tdiscseq_before_timestamptz(seq._inner, C.TimestampTz(t), C.bool(strict))
	return &TSequence{_inner: _cret}
}


// TdiscseqRestrictMinmax wraps MEOS C function tdiscseq_restrict_minmax.
func TdiscseqRestrictMinmax(seq *TSequence, min bool, atfunc bool) *TSequence {
	_cret := C.tdiscseq_restrict_minmax(seq._inner, C.bool(min), C.bool(atfunc))
	return &TSequence{_inner: _cret}
}


// TemporalBboxRestrictSet wraps MEOS C function temporal_bbox_restrict_set.
func TemporalBboxRestrictSet(temp *Temporal, set *Set) bool {
	_cret := C.temporal_bbox_restrict_set(temp._inner, set._inner)
	return bool(_cret)
}


// TemporalRestrictMinmax wraps MEOS C function temporal_restrict_minmax.
func TemporalRestrictMinmax(temp *Temporal, min bool, atfunc bool) *Temporal {
	_cret := C.temporal_restrict_minmax(temp._inner, C.bool(min), C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TemporalRestrictTimestamptz wraps MEOS C function temporal_restrict_timestamptz.
func TemporalRestrictTimestamptz(temp *Temporal, t int64, atfunc bool) *Temporal {
	_cret := C.temporal_restrict_timestamptz(temp._inner, C.TimestampTz(t), C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TemporalRestrictTstzset wraps MEOS C function temporal_restrict_tstzset.
func TemporalRestrictTstzset(temp *Temporal, s *Set, atfunc bool) *Temporal {
	_cret := C.temporal_restrict_tstzset(temp._inner, s._inner, C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TemporalRestrictTstzspan wraps MEOS C function temporal_restrict_tstzspan.
func TemporalRestrictTstzspan(temp *Temporal, s *Span, atfunc bool) *Temporal {
	_cret := C.temporal_restrict_tstzspan(temp._inner, s._inner, C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TemporalRestrictTstzspanset wraps MEOS C function temporal_restrict_tstzspanset.
func TemporalRestrictTstzspanset(temp *Temporal, ss *SpanSet, atfunc bool) *Temporal {
	_cret := C.temporal_restrict_tstzspanset(temp._inner, ss._inner, C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TemporalRestrictValues wraps MEOS C function temporal_restrict_values.
func TemporalRestrictValues(temp *Temporal, set *Set, atfunc bool) *Temporal {
	_cret := C.temporal_restrict_values(temp._inner, set._inner, C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TinstantAfterTimestamptz wraps MEOS C function tinstant_after_timestamptz.
func TinstantAfterTimestamptz(inst *TInstant, t int64, strict bool) *TInstant {
	_cret := C.tinstant_after_timestamptz(inst._inner, C.TimestampTz(t), C.bool(strict))
	return &TInstant{_inner: _cret}
}


// TinstantBeforeTimestamptz wraps MEOS C function tinstant_before_timestamptz.
func TinstantBeforeTimestamptz(inst *TInstant, t int64, strict bool) *TInstant {
	_cret := C.tinstant_before_timestamptz(inst._inner, C.TimestampTz(t), C.bool(strict))
	return &TInstant{_inner: _cret}
}


// TinstantRestrictTstzspan wraps MEOS C function tinstant_restrict_tstzspan.
func TinstantRestrictTstzspan(inst *TInstant, period *Span, atfunc bool) *TInstant {
	_cret := C.tinstant_restrict_tstzspan(inst._inner, period._inner, C.bool(atfunc))
	return &TInstant{_inner: _cret}
}


// TinstantRestrictTstzspanset wraps MEOS C function tinstant_restrict_tstzspanset.
func TinstantRestrictTstzspanset(inst *TInstant, ss *SpanSet, atfunc bool) *TInstant {
	_cret := C.tinstant_restrict_tstzspanset(inst._inner, ss._inner, C.bool(atfunc))
	return &TInstant{_inner: _cret}
}


// TinstantRestrictTimestamptz wraps MEOS C function tinstant_restrict_timestamptz.
func TinstantRestrictTimestamptz(inst *TInstant, t int64, atfunc bool) *TInstant {
	_cret := C.tinstant_restrict_timestamptz(inst._inner, C.TimestampTz(t), C.bool(atfunc))
	return &TInstant{_inner: _cret}
}


// TinstantRestrictTstzset wraps MEOS C function tinstant_restrict_tstzset.
func TinstantRestrictTstzset(inst *TInstant, s *Set, atfunc bool) *TInstant {
	_cret := C.tinstant_restrict_tstzset(inst._inner, s._inner, C.bool(atfunc))
	return &TInstant{_inner: _cret}
}


// TinstantRestrictValues wraps MEOS C function tinstant_restrict_values.
func TinstantRestrictValues(inst *TInstant, set *Set, atfunc bool) *TInstant {
	_cret := C.tinstant_restrict_values(inst._inner, set._inner, C.bool(atfunc))
	return &TInstant{_inner: _cret}
}


// TnumberRestrictSpan wraps MEOS C function tnumber_restrict_span.
func TnumberRestrictSpan(temp *Temporal, span *Span, atfunc bool) *Temporal {
	_cret := C.tnumber_restrict_span(temp._inner, span._inner, C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TnumberRestrictSpanset wraps MEOS C function tnumber_restrict_spanset.
func TnumberRestrictSpanset(temp *Temporal, ss *SpanSet, atfunc bool) *Temporal {
	_cret := C.tnumber_restrict_spanset(temp._inner, ss._inner, C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TnumberinstRestrictSpan wraps MEOS C function tnumberinst_restrict_span.
func TnumberinstRestrictSpan(inst *TInstant, span *Span, atfunc bool) *TInstant {
	_cret := C.tnumberinst_restrict_span(inst._inner, span._inner, C.bool(atfunc))
	return &TInstant{_inner: _cret}
}


// TnumberinstRestrictSpanset wraps MEOS C function tnumberinst_restrict_spanset.
func TnumberinstRestrictSpanset(inst *TInstant, ss *SpanSet, atfunc bool) *TInstant {
	_cret := C.tnumberinst_restrict_spanset(inst._inner, ss._inner, C.bool(atfunc))
	return &TInstant{_inner: _cret}
}


// TnumberseqsetRestrictSpan wraps MEOS C function tnumberseqset_restrict_span.
func TnumberseqsetRestrictSpan(ss *TSequenceSet, span *Span, atfunc bool) *TSequenceSet {
	_cret := C.tnumberseqset_restrict_span(ss._inner, span._inner, C.bool(atfunc))
	return &TSequenceSet{_inner: _cret}
}


// TnumberseqsetRestrictSpanset wraps MEOS C function tnumberseqset_restrict_spanset.
func TnumberseqsetRestrictSpanset(ss *TSequenceSet, spanset *SpanSet, atfunc bool) *TSequenceSet {
	_cret := C.tnumberseqset_restrict_spanset(ss._inner, spanset._inner, C.bool(atfunc))
	return &TSequenceSet{_inner: _cret}
}


// TsequenceAtTimestamptz wraps MEOS C function tsequence_at_timestamptz.
func TsequenceAtTimestamptz(seq *TSequence, t int64) *TInstant {
	_cret := C.tsequence_at_timestamptz(seq._inner, C.TimestampTz(t))
	return &TInstant{_inner: _cret}
}


// TsequenceRestrictTstzspan wraps MEOS C function tsequence_restrict_tstzspan.
func TsequenceRestrictTstzspan(seq *TSequence, s *Span, atfunc bool) *Temporal {
	_cret := C.tsequence_restrict_tstzspan(seq._inner, s._inner, C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TsequenceRestrictTstzspanset wraps MEOS C function tsequence_restrict_tstzspanset.
func TsequenceRestrictTstzspanset(seq *TSequence, ss *SpanSet, atfunc bool) *Temporal {
	_cret := C.tsequence_restrict_tstzspanset(seq._inner, ss._inner, C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TsequencesetAfterTimestamptz wraps MEOS C function tsequenceset_after_timestamptz.
func TsequencesetAfterTimestamptz(ss *TSequenceSet, t int64, strict bool) *TSequenceSet {
	_cret := C.tsequenceset_after_timestamptz(ss._inner, C.TimestampTz(t), C.bool(strict))
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetBeforeTimestamptz wraps MEOS C function tsequenceset_before_timestamptz.
func TsequencesetBeforeTimestamptz(ss *TSequenceSet, t int64, strict bool) *TSequenceSet {
	_cret := C.tsequenceset_before_timestamptz(ss._inner, C.TimestampTz(t), C.bool(strict))
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetRestrictMinmax wraps MEOS C function tsequenceset_restrict_minmax.
func TsequencesetRestrictMinmax(ss *TSequenceSet, min bool, atfunc bool) *TSequenceSet {
	_cret := C.tsequenceset_restrict_minmax(ss._inner, C.bool(min), C.bool(atfunc))
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetRestrictTstzspan wraps MEOS C function tsequenceset_restrict_tstzspan.
func TsequencesetRestrictTstzspan(ss *TSequenceSet, s *Span, atfunc bool) *TSequenceSet {
	_cret := C.tsequenceset_restrict_tstzspan(ss._inner, s._inner, C.bool(atfunc))
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetRestrictTstzspanset wraps MEOS C function tsequenceset_restrict_tstzspanset.
func TsequencesetRestrictTstzspanset(ss *TSequenceSet, ps *SpanSet, atfunc bool) *TSequenceSet {
	_cret := C.tsequenceset_restrict_tstzspanset(ss._inner, ps._inner, C.bool(atfunc))
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetRestrictTimestamptz wraps MEOS C function tsequenceset_restrict_timestamptz.
func TsequencesetRestrictTimestamptz(ss *TSequenceSet, t int64, atfunc bool) *Temporal {
	_cret := C.tsequenceset_restrict_timestamptz(ss._inner, C.TimestampTz(t), C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TsequencesetRestrictTstzset wraps MEOS C function tsequenceset_restrict_tstzset.
func TsequencesetRestrictTstzset(ss *TSequenceSet, s *Set, atfunc bool) *Temporal {
	_cret := C.tsequenceset_restrict_tstzset(ss._inner, s._inner, C.bool(atfunc))
	return &Temporal{_inner: _cret}
}


// TsequencesetRestrictValues wraps MEOS C function tsequenceset_restrict_values.
func TsequencesetRestrictValues(ss *TSequenceSet, s *Set, atfunc bool) *TSequenceSet {
	_cret := C.tsequenceset_restrict_values(ss._inner, s._inner, C.bool(atfunc))
	return &TSequenceSet{_inner: _cret}
}


// TinstantCmp wraps MEOS C function tinstant_cmp.
func TinstantCmp(inst1 *TInstant, inst2 *TInstant) int {
	_cret := C.tinstant_cmp(inst1._inner, inst2._inner)
	return int(_cret)
}


// TinstantEq wraps MEOS C function tinstant_eq.
func TinstantEq(inst1 *TInstant, inst2 *TInstant) bool {
	_cret := C.tinstant_eq(inst1._inner, inst2._inner)
	return bool(_cret)
}


// TsequenceCmp wraps MEOS C function tsequence_cmp.
func TsequenceCmp(seq1 *TSequence, seq2 *TSequence) int {
	_cret := C.tsequence_cmp(seq1._inner, seq2._inner)
	return int(_cret)
}


// TsequenceEq wraps MEOS C function tsequence_eq.
func TsequenceEq(seq1 *TSequence, seq2 *TSequence) bool {
	_cret := C.tsequence_eq(seq1._inner, seq2._inner)
	return bool(_cret)
}


// TsequencesetCmp wraps MEOS C function tsequenceset_cmp.
func TsequencesetCmp(ss1 *TSequenceSet, ss2 *TSequenceSet) int {
	_cret := C.tsequenceset_cmp(ss1._inner, ss2._inner)
	return int(_cret)
}


// TsequencesetEq wraps MEOS C function tsequenceset_eq.
func TsequencesetEq(ss1 *TSequenceSet, ss2 *TSequenceSet) bool {
	_cret := C.tsequenceset_eq(ss1._inner, ss2._inner)
	return bool(_cret)
}


// TnumberinstAbs wraps MEOS C function tnumberinst_abs.
func TnumberinstAbs(inst *TInstant) *TInstant {
	_cret := C.tnumberinst_abs(inst._inner)
	return &TInstant{_inner: _cret}
}


// TnumberinstDistance wraps MEOS C function tnumberinst_distance.
func TnumberinstDistance(inst1 *TInstant, inst2 *TInstant) float64 {
	_cret := C.tnumberinst_distance(inst1._inner, inst2._inner)
	return float64(_cret)
}


// TnumberseqAbs wraps MEOS C function tnumberseq_abs.
func TnumberseqAbs(seq *TSequence) *TSequence {
	_cret := C.tnumberseq_abs(seq._inner)
	return &TSequence{_inner: _cret}
}


// TnumberseqAngularDifference wraps MEOS C function tnumberseq_angular_difference.
func TnumberseqAngularDifference(seq *TSequence) *TSequence {
	_cret := C.tnumberseq_angular_difference(seq._inner)
	return &TSequence{_inner: _cret}
}


// TnumberseqDeltaValue wraps MEOS C function tnumberseq_delta_value.
func TnumberseqDeltaValue(seq *TSequence) *TSequence {
	_cret := C.tnumberseq_delta_value(seq._inner)
	return &TSequence{_inner: _cret}
}


// TnumberseqsetAbs wraps MEOS C function tnumberseqset_abs.
func TnumberseqsetAbs(ss *TSequenceSet) *TSequenceSet {
	_cret := C.tnumberseqset_abs(ss._inner)
	return &TSequenceSet{_inner: _cret}
}


// TnumberseqsetAngularDifference wraps MEOS C function tnumberseqset_angular_difference.
func TnumberseqsetAngularDifference(ss *TSequenceSet) *TSequence {
	_cret := C.tnumberseqset_angular_difference(ss._inner)
	return &TSequence{_inner: _cret}
}


// TnumberseqsetDeltaValue wraps MEOS C function tnumberseqset_delta_value.
func TnumberseqsetDeltaValue(ss *TSequenceSet) *TSequenceSet {
	_cret := C.tnumberseqset_delta_value(ss._inner)
	return &TSequenceSet{_inner: _cret}
}


// TnumberseqIntegral wraps MEOS C function tnumberseq_integral.
func TnumberseqIntegral(seq *TSequence) float64 {
	_cret := C.tnumberseq_integral(seq._inner)
	return float64(_cret)
}


// TnumberseqTwavg wraps MEOS C function tnumberseq_twavg.
func TnumberseqTwavg(seq *TSequence) float64 {
	_cret := C.tnumberseq_twavg(seq._inner)
	return float64(_cret)
}


// TnumberseqsetIntegral wraps MEOS C function tnumberseqset_integral.
func TnumberseqsetIntegral(ss *TSequenceSet) float64 {
	_cret := C.tnumberseqset_integral(ss._inner)
	return float64(_cret)
}


// TnumberseqsetTwavg wraps MEOS C function tnumberseqset_twavg.
func TnumberseqsetTwavg(ss *TSequenceSet) float64 {
	_cret := C.tnumberseqset_twavg(ss._inner)
	return float64(_cret)
}


// TemporalCompact wraps MEOS C function temporal_compact.
func TemporalCompact(temp *Temporal) *Temporal {
	_cret := C.temporal_compact(temp._inner)
	return &Temporal{_inner: _cret}
}


// TsequenceCompact wraps MEOS C function tsequence_compact.
func TsequenceCompact(seq *TSequence) *TSequence {
	_cret := C.tsequence_compact(seq._inner)
	return &TSequence{_inner: _cret}
}


// TsequencesetCompact wraps MEOS C function tsequenceset_compact.
func TsequencesetCompact(ss *TSequenceSet) *TSequenceSet {
	_cret := C.tsequenceset_compact(ss._inner)
	return &TSequenceSet{_inner: _cret}
}


// TemporalSkiplistMake wraps MEOS C function temporal_skiplist_make.
func TemporalSkiplistMake() *SkipList {
	_cret := C.temporal_skiplist_make()
	return &SkipList{_inner: _cret}
}


// SkiplistMake wraps MEOS C function skiplist_make.
func SkiplistMake(key_size uint, value_size uint, comp_fn unsafe.Pointer, merge_fn unsafe.Pointer) *SkipList {
	_cret := C.skiplist_make(C.size_t(key_size), C.size_t(value_size), (*[0]byte)(comp_fn), (*[0]byte)(merge_fn))
	return &SkipList{_inner: _cret}
}


// SkiplistSearch wraps MEOS C function skiplist_search.
func SkiplistSearch(list *SkipList, key unsafe.Pointer, value unsafe.Pointer) int {
	_cret := C.skiplist_search(list._inner, unsafe.Pointer(key), unsafe.Pointer(value))
	return int(_cret)
}


// SkiplistFree wraps MEOS C function skiplist_free.
func SkiplistFree(list *SkipList) {
	C.skiplist_free(list._inner)
}


// SkiplistSplice wraps MEOS C function skiplist_splice.
func SkiplistSplice(list *SkipList, keys unsafe.Pointer, values unsafe.Pointer, count int, func_ unsafe.Pointer, crossings bool, sktype SkipListType) {
	C.skiplist_splice(list._inner, (*unsafe.Pointer)(unsafe.Pointer(keys)), (*unsafe.Pointer)(unsafe.Pointer(values)), C.int(count), (*[0]byte)(func_), C.bool(crossings), C.SkipListType(sktype))
}


// TemporalSkiplistSplice wraps MEOS C function temporal_skiplist_splice.
func TemporalSkiplistSplice(list *SkipList, values unsafe.Pointer, count int, func_ unsafe.Pointer, crossings bool) {
	C.temporal_skiplist_splice(list._inner, (*unsafe.Pointer)(unsafe.Pointer(values)), C.int(count), (*[0]byte)(func_), C.bool(crossings))
}


// SkiplistValues wraps MEOS C function skiplist_values.
func SkiplistValues(list *SkipList) unsafe.Pointer {
	_cret := C.skiplist_values(list._inner)
	return unsafe.Pointer(_cret)
}


// SkiplistKeysValues wraps MEOS C function skiplist_keys_values.
func SkiplistKeysValues(list *SkipList, values unsafe.Pointer) unsafe.Pointer {
	_cret := C.skiplist_keys_values(list._inner, (*unsafe.Pointer)(unsafe.Pointer(values)))
	return unsafe.Pointer(_cret)
}


// TemporalAppTinstTransfn wraps MEOS C function temporal_app_tinst_transfn.
func TemporalAppTinstTransfn(state *Temporal, inst *TInstant, interp Interpolation, maxdist float64, maxt *Interval) *Temporal {
	_cret := C.temporal_app_tinst_transfn(state._inner, inst._inner, C.interpType(interp), C.double(maxdist), maxt._inner)
	return &Temporal{_inner: _cret}
}


// TemporalAppTseqTransfn wraps MEOS C function temporal_app_tseq_transfn.
func TemporalAppTseqTransfn(state *Temporal, seq *TSequence) *Temporal {
	_cret := C.temporal_app_tseq_transfn(state._inner, seq._inner)
	return &Temporal{_inner: _cret}
}

