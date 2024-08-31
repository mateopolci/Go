package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	)

func main(){
	var option int
	for {
		fmt.Println("----------------------------")
		fmt.Println("1: Calculator")
		fmt.Println("2: Caesar Cipher")
		fmt.Println("----------------------------")
		fmt.Print("Select an option: ")
		fmt.Scanf("%d\n", &option)
		switch option {
		case 1:
			math()
		case 2:
			cipher()
		}
	}
}

func math() {
	var (option, num1, num2, res int)
	fmt.Println("----------------------------")
	fmt.Println("1: Adding")
	fmt.Println("2: Subtracting")
	fmt.Println("3: Multiplying")
	fmt.Println("4: Dividing")
	fmt.Println("----------------------------")
	fmt.Print("Select an option: ")
	fmt.Scanf("%d\n", &option)
	fmt.Print("Enter the first number: ")
	fmt.Scanf("%d\n", &num1)
	fmt.Print("Enter the second number: ")
	fmt.Scanf("%d\n", &num2)
	
	switch option {
	case 1:
		fmt.Print("The sum is: ")
		res=num1+num2
	case 2:
		fmt.Print("The difference is: ")
		res=num1-num2
	case 3:
		fmt.Print("The product is: ")
		res=num1*num2
	case 4:
		fmt.Print("The quotient is: ")
		res=num1/num2
	}
	fmt.Println(res)
}

func cipher(){
	var option int
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Welcome to the Caesar Cipher module:")
	fmt.Println("----------------------------")
	fmt.Println("1: Cipher")
	fmt.Println("2: Decipher")
	fmt.Println("----------------------------")
	fmt.Print("Select an option: ")
	fmt.Scanf("%d", &option)
	
	var (
		key, phrase int
		alphabet[]rune
		shiftedAlphabet[]rune
	)

	//Función para crear el alfabeto original
	for i := 'a'; i <= 'z' ; i++ {
		alphabet = append(alphabet, i)
	}
	
	switch option {
		case 1:
			fmt.Printf("Enter a key value [0-26] to cipher: ")
			fmt.Scan(&key)

				//Crea el alfabeto desplazado en función de la llave
				for i := 'a'; i <= 'z' ; i++ {
					shiftedLetter := alphabet[(int(i)-'a'+key)%26]
					shiftedAlphabet = append(shiftedAlphabet, shiftedLetter)
				}

			//Mostrar el alfabeto original
			fmt.Println("Original Alphabet: ", string(alphabet))
			//Mostrar el alfabeto Desplazado
			fmt.Println("Shifted Alphabet: ", string(shiftedAlphabet))

			fmt.Printf("Enter the phrase to cipher: ")
			scanner.Scan()
			phrase := scanner.Text()
			fmt.Println(phrase)
		case 2:
			fmt.Println("Decipher function selected")
			fmt.Print("Enter a key value [0-26] to decipher: ")
			fmt.Scanf("%d", &key)
			fmt.Print("Enter the phrase to cipher: ")
			fmt.Scanf("%s", &phrase)
	}
}