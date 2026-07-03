package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

// greetHandler is an ordinary http.Handler: given a request, write a
// response. This is the entire shape of a Go HTTP server - no framework
// required for something this simple.
func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "hello, %s!", name)
}

func demoNetHTTP() {
	fmt.Println("--- net/http: a real client talking to a real server ---")

	// httptest.NewServer starts an actual local HTTP server on a real
	// (random, local) port - this is the standard way to test/demo HTTP
	// code without depending on the network or an external service.
	server := httptest.NewServer(http.HandlerFunc(greetHandler))
	defer server.Close()

	resp, err := http.Get(server.URL + "?name=gopher")
	if err != nil {
		fmt.Println("GET failed:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("reading body failed:", err)
		return
	}

	fmt.Println("status:", resp.Status)
	fmt.Println("body:", string(body))

	// A real, deployed server almost always uses http.ListenAndServe
	// instead of httptest.NewServer:
	//
	//   http.HandleFunc("/greet", greetHandler)
	//   log.Fatal(http.ListenAndServe(":8080", nil))
	//
	// ListenAndServe blocks forever serving requests - not something this
	// demo can call without hanging the rest of the program.
}

/*
Expected output (from demoNetHTTP, called via main.go):

--- net/http: a real client talking to a real server ---
status: 200 OK
body: hello, gopher!
*/
