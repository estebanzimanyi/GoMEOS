# GoMEOS code generator

`codegen.py` reads `tools/meos-idl.json` (produced by the
[MEOS-API](https://github.com/MobilityDB/MEOS-API) parser) and emits one
idiomatic Go wrapper per MEOS public C function into `tools/_preview/`.
The directory is prefixed with `_` so `go build ./...` ignores it; the
output is a Draft artifact that lets reviewers see the shape of the
generated surface without disturbing the hand-written package.

## Running

```
python3 tools/codegen.py
```

## Coverage today

2369 candidate functions across the six public headers.  2167 emit
cleanly under the heuristic rules; 170 are skipped as `Datum`-bearing
internal helpers (the hand-written surface exposes those through typed
overloads which the codegen cannot synthesise from the IDL); 32 remain
as TODO.

The covered shapes are: scalar inputs, wrapped opaque pointers
(`Temporal`, `STBox`, `TBox`, `Span`, `SpanSet`, `Set`, `GSERIALIZED`,
`TInstant`, `TSequence`, `TSequenceSet`, `Npoint`, `Nsegment`,
`SkipList`, `RTree`, `Match`, `BOX3D`, `GBOX`, `AFFINE`, `PJ_CONTEXT`,
`gsl_rng`, `Interval`), C and PostgreSQL text strings, output
parameters named `result` / `value` (scalar and wrapped-pointer
out-params surfaced as additional Go return values), counted-array
inputs (`T **` / `const T *` paired with `int count` or `size_t size`)
lowered to Go slices, counted-array returns (`T **` paired with
`int *count` or matching the input slice length) lowered to Go slices
including byte buffers and `char **` string arrays, and
`unsafe.Pointer` for `void *` arguments.

The remaining 32 TODOs need per-function metadata:

* `*_values` / `*_insts_p` / `*_sps` (~12 funcs) — array returns whose
  length comes from a sibling accessor (`set_num_values`,
  `temporal_num_instants`, etc.) the codegen cannot guess from the IDL
  alone.
* `*_split` family (`temporal_time_split`, `tfloat_value_split`,
  `tint_value_time_split`, `tgeo_space_split`, etc.) — parallel output
  arrays (`time_bins`, `value_bins`) sharing a single count, plus
  `GSERIALIZED ***` triple-pointer outputs.
* `tpoint_as_mvtgeom`, `tpointseq_make_coords` — multi-array shapes
  beyond a single counted-array companion.
* `tempsubtype_from_string`, `geom_min_bounding_radius` — output
  parameters named `subtype` / `radius` rather than the canonical
  `result` / `value` (could be hand-listed as named outputs once a
  metadata catalog exists).
* `skiplist_*` (function pointers, `void **`) — internal helpers; the
  hand-written surface does not expose them either.

A metadata catalog of `output_parameters` / `result_parameters` /
`array_length_for` tuples (mirroring `build_pymeos_functions.py`'s
`output_parameters` / `result_parameters` / `nullable_parameters`)
would close the remaining shapes without false positives.

## Refreshing the IDL

`tools/meos-idl.json` is vendored.  When MEOS bumps, regenerate with
MEOS-API against the new headers and copy the output back:

```
cp ../MEOS-API/output/meos-idl.json tools/meos-idl.json
python3 tools/codegen.py
```

The vendored IDL relies on MEOS-API PR #1 (the stdbool stub) being in
place so `bool` returns and `bool *` outputs do not get demoted to `int`.
