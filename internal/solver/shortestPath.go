package solver

import (
	pq "VK-contora/internal/priorityQueue"
	"container/heap"
	"errors"
)

var (
	noPathErr = errors.New("no path found")
)

func ShortestPath(labyrinth [][]int, start, end pq.Point) ([]pq.Point, error) {
	rows, cols := len(labyrinth), len(labyrinth[0])
	directions := []pq.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	cost := make([][]int, rows)
	for i := range cost {
		cost[i] = make([]int, cols)
		for j := range cost[i] {
			cost[i][j] = 1 << 30
		}
	}
	cost[start.X][start.Y] = 0

	prev := make(map[pq.Point]pq.Point)
	h := &pq.PriorityQueue{}
	heap.Init(h)
	heap.Push(h, pq.Node{Point: start, Cost: 0, Priority: 0})

	for h.Len() > 0 {
		current := heap.Pop(h).(pq.Node)
		if current.Point == end {
			break
		}

		for _, dir := range directions {
			neighbor := pq.Point{X: current.Point.X + dir.X, Y: current.Point.Y + dir.Y}
			if neighbor.X < 0 || neighbor.Y < 0 || neighbor.X >= rows || neighbor.Y >= cols {
				continue
			}
			if labyrinth[neighbor.X][neighbor.Y] == 0 {
				continue
			}
			newCost := current.Cost + labyrinth[neighbor.X][neighbor.Y]
			if newCost < cost[neighbor.X][neighbor.Y] {
				cost[neighbor.X][neighbor.Y] = newCost
				prev[neighbor] = current.Point
				heap.Push(h, pq.Node{Point: neighbor, Cost: newCost, Priority: newCost})
			}
		}
	}

	if cost[end.X][end.Y] == 1<<30 {
		return nil, noPathErr
	}

	var path []pq.Point
	for at := end; at != start; at = prev[at] {
		path = append([]pq.Point{at}, path...)
	}
	path = append([]pq.Point{start}, path...)
	return path, nil
}
