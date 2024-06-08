//go:build ruleguard
// +build ruleguard

package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

func helloWorld(m dsl.Matcher) {
	m.Match(`fmt.Println("HelloWorld")`).Report("HelloWorld発見")
}
