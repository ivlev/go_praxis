/*
1.51. Дано вещественное число a. пользуясь только операцией умножения, получить:
* 
а) а⁴ за две операции;
б) а⁶ за три операции;
в) а⁷ за четыре операции;
г) а⁸ за три операции;
д) а⁹ за четыре операции;
е) а¹⁰ за четыре операции;
ж) а¹³ за пять операций;
з) а¹⁵ за пять операций;
и) а²¹ за шесть операций;
к) а²⁸ за шесть операций;
л) а¹³ за шесть операций;
Не очень понимаю, что такое. в трактовке автора задачи, "операция"?
"операция" умножения? "Операция" на языке программирования?
Все рекомедуют воспользоваться "промежуточными" переменными.
*/

package main 

import (
	"math/rand"
	"fmt"
	_"math"
	"time"
)

	func get(min, max int64) (int64) { // генерация псевдослучайных (!!! - не путать с действительно случайными) чисел
	return rand.Int63n(max-min) + min
}

	func sqr(a int64) (int64) { // возведение в квадрат
	return a*a
}

	func sqr3(a int64) (int64) { // возведение в куб
	return sqr(a)*a
}

	func sqr4(a int64) (int64) { // возведение в четвертую степень
	return sqr3(a)*a
}

func sqr5(a int64) (int64) { // возведение в пятую степень
	return sqr4(a)*a
}

func main () {
	rand.Seed(time.Now().Unix()) // "сеим" случайности в основоной функции
	d := get(0, 10)
	a := int64(get(-d, d))
	fmt.Println("\n\tСлучайно выбрано число\t", a, "\n")
	fmt.Printf("%s\t%d\n", "а⁴\tза две операции ", sqr(a)*a*a)
	fmt.Printf("%s\t%d\n", "а⁶\tза три операции ", sqr(a)*sqr(a)*sqr(a))
	fmt.Printf("%s\t%d\n", "а⁷\tза четыре операции ", sqr(a)*sqr(a)*sqr(a)*a)
	fmt.Printf("%s\t%d\n", "а⁸\tза три операции ", sqr3(a)*sqr3(a)*sqr(a))
	fmt.Printf("%s\t%d\n", "а⁹\tза четыре операции ", sqr3(a)*sqr3(a)*sqr(a)*a)
	fmt.Printf("%s\t%d\n", "а¹³\tза пять операций ", sqr3(a)*sqr3(a)*sqr3(a)*sqr(a)*a*a)
	fmt.Printf("%s\t%d\n", "а¹⁵\tза пять операций ", sqr3(a)*sqr3(a)*sqr3(a)*sqr3(a)*sqr(a)*a)
	fmt.Printf("%s\t%d\n", "а²¹\tза шесть операций ", sqr4(a)*sqr4(a)*sqr3(a)*sqr3(a)*sqr3(a)*sqr3(a)*a)
	fmt.Printf("%s\t%d\n", "а²⁸\tза шесть операций ", sqr5(a)*sqr5(a)*sqr4(a)*sqr4(a)*sqr4(a)*sqr4(a)*sqr(a))
	fmt.Printf("%s\t%d\n\n", "а¹³\tза шесть операций ", sqr(a)*sqr(a)*sqr(a)*sqr(a)*sqr(a)*sqr(a)*a)
}

