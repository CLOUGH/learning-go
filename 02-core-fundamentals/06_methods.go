package main

import "fmt"

type Counter struct{ n int }

// Value receiver: gets a COPY of the Counter. Fine for reading.
func (c Counter) Value() int { return c.n }

// Pointer receiver: gets a pointer to the actual Counter. Required for
// anything that mutates it.
func (c *Counter) Inc() { c.n++ }

type Incrementer interface {
	Inc()
}

func demoMethodsAndReceivers() {
	fmt.Println("--- value vs. pointer receivers ---")

	c := Counter{}
	c.Inc() // Go automatically takes &c here, because c is an addressable variable
	c.Inc()
	fmt.Println("count after two Inc():", c.Value())

	fmt.Println()
	fmt.Println("--- the method-set gotcha: interfaces care about the difference ---")

	// *Counter satisfies Incrementer (Inc has a pointer receiver, and
	// &c is a *Counter).
	var inc Incrementer = &c
	inc.Inc()
	fmt.Println("via interface, count is now:", c.Value())

	// A plain Counter VALUE does NOT satisfy Incrementer - its method set
	// only contains value-receiver methods (just Value(), not Inc()).
	// Uncommenting the next line is a compile error:
	//
	//   var inc2 Incrementer = Counter{}
	//   // cannot use Counter{} (value of type Counter) as Incrementer value:
	//   // Counter does not implement Incrementer (method Inc has pointer receiver)
	//
	// This is why you'll see struct types in real Go code almost always
	// used consistently as either "always by pointer" or "always by
	// value" - mixing the two runs headfirst into this rule the moment
	// interfaces get involved.
	fmt.Println("(see the comment above for the compile error this would cause)")
}

/*
Expected output (from demoMethodsAndReceivers, called via main.go):

--- value vs. pointer receivers ---
count after two Inc(): 2

--- the method-set gotcha: interfaces care about the difference ---
via interface, count is now: 3
(see the comment above for the compile error this would cause)
*/
