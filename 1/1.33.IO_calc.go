/*
Известно количество жителей государства и площадб его территории.
Определить плотность населения в этом государстве. 

Я заранее сделал простой текстовый файл в котором необходимая и не очень
информация по странам разделена символом табуляции и положил его
в тот же каталог 1 главы: geo/1.33.CA.dat
 */

package main

import (
	"bufio"
	"fmt"
	"time"
	"os"
	"strconv"
	"strings"
	"math/rand"
)

func take(min, max int) int {
	rand.Seed(time.Now().Unix()) // А вот нельзя просто так взять и сгенерировать случайное число. Надо посеять "семеску" (seed)
	return rand.Intn(max - min) + min // И только потом - получить случайное число в нужном диапазоне
}

func main() {
	st := 0 // задаем порядковый номер первой строки файла
	file, _ := os.Open("geo/1.33.CA.dat") // открываем дескриптор файла со строчными данными по странам
	f := bufio.NewReader(file)  // считываем открытый файл в буфер
	stp := take(1, 169) // задаем случайную строку: не очень удачная идея - заранее задать количество строк, но с этим будем решать в следующий раз
	for {
		read_line, _ := f.ReadString('\n') // считываем файл построчно и переходим к следующей строке
		st = st + 1
		if st == stp { // когда доберемся до стучайно выхваченной строки
			dong := strings.Split(read_line, "\t") // разобъем её на массив по символу табуляции
			if a, err := strconv.Atoi(dong[2]); err == nil { // преобразуем элемент массива в третьей
				if b, err := strconv.Atoi(dong[3]); err == nil { // и четвертой позции в целое число
					fmt.Printf("%s %s\n", "Страна: ", dong[0]) // из первого элемента смастерим указатель
					fmt.Printf("%s %d %s\n" , "Плотность населения ", b/a, " чел/км²") // посчитаем плотность населения
					break
				}
			}
			file.Close()
		}
	}
}

