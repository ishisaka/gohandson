// List 3-18
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Data is interface for MyData
type Data interface {
	SetValue(vals map[string]string)
	PrintData()
}

// Mydata is structure
type Mydata struct {
	Name string
	Data []int
}

// Setvalue is Mydata method.
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

// PrintData is Mydata method.
func (md *Mydata) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

// Yourdata is structure
type Yourdate struct {
	Name string
	Mail string
	Age  int
}

// SetValue is Yourdata method.
func (yd *Yourdate) SetValue(vals map[string]string) {
	yd.Name = vals["name"]
	yd.Mail = vals["mail"]
	n, _ := strconv.Atoi(vals["age"])
	yd.Age = n
}

// PrintData is Yourdata method.
func (yd *Yourdate) PrintData() {
	fmt.Printf("I'm %s, (%d).\n", yd.Name, yd.Age)
	fmt.Printf("mail: %s.\n", yd.Mail)
}
func main() {
	ob := [2]Data{}
	ob[0] = new(Mydata)
	ob[0].SetValue(map[string]string{
		"name": "Sachiko",
		"data": "55, 66, 77",
	})
	ob[1] = new(Yourdate)
	ob[1].SetValue(map[string]string{
		"name": "Mami",
		"mail": "mami@muse.mo",
		"age":  "34",
	})
	for _, d := range ob {
		d.PrintData()
		fmt.Println()
	}
}
