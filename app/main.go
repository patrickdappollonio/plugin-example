// main.go
package main

import (
	"fmt"
	"log"
	"plugin"
)

func main() {
	// Load the plugin
	p, err := plugin.Open("./myplugin.so")
	if err != nil {
		log.Fatalf("Error opening plugin: %v\n", err)
	}

	// Look up the Greet symbol
	symGreet, err := p.Lookup("Greet")
	if err != nil {
		log.Fatalf("Error looking up symbol Greet: %v\n", err)
	}

	// Assert that the symbol is a function of the correct signature
	greetFunc, ok := symGreet.(func(string))
	if !ok {
		log.Fatalf("Symbol Greet has unexpected type: %T\n", symGreet)
	}

	// Now we can call the function in the plugin
	greetFunc("Gopher")

	fmt.Println("Main application finished calling the plugin's Greet function.")
}
