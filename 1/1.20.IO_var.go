/*

1.20. Указать значения величин s и k после выполнения следующих операторов присваивания:

a)	s:=14
	k:=-3
	d:=s+1
	s:=d
	k:=2*s

б)	s:=0
	k:=30
	d:=k-5
	k:=2*d
	s:=k-100

 */
package main

import (
	"fmt"
)

func a() { // если обявлен "проект", то попытка присвоить функции имя одной из переменных другого файла воспринимается как ошибка

	s:=14
	k:=-3
	d:=s+1
	s=d
	k=2*s

	fmt.Printf("%s%d\n\t%s%d\n", "a)\tПеременная s=", s, "Переменная k=",k)

}

func b() {

	s:=0
	k:=30
	d:=k-5
	k=2*d
	s=k-100

	fmt.Printf("%s%d\n\t%s%d\n", "б)\tПеременная s=", s, "Переменная k=",k)

}

func main() {
	a()
	b()
}
