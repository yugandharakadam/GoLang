package main

import ("fmt"
"test"
)

func main(){
	i,j := 10,20
	k := test.Add(i,j)
	fmt.Printf("Addition %d",k)
}