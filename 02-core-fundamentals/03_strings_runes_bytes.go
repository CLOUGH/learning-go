package main

import (
	"fmt"
	"strings"
)

func demoStringsRunesBytes() {
	fmt.Println("--- strings are immutable UTF-8 byte sequences ---")

	s := "héllo, 世界"

	// len() counts BYTES, not characters. "é" and each CJK character take
	// more than one byte in UTF-8, so byte length != character count.
	fmt.Println("len(s) in bytes:", len(s))

	// A plain indexed loop walks byte-by-byte, which can split a
	// multi-byte character in the middle - almost never what you want for
	// non-ASCII text.
	fmt.Print("byte-by-byte (wrong for non-ASCII): ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d ", s[i])
	}
	fmt.Println()

	// `range` over a string decodes UTF-8 for you: each iteration gives
	// you a rune (a Unicode code point, type int32) and the BYTE index it
	// started at - this is the correct way to iterate character by character.
	fmt.Print("range over string (correct): ")
	runeCount := 0
	for i, r := range s {
		fmt.Printf("[%d]=%c ", i, r)
		runeCount++
	}
	fmt.Println()
	fmt.Println("character count via range:", runeCount)

	// Converting explicitly between string, []rune, and []byte.
	runes := []rune(s)
	fmt.Println("as []rune, len =", len(runes)) // character count
	bytes := []byte(s)
	fmt.Println("as []byte, len =", len(bytes)) // same as len(s)
	backToString := string(runes)
	fmt.Println("back to string:", backToString)

	// A single-quoted literal like 'A' is a rune constant (an int32),
	// not a one-character string.
	var r rune = 'A'
	fmt.Printf("'A' as a rune is the integer %d\n", r)

	// Strings are immutable: there is no way to change a byte of `s` in
	// place. Building up a string via += in a loop reallocates every
	// time; strings.Builder avoids that by growing one buffer internally.
	var b strings.Builder
	for i := 0; i < 3; i++ {
		fmt.Fprintf(&b, "part%d-", i)
	}
	fmt.Println("built with strings.Builder:", b.String())
}
