/*
1.39. составить программу вычисления функции
 		
 				 2+y		
 			x + -----
				  x²
	z = --------------------
				1
		y + ----------
		     √ x²+10
и

	q = 2.8sin(x)+|y|
		
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
	x := float64(get(1, 500)) 
  	y := float64(get(501, 1000)) 
	fmt.Printf("%s%.0f\n%s%.0f\n", "x= ",  x, "y= ", y)
	fmt.Printf("%s%.8f\n", "z= ", (x+(2+y)/math.Pow(x, 2))/(y + (1/(math.Sqrt(math.Pow(x, 2)+10)))))
	fmt.Printf("%s%.8f\n", "q= ", 2.8*math.Sin(x) + math.Abs(y))
}
