// create a slice of integers [10,20,30,40,50]
// then delete the elements at index 1 and 2 (values 20 and 30)
package main
import "fmt"

func main(){
  slice := []int{10,20,30,40,50}
  fmt.Println("Initial slice:",slice)
  slice = append(slice[:1],slice[3:]...)
  fmt.Println("Current slice:", slice)
}
