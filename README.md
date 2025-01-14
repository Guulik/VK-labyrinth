# Тестовое задание в ВК на позицию Backend Developer

## Инструкция по запуску

1. Склонировать репозиторий через `git clone https://github.com/Guulik/VK-labyrinth.git`

2. Открыть терминал и перейти в папку с проектом.

3. Запустить проект `make run`

Команда `make test` запускает тесты. Тестовые варианты лабиринтов описаны в `/internal/solver/shortestPath_test.go`

Тесты для других модулей описаны в файлах `/internal/{pkg_name}/*_test.go`

# Описание алгоритма поиска кратчайшего пути
## Дейкстра - это база 
Алгоритм основан на алгоритме Дейкстры для нахождения кратчайшего пути в графе.
Лабиринт — это граф, где клетки являются вершинами, а переходы между соседними клетками
— рёбрами с весами, равными значениям клеток.

## Основные шаги:
### Инициализация:

- `cost` хранит минимальные стоимости до клеток. Начальная точка = 0, остальные = ∞ (2^30).
- `prev` сохраняет предыдущие клетки для восстановления пути.
- Приоритетная очередь используется для обработки клеток в порядке минимальной стоимости.

### Обработка клеток:
- Извлекается клетка с минимальной стоимостью.
- Для каждой соседней клетки рассчитывается новая стоимость.
- Если новая стоимость меньше текущей, она обновляется, а клетка добавляется в очередь.

### Завершение:
- Если финиш достигнут, путь восстанавливается в обратном порядке с использованием слайса `prev`.
- Если путь невозможен, возвращается ошибка "no path found".
## Роль приоритетной очереди
Приоритетная очередь (heap) используется для обработки клеток в порядке их минимальной стоимости.
Очередь позволяет быстро находить клетку с минимальной стоимостью
(операция извлечения занимает O(log n)), что значительно ускоряет алгоритм,
особенно в больших лабиринтах.