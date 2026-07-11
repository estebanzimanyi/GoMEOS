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

// BoolIn wraps MEOS C function bool_in.
func BoolIn(str string) bool {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.bool_in(_c_str)
	return bool(_cret)
}


// BoolOut wraps MEOS C function bool_out.
func BoolOut(b bool) string {
	_cret := C.bool_out(C.bool(b))
	return C.GoString(_cret)
}


// Float8Out wraps MEOS C function float8_out.
func Float8Out(num float64, maxdd int) string {
	_cret := C.float8_out(C.double(num), C.int(maxdd))
	return C.GoString(_cret)
}


// DateIn wraps MEOS C function date_in.
func DateIn(str string) int32 {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.date_in(_c_str)
	return int32(_cret)
}


// DateOut wraps MEOS C function date_out.
func DateOut(date int32) string {
	_cret := C.date_out(C.DateADT(date))
	return C.GoString(_cret)
}


// IntervalCmp wraps MEOS C function interval_cmp.
func IntervalCmp(interv1 *Interval, interv2 *Interval) int {
	_cret := C.interval_cmp(interv1._inner, interv2._inner)
	return int(_cret)
}


// IntervalIn wraps MEOS C function interval_in.
func IntervalIn(str string, typmod int32) *Interval {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.interval_in(_c_str, C.int32(typmod))
	return &Interval{_inner: _cret}
}


// IntervalOut wraps MEOS C function interval_out.
func IntervalOut(interv *Interval) string {
	_cret := C.interval_out(interv._inner)
	return C.GoString(_cret)
}


// TimeIn wraps MEOS C function time_in.
func TimeIn(str string, typmod int32) int64 {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.time_in(_c_str, C.int32(typmod))
	return int64(_cret)
}


// TimeOut wraps MEOS C function time_out.
func TimeOut(time int64) string {
	_cret := C.time_out(C.TimeADT(time))
	return C.GoString(_cret)
}


// TimestampIn wraps MEOS C function timestamp_in.
func TimestampIn(str string, typmod int32) int64 {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.timestamp_in(_c_str, C.int32(typmod))
	return int64(_cret)
}


// TimestampOut wraps MEOS C function timestamp_out.
func TimestampOut(ts int64) string {
	_cret := C.timestamp_out(C.Timestamp(ts))
	return C.GoString(_cret)
}


// TimestamptzIn wraps MEOS C function timestamptz_in.
func TimestamptzIn(str string, typmod int32) int64 {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.timestamptz_in(_c_str, C.int32(typmod))
	return int64(_cret)
}


// TimestamptzOut wraps MEOS C function timestamptz_out.
func TimestamptzOut(tstz int64) string {
	_cret := C.timestamptz_out(C.TimestampTz(tstz))
	return C.GoString(_cret)
}


// CstringToText wraps MEOS C function cstring_to_text.
func CstringToText(str string) string {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.cstring_to_text(_c_str)
	return C.GoString(C.text_to_cstring(_cret))
}


// TextToCstring wraps MEOS C function text_to_cstring.
func TextToCstring(txt string) string {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.text_to_cstring(_c_txt)
	return C.GoString(_cret)
}


// TextIn wraps MEOS C function text_in.
func TextIn(str string) string {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.text_in(_c_str)
	return C.GoString(C.text_to_cstring(_cret))
}


// TextOut wraps MEOS C function text_out.
func TextOut(txt string) string {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.text_out(_c_txt)
	return C.GoString(_cret)
}


// TODO text_cmp: unsupported param Oid
// func TextCmp(...) { /* not yet handled by codegen */ }


// TextCopy wraps MEOS C function text_copy.
func TextCopy(txt string) string {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.text_copy(_c_txt)
	return C.GoString(C.text_to_cstring(_cret))
}


// TextInitcap wraps MEOS C function text_initcap.
func TextInitcap(txt string) string {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.text_initcap(_c_txt)
	return C.GoString(C.text_to_cstring(_cret))
}


// TextLower wraps MEOS C function text_lower.
func TextLower(txt string) string {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.text_lower(_c_txt)
	return C.GoString(C.text_to_cstring(_cret))
}


// TextUpper wraps MEOS C function text_upper.
func TextUpper(txt string) string {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.text_upper(_c_txt)
	return C.GoString(C.text_to_cstring(_cret))
}


// TextcatTextText wraps MEOS C function textcat_text_text.
func TextcatTextText(txt1 string, txt2 string) string {
	_c_txt1 := C.cstring_to_text(C.CString(txt1))
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := C.cstring_to_text(C.CString(txt2))
	defer C.free(unsafe.Pointer(_c_txt2))
	_cret := C.textcat_text_text(_c_txt1, _c_txt2)
	return C.GoString(C.text_to_cstring(_cret))
}


// MeosErrno wraps MEOS C function meos_errno.
func MeosErrno() int {
	_cret := C.meos_errno()
	return int(_cret)
}


// MeosErrnoSet wraps MEOS C function meos_errno_set.
func MeosErrnoSet(err int) int {
	_cret := C.meos_errno_set(C.int(err))
	return int(_cret)
}


// MeosErrnoRestore wraps MEOS C function meos_errno_restore.
func MeosErrnoRestore(err int) int {
	_cret := C.meos_errno_restore(C.int(err))
	return int(_cret)
}


// MeosErrnoReset wraps MEOS C function meos_errno_reset.
func MeosErrnoReset() int {
	_cret := C.meos_errno_reset()
	return int(_cret)
}


// MeosArrayCreate wraps MEOS C function meos_array_create.
func MeosArrayCreate(elem_size int) *MeosArray {
	_cret := C.meos_array_create(C.int(elem_size))
	return &MeosArray{_inner: _cret}
}


// MeosArrayAdd wraps MEOS C function meos_array_add.
func MeosArrayAdd(array *MeosArray, value unsafe.Pointer) {
	C.meos_array_add(array._inner, unsafe.Pointer(value))
}


// MeosArrayGet wraps MEOS C function meos_array_get.
func MeosArrayGet(array *MeosArray, n int) unsafe.Pointer {
	_cret := C.meos_array_get(array._inner, C.int(n))
	return unsafe.Pointer(_cret)
}


// MeosArrayCount wraps MEOS C function meos_array_count.
func MeosArrayCount(array *MeosArray) int {
	_cret := C.meos_array_count(array._inner)
	return int(_cret)
}


// MeosArrayReset wraps MEOS C function meos_array_reset.
func MeosArrayReset(array *MeosArray) {
	C.meos_array_reset(array._inner)
}


// MeosArrayResetFree wraps MEOS C function meos_array_reset_free.
func MeosArrayResetFree(array *MeosArray) {
	C.meos_array_reset_free(array._inner)
}


// MeosArrayDestroy wraps MEOS C function meos_array_destroy.
func MeosArrayDestroy(array *MeosArray) {
	C.meos_array_destroy(array._inner)
}


// MeosArrayDestroyFree wraps MEOS C function meos_array_destroy_free.
func MeosArrayDestroyFree(array *MeosArray) {
	C.meos_array_destroy_free(array._inner)
}


// RtreeCreateIntspan wraps MEOS C function rtree_create_intspan.
func RtreeCreateIntspan() *RTree {
	_cret := C.rtree_create_intspan()
	return &RTree{_inner: _cret}
}


// RtreeCreateBigintspan wraps MEOS C function rtree_create_bigintspan.
func RtreeCreateBigintspan() *RTree {
	_cret := C.rtree_create_bigintspan()
	return &RTree{_inner: _cret}
}


// RtreeCreateFloatspan wraps MEOS C function rtree_create_floatspan.
func RtreeCreateFloatspan() *RTree {
	_cret := C.rtree_create_floatspan()
	return &RTree{_inner: _cret}
}


// RtreeCreateDatespan wraps MEOS C function rtree_create_datespan.
func RtreeCreateDatespan() *RTree {
	_cret := C.rtree_create_datespan()
	return &RTree{_inner: _cret}
}


// RtreeCreateTstzspan wraps MEOS C function rtree_create_tstzspan.
func RtreeCreateTstzspan() *RTree {
	_cret := C.rtree_create_tstzspan()
	return &RTree{_inner: _cret}
}


// RtreeCreateTBOX wraps MEOS C function rtree_create_tbox.
func RtreeCreateTBOX() *RTree {
	_cret := C.rtree_create_tbox()
	return &RTree{_inner: _cret}
}


// RtreeCreateSTBOX wraps MEOS C function rtree_create_stbox.
func RtreeCreateSTBOX() *RTree {
	_cret := C.rtree_create_stbox()
	return &RTree{_inner: _cret}
}


// RtreeCreateTpcbox wraps MEOS C function rtree_create_tpcbox.
func RtreeCreateTpcbox() *RTree {
	_cret := C.rtree_create_tpcbox()
	return &RTree{_inner: _cret}
}


// RtreeFree wraps MEOS C function rtree_free.
func RtreeFree(rtree *RTree) {
	C.rtree_free(rtree._inner)
}


// RtreeInsert wraps MEOS C function rtree_insert.
func RtreeInsert(rtree *RTree, box unsafe.Pointer, id int) {
	C.rtree_insert(rtree._inner, unsafe.Pointer(box), C.int(id))
}


// RtreeInsertTemporal wraps MEOS C function rtree_insert_temporal.
func RtreeInsertTemporal(rtree *RTree, temp *Temporal, id int) {
	C.rtree_insert_temporal(rtree._inner, temp._inner, C.int(id))
}


// RtreeInsertTemporalSplit wraps MEOS C function rtree_insert_temporal_split.
func RtreeInsertTemporalSplit(rtree *RTree, temp *Temporal, id int, maxboxes int) {
	C.rtree_insert_temporal_split(rtree._inner, temp._inner, C.int(id), C.int(maxboxes))
}


// RtreeSearch wraps MEOS C function rtree_search.
func RtreeSearch(rtree *RTree, op RTreeSearchOp, query unsafe.Pointer) (int, *MeosArray) {
	var _out_result C.MeosArray
	_cret := C.rtree_search(rtree._inner, C.RTreeSearchOp(op), unsafe.Pointer(query), &_out_result)
	return int(_cret), &MeosArray{_inner: &_out_result}
}


// RtreeSearchTemporal wraps MEOS C function rtree_search_temporal.
func RtreeSearchTemporal(rtree *RTree, op RTreeSearchOp, temp *Temporal) (int, *MeosArray) {
	var _out_result C.MeosArray
	_cret := C.rtree_search_temporal(rtree._inner, C.RTreeSearchOp(op), temp._inner, &_out_result)
	return int(_cret), &MeosArray{_inner: &_out_result}
}


// RtreeSearchTemporalDedup wraps MEOS C function rtree_search_temporal_dedup.
func RtreeSearchTemporalDedup(rtree *RTree, op RTreeSearchOp, temp *Temporal, maxboxes int) (int, *MeosArray) {
	var _out_result C.MeosArray
	_cret := C.rtree_search_temporal_dedup(rtree._inner, C.RTreeSearchOp(op), temp._inner, C.int(maxboxes), &_out_result)
	return int(_cret), &MeosArray{_inner: &_out_result}
}


// TODO meos_initialize_allocator: unsupported param meos_malloc_fn
// func MeosInitializeAllocator(...) { /* not yet handled by codegen */ }


// MeosInitializeNoexitErrorHandler wraps MEOS C function meos_initialize_noexit_error_handler.
func MeosInitializeNoexitErrorHandler() {
	C.meos_initialize_noexit_error_handler()
}


// MeosInitializeTimezone wraps MEOS C function meos_initialize_timezone.
func MeosInitializeTimezone(name string) {
	_c_name := C.CString(name)
	defer C.free(unsafe.Pointer(_c_name))
	C.meos_initialize_timezone(_c_name)
}


// MeosInitializeCollation wraps MEOS C function meos_initialize_collation.
func MeosInitializeCollation() {
	C.meos_initialize_collation()
}


// MeosFinalizeTimezone wraps MEOS C function meos_finalize_timezone.
func MeosFinalizeTimezone() {
	C.meos_finalize_timezone()
}


// MeosFinalizeCollation wraps MEOS C function meos_finalize_collation.
func MeosFinalizeCollation() {
	C.meos_finalize_collation()
}


// MeosFinalizeProjsrs wraps MEOS C function meos_finalize_projsrs.
func MeosFinalizeProjsrs() {
	C.meos_finalize_projsrs()
}


// MeosFinalizeWays wraps MEOS C function meos_finalize_ways.
func MeosFinalizeWays() {
	C.meos_finalize_ways()
}


// MeosInitializePointcloud wraps MEOS C function meos_initialize_pointcloud.
func MeosInitializePointcloud() {
	C.meos_initialize_pointcloud()
}


// MeosSetDatestyle wraps MEOS C function meos_set_datestyle.
func MeosSetDatestyle(newval string, extra unsafe.Pointer) bool {
	_c_newval := C.CString(newval)
	defer C.free(unsafe.Pointer(_c_newval))
	_cret := C.meos_set_datestyle(_c_newval, unsafe.Pointer(extra))
	return bool(_cret)
}


// MeosSetIntervalstyle wraps MEOS C function meos_set_intervalstyle.
func MeosSetIntervalstyle(newval string, extra int) bool {
	_c_newval := C.CString(newval)
	defer C.free(unsafe.Pointer(_c_newval))
	_cret := C.meos_set_intervalstyle(_c_newval, C.int(extra))
	return bool(_cret)
}


// MeosGetDatestyle wraps MEOS C function meos_get_datestyle.
func MeosGetDatestyle() string {
	_cret := C.meos_get_datestyle()
	return C.GoString(_cret)
}


// MeosGetIntervalstyle wraps MEOS C function meos_get_intervalstyle.
func MeosGetIntervalstyle() string {
	_cret := C.meos_get_intervalstyle()
	return C.GoString(_cret)
}


// MeosSetSpatialRefSysCsv wraps MEOS C function meos_set_spatial_ref_sys_csv.
func MeosSetSpatialRefSysCsv(path string) {
	_c_path := C.CString(path)
	defer C.free(unsafe.Pointer(_c_path))
	C.meos_set_spatial_ref_sys_csv(_c_path)
}


// MeosSetWaysCsv wraps MEOS C function meos_set_ways_csv.
func MeosSetWaysCsv(path string) {
	_c_path := C.CString(path)
	defer C.free(unsafe.Pointer(_c_path))
	C.meos_set_ways_csv(_c_path)
}


// MeosInitialize wraps MEOS C function meos_initialize.
func MeosInitialize() {
	C.meos_initialize()
}


// MeosFinalize wraps MEOS C function meos_finalize.
func MeosFinalize() {
	C.meos_finalize()
}


// BigintsetIn wraps MEOS C function bigintset_in.
func BigintsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.bigintset_in(_c_str)
	return &Set{_inner: _cret}
}


// BigintsetOut wraps MEOS C function bigintset_out.
func BigintsetOut(set *Set) string {
	_cret := C.bigintset_out(set._inner)
	return C.GoString(_cret)
}


// BigintspanExpand wraps MEOS C function bigintspan_expand.
func BigintspanExpand(s *Span, value int64) *Span {
	_cret := C.bigintspan_expand(s._inner, C.int64_t(value))
	return &Span{_inner: _cret}
}


// BigintspanIn wraps MEOS C function bigintspan_in.
func BigintspanIn(str string) *Span {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.bigintspan_in(_c_str)
	return &Span{_inner: _cret}
}


// BigintspanOut wraps MEOS C function bigintspan_out.
func BigintspanOut(s *Span) string {
	_cret := C.bigintspan_out(s._inner)
	return C.GoString(_cret)
}


// BigintspansetIn wraps MEOS C function bigintspanset_in.
func BigintspansetIn(str string) *SpanSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.bigintspanset_in(_c_str)
	return &SpanSet{_inner: _cret}
}


// BigintspansetOut wraps MEOS C function bigintspanset_out.
func BigintspansetOut(ss *SpanSet) string {
	_cret := C.bigintspanset_out(ss._inner)
	return C.GoString(_cret)
}


// DatesetIn wraps MEOS C function dateset_in.
func DatesetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.dateset_in(_c_str)
	return &Set{_inner: _cret}
}


// DatesetOut wraps MEOS C function dateset_out.
func DatesetOut(s *Set) string {
	_cret := C.dateset_out(s._inner)
	return C.GoString(_cret)
}


// DatespanIn wraps MEOS C function datespan_in.
func DatespanIn(str string) *Span {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.datespan_in(_c_str)
	return &Span{_inner: _cret}
}


// DatespanOut wraps MEOS C function datespan_out.
func DatespanOut(s *Span) string {
	_cret := C.datespan_out(s._inner)
	return C.GoString(_cret)
}


// DatespansetIn wraps MEOS C function datespanset_in.
func DatespansetIn(str string) *SpanSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.datespanset_in(_c_str)
	return &SpanSet{_inner: _cret}
}


// DatespansetOut wraps MEOS C function datespanset_out.
func DatespansetOut(ss *SpanSet) string {
	_cret := C.datespanset_out(ss._inner)
	return C.GoString(_cret)
}


// FloatsetIn wraps MEOS C function floatset_in.
func FloatsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.floatset_in(_c_str)
	return &Set{_inner: _cret}
}


// FloatsetOut wraps MEOS C function floatset_out.
func FloatsetOut(set *Set, maxdd int) string {
	_cret := C.floatset_out(set._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// FloatspanExpand wraps MEOS C function floatspan_expand.
func FloatspanExpand(s *Span, value float64) *Span {
	_cret := C.floatspan_expand(s._inner, C.double(value))
	return &Span{_inner: _cret}
}


// FloatspanIn wraps MEOS C function floatspan_in.
func FloatspanIn(str string) *Span {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.floatspan_in(_c_str)
	return &Span{_inner: _cret}
}


// FloatspanOut wraps MEOS C function floatspan_out.
func FloatspanOut(s *Span, maxdd int) string {
	_cret := C.floatspan_out(s._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// FloatspansetIn wraps MEOS C function floatspanset_in.
func FloatspansetIn(str string) *SpanSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.floatspanset_in(_c_str)
	return &SpanSet{_inner: _cret}
}


// FloatspansetOut wraps MEOS C function floatspanset_out.
func FloatspansetOut(ss *SpanSet, maxdd int) string {
	_cret := C.floatspanset_out(ss._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// IntsetIn wraps MEOS C function intset_in.
func IntsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.intset_in(_c_str)
	return &Set{_inner: _cret}
}


// IntsetOut wraps MEOS C function intset_out.
func IntsetOut(set *Set) string {
	_cret := C.intset_out(set._inner)
	return C.GoString(_cret)
}


// IntspanExpand wraps MEOS C function intspan_expand.
func IntspanExpand(s *Span, value int32) *Span {
	_cret := C.intspan_expand(s._inner, C.int32(value))
	return &Span{_inner: _cret}
}


// IntspanIn wraps MEOS C function intspan_in.
func IntspanIn(str string) *Span {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.intspan_in(_c_str)
	return &Span{_inner: _cret}
}


// IntspanOut wraps MEOS C function intspan_out.
func IntspanOut(s *Span) string {
	_cret := C.intspan_out(s._inner)
	return C.GoString(_cret)
}


// IntspansetIn wraps MEOS C function intspanset_in.
func IntspansetIn(str string) *SpanSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.intspanset_in(_c_str)
	return &SpanSet{_inner: _cret}
}


// IntspansetOut wraps MEOS C function intspanset_out.
func IntspansetOut(ss *SpanSet) string {
	_cret := C.intspanset_out(ss._inner)
	return C.GoString(_cret)
}


// SetAsHexwkb wraps MEOS C function set_as_hexwkb.
func SetAsHexwkb(s *Set, variant uint8, size_out unsafe.Pointer) string {
	_cret := C.set_as_hexwkb(s._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return C.GoString(_cret)
}


// SetAsWKB wraps MEOS C function set_as_wkb.
func SetAsWKB(s *Set, variant uint8, size_out unsafe.Pointer) unsafe.Pointer {
	_cret := C.set_as_wkb(s._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return unsafe.Pointer(_cret)
}


// SetFromHexwkb wraps MEOS C function set_from_hexwkb.
func SetFromHexwkb(hexwkb string) *Set {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	_cret := C.set_from_hexwkb(_c_hexwkb)
	return &Set{_inner: _cret}
}


// SetFromWKB wraps MEOS C function set_from_wkb.
func SetFromWKB(wkb unsafe.Pointer, size uint) *Set {
	_cret := C.set_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	return &Set{_inner: _cret}
}


// SpanAsHexwkb wraps MEOS C function span_as_hexwkb.
func SpanAsHexwkb(s *Span, variant uint8, size_out unsafe.Pointer) string {
	_cret := C.span_as_hexwkb(s._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return C.GoString(_cret)
}


// SpanAsWKB wraps MEOS C function span_as_wkb.
func SpanAsWKB(s *Span, variant uint8, size_out unsafe.Pointer) unsafe.Pointer {
	_cret := C.span_as_wkb(s._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return unsafe.Pointer(_cret)
}


// SpanFromHexwkb wraps MEOS C function span_from_hexwkb.
func SpanFromHexwkb(hexwkb string) *Span {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	_cret := C.span_from_hexwkb(_c_hexwkb)
	return &Span{_inner: _cret}
}


// SpanFromWKB wraps MEOS C function span_from_wkb.
func SpanFromWKB(wkb unsafe.Pointer, size uint) *Span {
	_cret := C.span_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	return &Span{_inner: _cret}
}


// SpansetAsHexwkb wraps MEOS C function spanset_as_hexwkb.
func SpansetAsHexwkb(ss *SpanSet, variant uint8, size_out unsafe.Pointer) string {
	_cret := C.spanset_as_hexwkb(ss._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return C.GoString(_cret)
}


// SpansetAsWKB wraps MEOS C function spanset_as_wkb.
func SpansetAsWKB(ss *SpanSet, variant uint8, size_out unsafe.Pointer) unsafe.Pointer {
	_cret := C.spanset_as_wkb(ss._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return unsafe.Pointer(_cret)
}


// SpansetFromHexwkb wraps MEOS C function spanset_from_hexwkb.
func SpansetFromHexwkb(hexwkb string) *SpanSet {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	_cret := C.spanset_from_hexwkb(_c_hexwkb)
	return &SpanSet{_inner: _cret}
}


// SpansetFromWKB wraps MEOS C function spanset_from_wkb.
func SpansetFromWKB(wkb unsafe.Pointer, size uint) *SpanSet {
	_cret := C.spanset_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	return &SpanSet{_inner: _cret}
}


// TextsetIn wraps MEOS C function textset_in.
func TextsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.textset_in(_c_str)
	return &Set{_inner: _cret}
}


// TextsetOut wraps MEOS C function textset_out.
func TextsetOut(set *Set) string {
	_cret := C.textset_out(set._inner)
	return C.GoString(_cret)
}


// TstzsetIn wraps MEOS C function tstzset_in.
func TstzsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tstzset_in(_c_str)
	return &Set{_inner: _cret}
}


// TstzsetOut wraps MEOS C function tstzset_out.
func TstzsetOut(set *Set) string {
	_cret := C.tstzset_out(set._inner)
	return C.GoString(_cret)
}


// TstzspanIn wraps MEOS C function tstzspan_in.
func TstzspanIn(str string) *Span {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tstzspan_in(_c_str)
	return &Span{_inner: _cret}
}


// TstzspanOut wraps MEOS C function tstzspan_out.
func TstzspanOut(s *Span) string {
	_cret := C.tstzspan_out(s._inner)
	return C.GoString(_cret)
}


// TstzspansetIn wraps MEOS C function tstzspanset_in.
func TstzspansetIn(str string) *SpanSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tstzspanset_in(_c_str)
	return &SpanSet{_inner: _cret}
}


// TstzspansetOut wraps MEOS C function tstzspanset_out.
func TstzspansetOut(ss *SpanSet) string {
	_cret := C.tstzspanset_out(ss._inner)
	return C.GoString(_cret)
}


// BigintsetMake wraps MEOS C function bigintset_make.
func BigintsetMake(values unsafe.Pointer, count int) *Set {
	_cret := C.bigintset_make((*C.int64_t)(unsafe.Pointer(values)), C.int(count))
	return &Set{_inner: _cret}
}


// BigintspanMake wraps MEOS C function bigintspan_make.
func BigintspanMake(lower int64, upper int64, lower_inc bool, upper_inc bool) *Span {
	_cret := C.bigintspan_make(C.int64_t(lower), C.int64_t(upper), C.bool(lower_inc), C.bool(upper_inc))
	return &Span{_inner: _cret}
}


// DatesetMake wraps MEOS C function dateset_make.
func DatesetMake(values unsafe.Pointer, count int) *Set {
	_cret := C.dateset_make((*C.DateADT)(unsafe.Pointer(values)), C.int(count))
	return &Set{_inner: _cret}
}


// DatespanMake wraps MEOS C function datespan_make.
func DatespanMake(lower int32, upper int32, lower_inc bool, upper_inc bool) *Span {
	_cret := C.datespan_make(C.DateADT(lower), C.DateADT(upper), C.bool(lower_inc), C.bool(upper_inc))
	return &Span{_inner: _cret}
}


// FloatsetMake wraps MEOS C function floatset_make.
func FloatsetMake(values unsafe.Pointer, count int) *Set {
	_cret := C.floatset_make((*C.double)(unsafe.Pointer(values)), C.int(count))
	return &Set{_inner: _cret}
}


// FloatspanMake wraps MEOS C function floatspan_make.
func FloatspanMake(lower float64, upper float64, lower_inc bool, upper_inc bool) *Span {
	_cret := C.floatspan_make(C.double(lower), C.double(upper), C.bool(lower_inc), C.bool(upper_inc))
	return &Span{_inner: _cret}
}


// IntsetMake wraps MEOS C function intset_make.
func IntsetMake(values unsafe.Pointer, count int) *Set {
	_cret := C.intset_make((*C.int)(unsafe.Pointer(values)), C.int(count))
	return &Set{_inner: _cret}
}


// IntspanMake wraps MEOS C function intspan_make.
func IntspanMake(lower int, upper int, lower_inc bool, upper_inc bool) *Span {
	_cret := C.intspan_make(C.int(lower), C.int(upper), C.bool(lower_inc), C.bool(upper_inc))
	return &Span{_inner: _cret}
}


// SetCopy wraps MEOS C function set_copy.
func SetCopy(s *Set) *Set {
	_cret := C.set_copy(s._inner)
	return &Set{_inner: _cret}
}


// SpanCopy wraps MEOS C function span_copy.
func SpanCopy(s *Span) *Span {
	_cret := C.span_copy(s._inner)
	return &Span{_inner: _cret}
}


// SpansetCopy wraps MEOS C function spanset_copy.
func SpansetCopy(ss *SpanSet) *SpanSet {
	_cret := C.spanset_copy(ss._inner)
	return &SpanSet{_inner: _cret}
}


// SpansetMake wraps MEOS C function spanset_make.
func SpansetMake(spans *Span, count int) *SpanSet {
	_cret := C.spanset_make(spans._inner, C.int(count))
	return &SpanSet{_inner: _cret}
}


// TextsetMake wraps MEOS C function textset_make.
func TextsetMake(values unsafe.Pointer, count int) *Set {
	_cret := C.textset_make((**C.text)(unsafe.Pointer(values)), C.int(count))
	return &Set{_inner: _cret}
}


// TstzsetMake wraps MEOS C function tstzset_make.
func TstzsetMake(values unsafe.Pointer, count int) *Set {
	_cret := C.tstzset_make((*C.TimestampTz)(unsafe.Pointer(values)), C.int(count))
	return &Set{_inner: _cret}
}


// TstzspanMake wraps MEOS C function tstzspan_make.
func TstzspanMake(lower int64, upper int64, lower_inc bool, upper_inc bool) *Span {
	_cret := C.tstzspan_make(C.TimestampTz(lower), C.TimestampTz(upper), C.bool(lower_inc), C.bool(upper_inc))
	return &Span{_inner: _cret}
}


// BigintToSet wraps MEOS C function bigint_to_set.
func BigintToSet(i int64) *Set {
	_cret := C.bigint_to_set(C.int64_t(i))
	return &Set{_inner: _cret}
}


// BigintToSpan wraps MEOS C function bigint_to_span.
func BigintToSpan(i int) *Span {
	_cret := C.bigint_to_span(C.int(i))
	return &Span{_inner: _cret}
}


// BigintToSpanset wraps MEOS C function bigint_to_spanset.
func BigintToSpanset(i int) *SpanSet {
	_cret := C.bigint_to_spanset(C.int(i))
	return &SpanSet{_inner: _cret}
}


// DateToSet wraps MEOS C function date_to_set.
func DateToSet(d int32) *Set {
	_cret := C.date_to_set(C.DateADT(d))
	return &Set{_inner: _cret}
}


// DateToSpan wraps MEOS C function date_to_span.
func DateToSpan(d int32) *Span {
	_cret := C.date_to_span(C.DateADT(d))
	return &Span{_inner: _cret}
}


// DateToSpanset wraps MEOS C function date_to_spanset.
func DateToSpanset(d int32) *SpanSet {
	_cret := C.date_to_spanset(C.DateADT(d))
	return &SpanSet{_inner: _cret}
}


// DatesetToTstzset wraps MEOS C function dateset_to_tstzset.
func DatesetToTstzset(s *Set) *Set {
	_cret := C.dateset_to_tstzset(s._inner)
	return &Set{_inner: _cret}
}


// DatespanToTstzspan wraps MEOS C function datespan_to_tstzspan.
func DatespanToTstzspan(s *Span) *Span {
	_cret := C.datespan_to_tstzspan(s._inner)
	return &Span{_inner: _cret}
}


// DatespansetToTstzspanset wraps MEOS C function datespanset_to_tstzspanset.
func DatespansetToTstzspanset(ss *SpanSet) *SpanSet {
	_cret := C.datespanset_to_tstzspanset(ss._inner)
	return &SpanSet{_inner: _cret}
}


// FloatToSet wraps MEOS C function float_to_set.
func FloatToSet(d float64) *Set {
	_cret := C.float_to_set(C.double(d))
	return &Set{_inner: _cret}
}


// FloatToSpan wraps MEOS C function float_to_span.
func FloatToSpan(d float64) *Span {
	_cret := C.float_to_span(C.double(d))
	return &Span{_inner: _cret}
}


// FloatToSpanset wraps MEOS C function float_to_spanset.
func FloatToSpanset(d float64) *SpanSet {
	_cret := C.float_to_spanset(C.double(d))
	return &SpanSet{_inner: _cret}
}


// FloatsetToIntset wraps MEOS C function floatset_to_intset.
func FloatsetToIntset(s *Set) *Set {
	_cret := C.floatset_to_intset(s._inner)
	return &Set{_inner: _cret}
}


// FloatspanToIntspan wraps MEOS C function floatspan_to_intspan.
func FloatspanToIntspan(s *Span) *Span {
	_cret := C.floatspan_to_intspan(s._inner)
	return &Span{_inner: _cret}
}


// FloatspanToBigintspan wraps MEOS C function floatspan_to_bigintspan.
func FloatspanToBigintspan(s *Span) *Span {
	_cret := C.floatspan_to_bigintspan(s._inner)
	return &Span{_inner: _cret}
}


// FloatspansetToIntspanset wraps MEOS C function floatspanset_to_intspanset.
func FloatspansetToIntspanset(ss *SpanSet) *SpanSet {
	_cret := C.floatspanset_to_intspanset(ss._inner)
	return &SpanSet{_inner: _cret}
}


// IntToSet wraps MEOS C function int_to_set.
func IntToSet(i int) *Set {
	_cret := C.int_to_set(C.int(i))
	return &Set{_inner: _cret}
}


// IntToSpan wraps MEOS C function int_to_span.
func IntToSpan(i int) *Span {
	_cret := C.int_to_span(C.int(i))
	return &Span{_inner: _cret}
}


// IntToSpanset wraps MEOS C function int_to_spanset.
func IntToSpanset(i int) *SpanSet {
	_cret := C.int_to_spanset(C.int(i))
	return &SpanSet{_inner: _cret}
}


// IntsetToFloatset wraps MEOS C function intset_to_floatset.
func IntsetToFloatset(s *Set) *Set {
	_cret := C.intset_to_floatset(s._inner)
	return &Set{_inner: _cret}
}


// IntspanToFloatspan wraps MEOS C function intspan_to_floatspan.
func IntspanToFloatspan(s *Span) *Span {
	_cret := C.intspan_to_floatspan(s._inner)
	return &Span{_inner: _cret}
}


// IntspanToBigintspan wraps MEOS C function intspan_to_bigintspan.
func IntspanToBigintspan(s *Span) *Span {
	_cret := C.intspan_to_bigintspan(s._inner)
	return &Span{_inner: _cret}
}


// BigintspanToIntspan wraps MEOS C function bigintspan_to_intspan.
func BigintspanToIntspan(s *Span) *Span {
	_cret := C.bigintspan_to_intspan(s._inner)
	return &Span{_inner: _cret}
}


// BigintspanToFloatspan wraps MEOS C function bigintspan_to_floatspan.
func BigintspanToFloatspan(s *Span) *Span {
	_cret := C.bigintspan_to_floatspan(s._inner)
	return &Span{_inner: _cret}
}


// IntspansetToFloatspanset wraps MEOS C function intspanset_to_floatspanset.
func IntspansetToFloatspanset(ss *SpanSet) *SpanSet {
	_cret := C.intspanset_to_floatspanset(ss._inner)
	return &SpanSet{_inner: _cret}
}


// SetToSpan wraps MEOS C function set_to_span.
func SetToSpan(s *Set) *Span {
	_cret := C.set_to_span(s._inner)
	return &Span{_inner: _cret}
}


// SetToSpanset wraps MEOS C function set_to_spanset.
func SetToSpanset(s *Set) *SpanSet {
	_cret := C.set_to_spanset(s._inner)
	return &SpanSet{_inner: _cret}
}


// SpanToSpanset wraps MEOS C function span_to_spanset.
func SpanToSpanset(s *Span) *SpanSet {
	_cret := C.span_to_spanset(s._inner)
	return &SpanSet{_inner: _cret}
}


// TextToSet wraps MEOS C function text_to_set.
func TextToSet(txt string) *Set {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.text_to_set(_c_txt)
	return &Set{_inner: _cret}
}


// TimestamptzToSet wraps MEOS C function timestamptz_to_set.
func TimestamptzToSet(t int64) *Set {
	_cret := C.timestamptz_to_set(C.TimestampTz(t))
	return &Set{_inner: _cret}
}


// TimestamptzToSpan wraps MEOS C function timestamptz_to_span.
func TimestamptzToSpan(t int64) *Span {
	_cret := C.timestamptz_to_span(C.TimestampTz(t))
	return &Span{_inner: _cret}
}


// TimestamptzToSpanset wraps MEOS C function timestamptz_to_spanset.
func TimestamptzToSpanset(t int64) *SpanSet {
	_cret := C.timestamptz_to_spanset(C.TimestampTz(t))
	return &SpanSet{_inner: _cret}
}


// TstzsetToDateset wraps MEOS C function tstzset_to_dateset.
func TstzsetToDateset(s *Set) *Set {
	_cret := C.tstzset_to_dateset(s._inner)
	return &Set{_inner: _cret}
}


// TstzspanToDatespan wraps MEOS C function tstzspan_to_datespan.
func TstzspanToDatespan(s *Span) *Span {
	_cret := C.tstzspan_to_datespan(s._inner)
	return &Span{_inner: _cret}
}


// TstzspansetToDatespanset wraps MEOS C function tstzspanset_to_datespanset.
func TstzspansetToDatespanset(ss *SpanSet) *SpanSet {
	_cret := C.tstzspanset_to_datespanset(ss._inner)
	return &SpanSet{_inner: _cret}
}


// BigintsetEndValue wraps MEOS C function bigintset_end_value.
func BigintsetEndValue(s *Set) int64 {
	_cret := C.bigintset_end_value(s._inner)
	return int64(_cret)
}


// BigintsetStartValue wraps MEOS C function bigintset_start_value.
func BigintsetStartValue(s *Set) int64 {
	_cret := C.bigintset_start_value(s._inner)
	return int64(_cret)
}


// BigintsetValueN wraps MEOS C function bigintset_value_n.
func BigintsetValueN(s *Set, n int) (bool, int64) {
	var _out_result C.int64_t
	_cret := C.bigintset_value_n(s._inner, C.int(n), &_out_result)
	return bool(_cret), int64(_out_result)
}


// BigintsetValues wraps MEOS C function bigintset_values.
func BigintsetValues(s *Set, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.bigintset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// BigintspanLower wraps MEOS C function bigintspan_lower.
func BigintspanLower(s *Span) int64 {
	_cret := C.bigintspan_lower(s._inner)
	return int64(_cret)
}


// BigintspanUpper wraps MEOS C function bigintspan_upper.
func BigintspanUpper(s *Span) int64 {
	_cret := C.bigintspan_upper(s._inner)
	return int64(_cret)
}


// BigintspanWidth wraps MEOS C function bigintspan_width.
func BigintspanWidth(s *Span) int64 {
	_cret := C.bigintspan_width(s._inner)
	return int64(_cret)
}


// BigintspansetLower wraps MEOS C function bigintspanset_lower.
func BigintspansetLower(ss *SpanSet) int64 {
	_cret := C.bigintspanset_lower(ss._inner)
	return int64(_cret)
}


// BigintspansetUpper wraps MEOS C function bigintspanset_upper.
func BigintspansetUpper(ss *SpanSet) int64 {
	_cret := C.bigintspanset_upper(ss._inner)
	return int64(_cret)
}


// BigintspansetWidth wraps MEOS C function bigintspanset_width.
func BigintspansetWidth(ss *SpanSet, boundspan bool) int64 {
	_cret := C.bigintspanset_width(ss._inner, C.bool(boundspan))
	return int64(_cret)
}


// DatesetEndValue wraps MEOS C function dateset_end_value.
func DatesetEndValue(s *Set) int32 {
	_cret := C.dateset_end_value(s._inner)
	return int32(_cret)
}


// DatesetStartValue wraps MEOS C function dateset_start_value.
func DatesetStartValue(s *Set) int32 {
	_cret := C.dateset_start_value(s._inner)
	return int32(_cret)
}


// DatesetValueN wraps MEOS C function dateset_value_n.
func DatesetValueN(s *Set, n int) (bool, int32) {
	var _out_result C.DateADT
	_cret := C.dateset_value_n(s._inner, C.int(n), &_out_result)
	return bool(_cret), int32(_out_result)
}


// DatesetValues wraps MEOS C function dateset_values.
func DatesetValues(s *Set, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.dateset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// DatespanDuration wraps MEOS C function datespan_duration.
func DatespanDuration(s *Span) *Interval {
	_cret := C.datespan_duration(s._inner)
	return &Interval{_inner: _cret}
}


// DatespanLower wraps MEOS C function datespan_lower.
func DatespanLower(s *Span) int32 {
	_cret := C.datespan_lower(s._inner)
	return int32(_cret)
}


// DatespanUpper wraps MEOS C function datespan_upper.
func DatespanUpper(s *Span) int32 {
	_cret := C.datespan_upper(s._inner)
	return int32(_cret)
}


// DatespansetDateN wraps MEOS C function datespanset_date_n.
func DatespansetDateN(ss *SpanSet, n int) (bool, int32) {
	var _out_result C.DateADT
	_cret := C.datespanset_date_n(ss._inner, C.int(n), &_out_result)
	return bool(_cret), int32(_out_result)
}


// DatespansetDates wraps MEOS C function datespanset_dates.
func DatespansetDates(ss *SpanSet) *Set {
	_cret := C.datespanset_dates(ss._inner)
	return &Set{_inner: _cret}
}


// DatespansetDuration wraps MEOS C function datespanset_duration.
func DatespansetDuration(ss *SpanSet, boundspan bool) *Interval {
	_cret := C.datespanset_duration(ss._inner, C.bool(boundspan))
	return &Interval{_inner: _cret}
}


// DatespansetEndDate wraps MEOS C function datespanset_end_date.
func DatespansetEndDate(ss *SpanSet) int32 {
	_cret := C.datespanset_end_date(ss._inner)
	return int32(_cret)
}


// DatespansetNumDates wraps MEOS C function datespanset_num_dates.
func DatespansetNumDates(ss *SpanSet) int {
	_cret := C.datespanset_num_dates(ss._inner)
	return int(_cret)
}


// DatespansetStartDate wraps MEOS C function datespanset_start_date.
func DatespansetStartDate(ss *SpanSet) int32 {
	_cret := C.datespanset_start_date(ss._inner)
	return int32(_cret)
}


// FloatsetEndValue wraps MEOS C function floatset_end_value.
func FloatsetEndValue(s *Set) float64 {
	_cret := C.floatset_end_value(s._inner)
	return float64(_cret)
}


// FloatsetStartValue wraps MEOS C function floatset_start_value.
func FloatsetStartValue(s *Set) float64 {
	_cret := C.floatset_start_value(s._inner)
	return float64(_cret)
}


// FloatsetValueN wraps MEOS C function floatset_value_n.
func FloatsetValueN(s *Set, n int) (bool, float64) {
	var _out_result C.double
	_cret := C.floatset_value_n(s._inner, C.int(n), &_out_result)
	return bool(_cret), float64(_out_result)
}


// FloatsetValues wraps MEOS C function floatset_values.
func FloatsetValues(s *Set, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.floatset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// FloatspanLower wraps MEOS C function floatspan_lower.
func FloatspanLower(s *Span) float64 {
	_cret := C.floatspan_lower(s._inner)
	return float64(_cret)
}


// FloatspanUpper wraps MEOS C function floatspan_upper.
func FloatspanUpper(s *Span) float64 {
	_cret := C.floatspan_upper(s._inner)
	return float64(_cret)
}


// FloatspanWidth wraps MEOS C function floatspan_width.
func FloatspanWidth(s *Span) float64 {
	_cret := C.floatspan_width(s._inner)
	return float64(_cret)
}


// FloatspansetLower wraps MEOS C function floatspanset_lower.
func FloatspansetLower(ss *SpanSet) float64 {
	_cret := C.floatspanset_lower(ss._inner)
	return float64(_cret)
}


// FloatspansetUpper wraps MEOS C function floatspanset_upper.
func FloatspansetUpper(ss *SpanSet) float64 {
	_cret := C.floatspanset_upper(ss._inner)
	return float64(_cret)
}


// FloatspansetWidth wraps MEOS C function floatspanset_width.
func FloatspansetWidth(ss *SpanSet, boundspan bool) float64 {
	_cret := C.floatspanset_width(ss._inner, C.bool(boundspan))
	return float64(_cret)
}


// IntsetEndValue wraps MEOS C function intset_end_value.
func IntsetEndValue(s *Set) int {
	_cret := C.intset_end_value(s._inner)
	return int(_cret)
}


// IntsetStartValue wraps MEOS C function intset_start_value.
func IntsetStartValue(s *Set) int {
	_cret := C.intset_start_value(s._inner)
	return int(_cret)
}


// IntsetValueN wraps MEOS C function intset_value_n.
func IntsetValueN(s *Set, n int) (bool, int) {
	var _out_result C.int
	_cret := C.intset_value_n(s._inner, C.int(n), &_out_result)
	return bool(_cret), int(_out_result)
}


// IntsetValues wraps MEOS C function intset_values.
func IntsetValues(s *Set, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.intset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// IntspanLower wraps MEOS C function intspan_lower.
func IntspanLower(s *Span) int {
	_cret := C.intspan_lower(s._inner)
	return int(_cret)
}


// IntspanUpper wraps MEOS C function intspan_upper.
func IntspanUpper(s *Span) int {
	_cret := C.intspan_upper(s._inner)
	return int(_cret)
}


// IntspanWidth wraps MEOS C function intspan_width.
func IntspanWidth(s *Span) int {
	_cret := C.intspan_width(s._inner)
	return int(_cret)
}


// IntspansetLower wraps MEOS C function intspanset_lower.
func IntspansetLower(ss *SpanSet) int {
	_cret := C.intspanset_lower(ss._inner)
	return int(_cret)
}


// IntspansetUpper wraps MEOS C function intspanset_upper.
func IntspansetUpper(ss *SpanSet) int {
	_cret := C.intspanset_upper(ss._inner)
	return int(_cret)
}


// IntspansetWidth wraps MEOS C function intspanset_width.
func IntspansetWidth(ss *SpanSet, boundspan bool) int {
	_cret := C.intspanset_width(ss._inner, C.bool(boundspan))
	return int(_cret)
}


// SetHash wraps MEOS C function set_hash.
func SetHash(s *Set) uint32 {
	_cret := C.set_hash(s._inner)
	return uint32(_cret)
}


// SetHashExtended wraps MEOS C function set_hash_extended.
func SetHashExtended(s *Set, seed uint64) uint64 {
	_cret := C.set_hash_extended(s._inner, C.uint64_t(seed))
	return uint64(_cret)
}


// SetNumValues wraps MEOS C function set_num_values.
func SetNumValues(s *Set) int {
	_cret := C.set_num_values(s._inner)
	return int(_cret)
}


// SpanHash wraps MEOS C function span_hash.
func SpanHash(s *Span) uint32 {
	_cret := C.span_hash(s._inner)
	return uint32(_cret)
}


// SpanHashExtended wraps MEOS C function span_hash_extended.
func SpanHashExtended(s *Span, seed uint64) uint64 {
	_cret := C.span_hash_extended(s._inner, C.uint64_t(seed))
	return uint64(_cret)
}


// SpanLowerInc wraps MEOS C function span_lower_inc.
func SpanLowerInc(s *Span) bool {
	_cret := C.span_lower_inc(s._inner)
	return bool(_cret)
}


// SpanUpperInc wraps MEOS C function span_upper_inc.
func SpanUpperInc(s *Span) bool {
	_cret := C.span_upper_inc(s._inner)
	return bool(_cret)
}


// SpansetEndSpan wraps MEOS C function spanset_end_span.
func SpansetEndSpan(ss *SpanSet) *Span {
	_cret := C.spanset_end_span(ss._inner)
	return &Span{_inner: _cret}
}


// SpansetHash wraps MEOS C function spanset_hash.
func SpansetHash(ss *SpanSet) uint32 {
	_cret := C.spanset_hash(ss._inner)
	return uint32(_cret)
}


// SpansetHashExtended wraps MEOS C function spanset_hash_extended.
func SpansetHashExtended(ss *SpanSet, seed uint64) uint64 {
	_cret := C.spanset_hash_extended(ss._inner, C.uint64_t(seed))
	return uint64(_cret)
}


// SpansetLowerInc wraps MEOS C function spanset_lower_inc.
func SpansetLowerInc(ss *SpanSet) bool {
	_cret := C.spanset_lower_inc(ss._inner)
	return bool(_cret)
}


// SpansetNumSpans wraps MEOS C function spanset_num_spans.
func SpansetNumSpans(ss *SpanSet) int {
	_cret := C.spanset_num_spans(ss._inner)
	return int(_cret)
}


// SpansetSpan wraps MEOS C function spanset_span.
func SpansetSpan(ss *SpanSet) *Span {
	_cret := C.spanset_span(ss._inner)
	return &Span{_inner: _cret}
}


// SpansetSpanN wraps MEOS C function spanset_span_n.
func SpansetSpanN(ss *SpanSet, i int) *Span {
	_cret := C.spanset_span_n(ss._inner, C.int(i))
	return &Span{_inner: _cret}
}


// SpansetSpanarr wraps MEOS C function spanset_spanarr.
func SpansetSpanarr(ss *SpanSet, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.spanset_spanarr(ss._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// SpansetStartSpan wraps MEOS C function spanset_start_span.
func SpansetStartSpan(ss *SpanSet) *Span {
	_cret := C.spanset_start_span(ss._inner)
	return &Span{_inner: _cret}
}


// SpansetUpperInc wraps MEOS C function spanset_upper_inc.
func SpansetUpperInc(ss *SpanSet) bool {
	_cret := C.spanset_upper_inc(ss._inner)
	return bool(_cret)
}


// TextsetEndValue wraps MEOS C function textset_end_value.
func TextsetEndValue(s *Set) string {
	_cret := C.textset_end_value(s._inner)
	return C.GoString(C.text_to_cstring(_cret))
}


// TextsetStartValue wraps MEOS C function textset_start_value.
func TextsetStartValue(s *Set) string {
	_cret := C.textset_start_value(s._inner)
	return C.GoString(C.text_to_cstring(_cret))
}


// TextsetValueN wraps MEOS C function textset_value_n.
func TextsetValueN(s *Set, n int) (bool, string) {
	var _out_result *C.text
	_cret := C.textset_value_n(s._inner, C.int(n), &_out_result)
	return bool(_cret), C.GoString(C.text_to_cstring(_out_result))
}


// TextsetValues wraps MEOS C function textset_values.
func TextsetValues(s *Set, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.textset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TstzsetEndValue wraps MEOS C function tstzset_end_value.
func TstzsetEndValue(s *Set) int64 {
	_cret := C.tstzset_end_value(s._inner)
	return int64(_cret)
}


// TstzsetStartValue wraps MEOS C function tstzset_start_value.
func TstzsetStartValue(s *Set) int64 {
	_cret := C.tstzset_start_value(s._inner)
	return int64(_cret)
}


// TstzsetValueN wraps MEOS C function tstzset_value_n.
func TstzsetValueN(s *Set, n int) (bool, int64) {
	var _out_result C.TimestampTz
	_cret := C.tstzset_value_n(s._inner, C.int(n), &_out_result)
	return bool(_cret), int64(_out_result)
}


// TstzsetValues wraps MEOS C function tstzset_values.
func TstzsetValues(s *Set, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tstzset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TstzspanDuration wraps MEOS C function tstzspan_duration.
func TstzspanDuration(s *Span) *Interval {
	_cret := C.tstzspan_duration(s._inner)
	return &Interval{_inner: _cret}
}


// TstzspanLower wraps MEOS C function tstzspan_lower.
func TstzspanLower(s *Span) int64 {
	_cret := C.tstzspan_lower(s._inner)
	return int64(_cret)
}


// TstzspanUpper wraps MEOS C function tstzspan_upper.
func TstzspanUpper(s *Span) int64 {
	_cret := C.tstzspan_upper(s._inner)
	return int64(_cret)
}


// TstzspansetDuration wraps MEOS C function tstzspanset_duration.
func TstzspansetDuration(ss *SpanSet, boundspan bool) *Interval {
	_cret := C.tstzspanset_duration(ss._inner, C.bool(boundspan))
	return &Interval{_inner: _cret}
}


// TstzspansetEndTimestamptz wraps MEOS C function tstzspanset_end_timestamptz.
func TstzspansetEndTimestamptz(ss *SpanSet) int64 {
	_cret := C.tstzspanset_end_timestamptz(ss._inner)
	return int64(_cret)
}


// TstzspansetLower wraps MEOS C function tstzspanset_lower.
func TstzspansetLower(ss *SpanSet) int64 {
	_cret := C.tstzspanset_lower(ss._inner)
	return int64(_cret)
}


// TstzspansetNumTimestamps wraps MEOS C function tstzspanset_num_timestamps.
func TstzspansetNumTimestamps(ss *SpanSet) int {
	_cret := C.tstzspanset_num_timestamps(ss._inner)
	return int(_cret)
}


// TstzspansetStartTimestamptz wraps MEOS C function tstzspanset_start_timestamptz.
func TstzspansetStartTimestamptz(ss *SpanSet) int64 {
	_cret := C.tstzspanset_start_timestamptz(ss._inner)
	return int64(_cret)
}


// TstzspansetTimestamps wraps MEOS C function tstzspanset_timestamps.
func TstzspansetTimestamps(ss *SpanSet) *Set {
	_cret := C.tstzspanset_timestamps(ss._inner)
	return &Set{_inner: _cret}
}


// TstzspansetTimestamptzN wraps MEOS C function tstzspanset_timestamptz_n.
func TstzspansetTimestamptzN(ss *SpanSet, n int) (bool, int64) {
	var _out_result C.TimestampTz
	_cret := C.tstzspanset_timestamptz_n(ss._inner, C.int(n), &_out_result)
	return bool(_cret), int64(_out_result)
}


// TstzspansetUpper wraps MEOS C function tstzspanset_upper.
func TstzspansetUpper(ss *SpanSet) int64 {
	_cret := C.tstzspanset_upper(ss._inner)
	return int64(_cret)
}


// BigintsetShiftScale wraps MEOS C function bigintset_shift_scale.
func BigintsetShiftScale(s *Set, shift int64, width int64, hasshift bool, haswidth bool) *Set {
	_cret := C.bigintset_shift_scale(s._inner, C.int64_t(shift), C.int64_t(width), C.bool(hasshift), C.bool(haswidth))
	return &Set{_inner: _cret}
}


// BigintspanShiftScale wraps MEOS C function bigintspan_shift_scale.
func BigintspanShiftScale(s *Span, shift int64, width int64, hasshift bool, haswidth bool) *Span {
	_cret := C.bigintspan_shift_scale(s._inner, C.int64_t(shift), C.int64_t(width), C.bool(hasshift), C.bool(haswidth))
	return &Span{_inner: _cret}
}


// BigintspansetShiftScale wraps MEOS C function bigintspanset_shift_scale.
func BigintspansetShiftScale(ss *SpanSet, shift int64, width int64, hasshift bool, haswidth bool) *SpanSet {
	_cret := C.bigintspanset_shift_scale(ss._inner, C.int64_t(shift), C.int64_t(width), C.bool(hasshift), C.bool(haswidth))
	return &SpanSet{_inner: _cret}
}


// DatesetShiftScale wraps MEOS C function dateset_shift_scale.
func DatesetShiftScale(s *Set, shift int, width int, hasshift bool, haswidth bool) *Set {
	_cret := C.dateset_shift_scale(s._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	return &Set{_inner: _cret}
}


// DatespanShiftScale wraps MEOS C function datespan_shift_scale.
func DatespanShiftScale(s *Span, shift int, width int, hasshift bool, haswidth bool) *Span {
	_cret := C.datespan_shift_scale(s._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	return &Span{_inner: _cret}
}


// DatespansetShiftScale wraps MEOS C function datespanset_shift_scale.
func DatespansetShiftScale(ss *SpanSet, shift int, width int, hasshift bool, haswidth bool) *SpanSet {
	_cret := C.datespanset_shift_scale(ss._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	return &SpanSet{_inner: _cret}
}


// FloatsetCeil wraps MEOS C function floatset_ceil.
func FloatsetCeil(s *Set) *Set {
	_cret := C.floatset_ceil(s._inner)
	return &Set{_inner: _cret}
}


// FloatsetDegrees wraps MEOS C function floatset_degrees.
func FloatsetDegrees(s *Set, normalize bool) *Set {
	_cret := C.floatset_degrees(s._inner, C.bool(normalize))
	return &Set{_inner: _cret}
}


// FloatsetFloor wraps MEOS C function floatset_floor.
func FloatsetFloor(s *Set) *Set {
	_cret := C.floatset_floor(s._inner)
	return &Set{_inner: _cret}
}


// FloatsetRadians wraps MEOS C function floatset_radians.
func FloatsetRadians(s *Set) *Set {
	_cret := C.floatset_radians(s._inner)
	return &Set{_inner: _cret}
}


// FloatsetShiftScale wraps MEOS C function floatset_shift_scale.
func FloatsetShiftScale(s *Set, shift float64, width float64, hasshift bool, haswidth bool) *Set {
	_cret := C.floatset_shift_scale(s._inner, C.double(shift), C.double(width), C.bool(hasshift), C.bool(haswidth))
	return &Set{_inner: _cret}
}


// FloatspanCeil wraps MEOS C function floatspan_ceil.
func FloatspanCeil(s *Span) *Span {
	_cret := C.floatspan_ceil(s._inner)
	return &Span{_inner: _cret}
}


// FloatspanDegrees wraps MEOS C function floatspan_degrees.
func FloatspanDegrees(s *Span, normalize bool) *Span {
	_cret := C.floatspan_degrees(s._inner, C.bool(normalize))
	return &Span{_inner: _cret}
}


// FloatspanFloor wraps MEOS C function floatspan_floor.
func FloatspanFloor(s *Span) *Span {
	_cret := C.floatspan_floor(s._inner)
	return &Span{_inner: _cret}
}


// FloatspanRadians wraps MEOS C function floatspan_radians.
func FloatspanRadians(s *Span) *Span {
	_cret := C.floatspan_radians(s._inner)
	return &Span{_inner: _cret}
}


// FloatspanRound wraps MEOS C function floatspan_round.
func FloatspanRound(s *Span, maxdd int) *Span {
	_cret := C.floatspan_round(s._inner, C.int(maxdd))
	return &Span{_inner: _cret}
}


// FloatspanShiftScale wraps MEOS C function floatspan_shift_scale.
func FloatspanShiftScale(s *Span, shift float64, width float64, hasshift bool, haswidth bool) *Span {
	_cret := C.floatspan_shift_scale(s._inner, C.double(shift), C.double(width), C.bool(hasshift), C.bool(haswidth))
	return &Span{_inner: _cret}
}


// FloatspansetCeil wraps MEOS C function floatspanset_ceil.
func FloatspansetCeil(ss *SpanSet) *SpanSet {
	_cret := C.floatspanset_ceil(ss._inner)
	return &SpanSet{_inner: _cret}
}


// FloatspansetFloor wraps MEOS C function floatspanset_floor.
func FloatspansetFloor(ss *SpanSet) *SpanSet {
	_cret := C.floatspanset_floor(ss._inner)
	return &SpanSet{_inner: _cret}
}


// FloatspansetDegrees wraps MEOS C function floatspanset_degrees.
func FloatspansetDegrees(ss *SpanSet, normalize bool) *SpanSet {
	_cret := C.floatspanset_degrees(ss._inner, C.bool(normalize))
	return &SpanSet{_inner: _cret}
}


// FloatspansetRadians wraps MEOS C function floatspanset_radians.
func FloatspansetRadians(ss *SpanSet) *SpanSet {
	_cret := C.floatspanset_radians(ss._inner)
	return &SpanSet{_inner: _cret}
}


// FloatspansetRound wraps MEOS C function floatspanset_round.
func FloatspansetRound(ss *SpanSet, maxdd int) *SpanSet {
	_cret := C.floatspanset_round(ss._inner, C.int(maxdd))
	return &SpanSet{_inner: _cret}
}


// FloatspansetShiftScale wraps MEOS C function floatspanset_shift_scale.
func FloatspansetShiftScale(ss *SpanSet, shift float64, width float64, hasshift bool, haswidth bool) *SpanSet {
	_cret := C.floatspanset_shift_scale(ss._inner, C.double(shift), C.double(width), C.bool(hasshift), C.bool(haswidth))
	return &SpanSet{_inner: _cret}
}


// IntsetShiftScale wraps MEOS C function intset_shift_scale.
func IntsetShiftScale(s *Set, shift int, width int, hasshift bool, haswidth bool) *Set {
	_cret := C.intset_shift_scale(s._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	return &Set{_inner: _cret}
}


// IntspanShiftScale wraps MEOS C function intspan_shift_scale.
func IntspanShiftScale(s *Span, shift int, width int, hasshift bool, haswidth bool) *Span {
	_cret := C.intspan_shift_scale(s._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	return &Span{_inner: _cret}
}


// IntspansetShiftScale wraps MEOS C function intspanset_shift_scale.
func IntspansetShiftScale(ss *SpanSet, shift int, width int, hasshift bool, haswidth bool) *SpanSet {
	_cret := C.intspanset_shift_scale(ss._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	return &SpanSet{_inner: _cret}
}


// TstzspanExpand wraps MEOS C function tstzspan_expand.
func TstzspanExpand(s *Span, interv *Interval) *Span {
	_cret := C.tstzspan_expand(s._inner, interv._inner)
	return &Span{_inner: _cret}
}


// SetRound wraps MEOS C function set_round.
func SetRound(s *Set, maxdd int) *Set {
	_cret := C.set_round(s._inner, C.int(maxdd))
	return &Set{_inner: _cret}
}


// TextcatTextTextset wraps MEOS C function textcat_text_textset.
func TextcatTextTextset(txt string, s *Set) *Set {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.textcat_text_textset(_c_txt, s._inner)
	return &Set{_inner: _cret}
}


// TextcatTextsetText wraps MEOS C function textcat_textset_text.
func TextcatTextsetText(s *Set, txt string) *Set {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.textcat_textset_text(s._inner, _c_txt)
	return &Set{_inner: _cret}
}


// TextsetInitcap wraps MEOS C function textset_initcap.
func TextsetInitcap(s *Set) *Set {
	_cret := C.textset_initcap(s._inner)
	return &Set{_inner: _cret}
}


// TextsetLower wraps MEOS C function textset_lower.
func TextsetLower(s *Set) *Set {
	_cret := C.textset_lower(s._inner)
	return &Set{_inner: _cret}
}


// TextsetUpper wraps MEOS C function textset_upper.
func TextsetUpper(s *Set) *Set {
	_cret := C.textset_upper(s._inner)
	return &Set{_inner: _cret}
}


// TimestamptzTprecision wraps MEOS C function timestamptz_tprecision.
func TimestamptzTprecision(t int64, duration *Interval, torigin int64) int64 {
	_cret := C.timestamptz_tprecision(C.TimestampTz(t), duration._inner, C.TimestampTz(torigin))
	return int64(_cret)
}


// TstzsetShiftScale wraps MEOS C function tstzset_shift_scale.
func TstzsetShiftScale(s *Set, shift *Interval, duration *Interval) *Set {
	_cret := C.tstzset_shift_scale(s._inner, shift._inner, duration._inner)
	return &Set{_inner: _cret}
}


// TstzsetTprecision wraps MEOS C function tstzset_tprecision.
func TstzsetTprecision(s *Set, duration *Interval, torigin int64) *Set {
	_cret := C.tstzset_tprecision(s._inner, duration._inner, C.TimestampTz(torigin))
	return &Set{_inner: _cret}
}


// TstzspanShiftScale wraps MEOS C function tstzspan_shift_scale.
func TstzspanShiftScale(s *Span, shift *Interval, duration *Interval) *Span {
	_cret := C.tstzspan_shift_scale(s._inner, shift._inner, duration._inner)
	return &Span{_inner: _cret}
}


// TstzspanTprecision wraps MEOS C function tstzspan_tprecision.
func TstzspanTprecision(s *Span, duration *Interval, torigin int64) *Span {
	_cret := C.tstzspan_tprecision(s._inner, duration._inner, C.TimestampTz(torigin))
	return &Span{_inner: _cret}
}


// TstzspansetShiftScale wraps MEOS C function tstzspanset_shift_scale.
func TstzspansetShiftScale(ss *SpanSet, shift *Interval, duration *Interval) *SpanSet {
	_cret := C.tstzspanset_shift_scale(ss._inner, shift._inner, duration._inner)
	return &SpanSet{_inner: _cret}
}


// TstzspansetTprecision wraps MEOS C function tstzspanset_tprecision.
func TstzspansetTprecision(ss *SpanSet, duration *Interval, torigin int64) *SpanSet {
	_cret := C.tstzspanset_tprecision(ss._inner, duration._inner, C.TimestampTz(torigin))
	return &SpanSet{_inner: _cret}
}


// SetCmp wraps MEOS C function set_cmp.
func SetCmp(s1 *Set, s2 *Set) int {
	_cret := C.set_cmp(s1._inner, s2._inner)
	return int(_cret)
}


// SetEq wraps MEOS C function set_eq.
func SetEq(s1 *Set, s2 *Set) bool {
	_cret := C.set_eq(s1._inner, s2._inner)
	return bool(_cret)
}


// SetGe wraps MEOS C function set_ge.
func SetGe(s1 *Set, s2 *Set) bool {
	_cret := C.set_ge(s1._inner, s2._inner)
	return bool(_cret)
}


// SetGt wraps MEOS C function set_gt.
func SetGt(s1 *Set, s2 *Set) bool {
	_cret := C.set_gt(s1._inner, s2._inner)
	return bool(_cret)
}


// SetLe wraps MEOS C function set_le.
func SetLe(s1 *Set, s2 *Set) bool {
	_cret := C.set_le(s1._inner, s2._inner)
	return bool(_cret)
}


// SetLt wraps MEOS C function set_lt.
func SetLt(s1 *Set, s2 *Set) bool {
	_cret := C.set_lt(s1._inner, s2._inner)
	return bool(_cret)
}


// SetNe wraps MEOS C function set_ne.
func SetNe(s1 *Set, s2 *Set) bool {
	_cret := C.set_ne(s1._inner, s2._inner)
	return bool(_cret)
}


// SpanCmp wraps MEOS C function span_cmp.
func SpanCmp(s1 *Span, s2 *Span) int {
	_cret := C.span_cmp(s1._inner, s2._inner)
	return int(_cret)
}


// SpanEq wraps MEOS C function span_eq.
func SpanEq(s1 *Span, s2 *Span) bool {
	_cret := C.span_eq(s1._inner, s2._inner)
	return bool(_cret)
}


// SpanGe wraps MEOS C function span_ge.
func SpanGe(s1 *Span, s2 *Span) bool {
	_cret := C.span_ge(s1._inner, s2._inner)
	return bool(_cret)
}


// SpanGt wraps MEOS C function span_gt.
func SpanGt(s1 *Span, s2 *Span) bool {
	_cret := C.span_gt(s1._inner, s2._inner)
	return bool(_cret)
}


// SpanLe wraps MEOS C function span_le.
func SpanLe(s1 *Span, s2 *Span) bool {
	_cret := C.span_le(s1._inner, s2._inner)
	return bool(_cret)
}


// SpanLt wraps MEOS C function span_lt.
func SpanLt(s1 *Span, s2 *Span) bool {
	_cret := C.span_lt(s1._inner, s2._inner)
	return bool(_cret)
}


// SpanNe wraps MEOS C function span_ne.
func SpanNe(s1 *Span, s2 *Span) bool {
	_cret := C.span_ne(s1._inner, s2._inner)
	return bool(_cret)
}


// SpansetCmp wraps MEOS C function spanset_cmp.
func SpansetCmp(ss1 *SpanSet, ss2 *SpanSet) int {
	_cret := C.spanset_cmp(ss1._inner, ss2._inner)
	return int(_cret)
}


// SpansetEq wraps MEOS C function spanset_eq.
func SpansetEq(ss1 *SpanSet, ss2 *SpanSet) bool {
	_cret := C.spanset_eq(ss1._inner, ss2._inner)
	return bool(_cret)
}


// SpansetGe wraps MEOS C function spanset_ge.
func SpansetGe(ss1 *SpanSet, ss2 *SpanSet) bool {
	_cret := C.spanset_ge(ss1._inner, ss2._inner)
	return bool(_cret)
}


// SpansetGt wraps MEOS C function spanset_gt.
func SpansetGt(ss1 *SpanSet, ss2 *SpanSet) bool {
	_cret := C.spanset_gt(ss1._inner, ss2._inner)
	return bool(_cret)
}


// SpansetLe wraps MEOS C function spanset_le.
func SpansetLe(ss1 *SpanSet, ss2 *SpanSet) bool {
	_cret := C.spanset_le(ss1._inner, ss2._inner)
	return bool(_cret)
}


// SpansetLt wraps MEOS C function spanset_lt.
func SpansetLt(ss1 *SpanSet, ss2 *SpanSet) bool {
	_cret := C.spanset_lt(ss1._inner, ss2._inner)
	return bool(_cret)
}


// SpansetNe wraps MEOS C function spanset_ne.
func SpansetNe(ss1 *SpanSet, ss2 *SpanSet) bool {
	_cret := C.spanset_ne(ss1._inner, ss2._inner)
	return bool(_cret)
}


// SetSpans wraps MEOS C function set_spans.
func SetSpans(s *Set, count unsafe.Pointer) *Span {
	_cret := C.set_spans(s._inner, (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// SetSplitEachNSpans wraps MEOS C function set_split_each_n_spans.
func SetSplitEachNSpans(s *Set, elems_per_span int, count unsafe.Pointer) *Span {
	_cret := C.set_split_each_n_spans(s._inner, C.int(elems_per_span), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// SetSplitNSpans wraps MEOS C function set_split_n_spans.
func SetSplitNSpans(s *Set, span_count int, count unsafe.Pointer) *Span {
	_cret := C.set_split_n_spans(s._inner, C.int(span_count), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// SpansetSpans wraps MEOS C function spanset_spans.
func SpansetSpans(ss *SpanSet, count unsafe.Pointer) *Span {
	_cret := C.spanset_spans(ss._inner, (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// SpansetSplitEachNSpans wraps MEOS C function spanset_split_each_n_spans.
func SpansetSplitEachNSpans(ss *SpanSet, elems_per_span int, count unsafe.Pointer) *Span {
	_cret := C.spanset_split_each_n_spans(ss._inner, C.int(elems_per_span), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// SpansetSplitNSpans wraps MEOS C function spanset_split_n_spans.
func SpansetSplitNSpans(ss *SpanSet, span_count int, count unsafe.Pointer) *Span {
	_cret := C.spanset_split_n_spans(ss._inner, C.int(span_count), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// AdjacentSpanBigint wraps MEOS C function adjacent_span_bigint.
func AdjacentSpanBigint(s *Span, i int64) bool {
	_cret := C.adjacent_span_bigint(s._inner, C.int64_t(i))
	return bool(_cret)
}


// AdjacentSpanDate wraps MEOS C function adjacent_span_date.
func AdjacentSpanDate(s *Span, d int32) bool {
	_cret := C.adjacent_span_date(s._inner, C.DateADT(d))
	return bool(_cret)
}


// AdjacentSpanFloat wraps MEOS C function adjacent_span_float.
func AdjacentSpanFloat(s *Span, d float64) bool {
	_cret := C.adjacent_span_float(s._inner, C.double(d))
	return bool(_cret)
}


// AdjacentSpanInt wraps MEOS C function adjacent_span_int.
func AdjacentSpanInt(s *Span, i int) bool {
	_cret := C.adjacent_span_int(s._inner, C.int(i))
	return bool(_cret)
}


// AdjacentSpanSpan wraps MEOS C function adjacent_span_span.
func AdjacentSpanSpan(s1 *Span, s2 *Span) bool {
	_cret := C.adjacent_span_span(s1._inner, s2._inner)
	return bool(_cret)
}


// AdjacentSpanSpanset wraps MEOS C function adjacent_span_spanset.
func AdjacentSpanSpanset(s *Span, ss *SpanSet) bool {
	_cret := C.adjacent_span_spanset(s._inner, ss._inner)
	return bool(_cret)
}


// AdjacentSpanTimestamptz wraps MEOS C function adjacent_span_timestamptz.
func AdjacentSpanTimestamptz(s *Span, t int64) bool {
	_cret := C.adjacent_span_timestamptz(s._inner, C.TimestampTz(t))
	return bool(_cret)
}


// AdjacentSpansetBigint wraps MEOS C function adjacent_spanset_bigint.
func AdjacentSpansetBigint(ss *SpanSet, i int64) bool {
	_cret := C.adjacent_spanset_bigint(ss._inner, C.int64_t(i))
	return bool(_cret)
}


// AdjacentSpansetDate wraps MEOS C function adjacent_spanset_date.
func AdjacentSpansetDate(ss *SpanSet, d int32) bool {
	_cret := C.adjacent_spanset_date(ss._inner, C.DateADT(d))
	return bool(_cret)
}


// AdjacentSpansetFloat wraps MEOS C function adjacent_spanset_float.
func AdjacentSpansetFloat(ss *SpanSet, d float64) bool {
	_cret := C.adjacent_spanset_float(ss._inner, C.double(d))
	return bool(_cret)
}


// AdjacentSpansetInt wraps MEOS C function adjacent_spanset_int.
func AdjacentSpansetInt(ss *SpanSet, i int) bool {
	_cret := C.adjacent_spanset_int(ss._inner, C.int(i))
	return bool(_cret)
}


// AdjacentSpansetTimestamptz wraps MEOS C function adjacent_spanset_timestamptz.
func AdjacentSpansetTimestamptz(ss *SpanSet, t int64) bool {
	_cret := C.adjacent_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return bool(_cret)
}


// AdjacentSpansetSpan wraps MEOS C function adjacent_spanset_span.
func AdjacentSpansetSpan(ss *SpanSet, s *Span) bool {
	_cret := C.adjacent_spanset_span(ss._inner, s._inner)
	return bool(_cret)
}


// AdjacentSpansetSpanset wraps MEOS C function adjacent_spanset_spanset.
func AdjacentSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) bool {
	_cret := C.adjacent_spanset_spanset(ss1._inner, ss2._inner)
	return bool(_cret)
}


// ContainedBigintSet wraps MEOS C function contained_bigint_set.
func ContainedBigintSet(i int64, s *Set) bool {
	_cret := C.contained_bigint_set(C.int64_t(i), s._inner)
	return bool(_cret)
}


// ContainedBigintSpan wraps MEOS C function contained_bigint_span.
func ContainedBigintSpan(i int64, s *Span) bool {
	_cret := C.contained_bigint_span(C.int64_t(i), s._inner)
	return bool(_cret)
}


// ContainedBigintSpanset wraps MEOS C function contained_bigint_spanset.
func ContainedBigintSpanset(i int64, ss *SpanSet) bool {
	_cret := C.contained_bigint_spanset(C.int64_t(i), ss._inner)
	return bool(_cret)
}


// ContainedDateSet wraps MEOS C function contained_date_set.
func ContainedDateSet(d int32, s *Set) bool {
	_cret := C.contained_date_set(C.DateADT(d), s._inner)
	return bool(_cret)
}


// ContainedDateSpan wraps MEOS C function contained_date_span.
func ContainedDateSpan(d int32, s *Span) bool {
	_cret := C.contained_date_span(C.DateADT(d), s._inner)
	return bool(_cret)
}


// ContainedDateSpanset wraps MEOS C function contained_date_spanset.
func ContainedDateSpanset(d int32, ss *SpanSet) bool {
	_cret := C.contained_date_spanset(C.DateADT(d), ss._inner)
	return bool(_cret)
}


// ContainedFloatSet wraps MEOS C function contained_float_set.
func ContainedFloatSet(d float64, s *Set) bool {
	_cret := C.contained_float_set(C.double(d), s._inner)
	return bool(_cret)
}


// ContainedFloatSpan wraps MEOS C function contained_float_span.
func ContainedFloatSpan(d float64, s *Span) bool {
	_cret := C.contained_float_span(C.double(d), s._inner)
	return bool(_cret)
}


// ContainedFloatSpanset wraps MEOS C function contained_float_spanset.
func ContainedFloatSpanset(d float64, ss *SpanSet) bool {
	_cret := C.contained_float_spanset(C.double(d), ss._inner)
	return bool(_cret)
}


// ContainedIntSet wraps MEOS C function contained_int_set.
func ContainedIntSet(i int, s *Set) bool {
	_cret := C.contained_int_set(C.int(i), s._inner)
	return bool(_cret)
}


// ContainedIntSpan wraps MEOS C function contained_int_span.
func ContainedIntSpan(i int, s *Span) bool {
	_cret := C.contained_int_span(C.int(i), s._inner)
	return bool(_cret)
}


// ContainedIntSpanset wraps MEOS C function contained_int_spanset.
func ContainedIntSpanset(i int, ss *SpanSet) bool {
	_cret := C.contained_int_spanset(C.int(i), ss._inner)
	return bool(_cret)
}


// ContainedSetSet wraps MEOS C function contained_set_set.
func ContainedSetSet(s1 *Set, s2 *Set) bool {
	_cret := C.contained_set_set(s1._inner, s2._inner)
	return bool(_cret)
}


// ContainedSpanSpan wraps MEOS C function contained_span_span.
func ContainedSpanSpan(s1 *Span, s2 *Span) bool {
	_cret := C.contained_span_span(s1._inner, s2._inner)
	return bool(_cret)
}


// ContainedSpanSpanset wraps MEOS C function contained_span_spanset.
func ContainedSpanSpanset(s *Span, ss *SpanSet) bool {
	_cret := C.contained_span_spanset(s._inner, ss._inner)
	return bool(_cret)
}


// ContainedSpansetSpan wraps MEOS C function contained_spanset_span.
func ContainedSpansetSpan(ss *SpanSet, s *Span) bool {
	_cret := C.contained_spanset_span(ss._inner, s._inner)
	return bool(_cret)
}


// ContainedSpansetSpanset wraps MEOS C function contained_spanset_spanset.
func ContainedSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) bool {
	_cret := C.contained_spanset_spanset(ss1._inner, ss2._inner)
	return bool(_cret)
}


// ContainedTextSet wraps MEOS C function contained_text_set.
func ContainedTextSet(txt string, s *Set) bool {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.contained_text_set(_c_txt, s._inner)
	return bool(_cret)
}


// ContainedTimestamptzSet wraps MEOS C function contained_timestamptz_set.
func ContainedTimestamptzSet(t int64, s *Set) bool {
	_cret := C.contained_timestamptz_set(C.TimestampTz(t), s._inner)
	return bool(_cret)
}


// ContainedTimestamptzSpan wraps MEOS C function contained_timestamptz_span.
func ContainedTimestamptzSpan(t int64, s *Span) bool {
	_cret := C.contained_timestamptz_span(C.TimestampTz(t), s._inner)
	return bool(_cret)
}


// ContainedTimestamptzSpanset wraps MEOS C function contained_timestamptz_spanset.
func ContainedTimestamptzSpanset(t int64, ss *SpanSet) bool {
	_cret := C.contained_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	return bool(_cret)
}


// ContainsSetBigint wraps MEOS C function contains_set_bigint.
func ContainsSetBigint(s *Set, i int64) bool {
	_cret := C.contains_set_bigint(s._inner, C.int64_t(i))
	return bool(_cret)
}


// ContainsSetDate wraps MEOS C function contains_set_date.
func ContainsSetDate(s *Set, d int32) bool {
	_cret := C.contains_set_date(s._inner, C.DateADT(d))
	return bool(_cret)
}


// ContainsSetFloat wraps MEOS C function contains_set_float.
func ContainsSetFloat(s *Set, d float64) bool {
	_cret := C.contains_set_float(s._inner, C.double(d))
	return bool(_cret)
}


// ContainsSetInt wraps MEOS C function contains_set_int.
func ContainsSetInt(s *Set, i int) bool {
	_cret := C.contains_set_int(s._inner, C.int(i))
	return bool(_cret)
}


// ContainsSetSet wraps MEOS C function contains_set_set.
func ContainsSetSet(s1 *Set, s2 *Set) bool {
	_cret := C.contains_set_set(s1._inner, s2._inner)
	return bool(_cret)
}


// ContainsSetText wraps MEOS C function contains_set_text.
func ContainsSetText(s *Set, t string) bool {
	_c_t := C.cstring_to_text(C.CString(t))
	defer C.free(unsafe.Pointer(_c_t))
	_cret := C.contains_set_text(s._inner, _c_t)
	return bool(_cret)
}


// ContainsSetTimestamptz wraps MEOS C function contains_set_timestamptz.
func ContainsSetTimestamptz(s *Set, t int64) bool {
	_cret := C.contains_set_timestamptz(s._inner, C.TimestampTz(t))
	return bool(_cret)
}


// ContainsSpanBigint wraps MEOS C function contains_span_bigint.
func ContainsSpanBigint(s *Span, i int64) bool {
	_cret := C.contains_span_bigint(s._inner, C.int64_t(i))
	return bool(_cret)
}


// ContainsSpanDate wraps MEOS C function contains_span_date.
func ContainsSpanDate(s *Span, d int32) bool {
	_cret := C.contains_span_date(s._inner, C.DateADT(d))
	return bool(_cret)
}


// ContainsSpanFloat wraps MEOS C function contains_span_float.
func ContainsSpanFloat(s *Span, d float64) bool {
	_cret := C.contains_span_float(s._inner, C.double(d))
	return bool(_cret)
}


// ContainsSpanInt wraps MEOS C function contains_span_int.
func ContainsSpanInt(s *Span, i int) bool {
	_cret := C.contains_span_int(s._inner, C.int(i))
	return bool(_cret)
}


// ContainsSpanSpan wraps MEOS C function contains_span_span.
func ContainsSpanSpan(s1 *Span, s2 *Span) bool {
	_cret := C.contains_span_span(s1._inner, s2._inner)
	return bool(_cret)
}


// ContainsSpanSpanset wraps MEOS C function contains_span_spanset.
func ContainsSpanSpanset(s *Span, ss *SpanSet) bool {
	_cret := C.contains_span_spanset(s._inner, ss._inner)
	return bool(_cret)
}


// ContainsSpanTimestamptz wraps MEOS C function contains_span_timestamptz.
func ContainsSpanTimestamptz(s *Span, t int64) bool {
	_cret := C.contains_span_timestamptz(s._inner, C.TimestampTz(t))
	return bool(_cret)
}


// ContainsSpansetBigint wraps MEOS C function contains_spanset_bigint.
func ContainsSpansetBigint(ss *SpanSet, i int64) bool {
	_cret := C.contains_spanset_bigint(ss._inner, C.int64_t(i))
	return bool(_cret)
}


// ContainsSpansetDate wraps MEOS C function contains_spanset_date.
func ContainsSpansetDate(ss *SpanSet, d int32) bool {
	_cret := C.contains_spanset_date(ss._inner, C.DateADT(d))
	return bool(_cret)
}


// ContainsSpansetFloat wraps MEOS C function contains_spanset_float.
func ContainsSpansetFloat(ss *SpanSet, d float64) bool {
	_cret := C.contains_spanset_float(ss._inner, C.double(d))
	return bool(_cret)
}


// ContainsSpansetInt wraps MEOS C function contains_spanset_int.
func ContainsSpansetInt(ss *SpanSet, i int) bool {
	_cret := C.contains_spanset_int(ss._inner, C.int(i))
	return bool(_cret)
}


// ContainsSpansetSpan wraps MEOS C function contains_spanset_span.
func ContainsSpansetSpan(ss *SpanSet, s *Span) bool {
	_cret := C.contains_spanset_span(ss._inner, s._inner)
	return bool(_cret)
}


// ContainsSpansetSpanset wraps MEOS C function contains_spanset_spanset.
func ContainsSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) bool {
	_cret := C.contains_spanset_spanset(ss1._inner, ss2._inner)
	return bool(_cret)
}


// ContainsSpansetTimestamptz wraps MEOS C function contains_spanset_timestamptz.
func ContainsSpansetTimestamptz(ss *SpanSet, t int64) bool {
	_cret := C.contains_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return bool(_cret)
}


// OverlapsSetSet wraps MEOS C function overlaps_set_set.
func OverlapsSetSet(s1 *Set, s2 *Set) bool {
	_cret := C.overlaps_set_set(s1._inner, s2._inner)
	return bool(_cret)
}


// OverlapsSpanSpan wraps MEOS C function overlaps_span_span.
func OverlapsSpanSpan(s1 *Span, s2 *Span) bool {
	_cret := C.overlaps_span_span(s1._inner, s2._inner)
	return bool(_cret)
}


// OverlapsSpanSpanset wraps MEOS C function overlaps_span_spanset.
func OverlapsSpanSpanset(s *Span, ss *SpanSet) bool {
	_cret := C.overlaps_span_spanset(s._inner, ss._inner)
	return bool(_cret)
}


// OverlapsSpansetSpan wraps MEOS C function overlaps_spanset_span.
func OverlapsSpansetSpan(ss *SpanSet, s *Span) bool {
	_cret := C.overlaps_spanset_span(ss._inner, s._inner)
	return bool(_cret)
}


// OverlapsSpansetSpanset wraps MEOS C function overlaps_spanset_spanset.
func OverlapsSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) bool {
	_cret := C.overlaps_spanset_spanset(ss1._inner, ss2._inner)
	return bool(_cret)
}


// SameSpanSpan wraps MEOS C function same_span_span.
func SameSpanSpan(s1 *Span, s2 *Span) bool {
	_cret := C.same_span_span(s1._inner, s2._inner)
	return bool(_cret)
}


// AfterDateSet wraps MEOS C function after_date_set.
func AfterDateSet(d int32, s *Set) bool {
	_cret := C.after_date_set(C.DateADT(d), s._inner)
	return bool(_cret)
}


// AfterDateSpan wraps MEOS C function after_date_span.
func AfterDateSpan(d int32, s *Span) bool {
	_cret := C.after_date_span(C.DateADT(d), s._inner)
	return bool(_cret)
}


// AfterDateSpanset wraps MEOS C function after_date_spanset.
func AfterDateSpanset(d int32, ss *SpanSet) bool {
	_cret := C.after_date_spanset(C.DateADT(d), ss._inner)
	return bool(_cret)
}


// AfterSetDate wraps MEOS C function after_set_date.
func AfterSetDate(s *Set, d int32) bool {
	_cret := C.after_set_date(s._inner, C.DateADT(d))
	return bool(_cret)
}


// AfterSetTimestamptz wraps MEOS C function after_set_timestamptz.
func AfterSetTimestamptz(s *Set, t int64) bool {
	_cret := C.after_set_timestamptz(s._inner, C.TimestampTz(t))
	return bool(_cret)
}


// AfterSpanDate wraps MEOS C function after_span_date.
func AfterSpanDate(s *Span, d int32) bool {
	_cret := C.after_span_date(s._inner, C.DateADT(d))
	return bool(_cret)
}


// AfterSpanTimestamptz wraps MEOS C function after_span_timestamptz.
func AfterSpanTimestamptz(s *Span, t int64) bool {
	_cret := C.after_span_timestamptz(s._inner, C.TimestampTz(t))
	return bool(_cret)
}


// AfterSpansetDate wraps MEOS C function after_spanset_date.
func AfterSpansetDate(ss *SpanSet, d int32) bool {
	_cret := C.after_spanset_date(ss._inner, C.DateADT(d))
	return bool(_cret)
}


// AfterSpansetTimestamptz wraps MEOS C function after_spanset_timestamptz.
func AfterSpansetTimestamptz(ss *SpanSet, t int64) bool {
	_cret := C.after_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return bool(_cret)
}


// AfterTimestamptzSet wraps MEOS C function after_timestamptz_set.
func AfterTimestamptzSet(t int64, s *Set) bool {
	_cret := C.after_timestamptz_set(C.TimestampTz(t), s._inner)
	return bool(_cret)
}


// AfterTimestamptzSpan wraps MEOS C function after_timestamptz_span.
func AfterTimestamptzSpan(t int64, s *Span) bool {
	_cret := C.after_timestamptz_span(C.TimestampTz(t), s._inner)
	return bool(_cret)
}


// AfterTimestamptzSpanset wraps MEOS C function after_timestamptz_spanset.
func AfterTimestamptzSpanset(t int64, ss *SpanSet) bool {
	_cret := C.after_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	return bool(_cret)
}


// BeforeDateSet wraps MEOS C function before_date_set.
func BeforeDateSet(d int32, s *Set) bool {
	_cret := C.before_date_set(C.DateADT(d), s._inner)
	return bool(_cret)
}


// BeforeDateSpan wraps MEOS C function before_date_span.
func BeforeDateSpan(d int32, s *Span) bool {
	_cret := C.before_date_span(C.DateADT(d), s._inner)
	return bool(_cret)
}


// BeforeDateSpanset wraps MEOS C function before_date_spanset.
func BeforeDateSpanset(d int32, ss *SpanSet) bool {
	_cret := C.before_date_spanset(C.DateADT(d), ss._inner)
	return bool(_cret)
}


// BeforeSetDate wraps MEOS C function before_set_date.
func BeforeSetDate(s *Set, d int32) bool {
	_cret := C.before_set_date(s._inner, C.DateADT(d))
	return bool(_cret)
}


// BeforeSetTimestamptz wraps MEOS C function before_set_timestamptz.
func BeforeSetTimestamptz(s *Set, t int64) bool {
	_cret := C.before_set_timestamptz(s._inner, C.TimestampTz(t))
	return bool(_cret)
}


// BeforeSpanDate wraps MEOS C function before_span_date.
func BeforeSpanDate(s *Span, d int32) bool {
	_cret := C.before_span_date(s._inner, C.DateADT(d))
	return bool(_cret)
}


// BeforeSpanTimestamptz wraps MEOS C function before_span_timestamptz.
func BeforeSpanTimestamptz(s *Span, t int64) bool {
	_cret := C.before_span_timestamptz(s._inner, C.TimestampTz(t))
	return bool(_cret)
}


// BeforeSpansetDate wraps MEOS C function before_spanset_date.
func BeforeSpansetDate(ss *SpanSet, d int32) bool {
	_cret := C.before_spanset_date(ss._inner, C.DateADT(d))
	return bool(_cret)
}


// BeforeSpansetTimestamptz wraps MEOS C function before_spanset_timestamptz.
func BeforeSpansetTimestamptz(ss *SpanSet, t int64) bool {
	_cret := C.before_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return bool(_cret)
}


// BeforeTimestamptzSet wraps MEOS C function before_timestamptz_set.
func BeforeTimestamptzSet(t int64, s *Set) bool {
	_cret := C.before_timestamptz_set(C.TimestampTz(t), s._inner)
	return bool(_cret)
}


// BeforeTimestamptzSpan wraps MEOS C function before_timestamptz_span.
func BeforeTimestamptzSpan(t int64, s *Span) bool {
	_cret := C.before_timestamptz_span(C.TimestampTz(t), s._inner)
	return bool(_cret)
}


// BeforeTimestamptzSpanset wraps MEOS C function before_timestamptz_spanset.
func BeforeTimestamptzSpanset(t int64, ss *SpanSet) bool {
	_cret := C.before_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	return bool(_cret)
}


// LeftBigintSet wraps MEOS C function left_bigint_set.
func LeftBigintSet(i int64, s *Set) bool {
	_cret := C.left_bigint_set(C.int64_t(i), s._inner)
	return bool(_cret)
}


// LeftBigintSpan wraps MEOS C function left_bigint_span.
func LeftBigintSpan(i int64, s *Span) bool {
	_cret := C.left_bigint_span(C.int64_t(i), s._inner)
	return bool(_cret)
}


// LeftBigintSpanset wraps MEOS C function left_bigint_spanset.
func LeftBigintSpanset(i int64, ss *SpanSet) bool {
	_cret := C.left_bigint_spanset(C.int64_t(i), ss._inner)
	return bool(_cret)
}


// LeftFloatSet wraps MEOS C function left_float_set.
func LeftFloatSet(d float64, s *Set) bool {
	_cret := C.left_float_set(C.double(d), s._inner)
	return bool(_cret)
}


// LeftFloatSpan wraps MEOS C function left_float_span.
func LeftFloatSpan(d float64, s *Span) bool {
	_cret := C.left_float_span(C.double(d), s._inner)
	return bool(_cret)
}


// LeftFloatSpanset wraps MEOS C function left_float_spanset.
func LeftFloatSpanset(d float64, ss *SpanSet) bool {
	_cret := C.left_float_spanset(C.double(d), ss._inner)
	return bool(_cret)
}


// LeftIntSet wraps MEOS C function left_int_set.
func LeftIntSet(i int, s *Set) bool {
	_cret := C.left_int_set(C.int(i), s._inner)
	return bool(_cret)
}


// LeftIntSpan wraps MEOS C function left_int_span.
func LeftIntSpan(i int, s *Span) bool {
	_cret := C.left_int_span(C.int(i), s._inner)
	return bool(_cret)
}


// LeftIntSpanset wraps MEOS C function left_int_spanset.
func LeftIntSpanset(i int, ss *SpanSet) bool {
	_cret := C.left_int_spanset(C.int(i), ss._inner)
	return bool(_cret)
}


// LeftSetBigint wraps MEOS C function left_set_bigint.
func LeftSetBigint(s *Set, i int64) bool {
	_cret := C.left_set_bigint(s._inner, C.int64_t(i))
	return bool(_cret)
}


// LeftSetFloat wraps MEOS C function left_set_float.
func LeftSetFloat(s *Set, d float64) bool {
	_cret := C.left_set_float(s._inner, C.double(d))
	return bool(_cret)
}


// LeftSetInt wraps MEOS C function left_set_int.
func LeftSetInt(s *Set, i int) bool {
	_cret := C.left_set_int(s._inner, C.int(i))
	return bool(_cret)
}


// LeftSetSet wraps MEOS C function left_set_set.
func LeftSetSet(s1 *Set, s2 *Set) bool {
	_cret := C.left_set_set(s1._inner, s2._inner)
	return bool(_cret)
}


// LeftSetText wraps MEOS C function left_set_text.
func LeftSetText(s *Set, txt string) bool {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.left_set_text(s._inner, _c_txt)
	return bool(_cret)
}


// LeftSpanBigint wraps MEOS C function left_span_bigint.
func LeftSpanBigint(s *Span, i int64) bool {
	_cret := C.left_span_bigint(s._inner, C.int64_t(i))
	return bool(_cret)
}


// LeftSpanFloat wraps MEOS C function left_span_float.
func LeftSpanFloat(s *Span, d float64) bool {
	_cret := C.left_span_float(s._inner, C.double(d))
	return bool(_cret)
}


// LeftSpanInt wraps MEOS C function left_span_int.
func LeftSpanInt(s *Span, i int) bool {
	_cret := C.left_span_int(s._inner, C.int(i))
	return bool(_cret)
}


// LeftSpanSpan wraps MEOS C function left_span_span.
func LeftSpanSpan(s1 *Span, s2 *Span) bool {
	_cret := C.left_span_span(s1._inner, s2._inner)
	return bool(_cret)
}


// LeftSpanSpanset wraps MEOS C function left_span_spanset.
func LeftSpanSpanset(s *Span, ss *SpanSet) bool {
	_cret := C.left_span_spanset(s._inner, ss._inner)
	return bool(_cret)
}


// LeftSpansetBigint wraps MEOS C function left_spanset_bigint.
func LeftSpansetBigint(ss *SpanSet, i int64) bool {
	_cret := C.left_spanset_bigint(ss._inner, C.int64_t(i))
	return bool(_cret)
}


// LeftSpansetFloat wraps MEOS C function left_spanset_float.
func LeftSpansetFloat(ss *SpanSet, d float64) bool {
	_cret := C.left_spanset_float(ss._inner, C.double(d))
	return bool(_cret)
}


// LeftSpansetInt wraps MEOS C function left_spanset_int.
func LeftSpansetInt(ss *SpanSet, i int) bool {
	_cret := C.left_spanset_int(ss._inner, C.int(i))
	return bool(_cret)
}


// LeftSpansetSpan wraps MEOS C function left_spanset_span.
func LeftSpansetSpan(ss *SpanSet, s *Span) bool {
	_cret := C.left_spanset_span(ss._inner, s._inner)
	return bool(_cret)
}


// LeftSpansetSpanset wraps MEOS C function left_spanset_spanset.
func LeftSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) bool {
	_cret := C.left_spanset_spanset(ss1._inner, ss2._inner)
	return bool(_cret)
}


// LeftTextSet wraps MEOS C function left_text_set.
func LeftTextSet(txt string, s *Set) bool {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.left_text_set(_c_txt, s._inner)
	return bool(_cret)
}


// OverafterDateSet wraps MEOS C function overafter_date_set.
func OverafterDateSet(d int32, s *Set) bool {
	_cret := C.overafter_date_set(C.DateADT(d), s._inner)
	return bool(_cret)
}


// OverafterDateSpan wraps MEOS C function overafter_date_span.
func OverafterDateSpan(d int32, s *Span) bool {
	_cret := C.overafter_date_span(C.DateADT(d), s._inner)
	return bool(_cret)
}


// OverafterDateSpanset wraps MEOS C function overafter_date_spanset.
func OverafterDateSpanset(d int32, ss *SpanSet) bool {
	_cret := C.overafter_date_spanset(C.DateADT(d), ss._inner)
	return bool(_cret)
}


// OverafterSetDate wraps MEOS C function overafter_set_date.
func OverafterSetDate(s *Set, d int32) bool {
	_cret := C.overafter_set_date(s._inner, C.DateADT(d))
	return bool(_cret)
}


// OverafterSetTimestamptz wraps MEOS C function overafter_set_timestamptz.
func OverafterSetTimestamptz(s *Set, t int64) bool {
	_cret := C.overafter_set_timestamptz(s._inner, C.TimestampTz(t))
	return bool(_cret)
}


// OverafterSpanDate wraps MEOS C function overafter_span_date.
func OverafterSpanDate(s *Span, d int32) bool {
	_cret := C.overafter_span_date(s._inner, C.DateADT(d))
	return bool(_cret)
}


// OverafterSpanTimestamptz wraps MEOS C function overafter_span_timestamptz.
func OverafterSpanTimestamptz(s *Span, t int64) bool {
	_cret := C.overafter_span_timestamptz(s._inner, C.TimestampTz(t))
	return bool(_cret)
}


// OverafterSpansetDate wraps MEOS C function overafter_spanset_date.
func OverafterSpansetDate(ss *SpanSet, d int32) bool {
	_cret := C.overafter_spanset_date(ss._inner, C.DateADT(d))
	return bool(_cret)
}


// OverafterSpansetTimestamptz wraps MEOS C function overafter_spanset_timestamptz.
func OverafterSpansetTimestamptz(ss *SpanSet, t int64) bool {
	_cret := C.overafter_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return bool(_cret)
}


// OverafterTimestamptzSet wraps MEOS C function overafter_timestamptz_set.
func OverafterTimestamptzSet(t int64, s *Set) bool {
	_cret := C.overafter_timestamptz_set(C.TimestampTz(t), s._inner)
	return bool(_cret)
}


// OverafterTimestamptzSpan wraps MEOS C function overafter_timestamptz_span.
func OverafterTimestamptzSpan(t int64, s *Span) bool {
	_cret := C.overafter_timestamptz_span(C.TimestampTz(t), s._inner)
	return bool(_cret)
}


// OverafterTimestamptzSpanset wraps MEOS C function overafter_timestamptz_spanset.
func OverafterTimestamptzSpanset(t int64, ss *SpanSet) bool {
	_cret := C.overafter_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	return bool(_cret)
}


// OverbeforeDateSet wraps MEOS C function overbefore_date_set.
func OverbeforeDateSet(d int32, s *Set) bool {
	_cret := C.overbefore_date_set(C.DateADT(d), s._inner)
	return bool(_cret)
}


// OverbeforeDateSpan wraps MEOS C function overbefore_date_span.
func OverbeforeDateSpan(d int32, s *Span) bool {
	_cret := C.overbefore_date_span(C.DateADT(d), s._inner)
	return bool(_cret)
}


// OverbeforeDateSpanset wraps MEOS C function overbefore_date_spanset.
func OverbeforeDateSpanset(d int32, ss *SpanSet) bool {
	_cret := C.overbefore_date_spanset(C.DateADT(d), ss._inner)
	return bool(_cret)
}


// OverbeforeSetDate wraps MEOS C function overbefore_set_date.
func OverbeforeSetDate(s *Set, d int32) bool {
	_cret := C.overbefore_set_date(s._inner, C.DateADT(d))
	return bool(_cret)
}


// OverbeforeSetTimestamptz wraps MEOS C function overbefore_set_timestamptz.
func OverbeforeSetTimestamptz(s *Set, t int64) bool {
	_cret := C.overbefore_set_timestamptz(s._inner, C.TimestampTz(t))
	return bool(_cret)
}


// OverbeforeSpanDate wraps MEOS C function overbefore_span_date.
func OverbeforeSpanDate(s *Span, d int32) bool {
	_cret := C.overbefore_span_date(s._inner, C.DateADT(d))
	return bool(_cret)
}


// OverbeforeSpanTimestamptz wraps MEOS C function overbefore_span_timestamptz.
func OverbeforeSpanTimestamptz(s *Span, t int64) bool {
	_cret := C.overbefore_span_timestamptz(s._inner, C.TimestampTz(t))
	return bool(_cret)
}


// OverbeforeSpansetDate wraps MEOS C function overbefore_spanset_date.
func OverbeforeSpansetDate(ss *SpanSet, d int32) bool {
	_cret := C.overbefore_spanset_date(ss._inner, C.DateADT(d))
	return bool(_cret)
}


// OverbeforeSpansetTimestamptz wraps MEOS C function overbefore_spanset_timestamptz.
func OverbeforeSpansetTimestamptz(ss *SpanSet, t int64) bool {
	_cret := C.overbefore_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return bool(_cret)
}


// OverbeforeTimestamptzSet wraps MEOS C function overbefore_timestamptz_set.
func OverbeforeTimestamptzSet(t int64, s *Set) bool {
	_cret := C.overbefore_timestamptz_set(C.TimestampTz(t), s._inner)
	return bool(_cret)
}


// OverbeforeTimestamptzSpan wraps MEOS C function overbefore_timestamptz_span.
func OverbeforeTimestamptzSpan(t int64, s *Span) bool {
	_cret := C.overbefore_timestamptz_span(C.TimestampTz(t), s._inner)
	return bool(_cret)
}


// OverbeforeTimestamptzSpanset wraps MEOS C function overbefore_timestamptz_spanset.
func OverbeforeTimestamptzSpanset(t int64, ss *SpanSet) bool {
	_cret := C.overbefore_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	return bool(_cret)
}


// OverleftBigintSet wraps MEOS C function overleft_bigint_set.
func OverleftBigintSet(i int64, s *Set) bool {
	_cret := C.overleft_bigint_set(C.int64_t(i), s._inner)
	return bool(_cret)
}


// OverleftBigintSpan wraps MEOS C function overleft_bigint_span.
func OverleftBigintSpan(i int64, s *Span) bool {
	_cret := C.overleft_bigint_span(C.int64_t(i), s._inner)
	return bool(_cret)
}


// OverleftBigintSpanset wraps MEOS C function overleft_bigint_spanset.
func OverleftBigintSpanset(i int64, ss *SpanSet) bool {
	_cret := C.overleft_bigint_spanset(C.int64_t(i), ss._inner)
	return bool(_cret)
}


// OverleftFloatSet wraps MEOS C function overleft_float_set.
func OverleftFloatSet(d float64, s *Set) bool {
	_cret := C.overleft_float_set(C.double(d), s._inner)
	return bool(_cret)
}


// OverleftFloatSpan wraps MEOS C function overleft_float_span.
func OverleftFloatSpan(d float64, s *Span) bool {
	_cret := C.overleft_float_span(C.double(d), s._inner)
	return bool(_cret)
}


// OverleftFloatSpanset wraps MEOS C function overleft_float_spanset.
func OverleftFloatSpanset(d float64, ss *SpanSet) bool {
	_cret := C.overleft_float_spanset(C.double(d), ss._inner)
	return bool(_cret)
}


// OverleftIntSet wraps MEOS C function overleft_int_set.
func OverleftIntSet(i int, s *Set) bool {
	_cret := C.overleft_int_set(C.int(i), s._inner)
	return bool(_cret)
}


// OverleftIntSpan wraps MEOS C function overleft_int_span.
func OverleftIntSpan(i int, s *Span) bool {
	_cret := C.overleft_int_span(C.int(i), s._inner)
	return bool(_cret)
}


// OverleftIntSpanset wraps MEOS C function overleft_int_spanset.
func OverleftIntSpanset(i int, ss *SpanSet) bool {
	_cret := C.overleft_int_spanset(C.int(i), ss._inner)
	return bool(_cret)
}


// OverleftSetBigint wraps MEOS C function overleft_set_bigint.
func OverleftSetBigint(s *Set, i int64) bool {
	_cret := C.overleft_set_bigint(s._inner, C.int64_t(i))
	return bool(_cret)
}


// OverleftSetFloat wraps MEOS C function overleft_set_float.
func OverleftSetFloat(s *Set, d float64) bool {
	_cret := C.overleft_set_float(s._inner, C.double(d))
	return bool(_cret)
}


// OverleftSetInt wraps MEOS C function overleft_set_int.
func OverleftSetInt(s *Set, i int) bool {
	_cret := C.overleft_set_int(s._inner, C.int(i))
	return bool(_cret)
}


// OverleftSetSet wraps MEOS C function overleft_set_set.
func OverleftSetSet(s1 *Set, s2 *Set) bool {
	_cret := C.overleft_set_set(s1._inner, s2._inner)
	return bool(_cret)
}


// OverleftSetText wraps MEOS C function overleft_set_text.
func OverleftSetText(s *Set, txt string) bool {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.overleft_set_text(s._inner, _c_txt)
	return bool(_cret)
}


// OverleftSpanBigint wraps MEOS C function overleft_span_bigint.
func OverleftSpanBigint(s *Span, i int64) bool {
	_cret := C.overleft_span_bigint(s._inner, C.int64_t(i))
	return bool(_cret)
}


// OverleftSpanFloat wraps MEOS C function overleft_span_float.
func OverleftSpanFloat(s *Span, d float64) bool {
	_cret := C.overleft_span_float(s._inner, C.double(d))
	return bool(_cret)
}


// OverleftSpanInt wraps MEOS C function overleft_span_int.
func OverleftSpanInt(s *Span, i int) bool {
	_cret := C.overleft_span_int(s._inner, C.int(i))
	return bool(_cret)
}


// OverleftSpanSpan wraps MEOS C function overleft_span_span.
func OverleftSpanSpan(s1 *Span, s2 *Span) bool {
	_cret := C.overleft_span_span(s1._inner, s2._inner)
	return bool(_cret)
}


// OverleftSpanSpanset wraps MEOS C function overleft_span_spanset.
func OverleftSpanSpanset(s *Span, ss *SpanSet) bool {
	_cret := C.overleft_span_spanset(s._inner, ss._inner)
	return bool(_cret)
}


// OverleftSpansetBigint wraps MEOS C function overleft_spanset_bigint.
func OverleftSpansetBigint(ss *SpanSet, i int64) bool {
	_cret := C.overleft_spanset_bigint(ss._inner, C.int64_t(i))
	return bool(_cret)
}


// OverleftSpansetFloat wraps MEOS C function overleft_spanset_float.
func OverleftSpansetFloat(ss *SpanSet, d float64) bool {
	_cret := C.overleft_spanset_float(ss._inner, C.double(d))
	return bool(_cret)
}


// OverleftSpansetInt wraps MEOS C function overleft_spanset_int.
func OverleftSpansetInt(ss *SpanSet, i int) bool {
	_cret := C.overleft_spanset_int(ss._inner, C.int(i))
	return bool(_cret)
}


// OverleftSpansetSpan wraps MEOS C function overleft_spanset_span.
func OverleftSpansetSpan(ss *SpanSet, s *Span) bool {
	_cret := C.overleft_spanset_span(ss._inner, s._inner)
	return bool(_cret)
}


// OverleftSpansetSpanset wraps MEOS C function overleft_spanset_spanset.
func OverleftSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) bool {
	_cret := C.overleft_spanset_spanset(ss1._inner, ss2._inner)
	return bool(_cret)
}


// OverleftTextSet wraps MEOS C function overleft_text_set.
func OverleftTextSet(txt string, s *Set) bool {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.overleft_text_set(_c_txt, s._inner)
	return bool(_cret)
}


// OverrightBigintSet wraps MEOS C function overright_bigint_set.
func OverrightBigintSet(i int64, s *Set) bool {
	_cret := C.overright_bigint_set(C.int64_t(i), s._inner)
	return bool(_cret)
}


// OverrightBigintSpan wraps MEOS C function overright_bigint_span.
func OverrightBigintSpan(i int64, s *Span) bool {
	_cret := C.overright_bigint_span(C.int64_t(i), s._inner)
	return bool(_cret)
}


// OverrightBigintSpanset wraps MEOS C function overright_bigint_spanset.
func OverrightBigintSpanset(i int64, ss *SpanSet) bool {
	_cret := C.overright_bigint_spanset(C.int64_t(i), ss._inner)
	return bool(_cret)
}


// OverrightFloatSet wraps MEOS C function overright_float_set.
func OverrightFloatSet(d float64, s *Set) bool {
	_cret := C.overright_float_set(C.double(d), s._inner)
	return bool(_cret)
}


// OverrightFloatSpan wraps MEOS C function overright_float_span.
func OverrightFloatSpan(d float64, s *Span) bool {
	_cret := C.overright_float_span(C.double(d), s._inner)
	return bool(_cret)
}


// OverrightFloatSpanset wraps MEOS C function overright_float_spanset.
func OverrightFloatSpanset(d float64, ss *SpanSet) bool {
	_cret := C.overright_float_spanset(C.double(d), ss._inner)
	return bool(_cret)
}


// OverrightIntSet wraps MEOS C function overright_int_set.
func OverrightIntSet(i int, s *Set) bool {
	_cret := C.overright_int_set(C.int(i), s._inner)
	return bool(_cret)
}


// OverrightIntSpan wraps MEOS C function overright_int_span.
func OverrightIntSpan(i int, s *Span) bool {
	_cret := C.overright_int_span(C.int(i), s._inner)
	return bool(_cret)
}


// OverrightIntSpanset wraps MEOS C function overright_int_spanset.
func OverrightIntSpanset(i int, ss *SpanSet) bool {
	_cret := C.overright_int_spanset(C.int(i), ss._inner)
	return bool(_cret)
}


// OverrightSetBigint wraps MEOS C function overright_set_bigint.
func OverrightSetBigint(s *Set, i int64) bool {
	_cret := C.overright_set_bigint(s._inner, C.int64_t(i))
	return bool(_cret)
}


// OverrightSetFloat wraps MEOS C function overright_set_float.
func OverrightSetFloat(s *Set, d float64) bool {
	_cret := C.overright_set_float(s._inner, C.double(d))
	return bool(_cret)
}


// OverrightSetInt wraps MEOS C function overright_set_int.
func OverrightSetInt(s *Set, i int) bool {
	_cret := C.overright_set_int(s._inner, C.int(i))
	return bool(_cret)
}


// OverrightSetSet wraps MEOS C function overright_set_set.
func OverrightSetSet(s1 *Set, s2 *Set) bool {
	_cret := C.overright_set_set(s1._inner, s2._inner)
	return bool(_cret)
}


// OverrightSetText wraps MEOS C function overright_set_text.
func OverrightSetText(s *Set, txt string) bool {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.overright_set_text(s._inner, _c_txt)
	return bool(_cret)
}


// OverrightSpanBigint wraps MEOS C function overright_span_bigint.
func OverrightSpanBigint(s *Span, i int64) bool {
	_cret := C.overright_span_bigint(s._inner, C.int64_t(i))
	return bool(_cret)
}


// OverrightSpanFloat wraps MEOS C function overright_span_float.
func OverrightSpanFloat(s *Span, d float64) bool {
	_cret := C.overright_span_float(s._inner, C.double(d))
	return bool(_cret)
}


// OverrightSpanInt wraps MEOS C function overright_span_int.
func OverrightSpanInt(s *Span, i int) bool {
	_cret := C.overright_span_int(s._inner, C.int(i))
	return bool(_cret)
}


// OverrightSpanSpan wraps MEOS C function overright_span_span.
func OverrightSpanSpan(s1 *Span, s2 *Span) bool {
	_cret := C.overright_span_span(s1._inner, s2._inner)
	return bool(_cret)
}


// OverrightSpanSpanset wraps MEOS C function overright_span_spanset.
func OverrightSpanSpanset(s *Span, ss *SpanSet) bool {
	_cret := C.overright_span_spanset(s._inner, ss._inner)
	return bool(_cret)
}


// OverrightSpansetBigint wraps MEOS C function overright_spanset_bigint.
func OverrightSpansetBigint(ss *SpanSet, i int64) bool {
	_cret := C.overright_spanset_bigint(ss._inner, C.int64_t(i))
	return bool(_cret)
}


// OverrightSpansetFloat wraps MEOS C function overright_spanset_float.
func OverrightSpansetFloat(ss *SpanSet, d float64) bool {
	_cret := C.overright_spanset_float(ss._inner, C.double(d))
	return bool(_cret)
}


// OverrightSpansetInt wraps MEOS C function overright_spanset_int.
func OverrightSpansetInt(ss *SpanSet, i int) bool {
	_cret := C.overright_spanset_int(ss._inner, C.int(i))
	return bool(_cret)
}


// OverrightSpansetSpan wraps MEOS C function overright_spanset_span.
func OverrightSpansetSpan(ss *SpanSet, s *Span) bool {
	_cret := C.overright_spanset_span(ss._inner, s._inner)
	return bool(_cret)
}


// OverrightSpansetSpanset wraps MEOS C function overright_spanset_spanset.
func OverrightSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) bool {
	_cret := C.overright_spanset_spanset(ss1._inner, ss2._inner)
	return bool(_cret)
}


// OverrightTextSet wraps MEOS C function overright_text_set.
func OverrightTextSet(txt string, s *Set) bool {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.overright_text_set(_c_txt, s._inner)
	return bool(_cret)
}


// RightBigintSet wraps MEOS C function right_bigint_set.
func RightBigintSet(i int64, s *Set) bool {
	_cret := C.right_bigint_set(C.int64_t(i), s._inner)
	return bool(_cret)
}


// RightBigintSpan wraps MEOS C function right_bigint_span.
func RightBigintSpan(i int64, s *Span) bool {
	_cret := C.right_bigint_span(C.int64_t(i), s._inner)
	return bool(_cret)
}


// RightBigintSpanset wraps MEOS C function right_bigint_spanset.
func RightBigintSpanset(i int64, ss *SpanSet) bool {
	_cret := C.right_bigint_spanset(C.int64_t(i), ss._inner)
	return bool(_cret)
}


// RightFloatSet wraps MEOS C function right_float_set.
func RightFloatSet(d float64, s *Set) bool {
	_cret := C.right_float_set(C.double(d), s._inner)
	return bool(_cret)
}


// RightFloatSpan wraps MEOS C function right_float_span.
func RightFloatSpan(d float64, s *Span) bool {
	_cret := C.right_float_span(C.double(d), s._inner)
	return bool(_cret)
}


// RightFloatSpanset wraps MEOS C function right_float_spanset.
func RightFloatSpanset(d float64, ss *SpanSet) bool {
	_cret := C.right_float_spanset(C.double(d), ss._inner)
	return bool(_cret)
}


// RightIntSet wraps MEOS C function right_int_set.
func RightIntSet(i int, s *Set) bool {
	_cret := C.right_int_set(C.int(i), s._inner)
	return bool(_cret)
}


// RightIntSpan wraps MEOS C function right_int_span.
func RightIntSpan(i int, s *Span) bool {
	_cret := C.right_int_span(C.int(i), s._inner)
	return bool(_cret)
}


// RightIntSpanset wraps MEOS C function right_int_spanset.
func RightIntSpanset(i int, ss *SpanSet) bool {
	_cret := C.right_int_spanset(C.int(i), ss._inner)
	return bool(_cret)
}


// RightSetBigint wraps MEOS C function right_set_bigint.
func RightSetBigint(s *Set, i int64) bool {
	_cret := C.right_set_bigint(s._inner, C.int64_t(i))
	return bool(_cret)
}


// RightSetFloat wraps MEOS C function right_set_float.
func RightSetFloat(s *Set, d float64) bool {
	_cret := C.right_set_float(s._inner, C.double(d))
	return bool(_cret)
}


// RightSetInt wraps MEOS C function right_set_int.
func RightSetInt(s *Set, i int) bool {
	_cret := C.right_set_int(s._inner, C.int(i))
	return bool(_cret)
}


// RightSetSet wraps MEOS C function right_set_set.
func RightSetSet(s1 *Set, s2 *Set) bool {
	_cret := C.right_set_set(s1._inner, s2._inner)
	return bool(_cret)
}


// RightSetText wraps MEOS C function right_set_text.
func RightSetText(s *Set, txt string) bool {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.right_set_text(s._inner, _c_txt)
	return bool(_cret)
}


// RightSpanBigint wraps MEOS C function right_span_bigint.
func RightSpanBigint(s *Span, i int64) bool {
	_cret := C.right_span_bigint(s._inner, C.int64_t(i))
	return bool(_cret)
}


// RightSpanFloat wraps MEOS C function right_span_float.
func RightSpanFloat(s *Span, d float64) bool {
	_cret := C.right_span_float(s._inner, C.double(d))
	return bool(_cret)
}


// RightSpanInt wraps MEOS C function right_span_int.
func RightSpanInt(s *Span, i int) bool {
	_cret := C.right_span_int(s._inner, C.int(i))
	return bool(_cret)
}


// RightSpanSpan wraps MEOS C function right_span_span.
func RightSpanSpan(s1 *Span, s2 *Span) bool {
	_cret := C.right_span_span(s1._inner, s2._inner)
	return bool(_cret)
}


// RightSpanSpanset wraps MEOS C function right_span_spanset.
func RightSpanSpanset(s *Span, ss *SpanSet) bool {
	_cret := C.right_span_spanset(s._inner, ss._inner)
	return bool(_cret)
}


// RightSpansetBigint wraps MEOS C function right_spanset_bigint.
func RightSpansetBigint(ss *SpanSet, i int64) bool {
	_cret := C.right_spanset_bigint(ss._inner, C.int64_t(i))
	return bool(_cret)
}


// RightSpansetFloat wraps MEOS C function right_spanset_float.
func RightSpansetFloat(ss *SpanSet, d float64) bool {
	_cret := C.right_spanset_float(ss._inner, C.double(d))
	return bool(_cret)
}


// RightSpansetInt wraps MEOS C function right_spanset_int.
func RightSpansetInt(ss *SpanSet, i int) bool {
	_cret := C.right_spanset_int(ss._inner, C.int(i))
	return bool(_cret)
}


// RightSpansetSpan wraps MEOS C function right_spanset_span.
func RightSpansetSpan(ss *SpanSet, s *Span) bool {
	_cret := C.right_spanset_span(ss._inner, s._inner)
	return bool(_cret)
}


// RightSpansetSpanset wraps MEOS C function right_spanset_spanset.
func RightSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) bool {
	_cret := C.right_spanset_spanset(ss1._inner, ss2._inner)
	return bool(_cret)
}


// RightTextSet wraps MEOS C function right_text_set.
func RightTextSet(txt string, s *Set) bool {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.right_text_set(_c_txt, s._inner)
	return bool(_cret)
}


// IntersectionBigintSet wraps MEOS C function intersection_bigint_set.
func IntersectionBigintSet(i int64, s *Set) *Set {
	_cret := C.intersection_bigint_set(C.int64_t(i), s._inner)
	return &Set{_inner: _cret}
}


// IntersectionDateSet wraps MEOS C function intersection_date_set.
func IntersectionDateSet(d int32, s *Set) *Set {
	_cret := C.intersection_date_set(C.DateADT(d), s._inner)
	return &Set{_inner: _cret}
}


// IntersectionFloatSet wraps MEOS C function intersection_float_set.
func IntersectionFloatSet(d float64, s *Set) *Set {
	_cret := C.intersection_float_set(C.double(d), s._inner)
	return &Set{_inner: _cret}
}


// IntersectionIntSet wraps MEOS C function intersection_int_set.
func IntersectionIntSet(i int, s *Set) *Set {
	_cret := C.intersection_int_set(C.int(i), s._inner)
	return &Set{_inner: _cret}
}


// IntersectionSetBigint wraps MEOS C function intersection_set_bigint.
func IntersectionSetBigint(s *Set, i int64) *Set {
	_cret := C.intersection_set_bigint(s._inner, C.int64_t(i))
	return &Set{_inner: _cret}
}


// IntersectionSetDate wraps MEOS C function intersection_set_date.
func IntersectionSetDate(s *Set, d int32) *Set {
	_cret := C.intersection_set_date(s._inner, C.DateADT(d))
	return &Set{_inner: _cret}
}


// IntersectionSetFloat wraps MEOS C function intersection_set_float.
func IntersectionSetFloat(s *Set, d float64) *Set {
	_cret := C.intersection_set_float(s._inner, C.double(d))
	return &Set{_inner: _cret}
}


// IntersectionSetInt wraps MEOS C function intersection_set_int.
func IntersectionSetInt(s *Set, i int) *Set {
	_cret := C.intersection_set_int(s._inner, C.int(i))
	return &Set{_inner: _cret}
}


// IntersectionSetSet wraps MEOS C function intersection_set_set.
func IntersectionSetSet(s1 *Set, s2 *Set) *Set {
	_cret := C.intersection_set_set(s1._inner, s2._inner)
	return &Set{_inner: _cret}
}


// IntersectionSetText wraps MEOS C function intersection_set_text.
func IntersectionSetText(s *Set, txt string) *Set {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.intersection_set_text(s._inner, _c_txt)
	return &Set{_inner: _cret}
}


// IntersectionSetTimestamptz wraps MEOS C function intersection_set_timestamptz.
func IntersectionSetTimestamptz(s *Set, t int64) *Set {
	_cret := C.intersection_set_timestamptz(s._inner, C.TimestampTz(t))
	return &Set{_inner: _cret}
}


// IntersectionSpanBigint wraps MEOS C function intersection_span_bigint.
func IntersectionSpanBigint(s *Span, i int64) *Span {
	_cret := C.intersection_span_bigint(s._inner, C.int64_t(i))
	return &Span{_inner: _cret}
}


// IntersectionSpanDate wraps MEOS C function intersection_span_date.
func IntersectionSpanDate(s *Span, d int32) *Span {
	_cret := C.intersection_span_date(s._inner, C.DateADT(d))
	return &Span{_inner: _cret}
}


// IntersectionSpanFloat wraps MEOS C function intersection_span_float.
func IntersectionSpanFloat(s *Span, d float64) *Span {
	_cret := C.intersection_span_float(s._inner, C.double(d))
	return &Span{_inner: _cret}
}


// IntersectionSpanInt wraps MEOS C function intersection_span_int.
func IntersectionSpanInt(s *Span, i int) *Span {
	_cret := C.intersection_span_int(s._inner, C.int(i))
	return &Span{_inner: _cret}
}


// IntersectionSpanSpan wraps MEOS C function intersection_span_span.
func IntersectionSpanSpan(s1 *Span, s2 *Span) *Span {
	_cret := C.intersection_span_span(s1._inner, s2._inner)
	return &Span{_inner: _cret}
}


// IntersectionSpanSpanset wraps MEOS C function intersection_span_spanset.
func IntersectionSpanSpanset(s *Span, ss *SpanSet) *SpanSet {
	_cret := C.intersection_span_spanset(s._inner, ss._inner)
	return &SpanSet{_inner: _cret}
}


// IntersectionSpanTimestamptz wraps MEOS C function intersection_span_timestamptz.
func IntersectionSpanTimestamptz(s *Span, t int64) *Span {
	_cret := C.intersection_span_timestamptz(s._inner, C.TimestampTz(t))
	return &Span{_inner: _cret}
}


// IntersectionSpansetBigint wraps MEOS C function intersection_spanset_bigint.
func IntersectionSpansetBigint(ss *SpanSet, i int64) *SpanSet {
	_cret := C.intersection_spanset_bigint(ss._inner, C.int64_t(i))
	return &SpanSet{_inner: _cret}
}


// IntersectionSpansetDate wraps MEOS C function intersection_spanset_date.
func IntersectionSpansetDate(ss *SpanSet, d int32) *SpanSet {
	_cret := C.intersection_spanset_date(ss._inner, C.DateADT(d))
	return &SpanSet{_inner: _cret}
}


// IntersectionSpansetFloat wraps MEOS C function intersection_spanset_float.
func IntersectionSpansetFloat(ss *SpanSet, d float64) *SpanSet {
	_cret := C.intersection_spanset_float(ss._inner, C.double(d))
	return &SpanSet{_inner: _cret}
}


// IntersectionSpansetInt wraps MEOS C function intersection_spanset_int.
func IntersectionSpansetInt(ss *SpanSet, i int) *SpanSet {
	_cret := C.intersection_spanset_int(ss._inner, C.int(i))
	return &SpanSet{_inner: _cret}
}


// IntersectionSpansetSpan wraps MEOS C function intersection_spanset_span.
func IntersectionSpansetSpan(ss *SpanSet, s *Span) *SpanSet {
	_cret := C.intersection_spanset_span(ss._inner, s._inner)
	return &SpanSet{_inner: _cret}
}


// IntersectionSpansetSpanset wraps MEOS C function intersection_spanset_spanset.
func IntersectionSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) *SpanSet {
	_cret := C.intersection_spanset_spanset(ss1._inner, ss2._inner)
	return &SpanSet{_inner: _cret}
}


// IntersectionSpansetTimestamptz wraps MEOS C function intersection_spanset_timestamptz.
func IntersectionSpansetTimestamptz(ss *SpanSet, t int64) *SpanSet {
	_cret := C.intersection_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return &SpanSet{_inner: _cret}
}


// IntersectionTextSet wraps MEOS C function intersection_text_set.
func IntersectionTextSet(txt string, s *Set) *Set {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.intersection_text_set(_c_txt, s._inner)
	return &Set{_inner: _cret}
}


// IntersectionTimestamptzSet wraps MEOS C function intersection_timestamptz_set.
func IntersectionTimestamptzSet(t int64, s *Set) *Set {
	_cret := C.intersection_timestamptz_set(C.TimestampTz(t), s._inner)
	return &Set{_inner: _cret}
}


// MinusBigintSet wraps MEOS C function minus_bigint_set.
func MinusBigintSet(i int64, s *Set) *Set {
	_cret := C.minus_bigint_set(C.int64_t(i), s._inner)
	return &Set{_inner: _cret}
}


// MinusBigintSpan wraps MEOS C function minus_bigint_span.
func MinusBigintSpan(i int64, s *Span) *SpanSet {
	_cret := C.minus_bigint_span(C.int64_t(i), s._inner)
	return &SpanSet{_inner: _cret}
}


// MinusBigintSpanset wraps MEOS C function minus_bigint_spanset.
func MinusBigintSpanset(i int64, ss *SpanSet) *SpanSet {
	_cret := C.minus_bigint_spanset(C.int64_t(i), ss._inner)
	return &SpanSet{_inner: _cret}
}


// MinusDateSet wraps MEOS C function minus_date_set.
func MinusDateSet(d int32, s *Set) *Set {
	_cret := C.minus_date_set(C.DateADT(d), s._inner)
	return &Set{_inner: _cret}
}


// MinusDateSpan wraps MEOS C function minus_date_span.
func MinusDateSpan(d int32, s *Span) *SpanSet {
	_cret := C.minus_date_span(C.DateADT(d), s._inner)
	return &SpanSet{_inner: _cret}
}


// MinusDateSpanset wraps MEOS C function minus_date_spanset.
func MinusDateSpanset(d int32, ss *SpanSet) *SpanSet {
	_cret := C.minus_date_spanset(C.DateADT(d), ss._inner)
	return &SpanSet{_inner: _cret}
}


// MinusFloatSet wraps MEOS C function minus_float_set.
func MinusFloatSet(d float64, s *Set) *Set {
	_cret := C.minus_float_set(C.double(d), s._inner)
	return &Set{_inner: _cret}
}


// MinusFloatSpan wraps MEOS C function minus_float_span.
func MinusFloatSpan(d float64, s *Span) *SpanSet {
	_cret := C.minus_float_span(C.double(d), s._inner)
	return &SpanSet{_inner: _cret}
}


// MinusFloatSpanset wraps MEOS C function minus_float_spanset.
func MinusFloatSpanset(d float64, ss *SpanSet) *SpanSet {
	_cret := C.minus_float_spanset(C.double(d), ss._inner)
	return &SpanSet{_inner: _cret}
}


// MinusIntSet wraps MEOS C function minus_int_set.
func MinusIntSet(i int, s *Set) *Set {
	_cret := C.minus_int_set(C.int(i), s._inner)
	return &Set{_inner: _cret}
}


// MinusIntSpan wraps MEOS C function minus_int_span.
func MinusIntSpan(i int, s *Span) *SpanSet {
	_cret := C.minus_int_span(C.int(i), s._inner)
	return &SpanSet{_inner: _cret}
}


// MinusIntSpanset wraps MEOS C function minus_int_spanset.
func MinusIntSpanset(i int, ss *SpanSet) *SpanSet {
	_cret := C.minus_int_spanset(C.int(i), ss._inner)
	return &SpanSet{_inner: _cret}
}


// MinusSetBigint wraps MEOS C function minus_set_bigint.
func MinusSetBigint(s *Set, i int64) *Set {
	_cret := C.minus_set_bigint(s._inner, C.int64_t(i))
	return &Set{_inner: _cret}
}


// MinusSetDate wraps MEOS C function minus_set_date.
func MinusSetDate(s *Set, d int32) *Set {
	_cret := C.minus_set_date(s._inner, C.DateADT(d))
	return &Set{_inner: _cret}
}


// MinusSetFloat wraps MEOS C function minus_set_float.
func MinusSetFloat(s *Set, d float64) *Set {
	_cret := C.minus_set_float(s._inner, C.double(d))
	return &Set{_inner: _cret}
}


// MinusSetInt wraps MEOS C function minus_set_int.
func MinusSetInt(s *Set, i int) *Set {
	_cret := C.minus_set_int(s._inner, C.int(i))
	return &Set{_inner: _cret}
}


// MinusSetSet wraps MEOS C function minus_set_set.
func MinusSetSet(s1 *Set, s2 *Set) *Set {
	_cret := C.minus_set_set(s1._inner, s2._inner)
	return &Set{_inner: _cret}
}


// MinusSetText wraps MEOS C function minus_set_text.
func MinusSetText(s *Set, txt string) *Set {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.minus_set_text(s._inner, _c_txt)
	return &Set{_inner: _cret}
}


// MinusSetTimestamptz wraps MEOS C function minus_set_timestamptz.
func MinusSetTimestamptz(s *Set, t int64) *Set {
	_cret := C.minus_set_timestamptz(s._inner, C.TimestampTz(t))
	return &Set{_inner: _cret}
}


// MinusSpanBigint wraps MEOS C function minus_span_bigint.
func MinusSpanBigint(s *Span, i int64) *SpanSet {
	_cret := C.minus_span_bigint(s._inner, C.int64_t(i))
	return &SpanSet{_inner: _cret}
}


// MinusSpanDate wraps MEOS C function minus_span_date.
func MinusSpanDate(s *Span, d int32) *SpanSet {
	_cret := C.minus_span_date(s._inner, C.DateADT(d))
	return &SpanSet{_inner: _cret}
}


// MinusSpanFloat wraps MEOS C function minus_span_float.
func MinusSpanFloat(s *Span, d float64) *SpanSet {
	_cret := C.minus_span_float(s._inner, C.double(d))
	return &SpanSet{_inner: _cret}
}


// MinusSpanInt wraps MEOS C function minus_span_int.
func MinusSpanInt(s *Span, i int) *SpanSet {
	_cret := C.minus_span_int(s._inner, C.int(i))
	return &SpanSet{_inner: _cret}
}


// MinusSpanSpan wraps MEOS C function minus_span_span.
func MinusSpanSpan(s1 *Span, s2 *Span) *SpanSet {
	_cret := C.minus_span_span(s1._inner, s2._inner)
	return &SpanSet{_inner: _cret}
}


// MinusSpanSpanset wraps MEOS C function minus_span_spanset.
func MinusSpanSpanset(s *Span, ss *SpanSet) *SpanSet {
	_cret := C.minus_span_spanset(s._inner, ss._inner)
	return &SpanSet{_inner: _cret}
}


// MinusSpanTimestamptz wraps MEOS C function minus_span_timestamptz.
func MinusSpanTimestamptz(s *Span, t int64) *SpanSet {
	_cret := C.minus_span_timestamptz(s._inner, C.TimestampTz(t))
	return &SpanSet{_inner: _cret}
}


// MinusSpansetBigint wraps MEOS C function minus_spanset_bigint.
func MinusSpansetBigint(ss *SpanSet, i int64) *SpanSet {
	_cret := C.minus_spanset_bigint(ss._inner, C.int64_t(i))
	return &SpanSet{_inner: _cret}
}


// MinusSpansetDate wraps MEOS C function minus_spanset_date.
func MinusSpansetDate(ss *SpanSet, d int32) *SpanSet {
	_cret := C.minus_spanset_date(ss._inner, C.DateADT(d))
	return &SpanSet{_inner: _cret}
}


// MinusSpansetFloat wraps MEOS C function minus_spanset_float.
func MinusSpansetFloat(ss *SpanSet, d float64) *SpanSet {
	_cret := C.minus_spanset_float(ss._inner, C.double(d))
	return &SpanSet{_inner: _cret}
}


// MinusSpansetInt wraps MEOS C function minus_spanset_int.
func MinusSpansetInt(ss *SpanSet, i int) *SpanSet {
	_cret := C.minus_spanset_int(ss._inner, C.int(i))
	return &SpanSet{_inner: _cret}
}


// MinusSpansetSpan wraps MEOS C function minus_spanset_span.
func MinusSpansetSpan(ss *SpanSet, s *Span) *SpanSet {
	_cret := C.minus_spanset_span(ss._inner, s._inner)
	return &SpanSet{_inner: _cret}
}


// MinusSpansetSpanset wraps MEOS C function minus_spanset_spanset.
func MinusSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) *SpanSet {
	_cret := C.minus_spanset_spanset(ss1._inner, ss2._inner)
	return &SpanSet{_inner: _cret}
}


// MinusSpansetTimestamptz wraps MEOS C function minus_spanset_timestamptz.
func MinusSpansetTimestamptz(ss *SpanSet, t int64) *SpanSet {
	_cret := C.minus_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return &SpanSet{_inner: _cret}
}


// MinusTextSet wraps MEOS C function minus_text_set.
func MinusTextSet(txt string, s *Set) *Set {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.minus_text_set(_c_txt, s._inner)
	return &Set{_inner: _cret}
}


// MinusTimestamptzSet wraps MEOS C function minus_timestamptz_set.
func MinusTimestamptzSet(t int64, s *Set) *Set {
	_cret := C.minus_timestamptz_set(C.TimestampTz(t), s._inner)
	return &Set{_inner: _cret}
}


// MinusTimestamptzSpan wraps MEOS C function minus_timestamptz_span.
func MinusTimestamptzSpan(t int64, s *Span) *SpanSet {
	_cret := C.minus_timestamptz_span(C.TimestampTz(t), s._inner)
	return &SpanSet{_inner: _cret}
}


// MinusTimestamptzSpanset wraps MEOS C function minus_timestamptz_spanset.
func MinusTimestamptzSpanset(t int64, ss *SpanSet) *SpanSet {
	_cret := C.minus_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	return &SpanSet{_inner: _cret}
}


// UnionBigintSet wraps MEOS C function union_bigint_set.
func UnionBigintSet(i int64, s *Set) *Set {
	_cret := C.gunion_bigint_set(C.int64_t(i), s._inner)
	return &Set{_inner: _cret}
}


// UnionBigintSpan wraps MEOS C function union_bigint_span.
func UnionBigintSpan(s *Span, i int64) *SpanSet {
	_cret := C.gunion_bigint_span(s._inner, C.int64_t(i))
	return &SpanSet{_inner: _cret}
}


// UnionBigintSpanset wraps MEOS C function union_bigint_spanset.
func UnionBigintSpanset(i int64, ss *SpanSet) *SpanSet {
	_cret := C.gunion_bigint_spanset(C.int64_t(i), ss._inner)
	return &SpanSet{_inner: _cret}
}


// UnionDateSet wraps MEOS C function union_date_set.
func UnionDateSet(d int32, s *Set) *Set {
	_cret := C.gunion_date_set(C.DateADT(d), s._inner)
	return &Set{_inner: _cret}
}


// UnionDateSpan wraps MEOS C function union_date_span.
func UnionDateSpan(s *Span, d int32) *SpanSet {
	_cret := C.gunion_date_span(s._inner, C.DateADT(d))
	return &SpanSet{_inner: _cret}
}


// UnionDateSpanset wraps MEOS C function union_date_spanset.
func UnionDateSpanset(d int32, ss *SpanSet) *SpanSet {
	_cret := C.gunion_date_spanset(C.DateADT(d), ss._inner)
	return &SpanSet{_inner: _cret}
}


// UnionFloatSet wraps MEOS C function union_float_set.
func UnionFloatSet(d float64, s *Set) *Set {
	_cret := C.gunion_float_set(C.double(d), s._inner)
	return &Set{_inner: _cret}
}


// UnionFloatSpan wraps MEOS C function union_float_span.
func UnionFloatSpan(s *Span, d float64) *SpanSet {
	_cret := C.gunion_float_span(s._inner, C.double(d))
	return &SpanSet{_inner: _cret}
}


// UnionFloatSpanset wraps MEOS C function union_float_spanset.
func UnionFloatSpanset(d float64, ss *SpanSet) *SpanSet {
	_cret := C.gunion_float_spanset(C.double(d), ss._inner)
	return &SpanSet{_inner: _cret}
}


// UnionIntSet wraps MEOS C function union_int_set.
func UnionIntSet(i int, s *Set) *Set {
	_cret := C.gunion_int_set(C.int(i), s._inner)
	return &Set{_inner: _cret}
}


// UnionIntSpan wraps MEOS C function union_int_span.
func UnionIntSpan(i int, s *Span) *SpanSet {
	_cret := C.gunion_int_span(C.int(i), s._inner)
	return &SpanSet{_inner: _cret}
}


// UnionIntSpanset wraps MEOS C function union_int_spanset.
func UnionIntSpanset(i int, ss *SpanSet) *SpanSet {
	_cret := C.gunion_int_spanset(C.int(i), ss._inner)
	return &SpanSet{_inner: _cret}
}


// UnionSetBigint wraps MEOS C function union_set_bigint.
func UnionSetBigint(s *Set, i int64) *Set {
	_cret := C.gunion_set_bigint(s._inner, C.int64_t(i))
	return &Set{_inner: _cret}
}


// UnionSetDate wraps MEOS C function union_set_date.
func UnionSetDate(s *Set, d int32) *Set {
	_cret := C.gunion_set_date(s._inner, C.DateADT(d))
	return &Set{_inner: _cret}
}


// UnionSetFloat wraps MEOS C function union_set_float.
func UnionSetFloat(s *Set, d float64) *Set {
	_cret := C.gunion_set_float(s._inner, C.double(d))
	return &Set{_inner: _cret}
}


// UnionSetInt wraps MEOS C function union_set_int.
func UnionSetInt(s *Set, i int) *Set {
	_cret := C.gunion_set_int(s._inner, C.int(i))
	return &Set{_inner: _cret}
}


// UnionSetSet wraps MEOS C function union_set_set.
func UnionSetSet(s1 *Set, s2 *Set) *Set {
	_cret := C.gunion_set_set(s1._inner, s2._inner)
	return &Set{_inner: _cret}
}


// UnionSetText wraps MEOS C function union_set_text.
func UnionSetText(s *Set, txt string) *Set {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.gunion_set_text(s._inner, _c_txt)
	return &Set{_inner: _cret}
}


// UnionSetTimestamptz wraps MEOS C function union_set_timestamptz.
func UnionSetTimestamptz(s *Set, t int64) *Set {
	_cret := C.gunion_set_timestamptz(s._inner, C.TimestampTz(t))
	return &Set{_inner: _cret}
}


// UnionSpanBigint wraps MEOS C function union_span_bigint.
func UnionSpanBigint(s *Span, i int64) *SpanSet {
	_cret := C.gunion_span_bigint(s._inner, C.int64_t(i))
	return &SpanSet{_inner: _cret}
}


// UnionSpanDate wraps MEOS C function union_span_date.
func UnionSpanDate(s *Span, d int32) *SpanSet {
	_cret := C.gunion_span_date(s._inner, C.DateADT(d))
	return &SpanSet{_inner: _cret}
}


// UnionSpanFloat wraps MEOS C function union_span_float.
func UnionSpanFloat(s *Span, d float64) *SpanSet {
	_cret := C.gunion_span_float(s._inner, C.double(d))
	return &SpanSet{_inner: _cret}
}


// UnionSpanInt wraps MEOS C function union_span_int.
func UnionSpanInt(s *Span, i int) *SpanSet {
	_cret := C.gunion_span_int(s._inner, C.int(i))
	return &SpanSet{_inner: _cret}
}


// UnionSpanSpan wraps MEOS C function union_span_span.
func UnionSpanSpan(s1 *Span, s2 *Span) *SpanSet {
	_cret := C.gunion_span_span(s1._inner, s2._inner)
	return &SpanSet{_inner: _cret}
}


// SuperUnionSpanSpan wraps MEOS C function super_union_span_span.
func SuperUnionSpanSpan(s1 *Span, s2 *Span) *Span {
	_cret := C.super_union_span_span(s1._inner, s2._inner)
	return &Span{_inner: _cret}
}


// UnionSpanSpanset wraps MEOS C function union_span_spanset.
func UnionSpanSpanset(s *Span, ss *SpanSet) *SpanSet {
	_cret := C.gunion_span_spanset(s._inner, ss._inner)
	return &SpanSet{_inner: _cret}
}


// UnionSpanTimestamptz wraps MEOS C function union_span_timestamptz.
func UnionSpanTimestamptz(s *Span, t int64) *SpanSet {
	_cret := C.gunion_span_timestamptz(s._inner, C.TimestampTz(t))
	return &SpanSet{_inner: _cret}
}


// UnionSpansetBigint wraps MEOS C function union_spanset_bigint.
func UnionSpansetBigint(ss *SpanSet, i int64) *SpanSet {
	_cret := C.gunion_spanset_bigint(ss._inner, C.int64_t(i))
	return &SpanSet{_inner: _cret}
}


// UnionSpansetDate wraps MEOS C function union_spanset_date.
func UnionSpansetDate(ss *SpanSet, d int32) *SpanSet {
	_cret := C.gunion_spanset_date(ss._inner, C.DateADT(d))
	return &SpanSet{_inner: _cret}
}


// UnionSpansetFloat wraps MEOS C function union_spanset_float.
func UnionSpansetFloat(ss *SpanSet, d float64) *SpanSet {
	_cret := C.gunion_spanset_float(ss._inner, C.double(d))
	return &SpanSet{_inner: _cret}
}


// UnionSpansetInt wraps MEOS C function union_spanset_int.
func UnionSpansetInt(ss *SpanSet, i int) *SpanSet {
	_cret := C.gunion_spanset_int(ss._inner, C.int(i))
	return &SpanSet{_inner: _cret}
}


// UnionSpansetSpan wraps MEOS C function union_spanset_span.
func UnionSpansetSpan(ss *SpanSet, s *Span) *SpanSet {
	_cret := C.gunion_spanset_span(ss._inner, s._inner)
	return &SpanSet{_inner: _cret}
}


// UnionSpansetSpanset wraps MEOS C function union_spanset_spanset.
func UnionSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) *SpanSet {
	_cret := C.gunion_spanset_spanset(ss1._inner, ss2._inner)
	return &SpanSet{_inner: _cret}
}


// UnionSpansetTimestamptz wraps MEOS C function union_spanset_timestamptz.
func UnionSpansetTimestamptz(ss *SpanSet, t int64) *SpanSet {
	_cret := C.gunion_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return &SpanSet{_inner: _cret}
}


// UnionTextSet wraps MEOS C function union_text_set.
func UnionTextSet(txt string, s *Set) *Set {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.gunion_text_set(_c_txt, s._inner)
	return &Set{_inner: _cret}
}


// UnionTimestamptzSet wraps MEOS C function union_timestamptz_set.
func UnionTimestamptzSet(t int64, s *Set) *Set {
	_cret := C.gunion_timestamptz_set(C.TimestampTz(t), s._inner)
	return &Set{_inner: _cret}
}


// UnionTimestamptzSpan wraps MEOS C function union_timestamptz_span.
func UnionTimestamptzSpan(t int64, s *Span) *SpanSet {
	_cret := C.gunion_timestamptz_span(C.TimestampTz(t), s._inner)
	return &SpanSet{_inner: _cret}
}


// UnionTimestamptzSpanset wraps MEOS C function union_timestamptz_spanset.
func UnionTimestamptzSpanset(t int64, ss *SpanSet) *SpanSet {
	_cret := C.gunion_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	return &SpanSet{_inner: _cret}
}


// DistanceBigintsetBigintset wraps MEOS C function distance_bigintset_bigintset.
func DistanceBigintsetBigintset(s1 *Set, s2 *Set) int64 {
	_cret := C.distance_bigintset_bigintset(s1._inner, s2._inner)
	return int64(_cret)
}


// DistanceBigintspanBigintspan wraps MEOS C function distance_bigintspan_bigintspan.
func DistanceBigintspanBigintspan(s1 *Span, s2 *Span) int64 {
	_cret := C.distance_bigintspan_bigintspan(s1._inner, s2._inner)
	return int64(_cret)
}


// DistanceBigintspansetBigintspan wraps MEOS C function distance_bigintspanset_bigintspan.
func DistanceBigintspansetBigintspan(ss *SpanSet, s *Span) int64 {
	_cret := C.distance_bigintspanset_bigintspan(ss._inner, s._inner)
	return int64(_cret)
}


// DistanceBigintspansetBigintspanset wraps MEOS C function distance_bigintspanset_bigintspanset.
func DistanceBigintspansetBigintspanset(ss1 *SpanSet, ss2 *SpanSet) int64 {
	_cret := C.distance_bigintspanset_bigintspanset(ss1._inner, ss2._inner)
	return int64(_cret)
}


// DistanceDatesetDateset wraps MEOS C function distance_dateset_dateset.
func DistanceDatesetDateset(s1 *Set, s2 *Set) int {
	_cret := C.distance_dateset_dateset(s1._inner, s2._inner)
	return int(_cret)
}


// DistanceDatespanDatespan wraps MEOS C function distance_datespan_datespan.
func DistanceDatespanDatespan(s1 *Span, s2 *Span) int {
	_cret := C.distance_datespan_datespan(s1._inner, s2._inner)
	return int(_cret)
}


// DistanceDatespansetDatespan wraps MEOS C function distance_datespanset_datespan.
func DistanceDatespansetDatespan(ss *SpanSet, s *Span) int {
	_cret := C.distance_datespanset_datespan(ss._inner, s._inner)
	return int(_cret)
}


// DistanceDatespansetDatespanset wraps MEOS C function distance_datespanset_datespanset.
func DistanceDatespansetDatespanset(ss1 *SpanSet, ss2 *SpanSet) int {
	_cret := C.distance_datespanset_datespanset(ss1._inner, ss2._inner)
	return int(_cret)
}


// DistanceFloatsetFloatset wraps MEOS C function distance_floatset_floatset.
func DistanceFloatsetFloatset(s1 *Set, s2 *Set) float64 {
	_cret := C.distance_floatset_floatset(s1._inner, s2._inner)
	return float64(_cret)
}


// DistanceFloatspanFloatspan wraps MEOS C function distance_floatspan_floatspan.
func DistanceFloatspanFloatspan(s1 *Span, s2 *Span) float64 {
	_cret := C.distance_floatspan_floatspan(s1._inner, s2._inner)
	return float64(_cret)
}


// DistanceFloatspansetFloatspan wraps MEOS C function distance_floatspanset_floatspan.
func DistanceFloatspansetFloatspan(ss *SpanSet, s *Span) float64 {
	_cret := C.distance_floatspanset_floatspan(ss._inner, s._inner)
	return float64(_cret)
}


// DistanceFloatspansetFloatspanset wraps MEOS C function distance_floatspanset_floatspanset.
func DistanceFloatspansetFloatspanset(ss1 *SpanSet, ss2 *SpanSet) float64 {
	_cret := C.distance_floatspanset_floatspanset(ss1._inner, ss2._inner)
	return float64(_cret)
}


// DistanceIntsetIntset wraps MEOS C function distance_intset_intset.
func DistanceIntsetIntset(s1 *Set, s2 *Set) int {
	_cret := C.distance_intset_intset(s1._inner, s2._inner)
	return int(_cret)
}


// DistanceIntspanIntspan wraps MEOS C function distance_intspan_intspan.
func DistanceIntspanIntspan(s1 *Span, s2 *Span) int {
	_cret := C.distance_intspan_intspan(s1._inner, s2._inner)
	return int(_cret)
}


// DistanceIntspansetIntspan wraps MEOS C function distance_intspanset_intspan.
func DistanceIntspansetIntspan(ss *SpanSet, s *Span) int {
	_cret := C.distance_intspanset_intspan(ss._inner, s._inner)
	return int(_cret)
}


// DistanceIntspansetIntspanset wraps MEOS C function distance_intspanset_intspanset.
func DistanceIntspansetIntspanset(ss1 *SpanSet, ss2 *SpanSet) int {
	_cret := C.distance_intspanset_intspanset(ss1._inner, ss2._inner)
	return int(_cret)
}


// DistanceSetBigint wraps MEOS C function distance_set_bigint.
func DistanceSetBigint(s *Set, i int64) int64 {
	_cret := C.distance_set_bigint(s._inner, C.int64_t(i))
	return int64(_cret)
}


// DistanceSetDate wraps MEOS C function distance_set_date.
func DistanceSetDate(s *Set, d int32) int {
	_cret := C.distance_set_date(s._inner, C.DateADT(d))
	return int(_cret)
}


// DistanceSetFloat wraps MEOS C function distance_set_float.
func DistanceSetFloat(s *Set, d float64) float64 {
	_cret := C.distance_set_float(s._inner, C.double(d))
	return float64(_cret)
}


// DistanceSetInt wraps MEOS C function distance_set_int.
func DistanceSetInt(s *Set, i int) int {
	_cret := C.distance_set_int(s._inner, C.int(i))
	return int(_cret)
}


// DistanceSetTimestamptz wraps MEOS C function distance_set_timestamptz.
func DistanceSetTimestamptz(s *Set, t int64) float64 {
	_cret := C.distance_set_timestamptz(s._inner, C.TimestampTz(t))
	return float64(_cret)
}


// DistanceSpanBigint wraps MEOS C function distance_span_bigint.
func DistanceSpanBigint(s *Span, i int64) int64 {
	_cret := C.distance_span_bigint(s._inner, C.int64_t(i))
	return int64(_cret)
}


// DistanceSpanDate wraps MEOS C function distance_span_date.
func DistanceSpanDate(s *Span, d int32) int {
	_cret := C.distance_span_date(s._inner, C.DateADT(d))
	return int(_cret)
}


// DistanceSpanFloat wraps MEOS C function distance_span_float.
func DistanceSpanFloat(s *Span, d float64) float64 {
	_cret := C.distance_span_float(s._inner, C.double(d))
	return float64(_cret)
}


// DistanceSpanInt wraps MEOS C function distance_span_int.
func DistanceSpanInt(s *Span, i int) int {
	_cret := C.distance_span_int(s._inner, C.int(i))
	return int(_cret)
}


// DistanceSpanTimestamptz wraps MEOS C function distance_span_timestamptz.
func DistanceSpanTimestamptz(s *Span, t int64) float64 {
	_cret := C.distance_span_timestamptz(s._inner, C.TimestampTz(t))
	return float64(_cret)
}


// DistanceSpansetBigint wraps MEOS C function distance_spanset_bigint.
func DistanceSpansetBigint(ss *SpanSet, i int64) int64 {
	_cret := C.distance_spanset_bigint(ss._inner, C.int64_t(i))
	return int64(_cret)
}


// DistanceSpansetDate wraps MEOS C function distance_spanset_date.
func DistanceSpansetDate(ss *SpanSet, d int32) int {
	_cret := C.distance_spanset_date(ss._inner, C.DateADT(d))
	return int(_cret)
}


// DistanceSpansetFloat wraps MEOS C function distance_spanset_float.
func DistanceSpansetFloat(ss *SpanSet, d float64) float64 {
	_cret := C.distance_spanset_float(ss._inner, C.double(d))
	return float64(_cret)
}


// DistanceSpansetInt wraps MEOS C function distance_spanset_int.
func DistanceSpansetInt(ss *SpanSet, i int) int {
	_cret := C.distance_spanset_int(ss._inner, C.int(i))
	return int(_cret)
}


// DistanceSpansetTimestamptz wraps MEOS C function distance_spanset_timestamptz.
func DistanceSpansetTimestamptz(ss *SpanSet, t int64) float64 {
	_cret := C.distance_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return float64(_cret)
}


// DistanceTstzsetTstzset wraps MEOS C function distance_tstzset_tstzset.
func DistanceTstzsetTstzset(s1 *Set, s2 *Set) float64 {
	_cret := C.distance_tstzset_tstzset(s1._inner, s2._inner)
	return float64(_cret)
}


// DistanceTstzspanTstzspan wraps MEOS C function distance_tstzspan_tstzspan.
func DistanceTstzspanTstzspan(s1 *Span, s2 *Span) float64 {
	_cret := C.distance_tstzspan_tstzspan(s1._inner, s2._inner)
	return float64(_cret)
}


// DistanceTstzspansetTstzspan wraps MEOS C function distance_tstzspanset_tstzspan.
func DistanceTstzspansetTstzspan(ss *SpanSet, s *Span) float64 {
	_cret := C.distance_tstzspanset_tstzspan(ss._inner, s._inner)
	return float64(_cret)
}


// DistanceTstzspansetTstzspanset wraps MEOS C function distance_tstzspanset_tstzspanset.
func DistanceTstzspansetTstzspanset(ss1 *SpanSet, ss2 *SpanSet) float64 {
	_cret := C.distance_tstzspanset_tstzspanset(ss1._inner, ss2._inner)
	return float64(_cret)
}


// BigintExtentTransfn wraps MEOS C function bigint_extent_transfn.
func BigintExtentTransfn(state *Span, i int64) *Span {
	_cret := C.bigint_extent_transfn(state._inner, C.int64_t(i))
	return &Span{_inner: _cret}
}


// BigintUnionTransfn wraps MEOS C function bigint_union_transfn.
func BigintUnionTransfn(state *Set, i int64) *Set {
	_cret := C.bigint_union_transfn(state._inner, C.int64_t(i))
	return &Set{_inner: _cret}
}


// DateExtentTransfn wraps MEOS C function date_extent_transfn.
func DateExtentTransfn(state *Span, d int32) *Span {
	_cret := C.date_extent_transfn(state._inner, C.DateADT(d))
	return &Span{_inner: _cret}
}


// DateUnionTransfn wraps MEOS C function date_union_transfn.
func DateUnionTransfn(state *Set, d int32) *Set {
	_cret := C.date_union_transfn(state._inner, C.DateADT(d))
	return &Set{_inner: _cret}
}


// FloatExtentTransfn wraps MEOS C function float_extent_transfn.
func FloatExtentTransfn(state *Span, d float64) *Span {
	_cret := C.float_extent_transfn(state._inner, C.double(d))
	return &Span{_inner: _cret}
}


// FloatUnionTransfn wraps MEOS C function float_union_transfn.
func FloatUnionTransfn(state *Set, d float64) *Set {
	_cret := C.float_union_transfn(state._inner, C.double(d))
	return &Set{_inner: _cret}
}


// IntExtentTransfn wraps MEOS C function int_extent_transfn.
func IntExtentTransfn(state *Span, i int) *Span {
	_cret := C.int_extent_transfn(state._inner, C.int(i))
	return &Span{_inner: _cret}
}


// IntUnionTransfn wraps MEOS C function int_union_transfn.
func IntUnionTransfn(state *Set, i int32) *Set {
	_cret := C.int_union_transfn(state._inner, C.int32(i))
	return &Set{_inner: _cret}
}


// SetExtentTransfn wraps MEOS C function set_extent_transfn.
func SetExtentTransfn(state *Span, s *Set) *Span {
	_cret := C.set_extent_transfn(state._inner, s._inner)
	return &Span{_inner: _cret}
}


// SetUnionFinalfn wraps MEOS C function set_union_finalfn.
func SetUnionFinalfn(state *Set) *Set {
	_cret := C.set_union_finalfn(state._inner)
	return &Set{_inner: _cret}
}


// SetUnionTransfn wraps MEOS C function set_union_transfn.
func SetUnionTransfn(state *Set, s *Set) *Set {
	_cret := C.set_union_transfn(state._inner, s._inner)
	return &Set{_inner: _cret}
}


// SpanExtentTransfn wraps MEOS C function span_extent_transfn.
func SpanExtentTransfn(state *Span, s *Span) *Span {
	_cret := C.span_extent_transfn(state._inner, s._inner)
	return &Span{_inner: _cret}
}


// SpanUnionTransfn wraps MEOS C function span_union_transfn.
func SpanUnionTransfn(state *SpanSet, s *Span) *SpanSet {
	_cret := C.span_union_transfn(state._inner, s._inner)
	return &SpanSet{_inner: _cret}
}


// SpansetExtentTransfn wraps MEOS C function spanset_extent_transfn.
func SpansetExtentTransfn(state *Span, ss *SpanSet) *Span {
	_cret := C.spanset_extent_transfn(state._inner, ss._inner)
	return &Span{_inner: _cret}
}


// SpansetUnionFinalfn wraps MEOS C function spanset_union_finalfn.
func SpansetUnionFinalfn(state *SpanSet) *SpanSet {
	_cret := C.spanset_union_finalfn(state._inner)
	return &SpanSet{_inner: _cret}
}


// SpansetUnionTransfn wraps MEOS C function spanset_union_transfn.
func SpansetUnionTransfn(state *SpanSet, ss *SpanSet) *SpanSet {
	_cret := C.spanset_union_transfn(state._inner, ss._inner)
	return &SpanSet{_inner: _cret}
}


// TextUnionTransfn wraps MEOS C function text_union_transfn.
func TextUnionTransfn(state *Set, txt string) *Set {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.text_union_transfn(state._inner, _c_txt)
	return &Set{_inner: _cret}
}


// TimestamptzExtentTransfn wraps MEOS C function timestamptz_extent_transfn.
func TimestamptzExtentTransfn(state *Span, t int64) *Span {
	_cret := C.timestamptz_extent_transfn(state._inner, C.TimestampTz(t))
	return &Span{_inner: _cret}
}


// TimestamptzUnionTransfn wraps MEOS C function timestamptz_union_transfn.
func TimestamptzUnionTransfn(state *Set, t int64) *Set {
	_cret := C.timestamptz_union_transfn(state._inner, C.TimestampTz(t))
	return &Set{_inner: _cret}
}


// BigintGetBin wraps MEOS C function bigint_get_bin.
func BigintGetBin(value int64, vsize int64, vorigin int64) int64 {
	_cret := C.bigint_get_bin(C.int64_t(value), C.int64_t(vsize), C.int64_t(vorigin))
	return int64(_cret)
}


// BigintspanBins wraps MEOS C function bigintspan_bins.
func BigintspanBins(s *Span, vsize int64, vorigin int64, count unsafe.Pointer) *Span {
	_cret := C.bigintspan_bins(s._inner, C.int64_t(vsize), C.int64_t(vorigin), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// BigintspansetBins wraps MEOS C function bigintspanset_bins.
func BigintspansetBins(ss *SpanSet, vsize int64, vorigin int64, count unsafe.Pointer) *Span {
	_cret := C.bigintspanset_bins(ss._inner, C.int64_t(vsize), C.int64_t(vorigin), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// DateGetBin wraps MEOS C function date_get_bin.
func DateGetBin(d int32, duration *Interval, torigin int32) int32 {
	_cret := C.date_get_bin(C.DateADT(d), duration._inner, C.DateADT(torigin))
	return int32(_cret)
}


// DatespanBins wraps MEOS C function datespan_bins.
func DatespanBins(s *Span, duration *Interval, torigin int32, count unsafe.Pointer) *Span {
	_cret := C.datespan_bins(s._inner, duration._inner, C.DateADT(torigin), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// DatespansetBins wraps MEOS C function datespanset_bins.
func DatespansetBins(ss *SpanSet, duration *Interval, torigin int32, count unsafe.Pointer) *Span {
	_cret := C.datespanset_bins(ss._inner, duration._inner, C.DateADT(torigin), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// FloatGetBin wraps MEOS C function float_get_bin.
func FloatGetBin(value float64, vsize float64, vorigin float64) float64 {
	_cret := C.float_get_bin(C.double(value), C.double(vsize), C.double(vorigin))
	return float64(_cret)
}


// FloatspanBins wraps MEOS C function floatspan_bins.
func FloatspanBins(s *Span, vsize float64, vorigin float64, count unsafe.Pointer) *Span {
	_cret := C.floatspan_bins(s._inner, C.double(vsize), C.double(vorigin), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// FloatspansetBins wraps MEOS C function floatspanset_bins.
func FloatspansetBins(ss *SpanSet, vsize float64, vorigin float64, count unsafe.Pointer) *Span {
	_cret := C.floatspanset_bins(ss._inner, C.double(vsize), C.double(vorigin), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// IntGetBin wraps MEOS C function int_get_bin.
func IntGetBin(value int, vsize int, vorigin int) int {
	_cret := C.int_get_bin(C.int(value), C.int(vsize), C.int(vorigin))
	return int(_cret)
}


// IntspanBins wraps MEOS C function intspan_bins.
func IntspanBins(s *Span, vsize int, vorigin int, count unsafe.Pointer) *Span {
	_cret := C.intspan_bins(s._inner, C.int(vsize), C.int(vorigin), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// IntspansetBins wraps MEOS C function intspanset_bins.
func IntspansetBins(ss *SpanSet, vsize int, vorigin int, count unsafe.Pointer) *Span {
	_cret := C.intspanset_bins(ss._inner, C.int(vsize), C.int(vorigin), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// TimestamptzGetBin wraps MEOS C function timestamptz_get_bin.
func TimestamptzGetBin(t int64, duration *Interval, torigin int64) int64 {
	_cret := C.timestamptz_get_bin(C.TimestampTz(t), duration._inner, C.TimestampTz(torigin))
	return int64(_cret)
}


// TstzspanBins wraps MEOS C function tstzspan_bins.
func TstzspanBins(s *Span, duration *Interval, origin int64, count unsafe.Pointer) *Span {
	_cret := C.tstzspan_bins(s._inner, duration._inner, C.TimestampTz(origin), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// TstzspansetBins wraps MEOS C function tstzspanset_bins.
func TstzspansetBins(ss *SpanSet, duration *Interval, torigin int64, count unsafe.Pointer) *Span {
	_cret := C.tstzspanset_bins(ss._inner, duration._inner, C.TimestampTz(torigin), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// TBOXAsHexwkb wraps MEOS C function tbox_as_hexwkb.
func TBOXAsHexwkb(box *TBox, variant uint8, size_out unsafe.Pointer) string {
	_cret := C.tbox_as_hexwkb(box._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return C.GoString(_cret)
}


// TBOXAsWKB wraps MEOS C function tbox_as_wkb.
func TBOXAsWKB(box *TBox, variant uint8, size_out unsafe.Pointer) unsafe.Pointer {
	_cret := C.tbox_as_wkb(box._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return unsafe.Pointer(_cret)
}


// TBOXFromHexwkb wraps MEOS C function tbox_from_hexwkb.
func TBOXFromHexwkb(hexwkb string) *TBox {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	_cret := C.tbox_from_hexwkb(_c_hexwkb)
	return &TBox{_inner: _cret}
}


// TBOXFromWKB wraps MEOS C function tbox_from_wkb.
func TBOXFromWKB(wkb unsafe.Pointer, size uint) *TBox {
	_cret := C.tbox_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	return &TBox{_inner: _cret}
}


// TBOXIn wraps MEOS C function tbox_in.
func TBOXIn(str string) *TBox {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tbox_in(_c_str)
	return &TBox{_inner: _cret}
}


// TBOXOut wraps MEOS C function tbox_out.
func TBOXOut(box *TBox, maxdd int) string {
	_cret := C.tbox_out(box._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// FloatTimestamptzToTBOX wraps MEOS C function float_timestamptz_to_tbox.
func FloatTimestamptzToTBOX(d float64, t int64) *TBox {
	_cret := C.float_timestamptz_to_tbox(C.double(d), C.TimestampTz(t))
	return &TBox{_inner: _cret}
}


// FloatTstzspanToTBOX wraps MEOS C function float_tstzspan_to_tbox.
func FloatTstzspanToTBOX(d float64, s *Span) *TBox {
	_cret := C.float_tstzspan_to_tbox(C.double(d), s._inner)
	return &TBox{_inner: _cret}
}


// IntTimestamptzToTBOX wraps MEOS C function int_timestamptz_to_tbox.
func IntTimestamptzToTBOX(i int, t int64) *TBox {
	_cret := C.int_timestamptz_to_tbox(C.int(i), C.TimestampTz(t))
	return &TBox{_inner: _cret}
}


// BigintTimestamptzToTBOX wraps MEOS C function bigint_timestamptz_to_tbox.
func BigintTimestamptzToTBOX(i int64, t int64) *TBox {
	_cret := C.bigint_timestamptz_to_tbox(C.int64_t(i), C.TimestampTz(t))
	return &TBox{_inner: _cret}
}


// IntTstzspanToTBOX wraps MEOS C function int_tstzspan_to_tbox.
func IntTstzspanToTBOX(i int, s *Span) *TBox {
	_cret := C.int_tstzspan_to_tbox(C.int(i), s._inner)
	return &TBox{_inner: _cret}
}


// BigintTstzspanToTBOX wraps MEOS C function bigint_tstzspan_to_tbox.
func BigintTstzspanToTBOX(i int64, s *Span) *TBox {
	_cret := C.bigint_tstzspan_to_tbox(C.int64_t(i), s._inner)
	return &TBox{_inner: _cret}
}


// NumspanTstzspanToTBOX wraps MEOS C function numspan_tstzspan_to_tbox.
func NumspanTstzspanToTBOX(span *Span, s *Span) *TBox {
	_cret := C.numspan_tstzspan_to_tbox(span._inner, s._inner)
	return &TBox{_inner: _cret}
}


// NumspanTimestamptzToTBOX wraps MEOS C function numspan_timestamptz_to_tbox.
func NumspanTimestamptzToTBOX(span *Span, t int64) *TBox {
	_cret := C.numspan_timestamptz_to_tbox(span._inner, C.TimestampTz(t))
	return &TBox{_inner: _cret}
}


// TBOXCopy wraps MEOS C function tbox_copy.
func TBOXCopy(box *TBox) *TBox {
	_cret := C.tbox_copy(box._inner)
	return &TBox{_inner: _cret}
}


// TBOXMake wraps MEOS C function tbox_make.
func TBOXMake(s *Span, p *Span) *TBox {
	_cret := C.tbox_make(s._inner, p._inner)
	return &TBox{_inner: _cret}
}


// FloatToTBOX wraps MEOS C function float_to_tbox.
func FloatToTBOX(d float64) *TBox {
	_cret := C.float_to_tbox(C.double(d))
	return &TBox{_inner: _cret}
}


// IntToTBOX wraps MEOS C function int_to_tbox.
func IntToTBOX(i int) *TBox {
	_cret := C.int_to_tbox(C.int(i))
	return &TBox{_inner: _cret}
}


// BigintToTBOX wraps MEOS C function bigint_to_tbox.
func BigintToTBOX(i int64) *TBox {
	_cret := C.bigint_to_tbox(C.int64_t(i))
	return &TBox{_inner: _cret}
}


// SetToTBOX wraps MEOS C function set_to_tbox.
func SetToTBOX(s *Set) *TBox {
	_cret := C.set_to_tbox(s._inner)
	return &TBox{_inner: _cret}
}


// SpanToTBOX wraps MEOS C function span_to_tbox.
func SpanToTBOX(s *Span) *TBox {
	_cret := C.span_to_tbox(s._inner)
	return &TBox{_inner: _cret}
}


// SpansetToTBOX wraps MEOS C function spanset_to_tbox.
func SpansetToTBOX(ss *SpanSet) *TBox {
	_cret := C.spanset_to_tbox(ss._inner)
	return &TBox{_inner: _cret}
}


// TBOXToIntspan wraps MEOS C function tbox_to_intspan.
func TBOXToIntspan(box *TBox) *Span {
	_cret := C.tbox_to_intspan(box._inner)
	return &Span{_inner: _cret}
}


// TBOXToBigintspan wraps MEOS C function tbox_to_bigintspan.
func TBOXToBigintspan(box *TBox) *Span {
	_cret := C.tbox_to_bigintspan(box._inner)
	return &Span{_inner: _cret}
}


// TBOXToFloatspan wraps MEOS C function tbox_to_floatspan.
func TBOXToFloatspan(box *TBox) *Span {
	_cret := C.tbox_to_floatspan(box._inner)
	return &Span{_inner: _cret}
}


// TBOXToTstzspan wraps MEOS C function tbox_to_tstzspan.
func TBOXToTstzspan(box *TBox) *Span {
	_cret := C.tbox_to_tstzspan(box._inner)
	return &Span{_inner: _cret}
}


// TimestamptzToTBOX wraps MEOS C function timestamptz_to_tbox.
func TimestamptzToTBOX(t int64) *TBox {
	_cret := C.timestamptz_to_tbox(C.TimestampTz(t))
	return &TBox{_inner: _cret}
}


// TBOXHash wraps MEOS C function tbox_hash.
func TBOXHash(box *TBox) uint32 {
	_cret := C.tbox_hash(box._inner)
	return uint32(_cret)
}


// TBOXHashExtended wraps MEOS C function tbox_hash_extended.
func TBOXHashExtended(box *TBox, seed uint64) uint64 {
	_cret := C.tbox_hash_extended(box._inner, C.uint64_t(seed))
	return uint64(_cret)
}


// TBOXHast wraps MEOS C function tbox_hast.
func TBOXHast(box *TBox) bool {
	_cret := C.tbox_hast(box._inner)
	return bool(_cret)
}


// TBOXHasx wraps MEOS C function tbox_hasx.
func TBOXHasx(box *TBox) bool {
	_cret := C.tbox_hasx(box._inner)
	return bool(_cret)
}


// TBOXTmax wraps MEOS C function tbox_tmax.
func TBOXTmax(box *TBox) (bool, int64) {
	var _out_result C.TimestampTz
	_cret := C.tbox_tmax(box._inner, &_out_result)
	return bool(_cret), int64(_out_result)
}


// TBOXTmaxInc wraps MEOS C function tbox_tmax_inc.
func TBOXTmaxInc(box *TBox) (bool, bool) {
	var _out_result C.bool
	_cret := C.tbox_tmax_inc(box._inner, &_out_result)
	return bool(_cret), bool(_out_result)
}


// TBOXTmin wraps MEOS C function tbox_tmin.
func TBOXTmin(box *TBox) (bool, int64) {
	var _out_result C.TimestampTz
	_cret := C.tbox_tmin(box._inner, &_out_result)
	return bool(_cret), int64(_out_result)
}


// TBOXTminInc wraps MEOS C function tbox_tmin_inc.
func TBOXTminInc(box *TBox) (bool, bool) {
	var _out_result C.bool
	_cret := C.tbox_tmin_inc(box._inner, &_out_result)
	return bool(_cret), bool(_out_result)
}


// TBOXXmax wraps MEOS C function tbox_xmax.
func TBOXXmax(box *TBox) (bool, float64) {
	var _out_result C.double
	_cret := C.tbox_xmax(box._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// TBOXXmaxInc wraps MEOS C function tbox_xmax_inc.
func TBOXXmaxInc(box *TBox) (bool, bool) {
	var _out_result C.bool
	_cret := C.tbox_xmax_inc(box._inner, &_out_result)
	return bool(_cret), bool(_out_result)
}


// TBOXXmin wraps MEOS C function tbox_xmin.
func TBOXXmin(box *TBox) (bool, float64) {
	var _out_result C.double
	_cret := C.tbox_xmin(box._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// TBOXXminInc wraps MEOS C function tbox_xmin_inc.
func TBOXXminInc(box *TBox) (bool, bool) {
	var _out_result C.bool
	_cret := C.tbox_xmin_inc(box._inner, &_out_result)
	return bool(_cret), bool(_out_result)
}


// TboxfloatXmax wraps MEOS C function tboxfloat_xmax.
func TboxfloatXmax(box *TBox) (bool, float64) {
	var _out_result C.double
	_cret := C.tboxfloat_xmax(box._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// TboxfloatXmin wraps MEOS C function tboxfloat_xmin.
func TboxfloatXmin(box *TBox) (bool, float64) {
	var _out_result C.double
	_cret := C.tboxfloat_xmin(box._inner, &_out_result)
	return bool(_cret), float64(_out_result)
}


// TboxintXmax wraps MEOS C function tboxint_xmax.
func TboxintXmax(box *TBox) (bool, int) {
	var _out_result C.int
	_cret := C.tboxint_xmax(box._inner, &_out_result)
	return bool(_cret), int(_out_result)
}


// TboxbigintXmax wraps MEOS C function tboxbigint_xmax.
func TboxbigintXmax(box *TBox) (bool, int64) {
	var _out_result C.int64_t
	_cret := C.tboxbigint_xmax(box._inner, &_out_result)
	return bool(_cret), int64(_out_result)
}


// TboxintXmin wraps MEOS C function tboxint_xmin.
func TboxintXmin(box *TBox) (bool, int) {
	var _out_result C.int
	_cret := C.tboxint_xmin(box._inner, &_out_result)
	return bool(_cret), int(_out_result)
}


// TboxbigintXmin wraps MEOS C function tboxbigint_xmin.
func TboxbigintXmin(box *TBox) (bool, int64) {
	var _out_result C.int64_t
	_cret := C.tboxbigint_xmin(box._inner, &_out_result)
	return bool(_cret), int64(_out_result)
}


// TfloatboxExpand wraps MEOS C function tfloatbox_expand.
func TfloatboxExpand(box *TBox, d float64) *TBox {
	_cret := C.tfloatbox_expand(box._inner, C.double(d))
	return &TBox{_inner: _cret}
}


// TintboxExpand wraps MEOS C function tintbox_expand.
func TintboxExpand(box *TBox, i int) *TBox {
	_cret := C.tintbox_expand(box._inner, C.int(i))
	return &TBox{_inner: _cret}
}


// TBOXExpandTime wraps MEOS C function tbox_expand_time.
func TBOXExpandTime(box *TBox, interv *Interval) *TBox {
	_cret := C.tbox_expand_time(box._inner, interv._inner)
	return &TBox{_inner: _cret}
}


// TBOXRound wraps MEOS C function tbox_round.
func TBOXRound(box *TBox, maxdd int) *TBox {
	_cret := C.tbox_round(box._inner, C.int(maxdd))
	return &TBox{_inner: _cret}
}


// TfloatboxShiftScale wraps MEOS C function tfloatbox_shift_scale.
func TfloatboxShiftScale(box *TBox, shift float64, width float64, hasshift bool, haswidth bool) *TBox {
	_cret := C.tfloatbox_shift_scale(box._inner, C.double(shift), C.double(width), C.bool(hasshift), C.bool(haswidth))
	return &TBox{_inner: _cret}
}


// TintboxShiftScale wraps MEOS C function tintbox_shift_scale.
func TintboxShiftScale(box *TBox, shift int, width int, hasshift bool, haswidth bool) *TBox {
	_cret := C.tintbox_shift_scale(box._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	return &TBox{_inner: _cret}
}


// TBOXShiftScaleTime wraps MEOS C function tbox_shift_scale_time.
func TBOXShiftScaleTime(box *TBox, shift *Interval, duration *Interval) *TBox {
	_cret := C.tbox_shift_scale_time(box._inner, shift._inner, duration._inner)
	return &TBox{_inner: _cret}
}


// TbigintboxExpand wraps MEOS C function tbigintbox_expand.
func TbigintboxExpand(box *TBox, i int64) *TBox {
	_cret := C.tbigintbox_expand(box._inner, C.int64_t(i))
	return &TBox{_inner: _cret}
}


// TbigintboxShiftScale wraps MEOS C function tbigintbox_shift_scale.
func TbigintboxShiftScale(box *TBox, shift int64, width int64, hasshift bool, haswidth bool) *TBox {
	_cret := C.tbigintbox_shift_scale(box._inner, C.int64_t(shift), C.int64_t(width), C.bool(hasshift), C.bool(haswidth))
	return &TBox{_inner: _cret}
}


// UnionTBOXTBOX wraps MEOS C function union_tbox_tbox.
func UnionTBOXTBOX(box1 *TBox, box2 *TBox, strict bool) *TBox {
	_cret := C.gunion_tbox_tbox(box1._inner, box2._inner, C.bool(strict))
	return &TBox{_inner: _cret}
}


// IntersectionTBOXTBOX wraps MEOS C function intersection_tbox_tbox.
func IntersectionTBOXTBOX(box1 *TBox, box2 *TBox) *TBox {
	_cret := C.intersection_tbox_tbox(box1._inner, box2._inner)
	return &TBox{_inner: _cret}
}


// AdjacentTBOXTBOX wraps MEOS C function adjacent_tbox_tbox.
func AdjacentTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	_cret := C.adjacent_tbox_tbox(box1._inner, box2._inner)
	return bool(_cret)
}


// ContainedTBOXTBOX wraps MEOS C function contained_tbox_tbox.
func ContainedTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	_cret := C.contained_tbox_tbox(box1._inner, box2._inner)
	return bool(_cret)
}


// ContainsTBOXTBOX wraps MEOS C function contains_tbox_tbox.
func ContainsTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	_cret := C.contains_tbox_tbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverlapsTBOXTBOX wraps MEOS C function overlaps_tbox_tbox.
func OverlapsTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	_cret := C.overlaps_tbox_tbox(box1._inner, box2._inner)
	return bool(_cret)
}


// SameTBOXTBOX wraps MEOS C function same_tbox_tbox.
func SameTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	_cret := C.same_tbox_tbox(box1._inner, box2._inner)
	return bool(_cret)
}


// AfterTBOXTBOX wraps MEOS C function after_tbox_tbox.
func AfterTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	_cret := C.after_tbox_tbox(box1._inner, box2._inner)
	return bool(_cret)
}


// BeforeTBOXTBOX wraps MEOS C function before_tbox_tbox.
func BeforeTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	_cret := C.before_tbox_tbox(box1._inner, box2._inner)
	return bool(_cret)
}


// LeftTBOXTBOX wraps MEOS C function left_tbox_tbox.
func LeftTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	_cret := C.left_tbox_tbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverafterTBOXTBOX wraps MEOS C function overafter_tbox_tbox.
func OverafterTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	_cret := C.overafter_tbox_tbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverbeforeTBOXTBOX wraps MEOS C function overbefore_tbox_tbox.
func OverbeforeTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	_cret := C.overbefore_tbox_tbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverleftTBOXTBOX wraps MEOS C function overleft_tbox_tbox.
func OverleftTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	_cret := C.overleft_tbox_tbox(box1._inner, box2._inner)
	return bool(_cret)
}


// OverrightTBOXTBOX wraps MEOS C function overright_tbox_tbox.
func OverrightTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	_cret := C.overright_tbox_tbox(box1._inner, box2._inner)
	return bool(_cret)
}


// RightTBOXTBOX wraps MEOS C function right_tbox_tbox.
func RightTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	_cret := C.right_tbox_tbox(box1._inner, box2._inner)
	return bool(_cret)
}


// TBOXCmp wraps MEOS C function tbox_cmp.
func TBOXCmp(box1 *TBox, box2 *TBox) int {
	_cret := C.tbox_cmp(box1._inner, box2._inner)
	return int(_cret)
}


// TBOXEq wraps MEOS C function tbox_eq.
func TBOXEq(box1 *TBox, box2 *TBox) bool {
	_cret := C.tbox_eq(box1._inner, box2._inner)
	return bool(_cret)
}


// TBOXGe wraps MEOS C function tbox_ge.
func TBOXGe(box1 *TBox, box2 *TBox) bool {
	_cret := C.tbox_ge(box1._inner, box2._inner)
	return bool(_cret)
}


// TBOXGt wraps MEOS C function tbox_gt.
func TBOXGt(box1 *TBox, box2 *TBox) bool {
	_cret := C.tbox_gt(box1._inner, box2._inner)
	return bool(_cret)
}


// TBOXLe wraps MEOS C function tbox_le.
func TBOXLe(box1 *TBox, box2 *TBox) bool {
	_cret := C.tbox_le(box1._inner, box2._inner)
	return bool(_cret)
}


// TBOXLt wraps MEOS C function tbox_lt.
func TBOXLt(box1 *TBox, box2 *TBox) bool {
	_cret := C.tbox_lt(box1._inner, box2._inner)
	return bool(_cret)
}


// TBOXNe wraps MEOS C function tbox_ne.
func TBOXNe(box1 *TBox, box2 *TBox) bool {
	_cret := C.tbox_ne(box1._inner, box2._inner)
	return bool(_cret)
}


// TboolFromMFJSON wraps MEOS C function tbool_from_mfjson.
func TboolFromMFJSON(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tbool_from_mfjson(_c_str)
	return &Temporal{_inner: _cret}
}


// TboolIn wraps MEOS C function tbool_in.
func TboolIn(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tbool_in(_c_str)
	return &Temporal{_inner: _cret}
}


// TboolOut wraps MEOS C function tbool_out.
func TboolOut(temp *Temporal) string {
	_cret := C.tbool_out(temp._inner)
	return C.GoString(_cret)
}


// TemporalAsHexwkb wraps MEOS C function temporal_as_hexwkb.
func TemporalAsHexwkb(temp *Temporal, variant uint8, size_out unsafe.Pointer) string {
	_cret := C.temporal_as_hexwkb(temp._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return C.GoString(_cret)
}


// TemporalAsMFJSON wraps MEOS C function temporal_as_mfjson.
func TemporalAsMFJSON(temp *Temporal, with_bbox bool, flags int, precision int, srs string) string {
	_c_srs := C.CString(srs)
	defer C.free(unsafe.Pointer(_c_srs))
	_cret := C.temporal_as_mfjson(temp._inner, C.bool(with_bbox), C.int(flags), C.int(precision), _c_srs)
	return C.GoString(_cret)
}


// TemporalAsWKB wraps MEOS C function temporal_as_wkb.
func TemporalAsWKB(temp *Temporal, variant uint8, size_out unsafe.Pointer) unsafe.Pointer {
	_cret := C.temporal_as_wkb(temp._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	return unsafe.Pointer(_cret)
}


// TemporalFromHexwkb wraps MEOS C function temporal_from_hexwkb.
func TemporalFromHexwkb(hexwkb string) *Temporal {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	_cret := C.temporal_from_hexwkb(_c_hexwkb)
	return &Temporal{_inner: _cret}
}


// TemporalFromWKB wraps MEOS C function temporal_from_wkb.
func TemporalFromWKB(wkb unsafe.Pointer, size uint) *Temporal {
	_cret := C.temporal_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	return &Temporal{_inner: _cret}
}


// TfloatFromMFJSON wraps MEOS C function tfloat_from_mfjson.
func TfloatFromMFJSON(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tfloat_from_mfjson(_c_str)
	return &Temporal{_inner: _cret}
}


// TfloatIn wraps MEOS C function tfloat_in.
func TfloatIn(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tfloat_in(_c_str)
	return &Temporal{_inner: _cret}
}


// TfloatOut wraps MEOS C function tfloat_out.
func TfloatOut(temp *Temporal, maxdd int) string {
	_cret := C.tfloat_out(temp._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// TintFromMFJSON wraps MEOS C function tint_from_mfjson.
func TintFromMFJSON(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tint_from_mfjson(_c_str)
	return &Temporal{_inner: _cret}
}


// TbigintFromMFJSON wraps MEOS C function tbigint_from_mfjson.
func TbigintFromMFJSON(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tbigint_from_mfjson(_c_str)
	return &Temporal{_inner: _cret}
}


// TintIn wraps MEOS C function tint_in.
func TintIn(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tint_in(_c_str)
	return &Temporal{_inner: _cret}
}


// TbigintIn wraps MEOS C function tbigint_in.
func TbigintIn(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tbigint_in(_c_str)
	return &Temporal{_inner: _cret}
}


// TintOut wraps MEOS C function tint_out.
func TintOut(temp *Temporal) string {
	_cret := C.tint_out(temp._inner)
	return C.GoString(_cret)
}


// TbigintOut wraps MEOS C function tbigint_out.
func TbigintOut(temp *Temporal) string {
	_cret := C.tbigint_out(temp._inner)
	return C.GoString(_cret)
}


// TtextFromMFJSON wraps MEOS C function ttext_from_mfjson.
func TtextFromMFJSON(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.ttext_from_mfjson(_c_str)
	return &Temporal{_inner: _cret}
}


// TtextIn wraps MEOS C function ttext_in.
func TtextIn(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.ttext_in(_c_str)
	return &Temporal{_inner: _cret}
}


// TtextOut wraps MEOS C function ttext_out.
func TtextOut(temp *Temporal) string {
	_cret := C.ttext_out(temp._inner)
	return C.GoString(_cret)
}


// TboolFromBaseTemp wraps MEOS C function tbool_from_base_temp.
func TboolFromBaseTemp(b bool, temp *Temporal) *Temporal {
	_cret := C.tbool_from_base_temp(C.bool(b), temp._inner)
	return &Temporal{_inner: _cret}
}


// TboolinstMake wraps MEOS C function tboolinst_make.
func TboolinstMake(b bool, t int64) *TInstant {
	_cret := C.tboolinst_make(C.bool(b), C.TimestampTz(t))
	return &TInstant{_inner: _cret}
}


// TboolseqFromBaseTstzset wraps MEOS C function tboolseq_from_base_tstzset.
func TboolseqFromBaseTstzset(b bool, s *Set) *TSequence {
	_cret := C.tboolseq_from_base_tstzset(C.bool(b), s._inner)
	return &TSequence{_inner: _cret}
}


// TboolseqFromBaseTstzspan wraps MEOS C function tboolseq_from_base_tstzspan.
func TboolseqFromBaseTstzspan(b bool, s *Span) *TSequence {
	_cret := C.tboolseq_from_base_tstzspan(C.bool(b), s._inner)
	return &TSequence{_inner: _cret}
}


// TboolseqsetFromBaseTstzspanset wraps MEOS C function tboolseqset_from_base_tstzspanset.
func TboolseqsetFromBaseTstzspanset(b bool, ss *SpanSet) *TSequenceSet {
	_cret := C.tboolseqset_from_base_tstzspanset(C.bool(b), ss._inner)
	return &TSequenceSet{_inner: _cret}
}


// TemporalCopy wraps MEOS C function temporal_copy.
func TemporalCopy(temp *Temporal) *Temporal {
	_cret := C.temporal_copy(temp._inner)
	return &Temporal{_inner: _cret}
}


// TfloatFromBaseTemp wraps MEOS C function tfloat_from_base_temp.
func TfloatFromBaseTemp(d float64, temp *Temporal) *Temporal {
	_cret := C.tfloat_from_base_temp(C.double(d), temp._inner)
	return &Temporal{_inner: _cret}
}


// TfloatinstMake wraps MEOS C function tfloatinst_make.
func TfloatinstMake(d float64, t int64) *TInstant {
	_cret := C.tfloatinst_make(C.double(d), C.TimestampTz(t))
	return &TInstant{_inner: _cret}
}


// TfloatseqFromBaseTstzset wraps MEOS C function tfloatseq_from_base_tstzset.
func TfloatseqFromBaseTstzset(d float64, s *Set) *TSequence {
	_cret := C.tfloatseq_from_base_tstzset(C.double(d), s._inner)
	return &TSequence{_inner: _cret}
}


// TfloatseqFromBaseTstzspan wraps MEOS C function tfloatseq_from_base_tstzspan.
func TfloatseqFromBaseTstzspan(d float64, s *Span, interp Interpolation) *TSequence {
	_cret := C.tfloatseq_from_base_tstzspan(C.double(d), s._inner, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TfloatseqsetFromBaseTstzspanset wraps MEOS C function tfloatseqset_from_base_tstzspanset.
func TfloatseqsetFromBaseTstzspanset(d float64, ss *SpanSet, interp Interpolation) *TSequenceSet {
	_cret := C.tfloatseqset_from_base_tstzspanset(C.double(d), ss._inner, C.interpType(interp))
	return &TSequenceSet{_inner: _cret}
}


// TintFromBaseTemp wraps MEOS C function tint_from_base_temp.
func TintFromBaseTemp(i int, temp *Temporal) *Temporal {
	_cret := C.tint_from_base_temp(C.int(i), temp._inner)
	return &Temporal{_inner: _cret}
}


// TbigintFromBaseTemp wraps MEOS C function tbigint_from_base_temp.
func TbigintFromBaseTemp(i int64, temp *Temporal) *Temporal {
	_cret := C.tbigint_from_base_temp(C.int64_t(i), temp._inner)
	return &Temporal{_inner: _cret}
}


// TintinstMake wraps MEOS C function tintinst_make.
func TintinstMake(i int, t int64) *TInstant {
	_cret := C.tintinst_make(C.int(i), C.TimestampTz(t))
	return &TInstant{_inner: _cret}
}


// TbigintinstMake wraps MEOS C function tbigintinst_make.
func TbigintinstMake(i int64, t int64) *TInstant {
	_cret := C.tbigintinst_make(C.int64_t(i), C.TimestampTz(t))
	return &TInstant{_inner: _cret}
}


// TintseqFromBaseTstzset wraps MEOS C function tintseq_from_base_tstzset.
func TintseqFromBaseTstzset(i int, s *Set) *TSequence {
	_cret := C.tintseq_from_base_tstzset(C.int(i), s._inner)
	return &TSequence{_inner: _cret}
}


// TbigintseqFromBaseTstzset wraps MEOS C function tbigintseq_from_base_tstzset.
func TbigintseqFromBaseTstzset(i int64, s *Set) *TSequence {
	_cret := C.tbigintseq_from_base_tstzset(C.int64_t(i), s._inner)
	return &TSequence{_inner: _cret}
}


// TintseqFromBaseTstzspan wraps MEOS C function tintseq_from_base_tstzspan.
func TintseqFromBaseTstzspan(i int, s *Span) *TSequence {
	_cret := C.tintseq_from_base_tstzspan(C.int(i), s._inner)
	return &TSequence{_inner: _cret}
}


// TbigintseqFromBaseTstzspan wraps MEOS C function tbigintseq_from_base_tstzspan.
func TbigintseqFromBaseTstzspan(i int64, s *Span) *TSequence {
	_cret := C.tbigintseq_from_base_tstzspan(C.int64_t(i), s._inner)
	return &TSequence{_inner: _cret}
}


// TintseqsetFromBaseTstzspanset wraps MEOS C function tintseqset_from_base_tstzspanset.
func TintseqsetFromBaseTstzspanset(i int, ss *SpanSet) *TSequenceSet {
	_cret := C.tintseqset_from_base_tstzspanset(C.int(i), ss._inner)
	return &TSequenceSet{_inner: _cret}
}


// TbigintseqsetFromBaseTstzspanset wraps MEOS C function tbigintseqset_from_base_tstzspanset.
func TbigintseqsetFromBaseTstzspanset(i int64, ss *SpanSet) *TSequenceSet {
	_cret := C.tbigintseqset_from_base_tstzspanset(C.int64_t(i), ss._inner)
	return &TSequenceSet{_inner: _cret}
}


// TsequenceMake wraps MEOS C function tsequence_make.
func TsequenceMake(instants unsafe.Pointer, count int, lower_inc bool, upper_inc bool, interp Interpolation, normalize bool) *TSequence {
	_cret := C.tsequence_make((**C.TInstant)(unsafe.Pointer(instants)), C.int(count), C.bool(lower_inc), C.bool(upper_inc), C.interpType(interp), C.bool(normalize))
	return &TSequence{_inner: _cret}
}


// TsequencesetMake wraps MEOS C function tsequenceset_make.
func TsequencesetMake(sequences unsafe.Pointer, count int, normalize bool) *TSequenceSet {
	_cret := C.tsequenceset_make((**C.TSequence)(unsafe.Pointer(sequences)), C.int(count), C.bool(normalize))
	return &TSequenceSet{_inner: _cret}
}


// TsequencesetMakeGaps wraps MEOS C function tsequenceset_make_gaps.
func TsequencesetMakeGaps(instants unsafe.Pointer, count int, interp Interpolation, maxt *Interval, maxdist float64) *TSequenceSet {
	_cret := C.tsequenceset_make_gaps((**C.TInstant)(unsafe.Pointer(instants)), C.int(count), C.interpType(interp), maxt._inner, C.double(maxdist))
	return &TSequenceSet{_inner: _cret}
}


// TtextFromBaseTemp wraps MEOS C function ttext_from_base_temp.
func TtextFromBaseTemp(txt string, temp *Temporal) *Temporal {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ttext_from_base_temp(_c_txt, temp._inner)
	return &Temporal{_inner: _cret}
}


// TtextinstMake wraps MEOS C function ttextinst_make.
func TtextinstMake(txt string, t int64) *TInstant {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ttextinst_make(_c_txt, C.TimestampTz(t))
	return &TInstant{_inner: _cret}
}


// TtextseqFromBaseTstzset wraps MEOS C function ttextseq_from_base_tstzset.
func TtextseqFromBaseTstzset(txt string, s *Set) *TSequence {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ttextseq_from_base_tstzset(_c_txt, s._inner)
	return &TSequence{_inner: _cret}
}


// TtextseqFromBaseTstzspan wraps MEOS C function ttextseq_from_base_tstzspan.
func TtextseqFromBaseTstzspan(txt string, s *Span) *TSequence {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ttextseq_from_base_tstzspan(_c_txt, s._inner)
	return &TSequence{_inner: _cret}
}


// TtextseqsetFromBaseTstzspanset wraps MEOS C function ttextseqset_from_base_tstzspanset.
func TtextseqsetFromBaseTstzspanset(txt string, ss *SpanSet) *TSequenceSet {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ttextseqset_from_base_tstzspanset(_c_txt, ss._inner)
	return &TSequenceSet{_inner: _cret}
}


// TboolToTint wraps MEOS C function tbool_to_tint.
func TboolToTint(temp *Temporal) *Temporal {
	_cret := C.tbool_to_tint(temp._inner)
	return &Temporal{_inner: _cret}
}


// TemporalToTstzspan wraps MEOS C function temporal_to_tstzspan.
func TemporalToTstzspan(temp *Temporal) *Span {
	_cret := C.temporal_to_tstzspan(temp._inner)
	return &Span{_inner: _cret}
}


// TfloatToTint wraps MEOS C function tfloat_to_tint.
func TfloatToTint(temp *Temporal) *Temporal {
	_cret := C.tfloat_to_tint(temp._inner)
	return &Temporal{_inner: _cret}
}


// TfloatToTbigint wraps MEOS C function tfloat_to_tbigint.
func TfloatToTbigint(temp *Temporal) *Temporal {
	_cret := C.tfloat_to_tbigint(temp._inner)
	return &Temporal{_inner: _cret}
}


// TintToTfloat wraps MEOS C function tint_to_tfloat.
func TintToTfloat(temp *Temporal) *Temporal {
	_cret := C.tint_to_tfloat(temp._inner)
	return &Temporal{_inner: _cret}
}


// TintToTbigint wraps MEOS C function tint_to_tbigint.
func TintToTbigint(temp *Temporal) *Temporal {
	_cret := C.tint_to_tbigint(temp._inner)
	return &Temporal{_inner: _cret}
}


// TbigintToTint wraps MEOS C function tbigint_to_tint.
func TbigintToTint(temp *Temporal) *Temporal {
	_cret := C.tbigint_to_tint(temp._inner)
	return &Temporal{_inner: _cret}
}


// TbigintToTfloat wraps MEOS C function tbigint_to_tfloat.
func TbigintToTfloat(temp *Temporal) *Temporal {
	_cret := C.tbigint_to_tfloat(temp._inner)
	return &Temporal{_inner: _cret}
}


// TnumberToSpan wraps MEOS C function tnumber_to_span.
func TnumberToSpan(temp *Temporal) *Span {
	_cret := C.tnumber_to_span(temp._inner)
	return &Span{_inner: _cret}
}


// TnumberToTBOX wraps MEOS C function tnumber_to_tbox.
func TnumberToTBOX(temp *Temporal) *TBox {
	_cret := C.tnumber_to_tbox(temp._inner)
	return &TBox{_inner: _cret}
}


// TboolEndValue wraps MEOS C function tbool_end_value.
func TboolEndValue(temp *Temporal) bool {
	_cret := C.tbool_end_value(temp._inner)
	return bool(_cret)
}


// TboolStartValue wraps MEOS C function tbool_start_value.
func TboolStartValue(temp *Temporal) bool {
	_cret := C.tbool_start_value(temp._inner)
	return bool(_cret)
}


// TboolValueAtTimestamptz wraps MEOS C function tbool_value_at_timestamptz.
func TboolValueAtTimestamptz(temp *Temporal, t int64, strict bool) (bool, bool) {
	var _out_value C.bool
	_cret := C.tbool_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_value)
	return bool(_cret), bool(_out_value)
}


// TboolValueN wraps MEOS C function tbool_value_n.
func TboolValueN(temp *Temporal, n int) (bool, bool) {
	var _out_result C.bool
	_cret := C.tbool_value_n(temp._inner, C.int(n), &_out_result)
	return bool(_cret), bool(_out_result)
}


// TboolValues wraps MEOS C function tbool_values.
func TboolValues(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tbool_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TemporalDuration wraps MEOS C function temporal_duration.
func TemporalDuration(temp *Temporal, boundspan bool) *Interval {
	_cret := C.temporal_duration(temp._inner, C.bool(boundspan))
	return &Interval{_inner: _cret}
}


// TemporalEndInstant wraps MEOS C function temporal_end_instant.
func TemporalEndInstant(temp *Temporal) *TInstant {
	_cret := C.temporal_end_instant(temp._inner)
	return &TInstant{_inner: _cret}
}


// TemporalEndSequence wraps MEOS C function temporal_end_sequence.
func TemporalEndSequence(temp *Temporal) *TSequence {
	_cret := C.temporal_end_sequence(temp._inner)
	return &TSequence{_inner: _cret}
}


// TemporalEndTimestamptz wraps MEOS C function temporal_end_timestamptz.
func TemporalEndTimestamptz(temp *Temporal) int64 {
	_cret := C.temporal_end_timestamptz(temp._inner)
	return int64(_cret)
}


// TemporalHash wraps MEOS C function temporal_hash.
func TemporalHash(temp *Temporal) uint32 {
	_cret := C.temporal_hash(temp._inner)
	return uint32(_cret)
}


// TemporalInstantN wraps MEOS C function temporal_instant_n.
func TemporalInstantN(temp *Temporal, n int) *TInstant {
	_cret := C.temporal_instant_n(temp._inner, C.int(n))
	return &TInstant{_inner: _cret}
}


// TemporalInstants wraps MEOS C function temporal_instants.
func TemporalInstants(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.temporal_instants(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TemporalInterp wraps MEOS C function temporal_interp.
func TemporalInterp(temp *Temporal) string {
	_cret := C.temporal_interp(temp._inner)
	return C.GoString(_cret)
}


// TemporalLowerInc wraps MEOS C function temporal_lower_inc.
func TemporalLowerInc(temp *Temporal) bool {
	_cret := C.temporal_lower_inc(temp._inner)
	return bool(_cret)
}


// TemporalMaxInstant wraps MEOS C function temporal_max_instant.
func TemporalMaxInstant(temp *Temporal) *TInstant {
	_cret := C.temporal_max_instant(temp._inner)
	return &TInstant{_inner: _cret}
}


// TemporalMinInstant wraps MEOS C function temporal_min_instant.
func TemporalMinInstant(temp *Temporal) *TInstant {
	_cret := C.temporal_min_instant(temp._inner)
	return &TInstant{_inner: _cret}
}


// TemporalNumInstants wraps MEOS C function temporal_num_instants.
func TemporalNumInstants(temp *Temporal) int {
	_cret := C.temporal_num_instants(temp._inner)
	return int(_cret)
}


// TemporalNumSequences wraps MEOS C function temporal_num_sequences.
func TemporalNumSequences(temp *Temporal) int {
	_cret := C.temporal_num_sequences(temp._inner)
	return int(_cret)
}


// TemporalNumTimestamps wraps MEOS C function temporal_num_timestamps.
func TemporalNumTimestamps(temp *Temporal) int {
	_cret := C.temporal_num_timestamps(temp._inner)
	return int(_cret)
}


// TemporalSegmDuration wraps MEOS C function temporal_segm_duration.
func TemporalSegmDuration(temp *Temporal, duration *Interval, atleast bool, strict bool) *TSequenceSet {
	_cret := C.temporal_segm_duration(temp._inner, duration._inner, C.bool(atleast), C.bool(strict))
	return &TSequenceSet{_inner: _cret}
}


// TemporalSegments wraps MEOS C function temporal_segments.
func TemporalSegments(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.temporal_segments(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TemporalSequenceN wraps MEOS C function temporal_sequence_n.
func TemporalSequenceN(temp *Temporal, i int) *TSequence {
	_cret := C.temporal_sequence_n(temp._inner, C.int(i))
	return &TSequence{_inner: _cret}
}


// TemporalSequences wraps MEOS C function temporal_sequences.
func TemporalSequences(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.temporal_sequences(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TemporalStartInstant wraps MEOS C function temporal_start_instant.
func TemporalStartInstant(temp *Temporal) *TInstant {
	_cret := C.temporal_start_instant(temp._inner)
	return &TInstant{_inner: _cret}
}


// TemporalStartSequence wraps MEOS C function temporal_start_sequence.
func TemporalStartSequence(temp *Temporal) *TSequence {
	_cret := C.temporal_start_sequence(temp._inner)
	return &TSequence{_inner: _cret}
}


// TemporalStartTimestamptz wraps MEOS C function temporal_start_timestamptz.
func TemporalStartTimestamptz(temp *Temporal) int64 {
	_cret := C.temporal_start_timestamptz(temp._inner)
	return int64(_cret)
}


// TemporalStops wraps MEOS C function temporal_stops.
func TemporalStops(temp *Temporal, maxdist float64, minduration *Interval) *TSequenceSet {
	_cret := C.temporal_stops(temp._inner, C.double(maxdist), minduration._inner)
	return &TSequenceSet{_inner: _cret}
}


// TemporalSubtype wraps MEOS C function temporal_subtype.
func TemporalSubtype(temp *Temporal) string {
	_cret := C.temporal_subtype(temp._inner)
	return C.GoString(_cret)
}


// TemporalBasetypeName wraps MEOS C function temporal_basetype_name.
func TemporalBasetypeName(temp *Temporal) string {
	_cret := C.temporal_basetype_name(temp._inner)
	return C.GoString(_cret)
}


// TemporalTime wraps MEOS C function temporal_time.
func TemporalTime(temp *Temporal) *SpanSet {
	_cret := C.temporal_time(temp._inner)
	return &SpanSet{_inner: _cret}
}


// TemporalTimestamps wraps MEOS C function temporal_timestamps.
func TemporalTimestamps(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.temporal_timestamps(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TemporalTimestamptzN wraps MEOS C function temporal_timestamptz_n.
func TemporalTimestamptzN(temp *Temporal, n int) (bool, int64) {
	var _out_result C.TimestampTz
	_cret := C.temporal_timestamptz_n(temp._inner, C.int(n), &_out_result)
	return bool(_cret), int64(_out_result)
}


// TemporalUpperInc wraps MEOS C function temporal_upper_inc.
func TemporalUpperInc(temp *Temporal) bool {
	_cret := C.temporal_upper_inc(temp._inner)
	return bool(_cret)
}


// TfloatEndValue wraps MEOS C function tfloat_end_value.
func TfloatEndValue(temp *Temporal) float64 {
	_cret := C.tfloat_end_value(temp._inner)
	return float64(_cret)
}


// TfloatMinValue wraps MEOS C function tfloat_min_value.
func TfloatMinValue(temp *Temporal) float64 {
	_cret := C.tfloat_min_value(temp._inner)
	return float64(_cret)
}


// TfloatMaxValue wraps MEOS C function tfloat_max_value.
func TfloatMaxValue(temp *Temporal) float64 {
	_cret := C.tfloat_max_value(temp._inner)
	return float64(_cret)
}


// TfloatStartValue wraps MEOS C function tfloat_start_value.
func TfloatStartValue(temp *Temporal) float64 {
	_cret := C.tfloat_start_value(temp._inner)
	return float64(_cret)
}


// TfloatValueAtTimestamptz wraps MEOS C function tfloat_value_at_timestamptz.
func TfloatValueAtTimestamptz(temp *Temporal, t int64, strict bool) (bool, float64) {
	var _out_value C.double
	_cret := C.tfloat_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_value)
	return bool(_cret), float64(_out_value)
}


// TfloatValueN wraps MEOS C function tfloat_value_n.
func TfloatValueN(temp *Temporal, n int) (bool, float64) {
	var _out_result C.double
	_cret := C.tfloat_value_n(temp._inner, C.int(n), &_out_result)
	return bool(_cret), float64(_out_result)
}


// TfloatValues wraps MEOS C function tfloat_values.
func TfloatValues(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tfloat_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TintEndValue wraps MEOS C function tint_end_value.
func TintEndValue(temp *Temporal) int {
	_cret := C.tint_end_value(temp._inner)
	return int(_cret)
}


// TbigintEndValue wraps MEOS C function tbigint_end_value.
func TbigintEndValue(temp *Temporal) int64 {
	_cret := C.tbigint_end_value(temp._inner)
	return int64(_cret)
}


// TintMaxValue wraps MEOS C function tint_max_value.
func TintMaxValue(temp *Temporal) int {
	_cret := C.tint_max_value(temp._inner)
	return int(_cret)
}


// TbigintMaxValue wraps MEOS C function tbigint_max_value.
func TbigintMaxValue(temp *Temporal) int64 {
	_cret := C.tbigint_max_value(temp._inner)
	return int64(_cret)
}


// TintMinValue wraps MEOS C function tint_min_value.
func TintMinValue(temp *Temporal) int {
	_cret := C.tint_min_value(temp._inner)
	return int(_cret)
}


// TbigintMinValue wraps MEOS C function tbigint_min_value.
func TbigintMinValue(temp *Temporal) int64 {
	_cret := C.tbigint_min_value(temp._inner)
	return int64(_cret)
}


// TintStartValue wraps MEOS C function tint_start_value.
func TintStartValue(temp *Temporal) int {
	_cret := C.tint_start_value(temp._inner)
	return int(_cret)
}


// TbigintStartValue wraps MEOS C function tbigint_start_value.
func TbigintStartValue(temp *Temporal) int64 {
	_cret := C.tbigint_start_value(temp._inner)
	return int64(_cret)
}


// TbigintValueAtTimestamptz wraps MEOS C function tbigint_value_at_timestamptz.
func TbigintValueAtTimestamptz(temp *Temporal, t int64, strict bool) (bool, int64) {
	var _out_value C.int64_t
	_cret := C.tbigint_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_value)
	return bool(_cret), int64(_out_value)
}


// TintValueAtTimestamptz wraps MEOS C function tint_value_at_timestamptz.
func TintValueAtTimestamptz(temp *Temporal, t int64, strict bool) (bool, int) {
	var _out_value C.int
	_cret := C.tint_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_value)
	return bool(_cret), int(_out_value)
}


// TintValueN wraps MEOS C function tint_value_n.
func TintValueN(temp *Temporal, n int) (bool, int) {
	var _out_result C.int
	_cret := C.tint_value_n(temp._inner, C.int(n), &_out_result)
	return bool(_cret), int(_out_result)
}


// TbigintValueN wraps MEOS C function tbigint_value_n.
func TbigintValueN(temp *Temporal, n int64) (bool, int64) {
	var _out_result C.int64_t
	_cret := C.tbigint_value_n(temp._inner, C.int64_t(n), &_out_result)
	return bool(_cret), int64(_out_result)
}


// TintValues wraps MEOS C function tint_values.
func TintValues(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tint_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TbigintValues wraps MEOS C function tbigint_values.
func TbigintValues(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tbigint_values(temp._inner, (*C.int32)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TnumberAvgValue wraps MEOS C function tnumber_avg_value.
func TnumberAvgValue(temp *Temporal) float64 {
	_cret := C.tnumber_avg_value(temp._inner)
	return float64(_cret)
}


// TnumberIntegral wraps MEOS C function tnumber_integral.
func TnumberIntegral(temp *Temporal) float64 {
	_cret := C.tnumber_integral(temp._inner)
	return float64(_cret)
}


// TnumberTwavg wraps MEOS C function tnumber_twavg.
func TnumberTwavg(temp *Temporal) float64 {
	_cret := C.tnumber_twavg(temp._inner)
	return float64(_cret)
}


// TnumberValuespans wraps MEOS C function tnumber_valuespans.
func TnumberValuespans(temp *Temporal) *SpanSet {
	_cret := C.tnumber_valuespans(temp._inner)
	return &SpanSet{_inner: _cret}
}


// TtextEndValue wraps MEOS C function ttext_end_value.
func TtextEndValue(temp *Temporal) string {
	_cret := C.ttext_end_value(temp._inner)
	return C.GoString(C.text_to_cstring(_cret))
}


// TtextMaxValue wraps MEOS C function ttext_max_value.
func TtextMaxValue(temp *Temporal) string {
	_cret := C.ttext_max_value(temp._inner)
	return C.GoString(C.text_to_cstring(_cret))
}


// TtextMinValue wraps MEOS C function ttext_min_value.
func TtextMinValue(temp *Temporal) string {
	_cret := C.ttext_min_value(temp._inner)
	return C.GoString(C.text_to_cstring(_cret))
}


// TtextStartValue wraps MEOS C function ttext_start_value.
func TtextStartValue(temp *Temporal) string {
	_cret := C.ttext_start_value(temp._inner)
	return C.GoString(C.text_to_cstring(_cret))
}


// TtextValueAtTimestamptz wraps MEOS C function ttext_value_at_timestamptz.
func TtextValueAtTimestamptz(temp *Temporal, t int64, strict bool) (bool, string) {
	var _out_value *C.text
	_cret := C.ttext_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_value)
	return bool(_cret), C.GoString(C.text_to_cstring(_out_value))
}


// TtextValueN wraps MEOS C function ttext_value_n.
func TtextValueN(temp *Temporal, n int) (bool, string) {
	var _out_result *C.text
	_cret := C.ttext_value_n(temp._inner, C.int(n), &_out_result)
	return bool(_cret), C.GoString(C.text_to_cstring(_out_result))
}


// TtextValues wraps MEOS C function ttext_values.
func TtextValues(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.ttext_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// FloatDegrees wraps MEOS C function float_degrees.
func FloatDegrees(value float64, normalize bool) float64 {
	_cret := C.float_degrees(C.double(value), C.bool(normalize))
	return float64(_cret)
}


// TemparrRound wraps MEOS C function temparr_round.
func TemparrRound(temp unsafe.Pointer, count int, maxdd int) unsafe.Pointer {
	_cret := C.temparr_round((**C.Temporal)(unsafe.Pointer(temp)), C.int(count), C.int(maxdd))
	return unsafe.Pointer(_cret)
}


// TemporalRound wraps MEOS C function temporal_round.
func TemporalRound(temp *Temporal, maxdd int) *Temporal {
	_cret := C.temporal_round(temp._inner, C.int(maxdd))
	return &Temporal{_inner: _cret}
}


// TemporalScaleTime wraps MEOS C function temporal_scale_time.
func TemporalScaleTime(temp *Temporal, duration *Interval) *Temporal {
	_cret := C.temporal_scale_time(temp._inner, duration._inner)
	return &Temporal{_inner: _cret}
}


// TemporalSetInterp wraps MEOS C function temporal_set_interp.
func TemporalSetInterp(temp *Temporal, interp Interpolation) *Temporal {
	_cret := C.temporal_set_interp(temp._inner, C.interpType(interp))
	return &Temporal{_inner: _cret}
}


// TemporalShiftScaleTime wraps MEOS C function temporal_shift_scale_time.
func TemporalShiftScaleTime(temp *Temporal, shift *Interval, duration *Interval) *Temporal {
	_cret := C.temporal_shift_scale_time(temp._inner, shift._inner, duration._inner)
	return &Temporal{_inner: _cret}
}


// TemporalShiftTime wraps MEOS C function temporal_shift_time.
func TemporalShiftTime(temp *Temporal, shift *Interval) *Temporal {
	_cret := C.temporal_shift_time(temp._inner, shift._inner)
	return &Temporal{_inner: _cret}
}


// TemporalToTinstant wraps MEOS C function temporal_to_tinstant.
func TemporalToTinstant(temp *Temporal) *TInstant {
	_cret := C.temporal_to_tinstant(temp._inner)
	return &TInstant{_inner: _cret}
}


// TemporalToTsequence wraps MEOS C function temporal_to_tsequence.
func TemporalToTsequence(temp *Temporal, interp Interpolation) *TSequence {
	_cret := C.temporal_to_tsequence(temp._inner, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TemporalToTsequenceset wraps MEOS C function temporal_to_tsequenceset.
func TemporalToTsequenceset(temp *Temporal, interp Interpolation) *TSequenceSet {
	_cret := C.temporal_to_tsequenceset(temp._inner, C.interpType(interp))
	return &TSequenceSet{_inner: _cret}
}


// TfloatCeil wraps MEOS C function tfloat_ceil.
func TfloatCeil(temp *Temporal) *Temporal {
	_cret := C.tfloat_ceil(temp._inner)
	return &Temporal{_inner: _cret}
}


// TfloatDegrees wraps MEOS C function tfloat_degrees.
func TfloatDegrees(temp *Temporal, normalize bool) *Temporal {
	_cret := C.tfloat_degrees(temp._inner, C.bool(normalize))
	return &Temporal{_inner: _cret}
}


// TfloatFloor wraps MEOS C function tfloat_floor.
func TfloatFloor(temp *Temporal) *Temporal {
	_cret := C.tfloat_floor(temp._inner)
	return &Temporal{_inner: _cret}
}


// TfloatRadians wraps MEOS C function tfloat_radians.
func TfloatRadians(temp *Temporal) *Temporal {
	_cret := C.tfloat_radians(temp._inner)
	return &Temporal{_inner: _cret}
}


// TfloatScaleValue wraps MEOS C function tfloat_scale_value.
func TfloatScaleValue(temp *Temporal, width float64) *Temporal {
	_cret := C.tfloat_scale_value(temp._inner, C.double(width))
	return &Temporal{_inner: _cret}
}


// TfloatShiftScaleValue wraps MEOS C function tfloat_shift_scale_value.
func TfloatShiftScaleValue(temp *Temporal, shift float64, width float64) *Temporal {
	_cret := C.tfloat_shift_scale_value(temp._inner, C.double(shift), C.double(width))
	return &Temporal{_inner: _cret}
}


// TfloatShiftValue wraps MEOS C function tfloat_shift_value.
func TfloatShiftValue(temp *Temporal, shift float64) *Temporal {
	_cret := C.tfloat_shift_value(temp._inner, C.double(shift))
	return &Temporal{_inner: _cret}
}


// TintScaleValue wraps MEOS C function tint_scale_value.
func TintScaleValue(temp *Temporal, width int) *Temporal {
	_cret := C.tint_scale_value(temp._inner, C.int(width))
	return &Temporal{_inner: _cret}
}


// TbigintScaleValue wraps MEOS C function tbigint_scale_value.
func TbigintScaleValue(temp *Temporal, width int64) *Temporal {
	_cret := C.tbigint_scale_value(temp._inner, C.int64_t(width))
	return &Temporal{_inner: _cret}
}


// TintShiftScaleValue wraps MEOS C function tint_shift_scale_value.
func TintShiftScaleValue(temp *Temporal, shift int, width int) *Temporal {
	_cret := C.tint_shift_scale_value(temp._inner, C.int(shift), C.int(width))
	return &Temporal{_inner: _cret}
}


// TbigintShiftScaleValue wraps MEOS C function tbigint_shift_scale_value.
func TbigintShiftScaleValue(temp *Temporal, shift int64, width int64) *Temporal {
	_cret := C.tbigint_shift_scale_value(temp._inner, C.int64_t(shift), C.int64_t(width))
	return &Temporal{_inner: _cret}
}


// TintShiftValue wraps MEOS C function tint_shift_value.
func TintShiftValue(temp *Temporal, shift int) *Temporal {
	_cret := C.tint_shift_value(temp._inner, C.int(shift))
	return &Temporal{_inner: _cret}
}


// TbigintShiftValue wraps MEOS C function tbigint_shift_value.
func TbigintShiftValue(temp *Temporal, shift int64) *Temporal {
	_cret := C.tbigint_shift_value(temp._inner, C.int64_t(shift))
	return &Temporal{_inner: _cret}
}


// TemporalAppendTinstant wraps MEOS C function temporal_append_tinstant.
func TemporalAppendTinstant(temp *Temporal, inst *TInstant, interp Interpolation, maxdist float64, maxt *Interval, expand bool) *Temporal {
	_cret := C.temporal_append_tinstant(temp._inner, inst._inner, C.interpType(interp), C.double(maxdist), maxt._inner, C.bool(expand))
	return &Temporal{_inner: _cret}
}


// TemporalAppendTsequence wraps MEOS C function temporal_append_tsequence.
func TemporalAppendTsequence(temp *Temporal, seq *TSequence, expand bool) *Temporal {
	_cret := C.temporal_append_tsequence(temp._inner, seq._inner, C.bool(expand))
	return &Temporal{_inner: _cret}
}


// TemporalDeleteTimestamptz wraps MEOS C function temporal_delete_timestamptz.
func TemporalDeleteTimestamptz(temp *Temporal, t int64, connect bool) *Temporal {
	_cret := C.temporal_delete_timestamptz(temp._inner, C.TimestampTz(t), C.bool(connect))
	return &Temporal{_inner: _cret}
}


// TemporalDeleteTstzset wraps MEOS C function temporal_delete_tstzset.
func TemporalDeleteTstzset(temp *Temporal, s *Set, connect bool) *Temporal {
	_cret := C.temporal_delete_tstzset(temp._inner, s._inner, C.bool(connect))
	return &Temporal{_inner: _cret}
}


// TemporalDeleteTstzspan wraps MEOS C function temporal_delete_tstzspan.
func TemporalDeleteTstzspan(temp *Temporal, s *Span, connect bool) *Temporal {
	_cret := C.temporal_delete_tstzspan(temp._inner, s._inner, C.bool(connect))
	return &Temporal{_inner: _cret}
}


// TemporalDeleteTstzspanset wraps MEOS C function temporal_delete_tstzspanset.
func TemporalDeleteTstzspanset(temp *Temporal, ss *SpanSet, connect bool) *Temporal {
	_cret := C.temporal_delete_tstzspanset(temp._inner, ss._inner, C.bool(connect))
	return &Temporal{_inner: _cret}
}


// TemporalInsert wraps MEOS C function temporal_insert.
func TemporalInsert(temp1 *Temporal, temp2 *Temporal, connect bool) *Temporal {
	_cret := C.temporal_insert(temp1._inner, temp2._inner, C.bool(connect))
	return &Temporal{_inner: _cret}
}


// TemporalMerge wraps MEOS C function temporal_merge.
func TemporalMerge(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.temporal_merge(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TemporalMergeArray wraps MEOS C function temporal_merge_array.
func TemporalMergeArray(temparr unsafe.Pointer, count int) *Temporal {
	_cret := C.temporal_merge_array((**C.Temporal)(unsafe.Pointer(temparr)), C.int(count))
	return &Temporal{_inner: _cret}
}


// TemporalUpdate wraps MEOS C function temporal_update.
func TemporalUpdate(temp1 *Temporal, temp2 *Temporal, connect bool) *Temporal {
	_cret := C.temporal_update(temp1._inner, temp2._inner, C.bool(connect))
	return &Temporal{_inner: _cret}
}


// TboolAtValue wraps MEOS C function tbool_at_value.
func TboolAtValue(temp *Temporal, b bool) *Temporal {
	_cret := C.tbool_at_value(temp._inner, C.bool(b))
	return &Temporal{_inner: _cret}
}


// TboolMinusValue wraps MEOS C function tbool_minus_value.
func TboolMinusValue(temp *Temporal, b bool) *Temporal {
	_cret := C.tbool_minus_value(temp._inner, C.bool(b))
	return &Temporal{_inner: _cret}
}


// TemporalAfterTimestamptz wraps MEOS C function temporal_after_timestamptz.
func TemporalAfterTimestamptz(temp *Temporal, t int64, strict bool) *Temporal {
	_cret := C.temporal_after_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict))
	return &Temporal{_inner: _cret}
}


// TemporalAtMax wraps MEOS C function temporal_at_max.
func TemporalAtMax(temp *Temporal) *Temporal {
	_cret := C.temporal_at_max(temp._inner)
	return &Temporal{_inner: _cret}
}


// TemporalAtMin wraps MEOS C function temporal_at_min.
func TemporalAtMin(temp *Temporal) *Temporal {
	_cret := C.temporal_at_min(temp._inner)
	return &Temporal{_inner: _cret}
}


// TemporalAtTimestamptz wraps MEOS C function temporal_at_timestamptz.
func TemporalAtTimestamptz(temp *Temporal, t int64) *Temporal {
	_cret := C.temporal_at_timestamptz(temp._inner, C.TimestampTz(t))
	return &Temporal{_inner: _cret}
}


// TemporalAtTstzset wraps MEOS C function temporal_at_tstzset.
func TemporalAtTstzset(temp *Temporal, s *Set) *Temporal {
	_cret := C.temporal_at_tstzset(temp._inner, s._inner)
	return &Temporal{_inner: _cret}
}


// TemporalAtTstzspan wraps MEOS C function temporal_at_tstzspan.
func TemporalAtTstzspan(temp *Temporal, s *Span) *Temporal {
	_cret := C.temporal_at_tstzspan(temp._inner, s._inner)
	return &Temporal{_inner: _cret}
}


// TemporalAtTstzspanset wraps MEOS C function temporal_at_tstzspanset.
func TemporalAtTstzspanset(temp *Temporal, ss *SpanSet) *Temporal {
	_cret := C.temporal_at_tstzspanset(temp._inner, ss._inner)
	return &Temporal{_inner: _cret}
}


// TemporalAtValues wraps MEOS C function temporal_at_values.
func TemporalAtValues(temp *Temporal, set *Set) *Temporal {
	_cret := C.temporal_at_values(temp._inner, set._inner)
	return &Temporal{_inner: _cret}
}


// TemporalBeforeTimestamptz wraps MEOS C function temporal_before_timestamptz.
func TemporalBeforeTimestamptz(temp *Temporal, t int64, strict bool) *Temporal {
	_cret := C.temporal_before_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict))
	return &Temporal{_inner: _cret}
}


// TemporalMinusMax wraps MEOS C function temporal_minus_max.
func TemporalMinusMax(temp *Temporal) *Temporal {
	_cret := C.temporal_minus_max(temp._inner)
	return &Temporal{_inner: _cret}
}


// TemporalMinusMin wraps MEOS C function temporal_minus_min.
func TemporalMinusMin(temp *Temporal) *Temporal {
	_cret := C.temporal_minus_min(temp._inner)
	return &Temporal{_inner: _cret}
}


// TemporalMinusTimestamptz wraps MEOS C function temporal_minus_timestamptz.
func TemporalMinusTimestamptz(temp *Temporal, t int64) *Temporal {
	_cret := C.temporal_minus_timestamptz(temp._inner, C.TimestampTz(t))
	return &Temporal{_inner: _cret}
}


// TemporalMinusTstzset wraps MEOS C function temporal_minus_tstzset.
func TemporalMinusTstzset(temp *Temporal, s *Set) *Temporal {
	_cret := C.temporal_minus_tstzset(temp._inner, s._inner)
	return &Temporal{_inner: _cret}
}


// TemporalMinusTstzspan wraps MEOS C function temporal_minus_tstzspan.
func TemporalMinusTstzspan(temp *Temporal, s *Span) *Temporal {
	_cret := C.temporal_minus_tstzspan(temp._inner, s._inner)
	return &Temporal{_inner: _cret}
}


// TemporalMinusTstzspanset wraps MEOS C function temporal_minus_tstzspanset.
func TemporalMinusTstzspanset(temp *Temporal, ss *SpanSet) *Temporal {
	_cret := C.temporal_minus_tstzspanset(temp._inner, ss._inner)
	return &Temporal{_inner: _cret}
}


// TemporalMinusValues wraps MEOS C function temporal_minus_values.
func TemporalMinusValues(temp *Temporal, set *Set) *Temporal {
	_cret := C.temporal_minus_values(temp._inner, set._inner)
	return &Temporal{_inner: _cret}
}


// TfloatAtValue wraps MEOS C function tfloat_at_value.
func TfloatAtValue(temp *Temporal, d float64) *Temporal {
	_cret := C.tfloat_at_value(temp._inner, C.double(d))
	return &Temporal{_inner: _cret}
}


// TfloatMinusValue wraps MEOS C function tfloat_minus_value.
func TfloatMinusValue(temp *Temporal, d float64) *Temporal {
	_cret := C.tfloat_minus_value(temp._inner, C.double(d))
	return &Temporal{_inner: _cret}
}


// TintAtValue wraps MEOS C function tint_at_value.
func TintAtValue(temp *Temporal, i int) *Temporal {
	_cret := C.tint_at_value(temp._inner, C.int(i))
	return &Temporal{_inner: _cret}
}


// TintMinusValue wraps MEOS C function tint_minus_value.
func TintMinusValue(temp *Temporal, i int) *Temporal {
	_cret := C.tint_minus_value(temp._inner, C.int(i))
	return &Temporal{_inner: _cret}
}


// TnumberAtSpan wraps MEOS C function tnumber_at_span.
func TnumberAtSpan(temp *Temporal, span *Span) *Temporal {
	_cret := C.tnumber_at_span(temp._inner, span._inner)
	return &Temporal{_inner: _cret}
}


// TnumberAtSpanset wraps MEOS C function tnumber_at_spanset.
func TnumberAtSpanset(temp *Temporal, ss *SpanSet) *Temporal {
	_cret := C.tnumber_at_spanset(temp._inner, ss._inner)
	return &Temporal{_inner: _cret}
}


// TnumberAtTBOX wraps MEOS C function tnumber_at_tbox.
func TnumberAtTBOX(temp *Temporal, box *TBox) *Temporal {
	_cret := C.tnumber_at_tbox(temp._inner, box._inner)
	return &Temporal{_inner: _cret}
}


// TnumberMinusSpan wraps MEOS C function tnumber_minus_span.
func TnumberMinusSpan(temp *Temporal, span *Span) *Temporal {
	_cret := C.tnumber_minus_span(temp._inner, span._inner)
	return &Temporal{_inner: _cret}
}


// TnumberMinusSpanset wraps MEOS C function tnumber_minus_spanset.
func TnumberMinusSpanset(temp *Temporal, ss *SpanSet) *Temporal {
	_cret := C.tnumber_minus_spanset(temp._inner, ss._inner)
	return &Temporal{_inner: _cret}
}


// TnumberMinusTBOX wraps MEOS C function tnumber_minus_tbox.
func TnumberMinusTBOX(temp *Temporal, box *TBox) *Temporal {
	_cret := C.tnumber_minus_tbox(temp._inner, box._inner)
	return &Temporal{_inner: _cret}
}


// TtextAtValue wraps MEOS C function ttext_at_value.
func TtextAtValue(temp *Temporal, txt string) *Temporal {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ttext_at_value(temp._inner, _c_txt)
	return &Temporal{_inner: _cret}
}


// TtextMinusValue wraps MEOS C function ttext_minus_value.
func TtextMinusValue(temp *Temporal, txt string) *Temporal {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ttext_minus_value(temp._inner, _c_txt)
	return &Temporal{_inner: _cret}
}


// TemporalCmp wraps MEOS C function temporal_cmp.
func TemporalCmp(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.temporal_cmp(temp1._inner, temp2._inner)
	return int(_cret)
}


// TemporalEq wraps MEOS C function temporal_eq.
func TemporalEq(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.temporal_eq(temp1._inner, temp2._inner)
	return bool(_cret)
}


// TemporalGe wraps MEOS C function temporal_ge.
func TemporalGe(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.temporal_ge(temp1._inner, temp2._inner)
	return bool(_cret)
}


// TemporalGt wraps MEOS C function temporal_gt.
func TemporalGt(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.temporal_gt(temp1._inner, temp2._inner)
	return bool(_cret)
}


// TemporalLe wraps MEOS C function temporal_le.
func TemporalLe(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.temporal_le(temp1._inner, temp2._inner)
	return bool(_cret)
}


// TemporalLt wraps MEOS C function temporal_lt.
func TemporalLt(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.temporal_lt(temp1._inner, temp2._inner)
	return bool(_cret)
}


// TemporalNe wraps MEOS C function temporal_ne.
func TemporalNe(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.temporal_ne(temp1._inner, temp2._inner)
	return bool(_cret)
}


// AlwaysEqBoolTbool wraps MEOS C function always_eq_bool_tbool.
func AlwaysEqBoolTbool(b bool, temp *Temporal) int {
	_cret := C.always_eq_bool_tbool(C.bool(b), temp._inner)
	return int(_cret)
}


// AlwaysEqFloatTfloat wraps MEOS C function always_eq_float_tfloat.
func AlwaysEqFloatTfloat(d float64, temp *Temporal) int {
	_cret := C.always_eq_float_tfloat(C.double(d), temp._inner)
	return int(_cret)
}


// AlwaysEqIntTint wraps MEOS C function always_eq_int_tint.
func AlwaysEqIntTint(i int, temp *Temporal) int {
	_cret := C.always_eq_int_tint(C.int(i), temp._inner)
	return int(_cret)
}


// AlwaysEqTboolBool wraps MEOS C function always_eq_tbool_bool.
func AlwaysEqTboolBool(temp *Temporal, b bool) int {
	_cret := C.always_eq_tbool_bool(temp._inner, C.bool(b))
	return int(_cret)
}


// AlwaysEqTemporalTemporal wraps MEOS C function always_eq_temporal_temporal.
func AlwaysEqTemporalTemporal(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_eq_temporal_temporal(temp1._inner, temp2._inner)
	return int(_cret)
}


// AlwaysEqTextTtext wraps MEOS C function always_eq_text_ttext.
func AlwaysEqTextTtext(txt string, temp *Temporal) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.always_eq_text_ttext(_c_txt, temp._inner)
	return int(_cret)
}


// AlwaysEqTfloatFloat wraps MEOS C function always_eq_tfloat_float.
func AlwaysEqTfloatFloat(temp *Temporal, d float64) int {
	_cret := C.always_eq_tfloat_float(temp._inner, C.double(d))
	return int(_cret)
}


// AlwaysEqTintInt wraps MEOS C function always_eq_tint_int.
func AlwaysEqTintInt(temp *Temporal, i int) int {
	_cret := C.always_eq_tint_int(temp._inner, C.int(i))
	return int(_cret)
}


// AlwaysEqBigintTbigint wraps MEOS C function always_eq_bigint_tbigint.
func AlwaysEqBigintTbigint(i int64, temp *Temporal) int {
	_cret := C.always_eq_bigint_tbigint(C.int64_t(i), temp._inner)
	return int(_cret)
}


// AlwaysEqTbigintBigint wraps MEOS C function always_eq_tbigint_bigint.
func AlwaysEqTbigintBigint(temp *Temporal, i int64) int {
	_cret := C.always_eq_tbigint_bigint(temp._inner, C.int64_t(i))
	return int(_cret)
}


// AlwaysEqTtextText wraps MEOS C function always_eq_ttext_text.
func AlwaysEqTtextText(temp *Temporal, txt string) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.always_eq_ttext_text(temp._inner, _c_txt)
	return int(_cret)
}


// AlwaysGeFloatTfloat wraps MEOS C function always_ge_float_tfloat.
func AlwaysGeFloatTfloat(d float64, temp *Temporal) int {
	_cret := C.always_ge_float_tfloat(C.double(d), temp._inner)
	return int(_cret)
}


// AlwaysGeIntTint wraps MEOS C function always_ge_int_tint.
func AlwaysGeIntTint(i int, temp *Temporal) int {
	_cret := C.always_ge_int_tint(C.int(i), temp._inner)
	return int(_cret)
}


// AlwaysGeTemporalTemporal wraps MEOS C function always_ge_temporal_temporal.
func AlwaysGeTemporalTemporal(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_ge_temporal_temporal(temp1._inner, temp2._inner)
	return int(_cret)
}


// AlwaysGeTextTtext wraps MEOS C function always_ge_text_ttext.
func AlwaysGeTextTtext(txt string, temp *Temporal) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.always_ge_text_ttext(_c_txt, temp._inner)
	return int(_cret)
}


// AlwaysGeTfloatFloat wraps MEOS C function always_ge_tfloat_float.
func AlwaysGeTfloatFloat(temp *Temporal, d float64) int {
	_cret := C.always_ge_tfloat_float(temp._inner, C.double(d))
	return int(_cret)
}


// AlwaysGeTintInt wraps MEOS C function always_ge_tint_int.
func AlwaysGeTintInt(temp *Temporal, i int) int {
	_cret := C.always_ge_tint_int(temp._inner, C.int(i))
	return int(_cret)
}


// AlwaysGeBigintTbigint wraps MEOS C function always_ge_bigint_tbigint.
func AlwaysGeBigintTbigint(i int64, temp *Temporal) int {
	_cret := C.always_ge_bigint_tbigint(C.int64_t(i), temp._inner)
	return int(_cret)
}


// AlwaysGeTbigintBigint wraps MEOS C function always_ge_tbigint_bigint.
func AlwaysGeTbigintBigint(temp *Temporal, i int64) int {
	_cret := C.always_ge_tbigint_bigint(temp._inner, C.int64_t(i))
	return int(_cret)
}


// AlwaysGeTtextText wraps MEOS C function always_ge_ttext_text.
func AlwaysGeTtextText(temp *Temporal, txt string) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.always_ge_ttext_text(temp._inner, _c_txt)
	return int(_cret)
}


// AlwaysGtFloatTfloat wraps MEOS C function always_gt_float_tfloat.
func AlwaysGtFloatTfloat(d float64, temp *Temporal) int {
	_cret := C.always_gt_float_tfloat(C.double(d), temp._inner)
	return int(_cret)
}


// AlwaysGtIntTint wraps MEOS C function always_gt_int_tint.
func AlwaysGtIntTint(i int, temp *Temporal) int {
	_cret := C.always_gt_int_tint(C.int(i), temp._inner)
	return int(_cret)
}


// AlwaysGtTemporalTemporal wraps MEOS C function always_gt_temporal_temporal.
func AlwaysGtTemporalTemporal(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_gt_temporal_temporal(temp1._inner, temp2._inner)
	return int(_cret)
}


// AlwaysGtTextTtext wraps MEOS C function always_gt_text_ttext.
func AlwaysGtTextTtext(txt string, temp *Temporal) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.always_gt_text_ttext(_c_txt, temp._inner)
	return int(_cret)
}


// AlwaysGtTfloatFloat wraps MEOS C function always_gt_tfloat_float.
func AlwaysGtTfloatFloat(temp *Temporal, d float64) int {
	_cret := C.always_gt_tfloat_float(temp._inner, C.double(d))
	return int(_cret)
}


// AlwaysGtTintInt wraps MEOS C function always_gt_tint_int.
func AlwaysGtTintInt(temp *Temporal, i int) int {
	_cret := C.always_gt_tint_int(temp._inner, C.int(i))
	return int(_cret)
}


// AlwaysGtBigintTbigint wraps MEOS C function always_gt_bigint_tbigint.
func AlwaysGtBigintTbigint(i int64, temp *Temporal) int {
	_cret := C.always_gt_bigint_tbigint(C.int64_t(i), temp._inner)
	return int(_cret)
}


// AlwaysGtTbigintBigint wraps MEOS C function always_gt_tbigint_bigint.
func AlwaysGtTbigintBigint(temp *Temporal, i int64) int {
	_cret := C.always_gt_tbigint_bigint(temp._inner, C.int64_t(i))
	return int(_cret)
}


// AlwaysGtTtextText wraps MEOS C function always_gt_ttext_text.
func AlwaysGtTtextText(temp *Temporal, txt string) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.always_gt_ttext_text(temp._inner, _c_txt)
	return int(_cret)
}


// AlwaysLeFloatTfloat wraps MEOS C function always_le_float_tfloat.
func AlwaysLeFloatTfloat(d float64, temp *Temporal) int {
	_cret := C.always_le_float_tfloat(C.double(d), temp._inner)
	return int(_cret)
}


// AlwaysLeIntTint wraps MEOS C function always_le_int_tint.
func AlwaysLeIntTint(i int, temp *Temporal) int {
	_cret := C.always_le_int_tint(C.int(i), temp._inner)
	return int(_cret)
}


// AlwaysLeTemporalTemporal wraps MEOS C function always_le_temporal_temporal.
func AlwaysLeTemporalTemporal(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_le_temporal_temporal(temp1._inner, temp2._inner)
	return int(_cret)
}


// AlwaysLeTextTtext wraps MEOS C function always_le_text_ttext.
func AlwaysLeTextTtext(txt string, temp *Temporal) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.always_le_text_ttext(_c_txt, temp._inner)
	return int(_cret)
}


// AlwaysLeTfloatFloat wraps MEOS C function always_le_tfloat_float.
func AlwaysLeTfloatFloat(temp *Temporal, d float64) int {
	_cret := C.always_le_tfloat_float(temp._inner, C.double(d))
	return int(_cret)
}


// AlwaysLeTintInt wraps MEOS C function always_le_tint_int.
func AlwaysLeTintInt(temp *Temporal, i int) int {
	_cret := C.always_le_tint_int(temp._inner, C.int(i))
	return int(_cret)
}


// AlwaysLeBigintTbigint wraps MEOS C function always_le_bigint_tbigint.
func AlwaysLeBigintTbigint(i int64, temp *Temporal) int {
	_cret := C.always_le_bigint_tbigint(C.int64_t(i), temp._inner)
	return int(_cret)
}


// AlwaysLeTbigintBigint wraps MEOS C function always_le_tbigint_bigint.
func AlwaysLeTbigintBigint(temp *Temporal, i int64) int {
	_cret := C.always_le_tbigint_bigint(temp._inner, C.int64_t(i))
	return int(_cret)
}


// AlwaysLeTtextText wraps MEOS C function always_le_ttext_text.
func AlwaysLeTtextText(temp *Temporal, txt string) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.always_le_ttext_text(temp._inner, _c_txt)
	return int(_cret)
}


// AlwaysLtFloatTfloat wraps MEOS C function always_lt_float_tfloat.
func AlwaysLtFloatTfloat(d float64, temp *Temporal) int {
	_cret := C.always_lt_float_tfloat(C.double(d), temp._inner)
	return int(_cret)
}


// AlwaysLtIntTint wraps MEOS C function always_lt_int_tint.
func AlwaysLtIntTint(i int, temp *Temporal) int {
	_cret := C.always_lt_int_tint(C.int(i), temp._inner)
	return int(_cret)
}


// AlwaysLtTemporalTemporal wraps MEOS C function always_lt_temporal_temporal.
func AlwaysLtTemporalTemporal(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_lt_temporal_temporal(temp1._inner, temp2._inner)
	return int(_cret)
}


// AlwaysLtTextTtext wraps MEOS C function always_lt_text_ttext.
func AlwaysLtTextTtext(txt string, temp *Temporal) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.always_lt_text_ttext(_c_txt, temp._inner)
	return int(_cret)
}


// AlwaysLtTfloatFloat wraps MEOS C function always_lt_tfloat_float.
func AlwaysLtTfloatFloat(temp *Temporal, d float64) int {
	_cret := C.always_lt_tfloat_float(temp._inner, C.double(d))
	return int(_cret)
}


// AlwaysLtTintInt wraps MEOS C function always_lt_tint_int.
func AlwaysLtTintInt(temp *Temporal, i int) int {
	_cret := C.always_lt_tint_int(temp._inner, C.int(i))
	return int(_cret)
}


// AlwaysLtBigintTbigint wraps MEOS C function always_lt_bigint_tbigint.
func AlwaysLtBigintTbigint(i int64, temp *Temporal) int {
	_cret := C.always_lt_bigint_tbigint(C.int64_t(i), temp._inner)
	return int(_cret)
}


// AlwaysLtTbigintBigint wraps MEOS C function always_lt_tbigint_bigint.
func AlwaysLtTbigintBigint(temp *Temporal, i int64) int {
	_cret := C.always_lt_tbigint_bigint(temp._inner, C.int64_t(i))
	return int(_cret)
}


// AlwaysLtTtextText wraps MEOS C function always_lt_ttext_text.
func AlwaysLtTtextText(temp *Temporal, txt string) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.always_lt_ttext_text(temp._inner, _c_txt)
	return int(_cret)
}


// AlwaysNeBoolTbool wraps MEOS C function always_ne_bool_tbool.
func AlwaysNeBoolTbool(b bool, temp *Temporal) int {
	_cret := C.always_ne_bool_tbool(C.bool(b), temp._inner)
	return int(_cret)
}


// AlwaysNeFloatTfloat wraps MEOS C function always_ne_float_tfloat.
func AlwaysNeFloatTfloat(d float64, temp *Temporal) int {
	_cret := C.always_ne_float_tfloat(C.double(d), temp._inner)
	return int(_cret)
}


// AlwaysNeIntTint wraps MEOS C function always_ne_int_tint.
func AlwaysNeIntTint(i int, temp *Temporal) int {
	_cret := C.always_ne_int_tint(C.int(i), temp._inner)
	return int(_cret)
}


// AlwaysNeTboolBool wraps MEOS C function always_ne_tbool_bool.
func AlwaysNeTboolBool(temp *Temporal, b bool) int {
	_cret := C.always_ne_tbool_bool(temp._inner, C.bool(b))
	return int(_cret)
}


// AlwaysNeTemporalTemporal wraps MEOS C function always_ne_temporal_temporal.
func AlwaysNeTemporalTemporal(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_ne_temporal_temporal(temp1._inner, temp2._inner)
	return int(_cret)
}


// AlwaysNeTextTtext wraps MEOS C function always_ne_text_ttext.
func AlwaysNeTextTtext(txt string, temp *Temporal) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.always_ne_text_ttext(_c_txt, temp._inner)
	return int(_cret)
}


// AlwaysNeTfloatFloat wraps MEOS C function always_ne_tfloat_float.
func AlwaysNeTfloatFloat(temp *Temporal, d float64) int {
	_cret := C.always_ne_tfloat_float(temp._inner, C.double(d))
	return int(_cret)
}


// AlwaysNeTintInt wraps MEOS C function always_ne_tint_int.
func AlwaysNeTintInt(temp *Temporal, i int) int {
	_cret := C.always_ne_tint_int(temp._inner, C.int(i))
	return int(_cret)
}


// AlwaysNeBigintTbigint wraps MEOS C function always_ne_bigint_tbigint.
func AlwaysNeBigintTbigint(i int64, temp *Temporal) int {
	_cret := C.always_ne_bigint_tbigint(C.int64_t(i), temp._inner)
	return int(_cret)
}


// AlwaysNeTbigintBigint wraps MEOS C function always_ne_tbigint_bigint.
func AlwaysNeTbigintBigint(temp *Temporal, i int64) int {
	_cret := C.always_ne_tbigint_bigint(temp._inner, C.int64_t(i))
	return int(_cret)
}


// AlwaysNeTtextText wraps MEOS C function always_ne_ttext_text.
func AlwaysNeTtextText(temp *Temporal, txt string) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.always_ne_ttext_text(temp._inner, _c_txt)
	return int(_cret)
}


// EverEqBoolTbool wraps MEOS C function ever_eq_bool_tbool.
func EverEqBoolTbool(b bool, temp *Temporal) int {
	_cret := C.ever_eq_bool_tbool(C.bool(b), temp._inner)
	return int(_cret)
}


// EverEqFloatTfloat wraps MEOS C function ever_eq_float_tfloat.
func EverEqFloatTfloat(d float64, temp *Temporal) int {
	_cret := C.ever_eq_float_tfloat(C.double(d), temp._inner)
	return int(_cret)
}


// EverEqIntTint wraps MEOS C function ever_eq_int_tint.
func EverEqIntTint(i int, temp *Temporal) int {
	_cret := C.ever_eq_int_tint(C.int(i), temp._inner)
	return int(_cret)
}


// EverEqTboolBool wraps MEOS C function ever_eq_tbool_bool.
func EverEqTboolBool(temp *Temporal, b bool) int {
	_cret := C.ever_eq_tbool_bool(temp._inner, C.bool(b))
	return int(_cret)
}


// EverEqTemporalTemporal wraps MEOS C function ever_eq_temporal_temporal.
func EverEqTemporalTemporal(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_eq_temporal_temporal(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverEqTextTtext wraps MEOS C function ever_eq_text_ttext.
func EverEqTextTtext(txt string, temp *Temporal) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ever_eq_text_ttext(_c_txt, temp._inner)
	return int(_cret)
}


// EverEqTfloatFloat wraps MEOS C function ever_eq_tfloat_float.
func EverEqTfloatFloat(temp *Temporal, d float64) int {
	_cret := C.ever_eq_tfloat_float(temp._inner, C.double(d))
	return int(_cret)
}


// EverEqTintInt wraps MEOS C function ever_eq_tint_int.
func EverEqTintInt(temp *Temporal, i int) int {
	_cret := C.ever_eq_tint_int(temp._inner, C.int(i))
	return int(_cret)
}


// EverEqBigintTbigint wraps MEOS C function ever_eq_bigint_tbigint.
func EverEqBigintTbigint(i int64, temp *Temporal) int {
	_cret := C.ever_eq_bigint_tbigint(C.int64_t(i), temp._inner)
	return int(_cret)
}


// EverEqTbigintBigint wraps MEOS C function ever_eq_tbigint_bigint.
func EverEqTbigintBigint(temp *Temporal, i int64) int {
	_cret := C.ever_eq_tbigint_bigint(temp._inner, C.int64_t(i))
	return int(_cret)
}


// EverEqTtextText wraps MEOS C function ever_eq_ttext_text.
func EverEqTtextText(temp *Temporal, txt string) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ever_eq_ttext_text(temp._inner, _c_txt)
	return int(_cret)
}


// EverGeFloatTfloat wraps MEOS C function ever_ge_float_tfloat.
func EverGeFloatTfloat(d float64, temp *Temporal) int {
	_cret := C.ever_ge_float_tfloat(C.double(d), temp._inner)
	return int(_cret)
}


// EverGeIntTint wraps MEOS C function ever_ge_int_tint.
func EverGeIntTint(i int, temp *Temporal) int {
	_cret := C.ever_ge_int_tint(C.int(i), temp._inner)
	return int(_cret)
}


// EverGeTemporalTemporal wraps MEOS C function ever_ge_temporal_temporal.
func EverGeTemporalTemporal(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_ge_temporal_temporal(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverGeTextTtext wraps MEOS C function ever_ge_text_ttext.
func EverGeTextTtext(txt string, temp *Temporal) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ever_ge_text_ttext(_c_txt, temp._inner)
	return int(_cret)
}


// EverGeTfloatFloat wraps MEOS C function ever_ge_tfloat_float.
func EverGeTfloatFloat(temp *Temporal, d float64) int {
	_cret := C.ever_ge_tfloat_float(temp._inner, C.double(d))
	return int(_cret)
}


// EverGeTintInt wraps MEOS C function ever_ge_tint_int.
func EverGeTintInt(temp *Temporal, i int) int {
	_cret := C.ever_ge_tint_int(temp._inner, C.int(i))
	return int(_cret)
}


// EverGeBigintTbigint wraps MEOS C function ever_ge_bigint_tbigint.
func EverGeBigintTbigint(i int64, temp *Temporal) int {
	_cret := C.ever_ge_bigint_tbigint(C.int64_t(i), temp._inner)
	return int(_cret)
}


// EverGeTbigintBigint wraps MEOS C function ever_ge_tbigint_bigint.
func EverGeTbigintBigint(temp *Temporal, i int64) int {
	_cret := C.ever_ge_tbigint_bigint(temp._inner, C.int64_t(i))
	return int(_cret)
}


// EverGeTtextText wraps MEOS C function ever_ge_ttext_text.
func EverGeTtextText(temp *Temporal, txt string) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ever_ge_ttext_text(temp._inner, _c_txt)
	return int(_cret)
}


// EverGtFloatTfloat wraps MEOS C function ever_gt_float_tfloat.
func EverGtFloatTfloat(d float64, temp *Temporal) int {
	_cret := C.ever_gt_float_tfloat(C.double(d), temp._inner)
	return int(_cret)
}


// EverGtIntTint wraps MEOS C function ever_gt_int_tint.
func EverGtIntTint(i int, temp *Temporal) int {
	_cret := C.ever_gt_int_tint(C.int(i), temp._inner)
	return int(_cret)
}


// EverGtTemporalTemporal wraps MEOS C function ever_gt_temporal_temporal.
func EverGtTemporalTemporal(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_gt_temporal_temporal(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverGtTextTtext wraps MEOS C function ever_gt_text_ttext.
func EverGtTextTtext(txt string, temp *Temporal) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ever_gt_text_ttext(_c_txt, temp._inner)
	return int(_cret)
}


// EverGtTfloatFloat wraps MEOS C function ever_gt_tfloat_float.
func EverGtTfloatFloat(temp *Temporal, d float64) int {
	_cret := C.ever_gt_tfloat_float(temp._inner, C.double(d))
	return int(_cret)
}


// EverGtTintInt wraps MEOS C function ever_gt_tint_int.
func EverGtTintInt(temp *Temporal, i int) int {
	_cret := C.ever_gt_tint_int(temp._inner, C.int(i))
	return int(_cret)
}


// EverGtBigintTbigint wraps MEOS C function ever_gt_bigint_tbigint.
func EverGtBigintTbigint(i int64, temp *Temporal) int {
	_cret := C.ever_gt_bigint_tbigint(C.int64_t(i), temp._inner)
	return int(_cret)
}


// EverGtTbigintBigint wraps MEOS C function ever_gt_tbigint_bigint.
func EverGtTbigintBigint(temp *Temporal, i int64) int {
	_cret := C.ever_gt_tbigint_bigint(temp._inner, C.int64_t(i))
	return int(_cret)
}


// EverGtTtextText wraps MEOS C function ever_gt_ttext_text.
func EverGtTtextText(temp *Temporal, txt string) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ever_gt_ttext_text(temp._inner, _c_txt)
	return int(_cret)
}


// EverLeFloatTfloat wraps MEOS C function ever_le_float_tfloat.
func EverLeFloatTfloat(d float64, temp *Temporal) int {
	_cret := C.ever_le_float_tfloat(C.double(d), temp._inner)
	return int(_cret)
}


// EverLeIntTint wraps MEOS C function ever_le_int_tint.
func EverLeIntTint(i int, temp *Temporal) int {
	_cret := C.ever_le_int_tint(C.int(i), temp._inner)
	return int(_cret)
}


// EverLeTemporalTemporal wraps MEOS C function ever_le_temporal_temporal.
func EverLeTemporalTemporal(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_le_temporal_temporal(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverLeTextTtext wraps MEOS C function ever_le_text_ttext.
func EverLeTextTtext(txt string, temp *Temporal) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ever_le_text_ttext(_c_txt, temp._inner)
	return int(_cret)
}


// EverLeTfloatFloat wraps MEOS C function ever_le_tfloat_float.
func EverLeTfloatFloat(temp *Temporal, d float64) int {
	_cret := C.ever_le_tfloat_float(temp._inner, C.double(d))
	return int(_cret)
}


// EverLeTintInt wraps MEOS C function ever_le_tint_int.
func EverLeTintInt(temp *Temporal, i int) int {
	_cret := C.ever_le_tint_int(temp._inner, C.int(i))
	return int(_cret)
}


// EverLeBigintTbigint wraps MEOS C function ever_le_bigint_tbigint.
func EverLeBigintTbigint(i int64, temp *Temporal) int {
	_cret := C.ever_le_bigint_tbigint(C.int64_t(i), temp._inner)
	return int(_cret)
}


// EverLeTbigintBigint wraps MEOS C function ever_le_tbigint_bigint.
func EverLeTbigintBigint(temp *Temporal, i int64) int {
	_cret := C.ever_le_tbigint_bigint(temp._inner, C.int64_t(i))
	return int(_cret)
}


// EverLeTtextText wraps MEOS C function ever_le_ttext_text.
func EverLeTtextText(temp *Temporal, txt string) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ever_le_ttext_text(temp._inner, _c_txt)
	return int(_cret)
}


// EverLtFloatTfloat wraps MEOS C function ever_lt_float_tfloat.
func EverLtFloatTfloat(d float64, temp *Temporal) int {
	_cret := C.ever_lt_float_tfloat(C.double(d), temp._inner)
	return int(_cret)
}


// EverLtIntTint wraps MEOS C function ever_lt_int_tint.
func EverLtIntTint(i int, temp *Temporal) int {
	_cret := C.ever_lt_int_tint(C.int(i), temp._inner)
	return int(_cret)
}


// EverLtTemporalTemporal wraps MEOS C function ever_lt_temporal_temporal.
func EverLtTemporalTemporal(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_lt_temporal_temporal(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverLtTextTtext wraps MEOS C function ever_lt_text_ttext.
func EverLtTextTtext(txt string, temp *Temporal) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ever_lt_text_ttext(_c_txt, temp._inner)
	return int(_cret)
}


// EverLtTfloatFloat wraps MEOS C function ever_lt_tfloat_float.
func EverLtTfloatFloat(temp *Temporal, d float64) int {
	_cret := C.ever_lt_tfloat_float(temp._inner, C.double(d))
	return int(_cret)
}


// EverLtTintInt wraps MEOS C function ever_lt_tint_int.
func EverLtTintInt(temp *Temporal, i int) int {
	_cret := C.ever_lt_tint_int(temp._inner, C.int(i))
	return int(_cret)
}


// EverLtBigintTbigint wraps MEOS C function ever_lt_bigint_tbigint.
func EverLtBigintTbigint(i int64, temp *Temporal) int {
	_cret := C.ever_lt_bigint_tbigint(C.int64_t(i), temp._inner)
	return int(_cret)
}


// EverLtTbigintBigint wraps MEOS C function ever_lt_tbigint_bigint.
func EverLtTbigintBigint(temp *Temporal, i int64) int {
	_cret := C.ever_lt_tbigint_bigint(temp._inner, C.int64_t(i))
	return int(_cret)
}


// EverLtTtextText wraps MEOS C function ever_lt_ttext_text.
func EverLtTtextText(temp *Temporal, txt string) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ever_lt_ttext_text(temp._inner, _c_txt)
	return int(_cret)
}


// EverNeBoolTbool wraps MEOS C function ever_ne_bool_tbool.
func EverNeBoolTbool(b bool, temp *Temporal) int {
	_cret := C.ever_ne_bool_tbool(C.bool(b), temp._inner)
	return int(_cret)
}


// EverNeFloatTfloat wraps MEOS C function ever_ne_float_tfloat.
func EverNeFloatTfloat(d float64, temp *Temporal) int {
	_cret := C.ever_ne_float_tfloat(C.double(d), temp._inner)
	return int(_cret)
}


// EverNeIntTint wraps MEOS C function ever_ne_int_tint.
func EverNeIntTint(i int, temp *Temporal) int {
	_cret := C.ever_ne_int_tint(C.int(i), temp._inner)
	return int(_cret)
}


// EverNeTboolBool wraps MEOS C function ever_ne_tbool_bool.
func EverNeTboolBool(temp *Temporal, b bool) int {
	_cret := C.ever_ne_tbool_bool(temp._inner, C.bool(b))
	return int(_cret)
}


// EverNeTemporalTemporal wraps MEOS C function ever_ne_temporal_temporal.
func EverNeTemporalTemporal(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_ne_temporal_temporal(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverNeTextTtext wraps MEOS C function ever_ne_text_ttext.
func EverNeTextTtext(txt string, temp *Temporal) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ever_ne_text_ttext(_c_txt, temp._inner)
	return int(_cret)
}


// EverNeTfloatFloat wraps MEOS C function ever_ne_tfloat_float.
func EverNeTfloatFloat(temp *Temporal, d float64) int {
	_cret := C.ever_ne_tfloat_float(temp._inner, C.double(d))
	return int(_cret)
}


// EverNeTintInt wraps MEOS C function ever_ne_tint_int.
func EverNeTintInt(temp *Temporal, i int) int {
	_cret := C.ever_ne_tint_int(temp._inner, C.int(i))
	return int(_cret)
}


// EverNeBigintTbigint wraps MEOS C function ever_ne_bigint_tbigint.
func EverNeBigintTbigint(i int64, temp *Temporal) int {
	_cret := C.ever_ne_bigint_tbigint(C.int64_t(i), temp._inner)
	return int(_cret)
}


// EverNeTbigintBigint wraps MEOS C function ever_ne_tbigint_bigint.
func EverNeTbigintBigint(temp *Temporal, i int64) int {
	_cret := C.ever_ne_tbigint_bigint(temp._inner, C.int64_t(i))
	return int(_cret)
}


// EverNeTtextText wraps MEOS C function ever_ne_ttext_text.
func EverNeTtextText(temp *Temporal, txt string) int {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.ever_ne_ttext_text(temp._inner, _c_txt)
	return int(_cret)
}


// TeqBoolTbool wraps MEOS C function teq_bool_tbool.
func TeqBoolTbool(b bool, temp *Temporal) *Temporal {
	_cret := C.teq_bool_tbool(C.bool(b), temp._inner)
	return &Temporal{_inner: _cret}
}


// TeqFloatTfloat wraps MEOS C function teq_float_tfloat.
func TeqFloatTfloat(d float64, temp *Temporal) *Temporal {
	_cret := C.teq_float_tfloat(C.double(d), temp._inner)
	return &Temporal{_inner: _cret}
}


// TeqIntTint wraps MEOS C function teq_int_tint.
func TeqIntTint(i int, temp *Temporal) *Temporal {
	_cret := C.teq_int_tint(C.int(i), temp._inner)
	return &Temporal{_inner: _cret}
}


// TeqTboolBool wraps MEOS C function teq_tbool_bool.
func TeqTboolBool(temp *Temporal, b bool) *Temporal {
	_cret := C.teq_tbool_bool(temp._inner, C.bool(b))
	return &Temporal{_inner: _cret}
}


// TeqTemporalTemporal wraps MEOS C function teq_temporal_temporal.
func TeqTemporalTemporal(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.teq_temporal_temporal(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TeqTextTtext wraps MEOS C function teq_text_ttext.
func TeqTextTtext(txt string, temp *Temporal) *Temporal {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.teq_text_ttext(_c_txt, temp._inner)
	return &Temporal{_inner: _cret}
}


// TeqTfloatFloat wraps MEOS C function teq_tfloat_float.
func TeqTfloatFloat(temp *Temporal, d float64) *Temporal {
	_cret := C.teq_tfloat_float(temp._inner, C.double(d))
	return &Temporal{_inner: _cret}
}


// TeqTintInt wraps MEOS C function teq_tint_int.
func TeqTintInt(temp *Temporal, i int) *Temporal {
	_cret := C.teq_tint_int(temp._inner, C.int(i))
	return &Temporal{_inner: _cret}
}


// TeqTtextText wraps MEOS C function teq_ttext_text.
func TeqTtextText(temp *Temporal, txt string) *Temporal {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.teq_ttext_text(temp._inner, _c_txt)
	return &Temporal{_inner: _cret}
}


// TgeFloatTfloat wraps MEOS C function tge_float_tfloat.
func TgeFloatTfloat(d float64, temp *Temporal) *Temporal {
	_cret := C.tge_float_tfloat(C.double(d), temp._inner)
	return &Temporal{_inner: _cret}
}


// TgeIntTint wraps MEOS C function tge_int_tint.
func TgeIntTint(i int, temp *Temporal) *Temporal {
	_cret := C.tge_int_tint(C.int(i), temp._inner)
	return &Temporal{_inner: _cret}
}


// TgeTemporalTemporal wraps MEOS C function tge_temporal_temporal.
func TgeTemporalTemporal(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tge_temporal_temporal(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TgeTextTtext wraps MEOS C function tge_text_ttext.
func TgeTextTtext(txt string, temp *Temporal) *Temporal {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.tge_text_ttext(_c_txt, temp._inner)
	return &Temporal{_inner: _cret}
}


// TgeTfloatFloat wraps MEOS C function tge_tfloat_float.
func TgeTfloatFloat(temp *Temporal, d float64) *Temporal {
	_cret := C.tge_tfloat_float(temp._inner, C.double(d))
	return &Temporal{_inner: _cret}
}


// TgeTintInt wraps MEOS C function tge_tint_int.
func TgeTintInt(temp *Temporal, i int) *Temporal {
	_cret := C.tge_tint_int(temp._inner, C.int(i))
	return &Temporal{_inner: _cret}
}


// TgeTtextText wraps MEOS C function tge_ttext_text.
func TgeTtextText(temp *Temporal, txt string) *Temporal {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.tge_ttext_text(temp._inner, _c_txt)
	return &Temporal{_inner: _cret}
}


// TgtFloatTfloat wraps MEOS C function tgt_float_tfloat.
func TgtFloatTfloat(d float64, temp *Temporal) *Temporal {
	_cret := C.tgt_float_tfloat(C.double(d), temp._inner)
	return &Temporal{_inner: _cret}
}


// TgtIntTint wraps MEOS C function tgt_int_tint.
func TgtIntTint(i int, temp *Temporal) *Temporal {
	_cret := C.tgt_int_tint(C.int(i), temp._inner)
	return &Temporal{_inner: _cret}
}


// TgtTemporalTemporal wraps MEOS C function tgt_temporal_temporal.
func TgtTemporalTemporal(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tgt_temporal_temporal(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TgtTextTtext wraps MEOS C function tgt_text_ttext.
func TgtTextTtext(txt string, temp *Temporal) *Temporal {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.tgt_text_ttext(_c_txt, temp._inner)
	return &Temporal{_inner: _cret}
}


// TgtTfloatFloat wraps MEOS C function tgt_tfloat_float.
func TgtTfloatFloat(temp *Temporal, d float64) *Temporal {
	_cret := C.tgt_tfloat_float(temp._inner, C.double(d))
	return &Temporal{_inner: _cret}
}


// TgtTintInt wraps MEOS C function tgt_tint_int.
func TgtTintInt(temp *Temporal, i int) *Temporal {
	_cret := C.tgt_tint_int(temp._inner, C.int(i))
	return &Temporal{_inner: _cret}
}


// TgtTtextText wraps MEOS C function tgt_ttext_text.
func TgtTtextText(temp *Temporal, txt string) *Temporal {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.tgt_ttext_text(temp._inner, _c_txt)
	return &Temporal{_inner: _cret}
}


// TleFloatTfloat wraps MEOS C function tle_float_tfloat.
func TleFloatTfloat(d float64, temp *Temporal) *Temporal {
	_cret := C.tle_float_tfloat(C.double(d), temp._inner)
	return &Temporal{_inner: _cret}
}


// TleIntTint wraps MEOS C function tle_int_tint.
func TleIntTint(i int, temp *Temporal) *Temporal {
	_cret := C.tle_int_tint(C.int(i), temp._inner)
	return &Temporal{_inner: _cret}
}


// TleTemporalTemporal wraps MEOS C function tle_temporal_temporal.
func TleTemporalTemporal(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tle_temporal_temporal(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TleTextTtext wraps MEOS C function tle_text_ttext.
func TleTextTtext(txt string, temp *Temporal) *Temporal {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.tle_text_ttext(_c_txt, temp._inner)
	return &Temporal{_inner: _cret}
}


// TleTfloatFloat wraps MEOS C function tle_tfloat_float.
func TleTfloatFloat(temp *Temporal, d float64) *Temporal {
	_cret := C.tle_tfloat_float(temp._inner, C.double(d))
	return &Temporal{_inner: _cret}
}


// TleTintInt wraps MEOS C function tle_tint_int.
func TleTintInt(temp *Temporal, i int) *Temporal {
	_cret := C.tle_tint_int(temp._inner, C.int(i))
	return &Temporal{_inner: _cret}
}


// TleTtextText wraps MEOS C function tle_ttext_text.
func TleTtextText(temp *Temporal, txt string) *Temporal {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.tle_ttext_text(temp._inner, _c_txt)
	return &Temporal{_inner: _cret}
}


// TltFloatTfloat wraps MEOS C function tlt_float_tfloat.
func TltFloatTfloat(d float64, temp *Temporal) *Temporal {
	_cret := C.tlt_float_tfloat(C.double(d), temp._inner)
	return &Temporal{_inner: _cret}
}


// TltIntTint wraps MEOS C function tlt_int_tint.
func TltIntTint(i int, temp *Temporal) *Temporal {
	_cret := C.tlt_int_tint(C.int(i), temp._inner)
	return &Temporal{_inner: _cret}
}


// TltTemporalTemporal wraps MEOS C function tlt_temporal_temporal.
func TltTemporalTemporal(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tlt_temporal_temporal(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TltTextTtext wraps MEOS C function tlt_text_ttext.
func TltTextTtext(txt string, temp *Temporal) *Temporal {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.tlt_text_ttext(_c_txt, temp._inner)
	return &Temporal{_inner: _cret}
}


// TltTfloatFloat wraps MEOS C function tlt_tfloat_float.
func TltTfloatFloat(temp *Temporal, d float64) *Temporal {
	_cret := C.tlt_tfloat_float(temp._inner, C.double(d))
	return &Temporal{_inner: _cret}
}


// TltTintInt wraps MEOS C function tlt_tint_int.
func TltTintInt(temp *Temporal, i int) *Temporal {
	_cret := C.tlt_tint_int(temp._inner, C.int(i))
	return &Temporal{_inner: _cret}
}


// TltTtextText wraps MEOS C function tlt_ttext_text.
func TltTtextText(temp *Temporal, txt string) *Temporal {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.tlt_ttext_text(temp._inner, _c_txt)
	return &Temporal{_inner: _cret}
}


// TneBoolTbool wraps MEOS C function tne_bool_tbool.
func TneBoolTbool(b bool, temp *Temporal) *Temporal {
	_cret := C.tne_bool_tbool(C.bool(b), temp._inner)
	return &Temporal{_inner: _cret}
}


// TneFloatTfloat wraps MEOS C function tne_float_tfloat.
func TneFloatTfloat(d float64, temp *Temporal) *Temporal {
	_cret := C.tne_float_tfloat(C.double(d), temp._inner)
	return &Temporal{_inner: _cret}
}


// TneIntTint wraps MEOS C function tne_int_tint.
func TneIntTint(i int, temp *Temporal) *Temporal {
	_cret := C.tne_int_tint(C.int(i), temp._inner)
	return &Temporal{_inner: _cret}
}


// TneTboolBool wraps MEOS C function tne_tbool_bool.
func TneTboolBool(temp *Temporal, b bool) *Temporal {
	_cret := C.tne_tbool_bool(temp._inner, C.bool(b))
	return &Temporal{_inner: _cret}
}


// TneTemporalTemporal wraps MEOS C function tne_temporal_temporal.
func TneTemporalTemporal(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tne_temporal_temporal(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TneTextTtext wraps MEOS C function tne_text_ttext.
func TneTextTtext(txt string, temp *Temporal) *Temporal {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.tne_text_ttext(_c_txt, temp._inner)
	return &Temporal{_inner: _cret}
}


// TneTfloatFloat wraps MEOS C function tne_tfloat_float.
func TneTfloatFloat(temp *Temporal, d float64) *Temporal {
	_cret := C.tne_tfloat_float(temp._inner, C.double(d))
	return &Temporal{_inner: _cret}
}


// TneTintInt wraps MEOS C function tne_tint_int.
func TneTintInt(temp *Temporal, i int) *Temporal {
	_cret := C.tne_tint_int(temp._inner, C.int(i))
	return &Temporal{_inner: _cret}
}


// TneTtextText wraps MEOS C function tne_ttext_text.
func TneTtextText(temp *Temporal, txt string) *Temporal {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.tne_ttext_text(temp._inner, _c_txt)
	return &Temporal{_inner: _cret}
}


// TemporalSpans wraps MEOS C function temporal_spans.
func TemporalSpans(temp *Temporal, count unsafe.Pointer) *Span {
	_cret := C.temporal_spans(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// TemporalSplitEachNSpans wraps MEOS C function temporal_split_each_n_spans.
func TemporalSplitEachNSpans(temp *Temporal, elem_count int, count unsafe.Pointer) *Span {
	_cret := C.temporal_split_each_n_spans(temp._inner, C.int(elem_count), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// TemporalSplitNSpans wraps MEOS C function temporal_split_n_spans.
func TemporalSplitNSpans(temp *Temporal, span_count int, count unsafe.Pointer) *Span {
	_cret := C.temporal_split_n_spans(temp._inner, C.int(span_count), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// TnumberSplitEachNTboxes wraps MEOS C function tnumber_split_each_n_tboxes.
func TnumberSplitEachNTboxes(temp *Temporal, elem_count int, count unsafe.Pointer) *TBox {
	_cret := C.tnumber_split_each_n_tboxes(temp._inner, C.int(elem_count), (*C.int)(unsafe.Pointer(count)))
	return &TBox{_inner: _cret}
}


// TnumberSplitNTboxes wraps MEOS C function tnumber_split_n_tboxes.
func TnumberSplitNTboxes(temp *Temporal, box_count int, count unsafe.Pointer) *TBox {
	_cret := C.tnumber_split_n_tboxes(temp._inner, C.int(box_count), (*C.int)(unsafe.Pointer(count)))
	return &TBox{_inner: _cret}
}


// TnumberTboxes wraps MEOS C function tnumber_tboxes.
func TnumberTboxes(temp *Temporal, count unsafe.Pointer) *TBox {
	_cret := C.tnumber_tboxes(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return &TBox{_inner: _cret}
}


// AdjacentNumspanTnumber wraps MEOS C function adjacent_numspan_tnumber.
func AdjacentNumspanTnumber(s *Span, temp *Temporal) bool {
	_cret := C.adjacent_numspan_tnumber(s._inner, temp._inner)
	return bool(_cret)
}


// AdjacentTBOXTnumber wraps MEOS C function adjacent_tbox_tnumber.
func AdjacentTBOXTnumber(box *TBox, temp *Temporal) bool {
	_cret := C.adjacent_tbox_tnumber(box._inner, temp._inner)
	return bool(_cret)
}


// AdjacentTemporalTemporal wraps MEOS C function adjacent_temporal_temporal.
func AdjacentTemporalTemporal(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.adjacent_temporal_temporal(temp1._inner, temp2._inner)
	return bool(_cret)
}


// AdjacentTemporalTstzspan wraps MEOS C function adjacent_temporal_tstzspan.
func AdjacentTemporalTstzspan(temp *Temporal, s *Span) bool {
	_cret := C.adjacent_temporal_tstzspan(temp._inner, s._inner)
	return bool(_cret)
}


// AdjacentTnumberNumspan wraps MEOS C function adjacent_tnumber_numspan.
func AdjacentTnumberNumspan(temp *Temporal, s *Span) bool {
	_cret := C.adjacent_tnumber_numspan(temp._inner, s._inner)
	return bool(_cret)
}


// AdjacentTnumberTBOX wraps MEOS C function adjacent_tnumber_tbox.
func AdjacentTnumberTBOX(temp *Temporal, box *TBox) bool {
	_cret := C.adjacent_tnumber_tbox(temp._inner, box._inner)
	return bool(_cret)
}


// AdjacentTnumberTnumber wraps MEOS C function adjacent_tnumber_tnumber.
func AdjacentTnumberTnumber(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.adjacent_tnumber_tnumber(temp1._inner, temp2._inner)
	return bool(_cret)
}


// AdjacentTstzspanTemporal wraps MEOS C function adjacent_tstzspan_temporal.
func AdjacentTstzspanTemporal(s *Span, temp *Temporal) bool {
	_cret := C.adjacent_tstzspan_temporal(s._inner, temp._inner)
	return bool(_cret)
}


// ContainedNumspanTnumber wraps MEOS C function contained_numspan_tnumber.
func ContainedNumspanTnumber(s *Span, temp *Temporal) bool {
	_cret := C.contained_numspan_tnumber(s._inner, temp._inner)
	return bool(_cret)
}


// ContainedTBOXTnumber wraps MEOS C function contained_tbox_tnumber.
func ContainedTBOXTnumber(box *TBox, temp *Temporal) bool {
	_cret := C.contained_tbox_tnumber(box._inner, temp._inner)
	return bool(_cret)
}


// ContainedTemporalTemporal wraps MEOS C function contained_temporal_temporal.
func ContainedTemporalTemporal(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.contained_temporal_temporal(temp1._inner, temp2._inner)
	return bool(_cret)
}


// ContainedTemporalTstzspan wraps MEOS C function contained_temporal_tstzspan.
func ContainedTemporalTstzspan(temp *Temporal, s *Span) bool {
	_cret := C.contained_temporal_tstzspan(temp._inner, s._inner)
	return bool(_cret)
}


// ContainedTnumberNumspan wraps MEOS C function contained_tnumber_numspan.
func ContainedTnumberNumspan(temp *Temporal, s *Span) bool {
	_cret := C.contained_tnumber_numspan(temp._inner, s._inner)
	return bool(_cret)
}


// ContainedTnumberTBOX wraps MEOS C function contained_tnumber_tbox.
func ContainedTnumberTBOX(temp *Temporal, box *TBox) bool {
	_cret := C.contained_tnumber_tbox(temp._inner, box._inner)
	return bool(_cret)
}


// ContainedTnumberTnumber wraps MEOS C function contained_tnumber_tnumber.
func ContainedTnumberTnumber(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.contained_tnumber_tnumber(temp1._inner, temp2._inner)
	return bool(_cret)
}


// ContainedTstzspanTemporal wraps MEOS C function contained_tstzspan_temporal.
func ContainedTstzspanTemporal(s *Span, temp *Temporal) bool {
	_cret := C.contained_tstzspan_temporal(s._inner, temp._inner)
	return bool(_cret)
}


// ContainsNumspanTnumber wraps MEOS C function contains_numspan_tnumber.
func ContainsNumspanTnumber(s *Span, temp *Temporal) bool {
	_cret := C.contains_numspan_tnumber(s._inner, temp._inner)
	return bool(_cret)
}


// ContainsTBOXTnumber wraps MEOS C function contains_tbox_tnumber.
func ContainsTBOXTnumber(box *TBox, temp *Temporal) bool {
	_cret := C.contains_tbox_tnumber(box._inner, temp._inner)
	return bool(_cret)
}


// ContainsTemporalTstzspan wraps MEOS C function contains_temporal_tstzspan.
func ContainsTemporalTstzspan(temp *Temporal, s *Span) bool {
	_cret := C.contains_temporal_tstzspan(temp._inner, s._inner)
	return bool(_cret)
}


// ContainsTemporalTemporal wraps MEOS C function contains_temporal_temporal.
func ContainsTemporalTemporal(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.contains_temporal_temporal(temp1._inner, temp2._inner)
	return bool(_cret)
}


// ContainsTnumberNumspan wraps MEOS C function contains_tnumber_numspan.
func ContainsTnumberNumspan(temp *Temporal, s *Span) bool {
	_cret := C.contains_tnumber_numspan(temp._inner, s._inner)
	return bool(_cret)
}


// ContainsTnumberTBOX wraps MEOS C function contains_tnumber_tbox.
func ContainsTnumberTBOX(temp *Temporal, box *TBox) bool {
	_cret := C.contains_tnumber_tbox(temp._inner, box._inner)
	return bool(_cret)
}


// ContainsTnumberTnumber wraps MEOS C function contains_tnumber_tnumber.
func ContainsTnumberTnumber(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.contains_tnumber_tnumber(temp1._inner, temp2._inner)
	return bool(_cret)
}


// ContainsTstzspanTemporal wraps MEOS C function contains_tstzspan_temporal.
func ContainsTstzspanTemporal(s *Span, temp *Temporal) bool {
	_cret := C.contains_tstzspan_temporal(s._inner, temp._inner)
	return bool(_cret)
}


// OverlapsNumspanTnumber wraps MEOS C function overlaps_numspan_tnumber.
func OverlapsNumspanTnumber(s *Span, temp *Temporal) bool {
	_cret := C.overlaps_numspan_tnumber(s._inner, temp._inner)
	return bool(_cret)
}


// OverlapsTBOXTnumber wraps MEOS C function overlaps_tbox_tnumber.
func OverlapsTBOXTnumber(box *TBox, temp *Temporal) bool {
	_cret := C.overlaps_tbox_tnumber(box._inner, temp._inner)
	return bool(_cret)
}


// OverlapsTemporalTemporal wraps MEOS C function overlaps_temporal_temporal.
func OverlapsTemporalTemporal(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.overlaps_temporal_temporal(temp1._inner, temp2._inner)
	return bool(_cret)
}


// OverlapsTemporalTstzspan wraps MEOS C function overlaps_temporal_tstzspan.
func OverlapsTemporalTstzspan(temp *Temporal, s *Span) bool {
	_cret := C.overlaps_temporal_tstzspan(temp._inner, s._inner)
	return bool(_cret)
}


// OverlapsTnumberNumspan wraps MEOS C function overlaps_tnumber_numspan.
func OverlapsTnumberNumspan(temp *Temporal, s *Span) bool {
	_cret := C.overlaps_tnumber_numspan(temp._inner, s._inner)
	return bool(_cret)
}


// OverlapsTnumberTBOX wraps MEOS C function overlaps_tnumber_tbox.
func OverlapsTnumberTBOX(temp *Temporal, box *TBox) bool {
	_cret := C.overlaps_tnumber_tbox(temp._inner, box._inner)
	return bool(_cret)
}


// OverlapsTnumberTnumber wraps MEOS C function overlaps_tnumber_tnumber.
func OverlapsTnumberTnumber(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.overlaps_tnumber_tnumber(temp1._inner, temp2._inner)
	return bool(_cret)
}


// OverlapsTstzspanTemporal wraps MEOS C function overlaps_tstzspan_temporal.
func OverlapsTstzspanTemporal(s *Span, temp *Temporal) bool {
	_cret := C.overlaps_tstzspan_temporal(s._inner, temp._inner)
	return bool(_cret)
}


// SameNumspanTnumber wraps MEOS C function same_numspan_tnumber.
func SameNumspanTnumber(s *Span, temp *Temporal) bool {
	_cret := C.same_numspan_tnumber(s._inner, temp._inner)
	return bool(_cret)
}


// SameTBOXTnumber wraps MEOS C function same_tbox_tnumber.
func SameTBOXTnumber(box *TBox, temp *Temporal) bool {
	_cret := C.same_tbox_tnumber(box._inner, temp._inner)
	return bool(_cret)
}


// SameTemporalTemporal wraps MEOS C function same_temporal_temporal.
func SameTemporalTemporal(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.same_temporal_temporal(temp1._inner, temp2._inner)
	return bool(_cret)
}


// SameTemporalTstzspan wraps MEOS C function same_temporal_tstzspan.
func SameTemporalTstzspan(temp *Temporal, s *Span) bool {
	_cret := C.same_temporal_tstzspan(temp._inner, s._inner)
	return bool(_cret)
}


// SameTnumberNumspan wraps MEOS C function same_tnumber_numspan.
func SameTnumberNumspan(temp *Temporal, s *Span) bool {
	_cret := C.same_tnumber_numspan(temp._inner, s._inner)
	return bool(_cret)
}


// SameTnumberTBOX wraps MEOS C function same_tnumber_tbox.
func SameTnumberTBOX(temp *Temporal, box *TBox) bool {
	_cret := C.same_tnumber_tbox(temp._inner, box._inner)
	return bool(_cret)
}


// SameTnumberTnumber wraps MEOS C function same_tnumber_tnumber.
func SameTnumberTnumber(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.same_tnumber_tnumber(temp1._inner, temp2._inner)
	return bool(_cret)
}


// SameTstzspanTemporal wraps MEOS C function same_tstzspan_temporal.
func SameTstzspanTemporal(s *Span, temp *Temporal) bool {
	_cret := C.same_tstzspan_temporal(s._inner, temp._inner)
	return bool(_cret)
}


// AfterTBOXTnumber wraps MEOS C function after_tbox_tnumber.
func AfterTBOXTnumber(box *TBox, temp *Temporal) bool {
	_cret := C.after_tbox_tnumber(box._inner, temp._inner)
	return bool(_cret)
}


// AfterTemporalTstzspan wraps MEOS C function after_temporal_tstzspan.
func AfterTemporalTstzspan(temp *Temporal, s *Span) bool {
	_cret := C.after_temporal_tstzspan(temp._inner, s._inner)
	return bool(_cret)
}


// AfterTemporalTemporal wraps MEOS C function after_temporal_temporal.
func AfterTemporalTemporal(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.after_temporal_temporal(temp1._inner, temp2._inner)
	return bool(_cret)
}


// AfterTnumberTBOX wraps MEOS C function after_tnumber_tbox.
func AfterTnumberTBOX(temp *Temporal, box *TBox) bool {
	_cret := C.after_tnumber_tbox(temp._inner, box._inner)
	return bool(_cret)
}


// AfterTnumberTnumber wraps MEOS C function after_tnumber_tnumber.
func AfterTnumberTnumber(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.after_tnumber_tnumber(temp1._inner, temp2._inner)
	return bool(_cret)
}


// AfterTstzspanTemporal wraps MEOS C function after_tstzspan_temporal.
func AfterTstzspanTemporal(s *Span, temp *Temporal) bool {
	_cret := C.after_tstzspan_temporal(s._inner, temp._inner)
	return bool(_cret)
}


// BeforeTBOXTnumber wraps MEOS C function before_tbox_tnumber.
func BeforeTBOXTnumber(box *TBox, temp *Temporal) bool {
	_cret := C.before_tbox_tnumber(box._inner, temp._inner)
	return bool(_cret)
}


// BeforeTemporalTstzspan wraps MEOS C function before_temporal_tstzspan.
func BeforeTemporalTstzspan(temp *Temporal, s *Span) bool {
	_cret := C.before_temporal_tstzspan(temp._inner, s._inner)
	return bool(_cret)
}


// BeforeTemporalTemporal wraps MEOS C function before_temporal_temporal.
func BeforeTemporalTemporal(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.before_temporal_temporal(temp1._inner, temp2._inner)
	return bool(_cret)
}


// BeforeTnumberTBOX wraps MEOS C function before_tnumber_tbox.
func BeforeTnumberTBOX(temp *Temporal, box *TBox) bool {
	_cret := C.before_tnumber_tbox(temp._inner, box._inner)
	return bool(_cret)
}


// BeforeTnumberTnumber wraps MEOS C function before_tnumber_tnumber.
func BeforeTnumberTnumber(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.before_tnumber_tnumber(temp1._inner, temp2._inner)
	return bool(_cret)
}


// BeforeTstzspanTemporal wraps MEOS C function before_tstzspan_temporal.
func BeforeTstzspanTemporal(s *Span, temp *Temporal) bool {
	_cret := C.before_tstzspan_temporal(s._inner, temp._inner)
	return bool(_cret)
}


// LeftTBOXTnumber wraps MEOS C function left_tbox_tnumber.
func LeftTBOXTnumber(box *TBox, temp *Temporal) bool {
	_cret := C.left_tbox_tnumber(box._inner, temp._inner)
	return bool(_cret)
}


// LeftNumspanTnumber wraps MEOS C function left_numspan_tnumber.
func LeftNumspanTnumber(s *Span, temp *Temporal) bool {
	_cret := C.left_numspan_tnumber(s._inner, temp._inner)
	return bool(_cret)
}


// LeftTnumberNumspan wraps MEOS C function left_tnumber_numspan.
func LeftTnumberNumspan(temp *Temporal, s *Span) bool {
	_cret := C.left_tnumber_numspan(temp._inner, s._inner)
	return bool(_cret)
}


// LeftTnumberTBOX wraps MEOS C function left_tnumber_tbox.
func LeftTnumberTBOX(temp *Temporal, box *TBox) bool {
	_cret := C.left_tnumber_tbox(temp._inner, box._inner)
	return bool(_cret)
}


// LeftTnumberTnumber wraps MEOS C function left_tnumber_tnumber.
func LeftTnumberTnumber(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.left_tnumber_tnumber(temp1._inner, temp2._inner)
	return bool(_cret)
}


// OverafterTBOXTnumber wraps MEOS C function overafter_tbox_tnumber.
func OverafterTBOXTnumber(box *TBox, temp *Temporal) bool {
	_cret := C.overafter_tbox_tnumber(box._inner, temp._inner)
	return bool(_cret)
}


// OverafterTemporalTstzspan wraps MEOS C function overafter_temporal_tstzspan.
func OverafterTemporalTstzspan(temp *Temporal, s *Span) bool {
	_cret := C.overafter_temporal_tstzspan(temp._inner, s._inner)
	return bool(_cret)
}


// OverafterTemporalTemporal wraps MEOS C function overafter_temporal_temporal.
func OverafterTemporalTemporal(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.overafter_temporal_temporal(temp1._inner, temp2._inner)
	return bool(_cret)
}


// OverafterTnumberTBOX wraps MEOS C function overafter_tnumber_tbox.
func OverafterTnumberTBOX(temp *Temporal, box *TBox) bool {
	_cret := C.overafter_tnumber_tbox(temp._inner, box._inner)
	return bool(_cret)
}


// OverafterTnumberTnumber wraps MEOS C function overafter_tnumber_tnumber.
func OverafterTnumberTnumber(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.overafter_tnumber_tnumber(temp1._inner, temp2._inner)
	return bool(_cret)
}


// OverafterTstzspanTemporal wraps MEOS C function overafter_tstzspan_temporal.
func OverafterTstzspanTemporal(s *Span, temp *Temporal) bool {
	_cret := C.overafter_tstzspan_temporal(s._inner, temp._inner)
	return bool(_cret)
}


// OverbeforeTBOXTnumber wraps MEOS C function overbefore_tbox_tnumber.
func OverbeforeTBOXTnumber(box *TBox, temp *Temporal) bool {
	_cret := C.overbefore_tbox_tnumber(box._inner, temp._inner)
	return bool(_cret)
}


// OverbeforeTemporalTstzspan wraps MEOS C function overbefore_temporal_tstzspan.
func OverbeforeTemporalTstzspan(temp *Temporal, s *Span) bool {
	_cret := C.overbefore_temporal_tstzspan(temp._inner, s._inner)
	return bool(_cret)
}


// OverbeforeTemporalTemporal wraps MEOS C function overbefore_temporal_temporal.
func OverbeforeTemporalTemporal(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.overbefore_temporal_temporal(temp1._inner, temp2._inner)
	return bool(_cret)
}


// OverbeforeTnumberTBOX wraps MEOS C function overbefore_tnumber_tbox.
func OverbeforeTnumberTBOX(temp *Temporal, box *TBox) bool {
	_cret := C.overbefore_tnumber_tbox(temp._inner, box._inner)
	return bool(_cret)
}


// OverbeforeTnumberTnumber wraps MEOS C function overbefore_tnumber_tnumber.
func OverbeforeTnumberTnumber(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.overbefore_tnumber_tnumber(temp1._inner, temp2._inner)
	return bool(_cret)
}


// OverbeforeTstzspanTemporal wraps MEOS C function overbefore_tstzspan_temporal.
func OverbeforeTstzspanTemporal(s *Span, temp *Temporal) bool {
	_cret := C.overbefore_tstzspan_temporal(s._inner, temp._inner)
	return bool(_cret)
}


// OverleftNumspanTnumber wraps MEOS C function overleft_numspan_tnumber.
func OverleftNumspanTnumber(s *Span, temp *Temporal) bool {
	_cret := C.overleft_numspan_tnumber(s._inner, temp._inner)
	return bool(_cret)
}


// OverleftTBOXTnumber wraps MEOS C function overleft_tbox_tnumber.
func OverleftTBOXTnumber(box *TBox, temp *Temporal) bool {
	_cret := C.overleft_tbox_tnumber(box._inner, temp._inner)
	return bool(_cret)
}


// OverleftTnumberNumspan wraps MEOS C function overleft_tnumber_numspan.
func OverleftTnumberNumspan(temp *Temporal, s *Span) bool {
	_cret := C.overleft_tnumber_numspan(temp._inner, s._inner)
	return bool(_cret)
}


// OverleftTnumberTBOX wraps MEOS C function overleft_tnumber_tbox.
func OverleftTnumberTBOX(temp *Temporal, box *TBox) bool {
	_cret := C.overleft_tnumber_tbox(temp._inner, box._inner)
	return bool(_cret)
}


// OverleftTnumberTnumber wraps MEOS C function overleft_tnumber_tnumber.
func OverleftTnumberTnumber(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.overleft_tnumber_tnumber(temp1._inner, temp2._inner)
	return bool(_cret)
}


// OverrightNumspanTnumber wraps MEOS C function overright_numspan_tnumber.
func OverrightNumspanTnumber(s *Span, temp *Temporal) bool {
	_cret := C.overright_numspan_tnumber(s._inner, temp._inner)
	return bool(_cret)
}


// OverrightTBOXTnumber wraps MEOS C function overright_tbox_tnumber.
func OverrightTBOXTnumber(box *TBox, temp *Temporal) bool {
	_cret := C.overright_tbox_tnumber(box._inner, temp._inner)
	return bool(_cret)
}


// OverrightTnumberNumspan wraps MEOS C function overright_tnumber_numspan.
func OverrightTnumberNumspan(temp *Temporal, s *Span) bool {
	_cret := C.overright_tnumber_numspan(temp._inner, s._inner)
	return bool(_cret)
}


// OverrightTnumberTBOX wraps MEOS C function overright_tnumber_tbox.
func OverrightTnumberTBOX(temp *Temporal, box *TBox) bool {
	_cret := C.overright_tnumber_tbox(temp._inner, box._inner)
	return bool(_cret)
}


// OverrightTnumberTnumber wraps MEOS C function overright_tnumber_tnumber.
func OverrightTnumberTnumber(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.overright_tnumber_tnumber(temp1._inner, temp2._inner)
	return bool(_cret)
}


// RightNumspanTnumber wraps MEOS C function right_numspan_tnumber.
func RightNumspanTnumber(s *Span, temp *Temporal) bool {
	_cret := C.right_numspan_tnumber(s._inner, temp._inner)
	return bool(_cret)
}


// RightTBOXTnumber wraps MEOS C function right_tbox_tnumber.
func RightTBOXTnumber(box *TBox, temp *Temporal) bool {
	_cret := C.right_tbox_tnumber(box._inner, temp._inner)
	return bool(_cret)
}


// RightTnumberNumspan wraps MEOS C function right_tnumber_numspan.
func RightTnumberNumspan(temp *Temporal, s *Span) bool {
	_cret := C.right_tnumber_numspan(temp._inner, s._inner)
	return bool(_cret)
}


// RightTnumberTBOX wraps MEOS C function right_tnumber_tbox.
func RightTnumberTBOX(temp *Temporal, box *TBox) bool {
	_cret := C.right_tnumber_tbox(temp._inner, box._inner)
	return bool(_cret)
}


// RightTnumberTnumber wraps MEOS C function right_tnumber_tnumber.
func RightTnumberTnumber(temp1 *Temporal, temp2 *Temporal) bool {
	_cret := C.right_tnumber_tnumber(temp1._inner, temp2._inner)
	return bool(_cret)
}


// TandBoolTbool wraps MEOS C function tand_bool_tbool.
func TandBoolTbool(b bool, temp *Temporal) *Temporal {
	_cret := C.tand_bool_tbool(C.bool(b), temp._inner)
	return &Temporal{_inner: _cret}
}


// TandTboolBool wraps MEOS C function tand_tbool_bool.
func TandTboolBool(temp *Temporal, b bool) *Temporal {
	_cret := C.tand_tbool_bool(temp._inner, C.bool(b))
	return &Temporal{_inner: _cret}
}


// TandTboolTbool wraps MEOS C function tand_tbool_tbool.
func TandTboolTbool(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tand_tbool_tbool(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TboolWhenTrue wraps MEOS C function tbool_when_true.
func TboolWhenTrue(temp *Temporal) *SpanSet {
	_cret := C.tbool_when_true(temp._inner)
	return &SpanSet{_inner: _cret}
}


// TnotTbool wraps MEOS C function tnot_tbool.
func TnotTbool(temp *Temporal) *Temporal {
	_cret := C.tnot_tbool(temp._inner)
	return &Temporal{_inner: _cret}
}


// TorBoolTbool wraps MEOS C function tor_bool_tbool.
func TorBoolTbool(b bool, temp *Temporal) *Temporal {
	_cret := C.tor_bool_tbool(C.bool(b), temp._inner)
	return &Temporal{_inner: _cret}
}


// TorTboolBool wraps MEOS C function tor_tbool_bool.
func TorTboolBool(temp *Temporal, b bool) *Temporal {
	_cret := C.tor_tbool_bool(temp._inner, C.bool(b))
	return &Temporal{_inner: _cret}
}


// TorTboolTbool wraps MEOS C function tor_tbool_tbool.
func TorTboolTbool(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tor_tbool_tbool(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// AddFloatTfloat wraps MEOS C function add_float_tfloat.
func AddFloatTfloat(d float64, tnumber *Temporal) *Temporal {
	_cret := C.add_float_tfloat(C.double(d), tnumber._inner)
	return &Temporal{_inner: _cret}
}


// AddIntTint wraps MEOS C function add_int_tint.
func AddIntTint(i int, tnumber *Temporal) *Temporal {
	_cret := C.add_int_tint(C.int(i), tnumber._inner)
	return &Temporal{_inner: _cret}
}


// AddTfloatFloat wraps MEOS C function add_tfloat_float.
func AddTfloatFloat(tnumber *Temporal, d float64) *Temporal {
	_cret := C.add_tfloat_float(tnumber._inner, C.double(d))
	return &Temporal{_inner: _cret}
}


// AddTintInt wraps MEOS C function add_tint_int.
func AddTintInt(tnumber *Temporal, i int) *Temporal {
	_cret := C.add_tint_int(tnumber._inner, C.int(i))
	return &Temporal{_inner: _cret}
}


// AddBigintTbigint wraps MEOS C function add_bigint_tbigint.
func AddBigintTbigint(i int64, tnumber *Temporal) *Temporal {
	_cret := C.add_bigint_tbigint(C.int64_t(i), tnumber._inner)
	return &Temporal{_inner: _cret}
}


// AddTbigintBigint wraps MEOS C function add_tbigint_bigint.
func AddTbigintBigint(tnumber *Temporal, i int64) *Temporal {
	_cret := C.add_tbigint_bigint(tnumber._inner, C.int64_t(i))
	return &Temporal{_inner: _cret}
}


// AddTnumberTnumber wraps MEOS C function add_tnumber_tnumber.
func AddTnumberTnumber(tnumber1 *Temporal, tnumber2 *Temporal) *Temporal {
	_cret := C.add_tnumber_tnumber(tnumber1._inner, tnumber2._inner)
	return &Temporal{_inner: _cret}
}


// DivFloatTfloat wraps MEOS C function div_float_tfloat.
func DivFloatTfloat(d float64, tnumber *Temporal) *Temporal {
	_cret := C.div_float_tfloat(C.double(d), tnumber._inner)
	return &Temporal{_inner: _cret}
}


// DivIntTint wraps MEOS C function div_int_tint.
func DivIntTint(i int, tnumber *Temporal) *Temporal {
	_cret := C.div_int_tint(C.int(i), tnumber._inner)
	return &Temporal{_inner: _cret}
}


// DivTfloatFloat wraps MEOS C function div_tfloat_float.
func DivTfloatFloat(tnumber *Temporal, d float64) *Temporal {
	_cret := C.div_tfloat_float(tnumber._inner, C.double(d))
	return &Temporal{_inner: _cret}
}


// DivTintInt wraps MEOS C function div_tint_int.
func DivTintInt(tnumber *Temporal, i int) *Temporal {
	_cret := C.div_tint_int(tnumber._inner, C.int(i))
	return &Temporal{_inner: _cret}
}


// DivBigintTbigint wraps MEOS C function div_bigint_tbigint.
func DivBigintTbigint(i int64, tnumber *Temporal) *Temporal {
	_cret := C.div_bigint_tbigint(C.int64_t(i), tnumber._inner)
	return &Temporal{_inner: _cret}
}


// DivTbigintBigint wraps MEOS C function div_tbigint_bigint.
func DivTbigintBigint(tnumber *Temporal, i int64) *Temporal {
	_cret := C.div_tbigint_bigint(tnumber._inner, C.int64_t(i))
	return &Temporal{_inner: _cret}
}


// DivTnumberTnumber wraps MEOS C function div_tnumber_tnumber.
func DivTnumberTnumber(tnumber1 *Temporal, tnumber2 *Temporal) *Temporal {
	_cret := C.div_tnumber_tnumber(tnumber1._inner, tnumber2._inner)
	return &Temporal{_inner: _cret}
}


// MulFloatTfloat wraps MEOS C function mul_float_tfloat.
func MulFloatTfloat(d float64, tnumber *Temporal) *Temporal {
	_cret := C.mul_float_tfloat(C.double(d), tnumber._inner)
	return &Temporal{_inner: _cret}
}


// MulIntTint wraps MEOS C function mul_int_tint.
func MulIntTint(i int, tnumber *Temporal) *Temporal {
	_cret := C.mul_int_tint(C.int(i), tnumber._inner)
	return &Temporal{_inner: _cret}
}


// MulTfloatFloat wraps MEOS C function mul_tfloat_float.
func MulTfloatFloat(tnumber *Temporal, d float64) *Temporal {
	_cret := C.mul_tfloat_float(tnumber._inner, C.double(d))
	return &Temporal{_inner: _cret}
}


// MulTintInt wraps MEOS C function mul_tint_int.
func MulTintInt(tnumber *Temporal, i int) *Temporal {
	_cret := C.mul_tint_int(tnumber._inner, C.int(i))
	return &Temporal{_inner: _cret}
}


// MulBigintTbigint wraps MEOS C function mul_bigint_tbigint.
func MulBigintTbigint(i int64, tnumber *Temporal) *Temporal {
	_cret := C.mul_bigint_tbigint(C.int64_t(i), tnumber._inner)
	return &Temporal{_inner: _cret}
}


// MulTbigintBigint wraps MEOS C function mul_tbigint_bigint.
func MulTbigintBigint(tnumber *Temporal, i int64) *Temporal {
	_cret := C.mul_tbigint_bigint(tnumber._inner, C.int64_t(i))
	return &Temporal{_inner: _cret}
}


// MulTnumberTnumber wraps MEOS C function mul_tnumber_tnumber.
func MulTnumberTnumber(tnumber1 *Temporal, tnumber2 *Temporal) *Temporal {
	_cret := C.mul_tnumber_tnumber(tnumber1._inner, tnumber2._inner)
	return &Temporal{_inner: _cret}
}


// SubFloatTfloat wraps MEOS C function sub_float_tfloat.
func SubFloatTfloat(d float64, tnumber *Temporal) *Temporal {
	_cret := C.sub_float_tfloat(C.double(d), tnumber._inner)
	return &Temporal{_inner: _cret}
}


// SubIntTint wraps MEOS C function sub_int_tint.
func SubIntTint(i int, tnumber *Temporal) *Temporal {
	_cret := C.sub_int_tint(C.int(i), tnumber._inner)
	return &Temporal{_inner: _cret}
}


// SubTfloatFloat wraps MEOS C function sub_tfloat_float.
func SubTfloatFloat(tnumber *Temporal, d float64) *Temporal {
	_cret := C.sub_tfloat_float(tnumber._inner, C.double(d))
	return &Temporal{_inner: _cret}
}


// SubTintInt wraps MEOS C function sub_tint_int.
func SubTintInt(tnumber *Temporal, i int) *Temporal {
	_cret := C.sub_tint_int(tnumber._inner, C.int(i))
	return &Temporal{_inner: _cret}
}


// SubBigintTbigint wraps MEOS C function sub_bigint_tbigint.
func SubBigintTbigint(i int64, tnumber *Temporal) *Temporal {
	_cret := C.sub_bigint_tbigint(C.int64_t(i), tnumber._inner)
	return &Temporal{_inner: _cret}
}


// SubTbigintBigint wraps MEOS C function sub_tbigint_bigint.
func SubTbigintBigint(tnumber *Temporal, i int64) *Temporal {
	_cret := C.sub_tbigint_bigint(tnumber._inner, C.int64_t(i))
	return &Temporal{_inner: _cret}
}


// SubTnumberTnumber wraps MEOS C function sub_tnumber_tnumber.
func SubTnumberTnumber(tnumber1 *Temporal, tnumber2 *Temporal) *Temporal {
	_cret := C.sub_tnumber_tnumber(tnumber1._inner, tnumber2._inner)
	return &Temporal{_inner: _cret}
}


// TemporalDerivative wraps MEOS C function temporal_derivative.
func TemporalDerivative(temp *Temporal) *Temporal {
	_cret := C.temporal_derivative(temp._inner)
	return &Temporal{_inner: _cret}
}


// TfloatExp wraps MEOS C function tfloat_exp.
func TfloatExp(temp *Temporal) *Temporal {
	_cret := C.tfloat_exp(temp._inner)
	return &Temporal{_inner: _cret}
}


// TfloatLn wraps MEOS C function tfloat_ln.
func TfloatLn(temp *Temporal) *Temporal {
	_cret := C.tfloat_ln(temp._inner)
	return &Temporal{_inner: _cret}
}


// TfloatLog10 wraps MEOS C function tfloat_log10.
func TfloatLog10(temp *Temporal) *Temporal {
	_cret := C.tfloat_log10(temp._inner)
	return &Temporal{_inner: _cret}
}


// TfloatSin wraps MEOS C function tfloat_sin.
func TfloatSin(temp *Temporal) *Temporal {
	_cret := C.tfloat_sin(temp._inner)
	return &Temporal{_inner: _cret}
}


// TfloatCos wraps MEOS C function tfloat_cos.
func TfloatCos(temp *Temporal) *Temporal {
	_cret := C.tfloat_cos(temp._inner)
	return &Temporal{_inner: _cret}
}


// TfloatTan wraps MEOS C function tfloat_tan.
func TfloatTan(temp *Temporal) *Temporal {
	_cret := C.tfloat_tan(temp._inner)
	return &Temporal{_inner: _cret}
}


// TnumberAbs wraps MEOS C function tnumber_abs.
func TnumberAbs(temp *Temporal) *Temporal {
	_cret := C.tnumber_abs(temp._inner)
	return &Temporal{_inner: _cret}
}


// TnumberTrend wraps MEOS C function tnumber_trend.
func TnumberTrend(temp *Temporal) *Temporal {
	_cret := C.tnumber_trend(temp._inner)
	return &Temporal{_inner: _cret}
}


// FloatAngularDifference wraps MEOS C function float_angular_difference.
func FloatAngularDifference(degrees1 float64, degrees2 float64) float64 {
	_cret := C.float_angular_difference(C.double(degrees1), C.double(degrees2))
	return float64(_cret)
}


// TnumberAngularDifference wraps MEOS C function tnumber_angular_difference.
func TnumberAngularDifference(temp *Temporal) *Temporal {
	_cret := C.tnumber_angular_difference(temp._inner)
	return &Temporal{_inner: _cret}
}


// TnumberDeltaValue wraps MEOS C function tnumber_delta_value.
func TnumberDeltaValue(temp *Temporal) *Temporal {
	_cret := C.tnumber_delta_value(temp._inner)
	return &Temporal{_inner: _cret}
}


// TextcatTextTtext wraps MEOS C function textcat_text_ttext.
func TextcatTextTtext(txt string, temp *Temporal) *Temporal {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.textcat_text_ttext(_c_txt, temp._inner)
	return &Temporal{_inner: _cret}
}


// TextcatTtextText wraps MEOS C function textcat_ttext_text.
func TextcatTtextText(temp *Temporal, txt string) *Temporal {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.textcat_ttext_text(temp._inner, _c_txt)
	return &Temporal{_inner: _cret}
}


// TextcatTtextTtext wraps MEOS C function textcat_ttext_ttext.
func TextcatTtextTtext(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.textcat_ttext_ttext(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// TtextInitcap wraps MEOS C function ttext_initcap.
func TtextInitcap(temp *Temporal) *Temporal {
	_cret := C.ttext_initcap(temp._inner)
	return &Temporal{_inner: _cret}
}


// TtextUpper wraps MEOS C function ttext_upper.
func TtextUpper(temp *Temporal) *Temporal {
	_cret := C.ttext_upper(temp._inner)
	return &Temporal{_inner: _cret}
}


// TtextLower wraps MEOS C function ttext_lower.
func TtextLower(temp *Temporal) *Temporal {
	_cret := C.ttext_lower(temp._inner)
	return &Temporal{_inner: _cret}
}


// TdistanceTfloatFloat wraps MEOS C function tdistance_tfloat_float.
func TdistanceTfloatFloat(temp *Temporal, d float64) *Temporal {
	_cret := C.tdistance_tfloat_float(temp._inner, C.double(d))
	return &Temporal{_inner: _cret}
}


// TdistanceTintInt wraps MEOS C function tdistance_tint_int.
func TdistanceTintInt(temp *Temporal, i int) *Temporal {
	_cret := C.tdistance_tint_int(temp._inner, C.int(i))
	return &Temporal{_inner: _cret}
}


// TdistanceTnumberTnumber wraps MEOS C function tdistance_tnumber_tnumber.
func TdistanceTnumberTnumber(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.tdistance_tnumber_tnumber(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// NadTboxfloatTboxfloat wraps MEOS C function nad_tboxfloat_tboxfloat.
func NadTboxfloatTboxfloat(box1 *TBox, box2 *TBox) float64 {
	_cret := C.nad_tboxfloat_tboxfloat(box1._inner, box2._inner)
	return float64(_cret)
}


// NadTboxintTboxint wraps MEOS C function nad_tboxint_tboxint.
func NadTboxintTboxint(box1 *TBox, box2 *TBox) int {
	_cret := C.nad_tboxint_tboxint(box1._inner, box2._inner)
	return int(_cret)
}


// NadTfloatFloat wraps MEOS C function nad_tfloat_float.
func NadTfloatFloat(temp *Temporal, d float64) float64 {
	_cret := C.nad_tfloat_float(temp._inner, C.double(d))
	return float64(_cret)
}


// NadTfloatTfloat wraps MEOS C function nad_tfloat_tfloat.
func NadTfloatTfloat(temp1 *Temporal, temp2 *Temporal) float64 {
	_cret := C.nad_tfloat_tfloat(temp1._inner, temp2._inner)
	return float64(_cret)
}


// NadTfloatTBOX wraps MEOS C function nad_tfloat_tbox.
func NadTfloatTBOX(temp *Temporal, box *TBox) float64 {
	_cret := C.nad_tfloat_tbox(temp._inner, box._inner)
	return float64(_cret)
}


// NadTintInt wraps MEOS C function nad_tint_int.
func NadTintInt(temp *Temporal, i int) int {
	_cret := C.nad_tint_int(temp._inner, C.int(i))
	return int(_cret)
}


// NadTintTBOX wraps MEOS C function nad_tint_tbox.
func NadTintTBOX(temp *Temporal, box *TBox) int {
	_cret := C.nad_tint_tbox(temp._inner, box._inner)
	return int(_cret)
}


// NadTintTint wraps MEOS C function nad_tint_tint.
func NadTintTint(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.nad_tint_tint(temp1._inner, temp2._inner)
	return int(_cret)
}


// TboolTandTransfn wraps MEOS C function tbool_tand_transfn.
func TboolTandTransfn(state *SkipList, temp *Temporal) *SkipList {
	_cret := C.tbool_tand_transfn(state._inner, temp._inner)
	return &SkipList{_inner: _cret}
}


// TboolTandCombinefn wraps MEOS C function tbool_tand_combinefn.
func TboolTandCombinefn(state1 *SkipList, state2 *SkipList) *SkipList {
	_cret := C.tbool_tand_combinefn(state1._inner, state2._inner)
	return &SkipList{_inner: _cret}
}


// TboolTorTransfn wraps MEOS C function tbool_tor_transfn.
func TboolTorTransfn(state *SkipList, temp *Temporal) *SkipList {
	_cret := C.tbool_tor_transfn(state._inner, temp._inner)
	return &SkipList{_inner: _cret}
}


// TboolTorCombinefn wraps MEOS C function tbool_tor_combinefn.
func TboolTorCombinefn(state1 *SkipList, state2 *SkipList) *SkipList {
	_cret := C.tbool_tor_combinefn(state1._inner, state2._inner)
	return &SkipList{_inner: _cret}
}


// TemporalExtentTransfn wraps MEOS C function temporal_extent_transfn.
func TemporalExtentTransfn(s *Span, temp *Temporal) *Span {
	_cret := C.temporal_extent_transfn(s._inner, temp._inner)
	return &Span{_inner: _cret}
}


// TemporalTaggFinalfn wraps MEOS C function temporal_tagg_finalfn.
func TemporalTaggFinalfn(state *SkipList) *Temporal {
	_cret := C.temporal_tagg_finalfn(state._inner)
	return &Temporal{_inner: _cret}
}


// TemporalTcountTransfn wraps MEOS C function temporal_tcount_transfn.
func TemporalTcountTransfn(state *SkipList, temp *Temporal) *SkipList {
	_cret := C.temporal_tcount_transfn(state._inner, temp._inner)
	return &SkipList{_inner: _cret}
}


// TemporalTcountCombinefn wraps MEOS C function temporal_tcount_combinefn.
func TemporalTcountCombinefn(state1 *SkipList, state2 *SkipList) *SkipList {
	_cret := C.temporal_tcount_combinefn(state1._inner, state2._inner)
	return &SkipList{_inner: _cret}
}


// TfloatTmaxTransfn wraps MEOS C function tfloat_tmax_transfn.
func TfloatTmaxTransfn(state *SkipList, temp *Temporal) *SkipList {
	_cret := C.tfloat_tmax_transfn(state._inner, temp._inner)
	return &SkipList{_inner: _cret}
}


// TfloatTmaxCombinefn wraps MEOS C function tfloat_tmax_combinefn.
func TfloatTmaxCombinefn(state1 *SkipList, state2 *SkipList) *SkipList {
	_cret := C.tfloat_tmax_combinefn(state1._inner, state2._inner)
	return &SkipList{_inner: _cret}
}


// TfloatTminTransfn wraps MEOS C function tfloat_tmin_transfn.
func TfloatTminTransfn(state *SkipList, temp *Temporal) *SkipList {
	_cret := C.tfloat_tmin_transfn(state._inner, temp._inner)
	return &SkipList{_inner: _cret}
}


// TfloatTminCombinefn wraps MEOS C function tfloat_tmin_combinefn.
func TfloatTminCombinefn(state1 *SkipList, state2 *SkipList) *SkipList {
	_cret := C.tfloat_tmin_combinefn(state1._inner, state2._inner)
	return &SkipList{_inner: _cret}
}


// TfloatTsumTransfn wraps MEOS C function tfloat_tsum_transfn.
func TfloatTsumTransfn(state *SkipList, temp *Temporal) *SkipList {
	_cret := C.tfloat_tsum_transfn(state._inner, temp._inner)
	return &SkipList{_inner: _cret}
}


// TfloatTsumCombinefn wraps MEOS C function tfloat_tsum_combinefn.
func TfloatTsumCombinefn(state1 *SkipList, state2 *SkipList) *SkipList {
	_cret := C.tfloat_tsum_combinefn(state1._inner, state2._inner)
	return &SkipList{_inner: _cret}
}


// TfloatWmaxTransfn wraps MEOS C function tfloat_wmax_transfn.
func TfloatWmaxTransfn(state *SkipList, temp *Temporal, interv *Interval) *SkipList {
	_cret := C.tfloat_wmax_transfn(state._inner, temp._inner, interv._inner)
	return &SkipList{_inner: _cret}
}


// TfloatWminTransfn wraps MEOS C function tfloat_wmin_transfn.
func TfloatWminTransfn(state *SkipList, temp *Temporal, interv *Interval) *SkipList {
	_cret := C.tfloat_wmin_transfn(state._inner, temp._inner, interv._inner)
	return &SkipList{_inner: _cret}
}


// TfloatWsumTransfn wraps MEOS C function tfloat_wsum_transfn.
func TfloatWsumTransfn(state *SkipList, temp *Temporal, interv *Interval) *SkipList {
	_cret := C.tfloat_wsum_transfn(state._inner, temp._inner, interv._inner)
	return &SkipList{_inner: _cret}
}


// TimestamptzTcountTransfn wraps MEOS C function timestamptz_tcount_transfn.
func TimestamptzTcountTransfn(state *SkipList, t int64) *SkipList {
	_cret := C.timestamptz_tcount_transfn(state._inner, C.TimestampTz(t))
	return &SkipList{_inner: _cret}
}


// TintTmaxTransfn wraps MEOS C function tint_tmax_transfn.
func TintTmaxTransfn(state *SkipList, temp *Temporal) *SkipList {
	_cret := C.tint_tmax_transfn(state._inner, temp._inner)
	return &SkipList{_inner: _cret}
}


// TintTmaxCombinefn wraps MEOS C function tint_tmax_combinefn.
func TintTmaxCombinefn(state1 *SkipList, state2 *SkipList) *SkipList {
	_cret := C.tint_tmax_combinefn(state1._inner, state2._inner)
	return &SkipList{_inner: _cret}
}


// TintTminTransfn wraps MEOS C function tint_tmin_transfn.
func TintTminTransfn(state *SkipList, temp *Temporal) *SkipList {
	_cret := C.tint_tmin_transfn(state._inner, temp._inner)
	return &SkipList{_inner: _cret}
}


// TintTminCombinefn wraps MEOS C function tint_tmin_combinefn.
func TintTminCombinefn(state1 *SkipList, state2 *SkipList) *SkipList {
	_cret := C.tint_tmin_combinefn(state1._inner, state2._inner)
	return &SkipList{_inner: _cret}
}


// TintTsumTransfn wraps MEOS C function tint_tsum_transfn.
func TintTsumTransfn(state *SkipList, temp *Temporal) *SkipList {
	_cret := C.tint_tsum_transfn(state._inner, temp._inner)
	return &SkipList{_inner: _cret}
}


// TintTsumCombinefn wraps MEOS C function tint_tsum_combinefn.
func TintTsumCombinefn(state1 *SkipList, state2 *SkipList) *SkipList {
	_cret := C.tint_tsum_combinefn(state1._inner, state2._inner)
	return &SkipList{_inner: _cret}
}


// TintWmaxTransfn wraps MEOS C function tint_wmax_transfn.
func TintWmaxTransfn(state *SkipList, temp *Temporal, interv *Interval) *SkipList {
	_cret := C.tint_wmax_transfn(state._inner, temp._inner, interv._inner)
	return &SkipList{_inner: _cret}
}


// TintWminTransfn wraps MEOS C function tint_wmin_transfn.
func TintWminTransfn(state *SkipList, temp *Temporal, interv *Interval) *SkipList {
	_cret := C.tint_wmin_transfn(state._inner, temp._inner, interv._inner)
	return &SkipList{_inner: _cret}
}


// TintWsumTransfn wraps MEOS C function tint_wsum_transfn.
func TintWsumTransfn(state *SkipList, temp *Temporal, interv *Interval) *SkipList {
	_cret := C.tint_wsum_transfn(state._inner, temp._inner, interv._inner)
	return &SkipList{_inner: _cret}
}


// TnumberExtentTransfn wraps MEOS C function tnumber_extent_transfn.
func TnumberExtentTransfn(box *TBox, temp *Temporal) *TBox {
	_cret := C.tnumber_extent_transfn(box._inner, temp._inner)
	return &TBox{_inner: _cret}
}


// TnumberTavgFinalfn wraps MEOS C function tnumber_tavg_finalfn.
func TnumberTavgFinalfn(state *SkipList) *Temporal {
	_cret := C.tnumber_tavg_finalfn(state._inner)
	return &Temporal{_inner: _cret}
}


// TnumberTavgTransfn wraps MEOS C function tnumber_tavg_transfn.
func TnumberTavgTransfn(state *SkipList, temp *Temporal) *SkipList {
	_cret := C.tnumber_tavg_transfn(state._inner, temp._inner)
	return &SkipList{_inner: _cret}
}


// TnumberTavgCombinefn wraps MEOS C function tnumber_tavg_combinefn.
func TnumberTavgCombinefn(state1 *SkipList, state2 *SkipList) *SkipList {
	_cret := C.tnumber_tavg_combinefn(state1._inner, state2._inner)
	return &SkipList{_inner: _cret}
}


// TnumberWavgTransfn wraps MEOS C function tnumber_wavg_transfn.
func TnumberWavgTransfn(state *SkipList, temp *Temporal, interv *Interval) *SkipList {
	_cret := C.tnumber_wavg_transfn(state._inner, temp._inner, interv._inner)
	return &SkipList{_inner: _cret}
}


// TstzsetTcountTransfn wraps MEOS C function tstzset_tcount_transfn.
func TstzsetTcountTransfn(state *SkipList, s *Set) *SkipList {
	_cret := C.tstzset_tcount_transfn(state._inner, s._inner)
	return &SkipList{_inner: _cret}
}


// TstzspanTcountTransfn wraps MEOS C function tstzspan_tcount_transfn.
func TstzspanTcountTransfn(state *SkipList, s *Span) *SkipList {
	_cret := C.tstzspan_tcount_transfn(state._inner, s._inner)
	return &SkipList{_inner: _cret}
}


// TstzspansetTcountTransfn wraps MEOS C function tstzspanset_tcount_transfn.
func TstzspansetTcountTransfn(state *SkipList, ss *SpanSet) *SkipList {
	_cret := C.tstzspanset_tcount_transfn(state._inner, ss._inner)
	return &SkipList{_inner: _cret}
}


// TemporalMergeTransfn wraps MEOS C function temporal_merge_transfn.
func TemporalMergeTransfn(state *SkipList, temp *Temporal) *SkipList {
	_cret := C.temporal_merge_transfn(state._inner, temp._inner)
	return &SkipList{_inner: _cret}
}


// TemporalMergeCombinefn wraps MEOS C function temporal_merge_combinefn.
func TemporalMergeCombinefn(state1 *SkipList, state2 *SkipList) *SkipList {
	_cret := C.temporal_merge_combinefn(state1._inner, state2._inner)
	return &SkipList{_inner: _cret}
}


// TtextTmaxTransfn wraps MEOS C function ttext_tmax_transfn.
func TtextTmaxTransfn(state *SkipList, temp *Temporal) *SkipList {
	_cret := C.ttext_tmax_transfn(state._inner, temp._inner)
	return &SkipList{_inner: _cret}
}


// TtextTmaxCombinefn wraps MEOS C function ttext_tmax_combinefn.
func TtextTmaxCombinefn(state1 *SkipList, state2 *SkipList) *SkipList {
	_cret := C.ttext_tmax_combinefn(state1._inner, state2._inner)
	return &SkipList{_inner: _cret}
}


// TtextTminTransfn wraps MEOS C function ttext_tmin_transfn.
func TtextTminTransfn(state *SkipList, temp *Temporal) *SkipList {
	_cret := C.ttext_tmin_transfn(state._inner, temp._inner)
	return &SkipList{_inner: _cret}
}


// TtextTminCombinefn wraps MEOS C function ttext_tmin_combinefn.
func TtextTminCombinefn(state1 *SkipList, state2 *SkipList) *SkipList {
	_cret := C.ttext_tmin_combinefn(state1._inner, state2._inner)
	return &SkipList{_inner: _cret}
}


// TemporalSimplifyDp wraps MEOS C function temporal_simplify_dp.
func TemporalSimplifyDp(temp *Temporal, dist float64, synchronized bool) *Temporal {
	_cret := C.temporal_simplify_dp(temp._inner, C.double(dist), C.bool(synchronized))
	return &Temporal{_inner: _cret}
}


// TemporalSimplifyMaxDist wraps MEOS C function temporal_simplify_max_dist.
func TemporalSimplifyMaxDist(temp *Temporal, dist float64, synchronized bool) *Temporal {
	_cret := C.temporal_simplify_max_dist(temp._inner, C.double(dist), C.bool(synchronized))
	return &Temporal{_inner: _cret}
}


// TemporalSimplifyMinDist wraps MEOS C function temporal_simplify_min_dist.
func TemporalSimplifyMinDist(temp *Temporal, dist float64) *Temporal {
	_cret := C.temporal_simplify_min_dist(temp._inner, C.double(dist))
	return &Temporal{_inner: _cret}
}


// TemporalSimplifyMinTdelta wraps MEOS C function temporal_simplify_min_tdelta.
func TemporalSimplifyMinTdelta(temp *Temporal, mint *Interval) *Temporal {
	_cret := C.temporal_simplify_min_tdelta(temp._inner, mint._inner)
	return &Temporal{_inner: _cret}
}


// TemporalTprecision wraps MEOS C function temporal_tprecision.
func TemporalTprecision(temp *Temporal, duration *Interval, origin int64) *Temporal {
	_cret := C.temporal_tprecision(temp._inner, duration._inner, C.TimestampTz(origin))
	return &Temporal{_inner: _cret}
}


// TemporalTsample wraps MEOS C function temporal_tsample.
func TemporalTsample(temp *Temporal, duration *Interval, origin int64, interp Interpolation) *Temporal {
	_cret := C.temporal_tsample(temp._inner, duration._inner, C.TimestampTz(origin), C.interpType(interp))
	return &Temporal{_inner: _cret}
}


// TemporalDyntimewarpDistance wraps MEOS C function temporal_dyntimewarp_distance.
func TemporalDyntimewarpDistance(temp1 *Temporal, temp2 *Temporal) float64 {
	_cret := C.temporal_dyntimewarp_distance(temp1._inner, temp2._inner)
	return float64(_cret)
}


// TemporalDyntimewarpPath wraps MEOS C function temporal_dyntimewarp_path.
func TemporalDyntimewarpPath(temp1 *Temporal, temp2 *Temporal, count unsafe.Pointer) *Match {
	_cret := C.temporal_dyntimewarp_path(temp1._inner, temp2._inner, (*C.int)(unsafe.Pointer(count)))
	return &Match{_inner: _cret}
}


// TemporalFrechetDistance wraps MEOS C function temporal_frechet_distance.
func TemporalFrechetDistance(temp1 *Temporal, temp2 *Temporal) float64 {
	_cret := C.temporal_frechet_distance(temp1._inner, temp2._inner)
	return float64(_cret)
}


// TemporalFrechetPath wraps MEOS C function temporal_frechet_path.
func TemporalFrechetPath(temp1 *Temporal, temp2 *Temporal, count unsafe.Pointer) *Match {
	_cret := C.temporal_frechet_path(temp1._inner, temp2._inner, (*C.int)(unsafe.Pointer(count)))
	return &Match{_inner: _cret}
}


// TemporalHausdorffDistance wraps MEOS C function temporal_hausdorff_distance.
func TemporalHausdorffDistance(temp1 *Temporal, temp2 *Temporal) float64 {
	_cret := C.temporal_hausdorff_distance(temp1._inner, temp2._inner)
	return float64(_cret)
}


// TemporalAverageHausdorffDistance wraps MEOS C function temporal_average_hausdorff_distance.
func TemporalAverageHausdorffDistance(temp1 *Temporal, temp2 *Temporal) float64 {
	_cret := C.temporal_average_hausdorff_distance(temp1._inner, temp2._inner)
	return float64(_cret)
}


// TemporalLcssDistance wraps MEOS C function temporal_lcss_distance.
func TemporalLcssDistance(temp1 *Temporal, temp2 *Temporal, epsilon float64) float64 {
	_cret := C.temporal_lcss_distance(temp1._inner, temp2._inner, C.double(epsilon))
	return float64(_cret)
}


// TemporalExtKalmanFilter wraps MEOS C function temporal_ext_kalman_filter.
func TemporalExtKalmanFilter(temp *Temporal, gate float64, q float64, variance float64, to_drop bool) *Temporal {
	_cret := C.temporal_ext_kalman_filter(temp._inner, C.double(gate), C.double(q), C.double(variance), C.bool(to_drop))
	return &Temporal{_inner: _cret}
}


// TemporalTimeBins wraps MEOS C function temporal_time_bins.
func TemporalTimeBins(temp *Temporal, duration *Interval, origin int64, count unsafe.Pointer) *Span {
	_cret := C.temporal_time_bins(temp._inner, duration._inner, C.TimestampTz(origin), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// TemporalTimeSplit wraps MEOS C function temporal_time_split.
func TemporalTimeSplit(temp *Temporal, duration *Interval, torigin int64, bins unsafe.Pointer, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.temporal_time_split(temp._inner, duration._inner, C.TimestampTz(torigin), (**C.TimestampTz)(unsafe.Pointer(bins)), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TfloatTimeBoxes wraps MEOS C function tfloat_time_boxes.
func TfloatTimeBoxes(temp *Temporal, duration *Interval, torigin int64, count unsafe.Pointer) *TBox {
	_cret := C.tfloat_time_boxes(temp._inner, duration._inner, C.TimestampTz(torigin), (*C.int)(unsafe.Pointer(count)))
	return &TBox{_inner: _cret}
}


// TfloatValueBins wraps MEOS C function tfloat_value_bins.
func TfloatValueBins(temp *Temporal, vsize float64, vorigin float64, count unsafe.Pointer) *Span {
	_cret := C.tfloat_value_bins(temp._inner, C.double(vsize), C.double(vorigin), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// TfloatValueBoxes wraps MEOS C function tfloat_value_boxes.
func TfloatValueBoxes(temp *Temporal, vsize float64, vorigin float64, count unsafe.Pointer) *TBox {
	_cret := C.tfloat_value_boxes(temp._inner, C.double(vsize), C.double(vorigin), (*C.int)(unsafe.Pointer(count)))
	return &TBox{_inner: _cret}
}


// TfloatValueSplit wraps MEOS C function tfloat_value_split.
func TfloatValueSplit(temp *Temporal, size float64, origin float64, bins unsafe.Pointer, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tfloat_value_split(temp._inner, C.double(size), C.double(origin), (**C.double)(unsafe.Pointer(bins)), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TfloatValueTimeBoxes wraps MEOS C function tfloat_value_time_boxes.
func TfloatValueTimeBoxes(temp *Temporal, vsize float64, duration *Interval, vorigin float64, torigin int64, count unsafe.Pointer) *TBox {
	_cret := C.tfloat_value_time_boxes(temp._inner, C.double(vsize), duration._inner, C.double(vorigin), C.TimestampTz(torigin), (*C.int)(unsafe.Pointer(count)))
	return &TBox{_inner: _cret}
}


// TfloatValueTimeSplit wraps MEOS C function tfloat_value_time_split.
func TfloatValueTimeSplit(temp *Temporal, vsize float64, duration *Interval, vorigin float64, torigin int64, value_bins unsafe.Pointer, time_bins unsafe.Pointer, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tfloat_value_time_split(temp._inner, C.double(vsize), duration._inner, C.double(vorigin), C.TimestampTz(torigin), (**C.double)(unsafe.Pointer(value_bins)), (**C.TimestampTz)(unsafe.Pointer(time_bins)), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TfloatboxTimeTiles wraps MEOS C function tfloatbox_time_tiles.
func TfloatboxTimeTiles(box *TBox, duration *Interval, torigin int64, count unsafe.Pointer) *TBox {
	_cret := C.tfloatbox_time_tiles(box._inner, duration._inner, C.TimestampTz(torigin), (*C.int)(unsafe.Pointer(count)))
	return &TBox{_inner: _cret}
}


// TfloatboxValueTiles wraps MEOS C function tfloatbox_value_tiles.
func TfloatboxValueTiles(box *TBox, vsize float64, vorigin float64, count unsafe.Pointer) *TBox {
	_cret := C.tfloatbox_value_tiles(box._inner, C.double(vsize), C.double(vorigin), (*C.int)(unsafe.Pointer(count)))
	return &TBox{_inner: _cret}
}


// TfloatboxValueTimeTiles wraps MEOS C function tfloatbox_value_time_tiles.
func TfloatboxValueTimeTiles(box *TBox, vsize float64, duration *Interval, vorigin float64, torigin int64, count unsafe.Pointer) *TBox {
	_cret := C.tfloatbox_value_time_tiles(box._inner, C.double(vsize), duration._inner, C.double(vorigin), C.TimestampTz(torigin), (*C.int)(unsafe.Pointer(count)))
	return &TBox{_inner: _cret}
}


// TintTimeBoxes wraps MEOS C function tint_time_boxes.
func TintTimeBoxes(temp *Temporal, duration *Interval, torigin int64, count unsafe.Pointer) *TBox {
	_cret := C.tint_time_boxes(temp._inner, duration._inner, C.TimestampTz(torigin), (*C.int)(unsafe.Pointer(count)))
	return &TBox{_inner: _cret}
}


// TintValueBins wraps MEOS C function tint_value_bins.
func TintValueBins(temp *Temporal, vsize int, vorigin int, count unsafe.Pointer) *Span {
	_cret := C.tint_value_bins(temp._inner, C.int(vsize), C.int(vorigin), (*C.int)(unsafe.Pointer(count)))
	return &Span{_inner: _cret}
}


// TintValueBoxes wraps MEOS C function tint_value_boxes.
func TintValueBoxes(temp *Temporal, vsize int, vorigin int, count unsafe.Pointer) *TBox {
	_cret := C.tint_value_boxes(temp._inner, C.int(vsize), C.int(vorigin), (*C.int)(unsafe.Pointer(count)))
	return &TBox{_inner: _cret}
}


// TintValueSplit wraps MEOS C function tint_value_split.
func TintValueSplit(temp *Temporal, vsize int, vorigin int, bins unsafe.Pointer, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tint_value_split(temp._inner, C.int(vsize), C.int(vorigin), (**C.int)(unsafe.Pointer(bins)), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TintValueTimeBoxes wraps MEOS C function tint_value_time_boxes.
func TintValueTimeBoxes(temp *Temporal, vsize int, duration *Interval, vorigin int, torigin int64, count unsafe.Pointer) *TBox {
	_cret := C.tint_value_time_boxes(temp._inner, C.int(vsize), duration._inner, C.int(vorigin), C.TimestampTz(torigin), (*C.int)(unsafe.Pointer(count)))
	return &TBox{_inner: _cret}
}


// TintValueTimeSplit wraps MEOS C function tint_value_time_split.
func TintValueTimeSplit(temp *Temporal, size int, duration *Interval, vorigin int, torigin int64, value_bins unsafe.Pointer, time_bins unsafe.Pointer, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tint_value_time_split(temp._inner, C.int(size), duration._inner, C.int(vorigin), C.TimestampTz(torigin), (**C.int)(unsafe.Pointer(value_bins)), (**C.TimestampTz)(unsafe.Pointer(time_bins)), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// TintboxTimeTiles wraps MEOS C function tintbox_time_tiles.
func TintboxTimeTiles(box *TBox, duration *Interval, torigin int64, count unsafe.Pointer) *TBox {
	_cret := C.tintbox_time_tiles(box._inner, duration._inner, C.TimestampTz(torigin), (*C.int)(unsafe.Pointer(count)))
	return &TBox{_inner: _cret}
}


// TintboxValueTiles wraps MEOS C function tintbox_value_tiles.
func TintboxValueTiles(box *TBox, xsize int, xorigin int, count unsafe.Pointer) *TBox {
	_cret := C.tintbox_value_tiles(box._inner, C.int(xsize), C.int(xorigin), (*C.int)(unsafe.Pointer(count)))
	return &TBox{_inner: _cret}
}


// TintboxValueTimeTiles wraps MEOS C function tintbox_value_time_tiles.
func TintboxValueTimeTiles(box *TBox, xsize int, duration *Interval, xorigin int, torigin int64, count unsafe.Pointer) *TBox {
	_cret := C.tintbox_value_time_tiles(box._inner, C.int(xsize), duration._inner, C.int(xorigin), C.TimestampTz(torigin), (*C.int)(unsafe.Pointer(count)))
	return &TBox{_inner: _cret}
}

