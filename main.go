package main

import "fmt"

type custom struct {
	firstname
	lastname
}

type custom1 struct {
	firstname
	lastname
}

func main() {

	fmt.Println("This is a simple text")
	// This is a modification from git hub

	custom.firstname = "Firstname String"
	custom.lastname = "Custom lastname"

	custom1.firstname = "This is the second struct of custom"
	custom.firstnam = " w"

}
