// NB - program copied verbatim from Francesc Campoy's
// FOSDEM talk "The State of Go"

package main

// No C code needed.
import "C"

import "fmt"

var V int

func F() { fmt.Printf("Hello, number %d\n", V) }
