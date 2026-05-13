"""GoMEOS code generator.

Drives idiomatic Go wrapper generation from meos-idl.json (produced by the
MEOS-API parser).  Output is one Go source file per MEOS public header,
emitted to ``generated/`` and consumed by the hand-written ergonomic surface
in the package root.

The generator is intentionally minimal in its first iteration: it covers the
scalar-in / opaque-pointer-out signatures (the bulk of the API) and leaves
edge-case shapes (multiple return values, output parameters, array
arguments, generic-method dispatch) for follow-up commits.  Functions whose
shape is not yet handled emit a ``// TODO`` placeholder so the diff against
the hand-written surface stays auditable.

Run from the repo root:

    python3 tools/codegen.py
"""

from __future__ import annotations

import json
import re
import sys
from dataclasses import dataclass
from pathlib import Path

HEADER_FILES = [
    "meos.h",
    "meos_catalog.h",
    "meos_geo.h",
    "meos_internal.h",
    "meos_internal_geo.h",
    "meos_npoint.h",
]

# Forward-declared opaque types we never wrap (mirrors the cdef skip list
# the PyMEOS-CFFI builder applies).
OPAQUE_TYPES = ("json_object", "GEOSContextHandle_t")

# C-name skip list: functions handled by hand or intentionally hidden.
SKIPPED_FUNCTIONS = {
    "py_error_handler",
    "meos_initialize_error_handler",
}


# Type mapping ----------------------------------------------------------

@dataclass
class TypeMapping:
    go_type: str           # Go-side parameter / return type
    c_cast: str | None     # ``C.<cast>(x)`` template, ``{}`` is the Go expr
    from_c: str | None     # Convert a C result to Go (``{}`` is the C expr)


# Scalars and well-known opaque pointers.  Pointer types map to the package's
# wrapper structs (``*STBox``, ``*Geom``, etc.) where the wrapper exposes an
# ``Inner()`` method returning the C pointer.
TYPE_MAP: dict[str, TypeMapping] = {
    "void":              TypeMapping("",         None,                                None),
    "bool":              TypeMapping("bool",     "C.bool({})",                        "bool({})"),
    "int":               TypeMapping("int",      "C.int({})",                         "int({})"),
    "int8":              TypeMapping("int8",     "C.int8({})",                        "int8({})"),
    "int16":             TypeMapping("int16",    "C.int16({})",                       "int16({})"),
    "int32":             TypeMapping("int32",    "C.int32({})",                       "int32({})"),
    "int32_t":           TypeMapping("int32",    "C.int32_t({})",                     "int32({})"),
    "int64":             TypeMapping("int64",    "C.int64({})",                       "int64({})"),
    "uint8":             TypeMapping("uint8",    "C.uint8({})",                       "uint8({})"),
    "uint8_t":           TypeMapping("uint8",    "C.uint8_t({})",                     "uint8({})"),
    "uint16":            TypeMapping("uint16",   "C.uint16({})",                      "uint16({})"),
    "uint32":            TypeMapping("uint32",   "C.uint32({})",                      "uint32({})"),
    "uint64":            TypeMapping("uint64",   "C.uint64({})",                      "uint64({})"),
    "double":            TypeMapping("float64",  "C.double({})",                      "float64({})"),
    "size_t":            TypeMapping("uint",     "C.size_t({})",                      "uint({})"),
    "DateADT":           TypeMapping("int32",    "C.DateADT({})",                     "int32({})"),
    "Timestamp":         TypeMapping("int64",    "C.Timestamp({})",                   "int64({})"),
    "TimestampTz":       TypeMapping("int64",    "C.TimestampTz({})",                 "int64({})"),
    "TimeOffset":        TypeMapping("int64",    "C.TimeOffset({})",                  "int64({})"),
    "interpType":        TypeMapping("Interpolation", "C.interpType({})",             "Interpolation({})"),
    "meosType":          TypeMapping("MeosType", "C.meosType({})",                    "MeosType({})"),
    "meosOper":          TypeMapping("MeosOper", "C.meosOper({})",                    "MeosOper({})"),
    "tempSubtype":       TypeMapping("TempSubtype", "C.tempSubtype({})",              "TempSubtype({})"),
    "errorLevel":        TypeMapping("ErrorLevel", "C.errorLevel({})",                "ErrorLevel({})"),
    "char *":            TypeMapping("string",   "C.CString({})",                     "C.GoString({})"),
    "const char *":      TypeMapping("string",   "C.CString({})",                     "C.GoString({})"),
    # PostgreSQL ``text`` is a varlena envelope around a Go string.  The
    # hand-written ``cstring2text`` / ``text2cstring`` helpers handle the
    # palloc/pfree, mirroring the same idiom used by the hand-written
    # surface.
    "text *":            TypeMapping("string",   "cstring2text({})",                  "text2cstring({})"),
    "const text *":      TypeMapping("string",   "cstring2text({})",                  "text2cstring({})"),
}

# Opaque MEOS struct pointers mapped to GoMEOS wrapper types.  When the
# pointer is an input, the caller passes the wrapper and the codegen emits
# ``.Inner()``.  When it is a return, the codegen wraps the pointer in a new
# instance.  ``Temporal`` is the interface type and uses ``CreateTemporal``.
WRAPPER_TYPES: dict[str, tuple[str, str]] = {
    # C struct name (without trailing space/star) -> (Go wrapper type, ctor expr).
    # The ctor expr uses a literal ``$res`` placeholder for the C result so we
    # do not collide with Python ``str.format`` parsing the embedded braces of
    # Go composite literals.
    "Temporal":          ("Temporal",          "CreateTemporal($res)"),
    "TInstant":          ("TInstant",          "TInstant{_inner: $res}"),
    "TSequence":         ("TSequence",         "TSequence{_inner: $res}"),
    "TSequenceSet":      ("TSequenceSet",      "TSequenceSet{_inner: $res}"),
    "STBox":             ("*STBox",            "&STBox{_inner: $res}"),
    "TBox":              ("*TBox",             "&TBox{_inner: $res}"),
    "Span":              ("*Span",             "&Span{_inner: $res}"),
    "SpanSet":           ("*SpanSet",          "&SpanSet{_inner: $res}"),
    "Set":               ("*Set",              "&Set{_inner: $res}"),
    "GSERIALIZED":       ("*Geom",             "&Geom{_inner: $res}"),
    "Interval":          ("timeutil.Timedelta", "IntervalToTimeDelta($res)"),
    "Npoint":            ("*Npoint",           "&Npoint{_inner: $res}"),
    "Nsegment":          ("*Nsegment",         "&Nsegment{_inner: $res}"),
    "SkipList":          ("*SkipList",         "&SkipList{_inner: $res}"),
    "RTree":             ("*RTree",            "&RTree{_inner: $res}"),
}


def _references_opaque(entry: dict) -> bool:
    if any(t in entry["returnType"]["c"] for t in OPAQUE_TYPES):
        return True
    return any(any(t in p["cType"] for t in OPAQUE_TYPES) for p in entry["params"])


def _is_datum_internal(entry: dict) -> bool:
    """Functions that take or return a raw Datum are MEOS-internal helpers
    that the hand-written surface re-exposes through typed overloads.  The
    codegen cannot produce those overloads automatically and reporting them
    as TODO would conflate metadata work with real shape gaps."""
    if "Datum" in entry["returnType"]["c"].split():
        return True
    for p in entry["params"]:
        if "Datum" in p["cType"].split():
            return True
    return False


def _strip_qualifiers(c_type: str) -> tuple[str, int]:
    """Return ``(base_type, pointer_level)`` stripped of ``const`` and ``*``."""
    s = c_type.replace("const ", "").strip()
    stars = s.count("*")
    return s.replace("*", "").strip(), stars


def _go_type_for(c_type: str) -> tuple[str | None, str | None, str | None]:
    """Look up a C type in the mapping tables.

    Returns ``(go_type, c_cast, from_c)``; any field may be ``None`` when the
    type is not yet handled.
    """
    if c_type in TYPE_MAP:
        m = TYPE_MAP[c_type]
        # Normalise scalar/string templates to the same ``$x`` placeholder
        # the wrapper-type entries use, so emit_function only deals with one
        # substitution dialect.
        c_cast = m.c_cast.replace("{}", "$x") if m.c_cast else None
        from_c = m.from_c.replace("{}", "$x") if m.from_c else None
        return m.go_type, c_cast, from_c
    base, stars = _strip_qualifiers(c_type)
    if stars == 1 and base in WRAPPER_TYPES:
        go_type, ctor = WRAPPER_TYPES[base]
        # Pass through the wrapper as an input; convert back as a return.
        c_cast = "$x.Inner()" if not go_type.startswith("*") else "$x._inner"
        if go_type == "Temporal":
            c_cast = "$x.Inner()"
        return go_type, c_cast, ctor.replace("$res", "$x")
    return None, None, None


# Name conversion -------------------------------------------------------

def _go_name(c_name: str) -> str:
    """Convert ``snake_case`` to ``PascalCase`` while keeping initialisms."""
    parts = c_name.split("_")
    out = []
    for p in parts:
        if not p:
            continue
        # Preserve common acronyms in upper case for readability.
        if p.upper() in {"WKB", "WKT", "MFJSON", "JSON", "SRID", "EWKT", "EWKB",
                         "STBOX", "TBOX", "ID", "X", "Y", "Z", "T", "MFJ"}:
            out.append(p.upper())
        else:
            out.append(p[:1].upper() + p[1:])
    return "".join(out) or c_name


_GO_RESERVED = {"type", "func", "interface", "select", "case", "chan", "goto",
                "package", "import", "go", "defer", "return", "range", "var",
                "const", "for", "if", "else", "switch", "break", "continue",
                "default", "fallthrough", "map", "struct", "string"}


def _go_param_name(c_name: str) -> str:
    if c_name in _GO_RESERVED:
        return c_name + "_"
    return c_name


# Emission --------------------------------------------------------------

@dataclass
class EmittedFunc:
    name: str          # Go-side name
    code: str          # full Go source for the function
    skipped: bool      # true if emitted as a TODO stub


def emit_function(entry: dict) -> EmittedFunc:
    c_name = entry["name"]
    go_name = _go_name(c_name)
    return_c = entry["returnType"]["c"]

    # Resolve return type.
    ret_go, _, ret_from_c = _go_type_for(return_c)
    if ret_go is None and return_c != "void":
        return EmittedFunc(go_name, _todo_stub(c_name, "unsupported return type " + return_c), True)

    # Resolve each parameter.
    go_args = []
    inner_args = []
    deferred = []
    for p in entry["params"]:
        pname = _go_param_name(p["name"])
        ptype = p["cType"]
        go_t, c_cast, _ = _go_type_for(ptype)
        if go_t is None:
            return EmittedFunc(go_name, _todo_stub(c_name, "unsupported param " + ptype), True)
        go_args.append(f"{pname} {go_t}")
        if c_cast is None:
            inner_args.append(pname)
        else:
            cast_expr = c_cast.replace("$x", pname)
            # Strings need a single C.CString allocation reused across the
            # call and the deferred free; bind to a local to avoid leaking.
            if go_t == "string":
                local = f"_c_{pname}"
                deferred.append(f"{local} := {cast_expr}")
                deferred.append(f"defer C.free(unsafe.Pointer({local}))")
                inner_args.append(local)
            else:
                inner_args.append(cast_expr)

    sig_args = ", ".join(go_args)
    ret_sig = "" if return_c == "void" else f" {ret_go}"
    call = f"C.{c_name}({', '.join(inner_args)})"

    body_lines = []
    # ``deferred`` emits as-is so the ordering of declarations (local bind
    # then defer) is preserved.
    body_lines.extend("\t" + d for d in deferred)
    if return_c == "void":
        body_lines.append(f"\t{call}")
    else:
        body_lines.append(f"\tres := {call}")
        body_lines.append(f"\treturn {ret_from_c.replace('$x', 'res')}")

    code = (
        f"// {go_name} wraps MEOS C function {c_name}.\n"
        f"func {go_name}({sig_args}){ret_sig} {{\n"
        + "\n".join(body_lines)
        + "\n}\n"
    )
    return EmittedFunc(go_name, code, False)


def _todo_stub(c_name: str, reason: str) -> str:
    return (
        f"// TODO {c_name}: {reason}\n"
        f"// func {_go_name(c_name)}(...) {{ /* not yet handled by codegen */ }}\n"
    )


# Driver ----------------------------------------------------------------

# Cgo preamble lives in its own file (_cgo.go) so the generated per-header
# files can share it without duplicating #include directives.
_CGO_FILE = """package generated

/*
#cgo darwin CFLAGS: -I/opt/homebrew/include
#cgo darwin LDFLAGS: -L/opt/homebrew/lib -lmeos -Wl,-rpath,/opt/homebrew/lib

#cgo linux CFLAGS: -I/usr/local/include/
#cgo linux LDFLAGS: -L/usr/local/lib -lmeos -Wl,-rpath,/usr/local/lib

#include <stddef.h>
#include "meos.h"
#include "meos_catalog.h"
#include "meos_geo.h"
#include "meos_internal.h"
#include "meos_internal_geo.h"
#include "meos_npoint.h"
*/
import "C"
"""

_PER_HEADER_PREAMBLE = """package generated

// #include <stddef.h>
import "C"
import (
\t"unsafe"

\t"github.com/leekchan/timeutil"
)

var _ = unsafe.Pointer(nil)
var _ = timeutil.Timedelta{}

"""


def generate(idl_path: Path, out_dir: Path) -> dict:
    idl = json.loads(idl_path.read_text())
    out_dir.mkdir(parents=True, exist_ok=True)

    entries_by_file: dict[str, list[dict]] = {h: [] for h in HEADER_FILES}
    for entry in idl["functions"]:
        if entry["file"] in entries_by_file:
            entries_by_file[entry["file"]].append(entry)

    stats = {"emitted": 0, "skipped": 0, "datum": 0, "by_header": {}}
    for header in HEADER_FILES:
        emitted_funcs = []
        local_emit = local_skip = local_datum = 0
        for entry in entries_by_file[header]:
            if entry["name"] in SKIPPED_FUNCTIONS:
                continue
            if _references_opaque(entry):
                continue
            if _is_datum_internal(entry):
                local_datum += 1
                continue
            ef = emit_function(entry)
            emitted_funcs.append(ef)
            if ef.skipped:
                local_skip += 1
            else:
                local_emit += 1
        stats["by_header"][header] = (local_emit, local_skip, local_datum)
        stats["emitted"] += local_emit
        stats["skipped"] += local_skip
        stats["datum"] += local_datum

        out_file = out_dir / f"meos_{header.replace('.h', '')}.go"
        body = _PER_HEADER_PREAMBLE + "\n\n".join(e.code for e in emitted_funcs) + "\n"
        out_file.write_text(body)

    # The single cgo preamble file with all #includes lives next to the
    # generated wrappers; Go's cgo machinery merges directives across files.
    (out_dir / "cgo.go").write_text(_CGO_FILE)
    return stats


if __name__ == "__main__":
    here = Path(__file__).parent
    # ``tools/preview/`` ships as a Draft artifact: the directory name is
    # prefixed with the standard Go-ignored ``_`` so ``go build ./...`` from
    # the repo root skips it.  Renaming to ``generated/`` is the staged step
    # that retires the hand-written wrappers once the remaining TODO shapes
    # are covered.
    stats = generate(here / "meos-idl.json", here / "_preview")
    print(f"Emitted {stats['emitted']} idiomatic wrappers")
    print(f"Skipped {stats['skipped']} as TODO (unsupported signature shape)")
    print(f"Excluded {stats['datum']} as Datum-bearing internal helpers")
    for h, (e, s, d) in stats["by_header"].items():
        print(f"  {h}: {e} emitted, {s} TODO, {d} Datum")
