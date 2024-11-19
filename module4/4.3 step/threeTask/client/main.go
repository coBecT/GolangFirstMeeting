package main

import (
	"net/http"
	"strings"
)

const (
	urGet  = "http://localhost:3333/get"
	urPost = "http://localhost:3333/post"
)

func main() {

	getHttp, err := http.Get(urGet)
	if err != nil {
	}

	getPost, err := http.Post(urPost, "application/json", strings.NewReader(`{}`))
}
