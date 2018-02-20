package main
import (
	"fmt"
)
type stt struct{
	a int
	b int
	s string
}
func main() {
	struct1 := stt{1 , 2 ,  "asda"}
	fmt.Println(struct1)
}