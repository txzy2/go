# Go Математический Калькулятор

Проект представляет собой математический калькулятор на языке Go с поддержкой различных операций и статистических функций.

## 📋 Описание

Этот проект демонстрирует использование Go для создания математического калькулятора с модульной архитектурой. Проект включает в себя:

- **Калькулятор** с поддержкой базовых арифметических операций (сложение, вычитание, умножение, деление)
- **Статистические функции** (поиск максимума, вычисление степеней)
- **Генератор массивов** для создания тестовых данных
- **Система валидации** для проверки входных данных
- **Централизованная обработка ошибок** с предопределенными типами
- **Полное покрытие тестами** для всех компонентов
- **Docker поддержка** для контейнеризации

## 🏗️ Архитектура проекта

```
go/
├── cmd/first/           # Точка входа приложения
│   └── main.go
├── internal/            # Внутренние пакеты
│   ├── arr/            # Работа с массивами
│   │   ├── arr.go
│   │   └── arr_test.go
│   └── mathutil/       # Математические утилиты
│       ├── calculator.go
│       ├── math.go
│       ├── mathutil_test.go
│       ├── statistics.go
│       └── operations/ # Операции калькулятора
│           ├── interfaces.go
│           ├── sum.go
│           ├── minus.go
│           ├── multiply.go
│           ├── divide.go
│           └── operations_test.go
├── pkg/                 # Публичные пакеты
│   ├── types/          # Типы и ошибки
│   │   └── errors.go
│   └── validators/     # Валидация данных
│       └── validator.go
├── tmp/                # Временные файлы
│   └── main
├── Dockerfile          # Docker конфигурация
├── Makefile           # Команды сборки
└── go.mod            # Зависимости Go
```

## 🚀 Быстрый старт

### Предварительные требования

- Go 1.21.0 или выше
- Docker (опционально)

### Установка и запуск

1. **Клонирование репозитория:**
   ```bash
   git clone <repository-url>
   cd go
   ```

2. **Запуск приложения:**
   ```bash
   # Сборка и запуск
   make run
   
   # Или напрямую
   go run ./cmd/first
   ```

3. **Запуск тестов:**
   ```bash
   make test
   # или
   go test -v ./...
   ```

## 🐳 Docker

### Сборка и запуск в Docker

```bash
# Сборка Docker образа
make docker-build

# Запуск в Docker
make docker-run

# Комбинированная команда (сборка + запуск)
make dbr
```

## 📚 API и функции

### Калькулятор

Основной функционал калькулятора реализован в пакете `mathutil`:

```go
// Создание калькулятора
calculator, err := mathutil.NewCalculator(nums, mathutil.MINUS)

// Выполнение вычислений
result, err := calculator.Calculate()
```

**Поддерживаемые операции:**
- `SUM` (+) - сложение
- `MINUS` (-) - вычитание  
- `MULTIPLY` (*) - умножение
- `DIVIDE` (/) - деление

### Статистические функции

```go
// Поиск максимального значения
max, err := mathutil.Max(numbers)

// Вычисление квадратов и кубов
squares, cubes, err := mathutil.Pow(numbers)
```

### Работа с массивами

```go
// Генерация случайного массива
numbers := arr.Gen_arr(10)

// Создание объекта массива
arrObj := arr.NewArr(numbers)
length := arrObj.Get_len()
```

### Валидация данных

```go
import "lab/first/pkg/validators"

// Проверка длины массива
valid, err := validators.CheckLength(numbers)

// Проверка для бинарных операций (деление, вычитание)
valid, err := validators.CheckBinaryOperation(numbers)

// Проверка для унарных операций (сложение, умножение)
valid, err := validators.CheckUnaryOperation(numbers)
```

## 🧪 Тестирование

Проект включает полное покрытие тестами:

- **Unit тесты** для всех математических функций
- **Тесты калькулятора** с различными операциями
- **Тесты валидации** входных данных (бинарные и унарные операции)
- **Тесты обработки ошибок** с предопределенными типами
- **Тесты работы с массивами**

Запуск тестов:
```bash
go test -v ./...
```

## 🔧 Makefile команды

| Команда | Описание |
|---------|----------|
| `make test` | Запуск всех тестов |
| `make build` | Сборка приложения |
| `make run` | Сборка и запуск |
| `make docker-build` | Сборка Docker образа |
| `make docker-run` | Запуск в Docker |
| `make br` | Быстрая сборка и запуск |
| `make dbr` | Быстрая Docker сборка и запуск |

## 📦 Модули и зависимости

Проект использует Go модули. Основные зависимости указаны в `go.mod`:

```go
module lab/first
go 1.21.0
```

## 🏛️ Паттерны проектирования

Проект демонстрирует следующие паттерны:

- **Factory Pattern** - для создания калькуляторов операций
- **Strategy Pattern** - для различных математических операций
- **Interface Segregation** - четкое разделение интерфейсов
- **Error Handling** - централизованная обработка ошибок
- **Validation Pattern** - отделение логики валидации в отдельный пакет
- **Package Organization** - разделение на internal и pkg пакеты

## 🚨 Обработка ошибок

Проект включает централизованную систему обработки ошибок в пакете `pkg/types`:

- `ErrInvalidInput` - неверные входные данные (пустой массив)
- `ErrDivisionByZero` - деление на ноль
- `ErrUnknownOperator` - неизвестная операция

Все ошибки определены в одном месте и используются во всех модулях проекта для обеспечения консистентности.

## 📈 Производительность

Приложение измеряет время выполнения и выводит его в консоль:

```
Runtime: 123.456µs
```

## 🔍 Пример использования

```go
package main

import (
    "fmt"
    "lab/first/internal/arr"
    "lab/first/internal/mathutil"
    "lab/first/pkg/validators"
)

func main() {
    // Генерация массива чисел
    nums := arr.Gen_arr(5)
    
    // Валидация входных данных
    valid, err := validators.CheckBinaryOperation(nums)
    if err != nil {
        fmt.Printf("Ошибка валидации: %v\n", err)
        return
    }
    if !valid {
        fmt.Println("Недостаточно данных для операции")
        return
    }
    
    // Создание калькулятора для вычитания
    calculator, err := mathutil.NewCalculator(nums, mathutil.MINUS)
    if err != nil {
        fmt.Printf("Ошибка: %v\n", err)
        return
    }
    
    // Выполнение вычислений
    result, err := calculator.Calculate()
    if err != nil {
        fmt.Printf("Ошибка вычисления: %v\n", err)
    } else {
        fmt.Printf("Результат: %v\n", result)
    }
    
    // Поиск максимума
    max, _ := mathutil.Max(nums)
    fmt.Printf("Максимум: %v\n", max)
}
```

## 📝 Лицензия

Этот проект создан в образовательных целях.

## 🤝 Вклад в проект

1. Форкните репозиторий
2. Создайте ветку для новой функции (`git checkout -b feature/AmazingFeature`)
3. Зафиксируйте изменения (`git commit -m 'Add some AmazingFeature'`)
4. Отправьте в ветку (`git push origin feature/AmazingFeature`)
5. Откройте Pull Request
