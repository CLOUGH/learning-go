package main

import "fmt"

type point struct{ X, Y int }

// Go is always pass-by-value: this function receives a COPY of the
// struct, so mutating it has no effect on the caller's variable.
func moveByValue(p point) {
	p.X += 10
}

// A pointer receiver/parameter lets a function reach back into the
// caller's actual variable instead of a copy.
func moveByPointer(p *point) {
	p.X += 10
}

func demoPointers() {
	fmt.Println("--- pointers: & takes an address, * dereferences it ---")

	p := point{X: 1, Y: 2}

	moveByValue(p)
	fmt.Println("after moveByValue, p is unchanged:", p)

	moveByPointer(&p)
	fmt.Println("after moveByPointer, p.X actually changed:", p)

	// new(T) allocates a zeroed T and returns a *T - rarely used directly
	// in idiomatic Go (a `&T{}` literal is more common), but worth knowing.
	pp := new(point)
	pp.X = 5 // Go automatically dereferences for field access: sugar for (*pp).X
	fmt.Println("new(point):", *pp)

	// A nil pointer holds no address. Dereferencing one panics - this is
	// Go's version of a null pointer exception, and the most common
	// reason for a "invalid memory address or nil pointer dereference"
	// panic you'll see in a stack trace.
	var nilPtr *point
	fmt.Println("nil pointer itself is printable:", nilPtr)
	safeDeref := func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("recovered: %v", r)
			}
		}()
		fmt.Println(nilPtr.X) // this line panics
		return nil
	}
	if err := safeDeref(); err != nil {
		fmt.Println(err)
	}

	// Rule of thumb: use a pointer receiver/parameter when the function
	// needs to mutate the caller's data, or the value is large enough
	// that copying it is wasteful. Use a plain value when the data is
	// small and immutability is desirable (harder to reason about
	// accidental mutation otherwise).
}

/*
Expected output (from demoPointers, called via main.go):

--- pointers: & takes an address, * dereferences it ---
after moveByValue, p is unchanged: {1 2}
after moveByPointer, p.X actually changed: {11 2}
new(point): {5 0}
nil pointer itself is printable: <nil>
recovered: runtime error: invalid memory address or nil pointer dereference
*/
