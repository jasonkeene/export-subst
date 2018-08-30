// This package tests if export substitution occured.
package main

import (
	"fmt"
	"os"
)

const (
	fixture      = "$Format:%h$"
	expectedData = "\x24\x46\x6f\x72\x6d\x61\x74\x3a\x25\x68\x24"
)

func main() {
	if fixture != expectedData {
		fmt.Printf("expected to get: %q, got: %q\n", expectedData, fixture)
		os.Exit(1)
	}
}
