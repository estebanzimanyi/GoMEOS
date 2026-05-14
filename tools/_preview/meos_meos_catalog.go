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


// TODO meostype_name: unsupported param MeosType
// func MeostypeName(...) { /* not yet handled by codegen */ }


// TODO temptype_basetype: unsupported return type MeosType
// func TemptypeBasetype(...) { /* not yet handled by codegen */ }


// TODO settype_basetype: unsupported return type MeosType
// func SettypeBasetype(...) { /* not yet handled by codegen */ }


// TODO spantype_basetype: unsupported return type MeosType
// func SpantypeBasetype(...) { /* not yet handled by codegen */ }


// TODO spantype_spansettype: unsupported return type MeosType
// func SpantypeSpansettype(...) { /* not yet handled by codegen */ }


// TODO spansettype_spantype: unsupported return type MeosType
// func SpansettypeSpantype(...) { /* not yet handled by codegen */ }


// TODO basetype_spantype: unsupported return type MeosType
// func BasetypeSpantype(...) { /* not yet handled by codegen */ }


// TODO basetype_settype: unsupported return type MeosType
// func BasetypeSettype(...) { /* not yet handled by codegen */ }


// TODO tnumber_basetype: unsupported param MeosType
// func TnumberBasetype(...) { /* not yet handled by codegen */ }


// TODO geo_basetype: unsupported param MeosType
// func GeoBasetype(...) { /* not yet handled by codegen */ }


// TODO meos_basetype: unsupported param MeosType
// func MeosBasetype(...) { /* not yet handled by codegen */ }


// TODO alphanum_basetype: unsupported param MeosType
// func AlphanumBasetype(...) { /* not yet handled by codegen */ }


// TODO alphanum_temptype: unsupported param MeosType
// func AlphanumTemptype(...) { /* not yet handled by codegen */ }


// TODO time_type: unsupported param MeosType
// func TimeType(...) { /* not yet handled by codegen */ }


// TODO set_basetype: unsupported param MeosType
// func SetBasetype(...) { /* not yet handled by codegen */ }


// TODO set_type: unsupported param MeosType
// func SetType(...) { /* not yet handled by codegen */ }


// TODO numset_type: unsupported param MeosType
// func NumsetType(...) { /* not yet handled by codegen */ }


// TODO ensure_numset_type: unsupported param MeosType
// func EnsureNumsetType(...) { /* not yet handled by codegen */ }


// TODO timeset_type: unsupported param MeosType
// func TimesetType(...) { /* not yet handled by codegen */ }


// TODO set_spantype: unsupported param MeosType
// func SetSpantype(...) { /* not yet handled by codegen */ }


// TODO ensure_set_spantype: unsupported param MeosType
// func EnsureSetSpantype(...) { /* not yet handled by codegen */ }


// TODO alphanumset_type: unsupported param MeosType
// func AlphanumsetType(...) { /* not yet handled by codegen */ }


// TODO geoset_type: unsupported param MeosType
// func GeosetType(...) { /* not yet handled by codegen */ }


// TODO ensure_geoset_type: unsupported param MeosType
// func EnsureGeosetType(...) { /* not yet handled by codegen */ }


// TODO spatialset_type: unsupported param MeosType
// func SpatialsetType(...) { /* not yet handled by codegen */ }


// TODO ensure_spatialset_type: unsupported param MeosType
// func EnsureSpatialsetType(...) { /* not yet handled by codegen */ }


// TODO span_basetype: unsupported param MeosType
// func SpanBasetype(...) { /* not yet handled by codegen */ }


// TODO span_canon_basetype: unsupported param MeosType
// func SpanCanonBasetype(...) { /* not yet handled by codegen */ }


// TODO span_type: unsupported param MeosType
// func SpanType(...) { /* not yet handled by codegen */ }


// TODO type_span_bbox: unsupported param MeosType
// func TypeSpanBbox(...) { /* not yet handled by codegen */ }


// TODO span_tbox_type: unsupported param MeosType
// func SpanTBOXType(...) { /* not yet handled by codegen */ }


// TODO ensure_span_tbox_type: unsupported param MeosType
// func EnsureSpanTBOXType(...) { /* not yet handled by codegen */ }


// TODO numspan_basetype: unsupported param MeosType
// func NumspanBasetype(...) { /* not yet handled by codegen */ }


// TODO numspan_type: unsupported param MeosType
// func NumspanType(...) { /* not yet handled by codegen */ }


// TODO ensure_numspan_type: unsupported param MeosType
// func EnsureNumspanType(...) { /* not yet handled by codegen */ }


// TODO timespan_basetype: unsupported param MeosType
// func TimespanBasetype(...) { /* not yet handled by codegen */ }


// TODO timespan_type: unsupported param MeosType
// func TimespanType(...) { /* not yet handled by codegen */ }


// TODO spanset_type: unsupported param MeosType
// func SpansetType(...) { /* not yet handled by codegen */ }


// TODO timespanset_type: unsupported param MeosType
// func TimespansetType(...) { /* not yet handled by codegen */ }


// TODO ensure_timespanset_type: unsupported param MeosType
// func EnsureTimespansetType(...) { /* not yet handled by codegen */ }


// TODO temporal_type: unsupported param MeosType
// func TemporalType(...) { /* not yet handled by codegen */ }


// TODO temporal_basetype: unsupported param MeosType
// func TemporalBasetype(...) { /* not yet handled by codegen */ }


// TODO temptype_supports_linear: unsupported param MeosType
// func TemptypeSupportsLinear(...) { /* not yet handled by codegen */ }


// TODO basetype_byvalue: unsupported param MeosType
// func BasetypeByvalue(...) { /* not yet handled by codegen */ }


// TODO basetype_varlength: unsupported param MeosType
// func BasetypeVarlength(...) { /* not yet handled by codegen */ }


// TODO meostype_length: unsupported param MeosType
// func MeostypeLength(...) { /* not yet handled by codegen */ }


// TODO talphanum_type: unsupported param MeosType
// func TalphanumType(...) { /* not yet handled by codegen */ }


// TODO talpha_type: unsupported param MeosType
// func TalphaType(...) { /* not yet handled by codegen */ }


// TODO tnumber_type: unsupported param MeosType
// func TnumberType(...) { /* not yet handled by codegen */ }


// TODO ensure_tnumber_type: unsupported param MeosType
// func EnsureTnumberType(...) { /* not yet handled by codegen */ }


// TODO ensure_tnumber_basetype: unsupported param MeosType
// func EnsureTnumberBasetype(...) { /* not yet handled by codegen */ }


// TODO tnumber_spantype: unsupported param MeosType
// func TnumberSpantype(...) { /* not yet handled by codegen */ }


// TODO spatial_basetype: unsupported param MeosType
// func SpatialBasetype(...) { /* not yet handled by codegen */ }


// TODO tspatial_type: unsupported param MeosType
// func TspatialType(...) { /* not yet handled by codegen */ }


// TODO ensure_tspatial_type: unsupported param MeosType
// func EnsureTspatialType(...) { /* not yet handled by codegen */ }


// TODO tpoint_type: unsupported param MeosType
// func TpointType(...) { /* not yet handled by codegen */ }


// TODO ensure_tpoint_type: unsupported param MeosType
// func EnsureTpointType(...) { /* not yet handled by codegen */ }


// TODO tgeo_type: unsupported param MeosType
// func TgeoType(...) { /* not yet handled by codegen */ }


// TODO ensure_tgeo_type: unsupported param MeosType
// func EnsureTgeoType(...) { /* not yet handled by codegen */ }


// TODO tgeo_type_all: unsupported param MeosType
// func TgeoTypeAll(...) { /* not yet handled by codegen */ }


// TODO ensure_tgeo_type_all: unsupported param MeosType
// func EnsureTgeoTypeAll(...) { /* not yet handled by codegen */ }


// TODO tgeometry_type: unsupported param MeosType
// func TgeometryType(...) { /* not yet handled by codegen */ }


// TODO ensure_tgeometry_type: unsupported param MeosType
// func EnsureTgeometryType(...) { /* not yet handled by codegen */ }


// TODO tgeodetic_type: unsupported param MeosType
// func TgeodeticType(...) { /* not yet handled by codegen */ }


// TODO ensure_tgeodetic_type: unsupported param MeosType
// func EnsureTgeodeticType(...) { /* not yet handled by codegen */ }


// TODO ensure_tnumber_tpoint_type: unsupported param MeosType
// func EnsureTnumberTpointType(...) { /* not yet handled by codegen */ }

