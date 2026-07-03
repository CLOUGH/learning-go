package vetkata

import "fmt"

// Describe has a format-verb bug: %d expects an integer, but Name is a
// string. The code compiles and runs (Printf-family functions can't be
// type-checked by the compiler - the format string is just data), but
// `go vet` specifically understands Printf-style verbs and will flag the
// mismatch. Find it and fix the verb.
func Describe(name string, age int) string {
	return fmt.Sprintf("%d is %d years old", name, age)
}
