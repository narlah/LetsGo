package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(4))
}

func generateParenthesis(n int) []string {
	var a []string
	return recGenerateParenthesis(n, "", 0, 0, a)
}

func recGenerateParenthesis(n int, currentResult string, currentlyOpened int, currentlyClosed int, res []string) []string {
	if currentlyClosed == n && currentlyOpened == n {
		return append(res, currentResult)
	}
	if currentlyOpened < n {
		res = recGenerateParenthesis(n, currentResult+"(", currentlyOpened+1, currentlyClosed, res)
	}
	if currentlyClosed < n && currentlyOpened > currentlyClosed {
		res = recGenerateParenthesis(n, currentResult+")", currentlyOpened, currentlyClosed+1, res)
	}
	return res
}