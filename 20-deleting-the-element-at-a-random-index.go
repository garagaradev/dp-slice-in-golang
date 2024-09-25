// create a slice of integers [10, 20, 30, 40, 50]
// then delete an element at a random index
package main
import (
  "fmt"
  "math/rand"
  "time"
)

func main(){
  rand.Seed(time.Now().UnixNano())
  slice := []int{10, 20, 30, 40, 50}
  fmt.Println("Initial slice:", slice)
  randomIndex := rand.Intn(len(slice))
  fmt.Printf("Deleting element at index %d\n", randomIndex)
  slice = append(slice[:randomIndex], slice[randomIndex+1:]...)
  fmt.Println("Current slice:", slice)
}
