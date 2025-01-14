package input

import (
	pq "VK-contora/internal/priorityQueue"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestRead_Happy(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedLab   [][]int
		expectedStart pq.Point
		expectedEnd   pq.Point
	}{
		{
			name: "Valid",
			input: `3 3
1 2 0
2 0 1
9 1 0
0 0 2 1
`,
			expectedLab: [][]int{
				{1, 2, 0},
				{2, 0, 1},
				{9, 1, 0},
			},
			expectedStart: pq.Point{X: 0, Y: 0},
			expectedEnd:   pq.Point{X: 2, Y: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile, err := os.CreateTemp("", "test_input")
			require.NoError(t, err)
			defer os.Remove(tmpFile.Name())

			_, err = tmpFile.WriteString(tt.input)
			require.NoError(t, err)

			_, err = tmpFile.Seek(0, 0)
			require.NoError(t, err)

			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()

			os.Stdin = tmpFile

			lab, start, end, err := Read()

			require.NoError(t, err)
			require.Equal(t, tt.expectedLab, lab)
			require.Equal(t, tt.expectedStart, start)
			require.Equal(t, tt.expectedEnd, end)

		})
	}
}

func TestRead_Bad(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedErr string
	}{
		{
			name: "Missing labyrinth size",
			input: `
1 2 0
2 0 1
9 1 0
0 0 2 1
`,
			expectedErr: "invalid labyrinth size format",
		},
		{
			name:        "Empty",
			input:       ``,
			expectedErr: "empty",
		},
		{
			name: "Invalid labyrinth size format",
			input: `3 a
1 2 0
2 0 1
9 1 0
0 0 2 1
`,
			expectedErr: "invalid labyrinth size values",
		},
		{
			name: "Row with incorrect length",
			input: `3 3
1 2
2 0 1
9 1 0
0 0 2 1
`,
			expectedErr: "row 0 has incorrect length",
		},
		{
			name: "Invalid cell value",
			input: `3 3
1 2 0
2 0 1
9 x 0
0 0 2 1
`,
			expectedErr: "invalid cell value at (2, 1)",
		},
		{
			name: "Missing start point",
			input: `3 3
1 2 0
2 0 1
9 1 0
`,
			expectedErr: "missing start point",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile, err := os.CreateTemp("", "test_input")
			require.NoError(t, err)
			defer os.Remove(tmpFile.Name())

			_, err = tmpFile.WriteString(tt.input)
			require.NoError(t, err)

			_, err = tmpFile.Seek(0, 0)
			require.NoError(t, err)

			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()

			os.Stdin = tmpFile

			lab, start, end, err := Read()

			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedErr)

			require.Equal(t, [][]int(nil), lab)
			require.Equal(t, invalidPoint, start)
			require.Equal(t, invalidPoint, end)
		})
	}
}
