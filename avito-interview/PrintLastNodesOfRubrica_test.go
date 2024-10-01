package avito_interview

import "testing"

func Test_printTree(t *testing.T) {
	tests := []struct {
		name  string
		input Tree
	}{
		{
			name: "baas",
			input: Tree{
				Node{
					title: "Вещи",
					children: &Tree{
						Node{
							title: "Одежда",
							children: &Tree{
								Node{
									title: "Мужская",
								},
								Node{
									title: "Женская",
								},
							},
						},
					},
				},
				Node{
					title: "Хобби",
					children: &Tree{
						Node{
							title: "Велосипеды",
							children: &Tree{
								Node{
									title: "Горные",
								},
							},
						},
						Node{
							title: "Мангалы",
						},
					},
				},
				Node{
					title: "Транспорт",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printTree(tt.input)
		})
	}
}
