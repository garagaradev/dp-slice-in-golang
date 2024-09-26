//create a slice of integers [10, 20, 30, 40, 50]
//then replace the element at index 3 with 45.
package main
import "fmt"

func main(){
  slice := []int{10, 20, 30, 40, 50}
  fmt.Println("Initial slice:", slice)
  slice[3] = 45
  fmt.Println("Current slice:", slice)
  fmt.Println("Lol. This is too easy!")
}
