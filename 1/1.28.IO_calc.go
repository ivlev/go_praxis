/*
1.28. Дана длина ребра куба. Найти объем куба и площадь его боковой поверхно-
сти.
 */

package main
import (
	"fmt"
	"math"
)

const r float64 = 6350
var h, c float64
var ch rune
var m string

func main() {

	fmt.Printf("%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s",
		"Пожалуйста, выберите единицу измерения выбрав соотвествующую цифру",
		"микрометры:\t1",
		"милиметры:\t2",
		"сантиметры:\t3",
		"метры:\t\t4",
		"километры:\t5",
		"световые года:\t6\n")
	fmt.Print("Введите выбранный номер: ")
	fmt.Scanf("%d", &ch)

	switch ch {
	default: m=" сам-сначала-определись-чего\n"
		c=1e+20
	case 1: m=" микрометрах"
		c=1e+6
	case 2: m=" милиметрах"
		c=1e+3
	case 3: m=" сантиметрах"
		c=1e+2
	case 4: m=" метрах"
		c=1
	case 5: m=" километрах"
		c=1e-3
	case 6: m=" световых годах"
		c=9.461e-15
	}
	fmt.Printf("%s%s%s","Ребро куба в", m, " равно ")
	fmt.Scanf("%f",&h)
	fmt.Printf("%s%.2f%s", "Объем куба равен ", math.Pow(h/c,3), " куб. м\n")
	fmt.Printf("%s%0.2f%s", "Площадь куба равна ", math.Pow(h/c,2)*6, " кв. м.\n")
}