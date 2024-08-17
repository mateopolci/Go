// Ecriba un programa que imprima los números del 1 al 100 e imprimir:
// "Fizz" si el número es divisible por 3.
// "Buzz" si el número es divisible por 5.
// "FizzBuzz" si el número es divisible por 3 y por 5.
// El número si ninguno de los casos anteriores coincide.
// Pruebe usar `switch`.

package main

import "fmt"

func main(){
    
    for i := 1; i <= 100; i++ {
        switch {
        case i % 3 == 0 && i % 5 == 0:
            fmt.Println("FizzBuzz")
        case i % 3 == 0:
            fmt.Println("Fizz")
        case i % 5 == 0:
            fmt.Println("Buzz")
        default:
            fmt.Println(i)
        }
    }
}