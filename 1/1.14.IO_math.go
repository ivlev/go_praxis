/* Простейшие программы. Арифметические выражения.
1.14 Записать согласно правилам изучаемого языка программирования
следующие выражения.

а) 2x				з) 3 квадратных корня x
б) sinx				и) sinacosb + cosasinb
в) а в квадрате			к) a квадратных корней 2b
г) квадратный корень x		л) 3sin2acos3b
д) модуль n			м) -5 квадратных корней из x + квадратный корень из y
е) 5cosy
ж) -7,5 x в квадрате

 */

package main

import (
	"fmt"
	"math"
)
var a, b, n, x, y float64

func main() {
	fmt.Println("Введите необходимые числа")
	fmt.Print("a=")
	fmt.Scanf("%f\n", &a)
	fmt.Print("b=")
	fmt.Scanf("%f\n", &b)
	fmt.Print("n=")
	fmt.Scanf("%f\n", &n)
	fmt.Print("x=")
	fmt.Scanf("%f\n", &x)
	fmt.Print("y=")
	fmt.Scanf("%f\n", &y)

	fmt.Printf("%s%.2f\n","2*x = ", 2*x)
	fmt.Printf("%s%.2f\n", "sin(X) = ", math.Sin(x))
	fmt.Printf("%s%.2f\n", "а в квадрате = ", math.Pow(a, 2))
	fmt.Printf("%s%.2f\n", "|n| = ", math.Abs(n))
	fmt.Printf("%s%.2f\n", "5 * cos(Y) = ", 5*math.Cos(y))
	fmt.Printf("%s%.2f\n", "-7.5 * X в квадрате = ", -7.5*math.Pow(x, 2))
	fmt.Printf("%s%.2f\n", "3 * квадратный корень X = ", 3*math.Sqrt(x))
	fmt.Printf("%s%.2f\n", "sin(a)cos(b) + cos(a)sin(b) = ", math.Sin(a)*math.Cos(b)+math.Cos(a)*math.Sin(b))
	fmt.Printf("%s%.2f\n", "a квадратных корней 2*b = ", a*math.Sqrt(2*b))
	fmt.Printf("%s%.2f\n", "3*sin(2*a)*cos(3*b) = ", 3*math.Sin(2*a)*math.Cos(3*b))
	fmt.Printf("%s%.2f\n", "sin(a)cos(b) + cos(a)sin(b) = ", math.Sin(a)*math.Cos(b)+math.Cos(a)*math.Sin(b))
	fmt.Printf("%s%.2f\n", "-5 квадратных корней из x + квадратный корень из y = ", -5*math.Sqrt(x + math.Sqrt(y)))
}
