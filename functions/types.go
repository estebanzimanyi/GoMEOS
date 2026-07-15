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
*/
import "C"
import (
	"unsafe"
)

var _ = unsafe.Pointer(nil)

// -------------------- opaque handle types --------------------
type Temporal struct { _inner *C.Temporal }
func (x *Temporal) Inner() *C.Temporal { if x == nil { return nil }; return x._inner }
type TInstant struct { _inner *C.TInstant }
func (x *TInstant) Inner() *C.TInstant { if x == nil { return nil }; return x._inner }
type TSequence struct { _inner *C.TSequence }
func (x *TSequence) Inner() *C.TSequence { if x == nil { return nil }; return x._inner }
type TSequenceSet struct { _inner *C.TSequenceSet }
func (x *TSequenceSet) Inner() *C.TSequenceSet { if x == nil { return nil }; return x._inner }
type STBox struct { _inner *C.STBox }
func (x *STBox) Inner() *C.STBox { if x == nil { return nil }; return x._inner }
type TBox struct { _inner *C.TBox }
func (x *TBox) Inner() *C.TBox { if x == nil { return nil }; return x._inner }
type Span struct { _inner *C.Span }
func (x *Span) Inner() *C.Span { if x == nil { return nil }; return x._inner }
type SpanSet struct { _inner *C.SpanSet }
func (x *SpanSet) Inner() *C.SpanSet { if x == nil { return nil }; return x._inner }
type Set struct { _inner *C.Set }
func (x *Set) Inner() *C.Set { if x == nil { return nil }; return x._inner }
type Geom struct { _inner *C.GSERIALIZED }
func (x *Geom) Inner() *C.GSERIALIZED { if x == nil { return nil }; return x._inner }
type Interval struct { _inner *C.Interval }
func (x *Interval) Inner() *C.Interval { if x == nil { return nil }; return x._inner }
type Npoint struct { _inner *C.Npoint }
func (x *Npoint) Inner() *C.Npoint { if x == nil { return nil }; return x._inner }
type Nsegment struct { _inner *C.Nsegment }
func (x *Nsegment) Inner() *C.Nsegment { if x == nil { return nil }; return x._inner }
type Cbuffer struct { _inner *C.Cbuffer }
func (x *Cbuffer) Inner() *C.Cbuffer { if x == nil { return nil }; return x._inner }
type Pose struct { _inner *C.Pose }
func (x *Pose) Inner() *C.Pose { if x == nil { return nil }; return x._inner }
type Jsonb struct { _inner *C.Jsonb }
func (x *Jsonb) Inner() *C.Jsonb { if x == nil { return nil }; return x._inner }
type JsonPath struct { _inner *C.JsonPath }
func (x *JsonPath) Inner() *C.JsonPath { if x == nil { return nil }; return x._inner }
type Pcpoint struct { _inner *C.Pcpoint }
func (x *Pcpoint) Inner() *C.Pcpoint { if x == nil { return nil }; return x._inner }
type Pcpatch struct { _inner *C.Pcpatch }
func (x *Pcpatch) Inner() *C.Pcpatch { if x == nil { return nil }; return x._inner }
type TPCBox struct { _inner *C.TPCBox }
func (x *TPCBox) Inner() *C.TPCBox { if x == nil { return nil }; return x._inner }
type MeosArray struct { _inner *C.MeosArray }
func (x *MeosArray) Inner() *C.MeosArray { if x == nil { return nil }; return x._inner }
type PCSchema struct { _inner *C.PCSCHEMA }
func (x *PCSchema) Inner() *C.PCSCHEMA { if x == nil { return nil }; return x._inner }
type SkipList struct { _inner *C.SkipList }
func (x *SkipList) Inner() *C.SkipList { if x == nil { return nil }; return x._inner }
type RTree struct { _inner *C.RTree }
func (x *RTree) Inner() *C.RTree { if x == nil { return nil }; return x._inner }
type Match struct { _inner *C.Match }
func (x *Match) Inner() *C.Match { if x == nil { return nil }; return x._inner }
type Box3D struct { _inner *C.BOX3D }
func (x *Box3D) Inner() *C.BOX3D { if x == nil { return nil }; return x._inner }
type GBox struct { _inner *C.GBOX }
func (x *GBox) Inner() *C.GBOX { if x == nil { return nil }; return x._inner }
type AFFINE struct { _inner *C.AFFINE }
func (x *AFFINE) Inner() *C.AFFINE { if x == nil { return nil }; return x._inner }
type PJContext struct { _inner *C.PJ_CONTEXT }
func (x *PJContext) Inner() *C.PJ_CONTEXT { if x == nil { return nil }; return x._inner }
type GslRng struct { _inner *C.gsl_rng }
func (x *GslRng) Inner() *C.gsl_rng { if x == nil { return nil }; return x._inner }

// -------------------- enums --------------------
type TempSubtype C.tempSubtype
const (
	TempSubtype_ANYTEMPSUBTYPE TempSubtype = 0
	TempSubtype_TINSTANT TempSubtype = 1
	TempSubtype_TSEQUENCE TempSubtype = 2
	TempSubtype_TSEQUENCESET TempSubtype = 3
)
type Interpolation C.interpType
const (
	Interpolation_INTERP_NONE Interpolation = 0
	Interpolation_DISCRETE Interpolation = 1
	Interpolation_STEP Interpolation = 2
	Interpolation_LINEAR Interpolation = 3
)
type RTreeSearchOp C.RTreeSearchOp
const (
	RTreeSearchOp_RTREE_OVERLAPS RTreeSearchOp = 0
	RTreeSearchOp_RTREE_CONTAINS RTreeSearchOp = 1
	RTreeSearchOp_RTREE_CONTAINED_BY RTreeSearchOp = 2
)
type MeosType C.MeosType
const (
	MeosType_T_UNKNOWN MeosType = 0
	MeosType_T_BOOL MeosType = 1
	MeosType_T_DATE MeosType = 2
	MeosType_T_DATEMULTIRANGE MeosType = 3
	MeosType_T_DATERANGE MeosType = 4
	MeosType_T_DATESET MeosType = 5
	MeosType_T_DATESPAN MeosType = 6
	MeosType_T_DATESPANSET MeosType = 7
	MeosType_T_DOUBLE2 MeosType = 8
	MeosType_T_DOUBLE3 MeosType = 9
	MeosType_T_DOUBLE4 MeosType = 10
	MeosType_T_FLOAT8 MeosType = 11
	MeosType_T_FLOATSET MeosType = 12
	MeosType_T_FLOATSPAN MeosType = 13
	MeosType_T_FLOATSPANSET MeosType = 14
	MeosType_T_INT4 MeosType = 15
	MeosType_T_INT4MULTIRANGE MeosType = 16
	MeosType_T_INT4RANGE MeosType = 17
	MeosType_T_INTSET MeosType = 18
	MeosType_T_INTSPAN MeosType = 19
	MeosType_T_INTSPANSET MeosType = 20
	MeosType_T_INT8 MeosType = 21
	MeosType_T_INT8MULTIRANGE MeosType = 52
	MeosType_T_INT8RANGE MeosType = 53
	MeosType_T_BIGINTSET MeosType = 22
	MeosType_T_BIGINTSPAN MeosType = 23
	MeosType_T_BIGINTSPANSET MeosType = 24
	MeosType_T_STBOX MeosType = 25
	MeosType_T_TBOOL MeosType = 26
	MeosType_T_TBOX MeosType = 27
	MeosType_T_TDOUBLE2 MeosType = 28
	MeosType_T_TDOUBLE3 MeosType = 29
	MeosType_T_TDOUBLE4 MeosType = 30
	MeosType_T_TEXT MeosType = 31
	MeosType_T_TEXTSET MeosType = 32
	MeosType_T_TFLOAT MeosType = 33
	MeosType_T_TIMESTAMPTZ MeosType = 34
	MeosType_T_TINT MeosType = 35
	MeosType_T_TSTZMULTIRANGE MeosType = 36
	MeosType_T_TSTZRANGE MeosType = 37
	MeosType_T_TSTZSET MeosType = 38
	MeosType_T_TSTZSPAN MeosType = 39
	MeosType_T_TSTZSPANSET MeosType = 40
	MeosType_T_TTEXT MeosType = 41
	MeosType_T_GEOMETRY MeosType = 42
	MeosType_T_GEOMSET MeosType = 43
	MeosType_T_GEOGRAPHY MeosType = 44
	MeosType_T_GEOGSET MeosType = 45
	MeosType_T_TGEOMPOINT MeosType = 46
	MeosType_T_TGEOGPOINT MeosType = 47
	MeosType_T_NPOINT MeosType = 48
	MeosType_T_NPOINTSET MeosType = 49
	MeosType_T_NSEGMENT MeosType = 50
	MeosType_T_TNPOINT MeosType = 51
	MeosType_T_POSE MeosType = 54
	MeosType_T_POSESET MeosType = 55
	MeosType_T_TPOSE MeosType = 56
	MeosType_T_CBUFFER MeosType = 57
	MeosType_T_CBUFFERSET MeosType = 58
	MeosType_T_TCBUFFER MeosType = 59
	MeosType_T_TGEOMETRY MeosType = 60
	MeosType_T_TGEOGRAPHY MeosType = 61
	MeosType_T_TRGEOMETRY MeosType = 62
	MeosType_T_JSONB MeosType = 63
	MeosType_T_JSONPATH MeosType = 64
	MeosType_T_JSONBSET MeosType = 65
	MeosType_T_TJSONB MeosType = 66
	MeosType_T_TBIGINT MeosType = 67
	MeosType_T_H3INDEX MeosType = 68
	MeosType_T_H3INDEXSET MeosType = 69
	MeosType_T_TH3INDEX MeosType = 70
	MeosType_T_QUADBIN MeosType = 71
	MeosType_T_QUADBINSET MeosType = 72
	MeosType_T_TQUADBIN MeosType = 73
	MeosType_T_PCPOINT MeosType = 74
	MeosType_T_PCPOINTSET MeosType = 75
	MeosType_T_TPCPOINT MeosType = 76
	MeosType_T_PCPATCH MeosType = 77
	MeosType_T_PCPATCHSET MeosType = 78
	MeosType_T_TPCPATCH MeosType = 79
	MeosType_T_TPCBOX MeosType = 80
	MeosType_T_RAQUET MeosType = 81
	MeosType_NUM_MEOS_TYPES MeosType = 82
)
type MeosOper C.MeosOper
const (
	MeosOper_UNKNOWN_OP MeosOper = 0
	MeosOper_EQ_OP MeosOper = 1
	MeosOper_NE_OP MeosOper = 2
	MeosOper_LT_OP MeosOper = 3
	MeosOper_LE_OP MeosOper = 4
	MeosOper_GT_OP MeosOper = 5
	MeosOper_GE_OP MeosOper = 6
	MeosOper_ADJACENT_OP MeosOper = 7
	MeosOper_UNION_OP MeosOper = 8
	MeosOper_MINUS_OP MeosOper = 9
	MeosOper_INTERSECT_OP MeosOper = 10
	MeosOper_OVERLAPS_OP MeosOper = 11
	MeosOper_CONTAINS_OP MeosOper = 12
	MeosOper_CONTAINED_OP MeosOper = 13
	MeosOper_SAME_OP MeosOper = 14
	MeosOper_LEFT_OP MeosOper = 15
	MeosOper_OVERLEFT_OP MeosOper = 16
	MeosOper_RIGHT_OP MeosOper = 17
	MeosOper_OVERRIGHT_OP MeosOper = 18
	MeosOper_BELOW_OP MeosOper = 19
	MeosOper_OVERBELOW_OP MeosOper = 20
	MeosOper_ABOVE_OP MeosOper = 21
	MeosOper_OVERABOVE_OP MeosOper = 22
	MeosOper_FRONT_OP MeosOper = 23
	MeosOper_OVERFRONT_OP MeosOper = 24
	MeosOper_BACK_OP MeosOper = 25
	MeosOper_OVERBACK_OP MeosOper = 26
	MeosOper_BEFORE_OP MeosOper = 27
	MeosOper_OVERBEFORE_OP MeosOper = 28
	MeosOper_AFTER_OP MeosOper = 29
	MeosOper_OVERAFTER_OP MeosOper = 30
	MeosOper_EVEREQ_OP MeosOper = 31
	MeosOper_EVERNE_OP MeosOper = 32
	MeosOper_EVERLT_OP MeosOper = 33
	MeosOper_EVERLE_OP MeosOper = 34
	MeosOper_EVERGT_OP MeosOper = 35
	MeosOper_EVERGE_OP MeosOper = 36
	MeosOper_ALWAYSEQ_OP MeosOper = 37
	MeosOper_ALWAYSNE_OP MeosOper = 38
	MeosOper_ALWAYSLT_OP MeosOper = 39
	MeosOper_ALWAYSLE_OP MeosOper = 40
	MeosOper_ALWAYSGT_OP MeosOper = 41
	MeosOper_ALWAYSGE_OP MeosOper = 42
	MeosOper_TEMPCONTAINS_OP MeosOper = 43
	MeosOper_TEMPCONTAINED_OP MeosOper = 44
)
type SkipListType C.SkipListType
const (
	SkipListType_SKIPLIST_TEMPORAL SkipListType = 0
	SkipListType_SKIPLIST_KEYVALUE SkipListType = 1
)
type NullHandleType C.nullHandleType
const (
	NullHandleType_NULL_INVALID NullHandleType = 0
	NullHandleType_NULL_ERROR NullHandleType = 1
	NullHandleType_NULL_JSON_NULL NullHandleType = 2
	NullHandleType_NULL_DELETE NullHandleType = 3
	NullHandleType_NULL_RETURN NullHandleType = 4
)
