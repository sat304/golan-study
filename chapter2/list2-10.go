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
		fmt.Println(n)
	} else {
		return
	}
	t := 0
	c := 1
	for c <= n {
		t += c
		c++
	}
	fmt.Println(t, "です。")
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
