package avito_interview

import "fmt"

type Dependencies map[string][]string

func PrintDeps(deps Dependencies) {
	visited := make(map[string]bool)

	var visit func(lib string)
	visit = func(lib string) {
		if visited[lib] {
			return
		}
		visited[lib] = true
		// Рекурсивно посещаем все зависимости
		for _, dep := range deps[lib] {
			visit(dep)
		}
		// Печатаем библиотеку после всех зависимостей
		fmt.Println(lib)
	}

	// Обходим все библиотеки, чтобы обработать их зависимости
	for lib := range deps {
		visit(lib)
	}
}
