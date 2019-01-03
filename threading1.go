package main

import (
"fmt"
"strconv"
"time"
"sync"
)

func main(){

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(){
		simple("first",10)
		wg.Done()
	}()
	wg.Wait()
	go func(){
		simple("second",20)
		wg.Done()
	}()
	wg.Wait()

}

func simple(str string,cnt int){
	for i := 1;i<10;i++ {
		fmt.Println(str + strconv.Itoa(i))
		time.Sleep(time.Millisecond * time.Duration(cnt))
	}
}