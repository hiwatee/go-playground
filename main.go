package main

import "log"

func main() {
	n := 3
	log.Println(CreateDividedGeo(0, 100, 0, 100, n))
}

// ------ 長方形座標をNxNで分割する -------

type Range struct {
	Start float64
	End   float64
}

type Box struct {
	StartX  float64
	EndX    float64
	StartY  float64
	EndY    float64
	CenterX float64
	CenterY float64
}

type Boxes [][]Box

func CreateDividedGeo(startX float64, endX float64, startY float64, endY float64, number int) Boxes {
	arrX := CreateMidRangeDistance(startX, endX, number)
	arrY := CreateMidRangeDistance(startY, endY, number)
	return CreateSquare(arrX, arrY)
}

func CreateSquare(axisX []Range, axisY []Range) Boxes {
	n := len(axisX)

	graph := make([][]Box, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]Box, n)
	}

	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			graph[j][i] = Box{
				StartX:  axisX[i].Start,
				EndX:    axisX[i].End,
				StartY:  axisY[j].Start,
				EndY:    axisY[j].End,
				CenterX: (axisX[i].Start + axisX[i].End) / 2,
				CenterY: (axisY[j].Start + axisY[j].End) / 2,
			}
		}
	}

	return graph
}

func CreateMidRangeDistance(start float64, end float64, number int) []Range {
	n := number + 1
	array := make([]Range, (n - 1))
	unit := (end - start) / float64((n - 1))
	for i := 0; i < (n - 1); i++ {
		array[i] = Range{
			Start: start + unit*float64(i),
			End:   start + unit*float64((i+1)),
		}
	}

	return array
}
