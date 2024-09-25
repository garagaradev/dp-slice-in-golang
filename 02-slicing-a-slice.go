//Create a slice of 5 integers [10,20,30,40,50] 
//then extract the elements from index 1 to 3

package main
import "fmt"

func main(){
  slice := []int{10,20,30,40,50}
  subSlice := slice[1:4] // index 1 is inclusive, meanwhile index 4 is exclusive
  fmt.Println(subSlice)
}
