package main

import "fmt"

func demoVariablesAndTypes() {
	fmt.Println("--- variables, zero values, const/iota, type conversion ---")

	// `var` with an explicit type, and `:=` with inference. Both are
	// idiomatic; `:=` is far more common inside function bodies.
	var name string = "gopher"
	age := 5

	// Every type has a zero value - there's no "undefined"/"uninitialized".
	var (
		zeroInt    int     // 0
		zeroFloat  float64 // 0
		zeroBool   bool    // false
		zeroString string  // "" (not nil)
	)
	fmt.Println("zero values:", zeroInt, zeroFloat, zeroBool, fmt.Sprintf("%q", zeroString))

	// The numeric type zoo. Pick int/float64 by default; reach for a
	// specific width (int32, uint8, ...) only when it matters - matching
	// an external format, saving memory in bulk, etc.
	var i8 int8 = 127  // -128..127
	var u8 uint8 = 255 // byte is an alias for uint8
	var f32 float32 = 3.14
	fmt.Println("sized types:", i8, u8, f32)

	// Go never converts types implicitly - not even int to float64.
	// Conversion is always explicit: T(value).
	var x int = 10
	var y float64 = float64(x) / 3
	fmt.Println("explicit conversion:", y)

	// const values known at compile time. iota gives you an auto-incrementing
	// counter inside a const block - the idiomatic way to build enums.
	type Weekday int
	const (
		Sunday  Weekday = iota // 0
		Monday                 // 1 (implicitly repeats "= iota")
		Tuesday                // 2
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println("iota enum:", Sunday, Wednesday, Saturday)

	// Untyped constants: `const Pi = 3.14159` has no fixed type until used;
	// it adapts to whatever context needs it (float32, float64, ...).
	const Pi = 3.14159
	var asFloat32 float32 = Pi
	var asFloat64 float64 = Pi
	fmt.Println("untyped constant adapts:", asFloat32, asFloat64)

	fmt.Println(name, "is", age, "years old")
}

/*
Expected output (from demoVariablesAndTypes, called via main.go):

--- variables, zero values, const/iota, type conversion ---
zero values: 0 0 false ""
sized types: 127 255 3.14
explicit conversion: 3.3333333333333335
iota enum: 0 3 6
untyped constant adapts: 3.14159 3.14159
gopher is 5 years old
*/
