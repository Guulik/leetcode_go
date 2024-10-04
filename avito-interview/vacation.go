package avito_interview

type DayWithMeetings struct {
	Day      int
	Meetings int
}

// FindVacationDay находит оптимальный день начала отпуска и количество пропущенных встреч
func FindVacationDay(daysWithMeetings []DayWithMeetings, periodLength int, vacationLength int) []int {
	// Создаем массив для хранения количества встреч
	meetings := make([]int, periodLength+1)

	// Заполняем массив meetings
	for _, entry := range daysWithMeetings {
		meetings[entry.Day] = entry.Meetings
	}

	// Префиксный массив для быстрого подсчета встреч
	prefixSum := make([]int, periodLength+1)
	for i := 1; i <= periodLength; i++ {
		prefixSum[i] = prefixSum[i-1] + meetings[i]
	}

	minMeetings := int(^uint(0) >> 1) // максимальное значение для int
	startDay := 0

	// Ищем оптимальный день для начала отпуска
	for i := 1; i <= periodLength-vacationLength+1; i++ {
		// Подсчитываем встречи за период с i по i + vacationLength - 1
		currentMeetings := prefixSum[i+vacationLength-1] - prefixSum[i-1]

		// Обновляем минимальное количество встреч
		if currentMeetings < minMeetings {
			minMeetings = currentMeetings
			startDay = i
		}
	}

	return []int{startDay, minMeetings}
}
