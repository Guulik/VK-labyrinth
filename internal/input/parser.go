package input

import (
	pq "VK-contora/internal/priorityQueue"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	negativeErr = errors.New("one of labyrinth conditions are negative")
	errorParse  = errors.New("failed to parse coord #")
)

func parsePoints(input string) (pq.Point, pq.Point, error) {
	var (
		x, y int
		err  error
	)

	coords := strings.Split(input, " ")
	x, err = strconv.Atoi(coords[0])
	if err != nil {
		return invalidPoint, invalidPoint, errors.Join(errorParse, fmt.Errorf("0: %w", err))
	}
	y, err = strconv.Atoi(coords[1])
	if err != nil {
		return invalidPoint, invalidPoint, errors.Join(errorParse, fmt.Errorf("1: %w", err))
	}
	if x < 0 || y < 0 {
		return invalidPoint, invalidPoint, negativeErr
	}
	start := pq.Point{X: x, Y: y}

	x, err = strconv.Atoi(coords[2])
	if err != nil {
		return invalidPoint, invalidPoint, errors.Join(errorParse, fmt.Errorf("2: %w", err))
	}
	y, err = strconv.Atoi(coords[3])
	if err != nil {
		return invalidPoint, invalidPoint, errors.Join(errorParse, fmt.Errorf("3: %w", err))
	}
	if x < 0 || y < 0 {
		return invalidPoint, invalidPoint, negativeErr
	}
	end := pq.Point{X: x, Y: y}

	return start, end, nil
}
