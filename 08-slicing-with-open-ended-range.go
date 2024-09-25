// create a slice of integers [10,20,30,40,50]
// then extract the elements from index 2 to the end

package main
import "fmt"

func main(){
  slice := []int{10,20,30,40,50}
  fmt.Println("Slice:",slice)
  fmt.Println("Extract the elements from index 2 to the end")
  subSlice := slice[2:]
  fmt.Println("subSlice:",subSlice)
}
