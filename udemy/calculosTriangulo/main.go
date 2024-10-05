package main

import "fmt"

func main() {
	var base, altura, hipotenusa float32

	fmt.Print("Ingrese la base del triángulo rectángulo: ")
	fmt.Scanf("%f\n",&base)
	fmt.Print("Ingrese la altura del triángulo rectángulo: ")
	fmt.Scanf("%f\n",&altura)
	fmt.Print("Ingrese la hipotenusa del triángulo rectángulo: ")
	fmt.Scanf("%f\n",&hipotenusa)
	
	area := (base*altura)/2
	perimetro := base+altura+hipotenusa
	
	fmt.Println("----------------------------------------------------------")
	fmt.Printf("El área del triángulo rectángulo es: %f \n", area)
	fmt.Printf("El perímetro del triángulo rectángulo es: %f \n",perimetro)
	fmt.Println("----------------------------------------------------------")
}
