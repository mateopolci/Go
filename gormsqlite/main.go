package main

import (
	"fmt"
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Contacts struct {
	Id        int
	FirstName string
	LastName  string
	Phone     string
	Adress    string
}

func main() {
	//Abrir la conexión a la db
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("DB CONNECTION FAILED", err)
	}

	//Creamos una variable para almacenar el contacto
	var contact1 Contacts
	var contact2 Contacts
	var contact3 Contacts

	//Intentamos recuperar el primer contacto con id=1
	result1 := db.First(&contact1, 1)
	result2 := db.First(&contact2, 2)
	result3 := db.First(&contact3, 3)
	
	if result1.Error != nil {
		log.Fatal("No se encontró el contacto con id=1", result1.Error)
		return
	}
	if result2.Error != nil {
		log.Fatal("No se encontró el contacto con id=1", result2.Error)
		return
	}
	if result3.Error != nil {
		log.Fatal("No se encontró el contacto con id=1", result3.Error)
		return
	}

	//Imprimimos resultado
	fmt.Println("---------------------------------")
	fmt.Println("ID:", contact1.Id)
	fmt.Println("First Name:", contact1.FirstName)
	fmt.Println("Last Name:", contact1.LastName)
	fmt.Println("Phone:", contact1.Phone)
	fmt.Println("Address:", contact1.Adress)
	fmt.Println("---------------------------------")
	fmt.Println("ID:", contact2.Id)
	fmt.Println("First Name:", contact2.FirstName)
	fmt.Println("Last Name:", contact2.LastName)
	fmt.Println("Phone:", contact2.Phone)
	fmt.Println("Address:", contact2.Adress)
	fmt.Println("---------------------------------")
	fmt.Println("ID:", contact3.Id)
	fmt.Println("First Name:", contact3.FirstName)
	fmt.Println("Last Name:", contact3.LastName)
	fmt.Println("Phone:", contact3.Phone)
	fmt.Println("Address:", contact3.Adress)
	fmt.Println("---------------------------------")
}
