package avito_interview

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindRoute(t *testing.T) {
	tests := []struct {
		name    string
		tickets []Ticket
		want    []Ticket
	}{
		{
			name: "basa",
			tickets: []Ticket{
				{From: "London", To: "Moscow"},
				{From: "NY", To: "London"},
				{From: "Moscow", To: "SPb"},
			},
			want: []Ticket{
				{From: "NY", To: "London"},
				{From: "London", To: "Moscow"},
				{From: "Moscow", To: "SPb"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(FindRoute(tt.tickets))
			require.Equal(t, tt.want, FindRoute(tt.tickets))
		})
	}
}
