/*
1.24. Составить программу:
а) вычисления значения функции x=√(2a+sin|3a|/3.56) при любом значении а;
б) вычисления значения функции y=sin((3.2+√1+x)/5x) при любом значении х.
 */

package main

import (
	"fmt"
	"math"
)

var x, a float64

func input() { // функция ввода "на вырост". Вдруг потребуются дополнения.
	for i:=0;i<=1;i++{
		switch i {
		case 0: fmt.Print("a=")
			fmt.Scanf("%f",&a) // при отрицательных значениях мы получаем значение NaN
		case 1: fmt.Print("x=")
			fmt.Scanf("%f",&x) // при отрицатеьных значениях мы получаем значение NaN
		}
	}
}

func X() {

	fmt.Printf("%s%f\n", "√(2a+sin|3a|/3.56)=", math.Sqrt(2*a+math.Sin(math.Abs(3*a)/3.56)))
}
func Y() {
	fmt.Printf("%s%f\n", "y=sin((3.2+√1+x)/5x)=", math.Sin((3.2+math.Sqrt(1+x))/5*x))
}

func main() {
	input()
	X()
	Y()
}
