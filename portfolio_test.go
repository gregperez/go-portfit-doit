package main

import (
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	// GIVEN
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// WHEN
	os.Args[1] = "2021-12-01"
	os.Args[2] = "2021-12-15"

	// THEN
	main()
}