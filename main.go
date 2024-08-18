package main

import "fmt"

func main() {
    months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
    quarter2 := months[3:6]
    fmt.Println("Slice original:", quarter2, len(quarter2), cap(quarter2))
    quarter2extended := quarter2[:4]
    fmt.Println("Extendido:", quarter2extended, len(quarter2extended), cap(quarter2extended))
}