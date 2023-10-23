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

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
