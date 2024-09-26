// write a function that rotates the slice k positions to the right.
// for example: [1,2,3,4,5] and k = 2, result should be [4,5,1,2,3]
package main
import "fmt"

func rotateSlice(slice []int, k int) []int{
  n := len(slice)
  k = k % n // in case k is greater than the length of the slice
  return append(slice[n-k:],slice[:n-k]...)
}

func main(){
  slice := []int{1,2,3,4,5}
  fmt.Println("Initial slice:", slice)
  slice = rotateSlice(slice, 2)
  fmt.Println("Current slice:", slice)
}
