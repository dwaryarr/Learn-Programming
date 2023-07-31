package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpDwi NoKTP = "1234567890"
	var MarriedStatus Married = false
	fmt.Println(noKtpDwi)
	fmt.Println(MarriedStatus)

}
