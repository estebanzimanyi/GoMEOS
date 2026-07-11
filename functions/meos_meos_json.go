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

// JSONIn wraps MEOS C function json_in.
func JSONIn(str string) string {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.json_in(_c_str)
	return C.GoString(C.text_to_cstring(_cret))
}


// JSONOut wraps MEOS C function json_out.
func JSONOut(js string) string {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	_cret := C.json_out(_c_js)
	return C.GoString(_cret)
}


// JsonbFromText wraps MEOS C function jsonb_from_text.
func JsonbFromText(txt string, unique_keys bool) *Jsonb {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_cret := C.jsonb_from_text(_c_txt, C.bool(unique_keys))
	return &Jsonb{_inner: _cret}
}


// JsonbIn wraps MEOS C function jsonb_in.
func JsonbIn(str string) *Jsonb {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.jsonb_in(_c_str)
	return &Jsonb{_inner: _cret}
}


// JsonbOut wraps MEOS C function jsonb_out.
func JsonbOut(jb *Jsonb) string {
	_cret := C.jsonb_out(jb._inner)
	return C.GoString(_cret)
}


// JSONMake wraps MEOS C function json_make.
func JSONMake(keys_vals unsafe.Pointer, count int) string {
	_cret := C.json_make((**C.text)(unsafe.Pointer(keys_vals)), C.int(count))
	return C.GoString(C.text_to_cstring(_cret))
}


// JSONMakeTwoArg wraps MEOS C function json_make_two_arg.
func JSONMakeTwoArg(keys unsafe.Pointer, values unsafe.Pointer, count int) string {
	_cret := C.json_make_two_arg((**C.text)(unsafe.Pointer(keys)), (**C.text)(unsafe.Pointer(values)), C.int(count))
	return C.GoString(C.text_to_cstring(_cret))
}


// JsonbCopy wraps MEOS C function jsonb_copy.
func JsonbCopy(jb *Jsonb) *Jsonb {
	_cret := C.jsonb_copy(jb._inner)
	return &Jsonb{_inner: _cret}
}


// JsonbMake wraps MEOS C function jsonb_make.
func JsonbMake(keys_vals unsafe.Pointer, count int) *Jsonb {
	_cret := C.jsonb_make((**C.text)(unsafe.Pointer(keys_vals)), C.int(count))
	return &Jsonb{_inner: _cret}
}


// JsonbMakeTwoArg wraps MEOS C function jsonb_make_two_arg.
func JsonbMakeTwoArg(keys unsafe.Pointer, values unsafe.Pointer, count int) *Jsonb {
	_cret := C.jsonb_make_two_arg((**C.text)(unsafe.Pointer(keys)), (**C.text)(unsafe.Pointer(values)), C.int(count))
	return &Jsonb{_inner: _cret}
}


// JsonbToBool wraps MEOS C function jsonb_to_bool.
func JsonbToBool(jb *Jsonb) bool {
	_cret := C.jsonb_to_bool(jb._inner)
	return bool(_cret)
}


// JsonbToCstring wraps MEOS C function jsonb_to_cstring.
func JsonbToCstring(jb *Jsonb) string {
	_cret := C.jsonb_to_cstring(jb._inner)
	return C.GoString(_cret)
}


// TODO jsonb_to_float4: unsupported return type float4
// func JsonbToFloat4(...) { /* not yet handled by codegen */ }


// TODO jsonb_to_float8: unsupported return type float8
// func JsonbToFloat8(...) { /* not yet handled by codegen */ }


// JsonbToInt16 wraps MEOS C function jsonb_to_int16.
func JsonbToInt16(jb *Jsonb) int16 {
	_cret := C.jsonb_to_int16(jb._inner)
	return int16(_cret)
}


// JsonbToInt32 wraps MEOS C function jsonb_to_int32.
func JsonbToInt32(jb *Jsonb) int32 {
	_cret := C.jsonb_to_int32(jb._inner)
	return int32(_cret)
}


// JsonbToInt64 wraps MEOS C function jsonb_to_int64.
func JsonbToInt64(jb *Jsonb) int64 {
	_cret := C.jsonb_to_int64(jb._inner)
	return int64(_cret)
}


// TODO jsonb_to_numeric: unsupported return type Numeric
// func JsonbToNumeric(...) { /* not yet handled by codegen */ }


// JsonbToText wraps MEOS C function jsonb_to_text.
func JsonbToText(jb *Jsonb) string {
	_cret := C.jsonb_to_text(jb._inner)
	return C.GoString(C.text_to_cstring(_cret))
}


// JSONArrayElement wraps MEOS C function json_array_element.
func JSONArrayElement(js string, element int) string {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	_cret := C.json_array_element(_c_js, C.int(element))
	return C.GoString(C.text_to_cstring(_cret))
}


// JSONArrayElementText wraps MEOS C function json_array_element_text.
func JSONArrayElementText(js string, element int) string {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	_cret := C.json_array_element_text(_c_js, C.int(element))
	return C.GoString(C.text_to_cstring(_cret))
}


// JSONArrayElements wraps MEOS C function json_array_elements.
func JSONArrayElements(js string, count unsafe.Pointer) unsafe.Pointer {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	_cret := C.json_array_elements(_c_js, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// JSONArrayElementsText wraps MEOS C function json_array_elements_text.
func JSONArrayElementsText(js string, count unsafe.Pointer) unsafe.Pointer {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	_cret := C.json_array_elements_text(_c_js, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// JSONArrayLength wraps MEOS C function json_array_length.
func JSONArrayLength(js string) int {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	_cret := C.json_array_length(_c_js)
	return int(_cret)
}


// JSONEach wraps MEOS C function json_each.
func JSONEach(js string, values unsafe.Pointer, count unsafe.Pointer) unsafe.Pointer {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	_cret := C.json_each(_c_js, (**C.text)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// JSONEachText wraps MEOS C function json_each_text.
func JSONEachText(js string, values unsafe.Pointer, count unsafe.Pointer) unsafe.Pointer {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	_cret := C.json_each_text(_c_js, (**C.text)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// JSONExtractPath wraps MEOS C function json_extract_path.
func JSONExtractPath(js string, path_elems unsafe.Pointer, path_len int) string {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	_cret := C.json_extract_path(_c_js, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len))
	return C.GoString(C.text_to_cstring(_cret))
}


// JSONExtractPathText wraps MEOS C function json_extract_path_text.
func JSONExtractPathText(js string, path_elems unsafe.Pointer, path_len int) string {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	_cret := C.json_extract_path_text(_c_js, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len))
	return C.GoString(C.text_to_cstring(_cret))
}


// JSONObjectField wraps MEOS C function json_object_field.
func JSONObjectField(js string, key string) string {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.json_object_field(_c_js, _c_key)
	return C.GoString(C.text_to_cstring(_cret))
}


// JSONObjectFieldText wraps MEOS C function json_object_field_text.
func JSONObjectFieldText(js string, key string) string {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.json_object_field_text(_c_js, _c_key)
	return C.GoString(C.text_to_cstring(_cret))
}


// JSONObjectKeys wraps MEOS C function json_object_keys.
func JSONObjectKeys(js string, count unsafe.Pointer) unsafe.Pointer {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	_cret := C.json_object_keys(_c_js, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// JSONTypeof wraps MEOS C function json_typeof.
func JSONTypeof(js string) string {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	_cret := C.json_typeof(_c_js)
	return C.GoString(C.text_to_cstring(_cret))
}


// JsonbArrayElement wraps MEOS C function jsonb_array_element.
func JsonbArrayElement(jb *Jsonb, element int) *Jsonb {
	_cret := C.jsonb_array_element(jb._inner, C.int(element))
	return &Jsonb{_inner: _cret}
}


// JsonbArrayElementText wraps MEOS C function jsonb_array_element_text.
func JsonbArrayElementText(jb *Jsonb, element int) string {
	_cret := C.jsonb_array_element_text(jb._inner, C.int(element))
	return C.GoString(C.text_to_cstring(_cret))
}


// JsonbArrayElements wraps MEOS C function jsonb_array_elements.
func JsonbArrayElements(jb *Jsonb, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.jsonb_array_elements(jb._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// JsonbArrayElementsText wraps MEOS C function jsonb_array_elements_text.
func JsonbArrayElementsText(jb *Jsonb, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.jsonb_array_elements_text(jb._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// JsonbArrayLength wraps MEOS C function jsonb_array_length.
func JsonbArrayLength(jb *Jsonb) int {
	_cret := C.jsonb_array_length(jb._inner)
	return int(_cret)
}


// JsonbContained wraps MEOS C function jsonb_contained.
func JsonbContained(jb1 *Jsonb, jb2 *Jsonb) bool {
	_cret := C.jsonb_contained(jb1._inner, jb2._inner)
	return bool(_cret)
}


// JsonbContains wraps MEOS C function jsonb_contains.
func JsonbContains(jb1 *Jsonb, jb2 *Jsonb) bool {
	_cret := C.jsonb_contains(jb1._inner, jb2._inner)
	return bool(_cret)
}


// JsonbEach wraps MEOS C function jsonb_each.
func JsonbEach(jb *Jsonb, values unsafe.Pointer, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.jsonb_each(jb._inner, (**C.Jsonb)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// JsonbEachText wraps MEOS C function jsonb_each_text.
func JsonbEachText(jb *Jsonb, values unsafe.Pointer, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.jsonb_each_text(jb._inner, (**C.text)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// JsonbExists wraps MEOS C function jsonb_exists.
func JsonbExists(jb *Jsonb, key string) bool {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.jsonb_exists(jb._inner, _c_key)
	return bool(_cret)
}


// JsonbExistsArray wraps MEOS C function jsonb_exists_array.
func JsonbExistsArray(jb *Jsonb, keys_elems unsafe.Pointer, keys_len int, any bool) bool {
	_cret := C.jsonb_exists_array(jb._inner, (**C.text)(unsafe.Pointer(keys_elems)), C.int(keys_len), C.bool(any))
	return bool(_cret)
}


// JsonbExtractPath wraps MEOS C function jsonb_extract_path.
func JsonbExtractPath(jb *Jsonb, path_elems unsafe.Pointer, path_len int) *Jsonb {
	_cret := C.jsonb_extract_path(jb._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len))
	return &Jsonb{_inner: _cret}
}


// JsonbExtractPathText wraps MEOS C function jsonb_extract_path_text.
func JsonbExtractPathText(jb *Jsonb, path_elems unsafe.Pointer, path_len int) string {
	_cret := C.jsonb_extract_path_text(jb._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len))
	return C.GoString(C.text_to_cstring(_cret))
}


// JsonbHash wraps MEOS C function jsonb_hash.
func JsonbHash(jb *Jsonb) uint32 {
	_cret := C.jsonb_hash(jb._inner)
	return uint32(_cret)
}


// JsonbHashExtended wraps MEOS C function jsonb_hash_extended.
func JsonbHashExtended(jb *Jsonb, seed uint64) uint64 {
	_cret := C.jsonb_hash_extended(jb._inner, C.uint64_t(seed))
	return uint64(_cret)
}


// JsonbObjectField wraps MEOS C function jsonb_object_field.
func JsonbObjectField(jb *Jsonb, key string) *Jsonb {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.jsonb_object_field(jb._inner, _c_key)
	return &Jsonb{_inner: _cret}
}


// JsonbObjectFieldText wraps MEOS C function jsonb_object_field_text.
func JsonbObjectFieldText(jb *Jsonb, key string) string {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.jsonb_object_field_text(jb._inner, _c_key)
	return C.GoString(C.text_to_cstring(_cret))
}


// JsonbObjectKeys wraps MEOS C function jsonb_object_keys.
func JsonbObjectKeys(jb *Jsonb, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.jsonb_object_keys(jb._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// JSONStripNulls wraps MEOS C function json_strip_nulls.
func JSONStripNulls(js string, strip_in_arrays bool) string {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	_cret := C.json_strip_nulls(_c_js, C.bool(strip_in_arrays))
	return C.GoString(C.text_to_cstring(_cret))
}


// JsonbConcat wraps MEOS C function jsonb_concat.
func JsonbConcat(jb1 *Jsonb, jb2 *Jsonb) *Jsonb {
	_cret := C.jsonb_concat(jb1._inner, jb2._inner)
	return &Jsonb{_inner: _cret}
}


// JsonbDelete wraps MEOS C function jsonb_delete.
func JsonbDelete(jb *Jsonb, key string) *Jsonb {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.jsonb_delete(jb._inner, _c_key)
	return &Jsonb{_inner: _cret}
}


// JsonbDeleteArray wraps MEOS C function jsonb_delete_array.
func JsonbDeleteArray(jb *Jsonb, keys_elems unsafe.Pointer, keys_len int) *Jsonb {
	_cret := C.jsonb_delete_array(jb._inner, (**C.text)(unsafe.Pointer(keys_elems)), C.int(keys_len))
	return &Jsonb{_inner: _cret}
}


// JsonbDeleteIndex wraps MEOS C function jsonb_delete_index.
func JsonbDeleteIndex(jb *Jsonb, idx int) *Jsonb {
	_cret := C.jsonb_delete_index(jb._inner, C.int(idx))
	return &Jsonb{_inner: _cret}
}


// JsonbDeletePath wraps MEOS C function jsonb_delete_path.
func JsonbDeletePath(jb *Jsonb, path_elems unsafe.Pointer, path_len int) *Jsonb {
	_cret := C.jsonb_delete_path(jb._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len))
	return &Jsonb{_inner: _cret}
}


// JsonbInsert wraps MEOS C function jsonb_insert.
func JsonbInsert(jb *Jsonb, path_elems unsafe.Pointer, path_len int, newjb *Jsonb, after bool) *Jsonb {
	_cret := C.jsonb_insert(jb._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len), newjb._inner, C.bool(after))
	return &Jsonb{_inner: _cret}
}


// JsonbPretty wraps MEOS C function jsonb_pretty.
func JsonbPretty(jb *Jsonb) string {
	_cret := C.jsonb_pretty(jb._inner)
	return C.GoString(C.text_to_cstring(_cret))
}


// JsonbSet wraps MEOS C function jsonb_set.
func JsonbSet(jb *Jsonb, path_elems unsafe.Pointer, path_len int, newjb *Jsonb, create bool) *Jsonb {
	_cret := C.jsonb_set(jb._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len), newjb._inner, C.bool(create))
	return &Jsonb{_inner: _cret}
}


// JsonbSetLax wraps MEOS C function jsonb_set_lax.
func JsonbSetLax(jb *Jsonb, path_elems unsafe.Pointer, path_len int, newjb *Jsonb, create bool, handle_null string) *Jsonb {
	_c_handle_null := C.cstring_to_text(C.CString(handle_null))
	defer C.free(unsafe.Pointer(_c_handle_null))
	_cret := C.jsonb_set_lax(jb._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len), newjb._inner, C.bool(create), _c_handle_null)
	return &Jsonb{_inner: _cret}
}


// JsonbStripNulls wraps MEOS C function jsonb_strip_nulls.
func JsonbStripNulls(jb *Jsonb, strip_in_arrays bool) *Jsonb {
	_cret := C.jsonb_strip_nulls(jb._inner, C.bool(strip_in_arrays))
	return &Jsonb{_inner: _cret}
}


// JsonbCmp wraps MEOS C function jsonb_cmp.
func JsonbCmp(jb1 *Jsonb, jb2 *Jsonb) int {
	_cret := C.jsonb_cmp(jb1._inner, jb2._inner)
	return int(_cret)
}


// JsonbEq wraps MEOS C function jsonb_eq.
func JsonbEq(jb1 *Jsonb, jb2 *Jsonb) bool {
	_cret := C.jsonb_eq(jb1._inner, jb2._inner)
	return bool(_cret)
}


// JsonbGe wraps MEOS C function jsonb_ge.
func JsonbGe(jb1 *Jsonb, jb2 *Jsonb) bool {
	_cret := C.jsonb_ge(jb1._inner, jb2._inner)
	return bool(_cret)
}


// JsonbGt wraps MEOS C function jsonb_gt.
func JsonbGt(jb1 *Jsonb, jb2 *Jsonb) bool {
	_cret := C.jsonb_gt(jb1._inner, jb2._inner)
	return bool(_cret)
}


// JsonbLe wraps MEOS C function jsonb_le.
func JsonbLe(jb1 *Jsonb, jb2 *Jsonb) bool {
	_cret := C.jsonb_le(jb1._inner, jb2._inner)
	return bool(_cret)
}


// JsonbLt wraps MEOS C function jsonb_lt.
func JsonbLt(jb1 *Jsonb, jb2 *Jsonb) bool {
	_cret := C.jsonb_lt(jb1._inner, jb2._inner)
	return bool(_cret)
}


// JsonbNe wraps MEOS C function jsonb_ne.
func JsonbNe(jb1 *Jsonb, jb2 *Jsonb) bool {
	_cret := C.jsonb_ne(jb1._inner, jb2._inner)
	return bool(_cret)
}


// JsonbPathExists wraps MEOS C function jsonb_path_exists.
func JsonbPathExists(jb *Jsonb, jp *JsonPath, vars *Jsonb, silent bool, tz bool) int {
	_cret := C.jsonb_path_exists(jb._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return int(_cret)
}


// JsonbPathMatch wraps MEOS C function jsonb_path_match.
func JsonbPathMatch(jb *Jsonb, jp *JsonPath, vars *Jsonb, silent bool, tz bool) bool {
	_cret := C.jsonb_path_match(jb._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return bool(_cret)
}


// JsonbPathQueryAll wraps MEOS C function jsonb_path_query_all.
func JsonbPathQueryAll(jb *Jsonb, jp *JsonPath, vars *Jsonb, silent bool, tz bool, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.jsonb_path_query_all(jb._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz), (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// JsonbPathQueryArray wraps MEOS C function jsonb_path_query_array.
func JsonbPathQueryArray(jb *Jsonb, jp *JsonPath, vars *Jsonb, silent bool, tz bool) *Jsonb {
	_cret := C.jsonb_path_query_array(jb._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return &Jsonb{_inner: _cret}
}


// JsonbPathQueryFirst wraps MEOS C function jsonb_path_query_first.
func JsonbPathQueryFirst(jb *Jsonb, jp *JsonPath, vars *Jsonb, silent bool, tz bool) *Jsonb {
	_cret := C.jsonb_path_query_first(jb._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return &Jsonb{_inner: _cret}
}


// JsonpathIn wraps MEOS C function jsonpath_in.
func JsonpathIn(str string) *JsonPath {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.jsonpath_in(_c_str)
	return &JsonPath{_inner: _cret}
}


// JsonpathCopy wraps MEOS C function jsonpath_copy.
func JsonpathCopy(jp *JsonPath) *JsonPath {
	_cret := C.jsonpath_copy(jp._inner)
	return &JsonPath{_inner: _cret}
}


// JsonpathOut wraps MEOS C function jsonpath_out.
func JsonpathOut(jp *JsonPath) string {
	_cret := C.jsonpath_out(jp._inner)
	return C.GoString(_cret)
}


// JsonbsetIn wraps MEOS C function jsonbset_in.
func JsonbsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.jsonbset_in(_c_str)
	return &Set{_inner: _cret}
}


// JsonbsetOut wraps MEOS C function jsonbset_out.
func JsonbsetOut(s *Set, maxdd int) string {
	_cret := C.jsonbset_out(s._inner, C.int(maxdd))
	return C.GoString(_cret)
}


// JsonbsetMake wraps MEOS C function jsonbset_make.
func JsonbsetMake(values unsafe.Pointer, count int) *Set {
	_cret := C.jsonbset_make((**C.Jsonb)(unsafe.Pointer(values)), C.int(count))
	return &Set{_inner: _cret}
}


// JsonbToSet wraps MEOS C function jsonb_to_set.
func JsonbToSet(jb *Jsonb) *Set {
	_cret := C.jsonb_to_set(jb._inner)
	return &Set{_inner: _cret}
}


// JsonbsetEndValue wraps MEOS C function jsonbset_end_value.
func JsonbsetEndValue(s *Set) *Jsonb {
	_cret := C.jsonbset_end_value(s._inner)
	return &Jsonb{_inner: _cret}
}


// JsonbsetStartValue wraps MEOS C function jsonbset_start_value.
func JsonbsetStartValue(s *Set) *Jsonb {
	_cret := C.jsonbset_start_value(s._inner)
	return &Jsonb{_inner: _cret}
}


// JsonbsetValueN wraps MEOS C function jsonbset_value_n.
func JsonbsetValueN(s *Set, n int) (bool, *Jsonb) {
	var _out_result *C.Jsonb
	_cret := C.jsonbset_value_n(s._inner, C.int(n), &_out_result)
	return bool(_cret), &Jsonb{_inner: _out_result}
}


// JsonbsetValues wraps MEOS C function jsonbset_values.
func JsonbsetValues(s *Set, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.jsonbset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// ConcatJsonbsetJsonb wraps MEOS C function concat_jsonbset_jsonb.
func ConcatJsonbsetJsonb(s *Set, jb *Jsonb, invert bool) *Set {
	_cret := C.concat_jsonbset_jsonb(s._inner, jb._inner, C.bool(invert))
	return &Set{_inner: _cret}
}


// JsonbsetArrayLength wraps MEOS C function jsonbset_array_length.
func JsonbsetArrayLength(set *Set) *Set {
	_cret := C.jsonbset_array_length(set._inner)
	return &Set{_inner: _cret}
}


// JsonbsetObjectField wraps MEOS C function jsonbset_object_field.
func JsonbsetObjectField(set *Set, key string, astext bool, null_handle NullHandleType) *Set {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.jsonbset_object_field(set._inner, _c_key, C.bool(astext), C.nullHandleType(null_handle))
	return &Set{_inner: _cret}
}


// JsonbsetArrayElement wraps MEOS C function jsonbset_array_element.
func JsonbsetArrayElement(set *Set, idx int, astext bool, null_handle NullHandleType) *Set {
	_cret := C.jsonbset_array_element(set._inner, C.int(idx), C.bool(astext), C.nullHandleType(null_handle))
	return &Set{_inner: _cret}
}


// JsonbsetDeleteIndex wraps MEOS C function jsonbset_delete_index.
func JsonbsetDeleteIndex(set *Set, idx int) *Set {
	_cret := C.jsonbset_delete_index(set._inner, C.int(idx))
	return &Set{_inner: _cret}
}


// JsonbsetDelete wraps MEOS C function jsonbset_delete.
func JsonbsetDelete(set *Set, key string) *Set {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.jsonbset_delete(set._inner, _c_key)
	return &Set{_inner: _cret}
}


// JsonbsetDeleteArray wraps MEOS C function jsonbset_delete_array.
func JsonbsetDeleteArray(set *Set, keys unsafe.Pointer, count int) *Set {
	_cret := C.jsonbset_delete_array(set._inner, (**C.text)(unsafe.Pointer(keys)), C.int(count))
	return &Set{_inner: _cret}
}


// JsonbsetExists wraps MEOS C function jsonbset_exists.
func JsonbsetExists(set *Set, key string) *Set {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.jsonbset_exists(set._inner, _c_key)
	return &Set{_inner: _cret}
}


// JsonbsetExistsArray wraps MEOS C function jsonbset_exists_array.
func JsonbsetExistsArray(set *Set, keys unsafe.Pointer, count int, any bool) *Set {
	_cret := C.jsonbset_exists_array(set._inner, (**C.text)(unsafe.Pointer(keys)), C.int(count), C.bool(any))
	return &Set{_inner: _cret}
}


// JsonbsetSet wraps MEOS C function jsonbset_set.
func JsonbsetSet(set *Set, keys unsafe.Pointer, count int, newjb *Jsonb, create bool, null_handle string, lax bool) *Set {
	_c_null_handle := C.cstring_to_text(C.CString(null_handle))
	defer C.free(unsafe.Pointer(_c_null_handle))
	_cret := C.jsonbset_set(set._inner, (**C.text)(unsafe.Pointer(keys)), C.int(count), newjb._inner, C.bool(create), _c_null_handle, C.bool(lax))
	return &Set{_inner: _cret}
}


// JsonbsetToAlphanumset wraps MEOS C function jsonbset_to_alphanumset.
func JsonbsetToAlphanumset(set *Set, key string, settype MeosType, null_handle NullHandleType) *Set {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.jsonbset_to_alphanumset(set._inner, _c_key, C.MeosType(settype), C.nullHandleType(null_handle))
	return &Set{_inner: _cret}
}


// JsonbsetToIntset wraps MEOS C function jsonbset_to_intset.
func JsonbsetToIntset(set *Set, key string, null_handle NullHandleType) *Set {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.jsonbset_to_intset(set._inner, _c_key, C.nullHandleType(null_handle))
	return &Set{_inner: _cret}
}


// JsonbsetToFloatset wraps MEOS C function jsonbset_to_floatset.
func JsonbsetToFloatset(set *Set, key string, null_handle NullHandleType) *Set {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.jsonbset_to_floatset(set._inner, _c_key, C.nullHandleType(null_handle))
	return &Set{_inner: _cret}
}


// JsonbsetToTextsetKey wraps MEOS C function jsonbset_to_textset_key.
func JsonbsetToTextsetKey(set *Set, key string, null_handle NullHandleType) *Set {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.jsonbset_to_textset_key(set._inner, _c_key, C.nullHandleType(null_handle))
	return &Set{_inner: _cret}
}


// JsonbsetStripNulls wraps MEOS C function jsonbset_strip_nulls.
func JsonbsetStripNulls(set *Set, strip_in_arrays bool) *Set {
	_cret := C.jsonbset_strip_nulls(set._inner, C.bool(strip_in_arrays))
	return &Set{_inner: _cret}
}


// JsonbsetPretty wraps MEOS C function jsonbset_pretty.
func JsonbsetPretty(set *Set) *Set {
	_cret := C.jsonbset_pretty(set._inner)
	return &Set{_inner: _cret}
}


// JsonbsetDeletePath wraps MEOS C function jsonbset_delete_path.
func JsonbsetDeletePath(set *Set, path_elems unsafe.Pointer, path_len int) *Set {
	_cret := C.jsonbset_delete_path(set._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len))
	return &Set{_inner: _cret}
}


// JsonbsetExtractPath wraps MEOS C function jsonbset_extract_path.
func JsonbsetExtractPath(set *Set, path_elems unsafe.Pointer, path_len int, astext bool, null_handle NullHandleType) *Set {
	_cret := C.jsonbset_extract_path(set._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len), C.bool(astext), C.nullHandleType(null_handle))
	return &Set{_inner: _cret}
}


// JsonbsetInsert wraps MEOS C function jsonbset_insert.
func JsonbsetInsert(set *Set, path_elems unsafe.Pointer, path_len int, newjb *Jsonb, after bool) *Set {
	_cret := C.jsonbset_insert(set._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len), newjb._inner, C.bool(after))
	return &Set{_inner: _cret}
}


// JsonbsetPathExists wraps MEOS C function jsonbset_path_exists.
func JsonbsetPathExists(set *Set, jp *JsonPath, vars *Jsonb, silent bool, tz bool) *Set {
	_cret := C.jsonbset_path_exists(set._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return &Set{_inner: _cret}
}


// JsonbsetPathMatch wraps MEOS C function jsonbset_path_match.
func JsonbsetPathMatch(set *Set, jp *JsonPath, vars *Jsonb, silent bool, tz bool) *Set {
	_cret := C.jsonbset_path_match(set._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return &Set{_inner: _cret}
}


// JsonbsetPathQueryArray wraps MEOS C function jsonbset_path_query_array.
func JsonbsetPathQueryArray(set *Set, jp *JsonPath, vars *Jsonb, silent bool, tz bool) *Set {
	_cret := C.jsonbset_path_query_array(set._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return &Set{_inner: _cret}
}


// JsonbsetPathQueryFirst wraps MEOS C function jsonbset_path_query_first.
func JsonbsetPathQueryFirst(set *Set, jp *JsonPath, vars *Jsonb, silent bool, tz bool) *Set {
	_cret := C.jsonbset_path_query_first(set._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return &Set{_inner: _cret}
}


// ContainedJsonbSet wraps MEOS C function contained_jsonb_set.
func ContainedJsonbSet(jb *Jsonb, s *Set) bool {
	_cret := C.contained_jsonb_set(jb._inner, s._inner)
	return bool(_cret)
}


// ContainsSetJsonb wraps MEOS C function contains_set_jsonb.
func ContainsSetJsonb(s *Set, jb *Jsonb) bool {
	_cret := C.contains_set_jsonb(s._inner, jb._inner)
	return bool(_cret)
}


// IntersectionJsonbSet wraps MEOS C function intersection_jsonb_set.
func IntersectionJsonbSet(jb *Jsonb, s *Set) *Set {
	_cret := C.intersection_jsonb_set(jb._inner, s._inner)
	return &Set{_inner: _cret}
}


// IntersectionSetJsonb wraps MEOS C function intersection_set_jsonb.
func IntersectionSetJsonb(s *Set, jb *Jsonb) *Set {
	_cret := C.intersection_set_jsonb(s._inner, jb._inner)
	return &Set{_inner: _cret}
}


// JsonbUnionTransfn wraps MEOS C function jsonb_union_transfn.
func JsonbUnionTransfn(state *Set, jb *Jsonb) *Set {
	_cret := C.jsonb_union_transfn(state._inner, jb._inner)
	return &Set{_inner: _cret}
}


// MinusJsonbSet wraps MEOS C function minus_jsonb_set.
func MinusJsonbSet(jb *Jsonb, s *Set) *Set {
	_cret := C.minus_jsonb_set(jb._inner, s._inner)
	return &Set{_inner: _cret}
}


// MinusSetJsonb wraps MEOS C function minus_set_jsonb.
func MinusSetJsonb(s *Set, jb *Jsonb) *Set {
	_cret := C.minus_set_jsonb(s._inner, jb._inner)
	return &Set{_inner: _cret}
}


// UnionJsonbSet wraps MEOS C function union_jsonb_set.
func UnionJsonbSet(jb *Jsonb, s *Set) *Set {
	_cret := C.gunion_jsonb_set(jb._inner, s._inner)
	return &Set{_inner: _cret}
}


// UnionSetJsonb wraps MEOS C function union_set_jsonb.
func UnionSetJsonb(s *Set, jb *Jsonb) *Set {
	_cret := C.gunion_set_jsonb(s._inner, jb._inner)
	return &Set{_inner: _cret}
}


// TjsonbFromMFJSON wraps MEOS C function tjsonb_from_mfjson.
func TjsonbFromMFJSON(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tjsonb_from_mfjson(_c_str)
	return &Temporal{_inner: _cret}
}


// TjsonbIn wraps MEOS C function tjsonb_in.
func TjsonbIn(str string) *Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tjsonb_in(_c_str)
	return &Temporal{_inner: _cret}
}


// TjsonbOut wraps MEOS C function tjsonb_out.
func TjsonbOut(temp *Temporal) string {
	_cret := C.tjsonb_out(temp._inner)
	return C.GoString(_cret)
}


// TjsonbinstIn wraps MEOS C function tjsonbinst_in.
func TjsonbinstIn(str string) *TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tjsonbinst_in(_c_str)
	return &TInstant{_inner: _cret}
}


// TjsonbseqIn wraps MEOS C function tjsonbseq_in.
func TjsonbseqIn(str string, interp Interpolation) *TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tjsonbseq_in(_c_str, C.interpType(interp))
	return &TSequence{_inner: _cret}
}


// TjsonbseqsetIn wraps MEOS C function tjsonbseqset_in.
func TjsonbseqsetIn(str string) *TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.tjsonbseqset_in(_c_str)
	return &TSequenceSet{_inner: _cret}
}


// TjsonbFromBaseTemp wraps MEOS C function tjsonb_from_base_temp.
func TjsonbFromBaseTemp(jsonb *Jsonb, temp *Temporal) *Temporal {
	_cret := C.tjsonb_from_base_temp(jsonb._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TjsonbinstMake wraps MEOS C function tjsonbinst_make.
func TjsonbinstMake(jsonb *Jsonb, t int64) *TInstant {
	_cret := C.tjsonbinst_make(jsonb._inner, C.TimestampTz(t))
	return &TInstant{_inner: _cret}
}


// TjsonbseqFromBaseTstzset wraps MEOS C function tjsonbseq_from_base_tstzset.
func TjsonbseqFromBaseTstzset(jsonb *Jsonb, s *Set) *TSequence {
	_cret := C.tjsonbseq_from_base_tstzset(jsonb._inner, s._inner)
	return &TSequence{_inner: _cret}
}


// TjsonbseqFromBaseTstzspan wraps MEOS C function tjsonbseq_from_base_tstzspan.
func TjsonbseqFromBaseTstzspan(jsonb *Jsonb, sp *Span) *TSequence {
	_cret := C.tjsonbseq_from_base_tstzspan(jsonb._inner, sp._inner)
	return &TSequence{_inner: _cret}
}


// TjsonbseqsetFromBaseTstzspanset wraps MEOS C function tjsonbseqset_from_base_tstzspanset.
func TjsonbseqsetFromBaseTstzspanset(jsonb *Jsonb, ss *SpanSet) *TSequenceSet {
	_cret := C.tjsonbseqset_from_base_tstzspanset(jsonb._inner, ss._inner)
	return &TSequenceSet{_inner: _cret}
}


// TjsonbToTtext wraps MEOS C function tjsonb_to_ttext.
func TjsonbToTtext(temp *Temporal) *Temporal {
	_cret := C.tjsonb_to_ttext(temp._inner)
	return &Temporal{_inner: _cret}
}


// TtextToTjsonb wraps MEOS C function ttext_to_tjsonb.
func TtextToTjsonb(temp *Temporal) *Temporal {
	_cret := C.ttext_to_tjsonb(temp._inner)
	return &Temporal{_inner: _cret}
}


// TjsonbEndValue wraps MEOS C function tjsonb_end_value.
func TjsonbEndValue(temp *Temporal) *Jsonb {
	_cret := C.tjsonb_end_value(temp._inner)
	return &Jsonb{_inner: _cret}
}


// TjsonbStartValue wraps MEOS C function tjsonb_start_value.
func TjsonbStartValue(temp *Temporal) *Jsonb {
	_cret := C.tjsonb_start_value(temp._inner)
	return &Jsonb{_inner: _cret}
}


// TjsonbValueAtTimestamptz wraps MEOS C function tjsonb_value_at_timestamptz.
func TjsonbValueAtTimestamptz(temp *Temporal, t int64, strict bool) (bool, *Jsonb) {
	var _out_value *C.Jsonb
	_cret := C.tjsonb_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_value)
	return bool(_cret), &Jsonb{_inner: _out_value}
}


// TjsonbValueN wraps MEOS C function tjsonb_value_n.
func TjsonbValueN(temp *Temporal, n int) (bool, *Jsonb) {
	var _out_result *C.Jsonb
	_cret := C.tjsonb_value_n(temp._inner, C.int(n), &_out_result)
	return bool(_cret), &Jsonb{_inner: _out_result}
}


// TjsonbValues wraps MEOS C function tjsonb_values.
func TjsonbValues(temp *Temporal, count unsafe.Pointer) unsafe.Pointer {
	_cret := C.tjsonb_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	return unsafe.Pointer(_cret)
}


// ConcatTjsonbJsonb wraps MEOS C function concat_tjsonb_jsonb.
func ConcatTjsonbJsonb(temp *Temporal, jb *Jsonb, invert bool) *Temporal {
	_cret := C.concat_tjsonb_jsonb(temp._inner, jb._inner, C.bool(invert))
	return &Temporal{_inner: _cret}
}


// ConcatTjsonbTjsonb wraps MEOS C function concat_tjsonb_tjsonb.
func ConcatTjsonbTjsonb(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.concat_tjsonb_tjsonb(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// ContainsTjsonbJsonb wraps MEOS C function contains_tjsonb_jsonb.
func ContainsTjsonbJsonb(temp *Temporal, jb *Jsonb, invert bool) *Temporal {
	_cret := C.contains_tjsonb_jsonb(temp._inner, jb._inner, C.bool(invert))
	return &Temporal{_inner: _cret}
}


// ContainsTjsonbTjsonb wraps MEOS C function contains_tjsonb_tjsonb.
func ContainsTjsonbTjsonb(temp1 *Temporal, temp2 *Temporal) *Temporal {
	_cret := C.contains_tjsonb_tjsonb(temp1._inner, temp2._inner)
	return &Temporal{_inner: _cret}
}


// NullHandleTypeFromString wraps MEOS C function null_handle_type_from_string.
func NullHandleTypeFromString(str string) NullHandleType {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	_cret := C.null_handle_type_from_string(_c_str)
	return NullHandleType(_cret)
}


// TjsonArrayElement wraps MEOS C function tjson_array_element.
func TjsonArrayElement(temp *Temporal, idx int, null_handle NullHandleType) *Temporal {
	_cret := C.tjson_array_element(temp._inner, C.int(idx), C.nullHandleType(null_handle))
	return &Temporal{_inner: _cret}
}


// TjsonArrayLength wraps MEOS C function tjson_array_length.
func TjsonArrayLength(temp *Temporal) *Temporal {
	_cret := C.tjson_array_length(temp._inner)
	return &Temporal{_inner: _cret}
}


// TjsonExtractPath wraps MEOS C function tjson_extract_path.
func TjsonExtractPath(temp *Temporal, path_elems unsafe.Pointer, path_len int, null_handle NullHandleType) *Temporal {
	_cret := C.tjson_extract_path(temp._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len), C.nullHandleType(null_handle))
	return &Temporal{_inner: _cret}
}


// TjsonObjectField wraps MEOS C function tjson_object_field.
func TjsonObjectField(temp *Temporal, key string, astext bool, null_handle NullHandleType) *Temporal {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.tjson_object_field(temp._inner, _c_key, C.bool(astext), C.nullHandleType(null_handle))
	return &Temporal{_inner: _cret}
}


// TjsonStripNulls wraps MEOS C function tjson_strip_nulls.
func TjsonStripNulls(temp *Temporal, strip_in_arrays bool) *Temporal {
	_cret := C.tjson_strip_nulls(temp._inner, C.bool(strip_in_arrays))
	return &Temporal{_inner: _cret}
}


// TjsonbArrayElement wraps MEOS C function tjsonb_array_element.
func TjsonbArrayElement(temp *Temporal, idx int, astext bool, null_handle NullHandleType) *Temporal {
	_cret := C.tjsonb_array_element(temp._inner, C.int(idx), C.bool(astext), C.nullHandleType(null_handle))
	return &Temporal{_inner: _cret}
}


// TjsonbArrayLength wraps MEOS C function tjsonb_array_length.
func TjsonbArrayLength(temp *Temporal) *Temporal {
	_cret := C.tjsonb_array_length(temp._inner)
	return &Temporal{_inner: _cret}
}


// TjsonbDelete wraps MEOS C function tjsonb_delete.
func TjsonbDelete(temp *Temporal, key string) *Temporal {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.tjsonb_delete(temp._inner, _c_key)
	return &Temporal{_inner: _cret}
}


// TjsonbDeleteArray wraps MEOS C function tjsonb_delete_array.
func TjsonbDeleteArray(temp *Temporal, keys unsafe.Pointer, count int) *Temporal {
	_cret := C.tjsonb_delete_array(temp._inner, (**C.text)(unsafe.Pointer(keys)), C.int(count))
	return &Temporal{_inner: _cret}
}


// TjsonbDeleteIndex wraps MEOS C function tjsonb_delete_index.
func TjsonbDeleteIndex(temp *Temporal, idx int) *Temporal {
	_cret := C.tjsonb_delete_index(temp._inner, C.int(idx))
	return &Temporal{_inner: _cret}
}


// TjsonbDeletePath wraps MEOS C function tjsonb_delete_path.
func TjsonbDeletePath(temp *Temporal, path_elems unsafe.Pointer, path_len int) *Temporal {
	_cret := C.tjsonb_delete_path(temp._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len))
	return &Temporal{_inner: _cret}
}


// TjsonbExists wraps MEOS C function tjsonb_exists.
func TjsonbExists(temp *Temporal, key string) *Temporal {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.tjsonb_exists(temp._inner, _c_key)
	return &Temporal{_inner: _cret}
}


// TjsonbExistsArray wraps MEOS C function tjsonb_exists_array.
func TjsonbExistsArray(temp *Temporal, keys unsafe.Pointer, count int, any bool) *Temporal {
	_cret := C.tjsonb_exists_array(temp._inner, (**C.text)(unsafe.Pointer(keys)), C.int(count), C.bool(any))
	return &Temporal{_inner: _cret}
}


// TjsonbExtractPath wraps MEOS C function tjsonb_extract_path.
func TjsonbExtractPath(temp *Temporal, path_elems unsafe.Pointer, path_len int, astext bool, null_handle NullHandleType) *Temporal {
	_cret := C.tjsonb_extract_path(temp._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len), C.bool(astext), C.nullHandleType(null_handle))
	return &Temporal{_inner: _cret}
}


// TjsonbInsert wraps MEOS C function tjsonb_insert.
func TjsonbInsert(temp *Temporal, keys unsafe.Pointer, count int, newjb *Jsonb, after bool) *Temporal {
	_cret := C.tjsonb_insert(temp._inner, (**C.text)(unsafe.Pointer(keys)), C.int(count), newjb._inner, C.bool(after))
	return &Temporal{_inner: _cret}
}


// TjsonbObjectField wraps MEOS C function tjsonb_object_field.
func TjsonbObjectField(temp *Temporal, key string, astext bool, null_handle NullHandleType) *Temporal {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.tjsonb_object_field(temp._inner, _c_key, C.bool(astext), C.nullHandleType(null_handle))
	return &Temporal{_inner: _cret}
}


// TjsonbPathExists wraps MEOS C function tjsonb_path_exists.
func TjsonbPathExists(temp *Temporal, jp *JsonPath, vars *Jsonb, silent bool, tz bool) *Temporal {
	_cret := C.tjsonb_path_exists(temp._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return &Temporal{_inner: _cret}
}


// TjsonbPathMatch wraps MEOS C function tjsonb_path_match.
func TjsonbPathMatch(temp *Temporal, jp *JsonPath, vars *Jsonb, silent bool, tz bool) *Temporal {
	_cret := C.tjsonb_path_match(temp._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return &Temporal{_inner: _cret}
}


// TjsonbPathQueryArray wraps MEOS C function tjsonb_path_query_array.
func TjsonbPathQueryArray(temp *Temporal, jp *JsonPath, vars *Jsonb, silent bool, tz bool) *Temporal {
	_cret := C.tjsonb_path_query_array(temp._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return &Temporal{_inner: _cret}
}


// TjsonbPathQueryFirst wraps MEOS C function tjsonb_path_query_first.
func TjsonbPathQueryFirst(temp *Temporal, jp *JsonPath, vars *Jsonb, silent bool, tz bool) *Temporal {
	_cret := C.tjsonb_path_query_first(temp._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return &Temporal{_inner: _cret}
}


// TjsonbPretty wraps MEOS C function tjsonb_pretty.
func TjsonbPretty(temp *Temporal) *Temporal {
	_cret := C.tjsonb_pretty(temp._inner)
	return &Temporal{_inner: _cret}
}


// TjsonbSet wraps MEOS C function tjsonb_set.
func TjsonbSet(temp *Temporal, keys unsafe.Pointer, count int, newjb *Jsonb, create bool, handle_null string, lax bool) *Temporal {
	_c_handle_null := C.cstring_to_text(C.CString(handle_null))
	defer C.free(unsafe.Pointer(_c_handle_null))
	_cret := C.tjsonb_set(temp._inner, (**C.text)(unsafe.Pointer(keys)), C.int(count), newjb._inner, C.bool(create), _c_handle_null, C.bool(lax))
	return &Temporal{_inner: _cret}
}


// TjsonbStripNulls wraps MEOS C function tjsonb_strip_nulls.
func TjsonbStripNulls(temp *Temporal, strip_in_arrays bool) *Temporal {
	_cret := C.tjsonb_strip_nulls(temp._inner, C.bool(strip_in_arrays))
	return &Temporal{_inner: _cret}
}


// TjsonbToTbool wraps MEOS C function tjsonb_to_tbool.
func TjsonbToTbool(temp *Temporal, key string, null_handle NullHandleType) *Temporal {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.tjsonb_to_tbool(temp._inner, _c_key, C.nullHandleType(null_handle))
	return &Temporal{_inner: _cret}
}


// TjsonbToTfloat wraps MEOS C function tjsonb_to_tfloat.
func TjsonbToTfloat(temp *Temporal, key string, interp Interpolation, null_handle NullHandleType) *Temporal {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.tjsonb_to_tfloat(temp._inner, _c_key, C.interpType(interp), C.nullHandleType(null_handle))
	return &Temporal{_inner: _cret}
}


// TjsonbToTint wraps MEOS C function tjsonb_to_tint.
func TjsonbToTint(temp *Temporal, key string, null_handle NullHandleType) *Temporal {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.tjsonb_to_tint(temp._inner, _c_key, C.nullHandleType(null_handle))
	return &Temporal{_inner: _cret}
}


// TjsonbToTtextKey wraps MEOS C function tjsonb_to_ttext_key.
func TjsonbToTtextKey(temp *Temporal, key string, null_handle NullHandleType) *Temporal {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	_cret := C.tjsonb_to_ttext_key(temp._inner, _c_key, C.nullHandleType(null_handle))
	return &Temporal{_inner: _cret}
}


// TjsonbAtValue wraps MEOS C function tjsonb_at_value.
func TjsonbAtValue(temp *Temporal, jsb *Jsonb) *Temporal {
	_cret := C.tjsonb_at_value(temp._inner, jsb._inner)
	return &Temporal{_inner: _cret}
}


// TjsonbMinusValue wraps MEOS C function tjsonb_minus_value.
func TjsonbMinusValue(temp *Temporal, jsb *Jsonb) *Temporal {
	_cret := C.tjsonb_minus_value(temp._inner, jsb._inner)
	return &Temporal{_inner: _cret}
}


// AlwaysEqJsonbTjsonb wraps MEOS C function always_eq_jsonb_tjsonb.
func AlwaysEqJsonbTjsonb(jb *Jsonb, temp *Temporal) int {
	_cret := C.always_eq_jsonb_tjsonb(jb._inner, temp._inner)
	return int(_cret)
}


// AlwaysEqTjsonbJsonb wraps MEOS C function always_eq_tjsonb_jsonb.
func AlwaysEqTjsonbJsonb(temp *Temporal, jb *Jsonb) int {
	_cret := C.always_eq_tjsonb_jsonb(temp._inner, jb._inner)
	return int(_cret)
}


// AlwaysEqTjsonbTjsonb wraps MEOS C function always_eq_tjsonb_tjsonb.
func AlwaysEqTjsonbTjsonb(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_eq_tjsonb_tjsonb(temp1._inner, temp2._inner)
	return int(_cret)
}


// AlwaysNeJsonbTjsonb wraps MEOS C function always_ne_jsonb_tjsonb.
func AlwaysNeJsonbTjsonb(jb *Jsonb, temp *Temporal) int {
	_cret := C.always_ne_jsonb_tjsonb(jb._inner, temp._inner)
	return int(_cret)
}


// AlwaysNeTjsonbJsonb wraps MEOS C function always_ne_tjsonb_jsonb.
func AlwaysNeTjsonbJsonb(temp *Temporal, jb *Jsonb) int {
	_cret := C.always_ne_tjsonb_jsonb(temp._inner, jb._inner)
	return int(_cret)
}


// AlwaysNeTjsonbTjsonb wraps MEOS C function always_ne_tjsonb_tjsonb.
func AlwaysNeTjsonbTjsonb(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.always_ne_tjsonb_tjsonb(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverEqJsonbTjsonb wraps MEOS C function ever_eq_jsonb_tjsonb.
func EverEqJsonbTjsonb(jb *Jsonb, temp *Temporal) int {
	_cret := C.ever_eq_jsonb_tjsonb(jb._inner, temp._inner)
	return int(_cret)
}


// EverEqTjsonbJsonb wraps MEOS C function ever_eq_tjsonb_jsonb.
func EverEqTjsonbJsonb(temp *Temporal, jb *Jsonb) int {
	_cret := C.ever_eq_tjsonb_jsonb(temp._inner, jb._inner)
	return int(_cret)
}


// EverEqTjsonbTjsonb wraps MEOS C function ever_eq_tjsonb_tjsonb.
func EverEqTjsonbTjsonb(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_eq_tjsonb_tjsonb(temp1._inner, temp2._inner)
	return int(_cret)
}


// EverNeJsonbTjsonb wraps MEOS C function ever_ne_jsonb_tjsonb.
func EverNeJsonbTjsonb(jb *Jsonb, temp *Temporal) int {
	_cret := C.ever_ne_jsonb_tjsonb(jb._inner, temp._inner)
	return int(_cret)
}


// EverNeTjsonbJsonb wraps MEOS C function ever_ne_tjsonb_jsonb.
func EverNeTjsonbJsonb(temp *Temporal, jb *Jsonb) int {
	_cret := C.ever_ne_tjsonb_jsonb(temp._inner, jb._inner)
	return int(_cret)
}


// EverNeTjsonbTjsonb wraps MEOS C function ever_ne_tjsonb_tjsonb.
func EverNeTjsonbTjsonb(temp1 *Temporal, temp2 *Temporal) int {
	_cret := C.ever_ne_tjsonb_tjsonb(temp1._inner, temp2._inner)
	return int(_cret)
}


// TeqJsonbTjsonb wraps MEOS C function teq_jsonb_tjsonb.
func TeqJsonbTjsonb(jb *Jsonb, temp *Temporal) *Temporal {
	_cret := C.teq_jsonb_tjsonb(jb._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TeqTjsonbJsonb wraps MEOS C function teq_tjsonb_jsonb.
func TeqTjsonbJsonb(temp *Temporal, jb *Jsonb) *Temporal {
	_cret := C.teq_tjsonb_jsonb(temp._inner, jb._inner)
	return &Temporal{_inner: _cret}
}


// TneJsonbTjsonb wraps MEOS C function tne_jsonb_tjsonb.
func TneJsonbTjsonb(jb *Jsonb, temp *Temporal) *Temporal {
	_cret := C.tne_jsonb_tjsonb(jb._inner, temp._inner)
	return &Temporal{_inner: _cret}
}


// TneTjsonbJsonb wraps MEOS C function tne_tjsonb_jsonb.
func TneTjsonbJsonb(temp *Temporal, jb *Jsonb) *Temporal {
	_cret := C.tne_tjsonb_jsonb(temp._inner, jb._inner)
	return &Temporal{_inner: _cret}
}

