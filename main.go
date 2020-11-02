// Single line comment
/* Multi-
line comment */

/* A build tag is a line comment starting with // +build
and can be execute by go build -tags="foo bar" command.
Build tags are placed before the package clause near or at the top of the file
followed by a blank line or other line comments. */
// +build prod, dev, test

// A package clause starts every source file.
// Main is a special name declaring an executable rather than a library.
package main

import "./foo"

func main() {

	foo.BeyondHello()

}
