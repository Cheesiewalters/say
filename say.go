// Package say provides primitives for saying hello
package say

import "fmt"

// Hello returns a string saying hello to the name being passed in as an argument
func Hello(name string) string {
	return fmt.Sprintf("Hello %s.", name)
}
