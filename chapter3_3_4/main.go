// List3-21, 3-22
package main

import (
	"fmt"
	"reflect"
)

// General is all data type
type General interface{}

//GData is holding General value
type GData interface {
	Set(m string, g General) GData
	Print()
}

// NData is structure
type NData struct {
	Name string
	Data int
}

// Set is NData method
func (nd *NData) Set(m string, g General) GData {
	nd.Name = m
	// nd.Data = g.(int)
	if reflect.TypeOf(g).Kind() == reflect.Int {
		nd.Data = g.(int)
	}
	return nd
}

// Print is NData method
func (nd *NData) Print() {
	fmt.Printf("<<%s>> value: %d\n", nd.Name, nd.Data)
}

//SData is Structure
type SData struct {
	Name string
	Data string
}

// Set is SData method
func (sd *SData) Set(m string, g General) GData {
	sd.Name = m
	// sd.Data = g.(string)
	if reflect.TypeOf(g).Kind() == reflect.String {
		sd.Data = g.(string)
	}
	return sd
}

// Print is SData method
func (sd *SData) Print() {
	fmt.Printf("<<%s>> value; %s\n", sd.Name, sd.Data)
}

func main() {
	var data = []GData{}
	data = append(data, new(NData).Set("Taro", 123))
	data = append(data, new(SData).Set("Jiro", "Hello!"))
	data = append(data, new(NData).Set("Hanako", 98700))
	data = append(data, new(SData).Set("Sachiko", []string{"happy?"}))
	for _, ob := range data {
		ob.Print()
	}
}
