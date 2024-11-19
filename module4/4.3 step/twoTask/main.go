package main

import (
	"fmt"
	"net/http"
)

func hand(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello,%s!", r.URL.Query().Get("name"))))
}
func main() {

	http.HandleFunc("/api/user", hand)

	http.ListenAndServe(":9000", nil)
}
