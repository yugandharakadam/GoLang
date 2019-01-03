package main

import (
	"fmt"
	http "net/http"
	"io"
	"log"
)

func main()  {
	
	hellohl := func(w http.ResponseWriter, req *http.Request){
		io.WriteString(w,"<h1>Index !!</h1>")
	}
	hellohl1 := func(w http.ResponseWriter, req *http.Request){
		io.WriteString(w,"<h1>Hello !!</h1>")
	}

	http.HandleFunc("/index",hellohl)
	http.HandleFunc("/",hellohl1)
	
	log.Fatal(http.ListenAndServe(":8080",nil))
	fmt.Println("Server started on 8080")
}