package main

import (
	"fmt"
	"log"
	"net/http"
	// "time"
)

const port string = "9090"

func welcome(w http.ResponseWriter, r *http.Request) {
	dataHTML := `<h1>Kumpulan Bahasa Pemrograman Keren</h1>
		<ul>
			<li>Golang</li>
			<li>PHP</li>
			<li>JavaScript</li>
			<li>Rust</li>
			<li>Java</li>
		</ul>
		Tutorial by : <a href="https://kodingin.com">Kodingin</a>
		`
	if r.Method == "GET" {
		fmt.Fprintf(w, dataHTML)
	}
}

func main() {
	fmt.Println("Berjalan di Port :", port)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Sedang Akses URL %v dengan method %v", r.URL, r.Method)
	// })

	// http.Handle("/handle", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Sedang Akses URL %v dengan method %v", r.URL, r.Method)
	// }))

	http.HandleFunc("/welcome", welcome)

	// server := &http.Server{
	// 	Addr:         ":" + port,
	// 	ReadTimeout:  30 * time.Second,
	// 	WriteTimeout: 30 * time.Second,
	// 	IdleTimeout:  10 * time.Second,
	// }
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}