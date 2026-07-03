package main

import "fmt"

func main() {
	demoStringsStrconv()
	fmt.Println()

	demoOSIOBufio()
	fmt.Println()

	demoTime()
	fmt.Println()

	demoSort()
	fmt.Println()

	demoEncodingJSON()
	fmt.Println()

	demoNetHTTP()
	fmt.Println()

	demoLogSlog()
	fmt.Println()

	demoFlag()
}

/*
Expected output: `go run ./11-standard-library` prints each demoXxx()'s
own output in the order called above - see the bottom of
01_strings_strconv.go through 08_flag.go for each section's exact
output (two sections have real caveats: 02_os_io_bufio.go's os.Args/HOME
line is machine-dependent, and 07_log_slog.go's timestamps reflect
whenever you actually run it).
*/
