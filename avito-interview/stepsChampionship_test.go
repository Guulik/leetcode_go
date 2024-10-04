package avito_interview

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindChampions(t *testing.T) {

	tests := []struct {
		name       string
		statistics [][]StepRecord
		want       ChampionResult
	}{
		{
			name: "gds",
			statistics: [][]StepRecord{
				{{UserId: 1, Steps: 1000}, {UserId: 2, Steps: 1500}},
				{{UserId: 2, Steps: 1000}},
			},
			want: ChampionResult{
				UserIds: []int{2},
				Steps:   2500,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, FindChampions(tt.statistics))
		})
	}
}
