package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if player == "Pepper" {
		fmt.Fprint(w, "20")
		return
	}
	if player == "Floyd" {
		fmt.Fprint(w, "10")
	}
}

func main() {
	handler := http.HandlerFunc(PlayerServer)
	log.Printf("Serving on 0.0.0.0:5000\n\n")
	log.Fatal(http.ListenAndServe(":5000", handler))
}
