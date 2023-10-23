package main

import (
	"fmt"
	"math"
	"runtime"
	"time"

	"rsc.io/quote"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func sqrt2(x float64) string {
	// 用牛顿法实现平方根函数。
	// 这里的思路是首先猜测一个值 z 然后不断令 z^2 逼近 x 的值。
	// 通过 z -= (z*z - x) / (2*z) 不断调整 z 的值。
	z := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return fmt.Sprint(z)
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}

func main() {
	defer fmt.Println("world")
	// loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(sqrt2(2), sqrt2(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println(quote.Go())

	// switch
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// defer 栈 先进后出
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	// 指针
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(p)  // 查看指针 p 的值
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
	// struct
	type Vertex struct {
		X int
		Y int
	}
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)
	// 结构体指针
	p2 := &v
	p2.X = 1e9
	fmt.Println(v)
	var (
		v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
		v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
		v3 = Vertex{}      // X:0 Y:0
		p3 = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
	)

	fmt.Println(v1, p3, v2, v3)
}
