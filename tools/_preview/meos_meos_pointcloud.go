package generated

// #include <stddef.h>
import "C"
import (
	"unsafe"

	"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

// PcpointHexIn wraps MEOS C function pcpoint_hex_in.
func PcpointHexIn(str string) *Pcpoint {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.pcpoint_hex_in(_c_str)
	return &Pcpoint{_inner: res}
}


// PcpointHexOut wraps MEOS C function pcpoint_hex_out.
func PcpointHexOut(pt *Pcpoint, maxdd int) string {
	res := C.pcpoint_hex_out(pt._inner, C.int(maxdd))
	return C.GoString(res)
}


// PcpointFromHexwkb wraps MEOS C function pcpoint_from_hexwkb.
func PcpointFromHexwkb(hexwkb string) *Pcpoint {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	res := C.pcpoint_from_hexwkb(_c_hexwkb)
	return &Pcpoint{_inner: res}
}


// PcpointAsHexwkb wraps MEOS C function pcpoint_as_hexwkb.
func PcpointAsHexwkb(pt *Pcpoint) string {
	res := C.pcpoint_as_hexwkb(pt._inner)
	return C.GoString(res)
}


// PcpointCopy wraps MEOS C function pcpoint_copy.
func PcpointCopy(pt *Pcpoint) *Pcpoint {
	res := C.pcpoint_copy(pt._inner)
	return &Pcpoint{_inner: res}
}


// PcpointGetPcid wraps MEOS C function pcpoint_get_pcid.
func PcpointGetPcid(pt *Pcpoint) uint32 {
	res := C.pcpoint_get_pcid(pt._inner)
	return uint32(res)
}


// PcpointHash wraps MEOS C function pcpoint_hash.
func PcpointHash(pt *Pcpoint) int {
	res := C.pcpoint_hash(pt._inner)
	return int(res)
}


// PcpointHashExtended wraps MEOS C function pcpoint_hash_extended.
func PcpointHashExtended(pt *Pcpoint, seed int) int {
	res := C.pcpoint_hash_extended(pt._inner, C.int(seed))
	return int(res)
}


// TODO pcpoint_get_x: unsupported param double *
// func PcpointGetX(...) { /* not yet handled by codegen */ }


// TODO pcpoint_get_y: unsupported param double *
// func PcpointGetY(...) { /* not yet handled by codegen */ }


// TODO pcpoint_get_z: unsupported param double *
// func PcpointGetZ(...) { /* not yet handled by codegen */ }


// TODO pcpoint_get_dim: unsupported param double *
// func PcpointGetDim(...) { /* not yet handled by codegen */ }


// PcpointToTpcbox wraps MEOS C function pcpoint_to_tpcbox.
func PcpointToTpcbox(pt *Pcpoint, schema *PCSchema) *TPCBox {
	res := C.pcpoint_to_tpcbox(pt._inner, schema._inner)
	return &TPCBox{_inner: res}
}


// MeosPcSchema wraps MEOS C function meos_pc_schema.
func MeosPcSchema(pcid uint32) *PCSchema {
	res := C.meos_pc_schema(C.uint32_t(pcid))
	return &PCSchema{_inner: res}
}


// MeosPcSchemaRegister wraps MEOS C function meos_pc_schema_register.
func MeosPcSchemaRegister(pcid uint32, schema *PCSchema) {
	C.meos_pc_schema_register(C.uint32_t(pcid), schema._inner)
}


// MeosPcSchemaRegisterXml wraps MEOS C function meos_pc_schema_register_xml.
func MeosPcSchemaRegisterXml(pcid uint32, schema *PCSchema, xml_text string) {
	_c_xml_text := C.CString(xml_text)
	defer C.free(unsafe.Pointer(_c_xml_text))
	C.meos_pc_schema_register_xml(C.uint32_t(pcid), schema._inner, _c_xml_text)
}


// MeosPcSchemaXml wraps MEOS C function meos_pc_schema_xml.
func MeosPcSchemaXml(pcid uint32) string {
	res := C.meos_pc_schema_xml(C.uint32_t(pcid))
	return C.GoString(res)
}


// MeosPcSchemaClear wraps MEOS C function meos_pc_schema_clear.
func MeosPcSchemaClear() {
	C.meos_pc_schema_clear()
}


// PcpointCmp wraps MEOS C function pcpoint_cmp.
func PcpointCmp(pt1 *Pcpoint, pt2 *Pcpoint) int {
	res := C.pcpoint_cmp(pt1._inner, pt2._inner)
	return int(res)
}


// PcpointEq wraps MEOS C function pcpoint_eq.
func PcpointEq(pt1 *Pcpoint, pt2 *Pcpoint) bool {
	res := C.pcpoint_eq(pt1._inner, pt2._inner)
	return bool(res)
}


// PcpointNe wraps MEOS C function pcpoint_ne.
func PcpointNe(pt1 *Pcpoint, pt2 *Pcpoint) bool {
	res := C.pcpoint_ne(pt1._inner, pt2._inner)
	return bool(res)
}


// PcpointLt wraps MEOS C function pcpoint_lt.
func PcpointLt(pt1 *Pcpoint, pt2 *Pcpoint) bool {
	res := C.pcpoint_lt(pt1._inner, pt2._inner)
	return bool(res)
}


// PcpointLe wraps MEOS C function pcpoint_le.
func PcpointLe(pt1 *Pcpoint, pt2 *Pcpoint) bool {
	res := C.pcpoint_le(pt1._inner, pt2._inner)
	return bool(res)
}


// PcpointGt wraps MEOS C function pcpoint_gt.
func PcpointGt(pt1 *Pcpoint, pt2 *Pcpoint) bool {
	res := C.pcpoint_gt(pt1._inner, pt2._inner)
	return bool(res)
}


// PcpointGe wraps MEOS C function pcpoint_ge.
func PcpointGe(pt1 *Pcpoint, pt2 *Pcpoint) bool {
	res := C.pcpoint_ge(pt1._inner, pt2._inner)
	return bool(res)
}


// PcpatchHexIn wraps MEOS C function pcpatch_hex_in.
func PcpatchHexIn(str string) *Pcpatch {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.pcpatch_hex_in(_c_str)
	return &Pcpatch{_inner: res}
}


// PcpatchHexOut wraps MEOS C function pcpatch_hex_out.
func PcpatchHexOut(pa *Pcpatch, maxdd int) string {
	res := C.pcpatch_hex_out(pa._inner, C.int(maxdd))
	return C.GoString(res)
}


// PcpatchFromHexwkb wraps MEOS C function pcpatch_from_hexwkb.
func PcpatchFromHexwkb(hexwkb string) *Pcpatch {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	res := C.pcpatch_from_hexwkb(_c_hexwkb)
	return &Pcpatch{_inner: res}
}


// PcpatchAsHexwkb wraps MEOS C function pcpatch_as_hexwkb.
func PcpatchAsHexwkb(pa *Pcpatch) string {
	res := C.pcpatch_as_hexwkb(pa._inner)
	return C.GoString(res)
}


// PcpatchCopy wraps MEOS C function pcpatch_copy.
func PcpatchCopy(pa *Pcpatch) *Pcpatch {
	res := C.pcpatch_copy(pa._inner)
	return &Pcpatch{_inner: res}
}


// PcpatchGetPcid wraps MEOS C function pcpatch_get_pcid.
func PcpatchGetPcid(pa *Pcpatch) uint32 {
	res := C.pcpatch_get_pcid(pa._inner)
	return uint32(res)
}


// PcpatchNpoints wraps MEOS C function pcpatch_npoints.
func PcpatchNpoints(pa *Pcpatch) uint32 {
	res := C.pcpatch_npoints(pa._inner)
	return uint32(res)
}


// PcpatchHash wraps MEOS C function pcpatch_hash.
func PcpatchHash(pa *Pcpatch) int {
	res := C.pcpatch_hash(pa._inner)
	return int(res)
}


// PcpatchHashExtended wraps MEOS C function pcpatch_hash_extended.
func PcpatchHashExtended(pa *Pcpatch, seed int) int {
	res := C.pcpatch_hash_extended(pa._inner, C.int(seed))
	return int(res)
}


// PcpatchCmp wraps MEOS C function pcpatch_cmp.
func PcpatchCmp(pa1 *Pcpatch, pa2 *Pcpatch) int {
	res := C.pcpatch_cmp(pa1._inner, pa2._inner)
	return int(res)
}


// PcpatchEq wraps MEOS C function pcpatch_eq.
func PcpatchEq(pa1 *Pcpatch, pa2 *Pcpatch) bool {
	res := C.pcpatch_eq(pa1._inner, pa2._inner)
	return bool(res)
}


// PcpatchNe wraps MEOS C function pcpatch_ne.
func PcpatchNe(pa1 *Pcpatch, pa2 *Pcpatch) bool {
	res := C.pcpatch_ne(pa1._inner, pa2._inner)
	return bool(res)
}


// PcpatchLt wraps MEOS C function pcpatch_lt.
func PcpatchLt(pa1 *Pcpatch, pa2 *Pcpatch) bool {
	res := C.pcpatch_lt(pa1._inner, pa2._inner)
	return bool(res)
}


// PcpatchLe wraps MEOS C function pcpatch_le.
func PcpatchLe(pa1 *Pcpatch, pa2 *Pcpatch) bool {
	res := C.pcpatch_le(pa1._inner, pa2._inner)
	return bool(res)
}


// PcpatchGt wraps MEOS C function pcpatch_gt.
func PcpatchGt(pa1 *Pcpatch, pa2 *Pcpatch) bool {
	res := C.pcpatch_gt(pa1._inner, pa2._inner)
	return bool(res)
}


// PcpatchGe wraps MEOS C function pcpatch_ge.
func PcpatchGe(pa1 *Pcpatch, pa2 *Pcpatch) bool {
	res := C.pcpatch_ge(pa1._inner, pa2._inner)
	return bool(res)
}


// PcpointsetIn wraps MEOS C function pcpointset_in.
func PcpointsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.pcpointset_in(_c_str)
	return &Set{_inner: res}
}


// PcpointsetOut wraps MEOS C function pcpointset_out.
func PcpointsetOut(s *Set, maxdd int) string {
	res := C.pcpointset_out(s._inner, C.int(maxdd))
	return C.GoString(res)
}


// PcpointsetMake wraps MEOS C function pcpointset_make.
func PcpointsetMake(values []*Pcpoint) *Set {
	_c_values := make([]*C.Pcpoint, len(values))
	for _i, _v := range values { _c_values[_i] = _v._inner }
	res := C.pcpointset_make((**C.Pcpoint)(unsafe.Pointer(&_c_values[0])), C.int(len(values)))
	return &Set{_inner: res}
}


// PcpointToSet wraps MEOS C function pcpoint_to_set.
func PcpointToSet(pt *Pcpoint) *Set {
	res := C.pcpoint_to_set(pt._inner)
	return &Set{_inner: res}
}


// PcpointsetStartValue wraps MEOS C function pcpointset_start_value.
func PcpointsetStartValue(s *Set) *Pcpoint {
	res := C.pcpointset_start_value(s._inner)
	return &Pcpoint{_inner: res}
}


// PcpointsetEndValue wraps MEOS C function pcpointset_end_value.
func PcpointsetEndValue(s *Set) *Pcpoint {
	res := C.pcpointset_end_value(s._inner)
	return &Pcpoint{_inner: res}
}


// PcpointsetValueN wraps MEOS C function pcpointset_value_n.
func PcpointsetValueN(s *Set, n int) (bool, *Pcpoint) {
	var _out_result *C.Pcpoint
	res := C.pcpointset_value_n(s._inner, C.int(n), &_out_result)
	return bool(res), &Pcpoint{_inner: _out_result}
}


// TODO pcpointset_values: unsupported return type Pcpoint **
// func PcpointsetValues(...) { /* not yet handled by codegen */ }


// ContainsSetPcpoint wraps MEOS C function contains_set_pcpoint.
func ContainsSetPcpoint(s *Set, pt *Pcpoint) bool {
	res := C.contains_set_pcpoint(s._inner, pt._inner)
	return bool(res)
}


// ContainedPcpointSet wraps MEOS C function contained_pcpoint_set.
func ContainedPcpointSet(pt *Pcpoint, s *Set) bool {
	res := C.contained_pcpoint_set(pt._inner, s._inner)
	return bool(res)
}


// IntersectionPcpointSet wraps MEOS C function intersection_pcpoint_set.
func IntersectionPcpointSet(pt *Pcpoint, s *Set) *Set {
	res := C.intersection_pcpoint_set(pt._inner, s._inner)
	return &Set{_inner: res}
}


// IntersectionSetPcpoint wraps MEOS C function intersection_set_pcpoint.
func IntersectionSetPcpoint(s *Set, pt *Pcpoint) *Set {
	res := C.intersection_set_pcpoint(s._inner, pt._inner)
	return &Set{_inner: res}
}


// MinusPcpointSet wraps MEOS C function minus_pcpoint_set.
func MinusPcpointSet(pt *Pcpoint, s *Set) *Set {
	res := C.minus_pcpoint_set(pt._inner, s._inner)
	return &Set{_inner: res}
}


// MinusSetPcpoint wraps MEOS C function minus_set_pcpoint.
func MinusSetPcpoint(s *Set, pt *Pcpoint) *Set {
	res := C.minus_set_pcpoint(s._inner, pt._inner)
	return &Set{_inner: res}
}


// UnionPcpointSet wraps MEOS C function union_pcpoint_set.
func UnionPcpointSet(pt *Pcpoint, s *Set) *Set {
	res := C.union_pcpoint_set(pt._inner, s._inner)
	return &Set{_inner: res}
}


// UnionSetPcpoint wraps MEOS C function union_set_pcpoint.
func UnionSetPcpoint(s *Set, pt *Pcpoint) *Set {
	res := C.union_set_pcpoint(s._inner, pt._inner)
	return &Set{_inner: res}
}


// PcpointUnionTransfn wraps MEOS C function pcpoint_union_transfn.
func PcpointUnionTransfn(state *Set, pt *Pcpoint) *Set {
	res := C.pcpoint_union_transfn(state._inner, pt._inner)
	return &Set{_inner: res}
}


// PcpatchsetIn wraps MEOS C function pcpatchset_in.
func PcpatchsetIn(str string) *Set {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.pcpatchset_in(_c_str)
	return &Set{_inner: res}
}


// PcpatchsetOut wraps MEOS C function pcpatchset_out.
func PcpatchsetOut(s *Set, maxdd int) string {
	res := C.pcpatchset_out(s._inner, C.int(maxdd))
	return C.GoString(res)
}


// PcpatchsetMake wraps MEOS C function pcpatchset_make.
func PcpatchsetMake(values []*Pcpatch) *Set {
	_c_values := make([]*C.Pcpatch, len(values))
	for _i, _v := range values { _c_values[_i] = _v._inner }
	res := C.pcpatchset_make((**C.Pcpatch)(unsafe.Pointer(&_c_values[0])), C.int(len(values)))
	return &Set{_inner: res}
}


// PcpatchToSet wraps MEOS C function pcpatch_to_set.
func PcpatchToSet(pa *Pcpatch) *Set {
	res := C.pcpatch_to_set(pa._inner)
	return &Set{_inner: res}
}


// PcpatchsetStartValue wraps MEOS C function pcpatchset_start_value.
func PcpatchsetStartValue(s *Set) *Pcpatch {
	res := C.pcpatchset_start_value(s._inner)
	return &Pcpatch{_inner: res}
}


// PcpatchsetEndValue wraps MEOS C function pcpatchset_end_value.
func PcpatchsetEndValue(s *Set) *Pcpatch {
	res := C.pcpatchset_end_value(s._inner)
	return &Pcpatch{_inner: res}
}


// PcpatchsetValueN wraps MEOS C function pcpatchset_value_n.
func PcpatchsetValueN(s *Set, n int) (bool, *Pcpatch) {
	var _out_result *C.Pcpatch
	res := C.pcpatchset_value_n(s._inner, C.int(n), &_out_result)
	return bool(res), &Pcpatch{_inner: _out_result}
}


// TODO pcpatchset_values: unsupported return type Pcpatch **
// func PcpatchsetValues(...) { /* not yet handled by codegen */ }


// ContainsSetPcpatch wraps MEOS C function contains_set_pcpatch.
func ContainsSetPcpatch(s *Set, pa *Pcpatch) bool {
	res := C.contains_set_pcpatch(s._inner, pa._inner)
	return bool(res)
}


// ContainedPcpatchSet wraps MEOS C function contained_pcpatch_set.
func ContainedPcpatchSet(pa *Pcpatch, s *Set) bool {
	res := C.contained_pcpatch_set(pa._inner, s._inner)
	return bool(res)
}


// IntersectionPcpatchSet wraps MEOS C function intersection_pcpatch_set.
func IntersectionPcpatchSet(pa *Pcpatch, s *Set) *Set {
	res := C.intersection_pcpatch_set(pa._inner, s._inner)
	return &Set{_inner: res}
}


// IntersectionSetPcpatch wraps MEOS C function intersection_set_pcpatch.
func IntersectionSetPcpatch(s *Set, pa *Pcpatch) *Set {
	res := C.intersection_set_pcpatch(s._inner, pa._inner)
	return &Set{_inner: res}
}


// MinusPcpatchSet wraps MEOS C function minus_pcpatch_set.
func MinusPcpatchSet(pa *Pcpatch, s *Set) *Set {
	res := C.minus_pcpatch_set(pa._inner, s._inner)
	return &Set{_inner: res}
}


// MinusSetPcpatch wraps MEOS C function minus_set_pcpatch.
func MinusSetPcpatch(s *Set, pa *Pcpatch) *Set {
	res := C.minus_set_pcpatch(s._inner, pa._inner)
	return &Set{_inner: res}
}


// UnionPcpatchSet wraps MEOS C function union_pcpatch_set.
func UnionPcpatchSet(pa *Pcpatch, s *Set) *Set {
	res := C.union_pcpatch_set(pa._inner, s._inner)
	return &Set{_inner: res}
}


// UnionSetPcpatch wraps MEOS C function union_set_pcpatch.
func UnionSetPcpatch(s *Set, pa *Pcpatch) *Set {
	res := C.union_set_pcpatch(s._inner, pa._inner)
	return &Set{_inner: res}
}


// PcpatchUnionTransfn wraps MEOS C function pcpatch_union_transfn.
func PcpatchUnionTransfn(state *Set, pa *Pcpatch) *Set {
	res := C.pcpatch_union_transfn(state._inner, pa._inner)
	return &Set{_inner: res}
}


// TpcboxIn wraps MEOS C function tpcbox_in.
func TpcboxIn(str string) *TPCBox {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	res := C.tpcbox_in(_c_str)
	return &TPCBox{_inner: res}
}


// TpcboxOut wraps MEOS C function tpcbox_out.
func TpcboxOut(box *TPCBox, maxdd int) string {
	res := C.tpcbox_out(box._inner, C.int(maxdd))
	return C.GoString(res)
}


// TpcboxMake wraps MEOS C function tpcbox_make.
func TpcboxMake(hasx bool, hasz bool, hast bool, geodetic bool, srid int32, pcid uint32, xmin float64, xmax float64, ymin float64, ymax float64, zmin float64, zmax float64, period *Span) *TPCBox {
	res := C.tpcbox_make(C.bool(hasx), C.bool(hasz), C.bool(hast), C.bool(geodetic), C.int32_t(srid), C.uint32_t(pcid), C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.double(zmin), C.double(zmax), period._inner)
	return &TPCBox{_inner: res}
}


// TpcboxCopy wraps MEOS C function tpcbox_copy.
func TpcboxCopy(box *TPCBox) *TPCBox {
	res := C.tpcbox_copy(box._inner)
	return &TPCBox{_inner: res}
}


// PcpatchToTpcbox wraps MEOS C function pcpatch_to_tpcbox.
func PcpatchToTpcbox(pa *Pcpatch, srid int32) *TPCBox {
	res := C.pcpatch_to_tpcbox(pa._inner, C.int32_t(srid))
	return &TPCBox{_inner: res}
}


// TpcboxHasx wraps MEOS C function tpcbox_hasx.
func TpcboxHasx(box *TPCBox) bool {
	res := C.tpcbox_hasx(box._inner)
	return bool(res)
}


// TpcboxHasz wraps MEOS C function tpcbox_hasz.
func TpcboxHasz(box *TPCBox) bool {
	res := C.tpcbox_hasz(box._inner)
	return bool(res)
}


// TpcboxHast wraps MEOS C function tpcbox_hast.
func TpcboxHast(box *TPCBox) bool {
	res := C.tpcbox_hast(box._inner)
	return bool(res)
}


// TpcboxGeodetic wraps MEOS C function tpcbox_geodetic.
func TpcboxGeodetic(box *TPCBox) bool {
	res := C.tpcbox_geodetic(box._inner)
	return bool(res)
}


// TpcboxXmin wraps MEOS C function tpcbox_xmin.
func TpcboxXmin(box *TPCBox) (bool, float64) {
	var _out_result C.double
	res := C.tpcbox_xmin(box._inner, &_out_result)
	return bool(res), float64(_out_result)
}


// TpcboxXmax wraps MEOS C function tpcbox_xmax.
func TpcboxXmax(box *TPCBox) (bool, float64) {
	var _out_result C.double
	res := C.tpcbox_xmax(box._inner, &_out_result)
	return bool(res), float64(_out_result)
}


// TpcboxYmin wraps MEOS C function tpcbox_ymin.
func TpcboxYmin(box *TPCBox) (bool, float64) {
	var _out_result C.double
	res := C.tpcbox_ymin(box._inner, &_out_result)
	return bool(res), float64(_out_result)
}


// TpcboxYmax wraps MEOS C function tpcbox_ymax.
func TpcboxYmax(box *TPCBox) (bool, float64) {
	var _out_result C.double
	res := C.tpcbox_ymax(box._inner, &_out_result)
	return bool(res), float64(_out_result)
}


// TpcboxZmin wraps MEOS C function tpcbox_zmin.
func TpcboxZmin(box *TPCBox) (bool, float64) {
	var _out_result C.double
	res := C.tpcbox_zmin(box._inner, &_out_result)
	return bool(res), float64(_out_result)
}


// TpcboxZmax wraps MEOS C function tpcbox_zmax.
func TpcboxZmax(box *TPCBox) (bool, float64) {
	var _out_result C.double
	res := C.tpcbox_zmax(box._inner, &_out_result)
	return bool(res), float64(_out_result)
}


// TpcboxTmin wraps MEOS C function tpcbox_tmin.
func TpcboxTmin(box *TPCBox) (bool, int64) {
	var _out_result C.TimestampTz
	res := C.tpcbox_tmin(box._inner, &_out_result)
	return bool(res), int64(_out_result)
}


// TpcboxTmax wraps MEOS C function tpcbox_tmax.
func TpcboxTmax(box *TPCBox) (bool, int64) {
	var _out_result C.TimestampTz
	res := C.tpcbox_tmax(box._inner, &_out_result)
	return bool(res), int64(_out_result)
}


// TpcboxSRID wraps MEOS C function tpcbox_srid.
func TpcboxSRID(box *TPCBox) int32 {
	res := C.tpcbox_srid(box._inner)
	return int32(res)
}


// TpcboxPcid wraps MEOS C function tpcbox_pcid.
func TpcboxPcid(box *TPCBox) uint32 {
	res := C.tpcbox_pcid(box._inner)
	return uint32(res)
}


// TpcboxToSTBOX wraps MEOS C function tpcbox_to_stbox.
func TpcboxToSTBOX(box *TPCBox) *STBox {
	res := C.tpcbox_to_stbox(box._inner)
	return &STBox{_inner: res}
}


// TpcboxExpand wraps MEOS C function tpcbox_expand.
func TpcboxExpand(box1 *TPCBox, box2 *TPCBox) {
	C.tpcbox_expand(box1._inner, box2._inner)
}


// TpcboxRound wraps MEOS C function tpcbox_round.
func TpcboxRound(box *TPCBox, maxdd int) *TPCBox {
	res := C.tpcbox_round(box._inner, C.int(maxdd))
	return &TPCBox{_inner: res}
}


// TpcboxSetSRID wraps MEOS C function tpcbox_set_srid.
func TpcboxSetSRID(box *TPCBox, srid int32) *TPCBox {
	res := C.tpcbox_set_srid(box._inner, C.int32_t(srid))
	return &TPCBox{_inner: res}
}


// UnionTpcboxTpcbox wraps MEOS C function union_tpcbox_tpcbox.
func UnionTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox, strict bool) *TPCBox {
	res := C.union_tpcbox_tpcbox(box1._inner, box2._inner, C.bool(strict))
	return &TPCBox{_inner: res}
}


// InterTpcboxTpcbox wraps MEOS C function inter_tpcbox_tpcbox.
func InterTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (bool, *TPCBox) {
	var _out_result C.TPCBox
	res := C.inter_tpcbox_tpcbox(box1._inner, box2._inner, &_out_result)
	return bool(res), &TPCBox{_inner: &_out_result}
}


// IntersectionTpcboxTpcbox wraps MEOS C function intersection_tpcbox_tpcbox.
func IntersectionTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) *TPCBox {
	res := C.intersection_tpcbox_tpcbox(box1._inner, box2._inner)
	return &TPCBox{_inner: res}
}


// ContainsTpcboxTpcbox wraps MEOS C function contains_tpcbox_tpcbox.
func ContainsTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.contains_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// ContainedTpcboxTpcbox wraps MEOS C function contained_tpcbox_tpcbox.
func ContainedTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.contained_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// OverlapsTpcboxTpcbox wraps MEOS C function overlaps_tpcbox_tpcbox.
func OverlapsTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.overlaps_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// SameTpcboxTpcbox wraps MEOS C function same_tpcbox_tpcbox.
func SameTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.same_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// AdjacentTpcboxTpcbox wraps MEOS C function adjacent_tpcbox_tpcbox.
func AdjacentTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.adjacent_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// TpcboxCmp wraps MEOS C function tpcbox_cmp.
func TpcboxCmp(box1 *TPCBox, box2 *TPCBox) int {
	res := C.tpcbox_cmp(box1._inner, box2._inner)
	return int(res)
}


// TpcboxEq wraps MEOS C function tpcbox_eq.
func TpcboxEq(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.tpcbox_eq(box1._inner, box2._inner)
	return bool(res)
}


// TpcboxNe wraps MEOS C function tpcbox_ne.
func TpcboxNe(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.tpcbox_ne(box1._inner, box2._inner)
	return bool(res)
}


// TpcboxLt wraps MEOS C function tpcbox_lt.
func TpcboxLt(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.tpcbox_lt(box1._inner, box2._inner)
	return bool(res)
}


// TpcboxLe wraps MEOS C function tpcbox_le.
func TpcboxLe(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.tpcbox_le(box1._inner, box2._inner)
	return bool(res)
}


// TpcboxGt wraps MEOS C function tpcbox_gt.
func TpcboxGt(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.tpcbox_gt(box1._inner, box2._inner)
	return bool(res)
}


// TpcboxGe wraps MEOS C function tpcbox_ge.
func TpcboxGe(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.tpcbox_ge(box1._inner, box2._inner)
	return bool(res)
}


// LeftTpcboxTpcbox wraps MEOS C function left_tpcbox_tpcbox.
func LeftTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.left_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// OverleftTpcboxTpcbox wraps MEOS C function overleft_tpcbox_tpcbox.
func OverleftTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.overleft_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// RightTpcboxTpcbox wraps MEOS C function right_tpcbox_tpcbox.
func RightTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.right_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// OverrightTpcboxTpcbox wraps MEOS C function overright_tpcbox_tpcbox.
func OverrightTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.overright_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// BelowTpcboxTpcbox wraps MEOS C function below_tpcbox_tpcbox.
func BelowTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.below_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// OverbelowTpcboxTpcbox wraps MEOS C function overbelow_tpcbox_tpcbox.
func OverbelowTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.overbelow_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// AboveTpcboxTpcbox wraps MEOS C function above_tpcbox_tpcbox.
func AboveTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.above_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// OveraboveTpcboxTpcbox wraps MEOS C function overabove_tpcbox_tpcbox.
func OveraboveTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.overabove_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// FrontTpcboxTpcbox wraps MEOS C function front_tpcbox_tpcbox.
func FrontTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.front_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// OverfrontTpcboxTpcbox wraps MEOS C function overfront_tpcbox_tpcbox.
func OverfrontTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.overfront_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// BackTpcboxTpcbox wraps MEOS C function back_tpcbox_tpcbox.
func BackTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.back_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// OverbackTpcboxTpcbox wraps MEOS C function overback_tpcbox_tpcbox.
func OverbackTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.overback_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// BeforeTpcboxTpcbox wraps MEOS C function before_tpcbox_tpcbox.
func BeforeTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.before_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// OverbeforeTpcboxTpcbox wraps MEOS C function overbefore_tpcbox_tpcbox.
func OverbeforeTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.overbefore_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// AfterTpcboxTpcbox wraps MEOS C function after_tpcbox_tpcbox.
func AfterTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.after_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// OverafterTpcboxTpcbox wraps MEOS C function overafter_tpcbox_tpcbox.
func OverafterTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.overafter_tpcbox_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// EnsureSamePcidTpcbox wraps MEOS C function ensure_same_pcid_tpcbox.
func EnsureSamePcidTpcbox(box1 *TPCBox, box2 *TPCBox) bool {
	res := C.ensure_same_pcid_tpcbox(box1._inner, box2._inner)
	return bool(res)
}


// TpointcloudinstMake wraps MEOS C function tpointcloudinst_make.
func TpointcloudinstMake(pt *Pcpoint, t int64) TInstant {
	res := C.tpointcloudinst_make(pt._inner, C.TimestampTz(t))
	return TInstant{_inner: res}
}


// EintersectsTpcpointGeo wraps MEOS C function eintersects_tpcpoint_geo.
func EintersectsTpcpointGeo(temp Temporal, gs *Geom) bool {
	res := C.eintersects_tpcpoint_geo(temp.Inner(), gs._inner)
	return bool(res)
}


// NadTpcpointGeo wraps MEOS C function nad_tpcpoint_geo.
func NadTpcpointGeo(temp Temporal, gs *Geom) float64 {
	res := C.nad_tpcpoint_geo(temp.Inner(), gs._inner)
	return float64(res)
}

