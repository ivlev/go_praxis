/*
1.30. Составить программу:
а) вычисления значения функции z = x³-2,5xy+1,78x²-2,5y+1 при любых значениях х и y;
б) вычисления значения функции x = 3,56(a+b)³-5,8b²+3,8a-1,5 при любых значениях a и b.
 */
package main

import (
	"fmt"
	"math"
)

var ch  string

func zexp(){
	fmt.Println("\tВы не хотите ничего решать")
}

func fp()  {
	var f,s float64
	var i rune
	fmt.Println("\tСчитаем первую функцию")
	for i=0;i<2;i++ {
		switch i {
		case 0: fmt.Print("\tx = ")
			fmt.Scanf("%f", &f)
		case 1: fmt.Print("\ty = ")
			fmt.Scanf("%f", &s)
		}
	}

	fmt.Printf("%s%.3f\n", "\tz = ", math.Pow(f, 3)-2.5*f*s+1.78*math.Pow(f,2)-2.5*s+1)
}

func sp(){
	var f,s float64
	var i rune
	fmt.Println("\tСчитаем вторую функцию")
	for i=0;i<2;i++ {
		switch i {
		case 0: fmt.Print("\ta = ")
			fmt.Scanf("%f", &f)
		case 1: fmt.Print("\tb = ")
			fmt.Scanf("%f", &s)
		}
	}

	fmt.Printf("%s%.3f\n", "\tx = ", 5.56*math.Pow(f+s,2)-5.8*math.Pow(s,2)+3.8*f-1.5)
}

func main() {
	fmt.Println("\tКакую функцию будете вычислять: \"а\" или \"б\"? (введите букву)")
	fmt.Println("\tа) z = x³-2,5xy+1,78x²-2,5y+1")
	fmt.Println("\tб) x = 3,56(a+b)³-5,8b²+3,8a-1,5")
	fmt.Print("\tВаш выбор: ")
	fmt.Scanf("%s", &ch)

	switch ch {
	case string('a'), string('а'), string('ф'), string('f'): fp() // не надо забывать о конвертации типа string() - в противном случае получаем ошибку при использовании строковых данных в операторе switch
	case string('b'), string('б'), string('и'), string(','): sp() // и небольшое удобство для русскоязычного пользователя, не печатающего вслепую и забывающего переключать раскладку
	default: zexp()
}
}
