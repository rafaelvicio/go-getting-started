package auth

import (
	"fmt"
	"net/http"
)

func autenticar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Gerando token...")
}
