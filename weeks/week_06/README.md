# 📅 Неделя 6: Интерфейсы в Go и Git collaboration

## 🎯 Цели недели
- Освоить концепцию интерфейсов в Go
- Понять полиморфизм и композицию интерфейсов
- Изучить advanced Git collaboration техники
- Создать гибкую архитектуру с интерфейсами

## 📚 Теория (Go 80% | Git 20%)

### Go (6 часов)
- **День 1-2**: Объявление и реализация интерфейсов
- **День 3-4**: Пустой интерфейс (interface{}), type assertions
- **День 5-6**: Композиция интерфейсов, интерфейсы в стандартной библиотеке
- **Выходные**: Рефакторинг существующего кода с интерфейсами

### Git (1.5 часа)
- **День 2**: Git hooks (pre-commit, pre-push)
- **День 4**: Code review best practices
- **День 6**: Conventional commits и автоматическое версионирование

## 🛠 Практические задания

### Go
1. **Базовые интерфейсы**:
   - Интерфейс Shape с методами Area() и Perimeter()
   - Реализация для различных фигур
   - Полиморфное использование

2. **Продвинутые интерфейсы**:
   - Интерфейс io.Reader и io.Writer
   - Создание custom Reader/Writer
   - Type assertions и type switches

3. **Архитектурные паттерны**:
   - Repository pattern с интерфейсами
   - Dependency injection
   - Mocking для тестирования

### Git
1. **Git hooks**:
   - Настройка pre-commit hook для проверки кода
   - Pre-push hook для запуска тестов
   - Automatic code formatting

2. **Collaboration**:
   - Code review процесс
   - Conventional commits
   - Semantic versioning

## 📖 Ресурсы для изучения

### Видео
- [Go Interfaces Explained](https://www.youtube.com/watch?v=KB3ysH8cupY)
- [Git Hooks Tutorial](https://www.youtube.com/watch?v=fMYv6-SZsSo)

### Документация
- [Go Interfaces](https://golang.org/doc/effective_go.html#interfaces)
- [Git Hooks Documentation](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks)

## 🏠 Домашнее задание

1. **Go**: Создать систему платежей с интерфейсами:
   - Интерфейс PaymentProcessor
   - Реализации: CreditCard, PayPal, Crypto
   - Интерфейс Logger для логирования
   - Использование dependency injection

2. **Git**: 
   - Настроить Git hooks для проекта
   - Внедрить conventional commits
   - Создать automated release workflow

## ✅ Критерии успеха
- [ ] Понимание концепции интерфейсов
- [ ] Умение создавать полиморфный код
- [ ] Владение type assertions и type switches
- [ ] Настроены Git hooks
- [ ] Завершена система платежей

## 🎯 Проект недели
**Файловый менеджер с интерфейсами**:
- Интерфейс FileSystem (Read, Write, Delete)
- Реализации: LocalFS, MemoryFS, S3FS
- Интерфейс Logger (Info, Error, Debug)
- Тестирование с моками

**Время на неделю**: ~14-18 часов 