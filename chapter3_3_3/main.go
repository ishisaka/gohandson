// List 3-20
package main

import (
	"fmt"
)

// General is all type Data
type General interface{}

// GData is holding General Value
type GData interface {
	Set(m string, g General)
	Print()
}

//GDataImpl is Structure
type GDataImpl struct {
	Name string
	Data General
}

// Set is GDataImpl method
func (gd *GDataImpl) Set(m string, g General) {
	gd.Name = m
	gd.Data = g
}

// Print is GDataImpl method
func (gd *GDataImpl) Print() {
	fmt.Printf("<<%s>>", gd.Name)
	fmt.Println(gd.Data)
}

func main() {
	var data = []GDataImpl{}
	data = append(data, GDataImpl{"Taro", 123})
	data = append(data, GDataImpl{"Hanako", "hello!"})
	data = append(data, GDataImpl{"sachiko", []int{123, 456, 789}})
	for _, ob := range data {
		ob.Print()
	}
}
