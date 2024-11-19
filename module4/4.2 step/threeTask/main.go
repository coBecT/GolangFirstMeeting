package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// // Считайте с консоли две переменные, сначала name, затем age. Сделайте HTTP запрос на
// // http://127.0.0.1:8080/hello передав name и age в query параметрах, затем напечатайте ответ сервера (только тело).
func main() {
	buf := bufio.NewReader(os.Stdin)
	name, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input name", err)
	}
	name = strings.TrimSuffix(name, "\n")
	age, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input age", err)
	}
	age = strings.TrimSpace(age)
	//	base, err := url.Parse("http://127.0.0.1:8080")
	ur := "http://127.0.0.1:8080"
	base, err := url.Parse(ur)
	if err != nil {
		fmt.Println("Error parsing URL", err)
	}
	base.Path += "hello"
	params := url.Values{}
	params.Add("name", name)
	params.Add("age", age)
	keys := []string{"name", "age"}
	rewQuery := ""
	for _, key := range keys {
		rewQuery += fmt.Sprintf("%v=%v&", key, params.Get(key))
	}
	rewQuery = strings.TrimSuffix(rewQuery, "&")
	base.RawQuery = rewQuery
	res, err := http.Get(base.String())
	if err != nil {
		fmt.Println("Error fetching URL", err)
	}
	defer res.Body.Close()
	red, err := io.ReadAll(res.Body)
	fmt.Println(string(red))

}
