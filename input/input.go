package main

import "fmt"

func main() {
	var input string
	fmt.Print("Masukkan sesuatu: ")
	fmt.Scanln(&input)
	fmt.Printf("Anda memasukkan: %s\n", input)
}
