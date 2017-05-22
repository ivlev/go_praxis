/*
1.44. Даны длины сторон прямоугольного параллелепипеда. Найти его объем и площади боковой поверхности.
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

func vol(a, b, h float64) (v float64) {
	v = a*b*h
	return v
}

// Решил оформить отдельными функциями, для упрощенного редактирования и изменения

func sqr(a, b, h float64) (s float64) {
	s = 2*((a*b)+(a*h)+(b*h))
	return s
}



func main() {
	rand.Seed(time.Now().Unix()) // "сеим" случайности в основоной функции
	a := float64(get(1, 10000))
	b := float64(get(1, 10000))
	h := float64(get(1, 10000))
	fmt.Println("\nДан параллелепипед\n")
	fmt.Printf("%s%.0f\n%s%.0f\n%s%.0f\n\n", "a₁ (первая сторона) = ", a, "a₂ (вторая сторона) = ", b, "h (высота) = ", h)
	fmt.Printf("%s\t%.0f%s\n", "Объем параллелепипеда\t\t( a₁*a₂*h ): ", vol(a, b, h), " чего-то-там²")   // всё таки мы доверчивые люди
	fmt.Printf("%s\t%.0f%s\n\n", "Площадь параллелепипеда\t(2*(a₁*a₂+a₁*h+a₂*h)): ", sqr(a, b, h), " чего-то-там²") // по хорошему надо бы посчитать
} // правильную ли величину возвращает функция
