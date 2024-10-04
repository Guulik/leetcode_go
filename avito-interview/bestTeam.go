package avito_interview

import "math"

type Engineers struct {
	Backend  []int
	Frontend []int
	QA       []int
	Design   []int
}

// Функция для поиска команды с минимальной разницей уровней
func FindBestTeam(engineers Engineers) map[string]int {
	// Индексы для каждого типа инженеров
	i, j, k, l := 0, 0, 0, 0

	// Переменные для хранения текущей команды и минимальной разницы
	bestTeam := make(map[string]int)
	minDiff := math.MaxInt64

	for {

		if i >= len(engineers.Backend) || j >= len(engineers.Frontend) || k >= len(engineers.QA) || l >= len(engineers.Design) {
			break
		}
		// Получаем текущие уровни
		b, f := engineers.Backend[i], engineers.Frontend[j]
		q, d := engineers.QA[k], engineers.Design[l]

		// Определяем минимальный и максимальный уровни
		minLevel := min(b, f, q, d)
		maxLevel := max(b, f, q, d)
		diff := maxLevel - minLevel

		// Обновляем минимальную разницу и лучшую команду, если текущая разница меньше
		if diff < minDiff {
			minDiff = diff
			bestTeam["backend"] = b
			bestTeam["frontend"] = f
			bestTeam["qa"] = q
			bestTeam["design"] = d
		}

		// Если достигли конца любого из массивов, выходим из цикла
		if i >= len(engineers.Backend) || j >= len(engineers.Frontend) || k >= len(engineers.QA) || l >= len(engineers.Design) {
			break
		}

		// Увеличиваем индекс текущего наименьшего уровня
		if minLevel == b {
			i++
		} else if minLevel == f {
			j++
		} else if minLevel == q {
			k++
		} else {
			l++
		}
	}

	return bestTeam
}
