package main

import "fmt"

func main() {
	var i int = 21
	fmt.Printf("%v \n", i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%% \n")
	var j bool = true
	fmt.Printf("%t \n", j)
	fmt.Printf("%t \n", j)
	fmt.Printf("%v \n", "Я")
	fmt.Printf("%d \n", 21)
	fmt.Printf("%o \n", 21)
	fmt.Printf("%x \n", 15)
	fmt.Printf("%X \n", 15)
	fmt.Printf("%U \n", 'Я')
	var k float64 = 123.456
	fmt.Printf("%.6f \n", k)
	fmt.Printf("%E \n", k)
}
