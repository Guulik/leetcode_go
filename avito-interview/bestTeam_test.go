package avito_interview

import (
	"fmt"
	"testing"
)

func TestFindBestTeam(t *testing.T) {

	tests := []struct {
		name      string
		engineers Engineers
	}{
		{
			name: "bas",
			engineers: Engineers{
				Backend:  []int{1, 2, 2, 3},
				Frontend: []int{1, 3},
				QA:       []int{3, 4, 4},
				Design:   []int{2, 3},
			},
		},
		{name: "dva",
			engineers: Engineers{
				Backend:  []int{5},
				Frontend: []int{3, 6, 7, 10},
				QA:       []int{3, 9, 11, 18},
				Design:   []int{20},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(FindBestTeam(tt.engineers))
		})
	}
}
