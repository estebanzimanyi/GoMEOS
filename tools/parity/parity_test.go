// Package parity is the language-independent mirror of
// tools/portable_parity.py — the GoMEOS portable bare-name parity gate
// (RFC #920; MEOS-API cross-repo handoff PR #9).
//
// A binding is done when its exposed symbol set ⊇ portableAliases.bareNames,
// verified with the same prefix logic as MEOS-API portable_parity.py: a
// bare name is backed iff some referenced MEOS symbol == bareName or
// startsWith(bareName + "_"), falling back to the contract's verified
// explicitBacking prefixes (nearestApproachDistance ← the nad_* family).
// 0 unbacked, no per-binding exceptions, across all six in-scope
// user-facing type families (temporal, geo, cbuffer, npoint, pose, rgeo)
// — cbuffer/npoint/pose/rgeo are never excluded from the parity headline.
//
// "Exposed symbol set" = every MEOS C function the CGO layer references
// (C.<symbol>( ) in the hand-written root package AND the IDL-driven
// generated surface under functions (emitted by tools/codegen.py
// from tools/meos-idl.json — the MEOS-API parser's JSON output). The
// operators' own backing functions, reused by construction, never
// reimplemented.
//
// This package has no `import "C"`, so the gate runs in CI with no
// libmeos/CGO toolchain, exactly like MEOS.NET's managed test mirror.
// Its verdict is identical to the Python script's, by construction.
package parity

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
)

// inScopeFamilies — full user-facing temporal type families.
// cbuffer/npoint/pose/rgeo are NOT internals and are never excluded from
// the parity headline.
var inScopeFamilies = []string{
	"temporal", "geo", "cbuffer", "npoint", "pose", "rgeo",
}

// bindingBacking is inert at MEOS 1.4 (canonical tdistance_* is present);
// retained as a documented safety net for older/partial header scans.
var bindingBacking = map[string][]string{
	"tdistance": {"distance_tfloat", "distance_tint",
		"distance_tnumber", "distance_tpoint"},
}

var cgoRe = regexp.MustCompile(`\bC\.([A-Za-z_]\w*)\s*\(`)

var cgoPseudo = map[string]bool{
	"CString": true, "CBytes": true, "GoString": true, "GoBytes": true,
	"free": true, "malloc": true, "calloc": true,
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatal("could not locate repo root (go.mod)")
		}
		dir = parent
	}
}

func scanDir(t *testing.T, dir string, syms map[string]bool) int {
	t.Helper()
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0 // optional source (e.g. functions absent)
	}
	files := 0
	for _, e := range entries {
		n := e.Name()
		if e.IsDir() || !strings.HasSuffix(n, ".go") ||
			strings.HasSuffix(n, "_test.go") {
			continue
		}
		b, err := os.ReadFile(filepath.Join(dir, n))
		if err != nil {
			t.Fatalf("read %s: %v", n, err)
		}
		for _, m := range cgoRe.FindAllStringSubmatch(string(b), -1) {
			if !cgoPseudo[m[1]] {
				syms[m[1]] = true
			}
		}
		files++
	}
	return files
}

// exposedSymbols = hand-written root *.go ∪ IDL-driven functions/*.go.
func exposedSymbols(t *testing.T, repo string) map[string]bool {
	t.Helper()
	syms := map[string]bool{}
	scanDir(t, repo, syms)
	scanDir(t, filepath.Join(repo, "functions"), syms)
	if len(syms) == 0 {
		t.Fatal("no CGO symbols found — repo layout changed?")
	}
	return syms
}

type pair struct{ op, bare, fam string }

func contract(t *testing.T, repo string) ([]pair, map[string][]string) {
	t.Helper()
	p := filepath.Join(repo, "tools", "portable-aliases.json")
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatalf("vendored portable-aliases SoT missing: %v", err)
	}
	var doc struct {
		Families map[string][]struct {
			Operator string `json:"operator"`
			BareName string `json:"bareName"`
		} `json:"families"`
		ExplicitBacking map[string][]string `json:"explicitBacking"`
	}
	if err := json.Unmarshal(b, &doc); err != nil {
		t.Fatalf("parse contract: %v", err)
	}
	var pairs []pair
	for fam, lst := range doc.Families {
		for _, e := range lst {
			pairs = append(pairs, pair{e.Operator, e.BareName, fam})
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].bare < pairs[j].bare
	})
	return pairs, doc.ExplicitBacking
}

func matches(symbols map[string]bool, prefix string) []string {
	var hits []string
	for s := range symbols {
		if s == prefix || strings.HasPrefix(s, prefix+"_") {
			hits = append(hits, s)
		}
	}
	return hits
}

func backing(bare string, symbols map[string]bool,
	explicit map[string][]string) (hits []string, via string) {
	if h := matches(symbols, bare); len(h) > 0 {
		return h, "prefix"
	}
	for _, pref := range explicit[bare] {
		hits = append(hits, matches(symbols, pref)...)
	}
	if len(hits) > 0 {
		return hits, "explicit:" + strings.Join(explicit[bare], ",")
	}
	for _, pref := range bindingBacking[bare] {
		hits = append(hits, matches(symbols, pref)...)
	}
	if len(hits) > 0 {
		return hits, "version-bridge:" +
			strings.Join(bindingBacking[bare], ",")
	}
	return nil, ""
}

func familyOf(name string) string {
	n := strings.ToLower(name)
	switch {
	case strings.Contains(n, "rgeo"):
		return "rgeo"
	case strings.Contains(n, "cbuffer"):
		return "cbuffer"
	case strings.Contains(n, "npoint"):
		return "npoint"
	case strings.Contains(n, "pose"):
		return "pose"
	case strings.Contains(n, "geo"), strings.Contains(n, "geom"),
		strings.Contains(n, "geog"), strings.Contains(n, "point"),
		strings.Contains(n, "spatial"):
		return "geo"
	default:
		return "temporal"
	}
}

// TestExposedApiSupersetOfPortableBareNames is the Go mirror of
// tools/portable_parity.py --check: the exposed CGO symbol set (root +
// IDL-driven functions) must be a superset of
// portableAliases.bareNames (29/29, 0 unbacked) AND every in-scope
// user-facing family must be covered. cbuffer/npoint/pose/rgeo are full
// types and are never excluded from the parity headline; a family present
// in the surface but unbacked (regressed) or absent (pending) is a hard
// failure.
func TestExposedApiSupersetOfPortableBareNames(t *testing.T) {
	repo := repoRoot(t)
	symbols := exposedSymbols(t, repo)
	pairs, explicit := contract(t, repo)

	if len(pairs) != 29 {
		t.Fatalf("contract must carry exactly 29 operator→bareName "+
			"pairs, got %d", len(pairs))
	}

	famsPresent := map[string]bool{}
	for s := range symbols {
		famsPresent[familyOf(s)] = true
	}

	famTotals := map[string]int{}
	for _, f := range inScopeFamilies {
		famTotals[f] = 0
	}
	var unbacked []string
	for _, p := range pairs {
		hits, via := backing(p.bare, symbols, explicit)
		if len(hits) == 0 {
			unbacked = append(unbacked,
				p.bare+" ("+p.op+", "+p.fam+")")
			continue
		}
		t.Logf("backed: %-24s %-5s via %s (%d symbols)",
			p.bare, p.op, via, len(hits))
		for _, h := range hits {
			if _, ok := famTotals[familyOf(h)]; ok {
				famTotals[familyOf(h)]++
			}
		}
	}

	if len(unbacked) != 0 {
		t.Fatalf("unbacked canonical bare names (exposed symbol set must "+
			"be a superset, 0 unbacked): %s",
			strings.Join(unbacked, ", "))
	}

	var regressed, pending []string
	for _, f := range inScopeFamilies {
		switch {
		case famTotals[f] > 0: // covered
		case famsPresent[f]:
			regressed = append(regressed, f)
		default:
			pending = append(pending, f)
		}
	}
	if len(regressed) != 0 {
		t.Fatalf("in-scope families present in the surface but unbacked "+
			"(cbuffer/npoint/pose/rgeo are never excluded from the "+
			"parity headline): %s; coverage=%v",
			strings.Join(regressed, ", "), famTotals)
	}
	if len(pending) != 0 {
		t.Fatalf("in-scope user-facing families absent from the exposed "+
			"surface (never excluded from the parity headline): %s; "+
			"coverage=%v", strings.Join(pending, ", "), famTotals)
	}
	t.Logf("PASS: 29/29 bare names backed, 0 unbacked, all six in-scope "+
		"families covered: %v", famTotals)
}
