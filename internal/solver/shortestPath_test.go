package solver

import (
	"VK-contora/internal/input"
	pq "VK-contora/internal/priorityQueue"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestShortestPath(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []pq.Point
		wantErr error
	}{
		{
			name: "example",
			input: `3 3
1 2 0
2 0 1
9 1 0
0 0 2 1
`,
			want: []pq.Point{
				{0, 0},
				{1, 0},
				{2, 0},
				{2, 1},
			},
			wantErr: nil,
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
			lab, start, end, err := input.Read()
			require.NoError(t, err)

			path, err := ShortestPath(lab, start, end)
			require.Equal(t, tt.want, path)
			require.NoError(t, err)
		})
	}
}
