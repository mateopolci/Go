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

	fmt.Println("Welcome to Go Guess The Number")
	fmt.Println("Enter a difficulty: Easy, Medium or Hard")
	fmt.Println("Im the random number", randomNumber)
}

func difficultyPicker(){
	var difficulty string
	_, err := fmt.Scan(&difficulty)
	difficulty = strings.ToLower(difficulty)

	switch difficulty {
	case "easy":
		difficulty = 
	case "medium":
	case "hard":

	}

	if err != nil {
		log.Println("Oops! An error ocurred:", err)
	} else if difficulty == "easy" {
		fmt.Println("You've selected the Easy difficulty")
	} else if difficulty == "medium" {
		fmt.Println("You've selected the Medium difficulty")
	} else if difficulty == "hard" {
		fmt.Println("You've selected the Hard difficulty")
	} else {
		fmt.Println("Please select a valid difficulty")
	}
}