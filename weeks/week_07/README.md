# 📅 Неделя 7: Обработка ошибок и Git debugging

## 🎯 Цели недели
- Освоить правильную обработку ошибок в Go
- Изучить паттерны работы с ошибками
- Освоить Git bisect для поиска багов
- Создать надежный код с error handling

## 📚 Теория (Go 80% | Git 20%)

### Go (6 часов)
- **День 1-2**: Базовая обработка ошибок, error interface
- **День 3-4**: Создание кастомных ошибок, wrapping errors
- **День 5-6**: Паттерны error handling, defer/panic/recover
- **Выходные**: Добавление error handling в существующие проекты

### Git (1.5 часа)
- **День 2**: Git bisect для поиска багов
- **День 4**: Git blame и история изменений
- **День 6**: Git reflog для восстановления потерянных коммитов

## 🛠 Практические задания

### Go
1. **Базовая обработка ошибок**:
   - Функции возвращающие ошибки
   - Проверка и обработка ошибок
   - Логирование ошибок

2. **Кастомные ошибки**:
   - Создание собственных типов ошибок
   - Error wrapping с fmt.Errorf
   - Использование errors.Is и errors.As

3. **Panic и recover**:
   - Понимание когда использовать panic
   - Graceful recovery с recover
   - Defer для cleanup операций

### Git
1. **Debugging с Git**:
   - Использование git bisect для поиска регрессий
   - Git blame для анализа изменений
   - Исследование истории файлов

2. **Recovery операции**:
   - Восстановление удаленных коммитов с reflog
   - Отмена изменений разными способами
   - Rescue операции

## 📖 Ресурсы для изучения

### Видео
- [Error Handling in Go](https://www.youtube.com/watch?v=lsBF58Q-DnY)
- [Git Bisect Tutorial](https://www.youtube.com/watch?v=P3ZR_83-hR8)

### Статьи
- [Error Handling and Go](https://blog.golang.org/error-handling-and-go)
- [Git Bisect Documentation](https://git-scm.com/docs/git-bisect)

## 🏠 Домашнее задание

1. **Go**: Создать файловый сервис с надежным error handling:
   - Функции для работы с файлами
   - Кастомные типы ошибок для разных случаев
   - Graceful shutdown с defer
   - Логирование всех ошибок

2. **Git**: 
   - Практика git bisect на реальном проекте
   - Создание сценария с багом для поиска
   - Документирование процесса debugging

## ✅ Критерии успеха
- [ ] Понимание идиоматичной обработки ошибок
- [ ] Умение создавать кастомные ошибки
- [ ] Понимание panic/recover механизма
- [ ] Владение Git debugging техниками
- [ ] Завершен файловый сервис

## 🎯 Проект недели
**HTTP клиент с retry логикой**:
- Обработка network ошибок
- Exponential backoff для retry
- Кастомные ошибки для разных HTTP статусов
- Graceful timeout handling

**Время на неделю**: ~14-18 часов 