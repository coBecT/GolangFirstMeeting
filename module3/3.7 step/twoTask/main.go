package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const unixTime = "2006-01-02 15:04:05"

func main() {
	buf, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
	}
	buf = strings.TrimSuffix(buf, "\n")
	//	layout := time.RFC3339
	in, err := time.Parse(unixTime, buf)

	if in.Hour() > 13 {
		post := in.Add(time.Hour * 24)
		//w := bufio.NewWriter(os.Stdout)
		//w.WriteString(strconv.FormatInt(post.Unix(), 10) + "\n")
		//w.Flush()
		fmt.Print(post.Format(unixTime))

	} else {
		fmt.Print(buf)

	}

}

// 2020-05-15 14:00:00
//
//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//	"time"
//)
//
//const unixTime = "2006-01-02 15:04:05"
//
//func main() {
//	// Читаем ввод
//	buf, err := bufio.NewReader(os.Stdin).ReadString('\n')
//	if err != nil {
//	}
//	buf = strings.TrimSuffix(buf, "\n")
//	//var input string
//	//fmt.Scanln(&input)
//
//	// Парсим строку в объект времени
//	//	layout := "2006-01-02 15:04:05" // формат времени в Go (определённый специфический формат)
//	eventTime, err := time.Parse(unixTime, buf)
//	if err != nil {
//		fmt.Println("Ошибка при парсинге времени:", err)
//		return
//	}
//
//	// Проверяем, до обеда ли событие
//	if eventTime.Hour() < 13 {
//		// Если событие до 13:00, выводим исходное время
//		fmt.Println(eventTime.Format(unixTime))
//	} else {
//		// Если событие после 13:00, переносим на следующий день
//		nextDay := eventTime.Add(24 * time.Hour)
//		fmt.Println(nextDay.Format(unixTime))
//	}
//}
