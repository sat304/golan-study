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
	switch {
	case n%2 == 0:
		fmt.Println("even")
	case n%2 == 1:
		fmt.Println("odd")
	}
}
func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
