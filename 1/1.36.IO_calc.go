/*
 1.36.  Найти площадь кольца по заданному внешнему и внутреннему диаметру.

	На всякий случай: площадь окружности πr² или в случае с известным диаметром π(d/2)²
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

func main() {
	rand.Seed(time.Now().Unix()) // "сеим" случайности в основоной функции
	d1 := float64(get(1, 500))
	d2 := float64(get(501, 1000)) // не забываем, хотя задачу можно усложнить перегенерацией в случае если не будет соблюдаться условие d1<d2
	fmt.Printf("%s%.0f\n%s%.0f\n", "Внутренний диаметр ", d1, "Внешний диаметр ", d2)
	fmt.Printf("%s%.4f%s\n", "Площадь кольца ", math.Pi*math.Pow(d2/2, 2)-math.Pi*math.Pow(d1/2, 2), " чего-то²")
}