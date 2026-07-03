package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return fmt.Sprintf("hi, I'm %s", p.Name)
}

// Embedding: Employee has no field named "Person", but Person's fields
// and methods are PROMOTED - accessible directly on Employee, as if
// Employee "inherited" them. This is Go's deliberate alternative to
// class inheritance: composition, with the compiler doing the
// forwarding for you.
type Employee struct {
	Person  // embedded (anonymous field) - note: no field name, just the type
	Company string
}

// Employee tags struct fields with `json:"..."` metadata, read via
// reflection by encoding/json (and many other libraries). Tags are just
// a convention encoded as a string literal - the compiler doesn't
// interpret them itself.
type Config struct {
	Host string `json:"host"`
	Port int    `json:"port,omitempty"`
}

func demoStructsEmbedding() {
	fmt.Println("--- struct literals, embedding, promotion ---")

	e := Employee{
		Person:  Person{Name: "Ada", Age: 30},
		Company: "Analytical Engines Inc",
	}

	// Promoted field and method - reached directly, without "e.Person.Name".
	fmt.Println("promoted field:", e.Name)
	fmt.Println("promoted method:", e.Greet())

	// You can still go through the embedded field explicitly when needed
	// (e.g. to disambiguate if Employee also defined its own Name).
	fmt.Println("explicit path still works:", e.Person.Name)

	// Structs are comparable with == IF all their fields are comparable
	// (no slices, maps, or funcs among them). Comparison is field-by-field.
	p1 := Person{Name: "Ada", Age: 30}
	p2 := Person{Name: "Ada", Age: 30}
	fmt.Println("p1 == p2:", p1 == p2) // true - same field values

	// Anonymous structs: a one-off struct type with no name, handy for
	// small throwaway groupings (like a table-driven test case).
	point := struct{ X, Y int }{X: 1, Y: 2}
	fmt.Println("anonymous struct:", point)

	_ = Config{Host: "localhost", Port: 8080} // tags shown for reference; see encoding/json in the stdlib
}
