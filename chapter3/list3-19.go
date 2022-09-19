package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Data is interfacd for Mydata.
type Data interface {
	SetValue(vals map[string]string)
	PrintData()
}

// Mydata is strucure.
type Mydata struct {
	Name string
	Data []int
}

// SetValue is Mydata method.
func (md *Mydata) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	valt := strings.Split(vals["data"], " ")
	vali := []int{}
	for _, i := range valt {
		n, _ := strconv.Atoi(i)
		vali = append(vali, n)
	}
	md.Data = vali
}

// PrintData is mydata method.
func (md *Mydata) PrintData() {
	if md != nil {
		fmt.Println("Name:", md.Name)
		fmt.Println("Data:", md.Data)
	} else {
		fmt.Println("**This is Nil value.**")
	}
}

// Yourdata is structure.
type Yourdata struct {
	Name string
	Mail string
	Age  int
}

// Setvalue is Yourdata method.
func (md *Yourdata) SetValue(vals map[string]string) {
	md.Name = vals["Name"]
	md.Mail = vals["mail"]
	n, _ := strconv.Atoi(vals["age"])
	md.Age = n
}

// PrintData is yourdata method.
func (md *Yourdata) PrintData() {
	fmt.Printf("I'm %s. (%d).\n", md.Name, md.Age)
	fmt.Printf("mail: %s.\n", md.Mail)
}

func main() {
	var ob *Mydata
	ob.PrintData()
	ob = &Mydata{}
	ob.SetValue(map[string]string{
		"name": "jiro",
		"data": "123 456 789",
	})
	ob.PrintData()
	// ob := [2]Data{}
	// ob[0] = new(Mydata)
	// ob[0].SetValue(map[string]string{
	// 	"name": "Sachiko",
	// 	"data": "55, 66, 77",
	// })
	// ob[1] = new(Yourdata)
	// ob[1].SetValue(map[string]string{
	// 	"name": "Mami",
	// 	"mail": "mami@mume.mo",
	// 	"age":  "34",
	// })
	// for _, d := range ob {
	// 	d.PrintData()
	// 	fmt.Println()
	// }
}
