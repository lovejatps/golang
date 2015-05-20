package main

import (
    "fmt"
    "github.com/user/newmath"
)

func main(){
  a := []int{}
  fmt.Printf("%v\n", cap(a))
  var b = append(a, 1, 2)
  fmt.Printf("%v\n", cap(b))
  fmt.Printf("Hello, world. Sqrt(2) = %v\n", newmath.Sqrt(2))
}
