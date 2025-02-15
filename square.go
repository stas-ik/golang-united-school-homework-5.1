package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a uint
}

func (squareItem Square) End() Point {

	endPoint := Point{
		x: squareItem.start.x + int(squareItem.a),
		y: squareItem.start.y + int(squareItem.a),
	}

	return endPoint
}

func (squareItem Square) Area() uint {
	
	return squareItem.a * squareItem.a
}

func (squareItem Square) Perimeter() uint {

	return  squareItem.a * 4
}