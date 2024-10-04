package avito_interview

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindVacationDay(t *testing.T) {

	tests := []struct {
		name             string
		daysWithMeetings []DayWithMeetings
		periodLength     int
		vacationLength   int
		want             []int
	}{
		{
			name: "g",
			daysWithMeetings: []DayWithMeetings{
				{Day: 3, Meetings: 1},
				{Day: 4, Meetings: 3},
				{Day: 14, Meetings: 3},
				{Day: 21, Meetings: 3},
				{Day: 28, Meetings: 1},
			},
			periodLength:   30,
			vacationLength: 7,
			want:           []int{5, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, FindVacationDay(tt.daysWithMeetings, tt.periodLength, tt.vacationLength))
		})
	}
}
