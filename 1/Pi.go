/* Простейшие программы. Арифметические выражения.
1.1. Вывести на экран с точностью два знака число Пи
1.2 Вывести на экран с точностью два знака число е (основание натурального логарифма)
*/
package main

import "fmt"
import "math"

func main() {

var pi = math.Pi
var e = math.E

fmt.Printf("%.2f\n", pi)
fmt.Printf("%.2f\n", e)

}
