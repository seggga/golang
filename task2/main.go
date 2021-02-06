/*
2. Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга.
Площадь круга должна вводиться пользователем с клавиатуры.
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	var s, c, d float64
	var userInput string

	fmt.Println("Программа вычисляет диаметр и длину окружности по заданной площади круга")
	fmt.Printf("введите площадь круга S: ")
	fmt.Scanln(&userInput)

	// перевод введенных данных в числовой формат
	s, err := strconv.ParseFloat(userInput, 64)

	// проверка на наличие ошибки
	if err != nil {
		// есть ошибка, выдается сообщение, вычисления не производятся
		fmt.Printf("введенные данные не удается перевести в число.")
	} else {
		// ошибок нет, производятся вычисления и выводится результата
		d = math.Sqrt(4 * s / math.Pi)
		c = math.Pi * d

		fmt.Printf("Диаметр окружности составляет %.2f, длина окружности - %.2f единиц", d, c)
	}
}