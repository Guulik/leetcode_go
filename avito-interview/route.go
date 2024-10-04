package avito_interview

type Ticket struct {
	From string
	To   string
}

// FindRoute получает список билетов и возвращает маршрут в порядке следования
func FindRoute(tickets []Ticket) []Ticket {
	// Создаем карту для сопоставления отправных точек с пунктами назначения
	ticketMap := make(map[string]string)
	var start string

	// Заполняем карту и определяем стартовую точку
	for _, ticket := range tickets {
		ticketMap[ticket.From] = ticket.To
		start = ticket.From
	}

	// Находим стартовую точку
	for _, ticket := range tickets {
		if _, exists := ticketMap[ticket.To]; !exists {
			start = ticket.From
			break
		}
	}

	// Формируем маршрут
	var route []Ticket
	for current := start; current != ""; {
		next, exists := ticketMap[current]
		if !exists {
			break
		}
		route = append(route, Ticket{From: current, To: next})
		current = next
	}

	return route
}
