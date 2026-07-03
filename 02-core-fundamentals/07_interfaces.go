package main

import "fmt"

// Shape is satisfied implicitly: any type with an Area() float64 method
// satisfies Shape, with no "implements Shape" declaration anywhere.
type Shape interface {
	Area() float64
}

type circle struct{ radius float64 }

func (c circle) Area() float64 { return 3.14159 * c.radius * c.radius }

type square struct{ side float64 }

func (s square) Area() float64 { return s.side * s.side }

// Implementing fmt.Stringer (one method: String() string) is how a type
// controls its own %v/%s/Println formatting - the idiomatic pattern for
// "give me a readable representation of this type".
func (c circle) String() string { return fmt.Sprintf("circle(r=%.1f)", c.radius) }

// MyError is a custom error type. Satisfying the standard library's
// `error` interface only requires one method: Error() string.
type MyError struct{ Code int }

func (e *MyError) Error() string { return fmt.Sprintf("failed with code %d", e.Code) }

func mightFail(fail bool) error {
	if fail {
		var e *MyError // nil *MyError
		return e       // returning a nil pointer AS an error - see the gotcha below
	}
	return nil
}

func demoInterfaces() {
	fmt.Println("--- implicit satisfaction ---")
	shapes := []Shape{circle{radius: 2}, square{side: 3}}
	for _, s := range shapes {
		fmt.Printf("%v has area %.2f\n", s, s.Area()) // circle's String() kicks in automatically
	}

	fmt.Println()
	fmt.Println("--- any (the empty interface) + type assertions/switches ---")

	// `any` (alias for interface{}) can hold a value of ANY type - the
	// tradeoff is you've thrown away type information and must recover it.
	values := []any{42, "hello", 3.14, circle{radius: 1}}
	for _, v := range values {
		// A type switch is the idiomatic way to branch on the dynamic type.
		switch x := v.(type) {
		case int:
			fmt.Println("int:", x*2)
		case string:
			fmt.Println("string, uppercased length:", len(x))
		case float64:
			fmt.Println("float64:", x)
		default:
			fmt.Printf("something else: %v (%T)\n", x, x)
		}
	}

	// A single type assertion, with the safe "comma-ok" form - the
	// unchecked form (v.(int) without ", ok") panics if the assertion fails.
	var v any = "just a string"
	if n, ok := v.(int); ok {
		fmt.Println("was an int:", n)
	} else {
		fmt.Println("was not an int")
	}

	fmt.Println()
	fmt.Println("--- the typed-nil-in-an-interface trap ---")

	err := mightFail(true)
	// err is NOT nil here, even though the *MyError it wraps IS nil!
	// An interface value is a (type, value) pair internally; err holds
	// (type=*MyError, value=nil), and that pair itself isn't the nil
	// interface (which would be (type=nil, value=nil)).
	fmt.Println("err == nil?", err == nil) // false - surprising the first time you hit it
	if err != nil {
		fmt.Println("err is non-nil, so this branch runs, calling Error() on a nil *MyError:")
		// This actually panics inside Error() (e.Code dereferences a nil e) -
		// but fmt specifically recovers a nil-pointer panic from a
		// String()/Error() method and prints "<nil>" instead of crashing,
		// precisely because this situation is common enough to guard against.
		// Don't rely on that safety net elsewhere - fmt is unusual in providing it.
		fmt.Println(err)
	}
	// The fix: a function returning `error` should return a literal `nil`,
	// not a nil pointer of a concrete error type, unless it explicitly
	// means to signal "an error occurred, and it happens to be this
	// nil-valued *MyError" (rare, usually a bug).
}
