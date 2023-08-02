package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var ptr *int
	fmt.Println(ptr) //print nil
	in := 5
	var mynumber = &in
	fmt.Println(mynumber, *mynumber)
	in *= 5
	fmt.Println(mynumber, *mynumber, in)
}
