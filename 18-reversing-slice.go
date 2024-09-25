// create a slice of integers [10, 20, 30, 40, 50]
// then reverse the orders of the elements.
package main
import "fmt"

func reverseSlice(slice []int) []int {
  for i, j := 0, len(slice)-1; i < j; i,j = i+1, j-1 {
    slice[i], slice[j] = slice[j], slice[i]
  }
  return slice
}

func main(){
  slice:= []int{10, 20, 30, 40, 50}
  fmt.Println("Initial slice:", slice)
  slice = reverseSlice(slice)
  fmt.Println("Current slice:", slice)
}
