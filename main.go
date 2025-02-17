package main

import 	"fmt"
		"log"
		"net/http"

func index(w http.ResponseWriter, r *http.Request) {
	w.header().Set("Content-Type", "text/plain")
	fmt.FPrintf(w, "Aplicação Exemplo - 2.0")
}
	
func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServer("0.0.0.0:5000", nil))
}
