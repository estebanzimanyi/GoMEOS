package generated

// #include <stddef.h>
import "C"
import (
	"unsafe"

	"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

// DateIn wraps MEOS C function date_in.
func DateIn(str string) int32 {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.date_in(_c_str)
	return int32(res)
}


// DateOut wraps MEOS C function date_out.
func DateOut(d int32) string {
	res := C.date_out(C.DateADT(d))
	return C.GoString(res)
}


// IntervalCmp wraps MEOS C function interval_cmp.
func IntervalCmp(interv1 timeutil.Timedelta, interv2 timeutil.Timedelta) int {
	res := C.interval_cmp(interv1.Inner(), interv2.Inner())
	return int(res)
}


// IntervalIn wraps MEOS C function interval_in.
func IntervalIn(str string, typmod int32) timeutil.Timedelta {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.interval_in(_c_str, C.int32(typmod))
	return IntervalToTimeDelta(res)
}


// IntervalOut wraps MEOS C function interval_out.
func IntervalOut(interv timeutil.Timedelta) string {
	res := C.interval_out(interv.Inner())
	return C.GoString(res)
}


// TimeIn wraps MEOS C function time_in.
func TimeIn(str string, typmod int32) int64 {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.time_in(_c_str, C.int32(typmod))
	return int64(res)
}


// TimeOut wraps MEOS C function time_out.
func TimeOut(t int64) string {
	res := C.time_out(C.TimeADT(t))
	return C.GoString(res)
}


// TimestampIn wraps MEOS C function timestamp_in.
func TimestampIn(str string, typmod int32) int64 {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.timestamp_in(_c_str, C.int32(typmod))
	return int64(res)
}


// TimestampOut wraps MEOS C function timestamp_out.
func TimestampOut(t int64) string {
	res := C.timestamp_out(C.Timestamp(t))
	return C.GoString(res)
}


// TimestamptzIn wraps MEOS C function timestamptz_in.
func TimestamptzIn(str string, typmod int32) int64 {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.timestamptz_in(_c_str, C.int32(typmod))
	return int64(res)
}


// TimestamptzOut wraps MEOS C function timestamptz_out.
func TimestamptzOut(t int64) string {
	res := C.timestamptz_out(C.TimestampTz(t))
	return C.GoString(res)
}


// RtreeCreateIntspan wraps MEOS C function rtree_create_intspan.
func RtreeCreateIntspan() *RTree {
	res := C.rtree_create_intspan()
	return &RTree{_inner: res}
}


// RtreeCreateBigintspan wraps MEOS C function rtree_create_bigintspan.
func RtreeCreateBigintspan() *RTree {
	res := C.rtree_create_bigintspan()
	return &RTree{_inner: res}
}


// RtreeCreateFloatspan wraps MEOS C function rtree_create_floatspan.
func RtreeCreateFloatspan() *RTree {
	res := C.rtree_create_floatspan()
	return &RTree{_inner: res}
}


// RtreeCreateDatespan wraps MEOS C function rtree_create_datespan.
func RtreeCreateDatespan() *RTree {
	res := C.rtree_create_datespan()
	return &RTree{_inner: res}
}


// RtreeCreateTstzspan wraps MEOS C function rtree_create_tstzspan.
func RtreeCreateTstzspan() *RTree {
	res := C.rtree_create_tstzspan()
	return &RTree{_inner: res}
}


// RtreeCreateTBOX wraps MEOS C function rtree_create_tbox.
func RtreeCreateTBOX() *RTree {
	res := C.rtree_create_tbox()
	return &RTree{_inner: res}
}


// RtreeCreateSTBOX wraps MEOS C function rtree_create_stbox.
func RtreeCreateSTBOX() *RTree {
	res := C.rtree_create_stbox()
	return &RTree{_inner: res}
}


// RtreeFree wraps MEOS C function rtree_free.
func RtreeFree(rtree *RTree) {
	C.rtree_free(rtree._inner)
}


// RtreeInsert wraps MEOS C function rtree_insert.
func RtreeInsert(rtree *RTree, box unsafe.Pointer, id int64) {
	C.rtree_insert(rtree._inner, unsafe.Pointer(box), C.int64(id))
}


// RtreeSearch wraps MEOS C function rtree_search.
func RtreeSearch(rtree *RTree, query unsafe.Pointer) []int {
	var _out_count C.int
	res := C.rtree_search(rtree._inner, unsafe.Pointer(query), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((*C.int)(unsafe.Pointer(res)), _n)
	_out := make([]int, _n)
	for _i, _e := range _slice {
		_out[_i] = int(_e)
	}
	return _out
}


// MeosError wraps MEOS C function meos_error.
func MeosError(errlevel int, errcode int, format string) {
	_c_format := C.CString(format)
	defer C.free(unsafe.Pointer(_c_format))
	C.meos_error(C.int(errlevel), C.int(errcode), _c_format)
}


// MeosErrno wraps MEOS C function meos_errno.
func MeosErrno() int {
	res := C.meos_errno()
	return int(res)
}


// MeosErrnoSet wraps MEOS C function meos_errno_set.
func MeosErrnoSet(err int) int {
	res := C.meos_errno_set(C.int(err))
	return int(res)
}


// MeosErrnoRestore wraps MEOS C function meos_errno_restore.
func MeosErrnoRestore(err int) int {
	res := C.meos_errno_restore(C.int(err))
	return int(res)
}


// MeosErrnoReset wraps MEOS C function meos_errno_reset.
func MeosErrnoReset() int {
	res := C.meos_errno_reset()
	return int(res)
}


// MeosInitializeTimezone wraps MEOS C function meos_initialize_timezone.
func MeosInitializeTimezone(name string) {
	_c_name := C.CString(name)
	defer C.free(unsafe.Pointer(_c_name))
	C.meos_initialize_timezone(_c_name)
}


// MeosFinalizeTimezone wraps MEOS C function meos_finalize_timezone.
func MeosFinalizeTimezone() {
	C.meos_finalize_timezone()
}


// MeosFinalizeProjsrs wraps MEOS C function meos_finalize_projsrs.
func MeosFinalizeProjsrs() {
	C.meos_finalize_projsrs()
}


// MeosFinalizeWays wraps MEOS C function meos_finalize_ways.
func MeosFinalizeWays() {
	C.meos_finalize_ways()
}


// MeosSetDatestyle wraps MEOS C function meos_set_datestyle.
func MeosSetDatestyle(newval string, extra unsafe.Pointer) bool {
	_c_newval := C.CString(newval)
	defer C.free(unsafe.Pointer(_c_newval))
	res := C.meos_set_datestyle(_c_newval, unsafe.Pointer(extra))
	return bool(res)
}


// MeosSetIntervalstyle wraps MEOS C function meos_set_intervalstyle.
func MeosSetIntervalstyle(newval string, extra int) bool {
	_c_newval := C.CString(newval)
	defer C.free(unsafe.Pointer(_c_newval))
	res := C.meos_set_intervalstyle(_c_newval, C.int(extra))
	return bool(res)
}


// MeosGetDatestyle wraps MEOS C function meos_get_datestyle.
func MeosGetDatestyle() string {
	res := C.meos_get_datestyle()
	return C.GoString(res)
}


// MeosGetIntervalstyle wraps MEOS C function meos_get_intervalstyle.
func MeosGetIntervalstyle() string {
	res := C.meos_get_intervalstyle()
	return C.GoString(res)
}


// MeosSetSpatialRefSysCsv wraps MEOS C function meos_set_spatial_ref_sys_csv.
func MeosSetSpatialRefSysCsv(path string) {
	_c_path := C.CString(path)
	defer C.free(unsafe.Pointer(_c_path))
	C.meos_set_spatial_ref_sys_csv(_c_path)
}


// MeosInitialize wraps MEOS C function meos_initialize.
func MeosInitialize() {
	C.meos_initialize()
}


// MeosFinalize wraps MEOS C function meos_finalize.
func MeosFinalize() {
	C.meos_finalize()
}


// AddDateInt wraps MEOS C function add_date_int.
func AddDateInt(d int32, days int32) int32 {
	res := C.add_date_int(C.DateADT(d), C.int32(days))
	return int32(res)
}


// AddIntervalInterval wraps MEOS C function add_interval_interval.
func AddIntervalInterval(interv1 timeutil.Timedelta, interv2 timeutil.Timedelta) timeutil.Timedelta {
	res := C.add_interval_interval(interv1.Inner(), interv2.Inner())
	return IntervalToTimeDelta(res)
}


// AddTimestamptzInterval wraps MEOS C function add_timestamptz_interval.
func AddTimestamptzInterval(t int64, interv timeutil.Timedelta) int64 {
	res := C.add_timestamptz_interval(C.TimestampTz(t), interv.Inner())
	return int64(res)
}


// BoolIn wraps MEOS C function bool_in.
func BoolIn(str string) bool {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.bool_in(_c_str)
	return bool(res)
}


// BoolOut wraps MEOS C function bool_out.
func BoolOut(b bool) string {
	res := C.bool_out(C.bool(b))
	return C.GoString(res)
}


// Cstring2text wraps MEOS C function cstring2text.
func Cstring2text(str string) string {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.cstring2text(_c_str)
	return text2cstring(res)
}


// DateToTimestamp wraps MEOS C function date_to_timestamp.
func DateToTimestamp(dateVal int32) int64 {
	res := C.date_to_timestamp(C.DateADT(dateVal))
	return int64(res)
}


// DateToTimestamptz wraps MEOS C function date_to_timestamptz.
func DateToTimestamptz(d int32) int64 {
	res := C.date_to_timestamptz(C.DateADT(d))
	return int64(res)
}


// FloatExp wraps MEOS C function float_exp.
func FloatExp(d float64) float64 {
	res := C.float_exp(C.double(d))
	return float64(res)
}


// FloatLn wraps MEOS C function float_ln.
func FloatLn(d float64) float64 {
	res := C.float_ln(C.double(d))
	return float64(res)
}


// FloatLog10 wraps MEOS C function float_log10.
func FloatLog10(d float64) float64 {
	res := C.float_log10(C.double(d))
	return float64(res)
}


// Float8Out wraps MEOS C function float8_out.
func Float8Out(d float64, maxdd int) string {
	res := C.float8_out(C.double(d), C.int(maxdd))
	return C.GoString(res)
}


// FloatRound wraps MEOS C function float_round.
func FloatRound(d float64, maxdd int) float64 {
	res := C.float_round(C.double(d), C.int(maxdd))
	return float64(res)
}


// Int32Cmp wraps MEOS C function int32_cmp.
func Int32Cmp(l int32, r int32) int {
	res := C.int32_cmp(C.int32(l), C.int32(r))
	return int(res)
}


// Int64Cmp wraps MEOS C function int64_cmp.
func Int64Cmp(l int64, r int64) int {
	res := C.int64_cmp(C.int64(l), C.int64(r))
	return int(res)
}


// IntervalMake wraps MEOS C function interval_make.
func IntervalMake(years int32, months int32, weeks int32, days int32, hours int32, mins int32, secs float64) timeutil.Timedelta {
	res := C.interval_make(C.int32(years), C.int32(months), C.int32(weeks), C.int32(days), C.int32(hours), C.int32(mins), C.double(secs))
	return IntervalToTimeDelta(res)
}


// MinusDateDate wraps MEOS C function minus_date_date.
func MinusDateDate(d1 int32, d2 int32) int {
	res := C.minus_date_date(C.DateADT(d1), C.DateADT(d2))
	return int(res)
}


// MinusDateInt wraps MEOS C function minus_date_int.
func MinusDateInt(d int32, days int32) int32 {
	res := C.minus_date_int(C.DateADT(d), C.int32(days))
	return int32(res)
}


// MinusTimestamptzInterval wraps MEOS C function minus_timestamptz_interval.
func MinusTimestamptzInterval(t int64, interv timeutil.Timedelta) int64 {
	res := C.minus_timestamptz_interval(C.TimestampTz(t), interv.Inner())
	return int64(res)
}


// MinusTimestamptzTimestamptz wraps MEOS C function minus_timestamptz_timestamptz.
func MinusTimestamptzTimestamptz(t1 int64, t2 int64) timeutil.Timedelta {
	res := C.minus_timestamptz_timestamptz(C.TimestampTz(t1), C.TimestampTz(t2))
	return IntervalToTimeDelta(res)
}


// MulIntervalDouble wraps MEOS C function mul_interval_double.
func MulIntervalDouble(interv timeutil.Timedelta, factor float64) timeutil.Timedelta {
	res := C.mul_interval_double(interv.Inner(), C.double(factor))
	return IntervalToTimeDelta(res)
}


// PgDateIn wraps MEOS C function pg_date_in.
func PgDateIn(str string) int32 {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.pg_date_in(_c_str)
	return int32(res)
}


// PgDateOut wraps MEOS C function pg_date_out.
func PgDateOut(d int32) string {
	res := C.pg_date_out(C.DateADT(d))
	return C.GoString(res)
}


// PgIntervalCmp wraps MEOS C function pg_interval_cmp.
func PgIntervalCmp(interv1 timeutil.Timedelta, interv2 timeutil.Timedelta) int {
	res := C.pg_interval_cmp(interv1.Inner(), interv2.Inner())
	return int(res)
}


// PgIntervalIn wraps MEOS C function pg_interval_in.
func PgIntervalIn(str string, typmod int32) timeutil.Timedelta {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.pg_interval_in(_c_str, C.int32(typmod))
	return IntervalToTimeDelta(res)
}


// PgIntervalOut wraps MEOS C function pg_interval_out.
func PgIntervalOut(interv timeutil.Timedelta) string {
	res := C.pg_interval_out(interv.Inner())
	return C.GoString(res)
}


// PgTimestampIn wraps MEOS C function pg_timestamp_in.
func PgTimestampIn(str string, typmod int32) int64 {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.pg_timestamp_in(_c_str, C.int32(typmod))
	return int64(res)
}


// PgTimestampOut wraps MEOS C function pg_timestamp_out.
func PgTimestampOut(t int64) string {
	res := C.pg_timestamp_out(C.Timestamp(t))
	return C.GoString(res)
}


// PgTimestamptzIn wraps MEOS C function pg_timestamptz_in.
func PgTimestamptzIn(str string, typmod int32) int64 {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.pg_timestamptz_in(_c_str, C.int32(typmod))
	return int64(res)
}


// PgTimestamptzOut wraps MEOS C function pg_timestamptz_out.
func PgTimestamptzOut(t int64) string {
	res := C.pg_timestamptz_out(C.TimestampTz(t))
	return C.GoString(res)
}


// Text2cstring wraps MEOS C function text2cstring.
func Text2cstring(txt string) string {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.text2cstring(_c_txt)
	return C.GoString(res)
}


// TextCmp wraps MEOS C function text_cmp.
func TextCmp(txt1 string, txt2 string) int {
	_c_txt1 := cstring2text(txt1)
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := cstring2text(txt2)
	defer C.free(unsafe.Pointer(_c_txt2))
	res := C.text_cmp(_c_txt1, _c_txt2)
	return int(res)
}


// TextCopy wraps MEOS C function text_copy.
func TextCopy(txt string) string {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.text_copy(_c_txt)
	return text2cstring(res)
}


// TextIn wraps MEOS C function text_in.
func TextIn(str string) string {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.text_in(_c_str)
	return text2cstring(res)
}


// TextInitcap wraps MEOS C function text_initcap.
func TextInitcap(txt string) string {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.text_initcap(_c_txt)
	return text2cstring(res)
}


// TextLower wraps MEOS C function text_lower.
func TextLower(txt string) string {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.text_lower(_c_txt)
	return text2cstring(res)
}


// TextOut wraps MEOS C function text_out.
func TextOut(txt string) string {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.text_out(_c_txt)
	return C.GoString(res)
}


// TextUpper wraps MEOS C function text_upper.
func TextUpper(txt string) string {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.text_upper(_c_txt)
	return text2cstring(res)
}


// TextcatTextText wraps MEOS C function textcat_text_text.
func TextcatTextText(txt1 string, txt2 string) string {
	_c_txt1 := cstring2text(txt1)
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := cstring2text(txt2)
	defer C.free(unsafe.Pointer(_c_txt2))
	res := C.textcat_text_text(_c_txt1, _c_txt2)
	return text2cstring(res)
}


// TimestamptzShift wraps MEOS C function timestamptz_shift.
func TimestamptzShift(t int64, interv timeutil.Timedelta) int64 {
	res := C.timestamptz_shift(C.TimestampTz(t), interv.Inner())
	return int64(res)
}


// TimestampToDate wraps MEOS C function timestamp_to_date.
func TimestampToDate(t int64) int32 {
	res := C.timestamp_to_date(C.Timestamp(t))
	return int32(res)
}


// TimestamptzToDate wraps MEOS C function timestamptz_to_date.
func TimestamptzToDate(t int64) int32 {
	res := C.timestamptz_to_date(C.TimestampTz(t))
	return int32(res)
}


// BigintsetIn wraps MEOS C function bigintset_in.
func BigintsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.bigintset_in(_c_str)
	return &Set{_inner: res}
}


// BigintsetOut wraps MEOS C function bigintset_out.
func BigintsetOut(set *Set) string {
	res := C.bigintset_out(set._inner)
	return C.GoString(res)
}


// BigintspanExpand wraps MEOS C function bigintspan_expand.
func BigintspanExpand(s *Span, value int64) *Span {
	res := C.bigintspan_expand(s._inner, C.int64(value))
	return &Span{_inner: res}
}


// BigintspanIn wraps MEOS C function bigintspan_in.
func BigintspanIn(str string) *Span {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.bigintspan_in(_c_str)
	return &Span{_inner: res}
}


// BigintspanOut wraps MEOS C function bigintspan_out.
func BigintspanOut(s *Span) string {
	res := C.bigintspan_out(s._inner)
	return C.GoString(res)
}


// BigintspansetIn wraps MEOS C function bigintspanset_in.
func BigintspansetIn(str string) *SpanSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.bigintspanset_in(_c_str)
	return &SpanSet{_inner: res}
}


// BigintspansetOut wraps MEOS C function bigintspanset_out.
func BigintspansetOut(ss *SpanSet) string {
	res := C.bigintspanset_out(ss._inner)
	return C.GoString(res)
}


// DatesetIn wraps MEOS C function dateset_in.
func DatesetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.dateset_in(_c_str)
	return &Set{_inner: res}
}


// DatesetOut wraps MEOS C function dateset_out.
func DatesetOut(s *Set) string {
	res := C.dateset_out(s._inner)
	return C.GoString(res)
}


// DatespanIn wraps MEOS C function datespan_in.
func DatespanIn(str string) *Span {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.datespan_in(_c_str)
	return &Span{_inner: res}
}


// DatespanOut wraps MEOS C function datespan_out.
func DatespanOut(s *Span) string {
	res := C.datespan_out(s._inner)
	return C.GoString(res)
}


// DatespansetIn wraps MEOS C function datespanset_in.
func DatespansetIn(str string) *SpanSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.datespanset_in(_c_str)
	return &SpanSet{_inner: res}
}


// DatespansetOut wraps MEOS C function datespanset_out.
func DatespansetOut(ss *SpanSet) string {
	res := C.datespanset_out(ss._inner)
	return C.GoString(res)
}


// FloatsetIn wraps MEOS C function floatset_in.
func FloatsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.floatset_in(_c_str)
	return &Set{_inner: res}
}


// FloatsetOut wraps MEOS C function floatset_out.
func FloatsetOut(set *Set, maxdd int) string {
	res := C.floatset_out(set._inner, C.int(maxdd))
	return C.GoString(res)
}


// FloatspanExpand wraps MEOS C function floatspan_expand.
func FloatspanExpand(s *Span, value float64) *Span {
	res := C.floatspan_expand(s._inner, C.double(value))
	return &Span{_inner: res}
}


// FloatspanIn wraps MEOS C function floatspan_in.
func FloatspanIn(str string) *Span {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.floatspan_in(_c_str)
	return &Span{_inner: res}
}


// FloatspanOut wraps MEOS C function floatspan_out.
func FloatspanOut(s *Span, maxdd int) string {
	res := C.floatspan_out(s._inner, C.int(maxdd))
	return C.GoString(res)
}


// FloatspansetIn wraps MEOS C function floatspanset_in.
func FloatspansetIn(str string) *SpanSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.floatspanset_in(_c_str)
	return &SpanSet{_inner: res}
}


// FloatspansetOut wraps MEOS C function floatspanset_out.
func FloatspansetOut(ss *SpanSet, maxdd int) string {
	res := C.floatspanset_out(ss._inner, C.int(maxdd))
	return C.GoString(res)
}


// IntsetIn wraps MEOS C function intset_in.
func IntsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.intset_in(_c_str)
	return &Set{_inner: res}
}


// IntsetOut wraps MEOS C function intset_out.
func IntsetOut(set *Set) string {
	res := C.intset_out(set._inner)
	return C.GoString(res)
}


// IntspanExpand wraps MEOS C function intspan_expand.
func IntspanExpand(s *Span, value int32) *Span {
	res := C.intspan_expand(s._inner, C.int32(value))
	return &Span{_inner: res}
}


// IntspanIn wraps MEOS C function intspan_in.
func IntspanIn(str string) *Span {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.intspan_in(_c_str)
	return &Span{_inner: res}
}


// IntspanOut wraps MEOS C function intspan_out.
func IntspanOut(s *Span) string {
	res := C.intspan_out(s._inner)
	return C.GoString(res)
}


// IntspansetIn wraps MEOS C function intspanset_in.
func IntspansetIn(str string) *SpanSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.intspanset_in(_c_str)
	return &SpanSet{_inner: res}
}


// IntspansetOut wraps MEOS C function intspanset_out.
func IntspansetOut(ss *SpanSet) string {
	res := C.intspanset_out(ss._inner)
	return C.GoString(res)
}


// SetAsHexwkb wraps MEOS C function set_as_hexwkb.
func SetAsHexwkb(s *Set, variant uint8) (string, uint) {
	var _out_size_out C.size_t
	res := C.set_as_hexwkb(s._inner, C.uint8_t(variant), &_out_size_out)
	return C.GoString(res), uint(_out_size_out)
}


// SetAsWKB wraps MEOS C function set_as_wkb.
func SetAsWKB(s *Set, variant uint8) []uint8 {
	var _out_size_out C.size_t
	res := C.set_as_wkb(s._inner, C.uint8_t(variant), &_out_size_out)
	_n := int(_out_size_out)
	_slice := unsafe.Slice((*C.uint8_t)(unsafe.Pointer(res)), _n)
	_out := make([]uint8, _n)
	for _i, _e := range _slice {
		_out[_i] = uint8(_e)
	}
	return _out
}


// SetFromHexwkb wraps MEOS C function set_from_hexwkb.
func SetFromHexwkb(hexwkb string) *Set {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	res := C.set_from_hexwkb(_c_hexwkb)
	return &Set{_inner: res}
}


// SetFromWKB wraps MEOS C function set_from_wkb.
func SetFromWKB(wkb []byte) *Set {
	res := C.set_from_wkb((*C.uint8_t)(unsafe.Pointer(&wkb[0])), C.size_t(len(wkb)))
	return &Set{_inner: res}
}


// SpanAsHexwkb wraps MEOS C function span_as_hexwkb.
func SpanAsHexwkb(s *Span, variant uint8) (string, uint) {
	var _out_size_out C.size_t
	res := C.span_as_hexwkb(s._inner, C.uint8_t(variant), &_out_size_out)
	return C.GoString(res), uint(_out_size_out)
}


// SpanAsWKB wraps MEOS C function span_as_wkb.
func SpanAsWKB(s *Span, variant uint8) []uint8 {
	var _out_size_out C.size_t
	res := C.span_as_wkb(s._inner, C.uint8_t(variant), &_out_size_out)
	_n := int(_out_size_out)
	_slice := unsafe.Slice((*C.uint8_t)(unsafe.Pointer(res)), _n)
	_out := make([]uint8, _n)
	for _i, _e := range _slice {
		_out[_i] = uint8(_e)
	}
	return _out
}


// SpanFromHexwkb wraps MEOS C function span_from_hexwkb.
func SpanFromHexwkb(hexwkb string) *Span {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	res := C.span_from_hexwkb(_c_hexwkb)
	return &Span{_inner: res}
}


// SpanFromWKB wraps MEOS C function span_from_wkb.
func SpanFromWKB(wkb []byte) *Span {
	res := C.span_from_wkb((*C.uint8_t)(unsafe.Pointer(&wkb[0])), C.size_t(len(wkb)))
	return &Span{_inner: res}
}


// SpansetAsHexwkb wraps MEOS C function spanset_as_hexwkb.
func SpansetAsHexwkb(ss *SpanSet, variant uint8) (string, uint) {
	var _out_size_out C.size_t
	res := C.spanset_as_hexwkb(ss._inner, C.uint8_t(variant), &_out_size_out)
	return C.GoString(res), uint(_out_size_out)
}


// SpansetAsWKB wraps MEOS C function spanset_as_wkb.
func SpansetAsWKB(ss *SpanSet, variant uint8) []uint8 {
	var _out_size_out C.size_t
	res := C.spanset_as_wkb(ss._inner, C.uint8_t(variant), &_out_size_out)
	_n := int(_out_size_out)
	_slice := unsafe.Slice((*C.uint8_t)(unsafe.Pointer(res)), _n)
	_out := make([]uint8, _n)
	for _i, _e := range _slice {
		_out[_i] = uint8(_e)
	}
	return _out
}


// SpansetFromHexwkb wraps MEOS C function spanset_from_hexwkb.
func SpansetFromHexwkb(hexwkb string) *SpanSet {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	res := C.spanset_from_hexwkb(_c_hexwkb)
	return &SpanSet{_inner: res}
}


// SpansetFromWKB wraps MEOS C function spanset_from_wkb.
func SpansetFromWKB(wkb []byte) *SpanSet {
	res := C.spanset_from_wkb((*C.uint8_t)(unsafe.Pointer(&wkb[0])), C.size_t(len(wkb)))
	return &SpanSet{_inner: res}
}


// TextsetIn wraps MEOS C function textset_in.
func TextsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.textset_in(_c_str)
	return &Set{_inner: res}
}


// TextsetOut wraps MEOS C function textset_out.
func TextsetOut(set *Set) string {
	res := C.textset_out(set._inner)
	return C.GoString(res)
}


// TstzsetIn wraps MEOS C function tstzset_in.
func TstzsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tstzset_in(_c_str)
	return &Set{_inner: res}
}


// TstzsetOut wraps MEOS C function tstzset_out.
func TstzsetOut(set *Set) string {
	res := C.tstzset_out(set._inner)
	return C.GoString(res)
}


// TstzspanIn wraps MEOS C function tstzspan_in.
func TstzspanIn(str string) *Span {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tstzspan_in(_c_str)
	return &Span{_inner: res}
}


// TstzspanOut wraps MEOS C function tstzspan_out.
func TstzspanOut(s *Span) string {
	res := C.tstzspan_out(s._inner)
	return C.GoString(res)
}


// TstzspansetIn wraps MEOS C function tstzspanset_in.
func TstzspansetIn(str string) *SpanSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tstzspanset_in(_c_str)
	return &SpanSet{_inner: res}
}


// TstzspansetOut wraps MEOS C function tstzspanset_out.
func TstzspansetOut(ss *SpanSet) string {
	res := C.tstzspanset_out(ss._inner)
	return C.GoString(res)
}


// BigintsetMake wraps MEOS C function bigintset_make.
func BigintsetMake(values []int64) *Set {
	_c_values := make([]C.int64, len(values))
	for _i, _v := range values { _c_values[_i] = C.int64(_v) }
	res := C.bigintset_make(&_c_values[0], C.int(len(values)))
	return &Set{_inner: res}
}


// BigintspanMake wraps MEOS C function bigintspan_make.
func BigintspanMake(lower int64, upper int64, lower_inc bool, upper_inc bool) *Span {
	res := C.bigintspan_make(C.int64(lower), C.int64(upper), C.bool(lower_inc), C.bool(upper_inc))
	return &Span{_inner: res}
}


// DatesetMake wraps MEOS C function dateset_make.
func DatesetMake(values []int32) *Set {
	_c_values := make([]C.DateADT, len(values))
	for _i, _v := range values { _c_values[_i] = C.DateADT(_v) }
	res := C.dateset_make(&_c_values[0], C.int(len(values)))
	return &Set{_inner: res}
}


// DatespanMake wraps MEOS C function datespan_make.
func DatespanMake(lower int32, upper int32, lower_inc bool, upper_inc bool) *Span {
	res := C.datespan_make(C.DateADT(lower), C.DateADT(upper), C.bool(lower_inc), C.bool(upper_inc))
	return &Span{_inner: res}
}


// FloatsetMake wraps MEOS C function floatset_make.
func FloatsetMake(values []float64) *Set {
	_c_values := make([]C.double, len(values))
	for _i, _v := range values { _c_values[_i] = C.double(_v) }
	res := C.floatset_make(&_c_values[0], C.int(len(values)))
	return &Set{_inner: res}
}


// FloatspanMake wraps MEOS C function floatspan_make.
func FloatspanMake(lower float64, upper float64, lower_inc bool, upper_inc bool) *Span {
	res := C.floatspan_make(C.double(lower), C.double(upper), C.bool(lower_inc), C.bool(upper_inc))
	return &Span{_inner: res}
}


// IntsetMake wraps MEOS C function intset_make.
func IntsetMake(values []int) *Set {
	_c_values := make([]C.int, len(values))
	for _i, _v := range values { _c_values[_i] = C.int(_v) }
	res := C.intset_make(&_c_values[0], C.int(len(values)))
	return &Set{_inner: res}
}


// IntspanMake wraps MEOS C function intspan_make.
func IntspanMake(lower int, upper int, lower_inc bool, upper_inc bool) *Span {
	res := C.intspan_make(C.int(lower), C.int(upper), C.bool(lower_inc), C.bool(upper_inc))
	return &Span{_inner: res}
}


// SetCopy wraps MEOS C function set_copy.
func SetCopy(s *Set) *Set {
	res := C.set_copy(s._inner)
	return &Set{_inner: res}
}


// SpanCopy wraps MEOS C function span_copy.
func SpanCopy(s *Span) *Span {
	res := C.span_copy(s._inner)
	return &Span{_inner: res}
}


// SpansetCopy wraps MEOS C function spanset_copy.
func SpansetCopy(ss *SpanSet) *SpanSet {
	res := C.spanset_copy(ss._inner)
	return &SpanSet{_inner: res}
}


// SpansetMake wraps MEOS C function spanset_make.
func SpansetMake(spans *Span, count int) *SpanSet {
	res := C.spanset_make(spans._inner, C.int(count))
	return &SpanSet{_inner: res}
}


// TextsetMake wraps MEOS C function textset_make.
func TextsetMake(values []string) *Set {
	_c_values := make([]*C.text, len(values))
	for _i, _v := range values { _c_values[_i] = cstring2text(_v) }
	res := C.textset_make((**C.text)(unsafe.Pointer(&_c_values[0])), C.int(len(values)))
	return &Set{_inner: res}
}


// TstzsetMake wraps MEOS C function tstzset_make.
func TstzsetMake(values []int64) *Set {
	_c_values := make([]C.TimestampTz, len(values))
	for _i, _v := range values { _c_values[_i] = C.TimestampTz(_v) }
	res := C.tstzset_make(&_c_values[0], C.int(len(values)))
	return &Set{_inner: res}
}


// TstzspanMake wraps MEOS C function tstzspan_make.
func TstzspanMake(lower int64, upper int64, lower_inc bool, upper_inc bool) *Span {
	res := C.tstzspan_make(C.TimestampTz(lower), C.TimestampTz(upper), C.bool(lower_inc), C.bool(upper_inc))
	return &Span{_inner: res}
}


// BigintToSet wraps MEOS C function bigint_to_set.
func BigintToSet(i int64) *Set {
	res := C.bigint_to_set(C.int64(i))
	return &Set{_inner: res}
}


// BigintToSpan wraps MEOS C function bigint_to_span.
func BigintToSpan(i int) *Span {
	res := C.bigint_to_span(C.int(i))
	return &Span{_inner: res}
}


// BigintToSpanset wraps MEOS C function bigint_to_spanset.
func BigintToSpanset(i int) *SpanSet {
	res := C.bigint_to_spanset(C.int(i))
	return &SpanSet{_inner: res}
}


// DateToSet wraps MEOS C function date_to_set.
func DateToSet(d int32) *Set {
	res := C.date_to_set(C.DateADT(d))
	return &Set{_inner: res}
}


// DateToSpan wraps MEOS C function date_to_span.
func DateToSpan(d int32) *Span {
	res := C.date_to_span(C.DateADT(d))
	return &Span{_inner: res}
}


// DateToSpanset wraps MEOS C function date_to_spanset.
func DateToSpanset(d int32) *SpanSet {
	res := C.date_to_spanset(C.DateADT(d))
	return &SpanSet{_inner: res}
}


// DatesetToTstzset wraps MEOS C function dateset_to_tstzset.
func DatesetToTstzset(s *Set) *Set {
	res := C.dateset_to_tstzset(s._inner)
	return &Set{_inner: res}
}


// DatespanToTstzspan wraps MEOS C function datespan_to_tstzspan.
func DatespanToTstzspan(s *Span) *Span {
	res := C.datespan_to_tstzspan(s._inner)
	return &Span{_inner: res}
}


// DatespansetToTstzspanset wraps MEOS C function datespanset_to_tstzspanset.
func DatespansetToTstzspanset(ss *SpanSet) *SpanSet {
	res := C.datespanset_to_tstzspanset(ss._inner)
	return &SpanSet{_inner: res}
}


// FloatToSet wraps MEOS C function float_to_set.
func FloatToSet(d float64) *Set {
	res := C.float_to_set(C.double(d))
	return &Set{_inner: res}
}


// FloatToSpan wraps MEOS C function float_to_span.
func FloatToSpan(d float64) *Span {
	res := C.float_to_span(C.double(d))
	return &Span{_inner: res}
}


// FloatToSpanset wraps MEOS C function float_to_spanset.
func FloatToSpanset(d float64) *SpanSet {
	res := C.float_to_spanset(C.double(d))
	return &SpanSet{_inner: res}
}


// FloatsetToIntset wraps MEOS C function floatset_to_intset.
func FloatsetToIntset(s *Set) *Set {
	res := C.floatset_to_intset(s._inner)
	return &Set{_inner: res}
}


// FloatspanToIntspan wraps MEOS C function floatspan_to_intspan.
func FloatspanToIntspan(s *Span) *Span {
	res := C.floatspan_to_intspan(s._inner)
	return &Span{_inner: res}
}


// FloatspansetToIntspanset wraps MEOS C function floatspanset_to_intspanset.
func FloatspansetToIntspanset(ss *SpanSet) *SpanSet {
	res := C.floatspanset_to_intspanset(ss._inner)
	return &SpanSet{_inner: res}
}


// IntToSet wraps MEOS C function int_to_set.
func IntToSet(i int) *Set {
	res := C.int_to_set(C.int(i))
	return &Set{_inner: res}
}


// IntToSpan wraps MEOS C function int_to_span.
func IntToSpan(i int) *Span {
	res := C.int_to_span(C.int(i))
	return &Span{_inner: res}
}


// IntToSpanset wraps MEOS C function int_to_spanset.
func IntToSpanset(i int) *SpanSet {
	res := C.int_to_spanset(C.int(i))
	return &SpanSet{_inner: res}
}


// IntsetToFloatset wraps MEOS C function intset_to_floatset.
func IntsetToFloatset(s *Set) *Set {
	res := C.intset_to_floatset(s._inner)
	return &Set{_inner: res}
}


// IntspanToFloatspan wraps MEOS C function intspan_to_floatspan.
func IntspanToFloatspan(s *Span) *Span {
	res := C.intspan_to_floatspan(s._inner)
	return &Span{_inner: res}
}


// IntspansetToFloatspanset wraps MEOS C function intspanset_to_floatspanset.
func IntspansetToFloatspanset(ss *SpanSet) *SpanSet {
	res := C.intspanset_to_floatspanset(ss._inner)
	return &SpanSet{_inner: res}
}


// SetToSpan wraps MEOS C function set_to_span.
func SetToSpan(s *Set) *Span {
	res := C.set_to_span(s._inner)
	return &Span{_inner: res}
}


// SetToSpanset wraps MEOS C function set_to_spanset.
func SetToSpanset(s *Set) *SpanSet {
	res := C.set_to_spanset(s._inner)
	return &SpanSet{_inner: res}
}


// SpanToSpanset wraps MEOS C function span_to_spanset.
func SpanToSpanset(s *Span) *SpanSet {
	res := C.span_to_spanset(s._inner)
	return &SpanSet{_inner: res}
}


// TextToSet wraps MEOS C function text_to_set.
func TextToSet(txt string) *Set {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.text_to_set(_c_txt)
	return &Set{_inner: res}
}


// TimestamptzToSet wraps MEOS C function timestamptz_to_set.
func TimestamptzToSet(t int64) *Set {
	res := C.timestamptz_to_set(C.TimestampTz(t))
	return &Set{_inner: res}
}


// TimestamptzToSpan wraps MEOS C function timestamptz_to_span.
func TimestamptzToSpan(t int64) *Span {
	res := C.timestamptz_to_span(C.TimestampTz(t))
	return &Span{_inner: res}
}


// TimestamptzToSpanset wraps MEOS C function timestamptz_to_spanset.
func TimestamptzToSpanset(t int64) *SpanSet {
	res := C.timestamptz_to_spanset(C.TimestampTz(t))
	return &SpanSet{_inner: res}
}


// TstzsetToDateset wraps MEOS C function tstzset_to_dateset.
func TstzsetToDateset(s *Set) *Set {
	res := C.tstzset_to_dateset(s._inner)
	return &Set{_inner: res}
}


// TstzspanToDatespan wraps MEOS C function tstzspan_to_datespan.
func TstzspanToDatespan(s *Span) *Span {
	res := C.tstzspan_to_datespan(s._inner)
	return &Span{_inner: res}
}


// TstzspansetToDatespanset wraps MEOS C function tstzspanset_to_datespanset.
func TstzspansetToDatespanset(ss *SpanSet) *SpanSet {
	res := C.tstzspanset_to_datespanset(ss._inner)
	return &SpanSet{_inner: res}
}


// BigintsetEndValue wraps MEOS C function bigintset_end_value.
func BigintsetEndValue(s *Set) int64 {
	res := C.bigintset_end_value(s._inner)
	return int64(res)
}


// BigintsetStartValue wraps MEOS C function bigintset_start_value.
func BigintsetStartValue(s *Set) int64 {
	res := C.bigintset_start_value(s._inner)
	return int64(res)
}


// BigintsetValueN wraps MEOS C function bigintset_value_n.
func BigintsetValueN(s *Set, n int) (bool, int64) {
	var _out_result C.int64
	res := C.bigintset_value_n(s._inner, C.int(n), &_out_result)
	return bool(res), int64({})
}


// TODO bigintset_values: unsupported return type int64 *
// func BigintsetValues(...) { /* not yet handled by codegen */ }


// BigintspanLower wraps MEOS C function bigintspan_lower.
func BigintspanLower(s *Span) int64 {
	res := C.bigintspan_lower(s._inner)
	return int64(res)
}


// BigintspanUpper wraps MEOS C function bigintspan_upper.
func BigintspanUpper(s *Span) int64 {
	res := C.bigintspan_upper(s._inner)
	return int64(res)
}


// BigintspanWidth wraps MEOS C function bigintspan_width.
func BigintspanWidth(s *Span) int64 {
	res := C.bigintspan_width(s._inner)
	return int64(res)
}


// BigintspansetLower wraps MEOS C function bigintspanset_lower.
func BigintspansetLower(ss *SpanSet) int64 {
	res := C.bigintspanset_lower(ss._inner)
	return int64(res)
}


// BigintspansetUpper wraps MEOS C function bigintspanset_upper.
func BigintspansetUpper(ss *SpanSet) int64 {
	res := C.bigintspanset_upper(ss._inner)
	return int64(res)
}


// BigintspansetWidth wraps MEOS C function bigintspanset_width.
func BigintspansetWidth(ss *SpanSet, boundspan bool) int64 {
	res := C.bigintspanset_width(ss._inner, C.bool(boundspan))
	return int64(res)
}


// DatesetEndValue wraps MEOS C function dateset_end_value.
func DatesetEndValue(s *Set) int32 {
	res := C.dateset_end_value(s._inner)
	return int32(res)
}


// DatesetStartValue wraps MEOS C function dateset_start_value.
func DatesetStartValue(s *Set) int32 {
	res := C.dateset_start_value(s._inner)
	return int32(res)
}


// DatesetValueN wraps MEOS C function dateset_value_n.
func DatesetValueN(s *Set, n int) (bool, int32) {
	var _out_result C.DateADT
	res := C.dateset_value_n(s._inner, C.int(n), &_out_result)
	return bool(res), int32({})
}


// TODO dateset_values: unsupported return type DateADT *
// func DatesetValues(...) { /* not yet handled by codegen */ }


// DatespanDuration wraps MEOS C function datespan_duration.
func DatespanDuration(s *Span) timeutil.Timedelta {
	res := C.datespan_duration(s._inner)
	return IntervalToTimeDelta(res)
}


// DatespanLower wraps MEOS C function datespan_lower.
func DatespanLower(s *Span) int32 {
	res := C.datespan_lower(s._inner)
	return int32(res)
}


// DatespanUpper wraps MEOS C function datespan_upper.
func DatespanUpper(s *Span) int32 {
	res := C.datespan_upper(s._inner)
	return int32(res)
}


// DatespansetDateN wraps MEOS C function datespanset_date_n.
func DatespansetDateN(ss *SpanSet, n int) (bool, int32) {
	var _out_result C.DateADT
	res := C.datespanset_date_n(ss._inner, C.int(n), &_out_result)
	return bool(res), int32({})
}


// DatespansetDates wraps MEOS C function datespanset_dates.
func DatespansetDates(ss *SpanSet) *Set {
	res := C.datespanset_dates(ss._inner)
	return &Set{_inner: res}
}


// DatespansetDuration wraps MEOS C function datespanset_duration.
func DatespansetDuration(ss *SpanSet, boundspan bool) timeutil.Timedelta {
	res := C.datespanset_duration(ss._inner, C.bool(boundspan))
	return IntervalToTimeDelta(res)
}


// DatespansetEndDate wraps MEOS C function datespanset_end_date.
func DatespansetEndDate(ss *SpanSet) int32 {
	res := C.datespanset_end_date(ss._inner)
	return int32(res)
}


// DatespansetNumDates wraps MEOS C function datespanset_num_dates.
func DatespansetNumDates(ss *SpanSet) int {
	res := C.datespanset_num_dates(ss._inner)
	return int(res)
}


// DatespansetStartDate wraps MEOS C function datespanset_start_date.
func DatespansetStartDate(ss *SpanSet) int32 {
	res := C.datespanset_start_date(ss._inner)
	return int32(res)
}


// FloatsetEndValue wraps MEOS C function floatset_end_value.
func FloatsetEndValue(s *Set) float64 {
	res := C.floatset_end_value(s._inner)
	return float64(res)
}


// FloatsetStartValue wraps MEOS C function floatset_start_value.
func FloatsetStartValue(s *Set) float64 {
	res := C.floatset_start_value(s._inner)
	return float64(res)
}


// FloatsetValueN wraps MEOS C function floatset_value_n.
func FloatsetValueN(s *Set, n int) (bool, float64) {
	var _out_result C.double
	res := C.floatset_value_n(s._inner, C.int(n), &_out_result)
	return bool(res), float64({})
}


// TODO floatset_values: unsupported return type double *
// func FloatsetValues(...) { /* not yet handled by codegen */ }


// FloatspanLower wraps MEOS C function floatspan_lower.
func FloatspanLower(s *Span) float64 {
	res := C.floatspan_lower(s._inner)
	return float64(res)
}


// FloatspanUpper wraps MEOS C function floatspan_upper.
func FloatspanUpper(s *Span) float64 {
	res := C.floatspan_upper(s._inner)
	return float64(res)
}


// FloatspanWidth wraps MEOS C function floatspan_width.
func FloatspanWidth(s *Span) float64 {
	res := C.floatspan_width(s._inner)
	return float64(res)
}


// FloatspansetLower wraps MEOS C function floatspanset_lower.
func FloatspansetLower(ss *SpanSet) float64 {
	res := C.floatspanset_lower(ss._inner)
	return float64(res)
}


// FloatspansetUpper wraps MEOS C function floatspanset_upper.
func FloatspansetUpper(ss *SpanSet) float64 {
	res := C.floatspanset_upper(ss._inner)
	return float64(res)
}


// FloatspansetWidth wraps MEOS C function floatspanset_width.
func FloatspansetWidth(ss *SpanSet, boundspan bool) float64 {
	res := C.floatspanset_width(ss._inner, C.bool(boundspan))
	return float64(res)
}


// IntsetEndValue wraps MEOS C function intset_end_value.
func IntsetEndValue(s *Set) int {
	res := C.intset_end_value(s._inner)
	return int(res)
}


// IntsetStartValue wraps MEOS C function intset_start_value.
func IntsetStartValue(s *Set) int {
	res := C.intset_start_value(s._inner)
	return int(res)
}


// IntsetValueN wraps MEOS C function intset_value_n.
func IntsetValueN(s *Set, n int) (bool, int) {
	var _out_result C.int
	res := C.intset_value_n(s._inner, C.int(n), &_out_result)
	return bool(res), int({})
}


// TODO intset_values: unsupported return type int *
// func IntsetValues(...) { /* not yet handled by codegen */ }


// IntspanLower wraps MEOS C function intspan_lower.
func IntspanLower(s *Span) int {
	res := C.intspan_lower(s._inner)
	return int(res)
}


// IntspanUpper wraps MEOS C function intspan_upper.
func IntspanUpper(s *Span) int {
	res := C.intspan_upper(s._inner)
	return int(res)
}


// IntspanWidth wraps MEOS C function intspan_width.
func IntspanWidth(s *Span) int {
	res := C.intspan_width(s._inner)
	return int(res)
}


// IntspansetLower wraps MEOS C function intspanset_lower.
func IntspansetLower(ss *SpanSet) int {
	res := C.intspanset_lower(ss._inner)
	return int(res)
}


// IntspansetUpper wraps MEOS C function intspanset_upper.
func IntspansetUpper(ss *SpanSet) int {
	res := C.intspanset_upper(ss._inner)
	return int(res)
}


// IntspansetWidth wraps MEOS C function intspanset_width.
func IntspansetWidth(ss *SpanSet, boundspan bool) int {
	res := C.intspanset_width(ss._inner, C.bool(boundspan))
	return int(res)
}


// SetHash wraps MEOS C function set_hash.
func SetHash(s *Set) uint32 {
	res := C.set_hash(s._inner)
	return uint32(res)
}


// SetHashExtended wraps MEOS C function set_hash_extended.
func SetHashExtended(s *Set, seed uint64) uint64 {
	res := C.set_hash_extended(s._inner, C.uint64(seed))
	return uint64(res)
}


// SetNumValues wraps MEOS C function set_num_values.
func SetNumValues(s *Set) int {
	res := C.set_num_values(s._inner)
	return int(res)
}


// SpanHash wraps MEOS C function span_hash.
func SpanHash(s *Span) uint32 {
	res := C.span_hash(s._inner)
	return uint32(res)
}


// SpanHashExtended wraps MEOS C function span_hash_extended.
func SpanHashExtended(s *Span, seed uint64) uint64 {
	res := C.span_hash_extended(s._inner, C.uint64(seed))
	return uint64(res)
}


// SpanLowerInc wraps MEOS C function span_lower_inc.
func SpanLowerInc(s *Span) bool {
	res := C.span_lower_inc(s._inner)
	return bool(res)
}


// SpanUpperInc wraps MEOS C function span_upper_inc.
func SpanUpperInc(s *Span) bool {
	res := C.span_upper_inc(s._inner)
	return bool(res)
}


// SpansetEndSpan wraps MEOS C function spanset_end_span.
func SpansetEndSpan(ss *SpanSet) *Span {
	res := C.spanset_end_span(ss._inner)
	return &Span{_inner: res}
}


// SpansetHash wraps MEOS C function spanset_hash.
func SpansetHash(ss *SpanSet) uint32 {
	res := C.spanset_hash(ss._inner)
	return uint32(res)
}


// SpansetHashExtended wraps MEOS C function spanset_hash_extended.
func SpansetHashExtended(ss *SpanSet, seed uint64) uint64 {
	res := C.spanset_hash_extended(ss._inner, C.uint64(seed))
	return uint64(res)
}


// SpansetLowerInc wraps MEOS C function spanset_lower_inc.
func SpansetLowerInc(ss *SpanSet) bool {
	res := C.spanset_lower_inc(ss._inner)
	return bool(res)
}


// SpansetNumSpans wraps MEOS C function spanset_num_spans.
func SpansetNumSpans(ss *SpanSet) int {
	res := C.spanset_num_spans(ss._inner)
	return int(res)
}


// SpansetSpan wraps MEOS C function spanset_span.
func SpansetSpan(ss *SpanSet) *Span {
	res := C.spanset_span(ss._inner)
	return &Span{_inner: res}
}


// SpansetSpanN wraps MEOS C function spanset_span_n.
func SpansetSpanN(ss *SpanSet, i int) *Span {
	res := C.spanset_span_n(ss._inner, C.int(i))
	return &Span{_inner: res}
}


// TODO spanset_spanarr: unsupported return type Span **
// func SpansetSpanarr(...) { /* not yet handled by codegen */ }


// SpansetStartSpan wraps MEOS C function spanset_start_span.
func SpansetStartSpan(ss *SpanSet) *Span {
	res := C.spanset_start_span(ss._inner)
	return &Span{_inner: res}
}


// SpansetUpperInc wraps MEOS C function spanset_upper_inc.
func SpansetUpperInc(ss *SpanSet) bool {
	res := C.spanset_upper_inc(ss._inner)
	return bool(res)
}


// TextsetEndValue wraps MEOS C function textset_end_value.
func TextsetEndValue(s *Set) string {
	res := C.textset_end_value(s._inner)
	return text2cstring(res)
}


// TextsetStartValue wraps MEOS C function textset_start_value.
func TextsetStartValue(s *Set) string {
	res := C.textset_start_value(s._inner)
	return text2cstring(res)
}


// TextsetValueN wraps MEOS C function textset_value_n.
func TextsetValueN(s *Set, n int) (bool, string) {
	var _out_result *C.text
	res := C.textset_value_n(s._inner, C.int(n), &_out_result)
	return bool(res), text2cstring(_out_result)
}


// TODO textset_values: unsupported return type text **
// func TextsetValues(...) { /* not yet handled by codegen */ }


// TstzsetEndValue wraps MEOS C function tstzset_end_value.
func TstzsetEndValue(s *Set) int64 {
	res := C.tstzset_end_value(s._inner)
	return int64(res)
}


// TstzsetStartValue wraps MEOS C function tstzset_start_value.
func TstzsetStartValue(s *Set) int64 {
	res := C.tstzset_start_value(s._inner)
	return int64(res)
}


// TstzsetValueN wraps MEOS C function tstzset_value_n.
func TstzsetValueN(s *Set, n int) (bool, int64) {
	var _out_result C.TimestampTz
	res := C.tstzset_value_n(s._inner, C.int(n), &_out_result)
	return bool(res), int64({})
}


// TODO tstzset_values: unsupported return type TimestampTz *
// func TstzsetValues(...) { /* not yet handled by codegen */ }


// TstzspanDuration wraps MEOS C function tstzspan_duration.
func TstzspanDuration(s *Span) timeutil.Timedelta {
	res := C.tstzspan_duration(s._inner)
	return IntervalToTimeDelta(res)
}


// TstzspanLower wraps MEOS C function tstzspan_lower.
func TstzspanLower(s *Span) int64 {
	res := C.tstzspan_lower(s._inner)
	return int64(res)
}


// TstzspanUpper wraps MEOS C function tstzspan_upper.
func TstzspanUpper(s *Span) int64 {
	res := C.tstzspan_upper(s._inner)
	return int64(res)
}


// TstzspansetDuration wraps MEOS C function tstzspanset_duration.
func TstzspansetDuration(ss *SpanSet, boundspan bool) timeutil.Timedelta {
	res := C.tstzspanset_duration(ss._inner, C.bool(boundspan))
	return IntervalToTimeDelta(res)
}


// TstzspansetEndTimestamptz wraps MEOS C function tstzspanset_end_timestamptz.
func TstzspansetEndTimestamptz(ss *SpanSet) int64 {
	res := C.tstzspanset_end_timestamptz(ss._inner)
	return int64(res)
}


// TstzspansetLower wraps MEOS C function tstzspanset_lower.
func TstzspansetLower(ss *SpanSet) int64 {
	res := C.tstzspanset_lower(ss._inner)
	return int64(res)
}


// TstzspansetNumTimestamps wraps MEOS C function tstzspanset_num_timestamps.
func TstzspansetNumTimestamps(ss *SpanSet) int {
	res := C.tstzspanset_num_timestamps(ss._inner)
	return int(res)
}


// TstzspansetStartTimestamptz wraps MEOS C function tstzspanset_start_timestamptz.
func TstzspansetStartTimestamptz(ss *SpanSet) int64 {
	res := C.tstzspanset_start_timestamptz(ss._inner)
	return int64(res)
}


// TstzspansetTimestamps wraps MEOS C function tstzspanset_timestamps.
func TstzspansetTimestamps(ss *SpanSet) *Set {
	res := C.tstzspanset_timestamps(ss._inner)
	return &Set{_inner: res}
}


// TstzspansetTimestamptzN wraps MEOS C function tstzspanset_timestamptz_n.
func TstzspansetTimestamptzN(ss *SpanSet, n int) (bool, int64) {
	var _out_result C.TimestampTz
	res := C.tstzspanset_timestamptz_n(ss._inner, C.int(n), &_out_result)
	return bool(res), int64({})
}


// TstzspansetUpper wraps MEOS C function tstzspanset_upper.
func TstzspansetUpper(ss *SpanSet) int64 {
	res := C.tstzspanset_upper(ss._inner)
	return int64(res)
}


// BigintsetShiftScale wraps MEOS C function bigintset_shift_scale.
func BigintsetShiftScale(s *Set, shift int64, width int64, hasshift bool, haswidth bool) *Set {
	res := C.bigintset_shift_scale(s._inner, C.int64(shift), C.int64(width), C.bool(hasshift), C.bool(haswidth))
	return &Set{_inner: res}
}


// BigintspanShiftScale wraps MEOS C function bigintspan_shift_scale.
func BigintspanShiftScale(s *Span, shift int64, width int64, hasshift bool, haswidth bool) *Span {
	res := C.bigintspan_shift_scale(s._inner, C.int64(shift), C.int64(width), C.bool(hasshift), C.bool(haswidth))
	return &Span{_inner: res}
}


// BigintspansetShiftScale wraps MEOS C function bigintspanset_shift_scale.
func BigintspansetShiftScale(ss *SpanSet, shift int64, width int64, hasshift bool, haswidth bool) *SpanSet {
	res := C.bigintspanset_shift_scale(ss._inner, C.int64(shift), C.int64(width), C.bool(hasshift), C.bool(haswidth))
	return &SpanSet{_inner: res}
}


// DatesetShiftScale wraps MEOS C function dateset_shift_scale.
func DatesetShiftScale(s *Set, shift int, width int, hasshift bool, haswidth bool) *Set {
	res := C.dateset_shift_scale(s._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	return &Set{_inner: res}
}


// DatespanShiftScale wraps MEOS C function datespan_shift_scale.
func DatespanShiftScale(s *Span, shift int, width int, hasshift bool, haswidth bool) *Span {
	res := C.datespan_shift_scale(s._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	return &Span{_inner: res}
}


// DatespansetShiftScale wraps MEOS C function datespanset_shift_scale.
func DatespansetShiftScale(ss *SpanSet, shift int, width int, hasshift bool, haswidth bool) *SpanSet {
	res := C.datespanset_shift_scale(ss._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	return &SpanSet{_inner: res}
}


// FloatsetCeil wraps MEOS C function floatset_ceil.
func FloatsetCeil(s *Set) *Set {
	res := C.floatset_ceil(s._inner)
	return &Set{_inner: res}
}


// FloatsetDegrees wraps MEOS C function floatset_degrees.
func FloatsetDegrees(s *Set, normalize bool) *Set {
	res := C.floatset_degrees(s._inner, C.bool(normalize))
	return &Set{_inner: res}
}


// FloatsetFloor wraps MEOS C function floatset_floor.
func FloatsetFloor(s *Set) *Set {
	res := C.floatset_floor(s._inner)
	return &Set{_inner: res}
}


// FloatsetRadians wraps MEOS C function floatset_radians.
func FloatsetRadians(s *Set) *Set {
	res := C.floatset_radians(s._inner)
	return &Set{_inner: res}
}


// FloatsetShiftScale wraps MEOS C function floatset_shift_scale.
func FloatsetShiftScale(s *Set, shift float64, width float64, hasshift bool, haswidth bool) *Set {
	res := C.floatset_shift_scale(s._inner, C.double(shift), C.double(width), C.bool(hasshift), C.bool(haswidth))
	return &Set{_inner: res}
}


// FloatspanCeil wraps MEOS C function floatspan_ceil.
func FloatspanCeil(s *Span) *Span {
	res := C.floatspan_ceil(s._inner)
	return &Span{_inner: res}
}


// FloatspanDegrees wraps MEOS C function floatspan_degrees.
func FloatspanDegrees(s *Span, normalize bool) *Span {
	res := C.floatspan_degrees(s._inner, C.bool(normalize))
	return &Span{_inner: res}
}


// FloatspanFloor wraps MEOS C function floatspan_floor.
func FloatspanFloor(s *Span) *Span {
	res := C.floatspan_floor(s._inner)
	return &Span{_inner: res}
}


// FloatspanRadians wraps MEOS C function floatspan_radians.
func FloatspanRadians(s *Span) *Span {
	res := C.floatspan_radians(s._inner)
	return &Span{_inner: res}
}


// FloatspanRound wraps MEOS C function floatspan_round.
func FloatspanRound(s *Span, maxdd int) *Span {
	res := C.floatspan_round(s._inner, C.int(maxdd))
	return &Span{_inner: res}
}


// FloatspanShiftScale wraps MEOS C function floatspan_shift_scale.
func FloatspanShiftScale(s *Span, shift float64, width float64, hasshift bool, haswidth bool) *Span {
	res := C.floatspan_shift_scale(s._inner, C.double(shift), C.double(width), C.bool(hasshift), C.bool(haswidth))
	return &Span{_inner: res}
}


// FloatspansetCeil wraps MEOS C function floatspanset_ceil.
func FloatspansetCeil(ss *SpanSet) *SpanSet {
	res := C.floatspanset_ceil(ss._inner)
	return &SpanSet{_inner: res}
}


// FloatspansetFloor wraps MEOS C function floatspanset_floor.
func FloatspansetFloor(ss *SpanSet) *SpanSet {
	res := C.floatspanset_floor(ss._inner)
	return &SpanSet{_inner: res}
}


// FloatspansetDegrees wraps MEOS C function floatspanset_degrees.
func FloatspansetDegrees(ss *SpanSet, normalize bool) *SpanSet {
	res := C.floatspanset_degrees(ss._inner, C.bool(normalize))
	return &SpanSet{_inner: res}
}


// FloatspansetRadians wraps MEOS C function floatspanset_radians.
func FloatspansetRadians(ss *SpanSet) *SpanSet {
	res := C.floatspanset_radians(ss._inner)
	return &SpanSet{_inner: res}
}


// FloatspansetRound wraps MEOS C function floatspanset_round.
func FloatspansetRound(ss *SpanSet, maxdd int) *SpanSet {
	res := C.floatspanset_round(ss._inner, C.int(maxdd))
	return &SpanSet{_inner: res}
}


// FloatspansetShiftScale wraps MEOS C function floatspanset_shift_scale.
func FloatspansetShiftScale(ss *SpanSet, shift float64, width float64, hasshift bool, haswidth bool) *SpanSet {
	res := C.floatspanset_shift_scale(ss._inner, C.double(shift), C.double(width), C.bool(hasshift), C.bool(haswidth))
	return &SpanSet{_inner: res}
}


// IntsetShiftScale wraps MEOS C function intset_shift_scale.
func IntsetShiftScale(s *Set, shift int, width int, hasshift bool, haswidth bool) *Set {
	res := C.intset_shift_scale(s._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	return &Set{_inner: res}
}


// IntspanShiftScale wraps MEOS C function intspan_shift_scale.
func IntspanShiftScale(s *Span, shift int, width int, hasshift bool, haswidth bool) *Span {
	res := C.intspan_shift_scale(s._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	return &Span{_inner: res}
}


// IntspansetShiftScale wraps MEOS C function intspanset_shift_scale.
func IntspansetShiftScale(ss *SpanSet, shift int, width int, hasshift bool, haswidth bool) *SpanSet {
	res := C.intspanset_shift_scale(ss._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	return &SpanSet{_inner: res}
}


// TstzspanExpand wraps MEOS C function tstzspan_expand.
func TstzspanExpand(s *Span, interv timeutil.Timedelta) *Span {
	res := C.tstzspan_expand(s._inner, interv.Inner())
	return &Span{_inner: res}
}


// SetRound wraps MEOS C function set_round.
func SetRound(s *Set, maxdd int) *Set {
	res := C.set_round(s._inner, C.int(maxdd))
	return &Set{_inner: res}
}


// TextcatTextTextset wraps MEOS C function textcat_text_textset.
func TextcatTextTextset(txt string, s *Set) *Set {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.textcat_text_textset(_c_txt, s._inner)
	return &Set{_inner: res}
}


// TextcatTextsetText wraps MEOS C function textcat_textset_text.
func TextcatTextsetText(s *Set, txt string) *Set {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.textcat_textset_text(s._inner, _c_txt)
	return &Set{_inner: res}
}


// TextsetInitcap wraps MEOS C function textset_initcap.
func TextsetInitcap(s *Set) *Set {
	res := C.textset_initcap(s._inner)
	return &Set{_inner: res}
}


// TextsetLower wraps MEOS C function textset_lower.
func TextsetLower(s *Set) *Set {
	res := C.textset_lower(s._inner)
	return &Set{_inner: res}
}


// TextsetUpper wraps MEOS C function textset_upper.
func TextsetUpper(s *Set) *Set {
	res := C.textset_upper(s._inner)
	return &Set{_inner: res}
}


// TimestamptzTprecision wraps MEOS C function timestamptz_tprecision.
func TimestamptzTprecision(t int64, duration timeutil.Timedelta, torigin int64) int64 {
	res := C.timestamptz_tprecision(C.TimestampTz(t), duration.Inner(), C.TimestampTz(torigin))
	return int64(res)
}


// TstzsetShiftScale wraps MEOS C function tstzset_shift_scale.
func TstzsetShiftScale(s *Set, shift timeutil.Timedelta, duration timeutil.Timedelta) *Set {
	res := C.tstzset_shift_scale(s._inner, shift.Inner(), duration.Inner())
	return &Set{_inner: res}
}


// TstzsetTprecision wraps MEOS C function tstzset_tprecision.
func TstzsetTprecision(s *Set, duration timeutil.Timedelta, torigin int64) *Set {
	res := C.tstzset_tprecision(s._inner, duration.Inner(), C.TimestampTz(torigin))
	return &Set{_inner: res}
}


// TstzspanShiftScale wraps MEOS C function tstzspan_shift_scale.
func TstzspanShiftScale(s *Span, shift timeutil.Timedelta, duration timeutil.Timedelta) *Span {
	res := C.tstzspan_shift_scale(s._inner, shift.Inner(), duration.Inner())
	return &Span{_inner: res}
}


// TstzspanTprecision wraps MEOS C function tstzspan_tprecision.
func TstzspanTprecision(s *Span, duration timeutil.Timedelta, torigin int64) *Span {
	res := C.tstzspan_tprecision(s._inner, duration.Inner(), C.TimestampTz(torigin))
	return &Span{_inner: res}
}


// TstzspansetShiftScale wraps MEOS C function tstzspanset_shift_scale.
func TstzspansetShiftScale(ss *SpanSet, shift timeutil.Timedelta, duration timeutil.Timedelta) *SpanSet {
	res := C.tstzspanset_shift_scale(ss._inner, shift.Inner(), duration.Inner())
	return &SpanSet{_inner: res}
}


// TstzspansetTprecision wraps MEOS C function tstzspanset_tprecision.
func TstzspansetTprecision(ss *SpanSet, duration timeutil.Timedelta, torigin int64) *SpanSet {
	res := C.tstzspanset_tprecision(ss._inner, duration.Inner(), C.TimestampTz(torigin))
	return &SpanSet{_inner: res}
}


// SetCmp wraps MEOS C function set_cmp.
func SetCmp(s1 *Set, s2 *Set) int {
	res := C.set_cmp(s1._inner, s2._inner)
	return int(res)
}


// SetEq wraps MEOS C function set_eq.
func SetEq(s1 *Set, s2 *Set) bool {
	res := C.set_eq(s1._inner, s2._inner)
	return bool(res)
}


// SetGe wraps MEOS C function set_ge.
func SetGe(s1 *Set, s2 *Set) bool {
	res := C.set_ge(s1._inner, s2._inner)
	return bool(res)
}


// SetGt wraps MEOS C function set_gt.
func SetGt(s1 *Set, s2 *Set) bool {
	res := C.set_gt(s1._inner, s2._inner)
	return bool(res)
}


// SetLe wraps MEOS C function set_le.
func SetLe(s1 *Set, s2 *Set) bool {
	res := C.set_le(s1._inner, s2._inner)
	return bool(res)
}


// SetLt wraps MEOS C function set_lt.
func SetLt(s1 *Set, s2 *Set) bool {
	res := C.set_lt(s1._inner, s2._inner)
	return bool(res)
}


// SetNe wraps MEOS C function set_ne.
func SetNe(s1 *Set, s2 *Set) bool {
	res := C.set_ne(s1._inner, s2._inner)
	return bool(res)
}


// SpanCmp wraps MEOS C function span_cmp.
func SpanCmp(s1 *Span, s2 *Span) int {
	res := C.span_cmp(s1._inner, s2._inner)
	return int(res)
}


// SpanEq wraps MEOS C function span_eq.
func SpanEq(s1 *Span, s2 *Span) bool {
	res := C.span_eq(s1._inner, s2._inner)
	return bool(res)
}


// SpanGe wraps MEOS C function span_ge.
func SpanGe(s1 *Span, s2 *Span) bool {
	res := C.span_ge(s1._inner, s2._inner)
	return bool(res)
}


// SpanGt wraps MEOS C function span_gt.
func SpanGt(s1 *Span, s2 *Span) bool {
	res := C.span_gt(s1._inner, s2._inner)
	return bool(res)
}


// SpanLe wraps MEOS C function span_le.
func SpanLe(s1 *Span, s2 *Span) bool {
	res := C.span_le(s1._inner, s2._inner)
	return bool(res)
}


// SpanLt wraps MEOS C function span_lt.
func SpanLt(s1 *Span, s2 *Span) bool {
	res := C.span_lt(s1._inner, s2._inner)
	return bool(res)
}


// SpanNe wraps MEOS C function span_ne.
func SpanNe(s1 *Span, s2 *Span) bool {
	res := C.span_ne(s1._inner, s2._inner)
	return bool(res)
}


// SpansetCmp wraps MEOS C function spanset_cmp.
func SpansetCmp(ss1 *SpanSet, ss2 *SpanSet) int {
	res := C.spanset_cmp(ss1._inner, ss2._inner)
	return int(res)
}


// SpansetEq wraps MEOS C function spanset_eq.
func SpansetEq(ss1 *SpanSet, ss2 *SpanSet) bool {
	res := C.spanset_eq(ss1._inner, ss2._inner)
	return bool(res)
}


// SpansetGe wraps MEOS C function spanset_ge.
func SpansetGe(ss1 *SpanSet, ss2 *SpanSet) bool {
	res := C.spanset_ge(ss1._inner, ss2._inner)
	return bool(res)
}


// SpansetGt wraps MEOS C function spanset_gt.
func SpansetGt(ss1 *SpanSet, ss2 *SpanSet) bool {
	res := C.spanset_gt(ss1._inner, ss2._inner)
	return bool(res)
}


// SpansetLe wraps MEOS C function spanset_le.
func SpansetLe(ss1 *SpanSet, ss2 *SpanSet) bool {
	res := C.spanset_le(ss1._inner, ss2._inner)
	return bool(res)
}


// SpansetLt wraps MEOS C function spanset_lt.
func SpansetLt(ss1 *SpanSet, ss2 *SpanSet) bool {
	res := C.spanset_lt(ss1._inner, ss2._inner)
	return bool(res)
}


// SpansetNe wraps MEOS C function spanset_ne.
func SpansetNe(ss1 *SpanSet, ss2 *SpanSet) bool {
	res := C.spanset_ne(ss1._inner, ss2._inner)
	return bool(res)
}


// SetSpans wraps MEOS C function set_spans.
func SetSpans(s *Set) *Span {
	res := C.set_spans(s._inner)
	return &Span{_inner: res}
}


// SetSplitEachNSpans wraps MEOS C function set_split_each_n_spans.
func SetSplitEachNSpans(s *Set, elems_per_span int) (*Span, int) {
	var _out_count C.int
	res := C.set_split_each_n_spans(s._inner, C.int(elems_per_span), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// SetSplitNSpans wraps MEOS C function set_split_n_spans.
func SetSplitNSpans(s *Set, span_count int) (*Span, int) {
	var _out_count C.int
	res := C.set_split_n_spans(s._inner, C.int(span_count), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// SpansetSpans wraps MEOS C function spanset_spans.
func SpansetSpans(ss *SpanSet) *Span {
	res := C.spanset_spans(ss._inner)
	return &Span{_inner: res}
}


// SpansetSplitEachNSpans wraps MEOS C function spanset_split_each_n_spans.
func SpansetSplitEachNSpans(ss *SpanSet, elems_per_span int) (*Span, int) {
	var _out_count C.int
	res := C.spanset_split_each_n_spans(ss._inner, C.int(elems_per_span), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// SpansetSplitNSpans wraps MEOS C function spanset_split_n_spans.
func SpansetSplitNSpans(ss *SpanSet, span_count int) (*Span, int) {
	var _out_count C.int
	res := C.spanset_split_n_spans(ss._inner, C.int(span_count), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// AdjacentSpanBigint wraps MEOS C function adjacent_span_bigint.
func AdjacentSpanBigint(s *Span, i int64) bool {
	res := C.adjacent_span_bigint(s._inner, C.int64(i))
	return bool(res)
}


// AdjacentSpanDate wraps MEOS C function adjacent_span_date.
func AdjacentSpanDate(s *Span, d int32) bool {
	res := C.adjacent_span_date(s._inner, C.DateADT(d))
	return bool(res)
}


// AdjacentSpanFloat wraps MEOS C function adjacent_span_float.
func AdjacentSpanFloat(s *Span, d float64) bool {
	res := C.adjacent_span_float(s._inner, C.double(d))
	return bool(res)
}


// AdjacentSpanInt wraps MEOS C function adjacent_span_int.
func AdjacentSpanInt(s *Span, i int) bool {
	res := C.adjacent_span_int(s._inner, C.int(i))
	return bool(res)
}


// AdjacentSpanSpan wraps MEOS C function adjacent_span_span.
func AdjacentSpanSpan(s1 *Span, s2 *Span) bool {
	res := C.adjacent_span_span(s1._inner, s2._inner)
	return bool(res)
}


// AdjacentSpanSpanset wraps MEOS C function adjacent_span_spanset.
func AdjacentSpanSpanset(s *Span, ss *SpanSet) bool {
	res := C.adjacent_span_spanset(s._inner, ss._inner)
	return bool(res)
}


// AdjacentSpanTimestamptz wraps MEOS C function adjacent_span_timestamptz.
func AdjacentSpanTimestamptz(s *Span, t int64) bool {
	res := C.adjacent_span_timestamptz(s._inner, C.TimestampTz(t))
	return bool(res)
}


// AdjacentSpansetBigint wraps MEOS C function adjacent_spanset_bigint.
func AdjacentSpansetBigint(ss *SpanSet, i int64) bool {
	res := C.adjacent_spanset_bigint(ss._inner, C.int64(i))
	return bool(res)
}


// AdjacentSpansetDate wraps MEOS C function adjacent_spanset_date.
func AdjacentSpansetDate(ss *SpanSet, d int32) bool {
	res := C.adjacent_spanset_date(ss._inner, C.DateADT(d))
	return bool(res)
}


// AdjacentSpansetFloat wraps MEOS C function adjacent_spanset_float.
func AdjacentSpansetFloat(ss *SpanSet, d float64) bool {
	res := C.adjacent_spanset_float(ss._inner, C.double(d))
	return bool(res)
}


// AdjacentSpansetInt wraps MEOS C function adjacent_spanset_int.
func AdjacentSpansetInt(ss *SpanSet, i int) bool {
	res := C.adjacent_spanset_int(ss._inner, C.int(i))
	return bool(res)
}


// AdjacentSpansetTimestamptz wraps MEOS C function adjacent_spanset_timestamptz.
func AdjacentSpansetTimestamptz(ss *SpanSet, t int64) bool {
	res := C.adjacent_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return bool(res)
}


// AdjacentSpansetSpan wraps MEOS C function adjacent_spanset_span.
func AdjacentSpansetSpan(ss *SpanSet, s *Span) bool {
	res := C.adjacent_spanset_span(ss._inner, s._inner)
	return bool(res)
}


// AdjacentSpansetSpanset wraps MEOS C function adjacent_spanset_spanset.
func AdjacentSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) bool {
	res := C.adjacent_spanset_spanset(ss1._inner, ss2._inner)
	return bool(res)
}


// ContainedBigintSet wraps MEOS C function contained_bigint_set.
func ContainedBigintSet(i int64, s *Set) bool {
	res := C.contained_bigint_set(C.int64(i), s._inner)
	return bool(res)
}


// ContainedBigintSpan wraps MEOS C function contained_bigint_span.
func ContainedBigintSpan(i int64, s *Span) bool {
	res := C.contained_bigint_span(C.int64(i), s._inner)
	return bool(res)
}


// ContainedBigintSpanset wraps MEOS C function contained_bigint_spanset.
func ContainedBigintSpanset(i int64, ss *SpanSet) bool {
	res := C.contained_bigint_spanset(C.int64(i), ss._inner)
	return bool(res)
}


// ContainedDateSet wraps MEOS C function contained_date_set.
func ContainedDateSet(d int32, s *Set) bool {
	res := C.contained_date_set(C.DateADT(d), s._inner)
	return bool(res)
}


// ContainedDateSpan wraps MEOS C function contained_date_span.
func ContainedDateSpan(d int32, s *Span) bool {
	res := C.contained_date_span(C.DateADT(d), s._inner)
	return bool(res)
}


// ContainedDateSpanset wraps MEOS C function contained_date_spanset.
func ContainedDateSpanset(d int32, ss *SpanSet) bool {
	res := C.contained_date_spanset(C.DateADT(d), ss._inner)
	return bool(res)
}


// ContainedFloatSet wraps MEOS C function contained_float_set.
func ContainedFloatSet(d float64, s *Set) bool {
	res := C.contained_float_set(C.double(d), s._inner)
	return bool(res)
}


// ContainedFloatSpan wraps MEOS C function contained_float_span.
func ContainedFloatSpan(d float64, s *Span) bool {
	res := C.contained_float_span(C.double(d), s._inner)
	return bool(res)
}


// ContainedFloatSpanset wraps MEOS C function contained_float_spanset.
func ContainedFloatSpanset(d float64, ss *SpanSet) bool {
	res := C.contained_float_spanset(C.double(d), ss._inner)
	return bool(res)
}


// ContainedIntSet wraps MEOS C function contained_int_set.
func ContainedIntSet(i int, s *Set) bool {
	res := C.contained_int_set(C.int(i), s._inner)
	return bool(res)
}


// ContainedIntSpan wraps MEOS C function contained_int_span.
func ContainedIntSpan(i int, s *Span) bool {
	res := C.contained_int_span(C.int(i), s._inner)
	return bool(res)
}


// ContainedIntSpanset wraps MEOS C function contained_int_spanset.
func ContainedIntSpanset(i int, ss *SpanSet) bool {
	res := C.contained_int_spanset(C.int(i), ss._inner)
	return bool(res)
}


// ContainedSetSet wraps MEOS C function contained_set_set.
func ContainedSetSet(s1 *Set, s2 *Set) bool {
	res := C.contained_set_set(s1._inner, s2._inner)
	return bool(res)
}


// ContainedSpanSpan wraps MEOS C function contained_span_span.
func ContainedSpanSpan(s1 *Span, s2 *Span) bool {
	res := C.contained_span_span(s1._inner, s2._inner)
	return bool(res)
}


// ContainedSpanSpanset wraps MEOS C function contained_span_spanset.
func ContainedSpanSpanset(s *Span, ss *SpanSet) bool {
	res := C.contained_span_spanset(s._inner, ss._inner)
	return bool(res)
}


// ContainedSpansetSpan wraps MEOS C function contained_spanset_span.
func ContainedSpansetSpan(ss *SpanSet, s *Span) bool {
	res := C.contained_spanset_span(ss._inner, s._inner)
	return bool(res)
}


// ContainedSpansetSpanset wraps MEOS C function contained_spanset_spanset.
func ContainedSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) bool {
	res := C.contained_spanset_spanset(ss1._inner, ss2._inner)
	return bool(res)
}


// ContainedTextSet wraps MEOS C function contained_text_set.
func ContainedTextSet(txt string, s *Set) bool {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.contained_text_set(_c_txt, s._inner)
	return bool(res)
}


// ContainedTimestamptzSet wraps MEOS C function contained_timestamptz_set.
func ContainedTimestamptzSet(t int64, s *Set) bool {
	res := C.contained_timestamptz_set(C.TimestampTz(t), s._inner)
	return bool(res)
}


// ContainedTimestamptzSpan wraps MEOS C function contained_timestamptz_span.
func ContainedTimestamptzSpan(t int64, s *Span) bool {
	res := C.contained_timestamptz_span(C.TimestampTz(t), s._inner)
	return bool(res)
}


// ContainedTimestamptzSpanset wraps MEOS C function contained_timestamptz_spanset.
func ContainedTimestamptzSpanset(t int64, ss *SpanSet) bool {
	res := C.contained_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	return bool(res)
}


// ContainsSetBigint wraps MEOS C function contains_set_bigint.
func ContainsSetBigint(s *Set, i int64) bool {
	res := C.contains_set_bigint(s._inner, C.int64(i))
	return bool(res)
}


// ContainsSetDate wraps MEOS C function contains_set_date.
func ContainsSetDate(s *Set, d int32) bool {
	res := C.contains_set_date(s._inner, C.DateADT(d))
	return bool(res)
}


// ContainsSetFloat wraps MEOS C function contains_set_float.
func ContainsSetFloat(s *Set, d float64) bool {
	res := C.contains_set_float(s._inner, C.double(d))
	return bool(res)
}


// ContainsSetInt wraps MEOS C function contains_set_int.
func ContainsSetInt(s *Set, i int) bool {
	res := C.contains_set_int(s._inner, C.int(i))
	return bool(res)
}


// ContainsSetSet wraps MEOS C function contains_set_set.
func ContainsSetSet(s1 *Set, s2 *Set) bool {
	res := C.contains_set_set(s1._inner, s2._inner)
	return bool(res)
}


// ContainsSetText wraps MEOS C function contains_set_text.
func ContainsSetText(s *Set, t string) bool {
	_c_t := cstring2text(t)
	defer C.free(unsafe.Pointer(_c_t))
	res := C.contains_set_text(s._inner, _c_t)
	return bool(res)
}


// ContainsSetTimestamptz wraps MEOS C function contains_set_timestamptz.
func ContainsSetTimestamptz(s *Set, t int64) bool {
	res := C.contains_set_timestamptz(s._inner, C.TimestampTz(t))
	return bool(res)
}


// ContainsSpanBigint wraps MEOS C function contains_span_bigint.
func ContainsSpanBigint(s *Span, i int64) bool {
	res := C.contains_span_bigint(s._inner, C.int64(i))
	return bool(res)
}


// ContainsSpanDate wraps MEOS C function contains_span_date.
func ContainsSpanDate(s *Span, d int32) bool {
	res := C.contains_span_date(s._inner, C.DateADT(d))
	return bool(res)
}


// ContainsSpanFloat wraps MEOS C function contains_span_float.
func ContainsSpanFloat(s *Span, d float64) bool {
	res := C.contains_span_float(s._inner, C.double(d))
	return bool(res)
}


// ContainsSpanInt wraps MEOS C function contains_span_int.
func ContainsSpanInt(s *Span, i int) bool {
	res := C.contains_span_int(s._inner, C.int(i))
	return bool(res)
}


// ContainsSpanSpan wraps MEOS C function contains_span_span.
func ContainsSpanSpan(s1 *Span, s2 *Span) bool {
	res := C.contains_span_span(s1._inner, s2._inner)
	return bool(res)
}


// ContainsSpanSpanset wraps MEOS C function contains_span_spanset.
func ContainsSpanSpanset(s *Span, ss *SpanSet) bool {
	res := C.contains_span_spanset(s._inner, ss._inner)
	return bool(res)
}


// ContainsSpanTimestamptz wraps MEOS C function contains_span_timestamptz.
func ContainsSpanTimestamptz(s *Span, t int64) bool {
	res := C.contains_span_timestamptz(s._inner, C.TimestampTz(t))
	return bool(res)
}


// ContainsSpansetBigint wraps MEOS C function contains_spanset_bigint.
func ContainsSpansetBigint(ss *SpanSet, i int64) bool {
	res := C.contains_spanset_bigint(ss._inner, C.int64(i))
	return bool(res)
}


// ContainsSpansetDate wraps MEOS C function contains_spanset_date.
func ContainsSpansetDate(ss *SpanSet, d int32) bool {
	res := C.contains_spanset_date(ss._inner, C.DateADT(d))
	return bool(res)
}


// ContainsSpansetFloat wraps MEOS C function contains_spanset_float.
func ContainsSpansetFloat(ss *SpanSet, d float64) bool {
	res := C.contains_spanset_float(ss._inner, C.double(d))
	return bool(res)
}


// ContainsSpansetInt wraps MEOS C function contains_spanset_int.
func ContainsSpansetInt(ss *SpanSet, i int) bool {
	res := C.contains_spanset_int(ss._inner, C.int(i))
	return bool(res)
}


// ContainsSpansetSpan wraps MEOS C function contains_spanset_span.
func ContainsSpansetSpan(ss *SpanSet, s *Span) bool {
	res := C.contains_spanset_span(ss._inner, s._inner)
	return bool(res)
}


// ContainsSpansetSpanset wraps MEOS C function contains_spanset_spanset.
func ContainsSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) bool {
	res := C.contains_spanset_spanset(ss1._inner, ss2._inner)
	return bool(res)
}


// ContainsSpansetTimestamptz wraps MEOS C function contains_spanset_timestamptz.
func ContainsSpansetTimestamptz(ss *SpanSet, t int64) bool {
	res := C.contains_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return bool(res)
}


// OverlapsSetSet wraps MEOS C function overlaps_set_set.
func OverlapsSetSet(s1 *Set, s2 *Set) bool {
	res := C.overlaps_set_set(s1._inner, s2._inner)
	return bool(res)
}


// OverlapsSpanSpan wraps MEOS C function overlaps_span_span.
func OverlapsSpanSpan(s1 *Span, s2 *Span) bool {
	res := C.overlaps_span_span(s1._inner, s2._inner)
	return bool(res)
}


// OverlapsSpanSpanset wraps MEOS C function overlaps_span_spanset.
func OverlapsSpanSpanset(s *Span, ss *SpanSet) bool {
	res := C.overlaps_span_spanset(s._inner, ss._inner)
	return bool(res)
}


// OverlapsSpansetSpan wraps MEOS C function overlaps_spanset_span.
func OverlapsSpansetSpan(ss *SpanSet, s *Span) bool {
	res := C.overlaps_spanset_span(ss._inner, s._inner)
	return bool(res)
}


// OverlapsSpansetSpanset wraps MEOS C function overlaps_spanset_spanset.
func OverlapsSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) bool {
	res := C.overlaps_spanset_spanset(ss1._inner, ss2._inner)
	return bool(res)
}


// AfterDateSet wraps MEOS C function after_date_set.
func AfterDateSet(d int32, s *Set) bool {
	res := C.after_date_set(C.DateADT(d), s._inner)
	return bool(res)
}


// AfterDateSpan wraps MEOS C function after_date_span.
func AfterDateSpan(d int32, s *Span) bool {
	res := C.after_date_span(C.DateADT(d), s._inner)
	return bool(res)
}


// AfterDateSpanset wraps MEOS C function after_date_spanset.
func AfterDateSpanset(d int32, ss *SpanSet) bool {
	res := C.after_date_spanset(C.DateADT(d), ss._inner)
	return bool(res)
}


// AfterSetDate wraps MEOS C function after_set_date.
func AfterSetDate(s *Set, d int32) bool {
	res := C.after_set_date(s._inner, C.DateADT(d))
	return bool(res)
}


// AfterSetTimestamptz wraps MEOS C function after_set_timestamptz.
func AfterSetTimestamptz(s *Set, t int64) bool {
	res := C.after_set_timestamptz(s._inner, C.TimestampTz(t))
	return bool(res)
}


// AfterSpanDate wraps MEOS C function after_span_date.
func AfterSpanDate(s *Span, d int32) bool {
	res := C.after_span_date(s._inner, C.DateADT(d))
	return bool(res)
}


// AfterSpanTimestamptz wraps MEOS C function after_span_timestamptz.
func AfterSpanTimestamptz(s *Span, t int64) bool {
	res := C.after_span_timestamptz(s._inner, C.TimestampTz(t))
	return bool(res)
}


// AfterSpansetDate wraps MEOS C function after_spanset_date.
func AfterSpansetDate(ss *SpanSet, d int32) bool {
	res := C.after_spanset_date(ss._inner, C.DateADT(d))
	return bool(res)
}


// AfterSpansetTimestamptz wraps MEOS C function after_spanset_timestamptz.
func AfterSpansetTimestamptz(ss *SpanSet, t int64) bool {
	res := C.after_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return bool(res)
}


// AfterTimestamptzSet wraps MEOS C function after_timestamptz_set.
func AfterTimestamptzSet(t int64, s *Set) bool {
	res := C.after_timestamptz_set(C.TimestampTz(t), s._inner)
	return bool(res)
}


// AfterTimestamptzSpan wraps MEOS C function after_timestamptz_span.
func AfterTimestamptzSpan(t int64, s *Span) bool {
	res := C.after_timestamptz_span(C.TimestampTz(t), s._inner)
	return bool(res)
}


// AfterTimestamptzSpanset wraps MEOS C function after_timestamptz_spanset.
func AfterTimestamptzSpanset(t int64, ss *SpanSet) bool {
	res := C.after_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	return bool(res)
}


// BeforeDateSet wraps MEOS C function before_date_set.
func BeforeDateSet(d int32, s *Set) bool {
	res := C.before_date_set(C.DateADT(d), s._inner)
	return bool(res)
}


// BeforeDateSpan wraps MEOS C function before_date_span.
func BeforeDateSpan(d int32, s *Span) bool {
	res := C.before_date_span(C.DateADT(d), s._inner)
	return bool(res)
}


// BeforeDateSpanset wraps MEOS C function before_date_spanset.
func BeforeDateSpanset(d int32, ss *SpanSet) bool {
	res := C.before_date_spanset(C.DateADT(d), ss._inner)
	return bool(res)
}


// BeforeSetDate wraps MEOS C function before_set_date.
func BeforeSetDate(s *Set, d int32) bool {
	res := C.before_set_date(s._inner, C.DateADT(d))
	return bool(res)
}


// BeforeSetTimestamptz wraps MEOS C function before_set_timestamptz.
func BeforeSetTimestamptz(s *Set, t int64) bool {
	res := C.before_set_timestamptz(s._inner, C.TimestampTz(t))
	return bool(res)
}


// BeforeSpanDate wraps MEOS C function before_span_date.
func BeforeSpanDate(s *Span, d int32) bool {
	res := C.before_span_date(s._inner, C.DateADT(d))
	return bool(res)
}


// BeforeSpanTimestamptz wraps MEOS C function before_span_timestamptz.
func BeforeSpanTimestamptz(s *Span, t int64) bool {
	res := C.before_span_timestamptz(s._inner, C.TimestampTz(t))
	return bool(res)
}


// BeforeSpansetDate wraps MEOS C function before_spanset_date.
func BeforeSpansetDate(ss *SpanSet, d int32) bool {
	res := C.before_spanset_date(ss._inner, C.DateADT(d))
	return bool(res)
}


// BeforeSpansetTimestamptz wraps MEOS C function before_spanset_timestamptz.
func BeforeSpansetTimestamptz(ss *SpanSet, t int64) bool {
	res := C.before_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return bool(res)
}


// BeforeTimestamptzSet wraps MEOS C function before_timestamptz_set.
func BeforeTimestamptzSet(t int64, s *Set) bool {
	res := C.before_timestamptz_set(C.TimestampTz(t), s._inner)
	return bool(res)
}


// BeforeTimestamptzSpan wraps MEOS C function before_timestamptz_span.
func BeforeTimestamptzSpan(t int64, s *Span) bool {
	res := C.before_timestamptz_span(C.TimestampTz(t), s._inner)
	return bool(res)
}


// BeforeTimestamptzSpanset wraps MEOS C function before_timestamptz_spanset.
func BeforeTimestamptzSpanset(t int64, ss *SpanSet) bool {
	res := C.before_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	return bool(res)
}


// LeftBigintSet wraps MEOS C function left_bigint_set.
func LeftBigintSet(i int64, s *Set) bool {
	res := C.left_bigint_set(C.int64(i), s._inner)
	return bool(res)
}


// LeftBigintSpan wraps MEOS C function left_bigint_span.
func LeftBigintSpan(i int64, s *Span) bool {
	res := C.left_bigint_span(C.int64(i), s._inner)
	return bool(res)
}


// LeftBigintSpanset wraps MEOS C function left_bigint_spanset.
func LeftBigintSpanset(i int64, ss *SpanSet) bool {
	res := C.left_bigint_spanset(C.int64(i), ss._inner)
	return bool(res)
}


// LeftFloatSet wraps MEOS C function left_float_set.
func LeftFloatSet(d float64, s *Set) bool {
	res := C.left_float_set(C.double(d), s._inner)
	return bool(res)
}


// LeftFloatSpan wraps MEOS C function left_float_span.
func LeftFloatSpan(d float64, s *Span) bool {
	res := C.left_float_span(C.double(d), s._inner)
	return bool(res)
}


// LeftFloatSpanset wraps MEOS C function left_float_spanset.
func LeftFloatSpanset(d float64, ss *SpanSet) bool {
	res := C.left_float_spanset(C.double(d), ss._inner)
	return bool(res)
}


// LeftIntSet wraps MEOS C function left_int_set.
func LeftIntSet(i int, s *Set) bool {
	res := C.left_int_set(C.int(i), s._inner)
	return bool(res)
}


// LeftIntSpan wraps MEOS C function left_int_span.
func LeftIntSpan(i int, s *Span) bool {
	res := C.left_int_span(C.int(i), s._inner)
	return bool(res)
}


// LeftIntSpanset wraps MEOS C function left_int_spanset.
func LeftIntSpanset(i int, ss *SpanSet) bool {
	res := C.left_int_spanset(C.int(i), ss._inner)
	return bool(res)
}


// LeftSetBigint wraps MEOS C function left_set_bigint.
func LeftSetBigint(s *Set, i int64) bool {
	res := C.left_set_bigint(s._inner, C.int64(i))
	return bool(res)
}


// LeftSetFloat wraps MEOS C function left_set_float.
func LeftSetFloat(s *Set, d float64) bool {
	res := C.left_set_float(s._inner, C.double(d))
	return bool(res)
}


// LeftSetInt wraps MEOS C function left_set_int.
func LeftSetInt(s *Set, i int) bool {
	res := C.left_set_int(s._inner, C.int(i))
	return bool(res)
}


// LeftSetSet wraps MEOS C function left_set_set.
func LeftSetSet(s1 *Set, s2 *Set) bool {
	res := C.left_set_set(s1._inner, s2._inner)
	return bool(res)
}


// LeftSetText wraps MEOS C function left_set_text.
func LeftSetText(s *Set, txt string) bool {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.left_set_text(s._inner, _c_txt)
	return bool(res)
}


// LeftSpanBigint wraps MEOS C function left_span_bigint.
func LeftSpanBigint(s *Span, i int64) bool {
	res := C.left_span_bigint(s._inner, C.int64(i))
	return bool(res)
}


// LeftSpanFloat wraps MEOS C function left_span_float.
func LeftSpanFloat(s *Span, d float64) bool {
	res := C.left_span_float(s._inner, C.double(d))
	return bool(res)
}


// LeftSpanInt wraps MEOS C function left_span_int.
func LeftSpanInt(s *Span, i int) bool {
	res := C.left_span_int(s._inner, C.int(i))
	return bool(res)
}


// LeftSpanSpan wraps MEOS C function left_span_span.
func LeftSpanSpan(s1 *Span, s2 *Span) bool {
	res := C.left_span_span(s1._inner, s2._inner)
	return bool(res)
}


// LeftSpanSpanset wraps MEOS C function left_span_spanset.
func LeftSpanSpanset(s *Span, ss *SpanSet) bool {
	res := C.left_span_spanset(s._inner, ss._inner)
	return bool(res)
}


// LeftSpansetBigint wraps MEOS C function left_spanset_bigint.
func LeftSpansetBigint(ss *SpanSet, i int64) bool {
	res := C.left_spanset_bigint(ss._inner, C.int64(i))
	return bool(res)
}


// LeftSpansetFloat wraps MEOS C function left_spanset_float.
func LeftSpansetFloat(ss *SpanSet, d float64) bool {
	res := C.left_spanset_float(ss._inner, C.double(d))
	return bool(res)
}


// LeftSpansetInt wraps MEOS C function left_spanset_int.
func LeftSpansetInt(ss *SpanSet, i int) bool {
	res := C.left_spanset_int(ss._inner, C.int(i))
	return bool(res)
}


// LeftSpansetSpan wraps MEOS C function left_spanset_span.
func LeftSpansetSpan(ss *SpanSet, s *Span) bool {
	res := C.left_spanset_span(ss._inner, s._inner)
	return bool(res)
}


// LeftSpansetSpanset wraps MEOS C function left_spanset_spanset.
func LeftSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) bool {
	res := C.left_spanset_spanset(ss1._inner, ss2._inner)
	return bool(res)
}


// LeftTextSet wraps MEOS C function left_text_set.
func LeftTextSet(txt string, s *Set) bool {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.left_text_set(_c_txt, s._inner)
	return bool(res)
}


// OverafterDateSet wraps MEOS C function overafter_date_set.
func OverafterDateSet(d int32, s *Set) bool {
	res := C.overafter_date_set(C.DateADT(d), s._inner)
	return bool(res)
}


// OverafterDateSpan wraps MEOS C function overafter_date_span.
func OverafterDateSpan(d int32, s *Span) bool {
	res := C.overafter_date_span(C.DateADT(d), s._inner)
	return bool(res)
}


// OverafterDateSpanset wraps MEOS C function overafter_date_spanset.
func OverafterDateSpanset(d int32, ss *SpanSet) bool {
	res := C.overafter_date_spanset(C.DateADT(d), ss._inner)
	return bool(res)
}


// OverafterSetDate wraps MEOS C function overafter_set_date.
func OverafterSetDate(s *Set, d int32) bool {
	res := C.overafter_set_date(s._inner, C.DateADT(d))
	return bool(res)
}


// OverafterSetTimestamptz wraps MEOS C function overafter_set_timestamptz.
func OverafterSetTimestamptz(s *Set, t int64) bool {
	res := C.overafter_set_timestamptz(s._inner, C.TimestampTz(t))
	return bool(res)
}


// OverafterSpanDate wraps MEOS C function overafter_span_date.
func OverafterSpanDate(s *Span, d int32) bool {
	res := C.overafter_span_date(s._inner, C.DateADT(d))
	return bool(res)
}


// OverafterSpanTimestamptz wraps MEOS C function overafter_span_timestamptz.
func OverafterSpanTimestamptz(s *Span, t int64) bool {
	res := C.overafter_span_timestamptz(s._inner, C.TimestampTz(t))
	return bool(res)
}


// OverafterSpansetDate wraps MEOS C function overafter_spanset_date.
func OverafterSpansetDate(ss *SpanSet, d int32) bool {
	res := C.overafter_spanset_date(ss._inner, C.DateADT(d))
	return bool(res)
}


// OverafterSpansetTimestamptz wraps MEOS C function overafter_spanset_timestamptz.
func OverafterSpansetTimestamptz(ss *SpanSet, t int64) bool {
	res := C.overafter_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return bool(res)
}


// OverafterTimestamptzSet wraps MEOS C function overafter_timestamptz_set.
func OverafterTimestamptzSet(t int64, s *Set) bool {
	res := C.overafter_timestamptz_set(C.TimestampTz(t), s._inner)
	return bool(res)
}


// OverafterTimestamptzSpan wraps MEOS C function overafter_timestamptz_span.
func OverafterTimestamptzSpan(t int64, s *Span) bool {
	res := C.overafter_timestamptz_span(C.TimestampTz(t), s._inner)
	return bool(res)
}


// OverafterTimestamptzSpanset wraps MEOS C function overafter_timestamptz_spanset.
func OverafterTimestamptzSpanset(t int64, ss *SpanSet) bool {
	res := C.overafter_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	return bool(res)
}


// OverbeforeDateSet wraps MEOS C function overbefore_date_set.
func OverbeforeDateSet(d int32, s *Set) bool {
	res := C.overbefore_date_set(C.DateADT(d), s._inner)
	return bool(res)
}


// OverbeforeDateSpan wraps MEOS C function overbefore_date_span.
func OverbeforeDateSpan(d int32, s *Span) bool {
	res := C.overbefore_date_span(C.DateADT(d), s._inner)
	return bool(res)
}


// OverbeforeDateSpanset wraps MEOS C function overbefore_date_spanset.
func OverbeforeDateSpanset(d int32, ss *SpanSet) bool {
	res := C.overbefore_date_spanset(C.DateADT(d), ss._inner)
	return bool(res)
}


// OverbeforeSetDate wraps MEOS C function overbefore_set_date.
func OverbeforeSetDate(s *Set, d int32) bool {
	res := C.overbefore_set_date(s._inner, C.DateADT(d))
	return bool(res)
}


// OverbeforeSetTimestamptz wraps MEOS C function overbefore_set_timestamptz.
func OverbeforeSetTimestamptz(s *Set, t int64) bool {
	res := C.overbefore_set_timestamptz(s._inner, C.TimestampTz(t))
	return bool(res)
}


// OverbeforeSpanDate wraps MEOS C function overbefore_span_date.
func OverbeforeSpanDate(s *Span, d int32) bool {
	res := C.overbefore_span_date(s._inner, C.DateADT(d))
	return bool(res)
}


// OverbeforeSpanTimestamptz wraps MEOS C function overbefore_span_timestamptz.
func OverbeforeSpanTimestamptz(s *Span, t int64) bool {
	res := C.overbefore_span_timestamptz(s._inner, C.TimestampTz(t))
	return bool(res)
}


// OverbeforeSpansetDate wraps MEOS C function overbefore_spanset_date.
func OverbeforeSpansetDate(ss *SpanSet, d int32) bool {
	res := C.overbefore_spanset_date(ss._inner, C.DateADT(d))
	return bool(res)
}


// OverbeforeSpansetTimestamptz wraps MEOS C function overbefore_spanset_timestamptz.
func OverbeforeSpansetTimestamptz(ss *SpanSet, t int64) bool {
	res := C.overbefore_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return bool(res)
}


// OverbeforeTimestamptzSet wraps MEOS C function overbefore_timestamptz_set.
func OverbeforeTimestamptzSet(t int64, s *Set) bool {
	res := C.overbefore_timestamptz_set(C.TimestampTz(t), s._inner)
	return bool(res)
}


// OverbeforeTimestamptzSpan wraps MEOS C function overbefore_timestamptz_span.
func OverbeforeTimestamptzSpan(t int64, s *Span) bool {
	res := C.overbefore_timestamptz_span(C.TimestampTz(t), s._inner)
	return bool(res)
}


// OverbeforeTimestamptzSpanset wraps MEOS C function overbefore_timestamptz_spanset.
func OverbeforeTimestamptzSpanset(t int64, ss *SpanSet) bool {
	res := C.overbefore_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	return bool(res)
}


// OverleftBigintSet wraps MEOS C function overleft_bigint_set.
func OverleftBigintSet(i int64, s *Set) bool {
	res := C.overleft_bigint_set(C.int64(i), s._inner)
	return bool(res)
}


// OverleftBigintSpan wraps MEOS C function overleft_bigint_span.
func OverleftBigintSpan(i int64, s *Span) bool {
	res := C.overleft_bigint_span(C.int64(i), s._inner)
	return bool(res)
}


// OverleftBigintSpanset wraps MEOS C function overleft_bigint_spanset.
func OverleftBigintSpanset(i int64, ss *SpanSet) bool {
	res := C.overleft_bigint_spanset(C.int64(i), ss._inner)
	return bool(res)
}


// OverleftFloatSet wraps MEOS C function overleft_float_set.
func OverleftFloatSet(d float64, s *Set) bool {
	res := C.overleft_float_set(C.double(d), s._inner)
	return bool(res)
}


// OverleftFloatSpan wraps MEOS C function overleft_float_span.
func OverleftFloatSpan(d float64, s *Span) bool {
	res := C.overleft_float_span(C.double(d), s._inner)
	return bool(res)
}


// OverleftFloatSpanset wraps MEOS C function overleft_float_spanset.
func OverleftFloatSpanset(d float64, ss *SpanSet) bool {
	res := C.overleft_float_spanset(C.double(d), ss._inner)
	return bool(res)
}


// OverleftIntSet wraps MEOS C function overleft_int_set.
func OverleftIntSet(i int, s *Set) bool {
	res := C.overleft_int_set(C.int(i), s._inner)
	return bool(res)
}


// OverleftIntSpan wraps MEOS C function overleft_int_span.
func OverleftIntSpan(i int, s *Span) bool {
	res := C.overleft_int_span(C.int(i), s._inner)
	return bool(res)
}


// OverleftIntSpanset wraps MEOS C function overleft_int_spanset.
func OverleftIntSpanset(i int, ss *SpanSet) bool {
	res := C.overleft_int_spanset(C.int(i), ss._inner)
	return bool(res)
}


// OverleftSetBigint wraps MEOS C function overleft_set_bigint.
func OverleftSetBigint(s *Set, i int64) bool {
	res := C.overleft_set_bigint(s._inner, C.int64(i))
	return bool(res)
}


// OverleftSetFloat wraps MEOS C function overleft_set_float.
func OverleftSetFloat(s *Set, d float64) bool {
	res := C.overleft_set_float(s._inner, C.double(d))
	return bool(res)
}


// OverleftSetInt wraps MEOS C function overleft_set_int.
func OverleftSetInt(s *Set, i int) bool {
	res := C.overleft_set_int(s._inner, C.int(i))
	return bool(res)
}


// OverleftSetSet wraps MEOS C function overleft_set_set.
func OverleftSetSet(s1 *Set, s2 *Set) bool {
	res := C.overleft_set_set(s1._inner, s2._inner)
	return bool(res)
}


// OverleftSetText wraps MEOS C function overleft_set_text.
func OverleftSetText(s *Set, txt string) bool {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.overleft_set_text(s._inner, _c_txt)
	return bool(res)
}


// OverleftSpanBigint wraps MEOS C function overleft_span_bigint.
func OverleftSpanBigint(s *Span, i int64) bool {
	res := C.overleft_span_bigint(s._inner, C.int64(i))
	return bool(res)
}


// OverleftSpanFloat wraps MEOS C function overleft_span_float.
func OverleftSpanFloat(s *Span, d float64) bool {
	res := C.overleft_span_float(s._inner, C.double(d))
	return bool(res)
}


// OverleftSpanInt wraps MEOS C function overleft_span_int.
func OverleftSpanInt(s *Span, i int) bool {
	res := C.overleft_span_int(s._inner, C.int(i))
	return bool(res)
}


// OverleftSpanSpan wraps MEOS C function overleft_span_span.
func OverleftSpanSpan(s1 *Span, s2 *Span) bool {
	res := C.overleft_span_span(s1._inner, s2._inner)
	return bool(res)
}


// OverleftSpanSpanset wraps MEOS C function overleft_span_spanset.
func OverleftSpanSpanset(s *Span, ss *SpanSet) bool {
	res := C.overleft_span_spanset(s._inner, ss._inner)
	return bool(res)
}


// OverleftSpansetBigint wraps MEOS C function overleft_spanset_bigint.
func OverleftSpansetBigint(ss *SpanSet, i int64) bool {
	res := C.overleft_spanset_bigint(ss._inner, C.int64(i))
	return bool(res)
}


// OverleftSpansetFloat wraps MEOS C function overleft_spanset_float.
func OverleftSpansetFloat(ss *SpanSet, d float64) bool {
	res := C.overleft_spanset_float(ss._inner, C.double(d))
	return bool(res)
}


// OverleftSpansetInt wraps MEOS C function overleft_spanset_int.
func OverleftSpansetInt(ss *SpanSet, i int) bool {
	res := C.overleft_spanset_int(ss._inner, C.int(i))
	return bool(res)
}


// OverleftSpansetSpan wraps MEOS C function overleft_spanset_span.
func OverleftSpansetSpan(ss *SpanSet, s *Span) bool {
	res := C.overleft_spanset_span(ss._inner, s._inner)
	return bool(res)
}


// OverleftSpansetSpanset wraps MEOS C function overleft_spanset_spanset.
func OverleftSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) bool {
	res := C.overleft_spanset_spanset(ss1._inner, ss2._inner)
	return bool(res)
}


// OverleftTextSet wraps MEOS C function overleft_text_set.
func OverleftTextSet(txt string, s *Set) bool {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.overleft_text_set(_c_txt, s._inner)
	return bool(res)
}


// OverrightBigintSet wraps MEOS C function overright_bigint_set.
func OverrightBigintSet(i int64, s *Set) bool {
	res := C.overright_bigint_set(C.int64(i), s._inner)
	return bool(res)
}


// OverrightBigintSpan wraps MEOS C function overright_bigint_span.
func OverrightBigintSpan(i int64, s *Span) bool {
	res := C.overright_bigint_span(C.int64(i), s._inner)
	return bool(res)
}


// OverrightBigintSpanset wraps MEOS C function overright_bigint_spanset.
func OverrightBigintSpanset(i int64, ss *SpanSet) bool {
	res := C.overright_bigint_spanset(C.int64(i), ss._inner)
	return bool(res)
}


// OverrightFloatSet wraps MEOS C function overright_float_set.
func OverrightFloatSet(d float64, s *Set) bool {
	res := C.overright_float_set(C.double(d), s._inner)
	return bool(res)
}


// OverrightFloatSpan wraps MEOS C function overright_float_span.
func OverrightFloatSpan(d float64, s *Span) bool {
	res := C.overright_float_span(C.double(d), s._inner)
	return bool(res)
}


// OverrightFloatSpanset wraps MEOS C function overright_float_spanset.
func OverrightFloatSpanset(d float64, ss *SpanSet) bool {
	res := C.overright_float_spanset(C.double(d), ss._inner)
	return bool(res)
}


// OverrightIntSet wraps MEOS C function overright_int_set.
func OverrightIntSet(i int, s *Set) bool {
	res := C.overright_int_set(C.int(i), s._inner)
	return bool(res)
}


// OverrightIntSpan wraps MEOS C function overright_int_span.
func OverrightIntSpan(i int, s *Span) bool {
	res := C.overright_int_span(C.int(i), s._inner)
	return bool(res)
}


// OverrightIntSpanset wraps MEOS C function overright_int_spanset.
func OverrightIntSpanset(i int, ss *SpanSet) bool {
	res := C.overright_int_spanset(C.int(i), ss._inner)
	return bool(res)
}


// OverrightSetBigint wraps MEOS C function overright_set_bigint.
func OverrightSetBigint(s *Set, i int64) bool {
	res := C.overright_set_bigint(s._inner, C.int64(i))
	return bool(res)
}


// OverrightSetFloat wraps MEOS C function overright_set_float.
func OverrightSetFloat(s *Set, d float64) bool {
	res := C.overright_set_float(s._inner, C.double(d))
	return bool(res)
}


// OverrightSetInt wraps MEOS C function overright_set_int.
func OverrightSetInt(s *Set, i int) bool {
	res := C.overright_set_int(s._inner, C.int(i))
	return bool(res)
}


// OverrightSetSet wraps MEOS C function overright_set_set.
func OverrightSetSet(s1 *Set, s2 *Set) bool {
	res := C.overright_set_set(s1._inner, s2._inner)
	return bool(res)
}


// OverrightSetText wraps MEOS C function overright_set_text.
func OverrightSetText(s *Set, txt string) bool {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.overright_set_text(s._inner, _c_txt)
	return bool(res)
}


// OverrightSpanBigint wraps MEOS C function overright_span_bigint.
func OverrightSpanBigint(s *Span, i int64) bool {
	res := C.overright_span_bigint(s._inner, C.int64(i))
	return bool(res)
}


// OverrightSpanFloat wraps MEOS C function overright_span_float.
func OverrightSpanFloat(s *Span, d float64) bool {
	res := C.overright_span_float(s._inner, C.double(d))
	return bool(res)
}


// OverrightSpanInt wraps MEOS C function overright_span_int.
func OverrightSpanInt(s *Span, i int) bool {
	res := C.overright_span_int(s._inner, C.int(i))
	return bool(res)
}


// OverrightSpanSpan wraps MEOS C function overright_span_span.
func OverrightSpanSpan(s1 *Span, s2 *Span) bool {
	res := C.overright_span_span(s1._inner, s2._inner)
	return bool(res)
}


// OverrightSpanSpanset wraps MEOS C function overright_span_spanset.
func OverrightSpanSpanset(s *Span, ss *SpanSet) bool {
	res := C.overright_span_spanset(s._inner, ss._inner)
	return bool(res)
}


// OverrightSpansetBigint wraps MEOS C function overright_spanset_bigint.
func OverrightSpansetBigint(ss *SpanSet, i int64) bool {
	res := C.overright_spanset_bigint(ss._inner, C.int64(i))
	return bool(res)
}


// OverrightSpansetFloat wraps MEOS C function overright_spanset_float.
func OverrightSpansetFloat(ss *SpanSet, d float64) bool {
	res := C.overright_spanset_float(ss._inner, C.double(d))
	return bool(res)
}


// OverrightSpansetInt wraps MEOS C function overright_spanset_int.
func OverrightSpansetInt(ss *SpanSet, i int) bool {
	res := C.overright_spanset_int(ss._inner, C.int(i))
	return bool(res)
}


// OverrightSpansetSpan wraps MEOS C function overright_spanset_span.
func OverrightSpansetSpan(ss *SpanSet, s *Span) bool {
	res := C.overright_spanset_span(ss._inner, s._inner)
	return bool(res)
}


// OverrightSpansetSpanset wraps MEOS C function overright_spanset_spanset.
func OverrightSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) bool {
	res := C.overright_spanset_spanset(ss1._inner, ss2._inner)
	return bool(res)
}


// OverrightTextSet wraps MEOS C function overright_text_set.
func OverrightTextSet(txt string, s *Set) bool {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.overright_text_set(_c_txt, s._inner)
	return bool(res)
}


// RightBigintSet wraps MEOS C function right_bigint_set.
func RightBigintSet(i int64, s *Set) bool {
	res := C.right_bigint_set(C.int64(i), s._inner)
	return bool(res)
}


// RightBigintSpan wraps MEOS C function right_bigint_span.
func RightBigintSpan(i int64, s *Span) bool {
	res := C.right_bigint_span(C.int64(i), s._inner)
	return bool(res)
}


// RightBigintSpanset wraps MEOS C function right_bigint_spanset.
func RightBigintSpanset(i int64, ss *SpanSet) bool {
	res := C.right_bigint_spanset(C.int64(i), ss._inner)
	return bool(res)
}


// RightFloatSet wraps MEOS C function right_float_set.
func RightFloatSet(d float64, s *Set) bool {
	res := C.right_float_set(C.double(d), s._inner)
	return bool(res)
}


// RightFloatSpan wraps MEOS C function right_float_span.
func RightFloatSpan(d float64, s *Span) bool {
	res := C.right_float_span(C.double(d), s._inner)
	return bool(res)
}


// RightFloatSpanset wraps MEOS C function right_float_spanset.
func RightFloatSpanset(d float64, ss *SpanSet) bool {
	res := C.right_float_spanset(C.double(d), ss._inner)
	return bool(res)
}


// RightIntSet wraps MEOS C function right_int_set.
func RightIntSet(i int, s *Set) bool {
	res := C.right_int_set(C.int(i), s._inner)
	return bool(res)
}


// RightIntSpan wraps MEOS C function right_int_span.
func RightIntSpan(i int, s *Span) bool {
	res := C.right_int_span(C.int(i), s._inner)
	return bool(res)
}


// RightIntSpanset wraps MEOS C function right_int_spanset.
func RightIntSpanset(i int, ss *SpanSet) bool {
	res := C.right_int_spanset(C.int(i), ss._inner)
	return bool(res)
}


// RightSetBigint wraps MEOS C function right_set_bigint.
func RightSetBigint(s *Set, i int64) bool {
	res := C.right_set_bigint(s._inner, C.int64(i))
	return bool(res)
}


// RightSetFloat wraps MEOS C function right_set_float.
func RightSetFloat(s *Set, d float64) bool {
	res := C.right_set_float(s._inner, C.double(d))
	return bool(res)
}


// RightSetInt wraps MEOS C function right_set_int.
func RightSetInt(s *Set, i int) bool {
	res := C.right_set_int(s._inner, C.int(i))
	return bool(res)
}


// RightSetSet wraps MEOS C function right_set_set.
func RightSetSet(s1 *Set, s2 *Set) bool {
	res := C.right_set_set(s1._inner, s2._inner)
	return bool(res)
}


// RightSetText wraps MEOS C function right_set_text.
func RightSetText(s *Set, txt string) bool {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.right_set_text(s._inner, _c_txt)
	return bool(res)
}


// RightSpanBigint wraps MEOS C function right_span_bigint.
func RightSpanBigint(s *Span, i int64) bool {
	res := C.right_span_bigint(s._inner, C.int64(i))
	return bool(res)
}


// RightSpanFloat wraps MEOS C function right_span_float.
func RightSpanFloat(s *Span, d float64) bool {
	res := C.right_span_float(s._inner, C.double(d))
	return bool(res)
}


// RightSpanInt wraps MEOS C function right_span_int.
func RightSpanInt(s *Span, i int) bool {
	res := C.right_span_int(s._inner, C.int(i))
	return bool(res)
}


// RightSpanSpan wraps MEOS C function right_span_span.
func RightSpanSpan(s1 *Span, s2 *Span) bool {
	res := C.right_span_span(s1._inner, s2._inner)
	return bool(res)
}


// RightSpanSpanset wraps MEOS C function right_span_spanset.
func RightSpanSpanset(s *Span, ss *SpanSet) bool {
	res := C.right_span_spanset(s._inner, ss._inner)
	return bool(res)
}


// RightSpansetBigint wraps MEOS C function right_spanset_bigint.
func RightSpansetBigint(ss *SpanSet, i int64) bool {
	res := C.right_spanset_bigint(ss._inner, C.int64(i))
	return bool(res)
}


// RightSpansetFloat wraps MEOS C function right_spanset_float.
func RightSpansetFloat(ss *SpanSet, d float64) bool {
	res := C.right_spanset_float(ss._inner, C.double(d))
	return bool(res)
}


// RightSpansetInt wraps MEOS C function right_spanset_int.
func RightSpansetInt(ss *SpanSet, i int) bool {
	res := C.right_spanset_int(ss._inner, C.int(i))
	return bool(res)
}


// RightSpansetSpan wraps MEOS C function right_spanset_span.
func RightSpansetSpan(ss *SpanSet, s *Span) bool {
	res := C.right_spanset_span(ss._inner, s._inner)
	return bool(res)
}


// RightSpansetSpanset wraps MEOS C function right_spanset_spanset.
func RightSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) bool {
	res := C.right_spanset_spanset(ss1._inner, ss2._inner)
	return bool(res)
}


// RightTextSet wraps MEOS C function right_text_set.
func RightTextSet(txt string, s *Set) bool {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.right_text_set(_c_txt, s._inner)
	return bool(res)
}


// IntersectionBigintSet wraps MEOS C function intersection_bigint_set.
func IntersectionBigintSet(i int64, s *Set) *Set {
	res := C.intersection_bigint_set(C.int64(i), s._inner)
	return &Set{_inner: res}
}


// IntersectionDateSet wraps MEOS C function intersection_date_set.
func IntersectionDateSet(d int32, s *Set) *Set {
	res := C.intersection_date_set(C.DateADT(d), s._inner)
	return &Set{_inner: res}
}


// IntersectionFloatSet wraps MEOS C function intersection_float_set.
func IntersectionFloatSet(d float64, s *Set) *Set {
	res := C.intersection_float_set(C.double(d), s._inner)
	return &Set{_inner: res}
}


// IntersectionIntSet wraps MEOS C function intersection_int_set.
func IntersectionIntSet(i int, s *Set) *Set {
	res := C.intersection_int_set(C.int(i), s._inner)
	return &Set{_inner: res}
}


// IntersectionSetBigint wraps MEOS C function intersection_set_bigint.
func IntersectionSetBigint(s *Set, i int64) *Set {
	res := C.intersection_set_bigint(s._inner, C.int64(i))
	return &Set{_inner: res}
}


// IntersectionSetDate wraps MEOS C function intersection_set_date.
func IntersectionSetDate(s *Set, d int32) *Set {
	res := C.intersection_set_date(s._inner, C.DateADT(d))
	return &Set{_inner: res}
}


// IntersectionSetFloat wraps MEOS C function intersection_set_float.
func IntersectionSetFloat(s *Set, d float64) *Set {
	res := C.intersection_set_float(s._inner, C.double(d))
	return &Set{_inner: res}
}


// IntersectionSetInt wraps MEOS C function intersection_set_int.
func IntersectionSetInt(s *Set, i int) *Set {
	res := C.intersection_set_int(s._inner, C.int(i))
	return &Set{_inner: res}
}


// IntersectionSetSet wraps MEOS C function intersection_set_set.
func IntersectionSetSet(s1 *Set, s2 *Set) *Set {
	res := C.intersection_set_set(s1._inner, s2._inner)
	return &Set{_inner: res}
}


// IntersectionSetText wraps MEOS C function intersection_set_text.
func IntersectionSetText(s *Set, txt string) *Set {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.intersection_set_text(s._inner, _c_txt)
	return &Set{_inner: res}
}


// IntersectionSetTimestamptz wraps MEOS C function intersection_set_timestamptz.
func IntersectionSetTimestamptz(s *Set, t int64) *Set {
	res := C.intersection_set_timestamptz(s._inner, C.TimestampTz(t))
	return &Set{_inner: res}
}


// IntersectionSpanBigint wraps MEOS C function intersection_span_bigint.
func IntersectionSpanBigint(s *Span, i int64) *Span {
	res := C.intersection_span_bigint(s._inner, C.int64(i))
	return &Span{_inner: res}
}


// IntersectionSpanDate wraps MEOS C function intersection_span_date.
func IntersectionSpanDate(s *Span, d int32) *Span {
	res := C.intersection_span_date(s._inner, C.DateADT(d))
	return &Span{_inner: res}
}


// IntersectionSpanFloat wraps MEOS C function intersection_span_float.
func IntersectionSpanFloat(s *Span, d float64) *Span {
	res := C.intersection_span_float(s._inner, C.double(d))
	return &Span{_inner: res}
}


// IntersectionSpanInt wraps MEOS C function intersection_span_int.
func IntersectionSpanInt(s *Span, i int) *Span {
	res := C.intersection_span_int(s._inner, C.int(i))
	return &Span{_inner: res}
}


// IntersectionSpanSpan wraps MEOS C function intersection_span_span.
func IntersectionSpanSpan(s1 *Span, s2 *Span) *Span {
	res := C.intersection_span_span(s1._inner, s2._inner)
	return &Span{_inner: res}
}


// IntersectionSpanSpanset wraps MEOS C function intersection_span_spanset.
func IntersectionSpanSpanset(s *Span, ss *SpanSet) *SpanSet {
	res := C.intersection_span_spanset(s._inner, ss._inner)
	return &SpanSet{_inner: res}
}


// IntersectionSpanTimestamptz wraps MEOS C function intersection_span_timestamptz.
func IntersectionSpanTimestamptz(s *Span, t int64) *Span {
	res := C.intersection_span_timestamptz(s._inner, C.TimestampTz(t))
	return &Span{_inner: res}
}


// IntersectionSpansetBigint wraps MEOS C function intersection_spanset_bigint.
func IntersectionSpansetBigint(ss *SpanSet, i int64) *SpanSet {
	res := C.intersection_spanset_bigint(ss._inner, C.int64(i))
	return &SpanSet{_inner: res}
}


// IntersectionSpansetDate wraps MEOS C function intersection_spanset_date.
func IntersectionSpansetDate(ss *SpanSet, d int32) *SpanSet {
	res := C.intersection_spanset_date(ss._inner, C.DateADT(d))
	return &SpanSet{_inner: res}
}


// IntersectionSpansetFloat wraps MEOS C function intersection_spanset_float.
func IntersectionSpansetFloat(ss *SpanSet, d float64) *SpanSet {
	res := C.intersection_spanset_float(ss._inner, C.double(d))
	return &SpanSet{_inner: res}
}


// IntersectionSpansetInt wraps MEOS C function intersection_spanset_int.
func IntersectionSpansetInt(ss *SpanSet, i int) *SpanSet {
	res := C.intersection_spanset_int(ss._inner, C.int(i))
	return &SpanSet{_inner: res}
}


// IntersectionSpansetSpan wraps MEOS C function intersection_spanset_span.
func IntersectionSpansetSpan(ss *SpanSet, s *Span) *SpanSet {
	res := C.intersection_spanset_span(ss._inner, s._inner)
	return &SpanSet{_inner: res}
}


// IntersectionSpansetSpanset wraps MEOS C function intersection_spanset_spanset.
func IntersectionSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) *SpanSet {
	res := C.intersection_spanset_spanset(ss1._inner, ss2._inner)
	return &SpanSet{_inner: res}
}


// IntersectionSpansetTimestamptz wraps MEOS C function intersection_spanset_timestamptz.
func IntersectionSpansetTimestamptz(ss *SpanSet, t int64) *SpanSet {
	res := C.intersection_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return &SpanSet{_inner: res}
}


// IntersectionTextSet wraps MEOS C function intersection_text_set.
func IntersectionTextSet(txt string, s *Set) *Set {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.intersection_text_set(_c_txt, s._inner)
	return &Set{_inner: res}
}


// IntersectionTimestamptzSet wraps MEOS C function intersection_timestamptz_set.
func IntersectionTimestamptzSet(t int64, s *Set) *Set {
	res := C.intersection_timestamptz_set(C.TimestampTz(t), s._inner)
	return &Set{_inner: res}
}


// MinusBigintSet wraps MEOS C function minus_bigint_set.
func MinusBigintSet(i int64, s *Set) *Set {
	res := C.minus_bigint_set(C.int64(i), s._inner)
	return &Set{_inner: res}
}


// MinusBigintSpan wraps MEOS C function minus_bigint_span.
func MinusBigintSpan(i int64, s *Span) *SpanSet {
	res := C.minus_bigint_span(C.int64(i), s._inner)
	return &SpanSet{_inner: res}
}


// MinusBigintSpanset wraps MEOS C function minus_bigint_spanset.
func MinusBigintSpanset(i int64, ss *SpanSet) *SpanSet {
	res := C.minus_bigint_spanset(C.int64(i), ss._inner)
	return &SpanSet{_inner: res}
}


// MinusDateSet wraps MEOS C function minus_date_set.
func MinusDateSet(d int32, s *Set) *Set {
	res := C.minus_date_set(C.DateADT(d), s._inner)
	return &Set{_inner: res}
}


// MinusDateSpan wraps MEOS C function minus_date_span.
func MinusDateSpan(d int32, s *Span) *SpanSet {
	res := C.minus_date_span(C.DateADT(d), s._inner)
	return &SpanSet{_inner: res}
}


// MinusDateSpanset wraps MEOS C function minus_date_spanset.
func MinusDateSpanset(d int32, ss *SpanSet) *SpanSet {
	res := C.minus_date_spanset(C.DateADT(d), ss._inner)
	return &SpanSet{_inner: res}
}


// MinusFloatSet wraps MEOS C function minus_float_set.
func MinusFloatSet(d float64, s *Set) *Set {
	res := C.minus_float_set(C.double(d), s._inner)
	return &Set{_inner: res}
}


// MinusFloatSpan wraps MEOS C function minus_float_span.
func MinusFloatSpan(d float64, s *Span) *SpanSet {
	res := C.minus_float_span(C.double(d), s._inner)
	return &SpanSet{_inner: res}
}


// MinusFloatSpanset wraps MEOS C function minus_float_spanset.
func MinusFloatSpanset(d float64, ss *SpanSet) *SpanSet {
	res := C.minus_float_spanset(C.double(d), ss._inner)
	return &SpanSet{_inner: res}
}


// MinusIntSet wraps MEOS C function minus_int_set.
func MinusIntSet(i int, s *Set) *Set {
	res := C.minus_int_set(C.int(i), s._inner)
	return &Set{_inner: res}
}


// MinusIntSpan wraps MEOS C function minus_int_span.
func MinusIntSpan(i int, s *Span) *SpanSet {
	res := C.minus_int_span(C.int(i), s._inner)
	return &SpanSet{_inner: res}
}


// MinusIntSpanset wraps MEOS C function minus_int_spanset.
func MinusIntSpanset(i int, ss *SpanSet) *SpanSet {
	res := C.minus_int_spanset(C.int(i), ss._inner)
	return &SpanSet{_inner: res}
}


// MinusSetBigint wraps MEOS C function minus_set_bigint.
func MinusSetBigint(s *Set, i int64) *Set {
	res := C.minus_set_bigint(s._inner, C.int64(i))
	return &Set{_inner: res}
}


// MinusSetDate wraps MEOS C function minus_set_date.
func MinusSetDate(s *Set, d int32) *Set {
	res := C.minus_set_date(s._inner, C.DateADT(d))
	return &Set{_inner: res}
}


// MinusSetFloat wraps MEOS C function minus_set_float.
func MinusSetFloat(s *Set, d float64) *Set {
	res := C.minus_set_float(s._inner, C.double(d))
	return &Set{_inner: res}
}


// MinusSetInt wraps MEOS C function minus_set_int.
func MinusSetInt(s *Set, i int) *Set {
	res := C.minus_set_int(s._inner, C.int(i))
	return &Set{_inner: res}
}


// MinusSetSet wraps MEOS C function minus_set_set.
func MinusSetSet(s1 *Set, s2 *Set) *Set {
	res := C.minus_set_set(s1._inner, s2._inner)
	return &Set{_inner: res}
}


// MinusSetText wraps MEOS C function minus_set_text.
func MinusSetText(s *Set, txt string) *Set {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.minus_set_text(s._inner, _c_txt)
	return &Set{_inner: res}
}


// MinusSetTimestamptz wraps MEOS C function minus_set_timestamptz.
func MinusSetTimestamptz(s *Set, t int64) *Set {
	res := C.minus_set_timestamptz(s._inner, C.TimestampTz(t))
	return &Set{_inner: res}
}


// MinusSpanBigint wraps MEOS C function minus_span_bigint.
func MinusSpanBigint(s *Span, i int64) *SpanSet {
	res := C.minus_span_bigint(s._inner, C.int64(i))
	return &SpanSet{_inner: res}
}


// MinusSpanDate wraps MEOS C function minus_span_date.
func MinusSpanDate(s *Span, d int32) *SpanSet {
	res := C.minus_span_date(s._inner, C.DateADT(d))
	return &SpanSet{_inner: res}
}


// MinusSpanFloat wraps MEOS C function minus_span_float.
func MinusSpanFloat(s *Span, d float64) *SpanSet {
	res := C.minus_span_float(s._inner, C.double(d))
	return &SpanSet{_inner: res}
}


// MinusSpanInt wraps MEOS C function minus_span_int.
func MinusSpanInt(s *Span, i int) *SpanSet {
	res := C.minus_span_int(s._inner, C.int(i))
	return &SpanSet{_inner: res}
}


// MinusSpanSpan wraps MEOS C function minus_span_span.
func MinusSpanSpan(s1 *Span, s2 *Span) *SpanSet {
	res := C.minus_span_span(s1._inner, s2._inner)
	return &SpanSet{_inner: res}
}


// MinusSpanSpanset wraps MEOS C function minus_span_spanset.
func MinusSpanSpanset(s *Span, ss *SpanSet) *SpanSet {
	res := C.minus_span_spanset(s._inner, ss._inner)
	return &SpanSet{_inner: res}
}


// MinusSpanTimestamptz wraps MEOS C function minus_span_timestamptz.
func MinusSpanTimestamptz(s *Span, t int64) *SpanSet {
	res := C.minus_span_timestamptz(s._inner, C.TimestampTz(t))
	return &SpanSet{_inner: res}
}


// MinusSpansetBigint wraps MEOS C function minus_spanset_bigint.
func MinusSpansetBigint(ss *SpanSet, i int64) *SpanSet {
	res := C.minus_spanset_bigint(ss._inner, C.int64(i))
	return &SpanSet{_inner: res}
}


// MinusSpansetDate wraps MEOS C function minus_spanset_date.
func MinusSpansetDate(ss *SpanSet, d int32) *SpanSet {
	res := C.minus_spanset_date(ss._inner, C.DateADT(d))
	return &SpanSet{_inner: res}
}


// MinusSpansetFloat wraps MEOS C function minus_spanset_float.
func MinusSpansetFloat(ss *SpanSet, d float64) *SpanSet {
	res := C.minus_spanset_float(ss._inner, C.double(d))
	return &SpanSet{_inner: res}
}


// MinusSpansetInt wraps MEOS C function minus_spanset_int.
func MinusSpansetInt(ss *SpanSet, i int) *SpanSet {
	res := C.minus_spanset_int(ss._inner, C.int(i))
	return &SpanSet{_inner: res}
}


// MinusSpansetSpan wraps MEOS C function minus_spanset_span.
func MinusSpansetSpan(ss *SpanSet, s *Span) *SpanSet {
	res := C.minus_spanset_span(ss._inner, s._inner)
	return &SpanSet{_inner: res}
}


// MinusSpansetSpanset wraps MEOS C function minus_spanset_spanset.
func MinusSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) *SpanSet {
	res := C.minus_spanset_spanset(ss1._inner, ss2._inner)
	return &SpanSet{_inner: res}
}


// MinusSpansetTimestamptz wraps MEOS C function minus_spanset_timestamptz.
func MinusSpansetTimestamptz(ss *SpanSet, t int64) *SpanSet {
	res := C.minus_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return &SpanSet{_inner: res}
}


// MinusTextSet wraps MEOS C function minus_text_set.
func MinusTextSet(txt string, s *Set) *Set {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.minus_text_set(_c_txt, s._inner)
	return &Set{_inner: res}
}


// MinusTimestamptzSet wraps MEOS C function minus_timestamptz_set.
func MinusTimestamptzSet(t int64, s *Set) *Set {
	res := C.minus_timestamptz_set(C.TimestampTz(t), s._inner)
	return &Set{_inner: res}
}


// MinusTimestamptzSpan wraps MEOS C function minus_timestamptz_span.
func MinusTimestamptzSpan(t int64, s *Span) *SpanSet {
	res := C.minus_timestamptz_span(C.TimestampTz(t), s._inner)
	return &SpanSet{_inner: res}
}


// MinusTimestamptzSpanset wraps MEOS C function minus_timestamptz_spanset.
func MinusTimestamptzSpanset(t int64, ss *SpanSet) *SpanSet {
	res := C.minus_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	return &SpanSet{_inner: res}
}


// UnionBigintSet wraps MEOS C function union_bigint_set.
func UnionBigintSet(i int64, s *Set) *Set {
	res := C.union_bigint_set(C.int64(i), s._inner)
	return &Set{_inner: res}
}


// UnionBigintSpan wraps MEOS C function union_bigint_span.
func UnionBigintSpan(s *Span, i int64) *SpanSet {
	res := C.union_bigint_span(s._inner, C.int64(i))
	return &SpanSet{_inner: res}
}


// UnionBigintSpanset wraps MEOS C function union_bigint_spanset.
func UnionBigintSpanset(i int64, ss *SpanSet) *SpanSet {
	res := C.union_bigint_spanset(C.int64(i), ss._inner)
	return &SpanSet{_inner: res}
}


// UnionDateSet wraps MEOS C function union_date_set.
func UnionDateSet(d int32, s *Set) *Set {
	res := C.union_date_set(C.DateADT(d), s._inner)
	return &Set{_inner: res}
}


// UnionDateSpan wraps MEOS C function union_date_span.
func UnionDateSpan(s *Span, d int32) *SpanSet {
	res := C.union_date_span(s._inner, C.DateADT(d))
	return &SpanSet{_inner: res}
}


// UnionDateSpanset wraps MEOS C function union_date_spanset.
func UnionDateSpanset(d int32, ss *SpanSet) *SpanSet {
	res := C.union_date_spanset(C.DateADT(d), ss._inner)
	return &SpanSet{_inner: res}
}


// UnionFloatSet wraps MEOS C function union_float_set.
func UnionFloatSet(d float64, s *Set) *Set {
	res := C.union_float_set(C.double(d), s._inner)
	return &Set{_inner: res}
}


// UnionFloatSpan wraps MEOS C function union_float_span.
func UnionFloatSpan(s *Span, d float64) *SpanSet {
	res := C.union_float_span(s._inner, C.double(d))
	return &SpanSet{_inner: res}
}


// UnionFloatSpanset wraps MEOS C function union_float_spanset.
func UnionFloatSpanset(d float64, ss *SpanSet) *SpanSet {
	res := C.union_float_spanset(C.double(d), ss._inner)
	return &SpanSet{_inner: res}
}


// UnionIntSet wraps MEOS C function union_int_set.
func UnionIntSet(i int, s *Set) *Set {
	res := C.union_int_set(C.int(i), s._inner)
	return &Set{_inner: res}
}


// UnionIntSpan wraps MEOS C function union_int_span.
func UnionIntSpan(i int, s *Span) *SpanSet {
	res := C.union_int_span(C.int(i), s._inner)
	return &SpanSet{_inner: res}
}


// UnionIntSpanset wraps MEOS C function union_int_spanset.
func UnionIntSpanset(i int, ss *SpanSet) *SpanSet {
	res := C.union_int_spanset(C.int(i), ss._inner)
	return &SpanSet{_inner: res}
}


// UnionSetBigint wraps MEOS C function union_set_bigint.
func UnionSetBigint(s *Set, i int64) *Set {
	res := C.union_set_bigint(s._inner, C.int64(i))
	return &Set{_inner: res}
}


// UnionSetDate wraps MEOS C function union_set_date.
func UnionSetDate(s *Set, d int32) *Set {
	res := C.union_set_date(s._inner, C.DateADT(d))
	return &Set{_inner: res}
}


// UnionSetFloat wraps MEOS C function union_set_float.
func UnionSetFloat(s *Set, d float64) *Set {
	res := C.union_set_float(s._inner, C.double(d))
	return &Set{_inner: res}
}


// UnionSetInt wraps MEOS C function union_set_int.
func UnionSetInt(s *Set, i int) *Set {
	res := C.union_set_int(s._inner, C.int(i))
	return &Set{_inner: res}
}


// UnionSetSet wraps MEOS C function union_set_set.
func UnionSetSet(s1 *Set, s2 *Set) *Set {
	res := C.union_set_set(s1._inner, s2._inner)
	return &Set{_inner: res}
}


// UnionSetText wraps MEOS C function union_set_text.
func UnionSetText(s *Set, txt string) *Set {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.union_set_text(s._inner, _c_txt)
	return &Set{_inner: res}
}


// UnionSetTimestamptz wraps MEOS C function union_set_timestamptz.
func UnionSetTimestamptz(s *Set, t int64) *Set {
	res := C.union_set_timestamptz(s._inner, C.TimestampTz(t))
	return &Set{_inner: res}
}


// UnionSpanBigint wraps MEOS C function union_span_bigint.
func UnionSpanBigint(s *Span, i int64) *SpanSet {
	res := C.union_span_bigint(s._inner, C.int64(i))
	return &SpanSet{_inner: res}
}


// UnionSpanDate wraps MEOS C function union_span_date.
func UnionSpanDate(s *Span, d int32) *SpanSet {
	res := C.union_span_date(s._inner, C.DateADT(d))
	return &SpanSet{_inner: res}
}


// UnionSpanFloat wraps MEOS C function union_span_float.
func UnionSpanFloat(s *Span, d float64) *SpanSet {
	res := C.union_span_float(s._inner, C.double(d))
	return &SpanSet{_inner: res}
}


// UnionSpanInt wraps MEOS C function union_span_int.
func UnionSpanInt(s *Span, i int) *SpanSet {
	res := C.union_span_int(s._inner, C.int(i))
	return &SpanSet{_inner: res}
}


// UnionSpanSpan wraps MEOS C function union_span_span.
func UnionSpanSpan(s1 *Span, s2 *Span) *SpanSet {
	res := C.union_span_span(s1._inner, s2._inner)
	return &SpanSet{_inner: res}
}


// UnionSpanSpanset wraps MEOS C function union_span_spanset.
func UnionSpanSpanset(s *Span, ss *SpanSet) *SpanSet {
	res := C.union_span_spanset(s._inner, ss._inner)
	return &SpanSet{_inner: res}
}


// UnionSpanTimestamptz wraps MEOS C function union_span_timestamptz.
func UnionSpanTimestamptz(s *Span, t int64) *SpanSet {
	res := C.union_span_timestamptz(s._inner, C.TimestampTz(t))
	return &SpanSet{_inner: res}
}


// UnionSpansetBigint wraps MEOS C function union_spanset_bigint.
func UnionSpansetBigint(ss *SpanSet, i int64) *SpanSet {
	res := C.union_spanset_bigint(ss._inner, C.int64(i))
	return &SpanSet{_inner: res}
}


// UnionSpansetDate wraps MEOS C function union_spanset_date.
func UnionSpansetDate(ss *SpanSet, d int32) *SpanSet {
	res := C.union_spanset_date(ss._inner, C.DateADT(d))
	return &SpanSet{_inner: res}
}


// UnionSpansetFloat wraps MEOS C function union_spanset_float.
func UnionSpansetFloat(ss *SpanSet, d float64) *SpanSet {
	res := C.union_spanset_float(ss._inner, C.double(d))
	return &SpanSet{_inner: res}
}


// UnionSpansetInt wraps MEOS C function union_spanset_int.
func UnionSpansetInt(ss *SpanSet, i int) *SpanSet {
	res := C.union_spanset_int(ss._inner, C.int(i))
	return &SpanSet{_inner: res}
}


// UnionSpansetSpan wraps MEOS C function union_spanset_span.
func UnionSpansetSpan(ss *SpanSet, s *Span) *SpanSet {
	res := C.union_spanset_span(ss._inner, s._inner)
	return &SpanSet{_inner: res}
}


// UnionSpansetSpanset wraps MEOS C function union_spanset_spanset.
func UnionSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) *SpanSet {
	res := C.union_spanset_spanset(ss1._inner, ss2._inner)
	return &SpanSet{_inner: res}
}


// UnionSpansetTimestamptz wraps MEOS C function union_spanset_timestamptz.
func UnionSpansetTimestamptz(ss *SpanSet, t int64) *SpanSet {
	res := C.union_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return &SpanSet{_inner: res}
}


// UnionTextSet wraps MEOS C function union_text_set.
func UnionTextSet(txt string, s *Set) *Set {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.union_text_set(_c_txt, s._inner)
	return &Set{_inner: res}
}


// UnionTimestamptzSet wraps MEOS C function union_timestamptz_set.
func UnionTimestamptzSet(t int64, s *Set) *Set {
	res := C.union_timestamptz_set(C.TimestampTz(t), s._inner)
	return &Set{_inner: res}
}


// UnionTimestamptzSpan wraps MEOS C function union_timestamptz_span.
func UnionTimestamptzSpan(t int64, s *Span) *SpanSet {
	res := C.union_timestamptz_span(C.TimestampTz(t), s._inner)
	return &SpanSet{_inner: res}
}


// UnionTimestamptzSpanset wraps MEOS C function union_timestamptz_spanset.
func UnionTimestamptzSpanset(t int64, ss *SpanSet) *SpanSet {
	res := C.union_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	return &SpanSet{_inner: res}
}


// DistanceBigintsetBigintset wraps MEOS C function distance_bigintset_bigintset.
func DistanceBigintsetBigintset(s1 *Set, s2 *Set) int64 {
	res := C.distance_bigintset_bigintset(s1._inner, s2._inner)
	return int64(res)
}


// DistanceBigintspanBigintspan wraps MEOS C function distance_bigintspan_bigintspan.
func DistanceBigintspanBigintspan(s1 *Span, s2 *Span) int64 {
	res := C.distance_bigintspan_bigintspan(s1._inner, s2._inner)
	return int64(res)
}


// DistanceBigintspansetBigintspan wraps MEOS C function distance_bigintspanset_bigintspan.
func DistanceBigintspansetBigintspan(ss *SpanSet, s *Span) int64 {
	res := C.distance_bigintspanset_bigintspan(ss._inner, s._inner)
	return int64(res)
}


// DistanceBigintspansetBigintspanset wraps MEOS C function distance_bigintspanset_bigintspanset.
func DistanceBigintspansetBigintspanset(ss1 *SpanSet, ss2 *SpanSet) int64 {
	res := C.distance_bigintspanset_bigintspanset(ss1._inner, ss2._inner)
	return int64(res)
}


// DistanceDatesetDateset wraps MEOS C function distance_dateset_dateset.
func DistanceDatesetDateset(s1 *Set, s2 *Set) int {
	res := C.distance_dateset_dateset(s1._inner, s2._inner)
	return int(res)
}


// DistanceDatespanDatespan wraps MEOS C function distance_datespan_datespan.
func DistanceDatespanDatespan(s1 *Span, s2 *Span) int {
	res := C.distance_datespan_datespan(s1._inner, s2._inner)
	return int(res)
}


// DistanceDatespansetDatespan wraps MEOS C function distance_datespanset_datespan.
func DistanceDatespansetDatespan(ss *SpanSet, s *Span) int {
	res := C.distance_datespanset_datespan(ss._inner, s._inner)
	return int(res)
}


// DistanceDatespansetDatespanset wraps MEOS C function distance_datespanset_datespanset.
func DistanceDatespansetDatespanset(ss1 *SpanSet, ss2 *SpanSet) int {
	res := C.distance_datespanset_datespanset(ss1._inner, ss2._inner)
	return int(res)
}


// DistanceFloatsetFloatset wraps MEOS C function distance_floatset_floatset.
func DistanceFloatsetFloatset(s1 *Set, s2 *Set) float64 {
	res := C.distance_floatset_floatset(s1._inner, s2._inner)
	return float64(res)
}


// DistanceFloatspanFloatspan wraps MEOS C function distance_floatspan_floatspan.
func DistanceFloatspanFloatspan(s1 *Span, s2 *Span) float64 {
	res := C.distance_floatspan_floatspan(s1._inner, s2._inner)
	return float64(res)
}


// DistanceFloatspansetFloatspan wraps MEOS C function distance_floatspanset_floatspan.
func DistanceFloatspansetFloatspan(ss *SpanSet, s *Span) float64 {
	res := C.distance_floatspanset_floatspan(ss._inner, s._inner)
	return float64(res)
}


// DistanceFloatspansetFloatspanset wraps MEOS C function distance_floatspanset_floatspanset.
func DistanceFloatspansetFloatspanset(ss1 *SpanSet, ss2 *SpanSet) float64 {
	res := C.distance_floatspanset_floatspanset(ss1._inner, ss2._inner)
	return float64(res)
}


// DistanceIntsetIntset wraps MEOS C function distance_intset_intset.
func DistanceIntsetIntset(s1 *Set, s2 *Set) int {
	res := C.distance_intset_intset(s1._inner, s2._inner)
	return int(res)
}


// DistanceIntspanIntspan wraps MEOS C function distance_intspan_intspan.
func DistanceIntspanIntspan(s1 *Span, s2 *Span) int {
	res := C.distance_intspan_intspan(s1._inner, s2._inner)
	return int(res)
}


// DistanceIntspansetIntspan wraps MEOS C function distance_intspanset_intspan.
func DistanceIntspansetIntspan(ss *SpanSet, s *Span) int {
	res := C.distance_intspanset_intspan(ss._inner, s._inner)
	return int(res)
}


// DistanceIntspansetIntspanset wraps MEOS C function distance_intspanset_intspanset.
func DistanceIntspansetIntspanset(ss1 *SpanSet, ss2 *SpanSet) int {
	res := C.distance_intspanset_intspanset(ss1._inner, ss2._inner)
	return int(res)
}


// DistanceSetBigint wraps MEOS C function distance_set_bigint.
func DistanceSetBigint(s *Set, i int64) int64 {
	res := C.distance_set_bigint(s._inner, C.int64(i))
	return int64(res)
}


// DistanceSetDate wraps MEOS C function distance_set_date.
func DistanceSetDate(s *Set, d int32) int {
	res := C.distance_set_date(s._inner, C.DateADT(d))
	return int(res)
}


// DistanceSetFloat wraps MEOS C function distance_set_float.
func DistanceSetFloat(s *Set, d float64) float64 {
	res := C.distance_set_float(s._inner, C.double(d))
	return float64(res)
}


// DistanceSetInt wraps MEOS C function distance_set_int.
func DistanceSetInt(s *Set, i int) int {
	res := C.distance_set_int(s._inner, C.int(i))
	return int(res)
}


// DistanceSetTimestamptz wraps MEOS C function distance_set_timestamptz.
func DistanceSetTimestamptz(s *Set, t int64) float64 {
	res := C.distance_set_timestamptz(s._inner, C.TimestampTz(t))
	return float64(res)
}


// DistanceSpanBigint wraps MEOS C function distance_span_bigint.
func DistanceSpanBigint(s *Span, i int64) int64 {
	res := C.distance_span_bigint(s._inner, C.int64(i))
	return int64(res)
}


// DistanceSpanDate wraps MEOS C function distance_span_date.
func DistanceSpanDate(s *Span, d int32) int {
	res := C.distance_span_date(s._inner, C.DateADT(d))
	return int(res)
}


// DistanceSpanFloat wraps MEOS C function distance_span_float.
func DistanceSpanFloat(s *Span, d float64) float64 {
	res := C.distance_span_float(s._inner, C.double(d))
	return float64(res)
}


// DistanceSpanInt wraps MEOS C function distance_span_int.
func DistanceSpanInt(s *Span, i int) int {
	res := C.distance_span_int(s._inner, C.int(i))
	return int(res)
}


// DistanceSpanTimestamptz wraps MEOS C function distance_span_timestamptz.
func DistanceSpanTimestamptz(s *Span, t int64) float64 {
	res := C.distance_span_timestamptz(s._inner, C.TimestampTz(t))
	return float64(res)
}


// DistanceSpansetBigint wraps MEOS C function distance_spanset_bigint.
func DistanceSpansetBigint(ss *SpanSet, i int64) int64 {
	res := C.distance_spanset_bigint(ss._inner, C.int64(i))
	return int64(res)
}


// DistanceSpansetDate wraps MEOS C function distance_spanset_date.
func DistanceSpansetDate(ss *SpanSet, d int32) int {
	res := C.distance_spanset_date(ss._inner, C.DateADT(d))
	return int(res)
}


// DistanceSpansetFloat wraps MEOS C function distance_spanset_float.
func DistanceSpansetFloat(ss *SpanSet, d float64) float64 {
	res := C.distance_spanset_float(ss._inner, C.double(d))
	return float64(res)
}


// DistanceSpansetInt wraps MEOS C function distance_spanset_int.
func DistanceSpansetInt(ss *SpanSet, i int) int {
	res := C.distance_spanset_int(ss._inner, C.int(i))
	return int(res)
}


// DistanceSpansetTimestamptz wraps MEOS C function distance_spanset_timestamptz.
func DistanceSpansetTimestamptz(ss *SpanSet, t int64) float64 {
	res := C.distance_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	return float64(res)
}


// DistanceTstzsetTstzset wraps MEOS C function distance_tstzset_tstzset.
func DistanceTstzsetTstzset(s1 *Set, s2 *Set) float64 {
	res := C.distance_tstzset_tstzset(s1._inner, s2._inner)
	return float64(res)
}


// DistanceTstzspanTstzspan wraps MEOS C function distance_tstzspan_tstzspan.
func DistanceTstzspanTstzspan(s1 *Span, s2 *Span) float64 {
	res := C.distance_tstzspan_tstzspan(s1._inner, s2._inner)
	return float64(res)
}


// DistanceTstzspansetTstzspan wraps MEOS C function distance_tstzspanset_tstzspan.
func DistanceTstzspansetTstzspan(ss *SpanSet, s *Span) float64 {
	res := C.distance_tstzspanset_tstzspan(ss._inner, s._inner)
	return float64(res)
}


// DistanceTstzspansetTstzspanset wraps MEOS C function distance_tstzspanset_tstzspanset.
func DistanceTstzspansetTstzspanset(ss1 *SpanSet, ss2 *SpanSet) float64 {
	res := C.distance_tstzspanset_tstzspanset(ss1._inner, ss2._inner)
	return float64(res)
}


// BigintExtentTransfn wraps MEOS C function bigint_extent_transfn.
func BigintExtentTransfn(state *Span, i int64) *Span {
	res := C.bigint_extent_transfn(state._inner, C.int64(i))
	return &Span{_inner: res}
}


// BigintUnionTransfn wraps MEOS C function bigint_union_transfn.
func BigintUnionTransfn(state *Set, i int64) *Set {
	res := C.bigint_union_transfn(state._inner, C.int64(i))
	return &Set{_inner: res}
}


// DateExtentTransfn wraps MEOS C function date_extent_transfn.
func DateExtentTransfn(state *Span, d int32) *Span {
	res := C.date_extent_transfn(state._inner, C.DateADT(d))
	return &Span{_inner: res}
}


// DateUnionTransfn wraps MEOS C function date_union_transfn.
func DateUnionTransfn(state *Set, d int32) *Set {
	res := C.date_union_transfn(state._inner, C.DateADT(d))
	return &Set{_inner: res}
}


// FloatExtentTransfn wraps MEOS C function float_extent_transfn.
func FloatExtentTransfn(state *Span, d float64) *Span {
	res := C.float_extent_transfn(state._inner, C.double(d))
	return &Span{_inner: res}
}


// FloatUnionTransfn wraps MEOS C function float_union_transfn.
func FloatUnionTransfn(state *Set, d float64) *Set {
	res := C.float_union_transfn(state._inner, C.double(d))
	return &Set{_inner: res}
}


// IntExtentTransfn wraps MEOS C function int_extent_transfn.
func IntExtentTransfn(state *Span, i int) *Span {
	res := C.int_extent_transfn(state._inner, C.int(i))
	return &Span{_inner: res}
}


// IntUnionTransfn wraps MEOS C function int_union_transfn.
func IntUnionTransfn(state *Set, i int32) *Set {
	res := C.int_union_transfn(state._inner, C.int32(i))
	return &Set{_inner: res}
}


// SetExtentTransfn wraps MEOS C function set_extent_transfn.
func SetExtentTransfn(state *Span, s *Set) *Span {
	res := C.set_extent_transfn(state._inner, s._inner)
	return &Span{_inner: res}
}


// SetUnionFinalfn wraps MEOS C function set_union_finalfn.
func SetUnionFinalfn(state *Set) *Set {
	res := C.set_union_finalfn(state._inner)
	return &Set{_inner: res}
}


// SetUnionTransfn wraps MEOS C function set_union_transfn.
func SetUnionTransfn(state *Set, s *Set) *Set {
	res := C.set_union_transfn(state._inner, s._inner)
	return &Set{_inner: res}
}


// SpanExtentTransfn wraps MEOS C function span_extent_transfn.
func SpanExtentTransfn(state *Span, s *Span) *Span {
	res := C.span_extent_transfn(state._inner, s._inner)
	return &Span{_inner: res}
}


// SpanUnionTransfn wraps MEOS C function span_union_transfn.
func SpanUnionTransfn(state *SpanSet, s *Span) *SpanSet {
	res := C.span_union_transfn(state._inner, s._inner)
	return &SpanSet{_inner: res}
}


// SpansetExtentTransfn wraps MEOS C function spanset_extent_transfn.
func SpansetExtentTransfn(state *Span, ss *SpanSet) *Span {
	res := C.spanset_extent_transfn(state._inner, ss._inner)
	return &Span{_inner: res}
}


// SpansetUnionFinalfn wraps MEOS C function spanset_union_finalfn.
func SpansetUnionFinalfn(state *SpanSet) *SpanSet {
	res := C.spanset_union_finalfn(state._inner)
	return &SpanSet{_inner: res}
}


// SpansetUnionTransfn wraps MEOS C function spanset_union_transfn.
func SpansetUnionTransfn(state *SpanSet, ss *SpanSet) *SpanSet {
	res := C.spanset_union_transfn(state._inner, ss._inner)
	return &SpanSet{_inner: res}
}


// TextUnionTransfn wraps MEOS C function text_union_transfn.
func TextUnionTransfn(state *Set, txt string) *Set {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.text_union_transfn(state._inner, _c_txt)
	return &Set{_inner: res}
}


// TimestamptzExtentTransfn wraps MEOS C function timestamptz_extent_transfn.
func TimestamptzExtentTransfn(state *Span, t int64) *Span {
	res := C.timestamptz_extent_transfn(state._inner, C.TimestampTz(t))
	return &Span{_inner: res}
}


// TimestamptzUnionTransfn wraps MEOS C function timestamptz_union_transfn.
func TimestamptzUnionTransfn(state *Set, t int64) *Set {
	res := C.timestamptz_union_transfn(state._inner, C.TimestampTz(t))
	return &Set{_inner: res}
}


// BigintGetBin wraps MEOS C function bigint_get_bin.
func BigintGetBin(value int64, vsize int64, vorigin int64) int64 {
	res := C.bigint_get_bin(C.int64(value), C.int64(vsize), C.int64(vorigin))
	return int64(res)
}


// BigintspanBins wraps MEOS C function bigintspan_bins.
func BigintspanBins(s *Span, vsize int64, vorigin int64) (*Span, int) {
	var _out_count C.int
	res := C.bigintspan_bins(s._inner, C.int64(vsize), C.int64(vorigin), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// BigintspansetBins wraps MEOS C function bigintspanset_bins.
func BigintspansetBins(ss *SpanSet, vsize int64, vorigin int64) (*Span, int) {
	var _out_count C.int
	res := C.bigintspanset_bins(ss._inner, C.int64(vsize), C.int64(vorigin), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// DateGetBin wraps MEOS C function date_get_bin.
func DateGetBin(d int32, duration timeutil.Timedelta, torigin int32) int32 {
	res := C.date_get_bin(C.DateADT(d), duration.Inner(), C.DateADT(torigin))
	return int32(res)
}


// DatespanBins wraps MEOS C function datespan_bins.
func DatespanBins(s *Span, duration timeutil.Timedelta, torigin int32) (*Span, int) {
	var _out_count C.int
	res := C.datespan_bins(s._inner, duration.Inner(), C.DateADT(torigin), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// DatespansetBins wraps MEOS C function datespanset_bins.
func DatespansetBins(ss *SpanSet, duration timeutil.Timedelta, torigin int32) (*Span, int) {
	var _out_count C.int
	res := C.datespanset_bins(ss._inner, duration.Inner(), C.DateADT(torigin), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// FloatGetBin wraps MEOS C function float_get_bin.
func FloatGetBin(value float64, vsize float64, vorigin float64) float64 {
	res := C.float_get_bin(C.double(value), C.double(vsize), C.double(vorigin))
	return float64(res)
}


// FloatspanBins wraps MEOS C function floatspan_bins.
func FloatspanBins(s *Span, vsize float64, vorigin float64) (*Span, int) {
	var _out_count C.int
	res := C.floatspan_bins(s._inner, C.double(vsize), C.double(vorigin), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// FloatspansetBins wraps MEOS C function floatspanset_bins.
func FloatspansetBins(ss *SpanSet, vsize float64, vorigin float64) (*Span, int) {
	var _out_count C.int
	res := C.floatspanset_bins(ss._inner, C.double(vsize), C.double(vorigin), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// IntGetBin wraps MEOS C function int_get_bin.
func IntGetBin(value int, vsize int, vorigin int) int {
	res := C.int_get_bin(C.int(value), C.int(vsize), C.int(vorigin))
	return int(res)
}


// IntspanBins wraps MEOS C function intspan_bins.
func IntspanBins(s *Span, vsize int, vorigin int) (*Span, int) {
	var _out_count C.int
	res := C.intspan_bins(s._inner, C.int(vsize), C.int(vorigin), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// IntspansetBins wraps MEOS C function intspanset_bins.
func IntspansetBins(ss *SpanSet, vsize int, vorigin int) (*Span, int) {
	var _out_count C.int
	res := C.intspanset_bins(ss._inner, C.int(vsize), C.int(vorigin), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// TimestamptzGetBin wraps MEOS C function timestamptz_get_bin.
func TimestamptzGetBin(t int64, duration timeutil.Timedelta, torigin int64) int64 {
	res := C.timestamptz_get_bin(C.TimestampTz(t), duration.Inner(), C.TimestampTz(torigin))
	return int64(res)
}


// TstzspanBins wraps MEOS C function tstzspan_bins.
func TstzspanBins(s *Span, duration timeutil.Timedelta, origin int64) (*Span, int) {
	var _out_count C.int
	res := C.tstzspan_bins(s._inner, duration.Inner(), C.TimestampTz(origin), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// TstzspansetBins wraps MEOS C function tstzspanset_bins.
func TstzspansetBins(ss *SpanSet, duration timeutil.Timedelta, torigin int64) (*Span, int) {
	var _out_count C.int
	res := C.tstzspanset_bins(ss._inner, duration.Inner(), C.TimestampTz(torigin), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// TBOXAsHexwkb wraps MEOS C function tbox_as_hexwkb.
func TBOXAsHexwkb(box *TBox, variant uint8) (string, uint) {
	var _out_size C.size_t
	res := C.tbox_as_hexwkb(box._inner, C.uint8_t(variant), &_out_size)
	return C.GoString(res), uint(_out_size)
}


// TBOXAsWKB wraps MEOS C function tbox_as_wkb.
func TBOXAsWKB(box *TBox, variant uint8) []uint8 {
	var _out_size_out C.size_t
	res := C.tbox_as_wkb(box._inner, C.uint8_t(variant), &_out_size_out)
	_n := int(_out_size_out)
	_slice := unsafe.Slice((*C.uint8_t)(unsafe.Pointer(res)), _n)
	_out := make([]uint8, _n)
	for _i, _e := range _slice {
		_out[_i] = uint8(_e)
	}
	return _out
}


// TBOXFromHexwkb wraps MEOS C function tbox_from_hexwkb.
func TBOXFromHexwkb(hexwkb string) *TBox {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	res := C.tbox_from_hexwkb(_c_hexwkb)
	return &TBox{_inner: res}
}


// TBOXFromWKB wraps MEOS C function tbox_from_wkb.
func TBOXFromWKB(wkb []byte) *TBox {
	res := C.tbox_from_wkb((*C.uint8_t)(unsafe.Pointer(&wkb[0])), C.size_t(len(wkb)))
	return &TBox{_inner: res}
}


// TBOXIn wraps MEOS C function tbox_in.
func TBOXIn(str string) *TBox {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tbox_in(_c_str)
	return &TBox{_inner: res}
}


// TBOXOut wraps MEOS C function tbox_out.
func TBOXOut(box *TBox, maxdd int) string {
	res := C.tbox_out(box._inner, C.int(maxdd))
	return C.GoString(res)
}


// FloatTimestamptzToTBOX wraps MEOS C function float_timestamptz_to_tbox.
func FloatTimestamptzToTBOX(d float64, t int64) *TBox {
	res := C.float_timestamptz_to_tbox(C.double(d), C.TimestampTz(t))
	return &TBox{_inner: res}
}


// FloatTstzspanToTBOX wraps MEOS C function float_tstzspan_to_tbox.
func FloatTstzspanToTBOX(d float64, s *Span) *TBox {
	res := C.float_tstzspan_to_tbox(C.double(d), s._inner)
	return &TBox{_inner: res}
}


// IntTimestamptzToTBOX wraps MEOS C function int_timestamptz_to_tbox.
func IntTimestamptzToTBOX(i int, t int64) *TBox {
	res := C.int_timestamptz_to_tbox(C.int(i), C.TimestampTz(t))
	return &TBox{_inner: res}
}


// IntTstzspanToTBOX wraps MEOS C function int_tstzspan_to_tbox.
func IntTstzspanToTBOX(i int, s *Span) *TBox {
	res := C.int_tstzspan_to_tbox(C.int(i), s._inner)
	return &TBox{_inner: res}
}


// NumspanTstzspanToTBOX wraps MEOS C function numspan_tstzspan_to_tbox.
func NumspanTstzspanToTBOX(span *Span, s *Span) *TBox {
	res := C.numspan_tstzspan_to_tbox(span._inner, s._inner)
	return &TBox{_inner: res}
}


// NumspanTimestamptzToTBOX wraps MEOS C function numspan_timestamptz_to_tbox.
func NumspanTimestamptzToTBOX(span *Span, t int64) *TBox {
	res := C.numspan_timestamptz_to_tbox(span._inner, C.TimestampTz(t))
	return &TBox{_inner: res}
}


// TBOXCopy wraps MEOS C function tbox_copy.
func TBOXCopy(box *TBox) *TBox {
	res := C.tbox_copy(box._inner)
	return &TBox{_inner: res}
}


// TBOXMake wraps MEOS C function tbox_make.
func TBOXMake(s *Span, p *Span) *TBox {
	res := C.tbox_make(s._inner, p._inner)
	return &TBox{_inner: res}
}


// FloatToTBOX wraps MEOS C function float_to_tbox.
func FloatToTBOX(d float64) *TBox {
	res := C.float_to_tbox(C.double(d))
	return &TBox{_inner: res}
}


// IntToTBOX wraps MEOS C function int_to_tbox.
func IntToTBOX(i int) *TBox {
	res := C.int_to_tbox(C.int(i))
	return &TBox{_inner: res}
}


// SetToTBOX wraps MEOS C function set_to_tbox.
func SetToTBOX(s *Set) *TBox {
	res := C.set_to_tbox(s._inner)
	return &TBox{_inner: res}
}


// SpanToTBOX wraps MEOS C function span_to_tbox.
func SpanToTBOX(s *Span) *TBox {
	res := C.span_to_tbox(s._inner)
	return &TBox{_inner: res}
}


// SpansetToTBOX wraps MEOS C function spanset_to_tbox.
func SpansetToTBOX(ss *SpanSet) *TBox {
	res := C.spanset_to_tbox(ss._inner)
	return &TBox{_inner: res}
}


// TBOXToIntspan wraps MEOS C function tbox_to_intspan.
func TBOXToIntspan(box *TBox) *Span {
	res := C.tbox_to_intspan(box._inner)
	return &Span{_inner: res}
}


// TBOXToFloatspan wraps MEOS C function tbox_to_floatspan.
func TBOXToFloatspan(box *TBox) *Span {
	res := C.tbox_to_floatspan(box._inner)
	return &Span{_inner: res}
}


// TBOXToTstzspan wraps MEOS C function tbox_to_tstzspan.
func TBOXToTstzspan(box *TBox) *Span {
	res := C.tbox_to_tstzspan(box._inner)
	return &Span{_inner: res}
}


// TimestamptzToTBOX wraps MEOS C function timestamptz_to_tbox.
func TimestamptzToTBOX(t int64) *TBox {
	res := C.timestamptz_to_tbox(C.TimestampTz(t))
	return &TBox{_inner: res}
}


// TBOXHash wraps MEOS C function tbox_hash.
func TBOXHash(box *TBox) uint32 {
	res := C.tbox_hash(box._inner)
	return uint32(res)
}


// TBOXHashExtended wraps MEOS C function tbox_hash_extended.
func TBOXHashExtended(box *TBox, seed uint64) uint64 {
	res := C.tbox_hash_extended(box._inner, C.uint64(seed))
	return uint64(res)
}


// TBOXHast wraps MEOS C function tbox_hast.
func TBOXHast(box *TBox) bool {
	res := C.tbox_hast(box._inner)
	return bool(res)
}


// TBOXHasx wraps MEOS C function tbox_hasx.
func TBOXHasx(box *TBox) bool {
	res := C.tbox_hasx(box._inner)
	return bool(res)
}


// TBOXTmax wraps MEOS C function tbox_tmax.
func TBOXTmax(box *TBox) (bool, int64) {
	var _out_result C.TimestampTz
	res := C.tbox_tmax(box._inner, &_out_result)
	return bool(res), int64({})
}


// TBOXTmaxInc wraps MEOS C function tbox_tmax_inc.
func TBOXTmaxInc(box *TBox) (bool, bool) {
	var _out_result C.bool
	res := C.tbox_tmax_inc(box._inner, &_out_result)
	return bool(res), bool({})
}


// TBOXTmin wraps MEOS C function tbox_tmin.
func TBOXTmin(box *TBox) (bool, int64) {
	var _out_result C.TimestampTz
	res := C.tbox_tmin(box._inner, &_out_result)
	return bool(res), int64({})
}


// TBOXTminInc wraps MEOS C function tbox_tmin_inc.
func TBOXTminInc(box *TBox) (bool, bool) {
	var _out_result C.bool
	res := C.tbox_tmin_inc(box._inner, &_out_result)
	return bool(res), bool({})
}


// TBOXXmax wraps MEOS C function tbox_xmax.
func TBOXXmax(box *TBox) (bool, float64) {
	var _out_result C.double
	res := C.tbox_xmax(box._inner, &_out_result)
	return bool(res), float64({})
}


// TBOXXmaxInc wraps MEOS C function tbox_xmax_inc.
func TBOXXmaxInc(box *TBox) (bool, bool) {
	var _out_result C.bool
	res := C.tbox_xmax_inc(box._inner, &_out_result)
	return bool(res), bool({})
}


// TBOXXmin wraps MEOS C function tbox_xmin.
func TBOXXmin(box *TBox) (bool, float64) {
	var _out_result C.double
	res := C.tbox_xmin(box._inner, &_out_result)
	return bool(res), float64({})
}


// TBOXXminInc wraps MEOS C function tbox_xmin_inc.
func TBOXXminInc(box *TBox) (bool, bool) {
	var _out_result C.bool
	res := C.tbox_xmin_inc(box._inner, &_out_result)
	return bool(res), bool({})
}


// TboxfloatXmax wraps MEOS C function tboxfloat_xmax.
func TboxfloatXmax(box *TBox) (bool, float64) {
	var _out_result C.double
	res := C.tboxfloat_xmax(box._inner, &_out_result)
	return bool(res), float64({})
}


// TboxfloatXmin wraps MEOS C function tboxfloat_xmin.
func TboxfloatXmin(box *TBox) (bool, float64) {
	var _out_result C.double
	res := C.tboxfloat_xmin(box._inner, &_out_result)
	return bool(res), float64({})
}


// TboxintXmax wraps MEOS C function tboxint_xmax.
func TboxintXmax(box *TBox) (bool, int) {
	var _out_result C.int
	res := C.tboxint_xmax(box._inner, &_out_result)
	return bool(res), int({})
}


// TboxintXmin wraps MEOS C function tboxint_xmin.
func TboxintXmin(box *TBox) (bool, int) {
	var _out_result C.int
	res := C.tboxint_xmin(box._inner, &_out_result)
	return bool(res), int({})
}


// TBOXExpandTime wraps MEOS C function tbox_expand_time.
func TBOXExpandTime(box *TBox, interv timeutil.Timedelta) *TBox {
	res := C.tbox_expand_time(box._inner, interv.Inner())
	return &TBox{_inner: res}
}


// TBOXRound wraps MEOS C function tbox_round.
func TBOXRound(box *TBox, maxdd int) *TBox {
	res := C.tbox_round(box._inner, C.int(maxdd))
	return &TBox{_inner: res}
}


// TBOXShiftScaleTime wraps MEOS C function tbox_shift_scale_time.
func TBOXShiftScaleTime(box *TBox, shift timeutil.Timedelta, duration timeutil.Timedelta) *TBox {
	res := C.tbox_shift_scale_time(box._inner, shift.Inner(), duration.Inner())
	return &TBox{_inner: res}
}


// TfloatboxExpand wraps MEOS C function tfloatbox_expand.
func TfloatboxExpand(box *TBox, d float64) *TBox {
	res := C.tfloatbox_expand(box._inner, C.double(d))
	return &TBox{_inner: res}
}


// TfloatboxShiftScale wraps MEOS C function tfloatbox_shift_scale.
func TfloatboxShiftScale(box *TBox, shift float64, width float64, hasshift bool, haswidth bool) *TBox {
	res := C.tfloatbox_shift_scale(box._inner, C.double(shift), C.double(width), C.bool(hasshift), C.bool(haswidth))
	return &TBox{_inner: res}
}


// TintboxExpand wraps MEOS C function tintbox_expand.
func TintboxExpand(box *TBox, i int) *TBox {
	res := C.tintbox_expand(box._inner, C.int(i))
	return &TBox{_inner: res}
}


// TintboxShiftScale wraps MEOS C function tintbox_shift_scale.
func TintboxShiftScale(box *TBox, shift int, width int, hasshift bool, haswidth bool) *TBox {
	res := C.tintbox_shift_scale(box._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	return &TBox{_inner: res}
}


// UnionTBOXTBOX wraps MEOS C function union_tbox_tbox.
func UnionTBOXTBOX(box1 *TBox, box2 *TBox, strict bool) *TBox {
	res := C.union_tbox_tbox(box1._inner, box2._inner, C.bool(strict))
	return &TBox{_inner: res}
}


// IntersectionTBOXTBOX wraps MEOS C function intersection_tbox_tbox.
func IntersectionTBOXTBOX(box1 *TBox, box2 *TBox) *TBox {
	res := C.intersection_tbox_tbox(box1._inner, box2._inner)
	return &TBox{_inner: res}
}


// AdjacentTBOXTBOX wraps MEOS C function adjacent_tbox_tbox.
func AdjacentTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	res := C.adjacent_tbox_tbox(box1._inner, box2._inner)
	return bool(res)
}


// ContainedTBOXTBOX wraps MEOS C function contained_tbox_tbox.
func ContainedTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	res := C.contained_tbox_tbox(box1._inner, box2._inner)
	return bool(res)
}


// ContainsTBOXTBOX wraps MEOS C function contains_tbox_tbox.
func ContainsTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	res := C.contains_tbox_tbox(box1._inner, box2._inner)
	return bool(res)
}


// OverlapsTBOXTBOX wraps MEOS C function overlaps_tbox_tbox.
func OverlapsTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	res := C.overlaps_tbox_tbox(box1._inner, box2._inner)
	return bool(res)
}


// SameTBOXTBOX wraps MEOS C function same_tbox_tbox.
func SameTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	res := C.same_tbox_tbox(box1._inner, box2._inner)
	return bool(res)
}


// AfterTBOXTBOX wraps MEOS C function after_tbox_tbox.
func AfterTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	res := C.after_tbox_tbox(box1._inner, box2._inner)
	return bool(res)
}


// BeforeTBOXTBOX wraps MEOS C function before_tbox_tbox.
func BeforeTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	res := C.before_tbox_tbox(box1._inner, box2._inner)
	return bool(res)
}


// LeftTBOXTBOX wraps MEOS C function left_tbox_tbox.
func LeftTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	res := C.left_tbox_tbox(box1._inner, box2._inner)
	return bool(res)
}


// OverafterTBOXTBOX wraps MEOS C function overafter_tbox_tbox.
func OverafterTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	res := C.overafter_tbox_tbox(box1._inner, box2._inner)
	return bool(res)
}


// OverbeforeTBOXTBOX wraps MEOS C function overbefore_tbox_tbox.
func OverbeforeTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	res := C.overbefore_tbox_tbox(box1._inner, box2._inner)
	return bool(res)
}


// OverleftTBOXTBOX wraps MEOS C function overleft_tbox_tbox.
func OverleftTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	res := C.overleft_tbox_tbox(box1._inner, box2._inner)
	return bool(res)
}


// OverrightTBOXTBOX wraps MEOS C function overright_tbox_tbox.
func OverrightTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	res := C.overright_tbox_tbox(box1._inner, box2._inner)
	return bool(res)
}


// RightTBOXTBOX wraps MEOS C function right_tbox_tbox.
func RightTBOXTBOX(box1 *TBox, box2 *TBox) bool {
	res := C.right_tbox_tbox(box1._inner, box2._inner)
	return bool(res)
}


// TBOXCmp wraps MEOS C function tbox_cmp.
func TBOXCmp(box1 *TBox, box2 *TBox) int {
	res := C.tbox_cmp(box1._inner, box2._inner)
	return int(res)
}


// TBOXEq wraps MEOS C function tbox_eq.
func TBOXEq(box1 *TBox, box2 *TBox) bool {
	res := C.tbox_eq(box1._inner, box2._inner)
	return bool(res)
}


// TBOXGe wraps MEOS C function tbox_ge.
func TBOXGe(box1 *TBox, box2 *TBox) bool {
	res := C.tbox_ge(box1._inner, box2._inner)
	return bool(res)
}


// TBOXGt wraps MEOS C function tbox_gt.
func TBOXGt(box1 *TBox, box2 *TBox) bool {
	res := C.tbox_gt(box1._inner, box2._inner)
	return bool(res)
}


// TBOXLe wraps MEOS C function tbox_le.
func TBOXLe(box1 *TBox, box2 *TBox) bool {
	res := C.tbox_le(box1._inner, box2._inner)
	return bool(res)
}


// TBOXLt wraps MEOS C function tbox_lt.
func TBOXLt(box1 *TBox, box2 *TBox) bool {
	res := C.tbox_lt(box1._inner, box2._inner)
	return bool(res)
}


// TBOXNe wraps MEOS C function tbox_ne.
func TBOXNe(box1 *TBox, box2 *TBox) bool {
	res := C.tbox_ne(box1._inner, box2._inner)
	return bool(res)
}


// TboolFromMFJSON wraps MEOS C function tbool_from_mfjson.
func TboolFromMFJSON(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tbool_from_mfjson(_c_str)
	return CreateTemporal(res)
}


// TboolIn wraps MEOS C function tbool_in.
func TboolIn(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tbool_in(_c_str)
	return CreateTemporal(res)
}


// TboolOut wraps MEOS C function tbool_out.
func TboolOut(temp Temporal) string {
	res := C.tbool_out(temp.Inner())
	return C.GoString(res)
}


// TemporalAsHexwkb wraps MEOS C function temporal_as_hexwkb.
func TemporalAsHexwkb(temp Temporal, variant uint8) (string, uint) {
	var _out_size_out C.size_t
	res := C.temporal_as_hexwkb(temp.Inner(), C.uint8_t(variant), &_out_size_out)
	return C.GoString(res), uint(_out_size_out)
}


// TemporalAsMFJSON wraps MEOS C function temporal_as_mfjson.
func TemporalAsMFJSON(temp Temporal, with_bbox bool, flags int, precision int, srs string) string {
	_c_srs := C.CString(srs)
	defer C.free(unsafe.Pointer(_c_srs))
	res := C.temporal_as_mfjson(temp.Inner(), C.bool(with_bbox), C.int(flags), C.int(precision), _c_srs)
	return C.GoString(res)
}


// TemporalAsWKB wraps MEOS C function temporal_as_wkb.
func TemporalAsWKB(temp Temporal, variant uint8) []uint8 {
	var _out_size_out C.size_t
	res := C.temporal_as_wkb(temp.Inner(), C.uint8_t(variant), &_out_size_out)
	_n := int(_out_size_out)
	_slice := unsafe.Slice((*C.uint8_t)(unsafe.Pointer(res)), _n)
	_out := make([]uint8, _n)
	for _i, _e := range _slice {
		_out[_i] = uint8(_e)
	}
	return _out
}


// TemporalFromHexwkb wraps MEOS C function temporal_from_hexwkb.
func TemporalFromHexwkb(hexwkb string) Temporal {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	res := C.temporal_from_hexwkb(_c_hexwkb)
	return CreateTemporal(res)
}


// TemporalFromWKB wraps MEOS C function temporal_from_wkb.
func TemporalFromWKB(wkb []byte) Temporal {
	res := C.temporal_from_wkb((*C.uint8_t)(unsafe.Pointer(&wkb[0])), C.size_t(len(wkb)))
	return CreateTemporal(res)
}


// TfloatFromMFJSON wraps MEOS C function tfloat_from_mfjson.
func TfloatFromMFJSON(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tfloat_from_mfjson(_c_str)
	return CreateTemporal(res)
}


// TfloatIn wraps MEOS C function tfloat_in.
func TfloatIn(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tfloat_in(_c_str)
	return CreateTemporal(res)
}


// TfloatOut wraps MEOS C function tfloat_out.
func TfloatOut(temp Temporal, maxdd int) string {
	res := C.tfloat_out(temp.Inner(), C.int(maxdd))
	return C.GoString(res)
}


// TintFromMFJSON wraps MEOS C function tint_from_mfjson.
func TintFromMFJSON(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tint_from_mfjson(_c_str)
	return CreateTemporal(res)
}


// TintIn wraps MEOS C function tint_in.
func TintIn(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tint_in(_c_str)
	return CreateTemporal(res)
}


// TintOut wraps MEOS C function tint_out.
func TintOut(temp Temporal) string {
	res := C.tint_out(temp.Inner())
	return C.GoString(res)
}


// TtextFromMFJSON wraps MEOS C function ttext_from_mfjson.
func TtextFromMFJSON(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.ttext_from_mfjson(_c_str)
	return CreateTemporal(res)
}


// TtextIn wraps MEOS C function ttext_in.
func TtextIn(str string) Temporal {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.ttext_in(_c_str)
	return CreateTemporal(res)
}


// TtextOut wraps MEOS C function ttext_out.
func TtextOut(temp Temporal) string {
	res := C.ttext_out(temp.Inner())
	return C.GoString(res)
}


// TboolFromBaseTemp wraps MEOS C function tbool_from_base_temp.
func TboolFromBaseTemp(b bool, temp Temporal) Temporal {
	res := C.tbool_from_base_temp(C.bool(b), temp.Inner())
	return CreateTemporal(res)
}


// TboolinstMake wraps MEOS C function tboolinst_make.
func TboolinstMake(b bool, t int64) TInstant {
	res := C.tboolinst_make(C.bool(b), C.TimestampTz(t))
	return TInstant{_inner: res}
}


// TboolseqFromBaseTstzset wraps MEOS C function tboolseq_from_base_tstzset.
func TboolseqFromBaseTstzset(b bool, s *Set) TSequence {
	res := C.tboolseq_from_base_tstzset(C.bool(b), s._inner)
	return TSequence{_inner: res}
}


// TboolseqFromBaseTstzspan wraps MEOS C function tboolseq_from_base_tstzspan.
func TboolseqFromBaseTstzspan(b bool, s *Span) TSequence {
	res := C.tboolseq_from_base_tstzspan(C.bool(b), s._inner)
	return TSequence{_inner: res}
}


// TboolseqsetFromBaseTstzspanset wraps MEOS C function tboolseqset_from_base_tstzspanset.
func TboolseqsetFromBaseTstzspanset(b bool, ss *SpanSet) TSequenceSet {
	res := C.tboolseqset_from_base_tstzspanset(C.bool(b), ss._inner)
	return TSequenceSet{_inner: res}
}


// TemporalCopy wraps MEOS C function temporal_copy.
func TemporalCopy(temp Temporal) Temporal {
	res := C.temporal_copy(temp.Inner())
	return CreateTemporal(res)
}


// TfloatFromBaseTemp wraps MEOS C function tfloat_from_base_temp.
func TfloatFromBaseTemp(d float64, temp Temporal) Temporal {
	res := C.tfloat_from_base_temp(C.double(d), temp.Inner())
	return CreateTemporal(res)
}


// TfloatinstMake wraps MEOS C function tfloatinst_make.
func TfloatinstMake(d float64, t int64) TInstant {
	res := C.tfloatinst_make(C.double(d), C.TimestampTz(t))
	return TInstant{_inner: res}
}


// TfloatseqFromBaseTstzset wraps MEOS C function tfloatseq_from_base_tstzset.
func TfloatseqFromBaseTstzset(d float64, s *Set) TSequence {
	res := C.tfloatseq_from_base_tstzset(C.double(d), s._inner)
	return TSequence{_inner: res}
}


// TfloatseqFromBaseTstzspan wraps MEOS C function tfloatseq_from_base_tstzspan.
func TfloatseqFromBaseTstzspan(d float64, s *Span, interp Interpolation) TSequence {
	res := C.tfloatseq_from_base_tstzspan(C.double(d), s._inner, C.interpType(interp))
	return TSequence{_inner: res}
}


// TfloatseqsetFromBaseTstzspanset wraps MEOS C function tfloatseqset_from_base_tstzspanset.
func TfloatseqsetFromBaseTstzspanset(d float64, ss *SpanSet, interp Interpolation) TSequenceSet {
	res := C.tfloatseqset_from_base_tstzspanset(C.double(d), ss._inner, C.interpType(interp))
	return TSequenceSet{_inner: res}
}


// TintFromBaseTemp wraps MEOS C function tint_from_base_temp.
func TintFromBaseTemp(i int, temp Temporal) Temporal {
	res := C.tint_from_base_temp(C.int(i), temp.Inner())
	return CreateTemporal(res)
}


// TintinstMake wraps MEOS C function tintinst_make.
func TintinstMake(i int, t int64) TInstant {
	res := C.tintinst_make(C.int(i), C.TimestampTz(t))
	return TInstant{_inner: res}
}


// TintseqFromBaseTstzset wraps MEOS C function tintseq_from_base_tstzset.
func TintseqFromBaseTstzset(i int, s *Set) TSequence {
	res := C.tintseq_from_base_tstzset(C.int(i), s._inner)
	return TSequence{_inner: res}
}


// TintseqFromBaseTstzspan wraps MEOS C function tintseq_from_base_tstzspan.
func TintseqFromBaseTstzspan(i int, s *Span) TSequence {
	res := C.tintseq_from_base_tstzspan(C.int(i), s._inner)
	return TSequence{_inner: res}
}


// TintseqsetFromBaseTstzspanset wraps MEOS C function tintseqset_from_base_tstzspanset.
func TintseqsetFromBaseTstzspanset(i int, ss *SpanSet) TSequenceSet {
	res := C.tintseqset_from_base_tstzspanset(C.int(i), ss._inner)
	return TSequenceSet{_inner: res}
}


// TsequenceMake wraps MEOS C function tsequence_make.
func TsequenceMake(instants []TInstant, lower_inc bool, upper_inc bool, interp Interpolation, normalize bool) TSequence {
	_c_instants := make([]*C.TInstant, len(instants))
	for _i, _v := range instants { _c_instants[_i] = _v._inner }
	res := C.tsequence_make((**C.TInstant)(unsafe.Pointer(&_c_instants[0])), C.int(len(instants)), C.bool(lower_inc), C.bool(upper_inc), C.interpType(interp), C.bool(normalize))
	return TSequence{_inner: res}
}


// TsequencesetMake wraps MEOS C function tsequenceset_make.
func TsequencesetMake(sequences []TSequence, normalize bool) TSequenceSet {
	_c_sequences := make([]*C.TSequence, len(sequences))
	for _i, _v := range sequences { _c_sequences[_i] = _v._inner }
	res := C.tsequenceset_make((**C.TSequence)(unsafe.Pointer(&_c_sequences[0])), C.int(len(sequences)), C.bool(normalize))
	return TSequenceSet{_inner: res}
}


// TsequencesetMakeGaps wraps MEOS C function tsequenceset_make_gaps.
func TsequencesetMakeGaps(instants []TInstant, interp Interpolation, maxt timeutil.Timedelta, maxdist float64) TSequenceSet {
	_c_instants := make([]*C.TInstant, len(instants))
	for _i, _v := range instants { _c_instants[_i] = _v._inner }
	res := C.tsequenceset_make_gaps((**C.TInstant)(unsafe.Pointer(&_c_instants[0])), C.int(len(instants)), C.interpType(interp), maxt.Inner(), C.double(maxdist))
	return TSequenceSet{_inner: res}
}


// TtextFromBaseTemp wraps MEOS C function ttext_from_base_temp.
func TtextFromBaseTemp(txt string, temp Temporal) Temporal {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ttext_from_base_temp(_c_txt, temp.Inner())
	return CreateTemporal(res)
}


// TtextinstMake wraps MEOS C function ttextinst_make.
func TtextinstMake(txt string, t int64) TInstant {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ttextinst_make(_c_txt, C.TimestampTz(t))
	return TInstant{_inner: res}
}


// TtextseqFromBaseTstzset wraps MEOS C function ttextseq_from_base_tstzset.
func TtextseqFromBaseTstzset(txt string, s *Set) TSequence {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ttextseq_from_base_tstzset(_c_txt, s._inner)
	return TSequence{_inner: res}
}


// TtextseqFromBaseTstzspan wraps MEOS C function ttextseq_from_base_tstzspan.
func TtextseqFromBaseTstzspan(txt string, s *Span) TSequence {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ttextseq_from_base_tstzspan(_c_txt, s._inner)
	return TSequence{_inner: res}
}


// TtextseqsetFromBaseTstzspanset wraps MEOS C function ttextseqset_from_base_tstzspanset.
func TtextseqsetFromBaseTstzspanset(txt string, ss *SpanSet) TSequenceSet {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ttextseqset_from_base_tstzspanset(_c_txt, ss._inner)
	return TSequenceSet{_inner: res}
}


// TboolToTint wraps MEOS C function tbool_to_tint.
func TboolToTint(temp Temporal) Temporal {
	res := C.tbool_to_tint(temp.Inner())
	return CreateTemporal(res)
}


// TemporalToTstzspan wraps MEOS C function temporal_to_tstzspan.
func TemporalToTstzspan(temp Temporal) *Span {
	res := C.temporal_to_tstzspan(temp.Inner())
	return &Span{_inner: res}
}


// TfloatToTint wraps MEOS C function tfloat_to_tint.
func TfloatToTint(temp Temporal) Temporal {
	res := C.tfloat_to_tint(temp.Inner())
	return CreateTemporal(res)
}


// TintToTfloat wraps MEOS C function tint_to_tfloat.
func TintToTfloat(temp Temporal) Temporal {
	res := C.tint_to_tfloat(temp.Inner())
	return CreateTemporal(res)
}


// TnumberToSpan wraps MEOS C function tnumber_to_span.
func TnumberToSpan(temp Temporal) *Span {
	res := C.tnumber_to_span(temp.Inner())
	return &Span{_inner: res}
}


// TnumberToTBOX wraps MEOS C function tnumber_to_tbox.
func TnumberToTBOX(temp Temporal) *TBox {
	res := C.tnumber_to_tbox(temp.Inner())
	return &TBox{_inner: res}
}


// TboolEndValue wraps MEOS C function tbool_end_value.
func TboolEndValue(temp Temporal) bool {
	res := C.tbool_end_value(temp.Inner())
	return bool(res)
}


// TboolStartValue wraps MEOS C function tbool_start_value.
func TboolStartValue(temp Temporal) bool {
	res := C.tbool_start_value(temp.Inner())
	return bool(res)
}


// TboolValueAtTimestamptz wraps MEOS C function tbool_value_at_timestamptz.
func TboolValueAtTimestamptz(temp Temporal, t int64, strict bool) (bool, bool) {
	var _out_value C.bool
	res := C.tbool_value_at_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(strict), &_out_value)
	return bool(res), bool({})
}


// TboolValueN wraps MEOS C function tbool_value_n.
func TboolValueN(temp Temporal, n int) (bool, bool) {
	var _out_result C.bool
	res := C.tbool_value_n(temp.Inner(), C.int(n), &_out_result)
	return bool(res), bool({})
}


// TboolValues wraps MEOS C function tbool_values.
func TboolValues(temp Temporal) []bool {
	var _out_count C.int
	res := C.tbool_values(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((*C.bool)(unsafe.Pointer(res)), _n)
	_out := make([]bool, _n)
	for _i, _e := range _slice {
		_out[_i] = bool(_e)
	}
	return _out
}


// TemporalDuration wraps MEOS C function temporal_duration.
func TemporalDuration(temp Temporal, boundspan bool) timeutil.Timedelta {
	res := C.temporal_duration(temp.Inner(), C.bool(boundspan))
	return IntervalToTimeDelta(res)
}


// TemporalEndInstant wraps MEOS C function temporal_end_instant.
func TemporalEndInstant(temp Temporal) TInstant {
	res := C.temporal_end_instant(temp.Inner())
	return TInstant{_inner: res}
}


// TemporalEndSequence wraps MEOS C function temporal_end_sequence.
func TemporalEndSequence(temp Temporal) TSequence {
	res := C.temporal_end_sequence(temp.Inner())
	return TSequence{_inner: res}
}


// TemporalEndTimestamptz wraps MEOS C function temporal_end_timestamptz.
func TemporalEndTimestamptz(temp Temporal) int64 {
	res := C.temporal_end_timestamptz(temp.Inner())
	return int64(res)
}


// TemporalHash wraps MEOS C function temporal_hash.
func TemporalHash(temp Temporal) uint32 {
	res := C.temporal_hash(temp.Inner())
	return uint32(res)
}


// TemporalInstantN wraps MEOS C function temporal_instant_n.
func TemporalInstantN(temp Temporal, n int) TInstant {
	res := C.temporal_instant_n(temp.Inner(), C.int(n))
	return TInstant{_inner: res}
}


// TemporalInstants wraps MEOS C function temporal_instants.
func TemporalInstants(temp Temporal) []TInstant {
	var _out_count C.int
	res := C.temporal_instants(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.TInstant)(unsafe.Pointer(res)), _n)
	_out := make([]TInstant, _n)
	for _i, _e := range _slice {
		_out[_i] = TInstant{_inner: _e}
	}
	return _out
}


// TemporalInterp wraps MEOS C function temporal_interp.
func TemporalInterp(temp Temporal) string {
	res := C.temporal_interp(temp.Inner())
	return C.GoString(res)
}


// TemporalLowerInc wraps MEOS C function temporal_lower_inc.
func TemporalLowerInc(temp Temporal) bool {
	res := C.temporal_lower_inc(temp.Inner())
	return bool(res)
}


// TemporalMaxInstant wraps MEOS C function temporal_max_instant.
func TemporalMaxInstant(temp Temporal) TInstant {
	res := C.temporal_max_instant(temp.Inner())
	return TInstant{_inner: res}
}


// TemporalMinInstant wraps MEOS C function temporal_min_instant.
func TemporalMinInstant(temp Temporal) TInstant {
	res := C.temporal_min_instant(temp.Inner())
	return TInstant{_inner: res}
}


// TemporalNumInstants wraps MEOS C function temporal_num_instants.
func TemporalNumInstants(temp Temporal) int {
	res := C.temporal_num_instants(temp.Inner())
	return int(res)
}


// TemporalNumSequences wraps MEOS C function temporal_num_sequences.
func TemporalNumSequences(temp Temporal) int {
	res := C.temporal_num_sequences(temp.Inner())
	return int(res)
}


// TemporalNumTimestamps wraps MEOS C function temporal_num_timestamps.
func TemporalNumTimestamps(temp Temporal) int {
	res := C.temporal_num_timestamps(temp.Inner())
	return int(res)
}


// TemporalSegmDuration wraps MEOS C function temporal_segm_duration.
func TemporalSegmDuration(temp Temporal, duration timeutil.Timedelta, atleast bool, strict bool) TSequenceSet {
	res := C.temporal_segm_duration(temp.Inner(), duration.Inner(), C.bool(atleast), C.bool(strict))
	return TSequenceSet{_inner: res}
}


// TemporalSegments wraps MEOS C function temporal_segments.
func TemporalSegments(temp Temporal) []TSequence {
	var _out_count C.int
	res := C.temporal_segments(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.TSequence)(unsafe.Pointer(res)), _n)
	_out := make([]TSequence, _n)
	for _i, _e := range _slice {
		_out[_i] = TSequence{_inner: _e}
	}
	return _out
}


// TemporalSequenceN wraps MEOS C function temporal_sequence_n.
func TemporalSequenceN(temp Temporal, i int) TSequence {
	res := C.temporal_sequence_n(temp.Inner(), C.int(i))
	return TSequence{_inner: res}
}


// TemporalSequences wraps MEOS C function temporal_sequences.
func TemporalSequences(temp Temporal) []TSequence {
	var _out_count C.int
	res := C.temporal_sequences(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.TSequence)(unsafe.Pointer(res)), _n)
	_out := make([]TSequence, _n)
	for _i, _e := range _slice {
		_out[_i] = TSequence{_inner: _e}
	}
	return _out
}


// TemporalStartInstant wraps MEOS C function temporal_start_instant.
func TemporalStartInstant(temp Temporal) TInstant {
	res := C.temporal_start_instant(temp.Inner())
	return TInstant{_inner: res}
}


// TemporalStartSequence wraps MEOS C function temporal_start_sequence.
func TemporalStartSequence(temp Temporal) TSequence {
	res := C.temporal_start_sequence(temp.Inner())
	return TSequence{_inner: res}
}


// TemporalStartTimestamptz wraps MEOS C function temporal_start_timestamptz.
func TemporalStartTimestamptz(temp Temporal) int64 {
	res := C.temporal_start_timestamptz(temp.Inner())
	return int64(res)
}


// TemporalStops wraps MEOS C function temporal_stops.
func TemporalStops(temp Temporal, maxdist float64, minduration timeutil.Timedelta) TSequenceSet {
	res := C.temporal_stops(temp.Inner(), C.double(maxdist), minduration.Inner())
	return TSequenceSet{_inner: res}
}


// TemporalSubtype wraps MEOS C function temporal_subtype.
func TemporalSubtype(temp Temporal) string {
	res := C.temporal_subtype(temp.Inner())
	return C.GoString(res)
}


// TemporalTime wraps MEOS C function temporal_time.
func TemporalTime(temp Temporal) *SpanSet {
	res := C.temporal_time(temp.Inner())
	return &SpanSet{_inner: res}
}


// TemporalTimestamps wraps MEOS C function temporal_timestamps.
func TemporalTimestamps(temp Temporal) []int64 {
	var _out_count C.int
	res := C.temporal_timestamps(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((*C.TimestampTz)(unsafe.Pointer(res)), _n)
	_out := make([]int64, _n)
	for _i, _e := range _slice {
		_out[_i] = int64(_e)
	}
	return _out
}


// TemporalTimestamptzN wraps MEOS C function temporal_timestamptz_n.
func TemporalTimestamptzN(temp Temporal, n int) (bool, int64) {
	var _out_result C.TimestampTz
	res := C.temporal_timestamptz_n(temp.Inner(), C.int(n), &_out_result)
	return bool(res), int64({})
}


// TemporalUpperInc wraps MEOS C function temporal_upper_inc.
func TemporalUpperInc(temp Temporal) bool {
	res := C.temporal_upper_inc(temp.Inner())
	return bool(res)
}


// TfloatAvgValue wraps MEOS C function tfloat_avg_value.
func TfloatAvgValue(temp Temporal) float64 {
	res := C.tfloat_avg_value(temp.Inner())
	return float64(res)
}


// TfloatEndValue wraps MEOS C function tfloat_end_value.
func TfloatEndValue(temp Temporal) float64 {
	res := C.tfloat_end_value(temp.Inner())
	return float64(res)
}


// TfloatMinValue wraps MEOS C function tfloat_min_value.
func TfloatMinValue(temp Temporal) float64 {
	res := C.tfloat_min_value(temp.Inner())
	return float64(res)
}


// TfloatMaxValue wraps MEOS C function tfloat_max_value.
func TfloatMaxValue(temp Temporal) float64 {
	res := C.tfloat_max_value(temp.Inner())
	return float64(res)
}


// TfloatStartValue wraps MEOS C function tfloat_start_value.
func TfloatStartValue(temp Temporal) float64 {
	res := C.tfloat_start_value(temp.Inner())
	return float64(res)
}


// TfloatValueAtTimestamptz wraps MEOS C function tfloat_value_at_timestamptz.
func TfloatValueAtTimestamptz(temp Temporal, t int64, strict bool) (bool, float64) {
	var _out_value C.double
	res := C.tfloat_value_at_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(strict), &_out_value)
	return bool(res), float64({})
}


// TfloatValueN wraps MEOS C function tfloat_value_n.
func TfloatValueN(temp Temporal, n int) (bool, float64) {
	var _out_result C.double
	res := C.tfloat_value_n(temp.Inner(), C.int(n), &_out_result)
	return bool(res), float64({})
}


// TfloatValues wraps MEOS C function tfloat_values.
func TfloatValues(temp Temporal) []float64 {
	var _out_count C.int
	res := C.tfloat_values(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((*C.double)(unsafe.Pointer(res)), _n)
	_out := make([]float64, _n)
	for _i, _e := range _slice {
		_out[_i] = float64(_e)
	}
	return _out
}


// TintEndValue wraps MEOS C function tint_end_value.
func TintEndValue(temp Temporal) int {
	res := C.tint_end_value(temp.Inner())
	return int(res)
}


// TintMaxValue wraps MEOS C function tint_max_value.
func TintMaxValue(temp Temporal) int {
	res := C.tint_max_value(temp.Inner())
	return int(res)
}


// TintMinValue wraps MEOS C function tint_min_value.
func TintMinValue(temp Temporal) int {
	res := C.tint_min_value(temp.Inner())
	return int(res)
}


// TintStartValue wraps MEOS C function tint_start_value.
func TintStartValue(temp Temporal) int {
	res := C.tint_start_value(temp.Inner())
	return int(res)
}


// TintValueAtTimestamptz wraps MEOS C function tint_value_at_timestamptz.
func TintValueAtTimestamptz(temp Temporal, t int64, strict bool) (bool, int) {
	var _out_value C.int
	res := C.tint_value_at_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(strict), &_out_value)
	return bool(res), int({})
}


// TintValueN wraps MEOS C function tint_value_n.
func TintValueN(temp Temporal, n int) (bool, int) {
	var _out_result C.int
	res := C.tint_value_n(temp.Inner(), C.int(n), &_out_result)
	return bool(res), int({})
}


// TintValues wraps MEOS C function tint_values.
func TintValues(temp Temporal) []int {
	var _out_count C.int
	res := C.tint_values(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((*C.int)(unsafe.Pointer(res)), _n)
	_out := make([]int, _n)
	for _i, _e := range _slice {
		_out[_i] = int(_e)
	}
	return _out
}


// TnumberAvgValue wraps MEOS C function tnumber_avg_value.
func TnumberAvgValue(temp Temporal) float64 {
	res := C.tnumber_avg_value(temp.Inner())
	return float64(res)
}


// TnumberIntegral wraps MEOS C function tnumber_integral.
func TnumberIntegral(temp Temporal) float64 {
	res := C.tnumber_integral(temp.Inner())
	return float64(res)
}


// TnumberTwavg wraps MEOS C function tnumber_twavg.
func TnumberTwavg(temp Temporal) float64 {
	res := C.tnumber_twavg(temp.Inner())
	return float64(res)
}


// TnumberValuespans wraps MEOS C function tnumber_valuespans.
func TnumberValuespans(temp Temporal) *SpanSet {
	res := C.tnumber_valuespans(temp.Inner())
	return &SpanSet{_inner: res}
}


// TtextEndValue wraps MEOS C function ttext_end_value.
func TtextEndValue(temp Temporal) string {
	res := C.ttext_end_value(temp.Inner())
	return text2cstring(res)
}


// TtextMaxValue wraps MEOS C function ttext_max_value.
func TtextMaxValue(temp Temporal) string {
	res := C.ttext_max_value(temp.Inner())
	return text2cstring(res)
}


// TtextMinValue wraps MEOS C function ttext_min_value.
func TtextMinValue(temp Temporal) string {
	res := C.ttext_min_value(temp.Inner())
	return text2cstring(res)
}


// TtextStartValue wraps MEOS C function ttext_start_value.
func TtextStartValue(temp Temporal) string {
	res := C.ttext_start_value(temp.Inner())
	return text2cstring(res)
}


// TtextValueAtTimestamptz wraps MEOS C function ttext_value_at_timestamptz.
func TtextValueAtTimestamptz(temp Temporal, t int64, strict bool) (bool, string) {
	var _out_value *C.text
	res := C.ttext_value_at_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(strict), &_out_value)
	return bool(res), text2cstring(_out_value)
}


// TtextValueN wraps MEOS C function ttext_value_n.
func TtextValueN(temp Temporal, n int) (bool, string) {
	var _out_result *C.text
	res := C.ttext_value_n(temp.Inner(), C.int(n), &_out_result)
	return bool(res), text2cstring(_out_result)
}


// TODO ttext_values: unsupported return type text **
// func TtextValues(...) { /* not yet handled by codegen */ }


// FloatDegrees wraps MEOS C function float_degrees.
func FloatDegrees(value float64, normalize bool) float64 {
	res := C.float_degrees(C.double(value), C.bool(normalize))
	return float64(res)
}


// TemparrRound wraps MEOS C function temparr_round.
func TemparrRound(temp []Temporal, maxdd int) []Temporal {
	_c_temp := make([]*C.Temporal, len(temp))
	for _i, _v := range temp { _c_temp[_i] = _v._inner }
	res := C.temparr_round((**C.Temporal)(unsafe.Pointer(&_c_temp[0])), C.int(len(temp)), C.int(maxdd))
	_n := len(temp)
	_slice := unsafe.Slice((**C.Temporal)(unsafe.Pointer(res)), _n)
	_out := make([]Temporal, _n)
	for _i, _e := range _slice {
		_out[_i] = CreateTemporal(_e)
	}
	return _out
}


// TemporalRound wraps MEOS C function temporal_round.
func TemporalRound(temp Temporal, maxdd int) Temporal {
	res := C.temporal_round(temp.Inner(), C.int(maxdd))
	return CreateTemporal(res)
}


// TemporalScaleTime wraps MEOS C function temporal_scale_time.
func TemporalScaleTime(temp Temporal, duration timeutil.Timedelta) Temporal {
	res := C.temporal_scale_time(temp.Inner(), duration.Inner())
	return CreateTemporal(res)
}


// TemporalSetInterp wraps MEOS C function temporal_set_interp.
func TemporalSetInterp(temp Temporal, interp Interpolation) Temporal {
	res := C.temporal_set_interp(temp.Inner(), C.interpType(interp))
	return CreateTemporal(res)
}


// TemporalShiftScaleTime wraps MEOS C function temporal_shift_scale_time.
func TemporalShiftScaleTime(temp Temporal, shift timeutil.Timedelta, duration timeutil.Timedelta) Temporal {
	res := C.temporal_shift_scale_time(temp.Inner(), shift.Inner(), duration.Inner())
	return CreateTemporal(res)
}


// TemporalShiftTime wraps MEOS C function temporal_shift_time.
func TemporalShiftTime(temp Temporal, shift timeutil.Timedelta) Temporal {
	res := C.temporal_shift_time(temp.Inner(), shift.Inner())
	return CreateTemporal(res)
}


// TemporalToTinstant wraps MEOS C function temporal_to_tinstant.
func TemporalToTinstant(temp Temporal) TInstant {
	res := C.temporal_to_tinstant(temp.Inner())
	return TInstant{_inner: res}
}


// TemporalToTsequence wraps MEOS C function temporal_to_tsequence.
func TemporalToTsequence(temp Temporal, interp Interpolation) TSequence {
	res := C.temporal_to_tsequence(temp.Inner(), C.interpType(interp))
	return TSequence{_inner: res}
}


// TemporalToTsequenceset wraps MEOS C function temporal_to_tsequenceset.
func TemporalToTsequenceset(temp Temporal, interp Interpolation) TSequenceSet {
	res := C.temporal_to_tsequenceset(temp.Inner(), C.interpType(interp))
	return TSequenceSet{_inner: res}
}


// TfloatCeil wraps MEOS C function tfloat_ceil.
func TfloatCeil(temp Temporal) Temporal {
	res := C.tfloat_ceil(temp.Inner())
	return CreateTemporal(res)
}


// TfloatDegrees wraps MEOS C function tfloat_degrees.
func TfloatDegrees(temp Temporal, normalize bool) Temporal {
	res := C.tfloat_degrees(temp.Inner(), C.bool(normalize))
	return CreateTemporal(res)
}


// TfloatFloor wraps MEOS C function tfloat_floor.
func TfloatFloor(temp Temporal) Temporal {
	res := C.tfloat_floor(temp.Inner())
	return CreateTemporal(res)
}


// TfloatRadians wraps MEOS C function tfloat_radians.
func TfloatRadians(temp Temporal) Temporal {
	res := C.tfloat_radians(temp.Inner())
	return CreateTemporal(res)
}


// TfloatScaleValue wraps MEOS C function tfloat_scale_value.
func TfloatScaleValue(temp Temporal, width float64) Temporal {
	res := C.tfloat_scale_value(temp.Inner(), C.double(width))
	return CreateTemporal(res)
}


// TfloatShiftScaleValue wraps MEOS C function tfloat_shift_scale_value.
func TfloatShiftScaleValue(temp Temporal, shift float64, width float64) Temporal {
	res := C.tfloat_shift_scale_value(temp.Inner(), C.double(shift), C.double(width))
	return CreateTemporal(res)
}


// TfloatShiftValue wraps MEOS C function tfloat_shift_value.
func TfloatShiftValue(temp Temporal, shift float64) Temporal {
	res := C.tfloat_shift_value(temp.Inner(), C.double(shift))
	return CreateTemporal(res)
}


// TintScaleValue wraps MEOS C function tint_scale_value.
func TintScaleValue(temp Temporal, width int) Temporal {
	res := C.tint_scale_value(temp.Inner(), C.int(width))
	return CreateTemporal(res)
}


// TintShiftScaleValue wraps MEOS C function tint_shift_scale_value.
func TintShiftScaleValue(temp Temporal, shift int, width int) Temporal {
	res := C.tint_shift_scale_value(temp.Inner(), C.int(shift), C.int(width))
	return CreateTemporal(res)
}


// TintShiftValue wraps MEOS C function tint_shift_value.
func TintShiftValue(temp Temporal, shift int) Temporal {
	res := C.tint_shift_value(temp.Inner(), C.int(shift))
	return CreateTemporal(res)
}


// TemporalAppendTinstant wraps MEOS C function temporal_append_tinstant.
func TemporalAppendTinstant(temp Temporal, inst TInstant, interp Interpolation, maxdist float64, maxt timeutil.Timedelta, expand bool) Temporal {
	res := C.temporal_append_tinstant(temp.Inner(), inst.Inner(), C.interpType(interp), C.double(maxdist), maxt.Inner(), C.bool(expand))
	return CreateTemporal(res)
}


// TemporalAppendTsequence wraps MEOS C function temporal_append_tsequence.
func TemporalAppendTsequence(temp Temporal, seq TSequence, expand bool) Temporal {
	res := C.temporal_append_tsequence(temp.Inner(), seq.Inner(), C.bool(expand))
	return CreateTemporal(res)
}


// TemporalDeleteTimestamptz wraps MEOS C function temporal_delete_timestamptz.
func TemporalDeleteTimestamptz(temp Temporal, t int64, connect bool) Temporal {
	res := C.temporal_delete_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(connect))
	return CreateTemporal(res)
}


// TemporalDeleteTstzset wraps MEOS C function temporal_delete_tstzset.
func TemporalDeleteTstzset(temp Temporal, s *Set, connect bool) Temporal {
	res := C.temporal_delete_tstzset(temp.Inner(), s._inner, C.bool(connect))
	return CreateTemporal(res)
}


// TemporalDeleteTstzspan wraps MEOS C function temporal_delete_tstzspan.
func TemporalDeleteTstzspan(temp Temporal, s *Span, connect bool) Temporal {
	res := C.temporal_delete_tstzspan(temp.Inner(), s._inner, C.bool(connect))
	return CreateTemporal(res)
}


// TemporalDeleteTstzspanset wraps MEOS C function temporal_delete_tstzspanset.
func TemporalDeleteTstzspanset(temp Temporal, ss *SpanSet, connect bool) Temporal {
	res := C.temporal_delete_tstzspanset(temp.Inner(), ss._inner, C.bool(connect))
	return CreateTemporal(res)
}


// TemporalInsert wraps MEOS C function temporal_insert.
func TemporalInsert(temp1 Temporal, temp2 Temporal, connect bool) Temporal {
	res := C.temporal_insert(temp1.Inner(), temp2.Inner(), C.bool(connect))
	return CreateTemporal(res)
}


// TemporalMerge wraps MEOS C function temporal_merge.
func TemporalMerge(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.temporal_merge(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TemporalMergeArray wraps MEOS C function temporal_merge_array.
func TemporalMergeArray(temparr []Temporal) Temporal {
	_c_temparr := make([]*C.Temporal, len(temparr))
	for _i, _v := range temparr { _c_temparr[_i] = _v._inner }
	res := C.temporal_merge_array((**C.Temporal)(unsafe.Pointer(&_c_temparr[0])), C.int(len(temparr)))
	return CreateTemporal(res)
}


// TemporalUpdate wraps MEOS C function temporal_update.
func TemporalUpdate(temp1 Temporal, temp2 Temporal, connect bool) Temporal {
	res := C.temporal_update(temp1.Inner(), temp2.Inner(), C.bool(connect))
	return CreateTemporal(res)
}


// TboolAtValue wraps MEOS C function tbool_at_value.
func TboolAtValue(temp Temporal, b bool) Temporal {
	res := C.tbool_at_value(temp.Inner(), C.bool(b))
	return CreateTemporal(res)
}


// TboolMinusValue wraps MEOS C function tbool_minus_value.
func TboolMinusValue(temp Temporal, b bool) Temporal {
	res := C.tbool_minus_value(temp.Inner(), C.bool(b))
	return CreateTemporal(res)
}


// TemporalAfterTimestamptz wraps MEOS C function temporal_after_timestamptz.
func TemporalAfterTimestamptz(temp Temporal, t int64, strict bool) Temporal {
	res := C.temporal_after_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(strict))
	return CreateTemporal(res)
}


// TemporalAtMax wraps MEOS C function temporal_at_max.
func TemporalAtMax(temp Temporal) Temporal {
	res := C.temporal_at_max(temp.Inner())
	return CreateTemporal(res)
}


// TemporalAtMin wraps MEOS C function temporal_at_min.
func TemporalAtMin(temp Temporal) Temporal {
	res := C.temporal_at_min(temp.Inner())
	return CreateTemporal(res)
}


// TemporalAtTimestamptz wraps MEOS C function temporal_at_timestamptz.
func TemporalAtTimestamptz(temp Temporal, t int64) Temporal {
	res := C.temporal_at_timestamptz(temp.Inner(), C.TimestampTz(t))
	return CreateTemporal(res)
}


// TemporalAtTstzset wraps MEOS C function temporal_at_tstzset.
func TemporalAtTstzset(temp Temporal, s *Set) Temporal {
	res := C.temporal_at_tstzset(temp.Inner(), s._inner)
	return CreateTemporal(res)
}


// TemporalAtTstzspan wraps MEOS C function temporal_at_tstzspan.
func TemporalAtTstzspan(temp Temporal, s *Span) Temporal {
	res := C.temporal_at_tstzspan(temp.Inner(), s._inner)
	return CreateTemporal(res)
}


// TemporalAtTstzspanset wraps MEOS C function temporal_at_tstzspanset.
func TemporalAtTstzspanset(temp Temporal, ss *SpanSet) Temporal {
	res := C.temporal_at_tstzspanset(temp.Inner(), ss._inner)
	return CreateTemporal(res)
}


// TemporalAtValues wraps MEOS C function temporal_at_values.
func TemporalAtValues(temp Temporal, set *Set) Temporal {
	res := C.temporal_at_values(temp.Inner(), set._inner)
	return CreateTemporal(res)
}


// TemporalBeforeTimestamptz wraps MEOS C function temporal_before_timestamptz.
func TemporalBeforeTimestamptz(temp Temporal, t int64, strict bool) Temporal {
	res := C.temporal_before_timestamptz(temp.Inner(), C.TimestampTz(t), C.bool(strict))
	return CreateTemporal(res)
}


// TemporalMinusMax wraps MEOS C function temporal_minus_max.
func TemporalMinusMax(temp Temporal) Temporal {
	res := C.temporal_minus_max(temp.Inner())
	return CreateTemporal(res)
}


// TemporalMinusMin wraps MEOS C function temporal_minus_min.
func TemporalMinusMin(temp Temporal) Temporal {
	res := C.temporal_minus_min(temp.Inner())
	return CreateTemporal(res)
}


// TemporalMinusTimestamptz wraps MEOS C function temporal_minus_timestamptz.
func TemporalMinusTimestamptz(temp Temporal, t int64) Temporal {
	res := C.temporal_minus_timestamptz(temp.Inner(), C.TimestampTz(t))
	return CreateTemporal(res)
}


// TemporalMinusTstzset wraps MEOS C function temporal_minus_tstzset.
func TemporalMinusTstzset(temp Temporal, s *Set) Temporal {
	res := C.temporal_minus_tstzset(temp.Inner(), s._inner)
	return CreateTemporal(res)
}


// TemporalMinusTstzspan wraps MEOS C function temporal_minus_tstzspan.
func TemporalMinusTstzspan(temp Temporal, s *Span) Temporal {
	res := C.temporal_minus_tstzspan(temp.Inner(), s._inner)
	return CreateTemporal(res)
}


// TemporalMinusTstzspanset wraps MEOS C function temporal_minus_tstzspanset.
func TemporalMinusTstzspanset(temp Temporal, ss *SpanSet) Temporal {
	res := C.temporal_minus_tstzspanset(temp.Inner(), ss._inner)
	return CreateTemporal(res)
}


// TemporalMinusValues wraps MEOS C function temporal_minus_values.
func TemporalMinusValues(temp Temporal, set *Set) Temporal {
	res := C.temporal_minus_values(temp.Inner(), set._inner)
	return CreateTemporal(res)
}


// TfloatAtValue wraps MEOS C function tfloat_at_value.
func TfloatAtValue(temp Temporal, d float64) Temporal {
	res := C.tfloat_at_value(temp.Inner(), C.double(d))
	return CreateTemporal(res)
}


// TfloatMinusValue wraps MEOS C function tfloat_minus_value.
func TfloatMinusValue(temp Temporal, d float64) Temporal {
	res := C.tfloat_minus_value(temp.Inner(), C.double(d))
	return CreateTemporal(res)
}


// TintAtValue wraps MEOS C function tint_at_value.
func TintAtValue(temp Temporal, i int) Temporal {
	res := C.tint_at_value(temp.Inner(), C.int(i))
	return CreateTemporal(res)
}


// TintMinusValue wraps MEOS C function tint_minus_value.
func TintMinusValue(temp Temporal, i int) Temporal {
	res := C.tint_minus_value(temp.Inner(), C.int(i))
	return CreateTemporal(res)
}


// TnumberAtSpan wraps MEOS C function tnumber_at_span.
func TnumberAtSpan(temp Temporal, span *Span) Temporal {
	res := C.tnumber_at_span(temp.Inner(), span._inner)
	return CreateTemporal(res)
}


// TnumberAtSpanset wraps MEOS C function tnumber_at_spanset.
func TnumberAtSpanset(temp Temporal, ss *SpanSet) Temporal {
	res := C.tnumber_at_spanset(temp.Inner(), ss._inner)
	return CreateTemporal(res)
}


// TnumberAtTBOX wraps MEOS C function tnumber_at_tbox.
func TnumberAtTBOX(temp Temporal, box *TBox) Temporal {
	res := C.tnumber_at_tbox(temp.Inner(), box._inner)
	return CreateTemporal(res)
}


// TnumberMinusSpan wraps MEOS C function tnumber_minus_span.
func TnumberMinusSpan(temp Temporal, span *Span) Temporal {
	res := C.tnumber_minus_span(temp.Inner(), span._inner)
	return CreateTemporal(res)
}


// TnumberMinusSpanset wraps MEOS C function tnumber_minus_spanset.
func TnumberMinusSpanset(temp Temporal, ss *SpanSet) Temporal {
	res := C.tnumber_minus_spanset(temp.Inner(), ss._inner)
	return CreateTemporal(res)
}


// TnumberMinusTBOX wraps MEOS C function tnumber_minus_tbox.
func TnumberMinusTBOX(temp Temporal, box *TBox) Temporal {
	res := C.tnumber_minus_tbox(temp.Inner(), box._inner)
	return CreateTemporal(res)
}


// TtextAtValue wraps MEOS C function ttext_at_value.
func TtextAtValue(temp Temporal, txt string) Temporal {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ttext_at_value(temp.Inner(), _c_txt)
	return CreateTemporal(res)
}


// TtextMinusValue wraps MEOS C function ttext_minus_value.
func TtextMinusValue(temp Temporal, txt string) Temporal {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ttext_minus_value(temp.Inner(), _c_txt)
	return CreateTemporal(res)
}


// TemporalCmp wraps MEOS C function temporal_cmp.
func TemporalCmp(temp1 Temporal, temp2 Temporal) int {
	res := C.temporal_cmp(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TemporalEq wraps MEOS C function temporal_eq.
func TemporalEq(temp1 Temporal, temp2 Temporal) bool {
	res := C.temporal_eq(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// TemporalGe wraps MEOS C function temporal_ge.
func TemporalGe(temp1 Temporal, temp2 Temporal) bool {
	res := C.temporal_ge(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// TemporalGt wraps MEOS C function temporal_gt.
func TemporalGt(temp1 Temporal, temp2 Temporal) bool {
	res := C.temporal_gt(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// TemporalLe wraps MEOS C function temporal_le.
func TemporalLe(temp1 Temporal, temp2 Temporal) bool {
	res := C.temporal_le(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// TemporalLt wraps MEOS C function temporal_lt.
func TemporalLt(temp1 Temporal, temp2 Temporal) bool {
	res := C.temporal_lt(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// TemporalNe wraps MEOS C function temporal_ne.
func TemporalNe(temp1 Temporal, temp2 Temporal) bool {
	res := C.temporal_ne(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// AlwaysEqBoolTbool wraps MEOS C function always_eq_bool_tbool.
func AlwaysEqBoolTbool(b bool, temp Temporal) int {
	res := C.always_eq_bool_tbool(C.bool(b), temp.Inner())
	return int(res)
}


// AlwaysEqFloatTfloat wraps MEOS C function always_eq_float_tfloat.
func AlwaysEqFloatTfloat(d float64, temp Temporal) int {
	res := C.always_eq_float_tfloat(C.double(d), temp.Inner())
	return int(res)
}


// AlwaysEqIntTint wraps MEOS C function always_eq_int_tint.
func AlwaysEqIntTint(i int, temp Temporal) int {
	res := C.always_eq_int_tint(C.int(i), temp.Inner())
	return int(res)
}


// AlwaysEqTboolBool wraps MEOS C function always_eq_tbool_bool.
func AlwaysEqTboolBool(temp Temporal, b bool) int {
	res := C.always_eq_tbool_bool(temp.Inner(), C.bool(b))
	return int(res)
}


// AlwaysEqTemporalTemporal wraps MEOS C function always_eq_temporal_temporal.
func AlwaysEqTemporalTemporal(temp1 Temporal, temp2 Temporal) int {
	res := C.always_eq_temporal_temporal(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AlwaysEqTextTtext wraps MEOS C function always_eq_text_ttext.
func AlwaysEqTextTtext(txt string, temp Temporal) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.always_eq_text_ttext(_c_txt, temp.Inner())
	return int(res)
}


// AlwaysEqTfloatFloat wraps MEOS C function always_eq_tfloat_float.
func AlwaysEqTfloatFloat(temp Temporal, d float64) int {
	res := C.always_eq_tfloat_float(temp.Inner(), C.double(d))
	return int(res)
}


// AlwaysEqTintInt wraps MEOS C function always_eq_tint_int.
func AlwaysEqTintInt(temp Temporal, i int) int {
	res := C.always_eq_tint_int(temp.Inner(), C.int(i))
	return int(res)
}


// AlwaysEqTtextText wraps MEOS C function always_eq_ttext_text.
func AlwaysEqTtextText(temp Temporal, txt string) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.always_eq_ttext_text(temp.Inner(), _c_txt)
	return int(res)
}


// AlwaysGeFloatTfloat wraps MEOS C function always_ge_float_tfloat.
func AlwaysGeFloatTfloat(d float64, temp Temporal) int {
	res := C.always_ge_float_tfloat(C.double(d), temp.Inner())
	return int(res)
}


// AlwaysGeIntTint wraps MEOS C function always_ge_int_tint.
func AlwaysGeIntTint(i int, temp Temporal) int {
	res := C.always_ge_int_tint(C.int(i), temp.Inner())
	return int(res)
}


// AlwaysGeTemporalTemporal wraps MEOS C function always_ge_temporal_temporal.
func AlwaysGeTemporalTemporal(temp1 Temporal, temp2 Temporal) int {
	res := C.always_ge_temporal_temporal(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AlwaysGeTextTtext wraps MEOS C function always_ge_text_ttext.
func AlwaysGeTextTtext(txt string, temp Temporal) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.always_ge_text_ttext(_c_txt, temp.Inner())
	return int(res)
}


// AlwaysGeTfloatFloat wraps MEOS C function always_ge_tfloat_float.
func AlwaysGeTfloatFloat(temp Temporal, d float64) int {
	res := C.always_ge_tfloat_float(temp.Inner(), C.double(d))
	return int(res)
}


// AlwaysGeTintInt wraps MEOS C function always_ge_tint_int.
func AlwaysGeTintInt(temp Temporal, i int) int {
	res := C.always_ge_tint_int(temp.Inner(), C.int(i))
	return int(res)
}


// AlwaysGeTtextText wraps MEOS C function always_ge_ttext_text.
func AlwaysGeTtextText(temp Temporal, txt string) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.always_ge_ttext_text(temp.Inner(), _c_txt)
	return int(res)
}


// AlwaysGtFloatTfloat wraps MEOS C function always_gt_float_tfloat.
func AlwaysGtFloatTfloat(d float64, temp Temporal) int {
	res := C.always_gt_float_tfloat(C.double(d), temp.Inner())
	return int(res)
}


// AlwaysGtIntTint wraps MEOS C function always_gt_int_tint.
func AlwaysGtIntTint(i int, temp Temporal) int {
	res := C.always_gt_int_tint(C.int(i), temp.Inner())
	return int(res)
}


// AlwaysGtTemporalTemporal wraps MEOS C function always_gt_temporal_temporal.
func AlwaysGtTemporalTemporal(temp1 Temporal, temp2 Temporal) int {
	res := C.always_gt_temporal_temporal(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AlwaysGtTextTtext wraps MEOS C function always_gt_text_ttext.
func AlwaysGtTextTtext(txt string, temp Temporal) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.always_gt_text_ttext(_c_txt, temp.Inner())
	return int(res)
}


// AlwaysGtTfloatFloat wraps MEOS C function always_gt_tfloat_float.
func AlwaysGtTfloatFloat(temp Temporal, d float64) int {
	res := C.always_gt_tfloat_float(temp.Inner(), C.double(d))
	return int(res)
}


// AlwaysGtTintInt wraps MEOS C function always_gt_tint_int.
func AlwaysGtTintInt(temp Temporal, i int) int {
	res := C.always_gt_tint_int(temp.Inner(), C.int(i))
	return int(res)
}


// AlwaysGtTtextText wraps MEOS C function always_gt_ttext_text.
func AlwaysGtTtextText(temp Temporal, txt string) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.always_gt_ttext_text(temp.Inner(), _c_txt)
	return int(res)
}


// AlwaysLeFloatTfloat wraps MEOS C function always_le_float_tfloat.
func AlwaysLeFloatTfloat(d float64, temp Temporal) int {
	res := C.always_le_float_tfloat(C.double(d), temp.Inner())
	return int(res)
}


// AlwaysLeIntTint wraps MEOS C function always_le_int_tint.
func AlwaysLeIntTint(i int, temp Temporal) int {
	res := C.always_le_int_tint(C.int(i), temp.Inner())
	return int(res)
}


// AlwaysLeTemporalTemporal wraps MEOS C function always_le_temporal_temporal.
func AlwaysLeTemporalTemporal(temp1 Temporal, temp2 Temporal) int {
	res := C.always_le_temporal_temporal(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AlwaysLeTextTtext wraps MEOS C function always_le_text_ttext.
func AlwaysLeTextTtext(txt string, temp Temporal) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.always_le_text_ttext(_c_txt, temp.Inner())
	return int(res)
}


// AlwaysLeTfloatFloat wraps MEOS C function always_le_tfloat_float.
func AlwaysLeTfloatFloat(temp Temporal, d float64) int {
	res := C.always_le_tfloat_float(temp.Inner(), C.double(d))
	return int(res)
}


// AlwaysLeTintInt wraps MEOS C function always_le_tint_int.
func AlwaysLeTintInt(temp Temporal, i int) int {
	res := C.always_le_tint_int(temp.Inner(), C.int(i))
	return int(res)
}


// AlwaysLeTtextText wraps MEOS C function always_le_ttext_text.
func AlwaysLeTtextText(temp Temporal, txt string) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.always_le_ttext_text(temp.Inner(), _c_txt)
	return int(res)
}


// AlwaysLtFloatTfloat wraps MEOS C function always_lt_float_tfloat.
func AlwaysLtFloatTfloat(d float64, temp Temporal) int {
	res := C.always_lt_float_tfloat(C.double(d), temp.Inner())
	return int(res)
}


// AlwaysLtIntTint wraps MEOS C function always_lt_int_tint.
func AlwaysLtIntTint(i int, temp Temporal) int {
	res := C.always_lt_int_tint(C.int(i), temp.Inner())
	return int(res)
}


// AlwaysLtTemporalTemporal wraps MEOS C function always_lt_temporal_temporal.
func AlwaysLtTemporalTemporal(temp1 Temporal, temp2 Temporal) int {
	res := C.always_lt_temporal_temporal(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AlwaysLtTextTtext wraps MEOS C function always_lt_text_ttext.
func AlwaysLtTextTtext(txt string, temp Temporal) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.always_lt_text_ttext(_c_txt, temp.Inner())
	return int(res)
}


// AlwaysLtTfloatFloat wraps MEOS C function always_lt_tfloat_float.
func AlwaysLtTfloatFloat(temp Temporal, d float64) int {
	res := C.always_lt_tfloat_float(temp.Inner(), C.double(d))
	return int(res)
}


// AlwaysLtTintInt wraps MEOS C function always_lt_tint_int.
func AlwaysLtTintInt(temp Temporal, i int) int {
	res := C.always_lt_tint_int(temp.Inner(), C.int(i))
	return int(res)
}


// AlwaysLtTtextText wraps MEOS C function always_lt_ttext_text.
func AlwaysLtTtextText(temp Temporal, txt string) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.always_lt_ttext_text(temp.Inner(), _c_txt)
	return int(res)
}


// AlwaysNeBoolTbool wraps MEOS C function always_ne_bool_tbool.
func AlwaysNeBoolTbool(b bool, temp Temporal) int {
	res := C.always_ne_bool_tbool(C.bool(b), temp.Inner())
	return int(res)
}


// AlwaysNeFloatTfloat wraps MEOS C function always_ne_float_tfloat.
func AlwaysNeFloatTfloat(d float64, temp Temporal) int {
	res := C.always_ne_float_tfloat(C.double(d), temp.Inner())
	return int(res)
}


// AlwaysNeIntTint wraps MEOS C function always_ne_int_tint.
func AlwaysNeIntTint(i int, temp Temporal) int {
	res := C.always_ne_int_tint(C.int(i), temp.Inner())
	return int(res)
}


// AlwaysNeTboolBool wraps MEOS C function always_ne_tbool_bool.
func AlwaysNeTboolBool(temp Temporal, b bool) int {
	res := C.always_ne_tbool_bool(temp.Inner(), C.bool(b))
	return int(res)
}


// AlwaysNeTemporalTemporal wraps MEOS C function always_ne_temporal_temporal.
func AlwaysNeTemporalTemporal(temp1 Temporal, temp2 Temporal) int {
	res := C.always_ne_temporal_temporal(temp1.Inner(), temp2.Inner())
	return int(res)
}


// AlwaysNeTextTtext wraps MEOS C function always_ne_text_ttext.
func AlwaysNeTextTtext(txt string, temp Temporal) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.always_ne_text_ttext(_c_txt, temp.Inner())
	return int(res)
}


// AlwaysNeTfloatFloat wraps MEOS C function always_ne_tfloat_float.
func AlwaysNeTfloatFloat(temp Temporal, d float64) int {
	res := C.always_ne_tfloat_float(temp.Inner(), C.double(d))
	return int(res)
}


// AlwaysNeTintInt wraps MEOS C function always_ne_tint_int.
func AlwaysNeTintInt(temp Temporal, i int) int {
	res := C.always_ne_tint_int(temp.Inner(), C.int(i))
	return int(res)
}


// AlwaysNeTtextText wraps MEOS C function always_ne_ttext_text.
func AlwaysNeTtextText(temp Temporal, txt string) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.always_ne_ttext_text(temp.Inner(), _c_txt)
	return int(res)
}


// EverEqBoolTbool wraps MEOS C function ever_eq_bool_tbool.
func EverEqBoolTbool(b bool, temp Temporal) int {
	res := C.ever_eq_bool_tbool(C.bool(b), temp.Inner())
	return int(res)
}


// EverEqFloatTfloat wraps MEOS C function ever_eq_float_tfloat.
func EverEqFloatTfloat(d float64, temp Temporal) int {
	res := C.ever_eq_float_tfloat(C.double(d), temp.Inner())
	return int(res)
}


// EverEqIntTint wraps MEOS C function ever_eq_int_tint.
func EverEqIntTint(i int, temp Temporal) int {
	res := C.ever_eq_int_tint(C.int(i), temp.Inner())
	return int(res)
}


// EverEqTboolBool wraps MEOS C function ever_eq_tbool_bool.
func EverEqTboolBool(temp Temporal, b bool) int {
	res := C.ever_eq_tbool_bool(temp.Inner(), C.bool(b))
	return int(res)
}


// EverEqTemporalTemporal wraps MEOS C function ever_eq_temporal_temporal.
func EverEqTemporalTemporal(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_eq_temporal_temporal(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverEqTextTtext wraps MEOS C function ever_eq_text_ttext.
func EverEqTextTtext(txt string, temp Temporal) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ever_eq_text_ttext(_c_txt, temp.Inner())
	return int(res)
}


// EverEqTfloatFloat wraps MEOS C function ever_eq_tfloat_float.
func EverEqTfloatFloat(temp Temporal, d float64) int {
	res := C.ever_eq_tfloat_float(temp.Inner(), C.double(d))
	return int(res)
}


// EverEqTintInt wraps MEOS C function ever_eq_tint_int.
func EverEqTintInt(temp Temporal, i int) int {
	res := C.ever_eq_tint_int(temp.Inner(), C.int(i))
	return int(res)
}


// EverEqTtextText wraps MEOS C function ever_eq_ttext_text.
func EverEqTtextText(temp Temporal, txt string) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ever_eq_ttext_text(temp.Inner(), _c_txt)
	return int(res)
}


// EverGeFloatTfloat wraps MEOS C function ever_ge_float_tfloat.
func EverGeFloatTfloat(d float64, temp Temporal) int {
	res := C.ever_ge_float_tfloat(C.double(d), temp.Inner())
	return int(res)
}


// EverGeIntTint wraps MEOS C function ever_ge_int_tint.
func EverGeIntTint(i int, temp Temporal) int {
	res := C.ever_ge_int_tint(C.int(i), temp.Inner())
	return int(res)
}


// EverGeTemporalTemporal wraps MEOS C function ever_ge_temporal_temporal.
func EverGeTemporalTemporal(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_ge_temporal_temporal(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverGeTextTtext wraps MEOS C function ever_ge_text_ttext.
func EverGeTextTtext(txt string, temp Temporal) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ever_ge_text_ttext(_c_txt, temp.Inner())
	return int(res)
}


// EverGeTfloatFloat wraps MEOS C function ever_ge_tfloat_float.
func EverGeTfloatFloat(temp Temporal, d float64) int {
	res := C.ever_ge_tfloat_float(temp.Inner(), C.double(d))
	return int(res)
}


// EverGeTintInt wraps MEOS C function ever_ge_tint_int.
func EverGeTintInt(temp Temporal, i int) int {
	res := C.ever_ge_tint_int(temp.Inner(), C.int(i))
	return int(res)
}


// EverGeTtextText wraps MEOS C function ever_ge_ttext_text.
func EverGeTtextText(temp Temporal, txt string) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ever_ge_ttext_text(temp.Inner(), _c_txt)
	return int(res)
}


// EverGtFloatTfloat wraps MEOS C function ever_gt_float_tfloat.
func EverGtFloatTfloat(d float64, temp Temporal) int {
	res := C.ever_gt_float_tfloat(C.double(d), temp.Inner())
	return int(res)
}


// EverGtIntTint wraps MEOS C function ever_gt_int_tint.
func EverGtIntTint(i int, temp Temporal) int {
	res := C.ever_gt_int_tint(C.int(i), temp.Inner())
	return int(res)
}


// EverGtTemporalTemporal wraps MEOS C function ever_gt_temporal_temporal.
func EverGtTemporalTemporal(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_gt_temporal_temporal(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverGtTextTtext wraps MEOS C function ever_gt_text_ttext.
func EverGtTextTtext(txt string, temp Temporal) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ever_gt_text_ttext(_c_txt, temp.Inner())
	return int(res)
}


// EverGtTfloatFloat wraps MEOS C function ever_gt_tfloat_float.
func EverGtTfloatFloat(temp Temporal, d float64) int {
	res := C.ever_gt_tfloat_float(temp.Inner(), C.double(d))
	return int(res)
}


// EverGtTintInt wraps MEOS C function ever_gt_tint_int.
func EverGtTintInt(temp Temporal, i int) int {
	res := C.ever_gt_tint_int(temp.Inner(), C.int(i))
	return int(res)
}


// EverGtTtextText wraps MEOS C function ever_gt_ttext_text.
func EverGtTtextText(temp Temporal, txt string) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ever_gt_ttext_text(temp.Inner(), _c_txt)
	return int(res)
}


// EverLeFloatTfloat wraps MEOS C function ever_le_float_tfloat.
func EverLeFloatTfloat(d float64, temp Temporal) int {
	res := C.ever_le_float_tfloat(C.double(d), temp.Inner())
	return int(res)
}


// EverLeIntTint wraps MEOS C function ever_le_int_tint.
func EverLeIntTint(i int, temp Temporal) int {
	res := C.ever_le_int_tint(C.int(i), temp.Inner())
	return int(res)
}


// EverLeTemporalTemporal wraps MEOS C function ever_le_temporal_temporal.
func EverLeTemporalTemporal(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_le_temporal_temporal(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverLeTextTtext wraps MEOS C function ever_le_text_ttext.
func EverLeTextTtext(txt string, temp Temporal) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ever_le_text_ttext(_c_txt, temp.Inner())
	return int(res)
}


// EverLeTfloatFloat wraps MEOS C function ever_le_tfloat_float.
func EverLeTfloatFloat(temp Temporal, d float64) int {
	res := C.ever_le_tfloat_float(temp.Inner(), C.double(d))
	return int(res)
}


// EverLeTintInt wraps MEOS C function ever_le_tint_int.
func EverLeTintInt(temp Temporal, i int) int {
	res := C.ever_le_tint_int(temp.Inner(), C.int(i))
	return int(res)
}


// EverLeTtextText wraps MEOS C function ever_le_ttext_text.
func EverLeTtextText(temp Temporal, txt string) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ever_le_ttext_text(temp.Inner(), _c_txt)
	return int(res)
}


// EverLtFloatTfloat wraps MEOS C function ever_lt_float_tfloat.
func EverLtFloatTfloat(d float64, temp Temporal) int {
	res := C.ever_lt_float_tfloat(C.double(d), temp.Inner())
	return int(res)
}


// EverLtIntTint wraps MEOS C function ever_lt_int_tint.
func EverLtIntTint(i int, temp Temporal) int {
	res := C.ever_lt_int_tint(C.int(i), temp.Inner())
	return int(res)
}


// EverLtTemporalTemporal wraps MEOS C function ever_lt_temporal_temporal.
func EverLtTemporalTemporal(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_lt_temporal_temporal(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverLtTextTtext wraps MEOS C function ever_lt_text_ttext.
func EverLtTextTtext(txt string, temp Temporal) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ever_lt_text_ttext(_c_txt, temp.Inner())
	return int(res)
}


// EverLtTfloatFloat wraps MEOS C function ever_lt_tfloat_float.
func EverLtTfloatFloat(temp Temporal, d float64) int {
	res := C.ever_lt_tfloat_float(temp.Inner(), C.double(d))
	return int(res)
}


// EverLtTintInt wraps MEOS C function ever_lt_tint_int.
func EverLtTintInt(temp Temporal, i int) int {
	res := C.ever_lt_tint_int(temp.Inner(), C.int(i))
	return int(res)
}


// EverLtTtextText wraps MEOS C function ever_lt_ttext_text.
func EverLtTtextText(temp Temporal, txt string) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ever_lt_ttext_text(temp.Inner(), _c_txt)
	return int(res)
}


// EverNeBoolTbool wraps MEOS C function ever_ne_bool_tbool.
func EverNeBoolTbool(b bool, temp Temporal) int {
	res := C.ever_ne_bool_tbool(C.bool(b), temp.Inner())
	return int(res)
}


// EverNeFloatTfloat wraps MEOS C function ever_ne_float_tfloat.
func EverNeFloatTfloat(d float64, temp Temporal) int {
	res := C.ever_ne_float_tfloat(C.double(d), temp.Inner())
	return int(res)
}


// EverNeIntTint wraps MEOS C function ever_ne_int_tint.
func EverNeIntTint(i int, temp Temporal) int {
	res := C.ever_ne_int_tint(C.int(i), temp.Inner())
	return int(res)
}


// EverNeTboolBool wraps MEOS C function ever_ne_tbool_bool.
func EverNeTboolBool(temp Temporal, b bool) int {
	res := C.ever_ne_tbool_bool(temp.Inner(), C.bool(b))
	return int(res)
}


// EverNeTemporalTemporal wraps MEOS C function ever_ne_temporal_temporal.
func EverNeTemporalTemporal(temp1 Temporal, temp2 Temporal) int {
	res := C.ever_ne_temporal_temporal(temp1.Inner(), temp2.Inner())
	return int(res)
}


// EverNeTextTtext wraps MEOS C function ever_ne_text_ttext.
func EverNeTextTtext(txt string, temp Temporal) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ever_ne_text_ttext(_c_txt, temp.Inner())
	return int(res)
}


// EverNeTfloatFloat wraps MEOS C function ever_ne_tfloat_float.
func EverNeTfloatFloat(temp Temporal, d float64) int {
	res := C.ever_ne_tfloat_float(temp.Inner(), C.double(d))
	return int(res)
}


// EverNeTintInt wraps MEOS C function ever_ne_tint_int.
func EverNeTintInt(temp Temporal, i int) int {
	res := C.ever_ne_tint_int(temp.Inner(), C.int(i))
	return int(res)
}


// EverNeTtextText wraps MEOS C function ever_ne_ttext_text.
func EverNeTtextText(temp Temporal, txt string) int {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.ever_ne_ttext_text(temp.Inner(), _c_txt)
	return int(res)
}


// TeqBoolTbool wraps MEOS C function teq_bool_tbool.
func TeqBoolTbool(b bool, temp Temporal) Temporal {
	res := C.teq_bool_tbool(C.bool(b), temp.Inner())
	return CreateTemporal(res)
}


// TeqFloatTfloat wraps MEOS C function teq_float_tfloat.
func TeqFloatTfloat(d float64, temp Temporal) Temporal {
	res := C.teq_float_tfloat(C.double(d), temp.Inner())
	return CreateTemporal(res)
}


// TeqIntTint wraps MEOS C function teq_int_tint.
func TeqIntTint(i int, temp Temporal) Temporal {
	res := C.teq_int_tint(C.int(i), temp.Inner())
	return CreateTemporal(res)
}


// TeqTboolBool wraps MEOS C function teq_tbool_bool.
func TeqTboolBool(temp Temporal, b bool) Temporal {
	res := C.teq_tbool_bool(temp.Inner(), C.bool(b))
	return CreateTemporal(res)
}


// TeqTemporalTemporal wraps MEOS C function teq_temporal_temporal.
func TeqTemporalTemporal(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.teq_temporal_temporal(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TeqTextTtext wraps MEOS C function teq_text_ttext.
func TeqTextTtext(txt string, temp Temporal) Temporal {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.teq_text_ttext(_c_txt, temp.Inner())
	return CreateTemporal(res)
}


// TeqTfloatFloat wraps MEOS C function teq_tfloat_float.
func TeqTfloatFloat(temp Temporal, d float64) Temporal {
	res := C.teq_tfloat_float(temp.Inner(), C.double(d))
	return CreateTemporal(res)
}


// TeqTintInt wraps MEOS C function teq_tint_int.
func TeqTintInt(temp Temporal, i int) Temporal {
	res := C.teq_tint_int(temp.Inner(), C.int(i))
	return CreateTemporal(res)
}


// TeqTtextText wraps MEOS C function teq_ttext_text.
func TeqTtextText(temp Temporal, txt string) Temporal {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.teq_ttext_text(temp.Inner(), _c_txt)
	return CreateTemporal(res)
}


// TgeFloatTfloat wraps MEOS C function tge_float_tfloat.
func TgeFloatTfloat(d float64, temp Temporal) Temporal {
	res := C.tge_float_tfloat(C.double(d), temp.Inner())
	return CreateTemporal(res)
}


// TgeIntTint wraps MEOS C function tge_int_tint.
func TgeIntTint(i int, temp Temporal) Temporal {
	res := C.tge_int_tint(C.int(i), temp.Inner())
	return CreateTemporal(res)
}


// TgeTemporalTemporal wraps MEOS C function tge_temporal_temporal.
func TgeTemporalTemporal(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tge_temporal_temporal(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TgeTextTtext wraps MEOS C function tge_text_ttext.
func TgeTextTtext(txt string, temp Temporal) Temporal {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.tge_text_ttext(_c_txt, temp.Inner())
	return CreateTemporal(res)
}


// TgeTfloatFloat wraps MEOS C function tge_tfloat_float.
func TgeTfloatFloat(temp Temporal, d float64) Temporal {
	res := C.tge_tfloat_float(temp.Inner(), C.double(d))
	return CreateTemporal(res)
}


// TgeTintInt wraps MEOS C function tge_tint_int.
func TgeTintInt(temp Temporal, i int) Temporal {
	res := C.tge_tint_int(temp.Inner(), C.int(i))
	return CreateTemporal(res)
}


// TgeTtextText wraps MEOS C function tge_ttext_text.
func TgeTtextText(temp Temporal, txt string) Temporal {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.tge_ttext_text(temp.Inner(), _c_txt)
	return CreateTemporal(res)
}


// TgtFloatTfloat wraps MEOS C function tgt_float_tfloat.
func TgtFloatTfloat(d float64, temp Temporal) Temporal {
	res := C.tgt_float_tfloat(C.double(d), temp.Inner())
	return CreateTemporal(res)
}


// TgtIntTint wraps MEOS C function tgt_int_tint.
func TgtIntTint(i int, temp Temporal) Temporal {
	res := C.tgt_int_tint(C.int(i), temp.Inner())
	return CreateTemporal(res)
}


// TgtTemporalTemporal wraps MEOS C function tgt_temporal_temporal.
func TgtTemporalTemporal(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tgt_temporal_temporal(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TgtTextTtext wraps MEOS C function tgt_text_ttext.
func TgtTextTtext(txt string, temp Temporal) Temporal {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.tgt_text_ttext(_c_txt, temp.Inner())
	return CreateTemporal(res)
}


// TgtTfloatFloat wraps MEOS C function tgt_tfloat_float.
func TgtTfloatFloat(temp Temporal, d float64) Temporal {
	res := C.tgt_tfloat_float(temp.Inner(), C.double(d))
	return CreateTemporal(res)
}


// TgtTintInt wraps MEOS C function tgt_tint_int.
func TgtTintInt(temp Temporal, i int) Temporal {
	res := C.tgt_tint_int(temp.Inner(), C.int(i))
	return CreateTemporal(res)
}


// TgtTtextText wraps MEOS C function tgt_ttext_text.
func TgtTtextText(temp Temporal, txt string) Temporal {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.tgt_ttext_text(temp.Inner(), _c_txt)
	return CreateTemporal(res)
}


// TleFloatTfloat wraps MEOS C function tle_float_tfloat.
func TleFloatTfloat(d float64, temp Temporal) Temporal {
	res := C.tle_float_tfloat(C.double(d), temp.Inner())
	return CreateTemporal(res)
}


// TleIntTint wraps MEOS C function tle_int_tint.
func TleIntTint(i int, temp Temporal) Temporal {
	res := C.tle_int_tint(C.int(i), temp.Inner())
	return CreateTemporal(res)
}


// TleTemporalTemporal wraps MEOS C function tle_temporal_temporal.
func TleTemporalTemporal(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tle_temporal_temporal(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TleTextTtext wraps MEOS C function tle_text_ttext.
func TleTextTtext(txt string, temp Temporal) Temporal {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.tle_text_ttext(_c_txt, temp.Inner())
	return CreateTemporal(res)
}


// TleTfloatFloat wraps MEOS C function tle_tfloat_float.
func TleTfloatFloat(temp Temporal, d float64) Temporal {
	res := C.tle_tfloat_float(temp.Inner(), C.double(d))
	return CreateTemporal(res)
}


// TleTintInt wraps MEOS C function tle_tint_int.
func TleTintInt(temp Temporal, i int) Temporal {
	res := C.tle_tint_int(temp.Inner(), C.int(i))
	return CreateTemporal(res)
}


// TleTtextText wraps MEOS C function tle_ttext_text.
func TleTtextText(temp Temporal, txt string) Temporal {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.tle_ttext_text(temp.Inner(), _c_txt)
	return CreateTemporal(res)
}


// TltFloatTfloat wraps MEOS C function tlt_float_tfloat.
func TltFloatTfloat(d float64, temp Temporal) Temporal {
	res := C.tlt_float_tfloat(C.double(d), temp.Inner())
	return CreateTemporal(res)
}


// TltIntTint wraps MEOS C function tlt_int_tint.
func TltIntTint(i int, temp Temporal) Temporal {
	res := C.tlt_int_tint(C.int(i), temp.Inner())
	return CreateTemporal(res)
}


// TltTemporalTemporal wraps MEOS C function tlt_temporal_temporal.
func TltTemporalTemporal(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tlt_temporal_temporal(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TltTextTtext wraps MEOS C function tlt_text_ttext.
func TltTextTtext(txt string, temp Temporal) Temporal {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.tlt_text_ttext(_c_txt, temp.Inner())
	return CreateTemporal(res)
}


// TltTfloatFloat wraps MEOS C function tlt_tfloat_float.
func TltTfloatFloat(temp Temporal, d float64) Temporal {
	res := C.tlt_tfloat_float(temp.Inner(), C.double(d))
	return CreateTemporal(res)
}


// TltTintInt wraps MEOS C function tlt_tint_int.
func TltTintInt(temp Temporal, i int) Temporal {
	res := C.tlt_tint_int(temp.Inner(), C.int(i))
	return CreateTemporal(res)
}


// TltTtextText wraps MEOS C function tlt_ttext_text.
func TltTtextText(temp Temporal, txt string) Temporal {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.tlt_ttext_text(temp.Inner(), _c_txt)
	return CreateTemporal(res)
}


// TneBoolTbool wraps MEOS C function tne_bool_tbool.
func TneBoolTbool(b bool, temp Temporal) Temporal {
	res := C.tne_bool_tbool(C.bool(b), temp.Inner())
	return CreateTemporal(res)
}


// TneFloatTfloat wraps MEOS C function tne_float_tfloat.
func TneFloatTfloat(d float64, temp Temporal) Temporal {
	res := C.tne_float_tfloat(C.double(d), temp.Inner())
	return CreateTemporal(res)
}


// TneIntTint wraps MEOS C function tne_int_tint.
func TneIntTint(i int, temp Temporal) Temporal {
	res := C.tne_int_tint(C.int(i), temp.Inner())
	return CreateTemporal(res)
}


// TneTboolBool wraps MEOS C function tne_tbool_bool.
func TneTboolBool(temp Temporal, b bool) Temporal {
	res := C.tne_tbool_bool(temp.Inner(), C.bool(b))
	return CreateTemporal(res)
}


// TneTemporalTemporal wraps MEOS C function tne_temporal_temporal.
func TneTemporalTemporal(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tne_temporal_temporal(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TneTextTtext wraps MEOS C function tne_text_ttext.
func TneTextTtext(txt string, temp Temporal) Temporal {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.tne_text_ttext(_c_txt, temp.Inner())
	return CreateTemporal(res)
}


// TneTfloatFloat wraps MEOS C function tne_tfloat_float.
func TneTfloatFloat(temp Temporal, d float64) Temporal {
	res := C.tne_tfloat_float(temp.Inner(), C.double(d))
	return CreateTemporal(res)
}


// TneTintInt wraps MEOS C function tne_tint_int.
func TneTintInt(temp Temporal, i int) Temporal {
	res := C.tne_tint_int(temp.Inner(), C.int(i))
	return CreateTemporal(res)
}


// TneTtextText wraps MEOS C function tne_ttext_text.
func TneTtextText(temp Temporal, txt string) Temporal {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.tne_ttext_text(temp.Inner(), _c_txt)
	return CreateTemporal(res)
}


// TemporalSpans wraps MEOS C function temporal_spans.
func TemporalSpans(temp Temporal) (*Span, int) {
	var _out_count C.int
	res := C.temporal_spans(temp.Inner(), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// TemporalSplitEachNSpans wraps MEOS C function temporal_split_each_n_spans.
func TemporalSplitEachNSpans(temp Temporal, elem_count int) (*Span, int) {
	var _out_count C.int
	res := C.temporal_split_each_n_spans(temp.Inner(), C.int(elem_count), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// TemporalSplitNSpans wraps MEOS C function temporal_split_n_spans.
func TemporalSplitNSpans(temp Temporal, span_count int) (*Span, int) {
	var _out_count C.int
	res := C.temporal_split_n_spans(temp.Inner(), C.int(span_count), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// TnumberSplitEachNTboxes wraps MEOS C function tnumber_split_each_n_tboxes.
func TnumberSplitEachNTboxes(temp Temporal, elem_count int) (*TBox, int) {
	var _out_count C.int
	res := C.tnumber_split_each_n_tboxes(temp.Inner(), C.int(elem_count), &_out_count)
	return &TBox{_inner: res}, int(_out_count)
}


// TnumberSplitNTboxes wraps MEOS C function tnumber_split_n_tboxes.
func TnumberSplitNTboxes(temp Temporal, box_count int) (*TBox, int) {
	var _out_count C.int
	res := C.tnumber_split_n_tboxes(temp.Inner(), C.int(box_count), &_out_count)
	return &TBox{_inner: res}, int(_out_count)
}


// TnumberTboxes wraps MEOS C function tnumber_tboxes.
func TnumberTboxes(temp Temporal) (*TBox, int) {
	var _out_count C.int
	res := C.tnumber_tboxes(temp.Inner(), &_out_count)
	return &TBox{_inner: res}, int(_out_count)
}


// AdjacentNumspanTnumber wraps MEOS C function adjacent_numspan_tnumber.
func AdjacentNumspanTnumber(s *Span, temp Temporal) bool {
	res := C.adjacent_numspan_tnumber(s._inner, temp.Inner())
	return bool(res)
}


// AdjacentTBOXTnumber wraps MEOS C function adjacent_tbox_tnumber.
func AdjacentTBOXTnumber(box *TBox, temp Temporal) bool {
	res := C.adjacent_tbox_tnumber(box._inner, temp.Inner())
	return bool(res)
}


// AdjacentTemporalTemporal wraps MEOS C function adjacent_temporal_temporal.
func AdjacentTemporalTemporal(temp1 Temporal, temp2 Temporal) bool {
	res := C.adjacent_temporal_temporal(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// AdjacentTemporalTstzspan wraps MEOS C function adjacent_temporal_tstzspan.
func AdjacentTemporalTstzspan(temp Temporal, s *Span) bool {
	res := C.adjacent_temporal_tstzspan(temp.Inner(), s._inner)
	return bool(res)
}


// AdjacentTnumberNumspan wraps MEOS C function adjacent_tnumber_numspan.
func AdjacentTnumberNumspan(temp Temporal, s *Span) bool {
	res := C.adjacent_tnumber_numspan(temp.Inner(), s._inner)
	return bool(res)
}


// AdjacentTnumberTBOX wraps MEOS C function adjacent_tnumber_tbox.
func AdjacentTnumberTBOX(temp Temporal, box *TBox) bool {
	res := C.adjacent_tnumber_tbox(temp.Inner(), box._inner)
	return bool(res)
}


// AdjacentTnumberTnumber wraps MEOS C function adjacent_tnumber_tnumber.
func AdjacentTnumberTnumber(temp1 Temporal, temp2 Temporal) bool {
	res := C.adjacent_tnumber_tnumber(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// AdjacentTstzspanTemporal wraps MEOS C function adjacent_tstzspan_temporal.
func AdjacentTstzspanTemporal(s *Span, temp Temporal) bool {
	res := C.adjacent_tstzspan_temporal(s._inner, temp.Inner())
	return bool(res)
}


// ContainedNumspanTnumber wraps MEOS C function contained_numspan_tnumber.
func ContainedNumspanTnumber(s *Span, temp Temporal) bool {
	res := C.contained_numspan_tnumber(s._inner, temp.Inner())
	return bool(res)
}


// ContainedTBOXTnumber wraps MEOS C function contained_tbox_tnumber.
func ContainedTBOXTnumber(box *TBox, temp Temporal) bool {
	res := C.contained_tbox_tnumber(box._inner, temp.Inner())
	return bool(res)
}


// ContainedTemporalTemporal wraps MEOS C function contained_temporal_temporal.
func ContainedTemporalTemporal(temp1 Temporal, temp2 Temporal) bool {
	res := C.contained_temporal_temporal(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// ContainedTemporalTstzspan wraps MEOS C function contained_temporal_tstzspan.
func ContainedTemporalTstzspan(temp Temporal, s *Span) bool {
	res := C.contained_temporal_tstzspan(temp.Inner(), s._inner)
	return bool(res)
}


// ContainedTnumberNumspan wraps MEOS C function contained_tnumber_numspan.
func ContainedTnumberNumspan(temp Temporal, s *Span) bool {
	res := C.contained_tnumber_numspan(temp.Inner(), s._inner)
	return bool(res)
}


// ContainedTnumberTBOX wraps MEOS C function contained_tnumber_tbox.
func ContainedTnumberTBOX(temp Temporal, box *TBox) bool {
	res := C.contained_tnumber_tbox(temp.Inner(), box._inner)
	return bool(res)
}


// ContainedTnumberTnumber wraps MEOS C function contained_tnumber_tnumber.
func ContainedTnumberTnumber(temp1 Temporal, temp2 Temporal) bool {
	res := C.contained_tnumber_tnumber(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// ContainedTstzspanTemporal wraps MEOS C function contained_tstzspan_temporal.
func ContainedTstzspanTemporal(s *Span, temp Temporal) bool {
	res := C.contained_tstzspan_temporal(s._inner, temp.Inner())
	return bool(res)
}


// ContainsNumspanTnumber wraps MEOS C function contains_numspan_tnumber.
func ContainsNumspanTnumber(s *Span, temp Temporal) bool {
	res := C.contains_numspan_tnumber(s._inner, temp.Inner())
	return bool(res)
}


// ContainsTBOXTnumber wraps MEOS C function contains_tbox_tnumber.
func ContainsTBOXTnumber(box *TBox, temp Temporal) bool {
	res := C.contains_tbox_tnumber(box._inner, temp.Inner())
	return bool(res)
}


// ContainsTemporalTstzspan wraps MEOS C function contains_temporal_tstzspan.
func ContainsTemporalTstzspan(temp Temporal, s *Span) bool {
	res := C.contains_temporal_tstzspan(temp.Inner(), s._inner)
	return bool(res)
}


// ContainsTemporalTemporal wraps MEOS C function contains_temporal_temporal.
func ContainsTemporalTemporal(temp1 Temporal, temp2 Temporal) bool {
	res := C.contains_temporal_temporal(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// ContainsTnumberNumspan wraps MEOS C function contains_tnumber_numspan.
func ContainsTnumberNumspan(temp Temporal, s *Span) bool {
	res := C.contains_tnumber_numspan(temp.Inner(), s._inner)
	return bool(res)
}


// ContainsTnumberTBOX wraps MEOS C function contains_tnumber_tbox.
func ContainsTnumberTBOX(temp Temporal, box *TBox) bool {
	res := C.contains_tnumber_tbox(temp.Inner(), box._inner)
	return bool(res)
}


// ContainsTnumberTnumber wraps MEOS C function contains_tnumber_tnumber.
func ContainsTnumberTnumber(temp1 Temporal, temp2 Temporal) bool {
	res := C.contains_tnumber_tnumber(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// ContainsTstzspanTemporal wraps MEOS C function contains_tstzspan_temporal.
func ContainsTstzspanTemporal(s *Span, temp Temporal) bool {
	res := C.contains_tstzspan_temporal(s._inner, temp.Inner())
	return bool(res)
}


// OverlapsNumspanTnumber wraps MEOS C function overlaps_numspan_tnumber.
func OverlapsNumspanTnumber(s *Span, temp Temporal) bool {
	res := C.overlaps_numspan_tnumber(s._inner, temp.Inner())
	return bool(res)
}


// OverlapsTBOXTnumber wraps MEOS C function overlaps_tbox_tnumber.
func OverlapsTBOXTnumber(box *TBox, temp Temporal) bool {
	res := C.overlaps_tbox_tnumber(box._inner, temp.Inner())
	return bool(res)
}


// OverlapsTemporalTemporal wraps MEOS C function overlaps_temporal_temporal.
func OverlapsTemporalTemporal(temp1 Temporal, temp2 Temporal) bool {
	res := C.overlaps_temporal_temporal(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// OverlapsTemporalTstzspan wraps MEOS C function overlaps_temporal_tstzspan.
func OverlapsTemporalTstzspan(temp Temporal, s *Span) bool {
	res := C.overlaps_temporal_tstzspan(temp.Inner(), s._inner)
	return bool(res)
}


// OverlapsTnumberNumspan wraps MEOS C function overlaps_tnumber_numspan.
func OverlapsTnumberNumspan(temp Temporal, s *Span) bool {
	res := C.overlaps_tnumber_numspan(temp.Inner(), s._inner)
	return bool(res)
}


// OverlapsTnumberTBOX wraps MEOS C function overlaps_tnumber_tbox.
func OverlapsTnumberTBOX(temp Temporal, box *TBox) bool {
	res := C.overlaps_tnumber_tbox(temp.Inner(), box._inner)
	return bool(res)
}


// OverlapsTnumberTnumber wraps MEOS C function overlaps_tnumber_tnumber.
func OverlapsTnumberTnumber(temp1 Temporal, temp2 Temporal) bool {
	res := C.overlaps_tnumber_tnumber(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// OverlapsTstzspanTemporal wraps MEOS C function overlaps_tstzspan_temporal.
func OverlapsTstzspanTemporal(s *Span, temp Temporal) bool {
	res := C.overlaps_tstzspan_temporal(s._inner, temp.Inner())
	return bool(res)
}


// SameNumspanTnumber wraps MEOS C function same_numspan_tnumber.
func SameNumspanTnumber(s *Span, temp Temporal) bool {
	res := C.same_numspan_tnumber(s._inner, temp.Inner())
	return bool(res)
}


// SameTBOXTnumber wraps MEOS C function same_tbox_tnumber.
func SameTBOXTnumber(box *TBox, temp Temporal) bool {
	res := C.same_tbox_tnumber(box._inner, temp.Inner())
	return bool(res)
}


// SameTemporalTemporal wraps MEOS C function same_temporal_temporal.
func SameTemporalTemporal(temp1 Temporal, temp2 Temporal) bool {
	res := C.same_temporal_temporal(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// SameTemporalTstzspan wraps MEOS C function same_temporal_tstzspan.
func SameTemporalTstzspan(temp Temporal, s *Span) bool {
	res := C.same_temporal_tstzspan(temp.Inner(), s._inner)
	return bool(res)
}


// SameTnumberNumspan wraps MEOS C function same_tnumber_numspan.
func SameTnumberNumspan(temp Temporal, s *Span) bool {
	res := C.same_tnumber_numspan(temp.Inner(), s._inner)
	return bool(res)
}


// SameTnumberTBOX wraps MEOS C function same_tnumber_tbox.
func SameTnumberTBOX(temp Temporal, box *TBox) bool {
	res := C.same_tnumber_tbox(temp.Inner(), box._inner)
	return bool(res)
}


// SameTnumberTnumber wraps MEOS C function same_tnumber_tnumber.
func SameTnumberTnumber(temp1 Temporal, temp2 Temporal) bool {
	res := C.same_tnumber_tnumber(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// SameTstzspanTemporal wraps MEOS C function same_tstzspan_temporal.
func SameTstzspanTemporal(s *Span, temp Temporal) bool {
	res := C.same_tstzspan_temporal(s._inner, temp.Inner())
	return bool(res)
}


// AfterTBOXTnumber wraps MEOS C function after_tbox_tnumber.
func AfterTBOXTnumber(box *TBox, temp Temporal) bool {
	res := C.after_tbox_tnumber(box._inner, temp.Inner())
	return bool(res)
}


// AfterTemporalTstzspan wraps MEOS C function after_temporal_tstzspan.
func AfterTemporalTstzspan(temp Temporal, s *Span) bool {
	res := C.after_temporal_tstzspan(temp.Inner(), s._inner)
	return bool(res)
}


// AfterTemporalTemporal wraps MEOS C function after_temporal_temporal.
func AfterTemporalTemporal(temp1 Temporal, temp2 Temporal) bool {
	res := C.after_temporal_temporal(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// AfterTnumberTBOX wraps MEOS C function after_tnumber_tbox.
func AfterTnumberTBOX(temp Temporal, box *TBox) bool {
	res := C.after_tnumber_tbox(temp.Inner(), box._inner)
	return bool(res)
}


// AfterTnumberTnumber wraps MEOS C function after_tnumber_tnumber.
func AfterTnumberTnumber(temp1 Temporal, temp2 Temporal) bool {
	res := C.after_tnumber_tnumber(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// AfterTstzspanTemporal wraps MEOS C function after_tstzspan_temporal.
func AfterTstzspanTemporal(s *Span, temp Temporal) bool {
	res := C.after_tstzspan_temporal(s._inner, temp.Inner())
	return bool(res)
}


// BeforeTBOXTnumber wraps MEOS C function before_tbox_tnumber.
func BeforeTBOXTnumber(box *TBox, temp Temporal) bool {
	res := C.before_tbox_tnumber(box._inner, temp.Inner())
	return bool(res)
}


// BeforeTemporalTstzspan wraps MEOS C function before_temporal_tstzspan.
func BeforeTemporalTstzspan(temp Temporal, s *Span) bool {
	res := C.before_temporal_tstzspan(temp.Inner(), s._inner)
	return bool(res)
}


// BeforeTemporalTemporal wraps MEOS C function before_temporal_temporal.
func BeforeTemporalTemporal(temp1 Temporal, temp2 Temporal) bool {
	res := C.before_temporal_temporal(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// BeforeTnumberTBOX wraps MEOS C function before_tnumber_tbox.
func BeforeTnumberTBOX(temp Temporal, box *TBox) bool {
	res := C.before_tnumber_tbox(temp.Inner(), box._inner)
	return bool(res)
}


// BeforeTnumberTnumber wraps MEOS C function before_tnumber_tnumber.
func BeforeTnumberTnumber(temp1 Temporal, temp2 Temporal) bool {
	res := C.before_tnumber_tnumber(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// BeforeTstzspanTemporal wraps MEOS C function before_tstzspan_temporal.
func BeforeTstzspanTemporal(s *Span, temp Temporal) bool {
	res := C.before_tstzspan_temporal(s._inner, temp.Inner())
	return bool(res)
}


// LeftTBOXTnumber wraps MEOS C function left_tbox_tnumber.
func LeftTBOXTnumber(box *TBox, temp Temporal) bool {
	res := C.left_tbox_tnumber(box._inner, temp.Inner())
	return bool(res)
}


// LeftNumspanTnumber wraps MEOS C function left_numspan_tnumber.
func LeftNumspanTnumber(s *Span, temp Temporal) bool {
	res := C.left_numspan_tnumber(s._inner, temp.Inner())
	return bool(res)
}


// LeftTnumberNumspan wraps MEOS C function left_tnumber_numspan.
func LeftTnumberNumspan(temp Temporal, s *Span) bool {
	res := C.left_tnumber_numspan(temp.Inner(), s._inner)
	return bool(res)
}


// LeftTnumberTBOX wraps MEOS C function left_tnumber_tbox.
func LeftTnumberTBOX(temp Temporal, box *TBox) bool {
	res := C.left_tnumber_tbox(temp.Inner(), box._inner)
	return bool(res)
}


// LeftTnumberTnumber wraps MEOS C function left_tnumber_tnumber.
func LeftTnumberTnumber(temp1 Temporal, temp2 Temporal) bool {
	res := C.left_tnumber_tnumber(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// OverafterTBOXTnumber wraps MEOS C function overafter_tbox_tnumber.
func OverafterTBOXTnumber(box *TBox, temp Temporal) bool {
	res := C.overafter_tbox_tnumber(box._inner, temp.Inner())
	return bool(res)
}


// OverafterTemporalTstzspan wraps MEOS C function overafter_temporal_tstzspan.
func OverafterTemporalTstzspan(temp Temporal, s *Span) bool {
	res := C.overafter_temporal_tstzspan(temp.Inner(), s._inner)
	return bool(res)
}


// OverafterTemporalTemporal wraps MEOS C function overafter_temporal_temporal.
func OverafterTemporalTemporal(temp1 Temporal, temp2 Temporal) bool {
	res := C.overafter_temporal_temporal(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// OverafterTnumberTBOX wraps MEOS C function overafter_tnumber_tbox.
func OverafterTnumberTBOX(temp Temporal, box *TBox) bool {
	res := C.overafter_tnumber_tbox(temp.Inner(), box._inner)
	return bool(res)
}


// OverafterTnumberTnumber wraps MEOS C function overafter_tnumber_tnumber.
func OverafterTnumberTnumber(temp1 Temporal, temp2 Temporal) bool {
	res := C.overafter_tnumber_tnumber(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// OverafterTstzspanTemporal wraps MEOS C function overafter_tstzspan_temporal.
func OverafterTstzspanTemporal(s *Span, temp Temporal) bool {
	res := C.overafter_tstzspan_temporal(s._inner, temp.Inner())
	return bool(res)
}


// OverbeforeTBOXTnumber wraps MEOS C function overbefore_tbox_tnumber.
func OverbeforeTBOXTnumber(box *TBox, temp Temporal) bool {
	res := C.overbefore_tbox_tnumber(box._inner, temp.Inner())
	return bool(res)
}


// OverbeforeTemporalTstzspan wraps MEOS C function overbefore_temporal_tstzspan.
func OverbeforeTemporalTstzspan(temp Temporal, s *Span) bool {
	res := C.overbefore_temporal_tstzspan(temp.Inner(), s._inner)
	return bool(res)
}


// OverbeforeTemporalTemporal wraps MEOS C function overbefore_temporal_temporal.
func OverbeforeTemporalTemporal(temp1 Temporal, temp2 Temporal) bool {
	res := C.overbefore_temporal_temporal(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// OverbeforeTnumberTBOX wraps MEOS C function overbefore_tnumber_tbox.
func OverbeforeTnumberTBOX(temp Temporal, box *TBox) bool {
	res := C.overbefore_tnumber_tbox(temp.Inner(), box._inner)
	return bool(res)
}


// OverbeforeTnumberTnumber wraps MEOS C function overbefore_tnumber_tnumber.
func OverbeforeTnumberTnumber(temp1 Temporal, temp2 Temporal) bool {
	res := C.overbefore_tnumber_tnumber(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// OverbeforeTstzspanTemporal wraps MEOS C function overbefore_tstzspan_temporal.
func OverbeforeTstzspanTemporal(s *Span, temp Temporal) bool {
	res := C.overbefore_tstzspan_temporal(s._inner, temp.Inner())
	return bool(res)
}


// OverleftNumspanTnumber wraps MEOS C function overleft_numspan_tnumber.
func OverleftNumspanTnumber(s *Span, temp Temporal) bool {
	res := C.overleft_numspan_tnumber(s._inner, temp.Inner())
	return bool(res)
}


// OverleftTBOXTnumber wraps MEOS C function overleft_tbox_tnumber.
func OverleftTBOXTnumber(box *TBox, temp Temporal) bool {
	res := C.overleft_tbox_tnumber(box._inner, temp.Inner())
	return bool(res)
}


// OverleftTnumberNumspan wraps MEOS C function overleft_tnumber_numspan.
func OverleftTnumberNumspan(temp Temporal, s *Span) bool {
	res := C.overleft_tnumber_numspan(temp.Inner(), s._inner)
	return bool(res)
}


// OverleftTnumberTBOX wraps MEOS C function overleft_tnumber_tbox.
func OverleftTnumberTBOX(temp Temporal, box *TBox) bool {
	res := C.overleft_tnumber_tbox(temp.Inner(), box._inner)
	return bool(res)
}


// OverleftTnumberTnumber wraps MEOS C function overleft_tnumber_tnumber.
func OverleftTnumberTnumber(temp1 Temporal, temp2 Temporal) bool {
	res := C.overleft_tnumber_tnumber(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// OverrightNumspanTnumber wraps MEOS C function overright_numspan_tnumber.
func OverrightNumspanTnumber(s *Span, temp Temporal) bool {
	res := C.overright_numspan_tnumber(s._inner, temp.Inner())
	return bool(res)
}


// OverrightTBOXTnumber wraps MEOS C function overright_tbox_tnumber.
func OverrightTBOXTnumber(box *TBox, temp Temporal) bool {
	res := C.overright_tbox_tnumber(box._inner, temp.Inner())
	return bool(res)
}


// OverrightTnumberNumspan wraps MEOS C function overright_tnumber_numspan.
func OverrightTnumberNumspan(temp Temporal, s *Span) bool {
	res := C.overright_tnumber_numspan(temp.Inner(), s._inner)
	return bool(res)
}


// OverrightTnumberTBOX wraps MEOS C function overright_tnumber_tbox.
func OverrightTnumberTBOX(temp Temporal, box *TBox) bool {
	res := C.overright_tnumber_tbox(temp.Inner(), box._inner)
	return bool(res)
}


// OverrightTnumberTnumber wraps MEOS C function overright_tnumber_tnumber.
func OverrightTnumberTnumber(temp1 Temporal, temp2 Temporal) bool {
	res := C.overright_tnumber_tnumber(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// RightNumspanTnumber wraps MEOS C function right_numspan_tnumber.
func RightNumspanTnumber(s *Span, temp Temporal) bool {
	res := C.right_numspan_tnumber(s._inner, temp.Inner())
	return bool(res)
}


// RightTBOXTnumber wraps MEOS C function right_tbox_tnumber.
func RightTBOXTnumber(box *TBox, temp Temporal) bool {
	res := C.right_tbox_tnumber(box._inner, temp.Inner())
	return bool(res)
}


// RightTnumberNumspan wraps MEOS C function right_tnumber_numspan.
func RightTnumberNumspan(temp Temporal, s *Span) bool {
	res := C.right_tnumber_numspan(temp.Inner(), s._inner)
	return bool(res)
}


// RightTnumberTBOX wraps MEOS C function right_tnumber_tbox.
func RightTnumberTBOX(temp Temporal, box *TBox) bool {
	res := C.right_tnumber_tbox(temp.Inner(), box._inner)
	return bool(res)
}


// RightTnumberTnumber wraps MEOS C function right_tnumber_tnumber.
func RightTnumberTnumber(temp1 Temporal, temp2 Temporal) bool {
	res := C.right_tnumber_tnumber(temp1.Inner(), temp2.Inner())
	return bool(res)
}


// TandBoolTbool wraps MEOS C function tand_bool_tbool.
func TandBoolTbool(b bool, temp Temporal) Temporal {
	res := C.tand_bool_tbool(C.bool(b), temp.Inner())
	return CreateTemporal(res)
}


// TandTboolBool wraps MEOS C function tand_tbool_bool.
func TandTboolBool(temp Temporal, b bool) Temporal {
	res := C.tand_tbool_bool(temp.Inner(), C.bool(b))
	return CreateTemporal(res)
}


// TandTboolTbool wraps MEOS C function tand_tbool_tbool.
func TandTboolTbool(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tand_tbool_tbool(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TboolWhenTrue wraps MEOS C function tbool_when_true.
func TboolWhenTrue(temp Temporal) *SpanSet {
	res := C.tbool_when_true(temp.Inner())
	return &SpanSet{_inner: res}
}


// TnotTbool wraps MEOS C function tnot_tbool.
func TnotTbool(temp Temporal) Temporal {
	res := C.tnot_tbool(temp.Inner())
	return CreateTemporal(res)
}


// TorBoolTbool wraps MEOS C function tor_bool_tbool.
func TorBoolTbool(b bool, temp Temporal) Temporal {
	res := C.tor_bool_tbool(C.bool(b), temp.Inner())
	return CreateTemporal(res)
}


// TorTboolBool wraps MEOS C function tor_tbool_bool.
func TorTboolBool(temp Temporal, b bool) Temporal {
	res := C.tor_tbool_bool(temp.Inner(), C.bool(b))
	return CreateTemporal(res)
}


// TorTboolTbool wraps MEOS C function tor_tbool_tbool.
func TorTboolTbool(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tor_tbool_tbool(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// AddFloatTfloat wraps MEOS C function add_float_tfloat.
func AddFloatTfloat(d float64, tnumber Temporal) Temporal {
	res := C.add_float_tfloat(C.double(d), tnumber.Inner())
	return CreateTemporal(res)
}


// AddIntTint wraps MEOS C function add_int_tint.
func AddIntTint(i int, tnumber Temporal) Temporal {
	res := C.add_int_tint(C.int(i), tnumber.Inner())
	return CreateTemporal(res)
}


// AddTfloatFloat wraps MEOS C function add_tfloat_float.
func AddTfloatFloat(tnumber Temporal, d float64) Temporal {
	res := C.add_tfloat_float(tnumber.Inner(), C.double(d))
	return CreateTemporal(res)
}


// AddTintInt wraps MEOS C function add_tint_int.
func AddTintInt(tnumber Temporal, i int) Temporal {
	res := C.add_tint_int(tnumber.Inner(), C.int(i))
	return CreateTemporal(res)
}


// AddTnumberTnumber wraps MEOS C function add_tnumber_tnumber.
func AddTnumberTnumber(tnumber1 Temporal, tnumber2 Temporal) Temporal {
	res := C.add_tnumber_tnumber(tnumber1.Inner(), tnumber2.Inner())
	return CreateTemporal(res)
}


// DivFloatTfloat wraps MEOS C function div_float_tfloat.
func DivFloatTfloat(d float64, tnumber Temporal) Temporal {
	res := C.div_float_tfloat(C.double(d), tnumber.Inner())
	return CreateTemporal(res)
}


// DivIntTint wraps MEOS C function div_int_tint.
func DivIntTint(i int, tnumber Temporal) Temporal {
	res := C.div_int_tint(C.int(i), tnumber.Inner())
	return CreateTemporal(res)
}


// DivTfloatFloat wraps MEOS C function div_tfloat_float.
func DivTfloatFloat(tnumber Temporal, d float64) Temporal {
	res := C.div_tfloat_float(tnumber.Inner(), C.double(d))
	return CreateTemporal(res)
}


// DivTintInt wraps MEOS C function div_tint_int.
func DivTintInt(tnumber Temporal, i int) Temporal {
	res := C.div_tint_int(tnumber.Inner(), C.int(i))
	return CreateTemporal(res)
}


// DivTnumberTnumber wraps MEOS C function div_tnumber_tnumber.
func DivTnumberTnumber(tnumber1 Temporal, tnumber2 Temporal) Temporal {
	res := C.div_tnumber_tnumber(tnumber1.Inner(), tnumber2.Inner())
	return CreateTemporal(res)
}


// MultFloatTfloat wraps MEOS C function mult_float_tfloat.
func MultFloatTfloat(d float64, tnumber Temporal) Temporal {
	res := C.mult_float_tfloat(C.double(d), tnumber.Inner())
	return CreateTemporal(res)
}


// MultIntTint wraps MEOS C function mult_int_tint.
func MultIntTint(i int, tnumber Temporal) Temporal {
	res := C.mult_int_tint(C.int(i), tnumber.Inner())
	return CreateTemporal(res)
}


// MultTfloatFloat wraps MEOS C function mult_tfloat_float.
func MultTfloatFloat(tnumber Temporal, d float64) Temporal {
	res := C.mult_tfloat_float(tnumber.Inner(), C.double(d))
	return CreateTemporal(res)
}


// MultTintInt wraps MEOS C function mult_tint_int.
func MultTintInt(tnumber Temporal, i int) Temporal {
	res := C.mult_tint_int(tnumber.Inner(), C.int(i))
	return CreateTemporal(res)
}


// MultTnumberTnumber wraps MEOS C function mult_tnumber_tnumber.
func MultTnumberTnumber(tnumber1 Temporal, tnumber2 Temporal) Temporal {
	res := C.mult_tnumber_tnumber(tnumber1.Inner(), tnumber2.Inner())
	return CreateTemporal(res)
}


// SubFloatTfloat wraps MEOS C function sub_float_tfloat.
func SubFloatTfloat(d float64, tnumber Temporal) Temporal {
	res := C.sub_float_tfloat(C.double(d), tnumber.Inner())
	return CreateTemporal(res)
}


// SubIntTint wraps MEOS C function sub_int_tint.
func SubIntTint(i int, tnumber Temporal) Temporal {
	res := C.sub_int_tint(C.int(i), tnumber.Inner())
	return CreateTemporal(res)
}


// SubTfloatFloat wraps MEOS C function sub_tfloat_float.
func SubTfloatFloat(tnumber Temporal, d float64) Temporal {
	res := C.sub_tfloat_float(tnumber.Inner(), C.double(d))
	return CreateTemporal(res)
}


// SubTintInt wraps MEOS C function sub_tint_int.
func SubTintInt(tnumber Temporal, i int) Temporal {
	res := C.sub_tint_int(tnumber.Inner(), C.int(i))
	return CreateTemporal(res)
}


// SubTnumberTnumber wraps MEOS C function sub_tnumber_tnumber.
func SubTnumberTnumber(tnumber1 Temporal, tnumber2 Temporal) Temporal {
	res := C.sub_tnumber_tnumber(tnumber1.Inner(), tnumber2.Inner())
	return CreateTemporal(res)
}


// TemporalDerivative wraps MEOS C function temporal_derivative.
func TemporalDerivative(temp Temporal) Temporal {
	res := C.temporal_derivative(temp.Inner())
	return CreateTemporal(res)
}


// TfloatExp wraps MEOS C function tfloat_exp.
func TfloatExp(temp Temporal) Temporal {
	res := C.tfloat_exp(temp.Inner())
	return CreateTemporal(res)
}


// TfloatLn wraps MEOS C function tfloat_ln.
func TfloatLn(temp Temporal) Temporal {
	res := C.tfloat_ln(temp.Inner())
	return CreateTemporal(res)
}


// TfloatLog10 wraps MEOS C function tfloat_log10.
func TfloatLog10(temp Temporal) Temporal {
	res := C.tfloat_log10(temp.Inner())
	return CreateTemporal(res)
}


// TnumberAbs wraps MEOS C function tnumber_abs.
func TnumberAbs(temp Temporal) Temporal {
	res := C.tnumber_abs(temp.Inner())
	return CreateTemporal(res)
}


// TnumberTrend wraps MEOS C function tnumber_trend.
func TnumberTrend(temp Temporal) Temporal {
	res := C.tnumber_trend(temp.Inner())
	return CreateTemporal(res)
}


// FloatAngularDifference wraps MEOS C function float_angular_difference.
func FloatAngularDifference(degrees1 float64, degrees2 float64) float64 {
	res := C.float_angular_difference(C.double(degrees1), C.double(degrees2))
	return float64(res)
}


// TnumberAngularDifference wraps MEOS C function tnumber_angular_difference.
func TnumberAngularDifference(temp Temporal) Temporal {
	res := C.tnumber_angular_difference(temp.Inner())
	return CreateTemporal(res)
}


// TnumberDeltaValue wraps MEOS C function tnumber_delta_value.
func TnumberDeltaValue(temp Temporal) Temporal {
	res := C.tnumber_delta_value(temp.Inner())
	return CreateTemporal(res)
}


// TextcatTextTtext wraps MEOS C function textcat_text_ttext.
func TextcatTextTtext(txt string, temp Temporal) Temporal {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.textcat_text_ttext(_c_txt, temp.Inner())
	return CreateTemporal(res)
}


// TextcatTtextText wraps MEOS C function textcat_ttext_text.
func TextcatTtextText(temp Temporal, txt string) Temporal {
	_c_txt := cstring2text(txt)
	defer C.free(unsafe.Pointer(_c_txt))
	res := C.textcat_ttext_text(temp.Inner(), _c_txt)
	return CreateTemporal(res)
}


// TextcatTtextTtext wraps MEOS C function textcat_ttext_ttext.
func TextcatTtextTtext(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.textcat_ttext_ttext(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// TtextInitcap wraps MEOS C function ttext_initcap.
func TtextInitcap(temp Temporal) Temporal {
	res := C.ttext_initcap(temp.Inner())
	return CreateTemporal(res)
}


// TtextUpper wraps MEOS C function ttext_upper.
func TtextUpper(temp Temporal) Temporal {
	res := C.ttext_upper(temp.Inner())
	return CreateTemporal(res)
}


// TtextLower wraps MEOS C function ttext_lower.
func TtextLower(temp Temporal) Temporal {
	res := C.ttext_lower(temp.Inner())
	return CreateTemporal(res)
}


// TdistanceTfloatFloat wraps MEOS C function tdistance_tfloat_float.
func TdistanceTfloatFloat(temp Temporal, d float64) Temporal {
	res := C.tdistance_tfloat_float(temp.Inner(), C.double(d))
	return CreateTemporal(res)
}


// TdistanceTintInt wraps MEOS C function tdistance_tint_int.
func TdistanceTintInt(temp Temporal, i int) Temporal {
	res := C.tdistance_tint_int(temp.Inner(), C.int(i))
	return CreateTemporal(res)
}


// TdistanceTnumberTnumber wraps MEOS C function tdistance_tnumber_tnumber.
func TdistanceTnumberTnumber(temp1 Temporal, temp2 Temporal) Temporal {
	res := C.tdistance_tnumber_tnumber(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}


// NadTboxfloatTboxfloat wraps MEOS C function nad_tboxfloat_tboxfloat.
func NadTboxfloatTboxfloat(box1 *TBox, box2 *TBox) float64 {
	res := C.nad_tboxfloat_tboxfloat(box1._inner, box2._inner)
	return float64(res)
}


// NadTboxintTboxint wraps MEOS C function nad_tboxint_tboxint.
func NadTboxintTboxint(box1 *TBox, box2 *TBox) int {
	res := C.nad_tboxint_tboxint(box1._inner, box2._inner)
	return int(res)
}


// NadTfloatFloat wraps MEOS C function nad_tfloat_float.
func NadTfloatFloat(temp Temporal, d float64) float64 {
	res := C.nad_tfloat_float(temp.Inner(), C.double(d))
	return float64(res)
}


// NadTfloatTfloat wraps MEOS C function nad_tfloat_tfloat.
func NadTfloatTfloat(temp1 Temporal, temp2 Temporal) float64 {
	res := C.nad_tfloat_tfloat(temp1.Inner(), temp2.Inner())
	return float64(res)
}


// NadTfloatTBOX wraps MEOS C function nad_tfloat_tbox.
func NadTfloatTBOX(temp Temporal, box *TBox) float64 {
	res := C.nad_tfloat_tbox(temp.Inner(), box._inner)
	return float64(res)
}


// NadTintInt wraps MEOS C function nad_tint_int.
func NadTintInt(temp Temporal, i int) int {
	res := C.nad_tint_int(temp.Inner(), C.int(i))
	return int(res)
}


// NadTintTBOX wraps MEOS C function nad_tint_tbox.
func NadTintTBOX(temp Temporal, box *TBox) int {
	res := C.nad_tint_tbox(temp.Inner(), box._inner)
	return int(res)
}


// NadTintTint wraps MEOS C function nad_tint_tint.
func NadTintTint(temp1 Temporal, temp2 Temporal) int {
	res := C.nad_tint_tint(temp1.Inner(), temp2.Inner())
	return int(res)
}


// TboolTandTransfn wraps MEOS C function tbool_tand_transfn.
func TboolTandTransfn(state *SkipList, temp Temporal) *SkipList {
	res := C.tbool_tand_transfn(state._inner, temp.Inner())
	return &SkipList{_inner: res}
}


// TboolTorTransfn wraps MEOS C function tbool_tor_transfn.
func TboolTorTransfn(state *SkipList, temp Temporal) *SkipList {
	res := C.tbool_tor_transfn(state._inner, temp.Inner())
	return &SkipList{_inner: res}
}


// TemporalExtentTransfn wraps MEOS C function temporal_extent_transfn.
func TemporalExtentTransfn(s *Span, temp Temporal) *Span {
	res := C.temporal_extent_transfn(s._inner, temp.Inner())
	return &Span{_inner: res}
}


// TemporalTaggFinalfn wraps MEOS C function temporal_tagg_finalfn.
func TemporalTaggFinalfn(state *SkipList) Temporal {
	res := C.temporal_tagg_finalfn(state._inner)
	return CreateTemporal(res)
}


// TemporalTcountTransfn wraps MEOS C function temporal_tcount_transfn.
func TemporalTcountTransfn(state *SkipList, temp Temporal) *SkipList {
	res := C.temporal_tcount_transfn(state._inner, temp.Inner())
	return &SkipList{_inner: res}
}


// TfloatTmaxTransfn wraps MEOS C function tfloat_tmax_transfn.
func TfloatTmaxTransfn(state *SkipList, temp Temporal) *SkipList {
	res := C.tfloat_tmax_transfn(state._inner, temp.Inner())
	return &SkipList{_inner: res}
}


// TfloatTminTransfn wraps MEOS C function tfloat_tmin_transfn.
func TfloatTminTransfn(state *SkipList, temp Temporal) *SkipList {
	res := C.tfloat_tmin_transfn(state._inner, temp.Inner())
	return &SkipList{_inner: res}
}


// TfloatTsumTransfn wraps MEOS C function tfloat_tsum_transfn.
func TfloatTsumTransfn(state *SkipList, temp Temporal) *SkipList {
	res := C.tfloat_tsum_transfn(state._inner, temp.Inner())
	return &SkipList{_inner: res}
}


// TfloatWmaxTransfn wraps MEOS C function tfloat_wmax_transfn.
func TfloatWmaxTransfn(state *SkipList, temp Temporal, interv timeutil.Timedelta) *SkipList {
	res := C.tfloat_wmax_transfn(state._inner, temp.Inner(), interv.Inner())
	return &SkipList{_inner: res}
}


// TfloatWminTransfn wraps MEOS C function tfloat_wmin_transfn.
func TfloatWminTransfn(state *SkipList, temp Temporal, interv timeutil.Timedelta) *SkipList {
	res := C.tfloat_wmin_transfn(state._inner, temp.Inner(), interv.Inner())
	return &SkipList{_inner: res}
}


// TfloatWsumTransfn wraps MEOS C function tfloat_wsum_transfn.
func TfloatWsumTransfn(state *SkipList, temp Temporal, interv timeutil.Timedelta) *SkipList {
	res := C.tfloat_wsum_transfn(state._inner, temp.Inner(), interv.Inner())
	return &SkipList{_inner: res}
}


// TimestamptzTcountTransfn wraps MEOS C function timestamptz_tcount_transfn.
func TimestamptzTcountTransfn(state *SkipList, t int64) *SkipList {
	res := C.timestamptz_tcount_transfn(state._inner, C.TimestampTz(t))
	return &SkipList{_inner: res}
}


// TintTmaxTransfn wraps MEOS C function tint_tmax_transfn.
func TintTmaxTransfn(state *SkipList, temp Temporal) *SkipList {
	res := C.tint_tmax_transfn(state._inner, temp.Inner())
	return &SkipList{_inner: res}
}


// TintTminTransfn wraps MEOS C function tint_tmin_transfn.
func TintTminTransfn(state *SkipList, temp Temporal) *SkipList {
	res := C.tint_tmin_transfn(state._inner, temp.Inner())
	return &SkipList{_inner: res}
}


// TintTsumTransfn wraps MEOS C function tint_tsum_transfn.
func TintTsumTransfn(state *SkipList, temp Temporal) *SkipList {
	res := C.tint_tsum_transfn(state._inner, temp.Inner())
	return &SkipList{_inner: res}
}


// TintWmaxTransfn wraps MEOS C function tint_wmax_transfn.
func TintWmaxTransfn(state *SkipList, temp Temporal, interv timeutil.Timedelta) *SkipList {
	res := C.tint_wmax_transfn(state._inner, temp.Inner(), interv.Inner())
	return &SkipList{_inner: res}
}


// TintWminTransfn wraps MEOS C function tint_wmin_transfn.
func TintWminTransfn(state *SkipList, temp Temporal, interv timeutil.Timedelta) *SkipList {
	res := C.tint_wmin_transfn(state._inner, temp.Inner(), interv.Inner())
	return &SkipList{_inner: res}
}


// TintWsumTransfn wraps MEOS C function tint_wsum_transfn.
func TintWsumTransfn(state *SkipList, temp Temporal, interv timeutil.Timedelta) *SkipList {
	res := C.tint_wsum_transfn(state._inner, temp.Inner(), interv.Inner())
	return &SkipList{_inner: res}
}


// TnumberExtentTransfn wraps MEOS C function tnumber_extent_transfn.
func TnumberExtentTransfn(box *TBox, temp Temporal) *TBox {
	res := C.tnumber_extent_transfn(box._inner, temp.Inner())
	return &TBox{_inner: res}
}


// TnumberTavgFinalfn wraps MEOS C function tnumber_tavg_finalfn.
func TnumberTavgFinalfn(state *SkipList) Temporal {
	res := C.tnumber_tavg_finalfn(state._inner)
	return CreateTemporal(res)
}


// TnumberTavgTransfn wraps MEOS C function tnumber_tavg_transfn.
func TnumberTavgTransfn(state *SkipList, temp Temporal) *SkipList {
	res := C.tnumber_tavg_transfn(state._inner, temp.Inner())
	return &SkipList{_inner: res}
}


// TnumberWavgTransfn wraps MEOS C function tnumber_wavg_transfn.
func TnumberWavgTransfn(state *SkipList, temp Temporal, interv timeutil.Timedelta) *SkipList {
	res := C.tnumber_wavg_transfn(state._inner, temp.Inner(), interv.Inner())
	return &SkipList{_inner: res}
}


// TstzsetTcountTransfn wraps MEOS C function tstzset_tcount_transfn.
func TstzsetTcountTransfn(state *SkipList, s *Set) *SkipList {
	res := C.tstzset_tcount_transfn(state._inner, s._inner)
	return &SkipList{_inner: res}
}


// TstzspanTcountTransfn wraps MEOS C function tstzspan_tcount_transfn.
func TstzspanTcountTransfn(state *SkipList, s *Span) *SkipList {
	res := C.tstzspan_tcount_transfn(state._inner, s._inner)
	return &SkipList{_inner: res}
}


// TstzspansetTcountTransfn wraps MEOS C function tstzspanset_tcount_transfn.
func TstzspansetTcountTransfn(state *SkipList, ss *SpanSet) *SkipList {
	res := C.tstzspanset_tcount_transfn(state._inner, ss._inner)
	return &SkipList{_inner: res}
}


// TtextTmaxTransfn wraps MEOS C function ttext_tmax_transfn.
func TtextTmaxTransfn(state *SkipList, temp Temporal) *SkipList {
	res := C.ttext_tmax_transfn(state._inner, temp.Inner())
	return &SkipList{_inner: res}
}


// TtextTminTransfn wraps MEOS C function ttext_tmin_transfn.
func TtextTminTransfn(state *SkipList, temp Temporal) *SkipList {
	res := C.ttext_tmin_transfn(state._inner, temp.Inner())
	return &SkipList{_inner: res}
}


// TemporalSimplifyDp wraps MEOS C function temporal_simplify_dp.
func TemporalSimplifyDp(temp Temporal, eps_dist float64, synchronized bool) Temporal {
	res := C.temporal_simplify_dp(temp.Inner(), C.double(eps_dist), C.bool(synchronized))
	return CreateTemporal(res)
}


// TemporalSimplifyMaxDist wraps MEOS C function temporal_simplify_max_dist.
func TemporalSimplifyMaxDist(temp Temporal, eps_dist float64, synchronized bool) Temporal {
	res := C.temporal_simplify_max_dist(temp.Inner(), C.double(eps_dist), C.bool(synchronized))
	return CreateTemporal(res)
}


// TemporalSimplifyMinDist wraps MEOS C function temporal_simplify_min_dist.
func TemporalSimplifyMinDist(temp Temporal, dist float64) Temporal {
	res := C.temporal_simplify_min_dist(temp.Inner(), C.double(dist))
	return CreateTemporal(res)
}


// TemporalSimplifyMinTdelta wraps MEOS C function temporal_simplify_min_tdelta.
func TemporalSimplifyMinTdelta(temp Temporal, mint timeutil.Timedelta) Temporal {
	res := C.temporal_simplify_min_tdelta(temp.Inner(), mint.Inner())
	return CreateTemporal(res)
}


// TemporalTprecision wraps MEOS C function temporal_tprecision.
func TemporalTprecision(temp Temporal, duration timeutil.Timedelta, origin int64) Temporal {
	res := C.temporal_tprecision(temp.Inner(), duration.Inner(), C.TimestampTz(origin))
	return CreateTemporal(res)
}


// TemporalTsample wraps MEOS C function temporal_tsample.
func TemporalTsample(temp Temporal, duration timeutil.Timedelta, origin int64, interp Interpolation) Temporal {
	res := C.temporal_tsample(temp.Inner(), duration.Inner(), C.TimestampTz(origin), C.interpType(interp))
	return CreateTemporal(res)
}


// TemporalDyntimewarpDistance wraps MEOS C function temporal_dyntimewarp_distance.
func TemporalDyntimewarpDistance(temp1 Temporal, temp2 Temporal) float64 {
	res := C.temporal_dyntimewarp_distance(temp1.Inner(), temp2.Inner())
	return float64(res)
}


// TemporalDyntimewarpPath wraps MEOS C function temporal_dyntimewarp_path.
func TemporalDyntimewarpPath(temp1 Temporal, temp2 Temporal) (*Match, int) {
	var _out_count C.int
	res := C.temporal_dyntimewarp_path(temp1.Inner(), temp2.Inner(), &_out_count)
	return &Match{_inner: res}, int(_out_count)
}


// TemporalFrechetDistance wraps MEOS C function temporal_frechet_distance.
func TemporalFrechetDistance(temp1 Temporal, temp2 Temporal) float64 {
	res := C.temporal_frechet_distance(temp1.Inner(), temp2.Inner())
	return float64(res)
}


// TemporalFrechetPath wraps MEOS C function temporal_frechet_path.
func TemporalFrechetPath(temp1 Temporal, temp2 Temporal) (*Match, int) {
	var _out_count C.int
	res := C.temporal_frechet_path(temp1.Inner(), temp2.Inner(), &_out_count)
	return &Match{_inner: res}, int(_out_count)
}


// TemporalHausdorffDistance wraps MEOS C function temporal_hausdorff_distance.
func TemporalHausdorffDistance(temp1 Temporal, temp2 Temporal) float64 {
	res := C.temporal_hausdorff_distance(temp1.Inner(), temp2.Inner())
	return float64(res)
}


// TemporalTimeBins wraps MEOS C function temporal_time_bins.
func TemporalTimeBins(temp Temporal, duration timeutil.Timedelta, origin int64) (*Span, int) {
	var _out_count C.int
	res := C.temporal_time_bins(temp.Inner(), duration.Inner(), C.TimestampTz(origin), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// TODO temporal_time_split: unsupported param TimestampTz **
// func TemporalTimeSplit(...) { /* not yet handled by codegen */ }


// TfloatTimeBoxes wraps MEOS C function tfloat_time_boxes.
func TfloatTimeBoxes(temp Temporal, duration timeutil.Timedelta, torigin int64) (*TBox, int) {
	var _out_count C.int
	res := C.tfloat_time_boxes(temp.Inner(), duration.Inner(), C.TimestampTz(torigin), &_out_count)
	return &TBox{_inner: res}, int(_out_count)
}


// TfloatValueBins wraps MEOS C function tfloat_value_bins.
func TfloatValueBins(temp Temporal, vsize float64, vorigin float64) (*Span, int) {
	var _out_count C.int
	res := C.tfloat_value_bins(temp.Inner(), C.double(vsize), C.double(vorigin), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// TfloatValueBoxes wraps MEOS C function tfloat_value_boxes.
func TfloatValueBoxes(temp Temporal, vsize float64, vorigin float64) (*TBox, int) {
	var _out_count C.int
	res := C.tfloat_value_boxes(temp.Inner(), C.double(vsize), C.double(vorigin), &_out_count)
	return &TBox{_inner: res}, int(_out_count)
}


// TODO tfloat_value_split: unsupported param double **
// func TfloatValueSplit(...) { /* not yet handled by codegen */ }


// TfloatValueTimeBoxes wraps MEOS C function tfloat_value_time_boxes.
func TfloatValueTimeBoxes(temp Temporal, vsize float64, duration timeutil.Timedelta, vorigin float64, torigin int64) (*TBox, int) {
	var _out_count C.int
	res := C.tfloat_value_time_boxes(temp.Inner(), C.double(vsize), duration.Inner(), C.double(vorigin), C.TimestampTz(torigin), &_out_count)
	return &TBox{_inner: res}, int(_out_count)
}


// TODO tfloat_value_time_split: unsupported param double **
// func TfloatValueTimeSplit(...) { /* not yet handled by codegen */ }


// TfloatboxTimeTiles wraps MEOS C function tfloatbox_time_tiles.
func TfloatboxTimeTiles(box *TBox, duration timeutil.Timedelta, torigin int64) (*TBox, int) {
	var _out_count C.int
	res := C.tfloatbox_time_tiles(box._inner, duration.Inner(), C.TimestampTz(torigin), &_out_count)
	return &TBox{_inner: res}, int(_out_count)
}


// TfloatboxValueTiles wraps MEOS C function tfloatbox_value_tiles.
func TfloatboxValueTiles(box *TBox, vsize float64, vorigin float64) (*TBox, int) {
	var _out_count C.int
	res := C.tfloatbox_value_tiles(box._inner, C.double(vsize), C.double(vorigin), &_out_count)
	return &TBox{_inner: res}, int(_out_count)
}


// TfloatboxValueTimeTiles wraps MEOS C function tfloatbox_value_time_tiles.
func TfloatboxValueTimeTiles(box *TBox, vsize float64, duration timeutil.Timedelta, vorigin float64, torigin int64) (*TBox, int) {
	var _out_count C.int
	res := C.tfloatbox_value_time_tiles(box._inner, C.double(vsize), duration.Inner(), C.double(vorigin), C.TimestampTz(torigin), &_out_count)
	return &TBox{_inner: res}, int(_out_count)
}


// TintTimeBoxes wraps MEOS C function tint_time_boxes.
func TintTimeBoxes(temp Temporal, duration timeutil.Timedelta, torigin int64) (*TBox, int) {
	var _out_count C.int
	res := C.tint_time_boxes(temp.Inner(), duration.Inner(), C.TimestampTz(torigin), &_out_count)
	return &TBox{_inner: res}, int(_out_count)
}


// TintValueBins wraps MEOS C function tint_value_bins.
func TintValueBins(temp Temporal, vsize int, vorigin int) (*Span, int) {
	var _out_count C.int
	res := C.tint_value_bins(temp.Inner(), C.int(vsize), C.int(vorigin), &_out_count)
	return &Span{_inner: res}, int(_out_count)
}


// TintValueBoxes wraps MEOS C function tint_value_boxes.
func TintValueBoxes(temp Temporal, vsize int, vorigin int) (*TBox, int) {
	var _out_count C.int
	res := C.tint_value_boxes(temp.Inner(), C.int(vsize), C.int(vorigin), &_out_count)
	return &TBox{_inner: res}, int(_out_count)
}


// TODO tint_value_split: unsupported param int **
// func TintValueSplit(...) { /* not yet handled by codegen */ }


// TintValueTimeBoxes wraps MEOS C function tint_value_time_boxes.
func TintValueTimeBoxes(temp Temporal, vsize int, duration timeutil.Timedelta, vorigin int, torigin int64) (*TBox, int) {
	var _out_count C.int
	res := C.tint_value_time_boxes(temp.Inner(), C.int(vsize), duration.Inner(), C.int(vorigin), C.TimestampTz(torigin), &_out_count)
	return &TBox{_inner: res}, int(_out_count)
}


// TODO tint_value_time_split: unsupported param int **
// func TintValueTimeSplit(...) { /* not yet handled by codegen */ }


// TintboxTimeTiles wraps MEOS C function tintbox_time_tiles.
func TintboxTimeTiles(box *TBox, duration timeutil.Timedelta, torigin int64) (*TBox, int) {
	var _out_count C.int
	res := C.tintbox_time_tiles(box._inner, duration.Inner(), C.TimestampTz(torigin), &_out_count)
	return &TBox{_inner: res}, int(_out_count)
}


// TintboxValueTiles wraps MEOS C function tintbox_value_tiles.
func TintboxValueTiles(box *TBox, xsize int, xorigin int) (*TBox, int) {
	var _out_count C.int
	res := C.tintbox_value_tiles(box._inner, C.int(xsize), C.int(xorigin), &_out_count)
	return &TBox{_inner: res}, int(_out_count)
}


// TintboxValueTimeTiles wraps MEOS C function tintbox_value_time_tiles.
func TintboxValueTimeTiles(box *TBox, xsize int, duration timeutil.Timedelta, xorigin int, torigin int64) (*TBox, int) {
	var _out_count C.int
	res := C.tintbox_value_time_tiles(box._inner, C.int(xsize), duration.Inner(), C.int(xorigin), C.TimestampTz(torigin), &_out_count)
	return &TBox{_inner: res}, int(_out_count)
}

