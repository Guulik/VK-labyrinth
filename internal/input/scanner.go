package input

import (
	pq "VK-contora/internal/priorityQueue"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	invalidPoint = pq.Point{X: -1, Y: -1}
)

func Read() ([][]int, pq.Point, pq.Point, error) {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return nil, invalidPoint, invalidPoint, fmt.Errorf("empty")
	}
	size := strings.Split(scanner.Text(), " ")
	if len(size) != 2 {
		return nil, invalidPoint, invalidPoint, fmt.Errorf("invalid labyrinth size format")
	}
	rows, err1 := strconv.Atoi(size[0])
	cols, err2 := strconv.Atoi(size[1])
	if err1 != nil || err2 != nil {
		return nil, invalidPoint, invalidPoint, fmt.Errorf("invalid labyrinth size values")
	}

	labyrinth := make([][]int, rows)
	for i := 0; i < rows; i++ {
		if !scanner.Scan() {
			return nil, invalidPoint, invalidPoint, fmt.Errorf("missing labyrinth rows")
		}
		row := strings.Split(scanner.Text(), " ")
		if len(row) != cols {
			return nil, invalidPoint, invalidPoint, fmt.Errorf("row %d has incorrect length", i)
		}
		labyrinth[i] = make([]int, cols)
		for j, val := range row {
			cell, err := strconv.Atoi(val)
			if err != nil || cell < 0 || cell > 9 {
				return nil, invalidPoint, invalidPoint, fmt.Errorf("invalid cell value at (%d, %d)", i, j)
			}
			labyrinth[i][j] = cell
		}
	}

	if !scanner.Scan() {
		return nil, invalidPoint, invalidPoint, fmt.Errorf("missing start point")
	}
	start, end, err := parsePoints(scanner.Text())
	if err != nil {
		return nil, invalidPoint, invalidPoint, err
	}

	return labyrinth, start, end, nil
}
