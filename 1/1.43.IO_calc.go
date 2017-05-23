/*
1.43. Даны два числа. Найти их сумму, разность, произведение, а так же частное от деления первого числа на второе.
*/

package main

import (
	"math/rand"
	"fmt"
	_"math"
	"time"
)

func get(min, max int) (int) { // генерация псевдослучайных (!!! - не путать с действительно случайными) чисел
	return rand.Intn(max-min) + min
}

// Надоело каждый раз писать формулу в конструкции fmt.Printf

func sum(a, b float64) (x float64) {
	x = a + b
	return x
}

// Решил оформить отдельными функциями, для упрощенного редактирования и изменения

func diff(a, b float64) (y float64) {
	y = a - b
	return y
}

func or(a, b float64) (z float64) {
	z = a * b
	return z
}

func div(a, b float64) (d float64) {
	d = a / b
	return d
}


func main() {
	rand.Seed(time.Now().Unix()) // "сеим" случайности в основоной функции
	a := float64(get(1, 10000))
	b := float64(get(1, 10000))
	fmt.Printf("%s%.0f\n%s%.0f\n", "a₁ (первое число) = ", a, "a₂ (второе число) = ", b)
	fmt.Printf("%s\t%.0f\n", "Сумма\t\t(a₁+a₂): ", sum(a, b))   // всё таки мы доверчивые люди
	fmt.Printf("%s\t%.0f\n", "Разность\t(a₁-a₂): ", diff(a, b)) // по хорошему надо бы посчитать
	fmt.Printf("%s\t%.0f\n", "Произведение\t(a₁*a₂): ",  or(a, b)) // по хорошему надо бы посчитать
	fmt.Printf("%s\t%.8f\n", "Частное\t\t(a₁/a₂): ", div(a, b)) // по хорошему надо бы посчитать
} // правильную ли величину возвращает функция
