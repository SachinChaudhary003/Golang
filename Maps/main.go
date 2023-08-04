package main

import "fmt"

func main() {

	maps := make(map[int]string)

	maps[1] = "name"
	maps[2] = "surname"
	maps[3] = "age"
	fmt.Println(maps)
	fmt.Println(maps[1])

	delete(maps, 3)
	fmt.Println(maps)
	fmt.Println("loops")
	for key, value := range maps {
		fmt.Print(key)
		fmt.Println(value)
	}

}
