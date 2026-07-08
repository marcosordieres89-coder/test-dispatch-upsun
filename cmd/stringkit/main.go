// Command stringkit is a tiny CLI over the stringkit package.
//
// Usage:
//
//	stringkit reverse "hello"
//	stringkit palindrome "racecar"
//	stringkit words "hello world"
package main

import (
	"fmt"
	"os"

	"github.com/marcosordieres89-coder/test-dispatch-upsun"
)

func main() {
	if len(os.Args) < 3 {
		usage()
		os.Exit(2)
	}

	cmd, arg := os.Args[1], os.Args[2]
	switch cmd {
	case "reverse":
		fmt.Println(stringkit.Reverse(arg))
	case "palindrome":
		fmt.Println(stringkit.IsPalindrome(arg))
	case "words":
		fmt.Println(stringkit.WordCount(arg))
	default:
		usage()
		os.Exit(2)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: stringkit <reverse|palindrome|words> <string>")
}
