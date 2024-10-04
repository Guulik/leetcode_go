package avito_interview

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCompareVersions(t *testing.T) {

	tests := []struct {
		name string
		v1   string
		v2   string
		want int
	}{
		{
			name: "-1",
			v1:   "v11.22.44",
			v2:   "v11.22.45",
			want: -1,
		}, {
			name: "0",
			v1:   "v11.22.44",
			v2:   "v11.22.44",
			want: 0,
		}, {
			name: "0",
			v1:   "v11.22.44",
			v2:   "v11.22.44.0",
			want: 0,
		}, {
			name: "1",
			v1:   "v1.12.3",
			v2:   "v1.3.4",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, CompareVersions(tt.v1, tt.v2))
		})
	}
}
