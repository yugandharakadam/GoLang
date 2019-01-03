package main
import ("fmt"
"os" )
func main1(){
    var s,str string
    for i := 1; i<len(os.Args); i++ {
        s += str + os.Args[i]
        str = " "
    }
    fmt.Printf("%s ",s)
    fmt.Printf("abc %d",123)

}