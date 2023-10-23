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

// 闭包
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// 实现一个 fibonacci 函数，它返回一个函数（闭包），该闭包返回一个斐波纳契数列 `(0, 1, 1, 2, 3, 5, ...)`。
func fibonacci() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f := f1
		f1 = f2
		f2 = f + f1
		return f
	}
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
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
	fmt.Println("===================Switch===================!")
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
	// fmt.Println("counting")

	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println(i)
	// }

	// fmt.Println("done")
	// 指针
	fmt.Println("===================指针===================!")
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(p)  // 查看指针 p 的值
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
	fmt.Println("===================struct===================!")
	// struct
	v := Vertex{1, 4}
	v.X = 3
	fmt.Println(v)
	fmt.Println(v.Abs())
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
	fmt.Println("===================数组===================!")
	// 数组
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	fmt.Println("===================切片===================!")
	// 切片
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a1 := names[0:2]
	b1 := names[1:3]
	fmt.Println(a1, b1)

	b1[0] = "XXX"
	fmt.Println(a1, b1)
	fmt.Println(names)

	fmt.Println("===================闭包===================!")
	// 闭包
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
