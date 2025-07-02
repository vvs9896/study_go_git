package main

import "fmt"

func main() {
	fmt.Println("=== ДЕМОНСТРАЦИЯ ПЕРЕМЕННЫХ В GO ===\n")

	// 1. Разные способы объявления переменных
	fmt.Println("1. Объявление переменных:")

	var name string = "Влад" // Полное объявление
	var age int = 25         // С указанием типа
	var city = "Калининград" // Автоматический вывод типа
	country := "Россия"      // Краткое объявление

	fmt.Printf("   Имя: %s (тип: %T)\n", name, name)
	fmt.Printf("   Возраст: %d (тип: %T)\n", age, age)
	fmt.Printf("   Город: %s (тип: %T)\n", city, city)
	fmt.Printf("   Страна: %s (тип: %T)\n", country, country)

	// 2. Нулевые значения
	fmt.Println("\n2. Нулевые значения:")

	var emptyString string
	var emptyInt int
	var emptyFloat float64
	var emptyBool bool

	fmt.Printf("   string: '%s' (длина: %d)\n", emptyString, len(emptyString))
	fmt.Printf("   int: %d\n", emptyInt)
	fmt.Printf("   float64: %f\n", emptyFloat)
	fmt.Printf("   bool: %t\n", emptyBool)

	// 3. Константы
	fmt.Println("\n3. Константы:")

	const Pi = 3.14159
	const AppName = "Go Learning App"
	const Version = "1.0.0"

	fmt.Printf("   π = %.5f\n", Pi)
	fmt.Printf("   Приложение: %s\n", AppName)
	fmt.Printf("   Версия: %s\n", Version)

	// 4. Множественное объявление
	fmt.Println("\n4. Множественное объявление:")

	var (
		firstName = "Влад"
		lastName  = "Иванов"
		birthYear = 1998
		isStudent = true
	)

	fmt.Printf("   Полное имя: %s %s\n", firstName, lastName)
	fmt.Printf("   Год рождения: %d\n", birthYear)
	fmt.Printf("   Студент: %t\n", isStudent)

	// 5. Конвертация типов
	fmt.Println("\n5. Конвертация типов:")

	var number int = 42
	floatNumber := float64(number)
	stringNumber := fmt.Sprintf("%d", number)

	fmt.Printf("   int: %d\n", number)
	fmt.Printf("   float64: %.2f\n", floatNumber)
	fmt.Printf("   string: '%s'\n", stringNumber)

	fmt.Println("\n=== КОНЕЦ ДЕМОНСТРАЦИИ ===")
}
