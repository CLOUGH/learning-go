package main

import (
	"errors"
	"fmt"
)

// A sentinel error: a specific error value callers can check for by
// identity. Exported sentinel errors (ErrNotFound-style) are a common
// stdlib pattern - see io.EOF, sql.ErrNoRows.
var ErrNotFound = errors.New("not found")

func lookup(id int) error {
	if id < 0 {
		// %w (not %v/%s) wraps the original error: it's still included in
		// the returned error's chain, recoverable via errors.Is/As, while
		// the message gains context about where it happened.
		return fmt.Errorf("lookup %d: %w", id, ErrNotFound)
	}
	return nil
}

// ValidationError is a custom error TYPE (as opposed to a sentinel VALUE) -
// used when the caller needs structured data out of the error, not just
// "which error was it".
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %q: %s", e.Field, e.Msg)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Msg: "must be non-negative"}
	}
	return nil
}

func processUser(age int) error {
	if err := validateAge(age); err != nil {
		return fmt.Errorf("processing user: %w", err)
	}
	return nil
}

func demoErrors() {
	fmt.Println("--- sentinel errors + errors.Is ---")

	err := lookup(-1)
	fmt.Println("error message:", err)
	// errors.Is walks the wrapped chain looking for a match by identity -
	// this is why it finds ErrNotFound even though `err`'s own message is
	// "lookup -1: not found", a different string/value entirely.
	fmt.Println("errors.Is(err, ErrNotFound):", errors.Is(err, ErrNotFound))
	fmt.Println("errors.Is(err, io/whatever unrelated err):", errors.Is(err, errors.New("not found"))) // false: different identity, same text

	fmt.Println()
	fmt.Println("--- custom error types + errors.As ---")

	err = processUser(-5)
	fmt.Println("error message:", err)

	// errors.As walks the chain looking for an error that matches the
	// TYPE of the target pointer, and if found, assigns it - this is how
	// you pull structured fields (like Field/Msg here) back out of an
	// error chain that's been wrapped one or more times.
	var validationErr *ValidationError
	if errors.As(err, &validationErr) {
		fmt.Printf("recovered structured error: field=%q msg=%q\n", validationErr.Field, validationErr.Msg)
	}

	fmt.Println()
	fmt.Println("--- plain equality (==) breaks the moment wrapping is involved ---")
	fmt.Println("err == ErrNotFound:", lookup(-1) == ErrNotFound) // false! always prefer errors.Is over ==
}
