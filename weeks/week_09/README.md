# 📅 Неделя 9: Тестирование в Go и командная работа с Git

## 🎯 Цели недели
- Освоить тестирование в Go (unit, integration, benchmark тесты)
- Понять TDD подход в разработке
- Изучить Git workflow для команд
- Создать полностью протестированное приложение

## 📚 Теория (Go 80% | Git 20%)

### Go (6 часов)
- **День 1-2**: Unit тесты, testing package, table-driven tests
- **День 3-4**: Mocking, testify library, integration тесты
- **День 5-6**: Benchmark тесты, code coverage, TDD cycle
- **Выходные**: Добавление тестов ко всем предыдущим проектам

### Git (1.5 часа)
- **День 2**: GitHub/GitLab workflow для команд
- **День 4**: Protected branches, required reviews
- **День 6**: Continuous Integration настройка

## 🛠 Практические задания

### Go
1. **Unit тесты**:
   - Тесты для функций и методов
   - Table-driven tests для множественных случаев
   - Тестирование error cases

2. **Mocking и интеграционные тесты**:
   - Создание мок-объектов
   - Тестирование с внешними зависимостями
   - HTTP handler тесты

3. **Benchmark и производительность**:
   - Benchmark тесты для критичных функций
   - Profiling памяти и CPU
   - Оптимизация на основе тестов

### Git
1. **Team workflow**:
   - Feature branch workflow
   - Pull request process
   - Code review guidelines

2. **CI/CD**:
   - GitHub Actions для автоматического тестирования
   - Test coverage reporting
   - Automated deployment

## 📖 Ресурсы для изучения

### Видео
- [Testing in Go](https://www.youtube.com/watch?v=ttKIMijlO-c)
- [Git Team Workflows](https://www.youtube.com/watch?v=aJnFGMclhU8)

### Библиотеки
- [Testify](https://github.com/stretchr/testify)
- [GoMock](https://github.com/golang/mock)
- [HTTPTest](https://golang.org/pkg/net/http/httptest/)

## 🏠 Домашнее задание

1. **Go**: Полностью протестировать Task Manager из недели 5:
   - Unit тесты для всех функций и методов
   - Integration тесты для HTTP endpoints
   - Мок-объекты для database layer
   - Benchmark тесты для критичных операций
   - Достичь 80%+ code coverage

2. **Git**: 
   - Настроить CI/CD pipeline
   - Создать protected main branch
   - Настроить автоматический запуск тестов
   - Implement code coverage reporting

## ✅ Критерии успеха
- [ ] Понимание различных типов тестов
- [ ] Умение писать эффективные unit тесты
- [ ] Владение mocking техниками
- [ ] Настроен CI/CD pipeline
- [ ] Task Manager покрыт тестами на 80%+

## 🎯 Проект недели
**TDD разработка URL Shortener**:
1. Написать тесты перед кодом
2. Реализовать минимальный код для прохождения тестов
3. Рефакторинг с сохранением тестов
4. Повторить цикл для всех функций

## 📊 Метрики качества
- Code coverage > 80%
- Все тесты проходят
- Benchmark тесты показывают приемлемую производительность
- CI/CD pipeline работает без ошибок

**Время на неделю**: ~16-20 часов 