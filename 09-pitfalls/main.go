package main

import (
	"fmt"
	"sync"
	"time"
)

// --- 1. Goroutine leaks ---

func leakyGoroutine() {
	fmt.Println("--- goroutine leak: this goroutine blocks forever ---")
	ch := make(chan int) // nobody will ever send on this
	go func() {
		<-ch // blocks here for the rest of the program's life - leaked
		fmt.Println("this line never runs")
	}()
	fmt.Println("leaked one goroutine (it will never be cleaned up)")
}

func fixedWithDone() {
	fmt.Println("--- fixed: goroutine also selects on a done channel ---")
	ch := make(chan int)
	done := make(chan struct{})

	go func() {
		select {
		case <-ch:
			fmt.Println("got a value")
		case <-done:
			fmt.Println("told to stop, exiting cleanly instead of leaking")
		}
	}()

	close(done) // give the goroutine a way out
	time.Sleep(10 * time.Millisecond)
}

// --- 2. Loop variable capture ---

func loopVariableGo122Plus() {
	fmt.Println("--- loop variable capture: Go 1.22+ gives each iteration its own copy ---")
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("captured i:", i) // correct on Go 1.22+: prints 0,1,2 in some order
		}()
	}
	wg.Wait()
}

func loopVariableOldStyleWorkaround() {
	fmt.Println("--- the pre-1.22 workaround, still common in the wild: pass i as a parameter ---")
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) { // shadows the loop var with a fresh parameter - works on ANY Go version
			defer wg.Done()
			fmt.Println("captured i:", i)
		}(i)
	}
	wg.Wait()
}

// --- 3. Deadlock ---

func demonstrateDeadlock() {
	fmt.Println("--- deadlock: unbuffered send with no receiver ---")
	ch := make(chan int)
	ch <- 1 // no other goroutine will ever receive this - the runtime detects
	// the whole program is stuck and crashes with:
	// fatal error: all goroutines are asleep - deadlock!
}

// --- 4. Panics from misusing closed channels ---

func closeOnlyOnceAndOnlyFromSender() {
	fmt.Println("--- sync.Once guards against double-close when ownership is unclear ---")
	ch := make(chan int)
	var closeOnce sync.Once
	closeSafely := func() { closeOnce.Do(func() { close(ch) }) }

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			closeSafely() // safe even though 3 goroutines all call this
		}()
	}
	wg.Wait()
	fmt.Println("channel closed exactly once, no panic")
}

// --- 5. Forgetting wg.Done() ---

func forgottenDoneWouldDeadlock() {
	fmt.Println("--- always `defer wg.Done()` immediately, so early returns/panics still count down ---")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done() // guarantees Done() runs no matter how this goroutine exits
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered from panic:", r)
			}
		}()
		panic("something went wrong mid-goroutine")
	}()
	wg.Wait()
	fmt.Println("Wait() returned - it didn't hang, because Done() was deferred")
}

func main() {
	leakyGoroutine()
	fixedWithDone()
	fmt.Println()

	loopVariableGo122Plus()
	loopVariableOldStyleWorkaround()
	fmt.Println()

	closeOnlyOnceAndOnlyFromSender()
	fmt.Println()

	forgottenDoneWouldDeadlock()
	fmt.Println()

	fmt.Println("(demonstrateDeadlock() is commented out below - it crashes the program on purpose)")
	// demonstrateDeadlock()
}
