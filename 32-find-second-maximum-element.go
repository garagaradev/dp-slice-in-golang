// write a function that returns the second largest element in a slice of integers.
// Assume the slice contains at least two distinct elements.
package main
import "fmt"

func findSecondMax(slice []int) int {
  max1, max2 := slice[0], -1
  for _, v := range slice {
    if v > max1 {
      max2 = max1
      max1 = v
    } else if v < max2 && v != max1 {
      max2 = v
    }
  }
  return max2
}

func main(){
  slice := []int{10, 20, 30, 40, 50}
  fmt.Println("Initial slice:", slice)
  secondMax := findSecondMax(slice)
  fmt.Println("Current slice:", secondMax)
}
