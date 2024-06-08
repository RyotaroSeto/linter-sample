//go:build ruleguard
// +build ruleguard

package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

func assignOp(m dsl.Matcher) {
	// This group defines 2 rules.
	//
	// Rules are executed in the order they're described, until
	// the first successfull match.
	//
	// This short-circuit logic allows us to write overlapping rules
	// without conflicts: only one of them will be executed.
	m.Match(`$x = $x + 1`).Report(`can rewrite as $x++`)
	m.Match(`$x = $x + $y`).Report(`can rewrite as $x += $y`)
}
