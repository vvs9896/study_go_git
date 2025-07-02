# 📅 Неделя 8: Горутины, каналы и Git submodules

## 🎯 Цели недели
- Освоить параллельное программирование с горутинами
- Понять каналы для коммуникации между горутинами
- Изучить Git submodules для управления зависимостями
- Создать concurrent приложения

## 📚 Теория (Go 80% | Git 20%)

### Go (6 часов)
- **День 1-2**: Горутины, ключевое слово go
- **День 3-4**: Каналы (channels), buffered vs unbuffered
- **День 5-6**: Select statement, worker pools
- **Выходные**: Создание concurrent веб-сервера

### Git (1.5 часа)
- **День 2**: Git submodules - добавление и управление
- **День 4**: Обновление submodules, работа с вложенными проектами
- **День 6**: Альтернативы submodules (subtree, vendor)

## 🛠 Практические задания

### Go
1. **Базовые горутины**:
   - Простые горутины с различными задачами
   - WaitGroup для синхронизации
   - Передача данных через каналы

2. **Продвинутые паттерны**:
   - Worker pool pattern
   - Fan-in/Fan-out паттерны
   - Pipeline паттерн с каналами

3. **Практические применения**:
   - Concurrent HTTP requests
   - Параллельная обработка файлов
   - Rate limiting с каналами

### Git
1. **Submodules**:
   - Добавление внешних репозиториев как submodules
   - Клонирование проектов с submodules
   - Обновление и синхронизация

2. **Dependency management**:
   - Сравнение подходов управления зависимостями
   - Vendoring vs submodules vs modules
   - Best practices

## 📖 Ресурсы для изучения

### Видео
- [Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)
- [Git Submodules Explained](https://www.youtube.com/watch?v=gSlXo2iLBro)

### Документация
- [Go Concurrency](https://golang.org/doc/effective_go.html#concurrency)
- [Git Submodules](https://git-scm.com/book/en/v2/Git-Tools-Submodules)

## 🏠 Домашнее задание

1. **Go**: Создать систему мониторинга веб-сайтов:
   - Горутины для проверки доступности сайтов
   - Каналы для сбора результатов
   - Worker pool для ограничения параллельности
   - Periodic reporting через каналы

2. **Git**: 
   - Организовать проект с использованием submodules
   - Создать shared библиотеку как submodule
   - Настроить автоматическое обновление

## ✅ Критерии успеха
- [ ] Понимание горутин и их lifecycle
- [ ] Уверенная работа с каналами
- [ ] Знание основных concurrency паттернов
- [ ] Умение работать с Git submodules
- [ ] Завершена система мониторинга

## 🎯 Проект недели
**Concurrent Web Scraper**:
- Worker pool для scraping
- Rate limiting для уважения к серверам
- Results aggregation через каналы
- Graceful shutdown с context

## 🔍 Важные концепции
- "Don't communicate by sharing memory; share memory by communicating"
- Channel ownership и closing patterns
- Context для cancellation
- Race conditions и их избежание

**Время на неделю**: ~16-20 часов 