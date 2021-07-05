package main

import "fmt"

func main() {
	// List 3-1
	// n := 123
	// p := &n
	// fmt.Println("number:", n)
	// fmt.Println("pointer:", p)
	// fmt.Println("value:", *p)

	// List 3-2
	// n := 123
	// p := &n
	// m := 10000
	// p2 := &m
	// fmt.Printf("p value:%d, address:%p\n", *p, p)
	// fmt.Printf("p2 value:%d, address:%p\n", *p2, p2)
	// pb := p
	// p = p2
	// p2 = pb
	// fmt.Printf("p value:%d, address:%p\n", *p, p)
	// fmt.Printf("p2 value:%d, address:%p\n", *p2, p2)

	// List 3-3
	// n := 123
	// p := &n
	// q := &p
	// m := 10000
	// p2 := &m
	// q2 := &p2
	// fmt.Printf("q value:%d, address:%p\n", **q, *q)
	// fmt.Printf("q2 value:%d, address:%p\n", **q2, *q2)
	// pb := p
	// p = p2
	// p2 = pb
	// fmt.Printf("q value:%d, address:%p\n", **q, *q)
	// fmt.Printf("q2 value:%d, address:%p\n", **q2, *q2)

	// List 3-4
	// n := 123
	// fmt.Printf("valu: %d.\n", n)
	// change1(n)
	// fmt.Printf("valu: %d.\n", n)
	// change2(&n)
	// fmt.Printf("valu: %d.\n", n)

	// List 3-5
	nr := []int{10, 20, 30}
	fmt.Println(nr)
	initial(&nr)
	fmt.Println(nr)
}

// List 3-4
// func change1(n int) {
// 	n *= 2
// }
//
// func change2(n *int) {
// 	*n *= 2
// }

// List 3-5
func initial(nr *[]int) {
	for i := 0; i < len(*nr); i++ {
		(*nr)[i] = 0
	}
}
