package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "data set")
}
func api(w http.ResponseWriter, r *http.Request) {
	//data := r.FormValue("data")
	msg := fmt.Sprintf("Ok %v")
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Query ok!")
	})
	http.HandleFunc("/index", index)
	http.HandleFunc("/ok", api)

	log.Println("Strating at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
