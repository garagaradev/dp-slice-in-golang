// Given a slice of integers [10,20,20,30,30,40,50]
// remove the duplicates
package main
import "fmt"

func removeDuplicates(slice []int) []int{
    keys := make(map[int]bool)
    result := []int{}
    
    for _, entry := range slice {
      if _, exists := keys[entry]; !exists {
      keys[entry] = true
      result = append(result, entry)
    }
  }
  return result
}

func main(){
  slice := []int{10,20,20,30,30,40,50}
  fmt.Println("Initial slice:", slice)
  slice = removeDuplicates(slice)
  fmt.Println("Current slice:", slice)

}
