# GoMEOS generation — the canonical per-binding generator policy

GoMEOS is a **generated** cgo binding. This document is the contract for how it is
generated, under the ecosystem-wide per-binding generator policy.

## The policy (ecosystem-wide)

Every MobilityDB language binding is a **pure projection of the MEOS-API catalog**, and
**each binding owns its own generator, in its own repo**, in a canonical layout — not a
single central generator-repo. The single source of truth is the **catalog**
(`MEOS-API/output/meos-idl.json`, generated from the MEOS C headers), not a generator
location. This mirrors how MEOS itself is built: independent, plug-and-play, CMake-gated
families — a binding is an independent module that owns its generation.

Each binding repo satisfies the same invariants:

1. **In-repo generator**, one clearly-designated location. For GoMEOS that is
   **`tools/codegen.py`** — it reads the vendored catalog and emits cgo wrappers.
2. **Own pin manifest** `tools/pin/compose-order.txt` — the canonical, dependency-ordered
   fold list of the open PRs that compose this binding's surface onto `main`.
3. **Vendored catalog**, version-pinned, read-only: `tools/meos-idl.json` (from a
   MobilityDB `ecosystem-pin-*`).
4. **Thin language projection** — language-neutral decisions (grouping, skip/classify,
   portable names, shape) belong upstream in the catalog, so per-language generators do
   not re-implement and drift.
5. **Full automation (North Star):** generate-then-retire toward a **zero hand-written**
   surface; anything that seems irreducible is either emitted by the generator or fixed at
   source in MEOS (export the symbol) — never hand-patched in the binding.

## GoMEOS scope: flat generated cgo wrappers

`tools/codegen.py` emits one `meos_<header>.go` per MEOS public header plus a single
`cgo.go` preamble — self-contained cgo (its own `#cgo` directives and `#include`s), no
hand-written `cast.h` dependency. The end state is **flat generated wrappers only**; the
idiomatic OO surface, if any, is also generated (no hand layer).

## Generate-then-retire — the green-CI version is the probe

Removing the hand-written binding happens **little by little, never wipe-first**:

1. align the generator + generate the full surface; `go build ./...` green;
2. **prove generated ⊇ hand** against the **last green-CI version** (the equivalence
   probe) — `go test` + parity, **family by family**;
3. retire the hand-written root `*.go` (and `cast.c`/`cast.h`) for that family;
4. repeat. End state: the hand-written root files are gone; the binding is the generated
   package. The green-CI baseline catches a generated gap before it ships.

## Pinning: this binding's catalog comes from a MobilityDB pin

GoMEOS's vendored `tools/meos-idl.json` is generated from a MobilityDB `ecosystem-pin-*`
(master ⊕ the MobilityDB compose-order) via the MEOS-API `run.py`. That pin is the
*catalog/surface* input; GoMEOS's own `tools/pin/compose-order.txt` governs *this repo's*
PR accumulate. See `tools/pin/compose-order.txt` for the composing set and the disposition
of every open PR.
