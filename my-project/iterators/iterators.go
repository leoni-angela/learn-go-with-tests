package iteration

import "strings"


const repeatCount = 5

// using Builder and WriteString for better performance (since strings are immutable in Go)
// takes alot of copying memory if we use += operator in a loop
func Repeat(char string) string {
	var repeated strings.Builder
	for i := 0; i< repeatCount; i++ {
		repeated.WriteString(char)
	}
	return repeated.String()
	// retrieve result with String() method
}

func ExampleRepeat(char string, n int) string {
	var repeated strings.Builder
	for i :=0; i<n; i++ {
		repeated.WriteString(char)
	}
	return repeated.String()
}