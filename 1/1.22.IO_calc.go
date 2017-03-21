/*
Вычисления по известным формулам

1.22. Составить программу:
а) вычисления значения функции y=7x²-3x+6 при любом значении x;
б) вычисления значения функции x=12a²+7a-16 при любом значении а.
 */

package main

import (
	"fmt"
	"math"
)

var x, a float64

func input() {
	for i:=0;i<=1;i++{
	switch i {
	case 0: fmt.Print("x=")
	fmt.Scanf("%f",&x)
	case 1: fmt.Print("a=")
	fmt.Scanf("%f",&a)
	}
	}
}

func Y() {
	fmt.Printf("%s%.4f\n","y=", 7*math.Pow(x,2)-3*x+6)
}

func X() {
	fmt.Printf("%s%.4f\n","x=", 12*math.Pow(a,2)+7*a-16)
}

func main ()  {
	input()
	Y()
	X()
}


