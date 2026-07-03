package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func demoOSIOBufio() {
	fmt.Println("--- os: environment, arguments, and process-level stuff ---")

	fmt.Println("os.Args[0] (the running binary's path):", truncate(os.Args[0]))

	if home, ok := os.LookupEnv("HOME"); ok {
		fmt.Println("HOME is set (value not printed - environment-dependent)")
		_ = home
	} else {
		fmt.Println("HOME is not set in this environment")
	}
	// os.Exit(code) would terminate the process immediately - no deferred
	// functions run. Reserved for main()'s very last word, not used here
	// since it would kill the rest of this demo program.

	fmt.Println()
	fmt.Println("--- io.Reader / io.Writer: the two interfaces everything else builds on ---")

	// strings.NewReader turns a string into an io.Reader - handy for demos
	// and tests, since it behaves exactly like reading from a file or
	// network connection would, without needing either.
	reader := strings.NewReader("line one\nline two\nline three\n")

	// io.Copy reads from a Reader and writes to a Writer until EOF - the
	// same function works whether the source is a file, a network
	// connection, or (as here) an in-memory string.
	var buf bytes.Buffer // implements io.Writer (and io.Reader)
	n, err := io.Copy(&buf, reader)
	fmt.Println("io.Copy copied", n, "bytes, err:", err)

	fmt.Println()
	fmt.Println("--- bufio.Scanner: the idiomatic way to read line by line ---")

	scanner := bufio.NewScanner(strings.NewReader(buf.String()))
	lineNum := 1
	for scanner.Scan() { // advances to the next token (line, by default)
		fmt.Printf("line %d: %q\n", lineNum, scanner.Text())
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("scan error:", err)
	}

	fmt.Println()
	fmt.Println("--- real file I/O: os.CreateTemp, WriteFile, ReadFile ---")

	tmp, err := os.CreateTemp("", "learn-go-demo-*.txt")
	if err != nil {
		fmt.Println("CreateTemp failed:", err)
		return
	}
	path := tmp.Name()
	defer os.Remove(path) // clean up - this is a scratch file for the demo only
	tmp.Close()

	content := []byte("written by os.WriteFile\n")
	if err := os.WriteFile(path, content, 0o644); err != nil {
		fmt.Println("WriteFile failed:", err)
		return
	}

	readBack, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("ReadFile failed:", err)
		return
	}
	fmt.Printf("read back: %q\n", string(readBack))
}

/*
Expected output (from demoOSIOBufio, called via main.go):

--- os: environment, arguments, and process-level stuff ---
os.Args[0] (the running binary's path): .../11-standard-library
HOME is set (value not printed - environment-dependent)

--- io.Reader / io.Writer: the two interfaces everything else builds on ---
io.Copy copied 29 bytes, err: <nil>

--- bufio.Scanner: the idiomatic way to read line by line ---
line 1: "line one"
line 2: "line two"
line 3: "line three"

--- real file I/O: os.CreateTemp, WriteFile, ReadFile ---
read back: "written by os.WriteFile\n"

The os.Args[0] line and whether HOME is set are inherently
machine/environment-dependent - the exact path and env result will differ
from what's shown here depending on where and how you run this.
*/

func truncate(s string) string {
	if len(s) > 20 {
		return "..." + s[len(s)-20:]
	}
	return s
}
