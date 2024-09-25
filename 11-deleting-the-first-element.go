// create a slice of integers [10, 20, 30, 40, 50]
// then delete the first element (value 10)
package main
import "fmt"

func main(){
  slice := []int{10,20,30,40,50}
  fmt.Println("Initial slice:", slice)
  slice = append(slice[1:])
  // alternative solution
  // slice = slice[1:]
  fmt.Println("Current slice:", slice)
}
