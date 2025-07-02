# 🌿 Основы Git - Неделя 1

## 🎯 Что такое Git?

Git - это распределенная система контроля версий (VCS), созданная Линусом Торвальдсом в 2005 году. Git позволяет отслеживать изменения в файлах и координировать работу между несколькими разработчиками.

### Ключевые преимущества Git:
- **Распределенность**: Каждый разработчик имеет полную копию истории
- **Скорость**: Большинство операций выполняются локально
- **Целостность данных**: Все изменения проверяются контрольными суммами
- **Ветвление**: Легкое создание и слияние веток
- **Отслеживание истории**: Полная история всех изменений

## 🛠 Установка и настройка Git

### Установка Git на Linux:

```bash
# Ubuntu/Debian
sudo apt update
sudo apt install git

# Проверка установки
git --version
```

### Первоначальная настройка:

```bash
# Настройка имени пользователя
git config --global user.name "Ваше Имя"

# Настройка email
git config --global user.email "your.email@example.com"

# Настройка редактора по умолчанию
git config --global core.editor "code --wait"

# Проверка настроек
git config --list
```

### Полезные настройки:

```bash
# Включение цветного вывода
git config --global color.ui auto

# Настройка алиасов для частых команд
git config --global alias.st status
git config --global alias.co checkout
git config --global alias.br branch
git config --global alias.cm commit

# Настройка автокоррекции
git config --global help.autocorrect 1
```

## 📁 Основные концепции Git

### Рабочая область (Working Directory)
- Папка с вашими файлами
- Здесь вы редактируете код

### Индекс (Staging Area)
- Промежуточная область
- Подготовка файлов к коммиту

### Репозиторий (Repository)
- Место хранения истории изменений
- Содержит все коммиты

```
Working Directory → Staging Area → Repository
    (git add)       (git commit)
```

## 🚀 Основные команды Git

### Инициализация репозитория:

```bash
# Создание нового репозитория
git init

# Клонирование существующего репозитория
git clone https://github.com/user/repo.git
```

### Отслеживание изменений:

```bash
# Проверка статуса файлов
git status

# Добавление файлов в индекс
git add filename.txt          # Конкретный файл
git add .                     # Все файлы в текущей папке
git add *.go                  # Все Go файлы

# Удаление файлов из индекса
git reset filename.txt

# Просмотр различий
git diff                      # Изменения в рабочей области
git diff --staged             # Изменения в индексе
```

### Создание коммитов:

```bash
# Создание коммита
git commit -m "Описание изменений"

# Коммит с добавлением всех измененных файлов
git commit -am "Быстрый коммит"

# Изменение последнего коммита
git commit --amend -m "Новое описание"
```

### Просмотр истории:

```bash
# Просмотр истории коммитов
git log

# Краткий формат
git log --oneline

# Графическое представление
git log --graph --oneline --all

# История конкретного файла
git log filename.txt
```

## 📝 Создание .gitignore

Файл `.gitignore` указывает Git, какие файлы не нужно отслеживать:

```bash
# Создание .gitignore для Go проекта
touch .gitignore
```

Содержимое `.gitignore` для Go:

```gitignore
# Бинарные файлы
*.exe
*.exe~
*.dll
*.so
*.dylib

# Тестовые файлы
*.test

# Файлы coverage
*.out

# Зависимости
vendor/

# Go workspace file
go.work

# IDE файлы
.vscode/
.idea/
*.swp
*.swo

# OS файлы
.DS_Store
Thumbs.db

# Логи
*.log

# Временные файлы
*.tmp
```

## 🌐 Работа с удаленными репозиториями

### Настройка удаленного репозитория:

```bash
# Добавление удаленного репозитория
git remote add origin https://github.com/username/repo.git

# Просмотр удаленных репозиториев
git remote -v

# Изменение URL удаленного репозитория
git remote set-url origin https://github.com/username/new-repo.git
```

### Отправка изменений:

```bash
# Отправка в удаленный репозиторий
git push origin main

# Первая отправка с установкой upstream
git push -u origin main
```

### Получение изменений:

```bash
# Получение изменений без слияния
git fetch origin

# Получение и слияние изменений
git pull origin main
```

## 🔑 Настройка SSH ключей

### Генерация SSH ключа:

```bash
# Генерация нового SSH ключа
ssh-keygen -t ed25519 -C "your.email@example.com"

# Добавление ключа в ssh-agent
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_ed25519

# Копирование публичного ключа
cat ~/.ssh/id_ed25519.pub
```

### Добавление ключа на GitHub:
1. Идите в Settings → SSH and GPG keys
2. Нажмите "New SSH key"
3. Вставьте содержимое файла `id_ed25519.pub`
4. Сохраните ключ

### Проверка соединения:

```bash
# Тест соединения с GitHub
ssh -T git@github.com
```

## 📊 Понимание состояний файлов

```
Untracked → Staged → Committed
    ↓         ↓         ↓
(git add) (git commit) (git push)
```

### Состояния файлов:
- **Untracked**: Новые файлы, не отслеживаемые Git
- **Modified**: Измененные отслеживаемые файлы
- **Staged**: Файлы, готовые к коммиту
- **Committed**: Файлы, сохраненные в репозиторий

## 🛡 Лучшие практики

### Правила хороших коммитов:
1. **Частые коммиты**: Делайте коммиты часто и логически
2. **Описательные сообщения**: Ясно объясняйте, что изменилось
3. **Один логический блок**: Один коммит = одно изменение
4. **Формат сообщений**:
   ```
   Краткое описание (до 50 символов)
   
   Подробное описание изменений, если необходимо.
   Объясните, что и почему, а не как.
   ```

### Структура сообщений коммитов:
```bash
# Хорошие примеры:
git commit -m "Add user authentication feature"
git commit -m "Fix memory leak in data processor"
git commit -m "Update README with installation instructions"

# Плохие примеры:
git commit -m "fix"
git commit -m "changes"
git commit -m "work in progress"
```

### Использование префиксов:
```bash
feat: новая функциональность
fix: исправление ошибки
docs: изменения документации
style: форматирование кода
refactor: рефакторинг без изменения функциональности
test: добавление тестов
chore: изменения инфраструктуры
```

## 🔍 Полезные команды для начинающих

```bash
# Отмена изменений в рабочей области
git checkout -- filename.txt

# Отмена всех изменений в рабочей области
git checkout -- .

# Удаление неотслеживаемых файлов
git clean -f

# Просмотр содержимого коммита
git show commit-hash

# Поиск в истории коммитов
git log --grep="search term"

# Просмотр изменений между коммитами
git diff commit1 commit2
```

## 🎯 Основные файлы Git

### .git/config
Локальные настройки репозитория

### .gitignore
Список игнорируемых файлов

### README.md
Документация проекта (рекомендуется)

## 📚 Полезные ресурсы

- [Git Book](https://git-scm.com/book/ru/v2) - полное руководство по Git
- [GitHub Guides](https://guides.github.com/) - руководства по GitHub
- [Git Cheat Sheet](https://education.github.com/git-cheat-sheet-education.pdf) - шпаргалка по командам
- [Learn Git Branching](https://learngitbranching.js.org/) - интерактивное обучение

## ✅ Чек-лист для самопроверки

- [ ] Git установлен и настроен
- [ ] Настроены имя пользователя и email
- [ ] Создан SSH ключ и добавлен на GitHub
- [ ] Понимаю основные концепции Git
- [ ] Умею создавать репозиторий
- [ ] Знаю основные команды: add, commit, push, pull
- [ ] Умею создавать .gitignore файл
- [ ] Понимаю состояния файлов в Git
- [ ] Знаю правила хороших коммитов

## 🚀 Практическое задание

Создайте свой первый Git репозиторий:

1. Создайте папку для проекта
2. Инициализируйте Git репозиторий
3. Создайте файл README.md
4. Добавьте .gitignore для Go
5. Сделайте первый коммит
6. Создайте репозиторий на GitHub
7. Подключите локальный репозиторий к GitHub
8. Отправьте код на GitHub

```bash
mkdir my-first-repo
cd my-first-repo
git init
echo "# My First Repository" > README.md
# Создайте .gitignore файл
git add .
git commit -m "Initial commit: Add README and gitignore"
git remote add origin git@github.com:username/my-first-repo.git
git push -u origin main
``` 