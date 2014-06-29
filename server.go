/*
 * Battleships
 * 
 * Author: mborowiec (gamza)
 *
 */
 
package main

import (
	"net/http"
	"html/template"
)

func start(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("battleship.html")).Execute(w, nil)
}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.HandleFunc("/", start)
	http.ListenAndServe(":3000", nil)
}
