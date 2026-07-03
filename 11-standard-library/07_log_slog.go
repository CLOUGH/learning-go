package main

import (
	"log"
	"log/slog"
	"os"
)

func demoLogSlog() {
	log.Println("--- log: the classic logger (this line is prefixed with a timestamp on stderr) ---")

	// log.Println/Printf write to stderr by default, each call prefixed
	// with a timestamp - unlike fmt.Println, which writes to stdout with
	// no prefix at all. That's why these lines are timestamped in your
	// terminal but every other demo in this lesson isn't.
	log.Printf("processed %d items", 3)

	// log.Fatal(...) would log the message, then call os.Exit(1)
	// immediately - no deferred functions run anywhere in the program.
	// Not called here since it would kill the rest of this demo.

	log.Println("--- log/slog: structured logging (Go 1.21+) ---")

	// slog.Info/Warn/Error take a message plus alternating key-value
	// pairs, rather than a single formatted string - the point is that a
	// log aggregator can parse "user_id" and "attempt" as actual fields,
	// not just grep the message text.
	slog.Info("user login", "user_id", 42, "attempt", 1)
	slog.Warn("retrying request", "url", "https://example.com", "attempt", 2)

	// A JSON handler writes each log line as a JSON object instead of
	// slog's default human-readable text - the format you'd actually want
	// in production, where logs get shipped to something that parses JSON.
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	jsonLogger.Info("structured as JSON", "request_id", "abc-123", "status", 200)
}

/*
Expected output (from demoLogSlog, called via main.go - timestamps will
be whatever moment you actually run this, not what's shown below):

2026/07/03 03:01:09 --- log: the classic logger (this line is prefixed with a timestamp on stderr) ---
2026/07/03 03:01:09 processed 3 items
2026/07/03 03:01:09 --- log/slog: structured logging (Go 1.21+) ---
2026/07/03 03:01:09 INFO user login user_id=42 attempt=1
2026/07/03 03:01:09 WARN retrying request url=https://example.com attempt=2
{"time":"2026-07-03T03:01:09.071784736-05:00","level":"INFO","msg":"structured as JSON","request_id":"abc-123","status":200}

Also note: `log.Println`/`log.Printf` AND the default `slog.Info`/
`slog.Warn` (via slog.Default()) all go to STDERR - only the explicit
`jsonLogger` constructed above with `slog.NewJSONHandler(os.Stdout, nil)`
writes to STDOUT. If you redirect the two streams separately
(`go run ./11-standard-library 2>err.txt`), everything in this file
except the final JSON line ends up in err.txt, not your terminal.
*/
