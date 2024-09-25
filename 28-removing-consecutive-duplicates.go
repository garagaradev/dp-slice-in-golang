// given a slice of integers [10, 20, 20, 30, 30, 30, 40, 50]
// remove consecutive duplicates
package main
import "fmt"

func removeConsecutiveDuplicates(slice []int) []int {
  if len(slice) == 0 {
    return slice
  }

  result := []int{slice[0]}
  for i := 1; i < len(slice); i++ {
    if slice[i] != slice[i-1]{
      result = append(result, slice[i])
    }
  }
  return result
}

func main(){
  slice := []int{10, 20, 20, 30, 30, 30, 40, 50}
  fmt.Println("Initial slice:", slice)
  slice = removeConsecutiveDuplicates(slice)
  fmt.Println("Current slice:", slice)

}
