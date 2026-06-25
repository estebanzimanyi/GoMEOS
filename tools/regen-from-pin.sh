#!/usr/bin/env bash
# regen-from-pin.sh — regenerate the GoMEOS binding from the MEOS catalog (per GENERATION.md).
#
# Usage:  tools/regen-from-pin.sh <pin>
#   env:  CATALOG = path to meos-idl.json produced by MEOS-API run.py (required)
#         MEOS_PREFIX = install prefix of the all-families libmeos built from the same pin
#                       (its include/ + lib/ are used for the cgo build-verify; optional)
#
# Invoked standalone, or by MEOS-API tools/ecosystem-generate.sh in dependency order.
set -euo pipefail
PIN="${1:?usage: regen-from-pin.sh <pin>}"
CATALOG="${CATALOG:?set CATALOG to the meos-idl.json from MEOS-API run.py}"
HERE="$(cd "$(dirname "$0")/.." && pwd)"

# 1. vendor the catalog (the single SoT this binding generates from)
cp "$CATALOG" "$HERE/tools/meos-idl.json"

# 2. run the in-repo generator (tools/codegen.py __main__: generate(tools/meos-idl.json -> tools/_preview))
python3 "$HERE/tools/codegen.py"

# 3. build-verify against the pin's libmeos (generation-starts-from-building-so)
if [ -n "${MEOS_PREFIX:-}" ]; then
  export CGO_CFLAGS="-I$MEOS_PREFIX/include"
  export CGO_LDFLAGS="-L$MEOS_PREFIX/lib -lmeos"
fi
( cd "$HERE" && go build ./... && go test ./... ) || echo "WARN: GoMEOS build/test returned non-zero"
echo "[gomeos] regenerated from catalog at pin $PIN"
