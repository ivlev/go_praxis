/*
1.52 Известна стоимость 1 кг конфет, печенья и яблок. Найти стоимость всей покупки,
если купили x кг конфет, у кг печенья и z кг яблок.

( 11 издание )
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"bufio"
)

func get(min, max int64) int64 { // генерация псевдослучайных (!!! - не путать с действительно случайными) чисел
return rand.Int63n(max-min) + min
}

func main() {
	rand.Seed(time.Now().Unix()) // "сеим" случайности в основной функции
//	cd := int(get(1, 320)) // количество покупок
//	fmt.Println(cd)
//	counts := make(map[string]int)
	file, _ := os.Open("foods/1.52.1.CD.dat") // открываем дескриптор файла со строчными данными по фруктам
	f := bufio.NewReader(file)  // считываем открытый файл в буфер
	fmt.Println(*f)

	for { // не очень хорошо заранее оговаривать количество продуктов
//		i := 12
//		cd := int(get(1, 320)) // количество покупок
		ReadLine, _ := f.ReadString('\n') // считываем файл построчно и переходим к следующей строке
		ReadLine = string(ReadLine)
// 		fmt.Println(cd)
//		if i == cd {
		fmt.Println(ReadLine)
//		}
	}

	file.Close()
}