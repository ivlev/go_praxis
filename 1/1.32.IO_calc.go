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
	"strings"
)

func take(min, max int) int {
	rand.Seed(time.Now().Unix())    // А вот нельзя просто так взять и сгенерировать случайное число. Надо посеять "семеску" (seed)
	return rand.Intn(max-min) + min // И только потом - получить случайное число в нужном диапазоне
}

func main () {

//	st := 0                               // задаем порядковый номер первой строки файла
	file, _ := os.Open("chemistry/density.dat") // открываем дескриптор файла со строчными данными по странам
	f := bufio.NewReader(file)            // считываем открытый файл в буфер
//	stp := take(0, 3132)                   // задаем случайную строку: не очень удачная идея - заранее задать количество строк, но с этим будем решать в следующий раз
	wght := take(60, 10000000)
	cpct := take(45, 1000000)
	dnst := float64(wght)/float64(cpct)
	fmt.Printf("%s%v%s\n%s%v%s\n%s%v%s\n", "Объем тела: ", cpct, " см³", "Масса: ", wght, " г", "Плотность: ", dnst, " г/см³")

	var validID= regexp.MustCompile(`\t\d{0,}[.]\d{0,}|\t\d{0,}`) // подстрока по шаблону которой извлекаются данные
	for {
		read_line, _ := f.ReadString('\n') // считываем файл построчно и переходим к следующей строке
//		st = st + 1
//		if st == stp { // когда доберемся до стучайно выхваченной строки
//			fmt.Println(read_line)
			str_numbers := validID.FindAllString(read_line, -1)
			str_numbers[0] = strings.Trim(str_numbers[0],"\t")
//			fmt.Printf("%T %v\n", str_numbers[0], str_numbers[0])
			if a, err := strconv.ParseFloat(str_numbers[0],64); err == nil {
//				fmt.Printf("%T %v\n", a, a)
				if dnst/a > 0.99 && dnst/a < 1 {
					fmt.Println(read_line)

				}
//				else if dnst/a > 1 {
//					fmt.Println("Ничего не обнаружено")
//					break
				}
			}
/*			if a, err := strconv.Atoi(str_numbers[0]); err == nil {
				//			dong := strings.Split(read_line, "\t") // разобъем её на массив по символу табуляции
				//			if a, err := strconv.Atoi(dong[2]); err == nil { // преобразуем элемент массива в третьей
				//				if b, err := strconv.Atoi(dong[3]); err == nil { // и четвертой позции в целое число
				//					fmt.Printf("%s %s\n", "Страна: ", dong[0]) // из первого элемента смастерим указатель
				//					fmt.Printf("%s %d %s\n" , "Плотность населения ", b/a, " чел/км²") // посчитаем плотность населения
			fmt.Printf("%T %v", a, a)
			*/
	file.Close()
			}



