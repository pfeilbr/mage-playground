//+build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

// Runs go mod download and then installs the binary.
func Build() error {
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	return sh.Run("go", "install", "./...")
}

// says hello
func Hello() {
	fmt.Printf("hello there you\n")
}
