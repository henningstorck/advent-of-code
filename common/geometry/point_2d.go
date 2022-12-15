package geometry

import "math"

type Point2D struct {
	X int
	Y int
}

func NewPoint2D(x, y int) Point2D {
	return Point2D{x, y}
}

func (point Point2D) Add(otherPoint Point2D) Point2D {
	return point.AddXY(otherPoint.X, otherPoint.Y)
}

func (point Point2D) AddXY(x, y int) Point2D {
	return Point2D{
		X: point.X + x,
		Y: point.Y + y,
	}
}

func (point Point2D) GetManhattenDistance(otherPoint Point2D) int {
	diffX := int(math.Abs(float64(point.X) - float64(otherPoint.X)))
	diffY := int(math.Abs(float64(point.Y) - float64(otherPoint.Y)))
	return diffX + diffY
}
