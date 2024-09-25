// create a slice of integers [10, 20, 30, 40, 50]
// then remove all value less than 30.
package main
import "fmt"

func removeLessThan(slice []int, threshold int) []int {
  result := []int{}
  for _, v := range slice {
    if v >= threshold {
      result = append(result,v)
    }
  }
  return result
}

func main(){
  slice := []int{10, 20, 30, 40, 50}
  fmt.Println("Initial slice:", slice)
  slice = removeLessThan(slice, 30)
  fmt.Println("Current slice:", slice)
}
