package SlidingWindow

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindLongestSbstrWoutRep(t *testing.T) {
	tests := []struct {
		name string
		str  string
		leng int
	}{
		{
			name: "abcaba",
			str:  "abcabcbb",
			leng: 3,
		},
		{
			name: "bububu",
			str:  "bbbb",
			leng: 1,
		},
		{
			name: "pew",
			str:  "pwwkew",
			leng: 3,
		},
		{
			name: "аубля",
			str:  "au",
			leng: 2,
		},
		{
			name: "аа",
			str:  "aa",
			leng: 1,
		},
		{
			name: "pusto",
			str:  "",
			leng: 0,
		},
		{
			name: "probel",
			str:  " ",
			leng: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			longest := FindLongestSbstrWoutRep(tt.str)
			require.Equal(t, tt.leng, longest)
		})
	}
}
