package main

import (
	"fmt"
	"os"
)

// func main() {
// 	var s, sep string
// 	for i := 1; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Print("Value: ")
		fmt.Println(os.Args[i])
		fmt.Print("Index: ")
		fmt.Println(i)
	}
}
