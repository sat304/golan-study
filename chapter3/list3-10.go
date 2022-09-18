package main

import "fmt"

// Mydata is structure.
type Mydata struct {
	Name string
	Data []int
}

func (md Mydata) PrintData() {
	fmt.Println("*** Mydata ***")
	fmt.Println("Name:", md.Name)
	fmt.Println("Data:", md.Data)
	fmt.Println("*** end ***")
}

func main() {
	taro := new(Mydata)
	fmt.Println(taro)
	taro.Name = "Taro"
	taro.Data = make([]int, 10)
	fmt.Println(taro)
	taro.PrintData()
}
