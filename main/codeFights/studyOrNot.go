package main

import (
	"fmt"
)

func main() {
	fmt.Println(studyOrSleep(35, 1))
}

func studyOrSleep(familiarity int, hoursRemaining int) int {
	//	res := 0  //got 23/24 , meh , lets just calculate it brutally
	//	newFam := float32(familiarity)
	//	if hoursRemaining > 8 {
	//		res = hoursRemaining - 8
	//		hoursRemaining -= res
	//		newFam = float32(float64(familiarity) * math.Pow(1.2 , float64(res)))
	//	}
	//	var unf = float32(100 - newFam)
	//	for i := hoursRemaining-1; i >= 0; i-- {
	//		coef := unf * 0.2
	//		if coef <= 5 {
	//			return res
	//		}
	//		res++
	//		unf -= coef
	//	}
	//	return res
	fam := float32(familiarity)
	unfam := float32(100 - fam)
	max := fam
	var pen float32 = 0
	res := 0
	for i := hoursRemaining-1; i >= 0; i-- {
		if i < 8 {
			pen++
		}
		coef := unfam * 0.2
		unfam -= coef
		fam += coef
		newMax := fam-pen*float32(5)
		if newMax > max {
			res = hoursRemaining - i
			max = newMax
		}

	}
	return res
}
