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

2369 candidate functions across the six public headers; 1967 emit cleanly
under the scalar / wrapped-pointer rules, 170 are skipped as `Datum`
internal helpers (the hand-written surface exposes those through typed
overloads which the codegen cannot synthesise from the IDL), and 232
remain as TODO until per-function metadata is encoded.

The 232 TODOs are all output-parameter or counted-array shapes:
`int *`, `double *`, `bool *`, `TimestampTz *`, `size_t *` (typically
`count` accumulators), and double-pointer arrays like `Temporal **` /
`GSERIALIZED **` / `TInstant **` / `TSequence **`.  Resolving them needs
a catalog of which parameters are outputs, which are array lengths, and
which are nullable, mirroring the `result_parameters` /
`output_parameters` / `nullable_parameters` sets PyMEOS-CFFI's
`build_pymeos_functions.py` carries.

## Refreshing the IDL

`tools/meos-idl.json` is vendored.  When MEOS bumps, regenerate with
MEOS-API against the new headers and copy the output back:

```
cp ../MEOS-API/output/meos-idl.json tools/meos-idl.json
python3 tools/codegen.py
```

The vendored IDL relies on MEOS-API PR #1 (the stdbool stub) being in
place so `bool` returns and `bool *` outputs do not get demoted to `int`.
