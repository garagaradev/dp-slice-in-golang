// create a slice of integers [10, 20, 30, 40, 50, 60, 70, 80, 90]
// then delete every third element from the slice
package main
import "fmt"

func removeEveryThird(slice []int) []int{
  result := []int{}
  for i, v := range slice {
    if(i+1)%3 != 0 {
      result = append(result, v)
    }
  }

  return result
}

func main(){
  slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
  fmt.Println("Initial slice:", slice)
  slice = removeEveryThird(slice)
  fmt.Println("Current slice:", slice)
}
