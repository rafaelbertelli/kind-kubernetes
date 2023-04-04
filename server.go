package main

import "net/http"

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	println("Chamou!!!")
	w.Write([]byte("<h1>Hello Full Cycle!!!</h1>"))
}
