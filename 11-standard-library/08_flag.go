package main

import (
	"flag"
	"fmt"
)

func demoFlag() {
	fmt.Println("--- flag: command-line flag parsing ---")

	// In a real command-line program, you'd declare package-level flags
	// and call flag.Parse() once in main(), which reads os.Args for you:
	//
	//   var name = flag.String("name", "world", "who to greet")
	//   flag.Parse()
	//   fmt.Println("hello,", *name)
	//
	// Since this demo runs as one function among several (and calling
	// flag.Parse() reads the REAL os.Args of whatever process is running
	// this whole lesson), it uses flag.NewFlagSet with an explicit slice
	// of arguments instead - the same API, pointed at args you supply
	// rather than the process's actual command line. This is also exactly
	// how you'd unit-test code that uses flags.
	fs := flag.NewFlagSet("demo", flag.ContinueOnError)
	name := fs.String("name", "world", "who to greet")
	count := fs.Int("count", 1, "how many times to greet")
	verbose := fs.Bool("verbose", false, "print extra detail")

	args := []string{"-name=gopher", "-count=3", "-verbose"}
	if err := fs.Parse(args); err != nil {
		fmt.Println("parse error:", err)
		return
	}

	fmt.Println("parsed -name:", *name)
	fmt.Println("parsed -count:", *count)
	fmt.Println("parsed -verbose:", *verbose)

	for i := 0; i < *count; i++ {
		if *verbose {
			fmt.Printf("[%d/%d] hello, %s!\n", i+1, *count, *name)
		} else {
			fmt.Println("hello,", *name)
		}
	}
}

/*
Expected output (from demoFlag, called via main.go):

--- flag: command-line flag parsing ---
parsed -name: gopher
parsed -count: 3
parsed -verbose: true
[1/3] hello, gopher!
[2/3] hello, gopher!
[3/3] hello, gopher!
*/
