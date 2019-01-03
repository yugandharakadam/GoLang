package main

import (
"fmt"
"log"
)

func main()  {
	var ans int
	fmt.Printf("Enter no")
	fmt.Scanln(&ans)
	aa,err := fmt.Scanln(&ans)

	if err != nil{
		log.Fatal("Err")
	}else{
		fmt.Printf("%d \n",ans)
		fmt.Printf("%d \n" ,aa)
	}
	
}