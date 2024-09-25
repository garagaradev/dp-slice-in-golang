// create a slice of integers [10, 20, 30, 40, 50]
// then delete the first n elements where n=3
package main
import "fmt"

func main(){
  slice := []int{10, 20, 30, 40, 50}
  fmt.Println("Initial slice:", slice)
  fmt.Println("Deleting the first element where n=3...")
  n := 3
  slice = slice[n:]
  fmt.Println("Current slice:", slice)
}
