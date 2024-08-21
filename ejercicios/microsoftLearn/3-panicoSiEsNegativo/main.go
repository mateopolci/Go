/*
Escriba un programa que pida a un usuario un número. 
Este programa solicita un número y lo imprime. Modifique el código de ejemplo para que:

Se pida continuamente un número entero. La condición de salida del bucle deba 
ser una entrada de usuario que sea un número negativo.

Se bloquee el programa cuando el usuario escriba un número negativo. Después, se imprima el 
error de seguimiento de la pila.

Cuando el número sea 0, se imprima 0 is neither negative nor positive. Siga pidiendo un número.

Cuando el número sea positivo, se imprima You entered: X (donde X sea el número escrito) 
y siga pidiendo un número.

Use el siguiente fragmento de código como punto de partida:
*/
package main

import "fmt"

func main() {
    val := 0
    fmt.Print("Enter number: ")
    fmt.Scanf("%d", &val)
    fmt.Println("You entered:", val)
}