package auth

import (
	"log"
	"net/http"
)

func Check(w http.ResponseWriter, r *http.Request) {
	log.Println("hit it")
}
