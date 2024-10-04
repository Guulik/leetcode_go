package avito_interview

import (
	"strconv"
	"strings"
)

func CompareVersions(v1, v2 string) int {
	// Убираем 'v' из начала строк
	v1 = strings.TrimPrefix(v1, "v")
	v2 = strings.TrimPrefix(v2, "v")

	// Разделяем строки на части по точке
	parts1 := strings.Split(v1, ".")
	parts2 := strings.Split(v2, ".")

	// Определяем максимальную длину частей
	maxLength := max(len(parts1), len(parts2))

	// Сравниваем каждую часть
	for i := 0; i < maxLength; i++ {
		// Получаем число из первой версии, если доступно
		num1 := 0
		if i < len(parts1) {
			num1, _ = strconv.Atoi(parts1[i])
		}

		// Получаем число из второй версии, если доступно
		num2 := 0
		if i < len(parts2) {
			num2, _ = strconv.Atoi(parts2[i])
		}

		// Сравниваем части
		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}
	}

	// Если все части равны
	return 0
}
