package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Group struct {
	ID       int       `json:"ID"`
	Number   string    `json:"Number"`
	Year     int       `json:"Year"`
	Students []Student `json:"Students"`
}

type Student struct {
	LastName   string `json:"LastName"`
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	Birthday   string `json:"Birthday"`
	Address    string `json:"Address"`
	Phone      string `json:"Phone"`
	Rating     []int  `json:"Rating"`
}

type out struct {
	Average float64 `json:"Average"`
}

func main() {

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	var group Group
	err = json.Unmarshal(data, &group)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	totalRating := 0
	for _, student := range group.Students {
		totalRating += len(student.Rating)
	}

	var resAv float64

	if len(group.Students) > 0 {
		resAv = float64(totalRating) / float64(len(group.Students))
	} else {
		resAv = 0.0
	}
	outp := out{Average: resAv}
	outData, err := json.MarshalIndent(outp, "", "    ")
	w := bufio.NewWriter(os.Stdout)
	w.WriteString(string(outData))
	w.Flush()

}

//file, err := os.Open("C:\\go-work\\src\\goPracticaStepik\\module3\\3.6 step\\oneTask\\file.json")
//if err != nil {
//	panic(err)
//}
//defer file.Close()
//data, err := io.ReadAll(file)

//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"io"
//	"os"
//)
//
//type Student struct {
//	LastName   string `json:"LastName"`
//	FirstName  string `json:"FirstName"`
//	MiddleName string `json:"MiddleName"`
//	Birthday   string `json:"Birthday"`
//	Address    string `json:"Address"`
//	Phone      string `json:"Phone"`
//	Rating     []int  `json:"Rating"`
//}
//
//type Group struct {
//	ID       int       `json:"ID"`
//	Number   string    `json:"Number"`
//	Year     int       `json:"Year"`
//	Students []Student `json:"Students"`
//}
//
//type Output struct {
//	Average float64 `json:"Average"`
//}
//
//func main() {
//	// Считывание данных из стандартного ввода
//	//	data, err := ioutil.ReadAll(os.Stdin)
//	file, err := os.Open("C:\\go-work\\src\\goPracticaStepik\\module3\\3.6 step\\oneTask\\file.json")
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//
//	data, err := io.ReadAll(file)
//	if err != nil {
//		fmt.Println("Error reading input:", err)
//		return
//	}
//
//	// Декодирование JSON в структуру Group
//	var group Group
//	err = json.Unmarshal(data, &group)
//	if err != nil {
//		fmt.Println("Error unmarshalling JSON:", err)
//		return
//	}
//
//	// Вычисление общего количества оценок и количества студентов
//	totalRatings := 0
//	for _, student := range group.Students {
//		totalRatings += len(student.Rating)
//	}
//
//	// Вычисление среднего количества оценок
//	var average float64
//	if len(group.Students) > 0 {
//		average = float64(totalRatings) / float64(len(group.Students))
//	} else {
//		average = 0.0
//	}
//
//	// Подготовка вывода
//	output := Output{Average: average}
//
//	// Кодирование результата в JSON
//	outputData, err := json.MarshalIndent(output, "", "    ")
//	if err != nil {
//		fmt.Println("Error marshalling output to JSON:", err)
//		return
//	}
//
//	// Вывод результата
//	fmt.Println(string(outputData))
//}
