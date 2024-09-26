//create a slice of integers [10, 20, 30, 40, 50]
//then insert the value 25 at index 2
package main
import "fmt"

func insertAt(slice []int, index, value int) []int {
  slice = append(slice[:index + 1],slice[index:]...)
  slice[index] = value
  return slice
}

func main(){
  slice := []int{10, 20, 30, 40, 50}
  fmt.Println("Initial slice:", slice)
  slice = insertAt(slice, 2, 25)
  fmt.Println("Current slice:", slice)
}
