/*
 1.34. составить программу решения линейного уравнения
 ax + b = 0 ( a != 0 )

Казалось бы всё просто: переносим в правую часть решение своится к банальному  ax = -b,
но решение не для всех вожможных "а" и "b".

 */

package main

import (
	"math/rand"
	"fmt"
	"time"
	_"math" // НА всякий случай "_" синтаксис для зарезервированых, но неиспользуемых пакетов
)


func get (min, max int) (int) { // генерация псевдослучайных (!!! - не путать с действительно случайными) чисел
	return rand.Intn(max-min) + min

}

func pn (c, d float64) (float64, float64) {
	f:=time.Now() // Три раза считываем время
	s:=time.Now()
	t:=time.Now()
	t1:=t.Sub(s) // Делаем два "среза" чтобы использовать их для увеличения случайности полученных чисел
	t2:=s.Sub(f)
	tr1:=float64(t2) // Приводим "время" к формату с плавающей точкой
	c = tr1*c

	f=time.Now()
	s=time.Now()
	t=time.Now()
	t1=f.Sub(s)
	t2=s.Sub(f)

	tr2:=float64(t1)

	d = tr2*d

	return c, d
}

func main ()  {
	rand.Seed(time.Now().Unix()) // "сеим" случайности в основоной функции
	c:= float64(get(1, 200))
	d:=float64(get(1, 100))
	pn(c, d)
	fmt.Printf("%s%.0f\n", "c=", c)
	fmt.Printf("%s%.0f\n", "d=", d)
	fmt.Printf("%s%f\n", "x= ", (-1*d)/c)
}