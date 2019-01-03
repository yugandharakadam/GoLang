package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main()  {
	url :="https://reqres.in/api/users/2"
	client := http.Client{}
	resp,err := client.Get(url)
	if err != nil {
		fmt.Println("Error %s",err);
	}else{
		fmt.Println(resp);
		data,_ := ioutil.ReadAll(resp.Body)
		fmt.Println("\n Body ",string(data))
	}
}