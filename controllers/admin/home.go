package backend

import (
	"log"
	"net/http"
)

func Home(rw http.ResponseWriter, req *http.Request) {
	log.Println("main")
}
