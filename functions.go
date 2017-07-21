package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func sub(x int, y int) int {
    return x-y
  }

func mult(x int, y int) int {
    return x*y
}

func bolme(x int, y int) int {
    return x/y
}

func kalan(x int, y int) int {
    return x%y
}


func main() {
  fmt.Println(add(42,42)) //84
  fmt.Println(sub(42,12)) //30
  fmt.Println(sub(12,42)) //-30
  fmt.Println(mult(5,100)) //500
  fmt.Println(bolme(100,5)) //20
  fmt.Println(bolme(5,100)) //0
  fmt.Println(kalan(102,5)) //2
}
