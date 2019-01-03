package main
import "fmt"

func main(){
	
	var arr[10] int
	arr[0]=0
	arr[1]=1
	for i := 2; i < len(arr);i++{
		arr[i] = arr[i-1] + arr[i-2]
	}
	for i := 0; i <len(arr);i++{
		fmt.Println(arr[i])
	}

}