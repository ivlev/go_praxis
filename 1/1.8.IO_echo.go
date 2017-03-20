/* Простейшие программы. Арифметические выражения.
1.8. составвить программу выводда на экран
в одну строку четырех любых чисел
с одним пробелом */

package main

import "fmt"

func main () {

var a[4] int

for i:=0; i<= 3; i++ {
fmt.Printf("%s%d%s", "Введите число № ", i+1, " ")
fmt.Scanf("%d%s", &a[i],"\n")
}

for i:=0; i<= 3; i++ {
fmt.Printf("%d ", a[i])
}
fmt.Println()
}
