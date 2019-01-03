package main

import (
"fmt"
"time"
)

func main()  {
	c := make(chan string, 1)
	go count("Value to be passed on channel",c)
	for i := 1;i<10;i++{
		fmt.Println("Listening : ",<-c)
	}
	fmt.Scanln()
}

func count (str string,c chan string){
	for i := 1;i<=5;i++{
		fmt.Println("count ",i)
		c <- str
		time.Sleep(time.Millisecond * 1500)
	}
}