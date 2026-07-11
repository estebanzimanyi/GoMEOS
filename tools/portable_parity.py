#!/usr/bin/env python3
"""Portable bare-name parity gate for GoMEOS.

The GoMEOS analogue of MEOS-API's portable_parity.py, MobilityDB's
`tools/portable_aliases/generate.py --check`, and the MEOS.NET / PyMEOS
parity gates. Per the cross-repo handoff (MEOS-API PR #9): a binding is
done when its **exposed symbol set ⊇ portableAliases.bareNames**, verified
with the *same prefix logic* as MEOS-API portable_parity.py, **0
unbacked**, no per-binding exceptions, across all six in-scope type
families.

"Exposed symbol set" for GoMEOS = every MEOS C function the CGO layer
references (`C.<symbol>(`):

  * the hand-written package at the repo root (`*.go`), and
  * the IDL-driven generated surface under `functions/*.go`, which
    `tools/codegen.py` emits from `tools/meos-idl.json` — the MEOS-API
    parser's JSON output. Measuring the generated surface is the exact
    GoMEOS analogue of MEOS.NET gating its generated P/Invoke file and
    PyMEOS/JMEOS gating their codegen: every binding derives the same
    dialect from the one catalog, so coverage is leveraged from the
    MEOS-API JSON end to end — never by parsing `meos.h`.

GoMEOS wraps MEOS directly, so these are the operators' *own* backing C
functions, reused by construction — never a reimplementation. A bare name
is *backed* iff some referenced symbol `== bareName` or
`startswith(bareName + "_")`, falling back to the contract's verified
`explicitBacking` prefixes (`nearestApproachDistance` <- the `nad_*`
family).

The portable-aliases contract is read from the catalog's folded-in
`portableAliases` when an --idl is given and carries it; otherwise from
the vendored, byte-identical SoT copy tools/portable-aliases.json (the
vendored MEOS 1.4 `tools/meos-idl.json` predates the MEOS-API #8 fold-in,
so the gate uses the SoT copy by default and stays self-contained).

Version bridge (now inert at MEOS 1.4, kept as a documented safety net):
pre-1.4 headers named the `<->` temporal-distance operator `distance_t*`;
MEOS 1.4 (vendored here via GoMEOS PR #3) uses the canonical `tdistance_*`,
so `tdistance` resolves directly by prefix and `BINDING_BACKING` no longer
fires. It is retained only so the gate still reports honestly if pointed
at an older header.

    python3 tools/portable_parity.py            # write report
    python3 tools/portable_parity.py --check    # exit non-zero on any gap

Writes tools/portable-parity.report.json.
"""

from __future__ import annotations

import argparse
import json
import re
import sys
from pathlib import Path

REPO = Path(__file__).resolve().parent.parent
PREVIEW = REPO / "functions"
VENDORED = Path(__file__).resolve().parent / "portable-aliases.json"
REPORT = Path(__file__).resolve().parent / "portable-parity.report.json"

# Full user-facing temporal type families — cbuffer/npoint/pose/rgeo are
# NOT internals and must never be excluded from the parity headline.
# Precedence keeps the broad geo/temporal buckets from swallowing them.
IN_SCOPE_FAMILIES = ["temporal", "geo", "cbuffer", "npoint", "pose", "rgeo"]

# Inert at MEOS 1.4 (canonical `tdistance_*` is present); retained as a
# documented safety net for older/partial header scans only.
BINDING_BACKING = {
    "tdistance": ["distance_tfloat", "distance_tint",
                  "distance_tnumber", "distance_tpoint"],
}

# Every MEOS C symbol referenced through CGO: `C.<symbol>(`.
_CGO_RE = re.compile(r"\bC\.([A-Za-z_]\w*)\s*\(")
_CGO_PSEUDO = {"CString", "CBytes", "GoString", "GoBytes", "free",
               "malloc", "calloc"}


def _scan(path: Path, syms: set[str]) -> None:
    for m in _CGO_RE.findall(path.read_text()):
        if m not in _CGO_PSEUDO:
            syms.add(m)


def exposed_symbols(repo: Path) -> list[str]:
    """MEOS C function names the CGO layer references — hand-written root
    package plus the IDL-driven generated surface (functions)."""
    syms: set[str] = set()
    for go in sorted(repo.glob("*.go")):
        if not go.name.endswith("_test.go"):
            _scan(go, syms)
    if PREVIEW.is_dir():
        for go in sorted(PREVIEW.glob("*.go")):
            if not go.name.endswith("_test.go"):
                _scan(go, syms)
    return sorted(syms)


def load_portable_aliases(idl_path: str | None) -> dict:
    """Prefer the catalog's folded-in portableAliases; else the vendored SoT."""
    if idl_path:
        idl = json.loads(Path(idl_path).read_text())
        pa = idl.get("portableAliases")
        if pa and pa.get("families"):
            return pa
    return json.loads(VENDORED.read_text())


def family_of(name: str) -> str:
    n = name.lower()
    if "rgeo" in n:
        return "rgeo"
    if "cbuffer" in n:
        return "cbuffer"
    if "npoint" in n:
        return "npoint"
    if "pose" in n:
        return "pose"
    if any(t in n for t in ("geo", "geom", "geog", "point", "spatial")):
        return "geo"
    return "temporal"


def build_parity(symbols: list[str], pa: dict) -> dict:
    fam_of = {p["bareName"]: (fam, p["operator"])
              for fam, lst in pa["families"].items() for p in lst}
    explicit = dict(pa.get("explicitBacking", {}))

    def matches(prefix: str) -> list[str]:
        return [s for s in symbols
                if s == prefix or s.startswith(prefix + "_")]

    fams_present = {family_of(s) for s in symbols}

    by_bare: dict[str, dict] = {}
    fam_totals: dict[str, int] = {f: 0 for f in IN_SCOPE_FAMILIES}
    for bare, (fam, op) in sorted(fam_of.items()):
        hits, via = matches(bare), "prefix"
        if not hits:
            for pref in explicit.get(bare, []):
                hits += matches(pref)
            if hits:
                via = "explicit:" + ",".join(explicit.get(bare, []))
        if not hits and bare in BINDING_BACKING:
            for pref in BINDING_BACKING[bare]:
                hits += matches(pref)
            if hits:
                via = "version-bridge:" + ",".join(BINDING_BACKING[bare])
        if not hits:
            via = None
        hist: dict[str, int] = {}
        for h in hits:
            k = family_of(h)
            hist[k] = hist.get(k, 0) + 1
            fam_totals[k] = fam_totals.get(k, 0) + 1
        by_bare[bare] = {
            "operator": op, "family": fam, "via": via,
            "backedBy": len(hits), "sample": sorted(hits)[:3],
            "familyCoverage": hist,
            "status": "backed" if hits else "needs-explicit-backing",
        }

    backed = [b for b, v in by_bare.items() if v["status"] == "backed"]
    unbacked = sorted(b for b, v in by_bare.items()
                      if v["status"] == "needs-explicit-backing")

    #  - covered  : has backings now
    #  - regressed : header carries the type's symbols but zero backings
    #                (a real exclusion — hard fail; never tolerated)
    #  - pending   : type absent from the scanned MEOS surface entirely
    fam_status: dict[str, str] = {}
    for f in IN_SCOPE_FAMILIES:
        if fam_totals.get(f, 0) > 0:
            fam_status[f] = "covered"
        elif f in fams_present:
            fam_status[f] = "regressed"
        else:
            fam_status[f] = "pending"
    regressed = [f for f, s in fam_status.items() if s == "regressed"]
    pending = [f for f, s in fam_status.items() if s == "pending"]

    total = len(by_bare)
    return {
        "exposedSymbols": len(symbols),
        "symbolSources": ["repo-root *.go (hand-written)",
                          "functions/*.go (IDL-driven codegen "
                          "from tools/meos-idl.json — MEOS-API output)"],
        "total": total,
        "backed": len(backed),
        "needsExplicitBacking": len(unbacked),
        "parityPct": round(len(backed) * 100 / total, 1) if total else 0,
        "unbacked": unbacked,
        "familyCoverage": fam_totals,
        "familyStatus": fam_status,
        "regressedFamilies": regressed,
        "pendingFamilies": pending,
        "byBareName": by_bare,
        "provenance": pa.get("provenance", {}),
        "scope": pa.get("scope", {}),
    }


def main() -> int:
    ap = argparse.ArgumentParser(description=__doc__)
    ap.add_argument("--idl", metavar="meos-idl.json", default=None,
                    help="catalog to read portableAliases from "
                         "(default: vendored tools/portable-aliases.json)")
    ap.add_argument("--check", action="store_true",
                    help="exit non-zero if any bare name is unbacked or any "
                         "in-scope family present in the surface is excluded "
                         "(CI gate)")
    args = ap.parse_args()

    symbols = exposed_symbols(REPO)
    pa = load_portable_aliases(args.idl)
    rep = build_parity(symbols, pa)
    REPORT.write_text(json.dumps(rep, indent=2) + "\n")

    src = ("idl.portableAliases" if args.idl
           and json.loads(Path(args.idl).read_text())
           .get("portableAliases", {}).get("families")
           else "vendored tools/portable-aliases.json")
    print(f"[portable-parity] {rep['backed']}/{rep['total']} bare names "
          f"backed in the exposed GoMEOS CGO symbol set "
          f"({rep['parityPct']}%); {rep['needsExplicitBacking']} unbacked "
          f"[contract: {src}]", file=sys.stderr)
    print(f"[portable-parity] six-family status {rep['familyStatus']} "
          f"-> {REPORT}", file=sys.stderr)
    for b in rep["unbacked"]:
        v = rep["byBareName"][b]
        print(f"  needs-explicit-backing: {b!r} ({v['operator']}, "
              f"{v['family']})", file=sys.stderr)

    # Hard gate = the handoff doc's literal "Done" for a binding:
    # 29/29 bare names backed, 0 unbacked, and every in-scope user-facing
    # family covered (cbuffer/npoint/pose/rgeo are never excluded).
    fail = bool(rep["unbacked"] or rep["regressedFamilies"]
                or rep["pendingFamilies"])
    if args.check:
        if rep["regressedFamilies"]:
            print("  EXCLUDED in-scope families present in surface but "
                  f"unbacked: {rep['regressedFamilies']}", file=sys.stderr)
        if rep["pendingFamilies"]:
            print("  in-scope families absent from the scanned surface: "
                  f"{rep['pendingFamilies']}", file=sys.stderr)
        verdict = ("FAIL" if fail else
                   f"PASS — {rep['backed']}/{rep['total']} = 100%, "
                   "0 unbacked, all six in-scope families covered")
        print(f"CHECK: {verdict}", file=sys.stderr)
        return 1 if fail else 0
    return 0


if __name__ == "__main__":
    sys.exit(main())
