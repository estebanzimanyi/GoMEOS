package generated

// #include <stddef.h>
import "C"
import (
	"unsafe"

	"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

// TemptypeSubtype wraps MEOS C function temptype_subtype.
func TemptypeSubtype(subtype TempSubtype) bool {
	res := C.temptype_subtype(C.tempSubtype(subtype))
	return bool(res)
}


// TemptypeSubtypeAll wraps MEOS C function temptype_subtype_all.
func TemptypeSubtypeAll(subtype TempSubtype) bool {
	res := C.temptype_subtype_all(C.tempSubtype(subtype))
	return bool(res)
}


// TempsubtypeName wraps MEOS C function tempsubtype_name.
func TempsubtypeName(subtype TempSubtype) string {
	res := C.tempsubtype_name(C.tempSubtype(subtype))
	return C.GoString(res)
}


// TempsubtypeFromString wraps MEOS C function tempsubtype_from_string.
func TempsubtypeFromString(str string) (bool, int16) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	var _out_subtype C.int16
	res := C.tempsubtype_from_string(_c_str, &_out_subtype)
	return bool(res), int16(_out_subtype)
}


// MeosoperName wraps MEOS C function meosoper_name.
func MeosoperName(oper MeosOper) string {
	res := C.meosoper_name(C.meosOper(oper))
	return C.GoString(res)
}


// MeosoperFromString wraps MEOS C function meosoper_from_string.
func MeosoperFromString(name string) MeosOper {
	_c_name := C.CString(name)
	defer C.free(unsafe.Pointer(_c_name))
	res := C.meosoper_from_string(_c_name)
	return MeosOper(res)
}


// InterptypeName wraps MEOS C function interptype_name.
func InterptypeName(interp Interpolation) string {
	res := C.interptype_name(C.interpType(interp))
	return C.GoString(res)
}


// InterptypeFromString wraps MEOS C function interptype_from_string.
func InterptypeFromString(interp_str string) Interpolation {
	_c_interp_str := C.CString(interp_str)
	defer C.free(unsafe.Pointer(_c_interp_str))
	res := C.interptype_from_string(_c_interp_str)
	return Interpolation(res)
}


// MeostypeName wraps MEOS C function meostype_name.
func MeostypeName(type_ MeosType) string {
	res := C.meostype_name(C.meosType(type_))
	return C.GoString(res)
}


// TemptypeBasetype wraps MEOS C function temptype_basetype.
func TemptypeBasetype(type_ MeosType) MeosType {
	res := C.temptype_basetype(C.meosType(type_))
	return MeosType(res)
}


// SettypeBasetype wraps MEOS C function settype_basetype.
func SettypeBasetype(type_ MeosType) MeosType {
	res := C.settype_basetype(C.meosType(type_))
	return MeosType(res)
}


// SpantypeBasetype wraps MEOS C function spantype_basetype.
func SpantypeBasetype(type_ MeosType) MeosType {
	res := C.spantype_basetype(C.meosType(type_))
	return MeosType(res)
}


// SpantypeSpansettype wraps MEOS C function spantype_spansettype.
func SpantypeSpansettype(type_ MeosType) MeosType {
	res := C.spantype_spansettype(C.meosType(type_))
	return MeosType(res)
}


// SpansettypeSpantype wraps MEOS C function spansettype_spantype.
func SpansettypeSpantype(type_ MeosType) MeosType {
	res := C.spansettype_spantype(C.meosType(type_))
	return MeosType(res)
}


// BasetypeSpantype wraps MEOS C function basetype_spantype.
func BasetypeSpantype(type_ MeosType) MeosType {
	res := C.basetype_spantype(C.meosType(type_))
	return MeosType(res)
}


// BasetypeSettype wraps MEOS C function basetype_settype.
func BasetypeSettype(type_ MeosType) MeosType {
	res := C.basetype_settype(C.meosType(type_))
	return MeosType(res)
}


// TnumberBasetype wraps MEOS C function tnumber_basetype.
func TnumberBasetype(type_ MeosType) bool {
	res := C.tnumber_basetype(C.meosType(type_))
	return bool(res)
}


// GeoBasetype wraps MEOS C function geo_basetype.
func GeoBasetype(type_ MeosType) bool {
	res := C.geo_basetype(C.meosType(type_))
	return bool(res)
}


// MeosBasetype wraps MEOS C function meos_basetype.
func MeosBasetype(type_ MeosType) bool {
	res := C.meos_basetype(C.meosType(type_))
	return bool(res)
}


// AlphanumBasetype wraps MEOS C function alphanum_basetype.
func AlphanumBasetype(type_ MeosType) bool {
	res := C.alphanum_basetype(C.meosType(type_))
	return bool(res)
}


// AlphanumTemptype wraps MEOS C function alphanum_temptype.
func AlphanumTemptype(type_ MeosType) bool {
	res := C.alphanum_temptype(C.meosType(type_))
	return bool(res)
}


// TimeType wraps MEOS C function time_type.
func TimeType(type_ MeosType) bool {
	res := C.time_type(C.meosType(type_))
	return bool(res)
}


// SetBasetype wraps MEOS C function set_basetype.
func SetBasetype(type_ MeosType) bool {
	res := C.set_basetype(C.meosType(type_))
	return bool(res)
}


// SetType wraps MEOS C function set_type.
func SetType(type_ MeosType) bool {
	res := C.set_type(C.meosType(type_))
	return bool(res)
}


// NumsetType wraps MEOS C function numset_type.
func NumsetType(type_ MeosType) bool {
	res := C.numset_type(C.meosType(type_))
	return bool(res)
}


// EnsureNumsetType wraps MEOS C function ensure_numset_type.
func EnsureNumsetType(type_ MeosType) bool {
	res := C.ensure_numset_type(C.meosType(type_))
	return bool(res)
}


// TimesetType wraps MEOS C function timeset_type.
func TimesetType(type_ MeosType) bool {
	res := C.timeset_type(C.meosType(type_))
	return bool(res)
}


// SetSpantype wraps MEOS C function set_spantype.
func SetSpantype(type_ MeosType) bool {
	res := C.set_spantype(C.meosType(type_))
	return bool(res)
}


// EnsureSetSpantype wraps MEOS C function ensure_set_spantype.
func EnsureSetSpantype(type_ MeosType) bool {
	res := C.ensure_set_spantype(C.meosType(type_))
	return bool(res)
}


// AlphanumsetType wraps MEOS C function alphanumset_type.
func AlphanumsetType(settype MeosType) bool {
	res := C.alphanumset_type(C.meosType(settype))
	return bool(res)
}


// GeosetType wraps MEOS C function geoset_type.
func GeosetType(type_ MeosType) bool {
	res := C.geoset_type(C.meosType(type_))
	return bool(res)
}


// EnsureGeosetType wraps MEOS C function ensure_geoset_type.
func EnsureGeosetType(type_ MeosType) bool {
	res := C.ensure_geoset_type(C.meosType(type_))
	return bool(res)
}


// SpatialsetType wraps MEOS C function spatialset_type.
func SpatialsetType(type_ MeosType) bool {
	res := C.spatialset_type(C.meosType(type_))
	return bool(res)
}


// EnsureSpatialsetType wraps MEOS C function ensure_spatialset_type.
func EnsureSpatialsetType(type_ MeosType) bool {
	res := C.ensure_spatialset_type(C.meosType(type_))
	return bool(res)
}


// SpanBasetype wraps MEOS C function span_basetype.
func SpanBasetype(type_ MeosType) bool {
	res := C.span_basetype(C.meosType(type_))
	return bool(res)
}


// SpanCanonBasetype wraps MEOS C function span_canon_basetype.
func SpanCanonBasetype(type_ MeosType) bool {
	res := C.span_canon_basetype(C.meosType(type_))
	return bool(res)
}


// SpanType wraps MEOS C function span_type.
func SpanType(type_ MeosType) bool {
	res := C.span_type(C.meosType(type_))
	return bool(res)
}


// TypeSpanBbox wraps MEOS C function type_span_bbox.
func TypeSpanBbox(type_ MeosType) bool {
	res := C.type_span_bbox(C.meosType(type_))
	return bool(res)
}


// SpanTBOXType wraps MEOS C function span_tbox_type.
func SpanTBOXType(type_ MeosType) bool {
	res := C.span_tbox_type(C.meosType(type_))
	return bool(res)
}


// EnsureSpanTBOXType wraps MEOS C function ensure_span_tbox_type.
func EnsureSpanTBOXType(type_ MeosType) bool {
	res := C.ensure_span_tbox_type(C.meosType(type_))
	return bool(res)
}


// NumspanBasetype wraps MEOS C function numspan_basetype.
func NumspanBasetype(type_ MeosType) bool {
	res := C.numspan_basetype(C.meosType(type_))
	return bool(res)
}


// NumspanType wraps MEOS C function numspan_type.
func NumspanType(type_ MeosType) bool {
	res := C.numspan_type(C.meosType(type_))
	return bool(res)
}


// EnsureNumspanType wraps MEOS C function ensure_numspan_type.
func EnsureNumspanType(type_ MeosType) bool {
	res := C.ensure_numspan_type(C.meosType(type_))
	return bool(res)
}


// TimespanBasetype wraps MEOS C function timespan_basetype.
func TimespanBasetype(type_ MeosType) bool {
	res := C.timespan_basetype(C.meosType(type_))
	return bool(res)
}


// TimespanType wraps MEOS C function timespan_type.
func TimespanType(type_ MeosType) bool {
	res := C.timespan_type(C.meosType(type_))
	return bool(res)
}


// SpansetType wraps MEOS C function spanset_type.
func SpansetType(type_ MeosType) bool {
	res := C.spanset_type(C.meosType(type_))
	return bool(res)
}


// TimespansetType wraps MEOS C function timespanset_type.
func TimespansetType(type_ MeosType) bool {
	res := C.timespanset_type(C.meosType(type_))
	return bool(res)
}


// EnsureTimespansetType wraps MEOS C function ensure_timespanset_type.
func EnsureTimespansetType(type_ MeosType) bool {
	res := C.ensure_timespanset_type(C.meosType(type_))
	return bool(res)
}


// TemporalType wraps MEOS C function temporal_type.
func TemporalType(type_ MeosType) bool {
	res := C.temporal_type(C.meosType(type_))
	return bool(res)
}


// TemporalBasetype wraps MEOS C function temporal_basetype.
func TemporalBasetype(type_ MeosType) bool {
	res := C.temporal_basetype(C.meosType(type_))
	return bool(res)
}


// TemptypeContinuous wraps MEOS C function temptype_continuous.
func TemptypeContinuous(type_ MeosType) bool {
	res := C.temptype_continuous(C.meosType(type_))
	return bool(res)
}


// BasetypeByvalue wraps MEOS C function basetype_byvalue.
func BasetypeByvalue(type_ MeosType) bool {
	res := C.basetype_byvalue(C.meosType(type_))
	return bool(res)
}


// BasetypeVarlength wraps MEOS C function basetype_varlength.
func BasetypeVarlength(type_ MeosType) bool {
	res := C.basetype_varlength(C.meosType(type_))
	return bool(res)
}


// BasetypeLength wraps MEOS C function basetype_length.
func BasetypeLength(type_ MeosType) int16 {
	res := C.basetype_length(C.meosType(type_))
	return int16(res)
}


// TalphanumType wraps MEOS C function talphanum_type.
func TalphanumType(type_ MeosType) bool {
	res := C.talphanum_type(C.meosType(type_))
	return bool(res)
}


// TalphaType wraps MEOS C function talpha_type.
func TalphaType(type_ MeosType) bool {
	res := C.talpha_type(C.meosType(type_))
	return bool(res)
}


// TnumberType wraps MEOS C function tnumber_type.
func TnumberType(type_ MeosType) bool {
	res := C.tnumber_type(C.meosType(type_))
	return bool(res)
}


// EnsureTnumberType wraps MEOS C function ensure_tnumber_type.
func EnsureTnumberType(type_ MeosType) bool {
	res := C.ensure_tnumber_type(C.meosType(type_))
	return bool(res)
}


// EnsureTnumberBasetype wraps MEOS C function ensure_tnumber_basetype.
func EnsureTnumberBasetype(type_ MeosType) bool {
	res := C.ensure_tnumber_basetype(C.meosType(type_))
	return bool(res)
}


// TnumberSpantype wraps MEOS C function tnumber_spantype.
func TnumberSpantype(type_ MeosType) bool {
	res := C.tnumber_spantype(C.meosType(type_))
	return bool(res)
}


// SpatialBasetype wraps MEOS C function spatial_basetype.
func SpatialBasetype(type_ MeosType) bool {
	res := C.spatial_basetype(C.meosType(type_))
	return bool(res)
}


// TspatialType wraps MEOS C function tspatial_type.
func TspatialType(type_ MeosType) bool {
	res := C.tspatial_type(C.meosType(type_))
	return bool(res)
}


// EnsureTspatialType wraps MEOS C function ensure_tspatial_type.
func EnsureTspatialType(type_ MeosType) bool {
	res := C.ensure_tspatial_type(C.meosType(type_))
	return bool(res)
}


// TpointType wraps MEOS C function tpoint_type.
func TpointType(type_ MeosType) bool {
	res := C.tpoint_type(C.meosType(type_))
	return bool(res)
}


// EnsureTpointType wraps MEOS C function ensure_tpoint_type.
func EnsureTpointType(type_ MeosType) bool {
	res := C.ensure_tpoint_type(C.meosType(type_))
	return bool(res)
}


// TgeoType wraps MEOS C function tgeo_type.
func TgeoType(type_ MeosType) bool {
	res := C.tgeo_type(C.meosType(type_))
	return bool(res)
}


// EnsureTgeoType wraps MEOS C function ensure_tgeo_type.
func EnsureTgeoType(type_ MeosType) bool {
	res := C.ensure_tgeo_type(C.meosType(type_))
	return bool(res)
}


// TgeoTypeAll wraps MEOS C function tgeo_type_all.
func TgeoTypeAll(type_ MeosType) bool {
	res := C.tgeo_type_all(C.meosType(type_))
	return bool(res)
}


// EnsureTgeoTypeAll wraps MEOS C function ensure_tgeo_type_all.
func EnsureTgeoTypeAll(type_ MeosType) bool {
	res := C.ensure_tgeo_type_all(C.meosType(type_))
	return bool(res)
}


// TgeometryType wraps MEOS C function tgeometry_type.
func TgeometryType(type_ MeosType) bool {
	res := C.tgeometry_type(C.meosType(type_))
	return bool(res)
}


// EnsureTgeometryType wraps MEOS C function ensure_tgeometry_type.
func EnsureTgeometryType(type_ MeosType) bool {
	res := C.ensure_tgeometry_type(C.meosType(type_))
	return bool(res)
}


// TgeodeticType wraps MEOS C function tgeodetic_type.
func TgeodeticType(type_ MeosType) bool {
	res := C.tgeodetic_type(C.meosType(type_))
	return bool(res)
}


// EnsureTgeodeticType wraps MEOS C function ensure_tgeodetic_type.
func EnsureTgeodeticType(type_ MeosType) bool {
	res := C.ensure_tgeodetic_type(C.meosType(type_))
	return bool(res)
}


// EnsureTnumberTpointType wraps MEOS C function ensure_tnumber_tpoint_type.
func EnsureTnumberTpointType(type_ MeosType) bool {
	res := C.ensure_tnumber_tpoint_type(C.meosType(type_))
	return bool(res)
}

