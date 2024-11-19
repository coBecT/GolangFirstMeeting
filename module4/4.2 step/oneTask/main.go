package main

import (
	"fmt"
	"net"
	"strings"
)

//client
//Подключитесь к адресу 127.0.0.1:8081 по протоколу TCP, считайте от сервера 3 сообщения, и выведите их в верхнем регистре. Рекомендуем использовать буфер в 1024 байта.

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("Error connecting:", err)
	}
	defer conn.Close()

	bufik := make([]byte, 1024)
	mapa := make(map[int]string)
	res := ""
	for i := 0; i < 3; i++ {
		n, err := conn.Read(bufik)
		if err != nil {
			fmt.Println("Error reading:", err)
		}
		res += string(bufik[:n])
		mapa[i] = strings.ToUpper(string(bufik[:n]))
	}

	for _, v := range mapa {
		fmt.Println(v)
	}
}
