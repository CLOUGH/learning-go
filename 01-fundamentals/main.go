package main

import "fmt"

// A struct: a typed bundle of fields. No classes, no inheritance.
type Point struct {
	X, Y int
}

// A method: a function with a receiver. Value receiver -> operates on a copy.
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// A pointer receiver -> can mutate the original.
func (p *Point) Scale(factor int) {
	p.X *= factor
	p.Y *= factor
}

// Interfaces are satisfied implicitly: anything with a String() string
// method satisfies fmt.Stringer, with no declared relationship to it.
type Shape interface {
	Area() float64
}

type Rectangle struct{ Width, Height float64 }

func (r Rectangle) Area() float64 { return r.Width * r.Height }

// Errors are ordinary return values.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide: cannot divide %v by zero", a)
	}
	return a / b, nil
}

func main() {
	// Variables, slices, maps.
	nums := []int{1, 2, 3, 4, 5}
	total := 0
	for _, n := range nums {
		total += n
	}
	fmt.Println("sum:", total)

	ages := map[string]int{"alice": 30, "bob": 25}
	if age, ok := ages["alice"]; ok {
		fmt.Println("alice is", age)
	}

	// Structs and methods.
	p := Point{X: 1, Y: 2}
	fmt.Println("point:", p.String())
	p.Scale(3)
	fmt.Println("scaled:", p.String())

	// Interfaces.
	var s Shape = Rectangle{Width: 3, Height: 4}
	fmt.Println("area:", s.Area())

	// Errors as values.
	if result, err := divide(10, 0); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", result)
	}

	// Closures: functions that capture variables from their surrounding
	// scope BY REFERENCE. This is the mechanism you must understand before
	// goroutines — a goroutine started from a closure sees whatever the
	// captured variable's value is when the goroutine actually runs, not
	// when `go` was called.
	counter := 0
	increment := func() int {
		counter++ // captures `counter` from main's scope
		return counter
	}
	fmt.Println(increment(), increment(), increment()) // 1 2 3
	fmt.Println("counter is now", counter)             // 3, mutated via the closure
}

/*
Expected output:

sum: 15
alice is 30
point: (1, 2)
scaled: (3, 6)
area: 12
error: divide: cannot divide 10 by zero
1 2 3
counter is now 3
*/
