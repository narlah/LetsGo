package main

import (
	"fmt"
)

type point struct {
	x int
	y int
}

func buriedTreasure(nutMap [][]int, startPosition []int, maxSteps int) int {
	startX := startPosition[0]
	startY := startPosition[1]
	maxX := len(nutMap) - 1
	maxY := len(nutMap[0]) - 1
	if startX > maxX || startY > maxY {
		return 0
	}
	topMax := 0
	var path = make([]point, maxSteps+1)

	var fu func(curX int, curY int, stepsLeft int, currMAx int)
	fu = func(curX int, curY int, stepsLeft int, currMAx int) {
		if curX > maxX || curY > maxY || curX < 0 || curY < 0 {
			return
		}
		flag := false
		for _, p := range path[:maxSteps-stepsLeft] {
			if p.x == curX && p.y == curY {
				flag = true
				break
			}
		}
		if !flag {
			currMAx += nutMap[curX][curY]
		}
		path[maxSteps-stepsLeft].x = curX
		path[maxSteps-stepsLeft].y = curY

		if stepsLeft == 0 {
			if topMax < currMAx {
				topMax = currMAx
			}
			return
		}

		fu(curX, curY+1, stepsLeft-1, currMAx)
		fu(curX-1, curY, stepsLeft-1, currMAx)
		fu(curX+1, curY, stepsLeft-1, currMAx)
		fu(curX, curY-1, stepsLeft-1, currMAx)
	}
	fu(startX, startY, maxSteps, 0)
	return topMax
}

func main() {
	nutMap := [][]int{
		{0, 4, 0, 0, 0, 0, 0, 0, 6, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 0, 5},
		//{0, 0, 0, 0, 0, 9},

		//{0, 5, 0, 0, 0, 0},
		//{0, 0, 3, 0, 0, 0},
		//{0, 0, 0, 0, 0, 0},
		//{0, 0, 4, 0, 0, 0},
	}
	startPoses := []int{0, 4}
	maxSteps := 12 //12 expected
	//defer profile.Start().Stop()
	fmt.Printf("squirel max sum %d in %d steps.", buriedTreasure(nutMap, startPoses, maxSteps), maxSteps)
}
