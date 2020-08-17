package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	// First we create a FuncMap with which to register the function.
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the templates text.
		"tl": strings.ToLower,
		"tu": strings.ToUpper,
	}

	// A simple templates definition to test our function.
	// We print the input text several ways:
	// 0 lowered-cased
	// 1 title-cased
	// 2 upper-cased and then printed with %q
	// - printed with %q or mentioned format string and then uppered-cased.

	//when you used printf, you must be a separate string with "|" on between they are
	//you must used double quotation mark "" to format string on templates

	const templateText = `
	Input: {{printf "%q" .}}
	Output 0 to lower: {{tl .}}
	Output 1 to upper: {{tu . | printf "%q"}}
	Output 2 with printf to upper: {{printf "%q" . | tu}}
	`

	// Create a templates, add the function map, and parse the text.
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// Run the templates to verify the output.
	err = tmpl.Execute(os.Stdout, "The Go Programming Language")
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

}
