//Create a slice of 3 integers [5,10,15]
//then print the length and capacity of the slice
package main
import "fmt"

func main(){
  slice := []int{5,10,15}
  fmt.Println("Slice:",slice)
  fmt.Println("Length:",len(slice))
  fmt.Println("Capacity:",cap(slice))
}
