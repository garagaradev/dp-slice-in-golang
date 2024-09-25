// create a slice of integers [10, 20, 30, 40, 50, 60, 70, 80, 90]
// then remove elements located at prime indices
package main
import "fmt"

func isPrime(n int) bool {
  if n < 2 {
    return false
  }

  for i := 2; i*i <= n; i++ {
    if n%i == 0 {
      return false
    }
  }
  return true
}

func removePrimeIndices(slice []int) []int {
  result := []int{}
  for i, v := range slice {
    if !isPrime(i){
      result = append(result, v)
    }
  }
  return result
}

func main(){
  slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
  fmt.Println("Initial slice:", slice)
  fmt.Println("Remove elements at prime indices...")
  slice = removePrimeIndices(slice)
  fmt.Println("Current slice:", slice)
}

