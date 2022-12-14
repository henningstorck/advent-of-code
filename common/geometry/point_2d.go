package geometry

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
