package main

func main() {
	var cond bool
	var boolVar bool

	println(cond == true)
	println(cond == false)
	println(cond != true)
	println(cond != false)

	// No warnings for these:
	println(cond == boolVar)
	println(cond == boolVar)
	println(cond != boolVar)
	println(cond != boolVar)
}
