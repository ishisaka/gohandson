package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a, b, c := 100, 200, 300
	fmt.Print("total: ")
	fmt.Println(a + b + c)

	// int32型xを宣言
	var x int32 = 100
	// int54に変換
	var y int64 = int64(x)
	// float32に変換
	var z float32 = float32(y)

	fmt.Println(z)

	// strconvによるテキストと整数値の返還
	var s string = strconv.Itoa(int(y))
	fmt.Println(s)
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("ERROR!")
		return
	}
	fmt.Println(i)

	// 定数
	const n int = 100
	m := float64(n) * 1.1
	fmt.Println(m)

	// 制御構文
	if x%2 == 0 {
		fmt.Println("偶数です")
	} else {
		fmt.Println("奇数です")
	}

	month := 12
	switch month {
	case 0:
		fmt.Println("月を示す数ではありません")
	case 1, 2, 12:
		fmt.Println("冬です")
	case 3, 4, 5:
		fmt.Println("春です")
	case 6, 7, 8:
		fmt.Println("夏です")
	case 9, 10, 11:
		fmt.Println("秋です")
	default:
		fmt.Println("月の数の指定がおかしいです")
	}

	// 配列
	ar := strings.Split("10 20 30 40 50", " ")
	t := 0
	for i := 0; i < len(ar); i++ {
		n, er := strconv.Atoi(ar[i])
		if er != nil {
			fmt.Println("ERROR")
			return
		}
		t += n
	}
	// 上をrangeで書き直す。
	t = 0
	for _, v := range ar {
		n, er := strconv.Atoi(v)
		if er != nil {
			fmt.Println("ERROR")
			return
		}
		t += n
	}

	fmt.Println("total:", t)

	// スライス
	sa := [5]int{10, 20, 30, 40, 50}
	sb := sa[0:3]
	fmt.Println(sa)
	fmt.Println(sb)
	// スライスは配列の参照。以下でsa、sbとも値が変化する。
	sa[0] = 100
	fmt.Println(sa)
	fmt.Println(sb)
	sb[1] = 200
	fmt.Println(sa)
	fmt.Println(sb)
	// スライスへの要素の追加
	sb = append(sb, 1000)
	fmt.Println(sa)
	fmt.Println(sb)
	sb = append(sb, 2000)
	fmt.Println(sa)
	fmt.Println(sb)
	sb = append(sb, 3000)
	fmt.Println(sa)
	fmt.Println(sb)
}
