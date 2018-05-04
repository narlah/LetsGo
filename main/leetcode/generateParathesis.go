package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(2))
}

func generateParenthesis(n int) []string {
	var a []string
	return recGenerateParenthesis(n, "", 0, 0, a)
}

func recGenerateParenthesis(n int, currentResult string, currentlyOpened int, currentlyClosed int, res []string) []string {
	if currentlyClosed == n && currentlyOpened == n {
		return append(res, currentResult)
	}
	var (
		k  []string
		k1 []string
	)
	if currentlyOpened < n {
		k = recGenerateParenthesis(n, currentResult+"(", currentlyOpened+1, currentlyClosed, res)
	}
	if currentlyClosed < n && currentlyOpened > currentlyClosed {
		k1 = recGenerateParenthesis(n, currentResult+")", currentlyOpened, currentlyClosed+1, res)
	}
	return append(k, k1...)
}
