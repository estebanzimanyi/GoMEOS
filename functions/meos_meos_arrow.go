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

// MeosTemporalToArrow wraps MEOS C function meos_temporal_to_arrow.
func MeosTemporalToArrow(temp *Temporal, out_schema unsafe.Pointer, out_array unsafe.Pointer) bool {
	_cret := C.meos_temporal_to_arrow(temp._inner, (*C.struct_ArrowSchema)(out_schema), (*C.struct_ArrowArray)(out_array))
	return bool(_cret)
}


// MeosTemporalFromArrow wraps MEOS C function meos_temporal_from_arrow.
func MeosTemporalFromArrow(schema unsafe.Pointer, array unsafe.Pointer) *Temporal {
	_cret := C.meos_temporal_from_arrow((*C.struct_ArrowSchema)(schema), (*C.struct_ArrowArray)(array))
	return &Temporal{_inner: _cret}
}


// MeosTemporalArrowRoundtrip wraps MEOS C function meos_temporal_arrow_roundtrip.
func MeosTemporalArrowRoundtrip(temp *Temporal) *Temporal {
	_cret := C.meos_temporal_arrow_roundtrip(temp._inner)
	return &Temporal{_inner: _cret}
}


// MeosSetToArrow wraps MEOS C function meos_set_to_arrow.
func MeosSetToArrow(s *Set, out_schema unsafe.Pointer, out_array unsafe.Pointer) bool {
	_cret := C.meos_set_to_arrow(s._inner, (*C.struct_ArrowSchema)(out_schema), (*C.struct_ArrowArray)(out_array))
	return bool(_cret)
}


// MeosSetFromArrow wraps MEOS C function meos_set_from_arrow.
func MeosSetFromArrow(schema unsafe.Pointer, array unsafe.Pointer) *Set {
	_cret := C.meos_set_from_arrow((*C.struct_ArrowSchema)(schema), (*C.struct_ArrowArray)(array))
	return &Set{_inner: _cret}
}


// MeosSetArrowRoundtrip wraps MEOS C function meos_set_arrow_roundtrip.
func MeosSetArrowRoundtrip(s *Set) *Set {
	_cret := C.meos_set_arrow_roundtrip(s._inner)
	return &Set{_inner: _cret}
}


// MeosSpanToArrow wraps MEOS C function meos_span_to_arrow.
func MeosSpanToArrow(s *Span, out_schema unsafe.Pointer, out_array unsafe.Pointer) bool {
	_cret := C.meos_span_to_arrow(s._inner, (*C.struct_ArrowSchema)(out_schema), (*C.struct_ArrowArray)(out_array))
	return bool(_cret)
}


// MeosSpanFromArrow wraps MEOS C function meos_span_from_arrow.
func MeosSpanFromArrow(schema unsafe.Pointer, array unsafe.Pointer) *Span {
	_cret := C.meos_span_from_arrow((*C.struct_ArrowSchema)(schema), (*C.struct_ArrowArray)(array))
	return &Span{_inner: _cret}
}


// MeosSpanArrowRoundtrip wraps MEOS C function meos_span_arrow_roundtrip.
func MeosSpanArrowRoundtrip(s *Span) *Span {
	_cret := C.meos_span_arrow_roundtrip(s._inner)
	return &Span{_inner: _cret}
}


// MeosSpansetToArrow wraps MEOS C function meos_spanset_to_arrow.
func MeosSpansetToArrow(ss *SpanSet, out_schema unsafe.Pointer, out_array unsafe.Pointer) bool {
	_cret := C.meos_spanset_to_arrow(ss._inner, (*C.struct_ArrowSchema)(out_schema), (*C.struct_ArrowArray)(out_array))
	return bool(_cret)
}


// MeosSpansetFromArrow wraps MEOS C function meos_spanset_from_arrow.
func MeosSpansetFromArrow(schema unsafe.Pointer, array unsafe.Pointer) *SpanSet {
	_cret := C.meos_spanset_from_arrow((*C.struct_ArrowSchema)(schema), (*C.struct_ArrowArray)(array))
	return &SpanSet{_inner: _cret}
}


// MeosSpansetArrowRoundtrip wraps MEOS C function meos_spanset_arrow_roundtrip.
func MeosSpansetArrowRoundtrip(ss *SpanSet) *SpanSet {
	_cret := C.meos_spanset_arrow_roundtrip(ss._inner)
	return &SpanSet{_inner: _cret}
}


// MeosTBOXToArrow wraps MEOS C function meos_tbox_to_arrow.
func MeosTBOXToArrow(box *TBox, out_schema unsafe.Pointer, out_array unsafe.Pointer) bool {
	_cret := C.meos_tbox_to_arrow(box._inner, (*C.struct_ArrowSchema)(out_schema), (*C.struct_ArrowArray)(out_array))
	return bool(_cret)
}


// MeosTBOXFromArrow wraps MEOS C function meos_tbox_from_arrow.
func MeosTBOXFromArrow(schema unsafe.Pointer, array unsafe.Pointer) *TBox {
	_cret := C.meos_tbox_from_arrow((*C.struct_ArrowSchema)(schema), (*C.struct_ArrowArray)(array))
	return &TBox{_inner: _cret}
}


// MeosTBOXArrowRoundtrip wraps MEOS C function meos_tbox_arrow_roundtrip.
func MeosTBOXArrowRoundtrip(box *TBox) *TBox {
	_cret := C.meos_tbox_arrow_roundtrip(box._inner)
	return &TBox{_inner: _cret}
}


// MeosSTBOXToArrow wraps MEOS C function meos_stbox_to_arrow.
func MeosSTBOXToArrow(box *STBox, out_schema unsafe.Pointer, out_array unsafe.Pointer) bool {
	_cret := C.meos_stbox_to_arrow(box._inner, (*C.struct_ArrowSchema)(out_schema), (*C.struct_ArrowArray)(out_array))
	return bool(_cret)
}


// MeosSTBOXFromArrow wraps MEOS C function meos_stbox_from_arrow.
func MeosSTBOXFromArrow(schema unsafe.Pointer, array unsafe.Pointer) *STBox {
	_cret := C.meos_stbox_from_arrow((*C.struct_ArrowSchema)(schema), (*C.struct_ArrowArray)(array))
	return &STBox{_inner: _cret}
}


// MeosSTBOXArrowRoundtrip wraps MEOS C function meos_stbox_arrow_roundtrip.
func MeosSTBOXArrowRoundtrip(box *STBox) *STBox {
	_cret := C.meos_stbox_arrow_roundtrip(box._inner)
	return &STBox{_inner: _cret}
}

