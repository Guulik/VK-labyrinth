package input

import (
	pq "VK-contora/internal/priorityQueue"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_parsePoint_Happy(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		point1 pq.Point
		point2 pq.Point
	}{
		{
			name:   "example",
			input:  "0 0 2 1",
			point1: pq.Point{X: 0, Y: 0},
			point2: pq.Point{X: 2, Y: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start, end, err := parsePoints(tt.input)
			require.NoError(t, err)
			require.Equal(t, tt.point1, start)
			require.Equal(t, tt.point2, end)
		})
	}
}

func Test_parsePoint_Bad(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		point1  pq.Point
		point2  pq.Point
		wantErr error
	}{
		{
			name:    "minus",
			input:   "-1 -3 -3 -7",
			point1:  invalidPoint,
			point2:  invalidPoint,
			wantErr: negativeErr,
		},
		{
			name:    "alphanumerical",
			input:   "-1 -72j -3r -7q",
			point1:  invalidPoint,
			point2:  invalidPoint,
			wantErr: errorParse,
		},
		{
			name:    "expression",
			input:   "1 153-72 -3 -7",
			point1:  invalidPoint,
			point2:  invalidPoint,
			wantErr: errorParse,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start, end, err := parsePoints(tt.input)
			require.ErrorIs(t, err, tt.wantErr)
			require.Equal(t, tt.point1, start)
			require.Equal(t, tt.point2, end)
		})
	}
}
