/* Простейшие программы. Арифметические выражения.
1.16. Перевести из линейной записи в обычную следующие выражения:

а) a / b / c; 			з) a / sin(b);
б) a * b / c; 			и) 1 / 2 * a * b * sin(x);
в) a / b * c; 			к) 2 * b * c * cos(a / 2) / ( b + c);5
г) a + b / c; 			л) 4 * R * sin(a / 2) * sin(b / 2) * sin(c / 2);
д) (a + b) / c;			м) (a * x + b) / (c * x + d);
е) a + b / b + c; 		н) 2 * sin((a + b)/2 * cos((a - b) / 2);
ж) ( a + b) / ( b + c );	о) abs(2 * sin(-3 * abs(x / 2)))

 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d, R, x float64
	var m[6] float64

	for i := range m {
		switch {
		case i == 0: fmt.Print("a = ")
		case i == 1: fmt.Print("b = ")
		case i == 2: fmt.Print("c = ")
		case i == 3: fmt.Print("d = ")
		case i == 4: fmt.Print("R = ")
		case i == 5: fmt.Print("x = ")
		}
		fmt.Scanf("%f", &m[i])
	}
	a = m[0]
	b = m[1]
	c = m[2]
	d = m[3]
	R = m[4]
	x = m[5]

	fmt.Printf("%s%.4f\n","а) ", a/b/c)
	fmt.Printf("%s%.4f\n","б) ", a*b/c)
	fmt.Printf("%s%.4f\n","в) ", a/b*c)
	fmt.Printf("%s%.4f\n","г) ", a+b/c)
	fmt.Printf("%s%.4f\n","д) ", (a+b)/c)
	fmt.Printf("%s%.4f\n","е) ", a+b/b+c)
	fmt.Printf("%s%.4f\n","ж) ", (a+b)/(b+c))
	fmt.Printf("%s%.4f\n","з) ", a/math.Sin(b))
	fmt.Printf("%s%.4f\n","и) ", 1/2*a*b*math.Sin(x))
	fmt.Printf("%s%.4f\n","к) ", 2*b*c*math.Cos(a/2)/(b+c))
	fmt.Printf("%s%.4f\n","л) ", 4*R*math.Sin(a/2)*math.Sin(b/2)*math.Sin(c/2))
	fmt.Printf("%s%.4f\n","м) ", (a*x+b)/(c*x+d))
	fmt.Printf("%s%.4f\n","н) ", 2*math.Sin((a+b)/2) * math.Cos((a-b)/2))
	fmt.Printf("%s%.4f\n","о) ", math.Abs(2*math.Sin(-3*math.Abs(x/2))))
}
