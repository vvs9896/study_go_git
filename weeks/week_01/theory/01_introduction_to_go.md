# 🚀 Урок 1: Введение в Go

## 🎯 Что такое Go?

**Go (Golang)** - это современный язык программирования, созданный Google в 2009 году. Он разработан для создания быстрых, надежных и простых в сопровождении программ.

## 💡 Почему Go?

### ✅ Преимущества:
- **Простота**: Минимальный синтаксис, легко изучить
- **Скорость**: Компилируется в машинный код, работает быстро
- **Concurrency**: Встроенная поддержка параллельного выполнения
- **Надежность**: Строгая типизация, сборщик мусора
- **Кроссплатформенность**: Работает на Windows, Linux, macOS

### 🎯 Где используется Go:
- **Веб-серверы** (Uber, Netflix, Twitch)
- **Микросервисы** (Docker, Kubernetes)
- **Облачные приложения** (DropBox, SoundCloud)
- **DevOps инструменты** (Terraform, Prometheus)

## 🔧 Установка Go

### Linux/Ubuntu:
```bash
# Скачиваем последнюю версию
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz

# Извлекаем в /usr/local
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz

# Добавляем в PATH (в файл ~/.bashrc или ~/.zshrc)
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export GOROOT=/usr/local/go

# Перезагружаем терминал или выполняем
source ~/.bashrc  # или ~/.zshrc
```

### Проверка установки:
```bash
go version
# Должно показать: go version go1.21.0 linux/amd64
```

## 👨‍💻 Настройка среды разработки

### Рекомендуемые редакторы:
1. **VS Code** + Go Extension
2. **GoLand** (JetBrains)
3. **Vim/Neovim** + vim-go

### Настройка VS Code:
```bash
# Установка VS Code (если нет)
sudo snap install --classic code

# Запуск VS Code
code .
```

В VS Code установи расширение "Go" от Google.

## 🌟 Первая программа "Hello, World!"

Создадим первую программу:

```go
// main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("Привет, Go!")
}
```

### Разбор кода:
- `package main` - определяет исполняемый пакет
- `import "fmt"` - импортирует пакет для форматированного ввода/вывода
- `func main()` - точка входа в программу
- `fmt.Println()` - функция вывода текста с переносом строки

### Запуск программы:
```bash
# Запуск без компиляции
go run main.go

# Компиляция в исполняемый файл
go build main.go
./main

# Компиляция и установка
go install main.go
```

## 🧠 Основные концепции

### 1. Пакеты (Packages)
- Каждый файл Go принадлежит пакету
- `package main` - специальный пакет для исполняемых программ
- Другие пакеты используются как библиотеки

### 2. Импорты (Imports)
```go
import "fmt"           // Один импорт
import (               // Множественный импорт
    "fmt"
    "os"
    "strings"
)
```

### 3. Функции (Functions)
```go
func functionName(parameters) returnType {
    // тело функции
    return value
}
```

## 📝 Практическое задание

1. Установи Go на свой компьютер
2. Создай файл `hello.go` с программой "Hello, World!"
3. Запусти программу командой `go run hello.go`
4. Модифицируй программу, чтобы она выводила твое имя

### Пример решения:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("Меня зовут Влад")
    fmt.Println("Я изучаю Go!")
}
```

## ❓ Вопросы для самопроверки

1. Что такое пакет в Go?
2. Зачем нужна функция `main()`?
3. Какая команда компилирует Go-программу?
4. В чем разница между `go run` и `go build`?

## 🔗 Полезные ссылки

- [Официальный сайт Go](https://golang.org/)
- [Tour of Go](https://tour.golang.org/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html) 