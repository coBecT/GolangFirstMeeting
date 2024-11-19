package main

import "net/http"

func hand(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello. web!"))
}

func main() {

	//	PORT := ":8080"
	http.HandleFunc("/get", hand)
	http.ListenAndServe(":8080", nil)

}

//func helloHandler(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Hello, web!"))
//}
//
//func main() {
//	http.HandleFunc("/get", helloHandler)
//
//	http.ListenAndServe(":8080", nil)
//}
