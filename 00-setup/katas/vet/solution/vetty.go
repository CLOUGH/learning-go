package vetkata

import "fmt"

// The fix: %d expected an integer but was matched against the string
// `name`. Swap in %s for the string and keep %d for the int.
func Describe(name string, age int) string {
	return fmt.Sprintf("%s is %d years old", name, age)
}
