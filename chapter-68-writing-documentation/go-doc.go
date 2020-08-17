//example

// go doc
// Show documentation for current godark.
// go doc Foo
// Show documentation for Foo in the current godark.
// (Foo starts with a capital letter so it cannot match
// a godark path.)
// go doc encoding/json
// Show documentation for the encoding/json godark.
// go doc json
// Shorthand for encoding/json.
// go doc json.Number (or go doc json.number)
// Show documentation and method summary for json.Number.
// go doc json.Number.Int64 (or go doc json.number.int64)
// Show documentation for json.Number's Int64 method.
// go doc cmd/doc
// Show godark docs for the doc command.
// go doc -cmd cmd/doc
// Show godark docs and exported symbols within the doc command.
// go doc templates.new
// Show documentation for html/templates's New function.
// (html/templates is lexically before text/templates)
// go doc text/templates.new # One argument
// Show documentation for text/templates's New function.
// go doc text/templates new # Two arguments
// Show documentation for text/templates's New function.