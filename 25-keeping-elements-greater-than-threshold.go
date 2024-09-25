// create a slice of integers [10, 20, 30, 40, 50]
// then keep only elements greater than 30
package main
import "fmt"

func retainGreaterThan(slice []int, threshold int) []int{
  result := []int{}
  for _, v := range slice {
    if v > threshold {
      result = append(result, v)
    }
  }
  return result
}

func main(){
  slice := []int{10, 20, 30, 40, 50}
  fmt.Println("Initial slice:", slice)
  n := 30
  fmt.Println("Keep only elements greater than", n)
  slice = retainGreaterThan(slice, n)
  fmt.Println("Current slice:", slice)
  
}
