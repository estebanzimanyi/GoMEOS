# GoMEOS portable bare-name parity (RFC #920)

The MobilityDB ecosystem defines a canonical, type-agnostic
operator → bare-name dialect so one program/query runs identically on
every engine and binding — a user learns one reference and assumes the
rest. The contract is **29 operator → bareName pairs** across eight
families (topology, time-position, space X/Y/Z, temporal-comparison,
distance, same); the single source of truth is MEOS-API
`meta/portable-aliases.json` (discussion MobilityDB#861 · RFC #920 ·
native MobilityDB#1075 · manual MobilityDB#1078 · cross-repo handoff
MEOS-API PR #9).

`tools/portable-aliases.json` is that contract, vendored
**byte-identical**. (`tools/portable_parity.py --idl tools/meos-idl.json`
prefers the catalog's folded-in `portableAliases` automatically once a
re-vendored MEOS-API catalog carries it; the MEOS 1.4 catalog vendored
today predates the MEOS-API #8 fold-in, so the SoT copy is used by
default and the gate stays self-contained.)

## Everything is leveraged from the MEOS-API JSON — never header parsing

This gate consumes only JSON-derived artifacts:

- **The contract** is MEOS-API's `meta/portable-aliases.json` (JSON),
  vendored byte-identical.
- **The binding surface** it measures is `tools/_preview/*.go`, which
  `tools/codegen.py` generates from `tools/meos-idl.json` — the output
  of the MEOS-API parser. This is the exact GoMEOS analogue of MEOS.NET
  gating its generated P/Invoke file and PyMEOS/JMEOS gating their
  codegen: every binding derives the *same* dialect from the *one*
  catalog.

The gate's only source-level step is reading `C.<symbol>(` tokens from
the generated/hand-written Go to learn which MEOS functions the binding
actually references. It does **not** parse `meos.h`; the C surface it
trusts is the one the MEOS-API JSON described.

## What "backed" means

GoMEOS wraps MEOS directly through CGO, and MEOS C already names every
operator's backing function `<bareName>...` (`&&` → `overlaps_*`,
`#<` → `tlt_*`, `~=` → `same_*`, `<->` → `tdistance_*`), so the binding
exposes the dialect **by construction**: every portable name reuses the
operator's *own* backing C function, never a reimplementation, with no
type-qualified or per-binding form.

A bare name is **backed** when the exposed CGO symbol set (every
`C.<symbol>(` in the hand-written root `*.go` and the IDL-driven
`tools/_preview/*.go`) contains a MEOS function whose name `== bareName`
or `startswith(bareName + "_")`, with one verified, *non-guessed*
fallback from the contract: `nearestApproachDistance` (`|=|`) ← the
`nad_*` family (`explicitBacking`).

The `tdistance` (`<->`) version bridge from the pre-1.4 era
(`distance_t*`) is now **inert**: MEOS 1.4 — vendored here via GoMEOS
PR #3 — provides the canonical `tdistance_*`, so `tdistance` resolves
directly by prefix. The bridge is retained only as a documented safety
net if the gate is pointed at an older header.

## Six-family scope — fully covered

`temporal`, `geo`, `cbuffer`, `npoint`, `pose`, `rgeo` are **all full
user-facing temporal types** and are **never excluded from the parity
headline**. Stacked on the MEOS 1.4 bump (GoMEOS PR #3), the IDL-driven
generated surface backs every one of the 29 bare names across **all
six** families. The gate hard-fails on any family that is *present in
the surface but unbacked* (a real exclusion — `regressed`) or *absent*
(`pending`); neither is tolerated.

Live result: **29 / 29 bare names backed, 0 unbacked, six / six families
covered** (`temporal`, `geo`, `cbuffer`, `npoint`, `pose`, `rgeo`).

## Verifying parity

```sh
python3 tools/portable_parity.py --check     # writes tools/portable-parity.report.json
CGO_ENABLED=0 go test ./tools/parity/ -v     # language-independent mirror
```

Both exit non-zero unless **29/29 bare names are backed, 0 unbacked,
all six in-scope families covered**. They are byte-for-byte equivalent
in verdict (the Go test mirrors the Python script — the analogue of
MEOS-API's `portable_parity.py`, MobilityDB's
`tools/portable_aliases/generate.py --check`, and the MEOS.NET / PyMEOS
parity gates). The same two checks run in the `portable-aliases parity`
CI workflow; neither needs libmeos or CGO.

## Provenance

Discussion MobilityDB#861 · RFC #920 · native MobilityDB#1075 · manual
MobilityDB#1078 · MEOS-API cross-repo handoff PR #9. Stacks on GoMEOS
PR #3 (`bump/meos-1.4`), which vendors the MEOS 1.4 headers and the
IDL-driven codegen.
