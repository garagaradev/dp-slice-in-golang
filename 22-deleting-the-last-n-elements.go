// create a slice of integers [10, 20, 30, 40, 50]
// then delete the last n elements where n = 2
package main
import "fmt"

func main(){
  slice := []int{10, 20, 30, 40, 50}
  fmt.Println("Initial slice:", slice)
  fmt.Println("Delete the last n elements where n=2...")
  n := 2
  slice = slice[:len(slice)-n]
  fmt.Println("Current slice:", slice)
}
