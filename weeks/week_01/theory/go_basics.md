# 🐹 Основы языка Go - Неделя 1

## 🎯 Что такое Go?

Go (также известный как Golang) - это язык программирования с открытым исходным кодом, разработанный Google в 2009 году. Создатели: Роберт Грайзмер, Роб Пайк и Кен Томпсон.

### Ключевые особенности Go:
- **Простота**: Минималистичный синтаксис
- **Быстрота**: Компилируется в машинный код
- **Параллельность**: Встроенная поддержка горутин
- **Типизация**: Статическая типизация с выводом типов
- **Сборка мусора**: Автоматическое управление памятью

## 🛠 Установка и настройка Go

### Шаг 1: Установка Go

#### Linux (Ubuntu/Debian):
```bash
# Скачиваем и устанавливаем Go
wget https://golang.org/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz

# Добавляем в PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

#### Проверка установки:
```bash
go version
# Должно вывести: go version go1.21.5 linux/amd64
```

### Шаг 2: Настройка рабочего пространства

```bash
# Создаем папку для Go проектов
mkdir -p ~/go-projects
cd ~/go-projects

# Инициализируем первый модуль
mkdir hello-world
cd hello-world
go mod init hello-world
```

### Шаг 3: Настройка VS Code

Установите расширение Go для VS Code:
1. Откройте VS Code
2. Перейдите в Extensions (Ctrl+Shift+X)
3. Найдите "Go" от Google
4. Установите расширение

## 📝 Синтаксис Go

### Базовая структура программы

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

**Разбор кода:**
- `package main` - объявление пакета main (точка входа)
- `import "fmt"` - импорт пакета для форматированного вывода
- `func main()` - главная функция программы
- `fmt.Println()` - функция вывода строки

### Переменные и константы

#### Объявление переменных:

```go
package main

import "fmt"

func main() {
    // Способ 1: Явное объявление типа
    var name string = "Влад"
    var age int = 25
    
    // Способ 2: Автоматический вывод типа
    var city = "Калининград"
    
    // Способ 3: Краткое объявление (только внутри функций)
    country := "Россия"
    
    fmt.Println("Имя:", name)
    fmt.Println("Возраст:", age)
    fmt.Println("Город:", city)
    fmt.Println("Страна:", country)
}
```

#### Нулевые значения:
```go
var i int       // 0
var f float64   // 0.0
var b bool      // false
var s string    // ""
```

#### Константы:
```go
const Pi = 3.14159
const (
    StatusOK = 200
    StatusNotFound = 404
)
```

### Базовые типы данных

```go
package main

import "fmt"

func main() {
    // Целые числа
    var age int = 25
    var bigNumber int64 = 9223372036854775807
    
    // Числа с плавающей точкой
    var price float64 = 99.99
    var temperature float32 = 23.5
    
    // Булевы значения
    var isStudent bool = true
    var hasJob bool = false
    
    // Строки
    var firstName string = "Влад"
    var lastName string = "Иванов"
    
    // Символы (руны)
    var firstLetter rune = 'В'
    
    fmt.Printf("Возраст: %d\n", age)
    fmt.Printf("Большое число: %d\n", bigNumber)
    fmt.Printf("Цена: %.2f\n", price)
    fmt.Printf("Температура: %.1f\n", temperature)
    fmt.Printf("Студент: %t\n", isStudent)
    fmt.Printf("Есть работа: %t\n", hasJob)
    fmt.Printf("Имя: %s %s\n", firstName, lastName)
    fmt.Printf("Первая буква: %c\n", firstLetter)
}
```

### Операторы

#### Арифметические операторы:
```go
a := 10
b := 3

sum := a + b        // 13
diff := a - b       // 7
product := a * b    // 30
quotient := a / b   // 3
remainder := a % b  // 1
```

#### Операторы сравнения:
```go
a := 10
b := 5

fmt.Println(a == b)  // false (равно)
fmt.Println(a != b)  // true (не равно)
fmt.Println(a > b)   // true (больше)
fmt.Println(a < b)   // false (меньше)
fmt.Println(a >= b)  // true (больше или равно)
fmt.Println(a <= b)  // false (меньше или равно)
```

#### Логические операторы:
```go
a := true
b := false

fmt.Println(a && b)  // false (И)
fmt.Println(a || b)  // true (ИЛИ)
fmt.Println(!a)      // false (НЕ)
```

### Конвертация типов

```go
package main

import "fmt"

func main() {
    var i int = 42
    var f float64 = float64(i)  // int в float64
    var u uint = uint(f)        // float64 в uint
    
    fmt.Printf("i = %d, f = %f, u = %d\n", i, f, u)
    
    // Конвертация в строку
    var s string = fmt.Sprintf("%d", i)
    fmt.Printf("Строка: %s\n", s)
}
```

## 📊 Форматированный вывод

### Printf форматирование:

```go
package main

import "fmt"

func main() {
    name := "Влад"
    age := 25
    height := 1.75
    isStudent := true
    
    // Основные спецификаторы
    fmt.Printf("Имя: %s\n", name)           // строка
    fmt.Printf("Возраст: %d\n", age)        // десятичное число
    fmt.Printf("Рост: %.2f м\n", height)    // число с плавающей точкой
    fmt.Printf("Студент: %t\n", isStudent)  // булево значение
    fmt.Printf("Возраст в 16-ричной: %x\n", age)  // шестнадцатеричное
    
    // Универсальный спецификатор
    fmt.Printf("Значение: %v\n", name)
    fmt.Printf("Тип и значение: %#v\n", name)
    fmt.Printf("Тип: %T\n", name)
}
```

## 💭 Комментарии

```go
package main

import "fmt"

// Это однострочный комментарий

/*
Это многострочный
комментарий
*/

func main() {
    // Комментарий к коду
    fmt.Println("Hello, World!") // Комментарий в конце строки
    
    /*
    Этот код можно закомментировать
    fmt.Println("Этот код не выполнится")
    */
}
```

## 🔧 Инструменты Go

### Основные команды:

```bash
# Запуск программы
go run main.go

# Компиляция программы
go build main.go

# Форматирование кода
go fmt main.go

# Загрузка зависимостей
go mod tidy

# Тестирование
go test

# Показать информацию о модуле
go mod graph
```

### Структура Go проекта:

```
my-project/
├── go.mod          # Файл модуля
├── go.sum          # Контрольные суммы зависимостей
├── main.go         # Главный файл
├── README.md       # Документация
└── pkg/            # Пакеты проекта
    └── utils/
        └── helper.go
```

## 🎯 Лучшие практики

1. **Именование**:
   - Используйте camelCase для переменных: `userName`, `isActive`
   - Используйте PascalCase для экспортируемых элементов: `UserName`, `IsActive`

2. **Форматирование**:
   - Всегда используйте `go fmt` для форматирования кода
   - Одна строка = одна инструкция

3. **Комментарии**:
   - Комментируйте экспортируемые функции и пакеты
   - Начинайте комментарии с имени функции/переменной

4. **Обработка ошибок**:
   - Всегда проверяйте ошибки
   - Обрабатывайте ошибки на том уровне, где можете их исправить

## 🔍 Полезные ресурсы

- [Go Tour](https://tour.golang.org/) - интерактивный учебник
- [Go by Example](https://gobyexample.com/) - примеры кода
- [Effective Go](https://golang.org/doc/effective_go.html) - лучшие практики
- [Go Documentation](https://golang.org/doc/) - официальная документация

## ✅ Чек-лист для самопроверки

- [ ] Go установлен и настроен
- [ ] VS Code с Go расширением работает
- [ ] Понимаю базовый синтаксис Go
- [ ] Умею объявлять переменные и константы
- [ ] Знаю основные типы данных
- [ ] Понимаю операторы и конвертацию типов
- [ ] Умею использовать форматированный вывод
- [ ] Знаю основные команды Go CLI 