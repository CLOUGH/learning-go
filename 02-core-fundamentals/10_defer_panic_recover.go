package main

import "fmt"

func deferOrder() {
	fmt.Println("defer runs LIFO - last deferred, first executed:")
	for i := 1; i <= 3; i++ {
		defer fmt.Println("deferred:", i)
	}
	fmt.Println("function body finishing")
	// Output order: body finishes first, then "deferred: 3", "deferred: 2", "deferred: 1"
}

func deferArgsEvaluatedImmediately() {
	x := 1
	defer fmt.Println("deferred saw x =", x) // x is captured/evaluated NOW, as 1
	x = 2
	fmt.Println("x is now", x, "but the defer above already locked in the old value")
}

func deferWithClosureSeesLatestValue() {
	x := 1
	defer func() { fmt.Println("deferred closure saw x =", x) }() // reads x when the closure RUNS, not when deferred
	x = 2
	fmt.Println("x is now", x)
}

// riskyDivide recovers from a divide-by-zero-style panic and turns it
// into a normal error return - this is the standard shape for "convert a
// panic at a boundary into an error the caller can handle normally".
// recover() only has any effect when called directly inside a deferred
// function; calling it anywhere else just returns nil.
func riskyDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()
	result = a / b // panics if b == 0: "runtime error: integer divide by zero"
	return result, nil
}

func demoDeferPanicRecover() {
	fmt.Println("--- defer: LIFO order ---")
	deferOrder()

	fmt.Println()
	fmt.Println("--- defer arguments are evaluated when `defer` runs, not later ---")
	deferArgsEvaluatedImmediately()

	fmt.Println()
	fmt.Println("--- unless the deferred call is a closure, which reads live variables ---")
	deferWithClosureSeesLatestValue()

	fmt.Println()
	fmt.Println("--- panic + recover: converting a panic into a normal error ---")
	if result, err := riskyDivide(10, 2); err == nil {
		fmt.Println("10 / 2 =", result)
	}
	if _, err := riskyDivide(10, 0); err != nil {
		fmt.Println("10 / 0 ->", err)
	}
	fmt.Println("program kept running - the panic never escaped riskyDivide")
}

/*
Expected output (from demoDeferPanicRecover, called via main.go):

--- defer: LIFO order ---
defer runs LIFO - last deferred, first executed:
function body finishing
deferred: 3
deferred: 2
deferred: 1

--- defer arguments are evaluated when `defer` runs, not later ---
x is now 2 but the defer above already locked in the old value
deferred saw x = 1

--- unless the deferred call is a closure, which reads live variables ---
x is now 2
deferred closure saw x = 2

--- panic + recover: converting a panic into a normal error ---
10 / 2 = 5
10 / 0 -> recovered from panic: runtime error: integer divide by zero
program kept running - the panic never escaped riskyDivide
*/
