package main

import (
	"fmt"
	"github.com/quincy0/poker/poker"
)

func main()  {
	in := poker.InitMethod()
	p1 := [][]int{{1,2},{12,3},{13,3}}
	p2 := [][]int{{11,2},{13,2},{12,4}}
	res := in.Compare(p1,p2)
	fmt.Println(res)
}
