package main

import (
	"fmt"
	"time"
)

func demoTime() {
	fmt.Println("--- time.Time: instants, and time.Duration: elapsed time ---")

	start := time.Now()
	time.Sleep(5 * time.Millisecond)
	elapsed := time.Since(start)
	fmt.Println("elapsed is at least 5ms:", elapsed >= 5*time.Millisecond)

	// Duration is just an int64 count of nanoseconds under the hood, with
	// arithmetic and comparison operators working exactly like any other
	// numeric type - and constants like time.Second make that arithmetic
	// readable instead of a wall of nanosecond digits.
	d := 90 * time.Minute
	fmt.Println("90 minutes as a Duration:", d)
	fmt.Println("...as hours (float):", d.Hours())
	fmt.Println("...as a Duration string again:", (2*time.Hour + 30*time.Minute).String())

	fmt.Println()
	fmt.Println("--- constructing and comparing specific times ---")

	launch := time.Date(2009, time.November, 10, 15, 0, 0, 0, time.UTC)
	fmt.Println("a specific moment:", launch)

	later := launch.Add(24 * time.Hour)
	fmt.Println("Before:", launch.Before(later))
	fmt.Println("After:", later.After(launch))
	fmt.Println("Sub gives back a Duration:", later.Sub(launch))

	fmt.Println()
	fmt.Println("--- formatting and parsing: Go's reference-time layout ---")

	// Go doesn't use strftime-style verbs (%Y-%m-%d). Instead, the layout
	// string IS an example of the format, applied to one specific
	// reference moment: Mon Jan 2 15:04:05 MST 2006 (1 2 3 4 5 6 7, if you
	// squint at the numbers). Whatever position "2006" sits in is where
	// the year goes; wherever "15" sits is where 24-hour-clock hours go.
	const layout = "2006-01-02 15:04:05"
	formatted := launch.Format(layout)
	fmt.Println("Format:", formatted)

	parsed, err := time.Parse(layout, formatted)
	fmt.Println("Parse round-trip equal:", parsed.Equal(launch), err)
}

/*
Expected output (from demoTime, called via main.go):

--- time.Time: instants, and time.Duration: elapsed time ---
elapsed is at least 5ms: true
90 minutes as a Duration: 1h30m0s
...as hours (float): 1.5
...as a Duration string again: 2h30m0s

--- constructing and comparing specific times ---
a specific moment: 2009-11-10 15:00:00 +0000 UTC
Before: true
After: true
Sub gives back a Duration: 24h0m0s

--- formatting and parsing: Go's reference-time layout ---
Format: 2009-11-10 15:00:00
Parse round-trip equal: true <nil>
*/
