package main

import (
	// "fmt"
	"log"
	"net/http"
	// "api-crud-go/config"
	"api-crud-go/mahasiswa"
	// "api-crud-go/models"
	"api-crud-go/utils"
	"fmt"
	"context"
)



func main() {

	http.HandleFunc("/mahasiswa", GetMahasiswa)
	err := http.ListenAndServe(":7000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

// GetMahasiswa
func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		mahasiswas, err := mahasiswa.GetAll(ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, mahasiswas, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}

