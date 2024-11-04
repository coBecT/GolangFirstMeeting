package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type GlobalId struct {
	Id int `json:"global_id"`
}

// тогда требуется создавать срез структур:

func main() {
	file, err := os.Open("C:\\go-work\\src\\goPracticaStepik\\module3\\3.6 step\\twoTask\\data-20190514T0100.json")
	if err != nil {
		fmt.Print("Error -> ", err.Error())
	}
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Print("Error -> ", err.Error())
	}

	var globalIds []GlobalId
	err = json.Unmarshal(data, &globalIds)
	if err != nil {
		fmt.Print("Error -> ", err.Error())
	}
	res := 0
	for _, val := range globalIds {
		res += val.Id
	}
	fmt.Println(res)
	//	fmt.Print(globalIds)
}
