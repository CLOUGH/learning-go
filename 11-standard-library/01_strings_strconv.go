package main

import (
	"fmt"
	"strconv"
	"strings"
)

func demoStringsStrconv() {
	fmt.Println("--- strings: the toolbox for almost everything text-shaped ---")

	s := "  Hello, Gophers!  "

	fmt.Println("Contains \"Gopher\":", strings.Contains(s, "Gopher"))
	fmt.Println("HasPrefix \"  Hello\":", strings.HasPrefix(s, "  Hello"))
	fmt.Println("HasSuffix \"!  \":", strings.HasSuffix(s, "!  "))
	fmt.Println("Index of \"Gopher\":", strings.Index(s, "Gopher"))

	trimmed := strings.TrimSpace(s)
	fmt.Printf("TrimSpace: %q\n", trimmed)

	fmt.Println("ToUpper:", strings.ToUpper(trimmed))
	fmt.Println("ToLower:", strings.ToLower(trimmed))
	fmt.Println("Replace one:", strings.Replace(trimmed, "o", "0", 1))
	fmt.Println("ReplaceAll:", strings.ReplaceAll(trimmed, "o", "0"))

	parts := strings.Split("a,b,,c", ",")
	fmt.Printf("Split on comma: %q (len=%d - note the empty string between b and c)\n", parts, len(parts))

	fmt.Println("Join with \" | \":", strings.Join(parts, " | "))

	// Fields splits on whitespace AND collapses runs of it - the usual
	// choice for "split this sentence into words", where Split would leave
	// you with empty strings for every extra space.
	fmt.Printf("Fields: %q\n", strings.Fields("  lots   of   space  "))

	var b strings.Builder
	for i := 0; i < 3; i++ {
		fmt.Fprintf(&b, "%d-", i)
	}
	fmt.Println("Builder result:", b.String())

	fmt.Println()
	fmt.Println("--- strconv: converting between strings and other basic types ---")

	n, err := strconv.Atoi("42")
	fmt.Println("Atoi(\"42\"):", n, err)

	_, err = strconv.Atoi("not-a-number")
	fmt.Println("Atoi(\"not-a-number\") error:", err)

	fmt.Println("Itoa(42):", strconv.Itoa(42))

	f, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println("ParseFloat(\"3.14\"):", f)

	ok, _ := strconv.ParseBool("true")
	fmt.Println("ParseBool(\"true\"):", ok)

	fmt.Println("FormatInt(255, 16):", strconv.FormatInt(255, 16)) // hex
	fmt.Println("Quote:", strconv.Quote("line1\nline2"))           // escapes control characters
}

/*
Expected output (from demoStringsStrconv, called via main.go):

--- strings: the toolbox for almost everything text-shaped ---
Contains "Gopher": true
HasPrefix "  Hello": true
HasSuffix "!  ": true
Index of "Gopher": 9
TrimSpace: "Hello, Gophers!"
ToUpper: HELLO, GOPHERS!
ToLower: hello, gophers!
Replace one: Hell0, Gophers!
ReplaceAll: Hell0, G0phers!
Split on comma: ["a" "b" "" "c"] (len=4 - note the empty string between b and c)
Join with " | ": a | b |  | c
Fields: ["lots" "of" "space"]
Builder result: 0-1-2-

--- strconv: converting between strings and other basic types ---
Atoi("42"): 42 <nil>
Atoi("not-a-number") error: strconv.Atoi: parsing "not-a-number": invalid syntax
Itoa(42): 42
ParseFloat("3.14"): 3.14
ParseBool("true"): true
FormatInt(255, 16): ff
Quote: "line1\nline2"
*/
