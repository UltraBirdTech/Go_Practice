package main
import "fmt"

func main(){
    var first int
    first = 0
    var second int
    second = 1
    var result int
    for i := 0; result < 1000; i++ {
        result = first + second
        fmt.Print(result)
        fmt.Printf("\n")
        first, second = second, result
    }
}
