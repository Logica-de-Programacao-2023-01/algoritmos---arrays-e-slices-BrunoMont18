package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80}
	var n1, n2 int

	fmt.Print("Escolha um número (0-7):")
	fmt.Scan(&n1)
	fmt.Print("Escolha um número (0-7):")
	fmt.Scan(&n2)

	slice[n1], slice[n2] = slice[n2], slice[n1]
	fmt.Println(slice)

}
