package main

import (
	"fmt"
)

func main() {
	// fmt.Print("Go runs on ")
	// os := runtime.GOOS
	// switch os {
	// case "darwin":
	// 	fmt.Println("OS X.")
	// case "linux":
	// 	fmt.Println("Linux.")
	// default:
	// 	// freebsd, openbsd,
	// 	// plan9, windows...
	// 	fmt.Printf("%s.\n", os)
	// }

	fmt.Print("Enter a number: ")
	var key int
	fmt.Scanf("%d", &key)

	switch {
	case key > 0:
		fmt.Println("Positive")
		fallthrough // Tiếp tục thực thi case tiếp theo
	case key < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}


}