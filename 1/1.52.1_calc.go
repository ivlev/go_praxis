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
	file, _ := os.Open("foods/1.52.1.FR.dat") // открываем дескриптор файла со строчными данными по странам
	f := bufio.NewReader(file)  // считываем открытый файл в буфер
//	fmt.Println(f)

	for i := 1; i<=321; i++ {
		cd := int(get(1, 320)) // количество покупок
		read_line, _ := f.ReadString('\n') // считываем файл построчно и переходим к следующей строке
		fmt.Println(cd)
		if i == cd {
		fmt.Printf("%d %s\n", i, read_line)
		}
	}

	file.Close()
}