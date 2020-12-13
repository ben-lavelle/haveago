package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value. Takes a pointer to testing.T as
// an argument, allowing package `testing` to report and log.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`) // Must contain the expected name
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error. All test functions' names begin Test
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
