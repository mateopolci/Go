// Number Guessing Game
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"strings"
)

func main(){
    // Crea una nueva fuente de aleatoriedad basada en el tiempo actual
    sec := rand.NewSource(time.Now().UnixNano())
    // Crea un nuevo generador de números aleatorios con la fuente
    r := rand.New(sec)
    // Genera un número aleatorio del 1 al 100
    randomNumber := r.Int31n(100) + 1

	fmt.Println("Bienvenido a Go Guess The Number")
	
	var difficulty string
	for{
		fmt.Println("Enter a difficulty: Easy, Medium or Hard")
		_, err := fmt.Scan(&difficulty)
		difficulty := strings.ToLower(difficulty)
		if err != nil {
			log.Println("Error reading difficulty:", err)
			break
		}
	}


	fmt.Println("Im the guessed difficulty", difficulty)
	fmt.Println("Im the random number", randomNumber)
}