package main

import "fmt"
type entry struct{
	vert, horizontal, total int
}
func main() {
	field := [][]string{{"0", "0", "E", "0"},
		{"W", "0", "W", "E"},
		{"0", "E", "0", "W"},
		{"0", "W", "0", "E"}}
	fmt.Println(bomber(field))
}
func bomber(field [][]string) int {
	if len(field)==0 {return 0}
	lenX,lenY  := len(field), len(field[0])
	var mem = make(map[int]map[int]*entry)
	for x:=0;x<lenX;x++{
		mem[x] = make(map[int]*entry)
	}
	max := 0
	for x:=0;x<lenX;x++  {
		for y:=0;y<lenY ;y++  {
			if field[x][y]=="0"{
				element, okE := mem[x][y]
				if !okE {

					mem[x][y]= &entry{-1,-1,0}
					element = mem[x][y]

				}
				if (*element).horizontal==-1{
					expandHorizontal(x,y,field, &mem)
				}
				if (*element).vert ==-1{
					expandVert(x,y,field,&mem)
				}
				if (*element).total > max {
					max = (*element).total
				}
			}
		}
	}
	return max
}

func expandHorizontal(x int , y int, field [][]string, mem *map[int]map[int]*entry) {
	bombs := 0
	elementsToUpdate := make([]int,y)
	elementsToUpdate = append(elementsToUpdate, y)
	l, r := y-1 , y+1
	for l>=0 && field[x][l]!="W" {
		if field[x][l]=="E" {
			bombs++
		}
		if field[x][l]=="0"{
			elementsToUpdate = append(elementsToUpdate,l)
		}
		l--
	}
	for r<len(field[0]) && field[x][r]!="W" {
		if field[x][r]=="E" {
			bombs++
		}
		if field[x][r]=="0"{
			elementsToUpdate = append(elementsToUpdate,r)
		}
		r++
	}
	for _,e := range elementsToUpdate {
		ele,ok := (*mem)[x][e]
		if !ok {
			(*mem)[x][e]= &entry{-1,bombs, bombs}
		} else {
			ele.horizontal = bombs
			ele.total +=ele.horizontal
		}

	}
}

func expandVert(x int , y int, field [][]string,mem *map[int]map[int]*entry) {
	bombs := 0
	elementsToUpdate := make([]int,x)
	elementsToUpdate = append(elementsToUpdate, x)
	t, b := x-1 , x+1
	for t>=0 && field[t][y]!="W" {
		if field[t][y]=="E" {
			bombs++
		}
		if field[t][y]=="0"{
			elementsToUpdate = append(elementsToUpdate,t)
		}
		t--
	}
	for b<len(field) && field[b][y]!="W" {
		if field[b][y]=="E" {
			bombs++
		}
		if field[b][y]=="0"{
			elementsToUpdate = append(elementsToUpdate,b)
		}
		b++
	}
	for _,e := range elementsToUpdate {
		ele,ok := (*mem)[e][y]
		if !ok {
			(*mem)[e][y] = &entry{bombs,-1, bombs}
		} else {
			ele.vert = bombs
			ele.total +=bombs
		}

	}
}

