// create a slice of integers [10, 15, 20, 25, 30, 35, 40]
// then keep only the even numbers
package main
import "fmt"

func retainEven(slice []int) []int {
  result := []int{}
  for _, v := range slice {
    if v%2 == 0 {
      result = append(result, v)
    }
  }
  return result
}

func main(){
  slice := []int{10,15,20,25,30,35,40}
  fmt.Println("Initial slice:", slice)
  slice = retainEven(slice)
  fmt.Println("Current slice:", slice)
}
