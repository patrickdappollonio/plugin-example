package main

import "fmt"

// Greet is the symbol we will export from our plugin.
// Note that the package name must be "main" for a plugin.
func Greet(name string) {
	fmt.Printf("Hello, %s! Greetings from the plugin.\n", name)
}
