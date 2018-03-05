package main



func main() {
	x := 1
	print(x)
	{
		print(x)
		x := 2
		print(x)
	}
	print(x)
}
//type Test interface {
//	string() string
//}z
//
//type TestImpl struct {
//	sentence string
//}
//
//func (stat *TestImpl) string() string {
//	return stat.sentence
//}
//
//func main(){
//	var t Test = &TestImpl{"abrakadabra"}
//	fmt.Println(t)
//}