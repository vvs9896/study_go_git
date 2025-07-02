package main

import "fmt"

func main() {
	fmt.Println("=== ПРОСТОЙ КАЛЬКУЛЯТОР ===\n")

	// Переменные для вычислений
	var a float64 = 15.5
	var b float64 = 4.2

	fmt.Printf("Число A: %.2f\n", a)
	fmt.Printf("Число B: %.2f\n", b)
	fmt.Println()

	// Арифметические операции
	sum := a + b
	difference := a - b
	product := a * b
	quotient := a / b

	fmt.Println("АРИФМЕТИЧЕСКИЕ ОПЕРАЦИИ:")
	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, sum)
	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, difference)
	fmt.Printf("%.2f × %.2f = %.2f\n", a, b, product)
	fmt.Printf("%.2f ÷ %.2f = %.2f\n", a, b, quotient)

	// Операции сравнения
	fmt.Println("\nОПЕРАЦИИ СРАВНЕНИЯ:")
	fmt.Printf("%.2f > %.2f: %t\n", a, b, a > b)
	fmt.Printf("%.2f < %.2f: %t\n", a, b, a < b)
	fmt.Printf("%.2f == %.2f: %t\n", a, b, a == b)
	fmt.Printf("%.2f != %.2f: %t\n", a, b, a != b)

	// Дополнительные вычисления
	fmt.Println("\nДОПОЛНИТЕЛЬНО:")
	average := (a + b) / 2
	fmt.Printf("Среднее арифметическое: %.2f\n", average)

	// Демонстрация целочисленного деления
	x := 17
	y := 5
	fmt.Printf("\nЦелочисленное деление: %d ÷ %d = %d (остаток: %d)\n", x, y, x/y, x%y)
}
