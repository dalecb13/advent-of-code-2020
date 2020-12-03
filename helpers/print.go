package helpers

import "fmt"

// PrintStringByteSpace prints the elements of the underlying byte array of a string and puts a space between each byte
func PrintStringByteSpace(s string) {
	fmt.Printf("%x \n", s)
}

// PrintUnambiguous prints the elements of the underlying byte array but leaves "unfamiliar" characters
func PrintUnambiguous(s string) {
	fmt.Printf("%q\n", s)
}

// PrintUnambiguousUnicode prints the underlying byte array as bytes and leaves "unfamiliar" characters as Unicode code points
func PrintUnambiguousUnicode(s string) {
	fmt.Printf("%+q\n", s)
}
