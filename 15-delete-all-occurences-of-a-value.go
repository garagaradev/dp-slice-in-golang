//create a slice of integers [10,20,30,40,50]
//then remove all occurences of the value 20 (FILTER)
package main
import "fmt"

func removeValue(slice []int, value int) []int {
  result := []int{}
  for _, v := range slice {
    if v != value {
      result = append(result, v)
    }
  }
  return result
}

func main(){
  slice := []int {10, 20, 30, 20, 40, 50}
  fmt.Println("Initial slice:", slice)
  slice = removeValue(slice, 20)
  fmt.Println("Current slice:",slice)
}
