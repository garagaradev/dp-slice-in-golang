// create a slice of integers [10, 20, 30, 40, 50]
// then remove all elements at odd indices (1, 3)
package main
import "fmt"

func removeOddIndices(slice []int) []int {
  result := []int{}
  for i := range slice {
    if i % 2 == 0 {
      result = append(result, slice[i])
    }
  }
  return result
}

func main(){
  slice := []int{10, 20, 30, 40, 50}
  fmt.Println("Initial slice:", slice)
  slice = removeOddIndices(slice)
  fmt.Println("Current slice:", slice)
}
