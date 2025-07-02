# 📅 Неделя 10: Веб-разработка на Go и масштабирование Git

## 🎯 Цели недели
- Освоить создание веб-приложений на Go
- Изучить HTTP сервер, middleware, routing
- Освоить Git для больших проектов
- Создать полноценное веб-приложение

## 📚 Теория (Go 80% | Git 20%)

### Go (6 часов)
- **День 1-2**: HTTP сервер, handlers, ServeMux
- **День 3-4**: Middleware, роутинг, популярные фреймворки
- **День 5-6**: Templates, статические файлы, sessions
- **Выходные**: Создание полноценного веб-приложения

### Git (1.5 часа)
- **День 2**: Git в больших проектах, monorepo vs multirepo
- **День 4**: Git LFS для больших файлов
- **День 6**: Git performance оптимизация

## 🛠 Практические задания

### Go
1. **Базовый HTTP сервер**:
   - Простой HTTP сервер с разными endpoints
   - Обработка GET, POST, PUT, DELETE запросов
   - JSON API для CRUD операций

2. **Middleware и роутинг**:
   - Создание кастомных middleware
   - Аутентификация и авторизация
   - Логирование запросов и ответов

3. **Полноценное веб-приложение**:
   - HTML templates с Go
   - Статические файлы (CSS, JS, images)
   - Session management
   - Form validation

### Git
1. **Масштабирование**:
   - Работа с большими репозиториями
   - Git LFS для медиа файлов
   - Оптимизация производительности

2. **Advanced workflows**:
   - Monorepo strategies
   - Multi-project coordination
   - Release management

## 📖 Ресурсы для изучения

### Видео
- [Building Web Apps with Go](https://www.youtube.com/watch?v=Vlie-srOU8c)
- [Git for Large Projects](https://www.youtube.com/watch?v=0SJCYPsef54)

### Фреймворки
- [Gin](https://github.com/gin-gonic/gin)
- [Echo](https://github.com/labstack/echo)
- [Gorilla Mux](https://github.com/gorilla/mux)

## 🏠 Домашнее задание

1. **Go**: Создать блог-платформу:
   - Пользователи могут регистрироваться и логиниться
   - Создание, редактирование, удаление постов
   - Комментарии к постам
   - Административная панель
   - Responsive веб-интерфейс

2. **Git**: 
   - Организовать проект как monorepo
   - Настроить Git LFS для изображений
   - Создать комплексный release process

## ✅ Критерии успеха
- [ ] Понимание HTTP протокола и веб-разработки
- [ ] Умение создавать middleware
- [ ] Владение template engine
- [ ] Знание Git для больших проектов
- [ ] Завершена блог-платформа

## 🎯 Проект недели
**Блог-платформа "GoBlog"**:
- Аутентификация пользователей
- CRUD операции для постов
- Rich text editor
- Поиск по постам
- Admin dashboard
- API + веб-интерфейс

## 🏗 Архитектура проекта
```
goblog/
├── cmd/server/main.go
├── internal/
│   ├── handlers/
│   ├── models/
│   ├── middleware/
│   └── database/
├── web/
│   ├── templates/
│   └── static/
├── api/
└── tests/
```

## 📈 Итоги первых 10 недель
После этой недели ты будешь:
- Уверенно программировать на Go
- Создавать веб-приложения и API
- Использовать Git на профессиональном уровне
- Готов к изучению продвинутых тем

**Время на неделю**: ~18-22 часа 