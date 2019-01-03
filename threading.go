package main

import (
"fmt"
"strconv"
"time"
)

func main(){
	go simple("first")
	simple("second")
	//fmt.Scanln()
}

func simple(str string){
	for i := 1;i<10;i++ {
		fmt.Println(str + strconv.Itoa(i))
		time.Sleep(5000)
	}
}