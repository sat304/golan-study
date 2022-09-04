package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	x := input("type 1-5")
	n, err := strconv.Atoi(x)
	if err == nil {
		fmt.Println(n)
	} else {
		return
	}
	t := 0
	switch n {
	case 5:
		t += 5
		fallthrough
	case 4:
		t += 4
		fallthrough
	case 3:
		t += 3
		fallthrough
	case 2:
		t += 2
		fallthrough
	case 1:
		t++
	default:
		fmt.Println("invalid input")
		return
	}
	fmt.Println("合計：", t, "です。")
}
func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
