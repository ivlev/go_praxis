package main

import ( 
	"fmt"
/*	"io"
	"os"
	"strings"*/
	)

func main() {
// r := strings.NewReader("Hello, Reader!")
fmt.Print("Введите сообщение: ")
var input string
fmt.Scanf("%s", &input)

output:=input
fmt.Print(output)
fmt.Println(" - вот такое сообщение вы ввели")

}
