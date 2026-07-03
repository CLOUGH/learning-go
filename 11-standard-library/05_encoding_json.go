package main

import (
	"encoding/json"
	"fmt"
)

// Struct tags (`json:"..."`) tell encoding/json which JSON key each field
// maps to - see 02-core-fundamentals/05_structs_embedding.go for tags in
// general. `omitempty` skips the field entirely when it's the zero value.
type Account struct {
	Username string   `json:"username"`
	Age      int      `json:"age"`
	Email    string   `json:"email,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}

func demoEncodingJSON() {
	fmt.Println("--- json.Marshal: Go value -> JSON bytes ---")

	a := Account{Username: "gopher", Age: 15, Tags: []string{"mascot", "blue"}}

	data, err := json.Marshal(a)
	fmt.Println("Marshal:", string(data), err)

	// Email is the zero value (""), and has `omitempty` - it's absent
	// from the output entirely, not present as `"email":""`.

	pretty, _ := json.MarshalIndent(a, "", "  ")
	fmt.Println("MarshalIndent:")
	fmt.Println(string(pretty))

	fmt.Println()
	fmt.Println("--- json.Unmarshal: JSON bytes -> Go value ---")

	input := `{"username":"ada","age":30,"email":"ada@example.com"}`
	var decoded Account
	if err := json.Unmarshal([]byte(input), &decoded); err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}
	fmt.Printf("decoded: %+v\n", decoded)

	fmt.Println()
	fmt.Println("--- unknown fields are silently ignored by default ---")

	inputWithExtra := `{"username":"grace","age":40,"extra_field":"ignored"}`
	var decoded2 Account
	json.Unmarshal([]byte(inputWithExtra), &decoded2)
	fmt.Printf("decoded (extra_field just vanishes): %+v\n", decoded2)
}

/*
Expected output (from demoEncodingJSON, called via main.go):

--- json.Marshal: Go value -> JSON bytes ---
Marshal: {"username":"gopher","age":15,"tags":["mascot","blue"]} <nil>
MarshalIndent:
{
  "username": "gopher",
  "age": 15,
  "tags": [
    "mascot",
    "blue"
  ]
}

--- json.Unmarshal: JSON bytes -> Go value ---
decoded: {Username:ada Age:30 Email:ada@example.com Tags:[]}

--- unknown fields are silently ignored by default ---
decoded (extra_field just vanishes): {Username:grace Age:40 Email: Tags:[]}
*/
