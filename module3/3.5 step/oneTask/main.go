package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {

	res := 0

	bufik := bufio.NewReader(os.Stdin)

	buf, err := bufik.ReadString('\n')
	if err == io.EOF {

	}

	es, _ := strconv.Atoi(buf)
	if es < 100 && es > 0 {
		fmt.Print(es)
		res += es
	}

	str := strconv.Itoa(res)
	w := bufio.NewWriter(os.Stdout)
	w.WriteString(str)
	w.Flush()
}
