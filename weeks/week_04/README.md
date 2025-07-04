# 📅 Неделя 4: Структуры и методы в Go, продвинутый Git

## 🎯 Цели недели
- Освоить создание и использование структур в Go
- Понять концепцию методов и receiver'ов
- Изучить продвинутые операции Git (stash, cherry-pick)
- Создать объектно-ориентированные программы

## 📚 Теория (Go 80% | Git 20%)

### Go (6 часов)
- **День 1-2**: Объявление структур, инициализация
- **День 3-4**: Методы, value vs pointer receivers
- **День 5-6**: Встраивание структур (embedding)
- **Выходные**: Создание системы с несколькими структурами

### Git (1.5 часа)
- **День 2**: Git stash для временного сохранения изменений
- **День 4**: Cherry-pick для переноса отдельных коммитов
- **День 6**: Interactive rebase для редактирования истории

## 🛠 Практические задания

### Go
1. **Базовые структуры**:
   - Структура Person с полями и методами
   - Структура Rectangle с методами площади и периметра
   - Структура BankAccount с методами депозита и снятия

2. **Методы и receiver'ы**:
   - Сравнение value и pointer receivers
   - Методы для модификации структур
   - Методы для вычислений без изменения состояния

3. **Композиция структур**:
   - Структура Employee встраивающая Person
   - Система управления библиотекой (Book, Author, Library)
   - Иерархия геометрических фигур

### Git
1. **Git stash**:
   - Сохранение незавершенной работы
   - Применение stash'а на другой ветке
   - Управление несколькими stash'ами

2. **Advanced operations**:
   - Cherry-pick коммитов между ветками
   - Interactive rebase для чистой истории
   - Amend для исправления последнего коммита

## 📖 Ресурсы для изучения

### Видео
- [Go Structs and Methods](https://www.youtube.com/watch?v=S5xvmNFKqz0)
- [Git Advanced Features](https://www.youtube.com/watch?v=ElRzTuYln0M)

### Практика
- [Go Struct Examples](https://gobyexample.com/structs)
- [Interactive Git Tutorial](https://ohmygit.org/)

## 🏠 Домашнее задание

1. **Go**: Создать систему управления студентами:
   - Структура Student (имя, ID, оценки)
   - Структура Course (название, преподаватель, студенты)
   - Методы для добавления оценок, вычисления среднего балла
   - Методы для поиска и фильтрации студентов

2. **Git**: 
   - Работать с feature branches для каждой структуры
   - Использовать stash при переключении между задачами
   - Создать чистую историю коммитов через rebase

## ✅ Критерии успеха
- [ ] Понимание создания и использования структур
- [ ] Четкое различие между value и pointer receivers
- [ ] Умение создавать композитные структуры
- [ ] Владение продвинутыми Git операциями
- [ ] Завершена система управления студентами

## 🎯 Мини-проект недели
**Система управления интернет-магазином**:
- Структуры: Product, Customer, Order, Cart
- Методы для добавления товаров в корзину
- Методы для расчета общей стоимости
- Система скидок и промокодов

## 🔍 Вопросы для самопроверки
1. Когда использовать pointer receiver, а когда value receiver?
2. Как работает embedding в Go?
3. Зачем нужен git stash?
4. Как безопасно переписать историю коммитов?

## 🛠 Дополнительные задачи
- Добавить JSON сериализацию для структур
- Создать интерфейс для общих методов
- Написать unit тесты для методов

**Время на неделю**: ~12-16 часов 