package main

import "fmt"

func main(){
	var f = 0
	var s = 1
	var r = 0
	fmt.Printf("%d ",f)
	fmt.Printf("%d ",s)
	for i := 1; i<=5 ; i++{
		
		r = f+s
		f = s
		s = r
		fmt.Printf("%d ",r)
	}
}