/*
1.42. Даны стороны прямоугольника. Найти его периметр и длину диагонали.

One more!

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

func per(a, b float64) (x float64) {
	x = (a + b) * 2
	return x
}

// Решил оформить отдельными функциями, для упрощенного редактирования и изменения

func diag(a, b float64) (y float64) {
	y = math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	return y
}

func main() {
	rand.Seed(time.Now().Unix()) // "сеим" случайности в основоной функции
	a := float64(get(1, 5000))
	b := float64(get(1, 10000))
	fmt.Printf("%s%.0f\n%s%.0f\n", "a₁ (первая сторона) = ", a, "a₂ (вторая сторона) = ", b)
	fmt.Printf("%s%.8f\n", "Периметр ( (a₁+a₂) * 2 ): ", per(a, b))   // всё таки мы доверчивые люди
	fmt.Printf("%s%.8f\n", "Диагональ ( √ a₁² + a₂² ): ", diag(a, b)) // по хорошему надо бы посчитать
} // правильную ли величину возвращает функция
