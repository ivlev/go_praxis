/*
1.32. Известны объём и масса тела. Определить плотность материала этого тела.

Задача очень скучная. Куда интереснее, если мы, зная объём и массу, будем
определять из чего сделано представленное нами тело.

Для этого надо где-то брать плотность материалов и используя данные делать
предположение об объекте.

Файл с 3133 протностями разных веществ в виде разделенного табуляциями файла находится 1/chemistry/density.dat


 */

package main

import (
	"bufio"
	"fmt"
	"time"
	"os"
	_ "strconv"
	_ "strings"
	"math/rand"
	"regexp"
	"strconv"
)

func take(min, max int) int {
	rand.Seed(time.Now().Unix())    // А вот нельзя просто так взять и сгенерировать случайное число. Надо посеять "семеску" (seed)
	return rand.Intn(max-min) + min // И только потом - получить случайное число в нужном диапазоне
}

func main () {
	var validID= regexp.MustCompile(`\d{0,}[.]\d{0,}|\d{1,}`)
	st := 0                               // задаем порядковый номер первой строки файла
	file, _ := os.Open("chemistry/density.dat") // открываем дескриптор файла со строчными данными по странам
	f := bufio.NewReader(file)            // считываем открытый файл в буфер
	stp := take(0, 3132)                   // задаем случайную строку: не очень удачная идея - заранее задать количество строк, но с этим будем решать в следующий раз
	for {
		read_line, _ := f.ReadString('\n') // считываем файл построчно и переходим к следующей строке
		st = st + 1
		if st == stp { // когда доберемся до стучайно выхваченной строки
			fmt.Println(read_line)
			str_numbers := validID.FindAllString(read_line, -2)
			fmt.Printf("%T %v\n", str_numbers[0], str_numbers[0])
			fmt.Printf("%T %v\n", str_numbers[1], str_numbers[1])
			if a, err := strconv.ParseFloat(str_numbers[0],64); err == nil {
				fmt.Printf("%T %v\n", a, a)
			}
/*			if a, err := strconv.Atoi(str_numbers[0]); err == nil {
				//			dong := strings.Split(read_line, "\t") // разобъем её на массив по символу табуляции
				//			if a, err := strconv.Atoi(dong[2]); err == nil { // преобразуем элемент массива в третьей
				//				if b, err := strconv.Atoi(dong[3]); err == nil { // и четвертой позции в целое число
				//					fmt.Printf("%s %s\n", "Страна: ", dong[0]) // из первого элемента смастерим указатель
				//					fmt.Printf("%s %d %s\n" , "Плотность населения ", b/a, " чел/км²") // посчитаем плотность населения
			fmt.Printf("%T %v", a, a)
			*/
				break
			}
		}
		file.Close()
	}
