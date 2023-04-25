package api

import (
	"fmt"
	"net/http"
)

func HelloW(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO SECURE CONNECTION")
}
