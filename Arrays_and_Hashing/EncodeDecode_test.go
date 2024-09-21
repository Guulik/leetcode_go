package Arrays_and_Hashing

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	tests := []struct {
		name    string
		strings []string
		encoded string
	}{
		{
			name:    "neetcode",
			strings: []string{"neet", "code", "love", "you"},
			encoded: "$4neet$4code$4love$3you",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := Encode(tt.strings)
			require.Equal(t, tt.encoded, encoded)
			decoded := Decode(encoded)
			require.Equal(t, tt.strings, decoded)
		})
	}
}
