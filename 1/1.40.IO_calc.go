/*
1.40. составить программу вычисления функции
* 
* 		 2
 		----- + b
		a²+25
	x = ----------------
		      a + b
		√b + -------
		         2
и
	    |a|+2sin(b)
	y= -------------
		5.5a
 */

package main

import (
	"math/rand"
	"fmt"
	"math"
	"time"
)

func get(min, max int) (int) { // генерация псевдослучайных (!!! - не путать с действительно случайными) чисел
	return rand.Intn(max-min) + min
}

// Надоело каждый раз писать формулу в конструкции fmt.Printf

func calc_x(a, b float64) (x float64) {
	x = (2/math.Pow(a, 2) + b) / math.Sqrt(b+(a+b)/2)
	return x
}

// Решил оформить отдельными функциями, для упрощенного редактирования и изменения

func calc_y(a, b float64) (y float64) {
	y = (math.Abs(a) + 2*math.Sin(b)) / (5.5 * a)
	return y
}

func main() {
	rand.Seed(time.Now().Unix()) // "сеим" случайности в основоной функции
	a := float64(get(1, 5000))
	b := float64(get(1, 10000))
	fmt.Printf("%s%.0f\n%s%.0f\n", "a = ", a, "b = ", b)
	fmt.Printf("%s%.8f\n", "x = ", calc_x(a, b)) // всё таки мы доверчивые люди
	fmt.Printf("%s%.8f\n", "y = ", calc_y(a, b)) // по хорошему надо бы посчитать
} // правильную ли величину возвращает функция