package main

import "strings"

var path string

func subMatch() {
	s := "Hello_world"
	println(strings.ReplaceAll(s, "_", "_")) // Hello_world
	part := "x"
	println(strings.ReplaceAll(path+"/local/bin/", part, part)) // /local/bin/
	println(strings.ReplaceAll(s, "_", ""))                     // Helloworld
}
