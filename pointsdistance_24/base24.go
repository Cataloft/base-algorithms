package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type Points struct {
	x float64
	y float64
}

func NewPoints(x, y float64) *Points {
	return &Points{
		x: x,
		y: y,
	}
}

func pointsDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func main() {
	var x1, y1, x2, y2 float64
	fmt.Println("Введите координаты 1 точки: ")
	fmt.Scan(&x1, &y1)
	fmt.Println("Введите координаты 2 точки: ")
	fmt.Scan(&x2, &y2)

	pointsA := NewPoints(x1, y1)
	pointsB := NewPoints(x2, y2)

	fmt.Println("Расстояния между двумя точками: ", pointsDistance(pointsA.x, pointsA.y, pointsB.x, pointsB.y))
}
