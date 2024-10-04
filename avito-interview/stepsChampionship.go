package avito_interview

type StepRecord struct {
	UserId int
	Steps  int
}

type ChampionResult struct {
	UserIds []int
	Steps   int
}

func FindChampions(statistics [][]StepRecord) ChampionResult {
	stepSums := make(map[int]int)
	userDays := make(map[int]int)

	for _, dailyStats := range statistics {
		uniqueUsers := make(map[int]bool)
		for _, record := range dailyStats {
			stepSums[record.UserId] += record.Steps
			uniqueUsers[record.UserId] = true
		}
		for userId := range uniqueUsers {
			userDays[userId]++
		}
	}

	var maxSteps int
	champions := []int{}
	for userId, totalSteps := range stepSums {
		if userDays[userId] == len(statistics) { // пользователь не пропустил ни одного дня
			if totalSteps > maxSteps {
				maxSteps = totalSteps
				champions = []int{userId}
			} else if totalSteps == maxSteps {
				champions = append(champions, userId)
			}
		}
	}

	return ChampionResult{
		UserIds: champions,
		Steps:   maxSteps,
	}
}
