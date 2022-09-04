package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	x := input("type a number")
	n, err := strconv.Atoi(x)
	if err == nil {
		fmt.Println("1から", n, "の偶数の合計は")
	} else {
		return
	}
	t := 0
	c := 0
	for {
		c++
		if c%2 == 1 {
			continue
		}
		if c > n {
			break
		}
		t += c

	}
	fmt.Println(t, "です。")
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
