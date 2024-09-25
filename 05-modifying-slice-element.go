//create a slice of integers [1,2,3,4] 
//then modify the element at index 2 to 10

package main
import "fmt"

func main(){
  slice := []int{1,2,3,4}
  fmt.Println("Initial slice:", slice)
  fmt.Println("Modify the element at index 2 to 10.")
  slice[2] = 10
  fmt.Println("Current Slice:", slice)

}
