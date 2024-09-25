// create a slice of integers [10,20,30,40,50] 
// then delete the element at index 2 (value 30)

package main
import "fmt"

func main(){
  slice:= []int{10,20,30,40,50}
  fmt.Println("initial slice:",slice)
  slice = append(slice[:2],slice[3:]...)
  fmt.Println("current slice:",slice)
}
