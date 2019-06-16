package main

import (
	"fmt"

	"github.com/shebbar/leftpad"
)



func main() {
	const s = ' '


	fmt.Println(leftpad.Format("hello", 15))
	fmt.Println(leftpad.Format("goodbye", 15))
	fmt.Println(leftpad.Format("como esta", 15))
	fmt.Println(leftpad.Format("Internalization", 15))
	fmt.Println(leftpad.FormatRune("hello", 15, s))
	fmt.Println(leftpad.FormatRune("goodbye", 15, s))
	fmt.Println(leftpad.FormatRune("como esta", 15, s))
}