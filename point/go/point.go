package point

import (
	"fmt"
	"math"
)

// Point with x and y
type Point struct {
	x, y int
}

func (p *Point) getX() int {
	return p.x
}

func (p *Point) getY() int {
	return p.y
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func makePoint(x, y int) *Point {
	return &Point{x, y}
}

func getQuadtant(p *Point) int {
	if p.x == 0 || p.y == 0 {
		return 0
	}

	if p.x > 0 {
		if p.y > 0 {
			return 1
		}
		return 4
	}

	if p.y > 0 {
		return 2
	}
	return 3

}

func getSymmetricalPoint(p *Point) *Point {
	return makePoint(-p.x, -p.y)
}

func calculateDistance(p1, p2 *Point) float64 {

	x1 := float64(p1.x)
	y1 := float64(p1.y)

	x2 := float64(p2.x)
	y2 := float64(p2.y)

	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// func main() {
// 	log.Println(getSymmetricalPoint(makePoint(10, 10)))
// 	log.Println(getSymmetricalPoint(makePoint(-10, 10)))
// 	log.Println(getSymmetricalPoint(makePoint(10, -10)))
// 	log.Println(getSymmetricalPoint(makePoint(-10, -10)))

// 	log.Println(calculateDistance(makePoint(-2, -3), makePoint(-4, 4)))
// }
