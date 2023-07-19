// Package packagetest intend to test the installation of riverutils
package packagetest

import "fmt"

// SayHellTo says hello to the parameter
//
// Sample usage SayHelloTo("yourname")
func SayHelloTo(name string) {
	fmt.Printf("Hello %s\n", name)
}
