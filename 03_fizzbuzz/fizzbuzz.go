package main
import "fmt"

func main(){
    for i := 1; i < 100; i++ {
        var result string
        if i % 3 == 0 {
            result += "fizz"
        }
        if i % 5 == 0 {
            result += "buzz"
        }
        if result != "" {
            fmt.Print(i)
            fmt.Printf(":" + result + "\n")
        }
    }

}
