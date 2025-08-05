// Scope determines variable accessibility from universe to block level.
// Shadowing occurs when inner scope variables hide outer ones with same names.
// Go has package, function, and block scopes.
// Understanding prevents bugs from accidentally creating new variables.

package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}