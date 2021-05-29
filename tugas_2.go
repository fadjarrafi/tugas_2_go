package main

import "fmt"

func main() {
	var angka = 10

	if angka%2 == 0 {
		fmt.Print(angka, " merupakan angka genap")
	} else {
		fmt.Print(angka, " merupakan angka ganjil")
	}
}
