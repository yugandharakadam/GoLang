package main

import (
	"fmt"
	simple "calc/simple"
	sc "calc/sc"
)

func main()  {
	fmt.Printf("Addition %d",simple.Add(10,20))
	fmt.Printf(" \n%d",simple.Sub(40,20))
	fmt.Printf("\n %d",simple.Divide(40,20))
	fmt.Printf("\n %d",simple.Mod(40,20))
	fmt.Printf(" \n%d",sc.Cos(40))
	fmt.Printf(" \n%d",sc.Sin(40))

}