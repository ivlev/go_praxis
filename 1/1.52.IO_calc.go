/*
1.52. Дано вещественное число a. пользуясь только операцией умножения, получить:

а) а³ и a¹⁰ за четыре операции;
б) а⁴ и а²⁰ за пять операций;
в) а⁵ и а¹³ за пять операций;
г) а⁵ и а¹⁹ за пять операций;
д) а², а⁵ и а¹⁷ за шесть операций;
е) а⁴, а¹² и а²⁸ за шесть операции.

*/

package main

import (
	"fmt"
	_ "math"
	"math/rand"
	"time"
)

func get(min, max int64) int64 { // генерация псевдослучайных (!!! - не путать с действительно случайными) чисел
	return rand.Int63n(max-min) + min
}

func sqr(a int64) int64 { // возведение в квадрат
	return a * a
}

func sqr3(a int64) int64 { // возведение в куб
	return sqr(a) * a
}

func sqr4(a int64) int64 { // возведение в четвертую степень
	return sqr(a) * sqr(a)
}

func sqr5(a int64) int64 { // возведение в пятую степень
	return sqr4(a) * a
}

func main() {
	rand.Seed(time.Now().Unix()) // "сеим" случайности в основной функции
	d := get(0, 50)
	a := int64(get(-d, d)) // для большего правдоподобия выбираем как "положительные" так и "отрицательные" числа
	fmt.Println("Случайно выбрано число: ", a)
	fmt.Printf("%s%d%s%d\n", "а³ и a¹⁰ за четыре операции ", sqr3(a), " и ", sqr4(a)*sqr(a)*sqr(a)*sqr(a)) // во второ строке поддерживается правило положительного резулльтата
	// для четных степеней отрицательных чисел
	fmt.Printf("%s%d%s%d\n", "а⁴ и а²⁰ за пять операций ", sqr4(a), " и ", sqr4(a)*sqr4(a)*sqr4(a)*sqr4(a)*sqr4(a))
	fmt.Printf("%s%d%s%d\n", "а⁵ и а¹³ за пять операций ", sqr5(a), " и ", sqr3(a)*sqr4(a)*sqr(a)*sqr(a)*sqr(a))
	fmt.Printf("%s%d%s%d\n", "а⁵ и а¹⁹ за пять операций ", sqr5(a), " и ", sqr5(a)*sqr5(a)*sqr5(a)*sqr(a)*sqr(a))
	fmt.Printf("%s%d%s%d%s%d\n", "а², а⁵ и а¹⁷ за шесть операций ", sqr(a), ", ", sqr5(a), " и ", sqr5(a)*sqr3(a)*sqr3(a)*sqr(a)*sqr(a))
	fmt.Printf("%s%d%s%d%s%d\n", "а⁴, а¹² и а²⁸ за шесть операции ", sqr4(a), ", ", sqr4(a)*sqr4(a)*sqr4(a), " и ", sqr5(a)*sqr5(a)*sqr5(a)*sqr5(a)*sqr5(a)*sqr3(a))

}
