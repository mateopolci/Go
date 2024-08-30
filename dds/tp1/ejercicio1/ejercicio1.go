package main

import (
	"fmt"
)

func main(){


}

func math(){
	var option int
	fmt.Println("Welcome to the math module, enter a number:")
	
	fmt.Println("1: Adding")
	fmt.Println("2: Subtracting")
	fmt.Println("3: Multiplying")
	fmt.Println("4: Dividing")
	
	switch option {
	case 1:
		fmt.Println("Add function")
	case 2:
		fmt.Println("Subtract function")
	case 3:
		fmt.Println("Multiply function")
	case 4:
		fmt.Println("Divide function")
	}
}

func cipher(){
	fmt.Println("Welcome to the Caesar Cipher module, enter a number:")
	fmt.Println("1: Cipher")
	fmt.Println("2: Decipher")

	var key int
	var option int

	switch option {
	case 1:
		fmt.Println("Cipher function selected, enter a key value to cipher")
		fmt.Scan(key)
	case 2:
		fmt.Println("Decipher function, enter a key value to decipher")
		fmt.Scan(key)
	}
}