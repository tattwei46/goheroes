package handler

import (
	"fmt"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}
