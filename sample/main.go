package main

import (
	"fmt"
	"net/http"
	api"sample/api"
)

func main() {
	http.HandleFunc("/", api.HelloW)
	fmt.Println("listeining on port 8085")
	err := http.ListenAndServeTLS(":8085", Cert, Key, nil)
	if err != nil {
		fmt.Println(err)
	}

}
