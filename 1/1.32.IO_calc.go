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
	file, _ := os.Open("chemistry/density.dat") // открываем дескриптор файла со строчными данными по странам
	f := bufio.NewReader(file)            	// считываем открытый файл в буфер
	wght := take(60, 10000000) 	// генерация случайного веса в граммах
	cpct := take(45, 1000000) 	// генерация случайного объёма в куб. см
	dnst := float64(wght)/float64(cpct)
	fmt.Printf("%s%v%s\n%s%v%s\n%s%v%s\n", "Объем тела: ", cpct, " см³", "Масса: ", wght, " г", "Плотность: ", dnst, " г/см³")

	var validID= regexp.MustCompile(`\t\d{0,}[.]\d{0,}|\t\d{0,}`) // подстрока по шаблону которой извлекаются данные
	for i:=0; i<3132;i++{
		read_line, _ := f.ReadString('\n') // считываем файл построчно и переходим к следующей строке
			str_numbers := validID.FindAllString(read_line, -1)
			str_numbers[0] = strings.Trim(str_numbers[0],"\t") // обрезаем из строки лишний символ "табуляция" ( \t ) - иначе невозможно преобразовать строку в число с плавающей точкой
			if a, err := strconv.ParseFloat(str_numbers[0],64); err == nil {
								if dnst/a > 0.99 && dnst/a < 1 {
					fmt.Println(read_line)
 								}
				} else { // а вот этот кусок почему-то не работает
				if dnst/a < 0.99 && dnst/a > 1{
					fmt.Println("Ничего похожего не обнаружено")
					break
				}
				}
				
			}
	file.Close()
			}



