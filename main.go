package main

import (
	// "log"
	// "net/http"
	// "golang/api-crud-go/config"


	"fmt"
)

type Profil struct {
	ID      int
	Name    string
	Age     int
	Address string
}

func main() {
	var profil Profil

	profil.ID = 1
	profil.Name = "Didik Prabowo"
	profil.Age = 20
	profil.Address = "Wuryantoto, Wonogiri"

	fmt.Println(profil)
	fmt.Println("===========================")
	fmt.Printf("Nama \t: %v\n", profil.Name)
	fmt.Printf("Umur \t: %d\n", profil.Age)
	fmt.Printf("Alamat \t: %v\n", profil.Address)
	fmt.Println("===========================")
}