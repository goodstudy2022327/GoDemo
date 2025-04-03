package main

import "fmt"

type Point struct {
	x int
	y int
}

func scalePoint(p *Point, factor float64) {
	p.x = int(float64(p.x) * factor)
	p.y = int(float64(p.y) * factor)
}

func main() {
	point := Point{x: 10, y: 20}
	fmt.Println("缩放前：", point)

	scalePoint(&point, 2.5)
	fmt.Println("缩放后：", point)
}
