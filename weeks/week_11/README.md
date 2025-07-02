# 📅 Неделя 11: Базы данных и ORM в Go

## 🎯 Цели недели
- Освоить работу с базами данных в Go
- Изучить SQL драйверы и ORM библиотеки
- Понять паттерны Repository и Active Record
- Создать приложение с полноценной базой данных

## 📚 Теория (Go 80% | Git 20%)

### Go (6 часов)
- **День 1-2**: database/sql package, SQL драйверы
- **День 3-4**: GORM ORM, миграции, связи между таблицами
- **День 5-6**: Connection pooling, транзакции, performance
- **Выходные**: Рефакторинг веб-приложений с настоящей БД

### Git (1.5 часа)
- **День 2**: Database migrations в Git workflow
- **День 4**: Environment-specific configurations
- **День 6**: Database seeding и fixtures

## 🛠 Практические задания

### Go
1. **Сырой SQL**:
   - Подключение к PostgreSQL/MySQL
   - CRUD операции с database/sql
   - Prepared statements и SQL injection защита

2. **ORM с GORM**:
   - Модели и автомиграции
   - Relationships (has many, belongs to, many to many)
   - Queries с условиями и joins

3. **Продвинутые техники**:
   - Транзакции и rollback
   - Connection pooling настройка
   - Database testing с testcontainers

### Git
1. **Database workflows**:
   - Versioning схемы базы данных
   - Migration files в Git
   - Environment configurations

## 📖 Ресурсы для изучения

### Библиотеки
- [GORM](https://gorm.io/)
- [Squirrel SQL Builder](https://github.com/Masterminds/squirrel)
- [Migrate](https://github.com/golang-migrate/migrate)

## 🏠 Домашнее задание

**Go**: Создать систему управления библиотекой:
- Модели: Book, Author, User, Loan
- CRUD операции через API
- Сложные запросы (поиск книг, история займов)
- Миграции и сиды для тестовых данных

## ✅ Критерии успеха
- [ ] Понимание работы с SQL в Go
- [ ] Уверенная работа с GORM
- [ ] Правильная организация миграций
- [ ] Завершена система библиотеки

**Время на неделю**: ~18-22 часа 