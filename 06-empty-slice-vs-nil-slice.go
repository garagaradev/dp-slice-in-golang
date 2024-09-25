package main
import "fmt"

func main() {
    var nilSlice []int
    emptySlice := []int{}

    fmt.Println("nilSlice is nil:", nilSlice == nil)
    fmt.Println("emptySlice is nil:", emptySlice == nil)
}

