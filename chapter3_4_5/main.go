// List 3-32
package main

import (
	"fmt"
	"strconv"
	"time"
)

func prmsg(n int, s string) {
	fmt.Println(s)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

func first(n int, c chan string) {
	const nm string = "first-"
	for i := 0; i < 10; i++ {
		a := nm + strconv.Itoa(i)
		prmsg(n, a)
		c <- a
	}
}

func secound(n int, c chan string) {
	for i := 0; i < 10; i++ {
		prmsg(n, "secound:["+<-c+"]")
	}
}

func main() {
	c := make(chan string)
	go first(10, c)
	secound(10, c)
	fmt.Println()
}
