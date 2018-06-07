package main

import ("fmt"
        "math/rand")

func add(x,y float64) float64{
    return x+y
}

func multiple(a,b string) (string, string) {
  return a,b
}


func main(){
  num1, num2 := 3.1415925, 2.8897878
  fmt.Println("A number from 1-100", rand.Intn(100))
  fmt.Println(add(num1,num2))
  fmt.Println(add(1,1))
  w1,w2 := "hey", "there"
  fmt.Println(multiple(w1,w2))
}


func pointers(){
  s := 15
  j := &s
  fmt.Println(s)
  fmt.Println(*j)
  

}
