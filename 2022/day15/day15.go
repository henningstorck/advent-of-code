package day15

import (
	"fmt"
	"math"

	"github.com/henningstorck/advent-of-code/common/functional"
	"github.com/henningstorck/advent-of-code/common/geometry"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/henningstorck/advent-of-code/common/set"
)

type Sensor struct {
	SensorPos geometry.Point2D
	BeaconPos geometry.Point2D
	Dist      int
}

func Solve(reader *inputreader.InputReader, filename string, row int) (int, int) {
	lines := reader.ReadLines(2022, 15, filename)
	sensors := parse(lines)
	return solvePart1(sensors, row), solvePart2(sensors, row*2)
}

func parse(lines []string) []Sensor {
	return functional.Map(lines, func(line string) Sensor {
		var sensorPos geometry.Point2D
		var beaconPos geometry.Point2D
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorPos.X, &sensorPos.Y, &beaconPos.X, &beaconPos.Y)
		dist := sensorPos.GetManhattenDistance(beaconPos)
		return Sensor{sensorPos, beaconPos, dist}
	})
}

func solvePart1(sensors []Sensor, row int) int {
	noBeacons := set.Set[geometry.Point2D]{}
	beacons := set.Set[geometry.Point2D]{}

	for _, sensor := range sensors {
		distRow := int(math.Abs(float64(sensor.SensorPos.Y) - float64(row)))

		for i := 0; i <= sensor.Dist-distRow; i++ {
			noBeacons.Add(geometry.NewPoint2D(sensor.SensorPos.X-i, row))
			noBeacons.Add(geometry.NewPoint2D(sensor.SensorPos.X+i, row))
		}

		if sensor.BeaconPos.Y == row {
			beacons.Add(sensor.BeaconPos)
		}
	}

	return len(noBeacons) - len(beacons)
}

func solvePart2(sensors []Sensor, searchSpace int) int {
	for _, sensor := range sensors {
		min := -sensor.Dist - 1
		max := sensor.Dist + 1

		for i := min; i <= max; i++ {
			y := sensor.SensorPos.Y + i

			xCoords := []int{
				sensor.SensorPos.X - int(math.Abs(float64(sensor.Dist+1-i))),
				sensor.SensorPos.X + int(math.Abs(float64(sensor.Dist+1-i))),
			}

			for _, x := range xCoords {
				if x < 0 || x > searchSpace || y < 0 || y > searchSpace {
					break
				}

				match := true

				for _, otherSensor := range sensors {
					pos := geometry.NewPoint2D(x, y)
					distManhattan := otherSensor.SensorPos.GetManhattenDistance(pos)

					if distManhattan <= otherSensor.Dist {
						match = false
						break
					}
				}

				if match {
					return x*4000000 + y
				}
			}
		}
	}

	return 0
}
