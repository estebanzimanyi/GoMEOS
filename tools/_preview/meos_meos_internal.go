package generated

// #include <stddef.h>
import "C"
import (
	"unsafe"

	"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

// GslGetGenerationRng wraps MEOS C function gsl_get_generation_rng.
func GslGetGenerationRng() *GslRng {
	res := C.gsl_get_generation_rng()
	return &GslRng{_inner: res}
}


// GslGetAggregationRng wraps MEOS C function gsl_get_aggregation_rng.
func GslGetAggregationRng() *GslRng {
	res := C.gsl_get_aggregation_rng()
	return &GslRng{_inner: res}
}


// FloatspanRoundSet wraps MEOS C function floatspan_round_set.
func FloatspanRoundSet(s *Span, maxdd int) *Span {
	var _out_result C.Span
	C.floatspan_round_set(s._inner, C.int(maxdd), &_out_result)
	return &Span{_inner: &_out_result}
}


// TODO set_in: unsupported param MeosType
// func SetIn(...) { /* not yet handled by codegen */ }


// SetOut wraps MEOS C function set_out.
func SetOut(s *Set, maxdd int) string {
	res := C.set_out(s._inner, C.int(maxdd))
	return C.GoString(res)
}


// TODO span_in: unsupported param MeosType
// func SpanIn(...) { /* not yet handled by codegen */ }


// SpanOut wraps MEOS C function span_out.
func SpanOut(s *Span, maxdd int) string {
	res := C.span_out(s._inner, C.int(maxdd))
	return C.GoString(res)
}


// TODO spanset_in: unsupported param MeosType
// func SpansetIn(...) { /* not yet handled by codegen */ }


// SpansetOut wraps MEOS C function spanset_out.
func SpansetOut(ss *SpanSet, maxdd int) string {
	res := C.spanset_out(ss._inner, C.int(maxdd))
	return C.GoString(res)
}


// SpansetMakeExp wraps MEOS C function spanset_make_exp.
func SpansetMakeExp(spans *Span, count int, maxcount int, normalize bool, order bool) *SpanSet {
	res := C.spanset_make_exp(spans._inner, C.int(count), C.int(maxcount), C.bool(normalize), C.bool(order))
	return &SpanSet{_inner: res}
}


// SpansetMakeFree wraps MEOS C function spanset_make_free.
func SpansetMakeFree(spans *Span, count int, normalize bool, order bool) *SpanSet {
	res := C.spanset_make_free(spans._inner, C.int(count), C.bool(normalize), C.bool(order))
	return &SpanSet{_inner: res}
}


// SetSpan wraps MEOS C function set_span.
func SetSpan(s *Set) *Span {
	res := C.set_span(s._inner)
	return &Span{_inner: res}
}


// SetSpanset wraps MEOS C function set_spanset.
func SetSpanset(s *Set) *SpanSet {
	res := C.set_spanset(s._inner)
	return &SpanSet{_inner: res}
}


// NumspanWidth wraps MEOS C function numspan_width.
func NumspanWidth(s *Span) int {
	res := C.numspan_width(s._inner)
	return int(res)
}


// NumspansetWidth wraps MEOS C function numspanset_width.
func NumspansetWidth(ss *SpanSet, boundspan bool) int {
	res := C.numspanset_width(ss._inner, C.bool(boundspan))
	return int(res)
}


// SetEndValue wraps MEOS C function set_end_value.
func SetEndValue(s *Set) int {
	res := C.set_end_value(s._inner)
	return int(res)
}


// SetMemSize wraps MEOS C function set_mem_size.
func SetMemSize(s *Set) int {
	res := C.set_mem_size(s._inner)
	return int(res)
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


// SetStartValue wraps MEOS C function set_start_value.
func SetStartValue(s *Set) int {
	res := C.set_start_value(s._inner)
	return int(res)
}


// SpansetLower wraps MEOS C function spanset_lower.
func SpansetLower(ss *SpanSet) int {
	res := C.spanset_lower(ss._inner)
	return int(res)
}


// SpansetMemSize wraps MEOS C function spanset_mem_size.
func SpansetMemSize(ss *SpanSet) int {
	res := C.spanset_mem_size(ss._inner)
	return int(res)
}


// SpansetSps wraps MEOS C function spanset_sps.
func SpansetSps(ss *SpanSet) []*Span {
	res := C.spanset_sps(ss._inner)
	_n := int(C.spanset_num_spans(ss.Inner()))
	_slice := unsafe.Slice((**C.Span)(unsafe.Pointer(res)), _n)
	_out := make([]*Span, _n)
	for _i, _e := range _slice {
		_out[_i] = &Span{_inner: _e}
	}
	return _out
}


// SpansetUpper wraps MEOS C function spanset_upper.
func SpansetUpper(ss *SpanSet) int {
	res := C.spanset_upper(ss._inner)
	return int(res)
}


// DatespanSetTstzspan wraps MEOS C function datespan_set_tstzspan.
func DatespanSetTstzspan(s1 *Span, s2 *Span) {
	C.datespan_set_tstzspan(s1._inner, s2._inner)
}


// FloatspanSetIntspan wraps MEOS C function floatspan_set_intspan.
func FloatspanSetIntspan(s1 *Span, s2 *Span) {
	C.floatspan_set_intspan(s1._inner, s2._inner)
}


// IntspanSetFloatspan wraps MEOS C function intspan_set_floatspan.
func IntspanSetFloatspan(s1 *Span, s2 *Span) {
	C.intspan_set_floatspan(s1._inner, s2._inner)
}


// SetCompact wraps MEOS C function set_compact.
func SetCompact(s *Set) *Set {
	res := C.set_compact(s._inner)
	return &Set{_inner: res}
}


// SpanExpand wraps MEOS C function span_expand.
func SpanExpand(s1 *Span, s2 *Span) {
	C.span_expand(s1._inner, s2._inner)
}


// SpansetCompact wraps MEOS C function spanset_compact.
func SpansetCompact(ss *SpanSet) *SpanSet {
	res := C.spanset_compact(ss._inner)
	return &SpanSet{_inner: res}
}


// TODO textcat_textset_text_common: unsupported param const int *
// func TextcatTextsetTextCommon(...) { /* not yet handled by codegen */ }


// TstzspanSetDatespan wraps MEOS C function tstzspan_set_datespan.
func TstzspanSetDatespan(s1 *Span, s2 *Span) {
	C.tstzspan_set_datespan(s1._inner, s2._inner)
}


// OvadjSpanSpan wraps MEOS C function ovadj_span_span.
func OvadjSpanSpan(s1 *Span, s2 *Span) bool {
	res := C.ovadj_span_span(s1._inner, s2._inner)
	return bool(res)
}


// LfnadjSpanSpan wraps MEOS C function lfnadj_span_span.
func LfnadjSpanSpan(s1 *Span, s2 *Span) bool {
	res := C.lfnadj_span_span(s1._inner, s2._inner)
	return bool(res)
}


// TODO bbox_type: unsupported param MeosType
// func BboxType(...) { /* not yet handled by codegen */ }


// TODO bbox_get_size: unsupported param MeosType
// func BboxGetSize(...) { /* not yet handled by codegen */ }


// TODO bbox_max_dims: unsupported param MeosType
// func BboxMaxDims(...) { /* not yet handled by codegen */ }


// TODO temporal_bbox_eq: unsupported param MeosType
// func TemporalBboxEq(...) { /* not yet handled by codegen */ }


// TODO temporal_bbox_cmp: unsupported param MeosType
// func TemporalBboxCmp(...) { /* not yet handled by codegen */ }


// BboxUnionSpanSpan wraps MEOS C function bbox_union_span_span.
func BboxUnionSpanSpan(s1 *Span, s2 *Span) *Span {
	var _out_result C.Span
	C.bbox_union_span_span(s1._inner, s2._inner, &_out_result)
	return &Span{_inner: &_out_result}
}


// InterSpanSpan wraps MEOS C function inter_span_span.
func InterSpanSpan(s1 *Span, s2 *Span) (bool, *Span) {
	var _out_result C.Span
	res := C.inter_span_span(s1._inner, s2._inner, &_out_result)
	return bool(res), &Span{_inner: &_out_result}
}


// MiSpanSpan wraps MEOS C function mi_span_span.
func MiSpanSpan(s1 *Span, s2 *Span) (int, *Span) {
	var _out_result C.Span
	res := C.mi_span_span(s1._inner, s2._inner, &_out_result)
	return int(res), &Span{_inner: &_out_result}
}


// SuperUnionSpanSpan wraps MEOS C function super_union_span_span.
func SuperUnionSpanSpan(s1 *Span, s2 *Span) *Span {
	res := C.super_union_span_span(s1._inner, s2._inner)
	return &Span{_inner: res}
}


// DistanceSetSet wraps MEOS C function distance_set_set.
func DistanceSetSet(s1 *Set, s2 *Set) int {
	res := C.distance_set_set(s1._inner, s2._inner)
	return int(res)
}


// DistanceSpanSpan wraps MEOS C function distance_span_span.
func DistanceSpanSpan(s1 *Span, s2 *Span) int {
	res := C.distance_span_span(s1._inner, s2._inner)
	return int(res)
}


// DistanceSpansetSpan wraps MEOS C function distance_spanset_span.
func DistanceSpansetSpan(ss *SpanSet, s *Span) int {
	res := C.distance_spanset_span(ss._inner, s._inner)
	return int(res)
}


// DistanceSpansetSpanset wraps MEOS C function distance_spanset_spanset.
func DistanceSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) int {
	res := C.distance_spanset_spanset(ss1._inner, ss2._inner)
	return int(res)
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
func TimestamptzSetTBOX(t int, box *TBox) {
	C.timestamptz_set_tbox(C.int(t), box._inner)
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
	res := C.inter_tbox_tbox(box1._inner, box2._inner, &_out_result)
	return bool(res), &TBox{_inner: &_out_result}
}


// TboolinstIn wraps MEOS C function tboolinst_in.
func TboolinstIn(str string) TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tboolinst_in(_c_str)
	return TInstant{_inner: res}
}


// TboolseqIn wraps MEOS C function tboolseq_in.
func TboolseqIn(str string, interp Interpolation) TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tboolseq_in(_c_str, C.interpType(interp))
	return TSequence{_inner: res}
}


// TboolseqsetIn wraps MEOS C function tboolseqset_in.
func TboolseqsetIn(str string) TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tboolseqset_in(_c_str)
	return TSequenceSet{_inner: res}
}


// TODO temporal_in: unsupported param MeosType
// func TemporalIn(...) { /* not yet handled by codegen */ }


// TemporalOut wraps MEOS C function temporal_out.
func TemporalOut(temp Temporal, maxdd int) string {
	res := C.temporal_out(temp.Inner(), C.int(maxdd))
	return C.GoString(res)
}


// TemparrOut wraps MEOS C function temparr_out.
func TemparrOut(temparr []Temporal, maxdd int) []string {
	_c_temparr := make([]*C.Temporal, len(temparr))
	for _i, _v := range temparr { _c_temparr[_i] = _v._inner }
	res := C.temparr_out((**C.Temporal)(unsafe.Pointer(&_c_temparr[0])), C.int(len(temparr)), C.int(maxdd))
	_n := len(temparr)
	_slice := unsafe.Slice((**C.char)(unsafe.Pointer(res)), _n)
	_out := make([]string, _n)
	for _i, _e := range _slice {
		_out[_i] = C.GoString(_e)
	}
	return _out
}


// TfloatinstIn wraps MEOS C function tfloatinst_in.
func TfloatinstIn(str string) TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tfloatinst_in(_c_str)
	return TInstant{_inner: res}
}


// TfloatseqIn wraps MEOS C function tfloatseq_in.
func TfloatseqIn(str string, interp Interpolation) TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tfloatseq_in(_c_str, C.interpType(interp))
	return TSequence{_inner: res}
}


// TfloatseqsetIn wraps MEOS C function tfloatseqset_in.
func TfloatseqsetIn(str string) TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tfloatseqset_in(_c_str)
	return TSequenceSet{_inner: res}
}


// TODO tinstant_in: unsupported param MeosType
// func TinstantIn(...) { /* not yet handled by codegen */ }


// TinstantOut wraps MEOS C function tinstant_out.
func TinstantOut(inst TInstant, maxdd int) string {
	res := C.tinstant_out(inst.Inner(), C.int(maxdd))
	return C.GoString(res)
}


// TintinstIn wraps MEOS C function tintinst_in.
func TintinstIn(str string) TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tintinst_in(_c_str)
	return TInstant{_inner: res}
}


// TintseqIn wraps MEOS C function tintseq_in.
func TintseqIn(str string, interp Interpolation) TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tintseq_in(_c_str, C.interpType(interp))
	return TSequence{_inner: res}
}


// TintseqsetIn wraps MEOS C function tintseqset_in.
func TintseqsetIn(str string) TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tintseqset_in(_c_str)
	return TSequenceSet{_inner: res}
}


// TODO tsequence_in: unsupported param MeosType
// func TsequenceIn(...) { /* not yet handled by codegen */ }


// TsequenceOut wraps MEOS C function tsequence_out.
func TsequenceOut(seq TSequence, maxdd int) string {
	res := C.tsequence_out(seq.Inner(), C.int(maxdd))
	return C.GoString(res)
}


// TODO tsequenceset_in: unsupported param MeosType
// func TsequencesetIn(...) { /* not yet handled by codegen */ }


// TsequencesetOut wraps MEOS C function tsequenceset_out.
func TsequencesetOut(ss TSequenceSet, maxdd int) string {
	res := C.tsequenceset_out(ss.Inner(), C.int(maxdd))
	return C.GoString(res)
}


// TtextinstIn wraps MEOS C function ttextinst_in.
func TtextinstIn(str string) TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.ttextinst_in(_c_str)
	return TInstant{_inner: res}
}


// TtextseqIn wraps MEOS C function ttextseq_in.
func TtextseqIn(str string, interp Interpolation) TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.ttextseq_in(_c_str, C.interpType(interp))
	return TSequence{_inner: res}
}


// TtextseqsetIn wraps MEOS C function ttextseqset_in.
func TtextseqsetIn(str string) TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.ttextseqset_in(_c_str)
	return TSequenceSet{_inner: res}
}


// TODO temporal_from_mfjson: unsupported param MeosType
// func TemporalFromMFJSON(...) { /* not yet handled by codegen */ }


// TinstantCopy wraps MEOS C function tinstant_copy.
func TinstantCopy(inst TInstant) TInstant {
	res := C.tinstant_copy(inst.Inner())
	return TInstant{_inner: res}
}


// TsequenceCopy wraps MEOS C function tsequence_copy.
func TsequenceCopy(seq TSequence) TSequence {
	res := C.tsequence_copy(seq.Inner())
	return TSequence{_inner: res}
}


// TsequenceMakeExp wraps MEOS C function tsequence_make_exp.
func TsequenceMakeExp(instants []TInstant, maxcount int, lower_inc bool, upper_inc bool, interp Interpolation, normalize bool) TSequence {
	_c_instants := make([]*C.TInstant, len(instants))
	for _i, _v := range instants { _c_instants[_i] = _v._inner }
	res := C.tsequence_make_exp((**C.TInstant)(unsafe.Pointer(&_c_instants[0])), C.int(len(instants)), C.int(maxcount), C.bool(lower_inc), C.bool(upper_inc), C.interpType(interp), C.bool(normalize))
	return TSequence{_inner: res}
}


// TsequenceMakeFree wraps MEOS C function tsequence_make_free.
func TsequenceMakeFree(instants []TInstant, lower_inc bool, upper_inc bool, interp Interpolation, normalize bool) TSequence {
	_c_instants := make([]*C.TInstant, len(instants))
	for _i, _v := range instants { _c_instants[_i] = _v._inner }
	res := C.tsequence_make_free((**C.TInstant)(unsafe.Pointer(&_c_instants[0])), C.int(len(instants)), C.bool(lower_inc), C.bool(upper_inc), C.interpType(interp), C.bool(normalize))
	return TSequence{_inner: res}
}


// TsequencesetCopy wraps MEOS C function tsequenceset_copy.
func TsequencesetCopy(ss TSequenceSet) TSequenceSet {
	res := C.tsequenceset_copy(ss.Inner())
	return TSequenceSet{_inner: res}
}


// TseqsetarrToTseqset wraps MEOS C function tseqsetarr_to_tseqset.
func TseqsetarrToTseqset(seqsets []TSequenceSet, totalseqs int) TSequenceSet {
	_c_seqsets := make([]*C.TSequenceSet, len(seqsets))
	for _i, _v := range seqsets { _c_seqsets[_i] = _v._inner }
	res := C.tseqsetarr_to_tseqset((**C.TSequenceSet)(unsafe.Pointer(&_c_seqsets[0])), C.int(len(seqsets)), C.int(totalseqs))
	return TSequenceSet{_inner: res}
}


// TsequencesetMakeExp wraps MEOS C function tsequenceset_make_exp.
func TsequencesetMakeExp(sequences []TSequence, maxcount int, normalize bool) TSequenceSet {
	_c_sequences := make([]*C.TSequence, len(sequences))
	for _i, _v := range sequences { _c_sequences[_i] = _v._inner }
	res := C.tsequenceset_make_exp((**C.TSequence)(unsafe.Pointer(&_c_sequences[0])), C.int(len(sequences)), C.int(maxcount), C.bool(normalize))
	return TSequenceSet{_inner: res}
}


// TsequencesetMakeFree wraps MEOS C function tsequenceset_make_free.
func TsequencesetMakeFree(sequences []TSequence, normalize bool) TSequenceSet {
	_c_sequences := make([]*C.TSequence, len(sequences))
	for _i, _v := range sequences { _c_sequences[_i] = _v._inner }
	res := C.tsequenceset_make_free((**C.TSequence)(unsafe.Pointer(&_c_sequences[0])), C.int(len(sequences)), C.bool(normalize))
	return TSequenceSet{_inner: res}
}


// TemporalSetTstzspan wraps MEOS C function temporal_set_tstzspan.
func TemporalSetTstzspan(temp Temporal, s *Span) {
	C.temporal_set_tstzspan(temp.Inner(), s._inner)
}


// TinstantSetTstzspan wraps MEOS C function tinstant_set_tstzspan.
func TinstantSetTstzspan(inst TInstant, s *Span) {
	C.tinstant_set_tstzspan(inst.Inner(), s._inner)
}


// TnumberSetTBOX wraps MEOS C function tnumber_set_tbox.
func TnumberSetTBOX(temp Temporal, box *TBox) {
	C.tnumber_set_tbox(temp.Inner(), box._inner)
}


// TnumberinstSetTBOX wraps MEOS C function tnumberinst_set_tbox.
func TnumberinstSetTBOX(inst TInstant, box *TBox) {
	C.tnumberinst_set_tbox(inst.Inner(), box._inner)
}


// TnumberseqSetTBOX wraps MEOS C function tnumberseq_set_tbox.
func TnumberseqSetTBOX(seq TSequence, box *TBox) {
	C.tnumberseq_set_tbox(seq.Inner(), box._inner)
}


// TnumberseqsetSetTBOX wraps MEOS C function tnumberseqset_set_tbox.
func TnumberseqsetSetTBOX(ss TSequenceSet, box *TBox) {
	C.tnumberseqset_set_tbox(ss.Inner(), box._inner)
}


// TsequenceSetTstzspan wraps MEOS C function tsequence_set_tstzspan.
func TsequenceSetTstzspan(seq TSequence, s *Span) {
	C.tsequence_set_tstzspan(seq.Inner(), s._inner)
}


// TsequencesetSetTstzspan wraps MEOS C function tsequenceset_set_tstzspan.
func TsequencesetSetTstzspan(ss TSequenceSet, s *Span) {
	C.tsequenceset_set_tstzspan(ss.Inner(), s._inner)
}


// TemporalEndInst wraps MEOS C function temporal_end_inst.
func TemporalEndInst(temp Temporal) TInstant {
	res := C.temporal_end_inst(temp.Inner())
	return TInstant{_inner: res}
}


// TemporalEndValue wraps MEOS C function temporal_end_value.
func TemporalEndValue(temp Temporal) int {
	res := C.temporal_end_value(temp.Inner())
	return int(res)
}


// TemporalInstN wraps MEOS C function temporal_inst_n.
func TemporalInstN(temp Temporal, n int) TInstant {
	res := C.temporal_inst_n(temp.Inner(), C.int(n))
	return TInstant{_inner: res}
}


// TemporalInstsP wraps MEOS C function temporal_insts_p.
func TemporalInstsP(temp Temporal) []TInstant {
	var _out_count C.int
	res := C.temporal_insts_p(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.TInstant)(unsafe.Pointer(res)), _n)
	_out := make([]TInstant, _n)
	for _i, _e := range _slice {
		_out[_i] = TInstant{_inner: _e}
	}
	return _out
}


// TemporalMaxInstP wraps MEOS C function temporal_max_inst_p.
func TemporalMaxInstP(temp Temporal) TInstant {
	res := C.temporal_max_inst_p(temp.Inner())
	return TInstant{_inner: res}
}


// TemporalMaxValue wraps MEOS C function temporal_max_value.
func TemporalMaxValue(temp Temporal) int {
	res := C.temporal_max_value(temp.Inner())
	return int(res)
}


// TemporalMemSize wraps MEOS C function temporal_mem_size.
func TemporalMemSize(temp Temporal) uint {
	res := C.temporal_mem_size(temp.Inner())
	return uint(res)
}


// TemporalMinInstP wraps MEOS C function temporal_min_inst_p.
func TemporalMinInstP(temp Temporal) TInstant {
	res := C.temporal_min_inst_p(temp.Inner())
	return TInstant{_inner: res}
}


// TemporalMinValue wraps MEOS C function temporal_min_value.
func TemporalMinValue(temp Temporal) int {
	res := C.temporal_min_value(temp.Inner())
	return int(res)
}


// TemporalSequencesP wraps MEOS C function temporal_sequences_p.
func TemporalSequencesP(temp Temporal) []TSequence {
	var _out_count C.int
	res := C.temporal_sequences_p(temp.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.TSequence)(unsafe.Pointer(res)), _n)
	_out := make([]TSequence, _n)
	for _i, _e := range _slice {
		_out[_i] = TSequence{_inner: _e}
	}
	return _out
}


// TemporalSetBbox wraps MEOS C function temporal_set_bbox.
func TemporalSetBbox(temp Temporal, box unsafe.Pointer) {
	C.temporal_set_bbox(temp.Inner(), unsafe.Pointer(box))
}


// TemporalStartInst wraps MEOS C function temporal_start_inst.
func TemporalStartInst(temp Temporal) TInstant {
	res := C.temporal_start_inst(temp.Inner())
	return TInstant{_inner: res}
}


// TemporalStartValue wraps MEOS C function temporal_start_value.
func TemporalStartValue(temp Temporal) int {
	res := C.temporal_start_value(temp.Inner())
	return int(res)
}


// TinstantHash wraps MEOS C function tinstant_hash.
func TinstantHash(inst TInstant) int {
	res := C.tinstant_hash(inst.Inner())
	return int(res)
}


// TinstantInsts wraps MEOS C function tinstant_insts.
func TinstantInsts(inst TInstant) []TInstant {
	var _out_count C.int
	res := C.tinstant_insts(inst.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.TInstant)(unsafe.Pointer(res)), _n)
	_out := make([]TInstant, _n)
	for _i, _e := range _slice {
		_out[_i] = TInstant{_inner: _e}
	}
	return _out
}


// TinstantSetBbox wraps MEOS C function tinstant_set_bbox.
func TinstantSetBbox(inst TInstant, box unsafe.Pointer) {
	C.tinstant_set_bbox(inst.Inner(), unsafe.Pointer(box))
}


// TinstantTime wraps MEOS C function tinstant_time.
func TinstantTime(inst TInstant) *SpanSet {
	res := C.tinstant_time(inst.Inner())
	return &SpanSet{_inner: res}
}


// TinstantTimestamps wraps MEOS C function tinstant_timestamps.
func TinstantTimestamps(inst TInstant) []int {
	var _out_count C.int
	res := C.tinstant_timestamps(inst.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((*C.int)(unsafe.Pointer(res)), _n)
	_out := make([]int, _n)
	for _i, _e := range _slice {
		_out[_i] = int(_e)
	}
	return _out
}


// TinstantValueP wraps MEOS C function tinstant_value_p.
func TinstantValueP(inst TInstant) int {
	res := C.tinstant_value_p(inst.Inner())
	return int(res)
}


// TinstantValue wraps MEOS C function tinstant_value.
func TinstantValue(inst TInstant) int {
	res := C.tinstant_value(inst.Inner())
	return int(res)
}


// TnumberSetSpan wraps MEOS C function tnumber_set_span.
func TnumberSetSpan(temp Temporal, span *Span) {
	C.tnumber_set_span(temp.Inner(), span._inner)
}


// TnumberinstValuespans wraps MEOS C function tnumberinst_valuespans.
func TnumberinstValuespans(inst TInstant) *SpanSet {
	res := C.tnumberinst_valuespans(inst.Inner())
	return &SpanSet{_inner: res}
}


// TnumberseqAvgVal wraps MEOS C function tnumberseq_avg_val.
func TnumberseqAvgVal(seq TSequence) float64 {
	res := C.tnumberseq_avg_val(seq.Inner())
	return float64(res)
}


// TnumberseqValuespans wraps MEOS C function tnumberseq_valuespans.
func TnumberseqValuespans(seq TSequence) *SpanSet {
	res := C.tnumberseq_valuespans(seq.Inner())
	return &SpanSet{_inner: res}
}


// TnumberseqsetAvgVal wraps MEOS C function tnumberseqset_avg_val.
func TnumberseqsetAvgVal(ss TSequenceSet) float64 {
	res := C.tnumberseqset_avg_val(ss.Inner())
	return float64(res)
}


// TnumberseqsetValuespans wraps MEOS C function tnumberseqset_valuespans.
func TnumberseqsetValuespans(ss TSequenceSet) *SpanSet {
	res := C.tnumberseqset_valuespans(ss.Inner())
	return &SpanSet{_inner: res}
}


// TODO tsequence_duration: unsupported return type int *
// func TsequenceDuration(...) { /* not yet handled by codegen */ }


// TsequenceEndTimestamptz wraps MEOS C function tsequence_end_timestamptz.
func TsequenceEndTimestamptz(seq TSequence) int {
	res := C.tsequence_end_timestamptz(seq.Inner())
	return int(res)
}


// TsequenceHash wraps MEOS C function tsequence_hash.
func TsequenceHash(seq TSequence) int {
	res := C.tsequence_hash(seq.Inner())
	return int(res)
}


// TsequenceInstsP wraps MEOS C function tsequence_insts_p.
func TsequenceInstsP(seq TSequence) []TInstant {
	res := C.tsequence_insts_p(seq.Inner())
	_n := int(C.temporal_num_instants((*C.Temporal)(unsafe.Pointer(seq.Inner()))))
	_slice := unsafe.Slice((**C.TInstant)(unsafe.Pointer(res)), _n)
	_out := make([]TInstant, _n)
	for _i, _e := range _slice {
		_out[_i] = TInstant{_inner: _e}
	}
	return _out
}


// TsequenceMaxInstP wraps MEOS C function tsequence_max_inst_p.
func TsequenceMaxInstP(seq TSequence) TInstant {
	res := C.tsequence_max_inst_p(seq.Inner())
	return TInstant{_inner: res}
}


// TsequenceMaxVal wraps MEOS C function tsequence_max_val.
func TsequenceMaxVal(seq TSequence) int {
	res := C.tsequence_max_val(seq.Inner())
	return int(res)
}


// TsequenceMinInstP wraps MEOS C function tsequence_min_inst_p.
func TsequenceMinInstP(seq TSequence) TInstant {
	res := C.tsequence_min_inst_p(seq.Inner())
	return TInstant{_inner: res}
}


// TsequenceMinVal wraps MEOS C function tsequence_min_val.
func TsequenceMinVal(seq TSequence) int {
	res := C.tsequence_min_val(seq.Inner())
	return int(res)
}


// TsequenceSegments wraps MEOS C function tsequence_segments.
func TsequenceSegments(seq TSequence) []TSequence {
	var _out_count C.int
	res := C.tsequence_segments(seq.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.TSequence)(unsafe.Pointer(res)), _n)
	_out := make([]TSequence, _n)
	for _i, _e := range _slice {
		_out[_i] = TSequence{_inner: _e}
	}
	return _out
}


// TsequenceSeqs wraps MEOS C function tsequence_seqs.
func TsequenceSeqs(seq TSequence) []TSequence {
	var _out_count C.int
	res := C.tsequence_seqs(seq.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.TSequence)(unsafe.Pointer(res)), _n)
	_out := make([]TSequence, _n)
	for _i, _e := range _slice {
		_out[_i] = TSequence{_inner: _e}
	}
	return _out
}


// TsequenceStartTimestamptz wraps MEOS C function tsequence_start_timestamptz.
func TsequenceStartTimestamptz(seq TSequence) int {
	res := C.tsequence_start_timestamptz(seq.Inner())
	return int(res)
}


// TsequenceTime wraps MEOS C function tsequence_time.
func TsequenceTime(seq TSequence) *SpanSet {
	res := C.tsequence_time(seq.Inner())
	return &SpanSet{_inner: res}
}


// TsequenceTimestamps wraps MEOS C function tsequence_timestamps.
func TsequenceTimestamps(seq TSequence) []int {
	var _out_count C.int
	res := C.tsequence_timestamps(seq.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((*C.int)(unsafe.Pointer(res)), _n)
	_out := make([]int, _n)
	for _i, _e := range _slice {
		_out[_i] = int(_e)
	}
	return _out
}


// TODO tsequenceset_duration: unsupported return type int *
// func TsequencesetDuration(...) { /* not yet handled by codegen */ }


// TsequencesetEndTimestamptz wraps MEOS C function tsequenceset_end_timestamptz.
func TsequencesetEndTimestamptz(ss TSequenceSet) int {
	res := C.tsequenceset_end_timestamptz(ss.Inner())
	return int(res)
}


// TsequencesetHash wraps MEOS C function tsequenceset_hash.
func TsequencesetHash(ss TSequenceSet) int {
	res := C.tsequenceset_hash(ss.Inner())
	return int(res)
}


// TsequencesetInstN wraps MEOS C function tsequenceset_inst_n.
func TsequencesetInstN(ss TSequenceSet, n int) TInstant {
	res := C.tsequenceset_inst_n(ss.Inner(), C.int(n))
	return TInstant{_inner: res}
}


// TsequencesetInstsP wraps MEOS C function tsequenceset_insts_p.
func TsequencesetInstsP(ss TSequenceSet) []TInstant {
	res := C.tsequenceset_insts_p(ss.Inner())
	_n := int(C.tsequenceset_num_instants(ss.Inner()))
	_slice := unsafe.Slice((**C.TInstant)(unsafe.Pointer(res)), _n)
	_out := make([]TInstant, _n)
	for _i, _e := range _slice {
		_out[_i] = TInstant{_inner: _e}
	}
	return _out
}


// TsequencesetMaxInstP wraps MEOS C function tsequenceset_max_inst_p.
func TsequencesetMaxInstP(ss TSequenceSet) TInstant {
	res := C.tsequenceset_max_inst_p(ss.Inner())
	return TInstant{_inner: res}
}


// TsequencesetMaxVal wraps MEOS C function tsequenceset_max_val.
func TsequencesetMaxVal(ss TSequenceSet) int {
	res := C.tsequenceset_max_val(ss.Inner())
	return int(res)
}


// TsequencesetMinInstP wraps MEOS C function tsequenceset_min_inst_p.
func TsequencesetMinInstP(ss TSequenceSet) TInstant {
	res := C.tsequenceset_min_inst_p(ss.Inner())
	return TInstant{_inner: res}
}


// TsequencesetMinVal wraps MEOS C function tsequenceset_min_val.
func TsequencesetMinVal(ss TSequenceSet) int {
	res := C.tsequenceset_min_val(ss.Inner())
	return int(res)
}


// TsequencesetNumInstants wraps MEOS C function tsequenceset_num_instants.
func TsequencesetNumInstants(ss TSequenceSet) int {
	res := C.tsequenceset_num_instants(ss.Inner())
	return int(res)
}


// TsequencesetNumTimestamps wraps MEOS C function tsequenceset_num_timestamps.
func TsequencesetNumTimestamps(ss TSequenceSet) int {
	res := C.tsequenceset_num_timestamps(ss.Inner())
	return int(res)
}


// TsequencesetSegments wraps MEOS C function tsequenceset_segments.
func TsequencesetSegments(ss TSequenceSet) []TSequence {
	var _out_count C.int
	res := C.tsequenceset_segments(ss.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.TSequence)(unsafe.Pointer(res)), _n)
	_out := make([]TSequence, _n)
	for _i, _e := range _slice {
		_out[_i] = TSequence{_inner: _e}
	}
	return _out
}


// TsequencesetSequencesP wraps MEOS C function tsequenceset_sequences_p.
func TsequencesetSequencesP(ss TSequenceSet) []TSequence {
	res := C.tsequenceset_sequences_p(ss.Inner())
	_n := int(C.temporal_num_sequences((*C.Temporal)(unsafe.Pointer(ss.Inner()))))
	_slice := unsafe.Slice((**C.TSequence)(unsafe.Pointer(res)), _n)
	_out := make([]TSequence, _n)
	for _i, _e := range _slice {
		_out[_i] = TSequence{_inner: _e}
	}
	return _out
}


// TsequencesetStartTimestamptz wraps MEOS C function tsequenceset_start_timestamptz.
func TsequencesetStartTimestamptz(ss TSequenceSet) int {
	res := C.tsequenceset_start_timestamptz(ss.Inner())
	return int(res)
}


// TsequencesetTime wraps MEOS C function tsequenceset_time.
func TsequencesetTime(ss TSequenceSet) *SpanSet {
	res := C.tsequenceset_time(ss.Inner())
	return &SpanSet{_inner: res}
}


// TsequencesetTimestamptzN wraps MEOS C function tsequenceset_timestamptz_n.
func TsequencesetTimestamptzN(ss TSequenceSet, n int) (bool, int) {
	var _out_result C.int
	res := C.tsequenceset_timestamptz_n(ss.Inner(), C.int(n), &_out_result)
	return bool(res), int(_out_result)
}


// TsequencesetTimestamps wraps MEOS C function tsequenceset_timestamps.
func TsequencesetTimestamps(ss TSequenceSet) []int {
	var _out_count C.int
	res := C.tsequenceset_timestamps(ss.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((*C.int)(unsafe.Pointer(res)), _n)
	_out := make([]int, _n)
	for _i, _e := range _slice {
		_out[_i] = int(_e)
	}
	return _out
}


// TemporalRestart wraps MEOS C function temporal_restart.
func TemporalRestart(temp Temporal, count int) {
	C.temporal_restart(temp.Inner(), C.int(count))
}


// TemporalTsequence wraps MEOS C function temporal_tsequence.
func TemporalTsequence(temp Temporal, interp Interpolation) TSequence {
	res := C.temporal_tsequence(temp.Inner(), C.interpType(interp))
	return TSequence{_inner: res}
}


// TemporalTsequenceset wraps MEOS C function temporal_tsequenceset.
func TemporalTsequenceset(temp Temporal, interp Interpolation) TSequenceSet {
	res := C.temporal_tsequenceset(temp.Inner(), C.interpType(interp))
	return TSequenceSet{_inner: res}
}


// TODO tinstant_shift_time: unsupported param const int *
// func TinstantShiftTime(...) { /* not yet handled by codegen */ }


// TinstantToTsequence wraps MEOS C function tinstant_to_tsequence.
func TinstantToTsequence(inst TInstant, interp Interpolation) TSequence {
	res := C.tinstant_to_tsequence(inst.Inner(), C.interpType(interp))
	return TSequence{_inner: res}
}


// TinstantToTsequenceFree wraps MEOS C function tinstant_to_tsequence_free.
func TinstantToTsequenceFree(inst TInstant, interp Interpolation) TSequence {
	res := C.tinstant_to_tsequence_free(inst.Inner(), C.interpType(interp))
	return TSequence{_inner: res}
}


// TinstantToTsequenceset wraps MEOS C function tinstant_to_tsequenceset.
func TinstantToTsequenceset(inst TInstant, interp Interpolation) TSequenceSet {
	res := C.tinstant_to_tsequenceset(inst.Inner(), C.interpType(interp))
	return TSequenceSet{_inner: res}
}


// TsequenceRestart wraps MEOS C function tsequence_restart.
func TsequenceRestart(seq TSequence, count int) {
	C.tsequence_restart(seq.Inner(), C.int(count))
}


// TsequenceSetInterp wraps MEOS C function tsequence_set_interp.
func TsequenceSetInterp(seq TSequence, interp Interpolation) Temporal {
	res := C.tsequence_set_interp(seq.Inner(), C.interpType(interp))
	return CreateTemporal(res)
}


// TODO tsequence_shift_scale_time: unsupported param const int *
// func TsequenceShiftScaleTime(...) { /* not yet handled by codegen */ }


// TsequenceSubseq wraps MEOS C function tsequence_subseq.
func TsequenceSubseq(seq TSequence, from int, to int, lower_inc bool, upper_inc bool) TSequence {
	res := C.tsequence_subseq(seq.Inner(), C.int(from), C.int(to), C.bool(lower_inc), C.bool(upper_inc))
	return TSequence{_inner: res}
}


// TsequenceToTinstant wraps MEOS C function tsequence_to_tinstant.
func TsequenceToTinstant(seq TSequence) TInstant {
	res := C.tsequence_to_tinstant(seq.Inner())
	return TInstant{_inner: res}
}


// TsequenceToTsequenceset wraps MEOS C function tsequence_to_tsequenceset.
func TsequenceToTsequenceset(seq TSequence) TSequenceSet {
	res := C.tsequence_to_tsequenceset(seq.Inner())
	return TSequenceSet{_inner: res}
}


// TsequenceToTsequencesetFree wraps MEOS C function tsequence_to_tsequenceset_free.
func TsequenceToTsequencesetFree(seq TSequence) TSequenceSet {
	res := C.tsequence_to_tsequenceset_free(seq.Inner())
	return TSequenceSet{_inner: res}
}


// TsequenceToTsequencesetInterp wraps MEOS C function tsequence_to_tsequenceset_interp.
func TsequenceToTsequencesetInterp(seq TSequence, interp Interpolation) TSequenceSet {
	res := C.tsequence_to_tsequenceset_interp(seq.Inner(), C.interpType(interp))
	return TSequenceSet{_inner: res}
}


// TsequencesetRestart wraps MEOS C function tsequenceset_restart.
func TsequencesetRestart(ss TSequenceSet, count int) {
	C.tsequenceset_restart(ss.Inner(), C.int(count))
}


// TsequencesetSetInterp wraps MEOS C function tsequenceset_set_interp.
func TsequencesetSetInterp(ss TSequenceSet, interp Interpolation) Temporal {
	res := C.tsequenceset_set_interp(ss.Inner(), C.interpType(interp))
	return CreateTemporal(res)
}


// TODO tsequenceset_shift_scale_time: unsupported param const int *
// func TsequencesetShiftScaleTime(...) { /* not yet handled by codegen */ }


// TsequencesetToDiscrete wraps MEOS C function tsequenceset_to_discrete.
func TsequencesetToDiscrete(ss TSequenceSet) TSequence {
	res := C.tsequenceset_to_discrete(ss.Inner())
	return TSequence{_inner: res}
}


// TsequencesetToLinear wraps MEOS C function tsequenceset_to_linear.
func TsequencesetToLinear(ss TSequenceSet) TSequenceSet {
	res := C.tsequenceset_to_linear(ss.Inner())
	return TSequenceSet{_inner: res}
}


// TsequencesetToStep wraps MEOS C function tsequenceset_to_step.
func TsequencesetToStep(ss TSequenceSet) TSequenceSet {
	res := C.tsequenceset_to_step(ss.Inner())
	return TSequenceSet{_inner: res}
}


// TsequencesetToTinstant wraps MEOS C function tsequenceset_to_tinstant.
func TsequencesetToTinstant(ss TSequenceSet) TInstant {
	res := C.tsequenceset_to_tinstant(ss.Inner())
	return TInstant{_inner: res}
}


// TsequencesetToTsequence wraps MEOS C function tsequenceset_to_tsequence.
func TsequencesetToTsequence(ss TSequenceSet) TSequence {
	res := C.tsequenceset_to_tsequence(ss.Inner())
	return TSequence{_inner: res}
}


// TinstantMerge wraps MEOS C function tinstant_merge.
func TinstantMerge(inst1 TInstant, inst2 TInstant) Temporal {
	res := C.tinstant_merge(inst1.Inner(), inst2.Inner())
	return CreateTemporal(res)
}


// TinstantMergeArray wraps MEOS C function tinstant_merge_array.
func TinstantMergeArray(instants []TInstant) Temporal {
	_c_instants := make([]*C.TInstant, len(instants))
	for _i, _v := range instants { _c_instants[_i] = _v._inner }
	res := C.tinstant_merge_array((**C.TInstant)(unsafe.Pointer(&_c_instants[0])), C.int(len(instants)))
	return CreateTemporal(res)
}


// TODO tsequence_append_tinstant: unsupported param const int *
// func TsequenceAppendTinstant(...) { /* not yet handled by codegen */ }


// TsequenceAppendTsequence wraps MEOS C function tsequence_append_tsequence.
func TsequenceAppendTsequence(seq1 TSequence, seq2 TSequence, expand bool) Temporal {
	res := C.tsequence_append_tsequence(seq1.Inner(), seq2.Inner(), C.bool(expand))
	return CreateTemporal(res)
}


// TsequenceDeleteTimestamptz wraps MEOS C function tsequence_delete_timestamptz.
func TsequenceDeleteTimestamptz(seq TSequence, t int, connect bool) Temporal {
	res := C.tsequence_delete_timestamptz(seq.Inner(), C.int(t), C.bool(connect))
	return CreateTemporal(res)
}


// TsequenceDeleteTstzset wraps MEOS C function tsequence_delete_tstzset.
func TsequenceDeleteTstzset(seq TSequence, s *Set, connect bool) Temporal {
	res := C.tsequence_delete_tstzset(seq.Inner(), s._inner, C.bool(connect))
	return CreateTemporal(res)
}


// TsequenceDeleteTstzspan wraps MEOS C function tsequence_delete_tstzspan.
func TsequenceDeleteTstzspan(seq TSequence, s *Span, connect bool) Temporal {
	res := C.tsequence_delete_tstzspan(seq.Inner(), s._inner, C.bool(connect))
	return CreateTemporal(res)
}


// TsequenceDeleteTstzspanset wraps MEOS C function tsequence_delete_tstzspanset.
func TsequenceDeleteTstzspanset(seq TSequence, ss *SpanSet, connect bool) Temporal {
	res := C.tsequence_delete_tstzspanset(seq.Inner(), ss._inner, C.bool(connect))
	return CreateTemporal(res)
}


// TsequenceInsert wraps MEOS C function tsequence_insert.
func TsequenceInsert(seq1 TSequence, seq2 TSequence, connect bool) Temporal {
	res := C.tsequence_insert(seq1.Inner(), seq2.Inner(), C.bool(connect))
	return CreateTemporal(res)
}


// TsequenceMerge wraps MEOS C function tsequence_merge.
func TsequenceMerge(seq1 TSequence, seq2 TSequence) Temporal {
	res := C.tsequence_merge(seq1.Inner(), seq2.Inner())
	return CreateTemporal(res)
}


// TsequenceMergeArray wraps MEOS C function tsequence_merge_array.
func TsequenceMergeArray(sequences []TSequence) Temporal {
	_c_sequences := make([]*C.TSequence, len(sequences))
	for _i, _v := range sequences { _c_sequences[_i] = _v._inner }
	res := C.tsequence_merge_array((**C.TSequence)(unsafe.Pointer(&_c_sequences[0])), C.int(len(sequences)))
	return CreateTemporal(res)
}


// TODO tsequenceset_append_tinstant: unsupported param const int *
// func TsequencesetAppendTinstant(...) { /* not yet handled by codegen */ }


// TsequencesetAppendTsequence wraps MEOS C function tsequenceset_append_tsequence.
func TsequencesetAppendTsequence(ss TSequenceSet, seq TSequence, expand bool) TSequenceSet {
	res := C.tsequenceset_append_tsequence(ss.Inner(), seq.Inner(), C.bool(expand))
	return TSequenceSet{_inner: res}
}


// TsequencesetDeleteTimestamptz wraps MEOS C function tsequenceset_delete_timestamptz.
func TsequencesetDeleteTimestamptz(ss TSequenceSet, t int) TSequenceSet {
	res := C.tsequenceset_delete_timestamptz(ss.Inner(), C.int(t))
	return TSequenceSet{_inner: res}
}


// TsequencesetDeleteTstzset wraps MEOS C function tsequenceset_delete_tstzset.
func TsequencesetDeleteTstzset(ss TSequenceSet, s *Set) TSequenceSet {
	res := C.tsequenceset_delete_tstzset(ss.Inner(), s._inner)
	return TSequenceSet{_inner: res}
}


// TsequencesetDeleteTstzspan wraps MEOS C function tsequenceset_delete_tstzspan.
func TsequencesetDeleteTstzspan(ss TSequenceSet, s *Span) TSequenceSet {
	res := C.tsequenceset_delete_tstzspan(ss.Inner(), s._inner)
	return TSequenceSet{_inner: res}
}


// TsequencesetDeleteTstzspanset wraps MEOS C function tsequenceset_delete_tstzspanset.
func TsequencesetDeleteTstzspanset(ss TSequenceSet, ps *SpanSet) TSequenceSet {
	res := C.tsequenceset_delete_tstzspanset(ss.Inner(), ps._inner)
	return TSequenceSet{_inner: res}
}


// TsequencesetInsert wraps MEOS C function tsequenceset_insert.
func TsequencesetInsert(ss1 TSequenceSet, ss2 TSequenceSet) TSequenceSet {
	res := C.tsequenceset_insert(ss1.Inner(), ss2.Inner())
	return TSequenceSet{_inner: res}
}


// TsequencesetMerge wraps MEOS C function tsequenceset_merge.
func TsequencesetMerge(ss1 TSequenceSet, ss2 TSequenceSet) TSequenceSet {
	res := C.tsequenceset_merge(ss1.Inner(), ss2.Inner())
	return TSequenceSet{_inner: res}
}


// TsequencesetMergeArray wraps MEOS C function tsequenceset_merge_array.
func TsequencesetMergeArray(seqsets []TSequenceSet) TSequenceSet {
	_c_seqsets := make([]*C.TSequenceSet, len(seqsets))
	for _i, _v := range seqsets { _c_seqsets[_i] = _v._inner }
	res := C.tsequenceset_merge_array((**C.TSequenceSet)(unsafe.Pointer(&_c_seqsets[0])), C.int(len(seqsets)))
	return TSequenceSet{_inner: res}
}


// TsequenceExpandBbox wraps MEOS C function tsequence_expand_bbox.
func TsequenceExpandBbox(seq TSequence, inst TInstant) {
	C.tsequence_expand_bbox(seq.Inner(), inst.Inner())
}


// TsequenceSetBbox wraps MEOS C function tsequence_set_bbox.
func TsequenceSetBbox(seq TSequence, box unsafe.Pointer) {
	C.tsequence_set_bbox(seq.Inner(), unsafe.Pointer(box))
}


// TsequencesetExpandBbox wraps MEOS C function tsequenceset_expand_bbox.
func TsequencesetExpandBbox(ss TSequenceSet, seq TSequence) {
	C.tsequenceset_expand_bbox(ss.Inner(), seq.Inner())
}


// TsequencesetSetBbox wraps MEOS C function tsequenceset_set_bbox.
func TsequencesetSetBbox(ss TSequenceSet, box unsafe.Pointer) {
	C.tsequenceset_set_bbox(ss.Inner(), unsafe.Pointer(box))
}


// TcontseqAfterTimestamptz wraps MEOS C function tcontseq_after_timestamptz.
func TcontseqAfterTimestamptz(seq TSequence, t int, strict bool) TSequence {
	res := C.tcontseq_after_timestamptz(seq.Inner(), C.int(t), C.bool(strict))
	return TSequence{_inner: res}
}


// TcontseqBeforeTimestamptz wraps MEOS C function tcontseq_before_timestamptz.
func TcontseqBeforeTimestamptz(seq TSequence, t int, strict bool) TSequence {
	res := C.tcontseq_before_timestamptz(seq.Inner(), C.int(t), C.bool(strict))
	return TSequence{_inner: res}
}


// TcontseqRestrictMinmax wraps MEOS C function tcontseq_restrict_minmax.
func TcontseqRestrictMinmax(seq TSequence, min bool, atfunc bool) TSequenceSet {
	res := C.tcontseq_restrict_minmax(seq.Inner(), C.bool(min), C.bool(atfunc))
	return TSequenceSet{_inner: res}
}


// TdiscseqAfterTimestamptz wraps MEOS C function tdiscseq_after_timestamptz.
func TdiscseqAfterTimestamptz(seq TSequence, t int, strict bool) TSequence {
	res := C.tdiscseq_after_timestamptz(seq.Inner(), C.int(t), C.bool(strict))
	return TSequence{_inner: res}
}


// TdiscseqBeforeTimestamptz wraps MEOS C function tdiscseq_before_timestamptz.
func TdiscseqBeforeTimestamptz(seq TSequence, t int, strict bool) TSequence {
	res := C.tdiscseq_before_timestamptz(seq.Inner(), C.int(t), C.bool(strict))
	return TSequence{_inner: res}
}


// TdiscseqRestrictMinmax wraps MEOS C function tdiscseq_restrict_minmax.
func TdiscseqRestrictMinmax(seq TSequence, min bool, atfunc bool) TSequence {
	res := C.tdiscseq_restrict_minmax(seq.Inner(), C.bool(min), C.bool(atfunc))
	return TSequence{_inner: res}
}


// TemporalBboxRestrictSet wraps MEOS C function temporal_bbox_restrict_set.
func TemporalBboxRestrictSet(temp Temporal, set *Set) bool {
	res := C.temporal_bbox_restrict_set(temp.Inner(), set._inner)
	return bool(res)
}


// TemporalRestrictMinmax wraps MEOS C function temporal_restrict_minmax.
func TemporalRestrictMinmax(temp Temporal, min bool, atfunc bool) Temporal {
	res := C.temporal_restrict_minmax(temp.Inner(), C.bool(min), C.bool(atfunc))
	return CreateTemporal(res)
}


// TemporalRestrictTimestamptz wraps MEOS C function temporal_restrict_timestamptz.
func TemporalRestrictTimestamptz(temp Temporal, t int, atfunc bool) Temporal {
	res := C.temporal_restrict_timestamptz(temp.Inner(), C.int(t), C.bool(atfunc))
	return CreateTemporal(res)
}


// TemporalRestrictTstzset wraps MEOS C function temporal_restrict_tstzset.
func TemporalRestrictTstzset(temp Temporal, s *Set, atfunc bool) Temporal {
	res := C.temporal_restrict_tstzset(temp.Inner(), s._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TemporalRestrictTstzspan wraps MEOS C function temporal_restrict_tstzspan.
func TemporalRestrictTstzspan(temp Temporal, s *Span, atfunc bool) Temporal {
	res := C.temporal_restrict_tstzspan(temp.Inner(), s._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TemporalRestrictTstzspanset wraps MEOS C function temporal_restrict_tstzspanset.
func TemporalRestrictTstzspanset(temp Temporal, ss *SpanSet, atfunc bool) Temporal {
	res := C.temporal_restrict_tstzspanset(temp.Inner(), ss._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TemporalRestrictValues wraps MEOS C function temporal_restrict_values.
func TemporalRestrictValues(temp Temporal, set *Set, atfunc bool) Temporal {
	res := C.temporal_restrict_values(temp.Inner(), set._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TinstantAfterTimestamptz wraps MEOS C function tinstant_after_timestamptz.
func TinstantAfterTimestamptz(inst TInstant, t int, strict bool) TInstant {
	res := C.tinstant_after_timestamptz(inst.Inner(), C.int(t), C.bool(strict))
	return TInstant{_inner: res}
}


// TinstantBeforeTimestamptz wraps MEOS C function tinstant_before_timestamptz.
func TinstantBeforeTimestamptz(inst TInstant, t int, strict bool) TInstant {
	res := C.tinstant_before_timestamptz(inst.Inner(), C.int(t), C.bool(strict))
	return TInstant{_inner: res}
}


// TinstantRestrictTstzspan wraps MEOS C function tinstant_restrict_tstzspan.
func TinstantRestrictTstzspan(inst TInstant, period *Span, atfunc bool) TInstant {
	res := C.tinstant_restrict_tstzspan(inst.Inner(), period._inner, C.bool(atfunc))
	return TInstant{_inner: res}
}


// TinstantRestrictTstzspanset wraps MEOS C function tinstant_restrict_tstzspanset.
func TinstantRestrictTstzspanset(inst TInstant, ss *SpanSet, atfunc bool) TInstant {
	res := C.tinstant_restrict_tstzspanset(inst.Inner(), ss._inner, C.bool(atfunc))
	return TInstant{_inner: res}
}


// TinstantRestrictTimestamptz wraps MEOS C function tinstant_restrict_timestamptz.
func TinstantRestrictTimestamptz(inst TInstant, t int, atfunc bool) TInstant {
	res := C.tinstant_restrict_timestamptz(inst.Inner(), C.int(t), C.bool(atfunc))
	return TInstant{_inner: res}
}


// TinstantRestrictTstzset wraps MEOS C function tinstant_restrict_tstzset.
func TinstantRestrictTstzset(inst TInstant, s *Set, atfunc bool) TInstant {
	res := C.tinstant_restrict_tstzset(inst.Inner(), s._inner, C.bool(atfunc))
	return TInstant{_inner: res}
}


// TinstantRestrictValues wraps MEOS C function tinstant_restrict_values.
func TinstantRestrictValues(inst TInstant, set *Set, atfunc bool) TInstant {
	res := C.tinstant_restrict_values(inst.Inner(), set._inner, C.bool(atfunc))
	return TInstant{_inner: res}
}


// TnumberRestrictSpan wraps MEOS C function tnumber_restrict_span.
func TnumberRestrictSpan(temp Temporal, span *Span, atfunc bool) Temporal {
	res := C.tnumber_restrict_span(temp.Inner(), span._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TnumberRestrictSpanset wraps MEOS C function tnumber_restrict_spanset.
func TnumberRestrictSpanset(temp Temporal, ss *SpanSet, atfunc bool) Temporal {
	res := C.tnumber_restrict_spanset(temp.Inner(), ss._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TnumberinstRestrictSpan wraps MEOS C function tnumberinst_restrict_span.
func TnumberinstRestrictSpan(inst TInstant, span *Span, atfunc bool) TInstant {
	res := C.tnumberinst_restrict_span(inst.Inner(), span._inner, C.bool(atfunc))
	return TInstant{_inner: res}
}


// TnumberinstRestrictSpanset wraps MEOS C function tnumberinst_restrict_spanset.
func TnumberinstRestrictSpanset(inst TInstant, ss *SpanSet, atfunc bool) TInstant {
	res := C.tnumberinst_restrict_spanset(inst.Inner(), ss._inner, C.bool(atfunc))
	return TInstant{_inner: res}
}


// TnumberseqsetRestrictSpan wraps MEOS C function tnumberseqset_restrict_span.
func TnumberseqsetRestrictSpan(ss TSequenceSet, span *Span, atfunc bool) TSequenceSet {
	res := C.tnumberseqset_restrict_span(ss.Inner(), span._inner, C.bool(atfunc))
	return TSequenceSet{_inner: res}
}


// TnumberseqsetRestrictSpanset wraps MEOS C function tnumberseqset_restrict_spanset.
func TnumberseqsetRestrictSpanset(ss TSequenceSet, spanset *SpanSet, atfunc bool) TSequenceSet {
	res := C.tnumberseqset_restrict_spanset(ss.Inner(), spanset._inner, C.bool(atfunc))
	return TSequenceSet{_inner: res}
}


// TsequenceAtTimestamptz wraps MEOS C function tsequence_at_timestamptz.
func TsequenceAtTimestamptz(seq TSequence, t int) TInstant {
	res := C.tsequence_at_timestamptz(seq.Inner(), C.int(t))
	return TInstant{_inner: res}
}


// TsequenceRestrictTstzspan wraps MEOS C function tsequence_restrict_tstzspan.
func TsequenceRestrictTstzspan(seq TSequence, s *Span, atfunc bool) Temporal {
	res := C.tsequence_restrict_tstzspan(seq.Inner(), s._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TsequenceRestrictTstzspanset wraps MEOS C function tsequence_restrict_tstzspanset.
func TsequenceRestrictTstzspanset(seq TSequence, ss *SpanSet, atfunc bool) Temporal {
	res := C.tsequence_restrict_tstzspanset(seq.Inner(), ss._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TsequencesetAfterTimestamptz wraps MEOS C function tsequenceset_after_timestamptz.
func TsequencesetAfterTimestamptz(ss TSequenceSet, t int, strict bool) TSequenceSet {
	res := C.tsequenceset_after_timestamptz(ss.Inner(), C.int(t), C.bool(strict))
	return TSequenceSet{_inner: res}
}


// TsequencesetBeforeTimestamptz wraps MEOS C function tsequenceset_before_timestamptz.
func TsequencesetBeforeTimestamptz(ss TSequenceSet, t int, strict bool) TSequenceSet {
	res := C.tsequenceset_before_timestamptz(ss.Inner(), C.int(t), C.bool(strict))
	return TSequenceSet{_inner: res}
}


// TsequencesetRestrictMinmax wraps MEOS C function tsequenceset_restrict_minmax.
func TsequencesetRestrictMinmax(ss TSequenceSet, min bool, atfunc bool) TSequenceSet {
	res := C.tsequenceset_restrict_minmax(ss.Inner(), C.bool(min), C.bool(atfunc))
	return TSequenceSet{_inner: res}
}


// TsequencesetRestrictTstzspan wraps MEOS C function tsequenceset_restrict_tstzspan.
func TsequencesetRestrictTstzspan(ss TSequenceSet, s *Span, atfunc bool) TSequenceSet {
	res := C.tsequenceset_restrict_tstzspan(ss.Inner(), s._inner, C.bool(atfunc))
	return TSequenceSet{_inner: res}
}


// TsequencesetRestrictTstzspanset wraps MEOS C function tsequenceset_restrict_tstzspanset.
func TsequencesetRestrictTstzspanset(ss TSequenceSet, ps *SpanSet, atfunc bool) TSequenceSet {
	res := C.tsequenceset_restrict_tstzspanset(ss.Inner(), ps._inner, C.bool(atfunc))
	return TSequenceSet{_inner: res}
}


// TsequencesetRestrictTimestamptz wraps MEOS C function tsequenceset_restrict_timestamptz.
func TsequencesetRestrictTimestamptz(ss TSequenceSet, t int, atfunc bool) Temporal {
	res := C.tsequenceset_restrict_timestamptz(ss.Inner(), C.int(t), C.bool(atfunc))
	return CreateTemporal(res)
}


// TsequencesetRestrictTstzset wraps MEOS C function tsequenceset_restrict_tstzset.
func TsequencesetRestrictTstzset(ss TSequenceSet, s *Set, atfunc bool) Temporal {
	res := C.tsequenceset_restrict_tstzset(ss.Inner(), s._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TsequencesetRestrictValues wraps MEOS C function tsequenceset_restrict_values.
func TsequencesetRestrictValues(ss TSequenceSet, s *Set, atfunc bool) TSequenceSet {
	res := C.tsequenceset_restrict_values(ss.Inner(), s._inner, C.bool(atfunc))
	return TSequenceSet{_inner: res}
}


// TinstantCmp wraps MEOS C function tinstant_cmp.
func TinstantCmp(inst1 TInstant, inst2 TInstant) int {
	res := C.tinstant_cmp(inst1.Inner(), inst2.Inner())
	return int(res)
}


// TinstantEq wraps MEOS C function tinstant_eq.
func TinstantEq(inst1 TInstant, inst2 TInstant) bool {
	res := C.tinstant_eq(inst1.Inner(), inst2.Inner())
	return bool(res)
}


// TsequenceCmp wraps MEOS C function tsequence_cmp.
func TsequenceCmp(seq1 TSequence, seq2 TSequence) int {
	res := C.tsequence_cmp(seq1.Inner(), seq2.Inner())
	return int(res)
}


// TsequenceEq wraps MEOS C function tsequence_eq.
func TsequenceEq(seq1 TSequence, seq2 TSequence) bool {
	res := C.tsequence_eq(seq1.Inner(), seq2.Inner())
	return bool(res)
}


// TsequencesetCmp wraps MEOS C function tsequenceset_cmp.
func TsequencesetCmp(ss1 TSequenceSet, ss2 TSequenceSet) int {
	res := C.tsequenceset_cmp(ss1.Inner(), ss2.Inner())
	return int(res)
}


// TsequencesetEq wraps MEOS C function tsequenceset_eq.
func TsequencesetEq(ss1 TSequenceSet, ss2 TSequenceSet) bool {
	res := C.tsequenceset_eq(ss1.Inner(), ss2.Inner())
	return bool(res)
}


// TnumberinstAbs wraps MEOS C function tnumberinst_abs.
func TnumberinstAbs(inst TInstant) TInstant {
	res := C.tnumberinst_abs(inst.Inner())
	return TInstant{_inner: res}
}


// TnumberseqAbs wraps MEOS C function tnumberseq_abs.
func TnumberseqAbs(seq TSequence) TSequence {
	res := C.tnumberseq_abs(seq.Inner())
	return TSequence{_inner: res}
}


// TnumberseqAngularDifference wraps MEOS C function tnumberseq_angular_difference.
func TnumberseqAngularDifference(seq TSequence) TSequence {
	res := C.tnumberseq_angular_difference(seq.Inner())
	return TSequence{_inner: res}
}


// TnumberseqDeltaValue wraps MEOS C function tnumberseq_delta_value.
func TnumberseqDeltaValue(seq TSequence) TSequence {
	res := C.tnumberseq_delta_value(seq.Inner())
	return TSequence{_inner: res}
}


// TnumberseqsetAbs wraps MEOS C function tnumberseqset_abs.
func TnumberseqsetAbs(ss TSequenceSet) TSequenceSet {
	res := C.tnumberseqset_abs(ss.Inner())
	return TSequenceSet{_inner: res}
}


// TnumberseqsetAngularDifference wraps MEOS C function tnumberseqset_angular_difference.
func TnumberseqsetAngularDifference(ss TSequenceSet) TSequence {
	res := C.tnumberseqset_angular_difference(ss.Inner())
	return TSequence{_inner: res}
}


// TnumberseqsetDeltaValue wraps MEOS C function tnumberseqset_delta_value.
func TnumberseqsetDeltaValue(ss TSequenceSet) TSequenceSet {
	res := C.tnumberseqset_delta_value(ss.Inner())
	return TSequenceSet{_inner: res}
}


// NadTBOXTBOX wraps MEOS C function nad_tbox_tbox.
func NadTBOXTBOX(box1 *TBox, box2 *TBox) float64 {
	res := C.nad_tbox_tbox(box1._inner, box2._inner)
	return float64(res)
}


// NadTnumberTBOX wraps MEOS C function nad_tnumber_tbox.
func NadTnumberTBOX(temp Temporal, box *TBox) float64 {
	res := C.nad_tnumber_tbox(temp.Inner(), box._inner)
	return float64(res)
}


// NadTnumberTnumber wraps MEOS C function nad_tnumber_tnumber.
func NadTnumberTnumber(temp1 Temporal, temp2 Temporal) float64 {
	res := C.nad_tnumber_tnumber(temp1.Inner(), temp2.Inner())
	return float64(res)
}


// TnumberseqIntegral wraps MEOS C function tnumberseq_integral.
func TnumberseqIntegral(seq TSequence) float64 {
	res := C.tnumberseq_integral(seq.Inner())
	return float64(res)
}


// TnumberseqTwavg wraps MEOS C function tnumberseq_twavg.
func TnumberseqTwavg(seq TSequence) float64 {
	res := C.tnumberseq_twavg(seq.Inner())
	return float64(res)
}


// TnumberseqsetIntegral wraps MEOS C function tnumberseqset_integral.
func TnumberseqsetIntegral(ss TSequenceSet) float64 {
	res := C.tnumberseqset_integral(ss.Inner())
	return float64(res)
}


// TnumberseqsetTwavg wraps MEOS C function tnumberseqset_twavg.
func TnumberseqsetTwavg(ss TSequenceSet) float64 {
	res := C.tnumberseqset_twavg(ss.Inner())
	return float64(res)
}


// TemporalCompact wraps MEOS C function temporal_compact.
func TemporalCompact(temp Temporal) Temporal {
	res := C.temporal_compact(temp.Inner())
	return CreateTemporal(res)
}


// TsequenceCompact wraps MEOS C function tsequence_compact.
func TsequenceCompact(seq TSequence) TSequence {
	res := C.tsequence_compact(seq.Inner())
	return TSequence{_inner: res}
}


// TsequencesetCompact wraps MEOS C function tsequenceset_compact.
func TsequencesetCompact(ss TSequenceSet) TSequenceSet {
	res := C.tsequenceset_compact(ss.Inner())
	return TSequenceSet{_inner: res}
}


// TemporalSkiplistMake wraps MEOS C function temporal_skiplist_make.
func TemporalSkiplistMake() *SkipList {
	res := C.temporal_skiplist_make()
	return &SkipList{_inner: res}
}


// SkiplistMake wraps MEOS C function skiplist_make.
func SkiplistMake(key_size uint, value_size uint, comp_fn unsafe.Pointer, merge_fn unsafe.Pointer) *SkipList {
	res := C.skiplist_make(C.size_t(key_size), C.size_t(value_size), comp_fn, merge_fn)
	return &SkipList{_inner: res}
}


// SkiplistSearch wraps MEOS C function skiplist_search.
func SkiplistSearch(list *SkipList, key unsafe.Pointer) (int, ) {
	var _out_value C.void
	res := C.skiplist_search(list._inner, unsafe.Pointer(key), &_out_value)
	return int(res), _out_value
}


// SkiplistFree wraps MEOS C function skiplist_free.
func SkiplistFree(list *SkipList) {
	C.skiplist_free(list._inner)
}


// SkiplistSplice wraps MEOS C function skiplist_splice.
func SkiplistSplice(list *SkipList, keys unsafe.Pointer, values unsafe.Pointer, count int, func_ unsafe.Pointer, crossings bool, sktype SkipListType) {
	C.skiplist_splice(list._inner, unsafe.Pointer(keys), unsafe.Pointer(values), C.int(count), func_, C.bool(crossings), C.SkipListType(sktype))
}


// TemporalSkiplistSplice wraps MEOS C function temporal_skiplist_splice.
func TemporalSkiplistSplice(list *SkipList, values unsafe.Pointer, count int, func_ unsafe.Pointer, crossings bool) {
	C.temporal_skiplist_splice(list._inner, unsafe.Pointer(values), C.int(count), func_, C.bool(crossings))
}


// SkiplistValues wraps MEOS C function skiplist_values.
func SkiplistValues(list *SkipList) unsafe.Pointer {
	res := C.skiplist_values(list._inner)
	return unsafe.Pointer(res)
}


// SkiplistKeysValues wraps MEOS C function skiplist_keys_values.
func SkiplistKeysValues(list *SkipList, values unsafe.Pointer) unsafe.Pointer {
	res := C.skiplist_keys_values(list._inner, unsafe.Pointer(values))
	return unsafe.Pointer(res)
}


// TODO temporal_app_tinst_transfn: unsupported param const int *
// func TemporalAppTinstTransfn(...) { /* not yet handled by codegen */ }


// TemporalAppTseqTransfn wraps MEOS C function temporal_app_tseq_transfn.
func TemporalAppTseqTransfn(state Temporal, seq TSequence) Temporal {
	res := C.temporal_app_tseq_transfn(state.Inner(), seq.Inner())
	return CreateTemporal(res)
}

