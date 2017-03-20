/*
1.18. Указать значение величины s после выполнения следующих операторов присваивания:
а) 	s:=5
	s:=57
б) 	s:=6
	s:=-5.2s
	s:=0
в)	s:=-7.5
	s:=2*s
г)	s:=45
	k:=-25
	s:=s+k
 */

/* будем считать, что это тренировка в приведении типов
есть две простые конструкции
var f float32 = float32(s) //и
f:=float23(s)

Для наглядности я воспользовался первой

Хотя существует пакет reflect ( https://golang.org/pkg/reflect/ )
преобразования черех отражение, с кототым ещё предстоит разобраться

https://ru.wikipedia.org/wiki/Рефлексия_(программирование)

 */
package main

import "fmt"
func a() {
	s:=5
	s=57
	fmt.Printf("%s%d\n\n","а) s:=5; s:=57\nв результате получаем ", s)
}

func b() {
	s:=6 				// Тип переменной назначается после её присвоения;
	var f float32 = float32(s) 	// Go - язык со строгой типизацией. Поэтому перенос в переменную другого типа - простейшая возможность перевода одного типа в другой;
	f=-5.2*f 			// дальше мы работаем с новой переменной забыв про первоначальную;
	f=0
	fmt.Printf("%s%.0f\n\n","б) s:=6; s:=-5.2s; s:=0\nв результате получаем ", f) 	// Итак, формально - это то же самое значение, а вот переменная другая / другого типа;
	// В толичие от школьников мы знаем, что переменная будет равна "0"
}
func v() {
	s:=-7.5
	s=2*s
	fmt.Printf("%s%.2f\n\n","в) s:=-7,5; s:=2*s\nв результате получаем ", s)
}
func g() {
	s:=45
	k:=-25
	fmt.Printf("%s%d\n\n","в) s:=-7,5; s:=2*s\nв результате получаем ", s+k)
}


func main() {
	a()
	b()
	v()
	g()
}
