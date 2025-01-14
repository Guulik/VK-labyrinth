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
		},
		{
			name: "medium 4x4",
			input: `4 4
1 2 3 1
0 0 1 1
1 1 9 1
1 1 1 1
0 0 3 3
`,
			want: []pq.Point{
				{0, 0},
				{0, 1},
				{0, 2},
				{1, 2},
				{1, 3},
				{2, 3},
				{3, 3},
			},
		},
		{
			name: "no way!!!",
			input: `4 4
0 0 1 1
0 0 1 1
0 0 1 1
1 1 1 1
0 0 3 3
`,
			want:    nil,
			wantErr: noPathErr,
		},
		{
			name: "difficult",
			input: `3 4
1 2 0 1
0 3 1 3
1 0 4 5
0 0 2 3
`,
			want: []pq.Point{
				{0, 0},
				{0, 1},
				{1, 1},
				{1, 2},
				{1, 3},
				{2, 3},
			},
			wantErr: nil,
		},
		{
			name: "one cell",
			input: `1 1
1
0 0 0 0
`,
			want: []pq.Point{
				{0, 0},
			},
		}, {
			name: "one row",
			input: `1 5
1 7 3 2 9
0 4 0 0
`,
			want: []pq.Point{
				{0, 4},
				{0, 3},
				{0, 2},
				{0, 1},
				{0, 0},
			},
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
			require.ErrorIs(t, tt.wantErr, err)
		})
	}
}
