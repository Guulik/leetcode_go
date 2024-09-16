package TwoPointers

func Trap(height []int) int {

	//идея: если с какой-то стороны стенка меньше, то будем двигаться по этой стороне.
	//каждый раз проверяем на превышение макимума стенки. Если текущий уровень ниже, чем максимальная стенка с рассматриваемой стороны,
	//то записываем эту разницу в площадь.
	// мы рассматриваем по-клеточно, потому что мы знаем, что с другой стороны стенка точно выше и она удержит нашу воду.
	if len(height) == 0 {
		return 0
	}

	L, R := 0, len(height)-1
	maxL, maxR := height[L], height[R]
	area := 0
	for L < R {
		if maxL <= maxR {
			L++
			maxL = max(maxL, height[L])
			area += maxL - height[L]
		} else {
			R--
			maxR = max(maxR, height[R])
			area += maxR - height[R]
		}
	}
	return area
}
