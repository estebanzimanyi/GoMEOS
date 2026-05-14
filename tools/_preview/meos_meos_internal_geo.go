package generated

// #include <stddef.h>
import "C"
import (
	"unsafe"

	"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

// ProjGetContext wraps MEOS C function proj_get_context.
func ProjGetContext() *PJContext {
	res := C.proj_get_context()
	return &PJContext{_inner: res}
}


// TODO point_round: unsupported return type int *
// func PointRound(...) { /* not yet handled by codegen */ }


// STBOXSet wraps MEOS C function stbox_set.
func STBOXSet(hasx bool, hasz bool, geodetic bool, srid int, xmin float64, xmax float64, ymin float64, ymax float64, zmin float64, zmax float64, s *Span, box *STBox) {
	C.stbox_set(C.bool(hasx), C.bool(hasz), C.bool(geodetic), C.int(srid), C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.double(zmin), C.double(zmax), s._inner, box._inner)
}


// TODO gbox_set_stbox: unsupported param const int *
// func GboxSetSTBOX(...) { /* not yet handled by codegen */ }


// TODO geo_set_stbox: unsupported param const int *
// func GeoSetSTBOX(...) { /* not yet handled by codegen */ }


// SpatialsetSetSTBOX wraps MEOS C function spatialset_set_stbox.
func SpatialsetSetSTBOX(set *Set, box *STBox) {
	C.spatialset_set_stbox(set._inner, box._inner)
}


// TODO stbox_set_box3d: unsupported param int *
// func STBOXSetBox3d(...) { /* not yet handled by codegen */ }


// TODO stbox_set_gbox: unsupported param int *
// func STBOXSetGbox(...) { /* not yet handled by codegen */ }


// TstzsetSetSTBOX wraps MEOS C function tstzset_set_stbox.
func TstzsetSetSTBOX(s *Set, box *STBox) {
	C.tstzset_set_stbox(s._inner, box._inner)
}


// TstzspanSetSTBOX wraps MEOS C function tstzspan_set_stbox.
func TstzspanSetSTBOX(s *Span, box *STBox) {
	C.tstzspan_set_stbox(s._inner, box._inner)
}


// TstzspansetSetSTBOX wraps MEOS C function tstzspanset_set_stbox.
func TstzspansetSetSTBOX(s *SpanSet, box *STBox) {
	C.tstzspanset_set_stbox(s._inner, box._inner)
}


// STBOXExpand wraps MEOS C function stbox_expand.
func STBOXExpand(box1 *STBox, box2 *STBox) {
	C.stbox_expand(box1._inner, box2._inner)
}


// InterSTBOXSTBOX wraps MEOS C function inter_stbox_stbox.
func InterSTBOXSTBOX(box1 *STBox, box2 *STBox) (bool, *STBox) {
	var _out_result C.STBox
	res := C.inter_stbox_stbox(box1._inner, box2._inner, &_out_result)
	return bool(res), &STBox{_inner: &_out_result}
}


// TgeogpointinstIn wraps MEOS C function tgeogpointinst_in.
func TgeogpointinstIn(str string) TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeogpointinst_in(_c_str)
	return TInstant{_inner: res}
}


// TgeogpointseqIn wraps MEOS C function tgeogpointseq_in.
func TgeogpointseqIn(str string, interp Interpolation) TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeogpointseq_in(_c_str, C.interpType(interp))
	return TSequence{_inner: res}
}


// TgeogpointseqsetIn wraps MEOS C function tgeogpointseqset_in.
func TgeogpointseqsetIn(str string) TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeogpointseqset_in(_c_str)
	return TSequenceSet{_inner: res}
}


// TgeompointinstIn wraps MEOS C function tgeompointinst_in.
func TgeompointinstIn(str string) TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeompointinst_in(_c_str)
	return TInstant{_inner: res}
}


// TgeompointseqIn wraps MEOS C function tgeompointseq_in.
func TgeompointseqIn(str string, interp Interpolation) TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeompointseq_in(_c_str, C.interpType(interp))
	return TSequence{_inner: res}
}


// TgeompointseqsetIn wraps MEOS C function tgeompointseqset_in.
func TgeompointseqsetIn(str string) TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeompointseqset_in(_c_str)
	return TSequenceSet{_inner: res}
}


// TgeographyinstIn wraps MEOS C function tgeographyinst_in.
func TgeographyinstIn(str string) TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeographyinst_in(_c_str)
	return TInstant{_inner: res}
}


// TgeographyseqIn wraps MEOS C function tgeographyseq_in.
func TgeographyseqIn(str string, interp Interpolation) TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeographyseq_in(_c_str, C.interpType(interp))
	return TSequence{_inner: res}
}


// TgeographyseqsetIn wraps MEOS C function tgeographyseqset_in.
func TgeographyseqsetIn(str string) TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeographyseqset_in(_c_str)
	return TSequenceSet{_inner: res}
}


// TgeometryinstIn wraps MEOS C function tgeometryinst_in.
func TgeometryinstIn(str string) TInstant {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeometryinst_in(_c_str)
	return TInstant{_inner: res}
}


// TgeometryseqIn wraps MEOS C function tgeometryseq_in.
func TgeometryseqIn(str string, interp Interpolation) TSequence {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeometryseq_in(_c_str, C.interpType(interp))
	return TSequence{_inner: res}
}


// TgeometryseqsetIn wraps MEOS C function tgeometryseqset_in.
func TgeometryseqsetIn(str string) TSequenceSet {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tgeometryseqset_in(_c_str)
	return TSequenceSet{_inner: res}
}


// TspatialSetSTBOX wraps MEOS C function tspatial_set_stbox.
func TspatialSetSTBOX(temp Temporal, box *STBox) {
	C.tspatial_set_stbox(temp.Inner(), box._inner)
}


// TspatialseqSetSTBOX wraps MEOS C function tspatialseq_set_stbox.
func TspatialseqSetSTBOX(seq TSequence, box *STBox) {
	C.tspatialseq_set_stbox(seq.Inner(), box._inner)
}


// TspatialseqsetSetSTBOX wraps MEOS C function tspatialseqset_set_stbox.
func TspatialseqsetSetSTBOX(ss TSequenceSet, box *STBox) {
	C.tspatialseqset_set_stbox(ss.Inner(), box._inner)
}


// TgeoRestrictElevation wraps MEOS C function tgeo_restrict_elevation.
func TgeoRestrictElevation(temp Temporal, s *Span, atfunc bool) Temporal {
	res := C.tgeo_restrict_elevation(temp.Inner(), s._inner, C.bool(atfunc))
	return CreateTemporal(res)
}


// TODO tgeo_restrict_geom: unsupported param const int *
// func TgeoRestrictGeom(...) { /* not yet handled by codegen */ }


// TgeoRestrictSTBOX wraps MEOS C function tgeo_restrict_stbox.
func TgeoRestrictSTBOX(temp Temporal, box *STBox, border_inc bool, atfunc bool) Temporal {
	res := C.tgeo_restrict_stbox(temp.Inner(), box._inner, C.bool(border_inc), C.bool(atfunc))
	return CreateTemporal(res)
}


// TODO tgeoinst_restrict_geom: unsupported param const int *
// func TgeoinstRestrictGeom(...) { /* not yet handled by codegen */ }


// TgeoinstRestrictSTBOX wraps MEOS C function tgeoinst_restrict_stbox.
func TgeoinstRestrictSTBOX(inst TInstant, box *STBox, border_inc bool, atfunc bool) TInstant {
	res := C.tgeoinst_restrict_stbox(inst.Inner(), box._inner, C.bool(border_inc), C.bool(atfunc))
	return TInstant{_inner: res}
}


// TODO tgeoseq_restrict_geom: unsupported param const int *
// func TgeoseqRestrictGeom(...) { /* not yet handled by codegen */ }


// TgeoseqRestrictSTBOX wraps MEOS C function tgeoseq_restrict_stbox.
func TgeoseqRestrictSTBOX(seq TSequence, box *STBox, border_inc bool, atfunc bool) Temporal {
	res := C.tgeoseq_restrict_stbox(seq.Inner(), box._inner, C.bool(border_inc), C.bool(atfunc))
	return CreateTemporal(res)
}


// TODO tgeoseqset_restrict_geom: unsupported param const int *
// func TgeoseqsetRestrictGeom(...) { /* not yet handled by codegen */ }


// TgeoseqsetRestrictSTBOX wraps MEOS C function tgeoseqset_restrict_stbox.
func TgeoseqsetRestrictSTBOX(ss TSequenceSet, box *STBox, border_inc bool, atfunc bool) TSequenceSet {
	res := C.tgeoseqset_restrict_stbox(ss.Inner(), box._inner, C.bool(border_inc), C.bool(atfunc))
	return TSequenceSet{_inner: res}
}


// TspatialinstSRID wraps MEOS C function tspatialinst_srid.
func TspatialinstSRID(inst TInstant) int {
	res := C.tspatialinst_srid(inst.Inner())
	return int(res)
}


// TpointseqAzimuth wraps MEOS C function tpointseq_azimuth.
func TpointseqAzimuth(seq TSequence) TSequenceSet {
	res := C.tpointseq_azimuth(seq.Inner())
	return TSequenceSet{_inner: res}
}


// TpointseqCumulativeLength wraps MEOS C function tpointseq_cumulative_length.
func TpointseqCumulativeLength(seq TSequence, prevlength float64) TSequence {
	res := C.tpointseq_cumulative_length(seq.Inner(), C.double(prevlength))
	return TSequence{_inner: res}
}


// TpointseqIsSimple wraps MEOS C function tpointseq_is_simple.
func TpointseqIsSimple(seq TSequence) bool {
	res := C.tpointseq_is_simple(seq.Inner())
	return bool(res)
}


// TpointseqLength wraps MEOS C function tpointseq_length.
func TpointseqLength(seq TSequence) float64 {
	res := C.tpointseq_length(seq.Inner())
	return float64(res)
}


// TODO tpointseq_linear_trajectory: unsupported return type int *
// func TpointseqLinearTrajectory(...) { /* not yet handled by codegen */ }


// TgeoseqStboxes wraps MEOS C function tgeoseq_stboxes.
func TgeoseqStboxes(seq TSequence) (*STBox, int) {
	var _out_count C.int
	res := C.tgeoseq_stboxes(seq.Inner(), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// TgeoseqSplitNStboxes wraps MEOS C function tgeoseq_split_n_stboxes.
func TgeoseqSplitNStboxes(seq TSequence, max_count int) (*STBox, int) {
	var _out_count C.int
	res := C.tgeoseq_split_n_stboxes(seq.Inner(), C.int(max_count), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// TpointseqsetAzimuth wraps MEOS C function tpointseqset_azimuth.
func TpointseqsetAzimuth(ss TSequenceSet) TSequenceSet {
	res := C.tpointseqset_azimuth(ss.Inner())
	return TSequenceSet{_inner: res}
}


// TpointseqsetCumulativeLength wraps MEOS C function tpointseqset_cumulative_length.
func TpointseqsetCumulativeLength(ss TSequenceSet) TSequenceSet {
	res := C.tpointseqset_cumulative_length(ss.Inner())
	return TSequenceSet{_inner: res}
}


// TpointseqsetIsSimple wraps MEOS C function tpointseqset_is_simple.
func TpointseqsetIsSimple(ss TSequenceSet) bool {
	res := C.tpointseqset_is_simple(ss.Inner())
	return bool(res)
}


// TpointseqsetLength wraps MEOS C function tpointseqset_length.
func TpointseqsetLength(ss TSequenceSet) float64 {
	res := C.tpointseqset_length(ss.Inner())
	return float64(res)
}


// TgeoseqsetStboxes wraps MEOS C function tgeoseqset_stboxes.
func TgeoseqsetStboxes(ss TSequenceSet) (*STBox, int) {
	var _out_count C.int
	res := C.tgeoseqset_stboxes(ss.Inner(), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// TgeoseqsetSplitNStboxes wraps MEOS C function tgeoseqset_split_n_stboxes.
func TgeoseqsetSplitNStboxes(ss TSequenceSet, max_count int) (*STBox, int) {
	var _out_count C.int
	res := C.tgeoseqset_split_n_stboxes(ss.Inner(), C.int(max_count), &_out_count)
	return &STBox{_inner: res}, int(_out_count)
}


// TgeominstTgeoginst wraps MEOS C function tgeominst_tgeoginst.
func TgeominstTgeoginst(inst TInstant, oper bool) TInstant {
	res := C.tgeominst_tgeoginst(inst.Inner(), C.bool(oper))
	return TInstant{_inner: res}
}


// TgeomseqTgeogseq wraps MEOS C function tgeomseq_tgeogseq.
func TgeomseqTgeogseq(seq TSequence, oper bool) TSequence {
	res := C.tgeomseq_tgeogseq(seq.Inner(), C.bool(oper))
	return TSequence{_inner: res}
}


// TgeomseqsetTgeogseqset wraps MEOS C function tgeomseqset_tgeogseqset.
func TgeomseqsetTgeogseqset(ss TSequenceSet, oper bool) TSequenceSet {
	res := C.tgeomseqset_tgeogseqset(ss.Inner(), C.bool(oper))
	return TSequenceSet{_inner: res}
}


// TgeomTgeog wraps MEOS C function tgeom_tgeog.
func TgeomTgeog(temp Temporal, oper bool) Temporal {
	res := C.tgeom_tgeog(temp.Inner(), C.bool(oper))
	return CreateTemporal(res)
}


// TgeoTpoint wraps MEOS C function tgeo_tpoint.
func TgeoTpoint(temp Temporal, oper bool) Temporal {
	res := C.tgeo_tpoint(temp.Inner(), C.bool(oper))
	return CreateTemporal(res)
}


// TspatialinstSetSRID wraps MEOS C function tspatialinst_set_srid.
func TspatialinstSetSRID(inst TInstant, srid int32) {
	C.tspatialinst_set_srid(inst.Inner(), C.int32_t(srid))
}


// TpointseqMakeSimple wraps MEOS C function tpointseq_make_simple.
func TpointseqMakeSimple(seq TSequence) []TSequence {
	var _out_count C.int
	res := C.tpointseq_make_simple(seq.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.TSequence)(unsafe.Pointer(res)), _n)
	_out := make([]TSequence, _n)
	for _i, _e := range _slice {
		_out[_i] = TSequence{_inner: _e}
	}
	return _out
}


// TspatialseqSetSRID wraps MEOS C function tspatialseq_set_srid.
func TspatialseqSetSRID(seq TSequence, srid int32) {
	C.tspatialseq_set_srid(seq.Inner(), C.int32_t(srid))
}


// TpointseqsetMakeSimple wraps MEOS C function tpointseqset_make_simple.
func TpointseqsetMakeSimple(ss TSequenceSet) []TSequence {
	var _out_count C.int
	res := C.tpointseqset_make_simple(ss.Inner(), &_out_count)
	_n := int(_out_count)
	_slice := unsafe.Slice((**C.TSequence)(unsafe.Pointer(res)), _n)
	_out := make([]TSequence, _n)
	for _i, _e := range _slice {
		_out[_i] = TSequence{_inner: _e}
	}
	return _out
}


// TspatialseqsetSetSRID wraps MEOS C function tspatialseqset_set_srid.
func TspatialseqsetSetSRID(ss TSequenceSet, srid int32) {
	C.tspatialseqset_set_srid(ss.Inner(), C.int32_t(srid))
}


// TODO tpointseq_twcentroid: unsupported return type int *
// func TpointseqTwcentroid(...) { /* not yet handled by codegen */ }


// TODO tpointseqset_twcentroid: unsupported return type int *
// func TpointseqsetTwcentroid(...) { /* not yet handled by codegen */ }

