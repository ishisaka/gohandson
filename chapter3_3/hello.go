// List 3-14, 3-15
package main

import (
	"fmt"
)

// Data is interface
type Data interface {
	Initial(name string, data []int)
	PrintData()
}

// Mydata is Struct
type Mydata struct {
	Name string
	Data []int
}

// Initial is init method
func (md *Mydata) Initial(name string, data []int) {
	md.Name = name
	md.Data = data
}

// PrintData is println all data.
func (md *Mydata) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

// List 3-16 Check is method
func (md *Mydata) Check() {
	fmt.Printf("Check! [%s]\n", md.Name)
}

func main() {
	var oh Mydata = Mydata{}
	oh.Initial("Sachiko", []int{55, 66, 77})
	oh.PrintData()
	oh.Check()

	// List 3-15 MyDataをData型として扱う
	var ob Data = new(Mydata)
	ob.Initial("Seiko", []int{155, 166, 177})
	ob.PrintData()
	// ob.Check() これはエラー
}
