package main

import "fmt"

func swap(x, y string)(string, string){ // swap fonksiyonu, dizgi ("string") türünde iki sonuc donduruyor. 
    return x, y
}

func main(){
  k, j := swap("bugunde", "olmedik!")
  fmt.Println(k, j)
}
