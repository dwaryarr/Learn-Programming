package main

import "fmt"

func main() {
	var name string
	name = "Dwi Arya"
	fmt.Println(name)

	// string
	name = "Dwi Arya R"
	fmt.Println(name)

	var friendName = "Siapa?"
	fmt.Println(friendName)

	// int (default = int64)
	var age = 20
	fmt.Println(age)

	// penggunaan := utk pengganti var
	country := "Indonesia"
	fmt.Println(country)

	// pembuatan multi variable
	var (
		firstName = "Dwi Arya"
		lastName  = "Ramadhani"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
