// create a slice of integers [10, 20, 30, 40, 50]
// then keep only the first and last elements
package main
import "fmt"

func main(){
  slice := []int{10, 20, 30, 40, 50}
  fmt.Println("Initial slice:", slice)
  slice = []int{slice[0], slice[len(slice)-1]}
  fmt.Println("Current slice:", slice)
}
