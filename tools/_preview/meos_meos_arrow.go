package generated

// #include <stddef.h>
import "C"
import (
	"unsafe"

	"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

// TODO meos_temporal_to_arrow: unsupported param struct ArrowSchema *
// func MeosTemporalToArrow(...) { /* not yet handled by codegen */ }


// TODO meos_temporal_from_arrow: unsupported param const struct ArrowSchema *
// func MeosTemporalFromArrow(...) { /* not yet handled by codegen */ }


// MeosTemporalArrowRoundtrip wraps MEOS C function meos_temporal_arrow_roundtrip.
func MeosTemporalArrowRoundtrip(temp Temporal) Temporal {
	res := C.meos_temporal_arrow_roundtrip(temp.Inner())
	return CreateTemporal(res)
}


// TODO meos_set_to_arrow: unsupported param struct ArrowSchema *
// func MeosSetToArrow(...) { /* not yet handled by codegen */ }


// TODO meos_set_from_arrow: unsupported param const struct ArrowSchema *
// func MeosSetFromArrow(...) { /* not yet handled by codegen */ }


// MeosSetArrowRoundtrip wraps MEOS C function meos_set_arrow_roundtrip.
func MeosSetArrowRoundtrip(s *Set) *Set {
	res := C.meos_set_arrow_roundtrip(s._inner)
	return &Set{_inner: res}
}


// TODO meos_span_to_arrow: unsupported param struct ArrowSchema *
// func MeosSpanToArrow(...) { /* not yet handled by codegen */ }


// TODO meos_span_from_arrow: unsupported param const struct ArrowSchema *
// func MeosSpanFromArrow(...) { /* not yet handled by codegen */ }


// MeosSpanArrowRoundtrip wraps MEOS C function meos_span_arrow_roundtrip.
func MeosSpanArrowRoundtrip(s *Span) *Span {
	res := C.meos_span_arrow_roundtrip(s._inner)
	return &Span{_inner: res}
}


// TODO meos_spanset_to_arrow: unsupported param struct ArrowSchema *
// func MeosSpansetToArrow(...) { /* not yet handled by codegen */ }


// TODO meos_spanset_from_arrow: unsupported param const struct ArrowSchema *
// func MeosSpansetFromArrow(...) { /* not yet handled by codegen */ }


// MeosSpansetArrowRoundtrip wraps MEOS C function meos_spanset_arrow_roundtrip.
func MeosSpansetArrowRoundtrip(ss *SpanSet) *SpanSet {
	res := C.meos_spanset_arrow_roundtrip(ss._inner)
	return &SpanSet{_inner: res}
}


// TODO meos_tbox_to_arrow: unsupported param struct ArrowSchema *
// func MeosTBOXToArrow(...) { /* not yet handled by codegen */ }


// TODO meos_tbox_from_arrow: unsupported param const struct ArrowSchema *
// func MeosTBOXFromArrow(...) { /* not yet handled by codegen */ }


// MeosTBOXArrowRoundtrip wraps MEOS C function meos_tbox_arrow_roundtrip.
func MeosTBOXArrowRoundtrip(box *TBox) *TBox {
	res := C.meos_tbox_arrow_roundtrip(box._inner)
	return &TBox{_inner: res}
}


// TODO meos_stbox_to_arrow: unsupported param struct ArrowSchema *
// func MeosSTBOXToArrow(...) { /* not yet handled by codegen */ }


// TODO meos_stbox_from_arrow: unsupported param const struct ArrowSchema *
// func MeosSTBOXFromArrow(...) { /* not yet handled by codegen */ }


// MeosSTBOXArrowRoundtrip wraps MEOS C function meos_stbox_arrow_roundtrip.
func MeosSTBOXArrowRoundtrip(box *STBox) *STBox {
	res := C.meos_stbox_arrow_roundtrip(box._inner)
	return &STBox{_inner: res}
}

