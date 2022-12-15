package geometry_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/common/geometry"
	"github.com/stretchr/testify/assert"
)

func TestPoint2DAdd(t *testing.T) {
	point1 := geometry.Point2D{2, 6}
	point2 := geometry.Point2D{-3, 5}
	result := point1.Add(point2)
	assert.Equal(t, geometry.Point2D{-1, 11}, result)
}

func TestPoint2DAddXY(t *testing.T) {
	point := geometry.Point2D{2, 6}
	result := point.AddXY(-3, 5)
	assert.Equal(t, geometry.Point2D{-1, 11}, result)
}

func TestGetManhattenDistance(t *testing.T) {
	point1 := geometry.Point2D{2, 6}
	point2 := geometry.Point2D{-3, 5}
	result := point1.GetManhattenDistance(point2)
	assert.Equal(t, 6, result)
}
