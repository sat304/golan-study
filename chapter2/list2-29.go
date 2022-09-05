package main

import "fmt"

func main() {
	data := "*新しい値*"
	m1 := modify(data)
	data = "+new data+"
	m2 := modify(data)

	fmt.Println(m1())
	fmt.Println(m2())

}

func modify(d string) func() []string {
	m := []string{"one", "two"}
	return func() []string {
		return append(m, d)
	}
}
