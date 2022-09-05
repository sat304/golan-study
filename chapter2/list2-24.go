package main

import "fmt"

func main() {
	m := []string{}
	m, _ = push(m, "apple")
	m, _ = push(m, "orange")
	m, _ = push(m, "banana")

	fmt.Println(m)
	m, v := pop(m)
	fmt.Println(m, v)
}

func push(a []string, v string) ([]string, int) {
	a = append(a, v)
	return a, len(a)
}

func pop(a []string) ([]string, string) {
	return a[:len(a)-1], a[len(a)-1]
}
