package generated

// #include <stddef.h>
import "C"
import (
	"unsafe"

	"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

// JSONIn wraps MEOS C function json_in.
func JSONIn(str string) string {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.json_in(_c_str)
	return text2cstring(res)
}


// JSONOut wraps MEOS C function json_out.
func JSONOut(js string) string {
	_c_js := cstring2text(js)
	defer C.free(unsafe.Pointer(_c_js))
	res := C.json_out(_c_js)
	return C.GoString(res)
}


// JsonbFromText wraps MEOS C function jsonb_from_text.
func JsonbFromText(txt string, unique_keys bool) *Jsonb {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.jsonb_from_text(_c_txt, C.bool(unique_keys))
	return &Jsonb{_inner: res}
}


// JsonbIn wraps MEOS C function jsonb_in.
func JsonbIn(str string) *Jsonb {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.jsonb_in(_c_str)
	return &Jsonb{_inner: res}
}


// JsonbOut wraps MEOS C function jsonb_out.
func JsonbOut(jb *Jsonb) string {
	res := C.jsonb_out(jb._inner)
	return C.GoString(res)
}


// JSONMake wraps MEOS C function json_make.
func JSONMake(keys_vals []string) string {
	_c_keys_vals := make([]*C.text, len(keys_vals))
	for _i, _v := range keys_vals { _c_keys_vals[_i] = cstring2text(_v) }
	res := C.json_make((**C.text)(unsafe.Pointer(&_c_keys_vals[0])), C.int(len(keys_vals)))
	return text2cstring(res)
}


// TODO json_make_two_arg: unsupported param text **
// func JSONMakeTwoArg(...) { /* not yet handled by codegen */ }


// JsonbCopy wraps MEOS C function jsonb_copy.
func JsonbCopy(jb *Jsonb) *Jsonb {
	res := C.jsonb_copy(jb._inner)
	return &Jsonb{_inner: res}
}


// JsonbMake wraps MEOS C function jsonb_make.
func JsonbMake(keys_vals []string) *Jsonb {
	_c_keys_vals := make([]*C.text, len(keys_vals))
	for _i, _v := range keys_vals { _c_keys_vals[_i] = cstring2text(_v) }
	res := C.jsonb_make((**C.text)(unsafe.Pointer(&_c_keys_vals[0])), C.int(len(keys_vals)))
	return &Jsonb{_inner: res}
}


// TODO jsonb_make_two_arg: unsupported param text **
// func JsonbMakeTwoArg(...) { /* not yet handled by codegen */ }


// JsonbToBool wraps MEOS C function jsonb_to_bool.
func JsonbToBool(jb *Jsonb) bool {
	res := C.jsonb_to_bool(jb._inner)
	return bool(res)
}


// JsonbToCstring wraps MEOS C function jsonb_to_cstring.
func JsonbToCstring(jb *Jsonb) string {
	res := C.jsonb_to_cstring(jb._inner)
	return C.GoString(res)
}


// JsonbToFloat4 wraps MEOS C function jsonb_to_float4.
func JsonbToFloat4(jb *Jsonb) int {
	res := C.jsonb_to_float4(jb._inner)
	return int(res)
}


// JsonbToFloat8 wraps MEOS C function jsonb_to_float8.
func JsonbToFloat8(jb *Jsonb) int {
	res := C.jsonb_to_float8(jb._inner)
	return int(res)
}


// JsonbToInt16 wraps MEOS C function jsonb_to_int16.
func JsonbToInt16(jb *Jsonb) int16 {
	res := C.jsonb_to_int16(jb._inner)
	return int16(res)
}


// JsonbToInt32 wraps MEOS C function jsonb_to_int32.
func JsonbToInt32(jb *Jsonb) int {
	res := C.jsonb_to_int32(jb._inner)
	return int(res)
}


// JsonbToInt64 wraps MEOS C function jsonb_to_int64.
func JsonbToInt64(jb *Jsonb) int64 {
	res := C.jsonb_to_int64(jb._inner)
	return int64(res)
}


// JsonbToNumeric wraps MEOS C function jsonb_to_numeric.
func JsonbToNumeric(jb *Jsonb) int {
	res := C.jsonb_to_numeric(jb._inner)
	return int(res)
}


// JsonbToText wraps MEOS C function jsonb_to_text.
func JsonbToText(jb *Jsonb) string {
	res := C.jsonb_to_text(jb._inner)
	return text2cstring(res)
}


// JSONArrayElement wraps MEOS C function json_array_element.
func JSONArrayElement(js string, element int) string {
	_c_js := cstring2text(js)
	defer C.free(unsafe.Pointer(_c_js))
	res := C.json_array_element(_c_js, C.int(element))
	return text2cstring(res)
}


// JSONArrayElementText wraps MEOS C function json_array_element_text.
func JSONArrayElementText(js string, element int) string {
	_c_js := cstring2text(js)
	defer C.free(unsafe.Pointer(_c_js))
	res := C.json_array_element_text(_c_js, C.int(element))
	return text2cstring(res)
}


// JSONArrayElements wraps MEOS C function json_array_elements.
func JSONArrayElements(js string) []string {
	_c_js := cstring2text(js)
	defer C.free(unsafe.Pointer(_c_js))
	var _out_count C.int
	res := C.json_array_elements(_c_js, &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.text)(unsafe.Pointer(res)), _n)
	_out := make([]string, _n)
	for _i, _e := range _slice {
		_out[_i] = text2cstring(_e)
	}
	return _out
}


// JSONArrayElementsText wraps MEOS C function json_array_elements_text.
func JSONArrayElementsText(js string) []string {
	_c_js := cstring2text(js)
	defer C.free(unsafe.Pointer(_c_js))
	var _out_count C.int
	res := C.json_array_elements_text(_c_js, &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.text)(unsafe.Pointer(res)), _n)
	_out := make([]string, _n)
	for _i, _e := range _slice {
		_out[_i] = text2cstring(_e)
	}
	return _out
}


// JSONArrayLength wraps MEOS C function json_array_length.
func JSONArrayLength(js string) int {
	_c_js := cstring2text(js)
	defer C.free(unsafe.Pointer(_c_js))
	res := C.json_array_length(_c_js)
	return int(res)
}


// TODO json_each: unsupported param text **
// func JSONEach(...) { /* not yet handled by codegen */ }


// TODO json_each_text: unsupported param text **
// func JSONEachText(...) { /* not yet handled by codegen */ }


// TODO json_extract_path: unsupported param text **
// func JSONExtractPath(...) { /* not yet handled by codegen */ }


// TODO json_extract_path_text: unsupported param text **
// func JSONExtractPathText(...) { /* not yet handled by codegen */ }


// JSONObjectField wraps MEOS C function json_object_field.
func JSONObjectField(js string, key string) string {
	_c_js := cstring2text(js)
	defer C.free(unsafe.Pointer(_c_js))
	_c_key := cstring2text(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.json_object_field(_c_js, _c_key)
	return text2cstring(res)
}


// JSONObjectFieldText wraps MEOS C function json_object_field_text.
func JSONObjectFieldText(js string, key string) string {
	_c_js := cstring2text(js)
	defer C.free(unsafe.Pointer(_c_js))
	_c_key := cstring2text(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.json_object_field_text(_c_js, _c_key)
	return text2cstring(res)
}


// JSONObjectKeys wraps MEOS C function json_object_keys.
func JSONObjectKeys(js string) []string {
	_c_js := cstring2text(js)
	defer C.free(unsafe.Pointer(_c_js))
	var _out_count C.int
	res := C.json_object_keys(_c_js, &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.text)(unsafe.Pointer(res)), _n)
	_out := make([]string, _n)
	for _i, _e := range _slice {
		_out[_i] = text2cstring(_e)
	}
	return _out
}


// JSONTypeof wraps MEOS C function json_typeof.
func JSONTypeof(js string) string {
	_c_js := cstring2text(js)
	defer C.free(unsafe.Pointer(_c_js))
	res := C.json_typeof(_c_js)
	return text2cstring(res)
}


// JsonbArrayElement wraps MEOS C function jsonb_array_element.
func JsonbArrayElement(jb *Jsonb, element int) *Jsonb {
	res := C.jsonb_array_element(jb._inner, C.int(element))
	return &Jsonb{_inner: res}
}


// JsonbArrayElementText wraps MEOS C function jsonb_array_element_text.
func JsonbArrayElementText(jb *Jsonb, element int) string {
	res := C.jsonb_array_element_text(jb._inner, C.int(element))
	return text2cstring(res)
}


// JsonbArrayElements wraps MEOS C function jsonb_array_elements.
func JsonbArrayElements(jb *Jsonb) []*Jsonb {
	var _out_count C.int
	res := C.jsonb_array_elements(jb._inner, &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.Jsonb)(unsafe.Pointer(res)), _n)
	_out := make([]*Jsonb, _n)
	for _i, _e := range _slice {
		_out[_i] = &Jsonb{_inner: _e}
	}
	return _out
}


// JsonbArrayElementsText wraps MEOS C function jsonb_array_elements_text.
func JsonbArrayElementsText(jb *Jsonb) []string {
	var _out_count C.int
	res := C.jsonb_array_elements_text(jb._inner, &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.text)(unsafe.Pointer(res)), _n)
	_out := make([]string, _n)
	for _i, _e := range _slice {
		_out[_i] = text2cstring(_e)
	}
	return _out
}


// JsonbArrayLength wraps MEOS C function jsonb_array_length.
func JsonbArrayLength(jb *Jsonb) int {
	res := C.jsonb_array_length(jb._inner)
	return int(res)
}


// JsonbContained wraps MEOS C function jsonb_contained.
func JsonbContained(jb1 *Jsonb, jb2 *Jsonb) bool {
	res := C.jsonb_contained(jb1._inner, jb2._inner)
	return bool(res)
}


// JsonbContains wraps MEOS C function jsonb_contains.
func JsonbContains(jb1 *Jsonb, jb2 *Jsonb) bool {
	res := C.jsonb_contains(jb1._inner, jb2._inner)
	return bool(res)
}


// TODO jsonb_each: unsupported param Jsonb **
// func JsonbEach(...) { /* not yet handled by codegen */ }


// TODO jsonb_each_text: unsupported param text **
// func JsonbEachText(...) { /* not yet handled by codegen */ }


// JsonbExists wraps MEOS C function jsonb_exists.
func JsonbExists(jb *Jsonb, key string) bool {
	_c_key := cstring2text(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.jsonb_exists(jb._inner, _c_key)
	return bool(res)
}


// TODO jsonb_exists_array: unsupported param text **
// func JsonbExistsArray(...) { /* not yet handled by codegen */ }


// TODO jsonb_extract_path: unsupported param text **
// func JsonbExtractPath(...) { /* not yet handled by codegen */ }


// TODO jsonb_extract_path_text: unsupported param text **
// func JsonbExtractPathText(...) { /* not yet handled by codegen */ }


// JsonbHash wraps MEOS C function jsonb_hash.
func JsonbHash(jb *Jsonb) int {
	res := C.jsonb_hash(jb._inner)
	return int(res)
}


// JsonbHashExtended wraps MEOS C function jsonb_hash_extended.
func JsonbHashExtended(jb *Jsonb, seed int) int {
	res := C.jsonb_hash_extended(jb._inner, C.int(seed))
	return int(res)
}


// JsonbObjectField wraps MEOS C function jsonb_object_field.
func JsonbObjectField(jb *Jsonb, key string) *Jsonb {
	_c_key := cstring2text(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.jsonb_object_field(jb._inner, _c_key)
	return &Jsonb{_inner: res}
}


// JsonbObjectFieldText wraps MEOS C function jsonb_object_field_text.
func JsonbObjectFieldText(jb *Jsonb, key string) string {
	_c_key := cstring2text(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.jsonb_object_field_text(jb._inner, _c_key)
	return text2cstring(res)
}


// JsonbObjectKeys wraps MEOS C function jsonb_object_keys.
func JsonbObjectKeys(jb *Jsonb) []string {
	var _out_count C.int
	res := C.jsonb_object_keys(jb._inner, &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.text)(unsafe.Pointer(res)), _n)
	_out := make([]string, _n)
	for _i, _e := range _slice {
		_out[_i] = text2cstring(_e)
	}
	return _out
}


// JSONStripNulls wraps MEOS C function json_strip_nulls.
func JSONStripNulls(js string, strip_in_arrays bool) string {
	_c_js := cstring2text(js)
	defer C.free(unsafe.Pointer(_c_js))
	res := C.json_strip_nulls(_c_js, C.bool(strip_in_arrays))
	return text2cstring(res)
}


// JsonbConcat wraps MEOS C function jsonb_concat.
func JsonbConcat(jb1 *Jsonb, jb2 *Jsonb) *Jsonb {
	res := C.jsonb_concat(jb1._inner, jb2._inner)
	return &Jsonb{_inner: res}
}


// JsonbDelete wraps MEOS C function jsonb_delete.
func JsonbDelete(jb *Jsonb, key string) *Jsonb {
	_c_key := cstring2text(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.jsonb_delete(jb._inner, _c_key)
	return &Jsonb{_inner: res}
}


// TODO jsonb_delete_array: unsupported param text **
// func JsonbDeleteArray(...) { /* not yet handled by codegen */ }


// JsonbDeleteIndex wraps MEOS C function jsonb_delete_index.
func JsonbDeleteIndex(jb *Jsonb, idx int) *Jsonb {
	res := C.jsonb_delete_index(jb._inner, C.int(idx))
	return &Jsonb{_inner: res}
}


// TODO jsonb_delete_path: unsupported param text **
// func JsonbDeletePath(...) { /* not yet handled by codegen */ }


// TODO jsonb_insert: unsupported param text **
// func JsonbInsert(...) { /* not yet handled by codegen */ }


// JsonbPretty wraps MEOS C function jsonb_pretty.
func JsonbPretty(jb *Jsonb) string {
	res := C.jsonb_pretty(jb._inner)
	return text2cstring(res)
}


// TODO jsonb_set: unsupported param text **
// func JsonbSet(...) { /* not yet handled by codegen */ }


// TODO jsonb_set_lax: unsupported param text **
// func JsonbSetLax(...) { /* not yet handled by codegen */ }


// JsonbStripNulls wraps MEOS C function jsonb_strip_nulls.
func JsonbStripNulls(jb *Jsonb, strip_in_arrays bool) *Jsonb {
	res := C.jsonb_strip_nulls(jb._inner, C.bool(strip_in_arrays))
	return &Jsonb{_inner: res}
}


// JsonbCmp wraps MEOS C function jsonb_cmp.
func JsonbCmp(jb1 *Jsonb, jb2 *Jsonb) int {
	res := C.jsonb_cmp(jb1._inner, jb2._inner)
	return int(res)
}


// JsonbEq wraps MEOS C function jsonb_eq.
func JsonbEq(jb1 *Jsonb, jb2 *Jsonb) bool {
	res := C.jsonb_eq(jb1._inner, jb2._inner)
	return bool(res)
}


// JsonbGe wraps MEOS C function jsonb_ge.
func JsonbGe(jb1 *Jsonb, jb2 *Jsonb) bool {
	res := C.jsonb_ge(jb1._inner, jb2._inner)
	return bool(res)
}


// JsonbGt wraps MEOS C function jsonb_gt.
func JsonbGt(jb1 *Jsonb, jb2 *Jsonb) bool {
	res := C.jsonb_gt(jb1._inner, jb2._inner)
	return bool(res)
}


// JsonbLe wraps MEOS C function jsonb_le.
func JsonbLe(jb1 *Jsonb, jb2 *Jsonb) bool {
	res := C.jsonb_le(jb1._inner, jb2._inner)
	return bool(res)
}


// JsonbLt wraps MEOS C function jsonb_lt.
func JsonbLt(jb1 *Jsonb, jb2 *Jsonb) bool {
	res := C.jsonb_lt(jb1._inner, jb2._inner)
	return bool(res)
}


// JsonbNe wraps MEOS C function jsonb_ne.
func JsonbNe(jb1 *Jsonb, jb2 *Jsonb) bool {
	res := C.jsonb_ne(jb1._inner, jb2._inner)
	return bool(res)
}


// JsonbPathExists wraps MEOS C function jsonb_path_exists.
func JsonbPathExists(jb *Jsonb, jp *JsonPath, vars *Jsonb, silent bool, tz bool) int {
	res := C.jsonb_path_exists(jb._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return int(res)
}


// JsonbPathMatch wraps MEOS C function jsonb_path_match.
func JsonbPathMatch(jb *Jsonb, jp *JsonPath, vars *Jsonb, silent bool, tz bool) bool {
	res := C.jsonb_path_match(jb._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return bool(res)
}


// JsonbPathQueryAll wraps MEOS C function jsonb_path_query_all.
func JsonbPathQueryAll(jb *Jsonb, jp *JsonPath, vars *Jsonb, silent bool, tz bool) []*Jsonb {
	var _out_count C.int
	res := C.jsonb_path_query_all(jb._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.Jsonb)(unsafe.Pointer(res)), _n)
	_out := make([]*Jsonb, _n)
	for _i, _e := range _slice {
		_out[_i] = &Jsonb{_inner: _e}
	}
	return _out
}


// JsonbPathQueryArray wraps MEOS C function jsonb_path_query_array.
func JsonbPathQueryArray(jb *Jsonb, jp *JsonPath, vars *Jsonb, silent bool, tz bool) *Jsonb {
	res := C.jsonb_path_query_array(jb._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return &Jsonb{_inner: res}
}


// JsonbPathQueryFirst wraps MEOS C function jsonb_path_query_first.
func JsonbPathQueryFirst(jb *Jsonb, jp *JsonPath, vars *Jsonb, silent bool, tz bool) *Jsonb {
	res := C.jsonb_path_query_first(jb._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return &Jsonb{_inner: res}
}


// JsonpathIn wraps MEOS C function jsonpath_in.
func JsonpathIn(str string) *JsonPath {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.jsonpath_in(_c_str)
	return &JsonPath{_inner: res}
}


// JsonpathCopy wraps MEOS C function jsonpath_copy.
func JsonpathCopy(jp *JsonPath) *JsonPath {
	res := C.jsonpath_copy(jp._inner)
	return &JsonPath{_inner: res}
}


// JsonpathOut wraps MEOS C function jsonpath_out.
func JsonpathOut(jp *JsonPath) string {
	res := C.jsonpath_out(jp._inner)
	return C.GoString(res)
}


// JsonbsetIn wraps MEOS C function jsonbset_in.
func JsonbsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.jsonbset_in(_c_str)
	return &Set{_inner: res}
}


// JsonbsetOut wraps MEOS C function jsonbset_out.
func JsonbsetOut(s *Set, maxdd int) string {
	res := C.jsonbset_out(s._inner, C.int(maxdd))
	return C.GoString(res)
}


// JsonbsetMake wraps MEOS C function jsonbset_make.
func JsonbsetMake(values []*Jsonb) *Set {
	_c_values := make([]*C.Jsonb, len(values))
	for _i, _v := range values { _c_values[_i] = _v._inner }
	res := C.jsonbset_make((**C.Jsonb)(unsafe.Pointer(&_c_values[0])), C.int(len(values)))
	return &Set{_inner: res}
}


// JsonbToSet wraps MEOS C function jsonb_to_set.
func JsonbToSet(jb *Jsonb) *Set {
	res := C.jsonb_to_set(jb._inner)
	return &Set{_inner: res}
}


// JsonbsetEndValue wraps MEOS C function jsonbset_end_value.
func JsonbsetEndValue(s *Set) *Jsonb {
	res := C.jsonbset_end_value(s._inner)
	return &Jsonb{_inner: res}
}


// JsonbsetStartValue wraps MEOS C function jsonbset_start_value.
func JsonbsetStartValue(s *Set) *Jsonb {
	res := C.jsonbset_start_value(s._inner)
	return &Jsonb{_inner: res}
}


// JsonbsetValueN wraps MEOS C function jsonbset_value_n.
func JsonbsetValueN(s *Set, n int) (bool, *Jsonb) {
	var _out_result *C.Jsonb
	res := C.jsonbset_value_n(s._inner, C.int(n), &_out_result)
	return bool(res), &Jsonb{_inner: _out_result}
}


// TODO jsonbset_values: unsupported return type Jsonb **
// func JsonbsetValues(...) { /* not yet handled by codegen */ }


// ConcatJsonbsetJsonb wraps MEOS C function concat_jsonbset_jsonb.
func ConcatJsonbsetJsonb(s *Set, jb *Jsonb, invert bool) *Set {
	res := C.concat_jsonbset_jsonb(s._inner, jb._inner, C.bool(invert))
	return &Set{_inner: res}
}


// JsonbsetArrayLength wraps MEOS C function jsonbset_array_length.
func JsonbsetArrayLength(set *Set) *Set {
	res := C.jsonbset_array_length(set._inner)
	return &Set{_inner: res}
}


// JsonbsetObjectField wraps MEOS C function jsonbset_object_field.
func JsonbsetObjectField(set *Set, key string, astext bool, null_handle NullHandleType) *Set {
	_c_key := cstring2text(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.jsonbset_object_field(set._inner, _c_key, C.bool(astext), C.nullHandleType(null_handle))
	return &Set{_inner: res}
}


// JsonbsetArrayElement wraps MEOS C function jsonbset_array_element.
func JsonbsetArrayElement(set *Set, idx int, astext bool, null_handle NullHandleType) *Set {
	res := C.jsonbset_array_element(set._inner, C.int(idx), C.bool(astext), C.nullHandleType(null_handle))
	return &Set{_inner: res}
}


// JsonbsetDeleteIndex wraps MEOS C function jsonbset_delete_index.
func JsonbsetDeleteIndex(set *Set, idx int) *Set {
	res := C.jsonbset_delete_index(set._inner, C.int(idx))
	return &Set{_inner: res}
}


// JsonbsetDelete wraps MEOS C function jsonbset_delete.
func JsonbsetDelete(set *Set, key string) *Set {
	_c_key := cstring2text(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.jsonbset_delete(set._inner, _c_key)
	return &Set{_inner: res}
}


// JsonbsetDeleteArray wraps MEOS C function jsonbset_delete_array.
func JsonbsetDeleteArray(set *Set, keys []string) *Set {
	_c_keys := make([]*C.text, len(keys))
	for _i, _v := range keys { _c_keys[_i] = cstring2text(_v) }
	res := C.jsonbset_delete_array(set._inner, (**C.text)(unsafe.Pointer(&_c_keys[0])), C.int(len(keys)))
	return &Set{_inner: res}
}


// JsonbsetExists wraps MEOS C function jsonbset_exists.
func JsonbsetExists(set *Set, key string) *Set {
	_c_key := cstring2text(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.jsonbset_exists(set._inner, _c_key)
	return &Set{_inner: res}
}


// JsonbsetExistsArray wraps MEOS C function jsonbset_exists_array.
func JsonbsetExistsArray(set *Set, keys []string, any bool) *Set {
	_c_keys := make([]*C.text, len(keys))
	for _i, _v := range keys { _c_keys[_i] = cstring2text(_v) }
	res := C.jsonbset_exists_array(set._inner, (**C.text)(unsafe.Pointer(&_c_keys[0])), C.int(len(keys)), C.bool(any))
	return &Set{_inner: res}
}


// JsonbsetSet wraps MEOS C function jsonbset_set.
func JsonbsetSet(set *Set, keys []string, newjb *Jsonb, create bool, null_handle string, lax bool) *Set {
	_c_keys := make([]*C.text, len(keys))
	for _i, _v := range keys { _c_keys[_i] = cstring2text(_v) }
	_c_null_handle := cstring2text(null_handle)
	defer C.free(unsafe.Pointer(_c_null_handle))
	res := C.jsonbset_set(set._inner, (**C.text)(unsafe.Pointer(&_c_keys[0])), C.int(len(keys)), newjb._inner, C.bool(create), _c_null_handle, C.bool(lax))
	return &Set{_inner: res}
}


// JsonbsetToAlphanumset wraps MEOS C function jsonbset_to_alphanumset.
func JsonbsetToAlphanumset(set *Set, key string, settype MeosType, null_handle NullHandleType) *Set {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.jsonbset_to_alphanumset(set._inner, _c_key, C.MeosType(settype), C.nullHandleType(null_handle))
	return &Set{_inner: res}
}


// JsonbsetToIntset wraps MEOS C function jsonbset_to_intset.
func JsonbsetToIntset(set *Set, key string, null_handle NullHandleType) *Set {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.jsonbset_to_intset(set._inner, _c_key, C.nullHandleType(null_handle))
	return &Set{_inner: res}
}


// JsonbsetToFloatset wraps MEOS C function jsonbset_to_floatset.
func JsonbsetToFloatset(set *Set, key string, null_handle NullHandleType) *Set {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.jsonbset_to_floatset(set._inner, _c_key, C.nullHandleType(null_handle))
	return &Set{_inner: res}
}


// JsonbsetToTextsetKey wraps MEOS C function jsonbset_to_textset_key.
func JsonbsetToTextsetKey(set *Set, key string, null_handle NullHandleType) *Set {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.jsonbset_to_textset_key(set._inner, _c_key, C.nullHandleType(null_handle))
	return &Set{_inner: res}
}


// JsonbsetStripNulls wraps MEOS C function jsonbset_strip_nulls.
func JsonbsetStripNulls(set *Set, strip_in_arrays bool) *Set {
	res := C.jsonbset_strip_nulls(set._inner, C.bool(strip_in_arrays))
	return &Set{_inner: res}
}


// JsonbsetPretty wraps MEOS C function jsonbset_pretty.
func JsonbsetPretty(set *Set) *Set {
	res := C.jsonbset_pretty(set._inner)
	return &Set{_inner: res}
}


// TODO jsonbset_delete_path: unsupported param text **
// func JsonbsetDeletePath(...) { /* not yet handled by codegen */ }


// TODO jsonbset_extract_path: unsupported param text **
// func JsonbsetExtractPath(...) { /* not yet handled by codegen */ }


// TODO jsonbset_insert: unsupported param text **
// func JsonbsetInsert(...) { /* not yet handled by codegen */ }


// JsonbsetPathExists wraps MEOS C function jsonbset_path_exists.
func JsonbsetPathExists(set *Set, jp *JsonPath, vars *Jsonb, silent bool, tz bool) *Set {
	res := C.jsonbset_path_exists(set._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return &Set{_inner: res}
}


// JsonbsetPathMatch wraps MEOS C function jsonbset_path_match.
func JsonbsetPathMatch(set *Set, jp *JsonPath, vars *Jsonb, silent bool, tz bool) *Set {
	res := C.jsonbset_path_match(set._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return &Set{_inner: res}
}


// JsonbsetPathQueryArray wraps MEOS C function jsonbset_path_query_array.
func JsonbsetPathQueryArray(set *Set, jp *JsonPath, vars *Jsonb, silent bool, tz bool) *Set {
	res := C.jsonbset_path_query_array(set._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return &Set{_inner: res}
}


// JsonbsetPathQueryFirst wraps MEOS C function jsonbset_path_query_first.
func JsonbsetPathQueryFirst(set *Set, jp *JsonPath, vars *Jsonb, silent bool, tz bool) *Set {
	res := C.jsonbset_path_query_first(set._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return &Set{_inner: res}
}


// ContainedJsonbSet wraps MEOS C function contained_jsonb_set.
func ContainedJsonbSet(jb *Jsonb, s *Set) bool {
	res := C.contained_jsonb_set(jb._inner, s._inner)
	return bool(res)
}


// ContainsSetJsonb wraps MEOS C function contains_set_jsonb.
func ContainsSetJsonb(s *Set, jb *Jsonb) bool {
	res := C.contains_set_jsonb(s._inner, jb._inner)
	return bool(res)
}


// IntersectionJsonbSet wraps MEOS C function intersection_jsonb_set.
func IntersectionJsonbSet(jb *Jsonb, s *Set) *Set {
	res := C.intersection_jsonb_set(jb._inner, s._inner)
	return &Set{_inner: res}
}


// IntersectionSetJsonb wraps MEOS C function intersection_set_jsonb.
func IntersectionSetJsonb(s *Set, jb *Jsonb) *Set {
	res := C.intersection_set_jsonb(s._inner, jb._inner)
	return &Set{_inner: res}
}


// JsonbUnionTransfn wraps MEOS C function jsonb_union_transfn.
func JsonbUnionTransfn(state *Set, jb *Jsonb) *Set {
	res := C.jsonb_union_transfn(state._inner, jb._inner)
	return &Set{_inner: res}
}


// MinusJsonbSet wraps MEOS C function minus_jsonb_set.
func MinusJsonbSet(jb *Jsonb, s *Set) *Set {
	res := C.minus_jsonb_set(jb._inner, s._inner)
	return &Set{_inner: res}
}


// MinusSetJsonb wraps MEOS C function minus_set_jsonb.
func MinusSetJsonb(s *Set, jb *Jsonb) *Set {
	res := C.minus_set_jsonb(s._inner, jb._inner)
	return &Set{_inner: res}
}


// UnionJsonbSet wraps MEOS C function union_jsonb_set.
func UnionJsonbSet(jb *Jsonb, s *Set) *Set {
	res := C.union_jsonb_set(jb._inner, s._inner)
	return &Set{_inner: res}
}


// UnionSetJsonb wraps MEOS C function union_set_jsonb.
func UnionSetJsonb(s *Set, jb *Jsonb) *Set {
	res := C.union_set_jsonb(s._inner, jb._inner)
	return &Set{_inner: res}
}


// TjsonbFromMFJSON wraps MEOS C function tjsonb_from_mfjson.
func TjsonbFromMFJSON(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tjsonb_from_mfjson(_c_str)
	return CreateTemporal(res)
}


// TjsonbIn wraps MEOS C function tjsonb_in.
func TjsonbIn(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tjsonb_in(_c_str)
	return CreateTemporal(res)
}


// TjsonbOut wraps MEOS C function tjsonb_out.
func TjsonbOut(temp Temporal) string {
	res := C.tjsonb_out(temp.Inner())
	return C.GoString(res)
}


// TjsonbinstIn wraps MEOS C function tjsonbinst_in.
func TjsonbinstIn(str string) TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tjsonbinst_in(_c_str)
	return TInstant{_inner: res}
}


// TjsonbseqIn wraps MEOS C function tjsonbseq_in.
func TjsonbseqIn(str string, interp Interpolation) TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tjsonbseq_in(_c_str, C.interpType(interp))
	return TSequence{_inner: res}
}


// TjsonbseqsetIn wraps MEOS C function tjsonbseqset_in.
func TjsonbseqsetIn(str string) TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tjsonbseqset_in(_c_str)
	return TSequenceSet{_inner: res}
}


// TjsonbFromBaseTemp wraps MEOS C function tjsonb_from_base_temp.
func TjsonbFromBaseTemp(jsonb *Jsonb, temp Temporal) Temporal {
	res := C.tjsonb_from_base_temp(jsonb._inner, temp.Inner())
	return CreateTemporal(res)
}


// TjsonbinstMake wraps MEOS C function tjsonbinst_make.
func TjsonbinstMake(jsonb *Jsonb, t int64) TInstant {
	res := C.tjsonbinst_make(jsonb._inner, C.TimestampTz(t))
	return TInstant{_inner: res}
}


// TjsonbseqFromBaseTstzset wraps MEOS C function tjsonbseq_from_base_tstzset.
func TjsonbseqFromBaseTstzset(jsonb *Jsonb, s *Set) TSequence {
	res := C.tjsonbseq_from_base_tstzset(jsonb._inner, s._inner)
	return TSequence{_inner: res}
}


// TjsonbseqFromBaseTstzspan wraps MEOS C function tjsonbseq_from_base_tstzspan.
func TjsonbseqFromBaseTstzspan(jsonb *Jsonb, sp *Span) TSequence {
	res := C.tjsonbseq_from_base_tstzspan(jsonb._inner, sp._inner)
	return TSequence{_inner: res}
}


// TjsonbseqsetFromBaseTstzspanset wraps MEOS C function tjsonbseqset_from_base_tstzspanset.
func TjsonbseqsetFromBaseTstzspanset(jsonb *Jsonb, ss *SpanSet) TSequenceSet {
	res := C.tjsonbseqset_from_base_tstzspanset(jsonb._inner, ss._inner)
	return TSequenceSet{_inner: res}
}


// TjsonbToTtext wraps MEOS C function tjsonb_to_ttext.
func TjsonbToTtext(temp Temporal) Temporal {
	res := C.tjsonb_to_ttext(temp.Inner())
	return CreateTemporal(res)
}


// TtextToTjsonb wraps MEOS C function ttext_to_tjsonb.
func TtextToTjsonb(temp Temporal) Temporal {
	res := C.ttext_to_tjsonb(temp.Inner())
	return CreateTemporal(res)
}


// TjsonbEndValue wraps MEOS C function tjsonb_end_value.
func TjsonbEndValue(temp Temporal) *Jsonb {
	res := C.tjsonb_end_value(temp.Inner())
	return &Jsonb{_inner: res}
}


// TjsonbStartValue wraps MEOS C function tjsonb_start_value.
func TjsonbStartValue(temp Temporal) *Jsonb {
	res := C.tjsonb_start_value(temp.Inner())
	return &Jsonb{_inner: res}
}


// TjsonbValueAtTimestamptz wraps MEOS C function tjsonb_value_at_timestamptz.
func TjsonbValueAtTimestamptz(temp Temporal, t int64, strict bool) (bool, *Jsonb) {
	var _out_value *C.Jsonb
	res := C.tjsonb_value_at_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(strict), &_out_value)
	return bool(res), &Jsonb{_inner: _out_value}
}


// TjsonbValueN wraps MEOS C function tjsonb_value_n.
func TjsonbValueN(temp Temporal, n int) (bool, *Jsonb) {
	var _out_result *C.Jsonb
	res := C.tjsonb_value_n(temp.Inner(), C.int(n), &_out_result)
	return bool(res), &Jsonb{_inner: _out_result}
}


// TjsonbValues wraps MEOS C function tjsonb_values.
func TjsonbValues(temp Temporal) []*Jsonb {
	var _out_count C.int
	res := C.tjsonb_values(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.Jsonb)(unsafe.Pointer(res)), _n)
	_out := make([]*Jsonb, _n)
	for _i, _e := range _slice {
		_out[_i] = &Jsonb{_inner: _e}
	}
	return _out
}


// ConcatTjsonbJsonb wraps MEOS C function concat_tjsonb_jsonb.
func ConcatTjsonbJsonb(temp Temporal, jb *Jsonb, invert bool) Temporal {
	res := C.concat_tjsonb_jsonb(temp.Inner(), jb._inner, C.bool(invert))
	return CreateTemporal(res)
}


// ConcatTjsonbTjsonb wraps MEOS C function concat_tjsonb_tjsonb.
func ConcatTjsonbTjsonb(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.concat_tjsonb_tjsonb(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// ContainsTjsonbJsonb wraps MEOS C function contains_tjsonb_jsonb.
func ContainsTjsonbJsonb(temp Temporal, jb *Jsonb, invert bool) Temporal {
	res := C.contains_tjsonb_jsonb(temp.Inner(), jb._inner, C.bool(invert))
	return CreateTemporal(res)
}


// ContainsTjsonbTjsonb wraps MEOS C function contains_tjsonb_tjsonb.
func ContainsTjsonbTjsonb(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.contains_tjsonb_tjsonb(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// NullHandleTypeFromString wraps MEOS C function null_handle_type_from_string.
func NullHandleTypeFromString(str string) NullHandleType {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.null_handle_type_from_string(_c_str)
	return NullHandleType(res)
}


// TjsonArrayElement wraps MEOS C function tjson_array_element.
func TjsonArrayElement(temp Temporal, idx int, null_handle NullHandleType) Temporal {
	res := C.tjson_array_element(temp.Inner(), C.int(idx), C.nullHandleType(null_handle))
	return CreateTemporal(res)
}


// TjsonArrayLength wraps MEOS C function tjson_array_length.
func TjsonArrayLength(temp Temporal) Temporal {
	res := C.tjson_array_length(temp.Inner())
	return CreateTemporal(res)
}


// TODO tjson_extract_path: unsupported param text **
// func TjsonExtractPath(...) { /* not yet handled by codegen */ }


// TjsonObjectField wraps MEOS C function tjson_object_field.
func TjsonObjectField(temp Temporal, key string, astext bool, null_handle NullHandleType) Temporal {
	_c_key := cstring2text(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.tjson_object_field(temp.Inner(), _c_key, C.bool(astext), C.nullHandleType(null_handle))
	return CreateTemporal(res)
}


// TjsonStripNulls wraps MEOS C function tjson_strip_nulls.
func TjsonStripNulls(temp Temporal, strip_in_arrays bool) Temporal {
	res := C.tjson_strip_nulls(temp.Inner(), C.bool(strip_in_arrays))
	return CreateTemporal(res)
}


// TjsonbArrayElement wraps MEOS C function tjsonb_array_element.
func TjsonbArrayElement(temp Temporal, idx int, astext bool, null_handle NullHandleType) Temporal {
	res := C.tjsonb_array_element(temp.Inner(), C.int(idx), C.bool(astext), C.nullHandleType(null_handle))
	return CreateTemporal(res)
}


// TjsonbArrayLength wraps MEOS C function tjsonb_array_length.
func TjsonbArrayLength(temp Temporal) Temporal {
	res := C.tjsonb_array_length(temp.Inner())
	return CreateTemporal(res)
}


// TjsonbDelete wraps MEOS C function tjsonb_delete.
func TjsonbDelete(temp Temporal, key string) Temporal {
	_c_key := cstring2text(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.tjsonb_delete(temp.Inner(), _c_key)
	return CreateTemporal(res)
}


// TjsonbDeleteArray wraps MEOS C function tjsonb_delete_array.
func TjsonbDeleteArray(temp Temporal, keys []string) Temporal {
	_c_keys := make([]*C.text, len(keys))
	for _i, _v := range keys { _c_keys[_i] = cstring2text(_v) }
	res := C.tjsonb_delete_array(temp.Inner(), (**C.text)(unsafe.Pointer(&_c_keys[0])), C.int(len(keys)))
	return CreateTemporal(res)
}


// TjsonbDeleteIndex wraps MEOS C function tjsonb_delete_index.
func TjsonbDeleteIndex(temp Temporal, idx int) Temporal {
	res := C.tjsonb_delete_index(temp.Inner(), C.int(idx))
	return CreateTemporal(res)
}


// TODO tjsonb_delete_path: unsupported param text **
// func TjsonbDeletePath(...) { /* not yet handled by codegen */ }


// TjsonbExists wraps MEOS C function tjsonb_exists.
func TjsonbExists(temp Temporal, key string) Temporal {
	_c_key := cstring2text(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.tjsonb_exists(temp.Inner(), _c_key)
	return CreateTemporal(res)
}


// TjsonbExistsArray wraps MEOS C function tjsonb_exists_array.
func TjsonbExistsArray(temp Temporal, keys []string, any bool) Temporal {
	_c_keys := make([]*C.text, len(keys))
	for _i, _v := range keys { _c_keys[_i] = cstring2text(_v) }
	res := C.tjsonb_exists_array(temp.Inner(), (**C.text)(unsafe.Pointer(&_c_keys[0])), C.int(len(keys)), C.bool(any))
	return CreateTemporal(res)
}


// TODO tjsonb_extract_path: unsupported param text **
// func TjsonbExtractPath(...) { /* not yet handled by codegen */ }


// TjsonbInsert wraps MEOS C function tjsonb_insert.
func TjsonbInsert(temp Temporal, keys []string, newjb *Jsonb, after bool) Temporal {
	_c_keys := make([]*C.text, len(keys))
	for _i, _v := range keys { _c_keys[_i] = cstring2text(_v) }
	res := C.tjsonb_insert(temp.Inner(), (**C.text)(unsafe.Pointer(&_c_keys[0])), C.int(len(keys)), newjb._inner, C.bool(after))
	return CreateTemporal(res)
}


// TjsonbObjectField wraps MEOS C function tjsonb_object_field.
func TjsonbObjectField(temp Temporal, key string, astext bool, null_handle NullHandleType) Temporal {
	_c_key := cstring2text(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.tjsonb_object_field(temp.Inner(), _c_key, C.bool(astext), C.nullHandleType(null_handle))
	return CreateTemporal(res)
}


// TjsonbPathExists wraps MEOS C function tjsonb_path_exists.
func TjsonbPathExists(temp Temporal, jp *JsonPath, vars *Jsonb, silent bool, tz bool) Temporal {
	res := C.tjsonb_path_exists(temp.Inner(), jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return CreateTemporal(res)
}


// TjsonbPathMatch wraps MEOS C function tjsonb_path_match.
func TjsonbPathMatch(temp Temporal, jp *JsonPath, vars *Jsonb, silent bool, tz bool) Temporal {
	res := C.tjsonb_path_match(temp.Inner(), jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return CreateTemporal(res)
}


// TjsonbPathQueryArray wraps MEOS C function tjsonb_path_query_array.
func TjsonbPathQueryArray(temp Temporal, jp *JsonPath, vars *Jsonb, silent bool, tz bool) Temporal {
	res := C.tjsonb_path_query_array(temp.Inner(), jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return CreateTemporal(res)
}


// TjsonbPathQueryFirst wraps MEOS C function tjsonb_path_query_first.
func TjsonbPathQueryFirst(temp Temporal, jp *JsonPath, vars *Jsonb, silent bool, tz bool) Temporal {
	res := C.tjsonb_path_query_first(temp.Inner(), jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	return CreateTemporal(res)
}


// TjsonbPretty wraps MEOS C function tjsonb_pretty.
func TjsonbPretty(temp Temporal) Temporal {
	res := C.tjsonb_pretty(temp.Inner())
	return CreateTemporal(res)
}


// TjsonbSet wraps MEOS C function tjsonb_set.
func TjsonbSet(temp Temporal, keys []string, newjb *Jsonb, create bool, handle_null string, lax bool) Temporal {
	_c_keys := make([]*C.text, len(keys))
	for _i, _v := range keys { _c_keys[_i] = cstring2text(_v) }
	_c_handle_null := cstring2text(handle_null)
	defer C.free(unsafe.Pointer(_c_handle_null))
	res := C.tjsonb_set(temp.Inner(), (**C.text)(unsafe.Pointer(&_c_keys[0])), C.int(len(keys)), newjb._inner, C.bool(create), _c_handle_null, C.bool(lax))
	return CreateTemporal(res)
}


// TjsonbStripNulls wraps MEOS C function tjsonb_strip_nulls.
func TjsonbStripNulls(temp Temporal, strip_in_arrays bool) Temporal {
	res := C.tjsonb_strip_nulls(temp.Inner(), C.bool(strip_in_arrays))
	return CreateTemporal(res)
}


// TjsonbToTbool wraps MEOS C function tjsonb_to_tbool.
func TjsonbToTbool(temp Temporal, key string, null_handle NullHandleType) Temporal {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.tjsonb_to_tbool(temp.Inner(), _c_key, C.nullHandleType(null_handle))
	return CreateTemporal(res)
}


// TjsonbToTfloat wraps MEOS C function tjsonb_to_tfloat.
func TjsonbToTfloat(temp Temporal, key string, interp Interpolation, null_handle NullHandleType) Temporal {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.tjsonb_to_tfloat(temp.Inner(), _c_key, C.interpType(interp), C.nullHandleType(null_handle))
	return CreateTemporal(res)
}


// TjsonbToTint wraps MEOS C function tjsonb_to_tint.
func TjsonbToTint(temp Temporal, key string, null_handle NullHandleType) Temporal {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.tjsonb_to_tint(temp.Inner(), _c_key, C.nullHandleType(null_handle))
	return CreateTemporal(res)
}


// TjsonbToTtextKey wraps MEOS C function tjsonb_to_ttext_key.
func TjsonbToTtextKey(temp Temporal, key string, null_handle NullHandleType) Temporal {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	res := C.tjsonb_to_ttext_key(temp.Inner(), _c_key, C.nullHandleType(null_handle))
	return CreateTemporal(res)
}


// TjsonbAtValue wraps MEOS C function tjsonb_at_value.
func TjsonbAtValue(temp Temporal, jsb *Jsonb) Temporal {
	res := C.tjsonb_at_value(temp.Inner(), jsb._inner)
	return CreateTemporal(res)
}


// TjsonbMinusValue wraps MEOS C function tjsonb_minus_value.
func TjsonbMinusValue(temp Temporal, jsb *Jsonb) Temporal {
	res := C.tjsonb_minus_value(temp.Inner(), jsb._inner)
	return CreateTemporal(res)
}


// AlwaysEqJsonbTjsonb wraps MEOS C function always_eq_jsonb_tjsonb.
func AlwaysEqJsonbTjsonb(jb *Jsonb, temp Temporal) int {
	res := C.always_eq_jsonb_tjsonb(jb._inner, temp.Inner())
	return int(res)
}


// AlwaysEqTjsonbJsonb wraps MEOS C function always_eq_tjsonb_jsonb.
func AlwaysEqTjsonbJsonb(temp Temporal, jb *Jsonb) int {
	res := C.always_eq_tjsonb_jsonb(temp.Inner(), jb._inner)
	return int(res)
}


// AlwaysEqTjsonbTjsonb wraps MEOS C function always_eq_tjsonb_tjsonb.
func AlwaysEqTjsonbTjsonb(temp1 Temporal, temp2 Temporal) int {
	res := C.always_eq_tjsonb_tjsonb(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AlwaysNeJsonbTjsonb wraps MEOS C function always_ne_jsonb_tjsonb.
func AlwaysNeJsonbTjsonb(jb *Jsonb, temp Temporal) int {
	res := C.always_ne_jsonb_tjsonb(jb._inner, temp.Inner())
	return int(res)
}


// AlwaysNeTjsonbJsonb wraps MEOS C function always_ne_tjsonb_jsonb.
func AlwaysNeTjsonbJsonb(temp Temporal, jb *Jsonb) int {
	res := C.always_ne_tjsonb_jsonb(temp.Inner(), jb._inner)
	return int(res)
}


// AlwaysNeTjsonbTjsonb wraps MEOS C function always_ne_tjsonb_tjsonb.
func AlwaysNeTjsonbTjsonb(temp1 Temporal, temp2 Temporal) int {
	res := C.always_ne_tjsonb_tjsonb(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverEqJsonbTjsonb wraps MEOS C function ever_eq_jsonb_tjsonb.
func EverEqJsonbTjsonb(jb *Jsonb, temp Temporal) int {
	res := C.ever_eq_jsonb_tjsonb(jb._inner, temp.Inner())
	return int(res)
}


// EverEqTjsonbJsonb wraps MEOS C function ever_eq_tjsonb_jsonb.
func EverEqTjsonbJsonb(temp Temporal, jb *Jsonb) int {
	res := C.ever_eq_tjsonb_jsonb(temp.Inner(), jb._inner)
	return int(res)
}


// EverEqTjsonbTjsonb wraps MEOS C function ever_eq_tjsonb_tjsonb.
func EverEqTjsonbTjsonb(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_eq_tjsonb_tjsonb(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverNeJsonbTjsonb wraps MEOS C function ever_ne_jsonb_tjsonb.
func EverNeJsonbTjsonb(jb *Jsonb, temp Temporal) int {
	res := C.ever_ne_jsonb_tjsonb(jb._inner, temp.Inner())
	return int(res)
}


// EverNeTjsonbJsonb wraps MEOS C function ever_ne_tjsonb_jsonb.
func EverNeTjsonbJsonb(temp Temporal, jb *Jsonb) int {
	res := C.ever_ne_tjsonb_jsonb(temp.Inner(), jb._inner)
	return int(res)
}


// EverNeTjsonbTjsonb wraps MEOS C function ever_ne_tjsonb_tjsonb.
func EverNeTjsonbTjsonb(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_ne_tjsonb_tjsonb(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TeqJsonbTjsonb wraps MEOS C function teq_jsonb_tjsonb.
func TeqJsonbTjsonb(jb *Jsonb, temp Temporal) Temporal {
	res := C.teq_jsonb_tjsonb(jb._inner, temp.Inner())
	return CreateTemporal(res)
}


// TeqTjsonbJsonb wraps MEOS C function teq_tjsonb_jsonb.
func TeqTjsonbJsonb(temp Temporal, jb *Jsonb) Temporal {
	res := C.teq_tjsonb_jsonb(temp.Inner(), jb._inner)
	return CreateTemporal(res)
}


// TneJsonbTjsonb wraps MEOS C function tne_jsonb_tjsonb.
func TneJsonbTjsonb(jb *Jsonb, temp Temporal) Temporal {
	res := C.tne_jsonb_tjsonb(jb._inner, temp.Inner())
	return CreateTemporal(res)
}


// TneTjsonbJsonb wraps MEOS C function tne_tjsonb_jsonb.
func TneTjsonbJsonb(temp Temporal, jb *Jsonb) Temporal {
	res := C.tne_tjsonb_jsonb(temp.Inner(), jb._inner)
	return CreateTemporal(res)
}

