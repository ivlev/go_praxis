/*
1.23. Составить программу вычисления значения функции

y=a²+10/√a²+1

при любом значении а.
 */

package main

import (
	"fmt"
	"math"
)

	var a, y float64

func main() {
	fmt.Print("Введите\ta=")
	fmt.Scanf("%f\n", &a)
	fmt.Printf("\t%s%f\n", "y=", (math.Pow(a,2)+10)/math.Sqrt(math.Pow(a,2)+1))
}