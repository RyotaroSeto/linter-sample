//go:build ruleguard
// +build ruleguard

package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

func dupArg(m dsl.Matcher) {
	m.Match(`strings.ReplaceAll($_, $x, $x)`).
		Report(`suspicious duplicated arg $x`)
}
