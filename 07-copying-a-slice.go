//create two integer slices.
//then copy the elements from the first slice to the second.
package main
import "fmt"

func main(){
  src := []int{1,2,3}
  dst := make([]int, len(src))
  copy(dst, src)
  fmt.Println("Source:",src)
  fmt.Println("Destination:",dst)
}
