package main

import (
	"bufio"
	"fmt"
	"os"
)

const Constant = "this is constant"

func main() {
	var name string = "Name"
	fmt.Println(name)

	var isbool bool = true
	fmt.Println(isbool)

	var interger int = 7806
	fmt.Println(interger)

	var fl float64 = 2.562045
	fmt.Println(fl)
	var auto = "input type detected"

	fmt.Println(auto)

	integ := 5
	fmt.Println(integ)

	fmt.Println(Constant)
	fmt.Println(" provide input")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println(input)

}
