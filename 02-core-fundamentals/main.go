package main

import (
	"fmt"

	"learn-go/02-core-fundamentals/shapes"
)

// init() runs automatically before main(), with no explicit call needed -
// commonly used for one-time setup (registering things, validating
// configuration). A package can have multiple init() funcs, even
// multiple in the same file; they run in the order they appear.
func init() {
	fmt.Println("init() ran before main() - no one called it explicitly")
}

func main() {
	demoVariablesAndTypes()
	fmt.Println()

	demoArraysSlicesMaps()
	fmt.Println()

	demoStringsRunesBytes()
	fmt.Println()

	demoPointers()
	fmt.Println()

	demoStructsEmbedding()
	fmt.Println()

	demoMethodsAndReceivers()
	fmt.Println()

	demoInterfaces()
	fmt.Println()

	demoGenerics()
	fmt.Println()

	demoIterators()
	fmt.Println()

	demoErrors()
	fmt.Println()

	demoDeferPanicRecover()
	fmt.Println()

	fmt.Println("--- packages: exported vs. unexported (see shapes/shapes.go) ---")
	c := shapes.NewCircle(2) // NewCircle and Circle are exported - reachable from here
	fmt.Println("circle area via exported API:", c.Area())
	// c.piApprox or shapes.piApprox would NOT compile - piApprox is
	// unexported, so it's invisible outside package shapes.
}
