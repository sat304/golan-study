package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	x := input("type data")
	ar := strings.Split(x, " ")
	t := 0

	for _, v := range ar {
		n, err := strconv.Atoi(v)
		if err != nil {
			goto err
		}
		t += n
	}
	fmt.Println("total:", t)
	return
err:
	fmt.Println("error!")
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
