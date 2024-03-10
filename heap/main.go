package main

import (
	"fmt"
	"sort"
)

func main() {
	infos := infos{
		{No: "1", Score: 100},
		{No: "2", Score: 90},
		{No: "10", Score: 10},
		{No: "4", Score: 70},
		{No: "8", Score: 30},
		{No: "5", Score: 60},
		{No: "6", Score: 50},
		{No: "7", Score: 40},
		{No: "3", Score: 80},
		{No: "9", Score: 20},
	}

	sort.Sort(infos)

	fmt.Println(infos)
}

type Info struct {
	No    string
	Score int
}

type infos []Info

func (in infos) Len() int {
	return len(in)
}

func (in infos) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}
func (in infos) Less(i, j int) bool {
	return in[i].Score > in[j].Score
}
