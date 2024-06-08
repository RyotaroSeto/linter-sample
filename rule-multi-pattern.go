//go:build ruleguard
// +build ruleguard

package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func boolLiteralInExpr(m dsl.Matcher) {
	// The following rule includes 4 patterns: 2 for each operator and bool value.
	// Every pattern in this list is called "an alternative".
	// Alternatives are tried out one by one, until we have a match.
	m.Match(`$x == true`,
		`$x != true`,
		`$x == false`,
		`$x != false`).
		Report(`omit bool literal in expression`)
}
