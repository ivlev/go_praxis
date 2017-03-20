/* Простейшие программы. Арифметические выражения.
1.15 получить линейную запись следующих выражений.

а) -1 / x в квадрате		ж) ( -b + 1/a ) / ( c/2 )
б) a / bc			з) 1 / 1 + ( a+b / 2 )
в) a/b * c			и) 1 / ( 1 + ( 1 / 2 + ( 1 / ( 2 + 3 /5 ) ) ) )
г) ( a + b ) / c		к) 2 в степени m
д) 5.45 * ( a + 2b ) / -2
е) -b + квадратный корень b в квадрате - 4ac / 2*a


 */
package main

import (
	"fmt"
	"math"
)

var a, b, c, x, m, r  float64

func main() {
	fmt.Println("Введите необходимые числа")
	fmt.Print("a=")
	fmt.Scanf("%f\n", &a)
	fmt.Print("b=")
	fmt.Scanf("%f\n", &b)
	fmt.Print("c=")
	fmt.Scanf("%f\n", &c)
	fmt.Print("x=")
	fmt.Scanf("%f\n", &x)
	fmt.Print("m=")
	fmt.Scanf("%f\n", &m)

	fmt.Printf("%s%f\n","-1/x = ", -1/x)
	fmt.Printf("%s%f\n","a / bc = ", a/(c*b))
	fmt.Printf("%s%f\n","( a/b ) * c = ", (a/b)*c)
	fmt.Printf("%s%f\n","5.45 * ( a + 2b ) / -2 = ", 5.45*((a+2*b)/-2))
	fmt.Printf("%s%f\n","-b + квадратный корень b в квадрате - 4ac / 2*a = ", (-b + math.Sqrt(math.Pow(b, 2)-4*a*c))/2*a)
	fmt.Printf("%s%f\n","( -b + 1/a ) / ( c/2 ) = ", (-b+(1/a))/(c/2))
	fmt.Printf("%s%f\n","( 1 / 1 + ( a+b / 2 ) = ", 1/(1+(a+b/2)))
	fmt.Printf("%s%T\n","тип результата 1.0 / ( 1.0 + ( 1.0 / 2.0 + ( 1.0 / ( 2.0 + 3.0 /5.0 ) ) ) ) = ", 1.0/(1.0+(1.0/(2.0+(1.0/(2.0+(3.0/5.0)))))))
	r = 1.0 / ( 1.0 + ( 1.0 / 2.0 + ( 1.0 / ( 2.0 + 3.0 /5.0 ) ) ) )
	fmt.Printf("%s%f\n","1.0 / ( 1.0 + ( 1.0 / 2.0 + ( 1.0 / ( 2.0 + 3.0 /5.0 ) ) ) ) = ", r)
	fmt.Printf("%s%f\n"," 2 в степени m  = ", math.Pow(2, m))
}

