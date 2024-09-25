//create a slice of integers [10,20,30,40,50]
//then delete the last element (value 50)
package main
import "fmt"

func main(){
  slice := []int{10,20,30,40,50}
  fmt.Println("Initial slice:", slice)
  slice = slice[:len(slice)-1]
  fmt.Println("Current slice:",slice)
}
