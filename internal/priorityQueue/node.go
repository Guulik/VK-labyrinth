package priorityQueue

type Point struct {
	X int
	Y int
}

type Node struct {
	Point    Point
	Cost     int
	Priority int
}
