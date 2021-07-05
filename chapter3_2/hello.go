package main

import (
	"fmt"
)

// List 3-6
// var mydata struct {
// 	Name string
// 	Data []int
// }
//
// func main() {
// 	mydata.Name = "Taro"
// 	mydata.Data = []int{10, 20, 30}
// 	fmt.Println(mydata)
// }

// List 3-7
// Mydata is structure
// type Mydata struct {
// 	Name string
// 	Data []int
// }
//
// func main() {
// 	taro := Mydata{"Taro", []int{10, 20, 30}}
// 	hanako := Mydata{
// 		Name: "Hanako",
// 		Data: []int{90, 80, 70},
// 	}
// 	fmt.Println(taro)
// 	fmt.Println(hanako)
// }

// List 3-9
// Mydata is structure
// type Mydata struct {
// 	Name string
// 	Data []int
// }
//
// func main() {
// 	taro := Mydata{"taro", []int{10, 20, 30}}
// 	fmt.Println(taro)
// 	rev(&taro)
// 	fmt.Println(taro)
// }
//
// func rev(md *Mydata) {
// 	od := (*md).Data
// 	nd := []int{}
// 	for i := len(od) - 1; i >= 0; i-- {
// 		nd = append(nd, od[i])
// 	}
// 	md.Data = nd
// }

// List 3-10
// Mydata is structure
// type Mydata struct {
// 	Name string
// 	Data []int
// }
// //
// func main() {
// 	taro := new(Mydata)
// 	fmt.Println(taro)
// 	taro.Name = "taro"
// 	taro.Data = make([]int, 5, 5)
// 	fmt.Println(taro)
// }

// List 3-11
// Mydata is structure
// type Mydata struct {
// 	Name string
// 	Data []int
// }
//
// // PrintData is println all data.
// func (md Mydata) PrintData() {
// 	fmt.Println("**** Mydata ****")
// 	fmt.Println("Name: ", md.Name)
// 	fmt.Println("Data: ", md.Data)
// 	fmt.Println("*** end ***")
// }
//
// func main() {
// 	taro := Mydata{"hanako", []int{98, 76, 54, 32, 10}}
// 	taro.PrintData()
// }

// List 3-12, 3-13
type intp int

func (num intp) IsPrime() bool {
	n := int(num)
	for i := 2; i <= (n / 2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func (num intp) PrimeFactor() []int {
	ar := []int{}
	x := int(num)
	n := 2
	for x > n {
		if x%n == 0 {
			x /= n
			ar = append(ar, n)
		} else {
			if n == 2 {
				n++
			} else {
				n += 2
			}
		}
	}
	ar = append(ar, x)
	return ar
}

func (num *intp) doParam() {
	pf := num.PrimeFactor()
	*num = intp(pf[len(pf)-1])
}

func main() {
	n := 1234
	x := intp(n)
	// fmt.Printf("%d [%t]. \n", x, x.IsPrime())
	// x *= 2
	// x++
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
	x.doParam()
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
	x++
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
}
