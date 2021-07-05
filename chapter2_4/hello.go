package main

import "fmt"

func main() {
	// List 2-24
	// m := []string{}
	// m, _ = push(m, "apple")
	// m, _ = push(m, "banana")
	// m, _ = push(m, "orange")
	// fmt.Println(m)
	// m, v := pop(m)
	// fmt.Println("Get "+v+" ->", m)

	// List 2-25
	// m := []string{"one", "two", "three"}
	// fmt.Println(m)
	// m = insert(m, "*", 2)
	// m = insert(m, "*", 1)
	// fmt.Println(m)

	// List 2-26
	// 	m := []string{
	// 		"one", "two", "three",
	// 	}
	// 	fmt.Println(m)
	// 	m = push(m, "1", "2", "3")
	// 	fmt.Println(m)

	//List 2-27
	// f := func(a []string) ([]string, string) {
	// 	return a[1:], a[0]
	// }
	// m := []string{
	// 	"one", "two", "three",
	// }
	// s := ""
	// fmt.Println(m)
	// for len(m) > 0 {
	// 	m, s = f(m)
	// 	fmt.Println(s+" ->", m)
	// }

	// List 2-28
	// modify := func(a []string, f func([]string) []string) []string {
	// 	return f(a)
	// }
	//
	// m := []string{
	// 	"1st", "2nd", "3rd",
	// }
	// fmt.Println(m)
	//
	// m1 := modify(m, func([]string) []string {
	// 	return append(m, m...)
	// })
	// fmt.Println(m1)
	//
	// m2 := modify(m, func([]string) []string {
	// 	return m[:len(m)-1]
	// })
	// fmt.Println(m2)
	//
	// m3 := modify(m, func([]string) []string {
	// 	return m[1:]
	// })
	// fmt.Println(m3)

	// List 2-29
	data := "*新しい値*"
	m1 := modify(data)
	data = "*new data*"
	m2 := modify(data)

	fmt.Println(m1())
	fmt.Println(m2())

	// List 2-30
	n := 123
	b := true
	a := "Hello"
	fmt.Printf("number:%d, bool:%t, string:%s.\n", n, b, a)
}

// List 2-24
// func push(a []string, v string) ([]string, int) {
// 	return append(a, v), len(a)
// }

// func pop(a []string) ([]string, string) {
// 	return a[:len(a)-1], a[len(a)-1]
// }

/* List 2-25 */
// func insert(a []string, v string, p int) (s []string) {
// 	s = append(a, "")
// 	s = append(s[:p+1], s[p:len(s)-1]...)
// 	s[p] = v
// 	return
// }

/* List 2-26 */
// func push(a []string, v ...string) (s []string) {
// 	s = append(a, v...)
// 	return
// }

// List 2-29
func modify(d string) func() []string {
	m := []string{
		"1st", "2nd",
	}
	return func() []string {
		return append(m, d)
	}
}
