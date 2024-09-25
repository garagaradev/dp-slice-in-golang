// Create a slice of integers [1,2,3]
// then append the numbers 4 and 5 to the slice
package main
import "fmt"

func main(){
  slice := []int{1,2,3}
  fmt.Println("initial slice:",slice)
  slice = append(slice, 4, 5)
  fmt.Println("current slice:",slice)
}
